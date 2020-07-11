package cmd

import (
	"fmt"
	"github.com/Bareksa/rest-api-boilerplate/models"
	"github.com/evalphobia/logrus_sentry"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_  "github.com/spf13/viper/remote"
	"os"
	"time"
)

var consulHost string

var rootCmd = &cobra.Command{
	// TODO: Don't forget to change Use field, and the descriptions.
	Use:   "default",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&consulHost, "consul", "", "app [cmd] --consul host:port")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "app [cmd] --consul host:port --verbose")
	log.SetFormatter(&log.JSONFormatter{})
	hook, err := logrus_sentry.NewAsyncSentryHook(viper.GetString("sentry.dsn"), []log.Level{
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
	})
	if err == nil {
		hook.Timeout = time.Duration(viper.GetInt("sentry.timeout")) * time.Second
		hook.StacktraceConfiguration.Enable = true
		log.AddHook(hook)
	}
}

func initConfig() {
	if consulHost == "" {
		fmt.Println("You must specify consul host. run `[cmd] --consul host:port`")
		os.Exit(1)
	}

	// Fist attempting to find remote config on Consul Key Value
	viper.AddRemoteProvider("consul", consulHost, "config-name")
	viper.SetConfigType("yaml") // Need to explicitly set this to yaml
	err := viper.ReadRemoteConfig()
	viper.AutomaticEnv()

	// If failed to get config on Consul Key Value, try to find a local config file
	if err != nil {
		fmt.Println(err)
		fmt.Println("attempting to use local config file....")
		viper.AddConfigPath(".")
		viper.SetConfigName(".app-config")
		viper.SetConfigType("yaml")
		viper.AutomaticEnv()
		err = viper.ReadInConfig()
		if err != nil {
			fmt.Println("failed to load local config", err)
			os.Exit(1)
		}else{
			fmt.Println("config local config loaded")
		}
	}else{

		err = viper.Unmarshal(&models.Config{})
		if err != nil {
			fmt.Printf("failed to unmarshal configuration from : %v\n", consulHost)
			os.Exit(1)
		}

		fmt.Println("remote config loaded")

		// Spinup a goroutine to watch remote changes forever
		go func() {
			for {
				time.Sleep(time.Second * 5) // delay after each request

				// currently, only tested with etcd support
				err := viper.WatchRemoteConfig()
				if err != nil {
					log.Errorf("unable to read remote config: %v", err)
					continue
				}

				// unmarshal new config into our runtime config struct. you can also use channel
				// to implement a signal to notify the system of the changes
				viper.Unmarshal(&models.Config{})
			}
		}()
	}
}


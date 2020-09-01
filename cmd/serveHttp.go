package cmd

import (
	"context"
	"fmt"
	"github.com/Bareksa/rest-api-boilerplate/infrastructures"
	"github.com/Bareksa/rest-api-boilerplate/routes"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var serveHttpCMD = &cobra.Command{
	Use:   "serveHttp",
	Short: "This command for serving http",
	Long:  "This command for serving http blah blah blah",
	Run: func(cmd *cobra.Command, args []string) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			oscall := <-c
			fmt.Printf("system call: %v", oscall)
			cancel()
		}()

		if err := serve(ctx); err != nil {
			fmt.Println(err)
		}
	},
}

func init()  {
	rootCmd.AddCommand(serveHttpCMD)
}

func serve(ctx context.Context) (err error){
	route := new(routes.Route)
	router := route.Init()
	consul := new(infrastructures.ServiceDiscovery)
	serviceID := uuid.New().String()
	consul.Register(serviceID, "boilerplate", hostname(), port())
	server := &http.Server{
		Addr:    port(),
		Handler: router,
	}

	go func() {
		fmt.Printf("server running on port %s\n", port())
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()

	fmt.Println("server stopped")
	ctxShutdown, cancel := context.WithTimeout(ctx, time.Second * 5)

	consul.DeRegister(serviceID)
	defer func() {
		cancel()
	}()

	if err = server.Shutdown(ctxShutdown); err != nil {
		fmt.Printf("failed to stop server: %v", err)
		os.Exit(1)
	}
	fmt.Println("server stop properly")
	if err == http.ErrServerClosed{
		err = nil
	}

	return
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	return ":" + port
}

func hostname() string{
	hostname, _ := os.Hostname()
	return hostname
}

package infrastructures

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"strconv"
)

type IServiceDiscovery interface {
	Register(serviceID, serviceName, serviceHostname, port string) error
	DeRegister(serviceID string)
}

type ServiceDiscovery struct {
	Consul *consulapi.Client
}

func createConnection() (*ServiceDiscovery, error){
	config := &consulapi.Config{
		Address: viper.GetString("consul.host"),
	}

	discovery, err := consulapi.NewClient(config)

	if err != nil {
		fmt.Printf("Failed create client discovery : %v", err)
	}

	return &ServiceDiscovery{
		Consul: discovery,
	}, nil
}

func (c *ServiceDiscovery) Register(serviceID, serviceName, serviceHostName, servicePort string) error  {
	consul, _ := createConnection()

	var register = new(consulapi.AgentServiceRegistration)
	register.ID = serviceID
	register.Name = serviceName
	register.Address = serviceHostName
	port, _ := strconv.Atoi(servicePort[1:len(servicePort)])
	register.Port = port

	register.Check = new(consulapi.AgentServiceCheck)
	register.Check.HTTP = fmt.Sprintf("http://%s:%v/v1/ping", serviceHostName, port)
	register.Check.Interval = "5s"
	register.Check.Timeout = "3s"
	return consul.Consul.Agent().ServiceRegister(register)
}

func (c *ServiceDiscovery) DeRegister(id string) error {
	consul, _ := createConnection()

	return consul.Consul.Agent().ServiceDeregister(id)
}

package sd

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
)

// ConsulRegister method.
func ConsulRegister(consulAddress string,
	consulPort string,
	advertiseAddress string,
	advertisePort string) (registrar sd.Registrar) {

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	// Service discovery domain. In this example we use consul.
	var client consulsd.Client
	{
		consulConfig := api.DefaultConfig()
		consulConfig.Address = consulAddress + ":" + consulPort
		theConsulClient, err := api.NewClient(consulConfig)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		client = consulsd.NewClient(theConsulClient)
	}

	check := api.AgentServiceCheck{
		HTTP:     "http://" + advertiseAddress + ":" + advertisePort + "/health",
		Interval: "10s",
		Timeout:  "1s",
		Notes:    "Basic health checks",
	}

	port, _ := strconv.Atoi(advertisePort)
	num := rand.Intn(100) // to make server ID unique
	asr := api.AgentServiceRegistration{
		ID:      "gizmo-serv-greeter-" + strconv.Itoa(num), // unique service ID
		Name:    "gizmo-serv-greeter",
		Tags:    []string{"gizmo", "greeter"},
		Port:    port,
		Address: advertiseAddress,
		Check:   &check,
	}

	registrar = consulsd.NewRegistrar(client, &asr, logger)
	return
}
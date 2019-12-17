package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"

	pb "bbxyard.com/msvc/pbs"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

const (
	svc_name = "go-micro-serv-greeter"
)

func main() {
	service := micro.NewService(
		micro.Name(svc_name+".client"),
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "Name to greeter",
			}),
	)

	var name string

	service.Init(
		micro.Action(func(c *cli.Context) {
			name = "GO-MICRO RPC CALL "
			if len(c.String("name")) > 0 {
				name = c.String("name")
			}
		}),
	)

	theClient := pb.NewGreeterService("go-micro-serv-greeter", client.DefaultClient)

	// client := pb.NewGreeterService(svc_name, service.Client())
	resp, err := theClient.Greeting(context.Background(), &pb.GreetingRequest{Name: name})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Greeting)
}

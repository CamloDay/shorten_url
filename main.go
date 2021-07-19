package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"shorten_url/api"
)

func main() {

	flag.String("ADDRESS", "localhost:8080", "address for the api")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		fmt.Printf("Error binding flags: %v\n", err)
		return
	}

	addr := viper.GetString("ADDRESS")

	fmt.Println("Starting Service")
	service := api.NewService()

	api.CreateRoutes(service, addr)

	fmt.Println("Starting Router")
	if err := service.Router.Run(addr); err != nil {
		fmt.Printf("Error starting router: %v\n", err)
	}

	fmt.Println("service stopped")
}

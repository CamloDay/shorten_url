package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

//go run cli/redirect/redirect_cli.go --SHORT='http://localhost:8080/03640c8a-727a-4f7f-804f-46f7006f3d49'
func main() {
	flag.String("SHORT", "", "short url")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		fmt.Println("Error parsing command line arguments: ", err)
		return
	}

	short := viper.GetString("SHORT")
	if short == "" {
		fmt.Println("SHORT flag required")
		return
	}

	req, err := http.NewRequest("GET", short, nil)
	if err != nil {
		fmt.Println("Error creating http request: ", err)
		return
	}

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending http request: ", err)
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return
	}

	fmt.Println(string(body))
}

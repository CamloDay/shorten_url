package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"shorten_url/api/handlers"
	"strings"
)

//go run cli/shorten/shorten_cli.go --ADDRESS='localhost:8080' --ORIGINAL='http://abc.com'
func main() {
	flag.String("ADDRESS", "localhost:8080", "address for the api")
	flag.String("ORIGINAL", "", "url to be shortened")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		fmt.Println("Error parsing command line arguments: ", err)
		return
	}

	original := viper.GetString("ORIGINAL")
	if original == "" {
		fmt.Println("ORIGINAL flag required")
		return
	}
	addr := viper.GetString("ADDRESS")

	payload := handlers.ShortRequest{
		Original: original,
	}
	p, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling payload: ", err)
		return
	}

	req, err := http.NewRequest("POST", "http://"+addr+"/shorten", strings.NewReader(string(p)))
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

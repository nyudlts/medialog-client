package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/nyudlts/medialog-client"
)

var (
	config      string
	environment string
	timeout     int
)

func init() {
	flag.StringVar(&config, "config", "", "")
	flag.StringVar(&environment, "environment", "", "")
	flag.IntVar(&timeout, "timeout", 20, "")
}

func main() {
	flag.Parse()

	mlc, err := medialog.NewClient(config, environment, 20)
	if err != nil {
		panic(err)
	}

	medialogHostInfo, err := mlc.GetHostInfo()
	if err != nil {
		panic(err)
	}

	hostJSON, err := json.MarshalIndent(medialogHostInfo, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", hostJSON)

	os.Exit(0)
}

package main

import (
	"flag"
	"log"
	"main/internal/app/apiserver"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	api := apiserver.New(config)

	if err := api.Start(); err != nil {
		log.Fatal(err)
	}
}
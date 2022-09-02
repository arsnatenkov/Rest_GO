package main

import (
	"Rest/internal/app/apiserver"
	"flag"
	"github.com/burntSushi/toml"
	"log"
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

	server := apiserver.NewAPIServer(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

//Config define struct for database conection
type Config struct {
	Server      string
	Database    string
	Connections int
}

func (config *Config) Read() {

	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatal(err)
	}
}

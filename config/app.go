package config

import (
	"fmt"
	"os"

	"github.com/tkanos/gonfig"
)

//Configuration app
type Configuration struct {
	DBUsername  string
	DBPassword  string
	DBHostXampp string
	DBName      string
	IsXampp     bool
}

var configuration Configuration

//Init config
func Init() {
	err := gonfig.GetConf("app.json", &configuration)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}
}

//GetConfig file configuration
func GetConfig() Configuration {
	return configuration
}

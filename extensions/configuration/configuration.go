package configuration

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Global *Configuration

type Configuration struct {
	ConnectionString string `yaml:"connectionString"`
	Database         string `yaml:"database"`
	Port             int    `yaml:"port"`
	AppName          string `yaml:"appname"`
}

func SetConfiguration() {
	env := "local" //os.Getenv("app_env")
	var t Configuration
	filePath := fmt.Sprintf("config.%s.yaml", env)
	value, _ := ioutil.ReadFile(filePath)
	err := yaml.Unmarshal(value, &t)
	if err != nil {
		log.Fatal(err)
		return
	}
	Global = &t
}

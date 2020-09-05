package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	EnvApp   = "APP_ENV"
	AppLocal = "local"
)

type HTTPServer struct {
	Port string `json:"port"`
}

type Configuration struct {
	HTTPServer HTTPServer
}

var Config = (func() Configuration {
	appEnv := os.Getenv(EnvApp)
	if appEnv == "" {
		appEnv = AppLocal
	}

	filePath := fmt.Sprintf("%s/config/config.%s.json", pwd(), appEnv)
	config := readJSONFile(filePath)
	return config
})()

func pwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("couldn't get working directory, possibly unsupported platform?")
	}
	// Replace forward slashes in case this is windows, URL parser errors
	return strings.Replace(cwd, "\\", "/", -1)
}

func readJSONFile(filePath string) Configuration {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Fail to read config file: ", err)
	}
	var config Configuration
	err = json.Unmarshal(fileBytes, &config)
	if err != nil {
		log.Fatal("Fail to unmarshall config file: ", err)
	}
	return config
}

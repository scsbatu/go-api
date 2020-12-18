package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// UserDataBase defines User DB config
type UserDataBase struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

// ApplicationConfig is the configuration
type ApplicationConfig struct {
	Service  string       `json:"service"`
	DataBase UserDataBase `json:"database"`
}

// Config stores the configuration
var Config *ApplicationConfig
var configFile *string

// LoadConfiguration loads configuration from file
// decodes the json config file into an instance of application config
// if the decoded config is valid it is set as config
func LoadConfiguration() error {
	if configFile == nil {
		return fmt.Errorf("config not initialized")
	}
	config := new(ApplicationConfig)
	file, err := os.Open(*configFile)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(file).Decode(config); err != nil {
		return err
	}
	Config = config
	return nil
}

// GetConfig returns a copy of init-ed application config instance
// if not already initialized it is initialized with "." as config path
func GetConfig() ApplicationConfig {
	if Config != nil {
		return *Config
	}
	err := Init(".")
	if err != nil {
		panic(err)
	}
	return *Config
}
func setupReload() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)
	go func() {
		for range c {
			log.Printf("Reloading configurations....\n")
			if configFile == nil {
				panic("Config file path not set!")
			}
			if err := LoadConfiguration(); err != nil {
				log.Printf("Error on reloading configurations, using old configuration : Error : %s\n", err.Error())
			}
		}
	}()
}

// Init will initialize app config with config file name
func Init(config string) error {
	if config == "" {
		config = "."
	}
	configFilepath := strings.TrimRight(config, "/") + "/server.config.json"
	configFile = &configFilepath
	setupReload()
	return LoadConfiguration()
}

package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

var Cnf Config

type Config struct {
	Server struct {
		Port string `yaml:"port" envconfig:"SERVER_PORT"`
		Host string `yaml:"host" envconfig:"SERVER_HOST"`
	} `yaml:"server"`
	Database struct {
		Port     string `yaml:"port" envconfig:"DB_PORT"`
		Host     string `yaml:"host" envconfig:"DB_HOST"`
		Username string `yaml:"user" envconfig:"DB_USERNAME"`
		Password string `yaml:"pass" envconfig:"DB_PASSWORD"`
		DbName   string `yaml:"dbname" envconfig:"DB_NAME"`
	} `yaml:"database"`
}

func readFileConfig(cfg *Config) {
	f, err := os.Open("./configs/config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		processError(err)
	}
}

func readEnvConfig(cfg *Config) {
	err := godotenv.Load()
	if err != nil {
		//Error loading .env file
		processError(err)
	}
	err = envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}

func InitConfig() {
	readFileConfig(&Cnf)
	readEnvConfig(&Cnf)
	// fmt.Printf("%+v", Cnf)
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

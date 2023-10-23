package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	HttpPort    string `yaml:"httpPort"`
	HttpsPort   string `yaml:"httpsPort"`
	Host        string `yaml:"host"`
	Env         string `yaml:"env"`
	CrtFile     string `yaml:"crtFile"`
	KeyFile     string `yaml:"keyFile"`
	MysqlUser   string `yaml:"mysqlUser"`
	MysqlPasswd string `yaml:"mysqlPasswd"`
	MysqlHost   string `yaml:"mysqlHost"`
	MysqlPort   string `yaml:"mysqlPort"`
	MysqlDbName string `yaml:"mysqlDbName"`
	Secret      string `yaml:"secret"`
}

var config *Config

func ParseConfig() *Config {
	config = &Config{}

	ofile, err := os.ReadFile("config.yaml")

	if err != nil {
		log.Fatalf("'%s'!", err)
	}

	err = yaml.Unmarshal(ofile, config)

	if err != nil {
		log.Printf("'%s'!", err)
	}

	return config
}

func GetConfig() *Config {
	return config
}
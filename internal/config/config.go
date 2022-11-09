package config

import (
	"caltoph/internal/logger"
	"os"

	"gopkg.in/yaml.v2"
)

var Config ServerConfig

type ServerConfig struct {
	Postgres_uri  string          `yaml:"postgres_uri"`
	Loglevel      string          `yaml:"loglevel"`
	Oidc_provider []Oidc_provider `yaml:"oidc_providers" json:"oidc_providers"`
}

type Oidc_provider struct {
	Name           string `yaml:"name" json:"name"`
	Url            string `yaml:"url" json:"url"`
	Url_authorize  string `yaml:"url_authorize" json:"url_authorize"`
	Url_token      string `yaml:"url_token" json:"-"`
	Url_userinfo   string `yaml:"url_userinfo" json:"-"`
	Client_id      string `yaml:"client_id" json:"client_id"`
	Client_secret  string `yaml:"client_secret" json:"-"`
	Username_claim string `yaml:"username_claim" json:"-"`
}

func Init(yamlFilePath string) ServerConfig {
	yamlConfig := ServerConfig{}
	//Read and parse config file if exists
	if yamlFilePath != "" {
		yamlFile, err := os.ReadFile(yamlFilePath)
		if err != nil {
			logger.FatalLogger.Println("Could not read config file")
			panic(err)
		}

		err2 := yaml.UnmarshalStrict(yamlFile, &yamlConfig)
		if err2 != nil {
			logger.FatalLogger.Println("Could not parse config file")
			panic(err2)
		}
	}

	env_postgres_uri, _ := os.LookupEnv("POSTGRES_URI")
	if env_postgres_uri != "" {
		yamlConfig.Postgres_uri = env_postgres_uri
	}

	env_loglevel, _ := os.LookupEnv("LOGLEVEL")
	if env_postgres_uri != "" {
		yamlConfig.Loglevel = env_loglevel
	}
	Config = yamlConfig
	return yamlConfig

}

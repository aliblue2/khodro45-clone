package configs

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Passwrod Passwordonfig
	Cors     CorsConfig
}

type ServerConfig struct {
	InternalPort string
	ExternalPort string
	RunMode      string
	Domain       string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode  string
}

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	Db                 int
	MinIdleConnections int
	PoolSize           int
	PoolTimeout        int
}

type Passwordonfig struct {
	IncludeChars     bool
	IncludeDigits    bool
	IncludeLowercase bool
	IncludeUppercase bool
	MinLength        int
	MaxLength        int
}

type CorsConfig struct {
	AllowedOrigins string
}

func GetConfig() *Config {
	configPath := getConfigPath(os.Getenv("APP_ENV"))
	viper, err := loadConfig(configPath, "yml")

	if err != nil {
		log.Fatalf("failed to load config")
	}

	config, err := parseConfig(viper)

	if err != nil {
		log.Fatalf("failed to parsed config")
		return nil
	}

	return config

}

func parseConfig(v *viper.Viper) (*Config, error) {
	var parsedCfg Config

	err := v.Unmarshal(&parsedCfg)

	if err != nil {
		log.Fatalf("unable to parse config : %v", err)
		return nil, err
	}

	return &parsedCfg, nil

}

func loadConfig(fileName string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(fileName)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not founded")
		}
		return nil, err
	}
	return v, nil
}

func getConfigPath(envPath string) string {
	switch envPath {
	case "docker":
		return "../configs/config-docker"
	case "production":
		return "../configs/config-production"
	default:
		return "../configs/config-development"

	}
}

package configs

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Password PasswordConfig
	Cors     CorsConfig
	Logger   LoggerConfig
}

type ServerConfig struct {
	InternalPort string
	ExternalPort string
	RunMode      string
	Domain       string
}

type PostgresConfig struct {
	Host               string
	Port               string
	User               string
	Password           string
	DbName             string
	SSLMode            string
	MaxIdleConns       int
	MaxOpenConns       int
	SetConnMaxLifetime time.Duration
}

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	Db                 string
	PoolSize           int
	PoolTimeout        time.Duration
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleCheckFrequency time.Duration
}

type PasswordConfig struct {
	IncludeChars     bool
	IncludeDigits    bool
	MinLength        int
	MaxLength        int
	IncludeUppercase bool
	IncludeLowercase bool
}

type CorsConfig struct {
	AllowOrigins string
}

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
	Logger   string
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

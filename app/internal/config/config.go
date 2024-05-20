package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Env      string `yaml:"env" env:"ENV" env-default:"local"`
	LogLevel string `yaml:"log_level" env-default:"debug"`
	Database
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"iddle_timeout" env-default:"60s"`
}

type Database struct {
	Username string `yaml:"username" env-default:"postgres"`
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	DBName   string `yaml:"dbname" env-default:"postgres"`
	SSLMode  string `yaml:"sslmode" env-default:"disable"`
	Password string
}

func New() *Config {
	return &Config{
		Database: Database{
			Username: viper.GetString("db.username"),
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
			Password: os.Getenv("DB_PASSWORD"),
		},
		HTTPServer: HTTPServer{
			Address:     viper.GetString("http_server.address"),
			Timeout:     viper.GetDuration("http_server.timeout"),
			IdleTimeout: viper.GetDuration("http_server.idle_timeout"),
		},
		Env:      viper.GetString("env"),
		LogLevel: viper.GetString("log_level"),
	}
}

func MustLoad() *Config {
	if err := initConfig(); err != nil {
		log.Fatalf("error loading configs: %s", err.Error())
	}

	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	cfg := New()

	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("error unmarshal configs: %s", err.Error())
	}

	return cfg
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("../config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Config{}); err != nil {
		return err
	}

	return nil
}

package config

import "github.com/spf13/viper"

// MongoDBConfig - config for mongodb database.
type MongoDBConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

// WebServer - config for web server.
type WebServer struct {
	Host string
	Port int
}

// Config - application configuration.
type Config struct {
	MongoDB MongoDBConfig
	Server  WebServer

	Env     string
	AppName string
}

func InitConfig() *Config {
	var cfg *Config = new(Config)

	// Init with default values.
	cfg.initDefault()

	// Init with file.
	cfg.initFileCfg()
	return cfg
}

func (c *Config) initDefault() {
	viper.SetDefault("mongodb.host", "localhost")
	viper.SetDefault("mongodb.port", 27017)
	viper.SetDefault("mongodb.database", "go-auth")
	viper.SetDefault("mongodb.username", "")
	viper.SetDefault("mongodb.password", "")

	viper.SetDefault("env", "dev")
	viper.SetDefault("appName", "go-auth")

	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", 8080)
}

func (c *Config) initFileCfg() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("config/")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
}

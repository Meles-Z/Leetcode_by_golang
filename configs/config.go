package config

import "github.com/spf13/viper"

type Configration struct {
	Auth     AuthenticationConfigration
	Server   ServerConfigration
	Database DatabaseConfigration
}
type DatabaseConfigration struct {
	Host     string
	Port     int
	Name     string
	Username string
	Password string
}

type ServerConfigration struct {
	Host string
	Port int
}

type AuthenticationConfigration struct {
	WebSecret string
	Secret    string
}

func LoadConfig() (*Configration, error) {
	var config Configration
	var err error
	viper.AddConfigPath("./")
	viper.SetConfigName(".")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.BindEnv("server.host", "SERVER_HOST"); err != nil {
		return nil, err
	}

	if err = viper.BindEnv("server.port", "SERVER_PORT"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.host", "DB_HOST"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.port", "DB_PORT"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.name", "DB_NAME"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.username", "DB_USERNAME"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.password", "DB_PASSWORD"); err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

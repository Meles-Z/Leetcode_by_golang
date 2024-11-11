package configs

import "github.com/spf13/viper"

type Config struct {
	Server ServerConfigration
	Db     DatabaseConfigration
	Auth   AuthConfigration
}

type ServerConfigration struct {
	Host string
	Port int
}

type AuthConfigration struct {
	Secret         string
	WebSecret      string
	MerchangSecret string
}

type DatabaseConfigration struct {
	Host     string
	Username string
	Password string
	Name     string
	Port     int
}

func LoadConfig() (*Config, error) {
	var config Config
	var err error

	viper.AddConfigPath("../")
	viper.SetConfigName(".")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err = viper.BindEnv("server.host", "DB_HOST"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("server.port", "DB_PORT"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.name", "DB_NAME"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.port", "DB_PORT"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.password", "DB_PASSWORD"); err != nil {
		return nil, err
	}

	if err = viper.BindEnv("auth.secret", "AUTH_SECRET"); err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

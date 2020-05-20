package main

import "fmt"

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (c PostgresConfig) Dialect() string {
	return "postgres"
}

func (c PostgresConfig) ConnectionInfo() string {
	if c.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s "+
			"sslmode=disable", c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s "+
		"sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Name)
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "123",
		Name:     "lenslocked_dev",
	}
}

type Config struct {
	Port    int
	Env     string
	Pepper  string `json:"pepper"`
	HMACKey string `json:"hmac_key"`
}

func (c Config) IsProd() bool {
	return c.Env == "prod"
}

func DefaultConfig() Config {
	return Config{
		Port:    3000,
		Env:     "dev",
		Pepper:  "!secret-random-string!",
		HMACKey: "secter-hmac-key",
	}
}

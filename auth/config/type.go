package config

import "time"

const (
	User = "user"
)

type (
	Config struct {
		Server  ServerConfig
		Service map[string]ServiceConfig
		Redis   RedisConfig
		Jwt     JwtConfig
	}

	ServerConfig struct {
		HostPort string
	}

	ServiceConfig struct {
		HostPort string
	}

	JwtConfig struct {
		SigningKey string
		ExpiresIn  time.Duration
	}

	RedisConfig struct {
		HostPort string
		Username string
		Password string
		DbNumber int
	}
)

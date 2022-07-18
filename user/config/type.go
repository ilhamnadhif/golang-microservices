package config

type (
	Config struct {
		Server   ServerConfig
		Database DatabaseConfig
		Redis    RedisConfig
		Jwt      JwtConfig
	}

	ServerConfig struct {
		HostPort string
	}

	JwtConfig struct {
		SigningKey string
	}

	DatabaseConfig struct {
		Driver   string
		HostPort string
		Username string
		Password string
		Database string
	}

	RedisConfig struct {
		HostPort string
		Username string
		Password string
		DbNumber int
	}
)

package main

const (
	EnvLocal      = "local"
	EnvProduction = "production"
)

type config struct {
	AllowedOrigins string `conf:"env:ALLOWED_ORIGINS"`
	Port           int    `conf:"env:PORT"`
	DatabaseURL    string `conf:"env:DATABASE_URL"`
}

var localConfig = config{
	AllowedOrigins: "http://localhost:3000",
	Port:           8080,
	DatabaseURL:    "postgresql://eventual-user:local-dev-password@localhost:5432/eventual",
}

var productionConfig = config{
	AllowedOrigins: "https://eventual-client.fly.dev",
	Port:           8080,
}

func Config(env string) config {
	switch env {
	case EnvLocal:
		return localConfig
	case EnvProduction:
		return productionConfig
	default:
		return config{}
	}
}

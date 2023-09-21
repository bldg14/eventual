package main

type config struct {
	AllowedOrigins string `conf:"env:ALLOWED_ORIGINS,default:http://localhost:3000"`
	Port           int    `conf:"env:PORT,default:8080"`
}

var productionConfig = config{
	AllowedOrigins: "https://eventual-client.fly.dev",
	Port:           8080,
}

func Config(env string) config {
	switch env {
	case "local":
		return config{}
	case "production":
		return productionConfig
	default:
		return config{}
	}
}

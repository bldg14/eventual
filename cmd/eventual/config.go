package main

type config struct {
	AllowedOrigins string `conf:"default:http://localhost:3000"`
	Port           int    `conf:"env:PORT,default:8080"`
}

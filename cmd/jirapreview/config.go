package jirapreview

import (
	"os"
)

type Config struct {
	JiraAPIUrl      string
	JiraLogin       string
	JiraToken       string
	PachcaAPIUnfurl string
	PachcaUserId    string
	PachcaToken     string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func NewConfig() *Config {
	return &Config{
	}
}

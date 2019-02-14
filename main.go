package main

import (
	"log"
	"os"
)

func GetEnv(key, fallback string) string {
	// Return the correct enviroment variable if found
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	// If the environment variable is not found utilise the default and throw a warning
	log.Println("Warning. Enviroment variable '%s' was not found. Falling back to the default value '%s'", key, fallback)
	return fallback
}

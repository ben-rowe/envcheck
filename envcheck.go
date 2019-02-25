package envcheck

import (
	"log"
	"os"
)

func GetEnv(key, fallback string, fatal bool) string {
	// Return the correct enviroment variable if found
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	
	// Do we want to log or kill the program?
	if fatal {
		// If the environment variable is not found exit the program.
		log.Fatal("Fatal. Enviroment variable '%s' was not found. Exiting.", key)
	}
	
	// If the environment variable is not found utilise the default and throw a warning
	log.Printf("Warning. Enviroment variable '%s' was not found. Falling back to the default value '%s'", key, fallback)
	
	return fallback
}

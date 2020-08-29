package utils

import "os"

// APIURL of cosmic cms
var APIURL = "https://api.cosmicjs.com/v1/"

// GetPortEnv function to get port from env
func GetPortEnv() string {
	var port string
	var ok bool

	if port, ok = os.LookupEnv("PORT"); !ok {
		port = "8080"
	}

	return ":" + port
}

// CheckIfEnvExists if env with key exists
func CheckIfEnvExists(key string) bool {
	var ok bool
	if _, ok = os.LookupEnv(key); !ok {
		return false
	}
	return true
}

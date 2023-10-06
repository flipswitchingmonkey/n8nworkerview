package main

import "os"

func getEnv(key string, fallback string) (envValue string) {
	value, exists := os.LookupEnv(key)
	if (exists) {
		envValue = value
	} else {
		envValue = fallback
	}
	return
}

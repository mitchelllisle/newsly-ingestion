package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func GetEnvOrString(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func GetEnvOrPanic(key string) string {
	value, exists := os.LookupEnv(key)
	if exists == false {
		err := errors.New(fmt.Sprintf("Key does not exist: %s", key))
		PanicOnError(err, "Unable to get environment variable")
	}
	return value
}

func PanicOnError(err error, message string) {
	if err != nil {
		log.Panicf("%s: %s", message, err)
	}
}

func ContinueOnError(err error, message string) {
	if err != nil {
		log.Printf("%s: %s", message, err)
	}
}

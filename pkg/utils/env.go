package utils

import (
	"log"
	"os"
	"strconv"
)

func GetEnvVariable(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Panicln("ENV mising, key: " + key)
	}
	return value
}

func GetEnvVariableAsBool(key string) bool {
	value := os.Getenv(k)
	if value == "" {
		log.Panicln("ENV missing, key: " + key)
	}

	boolean, err := strconv.ParseBool(value)
	if err != nil {
		log.Panicln("ENV err: [" + key + "]\n" + err.Error())
	}

	return boolean
}

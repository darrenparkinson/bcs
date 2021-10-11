package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func MustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		log.Fatalf("environment variable %q not set", envKey)
	}
	*target = v
}

func SaveAsJSON(filename string, data interface{}) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, file, 0644)
}

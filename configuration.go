package gopaci

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	BaseUrl  string
	Username string
	Password string
}

func ReadConfig() *Configuration {
	file, err := os.Open("conf.json")
	if err != nil {
		return nil
	}

	decoder := json.NewDecoder(file)
	configuration := &Configuration{}
	decoder.Decode(&configuration)

	return configuration
}

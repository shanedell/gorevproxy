package pkg

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func readJSON(fileData []byte) error {
	if err := json.Unmarshal(fileData, &config); err != nil {
		return err
	}

	return nil
}

func readYAML(fileData []byte) error {
	if err := yaml.Unmarshal(fileData, &config); err != nil {
		return err
	}

	return nil
}

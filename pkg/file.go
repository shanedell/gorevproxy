package pkg

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func ReadJSON(fileData []byte) error {
	if err := json.Unmarshal(fileData, &Config); err != nil {
		return err
	}

	return nil
}

func ReadYAML(fileData []byte) error {
	if err := yaml.Unmarshal(fileData, &Config); err != nil {
		return err
	}

	return nil
}

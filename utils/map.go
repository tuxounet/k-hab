package utils

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func LoadYamlFromString[R any](yamlStr string) (R, error) {

	var anyStruct R

	err := yaml.Unmarshal([]byte(yamlStr), &anyStruct)
	if err != nil {
		return anyStruct, err
	}

	return anyStruct, nil

}

func LoadJSONFromString[R any](jsonStr string) (R, error) {
	var anyStruct R

	err := (json.Unmarshal([]byte(jsonStr), &anyStruct))

	if err != nil {
		return anyStruct, err
	}

	return anyStruct, nil

}

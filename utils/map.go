package utils

import (
	"encoding/json"
	"strings"

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

func GetMapValue(anyMap any, path string) any {

	keys := strings.Split(path, ".")
	var value interface{}
	value = anyMap
	for _, key := range keys {
		value = value.(map[string]interface{})[key]
	}
	return value

}

package utils

import (
	"encoding/json"
	"strings"

	"gopkg.in/yaml.v3"
)

func LoadYamlFromString[R any](ctx *ScopeContext, yamlStr string) R {
	return ScopingWithReturnOnly(ctx, "utils", "LoadYamlFromString", func(ctx *ScopeContext) R {
		var anyJson R

		ctx.Must(yaml.Unmarshal([]byte(yamlStr), &anyJson))

		return anyJson
	})
}

func LoadJSONFromString[R any](ctx *ScopeContext, yamlStr string) R {
	return ScopingWithReturnOnly(ctx, "utils", "LoadJSONFromString", func(ctx *ScopeContext) R {
		var anyJson R

		ctx.Must(json.Unmarshal([]byte(yamlStr), &anyJson))

		return anyJson
	})
}

func GetMapValue(ctx *ScopeContext, m map[string]interface{}, path string) any {
	return ScopingWithReturnOnly(ctx, "utils", "GetMapValue", func(ctx *ScopeContext) any {
		ctx.Log.DebugF("GetMapValue:  %v", path)
		keys := strings.Split(path, ".")
		var value interface{}
		value = m
		for _, key := range keys {
			value = value.(map[string]interface{})[key]
		}
		return value
	})
}

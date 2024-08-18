package utils

import (
	"encoding/json"
	"strings"

	"gopkg.in/yaml.v3"
)

func LoadYamlFromString[R any](ctx *ScopeContext, yamlStr string) R {
	return ScopingWithReturn(ctx, "utils", "LoadYamlFromString", func(ctx *ScopeContext) R {
		var anyStruct R

		ctx.Must(yaml.Unmarshal([]byte(yamlStr), &anyStruct))

		return anyStruct
	})
}

func LoadJSONFromString[R any](ctx *ScopeContext, jsonStr string) R {
	return ScopingWithReturn(ctx, "utils", "LoadJSONFromString", func(ctx *ScopeContext) R {
		var anyStruct R

		ctx.Must(json.Unmarshal([]byte(jsonStr), &anyStruct))

		return anyStruct
	})
}

func GetMapValue(ctx *ScopeContext, m map[string]interface{}, path string) any {
	return ScopingWithReturn(ctx, "utils", "GetMapValue", func(ctx *ScopeContext) any {
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

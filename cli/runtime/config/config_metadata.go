package config

import (
	"gopkg.in/yaml.v3"
)

func ShouldReplace(node *yaml.Node, key string) (bool, error) {

	clientConfig, err := convertNodeToClientConfig(node)
	if err != nil {
		return false, err
	}

	if val, ok := clientConfig.ConfigMetadata.PatchStrategy[key]; ok {
		return val == "replace", nil
	}

	return false, nil
}

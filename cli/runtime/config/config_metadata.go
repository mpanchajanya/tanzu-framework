package config

import (
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
	"gopkg.in/yaml.v3"
)

func ShouldReplace(node *yaml.Node, key string) (bool, error) {

	clientConfig, err := convertFromNode[configapi.ClientConfig](node)
	if err != nil {
		return false, err
	}

	if val, ok := clientConfig.ConfigMetadata.PatchStrategy[key]; ok {
		return val == "replace", nil
	}

	return false, nil
}

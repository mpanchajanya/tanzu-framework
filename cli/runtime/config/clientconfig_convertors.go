package config

import (
	"github.com/pkg/errors"
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"

	"gopkg.in/yaml.v3"
)

func convertNodeToClientConfig(node *yaml.Node) (cfg *configapi.ClientConfig, err error) {
	err = node.Decode(&cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode nodeutils to client config")
	}
	return cfg, err
}

func convertNodeToMap(node *yaml.Node) (envs map[string]string, err error) {
	err = node.Decode(&envs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode nodeutils to client config")
	}
	return envs, err
}

func convertMapToNode(envs map[string]string) (*yaml.Node, error) {
	bytes, err := yaml.Marshal(envs)
	if err != nil {
		return nil, err
	}
	var node yaml.Node
	err = yaml.Unmarshal(bytes, &node)
	if err != nil {
		return nil, err
	}

	return &node, nil
}

func convertClientConfigToNode(cfg *configapi.ClientConfig) (*yaml.Node, error) {
	bytes, err := yaml.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	var node yaml.Node
	err = yaml.Unmarshal(bytes, &node)
	if err != nil {
		return nil, err
	}

	return &node, nil
}

func convertDiscoverySourceToNode(ds *configapi.PluginDiscovery) (*yaml.Node, error) {
	bytes, err := yaml.Marshal(ds)
	if err != nil {
		return nil, err
	}
	var node yaml.Node
	err = yaml.Unmarshal(bytes, &node)
	if err != nil {
		return nil, err
	}

	return &node, nil
}

func convertRepositoryToNode(repository *configapi.PluginRepository) (*yaml.Node, error) {
	bytes, err := yaml.Marshal(repository)
	if err != nil {
		return nil, err
	}
	var node yaml.Node
	err = yaml.Unmarshal(bytes, &node)
	if err != nil {
		return nil, err
	}

	return &node, nil
}

func convertContextToNode(c *configapi.Context) (*yaml.Node, error) {
	bytes, err := yaml.Marshal(c)
	if err != nil {
		return nil, err
	}
	var node yaml.Node
	err = yaml.Unmarshal(bytes, &node)
	if err != nil {
		return nil, err
	}

	return &node, nil

}

func convertServerToNode(s *configapi.Server) (*yaml.Node, error) {
	bytes, err := yaml.Marshal(s)
	if err != nil {
		return nil, err
	}

	var node yaml.Node
	err = yaml.Unmarshal(bytes, &node)
	if err != nil {
		return nil, err
	}

	return &node, nil
}

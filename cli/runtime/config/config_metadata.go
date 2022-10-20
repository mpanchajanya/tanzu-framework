package config

import (
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
	"github.com/vmware-tanzu/tanzu-framework/cli/runtime/config/nodeutils"
	"gopkg.in/yaml.v3"
)

func ShouldReplace(parent string) (bool, error) {

	return false, nil
}

//func getConfigMetadataPatchStrategiesByFeature(feature string) ()

func GetConfigMetadata() (*configapi.ConfigMetadata, error) {
	node, err := GetClientConfigNode()
	if err != nil {
		return nil, err
	}
	return getConfigMetadata(node)
}

func GetConfigMetadataPatchStrategy() (map[string]string, error) {
	configMetadata, err := GetConfigMetadata()
	if err != nil {
		return nil, err
	}

	return configMetadata.PatchStrategy, nil
}

func getConfigMetadata(node *yaml.Node) (*configapi.ConfigMetadata, error) {
	cfg, err := nodeutils.ConvertFromNode[configapi.ClientConfig](node)
	if err != nil {
		return nil, err
	}
	return cfg.ConfigMetadata, nil
}

func SetConfigMetadataPatchStrategy(key, value string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	err = setConfigMetadataPatchStrategy(node, key, value)
	if err != nil {
		return err
	}

	return PersistNode(node)

}

func SetConfigMetadataPatchStrategies(patchStrategies map[string]string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	err = setConfigMetadataPatchStrategies(node, patchStrategies)
	if err != nil {
		return err
	}

	return PersistNode(node)

}

func setConfigMetadataPatchStrategies(node *yaml.Node, patchStrategies map[string]string) error {
	for key, value := range patchStrategies {
		err := setConfigMetadataPatchStrategy(node, key, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func setConfigMetadataPatchStrategy(node *yaml.Node, key string, value string) error {
	//TODO: Validations for key and value
	configOptions := func(c *nodeutils.Config) {
		c.ForceCreate = true
		c.Keys = []nodeutils.Key{
			{Name: KeyConfigMetadata, Type: yaml.MappingNode},
			{Name: KeyPatchStrategy, Type: yaml.MappingNode},
		}
	}

	patchStrategyNode, err := nodeutils.FindNode(node.Content[0], configOptions)
	if err != nil {
		return err
	}

	if index := nodeutils.GetNodeIndex(patchStrategyNode.Content, key); index != -1 {
		patchStrategyNode.Content[index].Tag = "!!str"
		patchStrategyNode.Content[index].Value = value
	} else {
		patchStrategyNode.Content = append(patchStrategyNode.Content, nodeutils.CreateScalarNode(key, value)...)
	}
	return nil

}

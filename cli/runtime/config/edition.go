package config

import (
	"fmt"

	"github.com/vmware-tanzu/tanzu-framework/cli/runtime/config/nodeutils"
	"gopkg.in/yaml.v3"
)

// GetEdition returns the edition from the local configuration file
func GetEdition() (string, error) {
	node, err := GetClientConfigNode()
	if err != nil {
		return "", err
	}
	val, err := getEdition(node)
	if err != nil {
		return "", err
	}
	return val, nil
}

func SetEdition(val string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}
	err = setEdition(node, val)
	if err != nil {
		return err
	}
	return PersistNode(node)
}

func setEdition(node *yaml.Node, val string) error {
	configOptions := func(c *nodeutils.Config) {
		c.ForceCreate = true
		c.Keys = []nodeutils.Key{
			{Name: KeyClientOptions, Type: yaml.MappingNode},
			{Name: KeyCLI, Type: yaml.MappingNode},
			{Name: KeyEdition, Type: yaml.ScalarNode, Value: val},
		}
	}
	fmt.Println("hello")
	fmt.Println(node.Content[0])
	editionNode, err := nodeutils.FindNode(node.Content[0], configOptions)
	if err != nil {
		return err
	}
	editionNode.Value = val
	return nil
}

func getEdition(node *yaml.Node) (string, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return "", err
	}

	if cfg != nil && cfg.ClientOptions != nil && cfg.ClientOptions.CLI != nil {
		return string(cfg.ClientOptions.CLI.Edition), nil
	}
	return "", nil
}

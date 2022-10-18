package config

import (
	"github.com/pkg/errors"
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
	nodeutils "github.com/vmware-tanzu/tanzu-framework/cli/runtime/config/nodeutils"
	"gopkg.in/yaml.v3"
)

func GetUnstableVersionSelector() (configapi.VersionSelectorLevel, error) {
	node, err := GetClientConfigNode()
	if err != nil {
		return "", err
	}
	return getUnstableVersionSelector(node)
}

func getUnstableVersionSelector(node *yaml.Node) (configapi.VersionSelectorLevel, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return "", err
	}

	if cfg.ClientOptions != nil && cfg.ClientOptions.CLI != nil && cfg.ClientOptions.CLI.UnstableVersionSelector != "" {
		return cfg.ClientOptions.CLI.UnstableVersionSelector, nil
	}
	return "", errors.New("unstable version selector not found")
}
func SetUnstableVersionSelector(name string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	err = setUnstableVersionSelector(node, name)
	if err != nil {
		return err
	}

	return PersistNode(node)

}

func setUnstableVersionSelector(node *yaml.Node, name string) error {

	configOptions := func(c *nodeutils.Config) {
		c.ForceCreate = true
		c.Keys = []nodeutils.Key{
			{Name: KeyClientOptions, Type: yaml.MappingNode},
			{Name: KeyCLI, Type: yaml.MappingNode},
			{Name: KeyUnstableVersionSelector, Type: yaml.ScalarNode, Value: name},
		}
	}

	unstableVersionSelectorNode, err := nodeutils.FindNode(node.Content[0], configOptions)
	if err != nil {
		return err
	}
	unstableVersionSelectorNode.Value = name

	return nil

}

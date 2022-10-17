package config

import (
	"github.com/pkg/errors"
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
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

func setUnstableVersionSelector(node *yaml.Node, name string) error {
	clientOptionsNode := FindParentNode(node, KeyClientOptions)
	if clientOptionsNode == nil {
		//create cliClientOptions node and add to root node
		node.Content[0].Content = append(node.Content[0].Content, CreateMappingNode(KeyClientOptions)...)
		clientOptionsNode = FindParentNode(node, KeyClientOptions)
	}

	cliNode := FindNode(clientOptionsNode, KeyCLI)
	if cliNode == nil {
		//create cli node and add to root node
		clientOptionsNode.Content = append(clientOptionsNode.Content, CreateMappingNode(KeyCLI)...)
		cliNode = FindNode(clientOptionsNode, KeyCLI)
	}

	unstableVersionSelectorNode := FindNode(cliNode, KeyUnstableVersionSelector)
	if unstableVersionSelectorNode == nil {
		cliNode.Content = append(cliNode.Content, CreateScalarNode(KeyUnstableVersionSelector, name)...)
		unstableVersionSelectorNode = FindNode(cliNode, KeyUnstableVersionSelector)
	} else {
		unstableVersionSelectorNode.Content[0].Value = name
	}

	return nil

}

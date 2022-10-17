package config

import (
	"github.com/pkg/errors"
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
	"gopkg.in/yaml.v3"
)

func GetCLIDiscoverySources() ([]configapi.PluginDiscovery, error) {
	node, err := GetClientConfigNode()
	if err != nil {
		return nil, err
	}

	return getCLIDiscoverySources(node)
}

func getCLIDiscoverySources(node *yaml.Node) ([]configapi.PluginDiscovery, error) {

	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}

	if cfg.ClientOptions != nil && cfg.ClientOptions.CLI != nil && cfg.ClientOptions.CLI.DiscoverySources != nil {
		return cfg.ClientOptions.CLI.DiscoverySources, nil
	}

	return nil, errors.New("cli discovery sources not found")

}

func GetCLIDiscoverySource(name string) (*configapi.PluginDiscovery, error) {
	node, err := GetClientConfigNode()
	if err != nil {
		return nil, err
	}

	return getCLIDiscoverySource(node, name)
}

func getCLIDiscoverySource(node *yaml.Node, name string) (*configapi.PluginDiscovery, error) {

	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}

	if cfg.ClientOptions != nil && cfg.ClientOptions.CLI != nil && cfg.ClientOptions.CLI.DiscoverySources != nil {

		for _, discoverySource := range cfg.ClientOptions.CLI.DiscoverySources {
			_, discoverySourceName := getDiscoverySourceTypeAndName(discoverySource)

			if discoverySourceName == name {
				return &discoverySource, nil
			}

		}
	}

	return nil, errors.New("cli discovery source not found")

}

func SetCLIDiscoverySource(discoverySource configapi.PluginDiscovery) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	err = setCLIDiscoverySource(node, discoverySource)
	if err != nil {
		return err
	}
	return PersistNode(node)

}

func setCLIDiscoverySource(node *yaml.Node, discoverySource configapi.PluginDiscovery) error {

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

	err := setDiscoverySource(cliNode, discoverySource)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCLIDiscoverySource(name string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	err = deleteCLIDiscoverySource(node, name)
	if err != nil {
		return err
	}

	return PersistNode(node)

}

func deleteCLIDiscoverySource(node *yaml.Node, name string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	cliNode := FindParentSubNode(node, KeyClientOptions, KeyCLI)
	if cliNode == nil {
		return nil
	}

	cliDiscoverySourcesNode := FindNode(cliNode, KeyDiscoverySources)
	if cliDiscoverySourcesNode == nil {
		return nil
	}

	discoverySource, err := getCLIDiscoverySource(node, name)
	if err != nil {
		return nil
	}

	var result []*yaml.Node
	for _, discoverySourceNode := range cliDiscoverySourcesNode.Content {
		discoverySourceType, discoverySourceName := getDiscoverySourceTypeAndName(*discoverySource)
		if discoverySourceIndex := getNodeIndex(discoverySourceNode.Content, discoverySourceType); discoverySourceIndex != -1 {
			if discoverySourceFieldIndex := getNodeIndex(discoverySourceNode.Content[discoverySourceIndex].Content, "name"); discoverySourceFieldIndex != -1 && discoverySourceNode.Content[discoverySourceIndex].Content[discoverySourceFieldIndex].Value == discoverySourceName {
				continue
			}
		} else {
			result = append(result, discoverySourceNode)
		}
	}

	if len(result) == 0 {
		cliDiscoverySourcesNode.Kind = yaml.ScalarNode
		cliDiscoverySourcesNode.Tag = "!!seq"
	} else {
		cliDiscoverySourcesNode.Content = result
	}

	return nil

}

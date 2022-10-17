package config

import (
	"github.com/pkg/errors"
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
	"gopkg.in/yaml.v3"
)

func GetCLIRepositories() ([]configapi.PluginRepository, error) {
	node, err := GetClientConfigNode()
	if err != nil {
		return nil, err
	}

	return getCLIRepositories(node)
}

func getCLIRepositories(node *yaml.Node) ([]configapi.PluginRepository, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}

	if cfg.ClientOptions != nil && cfg.ClientOptions.CLI != nil && cfg.ClientOptions.CLI.Repositories != nil {
		return cfg.ClientOptions.CLI.Repositories, nil
	}

	return nil, errors.New("cli repositories not found")

}

func GetCLIRepository(name string) (*configapi.PluginRepository, error) {
	node, err := GetClientConfigNode()
	if err != nil {
		return nil, err
	}

	return getCLIRepository(node, name)
}

func getCLIRepository(node *yaml.Node, name string) (*configapi.PluginRepository, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}

	if cfg.ClientOptions != nil && cfg.ClientOptions.CLI != nil && cfg.ClientOptions.CLI.Repositories != nil {
		for _, repository := range cfg.ClientOptions.CLI.Repositories {
			_, repositoryName := getRepositoryTypeAndName(repository)
			if repositoryName == name {
				return &repository, nil
			}
		}
	}
	return nil, errors.New("cli repository not found")
}

func setCLIRepository(node *yaml.Node, repository configapi.PluginRepository) error {
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

	err := setRepository(cliNode, repository)
	if err != nil {
		return err
	}

	return PersistNode(node)

}

func DeleteCLIRepository(name string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	err = deleteCLIRepository(node, name)
	if err != nil {
		return err
	}

	return PersistNode(node)

}

func deleteCLIRepository(node *yaml.Node, name string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	cliNode := FindParentSubNode(node, KeyClientOptions, KeyCLI)
	if cliNode == nil {
		return nil
	}

	cliRepositoriesNode := FindNode(cliNode, KeyRepositories)
	if cliRepositoriesNode == nil {
		return nil
	}

	repository, err := getCLIRepository(node, name)
	if err != nil {
		return nil
	}

	var result []*yaml.Node
	for _, repositoryNode := range cliRepositoriesNode.Content {

		repositoryType, repositoryName := getRepositoryTypeAndName(*repository)

		if repositoryIndex := getNodeIndex(repositoryNode.Content, repositoryType); repositoryIndex != -1 {
			if repositoryFieldIndex := getNodeIndex(repositoryNode.Content[repositoryIndex].Content, "name"); repositoryFieldIndex != -1 && repositoryNode.Content[repositoryIndex].Content[repositoryFieldIndex].Value == repositoryName {
				continue
			}
		} else {
			result = append(result, repositoryNode)
		}

	}

	if len(result) == 0 {
		cliRepositoriesNode.Kind = yaml.ScalarNode
		cliRepositoriesNode.Tag = "!!seq"
	} else {
		cliRepositoriesNode.Content = result
	}

	return nil

}

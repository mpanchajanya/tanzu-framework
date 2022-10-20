package config

import (
	"github.com/pkg/errors"
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
	nodeutils "github.com/vmware-tanzu/tanzu-framework/cli/runtime/config/nodeutils"
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
	cfg, err := nodeutils.ConvertFromNode[configapi.ClientConfig](node)
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
	cfg, err := nodeutils.ConvertFromNode[configapi.ClientConfig](node)
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

func SetCLIRepository(repository configapi.PluginRepository) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	err = setCLIRepository(node, repository)
	if err != nil {
		return err
	}

	return PersistNode(node)
}

func setCLIRepository(node *yaml.Node, repository configapi.PluginRepository) error {

	configOptions := func(c *nodeutils.Config) {
		c.ForceCreate = true
		c.Keys = []nodeutils.Key{
			{Name: KeyClientOptions, Type: yaml.MappingNode},
			{Name: KeyCLI, Type: yaml.MappingNode},
			{Name: KeyRepositories, Type: yaml.SequenceNode},
		}
	}

	repositoriesNode, err := nodeutils.FindNode(node.Content[0], configOptions)
	if err != nil {
		return err
	}

	return setRepository(repositoriesNode, repository)

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

	configOptions := func(c *nodeutils.Config) {
		c.ForceCreate = false
		c.Keys = []nodeutils.Key{
			{Name: KeyClientOptions},
			{Name: KeyCLI},
			{Name: KeyRepositories},
		}
	}

	cliRepositoriesNode, err := nodeutils.FindNode(node.Content[0], configOptions)
	if err != nil {
		return err
	}

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

		if repositoryIndex := nodeutils.GetNodeIndex(repositoryNode.Content, repositoryType); repositoryIndex != -1 {
			if repositoryFieldIndex := nodeutils.GetNodeIndex(repositoryNode.Content[repositoryIndex].Content, "name"); repositoryFieldIndex != -1 && repositoryNode.Content[repositoryIndex].Content[repositoryFieldIndex].Value == repositoryName {
				continue
			}
		} else {
			result = append(result, repositoryNode)
		}

	}

	cliRepositoriesNode.Style = 0
	cliRepositoriesNode.Content = result

	return nil

}

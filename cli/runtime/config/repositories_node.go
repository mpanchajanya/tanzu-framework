package config

import (
	"github.com/pkg/errors"
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
	nodeutils "github.com/vmware-tanzu/tanzu-framework/cli/runtime/config/nodeutils"
	"gopkg.in/yaml.v3"
)

func setRepository(node *yaml.Node, repository configapi.PluginRepository) error {
	newNode, err := convertToNode[configapi.PluginRepository](&repository)
	if err != nil {
		return err
	}

	configOptions := func(c *nodeutils.Config) {
		c.ForceCreate = true
		c.Keys = []nodeutils.Key{
			{Name: KeyRepositories, Type: yaml.SequenceNode},
		}
	}

	repositoriesNode, err := nodeutils.FindNode(node.Content[0], configOptions)
	if err != nil {
		return err
	}

	exists := false
	var result []*yaml.Node
	for _, repositoryNode := range repositoriesNode.Content {

		repositoryType, repositoryName := getRepositoryTypeAndName(repository)

		if repositoryType == "" || repositoryName == "" {
			return errors.New("not found")
		}

		if repositoryIndex := nodeutils.GetNodeIndex(repositoryNode.Content, repositoryType); repositoryIndex != -1 {
			if repositoryFieldIndex := nodeutils.GetNodeIndex(repositoryNode.Content[repositoryIndex].Content, "name"); repositoryFieldIndex != -1 && repositoryNode.Content[repositoryIndex].Content[repositoryFieldIndex].Value == repositoryName {
				exists = true

				err = nodeutils.MergeNodes(newNode.Content[0], repositoryNode, nil)
				if err != nil {
					return err
				}
				result = append(result, repositoryNode)
			}
		}

	}

	if !exists {
		result = append(result, newNode.Content[0])
	}

	repositoriesNode.Content = result

	return nil

}

func getRepositoryTypeAndName(repository configapi.PluginRepository) (string, string) {

	if repository.GCPPluginRepository != nil || repository.GCPPluginRepository.Name != "" {
		return "gcpPluginRepository", repository.GCPPluginRepository.Name
	}
	return "", ""

}

package config

import (
	"github.com/pkg/errors"
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
	"gopkg.in/yaml.v3"
)

func setRepository(node *yaml.Node, repository configapi.PluginRepository) error {
	newNode, err := convertRepositoryToNode(&repository)
	if err != nil {
		return err
	}

	i := getNodeIndex(node.Content, KeyRepositories)
	if i == -1 {
		node.Content = append(node.Content, CreateSequenceNode(KeyRepositories)...)
		i = getNodeIndex(node.Content, KeyRepositories)
	}
	repositoriesNode := node.Content[i]

	exists := false
	var result []*yaml.Node
	for _, repositoryNode := range repositoriesNode.Content {

		repositoryType, repositoryName := getRepositoryTypeAndName(repository)

		if repositoryType == "" || repositoryName == "" {
			return errors.New("not found")
		}

		if repositoryIndex := getNodeIndex(repositoryNode.Content, repositoryType); repositoryIndex != -1 {
			if repositoryFieldIndex := getNodeIndex(repositoryNode.Content[repositoryIndex].Content, "name"); repositoryFieldIndex != -1 && repositoryNode.Content[repositoryIndex].Content[repositoryFieldIndex].Value == repositoryName {
				exists = true

				err = MergeNodes(newNode.Content[0], repositoryNode, nil)
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

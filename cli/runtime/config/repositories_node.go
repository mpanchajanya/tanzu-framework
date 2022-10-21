package config

import (
	"github.com/pkg/errors"
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
	nodeutils "github.com/vmware-tanzu/tanzu-framework/cli/runtime/config/nodeutils"
	"gopkg.in/yaml.v3"
)

func setRepository(repositoriesNode *yaml.Node, repository configapi.PluginRepository) (persist bool, err error) {
	newNode, err := nodeutils.ConvertToNode[configapi.PluginRepository](&repository)
	if err != nil {
		return persist, err
	}

	exists := false
	var result []*yaml.Node
	for _, repositoryNode := range repositoriesNode.Content {

		repositoryType, repositoryName := getRepositoryTypeAndName(repository)

		if repositoryType == "" || repositoryName == "" {
			return persist, errors.New("not found")
		}

		if repositoryIndex := nodeutils.GetNodeIndex(repositoryNode.Content, repositoryType); repositoryIndex != -1 {
			if repositoryFieldIndex := nodeutils.GetNodeIndex(repositoryNode.Content[repositoryIndex].Content, "name"); repositoryFieldIndex != -1 && repositoryNode.Content[repositoryIndex].Content[repositoryFieldIndex].Value == repositoryName {
				exists = true
				persist, err = nodeutils.NotEqual(newNode.Content[0], repositoryNode)
				if persist {
					err = nodeutils.MergeNodes(newNode.Content[0], repositoryNode)
					if err != nil {
						return persist, err
					}
				}

				result = append(result, repositoryNode)
			}
		}

	}

	if !exists {
		result = append(result, newNode.Content[0])
		persist = true
	}

	repositoriesNode.Content = result

	return persist, err

}

func getRepositoryTypeAndName(repository configapi.PluginRepository) (string, string) {

	if repository.GCPPluginRepository != nil && repository.GCPPluginRepository.Name != "" {
		return "gcpPluginRepository", repository.GCPPluginRepository.Name
	}
	return "", ""

}

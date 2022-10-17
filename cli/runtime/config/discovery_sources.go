package config

import (
	"github.com/pkg/errors"
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
	"gopkg.in/yaml.v3"
)

// DiscoveryType constants
const (
	DiscoveryTypeOCI        = "oci"
	DiscoveryTypeLocal      = "local"
	DiscoveryTypeGCP        = "gcp"
	DiscoveryTypeKubernetes = "kubernetes"
	DiscoveryTypeREST       = "rest"
)

func setDiscoverySource(node *yaml.Node, discoverySource configapi.PluginDiscovery) error {
	newNode, err := convertDiscoverySourceToNode(&discoverySource)
	if err != nil {
		return err
	}

	i := getNodeIndex(node.Content, KeyDiscoverySources)
	if i == -1 {
		node.Content = append(node.Content, CreateSequenceNode(KeyDiscoverySources)...)
		i = getNodeIndex(node.Content, KeyDiscoverySources)
	}
	discoverySourcesNode := node.Content[i]

	exists := false
	var result []*yaml.Node
	for _, discoverySourceNode := range discoverySourcesNode.Content {

		discoverySourceType, discoverySourceName := getDiscoverySourceTypeAndName(discoverySource)

		if discoverySourceType == "" || discoverySourceName == "" {
			return errors.New("not found")
		}

		if discoverySourceIndex := getNodeIndex(discoverySourceNode.Content, discoverySourceType); discoverySourceIndex != -1 {
			if discoverySourceFieldIndex := getNodeIndex(discoverySourceNode.Content[discoverySourceIndex].Content, "name"); discoverySourceFieldIndex != -1 && discoverySourceNode.Content[discoverySourceIndex].Content[discoverySourceFieldIndex].Value == discoverySourceName {
				exists = true

				err = MergeNodes(newNode.Content[0], discoverySourceNode, nil)
				if err != nil {
					return err
				}
				result = append(result, discoverySourceNode)
			}
		}

	}

	if !exists {
		result = append(result, newNode.Content[0])
	}

	discoverySourcesNode.Content = result

	return nil

}

func getDiscoverySourceTypeAndName(discoverySource configapi.PluginDiscovery) (string, string) {

	if discoverySource.GCP != nil || discoverySource.GCP.Name != "" {
		return DiscoveryTypeGCP, discoverySource.GCP.Name
	} else if discoverySource.OCI != nil || discoverySource.OCI.Name != "" {
		return DiscoveryTypeOCI, discoverySource.OCI.Name
	} else if discoverySource.Local != nil || discoverySource.Local.Name != "" {
		return DiscoveryTypeLocal, discoverySource.Local.Name
	} else if discoverySource.Kubernetes != nil || discoverySource.Kubernetes.Name != "" {
		return DiscoveryTypeKubernetes, discoverySource.Kubernetes.Name
	} else if discoverySource.REST != nil || discoverySource.REST.Name != "" {
		return DiscoveryTypeREST, discoverySource.REST.Name
	}

	return "", ""

}

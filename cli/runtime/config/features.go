package config

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"gopkg.in/yaml.v3"
)

func IsFeatureEnabled(plugin, key string) (bool, error) {
	node, err := GetClientConfigNode()
	if err != nil {
		return false, err
	}

	val, err := getFeature(node, plugin, key)
	if err != nil {
		return false, err
	}

	if strings.EqualFold(val, "true") {
		return true, nil
	}

	return false, nil
}

func getFeature(node *yaml.Node, plugin, key string) (string, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return "", err
	}

	if cfg.ClientOptions == nil || cfg.ClientOptions.Features == nil || cfg.ClientOptions.Features[plugin] == nil {
		return "", errors.New("not found")
	}

	if val, ok := cfg.ClientOptions.Features[plugin][key]; ok {
		return val, nil
	}

	return "", errors.New("not found")
}

func DeleteFeature(plugin, key string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	return deleteFeature(node, plugin, key)
}

func deleteFeature(node *yaml.Node, plugin, key string) error {

	featuresNode := FindParentSubNode(node, KeyClientOptions, KeyFeatures)
	if featuresNode == nil {
		return nil
	}

	pluginNode := FindNode(featuresNode, plugin)
	if pluginNode == nil {
		return nil
	}

	var currentPluginFeatures []*yaml.Node
	for _, pluginFeatureNode := range pluginNode.Content {
		if index := getNodeIndex(pluginFeatureNode.Content, key); index != -1 {
			continue
		}
		currentPluginFeatures = append(currentPluginFeatures, pluginFeatureNode)
	}

	if len(currentPluginFeatures) != 0 {
		pluginNode.Content = currentPluginFeatures
	}

	return nil
}

func SetFeature(plugin, key, value string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	return setFeature(node, plugin, key, value)
}

func setFeature(node *yaml.Node, plugin, key, value string) error {

	pluginNode := findPluginNode(plugin, node)

	if index := getNodeIndex(pluginNode.Content, key); index != -1 {
		pluginNode.Content[index].Value = value
	} else {
		pluginNode.Content = append(pluginNode.Content, CreateScalarNode(key, value)...)
	}
	return nil
}

func ConfigureDefaultFeatureFlagsIfMissing(plugin string, defaultFeatureFlags map[string]bool) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	pluginNode := findPluginNode(plugin, node)

	for key, value := range defaultFeatureFlags {
		val := strconv.FormatBool(value)
		if index := getNodeIndex(pluginNode.Content, key); index != -1 {
			pluginNode.Content[index].Value = val
		} else {
			pluginNode.Content = append(pluginNode.Content, CreateScalarNode(key, val)...)
		}

	}

	return nil

}

func findPluginNode(plugin string, node *yaml.Node) *yaml.Node {
	clientOptionsNode := FindParentNode(node, KeyClientOptions)
	if clientOptionsNode == nil {
		//create cliClientOptions node and add to root node
		node.Content[0].Content = append(node.Content[0].Content, CreateMappingNode(KeyClientOptions)...)
		clientOptionsNode = FindParentNode(node, KeyClientOptions)
	}

	featuresNode := FindNode(clientOptionsNode, KeyFeatures)
	if featuresNode == nil {
		//create features node and add to clientOptionsNode node
		clientOptionsNode.Content = append(clientOptionsNode.Content, CreateMappingNode(KeyFeatures)...)
		featuresNode = FindNode(clientOptionsNode, KeyFeatures)
	}

	pluginNode := FindNode(featuresNode, plugin)
	if pluginNode == nil {
		//create features node and add to clientOptionsNode node
		featuresNode.Content = append(featuresNode.Content, CreateMappingNode(plugin)...)
		pluginNode = FindNode(featuresNode, plugin)
	}
	return pluginNode
}

// IsFeatureActivated returns true if the given feature is activated
// User can set this CLI feature flag using `tanzu config set features.global.<feature> true`
func IsFeatureActivated(feature string) bool {
	cfg, err := GetClientConfig()
	if err != nil {
		return false
	}
	status, err := cfg.IsConfigFeatureActivated(feature)
	if err != nil {
		return false
	}
	return status
}

package config

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
	nodeutils "github.com/vmware-tanzu/tanzu-framework/cli/runtime/config/nodeutils"

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
	cfg, err := convertFromNode[configapi.ClientConfig](node)
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

	err = deleteFeature(node, plugin, key)
	if err != nil {
		return err
	}

	return PersistNode(node)

}

func deleteFeature(node *yaml.Node, plugin, key string) error {
	configOptions := func(c *nodeutils.Config) {

		c.Keys = []nodeutils.Key{
			{Name: KeyClientOptions},
			{Name: KeyFeatures},
			{Name: plugin},
		}
	}

	pluginNode, err := nodeutils.FindNode(node.Content[0], configOptions)
	if err != nil {
		return err
	}

	if pluginNode == nil {
		return nil
	}

	plugins, err := convertNodeToMap(pluginNode)
	if err != nil {
		return err
	}

	if _, ok := plugins[key]; ok {
		delete(plugins, key)
	}

	newPluginsNode, err := convertMapToNode(plugins)
	if err != nil {
		return err
	}

	pluginNode.Content = newPluginsNode.Content[0].Content

	return nil
}

func SetFeature(plugin, key, value string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	err = setFeature(node, plugin, key, value)
	if err != nil {
		return err
	}

	return PersistNode(node)
}

func setFeature(node *yaml.Node, plugin, key, value string) error {

	configOptions := func(c *nodeutils.Config) {
		c.ForceCreate = true
		c.Keys = []nodeutils.Key{
			{Name: KeyClientOptions, Type: yaml.MappingNode},
			{Name: KeyFeatures, Type: yaml.MappingNode},
			{Name: plugin, Type: yaml.MappingNode},
		}
	}

	pluginNode, err := nodeutils.FindNode(node.Content[0], configOptions)
	if err != nil {
		return err
	}

	if index := nodeutils.GetNodeIndex(pluginNode.Content, key); index != -1 {
		pluginNode.Content[index].Tag = "!!str"
		pluginNode.Content[index].Value = value
	} else {
		pluginNode.Content = append(pluginNode.Content, nodeutils.CreateScalarNode(key, value)...)
	}
	return nil
}

func ConfigureDefaultFeatureFlagsIfMissing(plugin string, defaultFeatureFlags map[string]bool) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	configOptions := func(c *nodeutils.Config) {
		c.ForceCreate = true
		c.Keys = []nodeutils.Key{
			{Name: KeyClientOptions, Type: yaml.MappingNode},
			{Name: KeyFeatures, Type: yaml.MappingNode},
			{Name: plugin, Type: yaml.MappingNode},
		}
	}

	pluginNode, err := nodeutils.FindNode(node.Content[0], configOptions)
	if err != nil {
		return err
	}

	for key, value := range defaultFeatureFlags {
		val := strconv.FormatBool(value)
		if index := nodeutils.GetNodeIndex(pluginNode.Content, key); index != -1 {
			pluginNode.Content[index].Value = val
		} else {
			pluginNode.Content = append(pluginNode.Content, nodeutils.CreateScalarNode(key, val)...)
		}
	}
	return nil
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

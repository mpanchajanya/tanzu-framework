package config

import (
	"github.com/pkg/errors"
	nodeutils "github.com/vmware-tanzu/tanzu-framework/cli/runtime/config/nodeutils"

	"gopkg.in/yaml.v3"
)

func GetAllEnvs() (map[string]string, error) {
	node, err := GetClientConfigNode()
	if err != nil {
		return nil, err
	}

	return getAllEnvs(node)

}

func getAllEnvs(node *yaml.Node) (map[string]string, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}

	if cfg.ClientOptions != nil && cfg.ClientOptions.Env != nil {
		return cfg.ClientOptions.Env, nil
	}

	return nil, errors.New("not found")

}

func GetEnv(key string) (string, error) {
	node, err := GetClientConfigNode()
	if err != nil {
		return "", err
	}

	return getEnv(node, key)
}

func getEnv(node *yaml.Node, key string) (string, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return "", err
	}

	if cfg.ClientOptions == nil && cfg.ClientOptions.Env == nil {
		return "", errors.New("not found")
	}

	if val, ok := cfg.ClientOptions.Env[key]; ok {
		return val, nil
	}

	return "", errors.New("not found")
}

func DeleteEnv(key string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}
	_, err = deleteEnv(node, key)
	if err != nil {
		return err
	}

	return PersistNode(node)
}

func deleteEnv(node *yaml.Node, key string) (ok bool, err error) {

	configOptions := func(c *nodeutils.Config) {
		c.Keys = []nodeutils.Key{
			{Name: KeyClientOptions},
			{Name: KeyEnv},
		}
	}

	envsNode, err := nodeutils.FindNode(node.Content[0], configOptions)
	if err != nil {
		return false, err
	}

	if envsNode == nil {
		return true, nil
	}

	envs, err := convertNodeToMap(envsNode)
	if err != nil {
		return false, err
	}

	if _, ok := envs[key]; ok {
		delete(envs, key)
	}

	newEnvsNode, err := convertMapToNode(envs)
	if err != nil {
		return false, err
	}

	envsNode.Content = newEnvsNode.Content[0].Content

	return true, nil
}

func SetEnv(key, value string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	err = setEnv(node, key, value)
	if err != nil {
		return err
	}

	return PersistNode(node)
}

func setEnv(node *yaml.Node, key, value string) error {

	configOptions := func(c *nodeutils.Config) {
		c.ForceCreate = true
		c.Keys = []nodeutils.Key{
			{Name: KeyClientOptions, Type: yaml.MappingNode},
			{Name: KeyEnv, Type: yaml.MappingNode},
		}
	}

	envsNode, err := nodeutils.FindNode(node.Content[0], configOptions)
	if err != nil {
		return err
	}

	envs, err := convertNodeToMap(envsNode)
	if err != nil {
		return err
	}

	envs[key] = value

	newEnvsNode, err := convertMapToNode(envs)
	if err != nil {
		return err
	}

	envsNode.Content = newEnvsNode.Content[0].Content

	return nil
}

// GetEnvConfigurations returns a map of configured environment variables
// to values as part of tanzu configuration file
// it returns nil if configuration is not yet defined
func GetEnvConfigurations() map[string]string {
	envs, err := GetAllEnvs()

	if err != nil {
		return make(map[string]string)
	}

	return envs
}

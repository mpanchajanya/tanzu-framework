package config

import (
	"github.com/pkg/errors"

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

	return nil
}

func deleteEnv(node *yaml.Node, key string) (ok bool, err error) {

	envsNode := FindParentSubNode(node, KeyClientOptions, KeyEnv)
	if envsNode == nil {
		return true, nil
	}

	var currentEnvs []*yaml.Node
	for _, envNode := range envsNode.Content {
		if index := getNodeIndex(envNode.Content, key); index != -1 {
			continue
		}
		currentEnvs = append(currentEnvs, envNode)
	}

	if len(currentEnvs) != 0 {
		envsNode.Content = currentEnvs
	}

	return true, nil
}

func SetEnv(key, value string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	return setEnv(node, key, value)
}

func setEnv(node *yaml.Node, key, value string) error {

	clientOptionsNode := FindParentNode(node, KeyClientOptions)
	if clientOptionsNode == nil {
		//create cliClientOptions node and add to root node
		node.Content[0].Content = append(node.Content[0].Content, CreateMappingNode(KeyClientOptions)...)
		clientOptionsNode = FindParentNode(node, KeyClientOptions)
	}

	envNode := FindNode(clientOptionsNode, KeyEnv)
	if envNode == nil {
		//create env node and add to clientOptionsNode node
		clientOptionsNode.Content = append(clientOptionsNode.Content, CreateMappingNode(KeyEnv)...)
		envNode = FindNode(clientOptionsNode, KeyEnv)
	}

	if index := getNodeIndex(envNode.Content, key); index != -1 {
		envNode.Content[index].Value = value
	} else {
		envNode.Content = append(envNode.Content, CreateScalarNode(key, value)...)
	}
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

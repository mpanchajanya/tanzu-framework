package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// GetClientConfigNode retrieves the config from the local directory with file lock
func GetClientConfigNode() (*yaml.Node, error) {
	// Acquire tanzu config lock
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	return GetClientConfigNodeNoLock()
}

// GetClientConfigNodeNoLock retrieves the config from the local directory without acquiring the lock
func GetClientConfigNodeNoLock() (*yaml.Node, error) {
	cfgPath, err := ClientConfigPath()
	if err != nil {
		return nil, errors.Wrap(err, "GetClientConfigNodeNoLock: failed getting client config path")
	}

	bytes, err := ioutil.ReadFile(cfgPath)
	if err != nil || len(bytes) == 0 {
		fmt.Errorf("failed to read in config: %v\n", err)
		node, err := NewClientConfigNode()
		if err != nil {
			return nil, errors.Wrap(err, "GetClientConfigNodeNoLock: failed to create new client config")
		}
		return node, nil
	}
	var node yaml.Node

	err = yaml.Unmarshal(bytes, &node)
	if err != nil {
		return nil, errors.Wrap(err, "GetClientConfigNodeNoLock: failed to construct struct from config data")
	}

	return &node, nil
}

func NewClientConfigNode() (*yaml.Node, error) {
	c := newClientConfig()
	node, err := convertClientConfigToNode(c)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func PersistNode(node *yaml.Node) error {
	cfgPath, err := ClientConfigPath()
	if err != nil {
		return errors.Wrap(err, "could not find config path")
	}

	cfgPathExists, err := fileExists(cfgPath)
	if err != nil {
		return errors.Wrap(err, "failed to check config path existence")
	}
	if !cfgPathExists {
		localDir, err := LocalDir()
		if err != nil {
			return errors.Wrap(err, "could not find local tanzu dir for OS")
		}
		if err := os.MkdirAll(localDir, 0755); err != nil {
			return errors.Wrap(err, "could not make local tanzu directory")
		}
	}

	data, err := yaml.Marshal(node)
	if err != nil {
		return errors.Wrap(err, "failed to marshal node")
	}

	err = os.WriteFile(cfgPath, data, 0644)
	if err != nil {
		return errors.Wrap(err, "failed to write the config to file.")
	}

	storeConfigToLegacyDir(data)

	return nil
}

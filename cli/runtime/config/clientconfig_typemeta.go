package config

import (
	"github.com/vmware-tanzu/tanzu-framework/cli/runtime/config/nodeutils"
	"gopkg.in/yaml.v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func setTypeMeta(node *yaml.Node, typeMeta metav1.TypeMeta) error {

	err := setKind(node, typeMeta.Kind)
	if err != nil {
		return err
	}

	err = setApiVersion(node, typeMeta.APIVersion)
	if err != nil {
		return err
	}
	return nil
}

func setKind(node *yaml.Node, kind string) error {
	configOptions := func(c *nodeutils.Config) {
		c.ForceCreate = true
		c.Keys = []nodeutils.Key{
			{Name: KeyKind, Type: yaml.ScalarNode, Value: kind},
		}
	}
	kindNode, err := nodeutils.FindNode(node.Content[0], configOptions)
	if err != nil {
		return err
	}
	kindNode.Value = kind
	return nil

}

func setApiVersion(node *yaml.Node, apiVersion string) error {
	configOptions := func(c *nodeutils.Config) {
		c.ForceCreate = true
		c.Keys = []nodeutils.Key{
			{Name: KeyApiVersion, Type: yaml.ScalarNode, Value: apiVersion},
		}
	}
	ApiVersionNode, err := nodeutils.FindNode(node.Content[0], configOptions)
	if err != nil {
		return err
	}

	ApiVersionNode.Value = apiVersion

	return nil

}

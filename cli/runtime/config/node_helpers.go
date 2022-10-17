package config

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

var (
	ErrDifferentArgumentsTypes = errors.New("src and dst must be of same type")
	ErrNonPointerArgument      = errors.New("dst must be a pointer")
)

func FindParentSubNode(configNode *yaml.Node, key, subKey string) *yaml.Node {
	if configNode.Content != nil || len(configNode.Content) == 0 {
		return nil
	}

	parentNodeIndex := getNodeIndex(configNode.Content[0].Content, key)
	if parentNodeIndex == -1 {
		return nil
	}
	parentNode := configNode.Content[0].Content[parentNodeIndex]

	if parentNode.Content != nil || len(parentNode.Content) == 0 {
		return nil
	}

	subNodeIndex := getNodeIndex(parentNode.Content[0].Content, subKey)
	if subNodeIndex == -1 {
		return nil
	}

	subNode := parentNode.Content[0].Content[subNodeIndex]

	return subNode
}

func FindParentNode(configNode *yaml.Node, key string) *yaml.Node {
	if configNode.Content == nil || len(configNode.Content) == 0 || configNode.Content[0] == nil || configNode.Content[0].Content == nil {
		return nil
	}
	configNode.Content[0].Style = 0
	parentNodeIndex := getNodeIndex(configNode.Content[0].Content, key)
	if parentNodeIndex == -1 {
		return nil
	}
	parentNode := configNode.Content[0].Content[parentNodeIndex]

	return parentNode
}

func FindNode(node *yaml.Node, key string) *yaml.Node {
	if node.Content == nil || len(node.Content) == 0 {
		return nil
	}
	nodeIndex := getNodeIndex(node.Content, key)
	if nodeIndex == -1 {
		return nil
	}
	foundNode := node.Content[nodeIndex]

	return foundNode
}

func getNodeIndex(node []*yaml.Node, key string) int {
	appIdx := -1
	for i, k := range node {
		if i%2 == 0 && k.Value == key {
			appIdx = i + 1
			break
		}
	}
	return appIdx
}

func CreateSequenceNode(key string) []*yaml.Node {
	keyNode := &yaml.Node{
		Kind:  yaml.ScalarNode,
		Style: 0,
		Value: key,
	}
	valueNode := &yaml.Node{
		Kind: yaml.SequenceNode,
	}

	return []*yaml.Node{keyNode, valueNode}
}

func CreateScalarNode(key, value string) []*yaml.Node {
	keyNode := &yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: key,
	}
	valueNode := &yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: value,
	}

	return []*yaml.Node{keyNode, valueNode}
}

func CreateMappingNode(key string) []*yaml.Node {
	keyNode := &yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: key,
	}
	valueNode := &yaml.Node{
		Kind: yaml.MappingNode,
	}

	return []*yaml.Node{keyNode, valueNode}
}

func EqualNodes(left, right *yaml.Node) (bool, error) {
	if left.Kind == yaml.ScalarNode && right.Kind == yaml.ScalarNode {
		return left.Value == right.Value, nil
	}
	return false, errors.New("equals on non-scalars not implemented!")
}

func MergeNodes(src, dst *yaml.Node, patchStrategy func(string) bool) error {
	if src.Kind != dst.Kind {
		return ErrDifferentArgumentsTypes
	}

	if dst != nil && reflect.ValueOf(dst).Kind() != reflect.Ptr {
		return ErrNonPointerArgument
	}

	switch src.Kind {
	case yaml.MappingNode:
		for i := 0; i < len(src.Content); i += 2 {
			fieldPath := src.Content[i].Value
			found := false
			for j := 0; j < len(dst.Content); j += 2 {
				fieldPath += dst.Content[i].Value
				if ok, _ := EqualNodes(src.Content[i], dst.Content[j]); ok {
					found = true
					fmt.Printf("Merginggg......%v - %v\n", src.Content[i].Value, dst.Content[j].Value)
					fmt.Println(fieldPath)
					if err := MergeNodes(src.Content[i+1], dst.Content[j+1], patchStrategy); err != nil {
						return errors.New("at key " + src.Content[i].Value + ": " + err.Error())
					}
					break
				}
			}
			if !found {
				fmt.Printf("Not Found Merginggg......%v\n", src.Content[i:i+2])
				fmt.Println(fieldPath)
				dst.Content = append(dst.Content, src.Content[i:i+2]...)
			}
		}
	case yaml.SequenceNode:
		dst.Content = append(dst.Content, src.Content...)
	case yaml.DocumentNode:
		err := MergeNodes(src.Content[0], dst.Content[0], patchStrategy)
		if err != nil {
			errors.New("at key " + src.Content[0].Value + ": " + err.Error())
		}
	case yaml.ScalarNode:
		if dst.Value != src.Value {
			dst.Value = src.Value
		}
	default:
		return errors.New("can only merge mapping and sequence nodes")
	}
	return nil
}

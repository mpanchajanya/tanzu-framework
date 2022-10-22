package nodeutils

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

func EqualNodes(left, right *yaml.Node) (bool, error) {
	if left.Kind == yaml.ScalarNode && right.Kind == yaml.ScalarNode {
		return left.Value == right.Value, nil
	}
	return false, errors.New("equals on non-scalars not implemented!")
}

func MergeNodes(src, dst *yaml.Node) error {
	if src.Kind != dst.Kind {
		return ErrDifferentArgumentsTypes
	}

	if dst != nil && reflect.ValueOf(dst).Kind() != reflect.Ptr {
		return ErrNonPointerArgument
	}

	switch src.Kind {
	case yaml.MappingNode:
		for i := 0; i < len(src.Content); i += 2 {
			found := false
			for j := 0; j < len(dst.Content); j += 2 {
				if ok, _ := EqualNodes(src.Content[i], dst.Content[j]); ok {
					found = true
					if err := MergeNodes(src.Content[i+1], dst.Content[j+1]); err != nil {
						return errors.New("at key " + src.Content[i].Value + ": " + err.Error())
					}
					break
				}
			}
			if !found {
				dst.Content = append(dst.Content, src.Content[i:i+2]...)
			}
		}
	case yaml.SequenceNode:
		if dst.Content[0].Kind == yaml.ScalarNode && src.Content[0].Kind == yaml.ScalarNode {
			dst.Content = append(dst.Content, src.Content...)
		}
	case yaml.DocumentNode:
		err := MergeNodes(src.Content[0], dst.Content[0])
		if err != nil {
			return errors.New("at key " + src.Content[0].Value + ": " + err.Error())
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

func ReplaceNodes(src, dst *yaml.Node, patchStrategyKey string, patchStrategies map[string]string) error {
	if src.Kind != dst.Kind {
		return ErrDifferentArgumentsTypes
	}

	if dst != nil && reflect.ValueOf(dst).Kind() != reflect.Ptr {
		return ErrNonPointerArgument
	}

	switch dst.Kind {
	case yaml.MappingNode:

		for i := 0; i < len(dst.Content); i += 2 {
			found := false
			key := patchStrategyKey
			for j := 0; j < len(src.Content); j += 2 {
				if ok, _ := EqualNodes(dst.Content[i], src.Content[j]); ok {
					key = key + "." + dst.Content[i].Value
					found = true
					fmt.Printf("found fields %v\n", key)
					if err := ReplaceNodes(src.Content[j+1], dst.Content[i+1], key, patchStrategies); err != nil {
						return errors.New("at key " + src.Content[i].Value + ": " + err.Error())
					}
					key = patchStrategyKey
					break
				}
			}
			if !found {
				key = key + "." + dst.Content[i].Value
				fmt.Printf("unknown found %v:-%v\n", key, dst.Content[i].Value)
				if patchStrategies[key] == "replace" {
					dst.Content = append(dst.Content[:i], dst.Content[i+1:]...)
					dst.Content = append(dst.Content[:i], dst.Content[i+1:]...)
					i--
					i--
				}

			}
		}
	case yaml.ScalarNode:
	case yaml.SequenceNode:
	case yaml.DocumentNode:
	default:
		return errors.New("can only merge mapping nodes")
	}
	return nil
}

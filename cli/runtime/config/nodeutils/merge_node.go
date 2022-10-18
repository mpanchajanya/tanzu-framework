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

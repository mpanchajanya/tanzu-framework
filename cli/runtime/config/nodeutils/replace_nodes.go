package nodeutils

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func ReplaceNodes(src, dst *yaml.Node, options *PatchStrategyOptions) error {
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
			key := options.Key
			for j := 0; j < len(src.Content); j += 2 {
				if ok, _ := EqualNodes(dst.Content[i], src.Content[j]); ok {
					found = true
					key = fmt.Sprintf("%v.%v", key, dst.Content[i].Value)
					fmt.Printf("found fields %v\n", key)
					if err := ReplaceNodes(src.Content[j+1], dst.Content[i+1], options); err != nil {
						return errors.New("at key " + src.Content[i].Value + ": " + err.Error())
					}
					key = options.Key
					break
				}
			}
			if !found {
				key = fmt.Sprintf("%v.%v", key, dst.Content[i].Value)
				fmt.Printf("unknown found %v:-%v\n", key, dst.Content[i].Value)
				if options.PatchStrategies[key] == "replace" {
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

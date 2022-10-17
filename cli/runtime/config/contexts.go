// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"

	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
	"gopkg.in/yaml.v3"
)

func GetContext(name string) (*configapi.Context, error) {
	node, err := GetClientConfigNode()
	if err != nil {
		return nil, err
	}

	return getContext(node, name)
}

func SetContext(c *configapi.Context, setCurrent bool) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	err = setContext(node, c)
	if err != nil {
		return err
	}

	if setCurrent {
		setCurrentContext(node, c)
	}

	s := convertContextToServer(c)
	err = setServer(node, s)
	if err != nil {
		return err
	}

	setCurrentServer(node, s.Name)

	return PersistNode(node)

}

func RemoveContext(name string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	ctx, err := getContext(node, name)
	if err != nil {
		return err
	}

	_, err = removeCurrentContext(node, ctx.Type)
	if err != nil {
		return err
	}
	_, err = removeContext(node, name)
	if err != nil {
		return err
	}

	_, err = removeServer(node, name)
	if err != nil {
		return err
	}

	err = removeCurrentServer(node, name)
	if err != nil {
		return err
	}

	return PersistNode(node)
}

func ContextExists(name string) (bool, error) {
	_, err := GetContext(name)
	return err == nil, nil
}

func GetCurrentContext(ctxType configapi.ContextType) (c *configapi.Context, err error) {
	node, err := GetClientConfigNode()
	if err != nil {
		return nil, err
	}

	return getCurrentContext(node, ctxType)
}

func SetCurrentContext(name string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	ctx, err := getContext(node, name)
	if err != nil {
		return err
	}

	setCurrentContext(node, ctx)
	setCurrentServer(node, name)

	return PersistNode(node)
}

func RemoveCurrentContext(ctxType configapi.ContextType) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	_, err = removeCurrentContext(node, ctxType)
	if err != nil {
		return err
	}

	c, err := getCurrentContext(node, ctxType)
	if err != nil {
		return err
	}

	err = removeCurrentServer(node, c.Name)
	if err != nil {
		return err
	}

	return PersistNode(node)

}

func EndpointFromContext(s *configapi.Context) (endpoint string, err error) {
	switch s.Type {
	case configapi.CtxTypeK8s:
		return s.ClusterOpts.Endpoint, nil
	case configapi.CtxTypeTMC:
		return s.GlobalOpts.Endpoint, nil
	default:
		return endpoint, fmt.Errorf("unknown server type %q", s.Type)
	}
}

func getContext(node *yaml.Node, name string) (*configapi.Context, error) {

	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}

	for _, ctx := range cfg.KnownContexts {
		if ctx.Name == name {
			return ctx, nil
		}
	}
	return nil, fmt.Errorf("could not find context %q", name)

}

func getCurrentContext(node *yaml.Node, ctxType configapi.ContextType) (*configapi.Context, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}
	return cfg.GetCurrentContext(ctxType)

}

func setContext(node *yaml.Node, c *configapi.Context) error {

	// Merge DiscoverSources separately
	copyOfDiscoverySources := c.DiscoverySources
	c.DiscoverySources = []configapi.PluginDiscovery{}
	fmt.Println(copyOfDiscoverySources)

	//convert context to node
	newContextNode, err := convertContextToNode(c)
	if err != nil {
		return err
	}

	contextsNode := FindParentNode(node, KeyContexts)

	if contextsNode == nil {
		//create context node and add to root node
		node.Content[0].Content = append(node.Content[0].Content, CreateSequenceNode(KeyContexts)...)
		contextsNode = FindParentNode(node, KeyContexts)
	}

	exists := false
	var result []*yaml.Node
	for _, contextNode := range contextsNode.Content {
		if index := getNodeIndex(contextNode.Content, "name"); index != -1 && contextNode.Content[index].Value == c.Name {
			exists = true

			for _, discoverySource := range copyOfDiscoverySources {
				err := setDiscoverySource(contextNode, discoverySource)
				if err != nil {
					return err
				}
			}
			//Merging
			err = MergeNodes(newContextNode.Content[0], contextNode, nil)
			if err != nil {
				return err
			}
			result = append(result, contextNode)
			continue
		}
		result = append(result, contextNode)
	}

	if !exists {
		result = append(result, newContextNode.Content[0])
	}

	contextsNode.Content = result

	return nil

}

func setCurrentContext(node *yaml.Node, ctx *configapi.Context) {
	currentContextNode := FindParentNode(node, KeyCurrentContext)

	if currentContextNode == nil {
		//create current context node and add to root node
		node.Content[0].Content = append(node.Content[0].Content, CreateMappingNode(KeyCurrentContext)...)
		currentContextNode = FindParentNode(node, KeyCurrentContext)
	}

	if index := getNodeIndex(currentContextNode.Content, string(ctx.Type)); index != -1 {
		currentContextNode.Content[index].Value = ctx.Name
	} else {
		currentContextNode.Content = append(currentContextNode.Content, CreateScalarNode(string(ctx.Type), ctx.Name)...)
	}

}

func removeCurrentContext(node *yaml.Node, ctxType configapi.ContextType) (ok bool, err error) {
	currentContextsNode := FindParentNode(node, KeyCurrentContext)
	if currentContextsNode == nil {
		return true, nil
	}

	var currentContexts []*yaml.Node
	for _, contextNode := range currentContextsNode.Content {
		if index := getNodeIndex(contextNode.Content, string(ctxType)); index != -1 {
			continue
		}
		currentContexts = append(currentContexts, contextNode)
	}

	if len(currentContexts) != 0 {
		currentContextsNode.Content = currentContexts
	}

	return true, nil
}

func removeContext(node *yaml.Node, name string) (ok bool, err error) {

	contextsNode := FindParentNode(node, KeyContexts)

	if contextsNode == nil {
		return true, nil
	}
	var contexts []*yaml.Node
	for _, contextNode := range contextsNode.Content {
		if index := getNodeIndex(contextNode.Content, "name"); index != -1 && contextNode.Content[index].Value == name {
			continue
		}
		contexts = append(contexts, contextNode)
	}

	if len(contexts) == 0 {
		contextsNode.Kind = yaml.ScalarNode
		contextsNode.Tag = "!!seq"
	} else {
		contextsNode.Content = contexts
	}

	return true, nil
}

package config

import (
	"fmt"
	"strings"

	"github.com/aunum/log"
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"

	"gopkg.in/yaml.v3"
)

func GetServer(name string) (*configapi.Server, error) {
	node, err := GetClientConfigNode()
	if err != nil {
		return nil, err
	}
	return getServer(node, name)
}

func ServerExists(name string) (bool, error) {
	node, err := GetClientConfigNode()
	if err != nil {
		return false, err
	}
	_, err = getServer(node, name)
	return err == nil, err
}

func GetCurrentServer() (*configapi.Server, error) {
	node, err := GetClientConfigNode()
	if err != nil {
		return nil, err
	}
	return getCurrentServer(node)
}

func SetCurrentServer(name string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	s, err := getServer(node, name)
	if err != nil {
		return err
	}

	setCurrentServer(node, name)

	// Front fill CurrentContext
	c := convertServerToContext(s)
	setCurrentContext(node, c)

	return PersistNode(node)
}

func RemoveCurrentServer(name string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	_, err = getServer(node, name)
	if err != nil {
		return err
	}

	err = removeCurrentServer(node, name)
	if err != nil {
		return err
	}

	//Front fill Context and CurrentContext
	c, err := getContext(node, name)
	if err != nil {
		return err
	}

	_, err = removeCurrentContext(node, c.Type)
	if err != nil {
		return err
	}

	//_, err = removeContext(node, c.Name)
	//if err != nil {
	//	return err
	//}

	return PersistNode(node)

}

func SetServer(s *configapi.Server, setCurrent bool) error {

	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	err = setServer(node, s)
	if err != nil {
		return err
	}

	if setCurrent {
		setCurrentServer(node, s.Name)
	}

	// Front fill Context and CurrentContext
	c := convertServerToContext(s)

	err = setContext(node, c)
	if err != nil {
		return err
	}

	if setCurrent {
		setCurrentContext(node, c)
	}

	return PersistNode(node)

}

func RemoveServer(name string) error {
	node, err := GetClientConfigNode()
	if err != nil {
		return err
	}

	_, err = getServer(node, name)
	if err != nil {
		return err

	}

	err = removeCurrentServer(node, name)
	if err != nil {
		return err

	}

	_, err = removeServer(node, name)
	if err != nil {
		return err

	}

	// Front fill Context and CurrentContext
	c, err := getContext(node, name)
	if err != nil {
		return err

	}

	_, err = removeCurrentContext(node, c.Type)
	if err != nil {
		return err

	}

	_, err = removeContext(node, name)
	if err != nil {
		return err

	}

	return PersistNode(node)
}

// GetDiscoverySources returns all discovery sources
// Includes standalone discovery sources and if server is available
// it also includes context based discovery sources as well
func GetDiscoverySources(serverName string) []configapi.PluginDiscovery {
	server, err := GetServer(serverName)
	if err != nil {
		log.Warningf("unknown server '%s', Unable to get server based discovery sources: %s", serverName, err.Error())
		return []configapi.PluginDiscovery{}
	}

	discoverySources := server.DiscoverySources
	// If current server type is management-cluster, then add
	// the default kubernetes discovery endpoint pointing to the
	// management-cluster kubeconfig
	if server.Type == configapi.ManagementClusterServerType {
		defaultClusterK8sDiscovery := configapi.PluginDiscovery{
			Kubernetes: &configapi.KubernetesDiscovery{
				Name:    fmt.Sprintf("default-%s", serverName),
				Path:    server.ManagementClusterOpts.Path,
				Context: server.ManagementClusterOpts.Context,
			},
		}
		discoverySources = append(discoverySources, defaultClusterK8sDiscovery)
	}

	// If the current server type is global, then add the default REST endpoint
	// for the discovery service
	if server.Type == configapi.GlobalServerType && server.GlobalOpts != nil {
		defaultRestDiscovery := configapi.PluginDiscovery{
			REST: &configapi.GenericRESTDiscovery{
				Name:     fmt.Sprintf("default-%s", serverName),
				Endpoint: appendURLScheme(server.GlobalOpts.Endpoint),
				BasePath: "v1alpha1/system/binaries/plugins",
			},
		}
		discoverySources = append(discoverySources, defaultRestDiscovery)
	}

	return discoverySources
}

func appendURLScheme(endpoint string) string {
	e := strings.Split(endpoint, ":")[0]
	if !strings.Contains(e, "https") {
		return fmt.Sprintf("https://%s", e)
	}
	return e
}

func setCurrentServer(node *yaml.Node, name string) {
	currentServerNode := FindParentNode(node, KeyCurrentServer)
	if currentServerNode == nil {
		node.Content[0].Content = append(node.Content[0].Content, CreateScalarNode(KeyCurrentServer, name)...)
		currentServerNode = FindParentNode(node, KeyCurrentServer)
	} else {
		currentServerNode.Value = name
	}
}

func getServer(node *yaml.Node, name string) (*configapi.Server, error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}

	for _, server := range cfg.KnownServers {
		if server.Name == name {
			return server, nil
		}
	}

	return nil, fmt.Errorf("could not find server %q", name)

}

func getCurrentServer(node *yaml.Node) (s *configapi.Server, err error) {
	cfg, err := convertNodeToClientConfig(node)
	if err != nil {
		return nil, err
	}
	for _, server := range cfg.KnownServers {
		if server.Name == cfg.CurrentServer {
			return server, nil
		}
	}
	return s, fmt.Errorf("current server %q not found in tanzu config", cfg.CurrentServer)
}

func removeCurrentServer(node *yaml.Node, name string) error {
	currentServerNode := FindParentNode(node, KeyCurrentServer)
	if currentServerNode == nil {
		return nil
	}

	for _, serverNode := range currentServerNode.Content {
		if index := getNodeIndex(serverNode.Content, name); index != -1 {
			serverNode.Content[index].Value = ""
		}
	}
	return nil
}

func removeServer(node *yaml.Node, name string) (ok bool, err error) {
	serversNode := FindParentNode(node, KeyServers)
	if serversNode == nil {
		return true, nil
	}

	var servers []*yaml.Node
	for _, serverNode := range serversNode.Content {
		if index := getNodeIndex(serverNode.Content, "name"); index != -1 && serverNode.Content[index].Value == name {
			continue
		}
		servers = append(servers, serverNode)
	}

	if len(servers) == 0 {
		serversNode.Kind = yaml.ScalarNode
		serversNode.Tag = "!!seq"
	} else {
		serversNode.Content = servers
	}

	return true, nil
}

func setServers(node *yaml.Node, servers []*configapi.Server) (bool, error) {

	for _, server := range servers {
		err := setServer(node, server)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func setServer(node *yaml.Node, s *configapi.Server) error {

	// Merge DiscoverSources separately
	copyOfDiscoverySources := s.DiscoverySources
	s.DiscoverySources = []configapi.PluginDiscovery{}
	fmt.Println(copyOfDiscoverySources)

	//convert server to node
	newNode, err := convertServerToNode(s)
	if err != nil {
		return err
	}

	serversNode := FindParentNode(node, KeyServers)

	if serversNode == nil {
		//create context node and add to root node
		node.Content[0].Content = append(node.Content[0].Content, CreateSequenceNode(KeyServers)...)
		serversNode = FindParentNode(node, KeyServers)
	}

	exists := false
	var result []*yaml.Node
	for _, serverNode := range serversNode.Content {
		if index := getNodeIndex(serverNode.Content, "name"); index != -1 && serverNode.Content[index].Value == s.Name {
			exists = true

			for _, discoverySource := range copyOfDiscoverySources {
				err := setDiscoverySource(serverNode, discoverySource)
				if err != nil {
					return err
				}
			}

			err = MergeNodes(newNode.Content[0], serverNode, nil)
			if err != nil {
				return err
			}
			result = append(result, serverNode)
			continue
		}
		result = append(result, serverNode)
	}

	if !exists {
		result = append(result, newNode.Content[0])
	}

	serversNode.Content = result

	return nil

}

// EndpointFromServer returns the endpoint from server.
func EndpointFromServer(s *configapi.Server) (endpoint string, err error) {
	switch s.Type {
	case configapi.ManagementClusterServerType:
		return s.ManagementClusterOpts.Endpoint, nil
	case configapi.GlobalServerType:
		return s.GlobalOpts.Endpoint, nil
	default:
		return endpoint, fmt.Errorf("unknown server type %q", s.Type)
	}
}

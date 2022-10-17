// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
	"gopkg.in/yaml.v3"
)

func setupDiscoverySources(t *testing.T) {
	LocalDirName = fmt.Sprintf(".tanzu-test")
	//cfg := &v1alpha1.ClientConfig{
	//	KnownContexts: []*v1alpha1.Context{
	//		{
	//			Name: "test-mc",
	//			Type: "k8s",
	//			ClusterOpts: &v1alpha1.ClusterServer{
	//				Endpoint:            "test-endpoint",
	//				Path:                "test-path",
	//				Context:             "test-context",
	//				IsManagementCluster: true,
	//			},
	//		},
	//		{
	//			Name: "test-tmc",
	//			Type: "tmc",
	//			GlobalOpts: &v1alpha1.GlobalServer{
	//				Endpoint: "test-endpoint",
	//			},
	//		},
	//	},
	//	CurrentContext: map[string]string{
	//		"k8s": "test-mc",
	//		"tmc": "test-tmc",
	//	},
	//}
	//
	//AcquireTanzuConfigLock()
	//defer ReleaseTanzuConfigLock()
	//err := PersistConfig(cfg)
	//require.NoError(t, err)
}

func TestSetDiscoverySource(t *testing.T) {
	setUpLocalDirName(t)
	// setUpClientConfigWithNoContexts(t)
	// cleanup()

	tests := []struct {
		name            string
		discoverySource v1alpha1.PluginDiscovery
		contextNode     *yaml.Node
		errStr          string
	}{
		{
			name: "success k8s",
			discoverySource: v1alpha1.PluginDiscovery{
				GCP: &v1alpha1.GCPDiscovery{
					Name:         "test",
					Bucket:       "updated-test-bucket",
					ManifestPath: "test-manifest-path",
				},
				ContextType: v1alpha1.CtxTypeTMC,
			},

			contextNode: &yaml.Node{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := setDiscoverySource(tc.contextNode, tc.discoverySource)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}

		})
	}

}

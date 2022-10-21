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

func TestSetDiscoverySource(t *testing.T) {
	func() {
		LocalDirName = fmt.Sprintf(".tanzu-test")
	}()

	defer func() {
		cleanupDir(LocalDirName)
	}()

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
			_, err := setDiscoverySource(tc.contextNode, tc.discoverySource)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}

		})
	}

}

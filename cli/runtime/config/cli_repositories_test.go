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

func TestSetRepository(t *testing.T) {
	func() {
		LocalDirName = fmt.Sprintf(".tanzu-test")
	}()

	defer func() {
		cleanupDir(LocalDirName)
	}()

	tests := []struct {
		name       string
		repository v1alpha1.PluginRepository
		cfg        *v1alpha1.ClientConfig
		repoNode   *yaml.Node
		errStr     string
	}{
		{
			name: "success k8s",
			repository: v1alpha1.PluginRepository{
				GCPPluginRepository: &v1alpha1.GCPPluginRepository{
					Name:       "test",
					BucketName: "bucket",
					RootPath:   "root-path",
				},
			},
			cfg:      &v1alpha1.ClientConfig{},
			repoNode: &yaml.Node{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			err := StoreClientConfig(tc.cfg)
			assert.NoError(t, err)

			err = SetCLIRepository(tc.repository)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}

		})
	}

}

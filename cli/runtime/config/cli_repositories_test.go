// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
)

func TestSetRepository(t *testing.T) {
	func() {
		LocalDirName = fmt.Sprintf(".tanzu-test")
	}()

	defer func() {
		cleanupDir(LocalDirName)
	}()

	tests := []struct {
		name    string
		cfg     *v1alpha1.ClientConfig
		in      v1alpha1.PluginRepository
		out     v1alpha1.PluginRepository
		persist bool
	}{
		{
			name: "should persist repository",
			cfg:  &v1alpha1.ClientConfig{},
			in: v1alpha1.PluginRepository{
				GCPPluginRepository: &v1alpha1.GCPPluginRepository{
					Name:       "test",
					BucketName: "bucket",
					RootPath:   "root-path",
				},
			},
			out: v1alpha1.PluginRepository{
				GCPPluginRepository: &v1alpha1.GCPPluginRepository{
					Name:       "test",
					BucketName: "bucket",
					RootPath:   "root-path",
				},
			},
			persist: true,
		},
		{
			name: "should not persist same repo",
			cfg: &v1alpha1.ClientConfig{
				ClientOptions: &v1alpha1.ClientOptions{
					CLI: &v1alpha1.CLIOptions{
						Repositories: []v1alpha1.PluginRepository{
							{
								GCPPluginRepository: &v1alpha1.GCPPluginRepository{
									Name:       "test",
									BucketName: "bucket",
									RootPath:   "root-path",
								},
							},
						},
					},
				},
			},
			in: v1alpha1.PluginRepository{
				GCPPluginRepository: &v1alpha1.GCPPluginRepository{
					Name:       "test",
					BucketName: "bucket",
					RootPath:   "root-path",
				},
			},
			out: v1alpha1.PluginRepository{
				GCPPluginRepository: &v1alpha1.GCPPluginRepository{
					Name:       "test",
					BucketName: "bucket",
					RootPath:   "root-path",
				},
			},
			persist: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			err := StoreClientConfig(tc.cfg)
			assert.NoError(t, err)

			persist, err := SetCLIRepository(tc.in)
			assert.Equal(t, tc.persist, persist)

			r, err := GetCLIRepository(tc.out.GCPPluginRepository.Name)
			assert.Equal(t, tc.out.GCPPluginRepository.Name, r.GCPPluginRepository.Name)
		})
	}

}

// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
)

func TestGetCLIDiscoverySources(t *testing.T) {

	// setup
	func() {
		LocalDirName = fmt.Sprintf(".tanzu-test")
	}()

	defer func() {
		cleanupDir(LocalDirName)
	}()

	tests := []struct {
		name   string
		in     *configapi.ClientConfig
		out    []configapi.PluginDiscovery
		errStr string
	}{
		{
			name: "success k8s",
			in: &configapi.ClientConfig{
				ClientOptions: &configapi.ClientOptions{
					CLI: &configapi.CLIOptions{
						DiscoverySources: []v1alpha1.PluginDiscovery{
							{
								GCP: &v1alpha1.GCPDiscovery{
									Name:         "test",
									Bucket:       "updated-test-bucket",
									ManifestPath: "test-manifest-path",
								},
								ContextType: v1alpha1.CtxTypeTMC,
							},
						},
					},
				},
			},
			out: []v1alpha1.PluginDiscovery{
				{
					GCP: &v1alpha1.GCPDiscovery{
						Name:         "test",
						Bucket:       "updated-test-bucket",
						ManifestPath: "test-manifest-path",
					},
					ContextType: v1alpha1.CtxTypeTMC,
				},
			},
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := StoreClientConfig(spec.in)
			if err != nil {
				fmt.Printf("StoreClientConfigV2 errors: %v\n", err)
			}

			c, err := GetCLIDiscoverySources()
			if err != nil {
				fmt.Printf("errors: %v\n", err)
			}

			assert.Equal(t, spec.out, c)
			assert.NoError(t, err)

		})
	}
}

func TestGetCLIDiscoverySource(t *testing.T) {

	// setup

	func() {
		LocalDirName = fmt.Sprintf(".tanzu-test")
	}()

	defer func() {
		cleanupDir(LocalDirName)
	}()

	tests := []struct {
		name string
		in   *configapi.ClientConfig
		out  *configapi.PluginDiscovery
	}{
		{
			name: "success test",
			in: &configapi.ClientConfig{
				ClientOptions: &configapi.ClientOptions{
					CLI: &configapi.CLIOptions{
						DiscoverySources: []v1alpha1.PluginDiscovery{
							{
								GCP: &v1alpha1.GCPDiscovery{
									Name:         "test",
									Bucket:       "updated-test-bucket",
									ManifestPath: "test-manifest-path",
								},
								ContextType: v1alpha1.CtxTypeTMC,
							},
						},
					},
				},
			},
			out: &v1alpha1.PluginDiscovery{
				GCP: &v1alpha1.GCPDiscovery{
					Name:         "test",
					Bucket:       "updated-test-bucket",
					ManifestPath: "test-manifest-path",
				},
				ContextType: v1alpha1.CtxTypeTMC,
			},
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := StoreClientConfig(spec.in)
			if err != nil {
				fmt.Printf("StoreClientConfigV2 errors: %v\n", err)
			}

			c, err := GetCLIDiscoverySource("test")
			if err != nil {
				fmt.Printf("errors: %v\n", err)
			}

			assert.Equal(t, spec.out, c)
			assert.NoError(t, err)

		})
	}
}

func TestSetCLIDiscoverySource(t *testing.T) {

	// setup
	func() {
		LocalDirName = fmt.Sprintf(".tanzu-test")
	}()

	defer func() {
		cleanupDir(LocalDirName)
	}()

	tests := []struct {
		name  string
		src   *configapi.ClientConfig
		input *configapi.PluginDiscovery
	}{
		{
			name: "success add test",
			src: &configapi.ClientConfig{
				ClientOptions: &configapi.ClientOptions{
					CLI: &configapi.CLIOptions{},
				},
			},
			input: &v1alpha1.PluginDiscovery{
				GCP: &v1alpha1.GCPDiscovery{
					Name:         "test",
					Bucket:       "test-bucket",
					ManifestPath: "test-manifest-path",
				},
				ContextType: v1alpha1.CtxTypeTMC,
			},
		},
		{
			name: "success update test",
			src: &configapi.ClientConfig{
				ClientOptions: &configapi.ClientOptions{
					CLI: &configapi.CLIOptions{
						DiscoverySources: []v1alpha1.PluginDiscovery{
							{
								GCP: &v1alpha1.GCPDiscovery{
									Name:         "test",
									Bucket:       "test-bucket",
									ManifestPath: "test-manifest-path",
								},
								ContextType: v1alpha1.CtxTypeTMC,
							},
						},
					},
				},
			},
			input: &v1alpha1.PluginDiscovery{
				GCP: &v1alpha1.GCPDiscovery{
					Name:         "test",
					Bucket:       "updated-test-bucket",
					ManifestPath: "test-manifest-path",
				},
				ContextType: v1alpha1.CtxTypeTMC,
			},
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := StoreClientConfig(spec.src)
			if err != nil {
				fmt.Printf("StoreClientConfigV2 errors: %v\n", err)
			}

			err = SetCLIDiscoverySource(*spec.input)
			if err != nil {
				fmt.Printf("errors: %v\n", err)
			}

			c, err := GetCLIDiscoverySource(spec.input.GCP.Name)

			assert.Equal(t, spec.input, c)
			assert.NoError(t, err)

		})
	}
}
func TestDeleteCLIDiscoverySource(t *testing.T) {

	// setup
	func() {
		LocalDirName = fmt.Sprintf(".tanzu-test")
	}()

	//defer func() {
	//	cleanupDir(LocalDirName)
	//}()

	tests := []struct {
		name    string
		src     *configapi.ClientConfig
		input   string
		deleted bool
	}{
		{
			name: "should return true on deleting non existing item",
			src: &configapi.ClientConfig{
				ClientOptions: &configapi.ClientOptions{
					CLI: &configapi.CLIOptions{},
				},
			},
			input:   "text-mc",
			deleted: true,
		},
		{
			name: "should return true on deleting existing item",
			src: &configapi.ClientConfig{
				ClientOptions: &configapi.ClientOptions{
					CLI: &configapi.CLIOptions{
						DiscoverySources: []v1alpha1.PluginDiscovery{
							{
								GCP: &v1alpha1.GCPDiscovery{
									Name:         "test",
									Bucket:       "test-bucket",
									ManifestPath: "test-manifest-path",
								},
								ContextType: v1alpha1.CtxTypeTMC,
							},
						},
					},
				},
			},
			input:   "test",
			deleted: true,
		},
		{
			name: "should return true on deleting non existing item",
			src: &configapi.ClientConfig{
				ClientOptions: &configapi.ClientOptions{
					CLI: &configapi.CLIOptions{
						DiscoverySources: []v1alpha1.PluginDiscovery{
							{
								GCP: &v1alpha1.GCPDiscovery{
									Name:         "test",
									Bucket:       "test-bucket",
									ManifestPath: "test-manifest-path",
								},
								ContextType: v1alpha1.CtxTypeTMC,
							},
						},
					},
				},
			},
			input:   "test-notfound",
			deleted: true,
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := StoreClientConfig(spec.src)
			if err != nil {
				fmt.Printf("StoreClientConfigV2 errors: %v\n", err)
			}

			ok, err := DeleteCLIDiscoverySource(spec.input)

			assert.Equal(t, spec.deleted, ok)
			assert.NoError(t, err)

		})
	}
}

func TestIntegrationSetGetDeleteCLIDiscoverySource(t *testing.T) {
	// setup
	func() {
		LocalDirName = fmt.Sprintf(".tanzu-test")
	}()

	defer func() {
		cleanupDir(LocalDirName)
	}()

	tests := []struct {
		name    string
		src     *configapi.ClientConfig
		input   string
		deleted bool
	}{
		{
			name: "should get delete set",
			src: &configapi.ClientConfig{
				ClientOptions: &configapi.ClientOptions{
					CLI: &configapi.CLIOptions{
						DiscoverySources: []v1alpha1.PluginDiscovery{
							{
								GCP: &v1alpha1.GCPDiscovery{
									Name:         "test",
									Bucket:       "test-bucket",
									ManifestPath: "test-manifest-path",
								},
								ContextType: v1alpha1.CtxTypeTMC,
							},
						},
					},
				},
			},
			input:   "test",
			deleted: true,
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := StoreClientConfig(spec.src)
			if err != nil {
				fmt.Printf("StoreClientConfigV2 errors: %v\n", err)
			}

			ds, err := GetCLIDiscoverySource(spec.input)
			assert.Equal(t, spec.src.ClientOptions.CLI.DiscoverySources[0].GCP, ds.GCP)
			assert.NoError(t, err)

			ok, err := DeleteCLIDiscoverySource(spec.input)
			assert.Equal(t, spec.deleted, ok)
			assert.NoError(t, err)

			err = SetCLIDiscoverySource(spec.src.ClientOptions.CLI.DiscoverySources[0])
			assert.NoError(t, err)
		})
	}

}

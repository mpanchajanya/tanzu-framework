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
			name: "success get all",
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
			name: "success get",
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
		name    string
		src     *configapi.ClientConfig
		input   *configapi.PluginDiscovery
		persist bool
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
			persist: true,
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
			persist: true,
		},
		{
			name: "should not persist same test",
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
					Bucket:       "test-bucket",
					ManifestPath: "test-manifest-path",
				},
				ContextType: v1alpha1.CtxTypeTMC,
			},
			persist: false,
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := StoreClientConfig(spec.src)
			if err != nil {
				fmt.Printf("StoreClientConfigV2 errors: %v\n", err)
			}

			persist, err := SetCLIDiscoverySource(*spec.input)
			if err != nil {
				fmt.Printf("errors: %v\n", err)
			}

			assert.Equal(t, persist, spec.persist)
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

			err = DeleteCLIDiscoverySource(spec.input)

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

			err = DeleteCLIDiscoverySource(spec.input)
			assert.NoError(t, err)

			_, err = SetCLIDiscoverySource(spec.src.ClientOptions.CLI.DiscoverySources[0])
			assert.NoError(t, err)
		})
	}

}

func TestSetCLIDiscoverySourceLocalMulti(t *testing.T) {
	// setup
	func() {
		LocalDirName = fmt.Sprintf(".tanzu-test")
	}()

	defer func() {
		cleanupDir(LocalDirName)
	}()

	src := &configapi.ClientConfig{
		ClientOptions: &configapi.ClientOptions{
			CLI: &configapi.CLIOptions{},
		},
	}

	input := v1alpha1.PluginDiscovery{
		Local: &v1alpha1.LocalDiscovery{
			Name: "admin-local",
			Path: "admin",
		},
	}

	input2 := v1alpha1.PluginDiscovery{
		Local: &v1alpha1.LocalDiscovery{
			Name: "default-local",
			Path: "standalone",
		},
		ContextType: "k8s",
	}

	updateInput2 := v1alpha1.PluginDiscovery{
		Local: &v1alpha1.LocalDiscovery{
			Name: "default-local",
			Path: "standalone-updated",
		},
		ContextType: "k8s",
	}
	err := StoreClientConfig(src)
	if err != nil {
		fmt.Printf("StoreClientConfigV2 errors: %v\n", err)
	}

	_, err = SetCLIDiscoverySource(input)
	assert.NoError(t, err)

	c, err := GetCLIDiscoverySource("admin-local")
	if err != nil {
		fmt.Printf("errors: %v\n", err)
	}

	assert.Equal(t, input.Local, c.Local)
	assert.NoError(t, err)

	_, err = SetCLIDiscoverySource(input2)
	assert.NoError(t, err)

	c, err = GetCLIDiscoverySource("default-local")
	if err != nil {
		fmt.Printf("errors: %v\n", err)
	}

	assert.Equal(t, input2.Local, c.Local)
	assert.NoError(t, err)

	//Update Input2
	_, err = SetCLIDiscoverySource(updateInput2)
	assert.NoError(t, err)

	c, err = GetCLIDiscoverySource("default-local")
	if err != nil {
		fmt.Printf("errors: %v\n", err)
	}

	assert.Equal(t, updateInput2.Local, c.Local)
	assert.NoError(t, err)

}

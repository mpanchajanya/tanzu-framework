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

func TestCLIDiscoverySources(t *testing.T) {

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
		out  []configapi.PluginDiscovery
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
func TestCLIDiscoverySource(t *testing.T) {

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
		name string
		in   *configapi.ClientConfig
		out  *configapi.PluginDiscovery
	}{
		{
			name: "success add test",
			in: &configapi.ClientConfig{
				ClientOptions: &configapi.ClientOptions{
					CLI: &configapi.CLIOptions{},
				},
			},
			out: &v1alpha1.PluginDiscovery{
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
			in: &configapi.ClientConfig{
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

			err = SetCLIDiscoverySource(*spec.out)
			if err != nil {
				fmt.Printf("errors: %v\n", err)
			}

			c, err := GetCLIDiscoverySource(spec.out.GCP.Name)

			assert.Equal(t, spec.out, c)
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
		name   string
		in     *configapi.ClientConfig
		out    *configapi.PluginDiscovery
		expect bool
	}{
		{
			name: "success add test",
			in: &configapi.ClientConfig{
				ClientOptions: &configapi.ClientOptions{
					CLI: &configapi.CLIOptions{},
				},
			},
			out: &v1alpha1.PluginDiscovery{
				GCP: &v1alpha1.GCPDiscovery{
					Name:         "test",
					Bucket:       "test-bucket",
					ManifestPath: "test-manifest-path",
				},
				ContextType: v1alpha1.CtxTypeTMC,
			},
			expect: false,
		},
		{
			name: "success update test",
			in: &configapi.ClientConfig{
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

			err = SetCLIDiscoverySource(*spec.out)
			if err != nil {
				fmt.Printf("errors: %v\n", err)
			}

			c, err := GetCLIDiscoverySource(spec.out.GCP.Name)

			assert.Equal(t, spec.out, c)
			assert.NoError(t, err)

		})
	}
}

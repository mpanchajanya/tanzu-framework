// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
	"github.com/vmware-tanzu/tanzu-framework/cli/runtime/config/nodeutils"
	"gopkg.in/yaml.v3"
)

func TestSetContextWithDiscoverySourceWithReplaceStrategy(t *testing.T) {
	// setup
	func() {
		LocalDirName = fmt.Sprintf(".tanzu-test")
	}()

	//defer func() {
	//	cleanupDir(LocalDirName)
	//}()

	tests := []struct {
		name    string
		src     *v1alpha1.ClientConfig
		ctx     *v1alpha1.Context
		current bool
		errStr  string
	}{
		{
			name: "should add new context with new discovery sources to empty client config",
			src:  &v1alpha1.ClientConfig{},
			ctx: &v1alpha1.Context{
				Name: "test-mc",
				Type: "k8s",
				ClusterOpts: &v1alpha1.ClusterServer{
					Endpoint:            "test-endpoint",
					Path:                "test-path",
					Context:             "test-context",
					IsManagementCluster: true,
				},
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
			current: true,
		},
		{
			name: "should update existing context",
			src: &v1alpha1.ClientConfig{
				KnownContexts: []*v1alpha1.Context{
					{
						Name: "test-mc",
						Type: "k8s",
						ClusterOpts: &v1alpha1.ClusterServer{
							Endpoint:            "test-endpoint",
							Path:                "test-path",
							Context:             "test-context",
							IsManagementCluster: true,
						},
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
				KnownServers: []*v1alpha1.Server{
					{
						Name: "test-mc",
						Type: v1alpha1.ManagementClusterServerType,
						ManagementClusterOpts: &v1alpha1.ManagementClusterServer{
							Endpoint: "test-endpoint",
							Path:     "test-path",
							Context:  "test-context",
						},
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
				CurrentServer: "test-mc",
				CurrentContext: map[v1alpha1.ContextType]string{
					v1alpha1.CtxTypeK8s: "test-mc",
				},
			},
			ctx: &v1alpha1.Context{
				Name: "test-mc",
				Type: "k8s",
				ClusterOpts: &v1alpha1.ClusterServer{
					Endpoint:            "updated-test-endpoint",
					Path:                "updated-test-path",
					Context:             "updated-test-context",
					IsManagementCluster: true,
				},
				DiscoverySources: []v1alpha1.PluginDiscovery{
					{
						GCP: &v1alpha1.GCPDiscovery{
							Name:         "test",
							Bucket:       "updated-test-bucket",
							ManifestPath: "updated-test-manifest-path",
						},
						ContextType: v1alpha1.CtxTypeTMC,
					},
				},
			},
			current: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			////setup data
			//node, err := nodeutils.ConvertToNode(tc.src)
			//assert.NoError(t, err)
			//err = PersistNode(node)
			//assert.NoError(t, err)

			err := SetContext(tc.ctx, tc.current)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}

			ok, err := ContextExists(tc.ctx.Name)
			assert.True(t, ok)
			assert.NoError(t, err)
		})
	}

}

func TestSetContextWithDiscoverySource(t *testing.T) {
	// setup
	func() {
		LocalDirName = fmt.Sprintf(".tanzu-test")
	}()

	//defer func() {
	//	cleanupDir(LocalDirName)
	//}()

	tests := []struct {
		name    string
		src     *v1alpha1.ClientConfig
		ctx     *v1alpha1.Context
		current bool
		errStr  string
	}{
		{
			name: "should add new context with new discovery sources to empty client config",
			src:  &v1alpha1.ClientConfig{},
			ctx: &v1alpha1.Context{
				Name: "test-mc",
				Type: "k8s",
				ClusterOpts: &v1alpha1.ClusterServer{
					Endpoint:            "test-endpoint",
					Path:                "test-path",
					Context:             "test-context",
					IsManagementCluster: true,
				},
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
			current: true,
		},
		{
			name: "should update existing context",
			src: &v1alpha1.ClientConfig{
				KnownContexts: []*v1alpha1.Context{
					{
						Name: "test-mc",
						Type: "k8s",
						ClusterOpts: &v1alpha1.ClusterServer{
							Endpoint:            "test-endpoint",
							Path:                "test-path",
							Context:             "test-context",
							IsManagementCluster: true,
						},
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
				KnownServers: []*v1alpha1.Server{
					{
						Name: "test-mc",
						Type: v1alpha1.ManagementClusterServerType,
						ManagementClusterOpts: &v1alpha1.ManagementClusterServer{
							Endpoint: "test-endpoint",
							Path:     "test-path",
							Context:  "test-context",
						},
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
				CurrentServer: "test-mc",
				CurrentContext: map[v1alpha1.ContextType]string{
					v1alpha1.CtxTypeK8s: "test-mc",
				},
			},
			ctx: &v1alpha1.Context{
				Name: "test-mc",
				Type: "k8s",
				ClusterOpts: &v1alpha1.ClusterServer{
					Endpoint:            "updated-test-endpoint",
					Path:                "updated-test-path",
					Context:             "updated-test-context",
					IsManagementCluster: true,
				},
				DiscoverySources: []v1alpha1.PluginDiscovery{
					{
						GCP: &v1alpha1.GCPDiscovery{
							Name:         "test",
							Bucket:       "updated-test-bucket",
							ManifestPath: "updated-test-manifest-path",
						},
						ContextType: v1alpha1.CtxTypeTMC,
					},
				},
			},
			current: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			//setup data
			//node, err := nodeutils.ConvertToNode(tc.src)
			//assert.NoError(t, err)
			//err = PersistNode(node)
			//assert.NoError(t, err)

			err := SetContext(tc.ctx, tc.current)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}

			ok, err := ContextExists(tc.ctx.Name)
			assert.True(t, ok)
			assert.NoError(t, err)
		})
	}

}

func TestGetContext(t *testing.T) {
	setupForGetContext(t)

	tcs := []struct {
		name    string
		ctxName string
		errStr  string
	}{
		{
			name:    "success k8s",
			ctxName: "test-mc",
		},
		{
			name:    "success tmc",
			ctxName: "test-tmc",
		},
		{
			name:    "failure",
			ctxName: "test",
			errStr:  "could not find context \"test\"",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			c, err := GetContext(tc.ctxName)
			if err != nil {
				fmt.Printf("errors: %v\n", err)
			}
			if tc.errStr == "" {
				assert.Equal(t, tc.ctxName, c.Name)
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}
		})
	}
}

func setupForGetContext(t *testing.T) {
	//setup
	cfg := &v1alpha1.ClientConfig{
		KnownContexts: []*v1alpha1.Context{
			{
				Name: "test-mc",
				Type: "k8s",
				ClusterOpts: &v1alpha1.ClusterServer{
					Endpoint:            "test-endpoint",
					Path:                "test-path",
					Context:             "test-context",
					IsManagementCluster: true,
				},
			},
			{
				Name: "test-tmc",
				Type: "tmc",
				GlobalOpts: &v1alpha1.GlobalServer{
					Endpoint: "test-endpoint",
				},
			},
		},
		CurrentContext: map[v1alpha1.ContextType]string{
			"k8s": "test-mc",
			"tmc": "test-tmc",
		},
	}

	func() {
		LocalDirName = fmt.Sprintf(".tanzu-test")
		err := StoreClientConfig(cfg)
		assert.NoError(t, err)
	}()

}

func TestContextExists(t *testing.T) {
	setupForGetContext(t)

	defer func() {
		cleanupDir(LocalDirName)
	}()

	tcs := []struct {
		name    string
		ctxName string
		ok      bool
	}{
		{
			name:    "success k8s",
			ctxName: "test-mc",
			ok:      true,
		},
		{
			name:    "success tmc",
			ctxName: "test-tmc",
			ok:      true,
		},
		{
			name:    "failure",
			ctxName: "test",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ok, err := ContextExists(tc.ctxName)
			assert.Equal(t, tc.ok, ok)
			assert.NoError(t, err)
		})
	}
}

func TestSetContext(t *testing.T) {
	// setup
	func() {
		LocalDirName = fmt.Sprintf(".tanzu-test")
	}()

	defer func() {
		cleanupDir(LocalDirName)
	}()

	tcs := []struct {
		name    string
		src     *v1alpha1.ClientConfig
		srcNode *yaml.Node
		ctx     *v1alpha1.Context
		current bool
		errStr  string
	}{
		{
			name:    "should add new context and set current to empty config",
			src:     &v1alpha1.ClientConfig{},
			srcNode: &yaml.Node{},
			ctx: &v1alpha1.Context{
				Name: "test-mc",
				Type: "k8s",
				ClusterOpts: &v1alpha1.ClusterServer{
					Endpoint:            "test-endpoint",
					Path:                "test-path",
					Context:             "test-context",
					IsManagementCluster: true,
				},
			},
			current: true,
		},
		{
			name: "should update existing context",
			src: &v1alpha1.ClientConfig{
				CurrentContext: map[v1alpha1.ContextType]string{
					v1alpha1.CtxTypeK8s: "test-mc",
				},
				CurrentServer: "test-mc",
				KnownContexts: []*v1alpha1.Context{
					{
						Name: "test-mc",
						Type: "k8s",
						ClusterOpts: &v1alpha1.ClusterServer{
							Endpoint:            "test-endpoint",
							Path:                "test-path",
							Context:             "test-context",
							IsManagementCluster: true,
						},
					},
				},
			},
			ctx: &v1alpha1.Context{
				Name: "test-mc",
				Type: "k8s",
				ClusterOpts: &v1alpha1.ClusterServer{
					Endpoint:            "test-endpoint-updated",
					Path:                "test-path",
					Context:             "test-context",
					IsManagementCluster: true,
				},
			},
		},
		{
			name: "success k8s not_current",
			src:  &v1alpha1.ClientConfig{},
			ctx: &v1alpha1.Context{
				Name: "test-mc2",
				Type: "k8s",
				ClusterOpts: &v1alpha1.ClusterServer{
					Endpoint:            "test-endpoint",
					Path:                "test-path",
					Context:             "test-context",
					IsManagementCluster: true,
				},
			},
		},
		{
			name: "success tmc current",
			src:  &v1alpha1.ClientConfig{},
			ctx: &v1alpha1.Context{
				Name: "test-tmc1",
				Type: "tmc",
				GlobalOpts: &v1alpha1.GlobalServer{
					Endpoint: "test-endpoint",
				},
			},
			current: true,
		},
		{
			name: "success tmc not_current",
			ctx: &v1alpha1.Context{
				Name: "test-tmc2",
				Type: "tmc",
				GlobalOpts: &v1alpha1.GlobalServer{
					Endpoint: "test-endpoint",
				},
			},
		},
		{
			name: "success update test-mc",
			ctx: &v1alpha1.Context{
				Name: "test-mc",
				Type: "k8s",
				ClusterOpts: &v1alpha1.ClusterServer{
					Endpoint:            "good-test-endpoint",
					Path:                "updated-test-path",
					Context:             "updated-test-context",
					IsManagementCluster: true,
				},
			},
		},
		{
			name: "success update tmc",
			ctx: &v1alpha1.Context{
				Name: "test-tmc",
				Type: "tmc",
				GlobalOpts: &v1alpha1.GlobalServer{
					Endpoint: "updated-test-endpoint",
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			//setup data
			node, err := nodeutils.ConvertToNode(tc.src)
			assert.NoError(t, err)
			err = PersistNode(node)
			assert.NoError(t, err)

			//perform test
			err = SetContext(tc.ctx, tc.current)
			fmt.Printf("eeeeee %v\n", err)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}

			ctx, err := GetContext(tc.ctx.Name)
			assert.NoError(t, err)
			assert.Equal(t, tc.ctx.Name, ctx.Name)

			s, err := GetServer(tc.ctx.Name)

			assert.NoError(t, err)
			assert.Equal(t, tc.ctx.Name, s.Name)

		})
	}
}

func TestRemoveContext(t *testing.T) {
	// setup
	setupForGetContext(t)

	defer func() {
		cleanupDir(LocalDirName)
	}()

	tcs := []struct {
		name    string
		ctxName string
		ctxType v1alpha1.ContextType
		errStr  string
	}{
		{
			name:    "success k8s",
			ctxName: "test-mc",
			ctxType: "k8s",
		},
		{
			name:    "success tmc",
			ctxName: "test-tmc",
			ctxType: "tmc",
		},
		{
			name:    "failure",
			ctxName: "test",
			errStr:  "context test not found",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if tc.errStr == "" {
				ok, err := ContextExists(tc.ctxName)
				require.True(t, ok)
				require.NoError(t, err)
			}

			err := RemoveContext(tc.ctxName)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}

			ok, err := ContextExists(tc.ctxName)
			assert.False(t, ok)
			assert.NoError(t, err)
			ok, err = ServerExists(tc.ctxName)
			assert.False(t, ok)
		})
	}
}

func TestSetCurrentContext(t *testing.T) {
	// setup
	setupForGetContext(t)

	defer func() {
		cleanupDir(LocalDirName)
	}()

	tcs := []struct {
		name    string
		ctxType v1alpha1.ContextType
		ctxName string
		errStr  string
	}{
		{
			name:    "success k8s",
			ctxName: "test-mc1",
			ctxType: "k8s",
		},
		{
			name:    "success tmc",
			ctxName: "test-tmc",
			ctxType: "tmc",
		},
		{
			name:    "failure",
			ctxName: "test",
			errStr:  "could not find context \"test\"",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := SetCurrentContext(tc.ctxName)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.errStr)
			}

			currCtx, err := GetCurrentContext(tc.ctxType)
			if tc.errStr == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.ctxName, currCtx.Name)
			} else {
				assert.Error(t, err)
			}
			currSrv, err := GetCurrentServer()
			assert.NoError(t, err)
			if tc.errStr == "" {
				assert.Equal(t, tc.ctxName, currSrv.Name)
			}
		})
	}
}

package config

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
)

func TestIsFeatureEnabled(t *testing.T) {
	// setup
	func() {
		LocalDirName = fmt.Sprintf(".tanzu-test")
	}()

	defer func() {
		cleanupDir(LocalDirName)
	}()

	tests := []struct {
		name    string
		feature map[string]v1alpha1.FeatureMap
		plugin  string
		key     string
	}{
		{
			name: "success context-aware-cli-for-plugins",
			feature: map[string]v1alpha1.FeatureMap{
				"global": {
					"context-aware-cli-for-plugins": "true",
				},
			},
			plugin: "global",
			key:    "context-aware-cli-for-plugins",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			cfg := &v1alpha1.ClientConfig{
				ClientOptions: &v1alpha1.ClientOptions{
					Features: tc.feature,
				},
			}

			err := StoreClientConfig(cfg)
			assert.NoError(t, err)

			ok, err := IsFeatureEnabled(tc.plugin, tc.key)

			assert.NoError(t, err)
			assert.Equal(t, ok, true)
		})
	}

}

func TestSetAndDeleteFeature(t *testing.T) {
	// setup
	func() {
		LocalDirName = fmt.Sprintf(".tanzu-test")
	}()

	defer func() {
		cleanupDir(LocalDirName)
	}()

	tests := []struct {
		name    string
		feature map[string]v1alpha1.FeatureMap
		plugin  string
		key     string
		value   bool
	}{
		{
			name: "success context-aware-cli-for-plugins",
			feature: map[string]v1alpha1.FeatureMap{
				"global": {
					"sample":                        "true",
					"context-aware-cli-for-plugins": "true",
				},
			},
			plugin: "global",
			key:    "context-aware-cli-for-plugins",
			value:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			cfg := &v1alpha1.ClientConfig{
				ClientOptions: &v1alpha1.ClientOptions{
					Features: tc.feature,
				},
			}

			err := StoreClientConfig(cfg)
			assert.NoError(t, err)

			err = SetFeature(tc.plugin, tc.key, strconv.FormatBool(tc.value))
			assert.NoError(t, err)

			ok, err := IsFeatureEnabled(tc.plugin, tc.key)
			assert.NoError(t, err)
			assert.Equal(t, ok, tc.value)

			err = DeleteFeature(tc.plugin, tc.key)
			assert.NoError(t, err)

			ok, err = IsFeatureEnabled(tc.plugin, tc.key)
			assert.Equal(t, ok, tc.value)

			err = SetFeature(tc.plugin, tc.key, strconv.FormatBool(tc.value))
			assert.NoError(t, err)
		})
	}

}

func TestSetFeature(t *testing.T) {
	// setup
	func() {
		LocalDirName = fmt.Sprintf(".tanzu-test")
	}()

	defer func() {
		cleanupDir(LocalDirName)
	}()

	tests := []struct {
		name    string
		feature map[string]v1alpha1.FeatureMap
		plugin  string
		key     string
		value   bool
	}{
		{
			name: "success context-aware-cli-for-plugins",
			feature: map[string]v1alpha1.FeatureMap{
				"global": {
					"context-aware-cli-for-plugins": "true",
				},
			},
			plugin: "global",
			key:    "context-aware-cli-for-plugins",
			value:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			cfg := &v1alpha1.ClientConfig{
				ClientOptions: &v1alpha1.ClientOptions{
					Features: tc.feature,
				},
			}

			err := StoreClientConfig(cfg)
			assert.NoError(t, err)

			err = SetFeature(tc.plugin, tc.key, strconv.FormatBool(tc.value))
			assert.NoError(t, err)

			ok, err := IsFeatureEnabled(tc.plugin, tc.key)
			assert.NoError(t, err)
			assert.Equal(t, ok, tc.value)
		})
	}

}

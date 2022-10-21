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
		persist bool
	}{
		{
			name: "success context-aware-cli-for-plugins",
			feature: map[string]v1alpha1.FeatureMap{
				"global": {
					"sample":                        "true",
					"context-aware-cli-for-plugins": "true",
				},
			},
			plugin:  "global",
			key:     "context-aware-cli-for-plugins",
			value:   false,
			persist: true,
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

			persist, err := SetFeature(tc.plugin, tc.key, strconv.FormatBool(tc.value))
			assert.NoError(t, err)
			assert.Equal(t, tc.persist, persist)

			ok, err := IsFeatureEnabled(tc.plugin, tc.key)
			assert.NoError(t, err)
			assert.Equal(t, ok, tc.value)

			err = DeleteFeature(tc.plugin, tc.key)
			assert.NoError(t, err)

			ok, err = IsFeatureEnabled(tc.plugin, tc.key)
			assert.Equal(t, ok, tc.value)

			persist, err = SetFeature(tc.plugin, tc.key, strconv.FormatBool(tc.value))
			assert.NoError(t, err)
			assert.Equal(t, tc.persist, persist)

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
		cfg     *v1alpha1.ClientConfig
		plugin  string
		key     string
		value   bool
		persist bool
	}{
		{
			name: "success context-aware-cli-for-plugins",
			cfg: &v1alpha1.ClientConfig{
				ClientOptions: &v1alpha1.ClientOptions{
					Features: map[string]v1alpha1.FeatureMap{
						"global": {
							"context-aware-cli-for-plugins": "true",
						},
					},
				},
			},
			plugin:  "global",
			key:     "context-aware-cli-for-plugins",
			value:   false,
			persist: true,
		},
		{
			name: "success context-aware-cli-for-plugins",
			cfg: &v1alpha1.ClientConfig{
				ClientOptions: &v1alpha1.ClientOptions{
					Features: map[string]v1alpha1.FeatureMap{
						"global": {
							"context-aware-cli-for-plugins": "true",
						},
					},
				},
			},
			plugin:  "global",
			key:     "context-aware-cli-for-plugins",
			value:   false,
			persist: true,
		},
		{
			name: "should not update the same feature value",
			cfg: &v1alpha1.ClientConfig{
				ClientOptions: &v1alpha1.ClientOptions{
					Features: map[string]v1alpha1.FeatureMap{
						"global": {
							"context-aware-cli-for-plugins": "true",
						},
					},
				},
			},
			plugin:  "global",
			key:     "context-aware-cli-for-plugins",
			value:   true,
			persist: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			err := StoreClientConfig(tc.cfg)
			assert.NoError(t, err)

			persist, err := SetFeature(tc.plugin, tc.key, strconv.FormatBool(tc.value))
			assert.NoError(t, err)
			assert.Equal(t, tc.persist, persist)

			ok, err := IsFeatureEnabled(tc.plugin, tc.key)
			assert.NoError(t, err)
			assert.Equal(t, ok, tc.value)
		})
	}

}

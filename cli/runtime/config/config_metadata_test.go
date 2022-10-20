package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
)

func TestGetConfigMetadata(t *testing.T) {

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
		out    string
		errStr string
	}{
		{
			name: "success k8s",
			in: &configapi.ClientConfig{
				ConfigMetadata: &configapi.ConfigMetadata{
					PatchStrategy: map[string]string{
						"contexts.clusterOpts":                      "replace",
						"contexts.discoverySources.gcp.annotations": "replace",
						"contexts.globalOpts.auth":                  "replace",
					},
				},
			},
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := SetConfigMetadataPatchStrategies(spec.in.ConfigMetadata.PatchStrategy)
			if err != nil {
				fmt.Printf("StoreClientConfigV2 errors: %v\n", err)
			}

			c, err := GetConfigMetadata()
			if err != nil {
				fmt.Printf("errors: %v\n", err)
			}
			assert.Equal(t, c, spec.in.ConfigMetadata)
			assert.NoError(t, err)

		})
	}
}

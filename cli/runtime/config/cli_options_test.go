// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
)

func TestGetEdition(t *testing.T) {

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
				ClientOptions: &configapi.ClientOptions{
					Env: map[string]string{
						"test": "test",
					},
				},
			},
			out: "test",
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := StoreClientConfig(spec.in)
			if err != nil {
				fmt.Printf("StoreClientConfigV2 errors: %v\n", err)
			}

			c, err := GetEnv("test")
			if err != nil {
				fmt.Printf("errors: %v\n", err)
			}

			assert.Equal(t, spec.out, c)
			assert.NoError(t, err)

		})
	}
}

func TestSetEdition(t *testing.T) {

	// setup
	func() {
		LocalDirName = fmt.Sprintf(".tanzu-test")
	}()

	defer func() {
		cleanupDir(LocalDirName)
	}()

	tests := []struct {
		name   string
		src    *configapi.ClientConfig
		out    string
		errStr string
	}{
		{
			name: "success tanzu",
			src: &configapi.ClientConfig{
				ClientOptions: &configapi.ClientOptions{
					CLI: &configapi.CLIOptions{},
				},
			},
			out: "tanzu",
		},
		{
			name: "success tanu no edition",
			src: &configapi.ClientConfig{
				ClientOptions: &configapi.ClientOptions{},
			},
			out: "tanzu",
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := StoreClientConfig(spec.src)
			if err != nil {
				fmt.Printf("StoreClientConfigV2 errors: %v\n", err)
			}
			err = SetEdition(spec.out)
			if err != nil {
				fmt.Printf("errors: %v\n", err)
			}

			c, err := GetEdition()

			assert.Equal(t, spec.out, c)
			assert.NoError(t, err)

		})
	}
}

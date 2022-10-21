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
		name    string
		src     *configapi.ClientConfig
		in      string
		out     string
		persist bool
		errStr  string
	}{
		{
			name: "should persist tanzu when empty edition",
			src: &configapi.ClientConfig{
				ClientOptions: &configapi.ClientOptions{
					CLI: &configapi.CLIOptions{},
				},
			},
			in:      "tanzu",
			out:     "tanzu",
			persist: true,
		},
		{
			name:    "should persist tanzu when empty client config",
			src:     &configapi.ClientConfig{},
			in:      "tanzu",
			out:     "tanzu",
			persist: true,
		},
		{
			name: "should update and persist tanzu ",
			src: &configapi.ClientConfig{ClientOptions: &configapi.ClientOptions{
				CLI: &configapi.CLIOptions{
					Edition: "old-tanzu",
				},
			}},
			in:      "tanzu",
			out:     "tanzu",
			persist: true,
		},
		{
			name: "should not persist tanzu ",
			src: &configapi.ClientConfig{ClientOptions: &configapi.ClientOptions{
				CLI: &configapi.CLIOptions{
					Edition: "tanzu",
				},
			}},
			in:      "tanzu",
			out:     "tanzu",
			persist: false,
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := StoreClientConfig(spec.src)
			if err != nil {
				fmt.Printf("StoreClientConfigV2 errors: %v\n", err)
			}
			persist, err := SetEdition(spec.in)
			if err != nil {
				fmt.Printf("errors: %v\n", err)
			}
			assert.Equal(t, spec.persist, persist)

			c, err := GetEdition()

			assert.Equal(t, spec.out, c)
			assert.NoError(t, err)

		})
	}
}

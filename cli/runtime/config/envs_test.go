// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
)

func TestGetAllEnvs(t *testing.T) {

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
		out    map[string]string
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
			out: map[string]string{
				"test": "test",
			},
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := StoreClientConfig(spec.in)
			if err != nil {
				fmt.Printf("StoreClientConfigV2 errors: %v\n", err)
			}

			c, err := GetAllEnvs()
			if err != nil {
				fmt.Printf("errors: %v\n", err)
			}

			assert.Equal(t, spec.out, c)
			assert.NoError(t, err)

		})
	}
}

func TestGetEnv(t *testing.T) {

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

func TestSetEnv(t *testing.T) {

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
			name: "success k8s",
			src: &configapi.ClientConfig{
				ClientOptions: &configapi.ClientOptions{
					Env: map[string]string{
						"test": "test",
					},
				},
			},
			out: "test2",
		},
		{
			name: "success k8s",
			src: &configapi.ClientConfig{
				ClientOptions: &configapi.ClientOptions{},
			},
			out: "test2",
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := StoreClientConfig(spec.src)
			if err != nil {
				fmt.Printf("StoreClientConfigV2 errors: %v\n", err)
			}
			err = SetEnv(spec.out, spec.out)
			if err != nil {
				fmt.Printf("errors: %v\n", err)
			}

			c, err := GetEnv(spec.out)

			assert.Equal(t, spec.out, c)
			assert.NoError(t, err)

		})
	}
}
func TestDeleteEnv(t *testing.T) {
	// setup
	func() {
		LocalDirName = fmt.Sprintf(".tanzu-test")
		cfg := &configapi.ClientConfig{
			ClientOptions: &configapi.ClientOptions{
				Env: map[string]string{
					"test":  "test",
					"test2": "test2",
					"test3": "test2",
					"test4": "test2",
				},
			},
		}
		err := StoreClientConfig(cfg)
		assert.NoError(t, err)
	}()

	defer func() {
		cleanupDir(LocalDirName)
	}()

	tests := []struct {
		name string
		in   string
		out  bool
	}{
		{
			name: "success delete test",
			in:   "test",
			out:  true,
		},
		{
			name: "success delete test2",
			in:   "test2",
			out:  true,
		},

		{
			name: "success delete test3",
			in:   "test3",
			out:  true,
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {

			err := DeleteEnv(spec.in)
			if err != nil {
				fmt.Printf("errors: %v\n", err)
			}

			c, err := GetEnv(spec.in)
			assert.Equal(t, spec.out, c == "")
		})
	}
}

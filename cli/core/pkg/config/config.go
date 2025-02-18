// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package config provides functions for the tanzu cli configuration
package config

import (
	"os"

	"github.com/vmware-tanzu/tanzu-framework/cli/runtime/config"
)

// ConfigureEnvVariables reads and configures provided environment variables
// as part of tanzu configuration file
func ConfigureEnvVariables() {
	envMap := config.GetEnvConfigurations()
	if envMap == nil {
		return
	}
	for variable, value := range envMap {
		// If environment variable is not already set
		// set the environment variable
		if os.Getenv(variable) == "" {
			os.Setenv(variable, value)
		}
	}
}

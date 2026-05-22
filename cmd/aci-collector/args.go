// SPDX-License-Identifier: Apache-2.0

// Copyright 2026 Cisco Systems, Inc. and their affiliates

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/brightpuddle/aci-collector/pkg/config"

	"github.com/alexflint/go-arg"
)

const resultZip = "aci-vetr-data.zip"

var version = "(dev)"

// Args are command line parameters.
type Args struct {
	URL               string            `arg:"--url,env:ACI_URL"           help:"APIC hostname or IP address"`
	Username          string            `arg:"--username,env:ACI_USERNAME" help:"APIC username"`
	Password          string            `arg:"--password,env:ACI_PASSWORD" help:"APIC password"`
	Output            string            `arg:"-o"                          help:"Output file"`
	ConfigFile        string            `arg:"-c,--config"                 help:"Path to YAML configuration file"`
	RequestRetryCount int               `arg:"--request-retry-count"       help:"Times to retry a failed request"       default:"3"`
	RetryDelay        int               `arg:"--retry-delay"               help:"Seconds to wait before retry"          default:"10"`
	BatchSize         int               `arg:"--batch-size"                help:"Max request to send in parallel"       default:"7"`
	PageSize          int               `arg:"--page-size"                 help:"Object per page for large datasets"    default:"1000"`
	SSHPort           int               `arg:"--ssh-port"                  help:"SSH port for CLI collection"           default:"22"`
	Confirm           bool              `arg:"-y"                          help:"Skip confirmation"`
	Verbose           bool              `arg:"-v,--verbose"                help:"Enable verbose (debug level) logging"`
	Class             string            `arg:"--class"                     help:"Collect a single class"                default:"all"`
	Query             map[string]string `arg:"-q"                          help:"Query(s) to filter single class query"`
}

// Description is the CLI description string.
func (Args) Description() string {
	return "ACI vetR collector"
}

// Version is the CLI version string.
func (Args) Version() string {
	return version
}

// readArgs collects the CLI args and returns a config.Config.
func readArgs() (*config.Config, error) {
	args := Args{Output: resultZip}
	arg.MustParse(&args)

	if args.ConfigFile != "" {
		cfg, err := config.ParseConfig(args.ConfigFile)
		if err != nil {
			return nil, err
		}
		if err := cfg.NormalizeAndPrompt(); err != nil {
			return nil, err
		}
		return cfg, nil
	}

	cfg := config.New()
	requestRetryCount := args.RequestRetryCount
	retryDelay := args.RetryDelay
	batchSize := args.BatchSize
	pageSize := args.PageSize
	sshPort := args.SSHPort
	confirm := args.Confirm
	verbose := args.Verbose

	cfg.Global.Verbose = args.Verbose
	cfg.Fabrics = []config.FabricConfig{{
		URL:               args.URL,
		Output:            args.Output,
		Username:          args.Username,
		Password:          args.Password,
		RequestRetryCount: &requestRetryCount,
		RetryDelay:        &retryDelay,
		BatchSize:         &batchSize,
		PageSize:          &pageSize,
		SSHPort:           &sshPort,
		Confirm:           &confirm,
		Verbose:           &verbose,
		Class:             args.Class,
		Query:             args.Query,
	}}

	if err := cfg.NormalizeAndPrompt(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Package ssh provides SSH-based CLI data collection from APIC nodes.

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
package ssh

import (
	"bytes"
	"fmt"
	"net"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/brightpuddle/aci-collector/pkg/archive"
	"github.com/brightpuddle/aci-collector/pkg/config"
	"github.com/brightpuddle/aci-collector/pkg/log"
	"github.com/brightpuddle/aci-collector/pkg/req"
)

// parseHost extracts the hostname or IP from a URL string, stripping any
// scheme prefix and port number.
func parseHost(rawURL string) string {
	rawURL, _ = strings.CutPrefix(rawURL, "https://")
	rawURL, _ = strings.CutPrefix(rawURL, "http://")
	rawURL = strings.TrimRight(rawURL, "/")
	host, _, err := net.SplitHostPort(rawURL)
	if err != nil {
		// No port present; use as-is.
		return rawURL
	}
	return host
}

// CollectCLI connects to the APIC via SSH and executes each CLIRequest,
// writing the output to the provided archive.
// The same username/password used for REST API collection is reused.
// Host key verification is intentionally skipped, matching the TLS behaviour
// of the REST API client.
func CollectCLI(cfg config.FabricConfig, arc archive.Writer) error {
	logger := getLogger(cfg)

	host := parseHost(cfg.URL)
	addr := net.JoinHostPort(host, fmt.Sprintf("%d", cfg.GetSSHPort()))

	sshCfg := &ssh.ClientConfig{
		User: cfg.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(cfg.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //nolint:gosec
		Timeout:         30 * time.Second,
	}

	logger.Info().Str("host", addr).Msg("Connecting via SSH for CLI collection...")
	client, err := ssh.Dial("tcp", addr, sshCfg)
	if err != nil {
		return fmt.Errorf("SSH connect to %s failed: %w", addr, err)
	}
	defer client.Close()

	for _, r := range req.CLIRequests {
		if err := runCommand(client, r, arc, logger); err != nil {
			logger.Warn().Err(err).Msgf("CLI command failed: %s", r.Command)
		}
	}
	return nil
}

func runCommand(client *ssh.Client, r req.CLIRequest, arc archive.Writer, logger log.Logger) error {
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("create SSH session: %w", err)
	}
	defer session.Close()

	var out bytes.Buffer
	session.Stdout = &out

	logger.Debug().Msgf("Running CLI command: %s", r.Command)
	if err := session.Run(r.Command); err != nil {
		return fmt.Errorf("run %q: %w", r.Command, err)
	}

	logger.Info().Msgf("CLI command complete: %s -> %s", r.Command, r.Filename)
	return arc.Add(r.Filename, out.Bytes())
}

func getLogger(cfg config.FabricConfig) log.Logger {
	if cfg.GetFabricName() != "" {
		return log.WithFabric(cfg.GetFabricName())
	}
	return log.New()
}

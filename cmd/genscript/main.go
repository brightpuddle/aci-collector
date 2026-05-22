// Command genscript generates the aci-collector.sh shell script
// from the embedded request data.

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
	"fmt"
	"os"
	"path/filepath"

	"github.com/brightpuddle/aci-collector/pkg/req"
)

const (
	tmpFolder = "/tmp/aci-collector"
)

func main() {
	// Determine output path - should be at repo root
	// When run via go generate from pkg/req, we need to go up two directories
	scriptPath := "aci-collector.sh"
	if _, err := os.Stat("../../go.mod"); err == nil {
		// We're in a subdirectory (e.g., pkg/req), write to repo root
		scriptPath = "../../aci-collector.sh"
	}

	// Get absolute path
	absPath, err := filepath.Abs(scriptPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error resolving path: %v\n", err)
		os.Exit(1)
	}

	// Open output file
	f, err := os.Create(absPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	// Write script header
	fmt.Fprintln(f, "#!/bin/bash")
	fmt.Fprintln(f, "")
	fmt.Fprintf(f, "rm -rf %s > /dev/null\n", tmpFolder)
	fmt.Fprintf(f, "mkdir %s\n", tmpFolder)
	fmt.Fprintln(f, "")
	fmt.Fprintln(f, "# Fetch data from the API")

	// Get requests
	reqs, err := req.GetRequests()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting requests: %v\n", err)
		os.Exit(1)
	}

	// Write icurl commands for each request
	for _, r := range reqs {
		cmd := fmt.Sprintf("icurl -kG https://localhost/api/class/%s.json", r.Class)

		// Add query parameters if present
		for k, v := range r.Query {
			cmd += fmt.Sprintf(" -d '%s=%s'", k, v)
		}

		// Add output redirection
		cmd += fmt.Sprintf(" > %s/%s.json", tmpFolder, r.Class)

		fmt.Fprintln(f, cmd)
	}

	// Write CLI commands (run locally on the APIC, no SSH needed)
	fmt.Fprintln(f, "")
	fmt.Fprintln(f, "# Collect CLI output")
	for _, r := range req.CLIRequests {
		fmt.Fprintf(f, "%s > %s/%s\n", r.Command, tmpFolder, r.Filename)
	}

	// Write script footer
	fmt.Fprintln(f, "")
	fmt.Fprintln(f, "# Zip result")
	fmt.Fprintf(f, "zip -mj ~/aci-vetr-data.zip %s/*.json %s/*.txt\n", tmpFolder, tmpFolder)
	fmt.Fprintln(f, "")
	fmt.Fprintln(f, "# Cleanup")
	fmt.Fprintln(f, "")
	fmt.Fprintf(f, "rm -rf %s\n", tmpFolder)
	fmt.Fprintln(f, "")
	fmt.Fprintln(f, "echo Collection complete")
	fmt.Fprintln(f, "echo Output writen to ~/aci-vetr-data.zip, i.e. user home folder")
	fmt.Fprintln(f, "echo Please provide aci-vetr-data.zip to Cisco for analysis.")

	// Make the script executable
	if err := os.Chmod(absPath, 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "Error making script executable: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s successfully\n", absPath)
}

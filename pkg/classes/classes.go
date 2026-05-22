// Package classes exports the list of ACI object classes collected by aci-collector.
// This is the canonical class list shared between the collector and analysis tools.
//
// To regenerate after modifying pkg/req/requests.go, run:
//
//	cd pkg/req && go generate

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
package classes

import "github.com/brightpuddle/aci-collector/pkg/req"

// Classes is the list of all ACI object classes queried by the collector.
// It is derived from pkg/req/requests.go and used by the analysis tool
// to validate that all DB queries reference classes that are actually collected.
var Classes = func() []string {
	reqs, _ := req.GetRequests()
	classes := make([]string, 0, len(reqs))
	for _, r := range reqs {
		classes = append(classes, r.Class)
	}
	return classes
}()

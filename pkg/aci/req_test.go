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

package aci

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

// TestSetRaw tests the Body::SetRaw method.
func TestSetRaw(t *testing.T) {
	name := Body{}.SetRaw("a", `{"name":"a"}`).Res().Get("a.name").Str
	assert.Equal(t, "a", name)
}

// TestQuery tests the Query function.
func TestQuery(t *testing.T) {
	defer gock.Off()
	client := testClient()

	gock.New(testURL).Get("/url").MatchParam("foo", "bar").Reply(200)
	_, err := client.Get("/url", Query("foo", "bar"))
	assert.NoError(t, err)

	// Test case for comma-separated parameters
	gock.New(testURL).Get("/url").MatchParam("foo", "bar,baz").Reply(200)
	_, err = client.Get("/url", Query("foo", "bar,baz"))
	assert.NoError(t, err)
}

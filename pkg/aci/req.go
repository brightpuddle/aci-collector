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
	"net/http"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// Body wraps SJSON for building JSON body strings.
// Usage example:
//
//	Body{}.Set("fvTenant.attributes.name", "mytenant").Str
type Body struct {
	Str string
}

// Set sets a JSON path to a value.
func (body Body) Set(path, value string) Body {
	res, _ := sjson.Set(body.Str, path, value)
	body.Str = res
	return body
}

// SetRaw sets a JSON path to a raw string value.
// This is primarily used for building up nested structures, e.g.:
//
//	Body{}.SetRaw("fvTenant.attributes", Body{}.Set("name", "mytenant").Str).Str
func (body Body) SetRaw(path, rawValue string) Body {
	res, _ := sjson.SetRaw(body.Str, path, rawValue)
	body.Str = res
	return body
}

// Res creates a Res object, i.e. a GJSON result object.
func (body Body) Res() Res {
	return gjson.Parse(body.Str)
}

// Req wraps http.Request for API requests.
type Req struct {
	// HTTPReq is the *http.Request obejct.
	HTTPReq *http.Request
	// Refresh indicates whether token refresh should be checked for this request.
	// Pass NoRefresh to disable Refresh check.
	Refresh bool
}

// NoRefresh prevents token refresh check.
// Primarily used by the Login and Refresh methods where this would be redundant.
func NoRefresh(req *Req) {
	req.Refresh = false
}

// Query sets an HTTP query parameter.
//
//	client.GetClass("fvBD", aci.Query("query-target-filter", `eq(fvBD.name,"bd-name")`))
//
// Or set multiple parameters:
//
//	client.GetClass("fvBD",
//	  aci.Query("rsp-subtree-include", "faults"),
//	  aci.Query("query-target-filter", `eq(fvBD.name,"bd-name")`))
func Query(k, v string) func(req *Req) {
	return func(req *Req) {
		q := req.HTTPReq.URL.Query()
		q.Add(k, v)
		req.HTTPReq.URL.RawQuery = q.Encode()
	}
}

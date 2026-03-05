// Package classes exports the list of ACI object classes collected by aci-collector.
// This is the canonical class list shared between the collector and analysis tools.
//
// To regenerate after modifying pkg/req/requests.go, run:
//
//	cd pkg/req && go generate
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

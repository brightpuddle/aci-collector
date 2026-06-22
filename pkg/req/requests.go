// Package requests defines the canonical list of ACI object classes and API
// queries collected by aci-collector. This file is the single source of truth;
// CI automatically propagates changes to brightpuddle/aci-collector via the
// sync-collector workflow.

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

package requests

// Request is an ACI REST API request.
type Request struct {
	Class string            // MO class
	Query map[string]string // Query parameters
}

// Requests contains all ACI API requests executed by the collector.
var Requests = []Request{
	{Class: "faultInst"},
	{Class: "eqptFlash"},
	{Class: "topSystem"},
	{Class: "isisDomPol"},
	{Class: "fabricSetupP"},
	{Class: "eqptStorage"},
	{Class: "aaaLoginDomain"},
	{Class: "aaaDomainAuth"},
	{Class: "pkiExportEncryptionKey"},
	{Class: "fvCtx"},
	{Class: "fabricNode"},
	{Class: "fabricLink"},
	{Class: "infraRsVlanNs"},
	{Class: "l2extDomP"},
	{Class: "l3extDomP"},
	{Class: "physDomP"},
	{Class: "l2extRsL2DomAtt"},
	{Class: "l3extRsL3DomAtt"},
	{Class: "fvRsDomAtt"},
	{Class: "maintMaintGrp"},
	{Class: "maintRsMgrpp"},
	{Class: "maintMaintP"},
	{Class: "fabricNodeBlk"},
	{Class: "firmwareRunning"},
	{Class: "firmwareCtrlrRunning"},
	{Class: "fabricHealthTotal"},
	{
		Class: "healthInst",
		Query: map[string]string{
			"query-target-filter": "wcard(healthInst.dn,\"^sys/health$\")",
		},
	},
	{Class: "infraSetPol"},
	{Class: "fvRsPathAtt"},
	{Class: "fvTenant"},
	{Class: "fvBD"},
	{Class: "vzBrCP"},
	{Class: "fvAEPg"},
	{Class: "l3extOut"},
	{Class: "epLoopProtectP"},
	{Class: "bgpRRNodePEp"},
	{Class: "eqptcapacityFSPartition"},
	{Class: "fvSubnet"},
	{Class: "fvRsBd"},
	{Class: "l3extRsPathL3OutAtt"},
	{Class: "ipv4Addr"},
	{Class: "ipv6Addr"},
	{Class: "eqptExtCh"},
	{Class: "coopPol"},
	{Class: "mcpInstPol"},
	{Class: "l3extLNodeP"},
	{Class: "l3extRsNodeL3OutAtt"},
	{Class: "fvnsVlanInstP"},
	{Class: "fvnsEncapBlk"},
	{Class: "infraAttEntityP"},
	{Class: "infraRsDomP"},
	{Class: "infraRsFuncToEpg"},
	{Class: "infraPortTrackPol"},
	{Class: "fabricRsTimePol"},
	{Class: "datetimePol"},
	{Class: "datetimeNtpProv"},
	{Class: "datetimeFormat"},
	{Class: "fcDomP"},
	{Class: "vmmDomP"},
	{Class: "infraRsAttEntP"},
	{Class: "apPlugin"},
	{Class: "l3IfPol"},
	{Class: "fabricExplicitGEp"},
	{Class: "fabricNodePEp"},
	{Class: "epControlP"},
	{Class: "fvRsCtx"},
	{Class: "fabricRsLeNodePGrp"},
	{Class: "fabricRsNodeCtrl"},
	{Class: "fabricNodeControl"},
	{Class: "fabricRsSpNodePGrp"},
	{Class: "configRsRemotePath"},
	{Class: "configRsExportScheduler"},
	{Class: "fvcapRule"},
	{
		Class: "fvCEp",
		Query: map[string]string{
			"rsp-subtree-include": "count",
		},
	},
	{Class: "vzFilter"},
	{Class: "vzRsSubjFiltAtt"},
	{Class: "fabricCtrlrConfigP"},
	{Class: "l3extInstP"},
	{Class: "l3extSubnet"},
	{Class: "fvRsCons"},
	{Class: "fvRsProv"},
	{Class: "fvAp"},
	{Class: "infraAccPortGrp"},
	{Class: "infraAccBndlGrp"},
	{Class: "infraAccPortP"},
	{Class: "infraNodeP"},
	{Class: "infraLeafS"},
	{Class: "mcpIfPol"},
	{Class: "infraRsMcpIfPol"},
	{Class: "infraRsAccBaseGrp"},
	{Class: "infraRsAccPortP"},
	{
		Class: "ctxClassCnt",
		Query: map[string]string{
			"rsp-subtree-class": "l2BD,fvEpP,l3Dom",
		},
	},
	{Class: "eqptcapacityVlanUsage5min"},
	{Class: "eqptcapacityL2Usage5min"},
	{Class: "eqptcapacityL2RemoteUsage5min"},
	{Class: "eqptcapacityL2TotalUsage5min"},
	{Class: "eqptcapacityL3Usage5min"},
	{Class: "eqptcapacityL3RemoteUsage5min"},
	{Class: "eqptcapacityL3TotalUsage5min"},
	{Class: "eqptcapacityL3TotalUsageCap5min"},
	{Class: "eqptcapacityPolUsage5min"},
	{Class: "infraWiNode"},
	{Class: "epIpAgingP"},
	{Class: "eqptFt"},
	{Class: "eqptFC"},
	{Class: "eqptSupC"},
	{Class: "eqptPsu"},
	{Class: "eqptLC"},
	{Class: "eqptSysC"},
	{Class: "cdpAdjEp"},
	{Class: "lldpAdjEp"},
	{Class: "infraHIfPol"},
	{Class: "mgmtStaticRoute"},
	{Class: "qosInstPol"},
	{Class: "licenseEntitlement"},
}

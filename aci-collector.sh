#!/bin/bash

rm -rf /tmp/aci-collector > /dev/null
mkdir /tmp/aci-collector

# Fetch data from the API
icurl -kG https://localhost/api/class/faultInst.json > /tmp/aci-collector/faultInst.json
icurl -kG https://localhost/api/class/eqptFlash.json > /tmp/aci-collector/eqptFlash.json
icurl -kG https://localhost/api/class/topSystem.json > /tmp/aci-collector/topSystem.json
icurl -kG https://localhost/api/class/isisDomPol.json > /tmp/aci-collector/isisDomPol.json
icurl -kG https://localhost/api/class/fabricSetupP.json > /tmp/aci-collector/fabricSetupP.json
icurl -kG https://localhost/api/class/eqptStorage.json > /tmp/aci-collector/eqptStorage.json
icurl -kG https://localhost/api/class/pkiExportEncryptionKey.json > /tmp/aci-collector/pkiExportEncryptionKey.json
icurl -kG https://localhost/api/class/fvCtx.json > /tmp/aci-collector/fvCtx.json
icurl -kG https://localhost/api/class/fabricNode.json > /tmp/aci-collector/fabricNode.json
icurl -kG https://localhost/api/class/fabricLink.json > /tmp/aci-collector/fabricLink.json
icurl -kG https://localhost/api/class/infraRsVlanNs.json > /tmp/aci-collector/infraRsVlanNs.json
icurl -kG https://localhost/api/class/l2extDomP.json > /tmp/aci-collector/l2extDomP.json
icurl -kG https://localhost/api/class/l3extDomP.json > /tmp/aci-collector/l3extDomP.json
icurl -kG https://localhost/api/class/physDomP.json > /tmp/aci-collector/physDomP.json
icurl -kG https://localhost/api/class/l2extRsL2DomAtt.json > /tmp/aci-collector/l2extRsL2DomAtt.json
icurl -kG https://localhost/api/class/l3extRsL3DomAtt.json > /tmp/aci-collector/l3extRsL3DomAtt.json
icurl -kG https://localhost/api/class/fvRsDomAtt.json > /tmp/aci-collector/fvRsDomAtt.json
icurl -kG https://localhost/api/class/maintMaintGrp.json > /tmp/aci-collector/maintMaintGrp.json
icurl -kG https://localhost/api/class/maintRsMgrpp.json > /tmp/aci-collector/maintRsMgrpp.json
icurl -kG https://localhost/api/class/maintMaintP.json > /tmp/aci-collector/maintMaintP.json
icurl -kG https://localhost/api/class/fabricNodeBlk.json > /tmp/aci-collector/fabricNodeBlk.json
icurl -kG https://localhost/api/class/firmwareRunning.json > /tmp/aci-collector/firmwareRunning.json
icurl -kG https://localhost/api/class/firmwareCtrlrRunning.json > /tmp/aci-collector/firmwareCtrlrRunning.json
icurl -kG https://localhost/api/class/fabricHealthTotal.json > /tmp/aci-collector/fabricHealthTotal.json
icurl -kG https://localhost/api/class/healthInst.json -d 'query-target-filter=wcard(healthInst.dn,"^sys/health$")' > /tmp/aci-collector/healthInst.json
icurl -kG https://localhost/api/class/infraSetPol.json > /tmp/aci-collector/infraSetPol.json
icurl -kG https://localhost/api/class/fvRsPathAtt.json > /tmp/aci-collector/fvRsPathAtt.json
icurl -kG https://localhost/api/class/fvTenant.json > /tmp/aci-collector/fvTenant.json
icurl -kG https://localhost/api/class/fvBD.json > /tmp/aci-collector/fvBD.json
icurl -kG https://localhost/api/class/vzBrCP.json > /tmp/aci-collector/vzBrCP.json
icurl -kG https://localhost/api/class/fvAEPg.json > /tmp/aci-collector/fvAEPg.json
icurl -kG https://localhost/api/class/l3extOut.json > /tmp/aci-collector/l3extOut.json
icurl -kG https://localhost/api/class/epLoopProtectP.json > /tmp/aci-collector/epLoopProtectP.json
icurl -kG https://localhost/api/class/bgpRRNodePEp.json > /tmp/aci-collector/bgpRRNodePEp.json
icurl -kG https://localhost/api/class/eqptcapacityFSPartition.json > /tmp/aci-collector/eqptcapacityFSPartition.json
icurl -kG https://localhost/api/class/fvSubnet.json > /tmp/aci-collector/fvSubnet.json
icurl -kG https://localhost/api/class/fvRsBd.json > /tmp/aci-collector/fvRsBd.json
icurl -kG https://localhost/api/class/l3extRsPathL3OutAtt.json > /tmp/aci-collector/l3extRsPathL3OutAtt.json
icurl -kG https://localhost/api/class/ipv4Addr.json > /tmp/aci-collector/ipv4Addr.json
icurl -kG https://localhost/api/class/ipv6Addr.json > /tmp/aci-collector/ipv6Addr.json
icurl -kG https://localhost/api/class/eqptExtCh.json > /tmp/aci-collector/eqptExtCh.json
icurl -kG https://localhost/api/class/coopPol.json > /tmp/aci-collector/coopPol.json
icurl -kG https://localhost/api/class/mcpInstPol.json > /tmp/aci-collector/mcpInstPol.json
icurl -kG https://localhost/api/class/l3extLNodeP.json > /tmp/aci-collector/l3extLNodeP.json
icurl -kG https://localhost/api/class/l3extRsNodeL3OutAtt.json > /tmp/aci-collector/l3extRsNodeL3OutAtt.json
icurl -kG https://localhost/api/class/fvnsVlanInstP.json > /tmp/aci-collector/fvnsVlanInstP.json
icurl -kG https://localhost/api/class/fvnsEncapBlk.json > /tmp/aci-collector/fvnsEncapBlk.json
icurl -kG https://localhost/api/class/infraAttEntityP.json > /tmp/aci-collector/infraAttEntityP.json
icurl -kG https://localhost/api/class/infraRsDomP.json > /tmp/aci-collector/infraRsDomP.json
icurl -kG https://localhost/api/class/infraRsFuncToEpg.json > /tmp/aci-collector/infraRsFuncToEpg.json
icurl -kG https://localhost/api/class/infraPortTrackPol.json > /tmp/aci-collector/infraPortTrackPol.json
icurl -kG https://localhost/api/class/fabricRsTimePol.json > /tmp/aci-collector/fabricRsTimePol.json
icurl -kG https://localhost/api/class/datetimePol.json > /tmp/aci-collector/datetimePol.json
icurl -kG https://localhost/api/class/datetimeNtpProv.json > /tmp/aci-collector/datetimeNtpProv.json
icurl -kG https://localhost/api/class/fcDomP.json > /tmp/aci-collector/fcDomP.json
icurl -kG https://localhost/api/class/vmmDomP.json > /tmp/aci-collector/vmmDomP.json
icurl -kG https://localhost/api/class/infraRsAttEntP.json > /tmp/aci-collector/infraRsAttEntP.json
icurl -kG https://localhost/api/class/apPlugin.json > /tmp/aci-collector/apPlugin.json
icurl -kG https://localhost/api/class/l3IfPol.json > /tmp/aci-collector/l3IfPol.json
icurl -kG https://localhost/api/class/fabricExplicitGEp.json > /tmp/aci-collector/fabricExplicitGEp.json
icurl -kG https://localhost/api/class/fabricNodePEp.json > /tmp/aci-collector/fabricNodePEp.json
icurl -kG https://localhost/api/class/epControlP.json > /tmp/aci-collector/epControlP.json
icurl -kG https://localhost/api/class/fvRsCtx.json > /tmp/aci-collector/fvRsCtx.json
icurl -kG https://localhost/api/class/fabricRsLeNodePGrp.json > /tmp/aci-collector/fabricRsLeNodePGrp.json
icurl -kG https://localhost/api/class/fabricRsNodeCtrl.json > /tmp/aci-collector/fabricRsNodeCtrl.json
icurl -kG https://localhost/api/class/fabricNodeControl.json > /tmp/aci-collector/fabricNodeControl.json
icurl -kG https://localhost/api/class/fabricRsSpNodePGrp.json > /tmp/aci-collector/fabricRsSpNodePGrp.json
icurl -kG https://localhost/api/class/configRsRemotePath.json > /tmp/aci-collector/configRsRemotePath.json
icurl -kG https://localhost/api/class/fvcapRule.json > /tmp/aci-collector/fvcapRule.json
icurl -kG https://localhost/api/class/fvCEp.json -d 'rsp-subtree-include=count' > /tmp/aci-collector/fvCEp.json
icurl -kG https://localhost/api/class/vzFilter.json > /tmp/aci-collector/vzFilter.json
icurl -kG https://localhost/api/class/fabricCtrlrConfigP.json > /tmp/aci-collector/fabricCtrlrConfigP.json
icurl -kG https://localhost/api/class/l3extInstP.json > /tmp/aci-collector/l3extInstP.json
icurl -kG https://localhost/api/class/fvRsCons.json > /tmp/aci-collector/fvRsCons.json
icurl -kG https://localhost/api/class/mcpIfPol.json > /tmp/aci-collector/mcpIfPol.json
icurl -kG https://localhost/api/class/infraRsMcpIfPol.json > /tmp/aci-collector/infraRsMcpIfPol.json
icurl -kG https://localhost/api/class/infraRsAccBaseGrp.json > /tmp/aci-collector/infraRsAccBaseGrp.json
icurl -kG https://localhost/api/class/infraRsAccPortP.json > /tmp/aci-collector/infraRsAccPortP.json
icurl -kG https://localhost/api/class/ctxClassCnt.json -d 'rsp-subtree-class=l2BD,fvEpP,l3Dom' > /tmp/aci-collector/ctxClassCnt.json
icurl -kG https://localhost/api/class/eqptcapacityVlanUsage5min.json > /tmp/aci-collector/eqptcapacityVlanUsage5min.json
icurl -kG https://localhost/api/class/eqptcapacityL2Usage5min.json > /tmp/aci-collector/eqptcapacityL2Usage5min.json
icurl -kG https://localhost/api/class/eqptcapacityL2RemoteUsage5min.json > /tmp/aci-collector/eqptcapacityL2RemoteUsage5min.json
icurl -kG https://localhost/api/class/eqptcapacityL2TotalUsage5min.json > /tmp/aci-collector/eqptcapacityL2TotalUsage5min.json
icurl -kG https://localhost/api/class/eqptcapacityL3Usage5min.json > /tmp/aci-collector/eqptcapacityL3Usage5min.json
icurl -kG https://localhost/api/class/eqptcapacityL3RemoteUsage5min.json > /tmp/aci-collector/eqptcapacityL3RemoteUsage5min.json
icurl -kG https://localhost/api/class/eqptcapacityL3TotalUsage5min.json > /tmp/aci-collector/eqptcapacityL3TotalUsage5min.json
icurl -kG https://localhost/api/class/eqptcapacityL3TotalUsageCap5min.json > /tmp/aci-collector/eqptcapacityL3TotalUsageCap5min.json
icurl -kG https://localhost/api/class/eqptcapacityPolUsage5min.json > /tmp/aci-collector/eqptcapacityPolUsage5min.json
icurl -kG https://localhost/api/class/infraWiNode.json > /tmp/aci-collector/infraWiNode.json
icurl -kG https://localhost/api/class/epIpAgingP.json > /tmp/aci-collector/epIpAgingP.json
icurl -kG https://localhost/api/class/eqptFt.json > /tmp/aci-collector/eqptFt.json
icurl -kG https://localhost/api/class/eqptFC.json > /tmp/aci-collector/eqptFC.json
icurl -kG https://localhost/api/class/eqptSupC.json > /tmp/aci-collector/eqptSupC.json
icurl -kG https://localhost/api/class/eqptPsu.json > /tmp/aci-collector/eqptPsu.json
icurl -kG https://localhost/api/class/eqptLC.json > /tmp/aci-collector/eqptLC.json
icurl -kG https://localhost/api/class/eqptSysC.json > /tmp/aci-collector/eqptSysC.json
icurl -kG https://localhost/api/class/cdpAdjEp.json > /tmp/aci-collector/cdpAdjEp.json
icurl -kG https://localhost/api/class/lldpAdjEp.json > /tmp/aci-collector/lldpAdjEp.json

# Zip result
zip -mj ~/aci-vetr-data.zip /tmp/aci-collector/*.json

# Cleanup

rm -rf /tmp/aci-collector

echo Collection complete
echo Output writen to ~/aci-vetr-data.zip, i.e. user home folder
echo Please provide aci-vetr-data.zip to Cisco for analysis.

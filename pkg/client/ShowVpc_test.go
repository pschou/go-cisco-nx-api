// Copyright 2018 Paul Greenberg (greenpau@outlook.com)
//            and Paul Schou     (github.com/pschou)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParseShowVpcJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowVpcResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.vpc",
			exp: &ShowVpcResponse{InsAPI: struct {
				Outputs struct {
					Output ShowVpcResponseResult "json:\"output\" xml:\"output\""
				} "json:\"outputs\" xml:\"outputs\""
				Sid     string "json:\"sid\" xml:\"sid\""
				Type    string "json:\"type\" xml:\"type\""
				Version string "json:\"version\" xml:\"version\""
			}{Outputs: struct {
				Output ShowVpcResponseResult "json:\"output\" xml:\"output\""
			}{Output: ShowVpcResponseResult{Body: ShowVpcResultBody{VpcDomainID: "100", VpcPeerStatus: "peer-ok", VpcPeerStatusReason: "SUCCESS", VpcPeerKeepaliveStatus: "peer-alive", VpcPeerConsistency: "consistent", VpcPerVlanPeerConsistency: "consistent", VpcPeerConsistencyStatus: "SUCCESS", VpcType2Consistency: "consistent", VpcType2ConsistencyStatus: "SUCCESS", VpcRole: "secondary", NumOfVpcs: 12, PeerGateway: "1", DualActiveExcludedVlans: "-", VpcGracefulConsistencyCheckStatus: "enabled", VpcAutoRecoveryStatus: "Enabled, timer is off.(timeout = 240s)", VpcDelayRestoreStatus: "Timer is off.(timeout = 30s)", VpcDelayRestoreSviStatus: "Timer is off.(timeout = 10s)", OperationalL3Peer: "Disabled", VpcPeerLinkHdr: "Start of VPC peer-link table", TablePeerlink: []struct {
				RowPeerlink []struct {
					PeerLinkID        string "json:\"peer-link-id\" xml:\"peer-link-id\""
					PeerlinkIfindex   string "json:\"peerlink-ifindex\" xml:\"peerlink-ifindex\""
					PeerLinkPortState string "json:\"peer-link-port-state\" xml:\"peer-link-port-state\""
					PeerUpVlanBitset  string "json:\"peer-up-vlan-bitset\" xml:\"peer-up-vlan-bitset\""
				} "json:\"ROW_peerlink\" xml:\"ROW_peerlink\""
			}{struct {
				RowPeerlink []struct {
					PeerLinkID        string "json:\"peer-link-id\" xml:\"peer-link-id\""
					PeerlinkIfindex   string "json:\"peerlink-ifindex\" xml:\"peerlink-ifindex\""
					PeerLinkPortState string "json:\"peer-link-port-state\" xml:\"peer-link-port-state\""
					PeerUpVlanBitset  string "json:\"peer-up-vlan-bitset\" xml:\"peer-up-vlan-bitset\""
				} "json:\"ROW_peerlink\" xml:\"ROW_peerlink\""
			}{RowPeerlink: []struct {
				PeerLinkID        string "json:\"peer-link-id\" xml:\"peer-link-id\""
				PeerlinkIfindex   string "json:\"peerlink-ifindex\" xml:\"peerlink-ifindex\""
				PeerLinkPortState string "json:\"peer-link-port-state\" xml:\"peer-link-port-state\""
				PeerUpVlanBitset  string "json:\"peer-up-vlan-bitset\" xml:\"peer-up-vlan-bitset\""
			}{struct {
				PeerLinkID        string "json:\"peer-link-id\" xml:\"peer-link-id\""
				PeerlinkIfindex   string "json:\"peerlink-ifindex\" xml:\"peerlink-ifindex\""
				PeerLinkPortState string "json:\"peer-link-port-state\" xml:\"peer-link-port-state\""
				PeerUpVlanBitset  string "json:\"peer-up-vlan-bitset\" xml:\"peer-up-vlan-bitset\""
			}{PeerLinkID: "1", PeerlinkIfindex: "Po100", PeerLinkPortState: "1", PeerUpVlanBitset: "1-501"}}}}, VpcEnd: []string{"End of table", "End of table"}, VpcHdr: "Start of vPC table", VpcNotEs: "vPC complex", TableVpc: []struct {
				RowVpc []struct {
					VpcID                int    "json:\"vpc-id\" xml:\"vpc-id\""
					VpcIfindex           string "json:\"vpc-ifindex\" xml:\"vpc-ifindex\""
					VpcPortState         string "json:\"vpc-port-state\" xml:\"vpc-port-state\""
					PhyPortIfRemoved     string "json:\"phy-port-if-removed\" xml:\"phy-port-if-removed\""
					VpcThruPeerlink      string "json:\"vpc-thru-peerlink\" xml:\"vpc-thru-peerlink\""
					VpcConsistency       string "json:\"vpc-consistency\" xml:\"vpc-consistency\""
					VpcConsistencyStatus string "json:\"vpc-consistency-status\" xml:\"vpc-consistency-status\""
					UpVlanBitset         string "json:\"up-vlan-bitset\" xml:\"up-vlan-bitset\""
					EsAttr               string "json:\"es-attr\" xml:\"es-attr\""
				} "json:\"ROW_vpc\" xml:\"ROW_vpc\""
			}{struct {
				RowVpc []struct {
					VpcID                int    "json:\"vpc-id\" xml:\"vpc-id\""
					VpcIfindex           string "json:\"vpc-ifindex\" xml:\"vpc-ifindex\""
					VpcPortState         string "json:\"vpc-port-state\" xml:\"vpc-port-state\""
					PhyPortIfRemoved     string "json:\"phy-port-if-removed\" xml:\"phy-port-if-removed\""
					VpcThruPeerlink      string "json:\"vpc-thru-peerlink\" xml:\"vpc-thru-peerlink\""
					VpcConsistency       string "json:\"vpc-consistency\" xml:\"vpc-consistency\""
					VpcConsistencyStatus string "json:\"vpc-consistency-status\" xml:\"vpc-consistency-status\""
					UpVlanBitset         string "json:\"up-vlan-bitset\" xml:\"up-vlan-bitset\""
					EsAttr               string "json:\"es-attr\" xml:\"es-attr\""
				} "json:\"ROW_vpc\" xml:\"ROW_vpc\""
			}{RowVpc: []struct {
				VpcID                int    "json:\"vpc-id\" xml:\"vpc-id\""
				VpcIfindex           string "json:\"vpc-ifindex\" xml:\"vpc-ifindex\""
				VpcPortState         string "json:\"vpc-port-state\" xml:\"vpc-port-state\""
				PhyPortIfRemoved     string "json:\"phy-port-if-removed\" xml:\"phy-port-if-removed\""
				VpcThruPeerlink      string "json:\"vpc-thru-peerlink\" xml:\"vpc-thru-peerlink\""
				VpcConsistency       string "json:\"vpc-consistency\" xml:\"vpc-consistency\""
				VpcConsistencyStatus string "json:\"vpc-consistency-status\" xml:\"vpc-consistency-status\""
				UpVlanBitset         string "json:\"up-vlan-bitset\" xml:\"up-vlan-bitset\""
				EsAttr               string "json:\"es-attr\" xml:\"es-attr\""
			}{struct {
				VpcID                int    "json:\"vpc-id\" xml:\"vpc-id\""
				VpcIfindex           string "json:\"vpc-ifindex\" xml:\"vpc-ifindex\""
				VpcPortState         string "json:\"vpc-port-state\" xml:\"vpc-port-state\""
				PhyPortIfRemoved     string "json:\"phy-port-if-removed\" xml:\"phy-port-if-removed\""
				VpcThruPeerlink      string "json:\"vpc-thru-peerlink\" xml:\"vpc-thru-peerlink\""
				VpcConsistency       string "json:\"vpc-consistency\" xml:\"vpc-consistency\""
				VpcConsistencyStatus string "json:\"vpc-consistency-status\" xml:\"vpc-consistency-status\""
				UpVlanBitset         string "json:\"up-vlan-bitset\" xml:\"up-vlan-bitset\""
				EsAttr               string "json:\"es-attr\" xml:\"es-attr\""
			}{VpcID: 1, VpcIfindex: "Po1", VpcPortState: "1", PhyPortIfRemoved: "disabled", VpcThruPeerlink: "0", VpcConsistency: "consistent", VpcConsistencyStatus: "SUCCESS", UpVlanBitset: "1-100,302-501", EsAttr: "DF: Invalid"}, struct {
				VpcID                int    "json:\"vpc-id\" xml:\"vpc-id\""
				VpcIfindex           string "json:\"vpc-ifindex\" xml:\"vpc-ifindex\""
				VpcPortState         string "json:\"vpc-port-state\" xml:\"vpc-port-state\""
				PhyPortIfRemoved     string "json:\"phy-port-if-removed\" xml:\"phy-port-if-removed\""
				VpcThruPeerlink      string "json:\"vpc-thru-peerlink\" xml:\"vpc-thru-peerlink\""
				VpcConsistency       string "json:\"vpc-consistency\" xml:\"vpc-consistency\""
				VpcConsistencyStatus string "json:\"vpc-consistency-status\" xml:\"vpc-consistency-status\""
				UpVlanBitset         string "json:\"up-vlan-bitset\" xml:\"up-vlan-bitset\""
				EsAttr               string "json:\"es-attr\" xml:\"es-attr\""
			}{VpcID: 2, VpcIfindex: "Po2", VpcPortState: "1", PhyPortIfRemoved: "disabled", VpcThruPeerlink: "0", VpcConsistency: "consistent", VpcConsistencyStatus: "SUCCESS", UpVlanBitset: "1-501", EsAttr: "DF: Invalid"}, struct {
				VpcID                int    "json:\"vpc-id\" xml:\"vpc-id\""
				VpcIfindex           string "json:\"vpc-ifindex\" xml:\"vpc-ifindex\""
				VpcPortState         string "json:\"vpc-port-state\" xml:\"vpc-port-state\""
				PhyPortIfRemoved     string "json:\"phy-port-if-removed\" xml:\"phy-port-if-removed\""
				VpcThruPeerlink      string "json:\"vpc-thru-peerlink\" xml:\"vpc-thru-peerlink\""
				VpcConsistency       string "json:\"vpc-consistency\" xml:\"vpc-consistency\""
				VpcConsistencyStatus string "json:\"vpc-consistency-status\" xml:\"vpc-consistency-status\""
				UpVlanBitset         string "json:\"up-vlan-bitset\" xml:\"up-vlan-bitset\""
				EsAttr               string "json:\"es-attr\" xml:\"es-attr\""
			}{VpcID: 3, VpcIfindex: "Po3", VpcPortState: "1", PhyPortIfRemoved: "disabled", VpcThruPeerlink: "0", VpcConsistency: "consistent", VpcConsistencyStatus: "SUCCESS", UpVlanBitset: "100", EsAttr: "DF: Invalid"}, struct {
				VpcID                int    "json:\"vpc-id\" xml:\"vpc-id\""
				VpcIfindex           string "json:\"vpc-ifindex\" xml:\"vpc-ifindex\""
				VpcPortState         string "json:\"vpc-port-state\" xml:\"vpc-port-state\""
				PhyPortIfRemoved     string "json:\"phy-port-if-removed\" xml:\"phy-port-if-removed\""
				VpcThruPeerlink      string "json:\"vpc-thru-peerlink\" xml:\"vpc-thru-peerlink\""
				VpcConsistency       string "json:\"vpc-consistency\" xml:\"vpc-consistency\""
				VpcConsistencyStatus string "json:\"vpc-consistency-status\" xml:\"vpc-consistency-status\""
				UpVlanBitset         string "json:\"up-vlan-bitset\" xml:\"up-vlan-bitset\""
				EsAttr               string "json:\"es-attr\" xml:\"es-attr\""
			}{VpcID: 4, VpcIfindex: "Po4", VpcPortState: "1", PhyPortIfRemoved: "disabled", VpcThruPeerlink: "0", VpcConsistency: "consistent", VpcConsistencyStatus: "SUCCESS", UpVlanBitset: "500", EsAttr: "DF: Invalid"}, struct {
				VpcID                int    "json:\"vpc-id\" xml:\"vpc-id\""
				VpcIfindex           string "json:\"vpc-ifindex\" xml:\"vpc-ifindex\""
				VpcPortState         string "json:\"vpc-port-state\" xml:\"vpc-port-state\""
				PhyPortIfRemoved     string "json:\"phy-port-if-removed\" xml:\"phy-port-if-removed\""
				VpcThruPeerlink      string "json:\"vpc-thru-peerlink\" xml:\"vpc-thru-peerlink\""
				VpcConsistency       string "json:\"vpc-consistency\" xml:\"vpc-consistency\""
				VpcConsistencyStatus string "json:\"vpc-consistency-status\" xml:\"vpc-consistency-status\""
				UpVlanBitset         string "json:\"up-vlan-bitset\" xml:\"up-vlan-bitset\""
				EsAttr               string "json:\"es-attr\" xml:\"es-attr\""
			}{VpcID: 5, VpcIfindex: "Po5", VpcPortState: "1", PhyPortIfRemoved: "disabled", VpcThruPeerlink: "0", VpcConsistency: "consistent", VpcConsistencyStatus: "SUCCESS", UpVlanBitset: "1-501", EsAttr: "DF: Invalid"}, struct {
				VpcID                int    "json:\"vpc-id\" xml:\"vpc-id\""
				VpcIfindex           string "json:\"vpc-ifindex\" xml:\"vpc-ifindex\""
				VpcPortState         string "json:\"vpc-port-state\" xml:\"vpc-port-state\""
				PhyPortIfRemoved     string "json:\"phy-port-if-removed\" xml:\"phy-port-if-removed\""
				VpcThruPeerlink      string "json:\"vpc-thru-peerlink\" xml:\"vpc-thru-peerlink\""
				VpcConsistency       string "json:\"vpc-consistency\" xml:\"vpc-consistency\""
				VpcConsistencyStatus string "json:\"vpc-consistency-status\" xml:\"vpc-consistency-status\""
				UpVlanBitset         string "json:\"up-vlan-bitset\" xml:\"up-vlan-bitset\""
				EsAttr               string "json:\"es-attr\" xml:\"es-attr\""
			}{VpcID: 6, VpcIfindex: "Po6", VpcPortState: "1", PhyPortIfRemoved: "disabled", VpcThruPeerlink: "0", VpcConsistency: "consistent", VpcConsistencyStatus: "SUCCESS", UpVlanBitset: "1-501", EsAttr: "DF: Invalid"}, struct {
				VpcID                int    "json:\"vpc-id\" xml:\"vpc-id\""
				VpcIfindex           string "json:\"vpc-ifindex\" xml:\"vpc-ifindex\""
				VpcPortState         string "json:\"vpc-port-state\" xml:\"vpc-port-state\""
				PhyPortIfRemoved     string "json:\"phy-port-if-removed\" xml:\"phy-port-if-removed\""
				VpcThruPeerlink      string "json:\"vpc-thru-peerlink\" xml:\"vpc-thru-peerlink\""
				VpcConsistency       string "json:\"vpc-consistency\" xml:\"vpc-consistency\""
				VpcConsistencyStatus string "json:\"vpc-consistency-status\" xml:\"vpc-consistency-status\""
				UpVlanBitset         string "json:\"up-vlan-bitset\" xml:\"up-vlan-bitset\""
				EsAttr               string "json:\"es-attr\" xml:\"es-attr\""
			}{VpcID: 7, VpcIfindex: "Po7", VpcPortState: "1", PhyPortIfRemoved: "disabled", VpcThruPeerlink: "0", VpcConsistency: "consistent", VpcConsistencyStatus: "SUCCESS", UpVlanBitset: "1-501", EsAttr: "DF: Invalid"}, struct {
				VpcID                int    "json:\"vpc-id\" xml:\"vpc-id\""
				VpcIfindex           string "json:\"vpc-ifindex\" xml:\"vpc-ifindex\""
				VpcPortState         string "json:\"vpc-port-state\" xml:\"vpc-port-state\""
				PhyPortIfRemoved     string "json:\"phy-port-if-removed\" xml:\"phy-port-if-removed\""
				VpcThruPeerlink      string "json:\"vpc-thru-peerlink\" xml:\"vpc-thru-peerlink\""
				VpcConsistency       string "json:\"vpc-consistency\" xml:\"vpc-consistency\""
				VpcConsistencyStatus string "json:\"vpc-consistency-status\" xml:\"vpc-consistency-status\""
				UpVlanBitset         string "json:\"up-vlan-bitset\" xml:\"up-vlan-bitset\""
				EsAttr               string "json:\"es-attr\" xml:\"es-attr\""
			}{VpcID: 8, VpcIfindex: "Po8", VpcPortState: "1", PhyPortIfRemoved: "disabled", VpcThruPeerlink: "0", VpcConsistency: "consistent", VpcConsistencyStatus: "SUCCESS", UpVlanBitset: "200", EsAttr: "DF: Invalid"}, struct {
				VpcID                int    "json:\"vpc-id\" xml:\"vpc-id\""
				VpcIfindex           string "json:\"vpc-ifindex\" xml:\"vpc-ifindex\""
				VpcPortState         string "json:\"vpc-port-state\" xml:\"vpc-port-state\""
				PhyPortIfRemoved     string "json:\"phy-port-if-removed\" xml:\"phy-port-if-removed\""
				VpcThruPeerlink      string "json:\"vpc-thru-peerlink\" xml:\"vpc-thru-peerlink\""
				VpcConsistency       string "json:\"vpc-consistency\" xml:\"vpc-consistency\""
				VpcConsistencyStatus string "json:\"vpc-consistency-status\" xml:\"vpc-consistency-status\""
				UpVlanBitset         string "json:\"up-vlan-bitset\" xml:\"up-vlan-bitset\""
				EsAttr               string "json:\"es-attr\" xml:\"es-attr\""
			}{VpcID: 9, VpcIfindex: "Po9", VpcPortState: "1", PhyPortIfRemoved: "disabled", VpcThruPeerlink: "0", VpcConsistency: "consistent", VpcConsistencyStatus: "SUCCESS", UpVlanBitset: "100", EsAttr: "DF: Invalid"}, struct {
				VpcID                int    "json:\"vpc-id\" xml:\"vpc-id\""
				VpcIfindex           string "json:\"vpc-ifindex\" xml:\"vpc-ifindex\""
				VpcPortState         string "json:\"vpc-port-state\" xml:\"vpc-port-state\""
				PhyPortIfRemoved     string "json:\"phy-port-if-removed\" xml:\"phy-port-if-removed\""
				VpcThruPeerlink      string "json:\"vpc-thru-peerlink\" xml:\"vpc-thru-peerlink\""
				VpcConsistency       string "json:\"vpc-consistency\" xml:\"vpc-consistency\""
				VpcConsistencyStatus string "json:\"vpc-consistency-status\" xml:\"vpc-consistency-status\""
				UpVlanBitset         string "json:\"up-vlan-bitset\" xml:\"up-vlan-bitset\""
				EsAttr               string "json:\"es-attr\" xml:\"es-attr\""
			}{VpcID: 10, VpcIfindex: "Po10", VpcPortState: "1", PhyPortIfRemoved: "disabled", VpcThruPeerlink: "0", VpcConsistency: "consistent", VpcConsistencyStatus: "SUCCESS", UpVlanBitset: "400", EsAttr: "DF: Invalid"}, struct {
				VpcID                int    "json:\"vpc-id\" xml:\"vpc-id\""
				VpcIfindex           string "json:\"vpc-ifindex\" xml:\"vpc-ifindex\""
				VpcPortState         string "json:\"vpc-port-state\" xml:\"vpc-port-state\""
				PhyPortIfRemoved     string "json:\"phy-port-if-removed\" xml:\"phy-port-if-removed\""
				VpcThruPeerlink      string "json:\"vpc-thru-peerlink\" xml:\"vpc-thru-peerlink\""
				VpcConsistency       string "json:\"vpc-consistency\" xml:\"vpc-consistency\""
				VpcConsistencyStatus string "json:\"vpc-consistency-status\" xml:\"vpc-consistency-status\""
				UpVlanBitset         string "json:\"up-vlan-bitset\" xml:\"up-vlan-bitset\""
				EsAttr               string "json:\"es-attr\" xml:\"es-attr\""
			}{VpcID: 30, VpcIfindex: "Po30", VpcPortState: "1", PhyPortIfRemoved: "disabled", VpcThruPeerlink: "0", VpcConsistency: "consistent", VpcConsistencyStatus: "SUCCESS", UpVlanBitset: "220", EsAttr: "DF: Invalid"}, struct {
				VpcID                int    "json:\"vpc-id\" xml:\"vpc-id\""
				VpcIfindex           string "json:\"vpc-ifindex\" xml:\"vpc-ifindex\""
				VpcPortState         string "json:\"vpc-port-state\" xml:\"vpc-port-state\""
				PhyPortIfRemoved     string "json:\"phy-port-if-removed\" xml:\"phy-port-if-removed\""
				VpcThruPeerlink      string "json:\"vpc-thru-peerlink\" xml:\"vpc-thru-peerlink\""
				VpcConsistency       string "json:\"vpc-consistency\" xml:\"vpc-consistency\""
				VpcConsistencyStatus string "json:\"vpc-consistency-status\" xml:\"vpc-consistency-status\""
				UpVlanBitset         string "json:\"up-vlan-bitset\" xml:\"up-vlan-bitset\""
				EsAttr               string "json:\"es-attr\" xml:\"es-attr\""
			}{VpcID: 50, VpcIfindex: "Po50", VpcPortState: "1", PhyPortIfRemoved: "disabled", VpcThruPeerlink: "0", VpcConsistency: "consistent", VpcConsistencyStatus: "SUCCESS", UpVlanBitset: "1-501", EsAttr: "DF: Invalid"}}}}}, Code: "200", Input: "show vpc", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"}},
			shouldFail: false,
			shouldErr:  false,
		},
	} {
		fp := fmt.Sprintf("%s/resp.%s.json", outputDir, test.input)
		content, err := ioutil.ReadFile(fp)
		if err != nil {
			t.Logf("FAIL: Test %d: failed reading '%s', error: %v", i, fp, err)
			testFailed++
			continue
		}
		dat, err := NewShowVpcFromBytes(content)
		//fmt.Printf("\n---\n%#v\n---\n", dat) //DEBUG
		if err != nil {
			if !test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but threw error: %v", i, test.input, err)
				testFailed++
				continue
			}
		} else {
			if test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to throw error, but passed: %v", i, test.input, *dat)
				testFailed++
				continue
			}
		}

		if dat != nil {
			if !reflect.DeepEqual(test.exp, dat) {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but failed due to mismatch", i, test.input)
				testFailed++
			}
		}

		if test.shouldFail {
			t.Logf("PASS: Test %d: input '%s', expected to fail, failed", i, test.input)
		} else {
			t.Logf("PASS: Test %d: input '%s', expected to pass, passed", i, test.input)
		}
	}
	if testFailed > 0 {
		t.Fatalf("Failed %d tests", testFailed)
	}
}

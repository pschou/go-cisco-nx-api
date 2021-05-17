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

func TestParseShowPortSecurityAddressJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowPortSecurityAddressResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.port-security.address",
			exp: &ShowPortSecurityAddressResponse{InsAPI: struct {
				Outputs struct {
					Output ShowPortSecurityAddressResponseResult "json:\"output\" xml:\"output\""
				} "json:\"outputs\" xml:\"outputs\""
				Sid     string "json:\"sid\" xml:\"sid\""
				Type    string "json:\"type\" xml:\"type\""
				Version string "json:\"version\" xml:\"version\""
			}{Outputs: struct {
				Output ShowPortSecurityAddressResponseResult "json:\"output\" xml:\"output\""
			}{Output: ShowPortSecurityAddressResponseResult{Body: ShowPortSecurityAddressResultBody{TotalAddr: 5, MaxSysLimit: 8187, TableEthPortSecMacAddrs: []struct {
				RowEthPortSecMacAddrs []struct {
					IfIndex      string "json:\"if_index\" xml:\"if_index\""
					VlanID       int    "json:\"vlan_id\" xml:\"vlan_id\""
					Type         string "json:\"type\" xml:\"type\""
					MacAddr      string "json:\"mac_addr\" xml:\"mac_addr\""
					RemainAge    int    "json:\"remain_age\" xml:\"remain_age\""
					RemoteLearnt int    "json:\"remote_learnt\" xml:\"remote_learnt\""
					RemoteAged   int    "json:\"remote_aged\" xml:\"remote_aged\""
					NumElems     []int  "json:\"num_elems\" xml:\"num_elems\""
					CmdAddrIndex []int  "json:\"cmd_addr_index\" xml:\"cmd_addr_index\""
				} "json:\"ROW_eth_port_sec_mac_addrs\" xml:\"ROW_eth_port_sec_mac_addrs\""
			}{

				{RowEthPortSecMacAddrs: []struct {
					IfIndex      string "json:\"if_index\" xml:\"if_index\""
					VlanID       int    "json:\"vlan_id\" xml:\"vlan_id\""
					Type         string "json:\"type\" xml:\"type\""
					MacAddr      string "json:\"mac_addr\" xml:\"mac_addr\""
					RemainAge    int    "json:\"remain_age\" xml:\"remain_age\""
					RemoteLearnt int    "json:\"remote_learnt\" xml:\"remote_learnt\""
					RemoteAged   int    "json:\"remote_aged\" xml:\"remote_aged\""
					NumElems     []int  "json:\"num_elems\" xml:\"num_elems\""
					CmdAddrIndex []int  "json:\"cmd_addr_index\" xml:\"cmd_addr_index\""
				}{

					{IfIndex: "Ethernet1/2", VlanID: 1, Type: "Static_Mac", MacAddr: "0000.1111.1115", RemainAge: 0, RemoteLearnt: 0, RemoteAged: 0, NumElems: []int{6}, CmdAddrIndex: []int{0}},

					{IfIndex: "Ethernet1/2", VlanID: 1, Type: "Static_Mac", MacAddr: "0000.1111.1111", RemainAge: 0, RemoteLearnt: 0, RemoteAged: 0, NumElems: []int{6}, CmdAddrIndex: []int{1}},

					{IfIndex: "Ethernet1/2", VlanID: 1, Type: "Static_Mac", MacAddr: "0000.1111.1112", RemainAge: 0, RemoteLearnt: 0, RemoteAged: 0, NumElems: []int{6}, CmdAddrIndex: []int{2}},

					{IfIndex: "Ethernet1/2", VlanID: 1, Type: "Static_Mac", MacAddr: "0000.1111.1113", RemainAge: 0, RemoteLearnt: 0, RemoteAged: 0, NumElems: []int{6}, CmdAddrIndex: []int{3}},

					{IfIndex: "Ethernet1/2", VlanID: 1, Type: "Static_Mac", MacAddr: "0000.1111.1114", RemainAge: 0, RemoteLearnt: 0, RemoteAged: 0, NumElems: []int{6}, CmdAddrIndex: []int{4}},

					{IfIndex: "Ethernet1/2", VlanID: 1, Type: "Sticky_Mac", MacAddr: "88F0.31F9.A341", RemainAge: 0, RemoteLearnt: 0, RemoteAged: 0, NumElems: []int{6, 6}, CmdAddrIndex: []int{5, 6}}}}}}, Code: "200", Input: "show port-security address", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"}},
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
		dat, err := NewShowPortSecurityAddressFromBytes(content)
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

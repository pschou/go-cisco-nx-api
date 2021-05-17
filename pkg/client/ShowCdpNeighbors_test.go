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

func TestParseShowCdpNeighborsJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowCdpNeighborsResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.cdp.neighbors",
			exp: &ShowCdpNeighborsResponse{InsAPI: struct {
				Outputs struct {
					Output ShowCdpNeighborsResponseResult "json:\"output\" xml:\"output\""
				} "json:\"outputs\" xml:\"outputs\""
				Sid     string "json:\"sid\" xml:\"sid\""
				Type    string "json:\"type\" xml:\"type\""
				Version string "json:\"version\" xml:\"version\""
			}{Outputs: struct {
				Output ShowCdpNeighborsResponseResult "json:\"output\" xml:\"output\""
			}{Output: ShowCdpNeighborsResponseResult{Body: ShowCdpNeighborsResultBody{NeighCount: 1, TableCdpNeighborBriefInfo: []struct {
				RowCdpNeighborBriefInfo []struct {
					Ifindex    int      "json:\"ifindex\" xml:\"ifindex\""
					DeviceID   string   "json:\"device_id\" xml:\"device_id\""
					IntfID     string   "json:\"intf_id\" xml:\"intf_id\""
					TTL        int      "json:\"ttl\" xml:\"ttl\""
					Capability []string "json:\"capability\" xml:\"capability\""
					PlatformID string   "json:\"platform_id\" xml:\"platform_id\""
					PortID     string   "json:\"port_id\" xml:\"port_id\""
				} "json:\"ROW_cdp_neighbor_brief_info\" xml:\"ROW_cdp_neighbor_brief_info\""
			}{struct {
				RowCdpNeighborBriefInfo []struct {
					Ifindex    int      "json:\"ifindex\" xml:\"ifindex\""
					DeviceID   string   "json:\"device_id\" xml:\"device_id\""
					IntfID     string   "json:\"intf_id\" xml:\"intf_id\""
					TTL        int      "json:\"ttl\" xml:\"ttl\""
					Capability []string "json:\"capability\" xml:\"capability\""
					PlatformID string   "json:\"platform_id\" xml:\"platform_id\""
					PortID     string   "json:\"port_id\" xml:\"port_id\""
				} "json:\"ROW_cdp_neighbor_brief_info\" xml:\"ROW_cdp_neighbor_brief_info\""
			}{RowCdpNeighborBriefInfo: []struct {
				Ifindex    int      "json:\"ifindex\" xml:\"ifindex\""
				DeviceID   string   "json:\"device_id\" xml:\"device_id\""
				IntfID     string   "json:\"intf_id\" xml:\"intf_id\""
				TTL        int      "json:\"ttl\" xml:\"ttl\""
				Capability []string "json:\"capability\" xml:\"capability\""
				PlatformID string   "json:\"platform_id\" xml:\"platform_id\""
				PortID     string   "json:\"port_id\" xml:\"port_id\""
			}{struct {
				Ifindex    int      "json:\"ifindex\" xml:\"ifindex\""
				DeviceID   string   "json:\"device_id\" xml:\"device_id\""
				IntfID     string   "json:\"intf_id\" xml:\"intf_id\""
				TTL        int      "json:\"ttl\" xml:\"ttl\""
				Capability []string "json:\"capability\" xml:\"capability\""
				PlatformID string   "json:\"platform_id\" xml:\"platform_id\""
				PortID     string   "json:\"port_id\" xml:\"port_id\""
			}{Ifindex: 439353856, DeviceID: "EOR-2(FGE18330EZZ)", IntfID: "Ethernet7/2", TTL: 171, Capability: []string{"router", "switch", "IGMP_cnd_filtering", "Supports-STP-Dispute"}, PlatformID: "N9K-C9508", PortID: "Ethernet1/36"}}}}}, Code: "200", Input: "show cdp neighbors interface e7/2", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"}},
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
		dat, err := NewShowCdpNeighborsFromBytes(content)
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

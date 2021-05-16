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

func TestParseShowIpArpJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowIpArpResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.ip.arp",
			exp: &ShowIpArpResponse{InsAPI: struct {
				Outputs struct {
					Output ShowIpArpResponseResult "json:\"output\" xml:\"output\""
				} "json:\"outputs\" xml:\"outputs\""
				Sid     string "json:\"sid\" xml:\"sid\""
				Type    string "json:\"type\" xml:\"type\""
				Version string "json:\"version\" xml:\"version\""
			}{Outputs: struct {
				Output ShowIpArpResponseResult "json:\"output\" xml:\"output\""
			}{Output: ShowIpArpResponseResult{Body: ShowIpArpResultBody{TableVrf: []struct {
				RowVrf []struct {
					TableAdj []struct {
						RowAdj []struct {
							Flags      string   "json:\"flags\" xml:\"flags\""
							IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
							IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
							MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
							TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
							Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
						} "json:\"ROW_adj\" xml:\"ROW_adj\""
					} "json:\"TABLE_adj\" xml:\"TABLE_adj\""
					CntTotal   int    "json:\"cnt-total\" xml:\"cnt-total\""
					VrfNameOut string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
				} "json:\"ROW_vrf\" xml:\"ROW_vrf\""
			}{struct {
				RowVrf []struct {
					TableAdj []struct {
						RowAdj []struct {
							Flags      string   "json:\"flags\" xml:\"flags\""
							IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
							IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
							MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
							TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
							Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
						} "json:\"ROW_adj\" xml:\"ROW_adj\""
					} "json:\"TABLE_adj\" xml:\"TABLE_adj\""
					CntTotal   int    "json:\"cnt-total\" xml:\"cnt-total\""
					VrfNameOut string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
				} "json:\"ROW_vrf\" xml:\"ROW_vrf\""
			}{RowVrf: []struct {
				TableAdj []struct {
					RowAdj []struct {
						Flags      string   "json:\"flags\" xml:\"flags\""
						IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
						IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
						MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
						TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
						Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
					} "json:\"ROW_adj\" xml:\"ROW_adj\""
				} "json:\"TABLE_adj\" xml:\"TABLE_adj\""
				CntTotal   int    "json:\"cnt-total\" xml:\"cnt-total\""
				VrfNameOut string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
			}{struct {
				TableAdj []struct {
					RowAdj []struct {
						Flags      string   "json:\"flags\" xml:\"flags\""
						IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
						IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
						MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
						TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
						Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
					} "json:\"ROW_adj\" xml:\"ROW_adj\""
				} "json:\"TABLE_adj\" xml:\"TABLE_adj\""
				CntTotal   int    "json:\"cnt-total\" xml:\"cnt-total\""
				VrfNameOut string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
			}{TableAdj: []struct {
				RowAdj []struct {
					Flags      string   "json:\"flags\" xml:\"flags\""
					IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
					IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
					MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
					TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
					Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
				} "json:\"ROW_adj\" xml:\"ROW_adj\""
			}{struct {
				RowAdj []struct {
					Flags      string   "json:\"flags\" xml:\"flags\""
					IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
					IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
					MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
					TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
					Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
				} "json:\"ROW_adj\" xml:\"ROW_adj\""
			}{RowAdj: []struct {
				Flags      string   "json:\"flags\" xml:\"flags\""
				IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
				IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
				MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
				TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
				Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
			}{struct {
				Flags      string   "json:\"flags\" xml:\"flags\""
				IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
				IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
				MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
				TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
				Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
			}{Flags: "", IntfOut: "Vlan253", IPAddrOut: "7.57.253.1", MAC: "f44e.0584.7ffc", TimeStamp: 0x10058b6c200, Incomplete: false}, struct {
				Flags      string   "json:\"flags\" xml:\"flags\""
				IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
				IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
				MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
				TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
				Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
			}{Flags: "", IntfOut: "Vlan50", IPAddrOut: "10.57.50.3", MAC: "f44e.0584.7ffc", TimeStamp: 0x3f18dbd600, Incomplete: false}, struct {
				Flags      string   "json:\"flags\" xml:\"flags\""
				IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
				IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
				MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
				TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
				Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
			}{Flags: "", IntfOut: "Ethernet1/2", IPAddrOut: "10.100.157.1", MAC: "f44e.0584.7ffc", TimeStamp: 0x1035f930400, Incomplete: false}, struct {
				Flags      string   "json:\"flags\" xml:\"flags\""
				IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
				IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
				MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
				TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
				Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
			}{Flags: "", IntfOut: "Ethernet1/4", IPAddrOut: "10.100.157.9", MAC: "f44e.0584.7ffc", TimeStamp: 0x1039b2dce00, Incomplete: false}, struct {
				Flags      string   "json:\"flags\" xml:\"flags\""
				IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
				IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
				MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
				TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
				Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
			}{Flags: "", IntfOut: "Ethernet1/7", IPAddrOut: "192.168.161.1", MAC: "f44e.0584.7ffc", TimeStamp: 0xfa86990800, Incomplete: false}, struct {
				Flags      string   "json:\"flags\" xml:\"flags\""
				IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
				IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
				MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
				TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
				Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
			}{Flags: "", IntfOut: "Ethernet1/1.50", IPAddrOut: "89.1.1.10", MAC: "", TimeStamp: 0x684ee1800, Incomplete: true}, struct {
				Flags      string   "json:\"flags\" xml:\"flags\""
				IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
				IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
				MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
				TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
				Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
			}{Flags: "", IntfOut: "Ethernet1/1.52", IPAddrOut: "89.1.3.10", MAC: "", TimeStamp: 0xee6b2800, Incomplete: true}}}}, CntTotal: 7, VrfNameOut: "default"}}}}}, Code: "200", Input: "show ip arp ", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"}},
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
		iparp, err := NewShowIpArpFromBytes(content)
		//fmt.Printf("\n---\n%#v\n---\n", iparp) //DEBUG
		//fmt.Printf("%#v\n", iparp) //DEBUG
		if err != nil {
			if !test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but threw error: %v", i, test.input, err)
				testFailed++
				continue
			}
		} else {
			if test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to throw error, but passed: %v", i, test.input, *iparp)
				testFailed++
				continue
			}
		}

		if iparp != nil {
			if !reflect.DeepEqual(test.exp, iparp) {
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

func TestParseShowIpArpResultJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowIpArpResponseResult
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "result.show.ip.arp",
			exp: &ShowIpArpResponseResult{Body: ShowIpArpResultBody{TableVrf: []struct {
				RowVrf []struct {
					TableAdj []struct {
						RowAdj []struct {
							Flags      string   "json:\"flags\" xml:\"flags\""
							IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
							IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
							MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
							TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
							Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
						} "json:\"ROW_adj\" xml:\"ROW_adj\""
					} "json:\"TABLE_adj\" xml:\"TABLE_adj\""
					CntTotal   int    "json:\"cnt-total\" xml:\"cnt-total\""
					VrfNameOut string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
				} "json:\"ROW_vrf\" xml:\"ROW_vrf\""
			}{struct {
				RowVrf []struct {
					TableAdj []struct {
						RowAdj []struct {
							Flags      string   "json:\"flags\" xml:\"flags\""
							IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
							IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
							MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
							TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
							Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
						} "json:\"ROW_adj\" xml:\"ROW_adj\""
					} "json:\"TABLE_adj\" xml:\"TABLE_adj\""
					CntTotal   int    "json:\"cnt-total\" xml:\"cnt-total\""
					VrfNameOut string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
				} "json:\"ROW_vrf\" xml:\"ROW_vrf\""
			}{RowVrf: []struct {
				TableAdj []struct {
					RowAdj []struct {
						Flags      string   "json:\"flags\" xml:\"flags\""
						IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
						IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
						MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
						TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
						Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
					} "json:\"ROW_adj\" xml:\"ROW_adj\""
				} "json:\"TABLE_adj\" xml:\"TABLE_adj\""
				CntTotal   int    "json:\"cnt-total\" xml:\"cnt-total\""
				VrfNameOut string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
			}{struct {
				TableAdj []struct {
					RowAdj []struct {
						Flags      string   "json:\"flags\" xml:\"flags\""
						IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
						IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
						MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
						TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
						Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
					} "json:\"ROW_adj\" xml:\"ROW_adj\""
				} "json:\"TABLE_adj\" xml:\"TABLE_adj\""
				CntTotal   int    "json:\"cnt-total\" xml:\"cnt-total\""
				VrfNameOut string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
			}{TableAdj: []struct {
				RowAdj []struct {
					Flags      string   "json:\"flags\" xml:\"flags\""
					IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
					IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
					MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
					TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
					Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
				} "json:\"ROW_adj\" xml:\"ROW_adj\""
			}{struct {
				RowAdj []struct {
					Flags      string   "json:\"flags\" xml:\"flags\""
					IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
					IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
					MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
					TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
					Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
				} "json:\"ROW_adj\" xml:\"ROW_adj\""
			}{RowAdj: []struct {
				Flags      string   "json:\"flags\" xml:\"flags\""
				IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
				IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
				MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
				TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
				Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
			}{struct {
				Flags      string   "json:\"flags\" xml:\"flags\""
				IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
				IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
				MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
				TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
				Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
			}{Flags: "", IntfOut: "Vlan253", IPAddrOut: "7.57.253.1", MAC: "f44e.0584.7ffc", TimeStamp: 0x10058b6c200, Incomplete: false}, struct {
				Flags      string   "json:\"flags\" xml:\"flags\""
				IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
				IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
				MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
				TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
				Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
			}{Flags: "", IntfOut: "Vlan50", IPAddrOut: "10.57.50.3", MAC: "f44e.0584.7ffc", TimeStamp: 0x3f18dbd600, Incomplete: false}, struct {
				Flags      string   "json:\"flags\" xml:\"flags\""
				IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
				IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
				MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
				TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
				Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
			}{Flags: "", IntfOut: "Ethernet1/2", IPAddrOut: "10.100.157.1", MAC: "f44e.0584.7ffc", TimeStamp: 0x1035f930400, Incomplete: false}, struct {
				Flags      string   "json:\"flags\" xml:\"flags\""
				IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
				IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
				MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
				TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
				Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
			}{Flags: "", IntfOut: "Ethernet1/4", IPAddrOut: "10.100.157.9", MAC: "f44e.0584.7ffc", TimeStamp: 0x1039b2dce00, Incomplete: false}, struct {
				Flags      string   "json:\"flags\" xml:\"flags\""
				IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
				IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
				MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
				TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
				Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
			}{Flags: "", IntfOut: "Ethernet1/7", IPAddrOut: "192.168.161.1", MAC: "f44e.0584.7ffc", TimeStamp: 0xfa86990800, Incomplete: false}, struct {
				Flags      string   "json:\"flags\" xml:\"flags\""
				IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
				IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
				MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
				TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
				Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
			}{Flags: "", IntfOut: "Ethernet1/1.50", IPAddrOut: "89.1.1.10", MAC: "", TimeStamp: 0x684ee1800, Incomplete: true}, struct {
				Flags      string   "json:\"flags\" xml:\"flags\""
				IntfOut    string   "json:\"intf-out\" xml:\"intf-out\""
				IPAddrOut  string   "json:\"ip-addr-out\" xml:\"ip-addr-out\""
				MAC        string   "json:\"mac,omitempty\" xml:\"mac,omitempty\""
				TimeStamp  Duration "json:\"time-stamp\" xml:\"time-stamp\""
				Incomplete bool     "json:\"incomplete,omitempty\" xml:\"incomplete,omitempty\""
			}{Flags: "", IntfOut: "Ethernet1/1.52", IPAddrOut: "89.1.3.10", MAC: "", TimeStamp: 0xee6b2800, Incomplete: true}}}}, CntTotal: 7, VrfNameOut: "default"}}}}}, Code: "200", Input: "show ip arp ", Msg: "Success"},
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
		iparp, err := NewShowIpArpResultFromBytes(content)
		//fmt.Printf("\n---\n%#v\n---\n", iparp) //DEBUG
		//fmt.Printf("%#v\n", iparp) //DEBUG
		if err != nil {
			if !test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but threw error: %v", i, test.input, err)
				testFailed++
				continue
			}
		} else {
			if test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to throw error, but passed: %v", i, test.input, *iparp)
				testFailed++
				continue
			}
		}

		if iparp != nil {
			if !reflect.DeepEqual(test.exp, iparp) {
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

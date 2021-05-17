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

func TestParseShowIpArpDetailVrfAllJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowIpArpDetailVrfAllResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.ip.eigrp.neighbors.vrf.all",
			exp: &ShowIpArpDetailVrfAllResponse{InsAPI: struct {
				Outputs struct {
					Output ShowIpArpDetailVrfAllResponseResult "json:\"output\" xml:\"output\""
				} "json:\"outputs\" xml:\"outputs\""
				Sid     string "json:\"sid\" xml:\"sid\""
				Type    string "json:\"type\" xml:\"type\""
				Version string "json:\"version\" xml:\"version\""
			}{Outputs: struct {
				Output ShowIpArpDetailVrfAllResponseResult "json:\"output\" xml:\"output\""
			}{Output: ShowIpArpDetailVrfAllResponseResult{Body: ShowIpArpDetailVrfAllResultBody{TableVrf: []struct {
				RowVrf []struct {
					VrfNameOut string "json:\"vrf-name-out\""
					CntTotal   int    "json:\"cnt-total\""
					TableAdj   []struct {
						RowAdj []struct {
							IntfOut   string    "json:\"intf-out\""
							IPAddrOut string    "json:\"ip-addr-out\""
							TimeStamp TimeStamp "json:\"time-stamp\""
							Mac       string    "json:\"mac\""
							PhyIntf   string    "json:\"phy-intf\""
						} "json:\"ROW_adj\""
					} "json:\"TABLE_adj\""
				} "json:\"ROW_vrf\""
			}(nil)}, Code: "200", Input: "show ip eigrp neighbors vrf all", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"}},
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
		dat, err := NewShowIpArpDetailVrfAllFromBytes(content)
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

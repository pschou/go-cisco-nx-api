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

func TestParseShowInterfaceCountersErrorsJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowInterfaceCountersErrorsResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.interface.counters.errors",
			exp: &ShowInterfaceCountersErrorsResponse{InsAPI: struct {
				Outputs struct {
					Output ShowInterfaceCountersErrorsResponseResult "json:\"output\" xml:\"output\""
				} "json:\"outputs\" xml:\"outputs\""
				Sid     string "json:\"sid\" xml:\"sid\""
				Type    string "json:\"type\" xml:\"type\""
				Version string "json:\"version\" xml:\"version\""
			}{Outputs: struct {
				Output ShowInterfaceCountersErrorsResponseResult "json:\"output\" xml:\"output\""
			}{Output: ShowInterfaceCountersErrorsResponseResult{Body: ShowInterfaceCountersErrorsResultBody{TableInterface: []struct {
				RowInterface []struct {
					Interface    string "json:\"interface\" xml:\"interface\""
					EthAlignErr  int    "json:\"eth_align_err\" xml:\"eth_align_err\""
					EthFcsErr    int    "json:\"eth_fcs_err\" xml:\"eth_fcs_err\""
					EthXmitErr   int    "json:\"eth_xmit_err\" xml:\"eth_xmit_err\""
					EthRcvErr    int    "json:\"eth_rcv_err\" xml:\"eth_rcv_err\""
					EthUndersize int    "json:\"eth_undersize\" xml:\"eth_undersize\""
					EthOutdisc   int    "json:\"eth_outdisc\" xml:\"eth_outdisc\""
				} "json:\"ROW_interface\" xml:\"ROW_interface\""
			}{

				{RowInterface: []struct {
					Interface    string "json:\"interface\" xml:\"interface\""
					EthAlignErr  int    "json:\"eth_align_err\" xml:\"eth_align_err\""
					EthFcsErr    int    "json:\"eth_fcs_err\" xml:\"eth_fcs_err\""
					EthXmitErr   int    "json:\"eth_xmit_err\" xml:\"eth_xmit_err\""
					EthRcvErr    int    "json:\"eth_rcv_err\" xml:\"eth_rcv_err\""
					EthUndersize int    "json:\"eth_undersize\" xml:\"eth_undersize\""
					EthOutdisc   int    "json:\"eth_outdisc\" xml:\"eth_outdisc\""
				}{

					{Interface: "Ethernet1/22", EthAlignErr: 0, EthFcsErr: 0, EthXmitErr: 0, EthRcvErr: 0, EthUndersize: 0, EthOutdisc: 0}}}}}, Code: "200", Input: "show interface eth1/22 status err-vlans", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"}},
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
		dat, err := NewShowInterfaceCountersErrorsFromBytes(content)
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

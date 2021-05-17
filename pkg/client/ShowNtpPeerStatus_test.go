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

func TestParseShowNtpPeerStatusJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowNtpPeerStatusResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.ntp.peer-status",
			exp: &ShowNtpPeerStatusResponse{InsAPI: struct {
				Outputs struct {
					Output ShowNtpPeerStatusResponseResult "json:\"output\" xml:\"output\""
				} "json:\"outputs\" xml:\"outputs\""
				Sid     string "json:\"sid\" xml:\"sid\""
				Type    string "json:\"type\" xml:\"type\""
				Version string "json:\"version\" xml:\"version\""
			}{Outputs: struct {
				Output ShowNtpPeerStatusResponseResult "json:\"output\" xml:\"output\""
			}{Output: ShowNtpPeerStatusResponseResult{Body: ShowNtpPeerStatusResultBody{Totalpeers: "Total peers : 5", TablePeersstatus: []struct {
				RowPeersstatus []struct {
					Syncmode string  "json:\"syncmode\" xml:\"syncmode\""
					Remote   string  "json:\"remote\" xml:\"remote\""
					Local    string  "json:\"local\" xml:\"local\""
					St       int     "json:\"st\" xml:\"st\""
					Poll     int     "json:\"poll\" xml:\"poll\""
					Reach    string  "json:\"reach\" xml:\"reach\""
					Delay    float32 "json:\"delay\" xml:\"delay\""
					Vrf      string  "json:\"vrf,omitempty\" xml:\"vrf,omitempty\""
				} "json:\"ROW_peersstatus\" xml:\"ROW_peersstatus\""
			}{struct {
				RowPeersstatus []struct {
					Syncmode string  "json:\"syncmode\" xml:\"syncmode\""
					Remote   string  "json:\"remote\" xml:\"remote\""
					Local    string  "json:\"local\" xml:\"local\""
					St       int     "json:\"st\" xml:\"st\""
					Poll     int     "json:\"poll\" xml:\"poll\""
					Reach    string  "json:\"reach\" xml:\"reach\""
					Delay    float32 "json:\"delay\" xml:\"delay\""
					Vrf      string  "json:\"vrf,omitempty\" xml:\"vrf,omitempty\""
				} "json:\"ROW_peersstatus\" xml:\"ROW_peersstatus\""
			}{RowPeersstatus: []struct {
				Syncmode string  "json:\"syncmode\" xml:\"syncmode\""
				Remote   string  "json:\"remote\" xml:\"remote\""
				Local    string  "json:\"local\" xml:\"local\""
				St       int     "json:\"st\" xml:\"st\""
				Poll     int     "json:\"poll\" xml:\"poll\""
				Reach    string  "json:\"reach\" xml:\"reach\""
				Delay    float32 "json:\"delay\" xml:\"delay\""
				Vrf      string  "json:\"vrf,omitempty\" xml:\"vrf,omitempty\""
			}{struct {
				Syncmode string  "json:\"syncmode\" xml:\"syncmode\""
				Remote   string  "json:\"remote\" xml:\"remote\""
				Local    string  "json:\"local\" xml:\"local\""
				St       int     "json:\"st\" xml:\"st\""
				Poll     int     "json:\"poll\" xml:\"poll\""
				Reach    string  "json:\"reach\" xml:\"reach\""
				Delay    float32 "json:\"delay\" xml:\"delay\""
				Vrf      string  "json:\"vrf,omitempty\" xml:\"vrf,omitempty\""
			}{Syncmode: "=", Remote: "1.1.1.1", Local: "1.1.1.1", St: 16, Poll: 32, Reach: "0", Delay: 0, Vrf: "default"}, struct {
				Syncmode string  "json:\"syncmode\" xml:\"syncmode\""
				Remote   string  "json:\"remote\" xml:\"remote\""
				Local    string  "json:\"local\" xml:\"local\""
				St       int     "json:\"st\" xml:\"st\""
				Poll     int     "json:\"poll\" xml:\"poll\""
				Reach    string  "json:\"reach\" xml:\"reach\""
				Delay    float32 "json:\"delay\" xml:\"delay\""
				Vrf      string  "json:\"vrf,omitempty\" xml:\"vrf,omitempty\""
			}{Syncmode: "+", Remote: "192:168:1::1", Local: "1::1", St: 16, Poll: 32, Reach: "0", Delay: 0, Vrf: "default"}, struct {
				Syncmode string  "json:\"syncmode\" xml:\"syncmode\""
				Remote   string  "json:\"remote\" xml:\"remote\""
				Local    string  "json:\"local\" xml:\"local\""
				St       int     "json:\"st\" xml:\"st\""
				Poll     int     "json:\"poll\" xml:\"poll\""
				Reach    string  "json:\"reach\" xml:\"reach\""
				Delay    float32 "json:\"delay\" xml:\"delay\""
				Vrf      string  "json:\"vrf,omitempty\" xml:\"vrf,omitempty\""
			}{Syncmode: "*", Remote: "127.127.1.0", Local: "1.1.1.1", St: 2, Poll: 32, Reach: "377", Delay: 0, Vrf: ""}, struct {
				Syncmode string  "json:\"syncmode\" xml:\"syncmode\""
				Remote   string  "json:\"remote\" xml:\"remote\""
				Local    string  "json:\"local\" xml:\"local\""
				St       int     "json:\"st\" xml:\"st\""
				Poll     int     "json:\"poll\" xml:\"poll\""
				Reach    string  "json:\"reach\" xml:\"reach\""
				Delay    float32 "json:\"delay\" xml:\"delay\""
				Vrf      string  "json:\"vrf,omitempty\" xml:\"vrf,omitempty\""
			}{Syncmode: "=", Remote: "192:168:1::2", Local: "1::1", St: 16, Poll: 32, Reach: "0", Delay: 0, Vrf: "default"}, struct {
				Syncmode string  "json:\"syncmode\" xml:\"syncmode\""
				Remote   string  "json:\"remote\" xml:\"remote\""
				Local    string  "json:\"local\" xml:\"local\""
				St       int     "json:\"st\" xml:\"st\""
				Poll     int     "json:\"poll\" xml:\"poll\""
				Reach    string  "json:\"reach\" xml:\"reach\""
				Delay    float32 "json:\"delay\" xml:\"delay\""
				Vrf      string  "json:\"vrf,omitempty\" xml:\"vrf,omitempty\""
			}{Syncmode: "=", Remote: "2.2.2.2", Local: "1.1.1.1", St: 16, Poll: 32, Reach: "0", Delay: 0, Vrf: "management"}}}}}, Code: "200", Input: "show ntp peer-status", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"}},
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
		dat, err := NewShowNtpPeerStatusFromBytes(content)
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

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

func TestParseShowModuleJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowModuleResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.module",
			exp: &ShowModuleResponse{InsAPI: struct {
				Outputs struct {
					Output ShowModuleResponseResult "json:\"output\" xml:\"output\""
				} "json:\"outputs\" xml:\"outputs\""
				Sid     string "json:\"sid\" xml:\"sid\""
				Type    string "json:\"type\" xml:\"type\""
				Version string "json:\"version\" xml:\"version\""
			}{Outputs: struct {
				Output ShowModuleResponseResult "json:\"output\" xml:\"output\""
			}{Output: ShowModuleResponseResult{Body: ShowModuleResultBody{TableModdiaginfo: []struct {
				RowModdiaginfo []struct {
					Diagstatus string "json:\"diagstatus\" xml:\"diagstatus\""
					Mod        int    "json:\"mod\" xml:\"mod\""
				} "json:\"ROW_moddiaginfo\" xml:\"ROW_moddiaginfo\""
			}{struct {
				RowModdiaginfo []struct {
					Diagstatus string "json:\"diagstatus\" xml:\"diagstatus\""
					Mod        int    "json:\"mod\" xml:\"mod\""
				} "json:\"ROW_moddiaginfo\" xml:\"ROW_moddiaginfo\""
			}{RowModdiaginfo: []struct {
				Diagstatus string "json:\"diagstatus\" xml:\"diagstatus\""
				Mod        int    "json:\"mod\" xml:\"mod\""
			}{struct {
				Diagstatus string "json:\"diagstatus\" xml:\"diagstatus\""
				Mod        int    "json:\"mod\" xml:\"mod\""
			}{Diagstatus: "Pass", Mod: 1}, struct {
				Diagstatus string "json:\"diagstatus\" xml:\"diagstatus\""
				Mod        int    "json:\"mod\" xml:\"mod\""
			}{Diagstatus: "Pass", Mod: 2}, struct {
				Diagstatus string "json:\"diagstatus\" xml:\"diagstatus\""
				Mod        int    "json:\"mod\" xml:\"mod\""
			}{Diagstatus: "Pass", Mod: 3}, struct {
				Diagstatus string "json:\"diagstatus\" xml:\"diagstatus\""
				Mod        int    "json:\"mod\" xml:\"mod\""
			}{Diagstatus: "Pass", Mod: 4}, struct {
				Diagstatus string "json:\"diagstatus\" xml:\"diagstatus\""
				Mod        int    "json:\"mod\" xml:\"mod\""
			}{Diagstatus: "Pass", Mod: 22}, struct {
				Diagstatus string "json:\"diagstatus\" xml:\"diagstatus\""
				Mod        int    "json:\"mod\" xml:\"mod\""
			}{Diagstatus: "Pass", Mod: 24}, struct {
				Diagstatus string "json:\"diagstatus\" xml:\"diagstatus\""
				Mod        int    "json:\"mod\" xml:\"mod\""
			}{Diagstatus: "Pass", Mod: 26}, struct {
				Diagstatus string "json:\"diagstatus\" xml:\"diagstatus\""
				Mod        int    "json:\"mod\" xml:\"mod\""
			}{Diagstatus: "Pass", Mod: 27}, struct {
				Diagstatus string "json:\"diagstatus\" xml:\"diagstatus\""
				Mod        int    "json:\"mod\" xml:\"mod\""
			}{Diagstatus: "Fail", Mod: 29}, struct {
				Diagstatus string "json:\"diagstatus\" xml:\"diagstatus\""
				Mod        int    "json:\"mod\" xml:\"mod\""
			}{Diagstatus: "Pass", Mod: 30}}}}, TableModinfo: []struct {
				RowModinfo []struct {
					Model   string "json:\"model\" xml:\"model\""
					Modinf  int    "json:\"modinf\" xml:\"modinf\""
					Modtype string "json:\"modtype\" xml:\"modtype\""
					Ports   int    "json:\"ports\" xml:\"ports\""
					Status  string "json:\"status\" xml:\"status\""
				} "json:\"ROW_modinfo\" xml:\"ROW_modinfo\""
			}{struct {
				RowModinfo []struct {
					Model   string "json:\"model\" xml:\"model\""
					Modinf  int    "json:\"modinf\" xml:\"modinf\""
					Modtype string "json:\"modtype\" xml:\"modtype\""
					Ports   int    "json:\"ports\" xml:\"ports\""
					Status  string "json:\"status\" xml:\"status\""
				} "json:\"ROW_modinfo\" xml:\"ROW_modinfo\""
			}{RowModinfo: []struct {
				Model   string "json:\"model\" xml:\"model\""
				Modinf  int    "json:\"modinf\" xml:\"modinf\""
				Modtype string "json:\"modtype\" xml:\"modtype\""
				Ports   int    "json:\"ports\" xml:\"ports\""
				Status  string "json:\"status\" xml:\"status\""
			}{struct {
				Model   string "json:\"model\" xml:\"model\""
				Modinf  int    "json:\"modinf\" xml:\"modinf\""
				Modtype string "json:\"modtype\" xml:\"modtype\""
				Ports   int    "json:\"ports\" xml:\"ports\""
				Status  string "json:\"status\" xml:\"status\""
			}{Model: "N9K-X9732C-EX", Modinf: 1, Modtype: "32x100G Ethernet Module", Ports: 32, Status: "ok"}, struct {
				Model   string "json:\"model\" xml:\"model\""
				Modinf  int    "json:\"modinf\" xml:\"modinf\""
				Modtype string "json:\"modtype\" xml:\"modtype\""
				Ports   int    "json:\"ports\" xml:\"ports\""
				Status  string "json:\"status\" xml:\"status\""
			}{Model: "N9K-X9732C-EXM", Modinf: 2, Modtype: "32x100G Ethernet Module", Ports: 32, Status: "ok"}, struct {
				Model   string "json:\"model\" xml:\"model\""
				Modinf  int    "json:\"modinf\" xml:\"modinf\""
				Modtype string "json:\"modtype\" xml:\"modtype\""
				Ports   int    "json:\"ports\" xml:\"ports\""
				Status  string "json:\"status\" xml:\"status\""
			}{Model: "N9K-X9788TC-FX", Modinf: 3, Modtype: "48x10G + 4x40/100G Ethernet Module", Ports: 52, Status: "ok"}, struct {
				Model   string "json:\"model\" xml:\"model\""
				Modinf  int    "json:\"modinf\" xml:\"modinf\""
				Modtype string "json:\"modtype\" xml:\"modtype\""
				Ports   int    "json:\"ports\" xml:\"ports\""
				Status  string "json:\"status\" xml:\"status\""
			}{Model: "N9K-X9736C-FX", Modinf: 4, Modtype: "36x40/100G Ethernet Module", Ports: 36, Status: "ok"}, struct {
				Model   string "json:\"model\" xml:\"model\""
				Modinf  int    "json:\"modinf\" xml:\"modinf\""
				Modtype string "json:\"modtype\" xml:\"modtype\""
				Ports   int    "json:\"ports\" xml:\"ports\""
				Status  string "json:\"status\" xml:\"status\""
			}{Model: "N9K-C9508-FM-E", Modinf: 22, Modtype: "8-slot Fabric Module", Ports: 0, Status: "ok"}, struct {
				Model   string "json:\"model\" xml:\"model\""
				Modinf  int    "json:\"modinf\" xml:\"modinf\""
				Modtype string "json:\"modtype\" xml:\"modtype\""
				Ports   int    "json:\"ports\" xml:\"ports\""
				Status  string "json:\"status\" xml:\"status\""
			}{Model: "N9K-C9508-FM-E", Modinf: 23, Modtype: "8-slot Fabric Module", Ports: 0, Status: "powered-dn"}, struct {
				Model   string "json:\"model\" xml:\"model\""
				Modinf  int    "json:\"modinf\" xml:\"modinf\""
				Modtype string "json:\"modtype\" xml:\"modtype\""
				Ports   int    "json:\"ports\" xml:\"ports\""
				Status  string "json:\"status\" xml:\"status\""
			}{Model: "N9K-C9508-FM-E", Modinf: 24, Modtype: "8-slot Fabric Module", Ports: 0, Status: "ok"}, struct {
				Model   string "json:\"model\" xml:\"model\""
				Modinf  int    "json:\"modinf\" xml:\"modinf\""
				Modtype string "json:\"modtype\" xml:\"modtype\""
				Ports   int    "json:\"ports\" xml:\"ports\""
				Status  string "json:\"status\" xml:\"status\""
			}{Model: "N9K-C9508-FM-E", Modinf: 26, Modtype: "8-slot Fabric Module", Ports: 0, Status: "ok"}, struct {
				Model   string "json:\"model\" xml:\"model\""
				Modinf  int    "json:\"modinf\" xml:\"modinf\""
				Modtype string "json:\"modtype\" xml:\"modtype\""
				Ports   int    "json:\"ports\" xml:\"ports\""
				Status  string "json:\"status\" xml:\"status\""
			}{Model: "N9K-SUP-A", Modinf: 27, Modtype: "Supervisor Module", Ports: 0, Status: "active *"}, struct {
				Model   string "json:\"model\" xml:\"model\""
				Modinf  int    "json:\"modinf\" xml:\"modinf\""
				Modtype string "json:\"modtype\" xml:\"modtype\""
				Ports   int    "json:\"ports\" xml:\"ports\""
				Status  string "json:\"status\" xml:\"status\""
			}{Model: "N9K-SC-A", Modinf: 29, Modtype: "System Controller", Ports: 0, Status: "standby"}, struct {
				Model   string "json:\"model\" xml:\"model\""
				Modinf  int    "json:\"modinf\" xml:\"modinf\""
				Modtype string "json:\"modtype\" xml:\"modtype\""
				Ports   int    "json:\"ports\" xml:\"ports\""
				Status  string "json:\"status\" xml:\"status\""
			}{Model: "N9K-SC-A", Modinf: 30, Modtype: "System Controller", Ports: 0, Status: "active"}}}}, TableModmacinfo: []struct {
				RowModmacinfo []struct {
					Mac       string "json:\"mac\" xml:\"mac\""
					Modmac    int    "json:\"modmac\" xml:\"modmac\""
					Serialnum string "json:\"serialnum\" xml:\"serialnum\""
				} "json:\"ROW_modmacinfo\" xml:\"ROW_modmacinfo\""
			}{struct {
				RowModmacinfo []struct {
					Mac       string "json:\"mac\" xml:\"mac\""
					Modmac    int    "json:\"modmac\" xml:\"modmac\""
					Serialnum string "json:\"serialnum\" xml:\"serialnum\""
				} "json:\"ROW_modmacinfo\" xml:\"ROW_modmacinfo\""
			}{RowModmacinfo: []struct {
				Mac       string "json:\"mac\" xml:\"mac\""
				Modmac    int    "json:\"modmac\" xml:\"modmac\""
				Serialnum string "json:\"serialnum\" xml:\"serialnum\""
			}{struct {
				Mac       string "json:\"mac\" xml:\"mac\""
				Modmac    int    "json:\"modmac\" xml:\"modmac\""
				Serialnum string "json:\"serialnum\" xml:\"serialnum\""
			}{Mac: " 00-a2-ee-31-b2-b4 to 00-a2-ee-31-b3-37  ", Modmac: 1, Serialnum: "SAL2039VA2U"}, struct {
				Mac       string "json:\"mac\" xml:\"mac\""
				Modmac    int    "json:\"modmac\" xml:\"modmac\""
				Serialnum string "json:\"serialnum\" xml:\"serialnum\""
			}{Mac: " 00-6b-f1-bf-0d-f0 to 00-6b-f1-bf-0e-73  ", Modmac: 2, Serialnum: "FOC20444YD2"}, struct {
				Mac       string "json:\"mac\" xml:\"mac\""
				Modmac    int    "json:\"modmac\" xml:\"modmac\""
				Serialnum string "json:\"serialnum\" xml:\"serialnum\""
			}{Mac: " a0-23-9f-0c-8f-28 to a0-23-9f-0c-8f-6b  ", Modmac: 3, Serialnum: "FOC21240U8B"}, struct {
				Mac       string "json:\"mac\" xml:\"mac\""
				Modmac    int    "json:\"modmac\" xml:\"modmac\""
				Serialnum string "json:\"serialnum\" xml:\"serialnum\""
			}{Mac: " 2c-d0-2d-3e-a5-08 to 2c-d0-2d-3e-a5-9b  ", Modmac: 4, Serialnum: "FOC210707HG"}, struct {
				Mac       string "json:\"mac\" xml:\"mac\""
				Modmac    int    "json:\"modmac\" xml:\"modmac\""
				Serialnum string "json:\"serialnum\" xml:\"serialnum\""
			}{Mac: " NA ", Modmac: 22, Serialnum: "SAL2035URFY"}, struct {
				Mac       string "json:\"mac\" xml:\"mac\""
				Modmac    int    "json:\"modmac\" xml:\"modmac\""
				Serialnum string "json:\"serialnum\" xml:\"serialnum\""
			}{Mac: " NA ", Modmac: 23, Serialnum: "SAL2035UUN3"}, struct {
				Mac       string "json:\"mac\" xml:\"mac\""
				Modmac    int    "json:\"modmac\" xml:\"modmac\""
				Serialnum string "json:\"serialnum\" xml:\"serialnum\""
			}{Mac: " NA ", Modmac: 24, Serialnum: "SAL2035URF2"}, struct {
				Mac       string "json:\"mac\" xml:\"mac\""
				Modmac    int    "json:\"modmac\" xml:\"modmac\""
				Serialnum string "json:\"serialnum\" xml:\"serialnum\""
			}{Mac: " NA ", Modmac: 26, Serialnum: "SAL2035URFP"}, struct {
				Mac       string "json:\"mac\" xml:\"mac\""
				Modmac    int    "json:\"modmac\" xml:\"modmac\""
				Serialnum string "json:\"serialnum\" xml:\"serialnum\""
			}{Mac: " cc-46-d6-9e-d2-e4 to cc-46-d6-9e-d2-f5  ", Modmac: 27, Serialnum: "SAL2015NQ3H"}, struct {
				Mac       string "json:\"mac\" xml:\"mac\""
				Modmac    int    "json:\"modmac\" xml:\"modmac\""
				Serialnum string "json:\"serialnum\" xml:\"serialnum\""
			}{Mac: " NA ", Modmac: 29, Serialnum: "SAL2035UTLL"}, struct {
				Mac       string "json:\"mac\" xml:\"mac\""
				Modmac    int    "json:\"modmac\" xml:\"modmac\""
				Serialnum string "json:\"serialnum\" xml:\"serialnum\""
			}{Mac: " NA ", Modmac: 30, Serialnum: "SAL2035UTK8"}}}}, TableModpwrinfo: []struct {
				RowModpwrinfo []struct {
					Modpwr  int    "json:\"modpwr\" xml:\"modpwr\""
					Pwrstat string "json:\"pwrstat\" xml:\"pwrstat\""
					Reason  string "json:\"reason\" xml:\"reason\""
				} "json:\"ROW_modpwrinfo\" xml:\"ROW_modpwrinfo\""
			}{struct {
				RowModpwrinfo []struct {
					Modpwr  int    "json:\"modpwr\" xml:\"modpwr\""
					Pwrstat string "json:\"pwrstat\" xml:\"pwrstat\""
					Reason  string "json:\"reason\" xml:\"reason\""
				} "json:\"ROW_modpwrinfo\" xml:\"ROW_modpwrinfo\""
			}{RowModpwrinfo: []struct {
				Modpwr  int    "json:\"modpwr\" xml:\"modpwr\""
				Pwrstat string "json:\"pwrstat\" xml:\"pwrstat\""
				Reason  string "json:\"reason\" xml:\"reason\""
			}{struct {
				Modpwr  int    "json:\"modpwr\" xml:\"modpwr\""
				Pwrstat string "json:\"pwrstat\" xml:\"pwrstat\""
				Reason  string "json:\"reason\" xml:\"reason\""
			}{Modpwr: 23, Pwrstat: "powered-dn", Reason: "Configured Power down"}}}}, TableModwwninfo: []struct {
				RowModwwninfo []struct {
					Hw       string "json:\"hw\" xml:\"hw\""
					Modwwn   int    "json:\"modwwn\" xml:\"modwwn\""
					Slottype string "json:\"slottype\" xml:\"slottype\""
					Sw       string "json:\"sw\" xml:\"sw\""
				} "json:\"ROW_modwwninfo\" xml:\"ROW_modwwninfo\""
			}{struct {
				RowModwwninfo []struct {
					Hw       string "json:\"hw\" xml:\"hw\""
					Modwwn   int    "json:\"modwwn\" xml:\"modwwn\""
					Slottype string "json:\"slottype\" xml:\"slottype\""
					Sw       string "json:\"sw\" xml:\"sw\""
				} "json:\"ROW_modwwninfo\" xml:\"ROW_modwwninfo\""
			}{RowModwwninfo: []struct {
				Hw       string "json:\"hw\" xml:\"hw\""
				Modwwn   int    "json:\"modwwn\" xml:\"modwwn\""
				Slottype string "json:\"slottype\" xml:\"slottype\""
				Sw       string "json:\"sw\" xml:\"sw\""
			}{struct {
				Hw       string "json:\"hw\" xml:\"hw\""
				Modwwn   int    "json:\"modwwn\" xml:\"modwwn\""
				Slottype string "json:\"slottype\" xml:\"slottype\""
				Sw       string "json:\"sw\" xml:\"sw\""
			}{Hw: "1.1", Modwwn: 1, Slottype: "LC1", Sw: "7.0(3)I7(4)"}, struct {
				Hw       string "json:\"hw\" xml:\"hw\""
				Modwwn   int    "json:\"modwwn\" xml:\"modwwn\""
				Slottype string "json:\"slottype\" xml:\"slottype\""
				Sw       string "json:\"sw\" xml:\"sw\""
			}{Hw: "0.1220", Modwwn: 2, Slottype: "LC2", Sw: "7.0(3)I7(4)"}, struct {
				Hw       string "json:\"hw\" xml:\"hw\""
				Modwwn   int    "json:\"modwwn\" xml:\"modwwn\""
				Slottype string "json:\"slottype\" xml:\"slottype\""
				Sw       string "json:\"sw\" xml:\"sw\""
			}{Hw: "0.1110", Modwwn: 3, Slottype: "LC3", Sw: "7.0(3)I7(4)"}, struct {
				Hw       string "json:\"hw\" xml:\"hw\""
				Modwwn   int    "json:\"modwwn\" xml:\"modwwn\""
				Slottype string "json:\"slottype\" xml:\"slottype\""
				Sw       string "json:\"sw\" xml:\"sw\""
			}{Hw: "0.1120", Modwwn: 4, Slottype: "LC4", Sw: "7.0(3)I7(4)"}, struct {
				Hw       string "json:\"hw\" xml:\"hw\""
				Modwwn   int    "json:\"modwwn\" xml:\"modwwn\""
				Slottype string "json:\"slottype\" xml:\"slottype\""
				Sw       string "json:\"sw\" xml:\"sw\""
			}{Hw: "1.1", Modwwn: 22, Slottype: "FM2", Sw: "7.0(3)I7(4)"}, struct {
				Hw       string "json:\"hw\" xml:\"hw\""
				Modwwn   int    "json:\"modwwn\" xml:\"modwwn\""
				Slottype string "json:\"slottype\" xml:\"slottype\""
				Sw       string "json:\"sw\" xml:\"sw\""
			}{Hw: "1.1", Modwwn: 24, Slottype: "FM4", Sw: "7.0(3)I7(4)"}, struct {
				Hw       string "json:\"hw\" xml:\"hw\""
				Modwwn   int    "json:\"modwwn\" xml:\"modwwn\""
				Slottype string "json:\"slottype\" xml:\"slottype\""
				Sw       string "json:\"sw\" xml:\"sw\""
			}{Hw: "1.1", Modwwn: 26, Slottype: "FM6", Sw: "7.0(3)I7(4)"}, struct {
				Hw       string "json:\"hw\" xml:\"hw\""
				Modwwn   int    "json:\"modwwn\" xml:\"modwwn\""
				Slottype string "json:\"slottype\" xml:\"slottype\""
				Sw       string "json:\"sw\" xml:\"sw\""
			}{Hw: "1.0", Modwwn: 27, Slottype: "SUP1", Sw: "7.0(3)I7(4)"}, struct {
				Hw       string "json:\"hw\" xml:\"hw\""
				Modwwn   int    "json:\"modwwn\" xml:\"modwwn\""
				Slottype string "json:\"slottype\" xml:\"slottype\""
				Sw       string "json:\"sw\" xml:\"sw\""
			}{Hw: "1.6", Modwwn: 29, Slottype: "SC1", Sw: "7.0(3)I7(4)"}, struct {
				Hw       string "json:\"hw\" xml:\"hw\""
				Modwwn   int    "json:\"modwwn\" xml:\"modwwn\""
				Slottype string "json:\"slottype\" xml:\"slottype\""
				Sw       string "json:\"sw\" xml:\"sw\""
			}{Hw: "1.6", Modwwn: 30, Slottype: "SC2", Sw: "7.0(3)I7(4)"}}}}}, Code: "200", Input: "show module", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"}},
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
		dat, err := NewShowModuleFromBytes(content)
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

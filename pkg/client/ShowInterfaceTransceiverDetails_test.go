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

func TestParseShowShowInterfaceTransceiverDetailsJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowInterfaceTransceiverDetailsResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.interface.transceiver.details",
			exp: &ShowInterfaceTransceiverDetailsResponse{InsAPI: struct {
				Outputs struct {
					Output ShowInterfaceTransceiverDetailsResponseResult "json:\"output\" xml:\"output\""
				} "json:\"outputs\" xml:\"outputs\""
				Sid     string "json:\"sid\" xml:\"sid\""
				Type    string "json:\"type\" xml:\"type\""
				Version string "json:\"version\" xml:\"version\""
			}{Outputs: struct {
				Output ShowInterfaceTransceiverDetailsResponseResult "json:\"output\" xml:\"output\""
			}{Output: ShowInterfaceTransceiverDetailsResponseResult{Body: ShowInterfaceTransceiverDetailsResultBody{TableInterface: []struct {
				RowInterface []struct {
					Interface       string "json:\"interface\" xml:\"interface\""
					Sfp             string "json:\"sfp\" xml:\"sfp\""
					Type            string "json:\"type,omitempty\" xml:\"type,omitempty\""
					Name            string "json:\"name,omitempty\" xml:\"name,omitempty\""
					PartNum         string "json:\"partnum,omitempty\" xml:\"partnum,omitempty\""
					Rev             string "json:\"rev,omitempty\" xml:\"rev,omitempty\""
					SerialNum       string "json:\"serialnum,omitempty\" xml:\"serialnum,omitempty\""
					NomBitrate      int    "json:\"nom_bitrate,omitempty\" xml:\"nom_bitrate,omitempty\""
					Len50OM3        int    "json:\"len_50_OM3,omitempty\" xml:\"len_50_OM3,omitempty\""
					CiscoID         string "json:\"ciscoid,omitempty\" xml:\"ciscoid,omitempty\""
					CiscoID1        int    "json:\"ciscoid_1,omitempty\" xml:\"ciscoid_1,omitempty\""
					CiscoPartNumber string "json:\"cisco_part_number,omitempty\" xml:\"cisco_part_number,omitempty\""
					CiscoProductID  string "json:\"cisco_product_id,omitempty\" xml:\"cisco_product_id,omitempty\""
					TABLELane       []struct {
						ROWLane []struct {
							LaneNumber    int      "json:\"lane_number\" xml:\"lane_number\""
							Temperature   int      "json:\"temperature\" xml:\"temperature\""
							TempAlrmHi    int      "json:\"temp_alrm_hi\" xml:\"temp_alrm_hi\""
							TempAlrmLo    int      "json:\"temp_alrm_lo\" xml:\"temp_alrm_lo\""
							TempWarnHi    int      "json:\"temp_warn_hi\" xml:\"temp_warn_hi\""
							TempWarnLo    int      "json:\"temp_warn_lo\" xml:\"temp_warn_lo\""
							Voltage       int      "json:\"voltage\" xml:\"voltage\""
							VoltAlrmHi    int      "json:\"volt_alrm_hi\" xml:\"volt_alrm_hi\""
							VoltAlrmLo    int      "json:\"volt_alrm_lo\" xml:\"volt_alrm_lo\""
							VoltWarnHi    int      "json:\"volt_warn_hi\" xml:\"volt_warn_hi\""
							VoltWarnLo    int      "json:\"volt_warn_lo\" xml:\"volt_warn_lo\""
							Current       int      "json:\"current\" xml:\"current\""
							CurrentAlrmHi int      "json:\"current_alrm_hi\" xml:\"current_alrm_hi\""
							CurrentAlrmLo int      "json:\"current_alrm_lo\" xml:\"current_alrm_lo\""
							CurrentWarnHi int      "json:\"current_warn_hi\" xml:\"current_warn_hi\""
							CurrentWarnLo int      "json:\"current_warn_lo\" xml:\"current_warn_lo\""
							TxPwr         int      "json:\"tx_pwr\" xml:\"tx_pwr\""
							TxPwrAlrmHi   int      "json:\"tx_pwr_alrm_hi\" xml:\"tx_pwr_alrm_hi\""
							TxPwrAlrmLo   int      "json:\"tx_pwr_alrm_lo\" xml:\"tx_pwr_alrm_lo\""
							TxPwrWarnHi   int      "json:\"tx_pwr_warn_hi\" xml:\"tx_pwr_warn_hi\""
							TxPwrWarnLo   int      "json:\"tx_pwr_warn_lo\" xml:\"tx_pwr_warn_lo\""
							RxPwr         int      "json:\"rx_pwr\" xml:\"rx_pwr\""
							RxPwrAlrmHi   int      "json:\"rx_pwr_alrm_hi\" xml:\"rx_pwr_alrm_hi\""
							RxPwrAlrmLo   int      "json:\"rx_pwr_alrm_lo\" xml:\"rx_pwr_alrm_lo\""
							RxPwrWarnHi   int      "json:\"rx_pwr_warn_hi\" xml:\"rx_pwr_warn_hi\""
							RxPwrWarnLo   int      "json:\"rx_pwr_warn_lo\" xml:\"rx_pwr_warn_lo\""
							XmitFaults    int      "json:\"xmit_faults\" xml:\"xmit_faults\""
							RxPwrFlag     []string "json:\"rx_pwr_flag\" xml:\"rx_pwr_flag\""
							TxPwrFlag     []string "json:\"tx_pwr_flag\" xml:\"tx_pwr_flag\""
							TempFlag      []string "json:\"temp_flag\" xml:\"temp_flag\""
							VoltFlag      []string "json:\"volt_flag\" xml:\"volt_flag\""
							CurrentFlag   []string "json:\"current_flag\" xml:\"current_flag\""
						} "json:\"ROW_lane\" xml:\"ROW_lane\""
					} "json:\"TABLE_lane,omitempty\" xml:\"TABLE_lane,omitempty\""
					CiscoVendorID string "json:\"cisco_vendor_id,omitempty\" xml:\"cisco_vendor_id,omitempty\""
				} "json:\"ROW_interface\" xml:\"ROW_interface\""
			}(nil)}, Code: "200", Input: "show ip route ", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"}},
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
		dat, err := NewShowInterfaceTransceiverDetailsFromBytes(content)
		//fmt.Printf("\n---\n%#v\n---\n", dat) //DEBUG
		//fmt.Printf("%#v\n", dat.Flat()) //DEBUG
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
func TestParseShowShowInterfaceTransceiverDetailsResultJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowInterfaceTransceiverDetailsResponseResult
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "result.show.interface.transceiver.details",
			exp: &ShowInterfaceTransceiverDetailsResponseResult{Body: ShowInterfaceTransceiverDetailsResultBody{TableInterface: []struct {
				RowInterface []struct {
					Interface       string "json:\"interface\" xml:\"interface\""
					Sfp             string "json:\"sfp\" xml:\"sfp\""
					Type            string "json:\"type,omitempty\" xml:\"type,omitempty\""
					Name            string "json:\"name,omitempty\" xml:\"name,omitempty\""
					PartNum         string "json:\"partnum,omitempty\" xml:\"partnum,omitempty\""
					Rev             string "json:\"rev,omitempty\" xml:\"rev,omitempty\""
					SerialNum       string "json:\"serialnum,omitempty\" xml:\"serialnum,omitempty\""
					NomBitrate      int    "json:\"nom_bitrate,omitempty\" xml:\"nom_bitrate,omitempty\""
					Len50OM3        int    "json:\"len_50_OM3,omitempty\" xml:\"len_50_OM3,omitempty\""
					CiscoID         string "json:\"ciscoid,omitempty\" xml:\"ciscoid,omitempty\""
					CiscoID1        int    "json:\"ciscoid_1,omitempty\" xml:\"ciscoid_1,omitempty\""
					CiscoPartNumber string "json:\"cisco_part_number,omitempty\" xml:\"cisco_part_number,omitempty\""
					CiscoProductID  string "json:\"cisco_product_id,omitempty\" xml:\"cisco_product_id,omitempty\""
					TABLELane       []struct {
						ROWLane []struct {
							LaneNumber    int      "json:\"lane_number\" xml:\"lane_number\""
							Temperature   int      "json:\"temperature\" xml:\"temperature\""
							TempAlrmHi    int      "json:\"temp_alrm_hi\" xml:\"temp_alrm_hi\""
							TempAlrmLo    int      "json:\"temp_alrm_lo\" xml:\"temp_alrm_lo\""
							TempWarnHi    int      "json:\"temp_warn_hi\" xml:\"temp_warn_hi\""
							TempWarnLo    int      "json:\"temp_warn_lo\" xml:\"temp_warn_lo\""
							Voltage       int      "json:\"voltage\" xml:\"voltage\""
							VoltAlrmHi    int      "json:\"volt_alrm_hi\" xml:\"volt_alrm_hi\""
							VoltAlrmLo    int      "json:\"volt_alrm_lo\" xml:\"volt_alrm_lo\""
							VoltWarnHi    int      "json:\"volt_warn_hi\" xml:\"volt_warn_hi\""
							VoltWarnLo    int      "json:\"volt_warn_lo\" xml:\"volt_warn_lo\""
							Current       int      "json:\"current\" xml:\"current\""
							CurrentAlrmHi int      "json:\"current_alrm_hi\" xml:\"current_alrm_hi\""
							CurrentAlrmLo int      "json:\"current_alrm_lo\" xml:\"current_alrm_lo\""
							CurrentWarnHi int      "json:\"current_warn_hi\" xml:\"current_warn_hi\""
							CurrentWarnLo int      "json:\"current_warn_lo\" xml:\"current_warn_lo\""
							TxPwr         int      "json:\"tx_pwr\" xml:\"tx_pwr\""
							TxPwrAlrmHi   int      "json:\"tx_pwr_alrm_hi\" xml:\"tx_pwr_alrm_hi\""
							TxPwrAlrmLo   int      "json:\"tx_pwr_alrm_lo\" xml:\"tx_pwr_alrm_lo\""
							TxPwrWarnHi   int      "json:\"tx_pwr_warn_hi\" xml:\"tx_pwr_warn_hi\""
							TxPwrWarnLo   int      "json:\"tx_pwr_warn_lo\" xml:\"tx_pwr_warn_lo\""
							RxPwr         int      "json:\"rx_pwr\" xml:\"rx_pwr\""
							RxPwrAlrmHi   int      "json:\"rx_pwr_alrm_hi\" xml:\"rx_pwr_alrm_hi\""
							RxPwrAlrmLo   int      "json:\"rx_pwr_alrm_lo\" xml:\"rx_pwr_alrm_lo\""
							RxPwrWarnHi   int      "json:\"rx_pwr_warn_hi\" xml:\"rx_pwr_warn_hi\""
							RxPwrWarnLo   int      "json:\"rx_pwr_warn_lo\" xml:\"rx_pwr_warn_lo\""
							XmitFaults    int      "json:\"xmit_faults\" xml:\"xmit_faults\""
							RxPwrFlag     []string "json:\"rx_pwr_flag\" xml:\"rx_pwr_flag\""
							TxPwrFlag     []string "json:\"tx_pwr_flag\" xml:\"tx_pwr_flag\""
							TempFlag      []string "json:\"temp_flag\" xml:\"temp_flag\""
							VoltFlag      []string "json:\"volt_flag\" xml:\"volt_flag\""
							CurrentFlag   []string "json:\"current_flag\" xml:\"current_flag\""
						} "json:\"ROW_lane\" xml:\"ROW_lane\""
					} "json:\"TABLE_lane,omitempty\" xml:\"TABLE_lane,omitempty\""
					CiscoVendorID string "json:\"cisco_vendor_id,omitempty\" xml:\"cisco_vendor_id,omitempty\""
				} "json:\"ROW_interface\" xml:\"ROW_interface\""
			}(nil)}, Code: "", Input: "", Msg: ""},
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
		dat, err := NewShowInterfaceTransceiverDetailsResultFromBytes(content)
		//fmt.Printf("\n---\n%#v\n---\n", dat) //DEBUG
		//fmt.Printf("%#v\n", dat.Flat()) //DEBUG
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

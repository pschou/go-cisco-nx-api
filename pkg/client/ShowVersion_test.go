// Copyright 2018 Paul Greenberg (greenpau@outlook.com)
//            and Paul Schou     (github.com/pschou)
//
// Licensed under the Apache License, ShowVersion 2.0 (the "License");
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

func TestParseShowVersionResponseJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowVersionResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.version",
			exp: &ShowVersionResponse{InsAPI: struct {
				Outputs struct {
					Output ShowVersionResponseResult "json:\"output\" xml:\"output\""
				} "json:\"outputs\" xml:\"outputs\""
				Sid         string "json:\"sid\" xml:\"sid\""
				Type        string "json:\"type\" xml:\"type\""
				ShowVersion string "json:\"version\" xml:\"version\""
			}{Outputs: struct {
				Output ShowVersionResponseResult "json:\"output\" xml:\"output\""
			}{Output: ShowVersionResponseResult{Body: ShowVersionResultBody{HeaderStr: "Cisco Nexus Operating System (NX-OS) Software\nTAC support: http://www.cisco.com/tac\nCopyright (C) 2002-2018, Cisco and/or its affiliates.\nAll rights reserved.\nThe copyrights to certain works contained in this software are\nowned by other third parties and used and distributed under their own\nlicenses, such as open source.  This software is provided \"as is,\" and unless\notherwise stated, there is no warranty, express or implied, including but not\nlimited to warranties of merchantability and fitness for a particular purpose.\nCertain components of this software are licensed under\nthe GNU General Public License (GPL) version 2.0 or \nGNU General Public License (GPL) version 3.0  or the GNU\nLesser General Public License (LGPL) Version 2.1 or \nLesser General Public License (LGPL) Version 2.0. \nA copy of each such license is available at\nhttp://www.opensource.org/licenses/gpl-2.0.php and\nhttp://opensource.org/licenses/gpl-3.0.html and\nhttp://www.opensource.org/licenses/lgpl-2.1.php and\nhttp://www.gnu.org/licenses/old-licenses/library.txt.\n", BiosVerStr: "08.32", KickstartVerStr: "7.0(3)I7(4)", NxosVerStr: "", BiosCmplTime: 0x58056600, KickFileName: "bootflash:///nxos.7.0.3.I7.4.bin", NxosFileName: "", KickCmplTime: 0x5b043070, NxosCmplTime: 0x0, KickTmstmp: 0x5b043690, NxosTmstmp: 0x0, ChassisID: "Nexus9000 C9508 (8 Slot) Chassis", CPUName: "Intel(R) Xeon(R) CPU E5-2403 0 @ 1.80GHz", Memory: 0xfa418c, MemType: "kB", ProcBoardID: "SAL2015NQ3H", HostName: "macsec2", BootflashSize: 0x14b0512, KernUptmDays: 0x0, KernUptmHrs: 0x5, KernUptmMins: 0x3, KernUptmSecs: 0x18, RrUsecs: 0xa6696, RrCtime: 0x5b05b244, RrReason: "Reset Requested by CLI command reload", RrSysVer: "7.0(3)I7(4)", RrService: "", Plugins: "", Manufacturer: "Cisco Systems, Inc.", TablePackageList: []struct {
				RowPackageList []struct {
					PackageID string "json:\"package_id\" xml:\"package_id\""
				} "json:\"ROW_package_list\" xml:\"ROW_package_list\""
			}{

				{RowPackageList: []struct {
					PackageID string "json:\"package_id\" xml:\"package_id\""
				}{

					{PackageID: ""}}}}}, Code: "200", Input: "show version", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", ShowVersion: "1.0"}},
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
		dat, err := NewShowVersionFromBytes(content)
		//fmt.Printf("\n---\n%#v\n---\n", dat) //DEBUG
		//fmt.Printf("%#v\n", dat) //DEBUG
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

func TestParseShowVersionResponseResultJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowVersionResponseResult
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "result.show.version",
			exp: &ShowVersionResponseResult{Body: ShowVersionResultBody{HeaderStr: "Cisco Nexus Operating System (NX-OS) Software\nTAC support: http://www.cisco.com/tac\nCopyright (C) 2002-2018, Cisco and/or its affiliates.\nAll rights reserved.\nThe copyrights to certain works contained in this software are\nowned by other third parties and used and distributed under their own\nlicenses, such as open source.  This software is provided \"as is,\" and unless\notherwise stated, there is no warranty, express or implied, including but not\nlimited to warranties of merchantability and fitness for a particular purpose.\nCertain components of this software are licensed under\nthe GNU General Public License (GPL) version 2.0 or \nGNU General Public License (GPL) version 3.0  or the GNU\nLesser General Public License (LGPL) Version 2.1 or \nLesser General Public License (LGPL) Version 2.0. \nA copy of each such license is available at\nhttp://www.opensource.org/licenses/gpl-2.0.php and\nhttp://opensource.org/licenses/gpl-3.0.html and\nhttp://www.opensource.org/licenses/lgpl-2.1.php and\nhttp://www.gnu.org/licenses/old-licenses/library.txt.\n", BiosVerStr: "08.32", KickstartVerStr: "7.0(3)I7(4)", NxosVerStr: "", BiosCmplTime: 0x58056600, KickFileName: "bootflash:///nxos.7.0.3.I7.4.bin", NxosFileName: "", KickCmplTime: 0x5b043070, NxosCmplTime: 0x0, KickTmstmp: 0x5b043690, NxosTmstmp: 0x0, ChassisID: "Nexus9000 C9508 (8 Slot) Chassis", CPUName: "Intel(R) Xeon(R) CPU E5-2403 0 @ 1.80GHz", Memory: 0xfa418c, MemType: "kB", ProcBoardID: "SAL2015NQ3H", HostName: "macsec2", BootflashSize: 0x14b0512, KernUptmDays: 0x0, KernUptmHrs: 0x5, KernUptmMins: 0x3, KernUptmSecs: 0x18, RrUsecs: 0xa6696, RrCtime: 0x5b05b244, RrReason: "Reset Requested by CLI command reload", RrSysVer: "7.0(3)I7(4)", RrService: "", Plugins: "", Manufacturer: "Cisco Systems, Inc.", TablePackageList: []struct {
				RowPackageList []struct {
					PackageID string "json:\"package_id\" xml:\"package_id\""
				} "json:\"ROW_package_list\" xml:\"ROW_package_list\""
			}{

				{RowPackageList: []struct {
					PackageID string "json:\"package_id\" xml:\"package_id\""
				}{

					{PackageID: ""}}}}}, Code: "200", Input: "show version", Msg: "Success"},
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
		dat, err := NewShowVersionResultFromBytes(content)
		//fmt.Printf("\n---\n%#v\n---\n", dat) //DEBUG
		//fmt.Printf("%#v\n", dat) //DEBUG
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

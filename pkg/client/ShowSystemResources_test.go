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

func TestParseShowSystemResourcesJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowSystemResourcesResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.system.resources",
			exp: &ShowSystemResourcesResponse{InsAPI: struct {
				Outputs struct {
					Output ShowSystemResourcesResponseResult "json:\"output\" xml:\"output\""
				} "json:\"outputs\" xml:\"outputs\""
				Sid     string "json:\"sid\" xml:\"sid\""
				Type    string "json:\"type\" xml:\"type\""
				Version string "json:\"version\" xml:\"version\""
			}{Outputs: struct {
				Output ShowSystemResourcesResponseResult "json:\"output\" xml:\"output\""
			}{Output: ShowSystemResourcesResponseResult{Body: ShowSystemResourcesResultBody{LoadAvg1Min: 0.8, LoadAvg5Min: 0.68, LoadAvg15Min: 0.67, ProcessesTotal: 827, ProcessesRunning: 2, CPUStateUser: 7.36, CPUStateKernel: 8.62, CPUStateIdle: 84.01, MemoryUsageTotal: 16400304, MemoryUsageUsed: 7069080, MemoryUsageFree: 9331224, CurrentMemoryStatus: "OK", TableCPUUsage: []struct {
				RowCPUUsage []struct {
					Cpuid  int     "json:\"cpuid\" xml:\"cpuid\""
					User   float32 "json:\"user\" xml:\"user\""
					Kernel float32 "json:\"kernel\" xml:\"kernel\""
					Idle   float32 "json:\"idle\" xml:\"idle\""
				} "json:\"ROW_cpu_usage\" xml:\"ROW_cpu_usage\""
			}{

				{RowCPUUsage: []struct {
					Cpuid  int     "json:\"cpuid\" xml:\"cpuid\""
					User   float32 "json:\"user\" xml:\"user\""
					Kernel float32 "json:\"kernel\" xml:\"kernel\""
					Idle   float32 "json:\"idle\" xml:\"idle\""
				}{

					{Cpuid: 0, User: 13.13, Kernel: 12.12, Idle: 74.74},

					{Cpuid: 1, User: 5.1, Kernel: 13.26, Idle: 81.63},

					{Cpuid: 2, User: 10.2, Kernel: 4.08, Idle: 85.71},

					{Cpuid: 3, User: 1.01, Kernel: 5.05, Idle: 93.93}}}}}, Code: "200", Input: "show system resources", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"}},
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
		dat, err := NewShowSystemResourcesFromBytes(content)
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

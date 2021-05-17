// Copyright 2018 Paul Greenberg (greenpau@outlook.com)
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
	"testing"
)

func TestParseShowSystemResourcesJsonOutputOrig(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"
	for i, test := range []struct {
		input        string
		exp          *SystemResources
		memoryStatus string
		cpuCount     int
		shouldFail   bool
		shouldErr    bool
	}{
		{
			input:        "show.system.resources.1",
			exp:          &SystemResources{},
			memoryStatus: "OK",
			cpuCount:     4,
			shouldFail:   false,
			shouldErr:    false,
		},
	} {
		fp := fmt.Sprintf("%s/resp.%s.json", outputDir, test.input)
		content, err := ioutil.ReadFile(fp)
		if err != nil {
			t.Logf("FAIL: Test %d: failed reading '%s', error: %v", i, fp, err)
			testFailed++
			continue
		}
		resources, err := NewSystemResourcesFromBytes(content)
		if err != nil {
			if !test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but threw error: %v", i, test.input, err)
				testFailed++
				continue
			}
		} else {
			if test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to throw error, but passed: %v", i, test.input, resources)
				testFailed++
				continue
			}
		}

		if (resources.Memory.Status != test.memoryStatus) && !test.shouldFail {
			t.Logf("FAIL: Test %d: input '%s', expected to pass, but failed due to memoryStatus [%s] != %s", i, test.input, resources.Memory.Status, test.memoryStatus)
			testFailed++
			continue
		}
		if (len(resources.CPUs) != test.cpuCount) && !test.shouldFail {
			t.Logf("FAIL: Test %d: input '%s', expected to pass, but failed due to cpuCount [%d] != %d", i, test.input, len(resources.CPUs), test.cpuCount)
			testFailed++
			continue
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

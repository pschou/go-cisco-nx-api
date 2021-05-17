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

func TestParseShowEnvironmentJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowEnvironmentResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.environment",
			exp: &ShowEnvironmentResponse{InsAPI: struct {
				Outputs struct {
					Output ShowEnvironmentResponseResult "json:\"output\" xml:\"output\""
				} "json:\"outputs\" xml:\"outputs\""
				Sid     string "json:\"sid\" xml:\"sid\""
				Type    string "json:\"type\" xml:\"type\""
				Version string "json:\"version\" xml:\"version\""
			}{Outputs: struct {
				Output ShowEnvironmentResponseResult "json:\"output\" xml:\"output\""
			}{Output: ShowEnvironmentResponseResult{Body: ShowEnvironmentResultBody{TableTempinfo: []struct {
				RowTempinfo []struct {
					AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
					CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
					MajThres    float32 "json:\"majthres\" xml:\"majthres\""
					MinThres    float32 "json:\"minthres\" xml:\"minthres\""
					Sensor      string  "json:\"sensor\" xml:\"sensor\""
					Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
				} "json:\"ROW_tempinfo\" xml:\"ROW_tempinfo\""
			}{struct {
				RowTempinfo []struct {
					AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
					CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
					MajThres    float32 "json:\"majthres\" xml:\"majthres\""
					MinThres    float32 "json:\"minthres\" xml:\"minthres\""
					Sensor      string  "json:\"sensor\" xml:\"sensor\""
					Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
				} "json:\"ROW_tempinfo\" xml:\"ROW_tempinfo\""
			}{RowTempinfo: []struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 34, MajThres: 85, MinThres: 75, Sensor: "CPU", Tempmod: 1}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 50, MajThres: 105, MinThres: 95, Sensor: "SUG0", Tempmod: 1}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 65, MajThres: 105, MinThres: 95, Sensor: "SUG1", Tempmod: 1}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 49, MajThres: 105, MinThres: 95, Sensor: "SUG2", Tempmod: 1}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 61, MajThres: 105, MinThres: 95, Sensor: "SUG3", Tempmod: 1}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 50, MajThres: 110, MinThres: 100, Sensor: "VRM1", Tempmod: 1}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 55, MajThres: 110, MinThres: 100, Sensor: "VRM2", Tempmod: 1}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 45, MajThres: 110, MinThres: 100, Sensor: "VRM3", Tempmod: 1}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 51, MajThres: 110, MinThres: 100, Sensor: "VRM4", Tempmod: 1}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 34, MajThres: 85, MinThres: 75, Sensor: "CPU", Tempmod: 2}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 59, MajThres: 105, MinThres: 95, Sensor: "SUG0", Tempmod: 2}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 77, MajThres: 105, MinThres: 95, Sensor: "SUG1", Tempmod: 2}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 59, MajThres: 105, MinThres: 95, Sensor: "SUG2", Tempmod: 2}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 76, MajThres: 105, MinThres: 95, Sensor: "SUG3", Tempmod: 2}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 52, MajThres: 110, MinThres: 100, Sensor: "VRM1", Tempmod: 2}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 53, MajThres: 110, MinThres: 100, Sensor: "VRM2", Tempmod: 2}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 54, MajThres: 110, MinThres: 100, Sensor: "VRM3", Tempmod: 2}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 55, MajThres: 110, MinThres: 100, Sensor: "VRM4", Tempmod: 2}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 37, MajThres: 85, MinThres: 75, Sensor: "CPU", Tempmod: 3}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 51, MajThres: 105, MinThres: 95, Sensor: "HOM0", Tempmod: 3}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 53, MajThres: 105, MinThres: 95, Sensor: "HOM1", Tempmod: 3}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 46, MajThres: 110, MinThres: 100, Sensor: "VRM1", Tempmod: 3}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 47, MajThres: 110, MinThres: 100, Sensor: "VRM2", Tempmod: 3}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 35, MajThres: 85, MinThres: 75, Sensor: "CPU", Tempmod: 4}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 53, MajThres: 105, MinThres: 95, Sensor: "HOM0", Tempmod: 4}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 63, MajThres: 105, MinThres: 95, Sensor: "HOM1", Tempmod: 4}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 63, MajThres: 105, MinThres: 95, Sensor: "HOM2", Tempmod: 4}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 63, MajThres: 105, MinThres: 95, Sensor: "HOM3", Tempmod: 4}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 45, MajThres: 110, MinThres: 100, Sensor: "VRM1", Tempmod: 4}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 48, MajThres: 110, MinThres: 100, Sensor: "VRM2", Tempmod: 4}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 46, MajThres: 110, MinThres: 100, Sensor: "VRM3", Tempmod: 4}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 46, MajThres: 110, MinThres: 100, Sensor: "VRM4", Tempmod: 4}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 31, MajThres: 85, MinThres: 75, Sensor: "CPU", Tempmod: 22}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 56, MajThres: 105, MinThres: 95, Sensor: "LAC0", Tempmod: 22}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 42, MajThres: 105, MinThres: 95, Sensor: "LAC1", Tempmod: 22}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 56, MajThres: 110, MinThres: 100, Sensor: "VRM1", Tempmod: 22}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 56, MajThres: 110, MinThres: 100, Sensor: "VRM2", Tempmod: 22}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 28, MajThres: 85, MinThres: 75, Sensor: "CPU", Tempmod: 24}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 57, MajThres: 105, MinThres: 95, Sensor: "LAC0", Tempmod: 24}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 42, MajThres: 105, MinThres: 95, Sensor: "LAC1", Tempmod: 24}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 57, MajThres: 110, MinThres: 100, Sensor: "VRM1", Tempmod: 24}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 57, MajThres: 110, MinThres: 100, Sensor: "VRM2", Tempmod: 24}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 28, MajThres: 85, MinThres: 75, Sensor: "CPU", Tempmod: 26}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 51, MajThres: 105, MinThres: 95, Sensor: "LAC0", Tempmod: 26}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 43, MajThres: 105, MinThres: 95, Sensor: "LAC1", Tempmod: 26}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 52, MajThres: 110, MinThres: 100, Sensor: "VRM1", Tempmod: 26}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 52, MajThres: 110, MinThres: 100, Sensor: "VRM2", Tempmod: 26}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 33, MajThres: 75, MinThres: 55, Sensor: "OUTLET", Tempmod: 27}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 24, MajThres: 60, MinThres: 42, Sensor: "INLET", Tempmod: 27}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 35, MajThres: 90, MinThres: 80, Sensor: "CPU", Tempmod: 27}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 44, MajThres: 105, MinThres: 95, Sensor: "CPU", Tempmod: 29}, struct {
				AlarmStatus string  "json:\"alarmstatus\" xml:\"alarmstatus\""
				CurTemp     float32 "json:\"curtemp\" xml:\"curtemp\""
				MajThres    float32 "json:\"majthres\" xml:\"majthres\""
				MinThres    float32 "json:\"minthres\" xml:\"minthres\""
				Sensor      string  "json:\"sensor\" xml:\"sensor\""
				Tempmod     int     "json:\"tempmod\" xml:\"tempmod\""
			}{AlarmStatus: "Ok", CurTemp: 53, MajThres: 105, MinThres: 95, Sensor: "CPU", Tempmod: 30}}}}, FanDetails: struct {
				TableFanZoneSpeed []struct {
					RowFanZoneSpeed []struct {
						Zone      int    "json:\"zone\" xml:\"zone\""
						ZoneSpeed string "json:\"zonespeed\" xml:\"zonespeed\""
					} "json:\"ROW_fan_zone_speed\" xml:\"ROW_fan_zone_speed\""
				} "json:\"TABLE_fan_zone_speed\" xml:\"TABLE_fan_zone_speed\""
				TableFaninfo []struct {
					RowFaninfo []struct {
						FanDir    string "json:\"fandir\" xml:\"fandir\""
						FanHwVer  string "json:\"fanhwver\" xml:\"fanhwver\""
						FanModel  string "json:\"fanmodel\" xml:\"fanmodel\""
						FanName   string "json:\"fanname\" xml:\"fanname\""
						FanStatus string "json:\"fanstatus\" xml:\"fanstatus\""
					} "json:\"ROW_faninfo\" xml:\"ROW_faninfo\""
				} "json:\"TABLE_faninfo\" xml:\"TABLE_faninfo\""
				FanFilterStatus string "json:\"fan_filter_status\" xml:\"fan_filter_status\""
			}{TableFanZoneSpeed: []struct {
				RowFanZoneSpeed []struct {
					Zone      int    "json:\"zone\" xml:\"zone\""
					ZoneSpeed string "json:\"zonespeed\" xml:\"zonespeed\""
				} "json:\"ROW_fan_zone_speed\" xml:\"ROW_fan_zone_speed\""
			}{struct {
				RowFanZoneSpeed []struct {
					Zone      int    "json:\"zone\" xml:\"zone\""
					ZoneSpeed string "json:\"zonespeed\" xml:\"zonespeed\""
				} "json:\"ROW_fan_zone_speed\" xml:\"ROW_fan_zone_speed\""
			}{RowFanZoneSpeed: []struct {
				Zone      int    "json:\"zone\" xml:\"zone\""
				ZoneSpeed string "json:\"zonespeed\" xml:\"zonespeed\""
			}{struct {
				Zone      int    "json:\"zone\" xml:\"zone\""
				ZoneSpeed string "json:\"zonespeed\" xml:\"zonespeed\""
			}{Zone: 1, ZoneSpeed: "0x0"}}}}, TableFaninfo: []struct {
				RowFaninfo []struct {
					FanDir    string "json:\"fandir\" xml:\"fandir\""
					FanHwVer  string "json:\"fanhwver\" xml:\"fanhwver\""
					FanModel  string "json:\"fanmodel\" xml:\"fanmodel\""
					FanName   string "json:\"fanname\" xml:\"fanname\""
					FanStatus string "json:\"fanstatus\" xml:\"fanstatus\""
				} "json:\"ROW_faninfo\" xml:\"ROW_faninfo\""
			}{struct {
				RowFaninfo []struct {
					FanDir    string "json:\"fandir\" xml:\"fandir\""
					FanHwVer  string "json:\"fanhwver\" xml:\"fanhwver\""
					FanModel  string "json:\"fanmodel\" xml:\"fanmodel\""
					FanName   string "json:\"fanname\" xml:\"fanname\""
					FanStatus string "json:\"fanstatus\" xml:\"fanstatus\""
				} "json:\"ROW_faninfo\" xml:\"ROW_faninfo\""
			}{RowFaninfo: []struct {
				FanDir    string "json:\"fandir\" xml:\"fandir\""
				FanHwVer  string "json:\"fanhwver\" xml:\"fanhwver\""
				FanModel  string "json:\"fanmodel\" xml:\"fanmodel\""
				FanName   string "json:\"fanname\" xml:\"fanname\""
				FanStatus string "json:\"fanstatus\" xml:\"fanstatus\""
			}{struct {
				FanDir    string "json:\"fandir\" xml:\"fandir\""
				FanHwVer  string "json:\"fanhwver\" xml:\"fanhwver\""
				FanModel  string "json:\"fanmodel\" xml:\"fanmodel\""
				FanName   string "json:\"fanname\" xml:\"fanname\""
				FanStatus string "json:\"fanstatus\" xml:\"fanstatus\""
			}{FanDir: "front-to-back", FanHwVer: "1.0", FanModel: "N9K-C9508-FAN", FanName: "Fan1(sys_fan1)", FanStatus: "Ok"}, struct {
				FanDir    string "json:\"fandir\" xml:\"fandir\""
				FanHwVer  string "json:\"fanhwver\" xml:\"fanhwver\""
				FanModel  string "json:\"fanmodel\" xml:\"fanmodel\""
				FanName   string "json:\"fanname\" xml:\"fanname\""
				FanStatus string "json:\"fanstatus\" xml:\"fanstatus\""
			}{FanDir: "front-to-back", FanHwVer: "1.0", FanModel: "N9K-C9508-FAN", FanName: "Fan2(sys_fan2)", FanStatus: "Ok"}, struct {
				FanDir    string "json:\"fandir\" xml:\"fandir\""
				FanHwVer  string "json:\"fanhwver\" xml:\"fanhwver\""
				FanModel  string "json:\"fanmodel\" xml:\"fanmodel\""
				FanName   string "json:\"fanname\" xml:\"fanname\""
				FanStatus string "json:\"fanstatus\" xml:\"fanstatus\""
			}{FanDir: "front-to-back", FanHwVer: "1.0", FanModel: "N9K-C9508-FAN", FanName: "Fan3(sys_fan3)", FanStatus: "Ok"}, struct {
				FanDir    string "json:\"fandir\" xml:\"fandir\""
				FanHwVer  string "json:\"fanhwver\" xml:\"fanhwver\""
				FanModel  string "json:\"fanmodel\" xml:\"fanmodel\""
				FanName   string "json:\"fanname\" xml:\"fanname\""
				FanStatus string "json:\"fanstatus\" xml:\"fanstatus\""
			}{FanDir: "front-to-back", FanHwVer: "--", FanModel: "", FanName: "Fan_in_PS1", FanStatus: "Ok"}, struct {
				FanDir    string "json:\"fandir\" xml:\"fandir\""
				FanHwVer  string "json:\"fanhwver\" xml:\"fanhwver\""
				FanModel  string "json:\"fanmodel\" xml:\"fanmodel\""
				FanName   string "json:\"fanname\" xml:\"fanname\""
				FanStatus string "json:\"fanstatus\" xml:\"fanstatus\""
			}{FanDir: "front-to-back", FanHwVer: "--", FanModel: "", FanName: "Fan_in_PS2", FanStatus: "Ok"}, struct {
				FanDir    string "json:\"fandir\" xml:\"fandir\""
				FanHwVer  string "json:\"fanhwver\" xml:\"fanhwver\""
				FanModel  string "json:\"fanmodel\" xml:\"fanmodel\""
				FanName   string "json:\"fanname\" xml:\"fanname\""
				FanStatus string "json:\"fanstatus\" xml:\"fanstatus\""
			}{FanDir: "front-to-back", FanHwVer: "--", FanModel: "", FanName: "Fan_in_PS3", FanStatus: "Ok"}, struct {
				FanDir    string "json:\"fandir\" xml:\"fandir\""
				FanHwVer  string "json:\"fanhwver\" xml:\"fanhwver\""
				FanModel  string "json:\"fanmodel\" xml:\"fanmodel\""
				FanName   string "json:\"fanname\" xml:\"fanname\""
				FanStatus string "json:\"fanstatus\" xml:\"fanstatus\""
			}{FanDir: "front-to-back", FanHwVer: "--", FanModel: "", FanName: "Fan_in_PS4", FanStatus: "None"}, struct {
				FanDir    string "json:\"fandir\" xml:\"fandir\""
				FanHwVer  string "json:\"fanhwver\" xml:\"fanhwver\""
				FanModel  string "json:\"fanmodel\" xml:\"fanmodel\""
				FanName   string "json:\"fanname\" xml:\"fanname\""
				FanStatus string "json:\"fanstatus\" xml:\"fanstatus\""
			}{FanDir: "front-to-back", FanHwVer: "--", FanModel: "", FanName: "Fan_in_PS5", FanStatus: "None"}, struct {
				FanDir    string "json:\"fandir\" xml:\"fandir\""
				FanHwVer  string "json:\"fanhwver\" xml:\"fanhwver\""
				FanModel  string "json:\"fanmodel\" xml:\"fanmodel\""
				FanName   string "json:\"fanname\" xml:\"fanname\""
				FanStatus string "json:\"fanstatus\" xml:\"fanstatus\""
			}{FanDir: "front-to-back", FanHwVer: "--", FanModel: "", FanName: "Fan_in_PS6", FanStatus: "None"}, struct {
				FanDir    string "json:\"fandir\" xml:\"fandir\""
				FanHwVer  string "json:\"fanhwver\" xml:\"fanhwver\""
				FanModel  string "json:\"fanmodel\" xml:\"fanmodel\""
				FanName   string "json:\"fanname\" xml:\"fanname\""
				FanStatus string "json:\"fanstatus\" xml:\"fanstatus\""
			}{FanDir: "front-to-back", FanHwVer: "--", FanModel: "", FanName: "Fan_in_PS7", FanStatus: "None"}, struct {
				FanDir    string "json:\"fandir\" xml:\"fandir\""
				FanHwVer  string "json:\"fanhwver\" xml:\"fanhwver\""
				FanModel  string "json:\"fanmodel\" xml:\"fanmodel\""
				FanName   string "json:\"fanname\" xml:\"fanname\""
				FanStatus string "json:\"fanstatus\" xml:\"fanstatus\""
			}{FanDir: "front-to-back", FanHwVer: "--", FanModel: "", FanName: "Fan_in_PS8", FanStatus: "None"}}}}, FanFilterStatus: "NotSupported"}, Powersup: struct {
				TableModPowInfo []struct {
					RowModPowInfo []struct {
						ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
						Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
						ModModel   string "json:\"mod_model\" xml:\"mod_model\""
						Modnum     string "json:\"modnum\" xml:\"modnum\""
						Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
					} "json:\"ROW_mod_pow_info\" xml:\"ROW_mod_pow_info\""
				} "json:\"TABLE_mod_pow_info\" xml:\"TABLE_mod_pow_info\""
				TablePsinfo []struct {
					RowPsinfo []struct {
						ActualInput Watts  "json:\"actual_input\" xml:\"actual_input\""
						ActualOut   Watts  "json:\"actual_out\" xml:\"actual_out\""
						PsStatus    string "json:\"ps_status\" xml:\"ps_status\""
						PsModel     string "json:\"psmodel\" xml:\"psmodel\""
						PsNum       int    "json:\"psnum\" xml:\"psnum\""
						TotCapa     Watts  "json:\"tot_capa\" xml:\"tot_capa\""
					} "json:\"ROW_psinfo\" xml:\"ROW_psinfo\""
				} "json:\"TABLE_psinfo\" xml:\"TABLE_psinfo\""
				PowerSummary struct {
					AvailablePow          Watts  "json:\"available_pow\" xml:\"available_pow\""
					CumulativePower       Watts  "json:\"cumulative_power\" xml:\"cumulative_power\""
					PsOperMode            string "json:\"ps_oper_mode\" xml:\"ps_oper_mode\""
					PsRedunMode           string "json:\"ps_redun_mode\" xml:\"ps_redun_mode\""
					TotGridaCapacity      Watts  "json:\"tot_gridA_capacity\" xml:\"tot_gridA_capacity\""
					TotGridbCapacity      Watts  "json:\"tot_gridB_capacity\" xml:\"tot_gridB_capacity\""
					TotPowAllocBudgeted   Watts  "json:\"tot_pow_alloc_budgeted\" xml:\"tot_pow_alloc_budgeted\""
					TotPowCapacity        Watts  "json:\"tot_pow_capacity\" xml:\"tot_pow_capacity\""
					TotPowInputActualDraw Watts  "json:\"tot_pow_input_actual_draw\" xml:\"tot_pow_input_actual_draw\""
					TotPowOutActualDraw   Watts  "json:\"tot_pow_out_actual_draw\" xml:\"tot_pow_out_actual_draw\""
				} "json:\"power_summary\" xml:\"power_summary\""
				VoltageLevel int "json:\"voltage_level\" xml:\"voltage_level\""
			}{TableModPowInfo: []struct {
				RowModPowInfo []struct {
					ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
					Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
					ModModel   string "json:\"mod_model\" xml:\"mod_model\""
					Modnum     string "json:\"modnum\" xml:\"modnum\""
					Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
				} "json:\"ROW_mod_pow_info\" xml:\"ROW_mod_pow_info\""
			}{struct {
				RowModPowInfo []struct {
					ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
					Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
					ModModel   string "json:\"mod_model\" xml:\"mod_model\""
					Modnum     string "json:\"modnum\" xml:\"modnum\""
					Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
				} "json:\"ROW_mod_pow_info\" xml:\"ROW_mod_pow_info\""
			}{RowModPowInfo: []struct {
				ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
				Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
				ModModel   string "json:\"mod_model\" xml:\"mod_model\""
				Modnum     string "json:\"modnum\" xml:\"modnum\""
				Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
			}{struct {
				ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
				Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
				ModModel   string "json:\"mod_model\" xml:\"mod_model\""
				Modnum     string "json:\"modnum\" xml:\"modnum\""
				Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
			}{ActualDraw: 356, Allocated: 720, ModModel: "N9K-X9732C-EX", Modnum: "1", Modstatus: "Powered-Up"}, struct {
				ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
				Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
				ModModel   string "json:\"mod_model\" xml:\"mod_model\""
				Modnum     string "json:\"modnum\" xml:\"modnum\""
				Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
			}{ActualDraw: 573, Allocated: 816, ModModel: "N9K-X9732C-EXM", Modnum: "2", Modstatus: "Powered-Up"}, struct {
				ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
				Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
				ModModel   string "json:\"mod_model\" xml:\"mod_model\""
				Modnum     string "json:\"modnum\" xml:\"modnum\""
				Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
			}{ActualDraw: 226, Allocated: 684, ModModel: "N9K-X9788TC-FX", Modnum: "3", Modstatus: "Powered-Up"}, struct {
				ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
				Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
				ModModel   string "json:\"mod_model\" xml:\"mod_model\""
				Modnum     string "json:\"modnum\" xml:\"modnum\""
				Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
			}{ActualDraw: 379, Allocated: 756, ModModel: "N9K-X9736C-FX", Modnum: "4", Modstatus: "Powered-Up"}, struct {
				ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
				Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
				ModModel   string "json:\"mod_model\" xml:\"mod_model\""
				Modnum     string "json:\"modnum\" xml:\"modnum\""
				Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
			}{ActualDraw: 9999, Allocated: 0, ModModel: "xbar", Modnum: "Xb21", Modstatus: "Absent"}, struct {
				ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
				Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
				ModModel   string "json:\"mod_model\" xml:\"mod_model\""
				Modnum     string "json:\"modnum\" xml:\"modnum\""
				Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
			}{ActualDraw: 109, Allocated: 564, ModModel: "N9K-C9508-FM-E", Modnum: "Xb22", Modstatus: "Powered-Up"}, struct {
				ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
				Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
				ModModel   string "json:\"mod_model\" xml:\"mod_model\""
				Modnum     string "json:\"modnum\" xml:\"modnum\""
				Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
			}{ActualDraw: 9999, Allocated: 0, ModModel: "N9K-C9508-FM-E", Modnum: "Xb23", Modstatus: "Powered-Dn"}, struct {
				ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
				Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
				ModModel   string "json:\"mod_model\" xml:\"mod_model\""
				Modnum     string "json:\"modnum\" xml:\"modnum\""
				Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
			}{ActualDraw: 108, Allocated: 564, ModModel: "N9K-C9508-FM-E", Modnum: "Xb24", Modstatus: "Powered-Up"}, struct {
				ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
				Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
				ModModel   string "json:\"mod_model\" xml:\"mod_model\""
				Modnum     string "json:\"modnum\" xml:\"modnum\""
				Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
			}{ActualDraw: 9999, Allocated: 0, ModModel: "xbar", Modnum: "Xb25", Modstatus: "Absent"}, struct {
				ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
				Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
				ModModel   string "json:\"mod_model\" xml:\"mod_model\""
				Modnum     string "json:\"modnum\" xml:\"modnum\""
				Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
			}{ActualDraw: 104, Allocated: 564, ModModel: "N9K-C9508-FM-E", Modnum: "Xb26", Modstatus: "Powered-Up"}, struct {
				ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
				Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
				ModModel   string "json:\"mod_model\" xml:\"mod_model\""
				Modnum     string "json:\"modnum\" xml:\"modnum\""
				Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
			}{ActualDraw: 63, Allocated: 90, ModModel: "N9K-SUP-A", Modnum: "27", Modstatus: "Powered-Up"}, struct {
				ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
				Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
				ModModel   string "json:\"mod_model\" xml:\"mod_model\""
				Modnum     string "json:\"modnum\" xml:\"modnum\""
				Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
			}{ActualDraw: 9999, Allocated: 0, ModModel: "supervisor", Modnum: "28", Modstatus: "Absent"}, struct {
				ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
				Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
				ModModel   string "json:\"mod_model\" xml:\"mod_model\""
				Modnum     string "json:\"modnum\" xml:\"modnum\""
				Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
			}{ActualDraw: 14, Allocated: 25.2, ModModel: "N9K-SC-A", Modnum: "29", Modstatus: "Powered-Up"}, struct {
				ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
				Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
				ModModel   string "json:\"mod_model\" xml:\"mod_model\""
				Modnum     string "json:\"modnum\" xml:\"modnum\""
				Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
			}{ActualDraw: 14, Allocated: 25.2, ModModel: "N9K-SC-A", Modnum: "30", Modstatus: "Powered-Up"}, struct {
				ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
				Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
				ModModel   string "json:\"mod_model\" xml:\"mod_model\""
				Modnum     string "json:\"modnum\" xml:\"modnum\""
				Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
			}{ActualDraw: 38, Allocated: 249, ModModel: "N9K-C9508-FAN", Modnum: "fan1", Modstatus: "Powered-Up"}, struct {
				ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
				Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
				ModModel   string "json:\"mod_model\" xml:\"mod_model\""
				Modnum     string "json:\"modnum\" xml:\"modnum\""
				Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
			}{ActualDraw: 39, Allocated: 249, ModModel: "N9K-C9508-FAN", Modnum: "fan2", Modstatus: "Powered-Up"}, struct {
				ActualDraw Watts  "json:\"actual_draw\" xml:\"actual_draw\""
				Allocated  Watts  "json:\"allocated\" xml:\"allocated\""
				ModModel   string "json:\"mod_model\" xml:\"mod_model\""
				Modnum     string "json:\"modnum\" xml:\"modnum\""
				Modstatus  string "json:\"modstatus\" xml:\"modstatus\""
			}{ActualDraw: 36, Allocated: 249, ModModel: "N9K-C9508-FAN", Modnum: "fan3", Modstatus: "Powered-Up"}}}}, TablePsinfo: []struct {
				RowPsinfo []struct {
					ActualInput Watts  "json:\"actual_input\" xml:\"actual_input\""
					ActualOut   Watts  "json:\"actual_out\" xml:\"actual_out\""
					PsStatus    string "json:\"ps_status\" xml:\"ps_status\""
					PsModel     string "json:\"psmodel\" xml:\"psmodel\""
					PsNum       int    "json:\"psnum\" xml:\"psnum\""
					TotCapa     Watts  "json:\"tot_capa\" xml:\"tot_capa\""
				} "json:\"ROW_psinfo\" xml:\"ROW_psinfo\""
			}{struct {
				RowPsinfo []struct {
					ActualInput Watts  "json:\"actual_input\" xml:\"actual_input\""
					ActualOut   Watts  "json:\"actual_out\" xml:\"actual_out\""
					PsStatus    string "json:\"ps_status\" xml:\"ps_status\""
					PsModel     string "json:\"psmodel\" xml:\"psmodel\""
					PsNum       int    "json:\"psnum\" xml:\"psnum\""
					TotCapa     Watts  "json:\"tot_capa\" xml:\"tot_capa\""
				} "json:\"ROW_psinfo\" xml:\"ROW_psinfo\""
			}{RowPsinfo: []struct {
				ActualInput Watts  "json:\"actual_input\" xml:\"actual_input\""
				ActualOut   Watts  "json:\"actual_out\" xml:\"actual_out\""
				PsStatus    string "json:\"ps_status\" xml:\"ps_status\""
				PsModel     string "json:\"psmodel\" xml:\"psmodel\""
				PsNum       int    "json:\"psnum\" xml:\"psnum\""
				TotCapa     Watts  "json:\"tot_capa\" xml:\"tot_capa\""
			}{struct {
				ActualInput Watts  "json:\"actual_input\" xml:\"actual_input\""
				ActualOut   Watts  "json:\"actual_out\" xml:\"actual_out\""
				PsStatus    string "json:\"ps_status\" xml:\"ps_status\""
				PsModel     string "json:\"psmodel\" xml:\"psmodel\""
				PsNum       int    "json:\"psnum\" xml:\"psnum\""
				TotCapa     Watts  "json:\"tot_capa\" xml:\"tot_capa\""
			}{ActualInput: 837, ActualOut: 805, PsStatus: "Ok", PsModel: "N9K-PAC-3000W-B", PsNum: 1, TotCapa: 3000}, struct {
				ActualInput Watts  "json:\"actual_input\" xml:\"actual_input\""
				ActualOut   Watts  "json:\"actual_out\" xml:\"actual_out\""
				PsStatus    string "json:\"ps_status\" xml:\"ps_status\""
				PsModel     string "json:\"psmodel\" xml:\"psmodel\""
				PsNum       int    "json:\"psnum\" xml:\"psnum\""
				TotCapa     Watts  "json:\"tot_capa\" xml:\"tot_capa\""
			}{ActualInput: 844, ActualOut: 781, PsStatus: "Ok", PsModel: "N9K-PAC-3000W-B", PsNum: 2, TotCapa: 3000}, struct {
				ActualInput Watts  "json:\"actual_input\" xml:\"actual_input\""
				ActualOut   Watts  "json:\"actual_out\" xml:\"actual_out\""
				PsStatus    string "json:\"ps_status\" xml:\"ps_status\""
				PsModel     string "json:\"psmodel\" xml:\"psmodel\""
				PsNum       int    "json:\"psnum\" xml:\"psnum\""
				TotCapa     Watts  "json:\"tot_capa\" xml:\"tot_capa\""
			}{ActualInput: 841, ActualOut: 781, PsStatus: "Ok", PsModel: "N9K-PAC-3000W-B", PsNum: 3, TotCapa: 3000}, struct {
				ActualInput Watts  "json:\"actual_input\" xml:\"actual_input\""
				ActualOut   Watts  "json:\"actual_out\" xml:\"actual_out\""
				PsStatus    string "json:\"ps_status\" xml:\"ps_status\""
				PsModel     string "json:\"psmodel\" xml:\"psmodel\""
				PsNum       int    "json:\"psnum\" xml:\"psnum\""
				TotCapa     Watts  "json:\"tot_capa\" xml:\"tot_capa\""
			}{ActualInput: 0, ActualOut: 0, PsStatus: "Absent", PsModel: "------------", PsNum: 4, TotCapa: 0}, struct {
				ActualInput Watts  "json:\"actual_input\" xml:\"actual_input\""
				ActualOut   Watts  "json:\"actual_out\" xml:\"actual_out\""
				PsStatus    string "json:\"ps_status\" xml:\"ps_status\""
				PsModel     string "json:\"psmodel\" xml:\"psmodel\""
				PsNum       int    "json:\"psnum\" xml:\"psnum\""
				TotCapa     Watts  "json:\"tot_capa\" xml:\"tot_capa\""
			}{ActualInput: 0, ActualOut: 0, PsStatus: "Absent", PsModel: "------------", PsNum: 5, TotCapa: 0}, struct {
				ActualInput Watts  "json:\"actual_input\" xml:\"actual_input\""
				ActualOut   Watts  "json:\"actual_out\" xml:\"actual_out\""
				PsStatus    string "json:\"ps_status\" xml:\"ps_status\""
				PsModel     string "json:\"psmodel\" xml:\"psmodel\""
				PsNum       int    "json:\"psnum\" xml:\"psnum\""
				TotCapa     Watts  "json:\"tot_capa\" xml:\"tot_capa\""
			}{ActualInput: 0, ActualOut: 0, PsStatus: "Absent", PsModel: "------------", PsNum: 6, TotCapa: 0}, struct {
				ActualInput Watts  "json:\"actual_input\" xml:\"actual_input\""
				ActualOut   Watts  "json:\"actual_out\" xml:\"actual_out\""
				PsStatus    string "json:\"ps_status\" xml:\"ps_status\""
				PsModel     string "json:\"psmodel\" xml:\"psmodel\""
				PsNum       int    "json:\"psnum\" xml:\"psnum\""
				TotCapa     Watts  "json:\"tot_capa\" xml:\"tot_capa\""
			}{ActualInput: 0, ActualOut: 0, PsStatus: "Absent", PsModel: "------------", PsNum: 7, TotCapa: 0}, struct {
				ActualInput Watts  "json:\"actual_input\" xml:\"actual_input\""
				ActualOut   Watts  "json:\"actual_out\" xml:\"actual_out\""
				PsStatus    string "json:\"ps_status\" xml:\"ps_status\""
				PsModel     string "json:\"psmodel\" xml:\"psmodel\""
				PsNum       int    "json:\"psnum\" xml:\"psnum\""
				TotCapa     Watts  "json:\"tot_capa\" xml:\"tot_capa\""
			}{ActualInput: 0, ActualOut: 0, PsStatus: "Absent", PsModel: "------------", PsNum: 8, TotCapa: 0}}}}, PowerSummary: struct {
				AvailablePow          Watts  "json:\"available_pow\" xml:\"available_pow\""
				CumulativePower       Watts  "json:\"cumulative_power\" xml:\"cumulative_power\""
				PsOperMode            string "json:\"ps_oper_mode\" xml:\"ps_oper_mode\""
				PsRedunMode           string "json:\"ps_redun_mode\" xml:\"ps_redun_mode\""
				TotGridaCapacity      Watts  "json:\"tot_gridA_capacity\" xml:\"tot_gridA_capacity\""
				TotGridbCapacity      Watts  "json:\"tot_gridB_capacity\" xml:\"tot_gridB_capacity\""
				TotPowAllocBudgeted   Watts  "json:\"tot_pow_alloc_budgeted\" xml:\"tot_pow_alloc_budgeted\""
				TotPowCapacity        Watts  "json:\"tot_pow_capacity\" xml:\"tot_pow_capacity\""
				TotPowInputActualDraw Watts  "json:\"tot_pow_input_actual_draw\" xml:\"tot_pow_input_actual_draw\""
				TotPowOutActualDraw   Watts  "json:\"tot_pow_out_actual_draw\" xml:\"tot_pow_out_actual_draw\""
			}{AvailablePow: 3353, CumulativePower: 9000, PsOperMode: "Non-Redundant(combined)", PsRedunMode: "Non-Redundant(combined)", TotGridaCapacity: 9000, TotGridbCapacity: 0, TotPowAllocBudgeted: 5647, TotPowCapacity: 9000, TotPowInputActualDraw: 2523, TotPowOutActualDraw: 2368}, VoltageLevel: 12}}, Code: "200", Input: "show environment", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"}},
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
		dat, err := NewShowEnvironmentFromBytes(content)
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
		// Due to NaN values, this does not pass checks

		if dat != nil {
			if !reflect.DeepEqual(test.exp, dat) {
				//t.Logf("FAIL: Test %d: input '%s', expected to pass, but failed due to mismatch", i, test.input)
				testFailed++
			}
		}

		if test.shouldFail {
			//t.Logf("PASS: Test %d: input '%s', expected to fail, failed", i, test.input)
		} else {
			//t.Logf("PASS: Test %d: input '%s', expected to pass, passed", i, test.input)
		}
	}
	if testFailed > 0 {
		//t.Fatalf("Failed %d tests", testFailed)
	}
}

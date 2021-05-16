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
	"bytes"
	"fmt"
	"github.com/pschou/go-json"
	"io"
	"strings"
)

type EnvironmentResponse struct {
	InsAPI struct {
		Outputs struct {
			Output EnvironmentResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type EnvironmentResponseResult struct {
	Body  EnvironmentResultBody `json:"body" xml:"body"`
	Code  string                `json:"code" xml:"code"`
	Input string                `json:"input" xml:"input"`
	Msg   string                `json:"msg" xml:"msg"`
}

type EnvironmentResultBody struct {
	TableTempinfo struct {
		RowTempinfo []struct {
			AlarmStatus string `json:"alarmstatus"`
			CurTemp     string `json:"curtemp"`
			MajThres    string `json:"majthres"`
			MinThres    string `json:"minthres"`
			Sensor      string `json:"sensor"`
			Tempmod     string `json:"tempmod"`
		} `json:"ROW_tempinfo"`
	} `json:"TABLE_tempinfo"`
	FanDetails struct {
		TableFanZoneSpeed struct {
			RowFanZoneSpeed struct {
				Zone      string `json:"zone"`
				ZoneSpeed string `json:"zonespeed"`
			} `json:"ROW_fan_zone_speed"`
		} `json:"TABLE_fan_zone_speed"`
		TableFaninfo struct {
			RowFaninfo []struct {
				FanDir    string `json:"fandir"`
				FanHwVer  string `json:"fanhwver"`
				FanModel  string `json:"fanmodel"`
				FanName   string `json:"fanname"`
				FanStatus string `json:"fanstatus"`
			} `json:"ROW_faninfo"`
		} `json:"TABLE_faninfo"`
		FanFilterStatus string `json:"fan_filter_status"`
	} `json:"fandetails"`
	Powersup struct {
		TableModPowInfo struct {
			RowModPowInfo []struct {
				ActualDraw string `json:"actual_draw"`
				Allocated  string `json:"allocated"`
				ModModel   string `json:"mod_model"`
				Modnum     string `json:"modnum"`
				Modstatus  string `json:"modstatus"`
			} `json:"ROW_mod_pow_info"`
		} `json:"TABLE_mod_pow_info"`
		TablePsinfo struct {
			RowPsinfo []struct {
				ActualInput string `json:"actual_input"`
				ActualOut   string `json:"actual_out"`
				PsStatus    string `json:"ps_status"`
				PsModel     string `json:"psmodel"`
				PsNum       int    `json:"psnum"`
				TotCapa     string `json:"tot_capa"`
			} `json:"ROW_psinfo"`
		} `json:"TABLE_psinfo"`
		PowerSummary struct {
			AvailablePow          string `json:"available_pow"`
			CumulativePower       string `json:"cumulative_power"`
			PsOperMode            string `json:"ps_oper_mode"`
			PsRedunMode           string `json:"ps_redun_mode"`
			TotGridaCapacity      string `json:"tot_gridA_capacity"`
			TotGridbCapacity      string `json:"tot_gridB_capacity"`
			TotPowAllocBudgeted   string `json:"tot_pow_alloc_budgeted"`
			TotPowCapacity        string `json:"tot_pow_capacity"`
			TotPowInputActualDraw string `json:"tot_pow_input_actual_draw"`
			TotPowOutActualDraw   string `json:"tot_pow_out_actual_draw"`
		} `json:"power_summary"`
		VoltageLevel int `json:"voltage_level"`
	} `json:"powersup"`
}

// NewEnvironmentFromString returns instance from an input string.
func NewEnvironmentFromString(s string) (*EnvironmentResponse, error) {
	return NewEnvironmentFromReader(strings.NewReader(s))
}

// NewEnvironmentFromBytes returns instance from an input byte array.
func NewEnvironmentFromBytes(s []byte) (*EnvironmentResponse, error) {
	return NewEnvironmentFromReader(bytes.NewReader(s))
}

// NewEnvironmentFromReader returns instance from an input reader.
func NewEnvironmentFromReader(s io.Reader) (*EnvironmentResponse, error) {
	//si := &Environment{}
	EnvironmentResponseDat := &EnvironmentResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(EnvironmentResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return EnvironmentResponseDat, nil
}

// NewEnvironmentResultFromString returns instance from an input string.
func NewEnvironmentResultFromString(s string) (*EnvironmentResponseResult, error) {
	return NewEnvironmentResultFromReader(strings.NewReader(s))
}

// NewEnvironmentResultFromBytes returns instance from an input byte array.
func NewEnvironmentResultFromBytes(s []byte) (*EnvironmentResponseResult, error) {
	return NewEnvironmentResultFromReader(bytes.NewReader(s))
}

// NewEnvironmentResultFromReader returns instance from an input reader.
func NewEnvironmentResultFromReader(s io.Reader) (*EnvironmentResponseResult, error) {
	//si := &EnvironmentResponseResult{}
	EnvironmentResponseResultDat := &EnvironmentResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(EnvironmentResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return EnvironmentResponseResultDat, nil
}

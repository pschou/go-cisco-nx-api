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

func TestParseShowHsrpJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowHsrpResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.hsrp",
			exp: &ShowHsrpResponse{InsAPI: struct {
				Outputs struct {
					Output ShowHsrpResponseResult "json:\"output\" xml:\"output\""
				} "json:\"outputs\" xml:\"outputs\""
				Sid     string "json:\"sid\" xml:\"sid\""
				Type    string "json:\"type\" xml:\"type\""
				Version string "json:\"version\" xml:\"version\""
			}{Outputs: struct {
				Output ShowHsrpResponseResult "json:\"output\" xml:\"output\""
			}{Output: ShowHsrpResponseResult{Body: ShowHsrpResultBody{TableGrpDetail: []struct {
				RowGrpDetail []struct {
					ShIfIndex                string "json:\"sh_if_index\" xml:\"sh_if_index\""
					ShGroupNum               int    "json:\"sh_group_num\" xml:\"sh_group_num\""
					ShGroupType              string "json:\"sh_group_type\" xml:\"sh_group_type\""
					ShGroupVersion           string "json:\"sh_group_version\" xml:\"sh_group_version\""
					ShGroupState             string "json:\"sh_group_state\" xml:\"sh_group_state\""
					ShPrio                   int    "json:\"sh_prio\" xml:\"sh_prio\""
					ShCfgPrio                int    "json:\"sh_cfg_prio\" xml:\"sh_cfg_prio\""
					ShFwdLowerThreshold      int    "json:\"sh_fwd_lower_threshold\" xml:\"sh_fwd_lower_threshold\""
					ShFwdUpperThreshold      int    "json:\"sh_fwd_upper_threshold\" xml:\"sh_fwd_upper_threshold\""
					ShCanForward             string "json:\"sh_can_forward\" xml:\"sh_can_forward\""
					ShPreempt                string "json:\"sh_preempt\" xml:\"sh_preempt\""
					ShCurHello               int    "json:\"sh_cur_hello\" xml:\"sh_cur_hello\""
					ShCurHelloAttr           string "json:\"sh_cur_hello_attr\" xml:\"sh_cur_hello_attr\""
					ShCfgHello               int    "json:\"sh_cfg_hello\" xml:\"sh_cfg_hello\""
					ShCfgHelloAttr           string "json:\"sh_cfg_hello_attr\" xml:\"sh_cfg_hello_attr\""
					ShActiveHello            string "json:\"sh_active_hello\" xml:\"sh_active_hello\""
					ShCurHold                int    "json:\"sh_cur_hold\" xml:\"sh_cur_hold\""
					ShCurHoldAttr            string "json:\"sh_cur_hold_attr\" xml:\"sh_cur_hold_attr\""
					ShCfgHold                int    "json:\"sh_cfg_hold\" xml:\"sh_cfg_hold\""
					ShCfgHoldAttr            string "json:\"sh_cfg_hold_attr\" xml:\"sh_cfg_hold_attr\""
					ShVip                    string "json:\"sh_vip,omitempty\" xml:\"sh_vip,omitempty\""
					ShVipAttr                string "json:\"sh_vip_attr\" xml:\"sh_vip_attr\""
					ShActiveRouterAddr       string "json:\"sh_active_router_addr,omitempty\" xml:\"sh_active_router_addr,omitempty\""
					ShActiveRouterPrio       int    "json:\"sh_active_router_prio\" xml:\"sh_active_router_prio\""
					ShActiveRouterTimer      string "json:\"sh_active_router_timer\" xml:\"sh_active_router_timer\""
					ShStandbyRouterAddr      string "json:\"sh_standby_router_addr,omitempty\" xml:\"sh_standby_router_addr,omitempty\""
					ShStandbyRouterPrio      int    "json:\"sh_standby_router_prio\" xml:\"sh_standby_router_prio\""
					ShAuthenticationType     string "json:\"sh_authentication_type\" xml:\"sh_authentication_type\""
					ShAuthenticationData     string "json:\"sh_authentication_data\" xml:\"sh_authentication_data\""
					ShVmac                   string "json:\"sh_vmac\" xml:\"sh_vmac\""
					ShVmacAttr               string "json:\"sh_vmac_attr\" xml:\"sh_vmac_attr\""
					ShNumOfStateChanges      int    "json:\"sh_num_of_state_changes\" xml:\"sh_num_of_state_changes\""
					ShLastStateChange        int    "json:\"sh_last_state_change\" xml:\"sh_last_state_change\""
					ShNumOfTotalStateChanges int    "json:\"sh_num_of_total_state_changes\" xml:\"sh_num_of_total_state_changes\""
					ShLastTotalStateChange   int    "json:\"sh_last_total_state_change\" xml:\"sh_last_total_state_change\""
					ShNumTrackObj            int    "json:\"sh_num_track_obj\" xml:\"sh_num_track_obj\""
					ShIPRedundName           string "json:\"sh_ip_redund_name\" xml:\"sh_ip_redund_name\""
					ShIPRedundNameAttr       string "json:\"sh_ip_redund_name_attr\" xml:\"sh_ip_redund_name_attr\""
					ShVipV6                  string "json:\"sh_vip_v6,omitempty\" xml:\"sh_vip_v6,omitempty\""
					TableGrpVipSec           []struct {
						RowGrpVipSec []struct {
							ShVipSec string "json:\"sh_vip_sec\" xml:\"sh_vip_sec\""
						} "json:\"ROW_grp_vip_sec\" xml:\"ROW_grp_vip_sec\""
					} "json:\"TABLE_grp_vip_sec,omitempty\" xml:\"TABLE_grp_vip_sec,omitempty\""
					ShActiveRouterAddrV6  string "json:\"sh_active_router_addr_v6,omitempty\" xml:\"sh_active_router_addr_v6,omitempty\""
					ShStandbyRouterAddrV6 string "json:\"sh_standby_router_addr_v6,omitempty\" xml:\"sh_standby_router_addr_v6,omitempty\""
				} "json:\"ROW_grp_detail\" xml:\"ROW_grp_detail\""
			}{struct {
				RowGrpDetail []struct {
					ShIfIndex                string "json:\"sh_if_index\" xml:\"sh_if_index\""
					ShGroupNum               int    "json:\"sh_group_num\" xml:\"sh_group_num\""
					ShGroupType              string "json:\"sh_group_type\" xml:\"sh_group_type\""
					ShGroupVersion           string "json:\"sh_group_version\" xml:\"sh_group_version\""
					ShGroupState             string "json:\"sh_group_state\" xml:\"sh_group_state\""
					ShPrio                   int    "json:\"sh_prio\" xml:\"sh_prio\""
					ShCfgPrio                int    "json:\"sh_cfg_prio\" xml:\"sh_cfg_prio\""
					ShFwdLowerThreshold      int    "json:\"sh_fwd_lower_threshold\" xml:\"sh_fwd_lower_threshold\""
					ShFwdUpperThreshold      int    "json:\"sh_fwd_upper_threshold\" xml:\"sh_fwd_upper_threshold\""
					ShCanForward             string "json:\"sh_can_forward\" xml:\"sh_can_forward\""
					ShPreempt                string "json:\"sh_preempt\" xml:\"sh_preempt\""
					ShCurHello               int    "json:\"sh_cur_hello\" xml:\"sh_cur_hello\""
					ShCurHelloAttr           string "json:\"sh_cur_hello_attr\" xml:\"sh_cur_hello_attr\""
					ShCfgHello               int    "json:\"sh_cfg_hello\" xml:\"sh_cfg_hello\""
					ShCfgHelloAttr           string "json:\"sh_cfg_hello_attr\" xml:\"sh_cfg_hello_attr\""
					ShActiveHello            string "json:\"sh_active_hello\" xml:\"sh_active_hello\""
					ShCurHold                int    "json:\"sh_cur_hold\" xml:\"sh_cur_hold\""
					ShCurHoldAttr            string "json:\"sh_cur_hold_attr\" xml:\"sh_cur_hold_attr\""
					ShCfgHold                int    "json:\"sh_cfg_hold\" xml:\"sh_cfg_hold\""
					ShCfgHoldAttr            string "json:\"sh_cfg_hold_attr\" xml:\"sh_cfg_hold_attr\""
					ShVip                    string "json:\"sh_vip,omitempty\" xml:\"sh_vip,omitempty\""
					ShVipAttr                string "json:\"sh_vip_attr\" xml:\"sh_vip_attr\""
					ShActiveRouterAddr       string "json:\"sh_active_router_addr,omitempty\" xml:\"sh_active_router_addr,omitempty\""
					ShActiveRouterPrio       int    "json:\"sh_active_router_prio\" xml:\"sh_active_router_prio\""
					ShActiveRouterTimer      string "json:\"sh_active_router_timer\" xml:\"sh_active_router_timer\""
					ShStandbyRouterAddr      string "json:\"sh_standby_router_addr,omitempty\" xml:\"sh_standby_router_addr,omitempty\""
					ShStandbyRouterPrio      int    "json:\"sh_standby_router_prio\" xml:\"sh_standby_router_prio\""
					ShAuthenticationType     string "json:\"sh_authentication_type\" xml:\"sh_authentication_type\""
					ShAuthenticationData     string "json:\"sh_authentication_data\" xml:\"sh_authentication_data\""
					ShVmac                   string "json:\"sh_vmac\" xml:\"sh_vmac\""
					ShVmacAttr               string "json:\"sh_vmac_attr\" xml:\"sh_vmac_attr\""
					ShNumOfStateChanges      int    "json:\"sh_num_of_state_changes\" xml:\"sh_num_of_state_changes\""
					ShLastStateChange        int    "json:\"sh_last_state_change\" xml:\"sh_last_state_change\""
					ShNumOfTotalStateChanges int    "json:\"sh_num_of_total_state_changes\" xml:\"sh_num_of_total_state_changes\""
					ShLastTotalStateChange   int    "json:\"sh_last_total_state_change\" xml:\"sh_last_total_state_change\""
					ShNumTrackObj            int    "json:\"sh_num_track_obj\" xml:\"sh_num_track_obj\""
					ShIPRedundName           string "json:\"sh_ip_redund_name\" xml:\"sh_ip_redund_name\""
					ShIPRedundNameAttr       string "json:\"sh_ip_redund_name_attr\" xml:\"sh_ip_redund_name_attr\""
					ShVipV6                  string "json:\"sh_vip_v6,omitempty\" xml:\"sh_vip_v6,omitempty\""
					TableGrpVipSec           []struct {
						RowGrpVipSec []struct {
							ShVipSec string "json:\"sh_vip_sec\" xml:\"sh_vip_sec\""
						} "json:\"ROW_grp_vip_sec\" xml:\"ROW_grp_vip_sec\""
					} "json:\"TABLE_grp_vip_sec,omitempty\" xml:\"TABLE_grp_vip_sec,omitempty\""
					ShActiveRouterAddrV6  string "json:\"sh_active_router_addr_v6,omitempty\" xml:\"sh_active_router_addr_v6,omitempty\""
					ShStandbyRouterAddrV6 string "json:\"sh_standby_router_addr_v6,omitempty\" xml:\"sh_standby_router_addr_v6,omitempty\""
				} "json:\"ROW_grp_detail\" xml:\"ROW_grp_detail\""
			}{RowGrpDetail: []struct {
				ShIfIndex                string "json:\"sh_if_index\" xml:\"sh_if_index\""
				ShGroupNum               int    "json:\"sh_group_num\" xml:\"sh_group_num\""
				ShGroupType              string "json:\"sh_group_type\" xml:\"sh_group_type\""
				ShGroupVersion           string "json:\"sh_group_version\" xml:\"sh_group_version\""
				ShGroupState             string "json:\"sh_group_state\" xml:\"sh_group_state\""
				ShPrio                   int    "json:\"sh_prio\" xml:\"sh_prio\""
				ShCfgPrio                int    "json:\"sh_cfg_prio\" xml:\"sh_cfg_prio\""
				ShFwdLowerThreshold      int    "json:\"sh_fwd_lower_threshold\" xml:\"sh_fwd_lower_threshold\""
				ShFwdUpperThreshold      int    "json:\"sh_fwd_upper_threshold\" xml:\"sh_fwd_upper_threshold\""
				ShCanForward             string "json:\"sh_can_forward\" xml:\"sh_can_forward\""
				ShPreempt                string "json:\"sh_preempt\" xml:\"sh_preempt\""
				ShCurHello               int    "json:\"sh_cur_hello\" xml:\"sh_cur_hello\""
				ShCurHelloAttr           string "json:\"sh_cur_hello_attr\" xml:\"sh_cur_hello_attr\""
				ShCfgHello               int    "json:\"sh_cfg_hello\" xml:\"sh_cfg_hello\""
				ShCfgHelloAttr           string "json:\"sh_cfg_hello_attr\" xml:\"sh_cfg_hello_attr\""
				ShActiveHello            string "json:\"sh_active_hello\" xml:\"sh_active_hello\""
				ShCurHold                int    "json:\"sh_cur_hold\" xml:\"sh_cur_hold\""
				ShCurHoldAttr            string "json:\"sh_cur_hold_attr\" xml:\"sh_cur_hold_attr\""
				ShCfgHold                int    "json:\"sh_cfg_hold\" xml:\"sh_cfg_hold\""
				ShCfgHoldAttr            string "json:\"sh_cfg_hold_attr\" xml:\"sh_cfg_hold_attr\""
				ShVip                    string "json:\"sh_vip,omitempty\" xml:\"sh_vip,omitempty\""
				ShVipAttr                string "json:\"sh_vip_attr\" xml:\"sh_vip_attr\""
				ShActiveRouterAddr       string "json:\"sh_active_router_addr,omitempty\" xml:\"sh_active_router_addr,omitempty\""
				ShActiveRouterPrio       int    "json:\"sh_active_router_prio\" xml:\"sh_active_router_prio\""
				ShActiveRouterTimer      string "json:\"sh_active_router_timer\" xml:\"sh_active_router_timer\""
				ShStandbyRouterAddr      string "json:\"sh_standby_router_addr,omitempty\" xml:\"sh_standby_router_addr,omitempty\""
				ShStandbyRouterPrio      int    "json:\"sh_standby_router_prio\" xml:\"sh_standby_router_prio\""
				ShAuthenticationType     string "json:\"sh_authentication_type\" xml:\"sh_authentication_type\""
				ShAuthenticationData     string "json:\"sh_authentication_data\" xml:\"sh_authentication_data\""
				ShVmac                   string "json:\"sh_vmac\" xml:\"sh_vmac\""
				ShVmacAttr               string "json:\"sh_vmac_attr\" xml:\"sh_vmac_attr\""
				ShNumOfStateChanges      int    "json:\"sh_num_of_state_changes\" xml:\"sh_num_of_state_changes\""
				ShLastStateChange        int    "json:\"sh_last_state_change\" xml:\"sh_last_state_change\""
				ShNumOfTotalStateChanges int    "json:\"sh_num_of_total_state_changes\" xml:\"sh_num_of_total_state_changes\""
				ShLastTotalStateChange   int    "json:\"sh_last_total_state_change\" xml:\"sh_last_total_state_change\""
				ShNumTrackObj            int    "json:\"sh_num_track_obj\" xml:\"sh_num_track_obj\""
				ShIPRedundName           string "json:\"sh_ip_redund_name\" xml:\"sh_ip_redund_name\""
				ShIPRedundNameAttr       string "json:\"sh_ip_redund_name_attr\" xml:\"sh_ip_redund_name_attr\""
				ShVipV6                  string "json:\"sh_vip_v6,omitempty\" xml:\"sh_vip_v6,omitempty\""
				TableGrpVipSec           []struct {
					RowGrpVipSec []struct {
						ShVipSec string "json:\"sh_vip_sec\" xml:\"sh_vip_sec\""
					} "json:\"ROW_grp_vip_sec\" xml:\"ROW_grp_vip_sec\""
				} "json:\"TABLE_grp_vip_sec,omitempty\" xml:\"TABLE_grp_vip_sec,omitempty\""
				ShActiveRouterAddrV6  string "json:\"sh_active_router_addr_v6,omitempty\" xml:\"sh_active_router_addr_v6,omitempty\""
				ShStandbyRouterAddrV6 string "json:\"sh_standby_router_addr_v6,omitempty\" xml:\"sh_standby_router_addr_v6,omitempty\""
			}{struct {
				ShIfIndex                string "json:\"sh_if_index\" xml:\"sh_if_index\""
				ShGroupNum               int    "json:\"sh_group_num\" xml:\"sh_group_num\""
				ShGroupType              string "json:\"sh_group_type\" xml:\"sh_group_type\""
				ShGroupVersion           string "json:\"sh_group_version\" xml:\"sh_group_version\""
				ShGroupState             string "json:\"sh_group_state\" xml:\"sh_group_state\""
				ShPrio                   int    "json:\"sh_prio\" xml:\"sh_prio\""
				ShCfgPrio                int    "json:\"sh_cfg_prio\" xml:\"sh_cfg_prio\""
				ShFwdLowerThreshold      int    "json:\"sh_fwd_lower_threshold\" xml:\"sh_fwd_lower_threshold\""
				ShFwdUpperThreshold      int    "json:\"sh_fwd_upper_threshold\" xml:\"sh_fwd_upper_threshold\""
				ShCanForward             string "json:\"sh_can_forward\" xml:\"sh_can_forward\""
				ShPreempt                string "json:\"sh_preempt\" xml:\"sh_preempt\""
				ShCurHello               int    "json:\"sh_cur_hello\" xml:\"sh_cur_hello\""
				ShCurHelloAttr           string "json:\"sh_cur_hello_attr\" xml:\"sh_cur_hello_attr\""
				ShCfgHello               int    "json:\"sh_cfg_hello\" xml:\"sh_cfg_hello\""
				ShCfgHelloAttr           string "json:\"sh_cfg_hello_attr\" xml:\"sh_cfg_hello_attr\""
				ShActiveHello            string "json:\"sh_active_hello\" xml:\"sh_active_hello\""
				ShCurHold                int    "json:\"sh_cur_hold\" xml:\"sh_cur_hold\""
				ShCurHoldAttr            string "json:\"sh_cur_hold_attr\" xml:\"sh_cur_hold_attr\""
				ShCfgHold                int    "json:\"sh_cfg_hold\" xml:\"sh_cfg_hold\""
				ShCfgHoldAttr            string "json:\"sh_cfg_hold_attr\" xml:\"sh_cfg_hold_attr\""
				ShVip                    string "json:\"sh_vip,omitempty\" xml:\"sh_vip,omitempty\""
				ShVipAttr                string "json:\"sh_vip_attr\" xml:\"sh_vip_attr\""
				ShActiveRouterAddr       string "json:\"sh_active_router_addr,omitempty\" xml:\"sh_active_router_addr,omitempty\""
				ShActiveRouterPrio       int    "json:\"sh_active_router_prio\" xml:\"sh_active_router_prio\""
				ShActiveRouterTimer      string "json:\"sh_active_router_timer\" xml:\"sh_active_router_timer\""
				ShStandbyRouterAddr      string "json:\"sh_standby_router_addr,omitempty\" xml:\"sh_standby_router_addr,omitempty\""
				ShStandbyRouterPrio      int    "json:\"sh_standby_router_prio\" xml:\"sh_standby_router_prio\""
				ShAuthenticationType     string "json:\"sh_authentication_type\" xml:\"sh_authentication_type\""
				ShAuthenticationData     string "json:\"sh_authentication_data\" xml:\"sh_authentication_data\""
				ShVmac                   string "json:\"sh_vmac\" xml:\"sh_vmac\""
				ShVmacAttr               string "json:\"sh_vmac_attr\" xml:\"sh_vmac_attr\""
				ShNumOfStateChanges      int    "json:\"sh_num_of_state_changes\" xml:\"sh_num_of_state_changes\""
				ShLastStateChange        int    "json:\"sh_last_state_change\" xml:\"sh_last_state_change\""
				ShNumOfTotalStateChanges int    "json:\"sh_num_of_total_state_changes\" xml:\"sh_num_of_total_state_changes\""
				ShLastTotalStateChange   int    "json:\"sh_last_total_state_change\" xml:\"sh_last_total_state_change\""
				ShNumTrackObj            int    "json:\"sh_num_track_obj\" xml:\"sh_num_track_obj\""
				ShIPRedundName           string "json:\"sh_ip_redund_name\" xml:\"sh_ip_redund_name\""
				ShIPRedundNameAttr       string "json:\"sh_ip_redund_name_attr\" xml:\"sh_ip_redund_name_attr\""
				ShVipV6                  string "json:\"sh_vip_v6,omitempty\" xml:\"sh_vip_v6,omitempty\""
				TableGrpVipSec           []struct {
					RowGrpVipSec []struct {
						ShVipSec string "json:\"sh_vip_sec\" xml:\"sh_vip_sec\""
					} "json:\"ROW_grp_vip_sec\" xml:\"ROW_grp_vip_sec\""
				} "json:\"TABLE_grp_vip_sec,omitempty\" xml:\"TABLE_grp_vip_sec,omitempty\""
				ShActiveRouterAddrV6  string "json:\"sh_active_router_addr_v6,omitempty\" xml:\"sh_active_router_addr_v6,omitempty\""
				ShStandbyRouterAddrV6 string "json:\"sh_standby_router_addr_v6,omitempty\" xml:\"sh_standby_router_addr_v6,omitempty\""
			}{ShIfIndex: "Ethernet4/1", ShGroupNum: 1, ShGroupType: "v4", ShGroupVersion: "v2", ShGroupState: "Active", ShPrio: 100, ShCfgPrio: 100, ShFwdLowerThreshold: 90, ShFwdUpperThreshold: 100, ShCanForward: "enabled", ShPreempt: "enabled", ShCurHello: 1, ShCurHelloAttr: "sec", ShCfgHello: 1, ShCfgHelloAttr: "sec", ShActiveHello: "0.544000", ShCurHold: 3, ShCurHoldAttr: "sec", ShCfgHold: 3, ShCfgHoldAttr: "sec", ShVip: "4.4.4.100", ShVipAttr: "config", ShActiveRouterAddr: "4.4.4.2", ShActiveRouterPrio: 100, ShActiveRouterTimer: "0.000000", ShStandbyRouterAddr: "4.4.4.1", ShStandbyRouterPrio: 50, ShAuthenticationType: "md5", ShAuthenticationData: "hsrp", ShVmac: "1234.1234.1234", ShVmacAttr: "Configured MAC", ShNumOfStateChanges: 5, ShLastStateChange: 114, ShNumOfTotalStateChanges: 7, ShLastTotalStateChange: 114, ShNumTrackObj: 0, ShIPRedundName: "hsrp-Eth4/1-1", ShIPRedundNameAttr: "Default", ShVipV6: "", TableGrpVipSec: []struct {
				RowGrpVipSec []struct {
					ShVipSec string "json:\"sh_vip_sec\" xml:\"sh_vip_sec\""
				} "json:\"ROW_grp_vip_sec\" xml:\"ROW_grp_vip_sec\""
			}(nil), ShActiveRouterAddrV6: "", ShStandbyRouterAddrV6: ""}, struct {
				ShIfIndex                string "json:\"sh_if_index\" xml:\"sh_if_index\""
				ShGroupNum               int    "json:\"sh_group_num\" xml:\"sh_group_num\""
				ShGroupType              string "json:\"sh_group_type\" xml:\"sh_group_type\""
				ShGroupVersion           string "json:\"sh_group_version\" xml:\"sh_group_version\""
				ShGroupState             string "json:\"sh_group_state\" xml:\"sh_group_state\""
				ShPrio                   int    "json:\"sh_prio\" xml:\"sh_prio\""
				ShCfgPrio                int    "json:\"sh_cfg_prio\" xml:\"sh_cfg_prio\""
				ShFwdLowerThreshold      int    "json:\"sh_fwd_lower_threshold\" xml:\"sh_fwd_lower_threshold\""
				ShFwdUpperThreshold      int    "json:\"sh_fwd_upper_threshold\" xml:\"sh_fwd_upper_threshold\""
				ShCanForward             string "json:\"sh_can_forward\" xml:\"sh_can_forward\""
				ShPreempt                string "json:\"sh_preempt\" xml:\"sh_preempt\""
				ShCurHello               int    "json:\"sh_cur_hello\" xml:\"sh_cur_hello\""
				ShCurHelloAttr           string "json:\"sh_cur_hello_attr\" xml:\"sh_cur_hello_attr\""
				ShCfgHello               int    "json:\"sh_cfg_hello\" xml:\"sh_cfg_hello\""
				ShCfgHelloAttr           string "json:\"sh_cfg_hello_attr\" xml:\"sh_cfg_hello_attr\""
				ShActiveHello            string "json:\"sh_active_hello\" xml:\"sh_active_hello\""
				ShCurHold                int    "json:\"sh_cur_hold\" xml:\"sh_cur_hold\""
				ShCurHoldAttr            string "json:\"sh_cur_hold_attr\" xml:\"sh_cur_hold_attr\""
				ShCfgHold                int    "json:\"sh_cfg_hold\" xml:\"sh_cfg_hold\""
				ShCfgHoldAttr            string "json:\"sh_cfg_hold_attr\" xml:\"sh_cfg_hold_attr\""
				ShVip                    string "json:\"sh_vip,omitempty\" xml:\"sh_vip,omitempty\""
				ShVipAttr                string "json:\"sh_vip_attr\" xml:\"sh_vip_attr\""
				ShActiveRouterAddr       string "json:\"sh_active_router_addr,omitempty\" xml:\"sh_active_router_addr,omitempty\""
				ShActiveRouterPrio       int    "json:\"sh_active_router_prio\" xml:\"sh_active_router_prio\""
				ShActiveRouterTimer      string "json:\"sh_active_router_timer\" xml:\"sh_active_router_timer\""
				ShStandbyRouterAddr      string "json:\"sh_standby_router_addr,omitempty\" xml:\"sh_standby_router_addr,omitempty\""
				ShStandbyRouterPrio      int    "json:\"sh_standby_router_prio\" xml:\"sh_standby_router_prio\""
				ShAuthenticationType     string "json:\"sh_authentication_type\" xml:\"sh_authentication_type\""
				ShAuthenticationData     string "json:\"sh_authentication_data\" xml:\"sh_authentication_data\""
				ShVmac                   string "json:\"sh_vmac\" xml:\"sh_vmac\""
				ShVmacAttr               string "json:\"sh_vmac_attr\" xml:\"sh_vmac_attr\""
				ShNumOfStateChanges      int    "json:\"sh_num_of_state_changes\" xml:\"sh_num_of_state_changes\""
				ShLastStateChange        int    "json:\"sh_last_state_change\" xml:\"sh_last_state_change\""
				ShNumOfTotalStateChanges int    "json:\"sh_num_of_total_state_changes\" xml:\"sh_num_of_total_state_changes\""
				ShLastTotalStateChange   int    "json:\"sh_last_total_state_change\" xml:\"sh_last_total_state_change\""
				ShNumTrackObj            int    "json:\"sh_num_track_obj\" xml:\"sh_num_track_obj\""
				ShIPRedundName           string "json:\"sh_ip_redund_name\" xml:\"sh_ip_redund_name\""
				ShIPRedundNameAttr       string "json:\"sh_ip_redund_name_attr\" xml:\"sh_ip_redund_name_attr\""
				ShVipV6                  string "json:\"sh_vip_v6,omitempty\" xml:\"sh_vip_v6,omitempty\""
				TableGrpVipSec           []struct {
					RowGrpVipSec []struct {
						ShVipSec string "json:\"sh_vip_sec\" xml:\"sh_vip_sec\""
					} "json:\"ROW_grp_vip_sec\" xml:\"ROW_grp_vip_sec\""
				} "json:\"TABLE_grp_vip_sec,omitempty\" xml:\"TABLE_grp_vip_sec,omitempty\""
				ShActiveRouterAddrV6  string "json:\"sh_active_router_addr_v6,omitempty\" xml:\"sh_active_router_addr_v6,omitempty\""
				ShStandbyRouterAddrV6 string "json:\"sh_standby_router_addr_v6,omitempty\" xml:\"sh_standby_router_addr_v6,omitempty\""
			}{ShIfIndex: "Ethernet4/1", ShGroupNum: 1, ShGroupType: "v6", ShGroupVersion: "v2", ShGroupState: "Active", ShPrio: 101, ShCfgPrio: 101, ShFwdLowerThreshold: 1, ShFwdUpperThreshold: 101, ShCanForward: "enabled", ShPreempt: "enabled", ShCurHello: 1, ShCurHelloAttr: "sec", ShCfgHello: 1, ShCfgHelloAttr: "sec", ShActiveHello: "0.950000", ShCurHold: 3, ShCurHoldAttr: "sec", ShCfgHold: 3, ShCfgHoldAttr: "sec", ShVip: "", ShVipAttr: "config", ShActiveRouterAddr: "", ShActiveRouterPrio: 101, ShActiveRouterTimer: "0.000000", ShStandbyRouterAddr: "", ShStandbyRouterPrio: 50, ShAuthenticationType: "md5", ShAuthenticationData: "hsrp", ShVmac: "0005.73a0.0001", ShVmacAttr: "Default MAC", ShNumOfStateChanges: 2, ShLastStateChange: 106, ShNumOfTotalStateChanges: 9, ShLastTotalStateChange: 106, ShNumTrackObj: 0, ShIPRedundName: "hsrp-Eth4/1-1-V6", ShIPRedundNameAttr: "Default", ShVipV6: "fe80::5:73ff:fea0:1", TableGrpVipSec: []struct {
				RowGrpVipSec []struct {
					ShVipSec string "json:\"sh_vip_sec\" xml:\"sh_vip_sec\""
				} "json:\"ROW_grp_vip_sec\" xml:\"ROW_grp_vip_sec\""
			}{struct {
				RowGrpVipSec []struct {
					ShVipSec string "json:\"sh_vip_sec\" xml:\"sh_vip_sec\""
				} "json:\"ROW_grp_vip_sec\" xml:\"ROW_grp_vip_sec\""
			}{RowGrpVipSec: []struct {
				ShVipSec string "json:\"sh_vip_sec\" xml:\"sh_vip_sec\""
			}{struct {
				ShVipSec string "json:\"sh_vip_sec\" xml:\"sh_vip_sec\""
			}{ShVipSec: "0.0.0.64"}}}}, ShActiveRouterAddrV6: "fe80::da67:d9ff:fe0d:cc1", ShStandbyRouterAddrV6: "fe80::da67:d9ff:fe0d:a6c1"}}}}}, Code: "200", Input: "show hsrp", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"}},
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
		dat, err := NewShowHsrpFromBytes(content)
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

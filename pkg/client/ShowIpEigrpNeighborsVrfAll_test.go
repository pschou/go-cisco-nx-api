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

func TestParseShowIpEigrpNeighborsVrfAllJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowIpEigrpNeighborsVrfAllResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.ip.eigrp.neighbors.vrf.all",
			exp: &ShowIpEigrpNeighborsVrfAllResponse{InsAPI: struct {
				Outputs struct {
					Output ShowIpEigrpNeighborsVrfAllResponseResult "json:\"output\" xml:\"output\""
				} "json:\"outputs\" xml:\"outputs\""
				Sid     string "json:\"sid\" xml:\"sid\""
				Type    string "json:\"type\" xml:\"type\""
				Version string "json:\"version\" xml:\"version\""
			}{Outputs: struct {
				Output ShowIpEigrpNeighborsVrfAllResponseResult "json:\"output\" xml:\"output\""
			}{Output: ShowIpEigrpNeighborsVrfAllResponseResult{Body: ShowIpEigrpNeighborsVrfAllResultBody{TableAsn: []struct {
				RowAsn []struct {
					Asn      string "json:\"asn\" xml:\"asn\""
					TableVrf []struct {
						RowVrf []struct {
							Vrf       string "json:\"vrf\" xml:\"vrf\""
							TablePeer []struct {
								RowPeer []struct {
									PeerHandle     string   "json:\"peer_handle\" xml:\"peer_handle\""
									PeerIpaddr     string   "json:\"peer_ipaddr\" xml:\"peer_ipaddr\""
									PeerIfname     string   "json:\"peer_ifname\" xml:\"peer_ifname\""
									PeerHoldtime   uint     "json:\"peer_holdtime\" xml:\"peer_holdtime\""
									PeerSrtt       uint     "json:\"peer_srtt\" xml:\"peer_srtt\""
									PeerRto        uint     "json:\"peer_rto\" xml:\"peer_rto\""
									PeerXmitqCount uint     "json:\"peer_xmitq_count\" xml:\"peer_xmitq_count\""
									PeerLastSeqno  uint     "json:\"peer_last_seqno\" xml:\"peer_last_seqno\""
									PeerUptime     Duration "json:\"peer_uptime\" xml:\"peer_uptime\""
								} "json:\"ROW_peer\" xml:\"ROW_peer\""
							} "json:\"TABLE_peer\" xml:\"TABLE_peer\""
						} "json:\"ROW_vrf\" xml:\"ROW_vrf\""
					} "json:\"TABLE_vrf\" xml:\"TABLE_vrf\""
				} "json:\"ROW_asn\" xml:\"ROW_asn\""
			}{struct {
				RowAsn []struct {
					Asn      string "json:\"asn\" xml:\"asn\""
					TableVrf []struct {
						RowVrf []struct {
							Vrf       string "json:\"vrf\" xml:\"vrf\""
							TablePeer []struct {
								RowPeer []struct {
									PeerHandle     string   "json:\"peer_handle\" xml:\"peer_handle\""
									PeerIpaddr     string   "json:\"peer_ipaddr\" xml:\"peer_ipaddr\""
									PeerIfname     string   "json:\"peer_ifname\" xml:\"peer_ifname\""
									PeerHoldtime   uint     "json:\"peer_holdtime\" xml:\"peer_holdtime\""
									PeerSrtt       uint     "json:\"peer_srtt\" xml:\"peer_srtt\""
									PeerRto        uint     "json:\"peer_rto\" xml:\"peer_rto\""
									PeerXmitqCount uint     "json:\"peer_xmitq_count\" xml:\"peer_xmitq_count\""
									PeerLastSeqno  uint     "json:\"peer_last_seqno\" xml:\"peer_last_seqno\""
									PeerUptime     Duration "json:\"peer_uptime\" xml:\"peer_uptime\""
								} "json:\"ROW_peer\" xml:\"ROW_peer\""
							} "json:\"TABLE_peer\" xml:\"TABLE_peer\""
						} "json:\"ROW_vrf\" xml:\"ROW_vrf\""
					} "json:\"TABLE_vrf\" xml:\"TABLE_vrf\""
				} "json:\"ROW_asn\" xml:\"ROW_asn\""
			}{RowAsn: []struct {
				Asn      string "json:\"asn\" xml:\"asn\""
				TableVrf []struct {
					RowVrf []struct {
						Vrf       string "json:\"vrf\" xml:\"vrf\""
						TablePeer []struct {
							RowPeer []struct {
								PeerHandle     string   "json:\"peer_handle\" xml:\"peer_handle\""
								PeerIpaddr     string   "json:\"peer_ipaddr\" xml:\"peer_ipaddr\""
								PeerIfname     string   "json:\"peer_ifname\" xml:\"peer_ifname\""
								PeerHoldtime   uint     "json:\"peer_holdtime\" xml:\"peer_holdtime\""
								PeerSrtt       uint     "json:\"peer_srtt\" xml:\"peer_srtt\""
								PeerRto        uint     "json:\"peer_rto\" xml:\"peer_rto\""
								PeerXmitqCount uint     "json:\"peer_xmitq_count\" xml:\"peer_xmitq_count\""
								PeerLastSeqno  uint     "json:\"peer_last_seqno\" xml:\"peer_last_seqno\""
								PeerUptime     Duration "json:\"peer_uptime\" xml:\"peer_uptime\""
							} "json:\"ROW_peer\" xml:\"ROW_peer\""
						} "json:\"TABLE_peer\" xml:\"TABLE_peer\""
					} "json:\"ROW_vrf\" xml:\"ROW_vrf\""
				} "json:\"TABLE_vrf\" xml:\"TABLE_vrf\""
			}{struct {
				Asn      string "json:\"asn\" xml:\"asn\""
				TableVrf []struct {
					RowVrf []struct {
						Vrf       string "json:\"vrf\" xml:\"vrf\""
						TablePeer []struct {
							RowPeer []struct {
								PeerHandle     string   "json:\"peer_handle\" xml:\"peer_handle\""
								PeerIpaddr     string   "json:\"peer_ipaddr\" xml:\"peer_ipaddr\""
								PeerIfname     string   "json:\"peer_ifname\" xml:\"peer_ifname\""
								PeerHoldtime   uint     "json:\"peer_holdtime\" xml:\"peer_holdtime\""
								PeerSrtt       uint     "json:\"peer_srtt\" xml:\"peer_srtt\""
								PeerRto        uint     "json:\"peer_rto\" xml:\"peer_rto\""
								PeerXmitqCount uint     "json:\"peer_xmitq_count\" xml:\"peer_xmitq_count\""
								PeerLastSeqno  uint     "json:\"peer_last_seqno\" xml:\"peer_last_seqno\""
								PeerUptime     Duration "json:\"peer_uptime\" xml:\"peer_uptime\""
							} "json:\"ROW_peer\" xml:\"ROW_peer\""
						} "json:\"TABLE_peer\" xml:\"TABLE_peer\""
					} "json:\"ROW_vrf\" xml:\"ROW_vrf\""
				} "json:\"TABLE_vrf\" xml:\"TABLE_vrf\""
			}{Asn: "1", TableVrf: []struct {
				RowVrf []struct {
					Vrf       string "json:\"vrf\" xml:\"vrf\""
					TablePeer []struct {
						RowPeer []struct {
							PeerHandle     string   "json:\"peer_handle\" xml:\"peer_handle\""
							PeerIpaddr     string   "json:\"peer_ipaddr\" xml:\"peer_ipaddr\""
							PeerIfname     string   "json:\"peer_ifname\" xml:\"peer_ifname\""
							PeerHoldtime   uint     "json:\"peer_holdtime\" xml:\"peer_holdtime\""
							PeerSrtt       uint     "json:\"peer_srtt\" xml:\"peer_srtt\""
							PeerRto        uint     "json:\"peer_rto\" xml:\"peer_rto\""
							PeerXmitqCount uint     "json:\"peer_xmitq_count\" xml:\"peer_xmitq_count\""
							PeerLastSeqno  uint     "json:\"peer_last_seqno\" xml:\"peer_last_seqno\""
							PeerUptime     Duration "json:\"peer_uptime\" xml:\"peer_uptime\""
						} "json:\"ROW_peer\" xml:\"ROW_peer\""
					} "json:\"TABLE_peer\" xml:\"TABLE_peer\""
				} "json:\"ROW_vrf\" xml:\"ROW_vrf\""
			}{struct {
				RowVrf []struct {
					Vrf       string "json:\"vrf\" xml:\"vrf\""
					TablePeer []struct {
						RowPeer []struct {
							PeerHandle     string   "json:\"peer_handle\" xml:\"peer_handle\""
							PeerIpaddr     string   "json:\"peer_ipaddr\" xml:\"peer_ipaddr\""
							PeerIfname     string   "json:\"peer_ifname\" xml:\"peer_ifname\""
							PeerHoldtime   uint     "json:\"peer_holdtime\" xml:\"peer_holdtime\""
							PeerSrtt       uint     "json:\"peer_srtt\" xml:\"peer_srtt\""
							PeerRto        uint     "json:\"peer_rto\" xml:\"peer_rto\""
							PeerXmitqCount uint     "json:\"peer_xmitq_count\" xml:\"peer_xmitq_count\""
							PeerLastSeqno  uint     "json:\"peer_last_seqno\" xml:\"peer_last_seqno\""
							PeerUptime     Duration "json:\"peer_uptime\" xml:\"peer_uptime\""
						} "json:\"ROW_peer\" xml:\"ROW_peer\""
					} "json:\"TABLE_peer\" xml:\"TABLE_peer\""
				} "json:\"ROW_vrf\" xml:\"ROW_vrf\""
			}{RowVrf: []struct {
				Vrf       string "json:\"vrf\" xml:\"vrf\""
				TablePeer []struct {
					RowPeer []struct {
						PeerHandle     string   "json:\"peer_handle\" xml:\"peer_handle\""
						PeerIpaddr     string   "json:\"peer_ipaddr\" xml:\"peer_ipaddr\""
						PeerIfname     string   "json:\"peer_ifname\" xml:\"peer_ifname\""
						PeerHoldtime   uint     "json:\"peer_holdtime\" xml:\"peer_holdtime\""
						PeerSrtt       uint     "json:\"peer_srtt\" xml:\"peer_srtt\""
						PeerRto        uint     "json:\"peer_rto\" xml:\"peer_rto\""
						PeerXmitqCount uint     "json:\"peer_xmitq_count\" xml:\"peer_xmitq_count\""
						PeerLastSeqno  uint     "json:\"peer_last_seqno\" xml:\"peer_last_seqno\""
						PeerUptime     Duration "json:\"peer_uptime\" xml:\"peer_uptime\""
					} "json:\"ROW_peer\" xml:\"ROW_peer\""
				} "json:\"TABLE_peer\" xml:\"TABLE_peer\""
			}{struct {
				Vrf       string "json:\"vrf\" xml:\"vrf\""
				TablePeer []struct {
					RowPeer []struct {
						PeerHandle     string   "json:\"peer_handle\" xml:\"peer_handle\""
						PeerIpaddr     string   "json:\"peer_ipaddr\" xml:\"peer_ipaddr\""
						PeerIfname     string   "json:\"peer_ifname\" xml:\"peer_ifname\""
						PeerHoldtime   uint     "json:\"peer_holdtime\" xml:\"peer_holdtime\""
						PeerSrtt       uint     "json:\"peer_srtt\" xml:\"peer_srtt\""
						PeerRto        uint     "json:\"peer_rto\" xml:\"peer_rto\""
						PeerXmitqCount uint     "json:\"peer_xmitq_count\" xml:\"peer_xmitq_count\""
						PeerLastSeqno  uint     "json:\"peer_last_seqno\" xml:\"peer_last_seqno\""
						PeerUptime     Duration "json:\"peer_uptime\" xml:\"peer_uptime\""
					} "json:\"ROW_peer\" xml:\"ROW_peer\""
				} "json:\"TABLE_peer\" xml:\"TABLE_peer\""
			}{Vrf: "default", TablePeer: []struct {
				RowPeer []struct {
					PeerHandle     string   "json:\"peer_handle\" xml:\"peer_handle\""
					PeerIpaddr     string   "json:\"peer_ipaddr\" xml:\"peer_ipaddr\""
					PeerIfname     string   "json:\"peer_ifname\" xml:\"peer_ifname\""
					PeerHoldtime   uint     "json:\"peer_holdtime\" xml:\"peer_holdtime\""
					PeerSrtt       uint     "json:\"peer_srtt\" xml:\"peer_srtt\""
					PeerRto        uint     "json:\"peer_rto\" xml:\"peer_rto\""
					PeerXmitqCount uint     "json:\"peer_xmitq_count\" xml:\"peer_xmitq_count\""
					PeerLastSeqno  uint     "json:\"peer_last_seqno\" xml:\"peer_last_seqno\""
					PeerUptime     Duration "json:\"peer_uptime\" xml:\"peer_uptime\""
				} "json:\"ROW_peer\" xml:\"ROW_peer\""
			}{struct {
				RowPeer []struct {
					PeerHandle     string   "json:\"peer_handle\" xml:\"peer_handle\""
					PeerIpaddr     string   "json:\"peer_ipaddr\" xml:\"peer_ipaddr\""
					PeerIfname     string   "json:\"peer_ifname\" xml:\"peer_ifname\""
					PeerHoldtime   uint     "json:\"peer_holdtime\" xml:\"peer_holdtime\""
					PeerSrtt       uint     "json:\"peer_srtt\" xml:\"peer_srtt\""
					PeerRto        uint     "json:\"peer_rto\" xml:\"peer_rto\""
					PeerXmitqCount uint     "json:\"peer_xmitq_count\" xml:\"peer_xmitq_count\""
					PeerLastSeqno  uint     "json:\"peer_last_seqno\" xml:\"peer_last_seqno\""
					PeerUptime     Duration "json:\"peer_uptime\" xml:\"peer_uptime\""
				} "json:\"ROW_peer\" xml:\"ROW_peer\""
			}{RowPeer: []struct {
				PeerHandle     string   "json:\"peer_handle\" xml:\"peer_handle\""
				PeerIpaddr     string   "json:\"peer_ipaddr\" xml:\"peer_ipaddr\""
				PeerIfname     string   "json:\"peer_ifname\" xml:\"peer_ifname\""
				PeerHoldtime   uint     "json:\"peer_holdtime\" xml:\"peer_holdtime\""
				PeerSrtt       uint     "json:\"peer_srtt\" xml:\"peer_srtt\""
				PeerRto        uint     "json:\"peer_rto\" xml:\"peer_rto\""
				PeerXmitqCount uint     "json:\"peer_xmitq_count\" xml:\"peer_xmitq_count\""
				PeerLastSeqno  uint     "json:\"peer_last_seqno\" xml:\"peer_last_seqno\""
				PeerUptime     Duration "json:\"peer_uptime\" xml:\"peer_uptime\""
			}{struct {
				PeerHandle     string   "json:\"peer_handle\" xml:\"peer_handle\""
				PeerIpaddr     string   "json:\"peer_ipaddr\" xml:\"peer_ipaddr\""
				PeerIfname     string   "json:\"peer_ifname\" xml:\"peer_ifname\""
				PeerHoldtime   uint     "json:\"peer_holdtime\" xml:\"peer_holdtime\""
				PeerSrtt       uint     "json:\"peer_srtt\" xml:\"peer_srtt\""
				PeerRto        uint     "json:\"peer_rto\" xml:\"peer_rto\""
				PeerXmitqCount uint     "json:\"peer_xmitq_count\" xml:\"peer_xmitq_count\""
				PeerLastSeqno  uint     "json:\"peer_last_seqno\" xml:\"peer_last_seqno\""
				PeerUptime     Duration "json:\"peer_uptime\" xml:\"peer_uptime\""
			}{PeerHandle: "0", PeerIpaddr: "10.1.0.1", PeerIfname: "Eth1/1", PeerHoldtime: 0xd, PeerSrtt: 0x6, PeerRto: 0x32, PeerXmitqCount: 0x0, PeerLastSeqno: 0x4, PeerUptime: 0x28e05727a00}}}}}}}}}}}}}, Code: "200", Input: "show ip eigrp neighbors vrf all", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"}},
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
		dat, err := NewShowIpEigrpNeighborsVrfAllFromBytes(content)
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

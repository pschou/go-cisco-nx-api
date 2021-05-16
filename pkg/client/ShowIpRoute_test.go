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

func TestParseShowIPRouteJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowIpRouteResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.ip.route",
			exp: &ShowIpRouteResponse{InsAPI: struct {
				Outputs struct {
					Output ShowIpRouteResponseResult "json:\"output\" xml:\"output\""
				} "json:\"outputs\" xml:\"outputs\""
				Sid     string "json:\"sid\" xml:\"sid\""
				Type    string "json:\"type\" xml:\"type\""
				Version string "json:\"version\" xml:\"version\""
			}{Outputs: struct {
				Output ShowIpRouteResponseResult "json:\"output\" xml:\"output\""
			}{Output: ShowIpRouteResponseResult{Body: ShowIpRouteResultBody{TableVrf: []struct {
				RowVrf []struct {
					TableAddrf []struct {
						RowAddrf []struct {
							TablePrefix []struct {
								RowPrefix []struct {
									TablePath []struct {
										RowPath []struct {
											ClientName string   "json:\"clientname\" xml:\"clientname\""
											IfName     string   "json:\"ifname\" xml:\"ifname\""
											Metric     int      "json:\"metric\" xml:\"metric\""
											Pref       int      "json:\"pref\" xml:\"pref\""
											UBest      bool     "json:\"ubest\" xml:\"ubest\""
											UpTime     Duration "json:\"uptime\" xml:\"uptime\""
										} "json:\"ROW_path\" xml:\"ROW_path\""
									} "json:\"TABLE_path\" xml:\"TABLE_path\""
									Attached   bool   "json:\"attached\" xml:\"attached\""
									IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
									MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
									UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
								} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
							} "json:\"TABLE_prefix\" xml:\"TABLE_prefix\""
							AddRf string "json:\"addrf\" xml:\"addrf\""
						} "json:\"ROW_addrf\" xml:\"ROW_addrf\""
					} "json:\"TABLE_addrf\" xml:\"TABLE_addrf\""
					VrfNameOut string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
				} "json:\"ROW_vrf\" xml:\"ROW_vrf\""
			}{struct {
				RowVrf []struct {
					TableAddrf []struct {
						RowAddrf []struct {
							TablePrefix []struct {
								RowPrefix []struct {
									TablePath []struct {
										RowPath []struct {
											ClientName string   "json:\"clientname\" xml:\"clientname\""
											IfName     string   "json:\"ifname\" xml:\"ifname\""
											Metric     int      "json:\"metric\" xml:\"metric\""
											Pref       int      "json:\"pref\" xml:\"pref\""
											UBest      bool     "json:\"ubest\" xml:\"ubest\""
											UpTime     Duration "json:\"uptime\" xml:\"uptime\""
										} "json:\"ROW_path\" xml:\"ROW_path\""
									} "json:\"TABLE_path\" xml:\"TABLE_path\""
									Attached   bool   "json:\"attached\" xml:\"attached\""
									IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
									MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
									UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
								} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
							} "json:\"TABLE_prefix\" xml:\"TABLE_prefix\""
							AddRf string "json:\"addrf\" xml:\"addrf\""
						} "json:\"ROW_addrf\" xml:\"ROW_addrf\""
					} "json:\"TABLE_addrf\" xml:\"TABLE_addrf\""
					VrfNameOut string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
				} "json:\"ROW_vrf\" xml:\"ROW_vrf\""
			}{RowVrf: []struct {
				TableAddrf []struct {
					RowAddrf []struct {
						TablePrefix []struct {
							RowPrefix []struct {
								TablePath []struct {
									RowPath []struct {
										ClientName string   "json:\"clientname\" xml:\"clientname\""
										IfName     string   "json:\"ifname\" xml:\"ifname\""
										Metric     int      "json:\"metric\" xml:\"metric\""
										Pref       int      "json:\"pref\" xml:\"pref\""
										UBest      bool     "json:\"ubest\" xml:\"ubest\""
										UpTime     Duration "json:\"uptime\" xml:\"uptime\""
									} "json:\"ROW_path\" xml:\"ROW_path\""
								} "json:\"TABLE_path\" xml:\"TABLE_path\""
								Attached   bool   "json:\"attached\" xml:\"attached\""
								IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
								MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
								UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
							} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
						} "json:\"TABLE_prefix\" xml:\"TABLE_prefix\""
						AddRf string "json:\"addrf\" xml:\"addrf\""
					} "json:\"ROW_addrf\" xml:\"ROW_addrf\""
				} "json:\"TABLE_addrf\" xml:\"TABLE_addrf\""
				VrfNameOut string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
			}{struct {
				TableAddrf []struct {
					RowAddrf []struct {
						TablePrefix []struct {
							RowPrefix []struct {
								TablePath []struct {
									RowPath []struct {
										ClientName string   "json:\"clientname\" xml:\"clientname\""
										IfName     string   "json:\"ifname\" xml:\"ifname\""
										Metric     int      "json:\"metric\" xml:\"metric\""
										Pref       int      "json:\"pref\" xml:\"pref\""
										UBest      bool     "json:\"ubest\" xml:\"ubest\""
										UpTime     Duration "json:\"uptime\" xml:\"uptime\""
									} "json:\"ROW_path\" xml:\"ROW_path\""
								} "json:\"TABLE_path\" xml:\"TABLE_path\""
								Attached   bool   "json:\"attached\" xml:\"attached\""
								IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
								MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
								UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
							} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
						} "json:\"TABLE_prefix\" xml:\"TABLE_prefix\""
						AddRf string "json:\"addrf\" xml:\"addrf\""
					} "json:\"ROW_addrf\" xml:\"ROW_addrf\""
				} "json:\"TABLE_addrf\" xml:\"TABLE_addrf\""
				VrfNameOut string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
			}{TableAddrf: []struct {
				RowAddrf []struct {
					TablePrefix []struct {
						RowPrefix []struct {
							TablePath []struct {
								RowPath []struct {
									ClientName string   "json:\"clientname\" xml:\"clientname\""
									IfName     string   "json:\"ifname\" xml:\"ifname\""
									Metric     int      "json:\"metric\" xml:\"metric\""
									Pref       int      "json:\"pref\" xml:\"pref\""
									UBest      bool     "json:\"ubest\" xml:\"ubest\""
									UpTime     Duration "json:\"uptime\" xml:\"uptime\""
								} "json:\"ROW_path\" xml:\"ROW_path\""
							} "json:\"TABLE_path\" xml:\"TABLE_path\""
							Attached   bool   "json:\"attached\" xml:\"attached\""
							IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
							MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
							UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
						} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
					} "json:\"TABLE_prefix\" xml:\"TABLE_prefix\""
					AddRf string "json:\"addrf\" xml:\"addrf\""
				} "json:\"ROW_addrf\" xml:\"ROW_addrf\""
			}{struct {
				RowAddrf []struct {
					TablePrefix []struct {
						RowPrefix []struct {
							TablePath []struct {
								RowPath []struct {
									ClientName string   "json:\"clientname\" xml:\"clientname\""
									IfName     string   "json:\"ifname\" xml:\"ifname\""
									Metric     int      "json:\"metric\" xml:\"metric\""
									Pref       int      "json:\"pref\" xml:\"pref\""
									UBest      bool     "json:\"ubest\" xml:\"ubest\""
									UpTime     Duration "json:\"uptime\" xml:\"uptime\""
								} "json:\"ROW_path\" xml:\"ROW_path\""
							} "json:\"TABLE_path\" xml:\"TABLE_path\""
							Attached   bool   "json:\"attached\" xml:\"attached\""
							IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
							MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
							UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
						} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
					} "json:\"TABLE_prefix\" xml:\"TABLE_prefix\""
					AddRf string "json:\"addrf\" xml:\"addrf\""
				} "json:\"ROW_addrf\" xml:\"ROW_addrf\""
			}{RowAddrf: []struct {
				TablePrefix []struct {
					RowPrefix []struct {
						TablePath []struct {
							RowPath []struct {
								ClientName string   "json:\"clientname\" xml:\"clientname\""
								IfName     string   "json:\"ifname\" xml:\"ifname\""
								Metric     int      "json:\"metric\" xml:\"metric\""
								Pref       int      "json:\"pref\" xml:\"pref\""
								UBest      bool     "json:\"ubest\" xml:\"ubest\""
								UpTime     Duration "json:\"uptime\" xml:\"uptime\""
							} "json:\"ROW_path\" xml:\"ROW_path\""
						} "json:\"TABLE_path\" xml:\"TABLE_path\""
						Attached   bool   "json:\"attached\" xml:\"attached\""
						IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
						MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
						UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
					} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
				} "json:\"TABLE_prefix\" xml:\"TABLE_prefix\""
				AddRf string "json:\"addrf\" xml:\"addrf\""
			}{struct {
				TablePrefix []struct {
					RowPrefix []struct {
						TablePath []struct {
							RowPath []struct {
								ClientName string   "json:\"clientname\" xml:\"clientname\""
								IfName     string   "json:\"ifname\" xml:\"ifname\""
								Metric     int      "json:\"metric\" xml:\"metric\""
								Pref       int      "json:\"pref\" xml:\"pref\""
								UBest      bool     "json:\"ubest\" xml:\"ubest\""
								UpTime     Duration "json:\"uptime\" xml:\"uptime\""
							} "json:\"ROW_path\" xml:\"ROW_path\""
						} "json:\"TABLE_path\" xml:\"TABLE_path\""
						Attached   bool   "json:\"attached\" xml:\"attached\""
						IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
						MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
						UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
					} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
				} "json:\"TABLE_prefix\" xml:\"TABLE_prefix\""
				AddRf string "json:\"addrf\" xml:\"addrf\""
			}{TablePrefix: []struct {
				RowPrefix []struct {
					TablePath []struct {
						RowPath []struct {
							ClientName string   "json:\"clientname\" xml:\"clientname\""
							IfName     string   "json:\"ifname\" xml:\"ifname\""
							Metric     int      "json:\"metric\" xml:\"metric\""
							Pref       int      "json:\"pref\" xml:\"pref\""
							UBest      bool     "json:\"ubest\" xml:\"ubest\""
							UpTime     Duration "json:\"uptime\" xml:\"uptime\""
						} "json:\"ROW_path\" xml:\"ROW_path\""
					} "json:\"TABLE_path\" xml:\"TABLE_path\""
					Attached   bool   "json:\"attached\" xml:\"attached\""
					IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
					MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
					UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
				} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
			}{struct {
				RowPrefix []struct {
					TablePath []struct {
						RowPath []struct {
							ClientName string   "json:\"clientname\" xml:\"clientname\""
							IfName     string   "json:\"ifname\" xml:\"ifname\""
							Metric     int      "json:\"metric\" xml:\"metric\""
							Pref       int      "json:\"pref\" xml:\"pref\""
							UBest      bool     "json:\"ubest\" xml:\"ubest\""
							UpTime     Duration "json:\"uptime\" xml:\"uptime\""
						} "json:\"ROW_path\" xml:\"ROW_path\""
					} "json:\"TABLE_path\" xml:\"TABLE_path\""
					Attached   bool   "json:\"attached\" xml:\"attached\""
					IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
					MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
					UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
				} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
			}{RowPrefix: []struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "static", IfName: "Null0", Metric: 0, Pref: 1, UBest: true, UpTime: 0x24d775c652200}}}}, Attached: false, IPPrefix: "7.57.0.0/16", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Vlan253", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d622c5f5400}}}}, Attached: true, IPPrefix: "7.57.253.0/30", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Vlan253", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d622c5f5400}}}}, Attached: true, IPPrefix: "7.57.253.2/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "bgp-65057", IfName: "", Metric: 0, Pref: 20, UBest: true, UpTime: 0x24d4b1f833600}, struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "bgp-65057", IfName: "", Metric: 0, Pref: 20, UBest: true, UpTime: 0x24d4b1f833600}}}}, Attached: false, IPPrefix: "7.57.255.1/32", MCastNHops: 0, UCastNHops: 2}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Lo0", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d76a994c400}, struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Lo0", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d76a994c400}}}}, Attached: true, IPPrefix: "7.57.255.2/32", MCastNHops: 0, UCastNHops: 2}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "ospf-1", IfName: "Vlan253", Metric: 42, Pref: 110, UBest: true, UpTime: 0x24d5f611ddc00}}}}, Attached: false, IPPrefix: "8.8.8.0/24", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "ospf-1", IfName: "Vlan253", Metric: 42, Pref: 110, UBest: true, UpTime: 0x24d5f611ddc00}}}}, Attached: false, IPPrefix: "9.9.9.0/24", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "static", IfName: "Null0", Metric: 0, Pref: 1, UBest: true, UpTime: 0x24d775c652200}}}}, Attached: false, IPPrefix: "10.0.0.0/8", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "static", IfName: "Null0", Metric: 0, Pref: 1, UBest: true, UpTime: 0x24d775c652200}}}}, Attached: false, IPPrefix: "10.57.0.0/16", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Vlan8", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d64806b3800}}}}, Attached: true, IPPrefix: "10.57.8.0/22", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Vlan8", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d64806b3800}}}}, Attached: true, IPPrefix: "10.57.8.3/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Vlan50", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d622c5f5400}}}}, Attached: true, IPPrefix: "10.57.50.0/24", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Vlan50", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d622c5f5400}}}}, Attached: true, IPPrefix: "10.57.50.4/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Eth1/1", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "10.100.90.128/30", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Eth1/1", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "10.100.90.129/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Eth1/2", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d656ed66000}}}}, Attached: true, IPPrefix: "10.100.157.0/30", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Eth1/2", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d656ed66000}}}}, Attached: true, IPPrefix: "10.100.157.2/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Eth1/4", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d656ed66000}}}}, Attached: true, IPPrefix: "10.100.157.8/30", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Eth1/4", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d656ed66000}}}}, Attached: true, IPPrefix: "10.100.157.10/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "static", IfName: "Eth1/5", Metric: 0, Pref: 1, UBest: true, UpTime: 0x24d70d7770a00}}}}, Attached: false, IPPrefix: "17.0.1.1/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "static", IfName: "Eth1/5", Metric: 0, Pref: 1, UBest: true, UpTime: 0x24d70d7770a00}}}}, Attached: false, IPPrefix: "17.0.1.2/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "static", IfName: "Eth1/5", Metric: 0, Pref: 1, UBest: true, UpTime: 0x24d70d7770a00}}}}, Attached: false, IPPrefix: "17.0.1.3/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "bgp-65057", IfName: "", Metric: 0, Pref: 20, UBest: true, UpTime: 0x24d4b1f833600}, struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "bgp-65057", IfName: "", Metric: 0, Pref: 20, UBest: true, UpTime: 0x24d4b1f833600}}}}, Attached: false, IPPrefix: "33.33.33.33/32", MCastNHops: 0, UCastNHops: 2}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Eth1/1.50", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "89.1.1.0/24", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Eth1/1.50", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "89.1.1.1/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Eth1/1.51", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "89.1.2.0/24", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Eth1/1.51", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "89.1.2.1/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Eth1/1.52", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "89.1.3.0/24", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Eth1/1.52", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "89.1.3.1/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Eth1/1.53", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "89.1.4.0/24", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Eth1/1.53", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "89.1.4.1/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Eth1/5", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d70d7770a00}}}}, Attached: true, IPPrefix: "94.1.1.0/24", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Eth1/5", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d70d7770a00}}}}, Attached: true, IPPrefix: "94.1.1.96/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Eth1/7", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d64bc060200}}}}, Attached: true, IPPrefix: "192.168.161.0/24", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Eth1/7", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d64bc060200}}}}, Attached: true, IPPrefix: "192.168.161.2/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{}}}, Attached: true, IPPrefix: "192.168.161.2/32", MCastNHops: 0, UCastNHops: 1}}}}, AddRf: "ipv4"}}}}, VrfNameOut: "default"}}}}}, Code: "200", Input: "show ip route ", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"}},
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
		dat, err := NewShowIpRouteFromBytes(content)
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

func TestParseShowIPRouteResponseJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *ShowIpRouteResponseResult
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "result.show.ip.route",
			exp: &ShowIpRouteResponseResult{Body: ShowIpRouteResultBody{TableVrf: []struct {
				RowVrf []struct {
					TableAddrf []struct {
						RowAddrf []struct {
							TablePrefix []struct {
								RowPrefix []struct {
									TablePath []struct {
										RowPath []struct {
											ClientName string   "json:\"clientname\" xml:\"clientname\""
											IfName     string   "json:\"ifname\" xml:\"ifname\""
											Metric     int      "json:\"metric\" xml:\"metric\""
											Pref       int      "json:\"pref\" xml:\"pref\""
											UBest      bool     "json:\"ubest\" xml:\"ubest\""
											UpTime     Duration "json:\"uptime\" xml:\"uptime\""
										} "json:\"ROW_path\" xml:\"ROW_path\""
									} "json:\"TABLE_path\" xml:\"TABLE_path\""
									Attached   bool   "json:\"attached\" xml:\"attached\""
									IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
									MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
									UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
								} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
							} "json:\"TABLE_prefix\" xml:\"TABLE_prefix\""
							AddRf string "json:\"addrf\" xml:\"addrf\""
						} "json:\"ROW_addrf\" xml:\"ROW_addrf\""
					} "json:\"TABLE_addrf\" xml:\"TABLE_addrf\""
					VrfNameOut string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
				} "json:\"ROW_vrf\" xml:\"ROW_vrf\""
			}{struct {
				RowVrf []struct {
					TableAddrf []struct {
						RowAddrf []struct {
							TablePrefix []struct {
								RowPrefix []struct {
									TablePath []struct {
										RowPath []struct {
											ClientName string   "json:\"clientname\" xml:\"clientname\""
											IfName     string   "json:\"ifname\" xml:\"ifname\""
											Metric     int      "json:\"metric\" xml:\"metric\""
											Pref       int      "json:\"pref\" xml:\"pref\""
											UBest      bool     "json:\"ubest\" xml:\"ubest\""
											UpTime     Duration "json:\"uptime\" xml:\"uptime\""
										} "json:\"ROW_path\" xml:\"ROW_path\""
									} "json:\"TABLE_path\" xml:\"TABLE_path\""
									Attached   bool   "json:\"attached\" xml:\"attached\""
									IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
									MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
									UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
								} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
							} "json:\"TABLE_prefix\" xml:\"TABLE_prefix\""
							AddRf string "json:\"addrf\" xml:\"addrf\""
						} "json:\"ROW_addrf\" xml:\"ROW_addrf\""
					} "json:\"TABLE_addrf\" xml:\"TABLE_addrf\""
					VrfNameOut string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
				} "json:\"ROW_vrf\" xml:\"ROW_vrf\""
			}{RowVrf: []struct {
				TableAddrf []struct {
					RowAddrf []struct {
						TablePrefix []struct {
							RowPrefix []struct {
								TablePath []struct {
									RowPath []struct {
										ClientName string   "json:\"clientname\" xml:\"clientname\""
										IfName     string   "json:\"ifname\" xml:\"ifname\""
										Metric     int      "json:\"metric\" xml:\"metric\""
										Pref       int      "json:\"pref\" xml:\"pref\""
										UBest      bool     "json:\"ubest\" xml:\"ubest\""
										UpTime     Duration "json:\"uptime\" xml:\"uptime\""
									} "json:\"ROW_path\" xml:\"ROW_path\""
								} "json:\"TABLE_path\" xml:\"TABLE_path\""
								Attached   bool   "json:\"attached\" xml:\"attached\""
								IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
								MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
								UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
							} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
						} "json:\"TABLE_prefix\" xml:\"TABLE_prefix\""
						AddRf string "json:\"addrf\" xml:\"addrf\""
					} "json:\"ROW_addrf\" xml:\"ROW_addrf\""
				} "json:\"TABLE_addrf\" xml:\"TABLE_addrf\""
				VrfNameOut string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
			}{struct {
				TableAddrf []struct {
					RowAddrf []struct {
						TablePrefix []struct {
							RowPrefix []struct {
								TablePath []struct {
									RowPath []struct {
										ClientName string   "json:\"clientname\" xml:\"clientname\""
										IfName     string   "json:\"ifname\" xml:\"ifname\""
										Metric     int      "json:\"metric\" xml:\"metric\""
										Pref       int      "json:\"pref\" xml:\"pref\""
										UBest      bool     "json:\"ubest\" xml:\"ubest\""
										UpTime     Duration "json:\"uptime\" xml:\"uptime\""
									} "json:\"ROW_path\" xml:\"ROW_path\""
								} "json:\"TABLE_path\" xml:\"TABLE_path\""
								Attached   bool   "json:\"attached\" xml:\"attached\""
								IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
								MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
								UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
							} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
						} "json:\"TABLE_prefix\" xml:\"TABLE_prefix\""
						AddRf string "json:\"addrf\" xml:\"addrf\""
					} "json:\"ROW_addrf\" xml:\"ROW_addrf\""
				} "json:\"TABLE_addrf\" xml:\"TABLE_addrf\""
				VrfNameOut string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
			}{TableAddrf: []struct {
				RowAddrf []struct {
					TablePrefix []struct {
						RowPrefix []struct {
							TablePath []struct {
								RowPath []struct {
									ClientName string   "json:\"clientname\" xml:\"clientname\""
									IfName     string   "json:\"ifname\" xml:\"ifname\""
									Metric     int      "json:\"metric\" xml:\"metric\""
									Pref       int      "json:\"pref\" xml:\"pref\""
									UBest      bool     "json:\"ubest\" xml:\"ubest\""
									UpTime     Duration "json:\"uptime\" xml:\"uptime\""
								} "json:\"ROW_path\" xml:\"ROW_path\""
							} "json:\"TABLE_path\" xml:\"TABLE_path\""
							Attached   bool   "json:\"attached\" xml:\"attached\""
							IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
							MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
							UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
						} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
					} "json:\"TABLE_prefix\" xml:\"TABLE_prefix\""
					AddRf string "json:\"addrf\" xml:\"addrf\""
				} "json:\"ROW_addrf\" xml:\"ROW_addrf\""
			}{struct {
				RowAddrf []struct {
					TablePrefix []struct {
						RowPrefix []struct {
							TablePath []struct {
								RowPath []struct {
									ClientName string   "json:\"clientname\" xml:\"clientname\""
									IfName     string   "json:\"ifname\" xml:\"ifname\""
									Metric     int      "json:\"metric\" xml:\"metric\""
									Pref       int      "json:\"pref\" xml:\"pref\""
									UBest      bool     "json:\"ubest\" xml:\"ubest\""
									UpTime     Duration "json:\"uptime\" xml:\"uptime\""
								} "json:\"ROW_path\" xml:\"ROW_path\""
							} "json:\"TABLE_path\" xml:\"TABLE_path\""
							Attached   bool   "json:\"attached\" xml:\"attached\""
							IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
							MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
							UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
						} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
					} "json:\"TABLE_prefix\" xml:\"TABLE_prefix\""
					AddRf string "json:\"addrf\" xml:\"addrf\""
				} "json:\"ROW_addrf\" xml:\"ROW_addrf\""
			}{RowAddrf: []struct {
				TablePrefix []struct {
					RowPrefix []struct {
						TablePath []struct {
							RowPath []struct {
								ClientName string   "json:\"clientname\" xml:\"clientname\""
								IfName     string   "json:\"ifname\" xml:\"ifname\""
								Metric     int      "json:\"metric\" xml:\"metric\""
								Pref       int      "json:\"pref\" xml:\"pref\""
								UBest      bool     "json:\"ubest\" xml:\"ubest\""
								UpTime     Duration "json:\"uptime\" xml:\"uptime\""
							} "json:\"ROW_path\" xml:\"ROW_path\""
						} "json:\"TABLE_path\" xml:\"TABLE_path\""
						Attached   bool   "json:\"attached\" xml:\"attached\""
						IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
						MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
						UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
					} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
				} "json:\"TABLE_prefix\" xml:\"TABLE_prefix\""
				AddRf string "json:\"addrf\" xml:\"addrf\""
			}{struct {
				TablePrefix []struct {
					RowPrefix []struct {
						TablePath []struct {
							RowPath []struct {
								ClientName string   "json:\"clientname\" xml:\"clientname\""
								IfName     string   "json:\"ifname\" xml:\"ifname\""
								Metric     int      "json:\"metric\" xml:\"metric\""
								Pref       int      "json:\"pref\" xml:\"pref\""
								UBest      bool     "json:\"ubest\" xml:\"ubest\""
								UpTime     Duration "json:\"uptime\" xml:\"uptime\""
							} "json:\"ROW_path\" xml:\"ROW_path\""
						} "json:\"TABLE_path\" xml:\"TABLE_path\""
						Attached   bool   "json:\"attached\" xml:\"attached\""
						IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
						MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
						UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
					} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
				} "json:\"TABLE_prefix\" xml:\"TABLE_prefix\""
				AddRf string "json:\"addrf\" xml:\"addrf\""
			}{TablePrefix: []struct {
				RowPrefix []struct {
					TablePath []struct {
						RowPath []struct {
							ClientName string   "json:\"clientname\" xml:\"clientname\""
							IfName     string   "json:\"ifname\" xml:\"ifname\""
							Metric     int      "json:\"metric\" xml:\"metric\""
							Pref       int      "json:\"pref\" xml:\"pref\""
							UBest      bool     "json:\"ubest\" xml:\"ubest\""
							UpTime     Duration "json:\"uptime\" xml:\"uptime\""
						} "json:\"ROW_path\" xml:\"ROW_path\""
					} "json:\"TABLE_path\" xml:\"TABLE_path\""
					Attached   bool   "json:\"attached\" xml:\"attached\""
					IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
					MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
					UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
				} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
			}{struct {
				RowPrefix []struct {
					TablePath []struct {
						RowPath []struct {
							ClientName string   "json:\"clientname\" xml:\"clientname\""
							IfName     string   "json:\"ifname\" xml:\"ifname\""
							Metric     int      "json:\"metric\" xml:\"metric\""
							Pref       int      "json:\"pref\" xml:\"pref\""
							UBest      bool     "json:\"ubest\" xml:\"ubest\""
							UpTime     Duration "json:\"uptime\" xml:\"uptime\""
						} "json:\"ROW_path\" xml:\"ROW_path\""
					} "json:\"TABLE_path\" xml:\"TABLE_path\""
					Attached   bool   "json:\"attached\" xml:\"attached\""
					IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
					MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
					UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
				} "json:\"ROW_prefix\" xml:\"ROW_prefix\""
			}{RowPrefix: []struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "static", IfName: "Null0", Metric: 0, Pref: 1, UBest: true, UpTime: 0x24d775c652200}}}}, Attached: false, IPPrefix: "7.57.0.0/16", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Vlan253", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d622c5f5400}}}}, Attached: true, IPPrefix: "7.57.253.0/30", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Vlan253", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d622c5f5400}}}}, Attached: true, IPPrefix: "7.57.253.2/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "bgp-65057", IfName: "", Metric: 0, Pref: 20, UBest: true, UpTime: 0x24d4b1f833600}, struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "bgp-65057", IfName: "", Metric: 0, Pref: 20, UBest: true, UpTime: 0x24d4b1f833600}}}}, Attached: false, IPPrefix: "7.57.255.1/32", MCastNHops: 0, UCastNHops: 2}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Lo0", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d76a994c400}, struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Lo0", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d76a994c400}}}}, Attached: true, IPPrefix: "7.57.255.2/32", MCastNHops: 0, UCastNHops: 2}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "ospf-1", IfName: "Vlan253", Metric: 42, Pref: 110, UBest: true, UpTime: 0x24d5f611ddc00}}}}, Attached: false, IPPrefix: "8.8.8.0/24", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "ospf-1", IfName: "Vlan253", Metric: 42, Pref: 110, UBest: true, UpTime: 0x24d5f611ddc00}}}}, Attached: false, IPPrefix: "9.9.9.0/24", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "static", IfName: "Null0", Metric: 0, Pref: 1, UBest: true, UpTime: 0x24d775c652200}}}}, Attached: false, IPPrefix: "10.0.0.0/8", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "static", IfName: "Null0", Metric: 0, Pref: 1, UBest: true, UpTime: 0x24d775c652200}}}}, Attached: false, IPPrefix: "10.57.0.0/16", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Vlan8", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d64806b3800}}}}, Attached: true, IPPrefix: "10.57.8.0/22", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Vlan8", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d64806b3800}}}}, Attached: true, IPPrefix: "10.57.8.3/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Vlan50", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d622c5f5400}}}}, Attached: true, IPPrefix: "10.57.50.0/24", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Vlan50", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d622c5f5400}}}}, Attached: true, IPPrefix: "10.57.50.4/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Eth1/1", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "10.100.90.128/30", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Eth1/1", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "10.100.90.129/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Eth1/2", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d656ed66000}}}}, Attached: true, IPPrefix: "10.100.157.0/30", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Eth1/2", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d656ed66000}}}}, Attached: true, IPPrefix: "10.100.157.2/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Eth1/4", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d656ed66000}}}}, Attached: true, IPPrefix: "10.100.157.8/30", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Eth1/4", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d656ed66000}}}}, Attached: true, IPPrefix: "10.100.157.10/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "static", IfName: "Eth1/5", Metric: 0, Pref: 1, UBest: true, UpTime: 0x24d70d7770a00}}}}, Attached: false, IPPrefix: "17.0.1.1/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "static", IfName: "Eth1/5", Metric: 0, Pref: 1, UBest: true, UpTime: 0x24d70d7770a00}}}}, Attached: false, IPPrefix: "17.0.1.2/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "static", IfName: "Eth1/5", Metric: 0, Pref: 1, UBest: true, UpTime: 0x24d70d7770a00}}}}, Attached: false, IPPrefix: "17.0.1.3/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "bgp-65057", IfName: "", Metric: 0, Pref: 20, UBest: true, UpTime: 0x24d4b1f833600}, struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "bgp-65057", IfName: "", Metric: 0, Pref: 20, UBest: true, UpTime: 0x24d4b1f833600}}}}, Attached: false, IPPrefix: "33.33.33.33/32", MCastNHops: 0, UCastNHops: 2}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Eth1/1.50", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "89.1.1.0/24", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Eth1/1.50", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "89.1.1.1/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Eth1/1.51", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "89.1.2.0/24", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Eth1/1.51", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "89.1.2.1/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Eth1/1.52", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "89.1.3.0/24", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Eth1/1.52", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "89.1.3.1/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Eth1/1.53", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "89.1.4.0/24", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Eth1/1.53", Metric: 0, Pref: 0, UBest: true, UpTime: 0x1fee55abcd000}}}}, Attached: true, IPPrefix: "89.1.4.1/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Eth1/5", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d70d7770a00}}}}, Attached: true, IPPrefix: "94.1.1.0/24", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Eth1/5", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d70d7770a00}}}}, Attached: true, IPPrefix: "94.1.1.96/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "direct", IfName: "Eth1/7", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d64bc060200}}}}, Attached: true, IPPrefix: "192.168.161.0/24", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{ClientName: "local", IfName: "Eth1/7", Metric: 0, Pref: 0, UBest: true, UpTime: 0x24d64bc060200}}}}, Attached: true, IPPrefix: "192.168.161.2/32", MCastNHops: 0, UCastNHops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						ClientName string   "json:\"clientname\" xml:\"clientname\""
						IfName     string   "json:\"ifname\" xml:\"ifname\""
						Metric     int      "json:\"metric\" xml:\"metric\""
						Pref       int      "json:\"pref\" xml:\"pref\""
						UBest      bool     "json:\"ubest\" xml:\"ubest\""
						UpTime     Duration "json:\"uptime\" xml:\"uptime\""
					} "json:\"ROW_path\" xml:\"ROW_path\""
				} "json:\"TABLE_path\" xml:\"TABLE_path\""
				Attached   bool   "json:\"attached\" xml:\"attached\""
				IPPrefix   string "json:\"ipprefix\" xml:\"ipprefix\""
				MCastNHops int    "json:\"mcast-nhops\" xml:\"mcast-nhops\""
				UCastNHops int    "json:\"ucast-nhops\" xml:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{struct {
				RowPath []struct {
					ClientName string   "json:\"clientname\" xml:\"clientname\""
					IfName     string   "json:\"ifname\" xml:\"ifname\""
					Metric     int      "json:\"metric\" xml:\"metric\""
					Pref       int      "json:\"pref\" xml:\"pref\""
					UBest      bool     "json:\"ubest\" xml:\"ubest\""
					UpTime     Duration "json:\"uptime\" xml:\"uptime\""
				} "json:\"ROW_path\" xml:\"ROW_path\""
			}{RowPath: []struct {
				ClientName string   "json:\"clientname\" xml:\"clientname\""
				IfName     string   "json:\"ifname\" xml:\"ifname\""
				Metric     int      "json:\"metric\" xml:\"metric\""
				Pref       int      "json:\"pref\" xml:\"pref\""
				UBest      bool     "json:\"ubest\" xml:\"ubest\""
				UpTime     Duration "json:\"uptime\" xml:\"uptime\""
			}{}}}, Attached: true, IPPrefix: "192.168.161.2/32", MCastNHops: 0, UCastNHops: 1}}}}, AddRf: "ipv4"}}}}, VrfNameOut: "default"}}}}}, Code: "200", Input: "show ip route ", Msg: "Success"},
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
		dat, err := NewShowIpRouteResultFromBytes(content)
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

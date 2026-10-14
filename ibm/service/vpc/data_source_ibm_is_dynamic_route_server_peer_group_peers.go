// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.0-a902401e-20260427-192904
 */

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsDynamicRouteServerPeerGroupPeers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsDynamicRouteServerPeerGroupPeersRead,

		Schema: map[string]*schema.Schema{
			"dynamic_route_server_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The dynamic route server identifier.",
			},
			"dynamic_route_server_peer_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The dynamic route server peer group identifier.",
			},
			"peers": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that this dynamic route server peer group peer was created.",
						},
						"creator": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of resource that created this peer:  - `dynamic_route_server`: The peer was created by a dynamic router server.  - `transit_gateway`:  The peer was created by a transit gateway.The enumerated values for this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this dynamic route server peer group peer.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID for this dynamic route server peer group peer.",
						},
						"lifecycle_reasons": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The reasons for the current `lifecycle_state` (if any).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
									},
									"message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "An explanation of the reason for this lifecycle state.",
									},
									"more_info": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A link to documentation about the reason for this lifecycle state.",
									},
								},
							},
						},
						"lifecycle_state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The lifecycle state of the dynamic route server peer group peer.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this dynamic route server peer group peer. The name must not be used by another peer in the peer group.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of this dynamic route server peer group peer connection:- `down`: not operational.- `up`: operating normally.",
						},
						"status_reasons": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The status reasons for the dynamic route server peer group peer.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The reasons for the current dynamic route server service connection status (if any).- `internal_error`- `peer_not_responding`- __TBD__.",
									},
									"message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "An explanation of the reason for this dynamic route server service connection's status.",
									},
									"more_info": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A link to documentation about this status reason.",
									},
								},
							},
						},
						"asn": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The autonomous system number (ASN) for this dynamic route server peer group address peer.",
						},
						"bidirectional_forwarding_detection": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The bidirectional forwarding detection (BFD) configuration for this dynamicroute server peer group peer.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"detect_multiplier": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The desired detection time multiplier for bidirectional forwarding detection control packets on this dynamic route server for this peer.",
									},
									"enabled": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether bidirectional forwarding detection (BFD) is enabled on this dynamic route server peer group peer.",
									},
									"mode": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The bidirectional forwarding detection mode of this peer:- `asynchronous`: Each peer sends BFD control packets independently. Session failure is detected when expected packets are not received within the detection interval as defined in[RFC 5880](https://www.rfc-editor.org/rfc/rfc5880.html?#section-3.2)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
									},
									"receive_interval": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum interval, in microseconds, between received bidirectional forwarding detection control packets that this dynamic route server is capable of supporting. The actual interval is negotiated between this dynamic route server and the peer.",
									},
									"role": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The bidirectional forwarding detection role used in session initialization:  - `active`: Actively initiates BFD control packets to bring up the session.  - `passive`: Waits for BFD control packets from the peer and does not initiate    the session.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
									},
									"sessions": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The bidirectional forwarding detection sessions for this peer.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"local": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The local peer for this bidirectional forwarding detection session.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"deleted": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"more_info": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "A link to documentation about deleted resources.",
																		},
																	},
																},
															},
															"href": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The URL for this dynamic route server member.",
															},
															"id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The unique identifier for this dynamic route server member.",
															},
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name for this dynamic route server member. The name is unique across all members in the dynamic route server.",
															},
															"resource_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The resource type.",
															},
															"virtual_network_interfaces": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The virtual network interfaces for this dynamic route server member.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"crn": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The CRN for this virtual network interface.",
																		},
																		"deleted": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"more_info": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "A link to documentation about deleted resources.",
																					},
																				},
																			},
																		},
																		"href": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The URL for this virtual network interface.",
																		},
																		"id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The unique identifier for this virtual network interface.",
																		},
																		"name": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.",
																		},
																		"primary_ip": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "The primary IP for this virtual network interface.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"address": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.",
																					},
																					"deleted": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"more_info": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "A link to documentation about deleted resources.",
																								},
																							},
																						},
																					},
																					"href": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "The URL for this reserved IP.",
																					},
																					"id": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "The unique identifier for this reserved IP.",
																					},
																					"name": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "The name for this reserved IP. The name is unique across all reserved IPs in a subnet.",
																					},
																					"resource_type": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "The resource type.",
																					},
																				},
																			},
																		},
																		"resource_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The resource type.",
																		},
																		"subnet": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "The associated subnet.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"crn": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "The CRN for this subnet.",
																					},
																					"deleted": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"more_info": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "A link to documentation about deleted resources.",
																								},
																							},
																						},
																					},
																					"href": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "The URL for this subnet.",
																					},
																					"id": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "The unique identifier for this subnet.",
																					},
																					"name": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "The name for this subnet. The name is unique across all subnets in the VPC.",
																					},
																					"resource_type": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "The resource type.",
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
												"remote": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The remote peer for this bidirectional forwarding detection session.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"deleted": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"more_info": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "A link to documentation about deleted resources.",
																		},
																	},
																},
															},
															"href": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The URL for this dynamic route server peer group peer.",
															},
															"id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The ID for this dynamic route server peer group peer.",
															},
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name for this dynamic route server peer group peer. The name must not be used by another peer in the peer group.",
															},
															"resource_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The resource type.",
															},
														},
													},
												},
												"state": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The current state of this bidirectional forwarding detection session as observed by the dynamic route server. The states are defined in [RFC 5880](https://www.rfc-editor.org/rfc/rfc5880.html#section-4.1).",
												},
											},
										},
									},
									"transmit_interval": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum interval, in microseconds that this dynamic route server prefer to use when transmitting bidirectional forwarding detection control to this peer. The actual interval is negotiated between this dynamic route server and the peer.",
									},
								},
							},
						},
						"endpoint": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The endpoint for a dynamic route server peer group address peer.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP address.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.",
									},
									"gateway": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The gateway IP address of the dynamic route server peer group address peer endpoint.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The IP address.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.",
												},
											},
										},
									},
									"deleted": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"more_info": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "A link to documentation about deleted resources.",
												},
											},
										},
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this reserved IP.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this reserved IP.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name for this reserved IP. The name is unique across all reserved IPs in a subnet.",
									},
									"resource_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type.",
									},
								},
							},
						},
						"priority": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The priority of the peer group peer. The priority is used to determine the preferred path for routing. A lower value indicates a higher priority.",
						},
						"sessions": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The BGP sessions for this peer group peer.Empty if `health_monitor.mode` is `none`.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"established_at": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The date and time that the BGP session was established. This property will be present only when the session `state` is `established`.",
									},
									"local": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The local peer for this BGP session.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deleted": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"more_info": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "A link to documentation about deleted resources.",
															},
														},
													},
												},
												"href": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The URL for this dynamic route server member.",
												},
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique identifier for this dynamic route server member.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name for this dynamic route server member. The name is unique across all members in the dynamic route server.",
												},
												"resource_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The resource type.",
												},
												"virtual_network_interfaces": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The virtual network interfaces for this dynamic route server member.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"crn": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The CRN for this virtual network interface.",
															},
															"deleted": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"more_info": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "A link to documentation about deleted resources.",
																		},
																	},
																},
															},
															"href": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The URL for this virtual network interface.",
															},
															"id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The unique identifier for this virtual network interface.",
															},
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.",
															},
															"primary_ip": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The primary IP for this virtual network interface.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"address": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.",
																		},
																		"deleted": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"more_info": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "A link to documentation about deleted resources.",
																					},
																				},
																			},
																		},
																		"href": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The URL for this reserved IP.",
																		},
																		"id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The unique identifier for this reserved IP.",
																		},
																		"name": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The name for this reserved IP. The name is unique across all reserved IPs in a subnet.",
																		},
																		"resource_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The resource type.",
																		},
																	},
																},
															},
															"resource_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The resource type.",
															},
															"subnet": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The associated subnet.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"crn": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The CRN for this subnet.",
																		},
																		"deleted": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"more_info": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "A link to documentation about deleted resources.",
																					},
																				},
																			},
																		},
																		"href": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The URL for this subnet.",
																		},
																		"id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The unique identifier for this subnet.",
																		},
																		"name": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The name for this subnet. The name is unique across all subnets in the VPC.",
																		},
																		"resource_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The resource type.",
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"protocol_state": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The state of the routing protocol with this dynamic route server peer group peer. The states follow the conventions defined in [RFC 4274](https://datatracker.ietf.org/doc/html/rfc4274#section-2.3).- `initializing`: The BGP session is being initialized.",
									},
									"remote": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The remote peer for this BGP session.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name for this dynamic route server peer group peer. The name must not be used by another peer in the peer group.",
												},
											},
										},
									},
								},
							},
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The administrative state used for this peer group peer:- `disabled`: The peer group peer is disabled, and the dynamic route server members will  not establish a BGP session with this peer group peer.- `enabled`: The peer group peer is enabled, and the dynamic route server members will  try to establish a BGP session with this peer group peer.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMIsDynamicRouteServerPeerGroupPeersRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_peers", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listDynamicRouteServerPeerGroupPeersOptions := &vpcv1.ListDynamicRouteServerPeerGroupPeersOptions{}

	listDynamicRouteServerPeerGroupPeersOptions.SetDynamicRouteServerID(d.Get("dynamic_route_server_id").(string))
	listDynamicRouteServerPeerGroupPeersOptions.SetDynamicRouteServerPeerGroupID(d.Get("dynamic_route_server_peer_group_id").(string))

	var pager *vpcv1.DynamicRouteServerPeerGroupPeersPager
	pager, err = vpcClient.NewDynamicRouteServerPeerGroupPeersPager(listDynamicRouteServerPeerGroupPeersOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_peers", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allItems, err := pager.GetAll()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DynamicRouteServerPeerGroupPeersPager.GetAll() failed %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_peers", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMIsDynamicRouteServerPeerGroupPeersID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allItems {
		modelMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerToMap(modelItem)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_peers", "read", "DynamicRouteServers-to-map").GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("peers", mapSlice); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting peers %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_peers", "read", "peers-set").GetDiag()
	}

	return nil
}

// dataSourceIBMIsDynamicRouteServerPeerGroupPeersID returns a reasonable ID for the list.
func dataSourceIBMIsDynamicRouteServerPeerGroupPeersID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerToMap(model vpcv1.DynamicRouteServerPeerGroupPeerIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddress); ok {
		return DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddress))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway); ok {
		return DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewayToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPeer); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.DynamicRouteServerPeerGroupPeer)
		if model.CreatedAt != nil {
			modelMap["created_at"] = model.CreatedAt.String()
		}
		modelMap["creator"] = *model.Creator
		modelMap["href"] = *model.Href
		modelMap["id"] = *model.ID
		lifecycleReasons := []map[string]interface{}{}
		for _, lifecycleReasonsItem := range model.LifecycleReasons {
			lifecycleReasonsItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			lifecycleReasons = append(lifecycleReasons, lifecycleReasonsItemMap)
		}
		modelMap["lifecycle_reasons"] = lifecycleReasons
		if model.LifecycleState != nil {
			modelMap["lifecycle_state"] = *model.LifecycleState
		}
		modelMap["name"] = *model.Name
		modelMap["resource_type"] = *model.ResourceType
		if model.Status != nil {
			modelMap["status"] = *model.Status
		}
		if model.StatusReasons != nil {
			statusReasons := []map[string]interface{}{}
			for _, statusReasonsItem := range model.StatusReasons {
				statusReasonsItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerStatusReasonToMap(&statusReasonsItem) // #nosec G601
				if err != nil {
					return modelMap, err
				}
				statusReasons = append(statusReasons, statusReasonsItemMap)
			}
			modelMap["status_reasons"] = statusReasons
		}
		if model.Asn != nil {
			modelMap["asn"] = flex.IntValue(model.Asn)
		}
		if model.BidirectionalForwardingDetection != nil {
			bidirectionalForwardingDetectionMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(model.BidirectionalForwardingDetection)
			if err != nil {
				return modelMap, err
			}
			modelMap["bidirectional_forwarding_detection"] = []map[string]interface{}{bidirectionalForwardingDetectionMap}
		}
		if model.Endpoint != nil {
			endpointMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressEndpointToMap(model.Endpoint)
			if err != nil {
				return modelMap, err
			}
			modelMap["endpoint"] = []map[string]interface{}{endpointMap}
		}
		if model.Priority != nil {
			modelMap["priority"] = flex.IntValue(model.Priority)
		}
		if model.Sessions != nil {
			sessions := []map[string]interface{}{}
			for _, sessionsItem := range model.Sessions {
				sessionsItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressSessionToMap(&sessionsItem) // #nosec G601
				if err != nil {
					return modelMap, err
				}
				sessions = append(sessions, sessionsItemMap)
			}
			modelMap["sessions"] = sessions
		}
		if model.State != nil {
			modelMap["state"] = *model.State
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupPeerIntf subtype encountered")
	}
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersLifecycleReasonToMap(model *vpcv1.LifecycleReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerStatusReasonToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerStatusReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetection) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["detect_multiplier"] = flex.IntValue(model.DetectMultiplier)
	modelMap["enabled"] = *model.Enabled
	modelMap["mode"] = *model.Mode
	modelMap["receive_interval"] = flex.IntValue(model.ReceiveInterval)
	modelMap["role"] = *model.Role
	sessions := []map[string]interface{}{}
	for _, sessionsItem := range model.Sessions {
		sessionsItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionToMap(&sessionsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		sessions = append(sessions, sessionsItemMap)
	}
	modelMap["sessions"] = sessions
	modelMap["transmit_interval"] = flex.IntValue(model.TransmitInterval)
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSession) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	localMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerMemberReferenceToMap(model.Local)
	if err != nil {
		return modelMap, err
	}
	modelMap["local"] = []map[string]interface{}{localMap}
	remoteMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerReferenceToMap(model.Remote)
	if err != nil {
		return modelMap, err
	}
	modelMap["remote"] = []map[string]interface{}{remoteMap}
	modelMap["state"] = *model.State
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerMemberReferenceToMap(model *vpcv1.DynamicRouteServerMemberReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	virtualNetworkInterfaces := []map[string]interface{}{}
	for _, virtualNetworkInterfacesItem := range model.VirtualNetworkInterfaces {
		virtualNetworkInterfacesItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersVirtualNetworkInterfaceReferenceToMap(&virtualNetworkInterfacesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		virtualNetworkInterfaces = append(virtualNetworkInterfaces, virtualNetworkInterfacesItemMap)
	}
	modelMap["virtual_network_interfaces"] = virtualNetworkInterfaces
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersVirtualNetworkInterfaceReferenceToMap(model *vpcv1.VirtualNetworkInterfaceReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	primaryIPMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersReservedIPReferenceToMap(model.PrimaryIP)
	if err != nil {
		return modelMap, err
	}
	modelMap["primary_ip"] = []map[string]interface{}{primaryIPMap}
	modelMap["resource_type"] = *model.ResourceType
	subnetMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersSubnetReferenceToMap(model.Subnet)
	if err != nil {
		return modelMap, err
	}
	modelMap["subnet"] = []map[string]interface{}{subnetMap}
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersReservedIPReferenceToMap(model *vpcv1.ReservedIPReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersSubnetReferenceToMap(model *vpcv1.SubnetReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerReferenceToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressEndpointToMap(model vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGateway); ok {
		return DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGateway))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointByReservedIP); ok {
		return DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressEndpointByReservedIPToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointByReservedIP))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpoint); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpoint)
		if model.Address != nil {
			modelMap["address"] = *model.Address
		}
		if model.Gateway != nil {
			gatewayMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(model.Gateway)
			if err != nil {
				return modelMap, err
			}
			modelMap["gateway"] = []map[string]interface{}{gatewayMap}
		}
		if model.Deleted != nil {
			deletedMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDeletedToMap(model.Deleted)
			if err != nil {
				return modelMap, err
			}
			modelMap["deleted"] = []map[string]interface{}{deletedMap}
		}
		if model.Href != nil {
			modelMap["href"] = *model.Href
		}
		if model.ID != nil {
			modelMap["id"] = *model.ID
		}
		if model.Name != nil {
			modelMap["name"] = *model.Name
		}
		if model.ResourceType != nil {
			modelMap["resource_type"] = *model.ResourceType
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIntf subtype encountered")
	}
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGateway) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	if model.Gateway != nil {
		gatewayMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(model.Gateway)
		if err != nil {
			return modelMap, err
		}
		modelMap["gateway"] = []map[string]interface{}{gatewayMap}
	}
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressEndpointByReservedIPToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointByReservedIP) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	gatewayMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersIPToMap(model.Gateway)
	if err != nil {
		return modelMap, err
	}
	modelMap["gateway"] = []map[string]interface{}{gatewayMap}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersIPToMap(model *vpcv1.IP) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressSessionToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressSession) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["established_at"] = model.EstablishedAt.String()
	localMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerMemberReferenceToMap(model.Local)
	if err != nil {
		return modelMap, err
	}
	modelMap["local"] = []map[string]interface{}{localMap}
	modelMap["protocol_state"] = *model.ProtocolState
	remoteMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerNameToMap(model.Remote)
	if err != nil {
		return modelMap, err
	}
	modelMap["remote"] = []map[string]interface{}{remoteMap}
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerNameToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerName) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddress) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	lifecycleReasons := []map[string]interface{}{}
	for _, lifecycleReasonsItem := range model.LifecycleReasons {
		lifecycleReasonsItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		lifecycleReasons = append(lifecycleReasons, lifecycleReasonsItemMap)
	}
	modelMap["lifecycle_reasons"] = lifecycleReasons
	if model.LifecycleState != nil {
		modelMap["lifecycle_state"] = *model.LifecycleState
	}
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	if model.StatusReasons != nil {
		statusReasons := []map[string]interface{}{}
		for _, statusReasonsItem := range model.StatusReasons {
			statusReasonsItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerStatusReasonToMap(&statusReasonsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			statusReasons = append(statusReasons, statusReasonsItemMap)
		}
		modelMap["status_reasons"] = statusReasons
	}
	modelMap["asn"] = flex.IntValue(model.Asn)
	if model.BidirectionalForwardingDetection != nil {
		bidirectionalForwardingDetectionMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(model.BidirectionalForwardingDetection)
		if err != nil {
			return modelMap, err
		}
		modelMap["bidirectional_forwarding_detection"] = []map[string]interface{}{bidirectionalForwardingDetectionMap}
	}
	modelMap["creator"] = *model.Creator
	endpointMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressEndpointToMap(model.Endpoint)
	if err != nil {
		return modelMap, err
	}
	modelMap["endpoint"] = []map[string]interface{}{endpointMap}
	modelMap["priority"] = flex.IntValue(model.Priority)
	sessions := []map[string]interface{}{}
	for _, sessionsItem := range model.Sessions {
		sessionsItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerAddressSessionToMap(&sessionsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		sessions = append(sessions, sessionsItemMap)
	}
	modelMap["sessions"] = sessions
	modelMap["state"] = *model.State
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewayToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	lifecycleReasons := []map[string]interface{}{}
	for _, lifecycleReasonsItem := range model.LifecycleReasons {
		lifecycleReasonsItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		lifecycleReasons = append(lifecycleReasons, lifecycleReasonsItemMap)
	}
	modelMap["lifecycle_reasons"] = lifecycleReasons
	if model.LifecycleState != nil {
		modelMap["lifecycle_state"] = *model.LifecycleState
	}
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	if model.StatusReasons != nil {
		statusReasons := []map[string]interface{}{}
		for _, statusReasonsItem := range model.StatusReasons {
			statusReasonsItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerStatusReasonToMap(&statusReasonsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			statusReasons = append(statusReasons, statusReasonsItemMap)
		}
		modelMap["status_reasons"] = statusReasons
	}
	modelMap["creator"] = *model.Creator
	endpointMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewayEndpointToMap(model.Endpoint)
	if err != nil {
		return modelMap, err
	}
	modelMap["endpoint"] = []map[string]interface{}{endpointMap}
	sessions := []map[string]interface{}{}
	for _, sessionsItem := range model.Sessions {
		sessionsItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewaySessionToMap(&sessionsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		sessions = append(sessions, sessionsItemMap)
	}
	modelMap["sessions"] = sessions
	modelMap["state"] = *model.State
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewayEndpointToMap(model vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReference); ok {
		return DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReferenceToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReference))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpoint); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpoint)
		if model.CRN != nil {
			modelMap["crn"] = *model.CRN
		}
		if model.ID != nil {
			modelMap["id"] = *model.ID
		}
		if model.ResourceType != nil {
			modelMap["resource_type"] = *model.ResourceType
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointIntf subtype encountered")
	}
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReferenceToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	modelMap["id"] = *model.ID
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewaySessionToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewaySession) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["established_at"] = model.EstablishedAt.String()
	localMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerMemberReferenceToMap(model.Local)
	if err != nil {
		return modelMap, err
	}
	modelMap["local"] = []map[string]interface{}{localMap}
	modelMap["protocol_state"] = *model.ProtocolState
	remoteMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteToMap(model.Remote)
	if err != nil {
		return modelMap, err
	}
	modelMap["remote"] = []map[string]interface{}{remoteMap}
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPeersDynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewaySessionRemote) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = *model.Name
	return modelMap, nil
}

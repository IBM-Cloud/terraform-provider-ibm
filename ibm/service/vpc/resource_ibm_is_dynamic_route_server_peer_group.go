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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func ResourceIBMIsDynamicRouteServerPeerGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsDynamicRouteServerPeerGroupCreate,
		ReadContext:   resourceIBMIsDynamicRouteServerPeerGroupRead,
		UpdateContext: resourceIBMIsDynamicRouteServerPeerGroupUpdate,
		DeleteContext: resourceIBMIsDynamicRouteServerPeerGroupDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"dynamic_route_server_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group", "dynamic_route_server_id"),
				Description:  "The dynamic route server identifier.",
			},
			"health_monitor": &schema.Schema{
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				Description: "The health monitoring configuration used for this dynamic route server peer group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The mode used for this health monitor:- `bfd`: Both BGP monitoring and Bidirectional Forwarding Detection (BFD)  are enabled. When a peer in this peer group becomes unreachable, routes  whose `next_hop` is this peer are automatically withdrawn. BFD provides  faster failure detection.- `bgp`: Only BGP monitoring is enabled. When a peer in this peer group  becomes unreachable, routes whose `next_hop` is this peer are automatically  withdrawn.- `none`: Monitoring is disabled. When a peer in this peer group becomes  unreachable, routes whose `next_hop` is this peer are not automatically  withdrawn.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
						},
						"roles": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The `roles` used for this dynamic route server peer group health monitor configuration:- `next_hop`: Health monitor configuration used for `next_hop` role peers  of this dynamic route server peer group.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group", "name"),
				Description:  "The name for this dynamic route server peer group.",
			},
			"peers": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Description: "A collection of peers of the same type having a similar network topology grouped together to apply identical policies and maintain consistent behavior.",
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
										Optional:    true,
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
							Required:    true,
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
										Optional:    true,
										Computed:    true,
										Description: "A link to documentation about this status reason.",
									},
								},
							},
						},
						"asn": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
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
																Optional:    true,
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
																			Optional:    true,
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
																						Optional:    true,
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
																						Optional:    true,
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
																Optional:    true,
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
							MaxItems:    1,
							Optional:    true,
							Description: "The endpoint for a dynamic route server peer group address peer.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The IP address.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.",
									},
									"gateway": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "The gateway IP address of the dynamic route server peer group address peer endpoint.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
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
										Optional:    true,
										Description: "The URL for this reserved IP.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
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
							Optional:    true,
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
													Optional:    true,
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
																Optional:    true,
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
																			Optional:    true,
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
																			Optional:    true,
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
													Optional:    true,
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
							Optional:    true,
							Description: "The administrative state used for this peer group peer:- `disabled`: The peer group peer is disabled, and the dynamic route server members will  not establish a BGP session with this peer group peer.- `enabled`: The peer group peer is enabled, and the dynamic route server members will  try to establish a BGP session with this peer group peer.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
						},
					},
				},
			},
			"prefix_watcher": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "The prefix watcher configuration used for this dynamic route server peer group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"excluded_prefixes": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "The prefixes that are excluded from `monitored_prefixes` for this prefix watcher.If empty, no prefixes are excluded from `monitored_prefixes`.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ge": &schema.Schema{
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The minimum matched prefix length. If non-zero, only routes within the `prefix` that have a prefix length greater or equal to this value are excluded.If zero, `ge` matching is not applied.",
									},
									"le": &schema.Schema{
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The maximum matched prefix length. If non-zero, only routes within the `prefix` that have a prefix length less than or equal to this value are excluded.If zero, `le` matching is not applied.",
									},
									"prefix": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The prefix excluded from `monitored_prefixes`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 address prefixes in the future.",
									},
								},
							},
						},
						"monitored_prefixes": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "The prefixes that are monitored by this prefix watcher.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ge": &schema.Schema{
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The minimum prefix length to match. If non-zero, the prefix watcher matches only routes within the prefix that have a prefix length greater or equal to this value.If zero, `ge` matching is not applied.",
									},
									"le": &schema.Schema{
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The maximum prefix length to match. If non-zero, the prefix watcher matches only routes within the prefix that have a prefix length less than or equal to this value.If zero, `le` matching is not applied.",
									},
									"prefix": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The IP address prefix matched by this prefix watcher.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 address prefixes in the future.",
									},
								},
							},
						},
					},
				},
			},
			"roles": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Description: "The roles assigned to peers in this dynamic route server peer group: - `next_hop`: The peer can serve as a next hop for routing traffic. - `router`: The peer can participate in route exchange and route filtering. The enumerated values for this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"state": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "enabled",
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group", "state"),
				Description:  "The administrative state used for the dynamic route server peer group:- `disabled`: The peer group is disabled, and the dynamic route server members will  not establish a BGP session with peers in this peer group.- `enabled`: The peer group is enabled, and the dynamic route server members will  try to establish a BGP session with peers in this peer group.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that this dynamic route server peer group was created.",
			},
			"creator": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The type of resource that created this dynamic route server peer group:  - `dynamic_route_server`: dynamic router server created peer group.  - `transit_gateway`:  transit gateway created peer group.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this dynamic route server peer group.",
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
							Optional:    true,
							Computed:    true,
							Description: "A link to documentation about the reason for this lifecycle state.",
						},
					},
				},
			},
			"lifecycle_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the dynamic route server peer group.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The aggregate session status of the peers in this dynamic route server peer group:  - `degraded`: One or more dynamic route server peer group peers    were unable to establish a session.  - `down`: All dynamic route server peer group peers were unable    to establish a session.  - `initializing`: The dynamic route server peer group is in the process of establishing     sessions with its peers.  - `unknown`: The aggregate session status of the dynamic route server peer group     peers could not be determined.  - `up`: All dynamic route server peer group peers have established    sessions and are operating normally.",
			},
			"status_reasons": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current dynamic route server peer group connection status (if any).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the status reason.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
						},
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the status reason.",
						},
						"more_info": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "A link to documentation about this status reason.",
						},
					},
				},
			},
			"is_dynamic_route_server_peer_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID for this dynamic route server peer group peer.",
			},
			"etag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceIBMIsDynamicRouteServerPeerGroupValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "dynamic_route_server_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[-0-9a-z_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
		validate.ValidateSchema{
			Identifier:                 "state",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "disabled, enabled",
			Regexp:                     `^[a-z][a-z0-9]*(_[a-z0-9]+)*$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_dynamic_route_server_peer_group", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsDynamicRouteServerPeerGroupCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createDynamicRouteServerPeerGroupOptions := &vpcv1.CreateDynamicRouteServerPeerGroupOptions{}

	createDynamicRouteServerPeerGroupOptions.SetDynamicRouteServerID(d.Get("dynamic_route_server_id").(string))
	healthMonitorModel, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPrototype(d.Get("health_monitor.0").(map[string]interface{}))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "create", "parse-health_monitor").GetDiag()
	}
	createDynamicRouteServerPeerGroupOptions.SetHealthMonitor(healthMonitorModel)
	var peers []vpcv1.DynamicRouteServerPeerGroupPeerPrototypeIntf
	for _, v := range d.Get("peers").([]interface{}) {
		value := v.(map[string]interface{})
		peersItem, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPeerPrototype(value)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "create", "parse-peers").GetDiag()
		}
		peers = append(peers, peersItem)
	}
	createDynamicRouteServerPeerGroupOptions.SetPeers(peers)
	var roles []string
	for _, v := range d.Get("roles").([]interface{}) {
		rolesItem := v.(string)
		roles = append(roles, rolesItem)
	}
	createDynamicRouteServerPeerGroupOptions.SetRoles(roles)
	if _, ok := d.GetOk("name"); ok {
		createDynamicRouteServerPeerGroupOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("prefix_watcher"); ok {
		prefixWatcherModel, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherPrototype(d.Get("prefix_watcher.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "create", "parse-prefix_watcher").GetDiag()
		}
		createDynamicRouteServerPeerGroupOptions.SetPrefixWatcher(prefixWatcherModel)
	}
	if _, ok := d.GetOk("state"); ok {
		createDynamicRouteServerPeerGroupOptions.SetState(d.Get("state").(string))
	}

	dynamicRouteServerPeerGroup, _, err := vpcClient.CreateDynamicRouteServerPeerGroupWithContext(context, createDynamicRouteServerPeerGroupOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateDynamicRouteServerPeerGroupWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *createDynamicRouteServerPeerGroupOptions.DynamicRouteServerID, *dynamicRouteServerPeerGroup.ID))

	return resourceIBMIsDynamicRouteServerPeerGroupRead(context, d, meta)
}

func resourceIBMIsDynamicRouteServerPeerGroupRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDynamicRouteServerPeerGroupOptions := &vpcv1.GetDynamicRouteServerPeerGroupOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "sep-id-parts").GetDiag()
	}

	getDynamicRouteServerPeerGroupOptions.SetDynamicRouteServerID(parts[0])
	getDynamicRouteServerPeerGroupOptions.SetID(parts[1])

	dynamicRouteServerPeerGroup, response, err := vpcClient.GetDynamicRouteServerPeerGroupWithContext(context, getDynamicRouteServerPeerGroupOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerPeerGroupWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	healthMonitorMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorToMap(dynamicRouteServerPeerGroup.HealthMonitor)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "health_monitor-to-map").GetDiag()
	}
	if err = d.Set("health_monitor", []map[string]interface{}{healthMonitorMap}); err != nil {
		err = fmt.Errorf("Error setting health_monitor: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "set-health_monitor").GetDiag()
	}
	if !core.IsNil(dynamicRouteServerPeerGroup.Name) {
		if err = d.Set("name", dynamicRouteServerPeerGroup.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "set-name").GetDiag()
		}
	}
	peers := []map[string]interface{}{}
	for _, peersItem := range dynamicRouteServerPeerGroup.Peers {
		peersItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerToMap(peersItem)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "peers-to-map").GetDiag()
		}
		peers = append(peers, peersItemMap)
	}
	if err = d.Set("peers", peers); err != nil {
		err = fmt.Errorf("Error setting peers: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "set-peers").GetDiag()
	}
	if !core.IsNil(dynamicRouteServerPeerGroup.PrefixWatcher) {
		prefixWatcherMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherToMap(dynamicRouteServerPeerGroup.PrefixWatcher)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "prefix_watcher-to-map").GetDiag()
		}
		if err = d.Set("prefix_watcher", []map[string]interface{}{prefixWatcherMap}); err != nil {
			err = fmt.Errorf("Error setting prefix_watcher: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "set-prefix_watcher").GetDiag()
		}
	}
	if err = d.Set("roles", dynamicRouteServerPeerGroup.Roles); err != nil {
		err = fmt.Errorf("Error setting roles: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "set-roles").GetDiag()
	}
	if !core.IsNil(dynamicRouteServerPeerGroup.State) {
		if err = d.Set("state", dynamicRouteServerPeerGroup.State); err != nil {
			err = fmt.Errorf("Error setting state: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "set-state").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(dynamicRouteServerPeerGroup.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("creator", dynamicRouteServerPeerGroup.Creator); err != nil {
		err = fmt.Errorf("Error setting creator: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "set-creator").GetDiag()
	}
	if err = d.Set("href", dynamicRouteServerPeerGroup.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "set-href").GetDiag()
	}
	lifecycleReasons := []map[string]interface{}{}
	for _, lifecycleReasonsItem := range dynamicRouteServerPeerGroup.LifecycleReasons {
		lifecycleReasonsItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "lifecycle_reasons-to-map").GetDiag()
		}
		lifecycleReasons = append(lifecycleReasons, lifecycleReasonsItemMap)
	}
	if err = d.Set("lifecycle_reasons", lifecycleReasons); err != nil {
		err = fmt.Errorf("Error setting lifecycle_reasons: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "set-lifecycle_reasons").GetDiag()
	}
	if err = d.Set("lifecycle_state", dynamicRouteServerPeerGroup.LifecycleState); err != nil {
		err = fmt.Errorf("Error setting lifecycle_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "set-lifecycle_state").GetDiag()
	}
	if err = d.Set("resource_type", dynamicRouteServerPeerGroup.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set("status", dynamicRouteServerPeerGroup.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "set-status").GetDiag()
	}
	statusReasons := []map[string]interface{}{}
	for _, statusReasonsItem := range dynamicRouteServerPeerGroup.StatusReasons {
		statusReasonsItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupStatusReasonToMap(&statusReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "status_reasons-to-map").GetDiag()
		}
		statusReasons = append(statusReasons, statusReasonsItemMap)
	}
	if err = d.Set("status_reasons", statusReasons); err != nil {
		err = fmt.Errorf("Error setting status_reasons: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "set-status_reasons").GetDiag()
	}
	if err = d.Set("is_dynamic_route_server_peer_group_id", dynamicRouteServerPeerGroup.ID); err != nil {
		err = fmt.Errorf("Error setting is_dynamic_route_server_peer_group_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "read", "set-is_dynamic_route_server_peer_group_id").GetDiag()
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting etag: %s", err), "ibm_is_dynamic_route_server_peer_group", "read", "set-etag").GetDiag()
	}

	return nil
}

func resourceIBMIsDynamicRouteServerPeerGroupUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateDynamicRouteServerPeerGroupOptions := &vpcv1.UpdateDynamicRouteServerPeerGroupOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "update", "sep-id-parts").GetDiag()
	}

	updateDynamicRouteServerPeerGroupOptions.SetDynamicRouteServerID(parts[0])
	updateDynamicRouteServerPeerGroupOptions.SetID(parts[1])

	hasChange := false

	patchVals := &vpcv1.DynamicRouteServerPeerGroupPatch{}
	if d.HasChange("dynamic_route_server_id") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "dynamic_route_server_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_is_dynamic_route_server_peer_group", "update", "dynamic_route_server_id-forces-new").GetDiag()
	}
	if d.HasChange("health_monitor") {
		healthMonitor, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPatch(d.Get("health_monitor.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "update", "parse-health_monitor").GetDiag()
		}
		patchVals.HealthMonitor = healthMonitor
		hasChange = true
	}
	if d.HasChange("name") {
		newName := d.Get("name").(string)
		patchVals.Name = &newName
		hasChange = true
	}
	if d.HasChange("prefix_watcher") {
		prefixWatcher, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherPatch(d.Get("prefix_watcher.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "update", "parse-prefix_watcher").GetDiag()
		}
		patchVals.PrefixWatcher = prefixWatcher
		hasChange = true
	}
	if d.HasChange("state") {
		newState := d.Get("state").(string)
		patchVals.State = &newState
		hasChange = true
	}
	updateDynamicRouteServerPeerGroupOptions.SetIfMatch(d.Get("etag").(string))

	if hasChange {
		// Fields with `nil` values are omitted from the generic map,
		// so we need to re-add them to support removing arguments
		// in merge-patch operations sent to the service.
		updateDynamicRouteServerPeerGroupOptions.DynamicRouteServerPeerGroupPatch = ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPatchAsPatch(patchVals, d)

		_, _, err = vpcClient.UpdateDynamicRouteServerPeerGroupWithContext(context, updateDynamicRouteServerPeerGroupOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateDynamicRouteServerPeerGroupWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMIsDynamicRouteServerPeerGroupRead(context, d, meta)
}

func resourceIBMIsDynamicRouteServerPeerGroupDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteDynamicRouteServerPeerGroupOptions := &vpcv1.DeleteDynamicRouteServerPeerGroupOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group", "delete", "sep-id-parts").GetDiag()
	}

	deleteDynamicRouteServerPeerGroupOptions.SetDynamicRouteServerID(parts[0])
	deleteDynamicRouteServerPeerGroupOptions.SetID(parts[1])

	deleteDynamicRouteServerPeerGroupOptions.SetIfMatch(d.Get("etag").(string))

	_, _, err = vpcClient.DeleteDynamicRouteServerPeerGroupWithContext(context, deleteDynamicRouteServerPeerGroupOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteDynamicRouteServerPeerGroupWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPrototype(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerPeerGroupHealthMonitorPrototypeIntf, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupHealthMonitorPrototype{}
	if modelMap["mode"] != nil && modelMap["mode"].(string) != "" {
		model.Mode = core.StringPtr(modelMap["mode"].(string))
	}
	if modelMap["roles"] != nil {
		roles := []string{}
		for _, rolesItem := range modelMap["roles"].([]interface{}) {
			roles = append(roles, rolesItem.(string))
		}
		model.Roles = roles
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPrototypeDynamicRouteServerPeerGroupHealthMonitorNextHopPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupHealthMonitorPrototypeDynamicRouteServerPeerGroupHealthMonitorNextHopPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupHealthMonitorPrototypeDynamicRouteServerPeerGroupHealthMonitorNextHopPrototype{}
	if modelMap["mode"] != nil && modelMap["mode"].(string) != "" {
		model.Mode = core.StringPtr(modelMap["mode"].(string))
	}
	roles := []string{}
	for _, rolesItem := range modelMap["roles"].([]interface{}) {
		roles = append(roles, rolesItem.(string))
	}
	model.Roles = roles
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPrototypeDynamicRouteServerPeerGroupHealthMonitorRouterPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupHealthMonitorPrototypeDynamicRouteServerPeerGroupHealthMonitorRouterPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupHealthMonitorPrototypeDynamicRouteServerPeerGroupHealthMonitorRouterPrototype{}
	if modelMap["mode"] != nil && modelMap["mode"].(string) != "" {
		model.Mode = core.StringPtr(modelMap["mode"].(string))
	}
	roles := []string{}
	for _, rolesItem := range modelMap["roles"].([]interface{}) {
		roles = append(roles, rolesItem.(string))
	}
	model.Roles = roles
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPeerPrototype(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerPeerGroupPeerPrototypeIntf, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPeerPrototype{}
	if modelMap["asn"] != nil {
		model.Asn = core.Int64Ptr(int64(modelMap["asn"].(int)))
	}
	if modelMap["endpoint"] != nil && len(modelMap["endpoint"].([]interface{})) > 0 {
		EndpointModel, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(modelMap["endpoint"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Endpoint = EndpointModel
	}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["priority"] != nil {
		model.Priority = core.Int64Ptr(int64(modelMap["priority"].(int)))
	}
	if modelMap["state"] != nil && modelMap["state"].(string) != "" {
		model.State = core.StringPtr(modelMap["state"].(string))
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeIntf, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototype{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	if modelMap["address"] != nil && modelMap["address"].(string) != "" {
		model.Address = core.StringPtr(modelMap["address"].(string))
	}
	if modelMap["gateway"] != nil && len(modelMap["gateway"].([]interface{})) > 0 {
		GatewayModel, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway(modelMap["gateway"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Gateway = GatewayModel
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway{}
	model.Address = core.StringPtr(modelMap["address"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentity(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityIntf, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentity{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID{}
	model.ID = core.StringPtr(modelMap["id"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByHref(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByHref, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByHref{}
	model.Href = core.StringPtr(modelMap["href"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype{}
	model.Address = core.StringPtr(modelMap["address"].(string))
	if modelMap["gateway"] != nil && len(modelMap["gateway"].([]interface{})) > 0 {
		GatewayModel, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway(modelMap["gateway"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Gateway = GatewayModel
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPeerPrototypeDynamicRouteServerPeerGroupPeerAddressPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPeerPrototypeDynamicRouteServerPeerGroupPeerAddressPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPeerPrototypeDynamicRouteServerPeerGroupPeerAddressPrototype{}
	model.Asn = core.Int64Ptr(int64(modelMap["asn"].(int)))
	EndpointModel, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(modelMap["endpoint"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Endpoint = EndpointModel
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["priority"] != nil {
		model.Priority = core.Int64Ptr(int64(modelMap["priority"].(int)))
	}
	if modelMap["state"] != nil && modelMap["state"].(string) != "" {
		model.State = core.StringPtr(modelMap["state"].(string))
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPrefixWatcherPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPrefixWatcherPrototype{}
	if modelMap["excluded_prefixes"] != nil {
		excludedPrefixes := []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{}
		for _, excludedPrefixesItem := range modelMap["excluded_prefixes"].([]interface{}) {
			excludedPrefixesItemModel, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(excludedPrefixesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			excludedPrefixes = append(excludedPrefixes, *excludedPrefixesItemModel)
		}
		model.ExcludedPrefixes = excludedPrefixes
	}
	if modelMap["monitored_prefixes"] != nil {
		monitoredPrefixes := []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototype{}
		for _, monitoredPrefixesItem := range modelMap["monitored_prefixes"].([]interface{}) {
			monitoredPrefixesItemModel, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototype(monitoredPrefixesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			monitoredPrefixes = append(monitoredPrefixes, *monitoredPrefixesItemModel)
		}
		model.MonitoredPrefixes = monitoredPrefixes
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{}
	if modelMap["ge"] != nil {
		model.Ge = core.Int64Ptr(int64(modelMap["ge"].(int)))
	}
	if modelMap["le"] != nil {
		model.Le = core.Int64Ptr(int64(modelMap["le"].(int)))
	}
	model.Prefix = core.StringPtr(modelMap["prefix"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototype{}
	if modelMap["ge"] != nil {
		model.Ge = core.Int64Ptr(int64(modelMap["ge"].(int)))
	}
	if modelMap["le"] != nil {
		model.Le = core.Int64Ptr(int64(modelMap["le"].(int)))
	}
	if modelMap["prefix"] != nil && modelMap["prefix"].(string) != "" {
		model.Prefix = core.StringPtr(modelMap["prefix"].(string))
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupHealthMonitorPatch(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupHealthMonitorPatch, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupHealthMonitorPatch{}
	if modelMap["mode"] != nil && modelMap["mode"].(string) != "" {
		model.Mode = core.StringPtr(modelMap["mode"].(string))
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherPatch(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPrefixWatcherPatch, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPrefixWatcherPatch{}
	if modelMap["excluded_prefixes"] != nil {
		excludedPrefixes := []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{}
		for _, excludedPrefixesItem := range modelMap["excluded_prefixes"].([]interface{}) {
			excludedPrefixesItemModel, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(excludedPrefixesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			excludedPrefixes = append(excludedPrefixes, *excludedPrefixesItemModel)
		}
		model.ExcludedPrefixes = excludedPrefixes
	}
	if modelMap["monitored_prefixes"] != nil {
		monitoredPrefixes := []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototype{}
		for _, monitoredPrefixesItem := range modelMap["monitored_prefixes"].([]interface{}) {
			monitoredPrefixesItemModel, err := ResourceIBMIsDynamicRouteServerPeerGroupMapToDynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototype(monitoredPrefixesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			monitoredPrefixes = append(monitoredPrefixes, *monitoredPrefixesItemModel)
		}
		model.MonitoredPrefixes = monitoredPrefixes
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorToMap(model vpcv1.DynamicRouteServerPeerGroupHealthMonitorIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupHealthMonitorNextHop); ok {
		return ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorNextHopToMap(model.(*vpcv1.DynamicRouteServerPeerGroupHealthMonitorNextHop))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupHealthMonitorRouter); ok {
		return ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorRouterToMap(model.(*vpcv1.DynamicRouteServerPeerGroupHealthMonitorRouter))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupHealthMonitor); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.DynamicRouteServerPeerGroupHealthMonitor)
		if model.Mode != nil {
			modelMap["mode"] = *model.Mode
		}
		if model.Roles != nil {
			modelMap["roles"] = model.Roles
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupHealthMonitorIntf subtype encountered")
	}
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorNextHopToMap(model *vpcv1.DynamicRouteServerPeerGroupHealthMonitorNextHop) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["mode"] = *model.Mode
	modelMap["roles"] = model.Roles
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorRouterToMap(model *vpcv1.DynamicRouteServerPeerGroupHealthMonitorRouter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["mode"] = *model.Mode
	modelMap["roles"] = model.Roles
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerToMap(model vpcv1.DynamicRouteServerPeerGroupPeerIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddress); ok {
		return ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddress))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway); ok {
		return ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerTransitGatewayToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway))
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
			lifecycleReasonsItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
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
				statusReasonsItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerStatusReasonToMap(&statusReasonsItem) // #nosec G601
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
			bidirectionalForwardingDetectionMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(model.BidirectionalForwardingDetection)
			if err != nil {
				return modelMap, err
			}
			modelMap["bidirectional_forwarding_detection"] = []map[string]interface{}{bidirectionalForwardingDetectionMap}
		}
		if model.Endpoint != nil {
			endpointMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressEndpointToMap(model.Endpoint)
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
				sessionsItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressSessionToMap(&sessionsItem) // #nosec G601
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

func ResourceIBMIsDynamicRouteServerPeerGroupLifecycleReasonToMap(model *vpcv1.LifecycleReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerStatusReasonToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerStatusReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetection) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["detect_multiplier"] = flex.IntValue(model.DetectMultiplier)
	modelMap["enabled"] = *model.Enabled
	modelMap["mode"] = *model.Mode
	modelMap["receive_interval"] = flex.IntValue(model.ReceiveInterval)
	modelMap["role"] = *model.Role
	sessions := []map[string]interface{}{}
	for _, sessionsItem := range model.Sessions {
		sessionsItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionToMap(&sessionsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		sessions = append(sessions, sessionsItemMap)
	}
	modelMap["sessions"] = sessions
	modelMap["transmit_interval"] = flex.IntValue(model.TransmitInterval)
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSession) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	localMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerMemberReferenceToMap(model.Local)
	if err != nil {
		return modelMap, err
	}
	modelMap["local"] = []map[string]interface{}{localMap}
	remoteMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerReferenceToMap(model.Remote)
	if err != nil {
		return modelMap, err
	}
	modelMap["remote"] = []map[string]interface{}{remoteMap}
	modelMap["state"] = *model.State
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerMemberReferenceToMap(model *vpcv1.DynamicRouteServerMemberReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDeletedToMap(model.Deleted)
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
		virtualNetworkInterfacesItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupVirtualNetworkInterfaceReferenceToMap(&virtualNetworkInterfacesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		virtualNetworkInterfaces = append(virtualNetworkInterfaces, virtualNetworkInterfacesItemMap)
	}
	modelMap["virtual_network_interfaces"] = virtualNetworkInterfaces
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupVirtualNetworkInterfaceReferenceToMap(model *vpcv1.VirtualNetworkInterfaceReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	primaryIPMap, err := ResourceIBMIsDynamicRouteServerPeerGroupReservedIPReferenceToMap(model.PrimaryIP)
	if err != nil {
		return modelMap, err
	}
	modelMap["primary_ip"] = []map[string]interface{}{primaryIPMap}
	modelMap["resource_type"] = *model.ResourceType
	subnetMap, err := ResourceIBMIsDynamicRouteServerPeerGroupSubnetReferenceToMap(model.Subnet)
	if err != nil {
		return modelMap, err
	}
	modelMap["subnet"] = []map[string]interface{}{subnetMap}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupReservedIPReferenceToMap(model *vpcv1.ReservedIPReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerPeerGroupSubnetReferenceToMap(model *vpcv1.SubnetReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerReferenceToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressEndpointToMap(model vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGateway); ok {
		return ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGateway))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointByReservedIP); ok {
		return ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressEndpointByReservedIPToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointByReservedIP))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpoint); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpoint)
		if model.Address != nil {
			modelMap["address"] = *model.Address
		}
		if model.Gateway != nil {
			gatewayMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(model.Gateway)
			if err != nil {
				return modelMap, err
			}
			modelMap["gateway"] = []map[string]interface{}{gatewayMap}
		}
		if model.Deleted != nil {
			deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGateway) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	if model.Gateway != nil {
		gatewayMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(model.Gateway)
		if err != nil {
			return modelMap, err
		}
		modelMap["gateway"] = []map[string]interface{}{gatewayMap}
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressEndpointByReservedIPToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointByReservedIP) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	gatewayMap, err := ResourceIBMIsDynamicRouteServerPeerGroupIPToMap(model.Gateway)
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

func ResourceIBMIsDynamicRouteServerPeerGroupIPToMap(model *vpcv1.IP) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressSessionToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressSession) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["established_at"] = model.EstablishedAt.String()
	localMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerMemberReferenceToMap(model.Local)
	if err != nil {
		return modelMap, err
	}
	modelMap["local"] = []map[string]interface{}{localMap}
	modelMap["protocol_state"] = *model.ProtocolState
	remoteMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerNameToMap(model.Remote)
	if err != nil {
		return modelMap, err
	}
	modelMap["remote"] = []map[string]interface{}{remoteMap}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerNameToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerName) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddress) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	lifecycleReasons := []map[string]interface{}{}
	for _, lifecycleReasonsItem := range model.LifecycleReasons {
		lifecycleReasonsItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
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
			statusReasonsItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerStatusReasonToMap(&statusReasonsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			statusReasons = append(statusReasons, statusReasonsItemMap)
		}
		modelMap["status_reasons"] = statusReasons
	}
	modelMap["asn"] = flex.IntValue(model.Asn)
	if model.BidirectionalForwardingDetection != nil {
		bidirectionalForwardingDetectionMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(model.BidirectionalForwardingDetection)
		if err != nil {
			return modelMap, err
		}
		modelMap["bidirectional_forwarding_detection"] = []map[string]interface{}{bidirectionalForwardingDetectionMap}
	}
	modelMap["creator"] = *model.Creator
	endpointMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressEndpointToMap(model.Endpoint)
	if err != nil {
		return modelMap, err
	}
	modelMap["endpoint"] = []map[string]interface{}{endpointMap}
	modelMap["priority"] = flex.IntValue(model.Priority)
	sessions := []map[string]interface{}{}
	for _, sessionsItem := range model.Sessions {
		sessionsItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerAddressSessionToMap(&sessionsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		sessions = append(sessions, sessionsItemMap)
	}
	modelMap["sessions"] = sessions
	modelMap["state"] = *model.State
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerTransitGatewayToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerTransitGateway) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	lifecycleReasons := []map[string]interface{}{}
	for _, lifecycleReasonsItem := range model.LifecycleReasons {
		lifecycleReasonsItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
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
			statusReasonsItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerStatusReasonToMap(&statusReasonsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			statusReasons = append(statusReasons, statusReasonsItemMap)
		}
		modelMap["status_reasons"] = statusReasons
	}
	modelMap["creator"] = *model.Creator
	endpointMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerTransitGatewayEndpointToMap(model.Endpoint)
	if err != nil {
		return modelMap, err
	}
	modelMap["endpoint"] = []map[string]interface{}{endpointMap}
	sessions := []map[string]interface{}{}
	for _, sessionsItem := range model.Sessions {
		sessionsItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerTransitGatewaySessionToMap(&sessionsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		sessions = append(sessions, sessionsItemMap)
	}
	modelMap["sessions"] = sessions
	modelMap["state"] = *model.State
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerTransitGatewayEndpointToMap(model vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReference); ok {
		return ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReferenceToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReference))
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

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReferenceToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewayEndpointTransitGatewayReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	modelMap["id"] = *model.ID
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerTransitGatewaySessionToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewaySession) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["established_at"] = model.EstablishedAt.String()
	localMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerMemberReferenceToMap(model.Local)
	if err != nil {
		return modelMap, err
	}
	modelMap["local"] = []map[string]interface{}{localMap}
	modelMap["protocol_state"] = *model.ProtocolState
	remoteMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteToMap(model.Remote)
	if err != nil {
		return modelMap, err
	}
	modelMap["remote"] = []map[string]interface{}{remoteMap}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPeerTransitGatewaySessionRemoteToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerTransitGatewaySessionRemote) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = *model.Name
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherToMap(model *vpcv1.DynamicRouteServerPeerGroupPrefixWatcher) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	excludedPrefixes := []map[string]interface{}{}
	for _, excludedPrefixesItem := range model.ExcludedPrefixes {
		excludedPrefixesItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixToMap(&excludedPrefixesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		excludedPrefixes = append(excludedPrefixes, excludedPrefixesItemMap)
	}
	modelMap["excluded_prefixes"] = excludedPrefixes
	monitoredPrefixes := []map[string]interface{}{}
	for _, monitoredPrefixesItem := range model.MonitoredPrefixes {
		monitoredPrefixesItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixToMap(&monitoredPrefixesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		monitoredPrefixes = append(monitoredPrefixes, monitoredPrefixesItemMap)
	}
	modelMap["monitored_prefixes"] = monitoredPrefixes
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixToMap(model *vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefix) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["ge"] = flex.IntValue(model.Ge)
	modelMap["le"] = flex.IntValue(model.Le)
	modelMap["prefix"] = *model.Prefix
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixToMap(model *vpcv1.DynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefix) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["ge"] = flex.IntValue(model.Ge)
	modelMap["le"] = flex.IntValue(model.Le)
	modelMap["prefix"] = *model.Prefix
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupStatusReasonToMap(model *vpcv1.DynamicRouteServerPeerGroupStatusReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPatchAsPatch(patchVals *vpcv1.DynamicRouteServerPeerGroupPatch, d *schema.ResourceData) map[string]interface{} {
	patch, _ := patchVals.AsPatch()
	var path string

	path = "health_monitor"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["health_monitor"] = nil
	} else if exists && patch["health_monitor"] != nil {
		ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorPatchAsPatch(patch["health_monitor"].(map[string]interface{}), d, fmt.Sprintf("%s.0", path))
	} else if !exists {
		delete(patch, "health_monitor")
	}
	path = "name"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["name"] = nil
	} else if !exists {
		delete(patch, "name")
	}
	path = "prefix_watcher"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["prefix_watcher"] = nil
	} else if exists && patch["prefix_watcher"] != nil {
		ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherPatchAsPatch(patch["prefix_watcher"].(map[string]interface{}), d, fmt.Sprintf("%s.0", path))
	} else if !exists {
		delete(patch, "prefix_watcher")
	}
	path = "state"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["state"] = nil
	} else if !exists {
		delete(patch, "state")
	}

	return patch
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherPatchAsPatch(patch map[string]interface{}, d *schema.ResourceData, rootPath string) {
	var path string

	path = rootPath + ".excluded_prefixes"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["excluded_prefixes"] = nil
	} else if exists && patch["excluded_prefixes"] != nil {
		excluded_prefixesList := patch["excluded_prefixes"].([]map[string]interface{})
		for i, excluded_prefixesItem := range excluded_prefixesList {
			ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeAsPatch(excluded_prefixesItem, d, fmt.Sprintf("%s.%d", path, i))
		}
	} else if !exists {
		delete(patch, "excluded_prefixes")
	}
	path = rootPath + ".monitored_prefixes"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["monitored_prefixes"] = nil
	} else if exists && patch["monitored_prefixes"] != nil {
		monitored_prefixesList := patch["monitored_prefixes"].([]map[string]interface{})
		for i, monitored_prefixesItem := range monitored_prefixesList {
			ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototypeAsPatch(monitored_prefixesItem, d, fmt.Sprintf("%s.%d", path, i))
		}
	} else if !exists {
		delete(patch, "monitored_prefixes")
	}
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherMonitoredPrefixPrototypeAsPatch(patch map[string]interface{}, d *schema.ResourceData, rootPath string) {
	var path string

	path = rootPath + ".ge"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["ge"] = nil
	} else if !exists {
		delete(patch, "ge")
	}
	path = rootPath + ".le"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["le"] = nil
	} else if !exists {
		delete(patch, "le")
	}
	path = rootPath + ".prefix"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["prefix"] = nil
	} else if !exists {
		delete(patch, "prefix")
	}
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeAsPatch(patch map[string]interface{}, d *schema.ResourceData, rootPath string) {
	var path string

	path = rootPath + ".ge"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["ge"] = nil
	} else if !exists {
		delete(patch, "ge")
	}
	path = rootPath + ".le"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["le"] = nil
	} else if !exists {
		delete(patch, "le")
	}
}

func ResourceIBMIsDynamicRouteServerPeerGroupDynamicRouteServerPeerGroupHealthMonitorPatchAsPatch(patch map[string]interface{}, d *schema.ResourceData, rootPath string) {
	var path string

	path = rootPath + ".mode"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["mode"] = nil
	} else if !exists {
		delete(patch, "mode")
	}
}

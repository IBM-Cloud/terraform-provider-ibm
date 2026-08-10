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

func ResourceIBMIsDynamicRouteServerPeerGroupPeer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsDynamicRouteServerPeerGroupPeerCreate,
		ReadContext:   resourceIBMIsDynamicRouteServerPeerGroupPeerRead,
		UpdateContext: resourceIBMIsDynamicRouteServerPeerGroupPeerUpdate,
		DeleteContext: resourceIBMIsDynamicRouteServerPeerGroupPeerDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"dynamic_route_server_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_peer", "dynamic_route_server_id"),
				Description:  "The dynamic route server identifier.",
			},
			"dynamic_route_server_peer_group_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_peer", "dynamic_route_server_peer_group_id"),
				Description:  "The dynamic route server peer group identifier.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_peer", "name"),
				Description:  "The name for this dynamic route server peer group peer. The name must not be used by another peer in the peer group.",
			},
			"asn": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The autonomous system number (ASN) for this dynamic route server peer group address peer.",
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
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      2,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_peer", "priority"),
				Description:  "The priority of the peer group peer. The priority is used to determine the preferred path for routing. A lower value indicates a higher priority.",
			},
			"state": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "enabled",
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_peer", "state"),
				Description:  "The administrative state used for this peer group peer:- `disabled`: The peer group peer is disabled, and the dynamic route server members will  not establish a BGP session with this peer group peer.- `enabled`: The peer group peer is enabled, and the dynamic route server members will  try to establish a BGP session with this peer group peer.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
			},
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
			"is_dynamic_route_server_peer_group_peer_id": &schema.Schema{
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

func ResourceIBMIsDynamicRouteServerPeerGroupPeerValidator() *validate.ResourceValidator {
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
			Identifier:                 "dynamic_route_server_peer_group_id",
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
			Identifier:                 "priority",
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			Optional:                   true,
			MinValue:                   "0",
			MaxValue:                   "4",
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

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_dynamic_route_server_peer_group_peer", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsDynamicRouteServerPeerGroupPeerCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	bodyModelMap := map[string]interface{}{}
	createDynamicRouteServerPeerGroupPeerOptions := &vpcv1.CreateDynamicRouteServerPeerGroupPeerOptions{}

	if _, ok := d.GetOk("asn"); ok {
		bodyModelMap["asn"] = d.Get("asn")
	}
	if _, ok := d.GetOk("endpoint"); ok {
		bodyModelMap["endpoint"] = d.Get("endpoint")
	}
	if _, ok := d.GetOk("name"); ok {
		bodyModelMap["name"] = d.Get("name")
	}
	if _, ok := d.GetOk("priority"); ok {
		bodyModelMap["priority"] = d.Get("priority")
	}
	if _, ok := d.GetOk("state"); ok {
		bodyModelMap["state"] = d.Get("state")
	}
	createDynamicRouteServerPeerGroupPeerOptions.SetDynamicRouteServerID(d.Get("dynamic_route_server_id").(string))
	createDynamicRouteServerPeerGroupPeerOptions.SetDynamicRouteServerPeerGroupID(d.Get("dynamic_route_server_peer_group_id").(string))
	convertedModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototype(bodyModelMap)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "create", "parse-request-body").GetDiag()
	}
	createDynamicRouteServerPeerGroupPeerOptions.DynamicRouteServerPeerGroupPeerPrototype = convertedModel

	dynamicRouteServerPeerGroupPeerIntf, _, err := vpcClient.CreateDynamicRouteServerPeerGroupPeerWithContext(context, createDynamicRouteServerPeerGroupPeerOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateDynamicRouteServerPeerGroupPeerWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_peer", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerPeerGroupPeer := dynamicRouteServerPeerGroupPeerIntf.(*vpcv1.DynamicRouteServerPeerGroupPeer)
	d.SetId(fmt.Sprintf("%s/%s/%s", *createDynamicRouteServerPeerGroupPeerOptions.DynamicRouteServerID, *createDynamicRouteServerPeerGroupPeerOptions.DynamicRouteServerPeerGroupID, *dynamicRouteServerPeerGroupPeer.ID))

	return resourceIBMIsDynamicRouteServerPeerGroupPeerRead(context, d, meta)
}

func resourceIBMIsDynamicRouteServerPeerGroupPeerRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDynamicRouteServerPeerGroupPeerOptions := &vpcv1.GetDynamicRouteServerPeerGroupPeerOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "sep-id-parts").GetDiag()
	}

	getDynamicRouteServerPeerGroupPeerOptions.SetDynamicRouteServerID(parts[0])
	getDynamicRouteServerPeerGroupPeerOptions.SetDynamicRouteServerPeerGroupID(parts[1])
	getDynamicRouteServerPeerGroupPeerOptions.SetID(parts[2])

	dynamicRouteServerPeerGroupPeerIntf, response, err := vpcClient.GetDynamicRouteServerPeerGroupPeerWithContext(context, getDynamicRouteServerPeerGroupPeerOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerPeerGroupPeerWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_peer", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerPeerGroupPeer := dynamicRouteServerPeerGroupPeerIntf.(*vpcv1.DynamicRouteServerPeerGroupPeer)
	if !core.IsNil(dynamicRouteServerPeerGroupPeer.Name) {
		if err = d.Set("name", dynamicRouteServerPeerGroupPeer.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPeer.Asn) {
		if err = d.Set("asn", flex.IntValue(dynamicRouteServerPeerGroupPeer.Asn)); err != nil {
			err = fmt.Errorf("Error setting asn: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-asn").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPeer.Endpoint) {
		endpointMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointToMap(dynamicRouteServerPeerGroupPeer.Endpoint)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "endpoint-to-map").GetDiag()
		}
		if err = d.Set("endpoint", []map[string]interface{}{endpointMap}); err != nil {
			err = fmt.Errorf("Error setting endpoint: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-endpoint").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPeer.Priority) {
		if err = d.Set("priority", flex.IntValue(dynamicRouteServerPeerGroupPeer.Priority)); err != nil {
			err = fmt.Errorf("Error setting priority: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-priority").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPeer.State) {
		if err = d.Set("state", dynamicRouteServerPeerGroupPeer.State); err != nil {
			err = fmt.Errorf("Error setting state: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-state").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPeer.CreatedAt) {
		if err = d.Set("created_at", flex.DateTimeToString(dynamicRouteServerPeerGroupPeer.CreatedAt)); err != nil {
			err = fmt.Errorf("Error setting created_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-created_at").GetDiag()
		}
	}
	if err = d.Set("creator", dynamicRouteServerPeerGroupPeer.Creator); err != nil {
		err = fmt.Errorf("Error setting creator: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-creator").GetDiag()
	}
	if err = d.Set("href", dynamicRouteServerPeerGroupPeer.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-href").GetDiag()
	}
	lifecycleReasons := []map[string]interface{}{}
	for _, lifecycleReasonsItem := range dynamicRouteServerPeerGroupPeer.LifecycleReasons {
		lifecycleReasonsItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "lifecycle_reasons-to-map").GetDiag()
		}
		lifecycleReasons = append(lifecycleReasons, lifecycleReasonsItemMap)
	}
	if err = d.Set("lifecycle_reasons", lifecycleReasons); err != nil {
		err = fmt.Errorf("Error setting lifecycle_reasons: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-lifecycle_reasons").GetDiag()
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPeer.LifecycleState) {
		if err = d.Set("lifecycle_state", dynamicRouteServerPeerGroupPeer.LifecycleState); err != nil {
			err = fmt.Errorf("Error setting lifecycle_state: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-lifecycle_state").GetDiag()
		}
	}
	if err = d.Set("resource_type", dynamicRouteServerPeerGroupPeer.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-resource_type").GetDiag()
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPeer.Status) {
		if err = d.Set("status", dynamicRouteServerPeerGroupPeer.Status); err != nil {
			err = fmt.Errorf("Error setting status: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-status").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPeer.StatusReasons) {
		statusReasons := []map[string]interface{}{}
		for _, statusReasonsItem := range dynamicRouteServerPeerGroupPeer.StatusReasons {
			statusReasonsItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerStatusReasonToMap(&statusReasonsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "status_reasons-to-map").GetDiag()
			}
			statusReasons = append(statusReasons, statusReasonsItemMap)
		}
		if err = d.Set("status_reasons", statusReasons); err != nil {
			err = fmt.Errorf("Error setting status_reasons: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-status_reasons").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPeer.BidirectionalForwardingDetection) {
		bidirectionalForwardingDetectionMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(dynamicRouteServerPeerGroupPeer.BidirectionalForwardingDetection)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "bidirectional_forwarding_detection-to-map").GetDiag()
		}
		if err = d.Set("bidirectional_forwarding_detection", []map[string]interface{}{bidirectionalForwardingDetectionMap}); err != nil {
			err = fmt.Errorf("Error setting bidirectional_forwarding_detection: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-bidirectional_forwarding_detection").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPeer.Sessions) {
		sessions := []map[string]interface{}{}
		for _, sessionsItem := range dynamicRouteServerPeerGroupPeer.Sessions {
			sessionsItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressSessionToMap(&sessionsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "sessions-to-map").GetDiag()
			}
			sessions = append(sessions, sessionsItemMap)
		}
		if err = d.Set("sessions", sessions); err != nil {
			err = fmt.Errorf("Error setting sessions: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-sessions").GetDiag()
		}
	}
	if err = d.Set("is_dynamic_route_server_peer_group_peer_id", dynamicRouteServerPeerGroupPeer.ID); err != nil {
		err = fmt.Errorf("Error setting is_dynamic_route_server_peer_group_peer_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-is_dynamic_route_server_peer_group_peer_id").GetDiag()
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting etag: %s", err), "ibm_is_dynamic_route_server_peer_group_peer", "read", "set-etag").GetDiag()
	}

	return nil
}

func resourceIBMIsDynamicRouteServerPeerGroupPeerUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateDynamicRouteServerPeerGroupPeerOptions := &vpcv1.UpdateDynamicRouteServerPeerGroupPeerOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "update", "sep-id-parts").GetDiag()
	}

	updateDynamicRouteServerPeerGroupPeerOptions.SetDynamicRouteServerID(parts[0])
	updateDynamicRouteServerPeerGroupPeerOptions.SetDynamicRouteServerPeerGroupID(parts[1])
	updateDynamicRouteServerPeerGroupPeerOptions.SetID(parts[2])

	hasChange := false

	patchVals := &vpcv1.DynamicRouteServerPeerGroupPeerPatch{}
	if d.HasChange("dynamic_route_server_id") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "dynamic_route_server_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_is_dynamic_route_server_peer_group_peer", "update", "dynamic_route_server_id-forces-new").GetDiag()
	}
	if d.HasChange("dynamic_route_server_peer_group_id") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "dynamic_route_server_peer_group_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_is_dynamic_route_server_peer_group_peer", "update", "dynamic_route_server_peer_group_id-forces-new").GetDiag()
	}
	if d.HasChange("asn") {
		newAsn := int64(d.Get("asn").(int))
		patchVals.Asn = &newAsn
		hasChange = true
	}
	if d.HasChange("name") {
		newName := d.Get("name").(string)
		patchVals.Name = &newName
		hasChange = true
	}
	if d.HasChange("priority") {
		newPriority := int64(d.Get("priority").(int))
		patchVals.Priority = &newPriority
		hasChange = true
	}
	if d.HasChange("state") {
		newState := d.Get("state").(string)
		patchVals.State = &newState
		hasChange = true
	}
	updateDynamicRouteServerPeerGroupPeerOptions.SetIfMatch(d.Get("etag").(string))

	if hasChange {
		// Fields with `nil` values are omitted from the generic map,
		// so we need to re-add them to support removing arguments
		// in merge-patch operations sent to the service.
		updateDynamicRouteServerPeerGroupPeerOptions.DynamicRouteServerPeerGroupPeerPatch = ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerPatchAsPatch(patchVals, d)

		_, _, err = vpcClient.UpdateDynamicRouteServerPeerGroupPeerWithContext(context, updateDynamicRouteServerPeerGroupPeerOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateDynamicRouteServerPeerGroupPeerWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_peer", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMIsDynamicRouteServerPeerGroupPeerRead(context, d, meta)
}

func resourceIBMIsDynamicRouteServerPeerGroupPeerDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteDynamicRouteServerPeerGroupPeerOptions := &vpcv1.DeleteDynamicRouteServerPeerGroupPeerOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_peer", "delete", "sep-id-parts").GetDiag()
	}

	deleteDynamicRouteServerPeerGroupPeerOptions.SetDynamicRouteServerID(parts[0])
	deleteDynamicRouteServerPeerGroupPeerOptions.SetDynamicRouteServerPeerGroupID(parts[1])
	deleteDynamicRouteServerPeerGroupPeerOptions.SetID(parts[2])

	deleteDynamicRouteServerPeerGroupPeerOptions.SetIfMatch(d.Get("etag").(string))

	_, _, err = vpcClient.DeleteDynamicRouteServerPeerGroupPeerWithContext(context, deleteDynamicRouteServerPeerGroupPeerOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteDynamicRouteServerPeerGroupPeerWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_peer", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeIntf, error) {
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
		GatewayModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway(modelMap["gateway"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Gateway = GatewayModel
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway{}
	model.Address = core.StringPtr(modelMap["address"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentity(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityIntf, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentity{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByID{}
	model.ID = core.StringPtr(modelMap["id"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByHref(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByHref, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeReservedIPIdentityByHref{}
	model.Href = core.StringPtr(modelMap["href"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointPrototypeDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayPrototype{}
	model.Address = core.StringPtr(modelMap["address"].(string))
	if modelMap["gateway"] != nil && len(modelMap["gateway"].([]interface{})) > 0 {
		GatewayModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway(modelMap["gateway"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Gateway = GatewayModel
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototype(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerPeerGroupPeerPrototypeIntf, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPeerPrototype{}
	if modelMap["asn"] != nil {
		model.Asn = core.Int64Ptr(int64(modelMap["asn"].(int)))
	}
	if modelMap["endpoint"] != nil && len(modelMap["endpoint"].([]interface{})) > 0 {
		EndpointModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(modelMap["endpoint"].([]interface{})[0].(map[string]interface{}))
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

func ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerPrototypeDynamicRouteServerPeerGroupPeerAddressPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPeerPrototypeDynamicRouteServerPeerGroupPeerAddressPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPeerPrototypeDynamicRouteServerPeerGroupPeerAddressPrototype{}
	model.Asn = core.Int64Ptr(int64(modelMap["asn"].(int)))
	EndpointModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerMapToDynamicRouteServerPeerGroupPeerAddressEndpointPrototype(modelMap["endpoint"].([]interface{})[0].(map[string]interface{}))
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

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointToMap(model vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGateway); ok {
		return ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGateway))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointByReservedIP); ok {
		return ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointByReservedIPToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointByReservedIP))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpoint); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpoint)
		if model.Address != nil {
			modelMap["address"] = *model.Address
		}
		if model.Gateway != nil {
			gatewayMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(model.Gateway)
			if err != nil {
				return modelMap, err
			}
			modelMap["gateway"] = []map[string]interface{}{gatewayMap}
		}
		if model.Deleted != nil {
			deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGateway) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointIPGateway) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	if model.Gateway != nil {
		gatewayMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointIPGatewayGatewayToMap(model.Gateway)
		if err != nil {
			return modelMap, err
		}
		modelMap["gateway"] = []map[string]interface{}{gatewayMap}
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressEndpointByReservedIPToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressEndpointByReservedIP) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	gatewayMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerIPToMap(model.Gateway)
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

func ResourceIBMIsDynamicRouteServerPeerGroupPeerIPToMap(model *vpcv1.IP) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerLifecycleReasonToMap(model *vpcv1.LifecycleReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerStatusReasonToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerStatusReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetection) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["detect_multiplier"] = flex.IntValue(model.DetectMultiplier)
	modelMap["enabled"] = *model.Enabled
	modelMap["mode"] = *model.Mode
	modelMap["receive_interval"] = flex.IntValue(model.ReceiveInterval)
	modelMap["role"] = *model.Role
	sessions := []map[string]interface{}{}
	for _, sessionsItem := range model.Sessions {
		sessionsItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionToMap(&sessionsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		sessions = append(sessions, sessionsItemMap)
	}
	modelMap["sessions"] = sessions
	modelMap["transmit_interval"] = flex.IntValue(model.TransmitInterval)
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSessionToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressBidirectionalForwardingDetectionSession) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	localMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerMemberReferenceToMap(model.Local)
	if err != nil {
		return modelMap, err
	}
	modelMap["local"] = []map[string]interface{}{localMap}
	remoteMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerReferenceToMap(model.Remote)
	if err != nil {
		return modelMap, err
	}
	modelMap["remote"] = []map[string]interface{}{remoteMap}
	modelMap["state"] = *model.State
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerMemberReferenceToMap(model *vpcv1.DynamicRouteServerMemberReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(model.Deleted)
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
		virtualNetworkInterfacesItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerVirtualNetworkInterfaceReferenceToMap(&virtualNetworkInterfacesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		virtualNetworkInterfaces = append(virtualNetworkInterfaces, virtualNetworkInterfacesItemMap)
	}
	modelMap["virtual_network_interfaces"] = virtualNetworkInterfaces
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerVirtualNetworkInterfaceReferenceToMap(model *vpcv1.VirtualNetworkInterfaceReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	primaryIPMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerReservedIPReferenceToMap(model.PrimaryIP)
	if err != nil {
		return modelMap, err
	}
	modelMap["primary_ip"] = []map[string]interface{}{primaryIPMap}
	modelMap["resource_type"] = *model.ResourceType
	subnetMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerSubnetReferenceToMap(model.Subnet)
	if err != nil {
		return modelMap, err
	}
	modelMap["subnet"] = []map[string]interface{}{subnetMap}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerReservedIPReferenceToMap(model *vpcv1.ReservedIPReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerPeerGroupPeerSubnetReferenceToMap(model *vpcv1.SubnetReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerReferenceToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerAddressSessionToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerAddressSession) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["established_at"] = model.EstablishedAt.String()
	localMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerMemberReferenceToMap(model.Local)
	if err != nil {
		return modelMap, err
	}
	modelMap["local"] = []map[string]interface{}{localMap}
	modelMap["protocol_state"] = *model.ProtocolState
	remoteMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerNameToMap(model.Remote)
	if err != nil {
		return modelMap, err
	}
	modelMap["remote"] = []map[string]interface{}{remoteMap}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerNameToMap(model *vpcv1.DynamicRouteServerPeerGroupPeerName) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPeerDynamicRouteServerPeerGroupPeerPatchAsPatch(patchVals *vpcv1.DynamicRouteServerPeerGroupPeerPatch, d *schema.ResourceData) map[string]interface{} {
	patch, _ := patchVals.AsPatch()
	var path string

	path = "asn"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["asn"] = nil
	} else if !exists {
		delete(patch, "asn")
	}
	path = "name"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["name"] = nil
	} else if !exists {
		delete(patch, "name")
	}
	path = "priority"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["priority"] = nil
	} else if !exists {
		delete(patch, "priority")
	}
	path = "state"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["state"] = nil
	} else if !exists {
		delete(patch, "state")
	}

	return patch
}

// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/ScaleFT/sshkeys"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/crypto/ssh"
)

const (
	isInstancePEM                       = "private_key"
	isInstancePassphrase                = "passphrase"
	isInstanceInitPassword              = "password"
	isInstanceInitKeys                  = "keys"
	isInstanceNicPrimaryIP              = "primary_ip"
	isInstanceNicReservedIpAddress      = "address"
	isInstanceNicReservedIpHref         = "href"
	isInstanceNicReservedIpAutoDelete   = "auto_delete"
	isInstanceNicReservedIpName         = "name"
	isInstanceNicReservedIpId           = "reserved_ip"
	isInstanceNicReservedIpResourceType = "resource_type"

	isInstanceReservation           = "reservation"
	isReservationDeleted            = "deleted"
	isReservationDeletedMoreInfo    = "more_info"
	isReservationAffinity           = "reservation_affinity"
	isReservationAffinityPool       = "pool"
	isReservationAffinityPolicyResp = "policy"
)

func DataSourceIBMISInstance() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISInstanceRead,

		Schema: map[string]*schema.Schema{

			isInstanceAvailablePolicyHostFailure: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The availability policy to use for this virtual server instance. The action to perform if the compute host experiences a failure.",
			},

			isInstanceName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance name",
			},
			// cluster changes
			"cluster_network": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "If present, the cluster network that this virtual server instance resides in.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this cluster network.",
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
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this cluster network.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this cluster network.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this cluster network. The name must not be used by another cluster network in the region.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"cluster_network_attachments": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The cluster network attachments for this virtual server instance.The cluster network attachments are ordered for consistent instance configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this instance cluster network attachment.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this instance cluster network attachment.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this instance cluster network attachment. The name is unique across all network attachments for the instance.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"confidential_compute_mode": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The confidential compute mode to use for this virtual server instance.If unspecified, the default confidential compute mode from the profile will be used.",
			},
			"enable_secure_boot": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether secure boot is enabled for this virtual server instance.If unspecified, the default secure boot mode from the profile will be used.",
			},
			isInstanceMetadataServiceEnabled: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether the metadata service endpoint is available to the virtual server instance",
			},
			isInstanceMetadataService: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The metadata service configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceMetadataServiceEnabled1: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether the metadata service endpoint will be available to the virtual server instance",
						},

						isInstanceMetadataServiceProtocol: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The communication protocol to use for the metadata service endpoint. Applies only when the metadata service is enabled.",
						},

						isInstanceMetadataServiceRespHopLimit: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The hop limit (IP time to live) for IP response packets from the metadata service",
						},
					},
				},
			},
			isInstancePEM: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance Private Key file",
			},

			isInstancePassphrase: {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Passphrase for Instance Private Key file",
			},

			isInstanceInitPassword: {
				Type:        schema.TypeString,
				Sensitive:   true,
				Computed:    true,
				Description: "password for Windows Instance",
			},

			isInstanceInitKeys: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Instance keys",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance key id",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance key name",
						},
					},
				},
			},

			isInstanceVPC: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "VPC id",
			},

			isInstanceZone: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zone name",
			},

			isInstanceProfile: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Profile info",
			},

			isInstanceTotalVolumeBandwidth: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The amount of bandwidth (in megabits per second) allocated exclusively to instance storage volumes",
			},

			isInstanceBandwidth: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total bandwidth (in megabits per second) shared across the instance's network interfaces and storage volumes",
			},

			isInstanceTotalNetworkBandwidth: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The amount of bandwidth (in megabits per second) allocated exclusively to instance network interfaces.",
			},

			isInstanceLifecycleState: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the virtual server instance.",
			},
			isInstanceLifecycleReasons: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current lifecycle_state (if any).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceLifecycleReasonsCode: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the reason for this lifecycle state.",
						},

						isInstanceLifecycleReasonsMessage: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the reason for this lifecycle state.",
						},

						isInstanceLifecycleReasonsMoreInfo: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about the reason for this lifecycle state.",
						},
					},
				},
			},

			isInstanceTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "list of tags for the instance",
			},
			isInstanceAccessTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "list of access tags for the instance",
			},
			isInstanceBootVolume: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Instance Boot Volume",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Boot Volume id",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Boot Volume name",
						},
						"device": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Boot Volume device",
						},
						"volume_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Boot Volume's volume id",
						},
						"volume_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Boot Volume's volume name",
						},
						"volume_crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Boot Volume's volume CRN",
						},
					},
				},
			},

			isInstanceCatalogOffering: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The catalog offering or offering version to use when provisioning this virtual server instance. If an offering is specified, the latest version of that offering will be used. The specified offering or offering version may be in a different account in the same enterprise, subject to IAM policies.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceCatalogOfferingOfferingCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Identifies a catalog offering by a unique CRN property",
						},
						isInstanceCatalogOfferingVersionCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Identifies a version of a catalog offering by a unique CRN property",
						},
						isInstanceCatalogOfferingPlanCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this catalog offering version's billing plan",
						},
						"deleted": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and provides some supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
					},
				},
			},

			isInstanceVolumeAttachments: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Instance Volume Attachments",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Volume Attachment id",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Volume Attachment name",
						},
						"volume_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Boot Volume's volume id",
						},
						"volume_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Boot Volume's volume name",
						},
						"volume_crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Boot Volume's volume CRN",
						},
					},
				},
			},

			isInstancePrimaryNetworkInterface: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Primary Network interface info",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Primary Network Interface id",
						},
						isInstanceNicName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Primary Network Interface name",
						},
						isInstanceNicPortSpeed: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Instance Primary Network Interface port speed",
						},
						isInstanceNicPrimaryIpv4Address: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Primary Network Interface IPV4 Address",
						},
						isInstanceNicPrimaryIP: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isInstanceNicReservedIpAddress: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP address to reserve, which must not already be reserved on the subnet.",
									},
									isInstanceNicReservedIpHref: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this reserved IP",
									},
									isInstanceNicReservedIpName: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The user-defined name for this reserved IP. If unspecified, the name will be a hyphenated list of randomly-selected words. Names must be unique within the subnet the reserved IP resides in. ",
									},
									isInstanceNicReservedIpId: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Identifies a reserved IP by a unique property.",
									},
									isInstanceNicReservedIpResourceType: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type",
									},
								},
							},
						},
						isInstanceNicSecurityGroups: {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Set:         schema.HashString,
							Description: "Instance Primary Network Interface Security groups",
						},
						isInstanceNicSubnet: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Primary Network Interface subnet",
						},
					},
				},
			},

			"primary_network_attachment": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The primary network attachment for this virtual server instance.",
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
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this network attachment.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this network attachment.",
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"primary_ip": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The primary IP address of the virtual network interface for the network attachment.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.",
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
													Description: "Link to documentation about deleted resources.",
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
							Description: "The subnet of the virtual network interface for the network attachment.",
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
													Description: "Link to documentation about deleted resources.",
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
						"virtual_network_interface": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The virtual network interface for this instance network attachment.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this virtual network interface.",
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

			isInstanceNetworkInterfaces: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Instance Network interface info",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Network Interface id",
						},
						isInstanceNicName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Network Interface name",
						},
						isInstanceNicPrimaryIpv4Address: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Network Interface IPV4 Address",
						},
						isInstanceNicPrimaryIP: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isInstanceNicReservedIpAddress: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP address to reserve, which must not already be reserved on the subnet.",
									},
									isInstanceNicReservedIpHref: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this reserved IP",
									},
									isInstanceNicReservedIpName: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The user-defined name for this reserved IP. If unspecified, the name will be a hyphenated list of randomly-selected words. Names must be unique within the subnet the reserved IP resides in. ",
									},
									isInstanceNicReservedIpId: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Identifies a reserved IP by a unique property.",
									},
									isInstanceNicReservedIpResourceType: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type",
									},
								},
							},
						},
						isInstanceNicSecurityGroups: {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Set:         schema.HashString,
							Description: "Instance Network Interface Security Groups",
						},
						isInstanceNicSubnet: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Network Interface subnet",
						},
					},
				},
			},

			"network_attachments": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The network attachments for this virtual server instance, including the primary network attachment.",
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
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this network attachment.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this network attachment.",
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"primary_ip": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The primary IP address of the virtual network interface for the network attachment.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.",
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
													Description: "Link to documentation about deleted resources.",
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
							Description: "The subnet of the virtual network interface for the network attachment.",
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
													Description: "Link to documentation about deleted resources.",
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
						"virtual_network_interface": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The virtual network interface for this instance network attachment.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this virtual network interface.",
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

			isInstanceImage: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Instance Image",
			},

			isInstanceVolumes: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "List of volumes",
			},

			isInstanceResourceGroup: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Instance resource group",
			},

			isInstanceCPU: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Instance vCPU",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceCPUArch: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance vCPU Architecture",
						},
						isInstanceCPUCount: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Instance vCPU count",
						},
						// Added for AMD support, manufacturer details.
						isInstanceCPUManufacturer: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance vCPU Manufacturer",
						},
					},
				},
			},

			isInstanceGpu: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Instance GPU",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceGpuCount: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Instance GPU Count",
						},
						isInstanceGpuMemory: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Instance GPU Memory",
						},
						isInstanceGpuManufacturer: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance GPU Manufacturer",
						},
						isInstanceGpuModel: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance GPU Model",
						},
					},
				},
			},

			isInstanceMemory: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Instance memory",
			},

			"numa_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of NUMA nodes this virtual server instance is provisioned on. This property may be absent if the instance's `status` is not `running`.",
			},

			isInstanceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "instance status",
			},

			isInstanceStatusReasons: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current status (if any).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceStatusReasonsCode: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the status reason",
						},

						isInstanceStatusReasonsMessage: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the status reason",
						},

						isInstanceStatusReasonsMoreInfo: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about this status reason",
						},
					},
				},
			},

			flex.ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},

			flex.ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			flex.ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			IsInstanceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			flex.ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			flex.ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
			isInstanceDisks: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of the instance's disks.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that the disk was created.",
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this instance disk.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this instance disk.",
						},
						"interface_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The disk interface used for attaching the disk.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this disk.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
						"size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size of the disk in GB (gigabytes).",
						},
					},
				},
			},
			"placement_target": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The placement restrictions for the virtual server instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this dedicated host group.",
						},
						"deleted": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this dedicated host group.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this dedicated host group.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique user-defined name for this dedicated host group. If unspecified, the name will be a hyphenated list of randomly-selected words.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of resource referenced.",
						},
					},
				},
			},
			"health_reasons": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current health_state (if any).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the reason for this health state.",
						},
						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the reason for this health state.",
						},
						"more_info": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about the reason for this health state.",
						},
					},
				},
			},
			"health_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The health of this resource",
			},
			isInstanceReservation: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reservation used by this virtual server instance",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isReservationId: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this reservation.",
						},
						isReservationCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this reservation.",
						},
						isReservationName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this reservation. The name is unique across all reservations in the region.",
						},
						isReservationHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this reservation.",
						},
						isReservationResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
						isReservationDeleted: &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isReservationDeletedMoreInfo: &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
					},
				},
			},
			isReservationAffinity: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isReservationAffinityPolicyResp: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reservation affinity policy to use for this virtual server instance.",
						},
						isReservationAffinityPool: &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The pool of reservations available for use by this virtual server instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isReservationId: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this reservation.",
									},
									isReservationCrn: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this reservation.",
									},
									isReservationName: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name for this reservation. The name is unique across all reservations in the region.",
									},
									isReservationHref: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this reservation.",
									},
									isReservationResourceType: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type.",
									},
									isReservationDeleted: &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												isReservationDeletedMoreInfo: &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Link to documentation about deleted resources.",
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
		},
	}
}

func resourceIbmIsInstanceCatalogOfferingVersionPlanReferenceDeletedToMap(catalogOfferingVersionPlanReferenceDeleted vpcv1.Deleted) map[string]interface{} {
	catalogOfferingVersionPlanReferenceDeletedMap := map[string]interface{}{}

	catalogOfferingVersionPlanReferenceDeletedMap["more_info"] = catalogOfferingVersionPlanReferenceDeleted.MoreInfo

	return catalogOfferingVersionPlanReferenceDeletedMap
}

func dataSourceIBMISInstanceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	name := d.Get(isInstanceName).(string)

	err := instanceGetByName(context, d, meta, name)
	if err != nil {
		return err
	}
	return nil
}

func instanceGetByName(context context.Context, d *schema.ResourceData, meta interface{}, name string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	listInstancesOptions := &vpcv1.ListInstancesOptions{
		Name: &name,
	}

	instances, _, err := sess.ListInstancesWithContext(context, listInstancesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstancesWithContext failed: %s", err.Error()), "(Data) ibm_is_instance", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	allrecs := instances.Instances

	if len(allrecs) == 0 {
		err = fmt.Errorf("No Instance found with name %s", name)
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstancesWithContext failed: %s", err.Error()), "(Data) ibm_is_instance", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	instance := allrecs[0]
	d.SetId(*instance.ID)
	id := *instance.ID

	// cluster changes

	clusterNetwork := []map[string]interface{}{}
	if !core.IsNil(instance.ClusterNetwork) {
		clusterNetworkMap, err := DataSourceIBMIsInstanceClusterNetworkReferenceToMap(instance.ClusterNetwork)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance", "read", "cluster_network-to-map").GetDiag()
		}
		clusterNetwork = append(clusterNetwork, clusterNetworkMap)
	}
	if err = d.Set("cluster_network", clusterNetwork); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cluster_network: %s", err), "(Data) ibm_is_instance", "read", "set-cluster_network").GetDiag()
	}

	clusterNetworkAttachments := []map[string]interface{}{}
	for _, clusterNetworkAttachmentsItem := range instance.ClusterNetworkAttachments {
		clusterNetworkAttachmentsItemMap, err := DataSourceIBMIsInstanceInstanceClusterNetworkAttachmentReferenceToMap(&clusterNetworkAttachmentsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance", "read", "cluster_network_attachments-to-map").GetDiag()
		}
		clusterNetworkAttachments = append(clusterNetworkAttachments, clusterNetworkAttachmentsItemMap)
	}
	if err = d.Set("cluster_network_attachments", clusterNetworkAttachments); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cluster_network_attachments: %s", err), "(Data) ibm_is_instance", "read", "set-cluster_network_attachments").GetDiag()
	}

	// catalog
	if instance.CatalogOffering != nil {
		versionCrn := *instance.CatalogOffering.Version.CRN
		catalogList := make([]map[string]interface{}, 0)
		catalogMap := map[string]interface{}{}
		catalogMap[isInstanceCatalogOfferingVersionCrn] = versionCrn
		if instance.CatalogOffering.Plan != nil {
			if instance.CatalogOffering.Plan.CRN != nil && *instance.CatalogOffering.Plan.CRN != "" {
				catalogMap[isInstanceCatalogOfferingPlanCrn] = *instance.CatalogOffering.Plan.CRN
			}
			if instance.CatalogOffering.Plan.Deleted != nil {
				deletedMap := resourceIbmIsInstanceCatalogOfferingVersionPlanReferenceDeletedToMap(*instance.CatalogOffering.Plan.Deleted)
				catalogMap["deleted"] = []map[string]interface{}{deletedMap}
			}
		}
		catalogList = append(catalogList, catalogMap)
		if err = d.Set("catalog_offering", catalogList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting catalog_offering: %s", err), "(Data) ibm_is_instance", "read", "set-catalog_offering").GetDiag()
		}
	}

	if err = d.Set("name", instance.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_instance", "read", "set-name").GetDiag()
	}
	if instance.Profile != nil {
		if err = d.Set("profile", *instance.Profile.Name); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting profile: %s", err), "(Data) ibm_is_instance", "read", "set-profile").GetDiag()
		}
	}
	if instance.MetadataService != nil {
		if err = d.Set("metadata_service_enabled", instance.MetadataService.Enabled); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting metadata_service_enabled: %s", err), "(Data) ibm_is_instance", "read", "set-metadata_service_enabled").GetDiag()
		}
		metadataService := []map[string]interface{}{}
		metadataServiceMap := map[string]interface{}{}

		metadataServiceMap[isInstanceMetadataServiceEnabled1] = instance.MetadataService.Enabled
		if instance.MetadataService.Protocol != nil {
			metadataServiceMap[isInstanceMetadataServiceProtocol] = instance.MetadataService.Protocol
		}
		if instance.MetadataService.ResponseHopLimit != nil {
			metadataServiceMap[isInstanceMetadataServiceRespHopLimit] = instance.MetadataService.ResponseHopLimit
		}

		metadataService = append(metadataService, metadataServiceMap)
		if err = d.Set(isInstanceMetadataService, metadataService); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting metadata_service: %s", err), "(Data) ibm_is_instance", "read", "set-metadata_service").GetDiag()
		}
	}

	if instance.AvailabilityPolicy != nil && instance.AvailabilityPolicy.HostFailure != nil {
		if err = d.Set(isInstanceAvailablePolicyHostFailure, *instance.AvailabilityPolicy.HostFailure); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting availability_policy_host_failure: %s", err), "(Data) ibm_is_instance", "read", "set-availability_policy_host_failure").GetDiag()
		}
	}
	cpuList := make([]map[string]interface{}, 0)
	if instance.Vcpu != nil {
		currentCPU := map[string]interface{}{}
		currentCPU[isInstanceCPUArch] = *instance.Vcpu.Architecture
		currentCPU[isInstanceCPUCount] = *instance.Vcpu.Count
		currentCPU[isInstanceCPUManufacturer] = *instance.Vcpu.Manufacturer // Added for AMD support, manufacturer details.
		cpuList = append(cpuList, currentCPU)
	}
	if err = d.Set(isInstanceCPU, cpuList); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vcpu: %s", err), "(Data) ibm_is_instance", "read", "set-vcpu").GetDiag()
	}
	if instance.PlacementTarget != nil {
		placementTargetMap := resourceIbmIsInstanceInstancePlacementToMap(*instance.PlacementTarget.(*vpcv1.InstancePlacementTarget))
		if err = d.Set("placement_target", []map[string]interface{}{placementTargetMap}); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting placement_target: %s", err), "(Data) ibm_is_instance", "read", "set-placement_target").GetDiag()
		}
	}
	if err = d.Set(isInstanceMemory, *instance.Memory); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting memory: %s", err), "(Data) ibm_is_instance", "read", "set-memory").GetDiag()
	}
	if instance.NumaCount != nil {
		if err = d.Set("numa_count", *instance.NumaCount); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting numa_count: %s", err), "(Data) ibm_is_instance", "read", "set-numa_count").GetDiag()
		}
	}
	gpuList := make([]map[string]interface{}, 0)
	if instance.Gpu != nil {
		currentGpu := map[string]interface{}{}
		currentGpu[isInstanceGpuManufacturer] = instance.Gpu.Manufacturer
		currentGpu[isInstanceGpuModel] = instance.Gpu.Model
		currentGpu[isInstanceGpuCount] = instance.Gpu.Count
		currentGpu[isInstanceGpuMemory] = instance.Gpu.Memory
		gpuList = append(gpuList, currentGpu)
		if err = d.Set(isInstanceGpu, gpuList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting gpu: %s", err), "(Data) ibm_is_instance", "read", "set-gpu").GetDiag()
		}
	}
	if instance.Bandwidth != nil {
		if err = d.Set(isInstanceBandwidth, int(*instance.Bandwidth)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting bandwidth: %s", err), "(Data) ibm_is_instance", "read", "set-bandwidth").GetDiag()
		}
	}

	if instance.TotalNetworkBandwidth != nil {
		if err = d.Set("total_network_bandwidth", flex.IntValue(instance.TotalNetworkBandwidth)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting total_network_bandwidth: %s", err), "(Data) ibm_is_instance", "read", "set-total_network_bandwidth").GetDiag()
		}
	}

	if instance.TotalVolumeBandwidth != nil {
		if err = d.Set("total_volume_bandwidth", flex.IntValue(instance.TotalVolumeBandwidth)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting total_volume_bandwidth: %s", err), "(Data) ibm_is_instance", "read", "set-total_volume_bandwidth").GetDiag()
		}
	}

	if instance.Disks != nil {
		if err = d.Set("disks", dataSourceInstanceFlattenDisks(instance.Disks)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting disks: %s", err), "(Data) ibm_is_instance", "read", "set-disks").GetDiag()
		}
	}

	if instance.PrimaryNetworkInterface != nil {
		primaryNicList := make([]map[string]interface{}, 0)
		currentPrimNic := map[string]interface{}{}
		currentPrimNic["id"] = *instance.PrimaryNetworkInterface.ID
		currentPrimNic[isInstanceNicName] = *instance.PrimaryNetworkInterface.Name

		// reserved ip changes
		primaryIpList := make([]map[string]interface{}, 0)
		currentPrimIp := map[string]interface{}{}
		if instance.PrimaryNetworkInterface.PrimaryIP.Address != nil {
			currentPrimNic[isInstanceNicPrimaryIpv4Address] = *instance.PrimaryNetworkInterface.PrimaryIP.Address
			currentPrimIp[isInstanceNicReservedIpAddress] = *instance.PrimaryNetworkInterface.PrimaryIP.Address
		}
		if instance.PrimaryNetworkInterface.PrimaryIP.Href != nil {
			currentPrimIp[isInstanceNicReservedIpHref] = *instance.PrimaryNetworkInterface.PrimaryIP.Href
		}
		if instance.PrimaryNetworkInterface.PrimaryIP.Name != nil {
			currentPrimIp[isInstanceNicReservedIpName] = *instance.PrimaryNetworkInterface.PrimaryIP.Name
		}
		if instance.PrimaryNetworkInterface.PrimaryIP.ID != nil {
			currentPrimIp[isInstanceNicReservedIpId] = *instance.PrimaryNetworkInterface.PrimaryIP.ID
		}
		if instance.PrimaryNetworkInterface.PrimaryIP.ResourceType != nil {
			currentPrimIp[isInstanceNicReservedIpResourceType] = *instance.PrimaryNetworkInterface.PrimaryIP.ResourceType
		}
		primaryIpList = append(primaryIpList, currentPrimIp)
		currentPrimNic[isInstanceNicPrimaryIP] = primaryIpList

		getnicoptions := &vpcv1.GetInstanceNetworkInterfaceOptions{
			InstanceID: &id,
			ID:         instance.PrimaryNetworkInterface.ID,
		}
		insnic, _, err := sess.GetInstanceNetworkInterfaceWithContext(context, getnicoptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceNetworkInterfaceWithContext failed: %s", err.Error()), "(Data) ibm_is_instance", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if insnic.PortSpeed != nil {
			currentPrimNic[isInstanceNicPortSpeed] = *insnic.PortSpeed
		}
		currentPrimNic[isInstanceNicSubnet] = *insnic.Subnet.ID
		if len(insnic.SecurityGroups) != 0 {
			secgrpList := []string{}
			for i := 0; i < len(insnic.SecurityGroups); i++ {
				secgrpList = append(secgrpList, string(*(insnic.SecurityGroups[i].ID)))
			}
			currentPrimNic[isInstanceNicSecurityGroups] = flex.NewStringSet(schema.HashString, secgrpList)
		}

		primaryNicList = append(primaryNicList, currentPrimNic)
		if err = d.Set("primary_network_interface", primaryNicList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_network_interface: %s", err), "(Data) ibm_is_instance", "read", "set-primary_network_interface").GetDiag()
		}
	}
	if err = d.Set("confidential_compute_mode", instance.ConfidentialComputeMode); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting confidential_compute_mode: %s", err), "(Data) ibm_is_instance", "read", "set-confidential_compute_mode").GetDiag()
	}
	primaryNetworkAttachment := []map[string]interface{}{}
	if instance.PrimaryNetworkAttachment != nil {
		modelMap, err := dataSourceIBMIsInstanceInstanceNetworkAttachmentReferenceToMap(instance.PrimaryNetworkAttachment)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance", "read", "primary_network_attachment-to-map").GetDiag()
		}
		primaryNetworkAttachment = append(primaryNetworkAttachment, modelMap)
	}
	if err = d.Set("primary_network_attachment", primaryNetworkAttachment); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_network_attachment: %s", err), "(Data) ibm_is_instance", "read", "set-primary_network_attachment").GetDiag()
	}

	if err = d.Set("enable_secure_boot", instance.EnableSecureBoot); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting enable_secure_boot: %s", err), "(Data) ibm_is_instance", "read", "set-enable_secure_boot").GetDiag()
	}
	if instance.NetworkInterfaces != nil {
		interfacesList := make([]map[string]interface{}, 0)
		for _, intfc := range instance.NetworkInterfaces {
			if *intfc.ID != *instance.PrimaryNetworkInterface.ID {
				currentNic := map[string]interface{}{}
				currentNic["id"] = *intfc.ID
				currentNic[isInstanceNicName] = *intfc.Name

				// reserved ip changes
				primaryIpList := make([]map[string]interface{}, 0)
				currentPrimIp := map[string]interface{}{}
				if intfc.PrimaryIP.Address != nil {
					currentPrimIp[isInstanceNicReservedIpAddress] = *intfc.PrimaryIP.Address
					currentNic[isInstanceNicPrimaryIpv4Address] = *intfc.PrimaryIP.Address
				}
				if intfc.PrimaryIP.Href != nil {
					currentPrimIp[isInstanceNicReservedIpHref] = *intfc.PrimaryIP.Href
				}
				if intfc.PrimaryIP.Name != nil {
					currentPrimIp[isInstanceNicReservedIpName] = *intfc.PrimaryIP.Name
				}
				if intfc.PrimaryIP.ID != nil {
					currentPrimIp[isInstanceNicReservedIpId] = *intfc.PrimaryIP.ID
				}
				if intfc.PrimaryIP.ResourceType != nil {
					currentPrimIp[isInstanceNicReservedIpResourceType] = *intfc.PrimaryIP.ResourceType
				}
				primaryIpList = append(primaryIpList, currentPrimIp)
				currentNic[isInstanceNicPrimaryIP] = primaryIpList

				getnicoptions := &vpcv1.GetInstanceNetworkInterfaceOptions{
					InstanceID: &id,
					ID:         intfc.ID,
				}
				insnic, _, err := sess.GetInstanceNetworkInterfaceWithContext(context, getnicoptions)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceNetworkInterfaceWithContext failed: %s", err.Error()), "(Data) ibm_is_instance", "read")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				currentNic[isInstanceNicSubnet] = *insnic.Subnet.ID
				if len(insnic.SecurityGroups) != 0 {
					secgrpList := []string{}
					for i := 0; i < len(insnic.SecurityGroups); i++ {
						secgrpList = append(secgrpList, string(*(insnic.SecurityGroups[i].ID)))
					}
					currentNic[isInstanceNicSecurityGroups] = flex.NewStringSet(schema.HashString, secgrpList)
				}
				interfacesList = append(interfacesList, currentNic)

			}
		}
		if err = d.Set("network_interfaces", interfacesList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting network_interfaces: %s", err), "(Data) ibm_is_instance", "read", "set-network_interfaces").GetDiag()
		}
	}
	networkAttachments := []map[string]interface{}{}
	if instance.NetworkAttachments != nil {
		for _, modelItem := range instance.NetworkAttachments {
			modelMap, err := dataSourceIBMIsInstanceInstanceNetworkAttachmentReferenceToMap(&modelItem)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance", "read", "network_attachments-to-map").GetDiag()
			}
			networkAttachments = append(networkAttachments, modelMap)
		}
	}
	if err = d.Set("network_attachments", networkAttachments); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting network_attachments: %s", err), "(Data) ibm_is_instance", "read", "set-network_attachments").GetDiag()
	}

	var rsaKey *rsa.PrivateKey
	if instance.Image != nil {
		if err = d.Set(isInstanceImage, *instance.Image.ID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting image: %s", err), "(Data) ibm_is_instance", "read", "set-image").GetDiag()
		}
		image := *instance.Image.Name
		res := strings.Contains(image, "windows")
		if res {
			if privatekey, ok := d.GetOk(isInstancePEM); ok {
				keyFlag := privatekey.(string)
				keybytes := []byte(keyFlag)

				if keyFlag != "" {
					block, err := pem.Decode(keybytes)
					if block == nil {
						converterr := fmt.Errorf("Failed to load the private key from the given key contents. Instead of the key file path, please make sure the private key is pem format (%v)", err)
						tfErr := flex.TerraformErrorf(converterr, fmt.Sprintf("Decode failed: %s", converterr.Error()), "(Data) ibm_is_instance", "read")
						log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
						return tfErr.GetDiag()
					}
					isEncrypted := false
					if block.Type == "OPENSSH PRIVATE KEY" {
						var err error
						isEncrypted, err = isOpenSSHPrivKeyEncrypted(block.Bytes)
						if err != nil {
							err = fmt.Errorf("Failed to check if the provided open ssh key is encrypted or not %s", err)
							tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isOpenSSHPrivKeyEncrypted failed: %s", err.Error()), "(Data) ibm_is_instance", "read")
							log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
							return tfErr.GetDiag()
						}
					} else {
						isEncrypted = x509.IsEncryptedPEMBlock(block)
					}
					passphrase := ""
					var privateKey interface{}
					if isEncrypted {
						if pass, ok := d.GetOk(isInstancePassphrase); ok {
							passphrase = pass.(string)
						} else {
							converterr := fmt.Errorf("Mandatory field 'passphrase' not provided")
							tfErr := flex.TerraformErrorf(converterr, fmt.Sprintf("passphrase failed: %s", converterr.Error()), "(Data) ibm_is_instance", "read")
							log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
							return tfErr.GetDiag()
						}
						var err error
						privateKey, err = sshkeys.ParseEncryptedRawPrivateKey(keybytes, []byte(passphrase))
						if err != nil {
							err = fmt.Errorf("Fail to decrypting the private key: %s", err)
							tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ParseEncryptedRawPrivateKey failed: %s", err.Error()), "(Data) ibm_is_instance", "read")
							log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
							return tfErr.GetDiag()
						}
					} else {
						var err error
						privateKey, err = sshkeys.ParseEncryptedRawPrivateKey(keybytes, nil)
						if err != nil {
							err = fmt.Errorf("Fail to decrypting the private key: %s", err)
							tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ParseEncryptedRawPrivateKey failed: %s", err.Error()), "(Data) ibm_is_instance", "read")
							log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
							return tfErr.GetDiag()
						}
					}
					var ok bool
					rsaKey, ok = privateKey.(*rsa.PrivateKey)
					if !ok {
						converterr := fmt.Errorf("Failed to convert to RSA private key")
						tfErr := flex.TerraformErrorf(converterr, fmt.Sprintf("privateKey.(*rsa.PrivateKey) failed: %s", converterr.Error()), "(Data) ibm_is_instance", "read")
						log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
						return tfErr.GetDiag()
					}
				}
			}
		}
	}

	getInstanceInitializationOptions := &vpcv1.GetInstanceInitializationOptions{
		ID: &id,
	}
	initParms, _, err := sess.GetInstanceInitializationWithContext(context, getInstanceInitializationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceInitializationWithContext failed: %s", err.Error()), "(Data) ibm_is_instance", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if initParms.Keys != nil {
		initKeyList := make([]map[string]interface{}, 0)
		for _, key := range initParms.Keys {
			initKey := map[string]interface{}{}
			id := ""
			if key.ID != nil {
				id = *key.ID
			}
			initKey["id"] = id
			name := ""
			if key.Name != nil {
				name = *key.Name
			}
			initKey["name"] = name
			initKeyList = append(initKeyList, initKey)
			break

		}
		if err = d.Set(isInstanceInitKeys, initKeyList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting keys: %s", err), "(Data) ibm_is_instance", "read", "set-keys").GetDiag()
		}
	}
	//set the lifecycle status, reasons
	if instance.LifecycleState != nil {
		if err = d.Set(isInstanceLifecycleState, *instance.LifecycleState); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_state: %s", err), "(Data) ibm_is_instance", "read", "set-lifecycle_state").GetDiag()
		}
	}
	if instance.LifecycleReasons != nil {
		if err = d.Set(isInstanceLifecycleReasons, dataSourceInstanceFlattenLifecycleReasons(instance.LifecycleReasons)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_reasons: %s", err), "(Data) ibm_is_instance", "read", "set-lifecycle_reasons").GetDiag()
		}
	}

	if initParms.Password != nil && initParms.Password.EncryptedPassword != nil {
		ciphertext := *initParms.Password.EncryptedPassword
		password := base64.StdEncoding.EncodeToString(ciphertext)
		if rsaKey != nil {
			rng := rand.Reader
			clearPassword, err := rsa.DecryptPKCS1v15(rng, rsaKey, ciphertext)
			if err != nil {
				err = fmt.Errorf("Can not decrypt the password with the given key, %s", err)
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DecryptPKCS1v15 failed: %s", err.Error()), "(Data) ibm_is_instance", "read")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			password = string(clearPassword)
		}
		if err = d.Set(isInstanceInitPassword, password); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting password: %s", err), "(Data) ibm_is_instance", "read", "set-password").GetDiag()
		}
	}

	if err = d.Set(isInstanceStatus, *instance.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_instance", "read", "set-status").GetDiag()
	}
	//set the status reasons
	if instance.StatusReasons != nil {
		statusReasonsList := make([]map[string]interface{}, 0)
		for _, sr := range instance.StatusReasons {
			currentSR := map[string]interface{}{}
			if sr.Code != nil && sr.Message != nil {
				currentSR[isInstanceStatusReasonsCode] = *sr.Code
				currentSR[isInstanceStatusReasonsMessage] = *sr.Message
				if sr.MoreInfo != nil {
					currentSR[isInstanceStatusReasonsMoreInfo] = *sr.MoreInfo
				}
				statusReasonsList = append(statusReasonsList, currentSR)
			}
		}
		if err = d.Set(isInstanceStatusReasons, statusReasonsList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status_reasons: %s", err), "(Data) ibm_is_instance", "read", "set-status_reasons").GetDiag()
		}
	}
	if err = d.Set(isInstanceVPC, *instance.VPC.ID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vpc: %s", err), "(Data) ibm_is_instance", "read", "set-vpc").GetDiag()
	}
	if err = d.Set(isInstanceZone, *instance.Zone.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting zone: %s", err), "(Data) ibm_is_instance", "read", "set-zone").GetDiag()
	}
	var volumes []string
	volumes = make([]string, 0)
	if instance.VolumeAttachments != nil {
		for _, volume := range instance.VolumeAttachments {
			if volume.Volume != nil && *volume.Volume.ID != *instance.BootVolumeAttachment.Volume.ID {
				volumes = append(volumes, *volume.Volume.ID)
			}
		}
	}
	if err = d.Set(isInstanceVolumes, flex.NewStringSet(schema.HashString, volumes)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting volumes: %s", err), "(Data) ibm_is_instance", "read", "set-volumes").GetDiag()
	}
	if instance.VolumeAttachments != nil {
		volList := make([]map[string]interface{}, 0)
		for _, volume := range instance.VolumeAttachments {
			vol := map[string]interface{}{}
			if volume.Volume != nil {
				vol["id"] = *volume.ID
				vol["volume_id"] = *volume.Volume.ID
				vol["name"] = *volume.Name
				vol["volume_name"] = *volume.Volume.Name
				vol["volume_crn"] = *volume.Volume.CRN
				volList = append(volList, vol)
			}
		}
		if err = d.Set(isInstanceVolumeAttachments, volList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting volume_attachments: %s", err), "(Data) ibm_is_instance", "read", "set-volume_attachments").GetDiag()
		}
	}
	if instance.BootVolumeAttachment != nil {
		bootVolList := make([]map[string]interface{}, 0)
		bootVol := map[string]interface{}{}
		bootVol["id"] = *instance.BootVolumeAttachment.ID
		bootVol["name"] = *instance.BootVolumeAttachment.Name
		if instance.BootVolumeAttachment.Device != nil {
			bootVol["device"] = *instance.BootVolumeAttachment.Device.ID
		}
		if instance.BootVolumeAttachment.Volume != nil {
			bootVol["volume_name"] = *instance.BootVolumeAttachment.Volume.Name
			bootVol["volume_id"] = *instance.BootVolumeAttachment.Volume.ID
			bootVol["volume_crn"] = *instance.BootVolumeAttachment.Volume.CRN
		}
		bootVolList = append(bootVolList, bootVol)
		if err = d.Set(isInstanceBootVolume, bootVolList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting boot_volume: %s", err), "(Data) ibm_is_instance", "read", "set-boot_volume").GetDiag()
		}
	}
	tags, err := flex.GetGlobalTagsUsingCRN(meta, *instance.CRN, "", isInstanceUserTagType)
	if err != nil {
		log.Printf(
			"[ERROR] Error on get of resource vpc Instance (%s) tags: %s", d.Id(), err)
	}
	if err = d.Set(isInstanceTags, tags); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tags: %s", err), "(Data) ibm_is_instance", "read", "set-tags").GetDiag()
	}
	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *instance.CRN, "", isInstanceAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource vpc Instance (%s) access tags: %s", d.Id(), err)
	}
	if err = d.Set(isInstanceAccessTags, accesstags); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_tags: %s", err), "(Data) ibm_is_instance", "read", "set-access_tags").GetDiag()
	}
	controller, err := flex.GetBaseController(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBaseController failed: %s", err.Error()), "(Data) ibm_is_instance", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/compute/vs"); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_controller_url: %s", err), "(Data) ibm_is_instance", "read", "set-resource_controller_url").GetDiag()
	}
	if err = d.Set(flex.ResourceName, instance.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_name: %s", err), "(Data) ibm_is_instance", "read", "set-resource_name").GetDiag()
	}
	if err = d.Set(flex.ResourceCRN, instance.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_crn: %s", err), "(Data) ibm_is_instance", "read", "set-resource_crn").GetDiag()
	}
	if err = d.Set(IsInstanceCRN, instance.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_instance", "read", "set-crn").GetDiag()
	}
	if err = d.Set(flex.ResourceStatus, instance.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_status: %s", err), "(Data) ibm_is_instance", "read", "set-resource_status").GetDiag()
	}
	if instance.ResourceGroup != nil {
		if err = d.Set(isInstanceResourceGroup, instance.ResourceGroup.ID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_is_instance", "read", "set-resource_group").GetDiag()
		}
		if err = d.Set(flex.ResourceGroupName, instance.ResourceGroup.Name); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group_name: %s", err), "(Data) ibm_is_instance", "read", "set-resource_group_name").GetDiag()
		}
	}
	if instance.HealthReasons != nil {
		healthReasonsList := make([]map[string]interface{}, 0)
		for _, sr := range instance.HealthReasons {
			currentSR := map[string]interface{}{}
			if sr.Code != nil && sr.Message != nil {
				currentSR["code"] = *sr.Code
				currentSR["message"] = *sr.Message
				if sr.MoreInfo != nil {
					currentSR["more_info"] = *sr.Message
				}
				healthReasonsList = append(healthReasonsList, currentSR)
			}
		}
		if err = d.Set("health_reasons", healthReasonsList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting health_reasons: %s", err), "(Data) ibm_is_instance", "read", "set-health_reasons").GetDiag()
		}
	}
	if err = d.Set("health_state", instance.HealthState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting health_state: %s", err), "(Data) ibm_is_instance", "read", "set-health_state").GetDiag()
	}
	if instance.ReservationAffinity != nil {
		reservationAffinity := []map[string]interface{}{}
		reservationAffinityMap := map[string]interface{}{}

		reservationAffinityMap[isReservationAffinityPolicyResp] = instance.ReservationAffinity.Policy
		if instance.ReservationAffinity.Pool != nil {
			poolList := make([]map[string]interface{}, 0)
			for _, pool := range instance.ReservationAffinity.Pool {
				res := map[string]interface{}{}

				res[isReservationId] = *pool.ID
				res[isReservationHref] = *pool.Href
				res[isReservationName] = *pool.Name
				res[isReservationCrn] = *pool.CRN
				res[isReservationResourceType] = *pool.ResourceType
				if pool.Deleted != nil {
					deletedList := []map[string]interface{}{}
					deletedMap := dataSourceReservationDeletedToMap(*pool.Deleted)
					deletedList = append(deletedList, deletedMap)
					res[isReservationDeleted] = deletedList
				}
				poolList = append(poolList, res)
			}
			reservationAffinityMap[isReservationAffinityPool] = poolList
		}
		reservationAffinity = append(reservationAffinity, reservationAffinityMap)
		if err = d.Set("reservation_affinity", reservationAffinity); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting reservation_affinity: %s", err), "(Data) ibm_is_instance", "read", "set-reservation_affinity").GetDiag()
		}
	}
	if instance.Reservation != nil {
		resList := make([]map[string]interface{}, 0)
		res := map[string]interface{}{}

		res[isReservationId] = *instance.Reservation.ID
		res[isReservationHref] = *instance.Reservation.Href
		res[isReservationName] = *instance.Reservation.Name
		res[isReservationCrn] = *instance.Reservation.CRN
		res[isReservationResourceType] = *instance.Reservation.ResourceType
		if instance.Reservation.Deleted != nil {
			deletedList := []map[string]interface{}{}
			deletedMap := dataSourceReservationDeletedToMap(*instance.Reservation.Deleted)
			deletedList = append(deletedList, deletedMap)
			res[isReservationDeleted] = deletedList
		}
		resList = append(resList, res)
		if err = d.Set("reservation", resList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting reservation: %s", err), "(Data) ibm_is_instance", "read", "set-reservation").GetDiag()
		}
	}
	return nil

}

func dataSourceReservationDeletedToMap(deletedItem vpcv1.Deleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

const opensshv1Magic = "openssh-key-v1"

type opensshPrivateKey struct {
	CipherName   string
	KdfName      string
	KdfOpts      string
	NumKeys      uint32
	PubKey       string
	PrivKeyBlock string
}

func isOpenSSHPrivKeyEncrypted(data []byte) (bool, error) {
	magic := append([]byte(opensshv1Magic), 0)
	if !bytes.Equal(magic, data[0:len(magic)]) {
		return false, errors.New("[ERROR] Invalid openssh private key format")
	}
	content := data[len(magic):]

	privKey := opensshPrivateKey{}

	if err := ssh.Unmarshal(content, &privKey); err != nil {
		return false, err
	}

	if privKey.KdfName == "none" && privKey.CipherName == "none" {
		return false, nil
	}
	return true, nil
}

func dataSourceInstanceFlattenDisks(result []vpcv1.InstanceDisk) (disks []map[string]interface{}) {
	for _, disksItem := range result {
		disks = append(disks, dataSourceInstanceDisksToMap(disksItem))
	}

	return disks
}

func dataSourceInstanceDisksToMap(disksItem vpcv1.InstanceDisk) (disksMap map[string]interface{}) {
	disksMap = map[string]interface{}{}

	if disksItem.CreatedAt != nil {
		disksMap["created_at"] = disksItem.CreatedAt.String()
	}
	if disksItem.Href != nil {
		disksMap["href"] = disksItem.Href
	}
	if disksItem.ID != nil {
		disksMap["id"] = disksItem.ID
	}
	if disksItem.InterfaceType != nil {
		disksMap["interface_type"] = disksItem.InterfaceType
	}
	if disksItem.Name != nil {
		disksMap["name"] = disksItem.Name
	}
	if disksItem.ResourceType != nil {
		disksMap["resource_type"] = disksItem.ResourceType
	}
	if disksItem.Size != nil {
		disksMap["size"] = disksItem.Size
	}

	return disksMap
}
func dataSourceInstanceFlattenLifecycleReasons(lifecycleReasons []vpcv1.InstanceLifecycleReason) (lifecycleReasonsList []map[string]interface{}) {
	lifecycleReasonsList = make([]map[string]interface{}, 0)
	for _, lr := range lifecycleReasons {
		currentLR := map[string]interface{}{}
		if lr.Code != nil && lr.Message != nil {
			currentLR[isInstanceLifecycleReasonsCode] = *lr.Code
			currentLR[isInstanceLifecycleReasonsMessage] = *lr.Message
			if lr.MoreInfo != nil {
				currentLR[isInstanceLifecycleReasonsMoreInfo] = *lr.MoreInfo
			}
			lifecycleReasonsList = append(lifecycleReasonsList, currentLR)
		}
	}
	return lifecycleReasonsList
}

func dataSourceIBMIsInstanceInstanceNetworkAttachmentReferenceToMap(model *vpcv1.InstanceNetworkAttachmentReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := dataSourceIBMIsInstanceInstanceNetworkAttachmentReferenceDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = model.Href
	modelMap["id"] = model.ID
	modelMap["name"] = model.Name
	primaryIPMap, err := dataSourceIBMIsInstanceReservedIPReferenceToMap(model.PrimaryIP)
	if err != nil {
		return modelMap, err
	}
	modelMap["primary_ip"] = []map[string]interface{}{primaryIPMap}
	modelMap["resource_type"] = model.ResourceType
	subnetMap, err := dataSourceIBMIsInstanceSubnetReferenceToMap(model.Subnet)
	if err != nil {
		return modelMap, err
	}
	modelMap["subnet"] = []map[string]interface{}{subnetMap}
	virtualNetworkInterfaceMap, err := dataSourceIBMIsInstanceVirtualNetworkInterfaceReferenceAttachmentContextToMap(model.VirtualNetworkInterface)
	if err != nil {
		return modelMap, err
	}
	modelMap["virtual_network_interface"] = []map[string]interface{}{virtualNetworkInterfaceMap}
	return modelMap, nil
}

func dataSourceIBMIsInstanceVirtualNetworkInterfaceReferenceAttachmentContextToMap(model *vpcv1.VirtualNetworkInterfaceReferenceAttachmentContext) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = model.CRN
	modelMap["href"] = model.Href
	modelMap["id"] = model.ID
	modelMap["name"] = model.Name
	modelMap["resource_type"] = model.ResourceType
	return modelMap, nil
}
func dataSourceIBMIsInstanceInstanceNetworkAttachmentReferenceDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = model.MoreInfo
	return modelMap, nil
}
func dataSourceIBMIsInstanceReservedIPReferenceToMap(model *vpcv1.ReservedIPReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = model.Address
	if model.Deleted != nil {
		deletedMap, err := dataSourceIBMIsInstanceReservedIPReferenceDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = model.Href
	modelMap["id"] = model.ID
	modelMap["name"] = model.Name
	modelMap["resource_type"] = model.ResourceType
	return modelMap, nil
}
func dataSourceIBMIsInstanceSubnetReferenceToMap(model *vpcv1.SubnetReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = model.CRN
	if model.Deleted != nil {
		deletedMap, err := dataSourceIBMIsInstanceSubnetReferenceDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = model.Href
	modelMap["id"] = model.ID
	modelMap["name"] = model.Name
	modelMap["resource_type"] = model.ResourceType
	return modelMap, nil
}
func dataSourceIBMIsInstanceSubnetReferenceDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = model.MoreInfo
	return modelMap, nil
}

func dataSourceIBMIsInstanceReservedIPReferenceDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = model.MoreInfo
	return modelMap, nil
}
func DataSourceIBMIsInstanceClusterNetworkReferenceToMap(model *vpcv1.ClusterNetworkReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsInstanceDeletedToMap(model.Deleted)
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

func DataSourceIBMIsInstanceInstanceClusterNetworkAttachmentReferenceToMap(model *vpcv1.InstanceClusterNetworkAttachmentReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}
func DataSourceIBMIsInstanceDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

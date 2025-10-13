// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isInstanceTemplateHref                 = "href"
	isInstanceTemplateCrn                  = "crn"
	isInstanceTemplateLimit                = "limit"
	isInstanceTemplateNext                 = "next"
	isInstanceTemplateTotalCount           = "total_count"
	isInstanceTemplatePortSpeed            = "port_speed"
	isInstanceTemplatePortType             = "type"
	isInstanceTemplatePortValue            = "value"
	isInstanceTemplateDeleteVol            = "delete_volume_on_instance_delete"
	isInstanceTemplateVol                  = "volume"
	isInstanceTemplateMemory               = "memory"
	isInstanceTemplateMemoryValue          = "value"
	isInstanceTemplateMemoryType           = "type"
	isInstanceTemplateMemoryValues         = "values"
	isInstanceTemplateMemoryDefault        = "default"
	isInstanceTemplateMemoryMin            = "min"
	isInstanceTemplateMemoryMax            = "max"
	isInstanceTemplateMemoryStep           = "step"
	isInstanceTemplateSocketCount          = "socket_count"
	isInstanceTemplateSocketValue          = "value"
	isInstanceTemplateSocketType           = "type"
	isInstanceTemplateSocketValues         = "values"
	isInstanceTemplateSocketDefault        = "default"
	isInstanceTemplateSocketMin            = "min"
	isInstanceTemplateSocketMax            = "max"
	isInstanceTemplateSocketStep           = "step"
	isInstanceTemplateVcpuArch             = "vcpu_architecture"
	isInstanceTemplateVcpuArchType         = "type"
	isInstanceTemplateVcpuArchValue        = "value"
	isInstanceTemplateVcpuCount            = "vcpu_count"
	isInstanceTemplateVcpuCountValue       = "value"
	isInstanceTemplateVcpuCountType        = "type"
	isInstanceTemplateVcpuCountValues      = "values"
	isInstanceTemplateVcpuCountDefault     = "default"
	isInstanceTemplateVcpuCountMin         = "min"
	isInstanceTemplateVcpuCountMax         = "max"
	isInstanceTemplateVcpuCountStep        = "step"
	isInstanceTemplateStart                = "start"
	isInstanceTemplateVersion              = "version"
	isInstanceTemplateBootVolumeAttachment = "boot_volume_attachment"
)

func DataSourceIBMISInstanceTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISInstanceTemplateRead,
		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"identifier", isInstanceTemplateName},
			},
			isInstanceTemplateName: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"identifier", isInstanceTemplateName},
			},

			// cluster changes
			"cluster_network_attachments": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The cluster network attachments to create for this virtual server instance. A cluster network attachment represents a device that is connected to a cluster network. The number of network attachments must match one of the values from the instance profile's `cluster_network_attachment_count` before the instance can be started.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_network_interface": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A cluster network interface for the instance cluster network attachment. This can bespecified using an existing cluster network interface that does not already have a `target`,or a prototype object for a new cluster network interface.This instance must reside in the same VPC as the specified cluster network interface. Thecluster network interface must reside in the same cluster network as the`cluster_network_interface` of any other `cluster_network_attachments` for this instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auto_delete": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether this cluster network interface will be automatically deleted when `target` is deleted.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name for this cluster network interface. The name must not be used by another interface in the cluster network. Names beginning with `ibm-` are reserved for provider-owned resources, and are not allowed. If unspecified, the name will be a hyphenated list of randomly-selected words.",
									},
									"primary_ip": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The primary IP address to bind to the cluster network interface. May be eithera cluster network subnet reserved IP identity, or a cluster network subnet reserved IPprototype object which will be used to create a new cluster network subnet reserved IP.If a cluster network subnet reserved IP identity is provided, the specified clusternetwork subnet reserved IP must be unbound.If a cluster network subnet reserved IP prototype object with an address is provided,the address must be available on the cluster network interface's cluster networksubnet. If no address is specified, an available address on the cluster network subnetwill be automatically selected and reserved.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique identifier for this cluster network subnet reserved IP.",
												},
												"href": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The URL for this cluster network subnet reserved IP.",
												},
												"address": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The IP address to reserve, which must not already be reserved on the subnet.If unspecified, an available address on the subnet will automatically be selected.",
												},
												"auto_delete": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Indicates whether this cluster network subnet reserved IP member will be automatically deleted when either `target` is deleted, or the cluster network subnet reserved IP is unbound.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name for this cluster network subnet reserved IP. The name must not be used by another reserved IP in the cluster network subnet. Names starting with `ibm-` are reserved for provider-owned resources, and are not allowed. If unspecified, the name will be a hyphenated list of randomly-selected words.",
												},
											},
										},
									},
									"subnet": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The associated cluster network subnet. Required if `primary_ip` does not specify acluster network subnet reserved IP identity.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique identifier for this cluster network subnet.",
												},
												"href": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The URL for this cluster network subnet.",
												},
											},
										},
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this cluster network interface.",
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this cluster network interface.",
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this cluster network attachment. Names must be unique within the instance the cluster network attachment resides in. If unspecified, the name will be a hyphenated list of randomly-selected words. Names starting with `ibm-` are reserved for provider-owned resources, and are not allowed.",
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
			isInstanceTemplateHref: {
				Type:     schema.TypeString,
				Computed: true,
			},
			isInstanceTemplateCrn: {
				Type:     schema.TypeString,
				Computed: true,
			},
			isInstanceTemplateVPC: {
				Type:     schema.TypeString,
				Computed: true,
			},
			isInstanceTemplateZone: {
				Type:     schema.TypeString,
				Computed: true,
			},
			isInstanceTemplateProfile: {
				Type:     schema.TypeString,
				Computed: true,
			},
			isInstanceTemplateKeys: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			isInstanceTotalVolumeBandwidth: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The amount of bandwidth (in megabits per second) allocated exclusively to instance storage volumes",
			},
			isInstanceDefaultTrustedProfileAutoLink: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If set to `true`, the system will create a link to the specified `target` trusted profile during instance creation. Regardless of whether a link is created by the system or manually using the IAM Identity service, it will be automatically deleted when the instance is deleted.",
			},
			isInstanceDefaultTrustedProfileTarget: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier or CRN of the default IAM trusted profile to use for this virtual server instance.",
			},
			isInstanceTemplateMetadataServiceEnabled: {
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
			isInstanceAvailablePolicyHostFailure: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The availability policy to use for this virtual server instance. The action to perform if the compute host experiences a failure.",
			},
			isInstanceTemplateVolumeAttachments: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceTemplateDeleteVol: {
							Type:     schema.TypeBool,
							Computed: true,
						},
						isInstanceTemplateName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplateVol: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplateVolAttVolPrototype: {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isInstanceTemplateVolAttVolIops: {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum I/O operations per second (IOPS) for the volume.",
									},
									isInstanceTemplateVolAttVolProfile: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The  globally unique name for the volume profile to use for this volume.",
									},
									"bandwidth": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									isInstanceTemplateVolAttVolCapacity: {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The capacity of the volume in gigabytes. The specified minimum and maximum capacity values for creating or updating volumes may expand in the future.",
									},
									isInstanceTemplateVolAttVolEncryptionKey: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.",
									},
									"source_snapshot": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The snapshot to use as a source for the volume's data.",
									},
									"allowed_use": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The usage constraints to be matched against the requested instance or bare metal server properties to determine compatibility.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"bare_metal_server": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The expression that must be satisfied by the properties of a bare metal server provisioned using the image data in this volume.",
												},
												"instance": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The expression that must be satisfied by the properties of a virtual server instance provisioned using this volume.",
												},
												"api_version": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The API version with which to evaluate the expressions.",
												},
											},
										},
									},
									isInstanceTemplateVolAttTags: {
										Type:        schema.TypeSet,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Set:         flex.ResourceIBMVPCHash,
										Description: "The user tags associated with this volume.",
									},
								},
							},
						},
					},
				},
			},
			isInstanceTemplateCatalogOffering: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The catalog offering or offering version to use when provisioning this virtual server instance template. If an offering is specified, the latest version of that offering will be used. The specified offering or offering version may be in a different account in the same enterprise, subject to IAM policies.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceTemplateCatalogOfferingOfferingCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Identifies a catalog offering by a unique CRN property",
						},
						isInstanceTemplateCatalogOfferingVersionCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Identifies a version of a catalog offering by a unique CRN property",
						},
						isInstanceTemplateCatalogOfferingPlanCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this catalog offering version's billing plan",
						},
					},
				},
			},
			isInstanceTemplatePrimaryNetworkInterface: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceTemplateNicName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplateNicPrimaryIpv4Address: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplateNicPrimaryIP: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isInstanceTemplateNicReservedIpAddress: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP address to reserve, which must not already be reserved on the subnet.",
									},
									isInstanceTemplateNicReservedIpName: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The user-defined name for this reserved IP. If unspecified, the name will be a hyphenated list of randomly-selected words. Names must be unique within the subnet the reserved IP resides in. ",
									},
									isInstanceTemplateNicReservedIpId: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Identifies a reserved IP by a unique property.",
									},
								},
							},
						},
						isInstanceTemplateNicSecurityGroups: {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						isInstanceTemplateNicSubnet: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			isInstanceTemplateNetworkInterfaces: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceTemplateNicName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplateNicPrimaryIpv4Address: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplateNicPrimaryIP: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isInstanceTemplateNicReservedIpAddress: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP address to reserve, which must not already be reserved on the subnet.",
									},
									isInstanceTemplateNicReservedIpName: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The user-defined name for this reserved IP. If unspecified, the name will be a hyphenated list of randomly-selected words. Names must be unique within the subnet the reserved IP resides in. ",
									},
									isInstanceTemplateNicReservedIpId: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Identifies a reserved IP by a unique property.",
									},
								},
							},
						},
						isInstanceTemplateNicSecurityGroups: {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						isInstanceTemplateNicSubnet: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"network_attachments": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The additional network attachments to create for the virtual server instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this network attachment. Names must be unique within the instance the network attachment resides in. If unspecified, the name will be a hyphenated list of randomly-selected words.",
						},
						"virtual_network_interface": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A virtual network interface for the instance network attachment. This can be specifiedusing an existing virtual network interface, or a prototype object for a new virtualnetwork interface.If an existing virtual network interface is specified, `enable_infrastructure_nat` must be`false`.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allow_ip_spoofing": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether source IP spoofing is allowed on this interface. If `false`, source IP spoofing is prevented on this interface. If `true`, source IP spoofing is allowed on this interface.",
									},
									"auto_delete": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether this virtual network interface will be automatically deleted when`target` is deleted.",
									},
									"enable_infrastructure_nat": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If `true`:- The VPC infrastructure performs any needed NAT operations.- `floating_ips` must not have more than one floating IP.If `false`:- Packets are passed unchanged to/from the virtual network interface,  allowing the workload to perform any needed NAT operations.- `allow_ip_spoofing` must be `false`.- If the virtual network interface is attached:  - The target `resource_type` must be `bare_metal_server_network_attachment`.  - The target `interface_type` must not be `hipersocket`.",
									},
									"ips": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Additional IP addresses to bind to the virtual network interface. Each item may be either a reserved IP identity, or as a reserved IP prototype object which will be used to create a new reserved IP. All IP addresses must be in the same subnet as the primary IP.If reserved IP identities are provided, the specified reserved IPs must be unbound.If reserved IP prototype objects with addresses are provided, the addresses must be available on the virtual network interface's subnet. For any prototype objects that do not specify an address, an available address on the subnet will be automatically selected and reserved.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique identifier for this reserved IP.",
												},
												"href": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The URL for this reserved IP.",
												},
												"address": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The IP address to reserve, which must not already be reserved on the subnet.If unspecified, an available address on the subnet will automatically be selected.",
												},
												"auto_delete": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Indicates whether this reserved IP member will be automatically deleted when either`target` is deleted, or the reserved IP is unbound.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name for this reserved IP. The name must not be used by another reserved IP in the subnet. Names starting with `ibm-` are reserved for provider-owned resources, and are not allowed. If unspecified, the name will be a hyphenated list of randomly-selected words.",
												},
											},
										},
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name for this virtual network interface. The name must not be used by another virtual network interface in the VPC. If unspecified, the name will be a hyphenated list of randomly-selected words. Names beginning with `ibm-` are reserved for provider-owned resources, and are not allowed.",
									},
									"primary_ip": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The primary IP address to bind to the virtual network interface. May be either areserved IP identity, or a reserved IP prototype object which will be used to create anew reserved IP.If a reserved IP identity is provided, the specified reserved IP must be unbound.If a reserved IP prototype object with an address is provided, the address must beavailable on the virtual network interface's subnet. If no address is specified,an available address on the subnet will be automatically selected and reserved.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique identifier for this reserved IP.",
												},
												"href": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The URL for this reserved IP.",
												},
												"address": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The IP address to reserve, which must not already be reserved on the subnet.If unspecified, an available address on the subnet will automatically be selected.",
												},
												"auto_delete": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Indicates whether this reserved IP member will be automatically deleted when either`target` is deleted, or the reserved IP is unbound.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name for this reserved IP. The name must not be used by another reserved IP in the subnet. Names starting with `ibm-` are reserved for provider-owned resources, and are not allowed. If unspecified, the name will be a hyphenated list of randomly-selected words.",
												},
											},
										},
									},
									"protocol_state_filtering_mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The protocol state filtering mode used for this virtual network interface.",
									},
									"resource_group": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The resource group to use for this virtual network interface. If unspecified, thevirtual server instance's resource group will be used.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique identifier for this resource group.",
												},
											},
										},
									},
									"security_groups": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The security groups to use for this virtual network interface. If unspecified, the default security group of the VPC for the subnet is used.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique identifier for this security group.",
												},
												"crn": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The security group's CRN.",
												},
												"href": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The security group's canonical URL.",
												},
											},
										},
									},
									"subnet": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The associated subnet. Required if `primary_ip` does not specify a reserved IP.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique identifier for this subnet.",
												},
												"crn": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The CRN for this subnet.",
												},
												"href": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The URL for this subnet.",
												},
											},
										},
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this virtual network interface.",
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this virtual network interface.",
									},
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this virtual network interface.",
									},
								},
							},
						},
					},
				},
			},
			"primary_network_attachment": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The primary network attachment to create for the virtual server instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this network attachment. Names must be unique within the instance the network attachment resides in. If unspecified, the name will be a hyphenated list of randomly-selected words.",
						},
						"virtual_network_interface": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A virtual network interface for the instance network attachment. This can be specifiedusing an existing virtual network interface, or a prototype object for a new virtualnetwork interface.If an existing virtual network interface is specified, `enable_infrastructure_nat` must be`false`.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allow_ip_spoofing": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether source IP spoofing is allowed on this interface. If `false`, source IP spoofing is prevented on this interface. If `true`, source IP spoofing is allowed on this interface.",
									},
									"auto_delete": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether this virtual network interface will be automatically deleted when`target` is deleted.",
									},
									"enable_infrastructure_nat": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If `true`:- The VPC infrastructure performs any needed NAT operations.- `floating_ips` must not have more than one floating IP.If `false`:- Packets are passed unchanged to/from the virtual network interface,  allowing the workload to perform any needed NAT operations.- `allow_ip_spoofing` must be `false`.- If the virtual network interface is attached:  - The target `resource_type` must be `bare_metal_server_network_attachment`.  - The target `interface_type` must not be `hipersocket`.",
									},
									"ips": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Additional IP addresses to bind to the virtual network interface. Each item may be either a reserved IP identity, or as a reserved IP prototype object which will be used to create a new reserved IP. All IP addresses must be in the same subnet as the primary IP.If reserved IP identities are provided, the specified reserved IPs must be unbound.If reserved IP prototype objects with addresses are provided, the addresses must be available on the virtual network interface's subnet. For any prototype objects that do not specify an address, an available address on the subnet will be automatically selected and reserved.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique identifier for this reserved IP.",
												},
												"href": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The URL for this reserved IP.",
												},
												"address": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The IP address to reserve, which must not already be reserved on the subnet.If unspecified, an available address on the subnet will automatically be selected.",
												},
												"auto_delete": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Indicates whether this reserved IP member will be automatically deleted when either`target` is deleted, or the reserved IP is unbound.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name for this reserved IP. The name must not be used by another reserved IP in the subnet. Names starting with `ibm-` are reserved for provider-owned resources, and are not allowed. If unspecified, the name will be a hyphenated list of randomly-selected words.",
												},
											},
										},
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name for this virtual network interface. The name must not be used by another virtual network interface in the VPC. If unspecified, the name will be a hyphenated list of randomly-selected words. Names beginning with `ibm-` are reserved for provider-owned resources, and are not allowed.",
									},
									"primary_ip": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The primary IP address to bind to the virtual network interface. May be either areserved IP identity, or a reserved IP prototype object which will be used to create anew reserved IP.If a reserved IP identity is provided, the specified reserved IP must be unbound.If a reserved IP prototype object with an address is provided, the address must beavailable on the virtual network interface's subnet. If no address is specified,an available address on the subnet will be automatically selected and reserved.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique identifier for this reserved IP.",
												},
												"href": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The URL for this reserved IP.",
												},
												"address": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The IP address to reserve, which must not already be reserved on the subnet.If unspecified, an available address on the subnet will automatically be selected.",
												},
												"auto_delete": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Indicates whether this reserved IP member will be automatically deleted when either`target` is deleted, or the reserved IP is unbound.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name for this reserved IP. The name must not be used by another reserved IP in the subnet. Names starting with `ibm-` are reserved for provider-owned resources, and are not allowed. If unspecified, the name will be a hyphenated list of randomly-selected words.",
												},
											},
										},
									},
									"protocol_state_filtering_mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The protocol state filtering mode used for this virtual network interface.",
									},
									"resource_group": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The resource group to use for this virtual network interface. If unspecified, thevirtual server instance's resource group will be used.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique identifier for this resource group.",
												},
											},
										},
									},
									"security_groups": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The security groups to use for this virtual network interface. If unspecified, the default security group of the VPC for the subnet is used.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique identifier for this security group.",
												},
												"crn": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The security group's CRN.",
												},
												"href": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The security group's canonical URL.",
												},
											},
										},
									},
									"subnet": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The associated subnet. Required if `primary_ip` does not specify a reserved IP.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique identifier for this subnet.",
												},
												"crn": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The CRN for this subnet.",
												},
												"href": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The URL for this subnet.",
												},
											},
										},
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this virtual network interface.",
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this virtual network interface.",
									},
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this virtual network interface.",
									},
								},
							},
						},
					},
				},
			},
			isInstanceTemplateUserData: {
				Type:     schema.TypeString,
				Computed: true,
			},
			isInstanceTemplateImage: {
				Type:     schema.TypeString,
				Computed: true,
			},
			isInstanceTemplateBootVolumeAttachment: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceTemplateDeleteVol: {
							Type:     schema.TypeBool,
							Computed: true,
						},
						isInstanceTemplateName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplateVol: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isInstanceTemplateBootSize: {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"bandwidth": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						isInstanceTemplateBootProfile: {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_snapshot": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The snapshot from which to clone the volume.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this snapshot.",
									},
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN of this snapshot.",
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this snapshot.",
									},
								},
							},
						},
						"allowed_use": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The usage constraints to be matched against the requested instance or bare metal server properties to determine compatibility.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bare_metal_server": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The expression that must be satisfied by the properties of a bare metal server provisioned using the image data in this volume.",
									},
									"instance": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The expression that must be satisfied by the properties of a virtual server instance provisioned using this volume.",
									},
									"api_version": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The API version with which to evaluate the expressions.",
									},
								},
							},
						},
						isInstanceTemplateBootVolumeTags: {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Set:         flex.ResourceIBMVPCHash,
							Description: "The user tags associated with this volume.",
						},
					},
				},
			},
			isInstanceTemplateResourceGroup: {
				Type:     schema.TypeString,
				Computed: true,
			},
			"placement_target": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The placement restrictions to use for the virtual server instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this dedicated host.",
						},
						"crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this dedicated host.",
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this dedicated host.",
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
						isReservationAffinityPool: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reservation associated with this template.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISInstanceTemplateRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	instanceC, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_template", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if idOk, ok := d.GetOk("identifier"); ok {
		id := idOk.(string)
		getInstanceTemplatesOptions := &vpcv1.GetInstanceTemplateOptions{
			ID: &id,
		}
		instTempl, _, err := instanceC.GetInstanceTemplateWithContext(context, getInstanceTemplatesOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceTemplateWithContext failed: %s", err.Error()), "(Data) ibm_is_instance_template", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		switch v := instTempl.(type) {
		case *vpcv1.InstanceTemplate:
			instanceTemplate := v
			d.SetId(*instanceTemplate.ID)
			if err = d.Set(isInstanceTemplateHref, instanceTemplate.Href); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_instance_template", "read", "set-href").GetDiag()
			}
			if err = d.Set(isInstanceTemplateCrn, instanceTemplate.CRN); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_instance_template", "read", "set-crn").GetDiag()
			}

			if err = d.Set(isInstanceTemplateName, instanceTemplate.Name); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_instance_template", "read", "set-name").GetDiag()
			}

			if err = d.Set(isInstanceTemplateUserData, instanceTemplate.UserData); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting user_data: %s", err), "(Data) ibm_is_instance_template", "read", "set-user_data").GetDiag()
			}

			if err = d.Set("confidential_compute_mode", instanceTemplate.ConfidentialComputeMode); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting confidential_compute_mode: %s", err), "(Data) ibm_is_instance_template", "read", "set-confidential_compute_mode").GetDiag()
			}
			// vni

			networkAttachments := []map[string]interface{}{}
			if instanceTemplate.NetworkAttachments != nil {
				for _, modelItem := range instanceTemplate.NetworkAttachments {
					modelMap, err := dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeToMap(&modelItem)
					if err != nil {
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_template", "read", "network_attachments-to-map").GetDiag()
					}
					networkAttachments = append(networkAttachments, modelMap)
				}
			}
			if err = d.Set("network_attachments", networkAttachments); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting network_attachments: %s", err), "(Data) ibm_is_instance_template", "read", "set-network_attachments").GetDiag()
			}

			primaryNetworkAttachment := []map[string]interface{}{}
			if instanceTemplate.PrimaryNetworkAttachment != nil {
				modelMap, err := dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeToMap(instanceTemplate.PrimaryNetworkAttachment)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_template", "read", "primary_network_attachment-to-map").GetDiag()
				}
				primaryNetworkAttachment = append(primaryNetworkAttachment, modelMap)
			}
			if err = d.Set("primary_network_attachment", primaryNetworkAttachment); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_network_attachment: %s", err), "(Data) ibm_is_instance_template", "read", "set-primary_network_attachment").GetDiag()
			}

			if err = d.Set("enable_secure_boot", instanceTemplate.EnableSecureBoot); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting enable_secure_boot: %s", err), "(Data) ibm_is_instance_template", "read", "set-enable_secure_boot").GetDiag()
			}
			if instanceTemplate.DefaultTrustedProfile != nil {
				if instanceTemplate.DefaultTrustedProfile.AutoLink != nil {
					if err = d.Set(isInstanceDefaultTrustedProfileAutoLink, instanceTemplate.DefaultTrustedProfile.AutoLink); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting default_trusted_profile_auto_link: %s", err), "(Data) ibm_is_instance_template", "read", "set-default_trusted_profile_auto_link").GetDiag()
					}
				}
				if instanceTemplate.DefaultTrustedProfile.Target != nil {
					switch reflect.TypeOf(instanceTemplate.DefaultTrustedProfile.Target).String() {
					case "*vpcv1.TrustedProfileIdentityTrustedProfileByID":
						{
							target := instanceTemplate.DefaultTrustedProfile.Target.(*vpcv1.TrustedProfileIdentityByID)
							if err = d.Set(isInstanceDefaultTrustedProfileTarget, target.ID); err != nil {
								return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting default_trusted_profile_target: %s", err), "(Data) ibm_is_instance_template", "read", "set-default_trusted_profile_target").GetDiag()
							}
						}
					case "*vpcv1.TrustedProfileIdentityTrustedProfileByCRN":
						{
							target := instanceTemplate.DefaultTrustedProfile.Target.(*vpcv1.TrustedProfileIdentityByCRN)
							if err = d.Set(isInstanceDefaultTrustedProfileTarget, target.CRN); err != nil {
								return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting default_trusted_profile_target: %s", err), "(Data) ibm_is_instance_template", "read", "set-default_trusted_profile_target").GetDiag()
							}
						}
					}
				}
			}

			// cluster changes
			if !core.IsNil(instanceTemplate.ClusterNetworkAttachments) {
				clusterNetworkAttachments := []map[string]interface{}{}
				for _, clusterNetworkAttachmentsItem := range instanceTemplate.ClusterNetworkAttachments {
					clusterNetworkAttachmentsItemMap, err := DataSourceIBMIsInstanceTemplateInstanceClusterNetworkAttachmentPrototypeInstanceContextToMap(&clusterNetworkAttachmentsItem) // #nosec G601
					if err != nil {
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_template", "read", "cluster_network_attachments-to-map").GetDiag()
					}
					clusterNetworkAttachments = append(clusterNetworkAttachments, clusterNetworkAttachmentsItemMap)
				}
				if err = d.Set("cluster_network_attachments", clusterNetworkAttachments); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cluster_network_attachments: %s", err), "(Data) ibm_is_instance_template", "read", "set-cluster_network_attachments").GetDiag()
				}
			}

			// catalog offering if any
			if instanceTemplate.CatalogOffering != nil {
				catOfferingList := make([]map[string]interface{}, 0)
				insTempCatalogOffering := instanceTemplate.CatalogOffering.(*vpcv1.InstanceCatalogOfferingPrototype)

				currentOffering := map[string]interface{}{}
				if insTempCatalogOffering.Offering != nil {
					offering := insTempCatalogOffering.Offering.(*vpcv1.CatalogOfferingIdentity)
					currentOffering[isInstanceTemplateCatalogOfferingOfferingCrn] = *offering.CRN
				}
				if insTempCatalogOffering.Version != nil {
					version := insTempCatalogOffering.Version.(*vpcv1.CatalogOfferingVersionIdentity)
					currentOffering[isInstanceTemplateCatalogOfferingVersionCrn] = *version.CRN
				}
				if insTempCatalogOffering.Plan != nil {
					plan := insTempCatalogOffering.Plan.(*vpcv1.CatalogOfferingVersionPlanIdentity)
					if plan.CRN != nil && *plan.CRN != "" {
						currentOffering[isInstanceTemplateCatalogOfferingPlanCrn] = *plan.CRN
					}
				}
				catOfferingList = append(catOfferingList, currentOffering)
				if err = d.Set(isInstanceTemplateCatalogOffering, catOfferingList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting catalog_offering: %s", err), "(Data) ibm_is_instance_template", "read", "set-catalog_offering").GetDiag()
				}
			}

			if instanceTemplate.AvailabilityPolicy != nil && instanceTemplate.AvailabilityPolicy.HostFailure != nil {
				if err = d.Set(isInstanceTemplateAvailablePolicyHostFailure, *instanceTemplate.AvailabilityPolicy.HostFailure); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting availability_policy_host_failure: %s", err), "(Data) ibm_is_instance_template", "read", "set-availability_policy_host_failure").GetDiag()
				}
			}
			if instanceTemplate.Keys != nil {
				keys := []string{}
				for _, intfc := range instanceTemplate.Keys {
					instanceKeyIntf := intfc.(*vpcv1.KeyIdentity)
					keys = append(keys, *instanceKeyIntf.ID)
				}
				if err = d.Set(isInstanceTemplateKeys, keys); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting keys: %s", err), "(Data) ibm_is_instance_template", "read", "set-keys").GetDiag()
				}
			}

			if instanceTemplate.MetadataService != nil {
				if err = d.Set(isInstanceTemplateMetadataServiceEnabled, instanceTemplate.MetadataService.Enabled); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting metadata_service_enabled: %s", err), "(Data) ibm_is_instance_template", "read", "set-metadata_service_enabled").GetDiag()
				}

				metadataService := []map[string]interface{}{}
				metadataServiceMap := map[string]interface{}{}

				metadataServiceMap[isInstanceMetadataServiceEnabled1] = instanceTemplate.MetadataService.Enabled
				if instanceTemplate.MetadataService.Protocol != nil {
					metadataServiceMap[isInstanceMetadataServiceProtocol] = instanceTemplate.MetadataService.Protocol
				}
				if instanceTemplate.MetadataService.ResponseHopLimit != nil {
					metadataServiceMap[isInstanceMetadataServiceRespHopLimit] = instanceTemplate.MetadataService.ResponseHopLimit
				}

				metadataService = append(metadataService, metadataServiceMap)
				if err = d.Set("metadata_service", metadataService); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting metadata_service: %s", err), "(Data) ibm_is_instance_template", "read", "set-metadata_service").GetDiag()
				}
			}

			if instanceTemplate.Profile != nil {
				instanceProfileIntf := instanceTemplate.Profile
				identity := instanceProfileIntf.(*vpcv1.InstanceProfileIdentity)
				if err = d.Set("profile", *identity.Name); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting profile: %s", err), "(Data) ibm_is_instance_template", "read", "set-profile").GetDiag()
				}
			}

			if instanceTemplate.PlacementTarget != nil {
				placementTargetList := []map[string]interface{}{}
				placementTargetMap := dataSourceInstanceTemplateCollectionTemplatesPlacementTargetToMap(*instanceTemplate.PlacementTarget.(*vpcv1.InstancePlacementTargetPrototype))
				placementTargetList = append(placementTargetList, placementTargetMap)
				if err = d.Set("placement_target", placementTargetList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting placement_target: %s", err), "(Data) ibm_is_instance_template", "read", "set-placement_target").GetDiag()
				}
			}

			if instanceTemplate.ReservationAffinity != nil {
				reservationAffinity := []map[string]interface{}{}
				reservationAffinityMap := map[string]interface{}{}

				reservationAffinityMap[isReservationAffinityPolicyResp] = instanceTemplate.ReservationAffinity.Policy
				if instanceTemplate.ReservationAffinity.Pool != nil && len(instanceTemplate.ReservationAffinity.Pool) > 0 {
					pool := instanceTemplate.ReservationAffinity.Pool[0]
					res := ""
					if idPool, ok := pool.(*vpcv1.ReservationIdentityByID); ok {
						res = *idPool.ID
					} else if crnPool, ok := pool.(*vpcv1.ReservationIdentityByCRN); ok {
						res = *crnPool.CRN
					} else if hrefPool, ok := pool.(*vpcv1.ReservationIdentityByHref); ok {
						res = *hrefPool.Href
					}
					reservationAffinityMap[isReservationAffinityPool] = res
				}
				reservationAffinity = append(reservationAffinity, reservationAffinityMap)
				if err = d.Set("reservation_affinity", reservationAffinity); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting reservation_affinity: %s", err), "(Data) ibm_is_instance_template", "read", "set-reservation_affinity").GetDiag()
				}
			}

			if instanceTemplate.TotalVolumeBandwidth != nil {
				if err = d.Set("total_volume_bandwidth", int(*instanceTemplate.TotalVolumeBandwidth)); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting total_volume_bandwidth: %s", err), "(Data) ibm_is_instance_template", "read", "set-total_volume_bandwidth").GetDiag()
				}
			}

			if instanceTemplate.PrimaryNetworkInterface != nil {
				interfaceList := make([]map[string]interface{}, 0)
				currentPrimNic := map[string]interface{}{}
				currentPrimNic[isInstanceTemplateNicName] = *instanceTemplate.PrimaryNetworkInterface.Name
				if instanceTemplate.PrimaryNetworkInterface.PrimaryIP != nil {
					primaryipIntf := instanceTemplate.PrimaryNetworkInterface.PrimaryIP
					primaryIpList := make([]map[string]interface{}, 0)
					currentPrimIp := map[string]interface{}{}
					switch reflect.TypeOf(primaryipIntf).String() {
					case "*vpcv1.NetworkInterfaceIPPrototype":
						{
							log.Printf("[INFO] NetworkInterfaceIPPrototype")
							primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototype)
							if primaryip.Address != nil {
								currentPrimNic[isInstanceTemplateNicPrimaryIpv4Address] = *primaryip.Address
								currentPrimIp[isInstanceTemplateNicReservedIpAddress] = *primaryip.Address
							}
							if primaryip.ID != nil {
								currentPrimIp[isInstanceTemplateNicReservedIpId] = *primaryip.ID
							}
						}
					case "*vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext":
						{
							log.Printf("[INFO] NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext")
							primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext)
							if primaryip.Address != nil {
								currentPrimNic[isInstanceTemplateNicPrimaryIpv4Address] = *primaryip.Address
								currentPrimIp[isInstanceTemplateNicReservedIpAddress] = *primaryip.Address
							}
						}
					case "*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity":
						{
							log.Printf("[INFO] NetworkInterfaceIPPrototypeReservedIPIdentity")
							primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity)
							if primaryip.ID != nil {
								currentPrimIp[isInstanceTemplateNicReservedIpId] = *primaryip.ID
							}
						}
					}
					primaryIpList = append(primaryIpList, currentPrimIp)
					currentPrimNic[isInstanceTemplateNicPrimaryIP] = primaryIpList
				}
				subInf := instanceTemplate.PrimaryNetworkInterface.Subnet
				subnetIdentity := subInf.(*vpcv1.SubnetIdentity)
				currentPrimNic[isInstanceTemplateNicSubnet] = *subnetIdentity.ID

				if len(instanceTemplate.PrimaryNetworkInterface.SecurityGroups) != 0 {
					secgrpList := []string{}
					for i := 0; i < len(instanceTemplate.PrimaryNetworkInterface.SecurityGroups); i++ {
						secGrpInf := instanceTemplate.PrimaryNetworkInterface.SecurityGroups[i]
						secGrpIdentity := secGrpInf.(*vpcv1.SecurityGroupIdentity)
						secgrpList = append(secgrpList, string(*secGrpIdentity.ID))
					}
					currentPrimNic[isInstanceTemplateNicSecurityGroups] = flex.NewStringSet(schema.HashString, secgrpList)
				}
				interfaceList = append(interfaceList, currentPrimNic)
				if err = d.Set("primary_network_interface", interfaceList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_network_interface: %s", err), "(Data) ibm_is_instance_template", "read", "set-primary_network_interface").GetDiag()
				}
			}

			if instanceTemplate.NetworkInterfaces != nil {
				interfacesList := make([]map[string]interface{}, 0)
				for _, intfc := range instanceTemplate.NetworkInterfaces {
					currentNic := map[string]interface{}{}
					currentNic[isInstanceTemplateNicName] = *intfc.Name
					if intfc.PrimaryIP != nil {
						primaryipIntf := intfc.PrimaryIP
						primaryIpList := make([]map[string]interface{}, 0)
						currentPrimIp := map[string]interface{}{}
						switch reflect.TypeOf(primaryipIntf).String() {
						case "*vpcv1.NetworkInterfaceIPPrototype":
							{
								primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototype)
								currentNic[isInstanceTemplateNicPrimaryIpv4Address] = primaryip.Address
								currentPrimIp[isInstanceTemplateNicReservedIpAddress] = primaryip.Address
							}
						case "*vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext":
							{
								primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext)
								currentNic[isInstanceTemplateNicPrimaryIpv4Address] = primaryip.Address
								currentPrimIp[isInstanceTemplateNicReservedIpAddress] = primaryip.Address
							}
						case "*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity":
							{
								primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity)
								currentPrimIp[isInstanceTemplateNicReservedIpId] = primaryip.ID
							}
						}
						primaryIpList = append(primaryIpList, currentPrimIp)
						currentNic[isInstanceTemplateNicPrimaryIP] = primaryIpList
					}
					//currentNic[isInstanceTemplateNicAllowIpSpoofing] = intfc.AllowIpSpoofing
					subInf := intfc.Subnet
					subnetIdentity := subInf.(*vpcv1.SubnetIdentity)
					currentNic[isInstanceTemplateNicSubnet] = *subnetIdentity.ID
					if len(intfc.SecurityGroups) != 0 {
						secgrpList := []string{}
						for i := 0; i < len(intfc.SecurityGroups); i++ {
							secGrpInf := intfc.SecurityGroups[i]
							secGrpIdentity := secGrpInf.(*vpcv1.SecurityGroupIdentity)
							secgrpList = append(secgrpList, string(*secGrpIdentity.ID))
						}
						currentNic[isInstanceTemplateNicSecurityGroups] = flex.NewStringSet(schema.HashString, secgrpList)
					}

					interfacesList = append(interfacesList, currentNic)
				}
				if err = d.Set("network_interfaces", interfacesList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting network_interfaces: %s", err), "(Data) ibm_is_instance_template", "read", "set-network_interfaces").GetDiag()
				}
			}

			if instanceTemplate.Image != nil {
				imageInf := instanceTemplate.Image
				imageIdentity := imageInf.(*vpcv1.ImageIdentity)
				if err = d.Set("image", imageIdentity.ID); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting image: %s", err), "(Data) ibm_is_instance_template", "read", "set-image").GetDiag()
				}
			}

			if instanceTemplate.VPC != nil {
				vpcInf := instanceTemplate.VPC
				vpcRef := vpcInf.(*vpcv1.VPCIdentity)
				if err = d.Set("vpc", vpcRef.ID); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vpc: %s", err), "(Data) ibm_is_instance_template", "read", "set-vpc").GetDiag()
				}
			}

			if instanceTemplate.Zone != nil {
				zoneInf := instanceTemplate.Zone
				zone := zoneInf.(*vpcv1.ZoneIdentity)
				if err = d.Set("zone", zone.Name); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting zone: %s", err), "(Data) ibm_is_instance_template", "read", "set-zone").GetDiag()
				}
			}

			interfacesList := make([]map[string]interface{}, 0)
			if instanceTemplate.VolumeAttachments != nil {
				for _, volume := range instanceTemplate.VolumeAttachments {
					volumeAttach := map[string]interface{}{}
					volumeAttach[isInstanceTemplateVolAttName] = *volume.Name
					volumeAttach[isInstanceTemplateDeleteVolume] = *volume.DeleteVolumeOnInstanceDelete
					volumeIntf := volume.Volume
					volumeInst := volumeIntf.(*vpcv1.VolumeAttachmentPrototypeVolume)
					newVolumeArr := []map[string]interface{}{}
					newVolume := map[string]interface{}{}

					if volumeInst.ID != nil {
						volumeAttach[isInstanceTemplateVolAttVolume] = *volumeInst.ID
					}
					// bandwidth changes
					if volumeInst.Bandwidth != nil {
						newVolume["bandwidth"] = volumeInst.Bandwidth
					}
					if volumeInst.Capacity != nil {
						newVolume[isInstanceTemplateVolAttVolCapacity] = *volumeInst.Capacity
					}
					if volumeInst.Profile != nil {
						profile := volumeInst.Profile.(*vpcv1.VolumeProfileIdentity)
						newVolume[isInstanceTemplateVolAttVolProfile] = profile.Name
					}

					if volumeInst.Iops != nil {
						newVolume[isInstanceTemplateVolAttVolIops] = *volumeInst.Iops
					}
					if volumeInst.EncryptionKey != nil {
						encryptionKey := volumeInst.EncryptionKey.(*vpcv1.EncryptionKeyIdentity)
						newVolume[isInstanceTemplateVolAttVolEncryptionKey] = *encryptionKey.CRN
					}
					if volumeInst.AllowedUse != nil {
						modelMap, err := DataSourceIBMIsVolumeAllowedUseToMap(volumeInst.AllowedUse)
						if err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allowed_use: %s", err), "(Data) ibm_is_instance_template", "read", "set-allowed_use").GetDiag()
						}
						newVolume["allowed_use"] = []map[string]interface{}{modelMap}
					}

					if volumeInst.SourceSnapshot != nil {
						sourceSnapshot := volumeInst.SourceSnapshot.(*vpcv1.SnapshotIdentity)
						newVolume["source_snapshot"] = *sourceSnapshot.ID
					}
					if volumeInst.UserTags != nil {
						newVolume[isInstanceTemplateVolAttTags] = instanceTemplate.BootVolumeAttachment.Volume.UserTags
					}
					newVolumeArr = append(newVolumeArr, newVolume)
					volumeAttach[isInstanceTemplateVolAttVolPrototype] = newVolumeArr

					interfacesList = append(interfacesList, volumeAttach)
				}
				if err = d.Set("volume_attachments", interfacesList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting volume_attachments: %s", err), "(Data) ibm_is_instance_template", "read", "set-volume_attachments").GetDiag()
				}
			}

			if instanceTemplate.BootVolumeAttachment != nil {
				bootVolList := make([]map[string]interface{}, 0)
				bootVol := map[string]interface{}{}

				bootVol[isInstanceTemplateDeleteVol] = *instanceTemplate.BootVolumeAttachment.DeleteVolumeOnInstanceDelete
				if instanceTemplate.BootVolumeAttachment.Volume != nil {
					volumeIntf := instanceTemplate.BootVolumeAttachment.Volume
					bootVol[isInstanceTemplateName] = volumeIntf.Name
					bootVol[isInstanceTemplateVol] = volumeIntf.Name
					// bandwidth changes
					if volumeIntf.Bandwidth != nil {
						bootVol["bandwidth"] = volumeIntf.Bandwidth
					}
					bootVol[isInstanceTemplateBootSize] = volumeIntf.Capacity
					if instanceTemplate.BootVolumeAttachment.Volume.Profile != nil {
						volProfIntf := instanceTemplate.BootVolumeAttachment.Volume.Profile
						volProfInst := volProfIntf.(*vpcv1.VolumeProfileIdentity)
						bootVol[isInstanceTemplateBootProfile] = volProfInst.Name
					}
					if instanceTemplate.BootVolumeAttachment.Volume.AllowedUse != nil {
						modelMap, err := DataSourceIBMIsVolumeAllowedUseToMap(instanceTemplate.BootVolumeAttachment.Volume.AllowedUse)
						if err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allowed_use: %s", err), "(Data) ibm_is_instance_template", "read", "set-allowed_use").GetDiag()
						}
						bootVol["allowed_use"] = []map[string]interface{}{modelMap}
					}
					if instanceTemplate.BootVolumeAttachment.Volume.UserTags != nil {
						bootVol[isInstanceTemplateBootVolumeTags] = instanceTemplate.BootVolumeAttachment.Volume.UserTags
					}
				}
				bootVolList = append(bootVolList, bootVol)
				if err = d.Set("boot_volume_attachment", bootVolList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting boot_volume_attachment: %s", err), "(Data) ibm_is_instance_template", "read", "set-boot_volume_attachment").GetDiag()
				}
			}

			if instanceTemplate.ResourceGroup != nil {
				rg := instanceTemplate.ResourceGroup
				if err = d.Set("resource_group", rg.ID); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_is_instance_template", "read", "set-resource_group").GetDiag()
				}
			}
		case *vpcv1.InstanceTemplateInstanceBySourceSnapshotInstanceTemplateContext:
			instanceTemplate := v
			d.SetId(*instanceTemplate.ID)
			if err = d.Set(isInstanceTemplateHref, instanceTemplate.Href); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_instance_template", "read", "set-href").GetDiag()
			}
			if err = d.Set(isInstanceTemplateCrn, instanceTemplate.CRN); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_instance_template", "read", "set-crn").GetDiag()
			}

			if err = d.Set(isInstanceTemplateName, instanceTemplate.Name); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_instance_template", "read", "set-name").GetDiag()
			}

			if err = d.Set(isInstanceTemplateUserData, instanceTemplate.UserData); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting user_data: %s", err), "(Data) ibm_is_instance_template", "read", "set-user_data").GetDiag()
			}

			if err = d.Set("confidential_compute_mode", instanceTemplate.ConfidentialComputeMode); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting confidential_compute_mode: %s", err), "(Data) ibm_is_instance_template", "read", "set-confidential_compute_mode").GetDiag()
			}
			// vni

			networkAttachments := []map[string]interface{}{}
			if instanceTemplate.NetworkAttachments != nil {
				for _, modelItem := range instanceTemplate.NetworkAttachments {
					modelMap, err := dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeToMap(&modelItem)
					if err != nil {
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_template", "read", "network_attachments-to-map").GetDiag()
					}
					networkAttachments = append(networkAttachments, modelMap)
				}
			}
			if err = d.Set("network_attachments", networkAttachments); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting network_attachments: %s", err), "(Data) ibm_is_instance_template", "read", "set-network_attachments").GetDiag()
			}

			primaryNetworkAttachment := []map[string]interface{}{}
			if instanceTemplate.PrimaryNetworkAttachment != nil {
				modelMap, err := dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeToMap(instanceTemplate.PrimaryNetworkAttachment)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_template", "read", "primary_network_attachment-to-map").GetDiag()
				}
				primaryNetworkAttachment = append(primaryNetworkAttachment, modelMap)
			}
			if err = d.Set("primary_network_attachment", primaryNetworkAttachment); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_network_attachment: %s", err), "(Data) ibm_is_instance_template", "read", "set-primary_network_attachment").GetDiag()
			}

			if err = d.Set("enable_secure_boot", instanceTemplate.EnableSecureBoot); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting enable_secure_boot: %s", err), "(Data) ibm_is_instance_template", "read", "set-enable_secure_boot").GetDiag()
			}
			if instanceTemplate.DefaultTrustedProfile != nil {
				if instanceTemplate.DefaultTrustedProfile.AutoLink != nil {
					if err = d.Set(isInstanceDefaultTrustedProfileAutoLink, instanceTemplate.DefaultTrustedProfile.AutoLink); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting default_trusted_profile_auto_link: %s", err), "(Data) ibm_is_instance_template", "read", "set-default_trusted_profile_auto_link").GetDiag()
					}
				}
				if instanceTemplate.DefaultTrustedProfile.Target != nil {
					switch reflect.TypeOf(instanceTemplate.DefaultTrustedProfile.Target).String() {
					case "*vpcv1.TrustedProfileIdentityTrustedProfileByID":
						{
							target := instanceTemplate.DefaultTrustedProfile.Target.(*vpcv1.TrustedProfileIdentityByID)
							if err = d.Set(isInstanceDefaultTrustedProfileTarget, target.ID); err != nil {
								return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting default_trusted_profile_target: %s", err), "(Data) ibm_is_instance_template", "read", "set-default_trusted_profile_target").GetDiag()
							}
						}
					case "*vpcv1.TrustedProfileIdentityTrustedProfileByCRN":
						{
							target := instanceTemplate.DefaultTrustedProfile.Target.(*vpcv1.TrustedProfileIdentityByCRN)
							if err = d.Set(isInstanceDefaultTrustedProfileTarget, target.CRN); err != nil {
								return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting default_trusted_profile_target: %s", err), "(Data) ibm_is_instance_template", "read", "set-default_trusted_profile_target").GetDiag()
							}
						}
					}
				}
			}

			// cluster changes
			if !core.IsNil(instanceTemplate.ClusterNetworkAttachments) {
				clusterNetworkAttachments := []map[string]interface{}{}
				for _, clusterNetworkAttachmentsItem := range instanceTemplate.ClusterNetworkAttachments {
					clusterNetworkAttachmentsItemMap, err := DataSourceIBMIsInstanceTemplateInstanceClusterNetworkAttachmentPrototypeInstanceContextToMap(&clusterNetworkAttachmentsItem) // #nosec G601
					if err != nil {
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_template", "read", "cluster_network_attachments-to-map").GetDiag()
					}
					clusterNetworkAttachments = append(clusterNetworkAttachments, clusterNetworkAttachmentsItemMap)
				}
				if err = d.Set("cluster_network_attachments", clusterNetworkAttachments); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cluster_network_attachments: %s", err), "(Data) ibm_is_instance_template", "read", "set-cluster_network_attachments").GetDiag()
				}
			}

			if instanceTemplate.AvailabilityPolicy != nil && instanceTemplate.AvailabilityPolicy.HostFailure != nil {
				if err = d.Set(isInstanceTemplateAvailablePolicyHostFailure, *instanceTemplate.AvailabilityPolicy.HostFailure); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting availability_policy_host_failure: %s", err), "(Data) ibm_is_instance_template", "read", "set-availability_policy_host_failure").GetDiag()
				}
			}
			if instanceTemplate.Keys != nil {
				keys := []string{}
				for _, intfc := range instanceTemplate.Keys {
					instanceKeyIntf := intfc.(*vpcv1.KeyIdentity)
					keys = append(keys, *instanceKeyIntf.ID)
				}
				if err = d.Set(isInstanceTemplateKeys, keys); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting keys: %s", err), "(Data) ibm_is_instance_template", "read", "set-keys").GetDiag()
				}
			}

			if instanceTemplate.MetadataService != nil {
				if err = d.Set(isInstanceTemplateMetadataServiceEnabled, instanceTemplate.MetadataService.Enabled); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting metadata_service_enabled: %s", err), "(Data) ibm_is_instance_template", "read", "set-metadata_service_enabled").GetDiag()
				}

				metadataService := []map[string]interface{}{}
				metadataServiceMap := map[string]interface{}{}

				metadataServiceMap[isInstanceMetadataServiceEnabled1] = instanceTemplate.MetadataService.Enabled
				if instanceTemplate.MetadataService.Protocol != nil {
					metadataServiceMap[isInstanceMetadataServiceProtocol] = instanceTemplate.MetadataService.Protocol
				}
				if instanceTemplate.MetadataService.ResponseHopLimit != nil {
					metadataServiceMap[isInstanceMetadataServiceRespHopLimit] = instanceTemplate.MetadataService.ResponseHopLimit
				}

				metadataService = append(metadataService, metadataServiceMap)
				if err = d.Set("metadata_service", metadataService); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting metadata_service: %s", err), "(Data) ibm_is_instance_template", "read", "set-metadata_service").GetDiag()
				}
			}

			if instanceTemplate.Profile != nil {
				instanceProfileIntf := instanceTemplate.Profile
				identity := instanceProfileIntf.(*vpcv1.InstanceProfileIdentity)
				if err = d.Set("profile", *identity.Name); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting profile: %s", err), "(Data) ibm_is_instance_template", "read", "set-profile").GetDiag()
				}
			}

			if instanceTemplate.PlacementTarget != nil {
				placementTargetList := []map[string]interface{}{}
				placementTargetMap := dataSourceInstanceTemplateCollectionTemplatesPlacementTargetToMap(*instanceTemplate.PlacementTarget.(*vpcv1.InstancePlacementTargetPrototype))
				placementTargetList = append(placementTargetList, placementTargetMap)
				if err = d.Set("placement_target", placementTargetList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting placement_target: %s", err), "(Data) ibm_is_instance_template", "read", "set-placement_target").GetDiag()
				}
			}

			if instanceTemplate.ReservationAffinity != nil {
				reservationAffinity := []map[string]interface{}{}
				reservationAffinityMap := map[string]interface{}{}

				reservationAffinityMap[isReservationAffinityPolicyResp] = instanceTemplate.ReservationAffinity.Policy
				if instanceTemplate.ReservationAffinity.Pool != nil && len(instanceTemplate.ReservationAffinity.Pool) > 0 {
					pool := instanceTemplate.ReservationAffinity.Pool[0]
					res := ""
					if idPool, ok := pool.(*vpcv1.ReservationIdentityByID); ok {
						res = *idPool.ID
					} else if crnPool, ok := pool.(*vpcv1.ReservationIdentityByCRN); ok {
						res = *crnPool.CRN
					} else if hrefPool, ok := pool.(*vpcv1.ReservationIdentityByHref); ok {
						res = *hrefPool.Href
					}
					reservationAffinityMap[isReservationAffinityPool] = res
				}
				reservationAffinity = append(reservationAffinity, reservationAffinityMap)
				if err = d.Set("reservation_affinity", reservationAffinity); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting reservation_affinity: %s", err), "(Data) ibm_is_instance_template", "read", "set-reservation_affinity").GetDiag()
				}
			}

			if instanceTemplate.TotalVolumeBandwidth != nil {
				if err = d.Set("total_volume_bandwidth", int(*instanceTemplate.TotalVolumeBandwidth)); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting total_volume_bandwidth: %s", err), "(Data) ibm_is_instance_template", "read", "set-total_volume_bandwidth").GetDiag()
				}
			}

			if instanceTemplate.PrimaryNetworkInterface != nil {
				interfaceList := make([]map[string]interface{}, 0)
				currentPrimNic := map[string]interface{}{}
				currentPrimNic[isInstanceTemplateNicName] = *instanceTemplate.PrimaryNetworkInterface.Name
				if instanceTemplate.PrimaryNetworkInterface.PrimaryIP != nil {
					primaryipIntf := instanceTemplate.PrimaryNetworkInterface.PrimaryIP
					primaryIpList := make([]map[string]interface{}, 0)
					currentPrimIp := map[string]interface{}{}
					switch reflect.TypeOf(primaryipIntf).String() {
					case "*vpcv1.NetworkInterfaceIPPrototype":
						{
							log.Printf("[INFO] NetworkInterfaceIPPrototype")
							primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototype)
							if primaryip.Address != nil {
								currentPrimNic[isInstanceTemplateNicPrimaryIpv4Address] = *primaryip.Address
								currentPrimIp[isInstanceTemplateNicReservedIpAddress] = *primaryip.Address
							}
							if primaryip.ID != nil {
								currentPrimIp[isInstanceTemplateNicReservedIpId] = *primaryip.ID
							}
						}
					case "*vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext":
						{
							log.Printf("[INFO] NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext")
							primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext)
							if primaryip.Address != nil {
								currentPrimNic[isInstanceTemplateNicPrimaryIpv4Address] = *primaryip.Address
								currentPrimIp[isInstanceTemplateNicReservedIpAddress] = *primaryip.Address
							}
						}
					case "*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity":
						{
							log.Printf("[INFO] NetworkInterfaceIPPrototypeReservedIPIdentity")
							primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity)
							if primaryip.ID != nil {
								currentPrimIp[isInstanceTemplateNicReservedIpId] = *primaryip.ID
							}
						}
					}
					primaryIpList = append(primaryIpList, currentPrimIp)
					currentPrimNic[isInstanceTemplateNicPrimaryIP] = primaryIpList
				}
				subInf := instanceTemplate.PrimaryNetworkInterface.Subnet
				subnetIdentity := subInf.(*vpcv1.SubnetIdentity)
				currentPrimNic[isInstanceTemplateNicSubnet] = *subnetIdentity.ID

				if len(instanceTemplate.PrimaryNetworkInterface.SecurityGroups) != 0 {
					secgrpList := []string{}
					for i := 0; i < len(instanceTemplate.PrimaryNetworkInterface.SecurityGroups); i++ {
						secGrpInf := instanceTemplate.PrimaryNetworkInterface.SecurityGroups[i]
						secGrpIdentity := secGrpInf.(*vpcv1.SecurityGroupIdentity)
						secgrpList = append(secgrpList, string(*secGrpIdentity.ID))
					}
					currentPrimNic[isInstanceTemplateNicSecurityGroups] = flex.NewStringSet(schema.HashString, secgrpList)
				}
				interfaceList = append(interfaceList, currentPrimNic)
				if err = d.Set("primary_network_interface", interfaceList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_network_interface: %s", err), "(Data) ibm_is_instance_template", "read", "set-primary_network_interface").GetDiag()
				}
			}

			if instanceTemplate.NetworkInterfaces != nil {
				interfacesList := make([]map[string]interface{}, 0)
				for _, intfc := range instanceTemplate.NetworkInterfaces {
					currentNic := map[string]interface{}{}
					currentNic[isInstanceTemplateNicName] = *intfc.Name
					if intfc.PrimaryIP != nil {
						primaryipIntf := intfc.PrimaryIP
						primaryIpList := make([]map[string]interface{}, 0)
						currentPrimIp := map[string]interface{}{}
						switch reflect.TypeOf(primaryipIntf).String() {
						case "*vpcv1.NetworkInterfaceIPPrototype":
							{
								primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototype)
								currentNic[isInstanceTemplateNicPrimaryIpv4Address] = primaryip.Address
								currentPrimIp[isInstanceTemplateNicReservedIpAddress] = primaryip.Address
							}
						case "*vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext":
							{
								primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext)
								currentNic[isInstanceTemplateNicPrimaryIpv4Address] = primaryip.Address
								currentPrimIp[isInstanceTemplateNicReservedIpAddress] = primaryip.Address
							}
						case "*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity":
							{
								primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity)
								currentPrimIp[isInstanceTemplateNicReservedIpId] = primaryip.ID
							}
						}
						primaryIpList = append(primaryIpList, currentPrimIp)
						currentNic[isInstanceTemplateNicPrimaryIP] = primaryIpList
					}
					//currentNic[isInstanceTemplateNicAllowIpSpoofing] = intfc.AllowIpSpoofing
					subInf := intfc.Subnet
					subnetIdentity := subInf.(*vpcv1.SubnetIdentity)
					currentNic[isInstanceTemplateNicSubnet] = *subnetIdentity.ID
					if len(intfc.SecurityGroups) != 0 {
						secgrpList := []string{}
						for i := 0; i < len(intfc.SecurityGroups); i++ {
							secGrpInf := intfc.SecurityGroups[i]
							secGrpIdentity := secGrpInf.(*vpcv1.SecurityGroupIdentity)
							secgrpList = append(secgrpList, string(*secGrpIdentity.ID))
						}
						currentNic[isInstanceTemplateNicSecurityGroups] = flex.NewStringSet(schema.HashString, secgrpList)
					}

					interfacesList = append(interfacesList, currentNic)
				}
				if err = d.Set("network_interfaces", interfacesList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting network_interfaces: %s", err), "(Data) ibm_is_instance_template", "read", "set-network_interfaces").GetDiag()
				}
			}

			if instanceTemplate.VPC != nil {
				vpcInf := instanceTemplate.VPC
				vpcRef := vpcInf.(*vpcv1.VPCIdentity)
				if err = d.Set("vpc", vpcRef.ID); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vpc: %s", err), "(Data) ibm_is_instance_template", "read", "set-vpc").GetDiag()
				}
			}

			if instanceTemplate.Zone != nil {
				zoneInf := instanceTemplate.Zone
				zone := zoneInf.(*vpcv1.ZoneIdentity)
				if err = d.Set("zone", zone.Name); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting zone: %s", err), "(Data) ibm_is_instance_template", "read", "set-zone").GetDiag()
				}
			}

			interfacesList := make([]map[string]interface{}, 0)
			if instanceTemplate.VolumeAttachments != nil {
				for _, volume := range instanceTemplate.VolumeAttachments {
					volumeAttach := map[string]interface{}{}
					volumeAttach[isInstanceTemplateVolAttName] = *volume.Name
					volumeAttach[isInstanceTemplateDeleteVolume] = *volume.DeleteVolumeOnInstanceDelete
					volumeIntf := volume.Volume
					volumeInst := volumeIntf.(*vpcv1.VolumeAttachmentPrototypeVolume)
					newVolumeArr := []map[string]interface{}{}
					newVolume := map[string]interface{}{}

					if volumeInst.ID != nil {
						volumeAttach[isInstanceTemplateVolAttVolume] = *volumeInst.ID
					}
					// bandwidth changes
					if volumeInst.Bandwidth != nil {
						newVolume["bandwidth"] = volumeInst.Bandwidth
					}
					if volumeInst.Capacity != nil {
						newVolume[isInstanceTemplateVolAttVolCapacity] = *volumeInst.Capacity
					}
					if volumeInst.Profile != nil {
						profile := volumeInst.Profile.(*vpcv1.VolumeProfileIdentity)
						newVolume[isInstanceTemplateVolAttVolProfile] = profile.Name
					}

					if volumeInst.Iops != nil {
						newVolume[isInstanceTemplateVolAttVolIops] = *volumeInst.Iops
					}
					if volumeInst.EncryptionKey != nil {
						encryptionKey := volumeInst.EncryptionKey.(*vpcv1.EncryptionKeyIdentity)
						newVolume[isInstanceTemplateVolAttVolEncryptionKey] = *encryptionKey.CRN
					}
					if volumeInst.AllowedUse != nil {
						modelMap, err := DataSourceIBMIsVolumeAllowedUseToMap(volumeInst.AllowedUse)
						if err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allowed_use: %s", err), "(Data) ibm_is_instance_template", "read", "set-allowed_use").GetDiag()
						}
						newVolume["allowed_use"] = []map[string]interface{}{modelMap}
					}

					if volumeInst.SourceSnapshot != nil {
						sourceSnapshot := volumeInst.SourceSnapshot.(*vpcv1.SnapshotIdentity)
						newVolume["source_snapshot"] = *sourceSnapshot.ID
					}
					if volumeInst.UserTags != nil {
						newVolume[isInstanceTemplateVolAttTags] = instanceTemplate.BootVolumeAttachment.Volume.UserTags
					}
					newVolumeArr = append(newVolumeArr, newVolume)
					volumeAttach[isInstanceTemplateVolAttVolPrototype] = newVolumeArr

					interfacesList = append(interfacesList, volumeAttach)
				}
				if err = d.Set("volume_attachments", interfacesList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting volume_attachments: %s", err), "(Data) ibm_is_instance_template", "read", "set-volume_attachments").GetDiag()
				}
			}

			if instanceTemplate.BootVolumeAttachment != nil {
				bootVolList := make([]map[string]interface{}, 0)
				bootVol := map[string]interface{}{}

				bootVol[isInstanceTemplateDeleteVol] = *instanceTemplate.BootVolumeAttachment.DeleteVolumeOnInstanceDelete
				if instanceTemplate.BootVolumeAttachment.Volume != nil {
					volumeIntf := instanceTemplate.BootVolumeAttachment.Volume
					bootVol[isInstanceTemplateName] = volumeIntf.Name
					bootVol[isInstanceTemplateVol] = volumeIntf.Name
					// bandwidth changes
					if volumeIntf.Bandwidth != nil {
						bootVol["bandwidth"] = volumeIntf.Bandwidth
					}
					bootVol[isInstanceTemplateBootSize] = volumeIntf.Capacity
					if instanceTemplate.BootVolumeAttachment.Volume.Profile != nil {
						volProfIntf := instanceTemplate.BootVolumeAttachment.Volume.Profile
						volProfInst := volProfIntf.(*vpcv1.VolumeProfileIdentity)
						bootVol[isInstanceTemplateBootProfile] = volProfInst.Name
					}
					if instanceTemplate.BootVolumeAttachment.Volume.UserTags != nil {
						bootVol[isInstanceTemplateBootVolumeTags] = instanceTemplate.BootVolumeAttachment.Volume.UserTags
					}
					if instanceTemplate.BootVolumeAttachment.Volume.SourceSnapshot != nil {
						sourceSnapshotMap, _ := DataSourceIBMIsInstanceTemplateSnapshotIdentityToMap(instanceTemplate.BootVolumeAttachment.Volume.SourceSnapshot)
						bootVol["source_snapshot"] = []map[string]interface{}{sourceSnapshotMap}
					}
					if instanceTemplate.BootVolumeAttachment.Volume.AllowedUse != nil {
						modelMap, err := DataSourceIBMIsVolumeAllowedUseToMap(instanceTemplate.BootVolumeAttachment.Volume.AllowedUse)
						if err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allowed_use: %s", err), "(Data) ibm_is_instance_template", "read", "set-allowed_use").GetDiag()
						}
						bootVol["allowed_use"] = []map[string]interface{}{modelMap}

					}
				}
				bootVolList = append(bootVolList, bootVol)
				if err = d.Set("boot_volume_attachment", bootVolList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting boot_volume_attachment: %s", err), "(Data) ibm_is_instance_template", "read", "set-boot_volume_attachment").GetDiag()
				}
			}

			if instanceTemplate.ResourceGroup != nil {
				rg := instanceTemplate.ResourceGroup
				if err = d.Set("resource_group", rg.ID); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_is_instance_template", "read", "set-resource_group").GetDiag()
				}
			}
		default:
			log.Println("Unknown or unsupported instance template context type")
		}

	} else if nameOk, ok := d.GetOk(isInstanceTemplateName); ok {
		name := nameOk.(string)
		listInstanceTemplatesOptions := &vpcv1.ListInstanceTemplatesOptions{}
		availableTemplates, _, err := instanceC.ListInstanceTemplatesWithContext(context, listInstanceTemplatesOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceTemplatesWithContext failed: %s", err.Error()), "(Data) ibm_is_instance_template", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		flag := false
		for _, instTempl := range availableTemplates.Templates {
			switch v := instTempl.(type) {
			case *vpcv1.InstanceTemplate:
				instanceTemplate := v
				if name == *instanceTemplate.Name {
					flag = true
					d.SetId(*instanceTemplate.ID)
					if err = d.Set("href", instanceTemplate.Href); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_instance_template", "read", "set-href").GetDiag()
					}
					if err = d.Set("crn", instanceTemplate.CRN); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_instance_template", "read", "set-crn").GetDiag()
					}
					if err = d.Set("name", instanceTemplate.Name); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_instance_template", "read", "set-name").GetDiag()
					}
					if err = d.Set("user_data", instanceTemplate.UserData); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting user_data: %s", err), "(Data) ibm_is_instance_template", "read", "set-user_data").GetDiag()
					}
					if err = d.Set("confidential_compute_mode", instanceTemplate.ConfidentialComputeMode); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting confidential_compute_mode: %s", err), "(Data) ibm_is_instance_template", "read", "set-confidential_compute_mode").GetDiag()
					}
					if err = d.Set("enable_secure_boot", instanceTemplate.EnableSecureBoot); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting enable_secure_boot: %s", err), "(Data) ibm_is_instance_template", "read", "set-enable_secure_boot").GetDiag()
					}
					// cluster changes
					if !core.IsNil(instanceTemplate.ClusterNetworkAttachments) {
						clusterNetworkAttachments := []map[string]interface{}{}
						for _, clusterNetworkAttachmentsItem := range instanceTemplate.ClusterNetworkAttachments {
							clusterNetworkAttachmentsItemMap, err := DataSourceIBMIsInstanceTemplateInstanceClusterNetworkAttachmentPrototypeInstanceContextToMap(&clusterNetworkAttachmentsItem) // #nosec G601
							if err != nil {
								return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_template", "read", "cluster_network_attachments-to-map").GetDiag()
							}
							clusterNetworkAttachments = append(clusterNetworkAttachments, clusterNetworkAttachmentsItemMap)
						}
						if err = d.Set("cluster_network_attachments", clusterNetworkAttachments); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cluster_network_attachments: %s", err), "(Data) ibm_is_instance_template", "read", "set-cluster_network_attachments").GetDiag()
						}
					}

					// catalog offering if any
					if instanceTemplate.CatalogOffering != nil {
						catOfferingList := make([]map[string]interface{}, 0)
						insTempCatalogOffering := instanceTemplate.CatalogOffering.(*vpcv1.InstanceCatalogOfferingPrototype)

						currentOffering := map[string]interface{}{}
						if insTempCatalogOffering.Offering != nil {
							offering := insTempCatalogOffering.Offering.(*vpcv1.CatalogOfferingIdentity)
							currentOffering[isInstanceTemplateCatalogOfferingOfferingCrn] = *offering.CRN
						}
						if insTempCatalogOffering.Version != nil {
							version := insTempCatalogOffering.Version.(*vpcv1.CatalogOfferingVersionIdentity)
							currentOffering[isInstanceTemplateCatalogOfferingVersionCrn] = *version.CRN
						}
						if insTempCatalogOffering.Plan != nil {
							plan := insTempCatalogOffering.Plan.(*vpcv1.CatalogOfferingVersionPlanIdentity)
							if plan.CRN != nil && *plan.CRN != "" {
								currentOffering[isInstanceTemplateCatalogOfferingPlanCrn] = *plan.CRN
							}
						}
						catOfferingList = append(catOfferingList, currentOffering)
						if err = d.Set(isInstanceTemplateCatalogOffering, catOfferingList); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting catalog_offering: %s", err), "(Data) ibm_is_instance_template", "read", "set-catalog_offering").GetDiag()
						}
					}

					// vni

					networkAttachments := []map[string]interface{}{}
					if instanceTemplate.NetworkAttachments != nil {
						for _, modelItem := range instanceTemplate.NetworkAttachments {
							modelMap, err := dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeToMap(&modelItem)
							if err != nil {
								return diag.FromErr(err)
							}
							networkAttachments = append(networkAttachments, modelMap)
						}
					}
					if err = d.Set("network_attachments", networkAttachments); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting network_attachments: %s", err), "(Data) ibm_is_instance_template", "read", "set-network_attachments").GetDiag()
					}

					primaryNetworkAttachment := []map[string]interface{}{}
					if instanceTemplate.PrimaryNetworkAttachment != nil {
						modelMap, err := dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeToMap(instanceTemplate.PrimaryNetworkAttachment)
						if err != nil {
							return diag.FromErr(err)
						}
						primaryNetworkAttachment = append(primaryNetworkAttachment, modelMap)
					}
					if err = d.Set("primary_network_attachment", primaryNetworkAttachment); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_network_attachment: %s", err), "(Data) ibm_is_instance_template", "read", "set-primary_network_attachment").GetDiag()
					}

					if instanceTemplate.DefaultTrustedProfile != nil {
						if instanceTemplate.DefaultTrustedProfile.AutoLink != nil {
							if err = d.Set(isInstanceDefaultTrustedProfileAutoLink, instanceTemplate.DefaultTrustedProfile.AutoLink); err != nil {
								return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting default_trusted_profile_auto_link: %s", err), "(Data) ibm_is_instance_template", "read", "set-default_trusted_profile_auto_link").GetDiag()
							}
						}
						if instanceTemplate.DefaultTrustedProfile.Target != nil {
							switch reflect.TypeOf(instanceTemplate.DefaultTrustedProfile.Target).String() {
							case "*vpcv1.TrustedProfileIdentityTrustedProfileByID":
								{
									target := instanceTemplate.DefaultTrustedProfile.Target.(*vpcv1.TrustedProfileIdentityByID)
									if err = d.Set(isInstanceDefaultTrustedProfileTarget, target.ID); err != nil {
										return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting default_trusted_profile_target: %s", err), "(Data) ibm_is_instance_template", "read", "set-default_trusted_profile_target").GetDiag()
									}
								}
							case "*vpcv1.TrustedProfileIdentityTrustedProfileByCRN":
								{
									target := instanceTemplate.DefaultTrustedProfile.Target.(*vpcv1.TrustedProfileIdentityByCRN)
									if err = d.Set(isInstanceDefaultTrustedProfileTarget, target.CRN); err != nil {
										return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting default_trusted_profile_target: %s", err), "(Data) ibm_is_instance_template", "read", "set-default_trusted_profile_target").GetDiag()
									}
								}
							}
						}
					}
					if instanceTemplate.Keys != nil {
						keys := []string{}
						for _, intfc := range instanceTemplate.Keys {
							instanceKeyIntf := intfc.(*vpcv1.KeyIdentity)
							keys = append(keys, *instanceKeyIntf.ID)
						}
						if err = d.Set(isInstanceTemplateKeys, keys); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting keys: %s", err), "(Data) ibm_is_instance_template", "read", "set-keys").GetDiag()
						}
					}

					if instanceTemplate.MetadataService != nil {
						if err = d.Set(isInstanceTemplateMetadataServiceEnabled, instanceTemplate.MetadataService.Enabled); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting metadata_service_enabled: %s", err), "(Data) ibm_is_instance_template", "read", "set-metadata_service_enabled").GetDiag()
						}
						metadataService := []map[string]interface{}{}
						metadataServiceMap := map[string]interface{}{}

						metadataServiceMap[isInstanceMetadataServiceEnabled1] = instanceTemplate.MetadataService.Enabled
						if instanceTemplate.MetadataService.Protocol != nil {
							metadataServiceMap[isInstanceMetadataServiceProtocol] = instanceTemplate.MetadataService.Protocol
						}
						if instanceTemplate.MetadataService.ResponseHopLimit != nil {
							metadataServiceMap[isInstanceMetadataServiceRespHopLimit] = instanceTemplate.MetadataService.ResponseHopLimit
						}

						metadataService = append(metadataService, metadataServiceMap)
						if err = d.Set("metadata_service", metadataService); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting metadata_service: %s", err), "(Data) ibm_is_instance_template", "read", "set-metadata_service").GetDiag()
						}
					}

					if instanceTemplate.Profile != nil {
						instanceProfileIntf := instanceTemplate.Profile
						identity := instanceProfileIntf.(*vpcv1.InstanceProfileIdentity)
						if err = d.Set("profile", *identity.Name); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting profile: %s", err), "(Data) ibm_is_instance_template", "read", "set-profile").GetDiag()
						}
					}

					if instanceTemplate.PlacementTarget != nil {
						placementTargetList := []map[string]interface{}{}
						placementTargetMap := dataSourceInstanceTemplateCollectionTemplatesPlacementTargetToMap(*instanceTemplate.PlacementTarget.(*vpcv1.InstancePlacementTargetPrototype))
						placementTargetList = append(placementTargetList, placementTargetMap)
						if err = d.Set("placement_target", placementTargetList); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting placement_target: %s", err), "(Data) ibm_is_instance_template", "read", "set-placement_target").GetDiag()
						}
					}

					if instanceTemplate.PrimaryNetworkInterface != nil {
						interfaceList := make([]map[string]interface{}, 0)
						currentPrimNic := map[string]interface{}{}
						currentPrimNic[isInstanceTemplateNicName] = *instanceTemplate.PrimaryNetworkInterface.Name
						if instanceTemplate.PrimaryNetworkInterface.PrimaryIP != nil {
							primaryipIntf := instanceTemplate.PrimaryNetworkInterface.PrimaryIP
							primaryIpList := make([]map[string]interface{}, 0)
							currentPrimIp := map[string]interface{}{}
							switch reflect.TypeOf(primaryipIntf).String() {
							case "*vpcv1.NetworkInterfaceIPPrototype":
								{
									primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototype)
									if primaryip.Address != nil {
										currentPrimNic[isInstanceTemplateNicPrimaryIpv4Address] = primaryip.Address
										currentPrimIp[isInstanceTemplateNicReservedIpAddress] = *primaryip.Address
									}
									if primaryip.ID != nil {
										currentPrimIp[isInstanceTemplateNicReservedIpId] = *primaryip.ID
									}
								}
							case "*vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext":
								{
									primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext)
									if primaryip.Address != nil {
										currentPrimNic[isInstanceTemplateNicPrimaryIpv4Address] = primaryip.Address
										currentPrimIp[isInstanceTemplateNicReservedIpAddress] = *primaryip.Address
									}
									if primaryip.Name != nil {
										currentPrimIp[isInstanceTemplateNicReservedIpName] = *primaryip.Name
									}
								}
							case "*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity":
								{
									primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity)
									if primaryip.ID != nil {
										currentPrimIp[isInstanceTemplateNicReservedIpId] = *primaryip.ID
									}
								}
							}
							primaryIpList = append(primaryIpList, currentPrimIp)
							currentPrimNic[isInstanceTemplateNicPrimaryIP] = primaryIpList
						}
						subInf := instanceTemplate.PrimaryNetworkInterface.Subnet
						subnetIdentity := subInf.(*vpcv1.SubnetIdentity)
						currentPrimNic[isInstanceTemplateNicSubnet] = *subnetIdentity.ID

						if len(instanceTemplate.PrimaryNetworkInterface.SecurityGroups) != 0 {
							secgrpList := []string{}
							for i := 0; i < len(instanceTemplate.PrimaryNetworkInterface.SecurityGroups); i++ {
								secGrpInf := instanceTemplate.PrimaryNetworkInterface.SecurityGroups[i]
								secGrpIdentity := secGrpInf.(*vpcv1.SecurityGroupIdentity)
								secgrpList = append(secgrpList, string(*secGrpIdentity.ID))
							}
							currentPrimNic[isInstanceTemplateNicSecurityGroups] = flex.NewStringSet(schema.HashString, secgrpList)
						}
						interfaceList = append(interfaceList, currentPrimNic)
						if err = d.Set("primary_network_interface", interfaceList); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_network_interface: %s", err), "(Data) ibm_is_instance_template", "read", "set-primary_network_interface").GetDiag()
						}
					}

					if instanceTemplate.NetworkInterfaces != nil {
						interfacesList := make([]map[string]interface{}, 0)
						for _, intfc := range instanceTemplate.NetworkInterfaces {
							currentNic := map[string]interface{}{}
							currentNic[isInstanceTemplateNicName] = *intfc.Name
							if intfc.PrimaryIP != nil {
								primaryipIntf := intfc.PrimaryIP
								switch reflect.TypeOf(primaryipIntf).String() {
								case "*vpcv1.NetworkInterfaceIPPrototype":
									{
										primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototype)
										currentNic[isInstanceTemplateNicPrimaryIpv4Address] = primaryip.Address

									}
								case "*vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext":
									{
										primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext)
										currentNic[isInstanceTemplateNicPrimaryIpv4Address] = primaryip.Address
									}
								}
							}
							//currentNic[isInstanceTemplateNicAllowIpSpoofing] = intfc.AllowIpSpoofing
							subInf := intfc.Subnet
							subnetIdentity := subInf.(*vpcv1.SubnetIdentity)
							currentNic[isInstanceTemplateNicSubnet] = *subnetIdentity.ID
							if len(intfc.SecurityGroups) != 0 {
								secgrpList := []string{}
								for i := 0; i < len(intfc.SecurityGroups); i++ {
									secGrpInf := intfc.SecurityGroups[i]
									secGrpIdentity := secGrpInf.(*vpcv1.SecurityGroupIdentity)
									secgrpList = append(secgrpList, string(*secGrpIdentity.ID))
								}
								currentNic[isInstanceTemplateNicSecurityGroups] = flex.NewStringSet(schema.HashString, secgrpList)
							}

							interfacesList = append(interfacesList, currentNic)
						}
						if err = d.Set("network_interfaces", interfacesList); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting network_interfaces: %s", err), "(Data) ibm_is_instance_template", "read", "set-network_interfaces").GetDiag()
						}
					}

					if instanceTemplate.TotalVolumeBandwidth != nil {
						if err = d.Set("total_volume_bandwidth", int(*instanceTemplate.TotalVolumeBandwidth)); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting total_volume_bandwidth: %s", err), "(Data) ibm_is_instance_template", "read", "set-total_volume_bandwidth").GetDiag()
						}
					}

					if instanceTemplate.Image != nil {
						imageInf := instanceTemplate.Image
						imageIdentity := imageInf.(*vpcv1.ImageIdentity)
						if err = d.Set("image", imageIdentity.ID); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting image: %s", err), "(Data) ibm_is_instance_template", "read", "set-image").GetDiag()
						}
					}

					if instanceTemplate.VPC != nil {
						vpcInf := instanceTemplate.VPC
						vpcRef := vpcInf.(*vpcv1.VPCIdentity)
						if err = d.Set("vpc", vpcRef.ID); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vpc: %s", err), "(Data) ibm_is_instance_template", "read", "set-vpc").GetDiag()
						}
					}

					if instanceTemplate.Zone != nil {
						zoneInf := instanceTemplate.Zone
						zone := zoneInf.(*vpcv1.ZoneIdentity)
						if err = d.Set("zone", zone.Name); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting zone: %s", err), "(Data) ibm_is_instance_template", "read", "set-zone").GetDiag()
						}
					}

					interfacesList := make([]map[string]interface{}, 0)
					if instanceTemplate.VolumeAttachments != nil {
						for _, volume := range instanceTemplate.VolumeAttachments {
							volumeAttach := map[string]interface{}{}
							volumeAttach[isInstanceTemplateVolAttName] = *volume.Name
							volumeAttach[isInstanceTemplateDeleteVolume] = *volume.DeleteVolumeOnInstanceDelete
							volumeIntf := volume.Volume
							volumeInst := volumeIntf.(*vpcv1.VolumeAttachmentPrototypeVolume)
							newVolumeArr := []map[string]interface{}{}
							newVolume := map[string]interface{}{}

							if volumeInst.ID != nil {
								volumeAttach[isInstanceTemplateVolAttVolume] = *volumeInst.ID
							}
							// bandwidth changes
							if volumeInst.Bandwidth != nil {
								newVolume["bandwidth"] = volumeInst.Bandwidth
							}
							if volumeInst.Capacity != nil {
								newVolume[isInstanceTemplateVolAttVolCapacity] = *volumeInst.Capacity
							}
							if volumeInst.Profile != nil {
								profile := volumeInst.Profile.(*vpcv1.VolumeProfileIdentity)
								newVolume[isInstanceTemplateVolAttVolProfile] = profile.Name
							}

							if volumeInst.Iops != nil {
								newVolume[isInstanceTemplateVolAttVolIops] = *volumeInst.Iops
							}
							if volumeInst.EncryptionKey != nil {
								encryptionKey := volumeInst.EncryptionKey.(*vpcv1.EncryptionKeyIdentity)
								newVolume[isInstanceTemplateVolAttVolEncryptionKey] = *encryptionKey.CRN
							}
							if volumeInst.AllowedUse != nil {
								modelMap, err := DataSourceIBMIsVolumeAllowedUseToMap(volumeInst.AllowedUse)
								if err != nil {
									return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allowed_use: %s", err), "(Data) ibm_is_instance_template", "read", "set-allowed_use").GetDiag()
								}
								newVolume["allowed_use"] = []map[string]interface{}{modelMap}
							}

							if volumeInst.SourceSnapshot != nil {
								sourceSnapshot := volumeInst.SourceSnapshot.(*vpcv1.SnapshotIdentity)
								newVolume["source_snapshot"] = *sourceSnapshot.ID
							}
							if volumeInst.UserTags != nil {
								newVolume[isInstanceTemplateVolAttTags] = volumeInst.UserTags
							}
							newVolumeArr = append(newVolumeArr, newVolume)
							volumeAttach[isInstanceTemplateVolAttVolPrototype] = newVolumeArr

							interfacesList = append(interfacesList, volumeAttach)
						}
						if err = d.Set("volume_attachments", interfacesList); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting volume_attachments: %s", err), "(Data) ibm_is_instance_template", "read", "set-volume_attachments").GetDiag()
						}
					}

					if instanceTemplate.BootVolumeAttachment != nil {
						bootVolList := make([]map[string]interface{}, 0)
						bootVol := map[string]interface{}{}

						bootVol[isInstanceTemplateDeleteVol] = *instanceTemplate.BootVolumeAttachment.DeleteVolumeOnInstanceDelete
						if instanceTemplate.BootVolumeAttachment.Volume != nil {
							volumeIntf := instanceTemplate.BootVolumeAttachment.Volume
							bootVol[isInstanceTemplateName] = volumeIntf.Name
							bootVol[isInstanceTemplateVol] = volumeIntf.Name
							// bandwidth changes
							if volumeIntf.Bandwidth != nil {
								bootVol["bandwidth"] = volumeIntf.Bandwidth
							}
							bootVol[isInstanceTemplateBootSize] = volumeIntf.Capacity
							if instanceTemplate.BootVolumeAttachment.Volume.Profile != nil {
								volProfIntf := instanceTemplate.BootVolumeAttachment.Volume.Profile
								volProfInst := volProfIntf.(*vpcv1.VolumeProfileIdentity)
								bootVol[isInstanceTemplateBootProfile] = volProfInst.Name
							}
							if instanceTemplate.BootVolumeAttachment.Volume.AllowedUse != nil {
								modelMap, err := DataSourceIBMIsVolumeAllowedUseToMap(instanceTemplate.BootVolumeAttachment.Volume.AllowedUse)
								if err != nil {
									return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allowed_use: %s", err), "(Data) ibm_is_instance_template", "read", "set-allowed_use").GetDiag()
								}
								bootVol["allowed_use"] = []map[string]interface{}{modelMap}

							}
							if instanceTemplate.BootVolumeAttachment.Volume.UserTags != nil {
								bootVol[isInstanceTemplateBootVolumeTags] = instanceTemplate.BootVolumeAttachment.Volume.UserTags
							}
						}
						bootVolList = append(bootVolList, bootVol)
						if err = d.Set("boot_volume_attachment", bootVolList); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting boot_volume_attachment: %s", err), "(Data) ibm_is_instance_template", "read", "set-boot_volume_attachment").GetDiag()
						}
					}

					if instanceTemplate.ResourceGroup != nil {
						rg := instanceTemplate.ResourceGroup
						if err = d.Set("resource_group", rg.ID); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_is_instance_template", "read", "set-resource_group").GetDiag()
						}
					}
				}
			case *vpcv1.InstanceTemplateInstanceBySourceSnapshotInstanceTemplateContext:
				instanceTemplate := v
				if name == *instanceTemplate.Name {
					flag = true
					d.SetId(*instanceTemplate.ID)
					if err = d.Set("href", instanceTemplate.Href); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_instance_template", "read", "set-href").GetDiag()
					}
					if err = d.Set("crn", instanceTemplate.CRN); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_instance_template", "read", "set-crn").GetDiag()
					}
					if err = d.Set("name", instanceTemplate.Name); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_instance_template", "read", "set-name").GetDiag()
					}
					if err = d.Set("user_data", instanceTemplate.UserData); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting user_data: %s", err), "(Data) ibm_is_instance_template", "read", "set-user_data").GetDiag()
					}
					if err = d.Set("confidential_compute_mode", instanceTemplate.ConfidentialComputeMode); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting confidential_compute_mode: %s", err), "(Data) ibm_is_instance_template", "read", "set-confidential_compute_mode").GetDiag()
					}
					if err = d.Set("enable_secure_boot", instanceTemplate.EnableSecureBoot); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting enable_secure_boot: %s", err), "(Data) ibm_is_instance_template", "read", "set-enable_secure_boot").GetDiag()
					}
					// cluster changes
					if !core.IsNil(instanceTemplate.ClusterNetworkAttachments) {
						clusterNetworkAttachments := []map[string]interface{}{}
						for _, clusterNetworkAttachmentsItem := range instanceTemplate.ClusterNetworkAttachments {
							clusterNetworkAttachmentsItemMap, err := DataSourceIBMIsInstanceTemplateInstanceClusterNetworkAttachmentPrototypeInstanceContextToMap(&clusterNetworkAttachmentsItem) // #nosec G601
							if err != nil {
								return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_template", "read", "cluster_network_attachments-to-map").GetDiag()
							}
							clusterNetworkAttachments = append(clusterNetworkAttachments, clusterNetworkAttachmentsItemMap)
						}
						if err = d.Set("cluster_network_attachments", clusterNetworkAttachments); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cluster_network_attachments: %s", err), "(Data) ibm_is_instance_template", "read", "set-cluster_network_attachments").GetDiag()
						}
					}

					// vni

					networkAttachments := []map[string]interface{}{}
					if instanceTemplate.NetworkAttachments != nil {
						for _, modelItem := range instanceTemplate.NetworkAttachments {
							modelMap, err := dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeToMap(&modelItem)
							if err != nil {
								return diag.FromErr(err)
							}
							networkAttachments = append(networkAttachments, modelMap)
						}
					}
					if err = d.Set("network_attachments", networkAttachments); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting network_attachments: %s", err), "(Data) ibm_is_instance_template", "read", "set-network_attachments").GetDiag()
					}

					primaryNetworkAttachment := []map[string]interface{}{}
					if instanceTemplate.PrimaryNetworkAttachment != nil {
						modelMap, err := dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeToMap(instanceTemplate.PrimaryNetworkAttachment)
						if err != nil {
							return diag.FromErr(err)
						}
						primaryNetworkAttachment = append(primaryNetworkAttachment, modelMap)
					}
					if err = d.Set("primary_network_attachment", primaryNetworkAttachment); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_network_attachment: %s", err), "(Data) ibm_is_instance_template", "read", "set-primary_network_attachment").GetDiag()
					}

					if instanceTemplate.DefaultTrustedProfile != nil {
						if instanceTemplate.DefaultTrustedProfile.AutoLink != nil {
							if err = d.Set(isInstanceDefaultTrustedProfileAutoLink, instanceTemplate.DefaultTrustedProfile.AutoLink); err != nil {
								return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting default_trusted_profile_auto_link: %s", err), "(Data) ibm_is_instance_template", "read", "set-default_trusted_profile_auto_link").GetDiag()
							}
						}
						if instanceTemplate.DefaultTrustedProfile.Target != nil {
							switch reflect.TypeOf(instanceTemplate.DefaultTrustedProfile.Target).String() {
							case "*vpcv1.TrustedProfileIdentityTrustedProfileByID":
								{
									target := instanceTemplate.DefaultTrustedProfile.Target.(*vpcv1.TrustedProfileIdentityByID)
									if err = d.Set(isInstanceDefaultTrustedProfileTarget, target.ID); err != nil {
										return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting default_trusted_profile_target: %s", err), "(Data) ibm_is_instance_template", "read", "set-default_trusted_profile_target").GetDiag()
									}
								}
							case "*vpcv1.TrustedProfileIdentityTrustedProfileByCRN":
								{
									target := instanceTemplate.DefaultTrustedProfile.Target.(*vpcv1.TrustedProfileIdentityByCRN)
									if err = d.Set(isInstanceDefaultTrustedProfileTarget, target.CRN); err != nil {
										return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting default_trusted_profile_target: %s", err), "(Data) ibm_is_instance_template", "read", "set-default_trusted_profile_target").GetDiag()
									}
								}
							}
						}
					}
					if instanceTemplate.Keys != nil {
						keys := []string{}
						for _, intfc := range instanceTemplate.Keys {
							instanceKeyIntf := intfc.(*vpcv1.KeyIdentity)
							keys = append(keys, *instanceKeyIntf.ID)
						}
						if err = d.Set(isInstanceTemplateKeys, keys); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting keys: %s", err), "(Data) ibm_is_instance_template", "read", "set-keys").GetDiag()
						}
					}

					if instanceTemplate.MetadataService != nil {
						if err = d.Set(isInstanceTemplateMetadataServiceEnabled, instanceTemplate.MetadataService.Enabled); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting metadata_service_enabled: %s", err), "(Data) ibm_is_instance_template", "read", "set-metadata_service_enabled").GetDiag()
						}
						metadataService := []map[string]interface{}{}
						metadataServiceMap := map[string]interface{}{}

						metadataServiceMap[isInstanceMetadataServiceEnabled1] = instanceTemplate.MetadataService.Enabled
						if instanceTemplate.MetadataService.Protocol != nil {
							metadataServiceMap[isInstanceMetadataServiceProtocol] = instanceTemplate.MetadataService.Protocol
						}
						if instanceTemplate.MetadataService.ResponseHopLimit != nil {
							metadataServiceMap[isInstanceMetadataServiceRespHopLimit] = instanceTemplate.MetadataService.ResponseHopLimit
						}

						metadataService = append(metadataService, metadataServiceMap)
						if err = d.Set("metadata_service", metadataService); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting metadata_service: %s", err), "(Data) ibm_is_instance_template", "read", "set-metadata_service").GetDiag()
						}
					}

					if instanceTemplate.Profile != nil {
						instanceProfileIntf := instanceTemplate.Profile
						identity := instanceProfileIntf.(*vpcv1.InstanceProfileIdentity)
						if err = d.Set("profile", *identity.Name); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting profile: %s", err), "(Data) ibm_is_instance_template", "read", "set-profile").GetDiag()
						}
					}

					if instanceTemplate.PlacementTarget != nil {
						placementTargetList := []map[string]interface{}{}
						placementTargetMap := dataSourceInstanceTemplateCollectionTemplatesPlacementTargetToMap(*instanceTemplate.PlacementTarget.(*vpcv1.InstancePlacementTargetPrototype))
						placementTargetList = append(placementTargetList, placementTargetMap)
						if err = d.Set("placement_target", placementTargetList); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting placement_target: %s", err), "(Data) ibm_is_instance_template", "read", "set-placement_target").GetDiag()
						}
					}

					if instanceTemplate.PrimaryNetworkInterface != nil {
						interfaceList := make([]map[string]interface{}, 0)
						currentPrimNic := map[string]interface{}{}
						currentPrimNic[isInstanceTemplateNicName] = *instanceTemplate.PrimaryNetworkInterface.Name
						if instanceTemplate.PrimaryNetworkInterface.PrimaryIP != nil {
							primaryipIntf := instanceTemplate.PrimaryNetworkInterface.PrimaryIP
							primaryIpList := make([]map[string]interface{}, 0)
							currentPrimIp := map[string]interface{}{}
							switch reflect.TypeOf(primaryipIntf).String() {
							case "*vpcv1.NetworkInterfaceIPPrototype":
								{
									primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototype)
									if primaryip.Address != nil {
										currentPrimNic[isInstanceTemplateNicPrimaryIpv4Address] = primaryip.Address
										currentPrimIp[isInstanceTemplateNicReservedIpAddress] = *primaryip.Address
									}
									if primaryip.ID != nil {
										currentPrimIp[isInstanceTemplateNicReservedIpId] = *primaryip.ID
									}
								}
							case "*vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext":
								{
									primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext)
									if primaryip.Address != nil {
										currentPrimNic[isInstanceTemplateNicPrimaryIpv4Address] = primaryip.Address
										currentPrimIp[isInstanceTemplateNicReservedIpAddress] = *primaryip.Address
									}
									if primaryip.Name != nil {
										currentPrimIp[isInstanceTemplateNicReservedIpName] = *primaryip.Name
									}
								}
							case "*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity":
								{
									primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity)
									if primaryip.ID != nil {
										currentPrimIp[isInstanceTemplateNicReservedIpId] = *primaryip.ID
									}
								}
							}
							primaryIpList = append(primaryIpList, currentPrimIp)
							currentPrimNic[isInstanceTemplateNicPrimaryIP] = primaryIpList
						}
						subInf := instanceTemplate.PrimaryNetworkInterface.Subnet
						subnetIdentity := subInf.(*vpcv1.SubnetIdentity)
						currentPrimNic[isInstanceTemplateNicSubnet] = *subnetIdentity.ID

						if len(instanceTemplate.PrimaryNetworkInterface.SecurityGroups) != 0 {
							secgrpList := []string{}
							for i := 0; i < len(instanceTemplate.PrimaryNetworkInterface.SecurityGroups); i++ {
								secGrpInf := instanceTemplate.PrimaryNetworkInterface.SecurityGroups[i]
								secGrpIdentity := secGrpInf.(*vpcv1.SecurityGroupIdentity)
								secgrpList = append(secgrpList, string(*secGrpIdentity.ID))
							}
							currentPrimNic[isInstanceTemplateNicSecurityGroups] = flex.NewStringSet(schema.HashString, secgrpList)
						}
						interfaceList = append(interfaceList, currentPrimNic)
						if err = d.Set("primary_network_interface", interfaceList); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_network_interface: %s", err), "(Data) ibm_is_instance_template", "read", "set-primary_network_interface").GetDiag()
						}
					}

					if instanceTemplate.NetworkInterfaces != nil {
						interfacesList := make([]map[string]interface{}, 0)
						for _, intfc := range instanceTemplate.NetworkInterfaces {
							currentNic := map[string]interface{}{}
							currentNic[isInstanceTemplateNicName] = *intfc.Name
							if intfc.PrimaryIP != nil {
								primaryipIntf := intfc.PrimaryIP
								switch reflect.TypeOf(primaryipIntf).String() {
								case "*vpcv1.NetworkInterfaceIPPrototype":
									{
										primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototype)
										currentNic[isInstanceTemplateNicPrimaryIpv4Address] = primaryip.Address

									}
								case "*vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext":
									{
										primaryip := primaryipIntf.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext)
										currentNic[isInstanceTemplateNicPrimaryIpv4Address] = primaryip.Address
									}
								}
							}
							//currentNic[isInstanceTemplateNicAllowIpSpoofing] = intfc.AllowIpSpoofing
							subInf := intfc.Subnet
							subnetIdentity := subInf.(*vpcv1.SubnetIdentity)
							currentNic[isInstanceTemplateNicSubnet] = *subnetIdentity.ID
							if len(intfc.SecurityGroups) != 0 {
								secgrpList := []string{}
								for i := 0; i < len(intfc.SecurityGroups); i++ {
									secGrpInf := intfc.SecurityGroups[i]
									secGrpIdentity := secGrpInf.(*vpcv1.SecurityGroupIdentity)
									secgrpList = append(secgrpList, string(*secGrpIdentity.ID))
								}
								currentNic[isInstanceTemplateNicSecurityGroups] = flex.NewStringSet(schema.HashString, secgrpList)
							}

							interfacesList = append(interfacesList, currentNic)
						}
						if err = d.Set("network_interfaces", interfacesList); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting network_interfaces: %s", err), "(Data) ibm_is_instance_template", "read", "set-network_interfaces").GetDiag()
						}
					}

					if instanceTemplate.TotalVolumeBandwidth != nil {
						if err = d.Set("total_volume_bandwidth", int(*instanceTemplate.TotalVolumeBandwidth)); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting total_volume_bandwidth: %s", err), "(Data) ibm_is_instance_template", "read", "set-total_volume_bandwidth").GetDiag()
						}
					}

					if instanceTemplate.VPC != nil {
						vpcInf := instanceTemplate.VPC
						vpcRef := vpcInf.(*vpcv1.VPCIdentity)
						if err = d.Set("vpc", vpcRef.ID); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vpc: %s", err), "(Data) ibm_is_instance_template", "read", "set-vpc").GetDiag()
						}
					}

					if instanceTemplate.Zone != nil {
						zoneInf := instanceTemplate.Zone
						zone := zoneInf.(*vpcv1.ZoneIdentity)
						if err = d.Set("zone", zone.Name); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting zone: %s", err), "(Data) ibm_is_instance_template", "read", "set-zone").GetDiag()
						}
					}

					interfacesList := make([]map[string]interface{}, 0)
					if instanceTemplate.VolumeAttachments != nil {
						for _, volume := range instanceTemplate.VolumeAttachments {
							volumeAttach := map[string]interface{}{}
							volumeAttach[isInstanceTemplateVolAttName] = *volume.Name
							volumeAttach[isInstanceTemplateDeleteVolume] = *volume.DeleteVolumeOnInstanceDelete
							volumeIntf := volume.Volume
							volumeInst := volumeIntf.(*vpcv1.VolumeAttachmentPrototypeVolume)
							newVolumeArr := []map[string]interface{}{}
							newVolume := map[string]interface{}{}

							if volumeInst.ID != nil {
								volumeAttach[isInstanceTemplateVolAttVolume] = *volumeInst.ID
							}
							// bandwidth changes
							if volumeInst.Bandwidth != nil {
								newVolume["bandwidth"] = volumeInst.Bandwidth
							}
							if volumeInst.Capacity != nil {
								newVolume[isInstanceTemplateVolAttVolCapacity] = *volumeInst.Capacity
							}
							if volumeInst.Profile != nil {
								profile := volumeInst.Profile.(*vpcv1.VolumeProfileIdentity)
								newVolume[isInstanceTemplateVolAttVolProfile] = profile.Name
							}

							if volumeInst.Iops != nil {
								newVolume[isInstanceTemplateVolAttVolIops] = *volumeInst.Iops
							}
							if volumeInst.EncryptionKey != nil {
								encryptionKey := volumeInst.EncryptionKey.(*vpcv1.EncryptionKeyIdentity)
								newVolume[isInstanceTemplateVolAttVolEncryptionKey] = *encryptionKey.CRN
							}
							if volumeInst.UserTags != nil {
								newVolume[isInstanceTemplateVolAttTags] = volumeInst.UserTags
							}
							newVolumeArr = append(newVolumeArr, newVolume)
							volumeAttach[isInstanceTemplateVolAttVolPrototype] = newVolumeArr

							interfacesList = append(interfacesList, volumeAttach)
						}
						if err = d.Set("volume_attachments", interfacesList); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting volume_attachments: %s", err), "(Data) ibm_is_instance_template", "read", "set-volume_attachments").GetDiag()
						}
					}

					if instanceTemplate.BootVolumeAttachment != nil {
						bootVolList := make([]map[string]interface{}, 0)
						bootVol := map[string]interface{}{}

						bootVol[isInstanceTemplateDeleteVol] = *instanceTemplate.BootVolumeAttachment.DeleteVolumeOnInstanceDelete
						if instanceTemplate.BootVolumeAttachment.Volume != nil {
							volumeIntf := instanceTemplate.BootVolumeAttachment.Volume
							bootVol[isInstanceTemplateName] = volumeIntf.Name
							bootVol[isInstanceTemplateVol] = volumeIntf.Name
							// bandwidth changes
							if volumeIntf.Bandwidth != nil {
								bootVol["bandwidth"] = volumeIntf.Bandwidth
							}
							bootVol[isInstanceTemplateBootSize] = volumeIntf.Capacity
							if instanceTemplate.BootVolumeAttachment.Volume.Profile != nil {
								volProfIntf := instanceTemplate.BootVolumeAttachment.Volume.Profile
								volProfInst := volProfIntf.(*vpcv1.VolumeProfileIdentity)
								bootVol[isInstanceTemplateBootProfile] = volProfInst.Name
							}
							if instanceTemplate.BootVolumeAttachment.Volume.UserTags != nil {
								bootVol[isInstanceTemplateBootVolumeTags] = instanceTemplate.BootVolumeAttachment.Volume.UserTags
							}
							if instanceTemplate.BootVolumeAttachment.Volume.SourceSnapshot != nil {
								sourceSnapshotMap, _ := DataSourceIBMIsInstanceTemplateSnapshotIdentityToMap(instanceTemplate.BootVolumeAttachment.Volume.SourceSnapshot)
								bootVol["source_snapshot"] = []map[string]interface{}{sourceSnapshotMap}
							}
							if instanceTemplate.BootVolumeAttachment.Volume.AllowedUse != nil {
								modelMap, err := DataSourceIBMIsVolumeAllowedUseToMap(instanceTemplate.BootVolumeAttachment.Volume.AllowedUse)
								if err != nil {
									return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allowed_use: %s", err), "(Data) ibm_is_instance_template", "read", "set-allowed_use").GetDiag()
								}
								bootVol["allowed_use"] = []map[string]interface{}{modelMap}

							}
						}
						bootVolList = append(bootVolList, bootVol)
						if err = d.Set("boot_volume_attachment", bootVolList); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting boot_volume_attachment: %s", err), "(Data) ibm_is_instance_template", "read", "set-boot_volume_attachment").GetDiag()
						}
					}

					if instanceTemplate.ResourceGroup != nil {
						rg := instanceTemplate.ResourceGroup
						if err = d.Set("resource_group", rg.ID); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_is_instance_template", "read", "set-resource_group").GetDiag()
						}
					}
				}
			default:
				log.Println("Unknown or unsupported instance template context type")
			}
		}
		if !flag {
			err = fmt.Errorf("No Instance Template found with name %s", name)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceTemplatesWithContext failed: %s", err.Error()), "(Data) ibm_is_instance_template", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return nil
}

func dataSourceInstanceTemplateCollectionTemplatePlacementTargetToMap(placementTargetItem vpcv1.InstancePlacementTargetPrototype) (placementTargetMap map[string]interface{}) {
	placementTargetMap = map[string]interface{}{}

	if placementTargetItem.ID != nil {
		placementTargetMap["id"] = placementTargetItem.ID
	}
	if placementTargetItem.CRN != nil {
		placementTargetMap["crn"] = placementTargetItem.CRN
	}
	if placementTargetItem.Href != nil {
		placementTargetMap["href"] = placementTargetItem.Href
	}

	return placementTargetMap
}

func dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeToMap(model *vpcv1.InstanceNetworkAttachmentPrototype) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	virtualNetworkInterfaceMap, err := dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceToMap(model.VirtualNetworkInterface)
	if err != nil {
		return modelMap, err
	}
	modelMap["virtual_network_interface"] = []map[string]interface{}{virtualNetworkInterfaceMap}
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByIDToMap(model *vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHrefToMap(model *vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = model.Href
	return modelMap, nil
}
func dataSourceIBMIsInstanceTemplateVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextToMap(model vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID); ok {
		return dataSourceIBMIsInstanceTemplateVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByIDToMap(model.(*vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByID))
	} else if _, ok := model.(*vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref); ok {
		return dataSourceIBMIsInstanceTemplateVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHrefToMap(model.(*vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextByHref))
	} else if _, ok := model.(*vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContext); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContext)
		if model.ID != nil {
			modelMap["id"] = model.ID
		}
		if model.Href != nil {
			modelMap["href"] = model.Href
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextIntf subtype encountered")
	}
}
func dataSourceIBMIsInstanceTemplateVirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContextToMap(model *vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Address != nil {
		modelMap["address"] = model.Address
	}
	if model.AutoDelete != nil {
		modelMap["auto_delete"] = model.AutoDelete
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}
func dataSourceIBMIsInstanceTemplateVirtualNetworkInterfaceIPPrototypeToMap(model vpcv1.VirtualNetworkInterfaceIPPrototypeIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContext); ok {
		return dataSourceIBMIsInstanceTemplateVirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContextToMap(model.(*vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPIdentityVirtualNetworkInterfaceIPsContext))
	} else if _, ok := model.(*vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext); ok {
		return dataSourceIBMIsInstanceTemplateVirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContextToMap(model.(*vpcv1.VirtualNetworkInterfaceIPPrototypeReservedIPPrototypeVirtualNetworkInterfaceIPsContext))
	} else if _, ok := model.(*vpcv1.VirtualNetworkInterfaceIPPrototype); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.VirtualNetworkInterfaceIPPrototype)
		if model.ID != nil {
			modelMap["id"] = model.ID
		}
		if model.Href != nil {
			modelMap["href"] = model.Href
		}
		if model.Address != nil {
			modelMap["address"] = model.Address
		}
		if model.AutoDelete != nil {
			modelMap["auto_delete"] = model.AutoDelete
		}
		if model.Name != nil {
			modelMap["name"] = model.Name
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.VirtualNetworkInterfaceIPPrototypeIntf subtype encountered")
	}
}
func dataSourceIBMIsInstanceTemplateVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextToMap(model vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID); ok {
		return dataSourceIBMIsInstanceTemplateVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByIDToMap(model.(*vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID))
	} else if _, ok := model.(*vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref); ok {
		return dataSourceIBMIsInstanceTemplateVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHrefToMap(model.(*vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref))
	} else if _, ok := model.(*vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContext); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContext)
		if model.ID != nil {
			modelMap["id"] = model.ID
		}
		if model.Href != nil {
			modelMap["href"] = model.Href
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextIntf subtype encountered")
	}
}
func dataSourceIBMIsInstanceTemplateVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByIDToMap(model *vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHrefToMap(model *vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextByHref) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = model.Href
	return modelMap, nil
}
func dataSourceIBMIsInstanceTemplateVirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContextToMap(model *vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Address != nil {
		modelMap["address"] = model.Address
	}
	if model.AutoDelete != nil {
		modelMap["auto_delete"] = model.AutoDelete
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}
func dataSourceIBMIsInstanceTemplateVirtualNetworkInterfacePrimaryIPPrototypeToMap(model vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContext); ok {
		return dataSourceIBMIsInstanceTemplateVirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContextToMap(model.(*vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPIdentityVirtualNetworkInterfacePrimaryIPContext))
	} else if _, ok := model.(*vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext); ok {
		return dataSourceIBMIsInstanceTemplateVirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContextToMap(model.(*vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeReservedIPPrototypeVirtualNetworkInterfacePrimaryIPContext))
	} else if _, ok := model.(*vpcv1.VirtualNetworkInterfacePrimaryIPPrototype); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.VirtualNetworkInterfacePrimaryIPPrototype)
		if model.ID != nil {
			modelMap["id"] = model.ID
		}
		if model.Href != nil {
			modelMap["href"] = model.Href
		}
		if model.Address != nil {
			modelMap["address"] = model.Address
		}
		if model.AutoDelete != nil {
			modelMap["auto_delete"] = model.AutoDelete
		}
		if model.Name != nil {
			modelMap["name"] = model.Name
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeIntf subtype encountered")
	}
}

func dataSourceIBMIsInstanceTemplateResourceGroupIdentityByIDToMap(model *vpcv1.ResourceGroupIdentityByID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID
	return modelMap, nil
}
func dataSourceIBMIsInstanceTemplateResourceGroupIdentityToMap(model vpcv1.ResourceGroupIdentityIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.ResourceGroupIdentityByID); ok {
		return dataSourceIBMIsInstanceTemplateResourceGroupIdentityByIDToMap(model.(*vpcv1.ResourceGroupIdentityByID))
	} else if _, ok := model.(*vpcv1.ResourceGroupIdentity); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.ResourceGroupIdentity)
		if model.ID != nil {
			modelMap["id"] = model.ID
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.ResourceGroupIdentityIntf subtype encountered")
	}
}

func dataSourceIBMIsInstanceTemplateSecurityGroupIdentityToMap(model vpcv1.SecurityGroupIdentityIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.SecurityGroupIdentityByID); ok {
		return dataSourceIBMIsInstanceTemplateSecurityGroupIdentityByIDToMap(model.(*vpcv1.SecurityGroupIdentityByID))
	} else if _, ok := model.(*vpcv1.SecurityGroupIdentityByCRN); ok {
		return dataSourceIBMIsInstanceTemplateSecurityGroupIdentityByCRNToMap(model.(*vpcv1.SecurityGroupIdentityByCRN))
	} else if _, ok := model.(*vpcv1.SecurityGroupIdentityByHref); ok {
		return dataSourceIBMIsInstanceTemplateSecurityGroupIdentityByHrefToMap(model.(*vpcv1.SecurityGroupIdentityByHref))
	} else if _, ok := model.(*vpcv1.SecurityGroupIdentity); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.SecurityGroupIdentity)
		if model.ID != nil {
			modelMap["id"] = model.ID
		}
		if model.CRN != nil {
			modelMap["crn"] = model.CRN
		}
		if model.Href != nil {
			modelMap["href"] = model.Href
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.SecurityGroupIdentityIntf subtype encountered")
	}
}

func dataSourceIBMIsInstanceTemplateSecurityGroupIdentityByIDToMap(model *vpcv1.SecurityGroupIdentityByID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateSecurityGroupIdentityByCRNToMap(model *vpcv1.SecurityGroupIdentityByCRN) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = model.CRN
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateSecurityGroupIdentityByHrefToMap(model *vpcv1.SecurityGroupIdentityByHref) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = model.Href
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateSubnetIdentityToMap(model vpcv1.SubnetIdentityIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.SubnetIdentityByID); ok {
		return dataSourceIBMIsInstanceTemplateSubnetIdentityByIDToMap(model.(*vpcv1.SubnetIdentityByID))
	} else if _, ok := model.(*vpcv1.SubnetIdentityByCRN); ok {
		return dataSourceIBMIsInstanceTemplateSubnetIdentityByCRNToMap(model.(*vpcv1.SubnetIdentityByCRN))
	} else if _, ok := model.(*vpcv1.SubnetIdentityByHref); ok {
		return dataSourceIBMIsInstanceTemplateSubnetIdentityByHrefToMap(model.(*vpcv1.SubnetIdentityByHref))
	} else if _, ok := model.(*vpcv1.SubnetIdentity); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.SubnetIdentity)
		if model.ID != nil {
			modelMap["id"] = model.ID
		}
		if model.CRN != nil {
			modelMap["crn"] = model.CRN
		}
		if model.Href != nil {
			modelMap["href"] = model.Href
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.SubnetIdentityIntf subtype encountered")
	}
}

func dataSourceIBMIsInstanceTemplateSubnetIdentityByIDToMap(model *vpcv1.SubnetIdentityByID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateSubnetIdentityByCRNToMap(model *vpcv1.SubnetIdentityByCRN) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = model.CRN
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateSubnetIdentityByHrefToMap(model *vpcv1.SubnetIdentityByHref) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = model.Href
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceToMap(model vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeInstanceNetworkAttachmentContext); ok {
		return dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeInstanceNetworkAttachmentContextToMap(model.(*vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeInstanceNetworkAttachmentContext))
	} else if _, ok := model.(*vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentity); ok {
		return dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityToMap(model.(*vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentity))
	} else if _, ok := model.(*vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterface); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterface)
		if model.AllowIPSpoofing != nil {
			modelMap["allow_ip_spoofing"] = model.AllowIPSpoofing
		}
		if model.AutoDelete != nil {
			modelMap["auto_delete"] = model.AutoDelete
		}
		if model.EnableInfrastructureNat != nil {
			modelMap["enable_infrastructure_nat"] = model.EnableInfrastructureNat
		}
		if model.Ips != nil {
			ips := []map[string]interface{}{}
			for _, ipsItem := range model.Ips {
				ipsItemMap, err := dataSourceIBMIsInstanceTemplateVirtualNetworkInterfaceIPPrototypeToMap(ipsItem)
				if err != nil {
					return modelMap, err
				}
				ips = append(ips, ipsItemMap)
			}
			modelMap["ips"] = ips
		}
		if model.Name != nil {
			modelMap["name"] = model.Name
		}
		if model.PrimaryIP != nil {
			primaryIPMap, err := dataSourceIBMIsInstanceTemplateVirtualNetworkInterfacePrimaryIPPrototypeToMap(model.PrimaryIP)
			if err != nil {
				return modelMap, err
			}
			modelMap["primary_ip"] = []map[string]interface{}{primaryIPMap}
		}
		if model.ProtocolStateFilteringMode != nil {
			modelMap["protocol_state_filtering_mode"] = model.ProtocolStateFilteringMode
		}
		if model.ResourceGroup != nil {
			resourceGroupMap, err := dataSourceIBMIsInstanceTemplateResourceGroupIdentityToMap(model.ResourceGroup)
			if err != nil {
				return modelMap, err
			}
			modelMap["resource_group"] = []map[string]interface{}{resourceGroupMap}
		}
		if model.SecurityGroups != nil {
			securityGroups := []map[string]interface{}{}
			for _, securityGroupsItem := range model.SecurityGroups {
				securityGroupsItemMap, err := dataSourceIBMIsInstanceTemplateSecurityGroupIdentityToMap(securityGroupsItem)
				if err != nil {
					return modelMap, err
				}
				securityGroups = append(securityGroups, securityGroupsItemMap)
			}
			modelMap["security_groups"] = securityGroups
		}
		if model.Subnet != nil {
			subnetMap, err := dataSourceIBMIsInstanceTemplateSubnetIdentityToMap(model.Subnet)
			if err != nil {
				return modelMap, err
			}
			modelMap["subnet"] = []map[string]interface{}{subnetMap}
		}
		if model.ID != nil {
			modelMap["id"] = model.ID
		}
		if model.Href != nil {
			modelMap["href"] = model.Href
		}
		if model.CRN != nil {
			modelMap["crn"] = model.CRN
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceIntf subtype encountered")
	}
}
func dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeInstanceNetworkAttachmentContextToMap(model *vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfacePrototypeInstanceNetworkAttachmentContext) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AllowIPSpoofing != nil {
		modelMap["allow_ip_spoofing"] = model.AllowIPSpoofing
	}
	if model.AutoDelete != nil {
		modelMap["auto_delete"] = model.AutoDelete
	}
	if model.EnableInfrastructureNat != nil {
		modelMap["enable_infrastructure_nat"] = model.EnableInfrastructureNat
	}
	if model.Ips != nil {
		ips := []map[string]interface{}{}
		for _, ipsItem := range model.Ips {
			ipsItemMap, err := dataSourceIBMIsInstanceTemplateVirtualNetworkInterfaceIPPrototypeToMap(ipsItem)
			if err != nil {
				return modelMap, err
			}
			ips = append(ips, ipsItemMap)
		}
		modelMap["ips"] = ips
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.PrimaryIP != nil {
		primaryIPMap, err := dataSourceIBMIsInstanceTemplateVirtualNetworkInterfacePrimaryIPPrototypeToMap(model.PrimaryIP)
		if err != nil {
			return modelMap, err
		}
		modelMap["primary_ip"] = []map[string]interface{}{primaryIPMap}
	}
	if model.ProtocolStateFilteringMode != nil {
		modelMap["protocol_state_filtering_mode"] = model.ProtocolStateFilteringMode
	}
	if model.ResourceGroup != nil {
		resourceGroupMap, err := dataSourceIBMIsInstanceTemplateResourceGroupIdentityToMap(model.ResourceGroup)
		if err != nil {
			return modelMap, err
		}
		modelMap["resource_group"] = []map[string]interface{}{resourceGroupMap}
	}
	if model.SecurityGroups != nil {
		securityGroups := []map[string]interface{}{}
		for _, securityGroupsItem := range model.SecurityGroups {
			securityGroupsItemMap, err := dataSourceIBMIsInstanceTemplateSecurityGroupIdentityToMap(securityGroupsItem)
			if err != nil {
				return modelMap, err
			}
			securityGroups = append(securityGroups, securityGroupsItemMap)
		}
		modelMap["security_groups"] = securityGroups
	}
	if model.Subnet != nil {
		subnetMap, err := dataSourceIBMIsInstanceTemplateSubnetIdentityToMap(model.Subnet)
		if err != nil {
			return modelMap, err
		}
		modelMap["subnet"] = []map[string]interface{}{subnetMap}
	}
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityToMap(model vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID); ok {
		return dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByIDToMap(model.(*vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID))
	} else if _, ok := model.(*vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref); ok {
		return dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHrefToMap(model.(*vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref))
	} else if _, ok := model.(*vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN); ok {
		return dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRNToMap(model.(*vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN))
	} else if _, ok := model.(*vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentity); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentity)
		if model.ID != nil {
			modelMap["id"] = model.ID
		}
		if model.Href != nil {
			modelMap["href"] = model.Href
		}
		if model.CRN != nil {
			modelMap["crn"] = model.CRN
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityIntf subtype encountered")
	}
}

func dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByIDToMap(model *vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHrefToMap(model *vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByHref) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = model.Href
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateInstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRNToMap(model *vpcv1.InstanceNetworkAttachmentPrototypeVirtualNetworkInterfaceVirtualNetworkInterfaceIdentityVirtualNetworkInterfaceIdentityByCRN) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = model.CRN
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateNetworkInterfacePrototypeToMap(model *vpcv1.NetworkInterfacePrototype) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AllowIPSpoofing != nil {
		modelMap["allow_ip_spoofing"] = model.AllowIPSpoofing
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.PrimaryIP != nil {
		primaryIPMap, err := dataSourceIBMIsInstanceTemplateNetworkInterfaceIPPrototypeToMap(model.PrimaryIP)
		if err != nil {
			return modelMap, err
		}
		modelMap["primary_ip"] = []map[string]interface{}{primaryIPMap}
	}
	if model.SecurityGroups != nil {
		securityGroups := []map[string]interface{}{}
		for _, securityGroupsItem := range model.SecurityGroups {
			securityGroupsItemMap, err := dataSourceIBMIsInstanceTemplateSecurityGroupIdentityToMap(securityGroupsItem)
			if err != nil {
				return modelMap, err
			}
			securityGroups = append(securityGroups, securityGroupsItemMap)
		}
		modelMap["security_groups"] = securityGroups
	}
	subnetMap, err := dataSourceIBMIsInstanceTemplateSubnetIdentityToMap(model.Subnet)
	if err != nil {
		return modelMap, err
	}
	modelMap["subnet"] = []map[string]interface{}{subnetMap}
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateNetworkInterfaceIPPrototypeToMap(model vpcv1.NetworkInterfaceIPPrototypeIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity); ok {
		return dataSourceIBMIsInstanceTemplateNetworkInterfaceIPPrototypeReservedIPIdentityToMap(model.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity))
	} else if _, ok := model.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext); ok {
		return dataSourceIBMIsInstanceTemplateNetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContextToMap(model.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext))
	} else if _, ok := model.(*vpcv1.NetworkInterfaceIPPrototype); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.NetworkInterfaceIPPrototype)
		if model.ID != nil {
			modelMap["id"] = model.ID
		}
		if model.Href != nil {
			modelMap["href"] = model.Href
		}
		if model.Address != nil {
			modelMap["address"] = model.Address
		}
		if model.AutoDelete != nil {
			modelMap["auto_delete"] = model.AutoDelete
		}
		if model.Name != nil {
			modelMap["name"] = model.Name
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.NetworkInterfaceIPPrototypeIntf subtype encountered")
	}
}

func dataSourceIBMIsInstanceTemplateNetworkInterfaceIPPrototypeReservedIPIdentityToMap(model vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentityIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentityByID); ok {
		return dataSourceIBMIsInstanceTemplateNetworkInterfaceIPPrototypeReservedIPIdentityByIDToMap(model.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentityByID))
	} else if _, ok := model.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentityByHref); ok {
		return dataSourceIBMIsInstanceTemplateNetworkInterfaceIPPrototypeReservedIPIdentityByHrefToMap(model.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentityByHref))
	} else if _, ok := model.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity)
		if model.ID != nil {
			modelMap["id"] = model.ID
		}
		if model.Href != nil {
			modelMap["href"] = model.Href
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentityIntf subtype encountered")
	}
}

func dataSourceIBMIsInstanceTemplateNetworkInterfaceIPPrototypeReservedIPIdentityByIDToMap(model *vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentityByID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateNetworkInterfaceIPPrototypeReservedIPIdentityByHrefToMap(model *vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentityByHref) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = model.Href
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateNetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContextToMap(model *vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Address != nil {
		modelMap["address"] = model.Address
	}
	if model.AutoDelete != nil {
		modelMap["auto_delete"] = model.AutoDelete
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateInstanceCatalogOfferingPrototypeToMap(model vpcv1.InstanceCatalogOfferingPrototypeIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.InstanceCatalogOfferingPrototypeCatalogOfferingByOffering); ok {
		return dataSourceIBMIsInstanceTemplateInstanceCatalogOfferingPrototypeCatalogOfferingByOfferingToMap(model.(*vpcv1.InstanceCatalogOfferingPrototypeCatalogOfferingByOffering))
	} else if _, ok := model.(*vpcv1.InstanceCatalogOfferingPrototypeCatalogOfferingByVersion); ok {
		return dataSourceIBMIsInstanceTemplateInstanceCatalogOfferingPrototypeCatalogOfferingByVersionToMap(model.(*vpcv1.InstanceCatalogOfferingPrototypeCatalogOfferingByVersion))
	} else if _, ok := model.(*vpcv1.InstanceCatalogOfferingPrototype); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.InstanceCatalogOfferingPrototype)
		if model.Offering != nil {
			offeringMap, err := dataSourceIBMIsInstanceTemplateCatalogOfferingIdentityToMap(model.Offering)
			if err != nil {
				return modelMap, err
			}
			modelMap["offering"] = []map[string]interface{}{offeringMap}
		}
		if model.Version != nil {
			versionMap, err := dataSourceIBMIsInstanceTemplateCatalogOfferingVersionIdentityToMap(model.Version)
			if err != nil {
				return modelMap, err
			}
			modelMap["version"] = []map[string]interface{}{versionMap}
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.InstanceCatalogOfferingPrototypeIntf subtype encountered")
	}
}

func dataSourceIBMIsInstanceTemplateCatalogOfferingIdentityToMap(model vpcv1.CatalogOfferingIdentityIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.CatalogOfferingIdentityCatalogOfferingByCRN); ok {
		return dataSourceIBMIsInstanceTemplateCatalogOfferingIdentityCatalogOfferingByCRNToMap(model.(*vpcv1.CatalogOfferingIdentityCatalogOfferingByCRN))
	} else if _, ok := model.(*vpcv1.CatalogOfferingIdentity); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.CatalogOfferingIdentity)
		if model.CRN != nil {
			modelMap["crn"] = model.CRN
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.CatalogOfferingIdentityIntf subtype encountered")
	}
}

func dataSourceIBMIsInstanceTemplateCatalogOfferingIdentityCatalogOfferingByCRNToMap(model *vpcv1.CatalogOfferingIdentityCatalogOfferingByCRN) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = model.CRN
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateCatalogOfferingVersionIdentityToMap(model vpcv1.CatalogOfferingVersionIdentityIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.CatalogOfferingVersionIdentityCatalogOfferingVersionByCRN); ok {
		return dataSourceIBMIsInstanceTemplateCatalogOfferingVersionIdentityCatalogOfferingVersionByCRNToMap(model.(*vpcv1.CatalogOfferingVersionIdentityCatalogOfferingVersionByCRN))
	} else if _, ok := model.(*vpcv1.CatalogOfferingVersionIdentity); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.CatalogOfferingVersionIdentity)
		if model.CRN != nil {
			modelMap["crn"] = model.CRN
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.CatalogOfferingVersionIdentityIntf subtype encountered")
	}
}

func dataSourceIBMIsInstanceTemplateCatalogOfferingVersionIdentityCatalogOfferingVersionByCRNToMap(model *vpcv1.CatalogOfferingVersionIdentityCatalogOfferingVersionByCRN) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = model.CRN
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateInstanceCatalogOfferingPrototypeCatalogOfferingByOfferingToMap(model *vpcv1.InstanceCatalogOfferingPrototypeCatalogOfferingByOffering) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	offeringMap, err := dataSourceIBMIsInstanceTemplateCatalogOfferingIdentityToMap(model.Offering)
	if err != nil {
		return modelMap, err
	}
	modelMap["offering"] = []map[string]interface{}{offeringMap}
	return modelMap, nil
}

func dataSourceIBMIsInstanceTemplateInstanceCatalogOfferingPrototypeCatalogOfferingByVersionToMap(model *vpcv1.InstanceCatalogOfferingPrototypeCatalogOfferingByVersion) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	versionMap, err := dataSourceIBMIsInstanceTemplateCatalogOfferingVersionIdentityToMap(model.Version)
	if err != nil {
		return modelMap, err
	}
	modelMap["version"] = []map[string]interface{}{versionMap}
	return modelMap, nil
}

func DataSourceIBMIsInstanceTemplateInstanceClusterNetworkAttachmentPrototypeInstanceContextToMap(model *vpcv1.InstanceClusterNetworkAttachmentPrototypeInstanceContext) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	clusterNetworkInterfaceMap, err := DataSourceIBMIsInstanceTemplateInstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceToMap(model.ClusterNetworkInterface)
	if err != nil {
		return modelMap, err
	}
	modelMap["cluster_network_interface"] = []map[string]interface{}{clusterNetworkInterfaceMap}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	return modelMap, nil
}

func DataSourceIBMIsInstanceTemplateInstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceToMap(model vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceInstanceClusterNetworkInterfacePrototypeInstanceClusterNetworkAttachment); ok {
		return DataSourceIBMIsInstanceTemplateInstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceInstanceClusterNetworkInterfacePrototypeInstanceClusterNetworkAttachmentToMap(model.(*vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceInstanceClusterNetworkInterfacePrototypeInstanceClusterNetworkAttachment))
	} else if _, ok := model.(*vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceClusterNetworkInterfaceIdentity); ok {
		return DataSourceIBMIsInstanceTemplateInstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceClusterNetworkInterfaceIdentityToMap(model.(*vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceClusterNetworkInterfaceIdentity))
	} else if _, ok := model.(*vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterface); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterface)
		if model.AutoDelete != nil {
			modelMap["auto_delete"] = *model.AutoDelete
		}
		if model.Name != nil {
			modelMap["name"] = *model.Name
		}
		if model.PrimaryIP != nil {
			primaryIPMap, err := DataSourceIBMIsInstanceTemplateClusterNetworkInterfacePrimaryIPPrototypeToMap(model.PrimaryIP)
			if err != nil {
				return modelMap, err
			}
			modelMap["primary_ip"] = []map[string]interface{}{primaryIPMap}
		}
		if model.Subnet != nil {
			subnetMap, err := DataSourceIBMIsInstanceTemplateClusterNetworkSubnetIdentityToMap(model.Subnet)
			if err != nil {
				return modelMap, err
			}
			modelMap["subnet"] = []map[string]interface{}{subnetMap}
		}
		if model.ID != nil {
			modelMap["id"] = *model.ID
		}
		if model.Href != nil {
			modelMap["href"] = *model.Href
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceIntf subtype encountered")
	}
}

func DataSourceIBMIsInstanceTemplateClusterNetworkInterfacePrimaryIPPrototypeToMap(model vpcv1.ClusterNetworkInterfacePrimaryIPPrototypeIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.ClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPIdentityClusterNetworkInterfacePrimaryIPContext); ok {
		return DataSourceIBMIsInstanceTemplateClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPIdentityClusterNetworkInterfacePrimaryIPContextToMap(model.(*vpcv1.ClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPIdentityClusterNetworkInterfacePrimaryIPContext))
	} else if _, ok := model.(*vpcv1.ClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPPrototypeClusterNetworkInterfacePrimaryIPContext); ok {
		return DataSourceIBMIsInstanceTemplateClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPPrototypeClusterNetworkInterfacePrimaryIPContextToMap(model.(*vpcv1.ClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPPrototypeClusterNetworkInterfacePrimaryIPContext))
	} else if _, ok := model.(*vpcv1.ClusterNetworkInterfacePrimaryIPPrototype); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.ClusterNetworkInterfacePrimaryIPPrototype)
		if model.ID != nil {
			modelMap["id"] = *model.ID
		}
		if model.Href != nil {
			modelMap["href"] = *model.Href
		}
		if model.Address != nil {
			modelMap["address"] = *model.Address
		}
		if model.AutoDelete != nil {
			modelMap["auto_delete"] = *model.AutoDelete
		}
		if model.Name != nil {
			modelMap["name"] = *model.Name
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.ClusterNetworkInterfacePrimaryIPPrototypeIntf subtype encountered")
	}
}

func DataSourceIBMIsInstanceTemplateClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPIdentityClusterNetworkInterfacePrimaryIPContextToMap(model vpcv1.ClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPIdentityClusterNetworkInterfacePrimaryIPContextIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.ClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPIdentityClusterNetworkInterfacePrimaryIPContextByID); ok {
		return DataSourceIBMIsInstanceTemplateClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPIdentityClusterNetworkInterfacePrimaryIPContextByIDToMap(model.(*vpcv1.ClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPIdentityClusterNetworkInterfacePrimaryIPContextByID))
	} else if _, ok := model.(*vpcv1.ClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPIdentityClusterNetworkInterfacePrimaryIPContextByHref); ok {
		return DataSourceIBMIsInstanceTemplateClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPIdentityClusterNetworkInterfacePrimaryIPContextByHrefToMap(model.(*vpcv1.ClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPIdentityClusterNetworkInterfacePrimaryIPContextByHref))
	} else if _, ok := model.(*vpcv1.ClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPIdentityClusterNetworkInterfacePrimaryIPContext); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.ClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPIdentityClusterNetworkInterfacePrimaryIPContext)
		if model.ID != nil {
			modelMap["id"] = *model.ID
		}
		if model.Href != nil {
			modelMap["href"] = *model.Href
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.ClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPIdentityClusterNetworkInterfacePrimaryIPContextIntf subtype encountered")
	}
}

func DataSourceIBMIsInstanceTemplateClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPIdentityClusterNetworkInterfacePrimaryIPContextByIDToMap(model *vpcv1.ClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPIdentityClusterNetworkInterfacePrimaryIPContextByID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	return modelMap, nil
}

func DataSourceIBMIsInstanceTemplateClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPIdentityClusterNetworkInterfacePrimaryIPContextByHrefToMap(model *vpcv1.ClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPIdentityClusterNetworkInterfacePrimaryIPContextByHref) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	return modelMap, nil
}

func DataSourceIBMIsInstanceTemplateClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPPrototypeClusterNetworkInterfacePrimaryIPContextToMap(model *vpcv1.ClusterNetworkInterfacePrimaryIPPrototypeClusterNetworkSubnetReservedIPPrototypeClusterNetworkInterfacePrimaryIPContext) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Address != nil {
		modelMap["address"] = *model.Address
	}
	if model.AutoDelete != nil {
		modelMap["auto_delete"] = *model.AutoDelete
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	return modelMap, nil
}

func DataSourceIBMIsInstanceTemplateClusterNetworkSubnetIdentityToMap(model vpcv1.ClusterNetworkSubnetIdentityIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.ClusterNetworkSubnetIdentityByID); ok {
		return DataSourceIBMIsInstanceTemplateClusterNetworkSubnetIdentityByIDToMap(model.(*vpcv1.ClusterNetworkSubnetIdentityByID))
	} else if _, ok := model.(*vpcv1.ClusterNetworkSubnetIdentityByHref); ok {
		return DataSourceIBMIsInstanceTemplateClusterNetworkSubnetIdentityByHrefToMap(model.(*vpcv1.ClusterNetworkSubnetIdentityByHref))
	} else if _, ok := model.(*vpcv1.ClusterNetworkSubnetIdentity); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.ClusterNetworkSubnetIdentity)
		if model.ID != nil {
			modelMap["id"] = *model.ID
		}
		if model.Href != nil {
			modelMap["href"] = *model.Href
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.ClusterNetworkSubnetIdentityIntf subtype encountered")
	}
}

func DataSourceIBMIsInstanceTemplateClusterNetworkSubnetIdentityByIDToMap(model *vpcv1.ClusterNetworkSubnetIdentityByID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	return modelMap, nil
}

func DataSourceIBMIsInstanceTemplateClusterNetworkSubnetIdentityByHrefToMap(model *vpcv1.ClusterNetworkSubnetIdentityByHref) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	return modelMap, nil
}

func DataSourceIBMIsInstanceTemplateInstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceInstanceClusterNetworkInterfacePrototypeInstanceClusterNetworkAttachmentToMap(model *vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceInstanceClusterNetworkInterfacePrototypeInstanceClusterNetworkAttachment) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AutoDelete != nil {
		modelMap["auto_delete"] = *model.AutoDelete
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.PrimaryIP != nil {
		primaryIPMap, err := DataSourceIBMIsInstanceTemplateClusterNetworkInterfacePrimaryIPPrototypeToMap(model.PrimaryIP)
		if err != nil {
			return modelMap, err
		}
		modelMap["primary_ip"] = []map[string]interface{}{primaryIPMap}
	}
	if model.Subnet != nil {
		subnetMap, err := DataSourceIBMIsInstanceTemplateClusterNetworkSubnetIdentityToMap(model.Subnet)
		if err != nil {
			return modelMap, err
		}
		modelMap["subnet"] = []map[string]interface{}{subnetMap}
	}
	return modelMap, nil
}

func DataSourceIBMIsInstanceTemplateInstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceClusterNetworkInterfaceIdentityToMap(model vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceClusterNetworkInterfaceIdentityIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceClusterNetworkInterfaceIdentityClusterNetworkInterfaceIdentityByID); ok {
		return DataSourceIBMIsInstanceTemplateInstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceClusterNetworkInterfaceIdentityClusterNetworkInterfaceIdentityByIDToMap(model.(*vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceClusterNetworkInterfaceIdentityClusterNetworkInterfaceIdentityByID))
	} else if _, ok := model.(*vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceClusterNetworkInterfaceIdentityClusterNetworkInterfaceIdentityByHref); ok {
		return DataSourceIBMIsInstanceTemplateInstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceClusterNetworkInterfaceIdentityClusterNetworkInterfaceIdentityByHrefToMap(model.(*vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceClusterNetworkInterfaceIdentityClusterNetworkInterfaceIdentityByHref))
	} else if _, ok := model.(*vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceClusterNetworkInterfaceIdentity); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceClusterNetworkInterfaceIdentity)
		if model.ID != nil {
			modelMap["id"] = *model.ID
		}
		if model.Href != nil {
			modelMap["href"] = *model.Href
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceClusterNetworkInterfaceIdentityIntf subtype encountered")
	}
}

func DataSourceIBMIsInstanceTemplateInstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceClusterNetworkInterfaceIdentityClusterNetworkInterfaceIdentityByIDToMap(model *vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceClusterNetworkInterfaceIdentityClusterNetworkInterfaceIdentityByID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	return modelMap, nil
}

func DataSourceIBMIsInstanceTemplateInstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceClusterNetworkInterfaceIdentityClusterNetworkInterfaceIdentityByHrefToMap(model *vpcv1.InstanceClusterNetworkAttachmentPrototypeClusterNetworkInterfaceClusterNetworkInterfaceIdentityClusterNetworkInterfaceIdentityByHref) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	return modelMap, nil
}

func DataSourceIBMIsInstanceTemplateSnapshotIdentityToMap(model vpcv1.SnapshotIdentityIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.SnapshotIdentityByID); ok {
		return DataSourceIBMIsInstanceTemplateSnapshotIdentityByIDToMap(model.(*vpcv1.SnapshotIdentityByID))
	} else if _, ok := model.(*vpcv1.SnapshotIdentityByCRN); ok {
		return DataSourceIBMIsInstanceTemplateSnapshotIdentityByCRNToMap(model.(*vpcv1.SnapshotIdentityByCRN))
	} else if _, ok := model.(*vpcv1.SnapshotIdentityByHref); ok {
		return DataSourceIBMIsInstanceTemplateSnapshotIdentityByHrefToMap(model.(*vpcv1.SnapshotIdentityByHref))
	} else if _, ok := model.(*vpcv1.SnapshotIdentity); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.SnapshotIdentity)
		if model.ID != nil {
			modelMap["id"] = *model.ID
		}
		if model.CRN != nil {
			modelMap["crn"] = *model.CRN
		}
		if model.Href != nil {
			modelMap["href"] = *model.Href
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.SnapshotIdentityIntf subtype encountered")
	}
}

func DataSourceIBMIsInstanceTemplateSnapshotIdentityByIDToMap(model *vpcv1.SnapshotIdentityByID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	return modelMap, nil
}

func DataSourceIBMIsInstanceTemplateSnapshotIdentityByCRNToMap(model *vpcv1.SnapshotIdentityByCRN) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	return modelMap, nil
}

func DataSourceIBMIsInstanceTemplateSnapshotIdentityByHrefToMap(model *vpcv1.SnapshotIdentityByHref) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	return modelMap, nil
}

func DataSourceIBMIsVolumeAllowedUseToMap(model *vpcv1.VolumeAllowedUsePrototype) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.BareMetalServer != nil {
		modelMap["bare_metal_server"] = *model.BareMetalServer
	}
	if model.Instance != nil {
		modelMap["instance"] = *model.Instance
	}
	if model.ApiVersion != nil {
		modelMap["api_version"] = *model.ApiVersion
	}
	return modelMap, nil
}

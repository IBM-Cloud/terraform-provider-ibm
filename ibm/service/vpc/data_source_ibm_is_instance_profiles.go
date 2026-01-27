// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isInstanceProfiles = "profiles"
)

func DataSourceIBMISInstanceProfiles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISInstanceProfilesRead,

		Schema: map[string]*schema.Schema{

			isInstanceProfiles: {
				Type:        schema.TypeList,
				Description: "List of instance profile maps",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"family": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The product family this virtual server instance profile belongs to.",
						},

						// spot changes
						"availability_class": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"default": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The default availability class for an instance with this profile.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"values": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The permitted values for this profile field.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value for this profile field.",
									},
								},
							},
						},
						// cluster changes
						"cluster_network_attachment_count": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"default": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"values": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The permitted values for this profile field.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"max": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum value for this profile field.",
									},
									"min": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum value for this profile field.",
									},
									"step": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"supported_cluster_network_profiles": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The cluster network profiles that support this instance profile.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this cluster network profile.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The globally unique name for this cluster network profile.",
									},
									"resource_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type.",
									},
								},
							},
						},
						"confidential_compute_modes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"default": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The default confidential compute mode for this profile.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"values": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The supported confidential compute modes.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"secure_boot_modes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"default": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "The default secure boot mode for this profile.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"values": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The supported `enable_secure_boot` values for an instance using this profile.",
										Elem: &schema.Schema{
											Type: schema.TypeBool,
										},
									},
								},
							},
						},
						"architecture": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The default OS architecture for an instance with this profile.",
						},
						"architecture_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for the OS architecture.",
						},

						"architecture_values": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The supported OS architecture(s) for an instance with this profile.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"bandwidth": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"value": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The value for this profile field.",
									},
									"default": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The default value for this profile field.",
									},
									"max": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum value for this profile field.",
									},
									"min": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum value for this profile field.",
									},
									"step": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The increment step value for this profile field.",
									},
									"values": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The permitted values for this profile field.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
								},
							},
						},
						"gpu_count": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "GPU count of this profile",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"value": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The value for this profile field.",
									},
									"default": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The default value for this profile field.",
									},
									"max": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum value for this profile field.",
									},
									"min": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum value for this profile field.",
									},
									"step": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The increment step value for this profile field.",
									},
									"values": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The permitted values for this profile field.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
								},
							},
						},
						"gpu_manufacturer": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "GPU manufacturer of this profile",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"values": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The possible GPU manufacturer(s) for an instance with this profile",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"gpu_memory": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "GPU memory of this profile",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"value": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The value for this profile field.",
									},
									"default": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The default value for this profile field.",
									},
									"max": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum value for this profile field.",
									},
									"min": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum value for this profile field.",
									},
									"step": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The increment step value for this profile field.",
									},
									"values": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The permitted values for this profile field.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
								},
							},
						},
						"gpu_model": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "GPU model of this profile",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"values": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The possible GPU model(s) for an instance with this profile",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"reservation_terms": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The type for this profile field",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"values": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The supported committed use terms for a reservation using this profile",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"total_volume_bandwidth": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The amount of bandwidth (in megabits per second) allocated exclusively to instance storage volumes. An increase in this value will result in a corresponding decrease to total_network_bandwidth.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"value": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The value for this profile field.",
									},
									"default": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The default value for this profile field.",
									},
									"max": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum value for this profile field.",
									},
									"min": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum value for this profile field.",
									},
									"step": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The increment step value for this profile field.",
									},
									"values": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The permitted values for this profile field.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
								},
							},
						},
						"disks": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Collection of the instance profile's disks.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"quantity": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The type for this profile field.",
												},
												"value": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The value for this profile field.",
												},
												"default": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The default value for this profile field.",
												},
												"max": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The maximum value for this profile field.",
												},
												"min": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The minimum value for this profile field.",
												},
												"step": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The increment step value for this profile field.",
												},
												"values": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The permitted values for this profile field.",
													Elem: &schema.Schema{
														Type: schema.TypeInt,
													},
												},
											},
										},
									},
									"size": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The type for this profile field.",
												},
												"value": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The value for this profile field.",
												},
												"default": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The default value for this profile field.",
												},
												"max": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The maximum value for this profile field.",
												},
												"min": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The minimum value for this profile field.",
												},
												"step": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The increment step value for this profile field.",
												},
												"values": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The permitted values for this profile field.",
													Elem: &schema.Schema{
														Type: schema.TypeInt,
													},
												},
											},
										},
									},
									"supported_interface_types": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"default": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The disk interface used for attaching the disk.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.",
												},
												"type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The type for this profile field.",
												},
												"values": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The supported disk interfaces used for attaching the disk.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this virtual server instance profile.",
						},
						"memory": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"value": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The value for this profile field.",
									},
									"default": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The default value for this profile field.",
									},
									"max": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum value for this profile field.",
									},
									"min": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum value for this profile field.",
									},
									"step": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The increment step value for this profile field.",
									},
									"values": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The permitted values for this profile field.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
								},
							},
						},
						"network_attachment_count": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum value for this profile field",
									},
									"min": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum value for this profile field",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
								},
							},
						},
						"network_interface_count": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum value for this profile field",
									},
									"min": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum value for this profile field",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
								},
							},
						},
						"numa_count": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"value": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The value for this profile field.",
									},
								},
							},
						},
						"port_speed": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"value": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The value for this profile field.",
									},
								},
							},
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the instance profile.",
						},
						"vcpu_architecture": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The VCPU architecture for an instance with this profile.",
									},
								},
							},
						},
						"vcpu_count": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"value": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The value for this profile field.",
									},
									"default": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The default value for this profile field.",
									},
									"max": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum value for this profile field.",
									},
									"min": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum value for this profile field.",
									},
									"step": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The increment step value for this profile field.",
									},
									"values": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The permitted values for this profile field.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
								},
							},
						},
						"volume_bandwidth_qos_modes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"default": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The default volume bandwidth QoS mode for this profile.",
									},
									"values": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The permitted volume bandwidth QoS modes for an instance using this profile.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"vcpu_manufacturer": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The VCPU manufacturer for an instance with this profile.",
									},
								},
							},
						},
						// shared core changes
						"vcpu_burst_limit": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The permitted value for VCPU burst limit percentage for an instance with this profile.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The value for this profile field.",
									},
								},
							},
						},
						"vcpu_percentage": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The permitted values for VCPU percentage for an instance with this profile.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"default": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The default value for this profile field.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"values": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The permitted values for this profile field.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
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

func dataSourceIBMISInstanceProfilesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	err := instanceProfilesList(context, d, meta)
	if err != nil {
		return err
	}
	return nil
}

func instanceProfilesList(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_profiles", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	listInstanceProfilesOptions := &vpcv1.ListInstanceProfilesOptions{}
	availableProfiles, _, err := sess.ListInstanceProfilesWithContext(context, listInstanceProfilesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceProfilesWithContext failed: %s", err.Error()), "(Data) ibm_is_instance_profiles", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	profilesInfo := make([]map[string]interface{}, 0)
	for _, profile := range availableProfiles.Profiles {

		l := map[string]interface{}{
			"name":   *profile.Name,
			"family": *profile.Family,
		}
		availabilityClassMap, err := DataSourceIBMIsInstanceProfilesInstanceProfileAvailabilityClassToMap(profile.AvailabilityClass)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_profiles", "read", "availability_class-to-map").GetDiag()
		}
		l["availability_class"] = []map[string]interface{}{availabilityClassMap}

		if profile.OsArchitecture != nil {
			if profile.OsArchitecture.Default != nil {
				l["architecture"] = *profile.OsArchitecture.Default
			}
			if profile.OsArchitecture.Type != nil {
				l["architecture_type"] = profile.OsArchitecture.Type
			}
			if profile.OsArchitecture.Values != nil {
				l["architecture_values"] = profile.OsArchitecture.Values
			}
		}
		if profile.Status != nil {
			l["status"] = *profile.Status
		}
		if profile.Bandwidth != nil {
			bandwidthList := []map[string]interface{}{}
			bandwidthMap := dataSourceInstanceProfileBandwidthToMap(*profile.Bandwidth.(*vpcv1.InstanceProfileBandwidth))
			bandwidthList = append(bandwidthList, bandwidthMap)
			l["bandwidth"] = bandwidthList
		}

		// cluster changes

		supportedClusterNetworkProfiles := []map[string]interface{}{}
		for _, supportedClusterNetworkProfilesItem := range profile.SupportedClusterNetworkProfiles {
			supportedClusterNetworkProfilesItemMap, err := DataSourceIBMIsInstanceProfilesClusterNetworkProfileReferenceToMap(&supportedClusterNetworkProfilesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_profiles", "read", "supported_cluster_network_profiles-to-map").GetDiag()
			}
			supportedClusterNetworkProfiles = append(supportedClusterNetworkProfiles, supportedClusterNetworkProfilesItemMap)
		}
		l["supported_cluster_network_profiles"] = supportedClusterNetworkProfiles

		clusterNetworkAttachmentCountMap, err := DataSourceIBMIsInstanceProfilesInstanceProfileClusterNetworkAttachmentCountToMap(profile.ClusterNetworkAttachmentCount)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_profiles", "read", "cluster_network_attachment_count-to-map").GetDiag()
		}
		l["cluster_network_attachment_count"] = []map[string]interface{}{clusterNetworkAttachmentCountMap}

		if profile.GpuCount != nil {
			l["gpu_count"] = dataSourceInstanceProfileFlattenGPUCount(*profile.GpuCount.(*vpcv1.InstanceProfileGpu))
		}

		if profile.GpuMemory != nil {
			l["gpu_memory"] = dataSourceInstanceProfileFlattenGPUMemory(*profile.GpuMemory.(*vpcv1.InstanceProfileGpuMemory))
		}

		if profile.GpuManufacturer != nil {
			l["gpu_manufacturer"] = dataSourceInstanceProfileFlattenGPUManufacturer(*profile.GpuManufacturer)
		}

		if profile.GpuModel != nil {
			l["gpu_model"] = dataSourceInstanceProfileFlattenGPUModel(*profile.GpuModel)
		}

		if profile.ReservationTerms != nil {
			l["reservation_terms"] = dataSourceInstanceProfileFlattenReservationTerms(*profile.ReservationTerms)
		}

		if profile.TotalVolumeBandwidth != nil {
			l["total_volume_bandwidth"] = dataSourceInstanceProfileFlattenTotalVolumeBandwidth(*profile.TotalVolumeBandwidth.(*vpcv1.InstanceProfileVolumeBandwidth))
		}

		if profile.Disks != nil {
			disksList := []map[string]interface{}{}
			for _, disksItem := range profile.Disks {
				disksList = append(disksList, dataSourceInstanceProfileDisksToMap(disksItem))
			}
			l["disks"] = disksList
		}
		if profile.Href != nil {
			l["href"] = profile.Href
		}
		confidentialComputeModes := []map[string]interface{}{}
		if profile.ConfidentialComputeModes != nil {
			modelMap, err := dataSourceIBMIsInstanceProfileInstanceProfileSupportedConfidentialComputeModesToMap(profile.ConfidentialComputeModes)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_profiles", "read", "confidential_compute_modes-to-map").GetDiag()
			}
			confidentialComputeModes = append(confidentialComputeModes, modelMap)
		}
		l["confidential_compute_modes"] = confidentialComputeModes

		secureBootModes := []map[string]interface{}{}
		if profile.SecureBootModes != nil {
			modelMap, err := dataSourceIBMIsInstanceProfileInstanceProfileSupportedSecureBootModesToMap(profile.SecureBootModes)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_profiles", "read", "secure_boot_modes-to-map").GetDiag()
			}
			secureBootModes = append(secureBootModes, modelMap)
		}
		l["secure_boot_modes"] = secureBootModes

		if profile.Memory != nil {
			memoryList := []map[string]interface{}{}
			memoryMap := dataSourceInstanceProfileMemoryToMap(*profile.Memory.(*vpcv1.InstanceProfileMemory))
			memoryList = append(memoryList, memoryMap)
			l["memory"] = memoryList
		}
		if profile.NetworkInterfaceCount != nil {
			networkInterfaceCountList := []map[string]interface{}{}
			networkInterfaceCountMap := dataSourceInstanceProfileNetworkInterfaceCount(*profile.NetworkInterfaceCount.(*vpcv1.InstanceProfileNetworkInterfaceCount))
			networkInterfaceCountList = append(networkInterfaceCountList, networkInterfaceCountMap)
			l["network_interface_count"] = networkInterfaceCountList
		}
		if profile.NetworkAttachmentCount != nil {
			networkAttachmentCountList := []map[string]interface{}{}
			networkAttachmentCountMap := dataSourceInstanceProfileNetworkAttachmentCount(*profile.NetworkAttachmentCount.(*vpcv1.InstanceProfileNetworkAttachmentCount))
			networkAttachmentCountList = append(networkAttachmentCountList, networkAttachmentCountMap)
			l["network_attachment_count"] = networkAttachmentCountList
		}
		if profile.NumaCount != nil {
			numaCountList := []map[string]interface{}{}
			numaCountMap := dataSourceInstanceProfileNumaCountToMap(*profile.NumaCount.(*vpcv1.InstanceProfileNumaCount))
			numaCountList = append(numaCountList, numaCountMap)
			l["numa_count"] = numaCountList
		}
		if profile.PortSpeed != nil {
			portSpeedList := []map[string]interface{}{}
			portSpeedMap := dataSourceInstanceProfilePortSpeedToMap(*profile.PortSpeed.(*vpcv1.InstanceProfilePortSpeed))
			portSpeedList = append(portSpeedList, portSpeedMap)
			l["port_speed"] = portSpeedList
		}
		if profile.VcpuArchitecture != nil {
			vcpuArchitectureList := []map[string]interface{}{}
			vcpuArchitectureMap := dataSourceInstanceProfileVcpuArchitectureToMap(*profile.VcpuArchitecture)
			vcpuArchitectureList = append(vcpuArchitectureList, vcpuArchitectureMap)
			l["vcpu_architecture"] = vcpuArchitectureList
		}
		if profile.VcpuCount != nil {
			vcpuCountList := []map[string]interface{}{}
			vcpuCountMap := dataSourceInstanceProfileVcpuCountToMap(*profile.VcpuCount.(*vpcv1.InstanceProfileVcpu))
			vcpuCountList = append(vcpuCountList, vcpuCountMap)
			l["vcpu_count"] = vcpuCountList
		}
		if profile.VolumeBandwidthQosModes != nil {
			volumeBandwidthQosModesList := []map[string]interface{}{}
			volumeBandwidthQosModesMap := dataSourceInstanceProfileVolumeBandwidthQoSModeToMap(*profile.VolumeBandwidthQosModes.(*vpcv1.InstanceProfileVolumeBandwidthQoSModes))
			volumeBandwidthQosModesList = append(volumeBandwidthQosModesList, volumeBandwidthQosModesMap)
			l["volume_bandwidth_qos_modes"] = volumeBandwidthQosModesList
		}
		// Changes for manufacturer for AMD Support.
		if profile.VcpuManufacturer != nil {
			vcpuManufacturerList := []map[string]interface{}{}
			vcpuManufacturerMap := dataSourceInstanceProfileVcpuManufacturerToMap(*profile.VcpuManufacturer.(*vpcv1.InstanceProfileVcpuManufacturer))
			vcpuManufacturerList = append(vcpuManufacturerList, vcpuManufacturerMap)
			l["vcpu_manufacturer"] = vcpuManufacturerList
		}

		if profile.Disks != nil {
			l[isInstanceDisks] = dataSourceInstanceProfileFlattenDisks(profile.Disks)
		}

		vcpuBurstLimitMap, err := DataSourceIBMIsInstanceProfilesInstanceProfileVcpuBurstLimitToMap(profile.VcpuBurstLimit.(*vpcv1.InstanceProfileVcpuBurstLimit))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vcpu_percentage: %s", err), "(Data) ibm_is_instance_profiles", "read", "set-vcpu_percentage").GetDiag()
		}
		l["vcpu_burst_limit"] = []map[string]interface{}{vcpuBurstLimitMap}
		vcpuPercentageMap, err := DataSourceIBMIsInstanceProfilesInstanceProfileVcpuPercentageToMap(profile.VcpuPercentage)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vcpu_percentage: %s", err), "(Data) ibm_is_instance_profiles", "read", "set-vcpu_burst_limit").GetDiag()
		}
		l["vcpu_percentage"] = []map[string]interface{}{vcpuPercentageMap}
		profilesInfo = append(profilesInfo, l)
	}
	d.SetId(dataSourceIBMISInstanceProfilesID(d))
	if err = d.Set("profiles", profilesInfo); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting profiles: %s", err), "(Data) ibm_is_instance_profiles", "read", "set-profiles").GetDiag()
	}
	return nil
}

// dataSourceIBMISInstanceProfilesID returns a reasonable ID for a Instance Profile list.
func dataSourceIBMISInstanceProfilesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMIsInstanceProfilesClusterNetworkProfileReferenceToMap(model *vpcv1.ClusterNetworkProfileReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

func DataSourceIBMIsInstanceProfilesInstanceProfileClusterNetworkAttachmentCountToMap(model vpcv1.InstanceProfileClusterNetworkAttachmentCountIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.InstanceProfileClusterNetworkAttachmentCountDependent); ok {
		return DataSourceIBMIsInstanceProfilesInstanceProfileClusterNetworkAttachmentCountDependentToMap(model.(*vpcv1.InstanceProfileClusterNetworkAttachmentCountDependent))
	} else if _, ok := model.(*vpcv1.InstanceProfileClusterNetworkAttachmentCountEnum); ok {
		return DataSourceIBMIsInstanceProfilesInstanceProfileClusterNetworkAttachmentCountEnumToMap(model.(*vpcv1.InstanceProfileClusterNetworkAttachmentCountEnum))
	} else if _, ok := model.(*vpcv1.InstanceProfileClusterNetworkAttachmentCountRange); ok {
		return DataSourceIBMIsInstanceProfilesInstanceProfileClusterNetworkAttachmentCountRangeToMap(model.(*vpcv1.InstanceProfileClusterNetworkAttachmentCountRange))
	} else if _, ok := model.(*vpcv1.InstanceProfileClusterNetworkAttachmentCount); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.InstanceProfileClusterNetworkAttachmentCount)
		if model.Type != nil {
			modelMap["type"] = *model.Type
		}
		if model.Default != nil {
			modelMap["default"] = flex.IntValue(model.Default)
		}
		if model.Values != nil {
			modelMap["values"] = model.Values
		}
		if model.Max != nil {
			modelMap["max"] = flex.IntValue(model.Max)
		}
		if model.Min != nil {
			modelMap["min"] = flex.IntValue(model.Min)
		}
		if model.Step != nil {
			modelMap["step"] = flex.IntValue(model.Step)
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.InstanceProfileClusterNetworkAttachmentCountIntf subtype encountered")
	}
}

func DataSourceIBMIsInstanceProfilesInstanceProfileClusterNetworkAttachmentCountDependentToMap(model *vpcv1.InstanceProfileClusterNetworkAttachmentCountDependent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = *model.Type
	return modelMap, nil
}

func DataSourceIBMIsInstanceProfilesInstanceProfileClusterNetworkAttachmentCountEnumToMap(model *vpcv1.InstanceProfileClusterNetworkAttachmentCountEnum) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Default != nil {
		modelMap["default"] = flex.IntValue(model.Default)
	}
	modelMap["type"] = *model.Type
	modelMap["values"] = model.Values
	return modelMap, nil
}

func DataSourceIBMIsInstanceProfilesInstanceProfileClusterNetworkAttachmentCountRangeToMap(model *vpcv1.InstanceProfileClusterNetworkAttachmentCountRange) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Max != nil {
		modelMap["max"] = flex.IntValue(model.Max)
	}
	if model.Min != nil {
		modelMap["min"] = flex.IntValue(model.Min)
	}
	if model.Step != nil {
		modelMap["step"] = flex.IntValue(model.Step)
	}
	modelMap["type"] = *model.Type
	return modelMap, nil
}

func DataSourceIBMIsInstanceProfilesInstanceProfileAvailabilityClassToMap(model vpcv1.InstanceProfileAvailabilityClassIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.InstanceProfileAvailabilityClassEnum); ok {
		return DataSourceIBMIsInstanceProfilesInstanceProfileAvailabilityClassEnumToMap(model.(*vpcv1.InstanceProfileAvailabilityClassEnum))
	} else if _, ok := model.(*vpcv1.InstanceProfileAvailabilityClassFixed); ok {
		return DataSourceIBMIsInstanceProfilesInstanceProfileAvailabilityClassFixedToMap(model.(*vpcv1.InstanceProfileAvailabilityClassFixed))
	} else if _, ok := model.(*vpcv1.InstanceProfileAvailabilityClass); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.InstanceProfileAvailabilityClass)
		if model.Default != nil {
			modelMap["default"] = *model.Default
		}
		if model.Type != nil {
			modelMap["type"] = *model.Type
		}
		if model.Values != nil {
			modelMap["values"] = model.Values
		}
		if model.Value != nil {
			modelMap["value"] = *model.Value
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.InstanceProfileAvailabilityClassIntf subtype encountered")
	}
}

func DataSourceIBMIsInstanceProfilesInstanceProfileAvailabilityClassEnumToMap(model *vpcv1.InstanceProfileAvailabilityClassEnum) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["default"] = *model.Default
	return modelMap, nil
}

func DataSourceIBMIsInstanceProfilesInstanceProfileVcpuBurstLimitToMap(model *vpcv1.InstanceProfileVcpuBurstLimit) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = *model.Type
	modelMap["value"] = flex.IntValue(model.Value)
	return modelMap, nil
}
func DataSourceIBMIsInstanceProfilesInstanceProfileVcpuPercentageToMap(model *vpcv1.InstanceProfileVcpuPercentage) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["default"] = flex.IntValue(model.Default)
	modelMap["type"] = *model.Type
	modelMap["values"] = model.Values
	return modelMap, nil
}

func DataSourceIBMIsInstanceProfilesInstanceProfileAvailabilityClassFixedToMap(model *vpcv1.InstanceProfileAvailabilityClassFixed) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = *model.Type
	modelMap["value"] = *model.Value
	return modelMap, nil
}

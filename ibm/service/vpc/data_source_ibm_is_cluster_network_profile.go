// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsClusterNetworkProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsClusterNetworkProfileRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The cluster network profile name.",
			},
			"family": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The product family this cluster network profile belongs to.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this cluster network profile.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"address_configuration_services": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The services providing address configuration for the cluster network profile. Possible values: dhcp, is, is_metadata.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"values": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The permitted values for this profile field",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field",
						},
					},
				},
			},
			"subnet_routing_supported": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The instance profiles that support this cluster network profile.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The value for this profile field",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field",
						},
					},
				},
			},
			"isolation_group_count": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The number of isolation groups supported by this cluster network profile.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The value for this profile field",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field",
						},
					},
				},
			},
			"supported_instance_profiles": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The instance profiles that support this cluster network profile.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this virtual server instance profile.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The globally unique name for this virtual server instance profile.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"zones": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Zones in this region that support this cluster network profile.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this zone.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The globally unique name for this zone.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMIsClusterNetworkProfileRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_cluster_network_profile", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getClusterNetworkProfileOptions := &vpcv1.GetClusterNetworkProfileOptions{}

	getClusterNetworkProfileOptions.SetName(d.Get("name").(string))

	clusterNetworkProfile, _, err := vpcClient.GetClusterNetworkProfileWithContext(context, getClusterNetworkProfileOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetClusterNetworkProfileWithContext failed: %s", err.Error()), "(Data) ibm_is_cluster_network_profile", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*clusterNetworkProfile.Name)

	if err = d.Set("family", clusterNetworkProfile.Family); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting family: %s", err), "(Data) ibm_is_cluster_network_profile", "read", "set-family").GetDiag()
	}

	if err = d.Set("href", clusterNetworkProfile.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_cluster_network_profile", "read", "set-href").GetDiag()
	}

	if err = d.Set("resource_type", clusterNetworkProfile.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_cluster_network_profile", "read", "set-resource_type").GetDiag()
	}

	supportedInstanceProfiles := []map[string]interface{}{}
	for _, supportedInstanceProfilesItem := range clusterNetworkProfile.SupportedInstanceProfiles {
		supportedInstanceProfilesItemMap, err := DataSourceIBMIsClusterNetworkProfileInstanceProfileReferenceToMap(&supportedInstanceProfilesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_cluster_network_profile", "read", "supported_instance_profiles-to-map").GetDiag()
		}
		supportedInstanceProfiles = append(supportedInstanceProfiles, supportedInstanceProfilesItemMap)
	}
	if err = d.Set("supported_instance_profiles", supportedInstanceProfiles); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting supported_instance_profiles: %s", err), "(Data) ibm_is_cluster_network_profile", "read", "set-supported_instance_profiles").GetDiag()
	}

	if clusterNetworkProfile.SubnetRoutingSupported != nil {
		if val, ok := clusterNetworkProfile.SubnetRoutingSupported.(*vpcv1.ClusterNetworkProfileSubnetRoutingSupported); ok {
			subnetRS := map[string]interface{}{
				"type":  *val.Type,
				"value": val.Value,
			}
			if err := d.Set("subnet_routing_supported", []map[string]interface{}{subnetRS}); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting subnet_routing_supported: %s", err), "(Data) ibm_is_cluster_network_profile", "read", "set-subnet_routing_supported").GetDiag()
			}
		}
	}

	if clusterNetworkProfile.AddressConfigurationServices != nil {
		if addressServices, ok := clusterNetworkProfile.AddressConfigurationServices.(*vpcv1.ClusterNetworkProfileAddressConfigurationServices); ok {
			values := []interface{}{}
			for _, v := range addressServices.Values {
				values = append(values, v)
			}
			addressCS := map[string]interface{}{
				"type":   *addressServices.Type,
				"values": values,
			}
			if err := d.Set("address_configuration_services", []map[string]interface{}{addressCS}); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting address_configuration_services: %s", err), "(Data) ibm_is_cluster_network_profile", "read", "set-address_configuration_services").GetDiag()
			}
		}
	}

	if clusterNetworkProfile.IsolationGroupCount != nil {
		if isolationGroupCount, ok := clusterNetworkProfile.IsolationGroupCount.(*vpcv1.ClusterNetworkProfileIsolationGroupCount); ok {
			isolationGC := map[string]interface{}{
				"type":  *isolationGroupCount.Type,
				"value": isolationGroupCount.Value,
			}
			if err := d.Set("isolation_group_count", []map[string]interface{}{isolationGC}); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting isolation_group_count: %s", err), "(Data) ibm_is_cluster_network_profile", "read", "set-isolation_group_count").GetDiag()
			}
		}
	}

	zones := []map[string]interface{}{}
	for _, zonesItem := range clusterNetworkProfile.Zones {
		zonesItemMap, err := DataSourceIBMIsClusterNetworkProfileZoneReferenceToMap(&zonesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_cluster_network_profile", "read", "zones-to-map").GetDiag()
		}
		zones = append(zones, zonesItemMap)
	}
	if err = d.Set("zones", zones); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting zones: %s", err), "(Data) ibm_is_cluster_network_profile", "read", "set-zones").GetDiag()
	}

	return nil
}

func DataSourceIBMIsClusterNetworkProfileInstanceProfileReferenceToMap(model *vpcv1.InstanceProfileReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

func DataSourceIBMIsClusterNetworkProfileZoneReferenceToMap(model *vpcv1.ZoneReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	modelMap["name"] = *model.Name
	return modelMap, nil
}

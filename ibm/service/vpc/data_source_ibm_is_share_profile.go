// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIbmIsShareProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmIsShareProfileRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The file share profile name.",
			},
			"allowed_access_protocols": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The default allowed access protocol modes for shares with this profile.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field.",
						},
						"values": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The possible allowed access protocols for shares with this profile.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"allowed_transit_encryption_modes": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The default allowed transit encryption modes for shares with this profile.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field.",
						},
						"values": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The allowed transit encryption modes for a share with this profile",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"availability_modes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The data availability mode of a share with this profile..",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The default data availability mode for this profile.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field.",
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The value for this profile field",
						},
						"values": {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "The permitted values for this profile field.",
						},
					},
				},
			},
			"bandwidth": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The permitted capacity range (in gigabytes) for a share with this profile.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The default capacity.",
						},
						"max": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The max capacity.",
						},
						"min": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The min capacity.",
						},
						"step": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The increment step value for this profile field.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field.",
						},
						"value": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The value for this profile field",
						},
						"values": {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Description: "The permitted values for this profile field.",
						},
					},
				},
			},
			"capacity": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The permitted capacity range (in gigabytes) for a share with this profile.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The default capacity.",
						},
						"max": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The max capacity.",
						},
						"min": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The min capacity.",
						},
						"step": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The increment step value for this profile field.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field.",
						},
						"value": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The value for this profile field",
						},
						"values": {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Description: "The permitted values for this profile field.",
						},
					},
				},
			},
			"iops": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The permitted IOPS range for a share with this profile.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The default iops.",
						},
						"max": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The max iops.",
						},
						"min": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The min iops.",
						},
						"step": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The increment step value for this profile field.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type for this profile field.",
						},
						"value": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The value for this profile field",
						},
						"values": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The permitted values for this profile field.",
						},
					},
				},
			},
			"family": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The product family this share profile belongs to.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this share profile.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
		},
	}
}

func dataSourceIbmIsShareProfileRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("vpcClient creation failed: %s", err.Error()), "(Data) ibm_is_share_profile", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getShareProfileOptions := &vpcv1.GetShareProfileOptions{}

	getShareProfileOptions.SetName(d.Get("name").(string))

	shareProfile, response, err := vpcClient.GetShareProfileWithContext(context, getShareProfileOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetShareProfileWithContext failed: %s\n%s", err.Error(), response), "(Data) ibm_is_share_profile", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if shareProfile.AllowedAccessProtocols != nil {
		allowedAccessprotocolsList := []map[string]interface{}{}
		allowedAccessprotocols := shareProfile.AllowedAccessProtocols.(*vpcv1.ShareProfileAllowedAccessProtocols)
		allowedAccessprotocolsMap := dataSourceShareProfileAllowedAccessProtocolToMap(*allowedAccessprotocols)
		allowedAccessprotocolsList = append(allowedAccessprotocolsList, allowedAccessprotocolsMap)
		if err = d.Set("allowed_access_protocols", allowedAccessprotocolsList); err != nil {
			err = fmt.Errorf("Error setting allowed_access_protocols: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_share_profile", "read", "set-allowed_access_protocols").GetDiag()
		}
	}
	if shareProfile.AllowedTransitEncryptionModes != nil {
		allowedTEMList := []map[string]interface{}{}
		allowedTEM := shareProfile.AllowedTransitEncryptionModes.(*vpcv1.ShareProfileAllowedTransitEncryptionModes)
		allowedTEMMap := dataSourceShareProfileAllowedTransitEncryptionToMap(*allowedTEM)
		allowedTEMList = append(allowedTEMList, allowedTEMMap)
		if err = d.Set("allowed_transit_encryption_modes", allowedTEMList); err != nil {
			err = fmt.Errorf("Error setting allowed_transit_encryption_modes: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_share_profile", "read", "set-allowed_transit_encryption_modes").GetDiag()
		}
	}
	if shareProfile.AvailabilityModes != nil {
		availabilityModesList := []map[string]interface{}{}
		availabilityModes := shareProfile.AvailabilityModes.(*vpcv1.ShareProfileAvailabilityModes)
		availabilityModesMap := dataSourceShareProfileAvailabilityModesToMap(*availabilityModes)
		availabilityModesList = append(availabilityModesList, availabilityModesMap)
		if err = d.Set("availability_modes", availabilityModesList); err != nil {
			err = fmt.Errorf("Error setting availability_modes: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_share_profile", "read", "set-availability_modes").GetDiag()
		}
	}
	if shareProfile.Bandwidth != nil {
		bandwidthList := []map[string]interface{}{}
		bandwidth := shareProfile.Bandwidth.(*vpcv1.ShareProfileBandwidth)
		bandwidthMap := dataSourceShareProfileBandwidthToMap(*bandwidth)
		bandwidthList = append(bandwidthList, bandwidthMap)
		if err = d.Set("bandwidth", bandwidthList); err != nil {
			err = fmt.Errorf("Error setting bandwidth: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_share_profile", "read", "set-bandwidth").GetDiag()
		}
	}
	if shareProfile.Capacity != nil {
		capacityList := []map[string]interface{}{}
		capacity := shareProfile.Capacity.(*vpcv1.ShareProfileCapacity)
		capacityMap := dataSourceShareProfileCapacityToMap(*capacity)
		capacityList = append(capacityList, capacityMap)
		if err = d.Set("capacity", capacityList); err != nil {
			err = fmt.Errorf("Error setting capacity: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_share_profile", "read", "set-capacity").GetDiag()
		}
	}
	if shareProfile.Iops != nil {
		iopsList := []map[string]interface{}{}
		iops := shareProfile.Iops.(*vpcv1.ShareProfileIops)
		iopsMap := dataSourceShareProfileIopsToMap(*iops)
		iopsList = append(iopsList, iopsMap)
		if err = d.Set("iops", iopsList); err != nil {
			err = fmt.Errorf("Error setting iops: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_share_profile", "read", "set-iops").GetDiag()
		}
	}
	d.SetId(*shareProfile.Name)
	if err = d.Set("family", shareProfile.Family); err != nil {
		err = fmt.Errorf("Error setting family: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_share_profile", "read", "set-family").GetDiag()
	}
	if err = d.Set("href", shareProfile.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_share_profile", "read", "set-href").GetDiag()
	}
	if err = d.Set("resource_type", shareProfile.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_share_profile", "read", "set-resource_type").GetDiag()
	}

	return nil
}
func dataSourceShareProfileIopsToMap(iops vpcv1.ShareProfileIops) (iopsMap map[string]interface{}) {
	iopsMap = map[string]interface{}{}
	if iops.Default != nil {
		iopsMap["default"] = int(*iops.Default)
	}
	iopsMap["max"] = iops.Max
	iopsMap["min"] = iops.Min
	iopsMap["step"] = iops.Step
	iopsMap["type"] = iops.Type
	if iops.Value != nil {
		iopsMap["value"] = int(*iops.Value)
	}
	values := []int{}
	if len(iops.Values) > 0 {
		for _, value := range iops.Values {
			values = append(values, int(value))
		}
		iopsMap["values"] = values
	}
	return iopsMap
}
func dataSourceShareProfileCapacityToMap(capacity vpcv1.ShareProfileCapacity) (capacityMap map[string]interface{}) {
	capacityMap = map[string]interface{}{}
	// if capacity.Default != nil {
	// 	capacityMap["default"] = int(*capacity.Default)
	// }
	if capacity.Max != nil {
		capacityMap["max"] = capacity.Max
	}
	if capacity.Min != nil {
		capacityMap["min"] = capacity.Min
	}
	if capacity.Step != nil {
		capacityMap["step"] = capacity.Step
	}
	if capacity.Type != nil {
		capacityMap["type"] = capacity.Type
	}
	if capacity.Value != nil {
		capacityMap["value"] = int(*capacity.Value)
	}
	values := []int{}
	if len(capacity.Values) > 0 {
		for _, value := range capacity.Values {
			values = append(values, int(value))
		}
		capacityMap["values"] = values
	}
	return capacityMap
}
func dataSourceShareProfileBandwidthToMap(bandwidth vpcv1.ShareProfileBandwidth) (bandwidthMap map[string]interface{}) {
	bandwidthMap = map[string]interface{}{}
	if bandwidth.Default != nil {
		bandwidthMap["default"] = int(*bandwidth.Default)
	}

	if bandwidth.Max != nil {
		bandwidthMap["max"] = *bandwidth.Max
	}
	if bandwidth.Min != nil {
		bandwidthMap["min"] = *bandwidth.Min
	}
	if bandwidth.Step != nil {
		bandwidthMap["step"] = *bandwidth.Step
	}
	if bandwidth.Type != nil {
		bandwidthMap["type"] = *bandwidth.Type
	}

	if bandwidth.Value != nil {
		bandwidthMap["value"] = bandwidth.Value
	}
	if bandwidth.Values != nil {
		bandwidthMap["values"] = bandwidth.Values
	}
	return bandwidthMap
}
func dataSourceShareProfileAllowedAccessProtocolToMap(allowedAccessProtocol vpcv1.ShareProfileAllowedAccessProtocols) (allowedAccessProtocolMap map[string]interface{}) {
	allowedAccessProtocolMap = map[string]interface{}{}

	if allowedAccessProtocol.Type != nil {
		allowedAccessProtocolMap["type"] = allowedAccessProtocol.Type
	}
	defaults := []string{}
	if len(allowedAccessProtocol.Default) > 0 {
		for _, value := range allowedAccessProtocol.Default {
			defaults = append(defaults, value)
		}
		allowedAccessProtocolMap["default"] = defaults
	}
	values := []string{}
	if len(allowedAccessProtocol.Values) > 0 {
		for _, value := range allowedAccessProtocol.Values {
			values = append(values, value)
		}
		allowedAccessProtocolMap["values"] = values
	}
	return allowedAccessProtocolMap
}

func dataSourceShareProfileAllowedTransitEncryptionToMap(transitEncryptionModes vpcv1.ShareProfileAllowedTransitEncryptionModes) (transitEncryptionModesMap map[string]interface{}) {
	transitEncryptionModesMap = map[string]interface{}{}

	if transitEncryptionModes.Type != nil {
		transitEncryptionModesMap["type"] = transitEncryptionModes.Type
	}
	defaults := []string{}
	if len(transitEncryptionModes.Default) > 0 {
		for _, value := range transitEncryptionModes.Default {
			defaults = append(defaults, value)
		}
		transitEncryptionModesMap["default"] = defaults
	}
	values := []string{}
	if len(transitEncryptionModes.Values) > 0 {
		for _, value := range transitEncryptionModes.Values {
			values = append(values, value)
		}
		transitEncryptionModesMap["values"] = values
	}
	return transitEncryptionModesMap
}
func dataSourceShareProfileAvailabilityModesToMap(availabilityModes vpcv1.ShareProfileAvailabilityModes) (availabilityModesMap map[string]interface{}) {
	availabilityModesMap = map[string]interface{}{}

	if availabilityModes.Type != nil {
		availabilityModesMap["type"] = availabilityModes.Type
	}
	if availabilityModes.Default != nil {
		availabilityModesMap["default"] = availabilityModes.Default
	}
	if availabilityModes.Value != nil {
		availabilityModesMap["value"] = availabilityModes.Value
	}
	values := []string{}
	if len(availabilityModes.Values) > 0 {
		for _, value := range availabilityModes.Values {
			values = append(values, value)
		}
		availabilityModesMap["values"] = values
	}
	return availabilityModesMap
}

// dataSourceIbmIsShareProfileID returns a reasonable ID for the list.
func dataSourceIbmIsShareProfileID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

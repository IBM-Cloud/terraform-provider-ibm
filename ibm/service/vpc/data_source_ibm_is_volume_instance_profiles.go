// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsVolumeInstanceProfiles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsVolumeInstanceProfilesRead,

		Schema: map[string]*schema.Schema{
			"identifier": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The volume identifier.",
			},
			"instance_profiles": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A page of instance profiles compatible with the volume.",
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
		},
	}
}

func dataSourceIBMIsVolumeInstanceProfilesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_volume_instance_profiles", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listvolumeInstanceProfilesOptions := &vpcv1.ListVolumeInstanceProfilesOptions{}

	listvolumeInstanceProfilesOptions.SetID(d.Get("identifier").(string))

	start := ""
	allrecs := []vpcv1.InstanceProfileReference{}
	for {
		if start != "" {
			listvolumeInstanceProfilesOptions.Start = &start
		}
		volumeInstanceProfile, response, err := vpcClient.ListVolumeInstanceProfiles(listvolumeInstanceProfilesOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_volume_instance_profiles", "read")
			log.Printf("[DEBUG]\n%s\n%s", tfErr.GetDebugMessage(), response)
			return tfErr.GetDiag()
		}
		start = flex.GetNext(volumeInstanceProfile.Next)
		allrecs = append(allrecs, volumeInstanceProfile.InstanceProfiles...)
		if start == "" {
			break
		}
	}

	d.SetId(dataSourceIBMIsVolumeInstanceProfilesID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allrecs {
		modelMap, err := DataSourceIBMIsVolumeInstanceProfilesInstanceProfileReferenceToMap(&modelItem)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_volume_instance_profiles", "read")
			return tfErr.GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("instance_profiles", mapSlice); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting instance_profiles %s", err), "(Data) ibm_is_volume_instance_profiles", "read")
		return tfErr.GetDiag()
	}

	return nil
}

// dataSourceIBMIsVolumeInstanceProfilesID returns a reasonable ID for the list.
func dataSourceIBMIsVolumeInstanceProfilesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMIsVolumeInstanceProfilesInstanceProfileReferenceToMap(model *vpcv1.InstanceProfileReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

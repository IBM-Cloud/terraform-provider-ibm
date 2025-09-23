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

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func DataSourceIBMIsImageInstanceProfiles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsImageInstanceProfilesRead,

		Schema: map[string]*schema.Schema{
			"identifier": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The image identifier.",
			},
			"instance_profiles": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A page of instance profiles compatible with the image.",
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

func dataSourceIBMIsImageInstanceProfilesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_image_instance_profiles", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listImageInstanceProfilesOptions := &vpcv1.ListImageInstanceProfilesOptions{}

	listImageInstanceProfilesOptions.SetID(d.Get("identifier").(string))

	start := ""
	allrecs := []vpcv1.InstanceProfileReference{}
	for {
		if start != "" {
			listImageInstanceProfilesOptions.Start = &start
		}
		imageInstanceProfile, response, err := vpcClient.ListImageInstanceProfiles(listImageInstanceProfilesOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_image_instance_profiles", "read")
			log.Printf("[DEBUG]\n%s\n%s", tfErr.GetDebugMessage(), response)
			return tfErr.GetDiag()
		}
		start = flex.GetNext(imageInstanceProfile.Next)
		allrecs = append(allrecs, imageInstanceProfile.InstanceProfiles...)
		if start == "" {
			break
		}
	}

	d.SetId(dataSourceIBMIsImageInstanceProfilesID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allrecs {
		modelMap, err := DataSourceIBMIsImageInstanceProfilesInstanceProfileReferenceToMap(&modelItem)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_image_instance_profiles", "read")
			return tfErr.GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("instance_profiles", mapSlice); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting instance_profiles %s", err), "(Data) ibm_is_image_instance_profiles", "read")
		return tfErr.GetDiag()
	}

	return nil
}

// dataSourceIBMIsImageInstanceProfilesID returns a reasonable ID for the list.
func dataSourceIBMIsImageInstanceProfilesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMIsImageInstanceProfilesInstanceProfileReferenceToMap(model *vpcv1.InstanceProfileReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

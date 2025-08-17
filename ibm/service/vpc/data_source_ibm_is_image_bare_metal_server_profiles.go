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

func DataSourceIBMIsImageBareMetalServerProfiles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsImageBareMetalServerProfilesRead,

		Schema: map[string]*schema.Schema{
			"identifier": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The image identifier.",
			},
			"bare_metal_server_profiles": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A page of bare metal server profiles compatible with the image.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this bare metal server profile.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this bare metal server profile.",
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

func dataSourceIBMIsImageBareMetalServerProfilesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_image_bare_metal_server_profiles", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listImageBareMetalServerProfilesOptions := &vpcv1.ListImageBareMetalServerProfilesOptions{}

	listImageBareMetalServerProfilesOptions.SetID(d.Get("identifier").(string))

	start := ""
	allrecs := []vpcv1.BareMetalServerProfileReference{}
	for {
		if start != "" {
			listImageBareMetalServerProfilesOptions.Start = &start
		}
		imageBareMetalServerProfiles, response, err := vpcClient.ListImageBareMetalServerProfiles(listImageBareMetalServerProfilesOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_image_bare_metal_server_profiles", "read")
			log.Printf("[DEBUG]\n%s\n%s", tfErr.GetDebugMessage(), response)
			return tfErr.GetDiag()
		}
		start = flex.GetNext(imageBareMetalServerProfiles.Next)
		allrecs = append(allrecs, imageBareMetalServerProfiles.BareMetalServerProfiles...)
		if start == "" {
			break
		}
	}

	d.SetId(dataSourceIBMIsImageBareMetalServerProfilesID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allrecs {
		modelMap, err := DataSourceIBMIsImageBareMetalServerProfilesBareMetalServerProfileReferenceToMap(&modelItem)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_image_bare_metal_server_profiles", "read")
			return tfErr.GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("bare_metal_server_profiles", mapSlice); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting bare_metal_server_profiles %s", err), "(Data) ibm_is_image_bare_metal_server_profiles", "read")
		return tfErr.GetDiag()
	}

	return nil
}

// dataSourceIBMIsImageBareMetalServerProfilesID returns a reasonable ID for the list.
func dataSourceIBMIsImageBareMetalServerProfilesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMIsImageBareMetalServerProfilesBareMetalServerProfileReferenceToMap(model *vpcv1.BareMetalServerProfileReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

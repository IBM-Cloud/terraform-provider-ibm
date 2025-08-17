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

func DataSourceIBMIsSnapshotInstanceProfiles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsSnapshotInstanceProfilesRead,

		Schema: map[string]*schema.Schema{
			"identifier": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The snapshot identifier.",
			},
			"instance_profiles": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A page of instance profiles compatible with the snapshot.",
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

func dataSourceIBMIsSnapshotInstanceProfilesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_snapshot_instance_profiles", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listSnapshotInstanceProfilesOptions := &vpcv1.ListSnapshotInstanceProfilesOptions{}

	listSnapshotInstanceProfilesOptions.SetID(d.Get("identifier").(string))

	start := ""
	allrecs := []vpcv1.InstanceProfileReference{}
	for {
		if start != "" {
			listSnapshotInstanceProfilesOptions.Start = &start
		}
		snapshotInstanceProfile, response, err := vpcClient.ListSnapshotInstanceProfiles(listSnapshotInstanceProfilesOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_snapshot_instance_profiles", "read")
			log.Printf("[DEBUG]\n%s\n%s", tfErr.GetDebugMessage(), response)
			return tfErr.GetDiag()
		}
		start = flex.GetNext(snapshotInstanceProfile.Next)
		allrecs = append(allrecs, snapshotInstanceProfile.InstanceProfiles...)
		if start == "" {
			break
		}
	}

	d.SetId(dataSourceIBMIsSnapshotInstanceProfilesID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allrecs {
		modelMap, err := DataSourceIBMIsSnapshotInstanceProfilesInstanceProfileReferenceToMap(&modelItem)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_snapshot_instance_profiles", "read")
			return tfErr.GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("instance_profiles", mapSlice); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting instance_profiles %s", err), "(Data) ibm_is_snapshot_instance_profiles", "read")
		return tfErr.GetDiag()
	}

	return nil
}

// dataSourceIBMIsSnapshotInstanceProfilesID returns a reasonable ID for the list.
func dataSourceIBMIsSnapshotInstanceProfilesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMIsSnapshotInstanceProfilesInstanceProfileReferenceToMap(model *vpcv1.InstanceProfileReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

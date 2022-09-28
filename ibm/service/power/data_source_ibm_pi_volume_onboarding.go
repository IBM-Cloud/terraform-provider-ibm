// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	//"fmt"

	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

func DataSourceIBMPIVolumeOnboarding() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIVolumeOnboardingReads,
		Schema: map[string]*schema.Schema{
			PIVolumeOnboardingID: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Volume onboarding ID",
				ValidateFunc: validation.NoZeroValues,
			},
			helpers.PICloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of consistency group at storage controller level",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Indicates the minimum period in seconds between multiple cycles",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Indicates the type of cycling mode used",
			},
			"input_volumes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Number of volumes in volume group",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"progress": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Indicates whether master/aux volume is playing the primary role",
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of remote-copy relationship names in a volume group",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"onboarded_volumes": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Number of volumes in volume group",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"volume_onboarding_failures": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of remote-copy relationship names in a volume group",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"failure_message": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Indicates whether master/aux volume is playing the primary role",
									},
									"volumes": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Number of volumes in volume group",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								}},
						},
					},
				},
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Indicates the relationship state",
			},
		},
	}
}

func dataSourceIBMPIVolumeOnboardingReads(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(helpers.PICloudInstanceId).(string)
	volOnboardClient := instance.NewIBMPIVolumeOnboardingClient(ctx, sess, cloudInstanceID)
	volOnboarding, err := volOnboardClient.Get(d.Get(PIVolumeOnboardingID).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*volOnboarding.ID)
	d.Set("creation_timestamp", volOnboarding.CreationTimestamp.String())
	d.Set("description", volOnboarding.Description)
	d.Set("input_volumes", volOnboarding.InputVolumes)
	d.Set("progress", volOnboarding.Progress)
	d.Set("status", volOnboarding.Status)
	d.Set("results", flattenVolumeOnboardingResults(volOnboarding.Results))

	return nil
}

func flattenVolumeOnboardingResults(res *models.VolumeOnboardingResults) []map[string]interface{} {
	log.Printf("Calling the flattenVolumeOnboardingResults")
	result := make([]map[string]interface{}, 0)
	result = append(result, map[string]interface{}{
		"onboarded_volumes":          res.OnboardedVolumes,
		"volume_onboarding_failures": flattenVolumeOnboardingFailures(res.VolumeOnboardingFailures),
	})

	return result
}

func flattenVolumeOnboardingFailures(list []*models.VolumeOnboardingFailure) (networks []map[string]interface{}) {
	if list != nil {
		result := make([]map[string]interface{}, len(list))
		for i, data := range list {
			l := map[string]interface{}{
				"failure_message": data.FailureMessage,
				"volumes":         data.Volumes,
			}
			result[i] = l
		}
		return result
	}
	return
}

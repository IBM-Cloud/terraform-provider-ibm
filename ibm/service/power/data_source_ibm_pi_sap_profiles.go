// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"log"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

// Attributes and Arguments defined in data_source_ibm_pi_sap_profile.go
func DataSourceIBMPISAPProfiles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPISAPProfilesRead,
		Schema: map[string]*schema.Schema{
			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			// Computed Attributes
			SapProfiles: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						SapProfileCertified: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Has certification been performed on profile",
						},
						SapProfileCores: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Amount of cores",
						},
						SapProfileMemory: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Amount of memory (in GB)",
						},
						SapProfilesID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "SAP Profile ID",
						},
						SapProfileType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of profile",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPISAPProfilesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)

	client := instance.NewIBMPISAPInstanceClient(ctx, sess, cloudInstanceID)
	sapProfiles, err := client.GetAllSAPProfiles(cloudInstanceID)
	if err != nil {
		log.Printf("[DEBUG] get all sap profiles failed %v", err)
		return diag.FromErr(err)
	}

	result := make([]map[string]interface{}, 0, len(sapProfiles.Profiles))
	for _, sapProfile := range sapProfiles.Profiles {
		profile := map[string]interface{}{
			SapProfileCertified: *sapProfile.Certified,
			SapProfileCores:     *sapProfile.Cores,
			SapProfileMemory:    *sapProfile.Memory,
			SapProfilesID:       *sapProfile.ProfileID,
			SapProfileType:      *sapProfile.Type,
		}
		result = append(result, profile)
	}

	var genID, _ = uuid.GenerateUUID()
	d.SetId(genID)
	d.Set(SapProfiles, result)

	return nil
}

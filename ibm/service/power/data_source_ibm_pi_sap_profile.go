// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

const (
	// Arguments
	PISapProfileID = "pi_sap_profile_id"

	// Attributes
	SapProfilesID       = "profile_id"
	SapProfiles         = "profiles"
	SapProfileCertified = "certified"
	SapProfileCores     = "cores"
	SapProfileMemory    = "memory"
	SapProfileType      = "type"
)

func DataSourceIBMPISAPProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPISAPProfileRead,
		Schema: map[string]*schema.Schema{
			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			PISapProfileID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "SAP Profile ID",
			},
			// Computed Attributes
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
			SapProfileType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of profile",
			},
		},
	}
}

func dataSourceIBMPISAPProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)
	profileID := d.Get(PISapProfileID).(string)

	client := instance.NewIBMPISAPInstanceClient(ctx, sess, cloudInstanceID)
	sapProfile, err := client.GetSAPProfile(profileID)
	if err != nil {
		log.Printf("[DEBUG] get sap profile failed %v", err)
		return diag.FromErr(err)
	}

	d.SetId(*sapProfile.ProfileID)
	d.Set(SapProfileCertified, *sapProfile.Certified)
	d.Set(SapProfileCores, *sapProfile.Cores)
	d.Set(SapProfileMemory, *sapProfile.Memory)
	d.Set(SapProfileType, *sapProfile.Type)

	return nil
}

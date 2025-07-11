// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceIBMPISoftwareTiers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPISoftwareTiersRead,
		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},

			// Attributes
			Attr_SupportedSoftwareTiers: {
				Computed:    true,
				Description: "List of supported software tiers (IBMi licensing).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_SupportedSystems: {
							Computed:    true,
							Description: "List of supported systems.",
							Elem:        schema.TypeString,
							Type:        schema.TypeList,
						},
						Attr_SoftwareTier: {
							Computed:    true,
							Description: "Software tier.",
							Type:        schema.TypeString,
						},
					},
				},
				Type: schema.TypeList,
			},
		},
	}
}

func dataSourceIBMPISoftwareTiersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)

	client := instance.NewIBMPIVSNClient(ctx, sess, cloudInstanceID)
	tiers, err := client.GetAllSoftwareTiers()
	if err != nil {
		return diag.Errorf("error on GET of virtual serial number software tiers: %v", err)
	}

	var genID, _ = uuid.GenerateUUID()
	d.SetId(genID)
	d.Set(Attr_SupportedSoftwareTiers, flattenSoftwareTiers(tiers))

	return nil
}

func flattenSoftwareTiers(tiers models.SupportedSoftwareTierList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, tier := range tiers {
		t := map[string]interface{}{
			Attr_SupportedSystems: tier.SupportedSystems,
			Attr_SoftwareTier:     tier.Tier,
		}
		result = append(result, t)
	}
	return result
}

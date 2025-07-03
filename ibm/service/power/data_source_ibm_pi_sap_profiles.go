// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"log"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceIBMPISAPProfiles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPISAPProfilesRead,
		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_FamilyFilter: {
				Description:  "SAP profile family filter.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"balanced", "compute", "memory", "sap-rise", "sap-rise-app", "small", "ultra-memory"}, false),
			},
			Arg_PrefixFilter: {
				Description:  "SAP profile prefix filter.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"bh1", "bh2", "ch1", "ch2", "mh1", "mh2", "umh", "ush1", "sh2", "sr2"}, false),
			},

			// Attributes
			Attr_Profiles: {
				Computed:    true,
				Description: "List of all the SAP Profiles.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_Certified: {
							Computed:    true,
							Description: "Has certification been performed on profile.",
							Type:        schema.TypeBool,
						},
						Attr_Cores: {
							Computed:    true,
							Description: "Amount of cores.",
							Type:        schema.TypeInt,
						},
						Attr_DefaultSystem: {
							Computed:    true,
							Description: "System to use if not provided",
							Type:        schema.TypeString,
						},
						Attr_FullSystemProfile: {
							Computed:    true,
							Description: "Requires full system for deployment.",
							Type:        schema.TypeBool,
						},
						Attr_Memory: {
							Computed:    true,
							Description: "Amount of memory (in GB).",
							Type:        schema.TypeInt,
						},
						Attr_ProfileID: {
							Computed:    true,
							Description: "SAP Profile ID.",
							Type:        schema.TypeString,
						},
						Attr_SAPS: {
							Computed:    true,
							Description: "SAP Application Performance Standard",
							Type:        schema.TypeInt,
						},
						Attr_SupportedSystems: {
							Computed:    true,
							Description: "List of supported systems.",
							Type:        schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						Attr_Type: {
							Computed:    true,
							Description: "Type of profile.",
							Type:        schema.TypeString,
						},
						Attr_WorkloadType: {
							Computed:    true,
							Description: "Workload Type.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Type: schema.TypeList,
						},
					},
				},
				Type: schema.TypeList,
			},
		},
	}
}

func dataSourceIBMPISAPProfilesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)

	client := instance.NewIBMPISAPInstanceClient(ctx, sess, cloudInstanceID)
	filters := map[string]string{}
	if v, ok := d.GetOk(Arg_FamilyFilter); ok {
		filters[Arg_FamilyFilter] = v.(string)
	}
	if v, ok := d.GetOk(Arg_PrefixFilter); ok {
		filters[Arg_PrefixFilter] = v.(string)
	}

	sapProfiles, err := client.GetAllSAPProfilesWithFilters(cloudInstanceID, filters)
	if err != nil {
		log.Printf("[DEBUG] get all sap profiles failed %v", err)
		return diag.FromErr(err)
	}

	result := make([]map[string]interface{}, 0, len(sapProfiles.Profiles))
	for _, sapProfile := range sapProfiles.Profiles {
		profile := map[string]interface{}{
			Attr_Certified:         *sapProfile.Certified,
			Attr_Cores:             *sapProfile.Cores,
			Attr_DefaultSystem:     sapProfile.DefaultSystem,
			Attr_FullSystemProfile: sapProfile.FullSystemProfile,
			Attr_Memory:            *sapProfile.Memory,
			Attr_ProfileID:         *sapProfile.ProfileID,
			Attr_SAPS:              sapProfile.Saps,
			Attr_SupportedSystems:  sapProfile.SupportedSystems,
			Attr_Type:              *sapProfile.Type,
			Attr_WorkloadType:      sapProfile.WorkloadTypes,
		}
		result = append(result, profile)
	}

	var genID, _ = uuid.GenerateUUID()
	d.SetId(genID)
	d.Set(Attr_Profiles, result)

	return nil
}

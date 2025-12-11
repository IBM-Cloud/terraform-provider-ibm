// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Datasource to list available console languages for an instance
func DataSourceIBMPIInstanceConsoleLanguages() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIInstanceConsoleLanguagesRead,
		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_InstanceID: {
				AtLeastOneOf:  []string{Arg_InstanceID, Arg_InstanceName},
				ConflictsWith: []string{Arg_InstanceName},
				Description:   "The ID of the PVM instance.",
				Optional:      true,
				Type:          schema.TypeString,
			},
			Arg_InstanceName: {
				AtLeastOneOf:  []string{Arg_InstanceID, Arg_InstanceName},
				ConflictsWith: []string{Arg_InstanceID},
				Deprecated:    "The pi_instance_name field is deprecated. Please use pi_instance_id instead",
				Description:   "The name of the PVM instance.",
				Optional:      true,
				Type:          schema.TypeString,
			},

			// Attributes
			Attr_ConsoleLanguages: {
				Computed:    true,
				Description: "List of all the Console Languages.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_Code: {
							Computed:    true,
							Description: "Language code.",
							Type:        schema.TypeString,
						},
						Attr_Language: {
							Computed:    true,
							Description: "Language description.",
							Type:        schema.TypeString,
						},
					},
				},
				Type: schema.TypeList,
			},
		},
	}
}

func dataSourceIBMPIInstanceConsoleLanguagesRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "(Data) ibm_pi_instance_console_languages", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	var instanceID string
	if v, ok := d.GetOk(Arg_InstanceID); ok {
		instanceID = v.(string)
	} else if v, ok := d.GetOk(Arg_InstanceName); ok {
		instanceID = v.(string)
	}

	client := instance.NewIBMPIInstanceClient(ctx, sess, cloudInstanceID)
	languages, err := client.GetConsoleLanguages(instanceID)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetConsoleLanguages failed: %s", err.Error()), "(Data) ibm_pi_instance_console_languages", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)

	if len(languages.ConsoleLanguages) > 0 {
		result := make([]map[string]any, 0, len(languages.ConsoleLanguages))
		for _, language := range languages.ConsoleLanguages {
			l := map[string]any{
				Attr_Code:     *language.Code,
				Attr_Language: language.Language,
			}
			result = append(result, l)
		}
		d.Set(Attr_ConsoleLanguages, result)
	}

	return nil
}

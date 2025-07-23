// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Datasource to get a virtual serial number in a power instance
func DataSourceIBMPIVirtualSerialNumber() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIVirtualSerialNumberRead,
		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_Serial: {
				Description:  "Virtual serial number.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},

			// Attributes
			Attr_Description: {
				Computed:    true,
				Description: "Description of virtual serial number.",
				Type:        schema.TypeString,
			},
			Attr_InstanceID: {
				Computed:    true,
				Description: "ID of PVM instance virtual serial number is attached to.",
				Type:        schema.TypeString,
			},
			Attr_SoftwareTier: {
				Computed:    true,
				Description: "Software tier for virtual serial number.",
				Type:        schema.TypeString,
			},
		},
	}
}

func dataSourceIBMPIVirtualSerialNumberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	client := instance.NewIBMPIVSNClient(ctx, sess, cloudInstanceID)

	vsnInput := d.Get(Arg_Serial).(string)
	virtualSerialNumberData, err := client.Get(vsnInput)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Get failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	id := *virtualSerialNumberData.Serial
	d.SetId(id)
	d.Set(Attr_Description, virtualSerialNumberData.Description)
	if virtualSerialNumberData.PvmInstanceID != nil {
		d.Set(Attr_InstanceID, virtualSerialNumberData.PvmInstanceID)
	}
	if virtualSerialNumberData.SoftwareTier != "" {
		d.Set(Attr_SoftwareTier, virtualSerialNumberData.SoftwareTier)
	}

	return nil
}

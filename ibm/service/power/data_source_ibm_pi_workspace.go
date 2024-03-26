// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DatasourceIBMPIWorkspace() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIWorkspaceRead,
		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},

			// Attributes
			Attr_WorkspaceCapabilities: {
				Computed:    true,
				Description: "Workspace Capabilities.",
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
				Type: schema.TypeMap,
			},
			Attr_WorkspaceDetails: {
				Computed:    true,
				Description: "Workspace information.",
				Type:        schema.TypeMap,
			},
			Attr_WorkspaceLocation: {
				Computed:    true,
				Description: "Workspace location.",
				Type:        schema.TypeMap,
			},
			Attr_WorkspaceName: {
				Computed:    true,
				Description: "Workspace name.",
				Type:        schema.TypeString,
			},
			Attr_WorkspaceStatus: {
				Computed:    true,
				Description: "Workspace status, active, critical, failed, provisioning.",
				Type:        schema.TypeString,
			},
			Attr_WorkspaceType: {
				Computed:    true,
				Description: "Workspace type, off-premises or on-premises.",
				Type:        schema.TypeString,
			},
		},
	}
}

func dataSourceIBMPIWorkspaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	client := instance.NewIBMPIWorkspacesClient(ctx, sess, cloudInstanceID)
	wsData, err := client.Get(cloudInstanceID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set(Attr_WorkspaceName, wsData.Name)
	d.Set(Attr_WorkspaceStatus, wsData.Status)
	d.Set(Attr_WorkspaceType, wsData.Type)
	d.Set(Attr_WorkspaceCapabilities, wsData.Capabilities)
	wsdetails := map[string]interface{}{
		Attr_CreationDate: wsData.Details.CreationDate.String(),
		Attr_CRN:          *wsData.Details.Crn,
	}
	d.Set(Attr_WorkspaceDetails, flex.Flatten(wsdetails))
	wslocation := map[string]interface{}{
		Attr_Region: *wsData.Location.Region,
		Attr_Type:   wsData.Location.Type,
		Attr_URL:    wsData.Location.URL,
	}
	d.Set(Attr_WorkspaceLocation, flex.Flatten(wslocation))
	d.SetId(*wsData.ID)
	return nil
}

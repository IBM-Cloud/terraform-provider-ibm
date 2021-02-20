/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSchematicsState() *schema.Resource {
	return &schema.Resource{
		Read: resourceIBMSchematicsStateRead,

		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of workspace",
			},
			"template_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of template",
			},
			"state_store": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state_store_json": {
				Type:     schema.TypeString,
				Computed: true,
			},
			ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this workspace",
			},
		},
	}

}

func resourceIBMSchematicsStateRead(d *schema.ResourceData, meta interface{}) error {
	scClient, err := meta.(ClientSession).SchematicsAPI()
	if err != nil {
		return err
	}

	wrkAPI := scClient.Workspaces()
	workspaceID := d.Get("workspace_id").(string)
	templateID := d.Get("template_id").(string)

	stateStore, err := wrkAPI.GetStateStore(workspaceID, templateID)
	if err != nil {
		return fmt.Errorf("Error retreiving statestore: %s", err)
	}
	statestr := fmt.Sprintf("%v", stateStore)
	d.SetId(fmt.Sprintf("%s/%s", workspaceID, templateID))
	d.Set("state_store", statestr)

	stateByte, err := json.MarshalIndent(stateStore, "", "")
	if err != nil {
		return err
	}

	stateStoreJson := string(stateByte[:])
	d.Set("state_store_json", stateStoreJson)

	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/schematics")

	return nil

}

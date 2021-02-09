/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

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

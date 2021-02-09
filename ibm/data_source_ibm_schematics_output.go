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

func dataSourceSchematicsOut() *schema.Resource {
	return &schema.Resource{
		Read: resourceIBMSchematicsOutRead,

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
			"type": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"output_values": {
				Type:     schema.TypeMap,
				Computed: true,
			},
			"output_json": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The json output in string",
			},
			ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this Workspace",
			},
		},
	}

}

func resourceIBMSchematicsOutRead(d *schema.ResourceData, meta interface{}) error {
	scClient, err := meta.(ClientSession).SchematicsAPI()
	if err != nil {
		return err
	}

	wrkAPI := scClient.Workspaces()
	workspaceID := d.Get("workspace_id").(string)
	templateID := d.Get("template_id").(string)

	out, err := wrkAPI.GetOutputValues(workspaceID)
	if err != nil {
		return fmt.Errorf("Error while retreiving outputs of workspace: %s", err)
	}

	var outputJSON string
	items := make(map[string]interface{})
	found := false
	for _, feilds := range out {
		if feilds.TemplateID == templateID {
			output := feilds.Output
			found = true
			outputByte, err := json.MarshalIndent(output, "", "")
			if err != nil {
				return err
			}

			outputJSON = string(outputByte[:])
			// items := map[string]interface{}

			for _, value := range output {
				for key, val := range value {
					val2 := val.Value
					items[key] = val2

				}
			}
		}
	}
	if !(found) {
		return fmt.Errorf("Error while fetching template id in workspace: %s", workspaceID)
	}

	d.Set("output_json", outputJSON)
	d.SetId(fmt.Sprintf("%s/%s", workspaceID, templateID))
	d.Set("output_values", Flatten(items))

	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/schematics")

	return nil

}

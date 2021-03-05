/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ibm

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func dataSourceIBMSchematicsState() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMSchematicsStateRead,

		Schema: map[string]*schema.Schema{
			"workspace_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The workspace ID for the workspace that you want to query.  You can run the GET /workspaces call if you need to look up the  workspace IDs in your IBM Cloud account.",
			},
			"template_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Template ID for which you want to get the values.  Use the GET /workspaces to look up the workspace IDs  or template IDs in your IBM Cloud account.",
			},
			"version": &schema.Schema{
				Type:        schema.TypeFloat,
				Computed:    true,
			},
			"terraform_version": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
			},
			"serial": &schema.Schema{
				Type:        schema.TypeFloat,
				Computed:    true,
			},
			"lineage": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
			},
			"modules": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
				},
			},
		},
	}
}

func dataSourceIBMSchematicsStateRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getWorkspaceTemplateStateOptions := &schematicsv1.GetWorkspaceTemplateStateOptions{}

	getWorkspaceTemplateStateOptions.SetWID(d.Get("workspace_id").(string))
	getWorkspaceTemplateStateOptions.SetTID(d.Get("template_id").(string))

	templateStateStore, response, err := schematicsClient.GetWorkspaceTemplateStateWithContext(context, getWorkspaceTemplateStateOptions)
	if err != nil {
		log.Printf("[DEBUG] GetWorkspaceTemplateStateWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId(dataSourceIBMSchematicsStateID(d))
	if err = d.Set("version", templateStateStore.Version); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting version: %s", err))
	}
	if err = d.Set("terraform_version", templateStateStore.TerraformVersion); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting terraform_version: %s", err))
	}
	if err = d.Set("serial", templateStateStore.Serial); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting serial: %s", err))
	}
	if err = d.Set("lineage", templateStateStore.Lineage); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting lineage: %s", err))
	}
	if err = d.Set("modules", templateStateStore.Modules); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting modules: %s", err))
	}

	return nil
}

// dataSourceIBMSchematicsStateID returns a reasonable ID for the list.
func dataSourceIBMSchematicsStateID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

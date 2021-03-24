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
	"fmt"
	"log"
	"time"

	"github.com/IBM/schematics-go-sdk/schematicsv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMSchematicsOutput() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMSchematicsOutputRead,

		Schema: map[string]*schema.Schema{
			"workspace_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the workspace for which you want to retrieve output values. To find the workspace ID, use the `GET /workspaces` API.",
			},
			"output_values": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "OutputValues -.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"folder": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Output variable name.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Output variable id.",
						},
						"output_values": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of Output values.",
							Elem: &schema.Schema{
								Type: schema.TypeMap,
							},
						},
						"value_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Output variable type.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMSchematicsOutputRead(d *schema.ResourceData, meta interface{}) error {
	schematicsClient, err := meta.(ClientSession).SchematicsV1()
	if err != nil {
		return err
	}

	getWorkspaceOutputsOptions := &schematicsv1.GetWorkspaceOutputsOptions{}

	getWorkspaceOutputsOptions.SetWID(d.Get("workspace_id").(string))

	outputValuesList, response, err := schematicsClient.GetWorkspaceOutputs(getWorkspaceOutputsOptions)
	if err != nil {
		log.Printf("[DEBUG] GetWorkspaceOutputs failed %s\n%s", err, response)
		return err
	}

	d.SetId(dataSourceIBMSchematicsOutputID(d))

	if outputValuesList != nil {
		err = d.Set("output_values", dataSourceOutputValuesListFlattenOutputValues(outputValuesList))
		if err != nil {
			return fmt.Errorf("Error setting output_values %s", err)
		}
	}

	return nil
}

// dataSourceIBMSchematicsOutputID returns a reasonable ID for the list.
func dataSourceIBMSchematicsOutputID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceOutputValuesListFlattenOutputValues(result []schematicsv1.OutputValuesItem) (outputValues []interface{}) {
	for _, outputValuesItem := range result {
		outputValues = append(outputValues, dataSourceOutputValuesListOutputValuesToMap(outputValuesItem))
	}

	return outputValues
}

func dataSourceOutputValuesListOutputValuesToMap(outputValuesItem schematicsv1.OutputValuesItem) (outputValuesMap map[string]interface{}) {
	outputValuesMap = map[string]interface{}{}

	if outputValuesItem.Folder != nil {
		outputValuesMap["folder"] = outputValuesItem.Folder
	}
	if outputValuesItem.ID != nil {
		outputValuesMap["id"] = outputValuesItem.ID
	}

	m := []Map{}

	for _, outputValues := range outputValuesItem.OutputValues {
		m = append(m, Flatten(outputValues.(map[string]interface{})))
	}

	if outputValuesItem.OutputValues != nil {
		outputValuesMap["output_values"] = m
	}
	if outputValuesItem.ValueType != nil {
		outputValuesMap["value_type"] = outputValuesItem.ValueType
	}

	return outputValuesMap
}

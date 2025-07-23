// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager

import (
	"context"
	"fmt"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
)

func DataSourceIbmSmCustomCredentialsConfiguration() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmCustomCredentialsConfigurationRead,

		Schema: map[string]*schema.Schema{
			"api_key_ref": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the IAM credentials secret that is used for setting up the custom credentials secret configuration.",
			},
			"code_engine": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The parameters required to configure Code Engine.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"job_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the Code Engine Job.",
						},
						"project_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Code Engine project.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region of the Code Engine project.",
						},
					},
				},
			},
			"code_engine_key_ref": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IAM API key used by the credentials system to access this Secrets Manager instance.",
			},
			"created_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier that is associated with the entity that created the configuration.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date when the configuration was created. The date format follows RFC 3339.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the configuration.",
			},
			"schema": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The schema that defines by the Code Engine job to be used as input and output formats for this custom credentials configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"credentials": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Custom credentials configuration schema credentials format.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the credential.",
									},
									"format": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The format of the credential, for example 'required:true, type:string'",
									},
								},
							},
						},
						"parameters": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Custom credentials configuration schema parameters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the parameter.",
									},
									"format": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The format of the parameter, for example 'required:true, type:string', 'type:int, required:false', 'type:enum[val1|val2|val3], required:true', 'required:true, type:boolean'",
									},
									"env_variable_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the environment variable associated with the configuration schema parameter.",
									},
								},
							},
						},
					},
				},
			},
			"task_timeout": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the maximum allowed time for a Code Engine task to be completed. After this time elapses, the task state will changed to failed.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date when the configuration was recently modified. The date format follows RFC 3339.",
			},
		},
	}
}

func dataSourceIbmSmCustomCredentialsConfigurationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, endpointsFile, err := getSecretsManagerSession(meta.(conns.ClientSession))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, "", fmt.Sprintf("(Data) %s", CustomCredentialsConfigResourceName), "read")
		return tfErr.GetDiag()
	}

	region := getRegion(secretsManagerClient, d)
	instanceId := d.Get("instance_id").(string)
	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, instanceId, region, getEndpointType(secretsManagerClient, d), endpointsFile)

	getConfigurationOptions := &secretsmanagerv2.GetConfigurationOptions{}

	getConfigurationOptions.SetName(d.Get("name").(string))

	configIntf, response, err := secretsManagerClient.GetConfigurationWithContext(context, getConfigurationOptions)
	if err != nil {
		log.Printf("[DEBUG] GetConfigurationWithContext failed %s\n%s", err, response)
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetConfigurationWithContext failed %s\n%s", err, response), fmt.Sprintf("(Data) %s", CustomCredentialsConfigResourceName), "read")
		return tfErr.GetDiag()
	}
	config, ok := configIntf.(*secretsmanagerv2.CustomCredentialsConfiguration)
	if !ok {
		tfErr := flex.TerraformErrorf(nil, fmt.Sprintf("Wrong configuration type: The provided configuration is not a Custom Credentials configuration."), fmt.Sprintf("(Data) %s", CustomCredentialsConfigResourceName), "read")
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", region, instanceId, *getConfigurationOptions.Name))

	if err = d.Set("region", region); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting region"), fmt.Sprintf("(Data) %s", CustomCredentialsConfigResourceName), "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("created_by", config.CreatedBy); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting created_by"), fmt.Sprintf("(Data) %s", CustomCredentialsConfigResourceName), "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("created_at", DateTimeToRFC3339(config.CreatedAt)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting created_at"), fmt.Sprintf("(Data) %s", CustomCredentialsConfigResourceName), "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("updated_at", DateTimeToRFC3339(config.UpdatedAt)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting updated_at"), fmt.Sprintf("(Data) %s", CustomCredentialsConfigResourceName), "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("api_key_ref", config.ApiKeyRef); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting api_key_ref"), fmt.Sprintf("(Data) %s", CustomCredentialsConfigResourceName), "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("task_timeout", config.TaskTimeout); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting task_timeout"), fmt.Sprintf("(Data) %s", CustomCredentialsConfigResourceName), "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("code_engine_key_ref", config.CodeEngineKeyRef); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting code_engine_key_ref"), fmt.Sprintf("(Data) %s", CustomCredentialsConfigResourceName), "read")
		return tfErr.GetDiag()
	}

	codeEngine := customCredentialsConfigurationCodeEngineToMap(config.CodeEngine)
	if err = d.Set("code_engine", codeEngine); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting code_engine"), fmt.Sprintf("(Data) %s", CustomCredentialsSecretResourceName), "read")
		return tfErr.GetDiag()
	}

	schema := customCredentialsConfigurationSchemaToMap(config.Schema)
	if err = d.Set("schema", schema); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting schema"), fmt.Sprintf("(Data) %s", CustomCredentialsSecretResourceName), "read")
		return tfErr.GetDiag()
	}
	return nil
}

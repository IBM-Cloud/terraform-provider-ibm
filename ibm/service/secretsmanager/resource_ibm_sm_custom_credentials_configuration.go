// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager

import (
	"context"
	"fmt"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
)

func ResourceIbmSmCustomCredentialsConfiguration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSmCustomCredentialsConfigurationCreate,
		ReadContext:   resourceIbmSmCustomCredentialsConfigurationRead,
		UpdateContext: resourceIbmSmCustomCredentialsConfigurationUpdate,
		DeleteContext: resourceIbmSmCustomCredentialsConfigurationDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"api_key_ref": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The ID of the IAM credentials secret that is used for setting up the custom credentials secret configuration.",
			},
			"code_engine": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				ForceNew:    true,
				Description: "The parameters required to configure Code Engine.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"job_name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the Code Engine Job.",
						},
						"project_id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The ID of the Code Engine project.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
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
				ForceNew:    true,
				Description: "A human-readable unique name to assign to your configuration.To protect your privacy, do not use personal data, such as your name or location, as an name for your secret.",
			},
			"schema": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The schema that defines the format of the input and output parameters (the credentials) of the Code Engine job.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"credentials": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The schema of the credentials.",
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
							Description: "The schema of the input parameters.",
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
				Optional:    true,
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

func resourceIbmSmCustomCredentialsConfigurationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, endpointsFile, err := getSecretsManagerSession(meta.(conns.ClientSession))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, "", CustomCredentialsConfigResourceName, "create")
		return tfErr.GetDiag()
	}

	region := getRegion(secretsManagerClient, d)
	instanceId := d.Get("instance_id").(string)
	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, instanceId, region, getEndpointType(secretsManagerClient, d), endpointsFile)

	createConfigurationOptions := &secretsmanagerv2.CreateConfigurationOptions{}

	configurationPrototypeModel, err := resourceIbmSmCustomCredentialsConfigurationMapToConfigurationPrototype(d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, "", CustomCredentialsConfigResourceName, "create")
		return tfErr.GetDiag()
	}
	createConfigurationOptions.SetConfigurationPrototype(configurationPrototypeModel)

	configurationIntf, response, err := secretsManagerClient.CreateConfigurationWithContext(context, createConfigurationOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateConfigurationWithContext failed %s\n%s", err, response)
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateConfigurationWithContext failed %s\n%s", err, response), CustomCredentialsConfigResourceName, "create")
		return tfErr.GetDiag()
	}
	configuration := configurationIntf.(*secretsmanagerv2.CustomCredentialsConfiguration)

	d.SetId(fmt.Sprintf("%s/%s/%s", region, instanceId, *configuration.Name))

	return resourceIbmSmCustomCredentialsConfigurationRead(context, d, meta)
}

func resourceIbmSmCustomCredentialsConfigurationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, endpointsFile, err := getSecretsManagerSession(meta.(conns.ClientSession))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, "", CustomCredentialsConfigResourceName, "read")
		return tfErr.GetDiag()
	}

	id := strings.Split(d.Id(), "/")
	if len(id) != 3 {
		tfErr := flex.TerraformErrorf(nil, "Wrong format of resource ID. To import Custom credentials configuration use the format `<region>/<instance_id>/<name>`", CustomCredentialsConfigResourceName, "read")
		return tfErr.GetDiag()
	}
	region := id[0]
	instanceId := id[1]
	configName := id[2]
	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, instanceId, region, getEndpointType(secretsManagerClient, d), endpointsFile)

	getConfigurationOptions := &secretsmanagerv2.GetConfigurationOptions{}

	getConfigurationOptions.SetName(configName)

	configurationIntf, response, err := secretsManagerClient.GetConfigurationWithContext(context, getConfigurationOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetConfigurationWithContext failed %s\n%s", err, response)
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetConfigurationWithContext failed %s\n%s", err, response), CustomCredentialsConfigResourceName, "read")
		return tfErr.GetDiag()
	}
	configuration := configurationIntf.(*secretsmanagerv2.CustomCredentialsConfiguration)

	if err = d.Set("instance_id", instanceId); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting instance_id"), CustomCredentialsConfigResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("region", region); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting region"), CustomCredentialsConfigResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("name", configuration.Name); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting name"), CustomCredentialsConfigResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("created_by", configuration.CreatedBy); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting created_by"), CustomCredentialsConfigResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("created_at", DateTimeToRFC3339(configuration.CreatedAt)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting created_at"), CustomCredentialsConfigResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("updated_at", DateTimeToRFC3339(configuration.UpdatedAt)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting updated_at"), CustomCredentialsConfigResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("api_key_ref", configuration.ApiKeyRef); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting api_key_ref"), CustomCredentialsConfigResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("code_engine_key_ref", configuration.CodeEngineKeyRef); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting code_engine_key_ref"), CustomCredentialsConfigResourceName, "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("task_timeout", configuration.TaskTimeout); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting task_timeout"), fmt.Sprintf("(Data) %s", CustomCredentialsConfigResourceName), "read")
		return tfErr.GetDiag()
	}
	codeEngine := customCredentialsConfigurationCodeEngineToMap(configuration.CodeEngine)
	if err = d.Set("code_engine", codeEngine); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting code_engine"), fmt.Sprintf("(Data) %s", CustomCredentialsSecretResourceName), "read")
		return tfErr.GetDiag()
	}
	schema := customCredentialsConfigurationSchemaToMap(configuration.Schema)
	if err = d.Set("schema", schema); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting schema"), fmt.Sprintf("(Data) %s", CustomCredentialsSecretResourceName), "read")
		return tfErr.GetDiag()
	}

	return nil
}

func resourceIbmSmCustomCredentialsConfigurationUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, endpointsFile, err := getSecretsManagerSession(meta.(conns.ClientSession))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, "", CustomCredentialsConfigResourceName, "update")
		return tfErr.GetDiag()
	}

	id := strings.Split(d.Id(), "/")
	region := id[0]
	instanceId := id[1]
	configName := id[2]
	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, instanceId, region, getEndpointType(secretsManagerClient, d), endpointsFile)

	updateConfigurationOptions := &secretsmanagerv2.UpdateConfigurationOptions{}

	updateConfigurationOptions.SetName(configName)
	updateConfigurationOptions.SetXSmAcceptConfigurationType("custom_credentials_configuration")

	hasChange := false

	patchVals := &secretsmanagerv2.ConfigurationPatch{}

	if d.HasChange("task_timeout") {
		patchVals.TaskTimeout = core.StringPtr(d.Get("task_timeout").(string))
		hasChange = true
	}

	if hasChange {
		updateConfigurationOptions.ConfigurationPatch, _ = patchVals.AsPatch()
		_, response, err := secretsManagerClient.UpdateConfigurationWithContext(context, updateConfigurationOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateConfigurationWithContext failed %s\n%s", err, response)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateConfigurationWithContext failed %s\n%s", err, response), CustomCredentialsConfigResourceName, "update")
			return tfErr.GetDiag()
		}
	}

	return resourceIbmSmCustomCredentialsConfigurationRead(context, d, meta)
}

func resourceIbmSmCustomCredentialsConfigurationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, endpointsFile, err := getSecretsManagerSession(meta.(conns.ClientSession))
	if err != nil {
		return diag.FromErr(err)
	}

	id := strings.Split(d.Id(), "/")
	region := id[0]
	instanceId := id[1]
	configName := id[2]
	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, instanceId, region, getEndpointType(secretsManagerClient, d), endpointsFile)

	deleteConfigurationOptions := &secretsmanagerv2.DeleteConfigurationOptions{}

	deleteConfigurationOptions.SetName(configName)

	response, err := secretsManagerClient.DeleteConfigurationWithContext(context, deleteConfigurationOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteConfigurationWithContext failed %s\n%s", err, response)
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteConfigurationWithContext failed %s\n%s", err, response), CustomCredentialsConfigResourceName, "delete")
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func resourceIbmSmCustomCredentialsConfigurationMapToConfigurationPrototype(d *schema.ResourceData) (secretsmanagerv2.ConfigurationPrototypeIntf, error) {
	model := &secretsmanagerv2.CustomCredentialsConfigurationPrototype{}

	model.ConfigType = core.StringPtr("custom_credentials_configuration")

	if _, ok := d.GetOk("name"); ok {
		model.Name = core.StringPtr(d.Get("name").(string))
	}
	if _, ok := d.GetOk("api_key_ref"); ok {
		model.ApiKeyRef = core.StringPtr(d.Get("api_key_ref").(string))
	}
	if _, ok := d.GetOk("task_timeout"); ok {
		model.TaskTimeout = core.StringPtr(d.Get("task_timeout").(string))
	}
	if _, ok := d.GetOk("code_engine"); ok {
		codeEngine := &secretsmanagerv2.CustomCredentialsConfigurationCodeEngine{}
		codeEngineMap := d.Get("code_engine").([]interface{})[0].(map[string]interface{})
		if codeEngineMap["job_name"] != nil {
			jobName := codeEngineMap["job_name"].(string)
			codeEngine.JobName = &jobName
		}
		if codeEngineMap["project_id"] != nil {
			projectId := codeEngineMap["project_id"].(string)
			codeEngine.ProjectID = &projectId
		}
		if codeEngineMap["region"] != nil {
			region := codeEngineMap["region"].(string)
			codeEngine.Region = &region
		}
		model.CodeEngine = codeEngine
	}
	return model, nil
}

func customCredentialsConfigurationSchemaToMap(schema *secretsmanagerv2.CustomCredentialsConfigurationSchema) []map[string]interface{} {
	schemaMapList := []map[string]interface{}{}
	schemaMap := map[string]interface{}{}
	if schema != nil {
		parameterMaps := []map[string]interface{}{}
		if schema.Parameters != nil {
			for _, parameter := range schema.Parameters {
				parameterMap := customCredentialsConfigurationSchemaParameterToMap(parameter)
				parameterMaps = append(parameterMaps, parameterMap)
			}
		}
		schemaMap["parameters"] = parameterMaps

		credentialMaps := []map[string]interface{}{}
		if schema.Credentials != nil {
			for _, credential := range schema.Credentials {
				credentialMap := customCredentialsConfigurationSchemaCredentialToMap(credential)
				credentialMaps = append(credentialMaps, credentialMap)
			}
		}
		schemaMap["credentials"] = credentialMaps
		schemaMapList = append(schemaMapList, schemaMap)
	}
	return schemaMapList
}

func customCredentialsConfigurationSchemaParameterToMap(parameter secretsmanagerv2.CustomCredentialsConfigurationSchemaParameter) map[string]interface{} {
	parameterMap := make(map[string]interface{})
	if parameter.Name != nil {
		parameterMap["name"] = parameter.Name
	}
	if parameter.EnvVariableName != nil {
		parameterMap["env_variable_name"] = parameter.EnvVariableName
	}
	if parameter.Format != nil {
		parameterMap["format"] = parameter.Format
	}
	return parameterMap
}

func customCredentialsConfigurationSchemaCredentialToMap(credential secretsmanagerv2.CustomCredentialsConfigurationSchemaCredentials) map[string]interface{} {
	credentialMap := make(map[string]interface{})
	if credential.Name != nil {
		credentialMap["name"] = credential.Name
	}
	if credential.Format != nil {
		credentialMap["format"] = credential.Format
	}
	return credentialMap
}

func customCredentialsConfigurationCodeEngineToMap(codeEngine *secretsmanagerv2.CustomCredentialsConfigurationCodeEngine) []map[string]interface{} {
	codeEngineMapList := []map[string]interface{}{}
	if codeEngine != nil {
		codeEngineMap := map[string]interface{}{}
		if codeEngine.JobName != nil {
			codeEngineMap["job_name"] = codeEngine.JobName
		}
		if codeEngine.ProjectID != nil {
			codeEngineMap["project_id"] = codeEngine.ProjectID
		}
		if codeEngine.Region != nil {
			codeEngineMap["region"] = codeEngine.Region
		}
		codeEngineMapList = append(codeEngineMapList, codeEngineMap)
	}
	return codeEngineMapList
}

// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager

import (
	"context"
	"errors"
	"fmt"
	"github.com/Mavrickk3/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func ResourceIbmSmCustomCredentialsSecret() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSmCustomCredentialsSecretCreate,
		ReadContext:   resourceIbmSmCustomCredentialsSecretRead,
		UpdateContext: resourceIbmSmCustomCredentialsSecretUpdate,
		DeleteContext: resourceIbmSmCustomCredentialsSecretDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"configuration": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Custom Credentials configuration.",
			},
			"credentials_content": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Sensitive:   true,
				Description: "The credentials that were generated for this secret.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"boolean_values": &schema.Schema{
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Credentials that have boolean values.",
							Elem: &schema.Schema{
								Type: schema.TypeBool,
							}},
						"integer_values": &schema.Schema{
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Credentials that have integer values.",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"string_values": &schema.Schema{
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Credentials that have string values.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"custom_metadata": &schema.Schema{
				Type:        schema.TypeMap,
				Optional:    true,
				Computed:    true,
				Description: "The secret metadata that a user can customize.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.",
			},
			"expiration_date": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date a secret is expired. The date format follows RFC 3339.",
			},
			"labels": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Labels that you can use to search for secrets in your instance.Up to 30 labels can be created.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "A human-readable name to assign to your secret.To protect your privacy, do not use personal data, such as your name or location, as a name for your secret.",
			},
			"parameters": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "The parameters that are passed to the Custom Credentials engine.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"boolean_values": &schema.Schema{
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Pararmeters that have boolean values.",
							Elem: &schema.Schema{
								Type: schema.TypeBool,
							}},
						"integer_values": &schema.Schema{
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Pararmeters that have integer values.",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"string_values": &schema.Schema{
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Pararmeters that have string values.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				}},
			"retrieved_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date when the data of the secret was last retrieved. The date format follows RFC 3339. Epoch date if there is no record of secret data retrieval.",
			},
			"secret_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "A UUID identifier, or `default` secret group.",
			},
			"secret_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A UUID identifier.",
			},
			"ttl": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringIsIntWithMinimum(86400),
				Description:  "The time-to-live or lease duration (in seconds) to assign to generated credentials. Minimum duration is 86400 seconds (one day).",
			},
			"version_custom_metadata": &schema.Schema{
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The secret version metadata that a user can customize.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"rotation": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "Determines whether Secrets Manager rotates your secrets automatically.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_rotate": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Determines whether Secrets Manager rotates your secret automatically.Default is `false`. If `auto_rotate` is set to `true` the service rotates your secret based on the defined interval.",
						},
						"interval": &schema.Schema{
							Type:             schema.TypeInt,
							Optional:         true,
							Computed:         true,
							Description:      "The length of the secret rotation time interval.",
							DiffSuppressFunc: rotationAttributesDiffSuppress,
						},
						"unit": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							Description:      "The units for the secret rotation time interval.",
							DiffSuppressFunc: rotationAttributesDiffSuppress,
						},
					},
				},
			},
			"created_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier that is associated with the entity that created the secret.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date when a resource was created. The date format follows RFC 3339.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A CRN that uniquely identifies an IBM Cloud resource.",
			},
			"downloaded": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether the secret data that is associated with a secret version was retrieved in a call to the service API.",
			},
			"locks_total": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of locks of the secret.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The secret state that is based on NIST SP 800-57. States are integers and correspond to the `Pre-activation = 0`, `Active = 1`,  `Suspended = 2`, `Deactivated = 3`, and `Destroyed = 5` values.",
			},
			"state_description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A text representation of the secret state.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date when a resource was recently modified. The date format follows RFC 3339.",
			},
			"versions_total": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of versions of the secret.",
			},
			"next_rotation_date": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date that the secret is scheduled for automatic rotation.The service automatically creates a new version of the secret on its next rotation date. This field exists only for secrets that have an existing rotation policy.",
			},
		},
	}
}

func resourceIbmSmCustomCredentialsSecretCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, endpointsFile, err := getSecretsManagerSession(meta.(conns.ClientSession))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, "", CustomCredentialsSecretResourceName, "create")
		return tfErr.GetDiag()
	}

	region := getRegion(secretsManagerClient, d)
	instanceId := d.Get("instance_id").(string)
	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, instanceId, region, getEndpointType(secretsManagerClient, d), endpointsFile)

	createSecretOptions := &secretsmanagerv2.CreateSecretOptions{}

	secretPrototypeModel, err := resourceIbmSmCustomCredentialsSecretMapToSecretPrototype(d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, "", CustomCredentialsSecretResourceName, "create")
		return tfErr.GetDiag()
	}
	createSecretOptions.SetSecretPrototype(secretPrototypeModel)

	secretIntf, response, err := secretsManagerClient.CreateSecretWithContext(context, createSecretOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateSecretWithContext failed %s\n%s", err, response)
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateSecretWithContext failed: %s\n%s", err.Error(), response), CustomCredentialsSecretResourceName, "create")
		return tfErr.GetDiag()
	}
	secret := secretIntf.(*secretsmanagerv2.CustomCredentialsSecret)
	d.SetId(fmt.Sprintf("%s/%s/%s", region, instanceId, *secret.ID))
	d.Set("secret_id", *secret.ID)

	_, err = waitForIbmSmCustomCredentialsSecretCreate(secretsManagerClient, d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error waiting for resource IbmSmCustomCredentialsSecret (%s) to be created: %s", d.Id(), err.Error()), CustomCredentialsSecretResourceName, "create")
		return tfErr.GetDiag()
	}

	return resourceIbmSmCustomCredentialsSecretRead(context, d, meta)
}

func waitForIbmSmCustomCredentialsSecretCreate(secretsManagerClient *secretsmanagerv2.SecretsManagerV2, d *schema.ResourceData) (interface{}, error) {
	getSecretOptions := &secretsmanagerv2.GetSecretOptions{}
	id := strings.Split(d.Id(), "/")
	secretId := id[2]
	getSecretOptions.SetID(secretId)

	stateConf := &resource.StateChangeConf{
		Pending: []string{"pre_activation"},
		Target:  []string{"active"},
		Refresh: func() (interface{}, string, error) {
			secretIntf, _, err := secretsManagerClient.GetSecret(getSecretOptions)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return nil, "", fmt.Errorf("The secret does not exist anymore")
				}
				return nil, "", err
			}
			secret := secretIntf.(*secretsmanagerv2.CustomCredentialsSecret)
			if *secret.StateDescription == "destroyed" {
				return secret, *secret.StateDescription, fmt.Errorf("Failed to get the secret %w", err)
			}
			return secret, *secret.StateDescription, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      0 * time.Second,
		MinTimeout: 5 * time.Second,
	}

	return stateConf.WaitForState()
}

func resourceIbmSmCustomCredentialsSecretRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, endpointsFile, err := getSecretsManagerSession(meta.(conns.ClientSession))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, "", CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}

	id := strings.Split(d.Id(), "/")
	if len(id) != 3 {
		tfErr := flex.TerraformErrorf(nil, "Wrong format of resource ID. To import a secret use the format `<region>/<instance_id>/<secret_id>`", CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	region := id[0]
	instanceId := id[1]
	secretId := id[2]
	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, instanceId, region, getEndpointType(secretsManagerClient, d), endpointsFile)

	getSecretOptions := &secretsmanagerv2.GetSecretOptions{}

	getSecretOptions.SetID(secretId)

	secretIntf, response, err := secretsManagerClient.GetSecretWithContext(context, getSecretOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetSecretWithContext failed %s\n%s", err, response)
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSecretWithContext failed %s\n%s", err, response), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	secret := secretIntf.(*secretsmanagerv2.CustomCredentialsSecret)

	if err = d.Set("secret_id", secretId); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting secret_id"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("instance_id", instanceId); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting instance_id"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("region", region); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting region"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("created_by", secret.CreatedBy); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting created_by"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("created_at", DateTimeToRFC3339(secret.CreatedAt)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting created_at"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("crn", secret.Crn); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting crn"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if secret.CustomMetadata != nil {
		d.Set("custom_metadata", secret.CustomMetadata)
	}
	if err = d.Set("description", secret.Description); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting description"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("downloaded", secret.Downloaded); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting downloaded"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if secret.Labels != nil {
		if err = d.Set("labels", secret.Labels); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting labels"), CustomCredentialsSecretResourceName, "read")
			return tfErr.GetDiag()
		}
	}
	if err = d.Set("locks_total", flex.IntValue(secret.LocksTotal)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting locks_total"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("name", secret.Name); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting name"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("secret_group_id", secret.SecretGroupID); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting secret_group_id"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("state", flex.IntValue(secret.State)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting state"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("state_description", secret.StateDescription); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting state_description"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("ttl", secret.TTL); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting ttl"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("updated_at", DateTimeToRFC3339(secret.UpdatedAt)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting updated_at"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("retrieved_at", DateTimeToRFC3339(secret.RetrievedAt)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting retrieved_at"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("versions_total", flex.IntValue(secret.VersionsTotal)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting versions_total"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("configuration", secret.Configuration); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting configuration"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if secret.Rotation != nil {
		rotationMap, err := customCredentialsSecretRotationPolicyToMap(secret.Rotation)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, "", CustomCredentialsSecretResourceName, "read")
			return tfErr.GetDiag()
		}
		if err = d.Set("rotation", []map[string]interface{}{rotationMap}); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting rotation"), CustomCredentialsSecretResourceName, "read")
			return tfErr.GetDiag()
		}
	}
	if err = d.Set("expiration_date", DateTimeToRFC3339(secret.ExpirationDate)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting expiration_date"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("next_rotation_date", DateTimeToRFC3339(secret.NextRotationDate)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting next_rotation_date"), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}
	parameters := []map[string]interface{}{}
	if secret.Parameters != nil {
		modelMap, err := customCredentialsSecretFieldsToMap(secret.Parameters)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, "", fmt.Sprintf("(Data) %s", CustomCredentialsSecretResourceName), "read")
			return tfErr.GetDiag()
		}
		parameters = append(parameters, modelMap)
	}
	if err = d.Set("parameters", parameters); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting parameters"), fmt.Sprintf("(Data) %s", CustomCredentialsSecretResourceName), "read")
		return tfErr.GetDiag()
	}

	credentials := []map[string]interface{}{}
	if secret.CredentialsContent != nil {
		modelMap, err := customCredentialsSecretFieldsToMap(secret.CredentialsContent)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, "", fmt.Sprintf("(Data) %s", CustomCredentialsSecretResourceName), "read")
			return tfErr.GetDiag()
		}
		credentials = append(credentials, modelMap)
		log.Printf("[DEBUG] credentials failed %v", credentials)
	}
	if err = d.Set("credentials_content", credentials); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting credentials_content"), fmt.Sprintf("(Data) %s", CustomCredentialsSecretResourceName), "read")
		return tfErr.GetDiag()
	}

	if err != nil {
		tfErr := flex.TerraformErrorf(err, "", CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}

	// Call get version metadata API to get the current version_custom_metadata
	getVersionMetdataOptions := &secretsmanagerv2.GetSecretVersionMetadataOptions{}
	getVersionMetdataOptions.SetSecretID(secretId)
	getVersionMetdataOptions.SetID("current")

	versionMetadataIntf, response, err := secretsManagerClient.GetSecretVersionMetadataWithContext(context, getVersionMetdataOptions)
	if err != nil {
		log.Printf("[DEBUG] GetSecretVersionMetadataWithContext failed %s\n%s", err, response)
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSecretVersionMetadataWithContext failed %s\n%s", err, response), CustomCredentialsSecretResourceName, "read")
		return tfErr.GetDiag()
	}

	versionMetadata := versionMetadataIntf.(*secretsmanagerv2.CustomCredentialsSecretVersionMetadata)
	if versionMetadata.VersionCustomMetadata != nil {
		if err = d.Set("version_custom_metadata", versionMetadata.VersionCustomMetadata); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting version_custom_metadata"), CustomCredentialsSecretResourceName, "read")
			return tfErr.GetDiag()
		}
	}

	return nil
}

func resourceIbmSmCustomCredentialsSecretUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, endpointsFile, err := getSecretsManagerSession(meta.(conns.ClientSession))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, "", CustomCredentialsSecretResourceName, "update")
		return tfErr.GetDiag()
	}

	id := strings.Split(d.Id(), "/")
	region := id[0]
	instanceId := id[1]
	secretId := id[2]
	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, instanceId, region, getEndpointType(secretsManagerClient, d), endpointsFile)

	updateSecretMetadataOptions := &secretsmanagerv2.UpdateSecretMetadataOptions{}

	updateSecretMetadataOptions.SetID(secretId)

	hasChange := false

	patchVals := &secretsmanagerv2.CustomCredentialsSecretMetadataPatch{}

	if d.HasChange("name") {
		patchVals.Name = core.StringPtr(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("description") {
		patchVals.Description = core.StringPtr(d.Get("description").(string))
		hasChange = true
	}
	if d.HasChange("labels") {
		labels := d.Get("labels").([]interface{})
		labelsParsed := make([]string, len(labels))
		for i, v := range labels {
			labelsParsed[i] = fmt.Sprint(v)
		}
		patchVals.Labels = labelsParsed
		hasChange = true
	}
	if d.HasChange("ttl") {
		patchVals.TTL = core.StringPtr(d.Get("ttl").(string))
		hasChange = true
	}
	if d.HasChange("custom_metadata") {
		patchVals.CustomMetadata = d.Get("custom_metadata").(map[string]interface{})
		hasChange = true
	}
	if d.HasChange("rotation") {
		RotationModel, err := resourceIbmSmCustomCredentialsSecretMapToRotationPolicy(d.Get("rotation").([]interface{})[0].(map[string]interface{}))
		if err != nil {
			log.Printf("[DEBUG] UpdateSecretMetadataWithContext failed: Reading Rotation parameter failed: %s", err)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSecretMetadataWithContext failed: Reading Rotation parameter failed: %s", err), CustomCredentialsSecretResourceName, "update")
			return tfErr.GetDiag()
		}
		patchVals.Rotation = RotationModel
		hasChange = true
	}

	if d.HasChange("parameters") {
		parameters, err := resourceIbmSmCustomCredentialsSecretMapToParameters(d.Get("parameters").([]interface{})[0].(map[string]interface{}))
		if err != nil {
			log.Printf("[DEBUG] UpdateSecretMetadataWithContext failed: Reading parameters failed: %s", err)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSecretMetadataWithContext failed: Reading parameters failed: %s", err), CustomCredentialsSecretResourceName, "update")
			return tfErr.GetDiag()
		}
		patchVals.Parameters = parameters
		hasChange = true
	}

	if hasChange {
		updateSecretMetadataOptions.SecretMetadataPatch, _ = patchVals.AsPatch()
		_, response, err := secretsManagerClient.UpdateSecretMetadataWithContext(context, updateSecretMetadataOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateSecretMetadataWithContext failed %s\n%s", err, response)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSecretMetadataWithContext failed %s\n%s", err, response), CustomCredentialsSecretResourceName, "update")
			return tfErr.GetDiag()
		}
	} else if d.HasChange("version_custom_metadata") {
		// Apply change to version_custom_metadata in current version
		secretVersionMetadataPatchModel := new(secretsmanagerv2.SecretVersionMetadataPatch)
		secretVersionMetadataPatchModel.VersionCustomMetadata = d.Get("version_custom_metadata").(map[string]interface{})
		secretVersionMetadataPatchModelAsPatch, _ := secretVersionMetadataAsPatchFunction(secretVersionMetadataPatchModel)

		updateSecretVersionOptions := &secretsmanagerv2.UpdateSecretVersionMetadataOptions{}
		updateSecretVersionOptions.SetSecretID(secretId)
		updateSecretVersionOptions.SetID("current")
		updateSecretVersionOptions.SetSecretVersionMetadataPatch(secretVersionMetadataPatchModelAsPatch)
		_, response, err := secretsManagerClient.UpdateSecretVersionMetadataWithContext(context, updateSecretVersionOptions)
		if err != nil {
			if hasChange {
				// Call the read function to update the Terraform state with the change already applied to the metadata
				resourceIbmSmCustomCredentialsSecretRead(context, d, meta)
			}
			log.Printf("[DEBUG] UpdateSecretVersionMetadataWithContext failed %s\n%s", err, response)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSecretVersionMetadataWithContext failed %s\n%s", err, response), CustomCredentialsSecretResourceName, "update")
			return tfErr.GetDiag()
		}
	}

	return resourceIbmSmCustomCredentialsSecretRead(context, d, meta)
}

func resourceIbmSmCustomCredentialsSecretDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, endpointsFile, err := getSecretsManagerSession(meta.(conns.ClientSession))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, "", CustomCredentialsSecretResourceName, "delete")
		return tfErr.GetDiag()
	}

	id := strings.Split(d.Id(), "/")
	region := id[0]
	instanceId := id[1]
	secretId := id[2]
	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, instanceId, region, getEndpointType(secretsManagerClient, d), endpointsFile)

	// Clear the data from versions. Start by getting the versions
	listVersionsOptions := &secretsmanagerv2.ListSecretVersionsOptions{}
	listVersionsOptions.SetSecretID(secretId)
	versionsResult, response, err := secretsManagerClient.ListSecretVersionsWithContext(context, listVersionsOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteSecretWithContext failed %s\n%s", err, response)
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListSecretVersionsWithContext failed %s\n%s", err, response), CustomCredentialsSecretResourceName, "delete")
		return tfErr.GetDiag()
	}
	versions := versionsResult.Versions
	deleteSecretVersionDataOptions := &secretsmanagerv2.DeleteSecretVersionDataOptions{}
	deleteSecretVersionDataOptions.SetSecretID(secretId)

	// Clear the version data from previous version if exists
	if len(versions) > 1 {
		previousVersion := versions[len(versions)-2].(*secretsmanagerv2.CustomCredentialsSecretVersionMetadata)
		if *previousVersion.PayloadAvailable {
			diagnostics := ibmSmCustomCredentialsClearVersion(secretsManagerClient, context, d, secretId, *previousVersion.ID)
			if diagnostics != nil {
				return diagnostics
			}
		}
	}

	// Clear the version data from the current version
	currentVersion := versions[len(versions)-1].(*secretsmanagerv2.CustomCredentialsSecretVersionMetadata)
	if *currentVersion.PayloadAvailable {
		diagnostics := ibmSmCustomCredentialsClearVersion(secretsManagerClient, context, d, secretId, *currentVersion.ID)
		if diagnostics != nil {
			return diagnostics
		}
	}

	// delete the secret
	deleteSecretOptions := &secretsmanagerv2.DeleteSecretOptions{}
	deleteSecretOptions.SetID(secretId)
	deleteSecretOptions.SetForceDelete(true)

	response, err = secretsManagerClient.DeleteSecretWithContext(context, deleteSecretOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteSecretWithContext failed %s\n%s", err, response)
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteSecretWithContext failed %s\n%s", err, response), CustomCredentialsSecretResourceName, "delete")
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ibmSmCustomCredentialsClearVersion(secretsManagerClient *secretsmanagerv2.SecretsManagerV2, context context.Context, d *schema.ResourceData, secretId, versionId string) diag.Diagnostics {
	deleteSecretVersionDataOptions := &secretsmanagerv2.DeleteSecretVersionDataOptions{}
	deleteSecretVersionDataOptions.SetSecretID(secretId)
	deleteSecretVersionDataOptions.SetID(versionId)
	response, err := secretsManagerClient.DeleteSecretVersionDataWithContext(context, deleteSecretVersionDataOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteSecretWithContext failed %s\n%s", err, response)
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteSecretVersionDataWithContext failed %s\n%s", err, response), CustomCredentialsSecretResourceName, "delete")
		return tfErr.GetDiag()
	}

	// Wait for the delete secret version task to start working. We don't have to wait for it to finish, because we
	// are going to force-delete the secret, but at least we want it to leave the queue. Once it started SM will continue
	// to retry even after the force-delete
	stateConf := &resource.StateChangeConf{
		Pending: []string{"queued"},
		Target:  []string{"credentials_deleted", "processing", "failed"},
		Refresh: func() (interface{}, string, error) {
			listSecretTasksOptions := &secretsmanagerv2.ListSecretTasksOptions{}
			listSecretTasksOptions.SetSecretID(secretId)

			taskList, response, err := secretsManagerClient.ListSecretTasksWithContext(context, listSecretTasksOptions)
			if err != nil {
				log.Printf("[DEBUG] ListSecretTasksWithContext failed %s\n%s", err, response)
				return nil, "", err
			}
			for _, task := range taskList.Tasks {
				if *task.Type == "delete_credentials" && *task.SecretVersionID == versionId {
					log.Printf("[DEBUG] Found the task with status %s for version %s", *task.Status, *task.SecretVersionID)
					return task, *task.Status, nil
				}
			}
			log.Printf("[DEBUG] clear version data - task not found. Secret ID: %s", secretId)
			// return a status of "failed" to stop the waiting loop
			return &secretsmanagerv2.SecretTask{}, "failed", nil
		},
		Timeout:    2 * time.Minute,
		Delay:      0 * time.Second,
		MinTimeout: 5 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error waiting for resource IbmSmCustomCredentialsSecret (%s) to be deleted: %s", d.Id(), err.Error()), CustomCredentialsSecretResourceName, "create")
		return tfErr.GetDiag()
	}
	return nil
}

func resourceIbmSmCustomCredentialsSecretMapToSecretPrototype(d *schema.ResourceData) (*secretsmanagerv2.CustomCredentialsSecretPrototype, error) {
	model := &secretsmanagerv2.CustomCredentialsSecretPrototype{}
	model.SecretType = core.StringPtr("custom_credentials")

	if _, ok := d.GetOk("name"); ok {
		model.Name = core.StringPtr(d.Get("name").(string))
	}
	if _, ok := d.GetOk("custom_metadata"); ok {
		model.CustomMetadata = d.Get("custom_metadata").(map[string]interface{})
	}
	if _, ok := d.GetOk("description"); ok {
		model.Description = core.StringPtr(d.Get("description").(string))
	}
	if _, ok := d.GetOk("configuration"); ok {
		model.Configuration = core.StringPtr(d.Get("configuration").(string))
	}
	if _, ok := d.GetOk("labels"); ok {
		labels := d.Get("labels").([]interface{})
		labelsParsed := make([]string, len(labels))
		for i, v := range labels {
			labelsParsed[i] = fmt.Sprint(v)
		}
		model.Labels = labelsParsed
	}
	if _, ok := d.GetOk("secret_group_id"); ok {
		model.SecretGroupID = core.StringPtr(d.Get("secret_group_id").(string))
	}
	if _, ok := d.GetOk("ttl"); ok {
		model.TTL = core.StringPtr(d.Get("ttl").(string))
	}
	if _, ok := d.GetOk("version_custom_metadata"); ok {
		model.VersionCustomMetadata = d.Get("version_custom_metadata").(map[string]interface{})
	}
	if _, ok := d.GetOk("rotation"); ok {
		RotationModel, err := resourceIbmSmCustomCredentialsSecretMapToRotationPolicy(d.Get("rotation").([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Rotation = RotationModel
	}
	if _, ok := d.GetOk("parameters"); ok {
		parameters, err := resourceIbmSmCustomCredentialsSecretMapToParameters(d.Get("parameters").([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Parameters = parameters
	}

	return model, nil
}

func resourceIbmSmCustomCredentialsSecretMapToRotationPolicy(modelMap map[string]interface{}) (secretsmanagerv2.RotationPolicyIntf, error) {
	model := &secretsmanagerv2.RotationPolicy{}
	if modelMap["auto_rotate"] != nil {
		model.AutoRotate = core.BoolPtr(modelMap["auto_rotate"].(bool))
	}
	if modelMap["interval"].(int) == 0 {
		model.Interval = nil
	} else {
		model.Interval = core.Int64Ptr(int64(modelMap["interval"].(int)))
	}
	if modelMap["unit"] != nil && modelMap["unit"].(string) != "" {
		model.Unit = core.StringPtr(modelMap["unit"].(string))
	}
	return model, nil
}

func customCredentialsSecretRotationPolicyToMap(model secretsmanagerv2.RotationPolicyIntf) (map[string]interface{}, error) {
	if _, ok := model.(*secretsmanagerv2.CommonRotationPolicy); ok {
		return customCredentialsSecretCommonRotationPolicyToMap(model.(*secretsmanagerv2.CommonRotationPolicy))
	} else if _, ok := model.(*secretsmanagerv2.RotationPolicy); ok {
		modelMap := make(map[string]interface{})
		model := model.(*secretsmanagerv2.RotationPolicy)
		if model.AutoRotate != nil {
			modelMap["auto_rotate"] = model.AutoRotate
		}
		if model.Interval != nil {
			modelMap["interval"] = flex.IntValue(model.Interval)
		}
		if model.Unit != nil {
			modelMap["unit"] = model.Unit
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized secretsmanagerv2.RotationPolicyIntf subtype encountered")
	}
}

func customCredentialsSecretCommonRotationPolicyToMap(model *secretsmanagerv2.CommonRotationPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["auto_rotate"] = model.AutoRotate
	if model.Interval != nil {
		modelMap["interval"] = flex.IntValue(model.Interval)
	}
	if model.Unit != nil {
		modelMap["unit"] = model.Unit
	}
	return modelMap, nil
}

func resourceIbmSmCustomCredentialsSecretMapToParameters(modelMap map[string]interface{}) (map[string]interface{}, error) {
	parameters := map[string]interface{}{}
	if modelMap["boolean_values"] != nil {
		boolValues := modelMap["boolean_values"].(map[string]interface{})
		for key, value := range boolValues {
			parameters[key] = value
		}
	}
	if modelMap["integer_values"] != nil {
		intValues := modelMap["integer_values"].(map[string]interface{})
		for key, value := range intValues {
			parameters[key] = value
		}
	}
	if modelMap["string_values"] != nil {
		stringValues := modelMap["string_values"].(map[string]interface{})
		for key, value := range stringValues {
			parameters[key] = value
		}
	}
	return parameters, nil
}

// Convert either a map of custom credentials parameters or a map of credentials to the model map with int, bool and string values
func customCredentialsSecretFieldsToMap(fields map[string]interface{}) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["boolean_values"] = make(map[string]bool)
	modelMap["integer_values"] = make(map[string]int)
	modelMap["string_values"] = make(map[string]string)
	for key, value := range fields {
		switch keyType := value.(type) {
		case float64:
			modelMap["integer_values"].(map[string]int)[key] = int(value.(float64))
		case string:
			modelMap["string_values"].(map[string]string)[key] = value.(string)
		case bool:
			modelMap["boolean_values"].(map[string]bool)[key] = value.(bool)
		default:
			return nil, errors.New(fmt.Sprintf("[ERROR] field type not supported. Key: %s, Value: %v, Type: %v", key, value, keyType))
		}
	}
	return modelMap, nil
}

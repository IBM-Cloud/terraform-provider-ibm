// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-go-sdk/secretsmanagerv2"
)

func ResourceIbmSmConfigurationIamCredentials() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSmConfigurationIamCredentialsCreate,
		ReadContext:   resourceIbmSmConfigurationIamCredentialsRead,
		UpdateContext: resourceIbmSmConfigurationIamCredentialsUpdate,
		DeleteContext: resourceIbmSmConfigurationIamCredentialsDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"config_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The configuration type.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "A human-readable unique name to assign to your configuration.To protect your privacy, do not use personal data, such as your name or location, as an name for your secret.",
			},
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "The API key that is generated for this secret.After the secret reaches the end of its lease (see the `ttl` field), the API key is deleted automatically. If you want to continue to use the same API key for future read operations, see the `reuse_api_key` field.",
			},
			"secret_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The secret type. Supported types are arbitrary, certificates (imported, public, and private), IAM credentials, key-value, and user credentials.",
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
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date when a resource was recently modified. The date format follows RFC 3339.",
			},
		},
	}
}

func resourceIbmSmConfigurationIamCredentialsCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, err := meta.(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return diag.FromErr(err)
	}

	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, d)

	createConfigurationOptions := &secretsmanagerv2.CreateConfigurationOptions{}

	configurationPrototypeModel, err := resourceIbmSmConfigurationIamCredentialsMapToConfigurationPrototype(d)
	if err != nil {
		return diag.FromErr(err)
	}
	createConfigurationOptions.SetConfigurationPrototype(configurationPrototypeModel)

	configurationIntf, response, err := secretsManagerClient.CreateConfigurationWithContext(context, createConfigurationOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateConfigurationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateConfigurationWithContext failed %s\n%s", err, response))
	}
	configuration := configurationIntf.(*secretsmanagerv2.IAMCredentialsConfiguration)

	d.SetId(*configuration.Name)

	return resourceIbmSmConfigurationIamCredentialsRead(context, d, meta)
}

func resourceIbmSmConfigurationIamCredentialsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, err := meta.(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return diag.FromErr(err)
	}

	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, d)

	getConfigurationOptions := &secretsmanagerv2.GetConfigurationOptions{}

	getConfigurationOptions.SetName(d.Id())

	configurationIntf, response, err := secretsManagerClient.GetConfigurationWithContext(context, getConfigurationOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetConfigurationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetConfigurationWithContext failed %s\n%s", err, response))
	}
	configuration := configurationIntf.(*secretsmanagerv2.IAMCredentialsConfiguration)

	if err = d.Set("config_type", configuration.ConfigType); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting config_type: %s", err))
	}
	if err = d.Set("secret_type", configuration.SecretType); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting secret_type: %s", err))
	}
	if err = d.Set("created_by", configuration.CreatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_by: %s", err))
	}
	if err = d.Set("created_at", flex.DateTimeToString(configuration.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}
	if err = d.Set("updated_at", flex.DateTimeToString(configuration.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}
	if err = d.Set("api_key", configuration.ApiKey); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting api_key: %s", err))
	}

	return nil
}

func resourceIbmSmConfigurationIamCredentialsUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, err := meta.(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return diag.FromErr(err)
	}

	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, d)

	updateConfigurationOptions := &secretsmanagerv2.UpdateConfigurationOptions{}

	updateConfigurationOptions.SetName(d.Get("name").(string))
	updateConfigurationOptions.SetXSmAcceptConfigurationType("iam_credentials_configuration")

	hasChange := false

	patchVals := &secretsmanagerv2.ConfigurationPatch{}

	if d.HasChange("api_key") {
		patchVals.ApiKey = core.StringPtr(d.Get("api_key").(string))
		hasChange = true
	}

	if hasChange {
		updateConfigurationOptions.ConfigurationPatch, _ = patchVals.AsPatch()
		_, response, err := secretsManagerClient.UpdateConfigurationWithContext(context, updateConfigurationOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateConfigurationWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateConfigurationWithContext failed %s\n%s", err, response))
		}
	}

	return resourceIbmSmConfigurationIamCredentialsRead(context, d, meta)
}

func resourceIbmSmConfigurationIamCredentialsDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, err := meta.(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return diag.FromErr(err)
	}

	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, d)

	deleteConfigurationOptions := &secretsmanagerv2.DeleteConfigurationOptions{}

	deleteConfigurationOptions.SetName(d.Id())

	response, err := secretsManagerClient.DeleteConfigurationWithContext(context, deleteConfigurationOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteConfigurationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteConfigurationWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func resourceIbmSmConfigurationIamCredentialsMapToConfigurationPrototype(d *schema.ResourceData) (secretsmanagerv2.ConfigurationPrototypeIntf, error) {
	model := &secretsmanagerv2.IAMCredentialsConfigurationPrototype{}

	model.ConfigType = core.StringPtr("iam_credentials_configuration")

	if _, ok := d.GetOk("name"); ok {
		model.Name = core.StringPtr(d.Get("name").(string))
	}
	if _, ok := d.GetOk("api_key"); ok {
		model.ApiKey = core.StringPtr(d.Get("api_key").(string))
	}
	return model, nil
}

// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/go-sdk-core/v5/core"
)

func ResourceIbmCodeEngineSecret() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmCodeEngineSecretCreate,
		ReadContext:   resourceIbmCodeEngineSecretRead,
		UpdateContext: resourceIbmCodeEngineSecretUpdate,
		DeleteContext: resourceIbmCodeEngineSecretDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_secret", "project_id"),
				Description:  "The ID of the project.",
			},
			"format": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_secret", "format"),
				Description:  "Specify the format of the secret.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_secret", "name"),
				Description:  "The name of the secret.",
			},
			"data": &schema.Schema{
				Type:      schema.TypeMap,
				Optional:  true,
				Sensitive: true,
				Elem:      &schema.Schema{Type: schema.TypeString},
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the resource was created.",
			},
			"entity_tag": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version of the secret instance, which is used to achieve optimistic locking.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "When you provision a new secret,  a URL is created identifying the location of the instance.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The identifier of the resource.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the secret.",
			},
			"secret_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The identifier of the resource.",
			},
			"etag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceIbmCodeEngineSecretValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "project_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$`,
			MinValueLength:             36,
			MaxValueLength:             36,
		},
		validate.ValidateSchema{
			Identifier:                 "format",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "basic_auth, generic, registry, ssh_auth, tls",
			Regexp:                     `^(generic|ssh_auth|basic_auth|tls|registry)$`,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$`,
			MinValueLength:             1,
			MaxValueLength:             253,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_code_engine_secret", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmCodeEngineSecretCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	createSecretOptions := &codeenginev2.CreateSecretOptions{}

	createSecretOptions.SetProjectID(d.Get("project_id").(string))
	createSecretOptions.SetFormat(d.Get("format").(string))
	createSecretOptions.SetName(d.Get("name").(string))
	if _, ok := d.GetOk("data"); ok {
		dataModel, err := resourceIbmCodeEngineSecretMapToSecretData(d.Get("data").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createSecretOptions.SetData(dataModel)
	}

	secret, response, err := codeEngineClient.CreateSecretWithContext(context, createSecretOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateSecretWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateSecretWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *createSecretOptions.ProjectID, *secret.Name))

	return resourceIbmCodeEngineSecretRead(context, d, meta)
}

func resourceIbmCodeEngineSecretRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	getSecretOptions := &codeenginev2.GetSecretOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	getSecretOptions.SetProjectID(parts[0])
	getSecretOptions.SetName(parts[1])

	secret, response, err := codeEngineClient.GetSecretWithContext(context, getSecretOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetSecretWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetSecretWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("project_id", secret.ProjectID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting project_id: %s", err))
	}
	if err = d.Set("format", secret.Format); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting format: %s", err))
	}
	if err = d.Set("name", secret.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if !core.IsNil(secret.Data) {
		data := make(map[string]string)
		for k, v := range secret.Data {
			data[k] = string(v)
		}
		if err = d.Set("data", data); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting data: %s", err))
		}
	}
	if !core.IsNil(secret.CreatedAt) {
		if err = d.Set("created_at", secret.CreatedAt); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
		}
	}
	if err = d.Set("entity_tag", secret.EntityTag); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting entity_tag: %s", err))
	}
	if !core.IsNil(secret.Href) {
		if err = d.Set("href", secret.Href); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
		}
	}
	if !core.IsNil(secret.ID) {
		if err = d.Set("id", secret.ID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting id: %s", err))
		}
	}
	if !core.IsNil(secret.ResourceType) {
		if err = d.Set("resource_type", secret.ResourceType); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting resource_type: %s", err))
		}
	}
	if !core.IsNil(secret.ID) {
		if err = d.Set("secret_id", secret.ID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting secret_id: %s", err))
		}
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting etag: %s", err))
	}

	return nil
}

func resourceIbmCodeEngineSecretUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	replaceSecretOptions := &codeenginev2.ReplaceSecretOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	replaceSecretOptions.SetProjectID(parts[0])
	replaceSecretOptions.SetName(parts[1])
	replaceSecretOptions.SetFormat(d.Get("format").(string))

	hasChange := false

	if d.HasChange("format") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "format"))
	}
	if d.HasChange("name") {
		replaceSecretOptions.SetName(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("data") {
		data, err := resourceIbmCodeEngineSecretMapToSecretData(d.Get("data").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		replaceSecretOptions.SetData(data)
		hasChange = true
	}
	replaceSecretOptions.SetIfMatch(d.Get("etag").(string))

	if hasChange {
		_, response, err := codeEngineClient.ReplaceSecretWithContext(context, replaceSecretOptions)
		if err != nil {
			log.Printf("[DEBUG] ReplaceSecretWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("ReplaceSecretWithContext failed %s\n%s", err, response))
		}
	}

	return resourceIbmCodeEngineSecretRead(context, d, meta)
}

func resourceIbmCodeEngineSecretDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteSecretOptions := &codeenginev2.DeleteSecretOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	deleteSecretOptions.SetProjectID(parts[0])
	deleteSecretOptions.SetName(parts[1])

	response, err := codeEngineClient.DeleteSecretWithContext(context, deleteSecretOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteSecretWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteSecretWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func resourceIbmCodeEngineSecretMapToSecretData(modelMap map[string]interface{}) (codeenginev2.SecretDataIntf, error) {
	model := &codeenginev2.SecretData{}

	for key, value := range modelMap {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)

		model.SetProperty(strKey, &strValue)
	}
	return model, nil
}

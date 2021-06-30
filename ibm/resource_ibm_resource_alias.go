// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

func resourceIbmResourceAlias() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmResourceAliasCreate,
		ReadContext:   resourceIbmResourceAliasRead,
		UpdateContext: resourceIbmResourceAliasUpdate,
		DeleteContext: resourceIbmResourceAliasDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: InvokeValidator("ibm_resource_alias", "name"),
				Description:  "The name of the alias. Must be 180 characters or less and cannot include any special characters other than `(space) - . _ :`.",
			},
			"source": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The short or long ID of resource instance.",
			},
			"target": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The CRN of target name(space) in a specific environment, for example, space in Dallas YP, CFEE instance etc.",
			},
			"guid": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "When you create a new alias, a globally unique identifier (GUID) is assigned. This GUID is a unique internal indentifier managed by the resource controller that corresponds to the alias.",
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "When you created a new alias, a relative URL path is created identifying the location of the alias.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date when the alias was created.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date when the alias was last updated.",
			},
			"deleted_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date when the alias was deleted.",
			},
			"created_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The subject who created the alias.",
			},
			"updated_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The subject who updated the alias.",
			},
			"deleted_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The subject who deleted the alias.",
			},
			"resource_instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the resource instance that is being aliased.",
			},
			"target_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN of the target namespace in the specific environment.",
			},
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An alpha-numeric value identifying the account ID.",
			},
			"resource_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique ID of the offering. This value is provided by and stored in the global catalog.",
			},
			"resource_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the resource group.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN of the alias. For more information about this format, see [Cloud Resource Names](https://cloud.ibm.com/docs/overview?topic=overview-crn).",
			},
			"region_instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the instance in the specific target environment, for example, `service_instance_id` in a given IBM Cloud environment.",
			},
			"region_instance_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN of the instance in the specific target environment.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the alias.",
			},
			"migrated": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "A boolean that dictates if the alias was migrated from a previous CF instance.",
			},
			"resource_instance_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The relative path to the resource instance.",
			},
			"resource_bindings_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The relative path to the resource bindings for the alias.",
			},
			"resource_keys_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The relative path to the resource keys for the alias.",
			},
		},
	}
}

func resourceIbmResourceAliasValidator() *ResourceValidator {
	validateSchema := make([]ValidateSchema, 1)
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: ValidateRegexp,
			Type:                       TypeString,
			Required:                   true,
			Regexp:                     `^([^[:ascii:]]|[a-zA-Z0-9-._: ])+$`,
			MaxValueLength:             180,
		},
	)

	resourceValidator := ResourceValidator{ResourceName: "ibm_resource_alias", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmResourceAliasCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceControllerClient, err := meta.(ClientSession).ResourceControllerV2API()
	if err != nil {
		return diag.FromErr(err)
	}

	createResourceAliasOptions := &resourcecontrollerv2.CreateResourceAliasOptions{}

	createResourceAliasOptions.SetName(d.Get("name").(string))
	createResourceAliasOptions.SetSource(d.Get("source").(string))
	createResourceAliasOptions.SetTarget(d.Get("target").(string))

	log.Printf("[DEBUG] Creating Resource Alias : %s", *createResourceAliasOptions.Name)
	resourceAlias, response, err := resourceControllerClient.CreateResourceAliasWithContext(context, createResourceAliasOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateResourceAliasWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateResourceAliasWithContext failed %s\n%s", err, response))
	}

	log.Println("[DEBUG] Resource Alias created successfully")
	d.SetId(*resourceAlias.ID)

	return resourceIbmResourceAliasRead(context, d, meta)
}

func resourceIbmResourceAliasRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceControllerClient, err := meta.(ClientSession).ResourceControllerV2API()
	if err != nil {
		return diag.FromErr(err)
	}

	getResourceAliasOptions := &resourcecontrollerv2.GetResourceAliasOptions{}
	getResourceAliasOptions.SetID(d.Id())

	log.Printf("[DEBUG] Reading Resource Alias : %s", *getResourceAliasOptions.ID)
	resourceAlias, response, err := resourceControllerClient.GetResourceAliasWithContext(context, getResourceAliasOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetResourceAliasWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetResourceAliasWithContext failed %s\n%s", err, response))
	}

	if _, ok := d.GetOk("source"); !ok {
		// when resource is imported, source can be extracted from ResourceInstanceURL
		resourceInstanceURL := *resourceAlias.ResourceInstanceURL
		resourceInstanceID := resourceInstanceURL[strings.LastIndex(resourceInstanceURL, "/")+1:]
		if err = d.Set("source", resourceInstanceID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting source: %s", err))
		}
	}
	if _, ok := d.GetOk("target"); !ok {
		// when resource is imported, target can be set as TargetCRN
		if err = d.Set("target", resourceAlias.TargetCRN); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting target: %s", err))
		}
	}
	if err = d.Set("name", resourceAlias.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("guid", resourceAlias.GUID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting guid: %s", err))
	}
	if err = d.Set("url", resourceAlias.URL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting url: %s", err))
	}
	if err = d.Set("created_at", dateTimeToString(resourceAlias.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}
	if err = d.Set("updated_at", dateTimeToString(resourceAlias.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}
	if err = d.Set("deleted_at", dateTimeToString(resourceAlias.DeletedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting deleted_at: %s", err))
	}
	if err = d.Set("created_by", resourceAlias.CreatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_by: %s", err))
	}
	if err = d.Set("updated_by", resourceAlias.UpdatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_by: %s", err))
	}
	if err = d.Set("deleted_by", resourceAlias.DeletedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting deleted_by: %s", err))
	}
	if err = d.Set("resource_instance_id", resourceAlias.ResourceInstanceID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_instance_id: %s", err))
	}
	if err = d.Set("target_crn", resourceAlias.TargetCRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting target_crn: %s", err))
	}
	if err = d.Set("account_id", resourceAlias.AccountID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting account_id: %s", err))
	}
	if err = d.Set("resource_id", resourceAlias.ResourceID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_id: %s", err))
	}
	if err = d.Set("resource_group_id", resourceAlias.ResourceGroupID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group_id: %s", err))
	}
	if err = d.Set("crn", resourceAlias.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("region_instance_id", resourceAlias.RegionInstanceID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting region_instance_id: %s", err))
	}
	if err = d.Set("region_instance_crn", resourceAlias.RegionInstanceCRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting region_instance_crn: %s", err))
	}
	if err = d.Set("state", resourceAlias.State); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting state: %s", err))
	}
	if err = d.Set("migrated", resourceAlias.Migrated); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting migrated: %s", err))
	}
	if err = d.Set("resource_instance_url", resourceAlias.ResourceInstanceURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_instance_url: %s", err))
	}
	if err = d.Set("resource_bindings_url", resourceAlias.ResourceBindingsURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_bindings_url: %s", err))
	}
	if err = d.Set("resource_keys_url", resourceAlias.ResourceKeysURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_keys_url: %s", err))
	}

	log.Println("[DEBUG] Resource Alias read successfully")
	return nil
}

func resourceIbmResourceAliasUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceControllerClient, err := meta.(ClientSession).ResourceControllerV2API()
	if err != nil {
		return diag.FromErr(err)
	}

	updateResourceAliasOptions := &resourcecontrollerv2.UpdateResourceAliasOptions{}
	updateResourceAliasOptions.SetID(d.Id())

	hasChange := false

	if d.HasChange("name") {
		updateResourceAliasOptions.SetName(d.Get("name").(string))
		hasChange = true
	}

	if hasChange {
		log.Printf("[DEBUG] Updating Resource Alias Name : %s", *updateResourceAliasOptions.Name)
		_, response, err := resourceControllerClient.UpdateResourceAliasWithContext(context, updateResourceAliasOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateResourceAliasWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateResourceAliasWithContext failed %s\n%s", err, response))
		}
		log.Println("[DEBUG] Resource Alias updated successfully")
	}

	return resourceIbmResourceAliasRead(context, d, meta)
}

func resourceIbmResourceAliasDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceControllerClient, err := meta.(ClientSession).ResourceControllerV2API()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteResourceAliasOptions := &resourcecontrollerv2.DeleteResourceAliasOptions{}
	deleteResourceAliasOptions.SetID(d.Id())

	log.Printf("[DEBUG] Deleting Resource Alias : %s", *deleteResourceAliasOptions.ID)
	response, err := resourceControllerClient.DeleteResourceAliasWithContext(context, deleteResourceAliasOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteResourceAliasWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteResourceAliasWithContext failed %s\n%s", err, response))
	}

	log.Println("[DEBUG] Resource Alias deleted successfully")
	d.SetId("")

	return nil
}

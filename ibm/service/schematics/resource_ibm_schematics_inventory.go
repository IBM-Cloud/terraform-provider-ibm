// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func ResourceIBMSchematicsInventory() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIBMSchematicsInventoryCreate,
		ReadContext:   ResourceIBMSchematicsInventoryRead,
		UpdateContext: ResourceIBMSchematicsInventoryUpdate,
		DeleteContext: ResourceIBMSchematicsInventoryDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_schematics_inventory", "name"),
				Description:  "The unique name of your Inventory definition. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of your Inventory definition. The description can be up to 2048 characters long in size.",
			},
			"location": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_schematics_inventory", "location"),
				Description:  "List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Resource-group name for the Inventory definition.   By default, Inventory definition will be created in Default Resource Group.",
			},
			"inventories_ini": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Input inventory of host and host group for the playbook, in the `.ini` file format.",
			},
			"resource_queries": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Input resource query definitions that is used to dynamically generate the inventory of host and host group for the playbook.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Inventory creation time.",
			},
			"created_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Email address of user who created the Inventory.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Inventory updation time.",
			},
			"updated_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Email address of user who updated the Inventory.",
			},
		},
	}
}

func ResourceIBMSchematicsInventoryValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 1)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.StringLenBetween,
			Type:                       validate.TypeString,
			Optional:                   true,
			MinValueLength:             3,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "location",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "eu-de, eu-gb, us-east, us-south",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_schematics_inventory", Schema: validateSchema}
	return &resourceValidator
}

func ResourceIBMSchematicsInventoryCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createInventoryOptions := &schematicsv1.CreateInventoryOptions{}

	if _, ok := d.GetOk("name"); ok {
		createInventoryOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("description"); ok {
		createInventoryOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("location"); ok {
		createInventoryOptions.SetLocation(d.Get("location").(string))
	}
	if _, ok := d.GetOk("resource_group"); ok {
		createInventoryOptions.SetResourceGroup(d.Get("resource_group").(string))
	}
	if _, ok := d.GetOk("inventories_ini"); ok {
		createInventoryOptions.SetInventoriesIni(d.Get("inventories_ini").(string))
	}
	if _, ok := d.GetOk("resource_queries"); ok {
		createInventoryOptions.SetResourceQueries(d.Get("resource_queries").([]string))
	}

	inventoryResourceRecord, response, err := schematicsClient.CreateInventoryWithContext(context, createInventoryOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateInventoryWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateInventoryWithContext failed %s\n%s", err, response))
	}

	d.SetId(*inventoryResourceRecord.ID)

	return ResourceIBMSchematicsInventoryRead(context, d, meta)
}

func ResourceIBMSchematicsInventoryRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getInventoryOptions := &schematicsv1.GetInventoryOptions{}

	getInventoryOptions.SetInventoryID(d.Id())

	inventoryResourceRecord, response, err := schematicsClient.GetInventoryWithContext(context, getInventoryOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetInventoryWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetInventoryWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("name", inventoryResourceRecord.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("description", inventoryResourceRecord.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
	}
	if err = d.Set("location", inventoryResourceRecord.Location); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting location: %s", err))
	}
	if err = d.Set("resource_group", inventoryResourceRecord.ResourceGroup); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group: %s", err))
	}
	if err = d.Set("inventories_ini", inventoryResourceRecord.InventoriesIni); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting inventories_ini: %s", err))
	}
	if inventoryResourceRecord.ResourceQueries != nil {
		if err = d.Set("resource_queries", inventoryResourceRecord.ResourceQueries); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting resource_queries: %s", err))
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(inventoryResourceRecord.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}
	if err = d.Set("created_by", inventoryResourceRecord.CreatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_by: %s", err))
	}
	if err = d.Set("updated_at", flex.DateTimeToString(inventoryResourceRecord.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}
	if err = d.Set("updated_by", inventoryResourceRecord.UpdatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_by: %s", err))
	}

	return nil
}

func ResourceIBMSchematicsInventoryUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	replaceInventoryOptions := &schematicsv1.ReplaceInventoryOptions{}

	replaceInventoryOptions.SetInventoryID(d.Id())
	if _, ok := d.GetOk("name"); ok {
		replaceInventoryOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("description"); ok {
		replaceInventoryOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("location"); ok {
		replaceInventoryOptions.SetLocation(d.Get("location").(string))
	}
	if _, ok := d.GetOk("resource_group"); ok {
		replaceInventoryOptions.SetResourceGroup(d.Get("resource_group").(string))
	}
	if _, ok := d.GetOk("inventories_ini"); ok {
		replaceInventoryOptions.SetInventoriesIni(d.Get("inventories_ini").(string))
	}
	if _, ok := d.GetOk("resource_queries"); ok {
		replaceInventoryOptions.SetResourceQueries(d.Get("resource_queries").([]string))
	}

	_, response, err := schematicsClient.ReplaceInventoryWithContext(context, replaceInventoryOptions)
	if err != nil {
		log.Printf("[DEBUG] ReplaceInventoryWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ReplaceInventoryWithContext failed %s\n%s", err, response))
	}

	return ResourceIBMSchematicsInventoryRead(context, d, meta)
}

func ResourceIBMSchematicsInventoryDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteInventoryOptions := &schematicsv1.DeleteInventoryOptions{}

	deleteInventoryOptions.SetInventoryID(d.Id())

	response, err := schematicsClient.DeleteInventoryWithContext(context, deleteInventoryOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteInventoryWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteInventoryWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

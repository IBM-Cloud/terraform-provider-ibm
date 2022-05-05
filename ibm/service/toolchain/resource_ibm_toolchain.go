// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package toolchain

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.ibm.com/org-ids/toolchain-go-sdk/toolchainv2"
)

func ResourceIBMToolchain() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIBMToolchainCreate,
		ReadContext:   ResourceIBMToolchainRead,
		UpdateContext: ResourceIBMToolchainUpdate,
		DeleteContext: ResourceIBMToolchainDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_toolchain", "name"),
				Description:  "Toolchain name.",
			},
			"resource_group_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_toolchain", "resource_group_id"),
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Describes the toolchain.",
			},
			"account_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"crn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"href": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceIBMToolchainValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 1)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexp,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([^\\x00-\\x7F]|[a-zA-Z0-9-._ ])+$`,
			MaxValueLength:             128,
		},
		validate.ValidateSchema{
			Identifier:                 "resource_group_id",
			ValidateFunctionIdentifier: validate.ValidateRegexp,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9a-f]{32}$`,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_toolchain", Schema: validateSchema}
	return &resourceValidator
}

func ResourceIBMToolchainCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	toolchainClient, err := meta.(conns.ClientSession).ToolchainV2()
	if err != nil {
		return diag.FromErr(err)
	}

	postToolchainOptions := &toolchainv2.PostToolchainOptions{}

	postToolchainOptions.SetName(d.Get("name").(string))
	postToolchainOptions.SetResourceGroupID(d.Get("resource_group_id").(string))
	if _, ok := d.GetOk("description"); ok {
		postToolchainOptions.SetDescription(d.Get("description").(string))
	}

	postToolchainResponse, response, err := toolchainClient.PostToolchainWithContext(context, postToolchainOptions)
	if err != nil {
		log.Printf("[DEBUG] PostToolchainWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("PostToolchainWithContext failed %s\n%s", err, response))
	}

	d.SetId(*postToolchainResponse.ID)

	return ResourceIBMToolchainRead(context, d, meta)
}

func ResourceIBMToolchainRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	toolchainClient, err := meta.(conns.ClientSession).ToolchainV2()
	if err != nil {
		return diag.FromErr(err)
	}

	getToolchainByIDOptions := &toolchainv2.GetToolchainByIDOptions{}

	getToolchainByIDOptions.SetToolchainID(d.Id())

	getToolchainByIDResponse, response, err := toolchainClient.GetToolchainByIDWithContext(context, getToolchainByIDOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetToolchainByIDWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetToolchainByIDWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("name", getToolchainByIDResponse.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("resource_group_id", getToolchainByIDResponse.ResourceGroupID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group_id: %s", err))
	}
	if err = d.Set("description", getToolchainByIDResponse.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
	}
	if err = d.Set("account_id", getToolchainByIDResponse.AccountID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting account_id: %s", err))
	}
	if err = d.Set("location", getToolchainByIDResponse.Location); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting location: %s", err))
	}
	if err = d.Set("crn", getToolchainByIDResponse.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("href", getToolchainByIDResponse.Href); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
	}
	if err = d.Set("created_at", flex.DateTimeToString(getToolchainByIDResponse.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}
	if err = d.Set("updated_at", flex.DateTimeToString(getToolchainByIDResponse.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}
	if err = d.Set("created_by", getToolchainByIDResponse.CreatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_by: %s", err))
	}
	if err = d.Set("tags", getToolchainByIDResponse.Tags); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting tags: %s", err))
	}

	return nil
}

func ResourceIBMToolchainUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	toolchainClient, err := meta.(conns.ClientSession).ToolchainV2()
	if err != nil {
		return diag.FromErr(err)
	}

	patchToolchainOptions := &toolchainv2.PatchToolchainOptions{}

	patchToolchainOptions.SetToolchainID(d.Id())

	hasChange := false

	if d.HasChange("resource_group_id") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "resource_group_id"))
	}
	if d.HasChange("name") {
		patchToolchainOptions.SetName(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("description") {
		patchToolchainOptions.SetDescription(d.Get("description").(string))
		hasChange = true
	}

	if hasChange {
		response, err := toolchainClient.PatchToolchainWithContext(context, patchToolchainOptions)
		if err != nil {
			log.Printf("[DEBUG] PatchToolchainWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("PatchToolchainWithContext failed %s\n%s", err, response))
		}
	}

	return ResourceIBMToolchainRead(context, d, meta)
}

func ResourceIBMToolchainDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	toolchainClient, err := meta.(conns.ClientSession).ToolchainV2()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteToolchainOptions := &toolchainv2.DeleteToolchainOptions{}

	deleteToolchainOptions.SetToolchainID(d.Id())

	response, err := toolchainClient.DeleteToolchainWithContext(context, deleteToolchainOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteToolchainWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteToolchainWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

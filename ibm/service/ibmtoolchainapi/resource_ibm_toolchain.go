// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibmtoolchainapi

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/org-ids/toolchain-go-sdk/ibmtoolchainapiv2"
)

func ResourceIbmToolchain() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIbmToolchainCreate,
		ReadContext:   ResourceIbmToolchainRead,
		UpdateContext: ResourceIbmToolchainUpdate,
		DeleteContext: ResourceIbmToolchainDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Toolchain name.",
			},
			"generator": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_toolchain", "generator"),
				Description:  "A description of who generated the toolchain.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Describes the toolchain.",
			},
			"key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "<strong>Deprecated: </strong><br><br>Key of this toolchain, can be used when querying for toolchain.",
			},
			"container": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"guid": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN for resource group based toolchains.",
			},
			"creator": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"status": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The status of the toolchain.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"detailed_status": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "A list of particular status issues related to the toolchain.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The status of the particular issue'ok' indicates a normal state, additional messages are possible but unlikely'warning' indicates a possible problem with the toolchain, but usage may continue'error' indicates furthur use of this toolchain would be problematic, user actions should be blocked.",
									},
									"status_line": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A short status that can be used to mark up a toolchain card or other location. Should always be included when the 'status' is not 'ok'.",
									},
									"message": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A short message indicating the problem. Should always be included when the 'status' is not 'ok'.",
									},
									"details": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A longer description of the problem, typically message/details would be displayed together in a UI to give the user an understanding of the issue. Should always be included when the 'status' is not 'ok'.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func ResourceIbmToolchainValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 1)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "generator",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "API, Bluemix, IBM Bluemix DevOps Services, IBM Cloud, IBM Cloud DevOps Services, otc_service",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_toolchain", Schema: validateSchema}
	return &resourceValidator
}

func ResourceIbmToolchainCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmToolchainApiClient, err := meta.(conns.ClientSession).IbmToolchainApiV2()
	if err != nil {
		return diag.FromErr(err)
	}

	createToolchainOptions := &ibmtoolchainapiv2.CreateToolchainOptions{}

	createToolchainOptions.SetName(d.Get("name").(string))
	createToolchainOptions.SetGenerator(d.Get("generator").(string))
	if _, ok := d.GetOk("description"); ok {
		createToolchainOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("key"); ok {
		createToolchainOptions.SetKey(d.Get("key").(string))
	}
	if _, ok := d.GetOk("container"); ok {
		container, err := ResourceIbmToolchainMapToContainer(d.Get("container.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createToolchainOptions.SetContainer(container)
	}

	toolchainResponse, response, err := ibmToolchainApiClient.CreateToolchainWithContext(context, createToolchainOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateToolchainWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateToolchainWithContext failed %s\n%s", err, response))
	}

	d.SetId(*toolchainResponse.ToolchainGuid)

	return ResourceIbmToolchainRead(context, d, meta)
}

func ResourceIbmToolchainRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmToolchainApiClient, err := meta.(conns.ClientSession).IbmToolchainApiV2()
	if err != nil {
		return diag.FromErr(err)
	}

	getToolchainOptions := &ibmtoolchainapiv2.GetToolchainOptions{}

	getToolchainOptions.SetToolchainGuid(d.Id())

	toolchainResponse, response, err := ibmToolchainApiClient.GetToolchainWithContext(context, getToolchainOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetToolchainWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetToolchainWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("name", toolchainResponse.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("generator", toolchainResponse.Generator); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting generator: %s", err))
	}
	if err = d.Set("description", toolchainResponse.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
	}
	if err = d.Set("key", toolchainResponse.Key); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting key: %s", err))
	}
	if toolchainResponse.Container != nil {
		containerMap, err := ResourceIbmToolchainContainerToMap(toolchainResponse.Container)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("container", []map[string]interface{}{containerMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting container: %s", err))
		}
	}
	if err = d.Set("crn", toolchainResponse.Crn); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("creator", toolchainResponse.Creator); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting creator: %s", err))
	}

	if err = d.Set("tags", toolchainResponse.Tags); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting tags: %s", err))
	}
	if toolchainResponse.Status != nil {
		statusMap, err := ResourceIbmToolchainToolchainResponseStatusToMap(toolchainResponse.Status)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("status", []map[string]interface{}{statusMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting status: %s", err))
		}
	}

	return nil
}

func ResourceIbmToolchainUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmToolchainApiClient, err := meta.(conns.ClientSession).IbmToolchainApiV2()
	if err != nil {
		return diag.FromErr(err)
	}

	patchToolchainOptions := &ibmtoolchainapiv2.PatchToolchainOptions{}

	patchToolchainOptions.SetToolchainGuid(d.Id())

	hasChange := false

	if d.HasChange("generator") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "generator"))
	}
	if d.HasChange("key") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "key"))
	}
	if d.HasChange("container") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "container"))
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
		response, err := ibmToolchainApiClient.PatchToolchainWithContext(context, patchToolchainOptions)
		if err != nil {
			log.Printf("[DEBUG] PatchToolchainWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("PatchToolchainWithContext failed %s\n%s", err, response))
		}
	}

	return ResourceIbmToolchainRead(context, d, meta)
}

func ResourceIbmToolchainDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmToolchainApiClient, err := meta.(conns.ClientSession).IbmToolchainApiV2()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteToolchainOptions := &ibmtoolchainapiv2.DeleteToolchainOptions{}

	deleteToolchainOptions.SetToolchainGuid(d.Id())

	response, err := ibmToolchainApiClient.DeleteToolchainWithContext(context, deleteToolchainOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteToolchainWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteToolchainWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func ResourceIbmToolchainMapToContainer(modelMap map[string]interface{}) (*ibmtoolchainapiv2.Container, error) {
	model := &ibmtoolchainapiv2.Container{}
	model.Guid = core.StringPtr(modelMap["guid"].(string))
	model.Type = core.StringPtr(modelMap["type"].(string))
	return model, nil
}

func ResourceIbmToolchainContainerToMap(model *ibmtoolchainapiv2.Container) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["guid"] = model.Guid
	modelMap["type"] = model.Type
	return modelMap, nil
}

func ResourceIbmToolchainToolchainResponseStatusToMap(model *ibmtoolchainapiv2.ToolchainResponseStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	detailedStatus := []map[string]interface{}{}
	for _, detailedStatusItem := range model.DetailedStatus {
		detailedStatusItemMap, err := ResourceIbmToolchainToolchainResponseStatusDetailedStatusItemToMap(&detailedStatusItem)
		if err != nil {
			return modelMap, err
		}
		detailedStatus = append(detailedStatus, detailedStatusItemMap)
	}
	modelMap["detailed_status"] = detailedStatus
	return modelMap, nil
}

func ResourceIbmToolchainToolchainResponseStatusDetailedStatusItemToMap(model *ibmtoolchainapiv2.ToolchainResponseStatusDetailedStatusItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["status"] = model.Status
	if model.StatusLine != nil {
		modelMap["status_line"] = model.StatusLine
	}
	if model.Message != nil {
		modelMap["message"] = model.Message
	}
	if model.Details != nil {
		modelMap["details"] = model.Details
	}
	return modelMap, nil
}

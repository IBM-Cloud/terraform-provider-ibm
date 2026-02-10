// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
	"github.com/go-openapi/strfmt"
)

func ResourceIbmLogsView() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmLogsViewCreate,
		ReadContext:   resourceIbmLogsViewRead,
		UpdateContext: resourceIbmLogsViewUpdate,
		DeleteContext: resourceIbmLogsViewDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_view", "name"),
				Description:  "View name.",
			},
			"search_query": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "View search query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"query": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "View search query.",
						},
						"syntax_type": &schema.Schema{
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validate.InvokeValidator("ibm_logs_view", "syntax_type"),
							Description:  "Syntax type for the query used in views.",
						},
					},
				},
			},
			"time_selection": &schema.Schema{
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				Description: "View time selection.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"quick_selection": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Quick time selection.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"caption": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Quick time selection caption.",
									},
									"seconds": &schema.Schema{
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Quick time selection amount in seconds.",
									},
								},
							},
						},
						"custom_selection": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Custom time selection.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"from_time": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Custom time selection starting timestamp.",
									},
									"to_time": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Custom time selection ending timestamp.",
									},
								},
							},
						},
					},
				},
			},
			"filters": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "View selected filters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filters": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Selected filters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Filter name.",
									},
									"selected_values": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "Filter selected values.",
										Elem:        &schema.Schema{Type: schema.TypeBool},
									},
								},
							},
						},
					},
				},
			},
			"folder_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "View folder ID.",
			},
			"tier": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_view", "tier"),
				Description:  "Type of view.",
			},
			"view_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "View Id.",
			},
		},
	}
}

func ResourceIbmLogsViewValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$`,
			MinValueLength:             1,
			MaxValueLength:             4096,
		},
		validate.ValidateSchema{
			Identifier:                 "syntax_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "lucene, dataprime",
		},
		validate.ValidateSchema{
			Identifier:                 "tier",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "all_logs, all_logs_templates, priority_insights, priority_insights_templates",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_logs_view", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmLogsViewCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_logs_view", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getLogsInstanceRegion(logsClient, d)
	instanceId := d.Get("instance_id").(string)
	logsClient, err = getClientWithLogsInstanceEndpoint(logsClient, meta, instanceId, region, getLogsInstanceEndpointType(logsClient, d))
	if err != nil {
		return diag.FromErr(fmt.Errorf("Unable to get updated logs instance client"))
	}

	createViewOptions := &logsv0.CreateViewOptions{}

	createViewOptions.SetName(d.Get("name").(string))
	timeSelectionModel, err := ResourceIbmLogsViewMapToApisViewsV1TimeSelection(d.Get("time_selection.0").(map[string]interface{}))
	if err != nil {
		return diag.FromErr(err)
	}
	createViewOptions.SetTimeSelection(timeSelectionModel)
	if _, ok := d.GetOk("search_query"); ok {
		searchQueryModel, err := ResourceIbmLogsViewMapToApisViewsV1SearchQuery(d.Get("search_query.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createViewOptions.SetSearchQuery(searchQueryModel)
	}
	if _, ok := d.GetOk("filters"); ok {
		filtersModel, err := ResourceIbmLogsViewMapToApisViewsV1SelectedFilters(d.Get("filters.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createViewOptions.SetFilters(filtersModel)
	}
	if _, ok := d.GetOk("folder_id"); ok {
		createViewOptions.SetFolderID(core.UUIDPtr(strfmt.UUID(d.Get("folder_id").(string))))
	}
	if _, ok := d.GetOk("tier"); ok {
		createViewOptions.SetTier(d.Get("tier").(string))
	}

	view, response, err := logsClient.CreateViewWithContext(context, createViewOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateViewWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateViewWithContext failed %s\n%s", err, response))
	}

	viewId := fmt.Sprintf("%s/%s/%d", region, instanceId, *view.ID)
	d.SetId(viewId)

	return resourceIbmLogsViewRead(context, d, meta)
}

func resourceIbmLogsViewRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_logs_view", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	logsClient, region, instanceId, viewId, err := updateClientURLWithInstanceEndpoint(d.Id(), meta, logsClient, d)
	if err != nil {
		return diag.FromErr(err)
	}
	getViewOptions := &logsv0.GetViewOptions{}

	viewIdInt, _ := strconv.ParseInt(viewId, 10, 64)
	getViewOptions.SetID(viewIdInt)

	view, response, err := logsClient.GetViewWithContext(context, getViewOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetViewWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetViewWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("view_id", viewId); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting view_id: %s", err))
	}
	if err = d.Set("instance_id", instanceId); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting instance_id: %s", err))
	}
	if err = d.Set("region", region); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting region: %s", err))
	}
	if err = d.Set("name", view.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if !core.IsNil(view.SearchQuery) {
		searchQueryMap, err := ResourceIbmLogsViewApisViewsV1SearchQueryToMap(view.SearchQuery)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("search_query", []map[string]interface{}{searchQueryMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting search_query: %s", err))
		}
	}
	timeSelectionMap, err := ResourceIbmLogsViewApisViewsV1TimeSelectionToMap(view.TimeSelection)
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("time_selection", []map[string]interface{}{timeSelectionMap}); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting time_selection: %s", err))
	}
	if !core.IsNil(view.Filters) {
		filtersMap, err := ResourceIbmLogsViewApisViewsV1SelectedFiltersToMap(view.Filters)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("filters", []map[string]interface{}{filtersMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting filters: %s", err))
		}
	}
	if !core.IsNil(view.FolderID) {
		if err = d.Set("folder_id", view.FolderID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting folder_id: %s", err))
		}
	}
	if !core.IsNil(view.Tier) {
		if err = d.Set("tier", view.Tier); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting tier: %s", err))
		}
	}

	return nil
}

func resourceIbmLogsViewUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_logs_view", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	logsClient, _, _, viewId, err := updateClientURLWithInstanceEndpoint(d.Id(), meta, logsClient, d)
	if err != nil {
		return diag.FromErr(err)
	}
	replaceViewOptions := &logsv0.ReplaceViewOptions{}

	viewIdInt, _ := strconv.ParseInt(viewId, 10, 64)
	replaceViewOptions.SetID(viewIdInt)

	hasChange := false

	if d.HasChange("name") || d.HasChange("time_selection") || d.HasChange("search_query") || d.HasChange("filters") || d.HasChange("folder_id") || d.HasChange("tier") {
		// Replace operation requires all fields to be sent
		replaceViewOptions.SetName(d.Get("name").(string))

		timeSelection, err := ResourceIbmLogsViewMapToApisViewsV1TimeSelection(d.Get("time_selection.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		replaceViewOptions.SetTimeSelection(timeSelection)

		searchQuery, err := ResourceIbmLogsViewMapToApisViewsV1SearchQuery(d.Get("search_query.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		replaceViewOptions.SetSearchQuery(searchQuery)

		if _, ok := d.GetOk("filters"); ok {
			filters, err := ResourceIbmLogsViewMapToApisViewsV1SelectedFilters(d.Get("filters.0").(map[string]interface{}))
			if err != nil {
				return diag.FromErr(err)
			}
			replaceViewOptions.SetFilters(filters)
		}

		if _, ok := d.GetOk("folder_id"); ok {
			replaceViewOptions.SetFolderID(core.UUIDPtr(strfmt.UUID(d.Get("folder_id").(string))))
		}

		if _, ok := d.GetOk("tier"); ok {
			replaceViewOptions.SetTier(d.Get("tier").(string))
		}

		hasChange = true
	}

	if hasChange {
		_, response, err := logsClient.ReplaceViewWithContext(context, replaceViewOptions)
		if err != nil {
			log.Printf("[DEBUG] ReplaceViewWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("ReplaceViewWithContext failed %s\n%s", err, response))
		}
	}

	return resourceIbmLogsViewRead(context, d, meta)
}

func resourceIbmLogsViewDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_logs_view", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	logsClient, _, _, viewId, err := updateClientURLWithInstanceEndpoint(d.Id(), meta, logsClient, d)
	if err != nil {
		return diag.FromErr(err)
	}

	deleteViewOptions := &logsv0.DeleteViewOptions{}

	viewIdInt, _ := strconv.ParseInt(viewId, 10, 64)
	deleteViewOptions.SetID(viewIdInt)

	response, err := logsClient.DeleteViewWithContext(context, deleteViewOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteViewWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteViewWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func ResourceIbmLogsViewMapToApisViewsV1TimeSelection(modelMap map[string]interface{}) (logsv0.ApisViewsV1TimeSelectionIntf, error) {
	model := &logsv0.ApisViewsV1TimeSelection{}
	if modelMap["quick_selection"] != nil && len(modelMap["quick_selection"].([]interface{})) > 0 {
		QuickSelectionModel, err := ResourceIbmLogsViewMapToApisViewsV1QuickTimeSelection(modelMap["quick_selection"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.QuickSelection = QuickSelectionModel
	}
	if modelMap["custom_selection"] != nil && len(modelMap["custom_selection"].([]interface{})) > 0 {
		CustomSelectionModel, err := ResourceIbmLogsViewMapToApisViewsV1CustomTimeSelection(modelMap["custom_selection"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.CustomSelection = CustomSelectionModel
	}
	return model, nil
}

func ResourceIbmLogsViewMapToApisViewsV1QuickTimeSelection(modelMap map[string]interface{}) (*logsv0.ApisViewsV1QuickTimeSelection, error) {
	model := &logsv0.ApisViewsV1QuickTimeSelection{}
	model.Caption = core.StringPtr(modelMap["caption"].(string))
	model.Seconds = core.Int64Ptr(int64(modelMap["seconds"].(int)))
	return model, nil
}

func ResourceIbmLogsViewMapToApisViewsV1CustomTimeSelection(modelMap map[string]interface{}) (*logsv0.ApisViewsV1CustomTimeSelection, error) {
	model := &logsv0.ApisViewsV1CustomTimeSelection{}
	fromDateTime, err := core.ParseDateTime(modelMap["from_time"].(string))
	if err != nil {
		return model, err
	}
	model.FromTime = &fromDateTime
	toDateTime, err := core.ParseDateTime(modelMap["to_time"].(string))
	if err != nil {
		return model, err
	}
	model.ToTime = &toDateTime
	return model, nil
}

func ResourceIbmLogsViewMapToApisViewsV1TimeSelectionSelectionTypeQuickSelection(modelMap map[string]interface{}) (*logsv0.ApisViewsV1TimeSelectionSelectionTypeQuickSelection, error) {
	model := &logsv0.ApisViewsV1TimeSelectionSelectionTypeQuickSelection{}
	if modelMap["quick_selection"] != nil && len(modelMap["quick_selection"].([]interface{})) > 0 {
		QuickSelectionModel, err := ResourceIbmLogsViewMapToApisViewsV1QuickTimeSelection(modelMap["quick_selection"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.QuickSelection = QuickSelectionModel
	}
	return model, nil
}

func ResourceIbmLogsViewMapToApisViewsV1TimeSelectionSelectionTypeCustomSelection(modelMap map[string]interface{}) (*logsv0.ApisViewsV1TimeSelectionSelectionTypeCustomSelection, error) {
	model := &logsv0.ApisViewsV1TimeSelectionSelectionTypeCustomSelection{}
	if modelMap["custom_selection"] != nil && len(modelMap["custom_selection"].([]interface{})) > 0 {
		CustomSelectionModel, err := ResourceIbmLogsViewMapToApisViewsV1CustomTimeSelection(modelMap["custom_selection"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.CustomSelection = CustomSelectionModel
	}
	return model, nil
}

func ResourceIbmLogsViewMapToApisViewsV1SearchQuery(modelMap map[string]interface{}) (*logsv0.ApisViewsV1SearchQuery, error) {
	model := &logsv0.ApisViewsV1SearchQuery{}
	model.Query = core.StringPtr(modelMap["query"].(string))
	if modelMap["syntax_type"] != nil && modelMap["syntax_type"].(string) != "" {
		model.SyntaxType = core.StringPtr(modelMap["syntax_type"].(string))
	}
	return model, nil
}

func ResourceIbmLogsViewMapToApisViewsV1SelectedFilters(modelMap map[string]interface{}) (*logsv0.ApisViewsV1SelectedFilters, error) {
	model := &logsv0.ApisViewsV1SelectedFilters{}
	if modelMap["filters"] != nil {
		filters := []logsv0.ApisViewsV1Filter{}
		for _, filtersItem := range modelMap["filters"].([]interface{}) {
			filtersItemModel, err := ResourceIbmLogsViewMapToApisViewsV1Filter(filtersItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			filters = append(filters, *filtersItemModel)
		}
		model.Filters = filters
	}
	return model, nil
}

func ResourceIbmLogsViewMapToApisViewsV1Filter(modelMap map[string]interface{}) (*logsv0.ApisViewsV1Filter, error) {
	model := &logsv0.ApisViewsV1Filter{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["selected_values"] != nil {
		SelectedValuesMap := make(map[string]bool)
		selectedValues := modelMap["selected_values"].(map[string]interface{})
		for key, value := range selectedValues {
			SelectedValuesMap[key] = value.(bool)
		}
		model.SelectedValues = SelectedValuesMap
	}

	// TODO: handle SelectedValues, map with entry type ''
	return model, nil
}

func ResourceIbmLogsViewApisViewsV1SearchQueryToMap(model *logsv0.ApisViewsV1SearchQuery) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["query"] = *model.Query
	if model.SyntaxType != nil {
		modelMap["syntax_type"] = model.SyntaxType
	}
	return modelMap, nil
}

func ResourceIbmLogsViewApisViewsV1TimeSelectionToMap(model logsv0.ApisViewsV1TimeSelectionIntf) (map[string]interface{}, error) {
	if _, ok := model.(*logsv0.ApisViewsV1TimeSelectionSelectionTypeQuickSelection); ok {
		return ResourceIbmLogsViewApisViewsV1TimeSelectionSelectionTypeQuickSelectionToMap(model.(*logsv0.ApisViewsV1TimeSelectionSelectionTypeQuickSelection))
	} else if _, ok := model.(*logsv0.ApisViewsV1TimeSelectionSelectionTypeCustomSelection); ok {
		return ResourceIbmLogsViewApisViewsV1TimeSelectionSelectionTypeCustomSelectionToMap(model.(*logsv0.ApisViewsV1TimeSelectionSelectionTypeCustomSelection))
	} else if _, ok := model.(*logsv0.ApisViewsV1TimeSelection); ok {
		modelMap := make(map[string]interface{})
		model := model.(*logsv0.ApisViewsV1TimeSelection)
		if model.QuickSelection != nil {
			quickSelectionMap, err := ResourceIbmLogsViewApisViewsV1QuickTimeSelectionToMap(model.QuickSelection)
			if err != nil {
				return modelMap, err
			}
			modelMap["quick_selection"] = []map[string]interface{}{quickSelectionMap}
		}
		if model.CustomSelection != nil {
			customSelectionMap, err := ResourceIbmLogsViewApisViewsV1CustomTimeSelectionToMap(model.CustomSelection)
			if err != nil {
				return modelMap, err
			}
			modelMap["custom_selection"] = []map[string]interface{}{customSelectionMap}
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized logsv0.ApisViewsV1TimeSelectionIntf subtype encountered")
	}
}

func ResourceIbmLogsViewApisViewsV1QuickTimeSelectionToMap(model *logsv0.ApisViewsV1QuickTimeSelection) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["caption"] = *model.Caption
	modelMap["seconds"] = flex.IntValue(model.Seconds)
	return modelMap, nil
}

func ResourceIbmLogsViewApisViewsV1CustomTimeSelectionToMap(model *logsv0.ApisViewsV1CustomTimeSelection) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["from_time"] = model.FromTime.String()
	modelMap["to_time"] = model.ToTime.String()
	return modelMap, nil
}

func ResourceIbmLogsViewApisViewsV1TimeSelectionSelectionTypeQuickSelectionToMap(model *logsv0.ApisViewsV1TimeSelectionSelectionTypeQuickSelection) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.QuickSelection != nil {
		quickSelectionMap, err := ResourceIbmLogsViewApisViewsV1QuickTimeSelectionToMap(model.QuickSelection)
		if err != nil {
			return modelMap, err
		}
		modelMap["quick_selection"] = []map[string]interface{}{quickSelectionMap}
	}
	return modelMap, nil
}

func ResourceIbmLogsViewApisViewsV1TimeSelectionSelectionTypeCustomSelectionToMap(model *logsv0.ApisViewsV1TimeSelectionSelectionTypeCustomSelection) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CustomSelection != nil {
		customSelectionMap, err := ResourceIbmLogsViewApisViewsV1CustomTimeSelectionToMap(model.CustomSelection)
		if err != nil {
			return modelMap, err
		}
		modelMap["custom_selection"] = []map[string]interface{}{customSelectionMap}
	}
	return modelMap, nil
}

func ResourceIbmLogsViewApisViewsV1SelectedFiltersToMap(model *logsv0.ApisViewsV1SelectedFilters) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Filters != nil {
		filters := []map[string]interface{}{}
		for _, filtersItem := range model.Filters {
			filtersItemMap, err := ResourceIbmLogsViewApisViewsV1FilterToMap(&filtersItem)
			if err != nil {
				return modelMap, err
			}
			filters = append(filters, filtersItemMap)
		}
		modelMap["filters"] = filters
	}
	return modelMap, nil
}

func ResourceIbmLogsViewApisViewsV1FilterToMap(model *logsv0.ApisViewsV1Filter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.SelectedValues != nil {
		selectedValues := make(map[string]interface{})
		for k, v := range model.SelectedValues {
			selectedValues[k] = v
		}
		modelMap["selected_values"] = selectedValues
	}
	return modelMap, nil
}

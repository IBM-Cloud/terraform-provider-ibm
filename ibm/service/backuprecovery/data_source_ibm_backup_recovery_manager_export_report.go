// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.1-067d600b-20250616-154447
 */

package backuprecovery

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmBackupRecoveryManagerExportReport() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryManagerExportReportRead,

		Schema: map[string]*schema.Schema{
			"report_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the id of the report.",
			},
			"file_path": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the absolute path for download",
			},
			"async": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies if the report should be generated asynchronously.",
			},
			"filters": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies list of global filters that are applicable to given components in the report.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attribute": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the attribute.",
						},
						"filter_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the type of the filter that needs to be applied.",
						},
						"in_filter_params": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the in filter that are applied on attributes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"attribute_data_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the data type of the attribute.",
									},
									"attribute_labels": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the optional label values for the attribute.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"bool_filter_values": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies list of boolean values to filter results on.",
										Elem: &schema.Schema{
											Type: schema.TypeBool,
										},
									},
									"int32_filter_values": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies list of int32 values to filter results on.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"int64_filter_values": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies list of int64 values to filter results on.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"string_filter_values": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies list of string values to filter results on.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"range_filter_params": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the filters that are applied on attributes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"lower_bound": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the lower bound value. If specified, all the results which are greater than this value will be returned.",
									},
									"upper_bound": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the upper bound value. If specified, all the results which are lesser than this value will be returned.",
									},
								},
							},
						},
						"systems_filter_params": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the systems filter. Specifying this will pre filter all the results provided list of system identifier before applying aggregations.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"system_ids": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies an array of system identifiers. System identifiers may be of format clusterid:clusterincarnationid or a regionid (applicable only in case of DMaaS).",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"system_names": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the optional system names labels.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"time_range_filter_params": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the time range filter. Specifying this will pre filter all the results on necessary resources like Protection Runs etc before applying aggregations. Currently, maximum allowed time range is 60 days.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"date_range": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Enum value for specifying the date range for a time filter. Considered only if lowerBound and upperBound are empty.",
									},
									"duration_hours": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the duration preceding the current time for which the data must be fetch i.e fetch data between currentTime and currentTime - durationHours. This filter is only considered if neither upperBound, lowerBound or dateRange is specified.",
									},
									"lower_bound": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the lower bound value. If specified, all the results which are greater than this value will be returned.",
									},
									"upper_bound": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the upper bound value. If specified, all the results which are lesser than this value will be returned.",
									},
								},
							},
						},
					},
				},
			},
			"layout": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The layout of the report which needs to be exported.",
			},
			"report_format": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The format in which the report needs to be exported.",
			},
			"timezone": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies timezone of the user. If nil, defaults to UTC. The time specified should be a location name in the IANA Time Zone database, for example, 'America/Los_Angeles'.",
			},
		},
	}
}

func dataSourceIbmBackupRecoveryManagerExportReportRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryManagerV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_export_report", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	exportReportOptions := &backuprecoveryv1.ExportReportOptions{}

	exportReportOptions.SetID(d.Get("report_id").(string))
	if _, ok := d.GetOk("async"); ok {
		exportReportOptions.SetAsync(d.Get("async").(bool))
	}
	if _, ok := d.GetOk("filters"); ok {
		var newFilters []backuprecoveryv1.AttributeFilter
		for _, v := range d.Get("filters").([]interface{}) {
			value := v.(map[string]interface{})
			newFiltersItem, err := ResourceIbmBackupRecoveryManagerExportReportMapToAttributeFilter(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_manager_export_report", "read", "parse-filters").GetDiag()
			}
			newFilters = append(newFilters, *newFiltersItem)
		}
		exportReportOptions.SetFilters(newFilters)
	}
	if _, ok := d.GetOk("layout"); ok {
		exportReportOptions.SetLayout(d.Get("layout").(string))
	}
	if _, ok := d.GetOk("report_format"); ok {
		exportReportOptions.SetReportFormat(d.Get("report_format").(string))
	}
	if _, ok := d.GetOk("timezone"); ok {
		exportReportOptions.SetTimezone(d.Get("timezone").(string))
	}

	typeString, _, err := backupRecoveryClient.ExportReportWithContext(context, exportReportOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ExportReportWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_manager_export_report", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryManagerExportReportID(d))

	err = saveReportExportToFile(typeString, d.Get("file_path").(string))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_download_agent", "read", "parse-linux_params").GetDiag()
	}

	return nil
}

func saveReportExportToFile(response io.ReadCloser, filePath string) error {
	outputFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, response)
	if err != nil {
		return err
	}

	err = response.Close()
	if err != nil {
		return err
	}

	return nil
}

// dataSourceIbmBackupRecoveryManagerExportReportID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryManagerExportReportID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func ResourceIbmBackupRecoveryManagerExportReportMapToAttributeFilter(modelMap map[string]interface{}) (*backuprecoveryv1.AttributeFilter, error) {
	model := &backuprecoveryv1.AttributeFilter{}
	model.Attribute = core.StringPtr(modelMap["attribute"].(string))
	model.FilterType = core.StringPtr(modelMap["filter_type"].(string))
	if modelMap["in_filter_params"] != nil && len(modelMap["in_filter_params"].([]interface{})) > 0 {
		InFilterParamsModel, err := ResourceIbmBackupRecoveryManagerExportReportMapToInFilterParams(modelMap["in_filter_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.InFilterParams = InFilterParamsModel
	}
	if modelMap["range_filter_params"] != nil && len(modelMap["range_filter_params"].([]interface{})) > 0 {
		RangeFilterParamsModel, err := ResourceIbmBackupRecoveryManagerExportReportMapToRangeFilterParams(modelMap["range_filter_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.RangeFilterParams = RangeFilterParamsModel
	}
	if modelMap["systems_filter_params"] != nil && len(modelMap["systems_filter_params"].([]interface{})) > 0 {
		SystemsFilterParamsModel, err := ResourceIbmBackupRecoveryManagerExportReportMapToSystemsFilterParams(modelMap["systems_filter_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.SystemsFilterParams = SystemsFilterParamsModel
	}
	if modelMap["time_range_filter_params"] != nil && len(modelMap["time_range_filter_params"].([]interface{})) > 0 {
		TimeRangeFilterParamsModel, err := ResourceIbmBackupRecoveryManagerExportReportMapToTimeRangeFilterParams(modelMap["time_range_filter_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.TimeRangeFilterParams = TimeRangeFilterParamsModel
	}
	return model, nil
}

func ResourceIbmBackupRecoveryManagerExportReportMapToInFilterParams(modelMap map[string]interface{}) (*backuprecoveryv1.InFilterParams, error) {
	model := &backuprecoveryv1.InFilterParams{}
	model.AttributeDataType = core.StringPtr(modelMap["attribute_data_type"].(string))
	if modelMap["attribute_labels"] != nil {
		attributeLabels := []string{}
		for _, attributeLabelsItem := range modelMap["attribute_labels"].([]interface{}) {
			attributeLabels = append(attributeLabels, attributeLabelsItem.(string))
		}
		model.AttributeLabels = attributeLabels
	}
	if modelMap["bool_filter_values"] != nil {
		boolFilterValues := []bool{}
		for _, boolFilterValuesItem := range modelMap["bool_filter_values"].([]interface{}) {
			boolFilterValues = append(boolFilterValues, boolFilterValuesItem.(bool))
		}
		model.BoolFilterValues = boolFilterValues
	}
	if modelMap["int32_filter_values"] != nil {
		int32FilterValues := []int64{}
		for _, int32FilterValuesItem := range modelMap["int32_filter_values"].([]interface{}) {
			int32FilterValues = append(int32FilterValues, int64(int32FilterValuesItem.(int)))
		}
		model.Int32FilterValues = int32FilterValues
	}
	if modelMap["int64_filter_values"] != nil {
		int64FilterValues := []int64{}
		for _, int64FilterValuesItem := range modelMap["int64_filter_values"].([]interface{}) {
			int64FilterValues = append(int64FilterValues, int64(int64FilterValuesItem.(int)))
		}
		model.Int64FilterValues = int64FilterValues
	}
	if modelMap["string_filter_values"] != nil {
		stringFilterValues := []string{}
		for _, stringFilterValuesItem := range modelMap["string_filter_values"].([]interface{}) {
			stringFilterValues = append(stringFilterValues, stringFilterValuesItem.(string))
		}
		model.StringFilterValues = stringFilterValues
	}
	return model, nil
}

func ResourceIbmBackupRecoveryManagerExportReportMapToRangeFilterParams(modelMap map[string]interface{}) (*backuprecoveryv1.RangeFilterParams, error) {
	model := &backuprecoveryv1.RangeFilterParams{}
	if modelMap["lower_bound"] != nil {
		model.LowerBound = core.Int64Ptr(int64(modelMap["lower_bound"].(int)))
	}
	if modelMap["upper_bound"] != nil {
		model.UpperBound = core.Int64Ptr(int64(modelMap["upper_bound"].(int)))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryManagerExportReportMapToSystemsFilterParams(modelMap map[string]interface{}) (*backuprecoveryv1.SystemsFilterParams, error) {
	model := &backuprecoveryv1.SystemsFilterParams{}
	systemIds := []string{}
	for _, systemIdsItem := range modelMap["system_ids"].([]interface{}) {
		systemIds = append(systemIds, systemIdsItem.(string))
	}
	model.SystemIds = systemIds
	if modelMap["system_names"] != nil {
		systemNames := []string{}
		for _, systemNamesItem := range modelMap["system_names"].([]interface{}) {
			systemNames = append(systemNames, systemNamesItem.(string))
		}
		model.SystemNames = systemNames
	}
	return model, nil
}

func ResourceIbmBackupRecoveryManagerExportReportMapToTimeRangeFilterParams(modelMap map[string]interface{}) (*backuprecoveryv1.TimeRangeFilterParams, error) {
	model := &backuprecoveryv1.TimeRangeFilterParams{}
	if modelMap["date_range"] != nil && modelMap["date_range"].(string) != "" {
		model.DateRange = core.StringPtr(modelMap["date_range"].(string))
	}
	if modelMap["duration_hours"] != nil {
		model.DurationHours = core.Int64Ptr(int64(modelMap["duration_hours"].(int)))
	}
	if modelMap["lower_bound"] != nil {
		model.LowerBound = core.Int64Ptr(int64(modelMap["lower_bound"].(int)))
	}
	if modelMap["upper_bound"] != nil {
		model.UpperBound = core.Int64Ptr(int64(modelMap["upper_bound"].(int)))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryManagerExportReportAttributeFilterToMap(model *backuprecoveryv1.AttributeFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["attribute"] = *model.Attribute
	modelMap["filter_type"] = *model.FilterType
	if model.InFilterParams != nil {
		inFilterParamsMap, err := ResourceIbmBackupRecoveryManagerExportReportInFilterParamsToMap(model.InFilterParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["in_filter_params"] = []map[string]interface{}{inFilterParamsMap}
	}
	if model.RangeFilterParams != nil {
		rangeFilterParamsMap, err := ResourceIbmBackupRecoveryManagerExportReportRangeFilterParamsToMap(model.RangeFilterParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["range_filter_params"] = []map[string]interface{}{rangeFilterParamsMap}
	}
	if model.SystemsFilterParams != nil {
		systemsFilterParamsMap, err := ResourceIbmBackupRecoveryManagerExportReportSystemsFilterParamsToMap(model.SystemsFilterParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["systems_filter_params"] = []map[string]interface{}{systemsFilterParamsMap}
	}
	if model.TimeRangeFilterParams != nil {
		timeRangeFilterParamsMap, err := ResourceIbmBackupRecoveryManagerExportReportTimeRangeFilterParamsToMap(model.TimeRangeFilterParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["time_range_filter_params"] = []map[string]interface{}{timeRangeFilterParamsMap}
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoveryManagerExportReportInFilterParamsToMap(model *backuprecoveryv1.InFilterParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["attribute_data_type"] = *model.AttributeDataType
	if model.AttributeLabels != nil {
		modelMap["attribute_labels"] = model.AttributeLabels
	}
	if model.BoolFilterValues != nil {
		modelMap["bool_filter_values"] = model.BoolFilterValues
	}
	if model.Int32FilterValues != nil {
		modelMap["int32_filter_values"] = model.Int32FilterValues
	}
	if model.Int64FilterValues != nil {
		modelMap["int64_filter_values"] = model.Int64FilterValues
	}
	if model.StringFilterValues != nil {
		modelMap["string_filter_values"] = model.StringFilterValues
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoveryManagerExportReportRangeFilterParamsToMap(model *backuprecoveryv1.RangeFilterParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LowerBound != nil {
		modelMap["lower_bound"] = flex.IntValue(model.LowerBound)
	}
	if model.UpperBound != nil {
		modelMap["upper_bound"] = flex.IntValue(model.UpperBound)
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoveryManagerExportReportSystemsFilterParamsToMap(model *backuprecoveryv1.SystemsFilterParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["system_ids"] = model.SystemIds
	if model.SystemNames != nil {
		modelMap["system_names"] = model.SystemNames
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoveryManagerExportReportTimeRangeFilterParamsToMap(model *backuprecoveryv1.TimeRangeFilterParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DateRange != nil {
		modelMap["date_range"] = *model.DateRange
	}
	if model.DurationHours != nil {
		modelMap["duration_hours"] = flex.IntValue(model.DurationHours)
	}
	if model.LowerBound != nil {
		modelMap["lower_bound"] = flex.IntValue(model.LowerBound)
	}
	if model.UpperBound != nil {
		modelMap["upper_bound"] = flex.IntValue(model.UpperBound)
	}
	return modelMap, nil
}

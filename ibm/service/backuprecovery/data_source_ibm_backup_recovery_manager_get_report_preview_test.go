// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.1-5136e54a-20241108-203028
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmBackupRecoveryManagerGetReportPreviewDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryManagerGetReportPreviewDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_manager_get_report_preview.backup_recovery_manager_get_report_preview_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_manager_get_report_preview.backup_recovery_manager_get_report_preview_instance", "backup_recovery_manager_get_report_preview_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryManagerGetReportPreviewDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_manager_get_report_preview" "backup_recovery_manager_get_report_preview_instance" {
			id = "id"
			componentIds = [ "componentIds" ]
			filters = [ { attribute="attribute", filter_type="In", in_filter_params={ attribute_data_type="Bool", attribute_labels=[ "attributeLabels" ], bool_filter_values=[ true ], int32_filter_values=[ 1 ], int64_filter_values=[ 1 ], string_filter_values=[ "stringFilterValues" ] }, range_filter_params={ lower_bound=1, upper_bound=1 }, systems_filter_params={ system_ids=[ "systemIds" ], system_names=[ "systemNames" ] }, time_range_filter_params={ date_range="Last1Hour", duration_hours=1, lower_bound=1, upper_bound=1 } } ]
			timezone = "timezone"
		}
	`)
}

func TestDataSourceIbmBackupRecoveryManagerGetReportPreviewComponentToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		aggregatedAttributesParamsModel := make(map[string]interface{})
		aggregatedAttributesParamsModel["aggregation_type"] = "sum"
		aggregatedAttributesParamsModel["attribute"] = "testString"
		aggregatedAttributesParamsModel["label"] = "testString"

		attributeAggregationsModel := make(map[string]interface{})
		attributeAggregationsModel["aggregated_attributes"] = []map[string]interface{}{aggregatedAttributesParamsModel}
		attributeAggregationsModel["grouped_attributes"] = []string{"testString"}

		xlsxAttributeCustomConfigParamsModel := make(map[string]interface{})
		xlsxAttributeCustomConfigParamsModel["attribute_name"] = "testString"
		xlsxAttributeCustomConfigParamsModel["custom_label"] = "testString"
		xlsxAttributeCustomConfigParamsModel["format"] = "Timestamp"

		xlsxCustomConfigParamsModel := make(map[string]interface{})
		xlsxCustomConfigParamsModel["attribute_config"] = []map[string]interface{}{xlsxAttributeCustomConfigParamsModel}

		customConfigParamsModel := make(map[string]interface{})
		customConfigParamsModel["xlsx_params"] = []map[string]interface{}{xlsxCustomConfigParamsModel}

		inFilterParamsModel := make(map[string]interface{})
		inFilterParamsModel["attribute_data_type"] = "Bool"
		inFilterParamsModel["attribute_labels"] = []string{"testString"}
		inFilterParamsModel["bool_filter_values"] = []bool{true}
		inFilterParamsModel["int32_filter_values"] = []int64{int64(38)}
		inFilterParamsModel["int64_filter_values"] = []int64{int64(26)}
		inFilterParamsModel["string_filter_values"] = []string{"testString"}

		rangeFilterParamsModel := make(map[string]interface{})
		rangeFilterParamsModel["lower_bound"] = int(26)
		rangeFilterParamsModel["upper_bound"] = int(26)

		systemsFilterParamsModel := make(map[string]interface{})
		systemsFilterParamsModel["system_ids"] = []string{"testString"}
		systemsFilterParamsModel["system_names"] = []string{"testString"}

		timeRangeFilterParamsModel := make(map[string]interface{})
		timeRangeFilterParamsModel["date_range"] = "Last1Hour"
		timeRangeFilterParamsModel["duration_hours"] = int(26)
		timeRangeFilterParamsModel["lower_bound"] = int(26)
		timeRangeFilterParamsModel["upper_bound"] = int(26)

		attributeFilterModel := make(map[string]interface{})
		attributeFilterModel["attribute"] = "testString"
		attributeFilterModel["filter_type"] = "In"
		attributeFilterModel["in_filter_params"] = []map[string]interface{}{inFilterParamsModel}
		attributeFilterModel["range_filter_params"] = []map[string]interface{}{rangeFilterParamsModel}
		attributeFilterModel["systems_filter_params"] = []map[string]interface{}{systemsFilterParamsModel}
		attributeFilterModel["time_range_filter_params"] = []map[string]interface{}{timeRangeFilterParamsModel}

		limitParamsModel := make(map[string]interface{})
		limitParamsModel["from"] = int(38)
		limitParamsModel["size"] = int(1)

		attributeSortModel := make(map[string]interface{})
		attributeSortModel["attribute"] = "testString"
		attributeSortModel["desc"] = true

		model := make(map[string]interface{})
		model["aggs"] = []map[string]interface{}{attributeAggregationsModel}
		model["config"] = []map[string]interface{}{customConfigParamsModel}
		model["data"] = []map[string]interface{}{map[string]interface{}{"anyKey": "anyValue"}}
		model["description"] = "testString"
		model["filters"] = []map[string]interface{}{attributeFilterModel}
		model["id"] = "testString"
		model["limit"] = []map[string]interface{}{limitParamsModel}
		model["name"] = "testString"
		model["report_type"] = "Failures"
		model["sort"] = []map[string]interface{}{attributeSortModel}

		assert.Equal(t, result, model)
	}

	aggregatedAttributesParamsModel := new(backuprecoveryv1.AggregatedAttributesParams)
	aggregatedAttributesParamsModel.AggregationType = core.StringPtr("sum")
	aggregatedAttributesParamsModel.Attribute = core.StringPtr("testString")
	aggregatedAttributesParamsModel.Label = core.StringPtr("testString")

	attributeAggregationsModel := new(backuprecoveryv1.AttributeAggregations)
	attributeAggregationsModel.AggregatedAttributes = []backuprecoveryv1.AggregatedAttributesParams{*aggregatedAttributesParamsModel}
	attributeAggregationsModel.GroupedAttributes = []string{"testString"}

	xlsxAttributeCustomConfigParamsModel := new(backuprecoveryv1.XlsxAttributeCustomConfigParams)
	xlsxAttributeCustomConfigParamsModel.AttributeName = core.StringPtr("testString")
	xlsxAttributeCustomConfigParamsModel.CustomLabel = core.StringPtr("testString")
	xlsxAttributeCustomConfigParamsModel.Format = core.StringPtr("Timestamp")

	xlsxCustomConfigParamsModel := new(backuprecoveryv1.XlsxCustomConfigParams)
	xlsxCustomConfigParamsModel.AttributeConfig = []backuprecoveryv1.XlsxAttributeCustomConfigParams{*xlsxAttributeCustomConfigParamsModel}

	customConfigParamsModel := new(backuprecoveryv1.CustomConfigParams)
	customConfigParamsModel.XlsxParams = xlsxCustomConfigParamsModel

	inFilterParamsModel := new(backuprecoveryv1.InFilterParams)
	inFilterParamsModel.AttributeDataType = core.StringPtr("Bool")
	inFilterParamsModel.AttributeLabels = []string{"testString"}
	inFilterParamsModel.BoolFilterValues = []bool{true}
	inFilterParamsModel.Int32FilterValues = []int64{int64(38)}
	inFilterParamsModel.Int64FilterValues = []int64{int64(26)}
	inFilterParamsModel.StringFilterValues = []string{"testString"}

	rangeFilterParamsModel := new(backuprecoveryv1.RangeFilterParams)
	rangeFilterParamsModel.LowerBound = core.Int64Ptr(int64(26))
	rangeFilterParamsModel.UpperBound = core.Int64Ptr(int64(26))

	systemsFilterParamsModel := new(backuprecoveryv1.SystemsFilterParams)
	systemsFilterParamsModel.SystemIds = []string{"testString"}
	systemsFilterParamsModel.SystemNames = []string{"testString"}

	timeRangeFilterParamsModel := new(backuprecoveryv1.TimeRangeFilterParams)
	timeRangeFilterParamsModel.DateRange = core.StringPtr("Last1Hour")
	timeRangeFilterParamsModel.DurationHours = core.Int64Ptr(int64(26))
	timeRangeFilterParamsModel.LowerBound = core.Int64Ptr(int64(26))
	timeRangeFilterParamsModel.UpperBound = core.Int64Ptr(int64(26))

	attributeFilterModel := new(backuprecoveryv1.AttributeFilter)
	attributeFilterModel.Attribute = core.StringPtr("testString")
	attributeFilterModel.FilterType = core.StringPtr("In")
	attributeFilterModel.InFilterParams = inFilterParamsModel
	attributeFilterModel.RangeFilterParams = rangeFilterParamsModel
	attributeFilterModel.SystemsFilterParams = systemsFilterParamsModel
	attributeFilterModel.TimeRangeFilterParams = timeRangeFilterParamsModel

	limitParamsModel := new(backuprecoveryv1.LimitParams)
	limitParamsModel.From = core.Int64Ptr(int64(38))
	limitParamsModel.Size = core.Int64Ptr(int64(1))

	attributeSortModel := new(backuprecoveryv1.AttributeSort)
	attributeSortModel.Attribute = core.StringPtr("testString")
	attributeSortModel.Desc = core.BoolPtr(true)

	model := new(backuprecoveryv1.Component)
	model.Aggs = attributeAggregationsModel
	model.Config = customConfigParamsModel
	model.Data = []map[string]interface{}{map[string]interface{}{"anyKey": "anyValue"}}
	model.Description = core.StringPtr("testString")
	model.Filters = []backuprecoveryv1.AttributeFilter{*attributeFilterModel}
	model.ID = core.StringPtr("testString")
	model.Limit = limitParamsModel
	model.Name = core.StringPtr("testString")
	model.ReportType = core.StringPtr("Failures")
	model.Sort = []backuprecoveryv1.AttributeSort{*attributeSortModel}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetReportPreviewComponentToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryManagerGetReportPreviewAttributeAggregationsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		aggregatedAttributesParamsModel := make(map[string]interface{})
		aggregatedAttributesParamsModel["aggregation_type"] = "sum"
		aggregatedAttributesParamsModel["attribute"] = "testString"
		aggregatedAttributesParamsModel["label"] = "testString"

		model := make(map[string]interface{})
		model["aggregated_attributes"] = []map[string]interface{}{aggregatedAttributesParamsModel}
		model["grouped_attributes"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	aggregatedAttributesParamsModel := new(backuprecoveryv1.AggregatedAttributesParams)
	aggregatedAttributesParamsModel.AggregationType = core.StringPtr("sum")
	aggregatedAttributesParamsModel.Attribute = core.StringPtr("testString")
	aggregatedAttributesParamsModel.Label = core.StringPtr("testString")

	model := new(backuprecoveryv1.AttributeAggregations)
	model.AggregatedAttributes = []backuprecoveryv1.AggregatedAttributesParams{*aggregatedAttributesParamsModel}
	model.GroupedAttributes = []string{"testString"}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetReportPreviewAttributeAggregationsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryManagerGetReportPreviewAggregatedAttributesParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["aggregation_type"] = "sum"
		model["attribute"] = "testString"
		model["label"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AggregatedAttributesParams)
	model.AggregationType = core.StringPtr("sum")
	model.Attribute = core.StringPtr("testString")
	model.Label = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetReportPreviewAggregatedAttributesParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryManagerGetReportPreviewCustomConfigParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		xlsxAttributeCustomConfigParamsModel := make(map[string]interface{})
		xlsxAttributeCustomConfigParamsModel["attribute_name"] = "testString"
		xlsxAttributeCustomConfigParamsModel["custom_label"] = "testString"
		xlsxAttributeCustomConfigParamsModel["format"] = "Timestamp"

		xlsxCustomConfigParamsModel := make(map[string]interface{})
		xlsxCustomConfigParamsModel["attribute_config"] = []map[string]interface{}{xlsxAttributeCustomConfigParamsModel}

		model := make(map[string]interface{})
		model["xlsx_params"] = []map[string]interface{}{xlsxCustomConfigParamsModel}

		assert.Equal(t, result, model)
	}

	xlsxAttributeCustomConfigParamsModel := new(backuprecoveryv1.XlsxAttributeCustomConfigParams)
	xlsxAttributeCustomConfigParamsModel.AttributeName = core.StringPtr("testString")
	xlsxAttributeCustomConfigParamsModel.CustomLabel = core.StringPtr("testString")
	xlsxAttributeCustomConfigParamsModel.Format = core.StringPtr("Timestamp")

	xlsxCustomConfigParamsModel := new(backuprecoveryv1.XlsxCustomConfigParams)
	xlsxCustomConfigParamsModel.AttributeConfig = []backuprecoveryv1.XlsxAttributeCustomConfigParams{*xlsxAttributeCustomConfigParamsModel}

	model := new(backuprecoveryv1.CustomConfigParams)
	model.XlsxParams = xlsxCustomConfigParamsModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetReportPreviewCustomConfigParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryManagerGetReportPreviewXlsxCustomConfigParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		xlsxAttributeCustomConfigParamsModel := make(map[string]interface{})
		xlsxAttributeCustomConfigParamsModel["attribute_name"] = "testString"
		xlsxAttributeCustomConfigParamsModel["custom_label"] = "testString"
		xlsxAttributeCustomConfigParamsModel["format"] = "Timestamp"

		model := make(map[string]interface{})
		model["attribute_config"] = []map[string]interface{}{xlsxAttributeCustomConfigParamsModel}

		assert.Equal(t, result, model)
	}

	xlsxAttributeCustomConfigParamsModel := new(backuprecoveryv1.XlsxAttributeCustomConfigParams)
	xlsxAttributeCustomConfigParamsModel.AttributeName = core.StringPtr("testString")
	xlsxAttributeCustomConfigParamsModel.CustomLabel = core.StringPtr("testString")
	xlsxAttributeCustomConfigParamsModel.Format = core.StringPtr("Timestamp")

	model := new(backuprecoveryv1.XlsxCustomConfigParams)
	model.AttributeConfig = []backuprecoveryv1.XlsxAttributeCustomConfigParams{*xlsxAttributeCustomConfigParamsModel}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetReportPreviewXlsxCustomConfigParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryManagerGetReportPreviewXlsxAttributeCustomConfigParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["attribute_name"] = "testString"
		model["custom_label"] = "testString"
		model["format"] = "Timestamp"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.XlsxAttributeCustomConfigParams)
	model.AttributeName = core.StringPtr("testString")
	model.CustomLabel = core.StringPtr("testString")
	model.Format = core.StringPtr("Timestamp")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetReportPreviewXlsxAttributeCustomConfigParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryManagerGetReportPreviewAttributeFilterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		inFilterParamsModel := make(map[string]interface{})
		inFilterParamsModel["attribute_data_type"] = "Bool"
		inFilterParamsModel["attribute_labels"] = []string{"testString"}
		inFilterParamsModel["bool_filter_values"] = []bool{true}
		inFilterParamsModel["int32_filter_values"] = []int64{int64(38)}
		inFilterParamsModel["int64_filter_values"] = []int64{int64(26)}
		inFilterParamsModel["string_filter_values"] = []string{"testString"}

		rangeFilterParamsModel := make(map[string]interface{})
		rangeFilterParamsModel["lower_bound"] = int(26)
		rangeFilterParamsModel["upper_bound"] = int(26)

		systemsFilterParamsModel := make(map[string]interface{})
		systemsFilterParamsModel["system_ids"] = []string{"testString"}
		systemsFilterParamsModel["system_names"] = []string{"testString"}

		timeRangeFilterParamsModel := make(map[string]interface{})
		timeRangeFilterParamsModel["date_range"] = "Last1Hour"
		timeRangeFilterParamsModel["duration_hours"] = int(26)
		timeRangeFilterParamsModel["lower_bound"] = int(26)
		timeRangeFilterParamsModel["upper_bound"] = int(26)

		model := make(map[string]interface{})
		model["attribute"] = "testString"
		model["filter_type"] = "In"
		model["in_filter_params"] = []map[string]interface{}{inFilterParamsModel}
		model["range_filter_params"] = []map[string]interface{}{rangeFilterParamsModel}
		model["systems_filter_params"] = []map[string]interface{}{systemsFilterParamsModel}
		model["time_range_filter_params"] = []map[string]interface{}{timeRangeFilterParamsModel}

		assert.Equal(t, result, model)
	}

	inFilterParamsModel := new(backuprecoveryv1.InFilterParams)
	inFilterParamsModel.AttributeDataType = core.StringPtr("Bool")
	inFilterParamsModel.AttributeLabels = []string{"testString"}
	inFilterParamsModel.BoolFilterValues = []bool{true}
	inFilterParamsModel.Int32FilterValues = []int64{int64(38)}
	inFilterParamsModel.Int64FilterValues = []int64{int64(26)}
	inFilterParamsModel.StringFilterValues = []string{"testString"}

	rangeFilterParamsModel := new(backuprecoveryv1.RangeFilterParams)
	rangeFilterParamsModel.LowerBound = core.Int64Ptr(int64(26))
	rangeFilterParamsModel.UpperBound = core.Int64Ptr(int64(26))

	systemsFilterParamsModel := new(backuprecoveryv1.SystemsFilterParams)
	systemsFilterParamsModel.SystemIds = []string{"testString"}
	systemsFilterParamsModel.SystemNames = []string{"testString"}

	timeRangeFilterParamsModel := new(backuprecoveryv1.TimeRangeFilterParams)
	timeRangeFilterParamsModel.DateRange = core.StringPtr("Last1Hour")
	timeRangeFilterParamsModel.DurationHours = core.Int64Ptr(int64(26))
	timeRangeFilterParamsModel.LowerBound = core.Int64Ptr(int64(26))
	timeRangeFilterParamsModel.UpperBound = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.AttributeFilter)
	model.Attribute = core.StringPtr("testString")
	model.FilterType = core.StringPtr("In")
	model.InFilterParams = inFilterParamsModel
	model.RangeFilterParams = rangeFilterParamsModel
	model.SystemsFilterParams = systemsFilterParamsModel
	model.TimeRangeFilterParams = timeRangeFilterParamsModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetReportPreviewAttributeFilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryManagerGetReportPreviewInFilterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["attribute_data_type"] = "Bool"
		model["attribute_labels"] = []string{"testString"}
		model["bool_filter_values"] = []bool{true}
		model["int32_filter_values"] = []int64{int64(38)}
		model["int64_filter_values"] = []int64{int64(26)}
		model["string_filter_values"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.InFilterParams)
	model.AttributeDataType = core.StringPtr("Bool")
	model.AttributeLabels = []string{"testString"}
	model.BoolFilterValues = []bool{true}
	model.Int32FilterValues = []int64{int64(38)}
	model.Int64FilterValues = []int64{int64(26)}
	model.StringFilterValues = []string{"testString"}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetReportPreviewInFilterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryManagerGetReportPreviewRangeFilterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["lower_bound"] = int(26)
		model["upper_bound"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.RangeFilterParams)
	model.LowerBound = core.Int64Ptr(int64(26))
	model.UpperBound = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetReportPreviewRangeFilterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryManagerGetReportPreviewSystemsFilterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["system_ids"] = []string{"testString"}
		model["system_names"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SystemsFilterParams)
	model.SystemIds = []string{"testString"}
	model.SystemNames = []string{"testString"}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetReportPreviewSystemsFilterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryManagerGetReportPreviewTimeRangeFilterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["date_range"] = "Last1Hour"
		model["duration_hours"] = int(26)
		model["lower_bound"] = int(26)
		model["upper_bound"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.TimeRangeFilterParams)
	model.DateRange = core.StringPtr("Last1Hour")
	model.DurationHours = core.Int64Ptr(int64(26))
	model.LowerBound = core.Int64Ptr(int64(26))
	model.UpperBound = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetReportPreviewTimeRangeFilterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryManagerGetReportPreviewLimitParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["from"] = int(38)
		model["size"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.LimitParams)
	model.From = core.Int64Ptr(int64(38))
	model.Size = core.Int64Ptr(int64(1))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetReportPreviewLimitParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryManagerGetReportPreviewAttributeSortToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["attribute"] = "testString"
		model["desc"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AttributeSort)
	model.Attribute = core.StringPtr("testString")
	model.Desc = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetReportPreviewAttributeSortToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

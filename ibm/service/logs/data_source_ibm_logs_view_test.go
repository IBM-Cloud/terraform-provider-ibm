// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsViewDataSourceBasic(t *testing.T) {
	viewName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsViewDataSourceConfigBasic(viewName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "logs_view_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "time_selection.#"),
				),
			},
		},
	})
}

func TestAccIbmLogsViewDataSourceAllArgs(t *testing.T) {
	viewName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	viewTier := "priority_insights"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsViewDataSourceConfig(viewName, viewTier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "logs_view_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "search_query.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "time_selection.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "filters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "tier"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsViewDataSourceConfigBasic(viewName string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_view" "logs_view_instance" {
			instance_id = "%s"
			region      = "%s"
			name        = "%s"
			time_selection {
				quick_selection {
					caption = "Last 1 hour"
					seconds = 3600
				}
			}
			tier = "priority_insights"
		}

		data "ibm_logs_view" "logs_view_instance" {
			instance_id = ibm_logs_view.logs_view_instance.instance_id
			region      = ibm_logs_view.logs_view_instance.region
			logs_view_id = ibm_logs_view.logs_view_instance.view_id
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, viewName)
}

func testAccCheckIbmLogsViewDataSourceConfig(viewName string, viewTier string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_view" "logs_view_instance" {
			instance_id = "%s"
			region      = "%s"
			name        = "%s"
			search_query {
				query = "error"
				syntax_type = "dataprime"
			}
			time_selection {
				quick_selection {
					caption = "Last 1 hour"
					seconds = 3600
				}
			}
			filters {
				filters {
					name = "applicationName"
					selected_values = {"cs-rest-test1":true,"demo":true}
				}
			}
			tier = "%s"
		}

		data "ibm_logs_view" "logs_view_instance" {
			instance_id = ibm_logs_view.logs_view_instance.instance_id
			region      = ibm_logs_view.logs_view_instance.region
			logs_view_id = ibm_logs_view.logs_view_instance.view_id
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, viewName, viewTier)
}

// Todo @kavya498: Fix unit testcases

// func TestDataSourceIbmLogsViewApisViewsV1SearchQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["query"] = "error"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisViewsV1SearchQuery)
// 	model.Query = core.StringPtr("error")

// 	result, err := logs.DataSourceIbmLogsViewApisViewsV1SearchQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsViewApisViewsV1TimeSelectionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisViewsV1QuickTimeSelectionModel := make(map[string]interface{})
// 		apisViewsV1QuickTimeSelectionModel["caption"] = "Last hour"
// 		apisViewsV1QuickTimeSelectionModel["seconds"] = int(3600)

// 		model := make(map[string]interface{})
// 		model["quick_selection"] = []map[string]interface{}{apisViewsV1QuickTimeSelectionModel}
// 		model["custom_selection"] = []map[string]interface{}{apisViewsV1CustomTimeSelectionModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisViewsV1QuickTimeSelectionModel := new(logsv0.ApisViewsV1QuickTimeSelection)
// 	apisViewsV1QuickTimeSelectionModel.Caption = core.StringPtr("Last hour")
// 	apisViewsV1QuickTimeSelectionModel.Seconds = core.Int64Ptr(int64(3600))

// 	model := new(logsv0.ApisViewsV1TimeSelection)
// 	model.QuickSelection = apisViewsV1QuickTimeSelectionModel
// 	model.CustomSelection = apisViewsV1CustomTimeSelectionModel

// 	result, err := logs.DataSourceIbmLogsViewApisViewsV1TimeSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsViewApisViewsV1QuickTimeSelectionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["caption"] = "Last hour"
// 		model["seconds"] = int(3600)

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisViewsV1QuickTimeSelection)
// 	model.Caption = core.StringPtr("Last hour")
// 	model.Seconds = core.Int64Ptr(int64(3600))

// 	result, err := logs.DataSourceIbmLogsViewApisViewsV1QuickTimeSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsViewApisViewsV1CustomTimeSelectionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["from_time"] = "2024-01-25T11:31:43.152Z"
// 		model["to_time"] = "2024-01-25T11:35:43.152Z"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisViewsV1CustomTimeSelection)
// 	model.FromTime = CreateMockDateTime("2024-01-25T11:31:43.152Z")
// 	model.ToTime = CreateMockDateTime("2024-01-25T11:35:43.152Z")

// 	result, err := logs.DataSourceIbmLogsViewApisViewsV1CustomTimeSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsViewApisViewsV1TimeSelectionSelectionTypeQuickSelectionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisViewsV1QuickTimeSelectionModel := make(map[string]interface{})
// 		apisViewsV1QuickTimeSelectionModel["caption"] = "Last hour"
// 		apisViewsV1QuickTimeSelectionModel["seconds"] = int(3600)

// 		model := make(map[string]interface{})
// 		model["quick_selection"] = []map[string]interface{}{apisViewsV1QuickTimeSelectionModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisViewsV1QuickTimeSelectionModel := new(logsv0.ApisViewsV1QuickTimeSelection)
// 	apisViewsV1QuickTimeSelectionModel.Caption = core.StringPtr("Last hour")
// 	apisViewsV1QuickTimeSelectionModel.Seconds = core.Int64Ptr(int64(3600))

// 	model := new(logsv0.ApisViewsV1TimeSelectionSelectionTypeQuickSelection)
// 	model.QuickSelection = apisViewsV1QuickTimeSelectionModel

// 	result, err := logs.DataSourceIbmLogsViewApisViewsV1TimeSelectionSelectionTypeQuickSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsViewApisViewsV1TimeSelectionSelectionTypeCustomSelectionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisViewsV1CustomTimeSelectionModel := make(map[string]interface{})
// 		apisViewsV1CustomTimeSelectionModel["from_time"] = "2024-01-25T11:31:43.152Z"
// 		apisViewsV1CustomTimeSelectionModel["to_time"] = "2024-01-25T11:35:43.152Z"

// 		model := make(map[string]interface{})
// 		model["custom_selection"] = []map[string]interface{}{apisViewsV1CustomTimeSelectionModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisViewsV1CustomTimeSelectionModel := new(logsv0.ApisViewsV1CustomTimeSelection)
// 	apisViewsV1CustomTimeSelectionModel.FromTime = CreateMockDateTime("2024-01-25T11:31:43.152Z")
// 	apisViewsV1CustomTimeSelectionModel.ToTime = CreateMockDateTime("2024-01-25T11:35:43.152Z")

// 	model := new(logsv0.ApisViewsV1TimeSelectionSelectionTypeCustomSelection)
// 	model.CustomSelection = apisViewsV1CustomTimeSelectionModel

// 	result, err := logs.DataSourceIbmLogsViewApisViewsV1TimeSelectionSelectionTypeCustomSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsViewApisViewsV1SelectedFiltersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisViewsV1FilterModel := make(map[string]interface{})
// 		apisViewsV1FilterModel["name"] = "applicationName"
// 		apisViewsV1FilterModel["selected_values"] = map[string]interface{}{"key1": fmt.Sprintf("%v", true)}

// 		model := make(map[string]interface{})
// 		model["filters"] = []map[string]interface{}{apisViewsV1FilterModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisViewsV1FilterModel := new(logsv0.ApisViewsV1Filter)
// 	apisViewsV1FilterModel.Name = core.StringPtr("applicationName")
// 	apisViewsV1FilterModel.SelectedValues = map[string]bool{"key1": true}

// 	model := new(logsv0.ApisViewsV1SelectedFilters)
// 	model.Filters = []logsv0.ApisViewsV1Filter{*apisViewsV1FilterModel}

// 	result, err := logs.DataSourceIbmLogsViewApisViewsV1SelectedFiltersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsViewApisViewsV1FilterToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["name"] = "applicationName"
// 		model["selected_values"] = map[string]interface{}{"key1": fmt.Sprintf("%v", true)}

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisViewsV1Filter)
// 	model.Name = core.StringPtr("applicationName")
// 	model.SelectedValues = map[string]bool{"key1": true}

// 	result, err := logs.DataSourceIbmLogsViewApisViewsV1FilterToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

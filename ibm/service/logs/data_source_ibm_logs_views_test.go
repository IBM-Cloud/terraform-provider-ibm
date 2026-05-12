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

func TestAccIbmLogsViewsDataSourceBasic(t *testing.T) {
	viewName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 1000))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsViewsDataSourceConfigBasic(viewName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_views.logs_views_instance", "id"),
				),
			},
		},
	})
}

func TestAccIbmLogsViewsDataSourceAllArgs(t *testing.T) {
	viewName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 1000))
	viewTier := "priority_insights"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsViewsDataSourceConfig(viewName, viewTier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_views.logs_views_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_views.logs_views_instance", "views.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_views.logs_views_instance", "views.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_views.logs_views_instance", "views.0.name"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsViewsDataSourceConfigBasic(viewName string) string {
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

		data "ibm_logs_views" "logs_views_instance" {
			instance_id = ibm_logs_view.logs_view_instance.instance_id
			region      = ibm_logs_view.logs_view_instance.region
			depends_on = [
				ibm_logs_view.logs_view_instance
			]
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, viewName)
}

func testAccCheckIbmLogsViewsDataSourceConfig(viewName string, viewTier string) string {
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

		data "ibm_logs_views" "logs_views_instance" {
			instance_id = ibm_logs_view.logs_view_instance.instance_id
			region      = ibm_logs_view.logs_view_instance.region
			depends_on = [
				ibm_logs_view.logs_view_instance
			]
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, viewName, viewTier)
}

// Todo @kavya498: Fix unit testcases
// func TestDataSourceIbmLogsViewsViewToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisViewsV1SearchQueryModel := make(map[string]interface{})
// 		apisViewsV1SearchQueryModel["query"] = "logs"

// 		apisViewsV1CustomTimeSelectionModel := make(map[string]interface{})
// 		apisViewsV1CustomTimeSelectionModel["from_time"] = "2024-01-25T11:31:43.152Z"
// 		apisViewsV1CustomTimeSelectionModel["to_time"] = "2024-01-25T11:37:13.238Z"

// 		apisViewsV1TimeSelectionModel := make(map[string]interface{})
// 		apisViewsV1TimeSelectionModel["custom_selection"] = []map[string]interface{}{apisViewsV1CustomTimeSelectionModel}

// 		apisViewsV1FilterModel := make(map[string]interface{})
// 		apisViewsV1FilterModel["name"] = "applicationName"
// 		apisViewsV1FilterModel["selected_values"] = map[string]interface{}{"key1": fmt.Sprintf("%v", true)}

// 		apisViewsV1SelectedFiltersModel := make(map[string]interface{})
// 		apisViewsV1SelectedFiltersModel["filters"] = []map[string]interface{}{apisViewsV1FilterModel}

// 		model := make(map[string]interface{})
// 		model["id"] = int(38)
// 		model["name"] = "testString"
// 		model["search_query"] = []map[string]interface{}{apisViewsV1SearchQueryModel}
// 		model["time_selection"] = []map[string]interface{}{apisViewsV1TimeSelectionModel}
// 		model["filters"] = []map[string]interface{}{apisViewsV1SelectedFiltersModel}
// 		model["folder_id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"

// 		assert.Equal(t, result, model)
// 	}

// 	apisViewsV1SearchQueryModel := new(logsv0.ApisViewsV1SearchQuery)
// 	apisViewsV1SearchQueryModel.Query = core.StringPtr("logs")

// 	apisViewsV1CustomTimeSelectionModel := new(logsv0.ApisViewsV1CustomTimeSelection)
// 	apisViewsV1CustomTimeSelectionModel.FromTime = CreateMockDateTime("2024-01-25T11:31:43.152Z")
// 	apisViewsV1CustomTimeSelectionModel.ToTime = CreateMockDateTime("2024-01-25T11:37:13.238Z")

// 	apisViewsV1TimeSelectionModel := new(logsv0.ApisViewsV1TimeSelectionSelectionTypeCustomSelection)
// 	apisViewsV1TimeSelectionModel.CustomSelection = apisViewsV1CustomTimeSelectionModel

// 	apisViewsV1FilterModel := new(logsv0.ApisViewsV1Filter)
// 	apisViewsV1FilterModel.Name = core.StringPtr("applicationName")
// 	apisViewsV1FilterModel.SelectedValues = map[string]bool{"key1": true}

// 	apisViewsV1SelectedFiltersModel := new(logsv0.ApisViewsV1SelectedFilters)
// 	apisViewsV1SelectedFiltersModel.Filters = []logsv0.ApisViewsV1Filter{*apisViewsV1FilterModel}

// 	model := new(logsv0.View)
// 	model.ID = core.Int64Ptr(int64(38))
// 	model.Name = core.StringPtr("testString")
// 	model.SearchQuery = apisViewsV1SearchQueryModel
// 	model.TimeSelection = apisViewsV1TimeSelectionModel
// 	model.Filters = apisViewsV1SelectedFiltersModel
// 	model.FolderID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")

// 	result, err := logs.DataSourceIbmLogsViewsViewToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsViewsApisViewsV1SearchQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["query"] = "error"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisViewsV1SearchQuery)
// 	model.Query = core.StringPtr("error")

// 	result, err := logs.DataSourceIbmLogsViewsApisViewsV1SearchQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsViewsApisViewsV1TimeSelectionToMap(t *testing.T) {
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

// 	result, err := logs.DataSourceIbmLogsViewsApisViewsV1TimeSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsViewsApisViewsV1QuickTimeSelectionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["caption"] = "Last hour"
// 		model["seconds"] = int(3600)

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisViewsV1QuickTimeSelection)
// 	model.Caption = core.StringPtr("Last hour")
// 	model.Seconds = core.Int64Ptr(int64(3600))

// 	result, err := logs.DataSourceIbmLogsViewsApisViewsV1QuickTimeSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsViewsApisViewsV1CustomTimeSelectionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["from_time"] = "2024-01-25T11:31:43.152Z"
// 		model["to_time"] = "2024-01-25T11:35:43.152Z"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisViewsV1CustomTimeSelection)
// 	model.FromTime = CreateMockDateTime("2024-01-25T11:31:43.152Z")
// 	model.ToTime = CreateMockDateTime("2024-01-25T11:35:43.152Z")

// 	result, err := logs.DataSourceIbmLogsViewsApisViewsV1CustomTimeSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsViewsApisViewsV1TimeSelectionSelectionTypeQuickSelectionToMap(t *testing.T) {
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

// 	result, err := logs.DataSourceIbmLogsViewsApisViewsV1TimeSelectionSelectionTypeQuickSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsViewsApisViewsV1TimeSelectionSelectionTypeCustomSelectionToMap(t *testing.T) {
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

// 	result, err := logs.DataSourceIbmLogsViewsApisViewsV1TimeSelectionSelectionTypeCustomSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsViewsApisViewsV1SelectedFiltersToMap(t *testing.T) {
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

// 	result, err := logs.DataSourceIbmLogsViewsApisViewsV1SelectedFiltersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsViewsApisViewsV1FilterToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["name"] = "applicationName"
// 		model["selected_values"] = map[string]interface{}{"key1": fmt.Sprintf("%v", true)}

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisViewsV1Filter)
// 	model.Name = core.StringPtr("applicationName")
// 	model.SelectedValues = map[string]bool{"key1": true}

// 	result, err := logs.DataSourceIbmLogsViewsApisViewsV1FilterToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

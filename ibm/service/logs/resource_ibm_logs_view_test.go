// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logs"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmLogsViewBasic(t *testing.T) {
	var conf logsv0.View
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsViewDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsViewConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsViewExists("ibm_logs_view.logs_view_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_view.logs_view_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsViewConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_view.logs_view_instance", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_view.logs_view_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsViewConfigBasic(name string) string {
	return fmt.Sprintf(`
	resource "ibm_logs_view" "logs_view_instance" {
		instance_id = "%s"
		region      = "%s"
		name        = "%s"
		filters {
		  filters {
			name = "applicationName"
			selected_values = {
			  demo = true
			}
		  }
		  filters {
			name = "subsystemName"
			selected_values = {
			  demo = true
			}
		  }
		  filters {
			name = "operationName"
			selected_values = {
			  demo = true
			}
		  }
		  filters {
			name = "serviceName"
			selected_values = {
			  demo = true
			}
		  }
		  filters {
			name = "severity"
			selected_values = {
			  demo = true
			}
		  }
		}
		search_query {
		  query = "logs"
		}
		time_selection {
		  custom_selection {
			from_time = "2024-01-25T11:31:43.152Z"
			to_time   = "2024-01-25T11:37:13.238Z"
		  }
		}
	}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name)
}

func testAccCheckIbmLogsViewExists(n string, obj logsv0.View) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
		if err != nil {
			return err
		}
		logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getViewOptions := &logsv0.GetViewOptions{}

		viewIdInt, _ := strconv.ParseInt(resourceID[2], 10, 64)
		getViewOptions.SetID(viewIdInt)

		view, _, err := logsClient.GetView(getViewOptions)
		if err != nil {
			return err
		}

		obj = *view
		return nil
	}
}

func testAccCheckIbmLogsViewDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_view" {
			continue
		}
		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getViewOptions := &logsv0.GetViewOptions{}

		viewIdInt, _ := strconv.ParseInt(resourceID[2], 10, 64)
		getViewOptions.SetID(viewIdInt)

		// Try to find the key
		_, response, err := logsClient.GetView(getViewOptions)

		if err == nil {
			return fmt.Errorf("logs_view still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_view (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmLogsViewApisViewsV1SearchQueryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["query"] = "error"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisViewsV1SearchQuery)
	model.Query = core.StringPtr("error")

	result, err := logs.ResourceIbmLogsViewApisViewsV1SearchQueryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

// Todo @kavya498: Fix unit testcases
// func TestResourceIbmLogsViewApisViewsV1TimeSelectionToMap(t *testing.T) {
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

// 	result, err := logs.ResourceIbmLogsViewApisViewsV1TimeSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

func TestResourceIbmLogsViewApisViewsV1QuickTimeSelectionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["caption"] = "Last hour"
		model["seconds"] = int(3600)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisViewsV1QuickTimeSelection)
	model.Caption = core.StringPtr("Last hour")
	model.Seconds = core.Int64Ptr(int64(3600))

	result, err := logs.ResourceIbmLogsViewApisViewsV1QuickTimeSelectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsViewApisViewsV1CustomTimeSelectionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["from_time"] = "2024-01-25T11:31:43.152Z"
		model["to_time"] = "2024-01-25T11:35:43.152Z"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisViewsV1CustomTimeSelection)
	model.FromTime = CreateMockDateTime("2024-01-25T11:31:43.152Z")
	model.ToTime = CreateMockDateTime("2024-01-25T11:35:43.152Z")

	result, err := logs.ResourceIbmLogsViewApisViewsV1CustomTimeSelectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsViewApisViewsV1TimeSelectionSelectionTypeQuickSelectionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisViewsV1QuickTimeSelectionModel := make(map[string]interface{})
		apisViewsV1QuickTimeSelectionModel["caption"] = "Last hour"
		apisViewsV1QuickTimeSelectionModel["seconds"] = int(3600)

		model := make(map[string]interface{})
		model["quick_selection"] = []map[string]interface{}{apisViewsV1QuickTimeSelectionModel}

		assert.Equal(t, result, model)
	}

	apisViewsV1QuickTimeSelectionModel := new(logsv0.ApisViewsV1QuickTimeSelection)
	apisViewsV1QuickTimeSelectionModel.Caption = core.StringPtr("Last hour")
	apisViewsV1QuickTimeSelectionModel.Seconds = core.Int64Ptr(int64(3600))

	model := new(logsv0.ApisViewsV1TimeSelectionSelectionTypeQuickSelection)
	model.QuickSelection = apisViewsV1QuickTimeSelectionModel

	result, err := logs.ResourceIbmLogsViewApisViewsV1TimeSelectionSelectionTypeQuickSelectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsViewApisViewsV1TimeSelectionSelectionTypeCustomSelectionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisViewsV1CustomTimeSelectionModel := make(map[string]interface{})
		apisViewsV1CustomTimeSelectionModel["from_time"] = "2024-01-25T11:31:43.152Z"
		apisViewsV1CustomTimeSelectionModel["to_time"] = "2024-01-25T11:35:43.152Z"

		model := make(map[string]interface{})
		model["custom_selection"] = []map[string]interface{}{apisViewsV1CustomTimeSelectionModel}

		assert.Equal(t, result, model)
	}

	apisViewsV1CustomTimeSelectionModel := new(logsv0.ApisViewsV1CustomTimeSelection)
	apisViewsV1CustomTimeSelectionModel.FromTime = CreateMockDateTime("2024-01-25T11:31:43.152Z")
	apisViewsV1CustomTimeSelectionModel.ToTime = CreateMockDateTime("2024-01-25T11:35:43.152Z")

	model := new(logsv0.ApisViewsV1TimeSelectionSelectionTypeCustomSelection)
	model.CustomSelection = apisViewsV1CustomTimeSelectionModel

	result, err := logs.ResourceIbmLogsViewApisViewsV1TimeSelectionSelectionTypeCustomSelectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsViewApisViewsV1SelectedFiltersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisViewsV1FilterModel := make(map[string]interface{})
		apisViewsV1FilterModel["name"] = "applicationName"
		apisViewsV1FilterModel["selected_values"] = map[string]interface{}{"key1": fmt.Sprintf("%v", true)}

		model := make(map[string]interface{})
		model["filters"] = []map[string]interface{}{apisViewsV1FilterModel}

		assert.Equal(t, result, model)
	}

	apisViewsV1FilterModel := new(logsv0.ApisViewsV1Filter)
	apisViewsV1FilterModel.Name = core.StringPtr("applicationName")
	apisViewsV1FilterModel.SelectedValues = map[string]bool{"key1": true}

	model := new(logsv0.ApisViewsV1SelectedFilters)
	model.Filters = []logsv0.ApisViewsV1Filter{*apisViewsV1FilterModel}

	result, err := logs.ResourceIbmLogsViewApisViewsV1SelectedFiltersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsViewApisViewsV1FilterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "applicationName"
		model["selected_values"] = map[string]interface{}{"key1": fmt.Sprintf("%v", true)}

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisViewsV1Filter)
	model.Name = core.StringPtr("applicationName")
	model.SelectedValues = map[string]bool{"key1": true}

	result, err := logs.ResourceIbmLogsViewApisViewsV1FilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsViewMapToApisViewsV1SearchQuery(t *testing.T) {
	checkResult := func(result *logsv0.ApisViewsV1SearchQuery) {
		model := new(logsv0.ApisViewsV1SearchQuery)
		model.Query = core.StringPtr("error")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["query"] = "error"

	result, err := logs.ResourceIbmLogsViewMapToApisViewsV1SearchQuery(model)
	assert.Nil(t, err)
	checkResult(result)
}

// Todo @kavya498: Fix unit testcases
// func TestResourceIbmLogsViewMapToApisViewsV1TimeSelection(t *testing.T) {
// 	checkResult := func(result logsv0.ApisViewsV1TimeSelectionIntf) {
// 		apisViewsV1QuickTimeSelectionModel := new(logsv0.ApisViewsV1QuickTimeSelection)
// 		apisViewsV1QuickTimeSelectionModel.Caption = core.StringPtr("Last hour")
// 		apisViewsV1QuickTimeSelectionModel.Seconds = core.Int64Ptr(int64(3600))

// 		model := new(logsv0.ApisViewsV1TimeSelection)
// 		model.QuickSelection = apisViewsV1QuickTimeSelectionModel
// 		model.CustomSelection = apisViewsV1CustomTimeSelectionModel

// 		assert.Equal(t, result, model)
// 	}

// 	apisViewsV1QuickTimeSelectionModel := make(map[string]interface{})
// 	apisViewsV1QuickTimeSelectionModel["caption"] = "Last hour"
// 	apisViewsV1QuickTimeSelectionModel["seconds"] = int(3600)

// 	model := make(map[string]interface{})
// 	model["quick_selection"] = []interface{}{apisViewsV1QuickTimeSelectionModel}
// 	model["custom_selection"] = []interface{}{apisViewsV1CustomTimeSelectionModel}

// 	result, err := logs.ResourceIbmLogsViewMapToApisViewsV1TimeSelection(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

func TestResourceIbmLogsViewMapToApisViewsV1QuickTimeSelection(t *testing.T) {
	checkResult := func(result *logsv0.ApisViewsV1QuickTimeSelection) {
		model := new(logsv0.ApisViewsV1QuickTimeSelection)
		model.Caption = core.StringPtr("Last hour")
		model.Seconds = core.Int64Ptr(int64(3600))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["caption"] = "Last hour"
	model["seconds"] = int(3600)

	result, err := logs.ResourceIbmLogsViewMapToApisViewsV1QuickTimeSelection(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsViewMapToApisViewsV1CustomTimeSelection(t *testing.T) {
	checkResult := func(result *logsv0.ApisViewsV1CustomTimeSelection) {
		model := new(logsv0.ApisViewsV1CustomTimeSelection)
		model.FromTime = CreateMockDateTime("2024-01-25T11:31:43.152Z")
		model.ToTime = CreateMockDateTime("2024-01-25T11:35:43.152Z")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["from_time"] = "2024-01-25T11:31:43.152Z"
	model["to_time"] = "2024-01-25T11:35:43.152Z"

	result, err := logs.ResourceIbmLogsViewMapToApisViewsV1CustomTimeSelection(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsViewMapToApisViewsV1TimeSelectionSelectionTypeQuickSelection(t *testing.T) {
	checkResult := func(result *logsv0.ApisViewsV1TimeSelectionSelectionTypeQuickSelection) {
		apisViewsV1QuickTimeSelectionModel := new(logsv0.ApisViewsV1QuickTimeSelection)
		apisViewsV1QuickTimeSelectionModel.Caption = core.StringPtr("Last hour")
		apisViewsV1QuickTimeSelectionModel.Seconds = core.Int64Ptr(int64(3600))

		model := new(logsv0.ApisViewsV1TimeSelectionSelectionTypeQuickSelection)
		model.QuickSelection = apisViewsV1QuickTimeSelectionModel

		assert.Equal(t, result, model)
	}

	apisViewsV1QuickTimeSelectionModel := make(map[string]interface{})
	apisViewsV1QuickTimeSelectionModel["caption"] = "Last hour"
	apisViewsV1QuickTimeSelectionModel["seconds"] = int(3600)

	model := make(map[string]interface{})
	model["quick_selection"] = []interface{}{apisViewsV1QuickTimeSelectionModel}

	result, err := logs.ResourceIbmLogsViewMapToApisViewsV1TimeSelectionSelectionTypeQuickSelection(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsViewMapToApisViewsV1TimeSelectionSelectionTypeCustomSelection(t *testing.T) {
	checkResult := func(result *logsv0.ApisViewsV1TimeSelectionSelectionTypeCustomSelection) {
		apisViewsV1CustomTimeSelectionModel := new(logsv0.ApisViewsV1CustomTimeSelection)
		apisViewsV1CustomTimeSelectionModel.FromTime = CreateMockDateTime("2024-01-25T11:31:43.152Z")
		apisViewsV1CustomTimeSelectionModel.ToTime = CreateMockDateTime("2024-01-25T11:35:43.152Z")

		model := new(logsv0.ApisViewsV1TimeSelectionSelectionTypeCustomSelection)
		model.CustomSelection = apisViewsV1CustomTimeSelectionModel

		assert.Equal(t, result, model)
	}

	apisViewsV1CustomTimeSelectionModel := make(map[string]interface{})
	apisViewsV1CustomTimeSelectionModel["from_time"] = "2024-01-25T11:31:43.152Z"
	apisViewsV1CustomTimeSelectionModel["to_time"] = "2024-01-25T11:35:43.152Z"

	model := make(map[string]interface{})
	model["custom_selection"] = []interface{}{apisViewsV1CustomTimeSelectionModel}

	result, err := logs.ResourceIbmLogsViewMapToApisViewsV1TimeSelectionSelectionTypeCustomSelection(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsViewMapToApisViewsV1SelectedFilters(t *testing.T) {
	checkResult := func(result *logsv0.ApisViewsV1SelectedFilters) {
		apisViewsV1FilterModel := new(logsv0.ApisViewsV1Filter)
		apisViewsV1FilterModel.Name = core.StringPtr("applicationName")
		apisViewsV1FilterModel.SelectedValues = map[string]bool{"key1": true}

		model := new(logsv0.ApisViewsV1SelectedFilters)
		model.Filters = []logsv0.ApisViewsV1Filter{*apisViewsV1FilterModel}

		assert.Equal(t, result, model)
	}

	apisViewsV1FilterModel := make(map[string]interface{})
	apisViewsV1FilterModel["name"] = "applicationName"
	apisViewsV1FilterModel["selected_values"] = map[string]interface{}{"key1": true}

	model := make(map[string]interface{})
	model["filters"] = []interface{}{apisViewsV1FilterModel}

	result, err := logs.ResourceIbmLogsViewMapToApisViewsV1SelectedFilters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsViewMapToApisViewsV1Filter(t *testing.T) {
	checkResult := func(result *logsv0.ApisViewsV1Filter) {
		model := new(logsv0.ApisViewsV1Filter)
		model.Name = core.StringPtr("applicationName")
		model.SelectedValues = map[string]bool{"key1": true}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "applicationName"
	model["selected_values"] = map[string]interface{}{"key1": true}

	result, err := logs.ResourceIbmLogsViewMapToApisViewsV1Filter(model)
	assert.Nil(t, err)
	checkResult(result)
}

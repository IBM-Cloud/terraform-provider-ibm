// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.92.0-af5c89a5-20240617-153232
 */

package configurationaggregator_test

import (
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/configurationaggregator"
	"github.com/IBM/configuration-aggregator-go-sdk/configurationaggregatorv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmConfigAggregatorResourceCollectionStatusDataSourceBasic(t *testing.T) {
	instanceID := "instance_id"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmConfigAggregatorResourceCollectionStatusDataSourceConfigBasic(instanceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_config_aggregator_resource_collection_status.config_aggregator_resource_collection_status_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_config_aggregator_resource_collection_status.config_aggregator_resource_collection_status_instance", "latest_reconciliation_mode"),
				),
			},
		},
	})
}

func testAccCheckIbmConfigAggregatorResourceCollectionStatusDataSourceConfigBasic(instanceID string) string {
	return fmt.Sprintf(`
		data "ibm_config_aggregator_resource_collection_status" "config_aggregator_resource_collection_status_instance" {
			instance_id="%s"
		}
	`, instanceID)
}

func TestDataSourceIbmConfigAggregatorResourceCollectionStatusReconciliationModeStatusToMap(t *testing.T) {
	refreshTime := strfmt.DateTime{}
	model := &configurationaggregatorv1.ReconciliationModeStatus{LastRefreshTime: &refreshTime}

	result := configurationaggregator.DataSourceIbmConfigAggregatorResourceCollectionStatusReconciliationModeStatusToMap(model)
	assert.Equal(t, map[string]interface{}{"last_refresh_time": refreshTime.String()}, result)
}

func TestDataSourceIbmConfigAggregatorResourceCollectionStatusStatusResponseReconciliationModesToMap(t *testing.T) {
	refreshTime := strfmt.DateTime{}
	modeStatus := &configurationaggregatorv1.ReconciliationModeStatus{LastRefreshTime: &refreshTime}
	modeMap := map[string]interface{}{"last_refresh_time": refreshTime.String()}

	model := &configurationaggregatorv1.StatusResponseReconciliationModes{
		Automatic: modeStatus,
		Realtime:  modeStatus,
	}

	result, err := configurationaggregator.DataSourceIbmConfigAggregatorResourceCollectionStatusStatusResponseReconciliationModesToMap(model)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{
		"automatic": []map[string]interface{}{modeMap},
		"realtime":  []map[string]interface{}{modeMap},
	}, result)
}

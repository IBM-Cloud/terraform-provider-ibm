// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.92.0-af5c89a5-20240617-153232
 */

package configurationaggregator_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
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

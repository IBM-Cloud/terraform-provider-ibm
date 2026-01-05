// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
*/

package distributionlistapi_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmDistributionListDestinationDataSourceBasic(t *testing.T) {
	eventNotificationDestinationAccountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))
	eventNotificationDestinationDestinationType := "event_notifications"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmDistributionListDestinationDataSourceConfigBasic(eventNotificationDestinationAccountID, eventNotificationDestinationDestinationType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_distribution_list_destination.distribution_list_destination_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_distribution_list_destination.distribution_list_destination_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_distribution_list_destination.distribution_list_destination_instance", "destination_id"),
					resource.TestCheckResourceAttrSet("data.ibm_distribution_list_destination.distribution_list_destination_instance", "destination_type"),
				),
			},
		},
	})
}

func testAccCheckIbmDistributionListDestinationDataSourceConfigBasic(eventNotificationDestinationAccountID string, eventNotificationDestinationDestinationType string) string {
	return fmt.Sprintf(`
		resource "ibm_distribution_list_destination" "distribution_list_destination_instance" {
			account_id = "%s"
			destination_type = "%s"
		}

		data "ibm_distribution_list_destination" "distribution_list_destination_instance" {
			account_id = ibm_distribution_list_destination.distribution_list_destination_instance.account_id
			destination_id = ibm_distribution_list_destination.distribution_list_destination_instance.destination_id
		}
	`, eventNotificationDestinationAccountID, eventNotificationDestinationDestinationType)
}


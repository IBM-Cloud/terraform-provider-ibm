// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmConnectionRegistrationTokenBasic(t *testing.T) {
	connectionID := fmt.Sprintf("tf_connection_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmConnectionRegistrationTokenConfigBasic(connectionID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_connection-registration-token.connection_registration_token_instance", "connection_id", connectionID),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_connection-registration-token.connection_registration_token",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmConnectionRegistrationTokenConfigBasic(connectionID string) string {
	return fmt.Sprintf(`
		resource "ibm_connection-registration-token" "connection_registration_token_instance" {
			connection_id = "%s"
		}
	`, connectionID)
}

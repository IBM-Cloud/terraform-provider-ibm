// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmEnRegistrationDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmEnRegistrationDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_en_registration.sm_en_registration", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_en_registration.sm_en_registration", "event_notifications_instance_crn"),
				),
			},
		},
	})
}

func testAccCheckIbmSmEnRegistrationDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_en_registration" "sm_en_registration_instance"{
  			instance_id   = "%s"
  			region        = "%s"
  			event_notifications_instance_crn = "%s"
  			event_notifications_source_description = "Terraform data source test."
  			event_notifications_source_name = "My Secrets Manager Terraform Test"
}
		data "ibm_sm_en_registration" "sm_en_registration" {
			instance_id = "%s"
			region = "%s"
			depends_on = [
				ibm_sm_en_registration.sm_en_registration_instance
  			]
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerENInstanceCrn, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

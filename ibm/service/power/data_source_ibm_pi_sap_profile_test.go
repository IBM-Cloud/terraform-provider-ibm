// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPISAPProfileDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPISAPProfileDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_sap_profile.test", "id"),
					resource.TestCheckResourceAttr("data.ibm_pi_sap_profile.test", "id", acc.Pi_sap_profile_id),
				),
			},
		},
	})
}

func testAccCheckIBMPISAPProfileDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_sap_profile" "test" {
			pi_cloud_instance_id = "%s"
			pi_sap_profile_id = "%s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_sap_profile_id)
}

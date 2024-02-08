// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIWorkspacesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIWorkspacesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_workspaces.test", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIWorkspacesDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_workspaces" "test" {
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_cloud_instance_id)
}

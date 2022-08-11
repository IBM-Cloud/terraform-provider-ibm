// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPICloudConnectionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICloudConnectionDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_cloud_connection.example", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pi_cloud_connection.example", "connection_mode"),
				),
			},
		},
	})
}

func testAccCheckIBMPICloudConnectionDataSourceConfig() string {
	return fmt.Sprintf(`
	data "ibm_pi_cloud_connection" "example" {
		pi_cloud_connection_name 	= "%s"
		pi_cloud_instance_id 		= "%s"
	}
	`, acc.PiCloudConnectionName, acc.Pi_cloud_instance_id)
}

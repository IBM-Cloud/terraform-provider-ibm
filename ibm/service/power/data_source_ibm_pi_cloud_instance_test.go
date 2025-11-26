// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMPICloudInstanceDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICloudInstanceDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_cloud_instance.testacc_ds_cloud_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPICloudInstanceDataSourceConfig() string {
	return fmt.Sprintf(`	
		data "ibm_pi_cloud_instance" "testacc_ds_cloud_instance" {
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_cloud_instance_id)
}

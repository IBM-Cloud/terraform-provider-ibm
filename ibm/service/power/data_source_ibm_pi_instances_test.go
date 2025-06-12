// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIInstancesDataSource_basic(t *testing.T) {
	instancesResData := "data.ibm_pi_instances.testacc_ds_instance"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstancesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(instancesResData, "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstancesDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_instances" "testacc_ds_instance" {
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_cloud_instance_id)
}

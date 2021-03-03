// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisDataSource_basic(t *testing.T) {
	instanceName := fmt.Sprintf(cisInstance)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCis(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:  testAccCheckIBMCisDataSourceConfig(instanceName),
				Destroy: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_cis.cis", "name", instanceName),
					resource.TestCheckResourceAttr("data.ibm_cis.cis", "service", "internet-svcs"),
					resource.TestCheckResourceAttr("data.ibm_cis.cis", "plan", "enterprise-usage"),
					resource.TestCheckResourceAttr("data.ibm_cis.cis", "location", "global"),
				),
			},
		},
	})
}

func testAccCheckIBMCisDataSourceConfig(instanceName string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}
	  
	  data "ibm_cis" "cis" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
	}
	`, cisResourceGroup, instanceName)

}

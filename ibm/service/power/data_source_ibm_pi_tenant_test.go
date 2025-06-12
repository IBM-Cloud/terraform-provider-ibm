// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPITenantDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPITenantDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_tenant.testacc_ds_tenant", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPITenantDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_tenant" "testacc_ds_tenant" {
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_cloud_instance_id)
}

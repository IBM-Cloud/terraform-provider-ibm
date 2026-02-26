// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package powerhaautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPhaGetLastOperationDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaGetLastOperationDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pha_get_last_operation.pha_get_last_operation_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_get_last_operation.pha_get_last_operation_instance", "pha_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_get_last_operation.pha_get_last_operation_instance", "deployment_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_get_last_operation.pha_get_last_operation_instance", "provision_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_get_last_operation.pha_get_last_operation_instance", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_get_last_operation.pha_get_last_operation_instance", "status"),
				),
			},
		},
	})
}

func testAccCheckIBMPhaGetLastOperationDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pha_get_last_operation" "pha_get_last_operation_instance" {
			pha_instance_id = "8ce2a099-a463-479a-9a1d-eedc19287a62"
		}
	`)
}

// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.116.0-df613dbc-20260803-154903
*/

package powerhaautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPhaLastOperationDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaLastOperationDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pha_last_operation.pha_last_operation_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_last_operation.pha_last_operation_instance", "pha_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_last_operation.pha_last_operation_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_last_operation.pha_last_operation_instance", "deployment_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_last_operation.pha_last_operation_instance", "provision_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_last_operation.pha_last_operation_instance", "resource_group"),
				),
			},
		},
	})
}

func testAccCheckIBMPhaLastOperationDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pha_last_operation" "pha_last_operation_instance" {
			pha_instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
			Accept-Language = "en-US"
			If-None-Match = "abcdef"
		}
	`)
}


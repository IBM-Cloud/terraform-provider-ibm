// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.0-d6dec9d7-20241008-212902
 */

package db2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmDb2ConnectionInfoDataSourceBasic(t *testing.T) {
	db2DeploymentId := "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-east%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3Af9455c22-07af-4a86-b9df-f02fd4774471%3A%3A"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmDb2ConnectionInfoDataSourceConfigBasic(db2DeploymentId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_db2_connection_info.db2_connection_info_instance", "deployment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_connection_info.db2_connection_info_instance", "x_deployment_id"),
				),
			},
		},
	})
}

func testAccCheckIbmDb2ConnectionInfoDataSourceConfigBasic(db2DeploymentId string) string {
	return fmt.Sprintf(`
		data "ibm_db2_connection_info" "db2_connection_info_instance" {
			deployment_id = "%[1]s"
            x_deployment_id = "crn:v1:staging:public:dashdb-for-transactions:us-east:a/e7e3e87b512f474381c0684a5ecbba03:f9455c22-07af-4a86-b9df-f02fd4774471::"
            
		}
	`, db2DeploymentId)
}

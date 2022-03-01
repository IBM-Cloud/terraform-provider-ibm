// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMDatabaseConnectionDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMDatabaseConnectionDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database_connections.database_connection", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_connections.database_connection", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_connections.database_connection", "user_type"),
					resource.TestCheckResourceAttrSet("data.ibm_database_connections.database_connection", "user_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_connections.database_connection", "endpoint_type"),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseConnectionDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_database_connections" "database_connection" {
			id = "crn:v1:bluemix:public:databases-for-postgresql:us-east:a/40ddc34a953a8c02f10987b59085b60e:dd922d62-2fda-43fa-ab1f-9f4d058d5893::"
			user_type = "database"
			user_id = "user_id"
			endpoint_type = "public"
		}
	`)
}

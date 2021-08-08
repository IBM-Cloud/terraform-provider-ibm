// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmCloudantDatabaseDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCloudantDatabaseDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cloudant_database.cloudant_database", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cloudant_database.cloudant_database", "db"),
					resource.TestCheckResourceAttrSet("data.ibm_cloudant_database.cloudant_database", "cluster.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cloudant_database.cloudant_database", "compact_running"),
					resource.TestCheckResourceAttrSet("data.ibm_cloudant_database.cloudant_database", "db_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cloudant_database.cloudant_database", "disk_format_version"),
					resource.TestCheckResourceAttrSet("data.ibm_cloudant_database.cloudant_database", "doc_count"),
					resource.TestCheckResourceAttrSet("data.ibm_cloudant_database.cloudant_database", "doc_del_count"),
					resource.TestCheckResourceAttrSet("data.ibm_cloudant_database.cloudant_database", "props.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cloudant_database.cloudant_database", "sizes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cloudant_database.cloudant_database", "update_seq"),
				),
			},
		},
	})
}

func testAccCheckIbmCloudantDatabaseDataSourceConfigBasic() string {
	return fmt.Sprintf(`
			data "ibm_resource_group" "cloudant" {
				is_default=true
	  		}
  
	  		resource "ibm_resource_instance" "cloudant_instance" {
				name              = "pr01"
				service           = "cloudantnosqldb"
				plan              = "standard"
				location          = "us-east"
				resource_group_id = data.ibm_resource_group.cloudant.id
	  		}

			resource "ibm_cloudant_database" "cloudant_database" {
				cloudant_guid = ibm_resource_instance.cloudant_instance.guid
				db = "db"
			}

			data "ibm_cloudant_database" "cloudant_database" {
				db = ibm_cloudant_database.cloudant_database.db
				cloudant_guid = ibm_cloudant_database.cloudant_database.cloudant_guid
			}
	`)
}

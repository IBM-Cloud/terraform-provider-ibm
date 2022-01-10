// Copyright IBM Corp. 2021, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudant_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCloudantDatabaseDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCloudantDatabaseDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cloudant_database.cloudant_database", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cloudant_database.cloudant_database", "db"),
					resource.TestCheckResourceAttrSet("data.ibm_cloudant_database.cloudant_database", "cluster.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cloudant_database.cloudant_database", "compact_running"),
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

func testAccCheckIBMCloudantDatabaseDataSourceConfigBasic() string {
	return `
			data "ibm_resource_group" "cloudant" {
				is_default=true
			}

			resource "ibm_cloudant" "cloudant_instance" {
				name              = "pr01"
				plan              = "standard"
				location          = "us-south"
				resource_group_id = data.ibm_resource_group.cloudant.id
			}

			resource "ibm_cloudant_database" "cloudant_database" {
				instance_crn = ibm_cloudant.cloudant_instance.crn
				db = "db"
			}

			data "ibm_cloudant_database" "cloudant_database" {
				db = ibm_cloudant_database.cloudant_database.db
				instance_crn = ibm_cloudant_database.cloudant_database.instance_crn
			}
	`
}

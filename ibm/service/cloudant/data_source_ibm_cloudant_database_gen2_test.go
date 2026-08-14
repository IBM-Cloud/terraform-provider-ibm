// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudant_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCloudantGen2DatabaseDataSourceBasic(t *testing.T) {
	serviceName := fmt.Sprintf("terraform-test-%s", acctest.RandString(8))
	db := fmt.Sprintf("tf_db_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCloudantGen2DatabaseDataSourceConfigBasic(serviceName, db),
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

func testAccCheckIBMCloudantGen2DatabaseDataSourceConfigBasic(serviceName, db string) string {
	return fmt.Sprintf(`

			resource "ibm_cloudant" "cloudant_instance" {
				name     = "%s"
				plan     = "standard-gen2"
				location = "%s"

				timeouts {
					create = "15m"
					update = "15m"
					delete = "15m"
				}
			}

			resource "ibm_cloudant_database" "cloudant_database" {
				instance_crn = ibm_cloudant.cloudant_instance.crn
				db = "%s"
			}

			data "ibm_cloudant_database" "cloudant_database" {
				db = ibm_cloudant_database.cloudant_database.db
				instance_crn = ibm_cloudant_database.cloudant_database.instance_crn
			}
	`, serviceName, acc.Region(), db)
}

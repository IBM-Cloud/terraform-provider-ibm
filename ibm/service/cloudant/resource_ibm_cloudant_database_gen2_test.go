// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudant_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

func TestAccIBMCloudantGen2DatabaseBasic(t *testing.T) {
	var conf cloudantv1.DatabaseInformation
	instanceName := fmt.Sprintf("tf_instance_%d", acctest.RandIntRange(10, 100))
	db := fmt.Sprintf("tf_db_%d", acctest.RandIntRange(10, 100))
	dbUpdate := fmt.Sprintf("tf_db_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCloudantDatabaseDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCloudantGen2DatabaseConfigBasic(instanceName, db),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCloudantDatabaseExists("ibm_cloudant_database.cloudant_database", conf),
					resource.TestCheckResourceAttr("ibm_cloudant_database.cloudant_database", "db", db),
				),
			},
			{
				Config: testAccCheckIBMCloudantGen2DatabaseConfigBasic(instanceName, dbUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloudant_database.cloudant_database", "db", dbUpdate),
				),
			},
		},
	})
}

func testAccCheckIBMCloudantGen2DatabaseConfigBasic(instanceName, db string) string {
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
			db           = "%s"
		}
	`, instanceName, acc.Region(), db)
}

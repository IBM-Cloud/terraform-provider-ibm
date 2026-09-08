// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.3-943fbc81-20260603-173645
*/

package brsmigration_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBrsMigrationDataSourceBasic(t *testing.T) {
	migrationName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	migrationBrsCrn := fmt.Sprintf("tf_brs_crn_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationDataSourceConfigBasic(migrationName, migrationBrsCrn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration.brs_migration_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration.brs_migration_instance", "migration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration.brs_migration_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration.brs_migration_instance", "brs_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration.brs_migration_instance", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration.brs_migration_instance", "created_at"),
				),
			},
		},
	})
}

func TestAccIbmBrsMigrationDataSourceAllArgs(t *testing.T) {
	migrationName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	migrationBrsCrn := fmt.Sprintf("tf_brs_crn_%d", acctest.RandIntRange(10, 100))
	migrationDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationDataSourceConfig(migrationName, migrationBrsCrn, migrationDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration.brs_migration_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration.brs_migration_instance", "migration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration.brs_migration_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration.brs_migration_instance", "brs_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration.brs_migration_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration.brs_migration_instance", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration.brs_migration_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration.brs_migration_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration.brs_migration_instance", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIbmBrsMigrationDataSourceConfigBasic(migrationName string, migrationBrsCrn string) string {
	return fmt.Sprintf(`
		resource "ibm_brs_migration" "brs_migration_instance" {
			name = "%s"
			brs_crn = "%s"
		}

		data "ibm_brs_migration" "brs_migration_instance" {
			migration_id = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
		}
	`, migrationName, migrationBrsCrn)
}

func testAccCheckIbmBrsMigrationDataSourceConfig(migrationName string, migrationBrsCrn string, migrationDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_brs_migration" "brs_migration_instance" {
			name = "%s"
			brs_crn = "%s"
			description = "%s"
		}

		data "ibm_brs_migration" "brs_migration_instance" {
			migration_id = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
		}
	`, migrationName, migrationBrsCrn, migrationDescription)
}

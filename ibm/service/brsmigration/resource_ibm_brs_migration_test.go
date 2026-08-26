// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package brsmigration_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.ibm.com/BackupAndRecovery/brs-migration-orchestrator/brsmigrationv2"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBrsMigrationBasic(t *testing.T) {
	var conf brsmigrationv2.Migration
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	brsCrn := fmt.Sprintf("tf_brs_crn_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	brsCrnUpdate := fmt.Sprintf("tf_brs_crn_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBrsMigrationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationConfigBasic(name, brsCrn),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBrsMigrationExists("ibm_brs_migration.brs_migration_instance", conf),
					resource.TestCheckResourceAttr("ibm_brs_migration.brs_migration_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_brs_migration.brs_migration_instance", "brs_crn", brsCrn),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationConfigBasic(nameUpdate, brsCrnUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_brs_migration.brs_migration_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_brs_migration.brs_migration_instance", "brs_crn", brsCrnUpdate),
				),
			},
		},
	})
}

func TestAccIbmBrsMigrationAllArgs(t *testing.T) {
	var conf brsmigrationv2.Migration
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	brsCrn := fmt.Sprintf("tf_brs_crn_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	brsCrnUpdate := fmt.Sprintf("tf_brs_crn_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBrsMigrationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationConfig(name, brsCrn, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBrsMigrationExists("ibm_brs_migration.brs_migration_instance", conf),
					resource.TestCheckResourceAttr("ibm_brs_migration.brs_migration_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_brs_migration.brs_migration_instance", "brs_crn", brsCrn),
					resource.TestCheckResourceAttr("ibm_brs_migration.brs_migration_instance", "description", description),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationConfig(nameUpdate, brsCrnUpdate, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_brs_migration.brs_migration_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_brs_migration.brs_migration_instance", "brs_crn", brsCrnUpdate),
					resource.TestCheckResourceAttr("ibm_brs_migration.brs_migration_instance", "description", descriptionUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_brs_migration.brs_migration_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBrsMigrationConfigBasic(name string, brsCrn string) string {
	return fmt.Sprintf(`
		resource "ibm_brs_migration" "brs_migration_instance" {
			name = "%s"
			brs_crn = "%s"
		}
	`, name, brsCrn)
}

func testAccCheckIbmBrsMigrationConfig(name string, brsCrn string, description string) string {
	return fmt.Sprintf(`

		resource "ibm_brs_migration" "brs_migration_instance" {
			name = "%s"
			brs_crn = "%s"
			description = "%s"
		}
	`, name, brsCrn, description)
}

func testAccCheckIbmBrsMigrationExists(n string, obj brsmigrationv2.Migration) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		brsMigrationClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BrsMigrationV2()
		if err != nil {
			return err
		}

		getMigrationOptions := &brsmigrationv2.GetMigrationOptions{}

		getMigrationOptions.SetMigrationID(rs.Primary.ID)

		migration, _, err := brsMigrationClient.GetMigration(getMigrationOptions)
		if err != nil {
			return err
		}

		obj = *migration
		return nil
	}
}

func testAccCheckIbmBrsMigrationDestroy(s *terraform.State) error {
	brsMigrationClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BrsMigrationV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_brs_migration" {
			continue
		}

		getMigrationOptions := &brsmigrationv2.GetMigrationOptions{}

		getMigrationOptions.SetMigrationID(rs.Primary.ID)

		// Try to find the key
		_, response, err := brsMigrationClient.GetMigration(getMigrationOptions)

		if err == nil {
			return fmt.Errorf("brs_migration still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for brs_migration (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

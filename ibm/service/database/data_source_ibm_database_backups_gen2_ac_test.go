// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

// TestAccIBMDatabaseBackupsGen2DataSourceBasic verifies that listing Independent
// Backups for a Gen2 deployment returns at least one backup with all expected
// fields populated.
func TestAccIBMDatabaseBackupsGen2DataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseBackupsGen2DataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database_backups.database_backups", "deployment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_backups.database_backups", "backups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_database_backups.database_backups", "backups.0.backup_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_backups.database_backups", "backups.0.deployment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_backups.database_backups", "backups.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_database_backups.database_backups", "backups.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_database_backups.database_backups", "backups.0.created_at"),
					resource.TestCheckResourceAttr("data.ibm_database_backups.database_backups", "backups.0.is_downloadable", "false"),
					resource.TestCheckResourceAttr("data.ibm_database_backups.database_backups", "backups.0.download_link", ""),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseBackupsGen2DataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_database_backups" "database_backups" {
			deployment_id = "%[1]s"
		}
	`, acc.Gen2DeploymentId)
}

// TestAccIBMDatabaseBackupsGen2DataSourceInvalidDeploymentID verifies the error
// path when the deployment_id does not correspond to a real resource instance.
func TestAccIBMDatabaseBackupsGen2DataSourceInvalidDeploymentID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMDatabaseBackupsGen2DataSourceInvalidDeploymentIDConfig(),
				ExpectError: regexp.MustCompile("failed to get resource instance|GetResourceInstance failed|not found"),
			},
		},
	})
}

func testAccCheckIBMDatabaseBackupsGen2DataSourceInvalidDeploymentIDConfig() string {
	return `
	data "ibm_database_backups" "invalid_deployment" {
		deployment_id = "crn:v1:bluemix:public:databases-for-mysql:us-south:a/00000000000000000000000000000000:00000000-0000-0000-0000-000000000000::"
	}
	`
}

// TestAccIBMDatabaseBackupsGen2DataSourceS2SWarning verifies that listing Gen2
// Independent Backups for an instance that has Independent Backups but no S2S
// authorizations emits a warning and still returns the backups list (non-blocking).
func TestAccIBMDatabaseBackupsGen2DataSourceS2SWarning(t *testing.T) {
	t.Parallel()
	name := fmt.Sprintf("tf-gen2-s2s-backups-%s", acctest.RandString(8))
	dsName := "data.ibm_database_backups.gen2_s2s_backups"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseBackupsGen2S2SConfig(name),
				Check: resource.ComposeTestCheckFunc(
					// backups list populated despite the S2S warning
					resource.TestCheckResourceAttrSet(dsName, "deployment_id"),
					resource.TestCheckResourceAttrSet(dsName, "backups.#"),
					resource.TestCheckResourceAttrSet(dsName, "backups.0.backup_id"),
					resource.TestCheckResourceAttrSet(dsName, "backups.0.deployment_id"),
					resource.TestCheckResourceAttrSet(dsName, "backups.0.type"),
					resource.TestCheckResourceAttrSet(dsName, "backups.0.status"),
					resource.TestCheckResourceAttrSet(dsName, "backups.0.created_at"),
					resource.TestCheckResourceAttr(dsName, "backups.0.is_downloadable", "false"),
					resource.TestCheckResourceAttr(dsName, "backups.0.download_link", ""),
				),
			},
		},
	})
}

// testAccCheckIBMDatabaseBackupsGen2S2SConfig creates a Gen2 MySQL instance in
// eu-fr2 without S2S authorizations, then lists its Independent Backups via the
// ibm_database_backups datasource. The S2S warning fires but must not prevent
// the backups list from being populated.
func testAccCheckIBMDatabaseBackupsGen2S2SConfig(name string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "gen2_s2s_backups" {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = %[1]q
  service           = "databases-for-mysql"
  plan              = "standard-gen2"
  location          = "eu-fr2"
  tags              = ["terraform", "s2s-backups-test"]

  group {
    group_id = "member"
    members {
      allocation_count = 2
    }
  }

  timeouts {
    create = "120m"
    update = "60m"
    delete = "15m"
  }
}

data "ibm_database_backups" "gen2_s2s_backups" {
  deployment_id = ibm_database.gen2_s2s_backups.id

  depends_on = [ibm_database.gen2_s2s_backups]
}
`, name)
}

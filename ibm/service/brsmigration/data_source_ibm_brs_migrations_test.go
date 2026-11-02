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

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/brsmigration"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.com/IBM/ibm-brs-migration-sdk-go/brsmigrationv1"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBrsMigrationsDataSourceBasic(t *testing.T) {
	migrationName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	migrationBrsCrn := fmt.Sprintf("tf_brs_crn_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationsDataSourceConfigBasic(migrationName, migrationBrsCrn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_brs_migrations.brs_migrations_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migrations.brs_migrations_instance", "migrations.#"),
					resource.TestCheckResourceAttr("data.ibm_brs_migrations.brs_migrations_instance", "migrations.0.name", migrationName),
					resource.TestCheckResourceAttr("data.ibm_brs_migrations.brs_migrations_instance", "migrations.0.brs_crn", migrationBrsCrn),
				),
			},
		},
	})
}

func TestAccIbmBrsMigrationsDataSourceAllArgs(t *testing.T) {
	migrationName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	migrationBrsCrn := fmt.Sprintf("tf_brs_crn_%d", acctest.RandIntRange(10, 100))
	migrationDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationsDataSourceConfig(migrationName, migrationBrsCrn, migrationDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_brs_migrations.brs_migrations_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migrations.brs_migrations_instance", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migrations.brs_migrations_instance", "migrations.#"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migrations.brs_migrations_instance", "migrations.0.id"),
					resource.TestCheckResourceAttr("data.ibm_brs_migrations.brs_migrations_instance", "migrations.0.name", migrationName),
					resource.TestCheckResourceAttr("data.ibm_brs_migrations.brs_migrations_instance", "migrations.0.brs_crn", migrationBrsCrn),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migrations.brs_migrations_instance", "migrations.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migrations.brs_migrations_instance", "migrations.0.state"),
					resource.TestCheckResourceAttr("data.ibm_brs_migrations.brs_migrations_instance", "migrations.0.description", migrationDescription),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migrations.brs_migrations_instance", "migrations.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migrations.brs_migrations_instance", "migrations.0.updated_at"),
				),
			},
		},
	})
}

func testAccCheckIbmBrsMigrationsDataSourceConfigBasic(migrationName string, migrationBrsCrn string) string {
	return fmt.Sprintf(`
		resource "ibm_brs_migration" "brs_migration_instance" {
			name = "%s"
			brs_crn = "%s"
		}

		data "ibm_brs_migrations" "brs_migrations_instance" {
			state = ibm_brs_migration.brs_migration_instance.state
		}
	`, migrationName, migrationBrsCrn)
}

func testAccCheckIbmBrsMigrationsDataSourceConfig(migrationName string, migrationBrsCrn string, migrationDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_brs_migration" "brs_migration_instance" {
			name = "%s"
			brs_crn = "%s"
			description = "%s"
		}

		data "ibm_brs_migrations" "brs_migrations_instance" {
			state = ibm_brs_migration.brs_migration_instance.state
		}
	`, migrationName, migrationBrsCrn, migrationDescription)
}

func TestDataSourceIbmBrsMigrationsMigrationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
		model["name"] = "prod-classic-to-vpc"
		model["brs_crn"] = "crn:v1:bluemix:public:backup-recovery:us-south:a/123456:abcdef01-2345-6789-abcd-ef0123456789::"
		model["crn"] = "crn:v1:bluemix:public:brs-migration:us-south:a/123456:a1b2c3d4-e5f6-7890-abcd-ef1234567890::"
		model["state"] = "active"
		model["description"] = "Migrate production Classic workloads to VPC"
		model["created_at"] = "2024-06-01T08:00:00.000Z"
		model["updated_at"] = "2024-06-10T12:00:00.000Z"

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.Migration)
	model.ID = core.StringPtr("mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab")
	model.Name = core.StringPtr("prod-classic-to-vpc")
	model.BrsCrn = core.StringPtr("crn:v1:bluemix:public:backup-recovery:us-south:a/123456:abcdef01-2345-6789-abcd-ef0123456789::")
	model.Crn = core.StringPtr("crn:v1:bluemix:public:brs-migration:us-south:a/123456:a1b2c3d4-e5f6-7890-abcd-ef1234567890::")
	model.State = core.StringPtr("active")
	model.Description = core.StringPtr("Migrate production Classic workloads to VPC")
	model.CreatedAt = CreateMockDateTime("2024-06-01T08:00:00.000Z")
	model.UpdatedAt = CreateMockDateTime("2024-06-10T12:00:00.000Z")

	result, err := brsmigration.DataSourceIbmBrsMigrationsMigrationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

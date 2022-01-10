// Copyright IBM Corp. 2021, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudant_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/cloudant"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

func TestAccIBMCloudantDatabaseBasic(t *testing.T) {
	var conf cloudantv1.DatabaseInformation
	instanceName := fmt.Sprintf("tf_instance_%d", acctest.RandIntRange(10, 100))
	db := fmt.Sprintf("tf_db_%d", acctest.RandIntRange(10, 100))
	dbUpdate := fmt.Sprintf("tf_db_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCloudantDatabaseDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCloudantDatabaseConfigBasic(instanceName, db),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCloudantDatabaseExists("ibm_cloudant_database.cloudant_database", conf),
					resource.TestCheckResourceAttr("ibm_cloudant_database.cloudant_database", "db", db),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCloudantDatabaseConfigBasic(instanceName, dbUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloudant_database.cloudant_database", "db", dbUpdate),
				),
			},
		},
	})
}

func TestAccIBMCloudantDatabaseAllArgs(t *testing.T) {
	var conf cloudantv1.DatabaseInformation
	instanceName := fmt.Sprintf("tf_instance_%d", acctest.RandIntRange(10, 100))
	db := fmt.Sprintf("tf_db_%d", acctest.RandIntRange(10, 100))
	partitioned := "true"
	shards := "16"
	dbUpdate := fmt.Sprintf("tf_db_%d", acctest.RandIntRange(10, 100))
	partitionedUpdate := "true"
	shardsUpdate := "16"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCloudantDatabaseDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCloudantDatabaseConfig(instanceName, db, partitioned, shards),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCloudantDatabaseExists("ibm_cloudant_database.cloudant_database", conf),
					resource.TestCheckResourceAttr("ibm_cloudant_database.cloudant_database", "db", db),
					resource.TestCheckResourceAttr("ibm_cloudant_database.cloudant_database", "partitioned", partitioned),
					resource.TestCheckResourceAttr("ibm_cloudant_database.cloudant_database", "shards", shards),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCloudantDatabaseConfig(instanceName, dbUpdate, partitionedUpdate, shardsUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloudant_database.cloudant_database", "db", dbUpdate),
					resource.TestCheckResourceAttr("ibm_cloudant_database.cloudant_database", "partitioned", partitionedUpdate),
					resource.TestCheckResourceAttr("ibm_cloudant_database.cloudant_database", "shards", shardsUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cloudant_database.cloudant_database",
				ImportState:       true,
				ImportStateVerify: true, ImportStateVerifyIgnore: []string{
					"partitioned", "shards"},
			},
		},
	})
}

func testAccCheckIBMCloudantDatabaseConfigBasic(instanceName, db string) string {
	return fmt.Sprintf(`

		data "ibm_resource_group" "cloudant" {
			is_default=true
		}

		resource "ibm_cloudant" "cloudant_instance" {
			name              = "%s"
			plan              = "standard"
			location          = "us-south"
			resource_group_id = data.ibm_resource_group.cloudant.id
		}

		resource "ibm_cloudant_database" "cloudant_database" {
			instance_crn = ibm_cloudant.cloudant_instance.crn
			db = "%s"
		}
	`, instanceName, db)
}

func testAccCheckIBMCloudantDatabaseConfig(instanceName, db string, partitioned string, shards string) string {
	return fmt.Sprintf(`

		data "ibm_resource_group" "cloudant" {
			is_default=true
		}

		resource "ibm_cloudant" "cloudant_instance" {
			name              = "%s"
			plan              = "standard"
			location          = "us-south"
			resource_group_id = data.ibm_resource_group.cloudant.id
		}

		resource "ibm_cloudant_database" "cloudant_database" {
			instance_crn = ibm_cloudant.cloudant_instance.crn
			db = "%s"
			partitioned = %s
			shards = %s
		}
	`, instanceName, db, partitioned, shards)
}

func testAccCheckIBMCloudantDatabaseExists(n string, obj cloudantv1.DatabaseInformation) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		instanceCRN := rs.Primary.Attributes["instance_crn"]
		cUrl, err := cloudant.GetCloudantInstanceUrl(instanceCRN, acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}

		cloudantClient, err := cloudant.GetCloudantClientForUrl(cUrl, acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}

		dbName := rs.Primary.Attributes["db"]
		getDatabaseInformationOptions := cloudantClient.NewGetDatabaseInformationOptions(dbName)

		documentResult, _, err := cloudantClient.GetDatabaseInformation(getDatabaseInformationOptions)
		if err != nil {
			return err
		}

		obj = *documentResult
		return nil
	}
}

func testAccCheckIBMCloudantDatabaseDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cloudant_database" {
			continue
		}

		instanceCRN := rs.Primary.Attributes["instance_crn"]
		cUrl, err := cloudant.GetCloudantInstanceUrl(instanceCRN, acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}

		cloudantClient, err := cloudant.GetCloudantClientForUrl(cUrl, acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}

		dbName := rs.Primary.Attributes["db"]
		getDatabaseInformationOptions := cloudantClient.NewGetDatabaseInformationOptions(dbName)

		// Try to find the key
		_, _, err = cloudantClient.GetDatabaseInformation(getDatabaseInformationOptions)
		if err == nil {
			return fmt.Errorf("cloudant_database still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

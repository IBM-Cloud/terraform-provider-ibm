// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

func TestAccIbmCloudantDatabaseBasic(t *testing.T) {
	var conf cloudantv1.DatabaseInformation
	instanceName := fmt.Sprintf("tf_instance_%d", acctest.RandIntRange(10, 100))
	db := fmt.Sprintf("tf_db_%d", acctest.RandIntRange(10, 100))
	dbUpdate := fmt.Sprintf("tf_db_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmCloudantDatabaseDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCloudantDatabaseConfigBasic(instanceName, db),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCloudantDatabaseExists("ibm_cloudant_database.cloudant_database", conf),
					resource.TestCheckResourceAttr("ibm_cloudant_database.cloudant_database", "db", db),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmCloudantDatabaseConfigBasic(instanceName, dbUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloudant_database.cloudant_database", "db", dbUpdate),
				),
			},
		},
	})
}

func TestAccIbmCloudantDatabaseAllArgs(t *testing.T) {
	var conf cloudantv1.DatabaseInformation
	instanceName := fmt.Sprintf("tf_instance_%d", acctest.RandIntRange(10, 100))
	db := fmt.Sprintf("tf_db_%d", acctest.RandIntRange(10, 100))
	partitioned := "true"
	q := "0"
	dbUpdate := fmt.Sprintf("tf_db_%d", acctest.RandIntRange(10, 100))
	partitionedUpdate := "true"
	qUpdate := "0"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmCloudantDatabaseDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCloudantDatabaseConfig(instanceName, db, partitioned, q),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCloudantDatabaseExists("ibm_cloudant_database.cloudant_database", conf),
					resource.TestCheckResourceAttr("ibm_cloudant_database.cloudant_database", "db", db),
					resource.TestCheckResourceAttr("ibm_cloudant_database.cloudant_database", "partitioned", partitioned),
					resource.TestCheckResourceAttr("ibm_cloudant_database.cloudant_database", "q", q),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmCloudantDatabaseConfig(instanceName, dbUpdate, partitionedUpdate, qUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloudant_database.cloudant_database", "db", dbUpdate),
					resource.TestCheckResourceAttr("ibm_cloudant_database.cloudant_database", "partitioned", partitionedUpdate),
					resource.TestCheckResourceAttr("ibm_cloudant_database.cloudant_database", "q", qUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cloudant_database.cloudant_database",
				ImportState:       true,
				ImportStateVerify: true, ImportStateVerifyIgnore: []string{
					"partitioned", "q"},
			},
		},
	})
}

func testAccCheckIbmCloudantDatabaseConfigBasic(instanceName, db string) string {
	return fmt.Sprintf(`

		data "ibm_resource_group" "cloudant" {
			is_default=true
	  	}
  
	  	resource "ibm_resource_instance" "cloudant_instance" {
			name              = "%s"
			service           = "cloudantnosqldb"
			plan              = "standard"
			location          = "us-east"
			resource_group_id = data.ibm_resource_group.cloudant.id
	  	}

		resource "ibm_cloudant_database" "cloudant_database" {
			cloudant_guid = ibm_resource_instance.cloudant_instance.guid
			db = "%s"
		}
	`, instanceName, db)
}

func testAccCheckIbmCloudantDatabaseConfig(instanceName, db string, partitioned string, q string) string {
	return fmt.Sprintf(`

	  	data "ibm_resource_group" "cloudant" {
			is_default=true
	  	}
	  
	  	resource "ibm_resource_instance" "cloudant_instance" {
			name              = "%s"
			service           = "cloudantnosqldb"
			plan              = "standard"
			location          = "us-east"
			resource_group_id = data.ibm_resource_group.cloudant.id
	  	}

		resource "ibm_cloudant_database" "cloudant_database" {
			cloudant_guid = ibm_resource_instance.cloudant_instance.guid
			db = "%s"
			partitioned = %s
			q = %s
		}
	`, instanceName, db, partitioned, q)
}

func testAccCheckIbmCloudantDatabaseExists(n string, obj cloudantv1.DatabaseInformation) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cloudantClient, err := testAccProvider.Meta().(ClientSession).CloudantV1()
		if err != nil {
			return err
		}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		cUrl, err := getCloudantInstanceUrl(parts[0], testAccProvider.Meta())
		if err != nil {
			return err
		}
		cloudantClient.Service.Options.URL = cUrl

		getDatabaseInformationOptions := &cloudantv1.GetDatabaseInformationOptions{}
		getDatabaseInformationOptions.SetDb(parts[1])

		documentResult, _, err := cloudantClient.GetDatabaseInformation(getDatabaseInformationOptions)
		if err != nil {
			return err
		}

		obj = *documentResult
		return nil
	}
}

func testAccCheckIbmCloudantDatabaseDestroy(s *terraform.State) error {
	cloudantClient, err := testAccProvider.Meta().(ClientSession).CloudantV1()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cloudant_database" {
			continue
		}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		cUrl, err := getCloudantInstanceUrl(parts[0], testAccProvider.Meta())
		if err != nil {
			return err
		}
		cloudantClient.Service.Options.URL = cUrl

		getDatabaseInformationOptions := &cloudantv1.GetDatabaseInformationOptions{}
		getDatabaseInformationOptions.SetDb(parts[1])

		// Try to find the key
		_, _, err = cloudantClient.GetDatabaseInformation(getDatabaseInformationOptions)
		if err == nil {
			return fmt.Errorf("cloudant_database still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

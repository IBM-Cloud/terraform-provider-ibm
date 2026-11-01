// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.2-f568b7d6-20261016-181021
*/

package secretsmanagerinstancemanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmDestinationDataSourceBasic(t *testing.T) {
	ibmCloudDatabaseDestinationInstanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	ibmCloudDatabaseDestinationName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	ibmCloudDatabaseDestinationType := "ibm_cloud_database"
	ibmCloudDatabaseDestinationCrn := fmt.Sprintf("tf_crn_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmDestinationDataSourceConfigBasic(ibmCloudDatabaseDestinationInstanceID, ibmCloudDatabaseDestinationName, ibmCloudDatabaseDestinationType, ibmCloudDatabaseDestinationCrn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "destination_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "crn"),
				),
			},
		},
	})
}

func TestAccIbmSmDestinationDataSourceAllArgs(t *testing.T) {
	ibmCloudDatabaseDestinationInstanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	ibmCloudDatabaseDestinationName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	ibmCloudDatabaseDestinationType := "ibm_cloud_database"
	ibmCloudDatabaseDestinationDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	ibmCloudDatabaseDestinationCrn := fmt.Sprintf("tf_crn_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmDestinationDataSourceConfig(ibmCloudDatabaseDestinationInstanceID, ibmCloudDatabaseDestinationName, ibmCloudDatabaseDestinationType, ibmCloudDatabaseDestinationDescription, ibmCloudDatabaseDestinationCrn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "destination_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "message"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destination.sm_destination_instance", "crn"),
				),
			},
		},
	})
}

func testAccCheckIbmSmDestinationDataSourceConfigBasic(ibmCloudDatabaseDestinationInstanceID string, ibmCloudDatabaseDestinationName string, ibmCloudDatabaseDestinationType string, ibmCloudDatabaseDestinationCrn string) string {
	return fmt.Sprintf(`
		resource "ibm_sm_destination" "sm_destination_instance" {
			instance_id = "%s"
			name = "%s"
			type = "%s"
			crn = "%s"
		}

		data "ibm_sm_destination" "sm_destination_instance" {
			instance_id = ibm_sm_destination.sm_destination_instance.instance_id
			destination_id = ibm_sm_destination.sm_destination_instance.destination_id
		}
	`, ibmCloudDatabaseDestinationInstanceID, ibmCloudDatabaseDestinationName, ibmCloudDatabaseDestinationType, ibmCloudDatabaseDestinationCrn)
}

func testAccCheckIbmSmDestinationDataSourceConfig(ibmCloudDatabaseDestinationInstanceID string, ibmCloudDatabaseDestinationName string, ibmCloudDatabaseDestinationType string, ibmCloudDatabaseDestinationDescription string, ibmCloudDatabaseDestinationCrn string) string {
	return fmt.Sprintf(`
		resource "ibm_sm_destination" "sm_destination_instance" {
			instance_id = "%s"
			name = "%s"
			type = "%s"
			description = "%s"
			crn = "%s"
		}

		data "ibm_sm_destination" "sm_destination_instance" {
			instance_id = ibm_sm_destination.sm_destination_instance.instance_id
			destination_id = ibm_sm_destination.sm_destination_instance.destination_id
		}
	`, ibmCloudDatabaseDestinationInstanceID, ibmCloudDatabaseDestinationName, ibmCloudDatabaseDestinationType, ibmCloudDatabaseDestinationDescription, ibmCloudDatabaseDestinationCrn)
}

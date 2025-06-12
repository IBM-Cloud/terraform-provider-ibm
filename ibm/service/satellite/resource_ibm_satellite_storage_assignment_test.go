// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM-Cloud/container-services-go-sdk/kubernetesserviceapiv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccSatelliteStorageAssignment_Basic(t *testing.T) {
	assignment_name := fmt.Sprintf("tf_assignment_%d", acctest.RandIntRange(10, 100))
	controller := "test-controller"
	config := "test-odf-remote"
	cluster := "test-cluster"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckSatelliteStorageConfigurationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSatelliteStorageAssignmentCreate(controller, assignment_name, config, cluster),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_satellite_storage_assignment", "assignment.#", "1"),
				),
			},
			{
				Config: testAccCheckSatelliteStorageAssignmentUpdate(controller, assignment_name, config, cluster),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_satellite_storage_assignment", "assignment.#", "2"),
				),
			},
			{
				ResourceName:      "ibm_satellite_storage_assignment.assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckSatelliteStorageAssignmentDestroy(s *terraform.State) error {
	satClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SatelliteClientSession()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_satellite_storage_assignment" {
			continue
		}

		ID := rs.Primary.ID
		getAssignmentOptions := &kubernetesserviceapiv1.GetAssignmentOptions{
			UUID: &ID,
		}

		_, _, err = satClient.GetAssignment(getAssignmentOptions)
		if err == nil {
			return fmt.Errorf("Storage Assignment still exists: %s", rs.Primary.ID)
		}

	}
	return nil
}

func testAccCheckSatelliteStorageAssignmentUpdate(controller string, assignment_name string, config string, cluster string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "res_group" {
		is_default = true
	}

	resource "ibm_satellite_storage_assignment" "assignment" {
		assignment_name = %s
		cluster = %s
		config = %s
		controller = %s
		update_config_revision = true
	}
	
	  
`, assignment_name, cluster, config, controller)
}

func testAccCheckSatelliteStorageAssignmentCreate(controller string, assignment_name string, config string, cluster string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "res_group" {
		is_default = true
	}

	resource "ibm_satellite_storage_assignment" "assignment" {
		assignment_name = %s
		cluster = %s
		config = %s
		controller = %s
	}
	
	  
`, assignment_name, cluster, config, controller)
}

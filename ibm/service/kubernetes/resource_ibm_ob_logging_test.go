// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMLogging_basic(t *testing.T) {
	clusterName := fmt.Sprintf("tf-clusters-%d", acctest.RandIntRange(10, 100))
	loggingName := fmt.Sprintf("tf-logging2-%d", acctest.RandIntRange(10, 100))
	ingestionKeyName := fmt.Sprintf("tf-keys-%d", acctest.RandIntRange(10, 100))
	updatedIngestionKey := fmt.Sprintf("tf-keys-updated-%d", acctest.RandIntRange(10, 100))
	updatedLoggingName := fmt.Sprintf("tf-logging-updated2-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLoggingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMLoggingBasic(clusterName, loggingName, ingestionKeyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMLoggingExists("ibm_ob_logging.test2"),
					resource.TestCheckResourceAttr(
						"ibm_ob_logging.test2", "instance_name", loggingName),
					resource.TestCheckResourceAttr(
						"ibm_ob_logging.test2", "private_endpoint", "false"),
				),
			},
			{
				Config: testAccCheckIBMLoggingUpdate(clusterName, loggingName, ingestionKeyName, updatedLoggingName, updatedIngestionKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_ob_logging.test2", "instance_name", updatedLoggingName),
					resource.TestCheckResourceAttr(
						"ibm_ob_logging.test2", "private_endpoint", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMLoggingExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		csClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
		if err != nil {
			return err
		}
		loggingID := rs.Primary.ID
		parts, err := flex.IdParts(loggingID)
		if err != nil {
			return err
		}

		clusterName := parts[0]
		instanceID := parts[1]

		targetEnv, err := getLoggingTarget()
		if err != nil {
			return err
		}

		_, err = csClient.Logging().GetLoggingConfig(clusterName, instanceID, targetEnv)

		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIBMLoggingDestroy(s *terraform.State) error {
	csClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_ob_logging" {
			continue
		}

		loggingID := rs.Primary.ID
		parts, err := flex.IdParts(loggingID)
		if err != nil {
			return err
		}

		clusterName := parts[0]
		instanceID := parts[1]

		targetEnv, err := getLoggingTarget()
		if err != nil {
			return err
		}

		_, err = csClient.Logging().GetLoggingConfig(clusterName, instanceID, targetEnv)

		if err == nil {
			return fmt.Errorf("Logging instance still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func getLoggingTarget() (v2.LoggingTargetHeader, error) {

	userDetails, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return v2.LoggingTargetHeader{}, err
	}

	accountID := userDetails.UserAccount

	targetEnv := v2.LoggingTargetHeader{
		AccountID: accountID,
	}

	return targetEnv, nil
}

func testAccCheckIBMLoggingBasic(clusterName, loggingName, ingestionKeyName string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "testacc_ds_resource_group" {
        is_default = "true"
      }
      
      resource "ibm_container_cluster" "testacc_cluster" {
        name       = "%s"
        datacenter = "%s"
        resource_group_id = data.ibm_resource_group.testacc_ds_resource_group.id
        default_pool_size = 1
		wait_till         = "MasterNodeReady"
        hardware        = "shared"
        machine_type    = "%s"
        timeouts {
          create = "720m"
          update = "720m"
        }
    } 

    resource "ibm_resource_instance" "instance" {
        name     = "%s"
        service  = "logdna"
        plan     = "7-day"
        location = "us-south"
    }
	
	resource "ibm_resource_key" "resourceKey" {
		name                 = "%s"
		resource_instance_id = ibm_resource_instance.instance.id
		role                 = "Manager"
	}

	resource "ibm_ob_logging" "test2" {
		depends_on = [ibm_resource_key.resourceKey]
        cluster = ibm_container_cluster.testacc_cluster.id
        instance_id = ibm_resource_instance.instance.guid
    }`, clusterName, acc.Datacenter, acc.MachineType, loggingName, ingestionKeyName)
}

func testAccCheckIBMLoggingUpdate(clusterName, loggingName, ingestionKeyName, instanceName, keyName string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "testacc_ds_resource_group" {
        is_default = "true"
      }
      
      resource "ibm_container_cluster" "testacc_cluster" {
        name       = "%s"
        datacenter = "%s"
        resource_group_id = data.ibm_resource_group.testacc_ds_resource_group.id
        default_pool_size = 1
		wait_till         = "MasterNodeReady"
        hardware        = "shared"
        machine_type    = "%s"
        timeouts {
          create = "720m"
          update = "720m"
        }
    } 

    resource "ibm_resource_instance" "instance" {
        name     = "%s"
        service  = "logdna"
        plan     = "7-day"
        location = "us-south"
    }
	
	resource "ibm_resource_key" "resourceKey" {
		name                 = "%s"
		resource_instance_id = ibm_resource_instance.instance.id
		role                 = "Manager"
	}

	resource "ibm_resource_instance" "instance2" {
		name     = "%s"
		service  = "logdna"
		plan     = "7-day"
		location = "us-south"
	}

	resource "ibm_resource_key" "resourceKey2" {
		name                 = "%s"
		resource_instance_id = ibm_resource_instance.instance2.id
		role                 = "Manager"
	}

	resource "ibm_ob_logging" "test2" {
		depends_on = [ibm_resource_key.resourceKey]
        cluster = ibm_container_cluster.testacc_cluster.id
        instance_id = ibm_resource_instance.instance2.guid
		private_endpoint = true
	}`, clusterName, acc.Datacenter, acc.MachineType, loggingName, ingestionKeyName, instanceName, keyName)
}

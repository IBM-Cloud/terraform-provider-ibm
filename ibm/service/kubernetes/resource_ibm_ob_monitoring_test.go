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
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMMonitoring_basic(t *testing.T) {
	clusterName := fmt.Sprintf("tf-cluster-%d", acctest.RandIntRange(10, 100))
	monitoringName := fmt.Sprintf("tf-monitoring-%d", acctest.RandIntRange(10, 100))
	ingestionKeyName := fmt.Sprintf("tf-key-%d", acctest.RandIntRange(10, 100))
	updatedIngestionKey := fmt.Sprintf("tf-key-updated-%d", acctest.RandIntRange(10, 100))
	updatedMonitoringName := fmt.Sprintf("tf-monitoring-updated-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMMonitoringDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMMonitoringBasic(clusterName, monitoringName, ingestionKeyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMMonitoringExists("ibm_ob_monitoring.test2"),
					resource.TestCheckResourceAttr(
						"ibm_ob_monitoring.test2", "instance_name", monitoringName),
					resource.TestCheckResourceAttr(
						"ibm_ob_monitoring.test2", "private_endpoint", "false"),
				),
			},
			{
				Config: testAccCheckIBMMonitoringUpdate(clusterName, monitoringName, ingestionKeyName, updatedMonitoringName, updatedIngestionKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_ob_monitoring.test2", "instance_name", updatedMonitoringName),
					resource.TestCheckResourceAttr(
						"ibm_ob_monitoring.test2", "private_endpoint", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMMonitoringExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		csClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
		if err != nil {
			return err
		}
		monitoringID := rs.Primary.ID
		parts, err := flex.IdParts(monitoringID)
		if err != nil {
			return err
		}

		clusterName := parts[0]
		instanceID := parts[1]

		targetEnv, err := getMonitoringTarget()
		if err != nil {
			return err
		}

		_, err = csClient.Monitoring().GetMonitoringConfig(clusterName, instanceID, targetEnv)

		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIBMMonitoringDestroy(s *terraform.State) error {
	csClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_ob_monitoring" {
			continue
		}

		monitoringID := rs.Primary.ID
		parts, err := flex.IdParts(monitoringID)
		if err != nil {
			return err
		}

		clusterName := parts[0]
		instanceID := parts[1]

		targetEnv, err := getMonitoringTarget()
		if err != nil {
			return err
		}

		_, err = csClient.Monitoring().GetMonitoringConfig(clusterName, instanceID, targetEnv)

		if err == nil {
			return fmt.Errorf("Monitoring instance still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func getMonitoringTarget() (v2.MonitoringTargetHeader, error) {

	userDetails, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return v2.MonitoringTargetHeader{}, err
	}

	accountID := userDetails.UserAccount

	targetEnv := v2.MonitoringTargetHeader{
		AccountID: accountID,
	}

	return targetEnv, nil
}

func testAccCheckIBMMonitoringBasic(clusterName, monitorName, ingestionKeyName string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "testacc_ds_resource_group" {
        name = "Default"
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
        service  = "sysdig-monitor"
        plan     = "graduated-tier"
        location = "us-south"
    }
	
	resource "ibm_resource_key" "resourceKey" {
		name                 = "%s"
		resource_instance_id = ibm_resource_instance.instance.id
		role                 = "Manager"
	}

	resource "ibm_ob_monitoring" "test2" {
		depends_on = [ibm_resource_key.resourceKey]
        cluster = ibm_container_cluster.testacc_cluster.id
        instance_id = ibm_resource_instance.instance.guid
    }`, clusterName, acc.Datacenter, acc.MachineType, monitorName, ingestionKeyName)
}

func testAccCheckIBMMonitoringUpdate(clusterName, monitorName, ingestionKeyName, instanceName, keyName string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "testacc_ds_resource_group" {
        name = "Default"
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
        service  = "sysdig-monitor"
        plan     = "graduated-tier"
        location = "us-south"
    }
	
	resource "ibm_resource_key" "resourceKey" {
		name                 = "%s"
		resource_instance_id = ibm_resource_instance.instance.id
		role                 = "Manager"
	}

	resource "ibm_resource_instance" "instance2" {
		name     = "%s"
		service  = "sysdig-monitor"
		plan     = "graduated-tier"
		location = "us-south"
	}

	resource "ibm_resource_key" "resourceKey2" {
		name                 = "%s"
		resource_instance_id = ibm_resource_instance.instance2.id
		role                 = "Manager"
	}

	resource "ibm_ob_monitoring" "test2" {
		depends_on = [ibm_resource_key.resourceKey]
        cluster = ibm_container_cluster.testacc_cluster.id
        instance_id = ibm_resource_instance.instance2.guid
		private_endpoint = true
	}`, clusterName, acc.Datacenter, acc.MachineType, monitorName, ingestionKeyName, instanceName, keyName)
}

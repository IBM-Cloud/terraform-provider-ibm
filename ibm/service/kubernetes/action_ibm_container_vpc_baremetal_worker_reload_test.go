// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMContainerVpcBareMetalWorkerReloadBasic(t *testing.T) {
	clusterName := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		ExternalProviders: map[string]resource.ExternalProvider{
			"null": {
				Source:            "hashicorp/null",
				VersionConstraint: "~> 3.0",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcBareMetalWorkerReloadBasic(vpcname, subnetname, clusterName, "mx3d.metal.64x512", "1h", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", clusterName),
					testAccCheckIBMContainerVpcWorkerState(clusterName, "deployed"),
				),
			},
		},
	})
}

func TestAccIBMContainerVpcBareMetalWorkerReloadInvalidTimeoutFormat(t *testing.T) {
	clusterName := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	timeout := "10abc"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		ExternalProviders: map[string]resource.ExternalProvider{
			"null": {
				Source:            "hashicorp/null",
				VersionConstraint: "~> 3.0",
			},
		},

		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMContainerVpcBareMetalWorkerReloadBasic(vpcname, subnetname, clusterName, "mx3d.metal.64x512", timeout, false),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`Invalid Timeout Format`),
			},
		},
	})
}

func TestAccIBMContainerVpcBareMetalWorkerReloadNoWait(t *testing.T) {
	clusterName := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		ExternalProviders: map[string]resource.ExternalProvider{
			"null": {
				Source:            "hashicorp/null",
				VersionConstraint: "~> 3.0",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcBareMetalWorkerReloadBasic(vpcname, subnetname, clusterName, "mx3d.metal.64x512", "1h", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_cluster.cluster", "name", clusterName),
					func(s *terraform.State) error {
						time.Sleep(5 * time.Second) //wait for 5 seconds for the worker to start reload
						return nil
					},
					testAccCheckIBMContainerVpcWorkerState(clusterName, "undeploying", "undeployed", "reload_pending", "reloading"),
				),
			},
		},
	})
}

// testAccCheckIBMContainerVpcWorkerState checks if the worker is in any of the expected states
func testAccCheckIBMContainerVpcWorkerState(clusterName string, expectedStates ...string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
		if err != nil {
			return fmt.Errorf("Error getting VPC container API client: %s", err)
		}

		vpcWorkerClient := vpcClient.Workers()

		workers, err := vpcWorkerClient.ListWorkers(clusterName, false, containerv2.ClusterTargetHeader{})
		if err != nil {
			return fmt.Errorf("Error listing workers for cluster %s: %s", clusterName, err)
		}

		if len(workers) == 0 {
			return fmt.Errorf("No workers found in cluster %s", clusterName)
		}

		workerID := workers[0].ID
		worker, err := vpcWorkerClient.Get(clusterName, workerID, containerv2.ClusterTargetHeader{})
		if err != nil {
			return fmt.Errorf("Error retrieving worker %s from cluster %s: %s", workerID, clusterName, err)
		}

		actualState := worker.LifeCycle.ActualState
		for _, expectedState := range expectedStates {
			if actualState == expectedState {
				return nil
			}
		}

		return fmt.Errorf("Worker %s is not in any of the expected states %v. Current state: %s", workerID, expectedStates, actualState)
	}
}

func testAccCheckIBMContainerVpcBareMetalWorkerReloadBasic(vpcname, subnetname, clusterName, flavor, customTimeout string, noWait bool) string {
	timeout := ""
	if customTimeout != "" {
		timeout = fmt.Sprintf(`timeout = "%s"`, customTimeout)
	}
	noWaitConfig := ""
	if noWait {
		noWaitConfig = fmt.Sprintf(`no_wait = "%s"`, "true")
	}
	return fmt.Sprintf(`
	data "ibm_resource_group" "resource_group" {
		is_default = true
	}

	resource "ibm_is_vpc" "vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "subnet" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.vpc.id
		zone                     = "us-south-1"
		total_ipv4_address_count = 256
	}

	resource "ibm_container_vpc_cluster" "cluster" {
		name              = "%s"
		vpc_id            = ibm_is_vpc.vpc.id
		flavor            = "%s"
		worker_count      = 1
		resource_group_id = data.ibm_resource_group.resource_group.id
		wait_till         = "OneWorkerNodeReady"
		zones {
			subnet_id = ibm_is_subnet.subnet.id
			name      = "us-south-1"
		}
	}

	# Get the worker from the cluster
	data "ibm_container_vpc_cluster" "cluster_data" {
		name              = ibm_container_vpc_cluster.cluster.name
		resource_group_id = data.ibm_resource_group.resource_group.id
	}

	# Reload the first bare metal worker using the worker ID (which is the bare metal server ID for bare metal workers)
	action "ibm_container_vpc_bare_metal_worker_reload" "test_reload" {
		config {
			cluster_name_id           = ibm_container_vpc_cluster.cluster.id
			bare_metal_server_id = data.ibm_container_vpc_cluster.cluster_data.workers[0]
			%s
			%s
		}
	}

	resource "null_resource" "trigger_reload" {
		provisioner "local-exec" {
			command = "echo 'All infrastructure ready. Triggering bare metal worker reload now...'"
		}

		lifecycle {
			action_trigger {
				events  = [after_create]
				actions = [action.ibm_container_vpc_bare_metal_worker_reload.test_reload]
			}
		}
	}
	`, vpcname, subnetname, clusterName, flavor, timeout, noWaitConfig)
}

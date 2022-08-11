// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
)

func TestAccIBMContainerVpcClusterWorkerPoolBasic(t *testing.T) {

	name := fmt.Sprintf("tf-vpc-worker-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerWorkerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMVpcContainerWorkerPoolBasic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "flavor", "cx2.2x4"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "zones.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "labels.%", "2"),
				),
			},
			{
				Config: testAccCheckIBMVpcContainerWorkerPoolUpdate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "flavor", "cx2.2x4"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "zones.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "labels.%", "3"),
				),
			},
			{
				ResourceName:      "ibm_container_vpc_worker_pool.test_pool",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMContainerVpcClusterWorkerPoolDedicatedHost(t *testing.T) {

	name := fmt.Sprintf("tf-vpc-worker-%d", acctest.RandIntRange(10, 100))
	hostpoolID := acc.HostPoolID
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMVpcContainerWorkerPoolDedicatedHostCreate(
					acc.ClusterName,
					name,
					"bx2d.4x16",
					acc.IksClusterSubnetID,
					acc.IksClusterVpcID,
					acc.IksClusterResourceGroupID,
					hostpoolID,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.vpc_worker_pool", "host_pool_id", hostpoolID),
				),
			},
		},
	})
}

func testAccCheckIBMVpcContainerWorkerPoolDestroy(s *terraform.State) error {

	wpClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_vpc_worker_pool" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cluster := parts[0]
		workerPoolID := parts[1]

		target := v2.ClusterTargetHeader{}

		// Try to find the key
		wp, err := wpClient.WorkerPools().GetWorkerPool(cluster, workerPoolID, target)

		if err == nil {
			if wp.ActualState == "deleted" && wp.DesiredState == "deleted" {
				return nil
			}
			return fmt.Errorf("Worker pool still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("[ERROR] Error waiting for worker pool (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMVpcContainerWorkerPoolBasic(name string) string {
	return fmt.Sprintf(`
	provider "ibm" {
		region="eu-de"
	}
	data "ibm_resource_group" "resource_group" {
		is_default=true
	}
	resource "ibm_is_vpc" "vpc" {
	  name = "%[1]s"
	}
	
	resource "ibm_is_subnet" "subnet1" {
	  name                     = "%[1]s-1"
	  vpc                      = ibm_is_vpc.vpc.id
	  zone                     = "eu-de-1"
	  total_ipv4_address_count = 256
	}
	
	resource "ibm_is_subnet" "subnet2" {
	  name                     = "%[1]s-2"
	  vpc                      = ibm_is_vpc.vpc.id
	  zone                     = "eu-de-2"
	  total_ipv4_address_count = 256
	}
	
	resource "ibm_container_vpc_cluster" "cluster" {
	  name              = "%[1]s"
	  vpc_id            = ibm_is_vpc.vpc.id
	  flavor            = "cx2.2x4"
	  worker_count      = 1
	  resource_group_id = data.ibm_resource_group.resource_group.id
	  wait_till         = "MasterNodeReady"
	  zones {
		subnet_id = ibm_is_subnet.subnet1.id
		name      = "eu-de-1"
	  }
	}
	resource "ibm_container_vpc_worker_pool" "test_pool" {
	  cluster           = ibm_container_vpc_cluster.cluster.id
	  worker_pool_name  = "%[1]s"
	  flavor            = "cx2.2x4"
	  vpc_id            = ibm_is_vpc.vpc.id
	  worker_count      = 1
	  resource_group_id = data.ibm_resource_group.resource_group.id
	  zones {
		name      = "eu-de-2"
		subnet_id = ibm_is_subnet.subnet2.id
	  }
	  labels = {
		"test"  = "test-pool"
		"test1" = "test-pool1"
	  }
	}
		`, name)
}

func testAccCheckIBMVpcContainerWorkerPoolUpdate(name string) string {
	return fmt.Sprintf(`
	provider "ibm" {
		region="eu-de"
	}
	data "ibm_resource_group" "resource_group" {
		is_default=true
	}
	resource "ibm_is_vpc" "vpc" {
	  name = "%[1]s"
	}
	resource "ibm_is_subnet" "subnet1" {
	  name                     = "%[1]s-1"
	  vpc                      = ibm_is_vpc.vpc.id
	  zone                     = "eu-de-1"
	  total_ipv4_address_count = 256
	}
	resource "ibm_is_subnet" "subnet2" {
	  name                     = "%[1]s-2"
	  vpc                      = ibm_is_vpc.vpc.id
	  zone                     = "eu-de-2"
	  total_ipv4_address_count = 256
	}
	resource "ibm_container_vpc_cluster" "cluster" {
	  name              = "%[1]s"
	  vpc_id            = ibm_is_vpc.vpc.id
	  flavor            = "cx2.2x4"
	  worker_count      = 1
	  resource_group_id = data.ibm_resource_group.resource_group.id
	  wait_till         = "MasterNodeReady"
	  zones {
		subnet_id = ibm_is_subnet.subnet1.id
		name      = "eu-de-1"
	  }
	}
	resource "ibm_container_vpc_worker_pool" "test_pool" {
	  cluster           = ibm_container_vpc_cluster.cluster.id
	  worker_pool_name  = "%[1]s"
	  flavor            = "cx2.2x4"
	  vpc_id            = ibm_is_vpc.vpc.id
	  worker_count      = 1
	  resource_group_id = data.ibm_resource_group.resource_group.id
	  zones {
		name      = "eu-de-2"
		subnet_id = ibm_is_subnet.subnet2.id
	  }
	  zones {
		subnet_id = ibm_is_subnet.subnet1.id
		name      = "eu-de-1"
	  }
	  labels = {
		"test"  = "test-pool"
		"test1" = "test-pool1"
		"test2" = "test-pool2"
	  }
	}
		`, name)
}

func TestAccIBMContainerVpcClusterWorkerPoolEnvvar(t *testing.T) {

	name := fmt.Sprintf("tf-vpc-worker-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerWorkerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMVpcContainerWorkerPoolEnvvar(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "flavor", "bx2.4x16"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "zones.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "kms_instance_id", acc.KmsInstanceID),
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "crk", acc.CrkID),
				),
			},
			{
				ResourceName:      "ibm_container_vpc_worker_pool.test_pool",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"kms_instance_id", "crk"},
			},
		},
	})
}

func testAccCheckIBMVpcContainerWorkerPoolDedicatedHostCreate(clusterName, name, flavor, subnetID, vpcID, rgroupID, hostpoolID string) string {
	return fmt.Sprintf(`
	resource "ibm_container_vpc_worker_pool" "vpc_worker_pool" {
		cluster = "%s"
		flavor = "%s"
		worker_pool_name = "%s"
		zones {
		  subnet_id = "%s"
		  name      = "us-south-1"
		}
		worker_count      = 1
		vpc_id = "%s"
		resource_group_id = "%s"
		host_pool_id      = "%s"
	  }
	`, clusterName, flavor, name, subnetID, vpcID, rgroupID, hostpoolID)
}

func testAccCheckIBMVpcContainerWorkerPoolEnvvar(name string) string {
	return fmt.Sprintf(`
	resource "ibm_container_vpc_worker_pool" "test_pool" {
	  cluster           = "%[2]s"
	  worker_pool_name  = "%[1]s"
	  flavor            = "bx2.4x16"
	  vpc_id            = "%[3]s"
	  worker_count      = 1
	  zones {
		subnet_id = "%[4]s"
		name      = "us-south-1"
	  }
	  kms_instance_id = "%[5]s"
	  crk = "%[6]s"
	}
		`, name, acc.IksClusterID, acc.IksClusterVpcID, acc.IksClusterSubnetID, acc.KmsInstanceID, acc.CrkID)
}

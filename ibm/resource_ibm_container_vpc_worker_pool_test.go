// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
)

func TestAccIBMContainerVpcClusterWorkerPoolBasic(t *testing.T) {

	name := fmt.Sprintf("tf-vpc-worker-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
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

func testAccCheckIBMVpcContainerWorkerPoolDestroy(s *terraform.State) error {

	wpClient, err := testAccProvider.Meta().(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_vpc_worker_pool" {
			continue
		}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cluster := parts[0]
		workerPoolID := parts[1]

		target := v2.ClusterTargetHeader{}

		// Try to find the key
		_, err = wpClient.WorkerPools().GetWorkerPool(cluster, workerPoolID, target)

		if err == nil {
			return fmt.Errorf("Worker pool still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for worker pool (%s) to be destroyed: %s", rs.Primary.ID, err)
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

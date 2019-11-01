package ibm

import (
	"fmt"
	"strings"
	"testing"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
<<<<<<< HEAD
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
=======
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
>>>>>>> 643f668... Add support for terraform v0.12
)

func TestAccIBMVpcContainerWorkerPool_basic(t *testing.T) {

	flavor := "c2.2x4"
	worker_count := 1
	name1 := acctest.RandInt()
	name2 := acctest.RandInt()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerWorkerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMVpcContainerWorkerPool_basic(flavor, worker_count, name1, name2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_worker_pool.test_pool", "flavor", flavor),
				),
			},
		},
	})
}

func TestAccIBMVpcContainerWorkerPool_importBasic(t *testing.T) {
	flavor := "c2.2x4"
	worker_count := 1
	name1 := acctest.RandInt()
	name2 := acctest.RandInt()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerWorkerPoolDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMVpcContainerWorkerPool_basic(flavor, worker_count, name1, name2),
			},

			resource.TestStep{
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

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for worker pool (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMVpcContainerWorkerPool_basic(flavor string, worker_count, name1, name2 int) string {
	return fmt.Sprintf(`
	provider "ibm" {
		generation = 1
	  }
	  
	  variable "name1" {
		default = "%d"
	  }
	  variable "name2" {
		default = "%d"
	  }
	 
	  locals {
		ZONE1 = "us-south-1"
		ZONE2 = "us-south-2"
	  }
	  
	  resource "ibm_is_vpc" "vpc1" {
		name = "vpc-${var.name1}"
	  }
	  
	  resource "ibm_is_subnet" "subnet1" {
		name                     = "subnet-${var.name1}"
		vpc                      = "${ibm_is_vpc.vpc1.id}"
		zone                     = "${local.ZONE1}"
		total_ipv4_address_count = 256
	  }
	  
	  resource "ibm_is_subnet" "subnet2" {
		name                     = "subnet-${var.name2}"
		vpc                      = "${ibm_is_vpc.vpc1.id}"
		zone                     = "${local.ZONE2}"
		total_ipv4_address_count = 256
	  }
	  
	  data "ibm_resource_group" "resource_group" {
		name = "Default"
	  }
	  
	  resource "ibm_container_vpc_cluster" "cluster" {
		name              = "cluster${var.name1}"
		vpc_id            = "${ibm_is_vpc.vpc1.id}"
		flavor            = "%s"
		worker_count      = "%d"
		resource_group_id = "${data.ibm_resource_group.resource_group.id}"
	  
		zones = [
		  {
			subnet_id = "${ibm_is_subnet.subnet1.id}"
			name      = "${local.ZONE1}"
		  },
		]
	  }
	  
	  resource "ibm_container_vpc_worker_pool" "cluster_pool" {
		cluster          = "${ibm_container_vpc_cluster.cluster.id}"
		worker_pool_name = "workerpool${var.name1}"
		flavor           = "%s"
		vpc_id           = "${ibm_is_vpc.vpc1.id}"
		worker_count     = "%d"
		resource_group_id = "${data.ibm_resource_group.resource_group.id}"
		zones = [
		  {
			name      = "${local.ZONE2}"
			subnet_id = "${ibm_is_subnet.subnet2.id}"
		  },
		]
	  }
	  
		`, name1, name2, flavor, worker_count, flavor, worker_count)
}

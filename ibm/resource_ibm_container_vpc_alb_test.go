package ibm

import (
	"fmt"
	"strings"
	"testing"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMVpcContainerALB_Basic(t *testing.T) {
	flavor := "c2.2x4"
	worker_count := 1
	name1 := acctest.RandIntRange(10, 100)
	name2 := acctest.RandIntRange(10, 100)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerALBDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMVpcContainerALB_basic(true, flavor, worker_count, name1, name2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_alb.alb", "enable", "true"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMVpcContainerALB_basic(false, flavor, worker_count, name1, name2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_alb.alb", "enable", "false"),
				),
			},
		},
	})
}

func testAccCheckIBMVpcContainerALBDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_vpc_alb" {
			continue
		}

		albID := rs.Primary.ID
		targetEnv := v2.ClusterTargetHeader{}

		csClient, err := testAccProvider.Meta().(ClientSession).VpcContainerAPI()
		if err != nil {
			return err
		}
		albAPI := csClient.Albs()
		_, err = albAPI.GetAlb(albID, targetEnv)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMVpcContainerALB_basic(enable bool, flavor string, worker_count, name1, name2 int) string {
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
	  
	  resource ibm_container_vpc_alb alb {
		alb_id = "${ibm_container_vpc_cluster.cluster.id}"
		enable = "%t"
	  }
	  `, name1, name2, flavor, worker_count, enable)
}

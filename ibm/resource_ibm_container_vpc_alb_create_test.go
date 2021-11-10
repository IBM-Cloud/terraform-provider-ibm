// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
)

func TestAccIBMContainerVPCClusterALBCreate(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-alb-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerALBCreateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMVpcContainerALBCreate(true, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_container_vpc_alb_create.alb", "enable_by_default", "true"),
					resource.TestCheckResourceAttr("ibm_container_vpc_alb_create.alb", "cluster", name),
					resource.TestCheckResourceAttr("ibm_container_vpc_alb_create.alb", "zone", "dal10"),
					resource.TestCheckResourceAttr("ibm_container_vpc_alb_create.alb", "type", "private"),
				),
			},
		},
	})
}

func testAccCheckIBMVpcContainerALBCreateDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_vpc_alb_create" {
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

		if err == nil {
			return fmt.Errorf("Instance still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMVpcContainerALBCreate(enable bool, name string) string {
	return fmt.Sprintf(`
	provider "ibm" {
		region="us-south"
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
	  zone                     = "dal10"
	  total_ipv4_address_count = 256
	}
	resource "ibm_container_vpc_cluster" "cluster" {
		name              = "%[1]s"
		vpc_id            = ibm_is_vpc.vpc.id
		flavor            = "cx2.2x4"
		worker_count      = 1
		resource_group_id = data.ibm_resource_group.resource_group.id
		zones {
			subnet_id = ibm_is_subnet.subnet1.id
			name      = "dal10"
		}
	}
	resource ibm_container_vpc_alb_create alb {
		cluster = "%[1]s"
		type = "private"
		zone = "dal10"
		enable_by_default = "%t"
	}
	`, name, enable)
}

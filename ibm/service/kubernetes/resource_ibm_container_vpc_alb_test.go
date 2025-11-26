// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
)

func TestAccIBMContainerVPCClusterALBBasic(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-alb-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerALBDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMVpcContainerALBBasic(true, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_vpc_alb.alb", "enable", "true"),
				),
			},
			{
				Config: testAccCheckIBMVpcContainerALBBasic(false, name),
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

		csClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
		if err != nil {
			return err
		}
		albAPI := csClient.Albs()
		_, err = albAPI.GetAlb(albID, targetEnv)

		if err == nil {
			return fmt.Errorf("Instance still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("[ERROR] Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMVpcContainerALBBasic(enable bool, name string) string {
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
	resource "ibm_container_vpc_cluster" "cluster" {
		name              = "%[1]s"
		vpc_id            = ibm_is_vpc.vpc.id
		flavor            = "cx2.2x4"
		worker_count      = 1
		resource_group_id = data.ibm_resource_group.resource_group.id
		zones {
			subnet_id = ibm_is_subnet.subnet1.id
			name      = "eu-de-1"
		}
	}
	  resource ibm_container_vpc_alb alb {
		alb_id = "${ibm_container_vpc_cluster.cluster.albs.0.id}"
		enable = "%t"
	  }
	  `, name, enable)
}

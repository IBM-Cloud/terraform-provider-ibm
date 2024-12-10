// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
)

func TestAccIBMContainerVpcClusterWorkerBasic(t *testing.T) {

	name := fmt.Sprintf("tf-vpc-worker-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerWorkerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMVpcContainerWorkerBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMVpcContainerExists(),
				),
			},
			{
				ResourceName:      "ibm_container_vpc_worker.test_worker",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMVpcContainerWorkerDestroy(s *terraform.State) error {

	//Destroy basically does nothing in this resource
	return nil
}

func testAccCheckIBMVpcContainerWorkerBasic(name string) string {
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

    data ibm_container_cluster_config "cluster_config: {
        cluster_name_id   = "%[1]s"
        resource_group_id = data.ibm_resource_group.resource_group.id
    }

    resource "ibm_container_vpc_worker" "test_worker" {
        cluster_name        = "%[1]s"
        replace_worker      = element(ibm_container_vpc_cluster.cluster.workers, 0)
        resource_group_id   = data.ibm_resource_group.resource_group.id
        kube_config_path    = data.ibm_container_cluster_config.cluster_config.config_file_path
        check_ptx_status    = false
    }  
		`, name)
}

func testAccCheckIBMVpcContainerExists() resource.TestCheckFunc {
	return func(s *terraform.State) error {

		wpClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
		if err != nil {
			return err
		}

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_container_vpc_worker" {
				continue
			}

			parts, err := flex.SepIdParts(rs.Primary.ID, "-")
			if err != nil {
				return err
			}
			if len(parts) < 2 {
				return fmt.Errorf("[ERROR] Incorrect ID %s: Id should be in kube-clusterID-* format", rs.Primary.ID)
			}
			cluster := parts[1]

			target := v2.ClusterTargetHeader{}

			_, err = wpClient.Workers().Get(cluster, rs.Primary.ID, target)
			if err != nil {
				return fmt.Errorf("[ERROR] Error getting container vpc worker node: %s", err)
			}
			return nil
		}
		return fmt.Errorf("[ERROR] Error getting container vpc worker resource")
	}
}

// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"log"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	v2 "github.com/Mavrickk3/bluemix-go/api/container/containerv2"
)

func TestAccIBMContainerVPCClusterALBCreate(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-alb-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMVpcContainerALBCreateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMVpcContainerALBCreate(true, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_container_vpc_alb_create.alb", "enable", "true"),
					resource.TestCheckResourceAttr("ibm_container_vpc_alb_create.alb", "zone", "us-south-1"),
					resource.TestCheckResourceAttr("ibm_container_vpc_alb_create.alb", "type", "private"),
				),
			},
		},
	})
}

func testAccCheckIBMVpcContainerALBCreateDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "ibm_container_vpc_alb_create" {
			albID := rs.Primary.ID
			targetEnv := v2.ClusterTargetHeader{
				ResourceGroup: acc.IksClusterResourceGroupID,
			}

			csClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
			if err != nil {
				return err
			}
			albAPI := csClient.Albs()
			albconfig, err := albAPI.GetAlb(albID, targetEnv)
			if err != nil {
				return err
			}
			log.Println("[WARN] No API to delete ALB : ", albconfig)
		}
		if rs.Type == "ibm_container_vpc_cluster" {
			csClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
			if err != nil {
				return err
			}
			targetEnv := v2.ClusterTargetHeader{
				ResourceGroup: acc.IksClusterResourceGroupID,
			}
			// Try to find the key
			_, err = csClient.Clusters().GetCluster(rs.Primary.ID, targetEnv)

			if err == nil {
				log.Printf("[DEBUG] ibm_container_vpc_cluster Cluster still exists: %s", rs.Primary.ID)
				return fmt.Errorf("Cluster still exists: %s", rs.Primary.ID)
			} else if !strings.Contains(err.Error(), "404") {
				log.Printf("[DEBUG] ibm_container_vpc_cluster Error waiting for cluster (%s) to be destroyed: %s", rs.Primary.ID, err)
				return fmt.Errorf("[ERROR] Error waiting for cluster (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
			log.Println("[DEBUG] ibm_container_vpc_cluster deleted")
		}

	}
	return nil
}

// You need to set up:
// IBM_CLUSTER_VPC_ID
// IBM_CLUSTER_VPC_SUBNET_ID
// IBM_CLUSTER_VPC_RESOURCE_GROUP_ID
func testAccCheckIBMVpcContainerALBCreate(enable bool, name string) string {
	config := fmt.Sprintf(`

	resource "ibm_container_vpc_cluster" "cluster" {
		name              = "%[1]s"
		vpc_id            = "%[3]s"
		flavor            = "cx2.2x4"
		worker_count      = 1
		resource_group_id = "%[4]s"
		zones {
			subnet_id = "%[5]s"
			name      = "us-south-1"
		}
	}
	resource ibm_container_vpc_alb_create alb {
		cluster = ibm_container_vpc_cluster.cluster.id
		type = "private"
		zone = "us-south-1"
		resource_group_id = "%[4]s"
		enable = "%[2]t"
	}
	`, name, enable, acc.IksClusterVpcID, acc.IksClusterResourceGroupID, acc.IksClusterSubnetID)
	fmt.Println(config)
	return config
}

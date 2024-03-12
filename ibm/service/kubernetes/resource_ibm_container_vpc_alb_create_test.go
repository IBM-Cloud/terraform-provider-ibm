// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMContainerVPCClusterALBCreate(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-alb-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		//CheckDestroy: testAccCheckIBMVpcContainerALBCreateDestroy,
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

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

func TestAccIBMContainerVPCClusterALBBasic(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-alb-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
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

func testAccCheckIBMVpcContainerALBBasic(enable bool, name string) string {
	return fmt.Sprintf(`
	resource "ibm_container_vpc_cluster" "cluster" {
		name              = "%[1]s"
		vpc_id            = "%[2]s"
		flavor            = "bx2.4x16"
		worker_count      = 1
		resource_group_id = "%[3]s"
		zones {
			subnet_id = "%[4]s"
			name      = "us-south-1"
		}
		wait_till = "IngressReady"
		ingress_config {
			ingress_status_report {
				enabled = true
				ignored_errors = ["ERRHPAIWC"]
			}
			ingress_health_checker_enabled = true
		}
	}
	  resource ibm_container_vpc_alb alb {
		alb_id = "${ibm_container_vpc_cluster.cluster.albs.0.id}"
		enable = "%[5]t"
	  }
	  `, name, acc.IksClusterVpcID, acc.IksClusterResourceGroupID, acc.IksClusterSubnetID, enable)
}

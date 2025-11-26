// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMContainerALB_Create(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerALBCreateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerALBCreate(clusterName, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_container_alb_create.alb", "enable", "true"),
					resource.TestCheckResourceAttr("ibm_container_alb_create.alb", "alb_type", "private"),
					resource.TestCheckResourceAttr("ibm_container_alb_create.alb", "vlan_id", acc.PrivateVlanID),
					resource.TestCheckResourceAttr("ibm_container_alb_create.alb", "zone", acc.Zone),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckIBMContainerALBCreateDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_alb_create" {
			continue
		}

		albID := rs.Primary.ID
		targetEnv := v1.ClusterTargetHeader{
			Region: "us-south",
		}

		csClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContainerAPI()
		if err != nil {
			return err
		}
		albAPI := csClient.Albs()
		resp, err := albAPI.GetALB(albID, targetEnv)

		if err == nil {
			return fmt.Errorf("Instance still exists: %s, values: %v", rs.Primary.ID, resp)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("[ERROR] Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

// you need to set up the followings
// IBM_PUBLIC_VLAN_ID
// IBM_PRIVATE_VLAN_ID
// IBM_DATACENTER
// IBM_WORKER_POOL_ZONE

func testAccCheckIBMContainerALBCreate(clusterName string, enable bool) string {
	config := fmt.Sprintf(`resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"
  default_pool_size = 1
  machine_type    = "%s"
  hardware        = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  timeouts {
    create = "120m"
	update = "120m"
  }
}

resource "ibm_container_alb_create" "alb" {
  cluster=ibm_container_cluster.testacc_cluster.name
  enable = "%t"
  alb_type = "private"
  vlan_id = ibm_container_cluster.testacc_cluster.private_vlan_id
  zone = "%s"
}`, clusterName, acc.Datacenter, acc.MachineType, acc.PublicVlanID, acc.PrivateVlanID, enable, acc.Zone)
	fmt.Println(config)
	return config
}

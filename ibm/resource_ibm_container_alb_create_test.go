// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"strings"
	"testing"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMContainerALB_Create(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerALBCreateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerALBCreate(clusterName, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_container_alb_create.alb", "enable_by_default", "true"),
					resource.TestCheckResourceAttr("ibm_container_alb_create.alb", "type", "private"),
					resource.TestCheckResourceAttr("ibm_container_alb_create.alb", "vlan_id", privateVlanID),
					resource.TestCheckResourceAttr("ibm_container_alb_create.alb", "zone", zone),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckIBMContainerALBCreate(clusterName, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_container_alb_create.alb", "enableByDefault", "false"),
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

		fmt.Println("DESTROY ibm_container_alb_create")
		albID := rs.Primary.ID
		fmt.Println("DESTROY ibm_container_alb_create ID: ", albID)
		targetEnv := v1.ClusterTargetHeader{
			Region: "us-south",
		}

		csClient, err := testAccProvider.Meta().(ClientSession).ContainerAPI()
		if err != nil {
			return err
		}
		albAPI := csClient.Albs()
		resp, err := albAPI.GetALB(albID, targetEnv)
		fmt.Println("DESTROY resp: ", resp)

		if err == nil {
			return fmt.Errorf("Instance still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

// func testAccCheckIBMContainerALBCreate(clusterName string, enable bool) string {
// 	config := fmt.Sprintf(`resource "ibm_container_alb_create" "alb" {
//   enable_by_default = "false"
//   type = "private"
//   vlan_id = "1900403"
//   zone = "dal10"
//   cluster="terraform-manual-test"
// }`)
// 	fmt.Println(config)
// 	return config
// }

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
    create = "720m"
	update = "720m"
  }
  subnet_id       = ["%s", "%s"]
}

resource "ibm_container_alb_create" "alb" {
  enable_by_default = "%t"
  type = "private"
  vlan_id = "%[5]s"
  zone = "%[7]s"
  cluster=ibm_container_cluster.testacc_cluster.id
}`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, privateSubnetID, publicSubnetID, enable, zone)
	fmt.Println(config)
	return config
}

// export IBM_PRIVATE_SUBNET_ID=2597322
// export IBM_PRIVATE_VLAN_ID=1900403
// export IBM_PUBLIC_SUBNET_ID=1576763
// export IBM_PUBLIC_VLAN_ID=1900401

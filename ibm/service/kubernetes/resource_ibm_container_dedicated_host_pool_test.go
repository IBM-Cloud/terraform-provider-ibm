// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMContainerDedicatedHostPoolBasic(t *testing.T) {

	dedicatedHostPoolName := fmt.Sprintf("tf-dedicated-host-pool-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerDedicatedHostPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerDedicatedHostPoolBasic(dedicatedHostPoolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_dedicated_host_pool.test_dhostpool", "name", dedicatedHostPoolName),
					resource.TestCheckResourceAttr(
						"ibm_container_dedicated_host_pool.test_dhostpool", "metro", "dal"),
					resource.TestCheckResourceAttr(
						"ibm_container_dedicated_host_pool.test_dhostpool", "flavor_class", "bx2d"),
					resource.TestCheckResourceAttr(
						"ibm_container_dedicated_host_pool.test_dhostpool", "host_count", "0"),
					resource.TestCheckResourceAttr(
						"ibm_container_dedicated_host_pool.test_dhostpool", "state", "created"),
					resource.TestCheckResourceAttr(
						"ibm_container_dedicated_host_pool.test_dhostpool", "zones.#", "0"),
				),
			},
			{
				ResourceName:      "ibm_container_dedicated_host_pool.test_dhostpool",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMContainerDedicatedHostPoolDestroy(s *terraform.State) error {

	client, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}
	dedicatedHostPoolAPI := client.DedicatedHostPool()
	targetEnv := v2.ClusterTargetHeader{}

	var retryCounter int = 0
	var returnErr error = nil
	for retryCounter < 2 {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_container_dedicated_host_pool" {
				continue
			}

			dhp, err := dedicatedHostPoolAPI.GetDedicatedHostPool(rs.Primary.ID, targetEnv)

			if err == nil {
				if dhp.State != "deleted" {
					returnErr = fmt.Errorf("Dedicated host pool still exists: %s", rs.Primary.ID)
					continue
				}
				return nil
			} else if !strings.Contains(err.Error(), "404") {
				return fmt.Errorf("[ERROR] Error waiting for dedicated host pool (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
		retryCounter++
		time.Sleep(time.Second)
	}

	return returnErr
}

func testAccCheckIBMContainerDedicatedHostPoolBasic(dedicatedHostPoolName string) string {
	return fmt.Sprintf(`
resource "ibm_container_dedicated_host_pool" "test_dhostpool" {
	name         = "%s"
	flavor_class = "bx2d"
	metro        = "dal"
}
`, dedicatedHostPoolName)
}

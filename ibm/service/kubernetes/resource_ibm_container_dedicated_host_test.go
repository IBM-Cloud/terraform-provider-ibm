// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"testing"
	"time"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	v2 "github.com/Mavrickk3/bluemix-go/api/container/containerv2"
	"github.com/Mavrickk3/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMContainerDedicatedHostBasic(t *testing.T) {

	dedicatedHostPoolName := fmt.Sprintf("tf-dedicated-host-pool-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerDedicatedHostDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerDedicatedHostBasic(dedicatedHostPoolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_dedicated_host.test_dhost", "placement_enabled", "true"),
					resource.TestCheckResourceAttrSet(
						"ibm_container_dedicated_host.test_dhost", "id"),
					resource.TestCheckResourceAttr(
						"ibm_container_dedicated_host.test_dhost", "life_cycle.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_dedicated_host.test_dhost", "resources.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_dedicated_host.test_dhost", "workers.#", "0"),
				),
			},
			{
				Config: testAccCheckIBMContainerDedicatedHostDisable(dedicatedHostPoolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_dedicated_host.test_dhost", "placement_enabled", "false"),
				),
			},
			{
				Config: testAccCheckIBMContainerDedicatedHostEnable(dedicatedHostPoolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_dedicated_host.test_dhost", "placement_enabled", "true"),
				),
			},
			{
				ResourceName:      "ibm_container_dedicated_host.test_dhost",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMContainerDedicatedHostDestroy(s *terraform.State) error {

	client, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}
	dedicatedHostAPI := client.DedicatedHost()
	targetEnv := v2.ClusterTargetHeader{}

	var (
		dhostID     string
		dhostPoolID string
	)

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "ibm_container_dedicated_host" {
			dhostID = rs.Primary.ID
		} else if rs.Type == "ibm_container_dedicated_host_pool" {
			dhostPoolID = rs.Primary.ID
		}
	}

	var retryCounter int = 0
	var returnErr error = nil
	for retryCounter < 2 {
		dhp, err := dedicatedHostAPI.GetDedicatedHost(dhostID, dhostPoolID, targetEnv)
		if err == nil {
			if dhp.Lifecycle.ActualState != "deleted" {
				returnErr = fmt.Errorf("Dedicated host still exists, dhostid %s, dhostpoolid %s", dhostID, dhostPoolID)
				continue
			}
			return nil
		} else {
			if apiErr, ok := err.(bmxerror.RequestFailure); ok {
				if apiErr.StatusCode() == 404 {
					return nil
				}
			}
			returnErr = err
		}

		retryCounter++
		time.Sleep(time.Second)
	}

	return returnErr
}

func testAccCheckIBMContainerDedicatedHostBasic(dedicatedHostPoolName string) string {
	return testAccCheckIBMContainerDedicatedHostPoolBasic(dedicatedHostPoolName) + `
resource "ibm_container_dedicated_host" "test_dhost" {
	flavor         = "bx2d.host.152x608"
	host_pool_id   = ibm_container_dedicated_host_pool.test_dhostpool.id
	zone           = "us-south-2"
}
`
}

func testAccCheckIBMContainerDedicatedHostDisable(dedicatedHostPoolName string) string {
	return testAccCheckIBMContainerDedicatedHostPoolBasic(dedicatedHostPoolName) + `
resource "ibm_container_dedicated_host" "test_dhost" {
	flavor            = "bx2d.host.152x608"
	host_pool_id      = ibm_container_dedicated_host_pool.test_dhostpool.id
	zone              = "us-south-2"
	placement_enabled = "false"
}
`
}

func testAccCheckIBMContainerDedicatedHostEnable(dedicatedHostPoolName string) string {
	return testAccCheckIBMContainerDedicatedHostPoolBasic(dedicatedHostPoolName) + `
resource "ibm_container_dedicated_host" "test_dhost" {
	flavor            = "bx2d.host.152x608"
	host_pool_id      = ibm_container_dedicated_host_pool.test_dhostpool.id
	zone              = "us-south-2"
	placement_enabled = "true"
}
`
}

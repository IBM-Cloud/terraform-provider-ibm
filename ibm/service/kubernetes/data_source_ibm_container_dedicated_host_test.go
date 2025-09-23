// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMContainerDedicatedHostDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-dedicated-host-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerDedicatedHostDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_container_dedicated_host.test_dhost_2", "placement_enabled", "true"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_dedicated_host.test_dhost_2", "id"),
					resource.TestCheckResourceAttr(
						"data.ibm_container_dedicated_host.test_dhost_2", "life_cycle.#", "1"),
					resource.TestCheckResourceAttr(
						"data.ibm_container_dedicated_host.test_dhost_2", "resources.#", "1"),
					resource.TestCheckResourceAttr(
						"data.ibm_container_dedicated_host.test_dhost_2", "workers.#", "0"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerDedicatedHostDataSourceConfig(name string) string {
	return testAccCheckIBMContainerDedicatedHostBasic(name) + `
	data "ibm_container_dedicated_host" "test_dhost_2" {
	    host_id = ibm_container_dedicated_host.test_dhost.host_id
	    host_pool_id = ibm_container_dedicated_host.test_dhost.host_pool_id
	}
`
}

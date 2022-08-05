// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMContainerDedicatedHostPoolDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-dedicated-host-pool-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerDedicatedHostPoolDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_dedicated_host_pool.test_dhostpool_2", "name", name),
					resource.TestCheckResourceAttr("data.ibm_container_dedicated_host_pool.test_dhostpool_2", "metro", "dal"),
					resource.TestCheckResourceAttr("data.ibm_container_dedicated_host_pool.test_dhostpool_2", "flavor_class", "bx2d"),
					resource.TestCheckResourceAttr("data.ibm_container_dedicated_host_pool.test_dhostpool_2", "host_count", "0"),
					resource.TestCheckResourceAttr("data.ibm_container_dedicated_host_pool.test_dhostpool_2", "state", "created"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerDedicatedHostPoolDataSourceConfig(name string) string {
	return testAccCheckIBMContainerDedicatedHostPoolBasic(name) + `
	data "ibm_container_dedicated_host_pool" "test_dhostpool_2" {
	    host_pool_id = ibm_container_dedicated_host_pool.test_dhostpool.id
	}
`
}

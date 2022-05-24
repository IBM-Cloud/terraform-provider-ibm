// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMContainerDedicatedHostFlavorDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerDedicatedHostFlavorDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_dedicated_host_flavor.test_dhost_flavor", "flavor_class", "bx2d"),
					resource.TestCheckResourceAttr("data.ibm_container_dedicated_host_flavor.test_dhost_flavor", "region", "us-south"),
					resource.TestCheckResourceAttr("data.ibm_container_dedicated_host_flavor.test_dhost_flavor", "deprecated", "false"),
					resource.TestCheckResourceAttr("data.ibm_container_dedicated_host_flavor.test_dhost_flavor", "max_vcpus", "152"),
					resource.TestCheckResourceAttr("data.ibm_container_dedicated_host_flavor.test_dhost_flavor", "max_memory", "0"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerDedicatedHostFlavorDataSourceConfig() string {
	return `
	data "ibm_container_dedicated_host_flavor" "test_dhost_flavor" {
	    host_flavor_id = "bx2d.host.152x608"
		zone           = "us-south-1"
	}
`
}

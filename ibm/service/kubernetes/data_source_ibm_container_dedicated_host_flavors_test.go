// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMContainerDedicatedHostFlavorsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerDedicatedHostFlavorsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_dedicated_host_flavors.test_dhost_flavors", "host_flavors.#", "1"),
					resource.TestCheckResourceAttrSet("data.ibm_container_dedicated_host_flavors.test_dhost_flavors", "host_flavors.0.host_flavor_id"),
					resource.TestCheckResourceAttrSet("data.ibm_container_dedicated_host_flavors.test_dhost_flavors", "host_flavors.0.flavor_class"),
					resource.TestCheckResourceAttrSet("data.ibm_container_dedicated_host_flavors.test_dhost_flavors", "host_flavors.0.region"),
					resource.TestCheckResourceAttrSet("data.ibm_container_dedicated_host_flavors.test_dhost_flavors", "host_flavors.0.deprecated"),
					resource.TestCheckResourceAttrSet("data.ibm_container_dedicated_host_flavors.test_dhost_flavors", "host_flavors.0.max_vcpus"),
					resource.TestCheckResourceAttrSet("data.ibm_container_dedicated_host_flavors.test_dhost_flavors", "host_flavors.0.max_memory"),
					resource.TestCheckResourceAttr("data.ibm_container_dedicated_host_flavors.test_dhost_flavors", "host_flavors.0.instance_storage.#", "1"),
					resource.TestCheckResourceAttrSet("data.ibm_container_dedicated_host_flavors.test_dhost_flavors", "host_flavors.0.instance_storage.0.count"),
					resource.TestCheckResourceAttrSet("data.ibm_container_dedicated_host_flavors.test_dhost_flavors", "host_flavors.0.instance_storage.0.size"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerDedicatedHostFlavorsDataSourceConfig() string {
	return `
	data "ibm_container_dedicated_host_flavors" "test_dhost_flavors" {
	  zone           = "us-south-1"
	}
`
}

// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmIsDedicatedHostDSBasic(t *testing.T) {
	//var conf vpcv1.DedicatedHost
	resName := "data.ibm_is_dedicated_host.dhost"
	grpname := fmt.Sprintf("tf-dhostgroup%d", acctest.RandIntRange(10, 100))
	dhname := fmt.Sprintf("tf-dhost%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmIsDedicatedHostDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIsDedicatedHostDSConfigBasic(dedicatedHostGroupClass, dedicatedHostGroupFamily, grpname, dedicatedHostProfileName, dhname),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "name"),
					resource.TestCheckResourceAttrSet(resName, "zone"),
					resource.TestCheckResourceAttrSet(resName, "host_group"),
					resource.TestCheckResourceAttr(resName, "disks.#", "2"),
					resource.TestCheckResourceAttrSet(resName, "disks.0.name"),
					resource.TestCheckResourceAttrSet(resName, "disks.0.size"),
				),
			},
		},
	})
}

func testAccCheckIbmIsDedicatedHostDSConfigBasic(class, family, grpname, profile, name string) string {
	return fmt.Sprintf(`
	
	data "ibm_resource_group" "default" {
		is_default=true
	}
	resource "ibm_is_dedicated_host_group" "dhgroup" {
		class = "%s"
		family = "%s"
		name = "%s"
		resource_group = data.ibm_resource_group.default.id
		zone = "us-south-2"
	}
	data "ibm_is_dedicated_host_group" "dgroup" {
		name = ibm_is_dedicated_host_group.dhgroup.name
	}

	resource "ibm_is_dedicated_host" "dedicated-host-test-01" {
		profile = "%s"
		host_group = data.ibm_is_dedicated_host_group.dgroup.id
		name = "%s"
	  }
	data "ibm_is_dedicated_host" "dhost"{
		name = "%s"
		host_group = ibm_is_dedicated_host.dedicated-host-test-01.host_group
	}
	`, class, family, grpname, profile, name, name)
}

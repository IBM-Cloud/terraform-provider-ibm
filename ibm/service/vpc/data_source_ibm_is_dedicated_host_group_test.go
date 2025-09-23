// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmIsDedicatedHostGroupDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tfdhgroup%d", acctest.RandIntRange(10, 100))
	resName := "data.ibm_is_dedicated_host_group.dgroup"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmIsDedicatedHostGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsDedicatedHostGroupDataSourceConfigBasic(acc.DedicatedHostGroupClass, acc.DedicatedHostGroupFamily, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", name),
					resource.TestCheckResourceAttr(resName, "class", acc.DedicatedHostGroupClass),
					resource.TestCheckResourceAttr(resName, "family", acc.DedicatedHostGroupFamily),
					resource.TestCheckResourceAttrSet(resName, "zone"),
				),
			},
		},
	})
}

func testAccCheckIbmIsDedicatedHostGroupDataSourceConfigBasic(class string, family string, name string) string {
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
	`, class, family, name)
}

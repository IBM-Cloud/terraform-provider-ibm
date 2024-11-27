// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIbmIsDedicatedHostGroupsDataSourceBasic(t *testing.T) {
	var conf vpcv1.DedicatedHostGroup
	name := fmt.Sprintf("tfdhgroups%d", acctest.RandIntRange(10, 100))
	resName := "data.ibm_is_dedicated_host_groups.dhgroups"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmIsDedicatedHostGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsDedicatedHostGroupConfigBasic(acc.DedicatedHostGroupClass, acc.DedicatedHostGroupFamily, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsDedicatedHostGroupExists("ibm_is_dedicated_host_group.is_dedicated_host_group", conf),
					resource.TestCheckResourceAttr("ibm_is_dedicated_host_group.is_dedicated_host_group", "class", acc.DedicatedHostGroupClass),
					resource.TestCheckResourceAttr("ibm_is_dedicated_host_group.is_dedicated_host_group", "family", acc.DedicatedHostGroupFamily),
					resource.TestCheckResourceAttr("ibm_is_dedicated_host_group.is_dedicated_host_group", "name", name),
				),
			},
			{
				Config: testAccCheckIbmIsDedicatedHostGroupsDataSourceConfigBasic(acc.DedicatedHostGroupClass, acc.DedicatedHostGroupFamily, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "host_groups.0.name"),
					resource.TestCheckResourceAttrSet(resName, "host_groups.0.class"),
					resource.TestCheckResourceAttrSet(resName, "host_groups.0.family"),
					resource.TestCheckResourceAttrSet(resName, "host_groups.0.zone"),
				),
			},
		},
	})
}

func testAccCheckIbmIsDedicatedHostGroupsDataSourceConfigBasic(class, family, name string) string {
	return testAccCheckIbmIsDedicatedHostGroupConfigBasic(class, family, name) + fmt.Sprintf(`
	
	data "ibm_is_dedicated_host_groups" "dhgroups" {
	}
	`)
}

// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIbmIsDedicatedHostsDSBasic(t *testing.T) {
	var conf vpcv1.DedicatedHost
	groupname := fmt.Sprintf("tfgroup%d", acctest.RandIntRange(10, 100))
	dhname := fmt.Sprintf("tfdhost%d", acctest.RandIntRange(10, 1000))
	dhresName := "ibm_is_dedicated_host.dhost"
	resName := "data.ibm_is_dedicated_hosts.dhosts"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmIsDedicatedHostDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsDedicatedHostConfigBasic(acc.DedicatedHostGroupClass, acc.DedicatedHostGroupFamily, groupname, acc.DedicatedHostProfileName, dhname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIbmIsDedicatedHostExists(dhresName, conf),
					resource.TestCheckResourceAttr(
						dhresName, "name", dhname),
				),
			},
			{
				Config: testAccCheckIbmIsDedicatedHostsDSConfigBasic(acc.DedicatedHostGroupClass, acc.DedicatedHostGroupFamily, groupname, acc.DedicatedHostProfileName, dhname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "dedicated_hosts.0.name"),
					resource.TestCheckResourceAttrSet(resName, "dedicated_hosts.0.available_memory"),
					resource.TestCheckResourceAttrSet(resName, "dedicated_hosts.0.memory"),
					resource.TestCheckResourceAttrSet(resName, "dedicated_hosts.0.host_group"),
					resource.TestCheckResourceAttrSet(resName, "dedicated_hosts.0.zone"),
					resource.TestCheckResourceAttrSet(resName, "dedicated_hosts.0.vcpu.0.manufacturer"),
				),
			},
		},
	})
}

func testAccCheckIbmIsDedicatedHostsDSConfigBasic(class string, family string, groupname string, profile string, dhname string) string {
	return testAccCheckIbmIsDedicatedHostConfigBasic(class, family, groupname, profile, dhname) + fmt.Sprintf(`
	 
	 data "ibm_is_dedicated_hosts" "dhosts"{
	 }
	 `)
}

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

func TestAccIBMISDedicatedHostDisksDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_dedicated_host_disks.test1"
	var conf vpcv1.DedicatedHost
	groupname := fmt.Sprintf("tf-dhostgroup%d", acctest.RandIntRange(10, 100))
	dhname := fmt.Sprintf("tf-dhost%d", acctest.RandIntRange(10, 100))
	dhresname := "ibm_is_dedicated_host.dhost"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsDedicatedHostConfigBasic(acc.DedicatedHostGroupClass, acc.DedicatedHostGroupFamily, groupname, acc.DedicatedHostProfileName, dhname),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsDedicatedHostExists(dhresname, conf),
					resource.TestCheckResourceAttr(dhresname, "name", dhname),
					resource.TestCheckResourceAttr(dhresname, "disks.#", "2"),
					resource.TestCheckResourceAttrSet(dhresname, "disks.0.name"),
					resource.TestCheckResourceAttrSet(dhresname, "disks.0.size"),
				),
			},
			{
				Config: testAccCheckIBMISDedicatedHostDisksDataSourceConfig(acc.DedicatedHostGroupClass, acc.DedicatedHostGroupFamily, groupname, acc.DedicatedHostProfileName, dhname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "disks.0.name"),
					resource.TestCheckResourceAttrSet(resName, "disks.0.size"),
				),
			},
		},
	})
}

func testAccCheckIBMISDedicatedHostDisksDataSourceConfig(class, family, groupname, DedicatedHostProfileName, dhname string) string {
	// status filter defaults to empty
	return testAccCheckIbmIsDedicatedHostConfigBasic(class, family, groupname, acc.DedicatedHostProfileName, dhname) + fmt.Sprintf(`
	  data "ibm_is_dedicated_host" "dhost" {
		name = "%s"
		host_group = ibm_is_dedicated_host.dhost.host_group
	  }
      data "ibm_is_dedicated_host_disks" "test1" {
		dedicated_host = data.ibm_is_dedicated_host.dhost.id
		
      }`, dhname)
}

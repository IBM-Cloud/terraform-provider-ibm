// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISDedicatedHostDiskDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_dedicated_host_disk.test1"
	var conf vpcv1.DedicatedHost
	groupname := fmt.Sprintf("tf-dhostgroup%d", acctest.RandIntRange(10, 100))
	dhname := fmt.Sprintf("tf-dhost%d", acctest.RandIntRange(10, 100))
	dhresname := "ibm_is_dedicated_host.dhost"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsDedicatedHostConfigBasic(dedicatedHostGroupClass, dedicatedHostGroupFamily, groupname, dedicatedHostProfileName, dhname),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsDedicatedHostExists(dhresname, conf),
					resource.TestCheckResourceAttr(dhresname, "name", dhname),
					resource.TestCheckResourceAttr(dhresname, "disks.#", "2"),
					resource.TestCheckResourceAttrSet(dhresname, "disks.0.name"),
					resource.TestCheckResourceAttrSet(dhresname, "disks.0.size"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMISDedicatedHostDiskDataSourceConfig(dedicatedHostGroupClass, dedicatedHostGroupFamily, groupname, dedicatedHostProfileName, dhname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "name"),
					resource.TestCheckResourceAttrSet(resName, "size"),
				),
			},
		},
	})
}

func testAccCheckIBMISDedicatedHostDiskDataSourceConfig(class, family, groupname, dedicatedHostProfileName, dhname string) string {
	return testAccCheckIbmIsDedicatedHostConfigBasic(class, family, groupname, dedicatedHostProfileName, dhname) + fmt.Sprintf(`
	data "ibm_is_dedicated_host" "dhost" {
		name = "%s"
		host_group = ibm_is_dedicated_host.dhost.host_group
	  }
      data "ibm_is_dedicated_host_disks" "test1" {
		dedicated_host = data.ibm_is_dedicated_host.dhost.id
		
      }
	  data "ibm_is_dedicated_host_disk" "test1" {
	    dedicated_host = data.ibm_is_dedicated_host.dhost.id
	    disk = data.ibm_is_dedicated_host_disks.test1.disks.0.id
	}`, dhname)
}

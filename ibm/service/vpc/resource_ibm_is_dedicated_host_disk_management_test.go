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

func TestAccIBMISDedicatedHostDiskManagement_basic(t *testing.T) {
	resName := "ibm_is_dedicated_host.dhost"
	var conf vpcv1.DedicatedHost
	groupname := fmt.Sprintf("tf-dhostgroup%d", acctest.RandIntRange(10, 100))
	dhname := fmt.Sprintf("tf-dhost%d", acctest.RandIntRange(10, 100))
	diskName := "tf-dhdisk01"
	diskResname := "data.ibm_is_dedicated_host_disk.test1"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsDedicatedHostConfigBasic(acc.DedicatedHostGroupClass, acc.DedicatedHostGroupFamily, groupname, acc.DedicatedHostProfileName, dhname),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsDedicatedHostExists(resName, conf),
					resource.TestCheckResourceAttr(resName, "name", dhname),
					resource.TestCheckResourceAttr(resName, "disks.#", "2"),
					resource.TestCheckResourceAttrSet(resName, "disks.0.name"),
					resource.TestCheckResourceAttrSet(resName, "disks.0.size"),
				),
			},
			{
				Config: testAccCheckIBMISDedicatedHostDiskManagementConfig(acc.DedicatedHostGroupClass, acc.DedicatedHostGroupFamily, groupname, acc.DedicatedHostProfileName, dhname, diskName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(diskResname, "name"),
					resource.TestCheckResourceAttrSet(diskResname, "size"),
					resource.TestCheckResourceAttr(diskResname, "name", diskName),
				),
			},
		},
	})
}

func testAccCheckIBMISDedicatedHostDiskManagementConfig(class, family, groupname, DedicatedHostProfileName, dhname, diskName string) string {
	return testAccCheckIbmIsDedicatedHostConfigBasic(class, family, groupname, acc.DedicatedHostProfileName, dhname) + fmt.Sprintf(`
	data "ibm_is_dedicated_host" "dhost" {
		name = "%s"
		host_group = ibm_is_dedicated_host.dhost.host_group
	  }
      data "ibm_is_dedicated_host_disks" "test1" {
		dedicated_host = data.ibm_is_dedicated_host.dhost.id
		
      }
	  data "ibm_is_dedicated_host_disk" "test1" {
	    dedicated_host = data.ibm_is_dedicated_host.dhost.id
	    disk = ibm_is_dedicated_host_disk_management.disks.disks.0.id
	  }
	  resource "ibm_is_dedicated_host_disk_management" "disks"{
		dedicated_host = data.ibm_is_dedicated_host.dhost.id
		disks {
		  name = "%s"
		  id = data.ibm_is_dedicated_host_disks.test1.disks.0.id
		}
	  }
	
	`, dhname, diskName)
}

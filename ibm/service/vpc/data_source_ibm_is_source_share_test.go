// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIbmIsSourceShareDataSourceBasic(t *testing.T) {
	shareName := fmt.Sprintf("tf-fs-name-%d", acctest.RandIntRange(10, 100))
	replicaName := fmt.Sprintf("tf-fsrep-name-%d", acctest.RandIntRange(10, 100))
	shareTargetName := fmt.Sprintf("tf-fs-tg-name-%d", acctest.RandIntRange(10, 100))
	shareTargetName1 := fmt.Sprintf("tf-fs-tg-name-%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tf-vpc-name-%d", acctest.RandIntRange(10, 100))
	vpcname1 := fmt.Sprintf("tf-vpc-name1-%d", acctest.RandIntRange(10, 100))
	size := acctest.RandIntRange(10, 50)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsSourceShareDataSourceConfigBasic(vpcname, vpcname1, shareName, size, shareTargetName, shareTargetName1, replicaName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_source_share.test", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_source_share.test", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_source_share.test", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_source_share.test", "encryption"),
					resource.TestCheckResourceAttrSet("data.ibm_is_source_share.test", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_source_share.test", "iops"),
					resource.TestCheckResourceAttrSet("data.ibm_is_source_share.test", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_source_share.test", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_source_share.test", "profile"),
					resource.TestCheckResourceAttrSet("data.ibm_is_source_share.test", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_is_source_share.test", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_source_share.test", "size"),
					resource.TestCheckResourceAttrSet("data.ibm_is_source_share.test", "zone"),
				),
			},
		},
	})
}

func testAccCheckIbmIsSourceShareDataSourceConfigBasic(vpcname, vpcname1, shareName string, size int, shareTargetName, shareTargetName1, replicaName string) string {
	return testAccCheckIbmIsShareConfigReplica(vpcname, vpcname1, shareName, size, shareTargetName, shareTargetName1, replicaName) + fmt.Sprintf(`
		
		data "ibm_is_source_share" "test" {
			share_replica = ibm_is_share.replica.id
		}
	`)
}
func TestAccIbmIsSourceShareDataSource404(t *testing.T) {
	srId := "8843-5fr454ft-f6-4565-9555-5f889f5f3f7777"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIbmIsSourceShareDataSourceConfig404(srId),
				ExpectError: regexp.MustCompile("GetShareSourceWithContext failed"),
			},
		},
	})
}

func testAccCheckIbmIsSourceShareDataSourceConfig404(srId string) string {
	return fmt.Sprintf(`
		
		data "ibm_is_source_share" "test" {
			share_replica = "%s"
		}
	`, srId)
}

// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMIsShareTargetDataSourceAllArgs(t *testing.T) {
	vpcName := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	targetName := fmt.Sprintf("tf-share-target-%d", acctest.RandIntRange(10, 100))
	shareName := fmt.Sprintf("tf-fs-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsShareTargetDataSourceConfigBasic(shareName, vpcName, targetName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_share_mount_target.is_share_target", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_mount_target.is_share_target", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_mount_target.is_share_target", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_mount_target.is_share_target", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_mount_target.is_share_target", "mount_path"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_mount_target.is_share_target", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_mount_target.is_share_target", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_mount_target.is_share_target", "vpc.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsShareTargetDataSourceConfigBasic(sname, vpcName, targetName string) string {
	return testAccCheckIbmIsShareTargetsDataSourceConfigBasic(sname, vpcName, targetName) + fmt.Sprintf(`
		
		data "ibm_is_share_mount_target" "is_share_target" {
			share = ibm_is_share.is_share.id
			share_target = data.ibm_is_mount_share_targets.is_share_targets.mount_targets.0.id
		}
	`)
}

func testAccCheckIbmIsShareTargetsDataSourceConfigBasic(sname, vpcName, targetName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_share" "is_share" {
			zone = "us-south-2"
			size = 200
			name = "%s"
			profile = "%s"
		}

		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		}

		resource "ibm_is_share_mount_target" "is_share_target" {
			share = ibm_is_share.is_share.id
			name = "%s"
			vpc = ibm_is_vpc.testacc_vpc.id
		}

		data "ibm_is_share_mount_targets" "is_share_targets" {
			share = ibm_is_share_mount_target.is_share_target.share
		}
	`, sname, acc.ShareProfileName, vpcName, targetName)
}

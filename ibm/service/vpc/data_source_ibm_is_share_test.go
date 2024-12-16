// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmIsShareDataSourceBasic(t *testing.T) {
	shareName := fmt.Sprintf("tf-fs-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsShareDataSourceConfigBasic(shareName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "encryption"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "iops"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "profile"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "size"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "zone"),
				),
			},
		},
	})
}
func TestAccIbmIsShareDataSource404(t *testing.T) {
	shareId := "8843-5fr454ft-f6-4565-9555-5f889f5f3f7777"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIbmIsShareDataSourceConfig404(shareId),
				ExpectError: regexp.MustCompile("GetShareWithContext failed"),
			},
		},
	})
}

func TestAccIbmIsShareDataSourceAllArgs(t *testing.T) {
	shareName := fmt.Sprintf("tf-fs-name-%d", acctest.RandIntRange(10, 100))
	shareSize := acctest.RandIntRange(10, 16000)
	shareTargetName := fmt.Sprintf("tf-fs-tg-name-%d", acctest.RandIntRange(10, 100))
	vpcName := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsShareDataSourceConfig(vpcName, shareName, shareSize, shareTargetName),
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "encryption"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "iops"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "profile"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "size"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "mount_targets.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "mount_targets.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "mount_targets.0.id"),
					resource.TestCheckResourceAttr("data.ibm_is_share.is_share", "mount_targets.0.name", shareTargetName),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "mount_targets.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "zone"),
					resource.TestCheckResourceAttr("data.ibm_is_share.is_share", "tags.0", "sr1"),
					resource.TestCheckResourceAttr("data.ibm_is_share.is_share", "tags.1", "sr2"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share.is_share", "accessor_binding_role"),
				),
			},
		},
	})
}

func testAccCheckIbmIsShareDataSourceConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_share" "is_share" {
			zone = "us-south-2"
			size = 200
			name = "%s"
			profile = "%s"
		}

		data "ibm_is_share" "is_share" {
			share = ibm_is_share.is_share.id
		}
	`, name, acc.ShareProfileName)
}
func testAccCheckIbmIsShareDataSourceConfig404(id string) string {
	return fmt.Sprintf(`
		data "ibm_is_share" "is_share" {
			share = "%s"
		}
	`, id)
}

func testAccCheckIbmIsShareDataSourceConfig(vpcName, shareName string, shareSize int, shareTargetName string) string {
	return fmt.Sprintf(`

		resource "ibm_is_vpc" "tfvpc" {
			name = "%s"
		}
		resource "ibm_is_share" "is_share" {
			allowed_transit_encryption_modes = ["user_managed", "none"]
			zone = "us-south-2"
			name = "%s"
			size = %d
			profile = "%s"
			mount_targets {
				name = "%s"
				vpc = ibm_is_vpc.tfvpc.id
			}
			tags = ["sr1", "sr2"]
		}
		data "ibm_is_share" "is_share" {
			share = ibm_is_share.is_share.id
		}
	`, vpcName, shareName, shareSize, acc.ShareProfileName, shareTargetName)
}

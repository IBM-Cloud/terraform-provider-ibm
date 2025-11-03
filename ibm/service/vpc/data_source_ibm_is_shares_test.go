// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmIsSharesDataSourceBasic(t *testing.T) {
	sname := fmt.Sprintf("tf-fs-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsSharesDataSourceConfigBasic(sname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_shares.is_shares", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_shares.is_shares", "shares.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_shares.is_shares", "total_count"),
				),
			},
		},
	})
}

func TestAccIbmIsSharesDataSourceAllArgs(t *testing.T) {
	shareName := fmt.Sprintf("tf-fs-name-%d", acctest.RandIntRange(10, 100))
	shareSize := acctest.RandIntRange(10, 16000)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsSharesDataSourceConfig(shareName, shareSize),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_shares.is_shares", "shares.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_shares.is_shares", "shares.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_shares.is_shares", "shares.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_shares.is_shares", "shares.0.encryption"),
					resource.TestCheckResourceAttrSet("data.ibm_is_shares.is_shares", "shares.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_shares.is_shares", "shares.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_shares.is_shares", "shares.0.iops"),
					resource.TestCheckResourceAttrSet("data.ibm_is_shares.is_shares", "shares.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_shares.is_shares", "shares.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_shares.is_shares", "shares.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_shares.is_shares", "shares.0.size"),
					resource.TestCheckResourceAttrSet("data.ibm_is_shares.is_shares", "shares.0.storage_generation"),
					resource.TestCheckResourceAttrSet("data.ibm_is_shares.is_shares", "shares.0.accessor_binding_role"),
					resource.TestCheckResourceAttrSet("data.ibm_is_shares.is_shares", "total_count"),
				),
			},
		},
	})
}

func testAccCheckIbmIsSharesDataSourceConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_share" "is_share" {
			zone = "us-south-2"
			size = 200
			name = "%s"
			profile = "%s"
		}

		data "ibm_is_shares" "is_shares" {
		}
	`, name, acc.ShareProfileName)
}

func testAccCheckIbmIsSharesDataSourceConfig(shareName string, shareSize int) string {
	return fmt.Sprintf(`
		resource "ibm_is_share" "is_share" {
			zone = "us-south-2"
			name = "%s"
			profile = "%s"
			size = %d

		}

		data "ibm_is_shares" "is_shares" {
		}
	`, shareName, acc.ShareProfileName, shareSize)
}

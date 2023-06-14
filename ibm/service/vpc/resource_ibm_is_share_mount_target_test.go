// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-beta-go-sdk/vpcbetav1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIbmIsShareMountTarget(t *testing.T) {
	var conf vpcbetav1.ShareMountTarget
	vpcname := fmt.Sprintf("tf-vpc-name-%d", acctest.RandIntRange(10, 100))
	targetName := fmt.Sprintf("tf-target-%d", acctest.RandIntRange(10, 100))
	targetNameUpdate := fmt.Sprintf("tf-target-%d", acctest.RandIntRange(10, 100))
	sname := fmt.Sprintf("tf-fs-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmIsShareMountTargetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsShareTargetConfig(vpcname, sname, targetName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsShareMountTargetExists("ibm_is_share_mount_target.is_share_target", conf),
					resource.TestCheckResourceAttr("ibm_is_share_mount_target.is_share_target", "name", targetName),
					resource.TestCheckResourceAttr("ibm_is_share_mount_target.is_share_target", "transit_encryption", "none"),
				),
			},
			{
				Config: testAccCheckIbmIsShareTargetConfig(vpcname, sname, targetNameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_share_mount_target.is_share_target", "name", targetNameUpdate),
				),
			},
		},
	})
}
func TestAccIBMIsShareMountTargetTransitEncryptionBasic(t *testing.T) {
	var conf vpcbetav1.ShareMountTarget
	vpcname := fmt.Sprintf("tf-vpc-name-%d", acctest.RandIntRange(10, 100))
	targetName := fmt.Sprintf("tf-target-%d", acctest.RandIntRange(10, 100))
	sname := fmt.Sprintf("tf-fs-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmIsShareTargetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsShareTargetTransitEncryptionConfigBasic(vpcname, sname, targetName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsShareTargetExists("ibm_is_share_target.is_share_target", conf),
					resource.TestCheckResourceAttr("ibm_is_share_target.is_share_target", "name", targetName),
					resource.TestCheckResourceAttr("ibm_is_share_target.is_share_target", "transit_encryption", "user_managed"),
				),
			},
		},
	})
}
func testAccCheckIbmIsShareTargetConfig(vpcName, sname, targetName string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default = "true"
	}
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
		vpc = ibm_is_vpc.testacc_vpc.id
		name = "%s"
	}
	`, sname, acc.ShareProfileName, vpcName, targetName)
}
func testAccCheckIBMIsShareTargetTransitEncryptionConfigBasic(vpcName, sname, targetName string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default = "true"
	}
	resource "ibm_is_share" "is_share" {
		zone = "us-south-2"
		size = 200
		name = "%s"
		profile = "%s"
	}
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_share_target" "is_share_target" {
		share = ibm_is_share.is_share.id
		vpc = ibm_is_vpc.testacc_vpc.id
		transit_encryption = "user_managed"
		name = "%s"
	}
	`, sname, acc.ShareProfileName, vpcName, targetName)
}
func testAccCheckIbmIsShareMountTargetExists(n string, obj vpcbetav1.ShareMountTarget) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1BetaAPI()
		if err != nil {
			return err
		}

		getShareTargetOptions := &vpcbetav1.GetShareMountTargetOptions{}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getShareTargetOptions.SetShareID(parts[0])
		getShareTargetOptions.SetID(parts[1])

		shareTarget, _, err := vpcClient.GetShareMountTarget(getShareTargetOptions)
		if err != nil {
			return err
		}

		obj = *shareTarget
		return nil
	}
}

func testAccCheckIbmIsShareMountTargetDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1BetaAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_share_mount_target" {
			continue
		}

		getShareTargetOptions := &vpcbetav1.GetShareMountTargetOptions{}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getShareTargetOptions.SetShareID(parts[0])
		getShareTargetOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetShareMountTarget(getShareTargetOptions)

		if err == nil {
			return fmt.Errorf("ShareTarget still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for ShareTarget (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

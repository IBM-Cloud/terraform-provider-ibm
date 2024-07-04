// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmIsShareDeleteAccessorBinding(t *testing.T) {
	var conf vpcv1.Share

	// name := fmt.Sprintf("tf-fs-name-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tf-vpc-name-%d", acctest.RandIntRange(10, 100))
	shareName := fmt.Sprintf("tf-share-%d", acctest.RandIntRange(10, 100))
	shareName1 := fmt.Sprintf("tf-share1-%d", acctest.RandIntRange(10, 100))
	tEMode1 := "user_managed"
	tEMode2 := "none"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmIsShareDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsShareDeleteAccessorBindingConfig(vpcname, subnetName, tEMode1, shareName, shareName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsShareExists("ibm_is_share.is_share", conf),
					resource.TestCheckResourceAttr("ibm_is_share.is_share", "name", shareName),
					resource.TestCheckResourceAttrSet("ibm_is_share.is_share", "id"),
					resource.TestCheckResourceAttr("ibm_is_share.is_share", "allowed_transit_encryption_modes.0", tEMode1),
					resource.TestCheckResourceAttr("ibm_is_share.is_share_accessor", "allowed_transit_encryption_modes.0", tEMode1),
					resource.TestCheckResourceAttr("ibm_is_share.is_share_accessor", "accessor_binding_role", "accessor"),
					resource.TestCheckResourceAttrSet("ibm_is_share.is_share_accessor", "origin_share.0.crn"),
					resource.TestCheckResourceAttrSet("ibm_is_share.is_share_accessor", "origin_share.0.id"),
					resource.TestCheckResourceAttrSet("ibm_is_share.is_share_accessor", "origin_share.0.resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_share.is_share_accessor", "origin_share.0.href"),
				),
			},
			{
				Config: testAccCheckIbmIsShareDeleteAccessorBindingConfig(vpcname, subnetName, tEMode2, shareName, shareName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsShareExists("ibm_is_share.is_share", conf),
					resource.TestCheckResourceAttr("ibm_is_share.is_share", "name", shareName),
					resource.TestCheckResourceAttrSet("ibm_is_share.is_share", "id"),
					resource.TestCheckResourceAttr("ibm_is_share.is_share", "allowed_transit_encryption_modes.0", tEMode2),
					resource.TestCheckResourceAttr("ibm_is_share.is_share_accessor", "allowed_transit_encryption_modes.0", tEMode2),
					resource.TestCheckResourceAttr("ibm_is_share.is_share", "accessor_binding_role", "origin"),
					resource.TestCheckResourceAttr("ibm_is_share.is_share_accessor", "accessor_binding_role", "accessor"),
					resource.TestCheckResourceAttrSet("ibm_is_share.is_share_accessor", "origin_share.0.crn"),
					resource.TestCheckResourceAttrSet("ibm_is_share.is_share_accessor", "origin_share.0.id"),
					resource.TestCheckResourceAttrSet("ibm_is_share.is_share_accessor", "origin_share.0.resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_share.is_share_accessor", "origin_share.0.href"),
					resource.TestCheckResourceAttrSet("ibm_is_share.is_share", "accessor.#"),
				),
			},
		},
	})
}

func testAccCheckIbmIsShareDeleteAccessorBindingConfig(vpcName, sname, tEMode, shareName, shareName1 string) string {
	return testAccCheckIbmIsShareConfigOriginShareConfig(vpcName, sname, tEMode, shareName, shareName1) + fmt.Sprintf(`
	
	data "ibm_is_share_accessor_bindings" "bindings" {
		share = ibm_is_share.is_share.id
	}
	resource "ibm_is_share_delete_accessor_binding" "delete_biinding" {
		share = ibm_is_share.is_share.id
		share_accessor_binding = data.ibm_is_share_accessor_bindings.bindings.accessor_bindings.0.id
	}
	`)
}

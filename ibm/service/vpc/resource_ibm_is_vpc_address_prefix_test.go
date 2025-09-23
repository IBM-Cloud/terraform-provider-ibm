// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISVPCAddressPrefix_basic(t *testing.T) {
	var vpcAddressPrefix string
	name := fmt.Sprintf("tfvpcuat-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpcnameuat-%d", acctest.RandIntRange(10, 100))
	prefixName := fmt.Sprintf("tfaddprename-%d", acctest.RandIntRange(10, 100))
	prefixName1 := fmt.Sprintf("tfaddprenamename-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCAddressPrefixDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCAddressPrefixConfig(name, prefixName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCAddressPrefixExists("ibm_is_vpc_address_prefix.testacc_vpc_address_prefix", vpcAddressPrefix),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_address_prefix.testacc_vpc_address_prefix", "name", prefixName),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_address_prefix.testacc_vpc_address_prefix", "is_default", "true"),
				),
			},
			{
				Config: testAccCheckIBMISVPCAddressPrefixConfig1(name1, prefixName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCAddressPrefixExists("ibm_is_vpc_address_prefix.testacc_vpc_address_prefix1", vpcAddressPrefix),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_address_prefix.testacc_vpc_address_prefix1", "name", prefixName1),
				),
			},
		},
	})
}

func TestAccIBMISVPCAddressPrefix_InvalidCidr(t *testing.T) {
	name2 := fmt.Sprintf("tfvpcuatnamename-%d", acctest.RandIntRange(10, 100))
	prefixName2 := fmt.Sprintf("tfaddprename-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMISVPCAddressPrefixConfig2(name2, prefixName2),
				ExpectError: regexp.MustCompile(fmt.Sprintf("the request is overlapping with reserved address ranges")),
			},
		},
	})
}

func testAccCheckIBMISVPCAddressPrefixDestroy(s *terraform.State) error {
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpc_address_prefix" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		vpcID := parts[0]
		addrPrefixID := parts[1]
		getvpcAddressPrefixOptions := &vpcv1.GetVPCAddressPrefixOptions{
			VPCID: &vpcID,
			ID:    &addrPrefixID,
		}
		_, _, err1 := sess.GetVPCAddressPrefix(getvpcAddressPrefixOptions)
		if err1 == nil {
			return fmt.Errorf("vpc still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISVPCAddressPrefixExists(n, vpcAddressPrefix string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		vpcID := parts[0]
		addrPrefixID := parts[1]
		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getvpcAddressPrefixOptions := &vpcv1.GetVPCAddressPrefixOptions{
			VPCID: &vpcID,
			ID:    &addrPrefixID,
		}
		addrPrefix, _, err := sess.GetVPCAddressPrefix(getvpcAddressPrefixOptions)
		if err != nil {
			return err
		}
		vpcAddressPrefix = *addrPrefix.ID
		return nil
	}
}

func testAccCheckIBMISVPCAddressPrefixConfig(name, prefixName string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
    name = "%s"
	address_prefix_management = "manual"
}
resource "ibm_is_vpc_address_prefix" "testacc_vpc_address_prefix" {
    name = "%s"
    zone = "%s"
    vpc = "${ibm_is_vpc.testacc_vpc.id}"
	cidr = "%s"
	is_default = true
}`, name, prefixName, acc.ISZoneName, acc.ISAddressPrefixCIDR)
}

func testAccCheckIBMISVPCAddressPrefixConfig1(name, prefixName string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc1" {
    name = "%s"
}
resource "ibm_is_vpc_address_prefix" "testacc_vpc_address_prefix1" {
    name = "%s"
    zone = "%s"
    vpc = "${ibm_is_vpc.testacc_vpc1.id}"
	cidr = "%s"
}`, name, prefixName, acc.ISZoneName, acc.ISAddressPrefixCIDR)
}

func testAccCheckIBMISVPCAddressPrefixConfig2(name, prefixName string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc2" {
    name = "%s"
}
resource "ibm_is_vpc_address_prefix" "testacc_vpc_address_prefix2" {
    name = "%s"
    zone = "%s"
    vpc = "${ibm_is_vpc.testacc_vpc2.id}"
	cidr = "127.0.0.0/8"
}`, name, prefixName, acc.ISZoneName)
}

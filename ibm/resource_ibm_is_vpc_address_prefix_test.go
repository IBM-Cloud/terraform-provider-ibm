/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMISVPCAddressPrefix_basic(t *testing.T) {
	var vpcAddressPrefix string
	name1 := fmt.Sprintf("tfvpcuat-%d", acctest.RandIntRange(10, 100))
	prefixName := fmt.Sprintf("tfaddprename-%d", acctest.RandIntRange(10, 100))
	prefixName1 := fmt.Sprintf("tfaddprenamename-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPCAddressPrefixDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCAddressPrefixConfig(name1, prefixName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCAddressPrefixExists("ibm_is_vpc_address_prefix.testacc_vpc_address_prefix", vpcAddressPrefix),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_address_prefix.testacc_vpc_address_prefix", "name", prefixName),
				),
			},
			{
				Config: testAccCheckIBMISVPCAddressPrefixConfig(name1, prefixName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCAddressPrefixExists("ibm_is_vpc_address_prefix.testacc_vpc_address_prefix", vpcAddressPrefix),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_address_prefix.testacc_vpc_address_prefix", "name", prefixName1),
				),
			},
		},
	})
}

func testAccCheckIBMISVPCAddressPrefixDestroy(s *terraform.State) error {
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_vpc_address_prefix" {
				continue
			}

			parts, err := idParts(rs.Primary.ID)
			if err != nil {
				return err
			}

			vpcID := parts[0]
			addrPrefixID := parts[1]
			getvpcAddressPrefixOptions := &vpcclassicv1.GetVPCAddressPrefixOptions{
				VPCID: &vpcID,
				ID:    &addrPrefixID,
			}
			_, _, err1 := sess.GetVPCAddressPrefix(getvpcAddressPrefixOptions)
			if err1 == nil {
				return fmt.Errorf("vpc still exists: %s", rs.Primary.ID)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_vpc_address_prefix" {
				continue
			}

			parts, err := idParts(rs.Primary.ID)
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
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		vpcID := parts[0]
		addrPrefixID := parts[1]
		userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()
		if userDetails.generation == 1 {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
			getvpcAddressPrefixOptions := &vpcclassicv1.GetVPCAddressPrefixOptions{
				VPCID: &vpcID,
				ID:    &addrPrefixID,
			}
			addrPrefix, _, err := sess.GetVPCAddressPrefix(getvpcAddressPrefixOptions)
			if err != nil {
				return err
			}
			vpcAddressPrefix = *addrPrefix.ID
		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getvpcAddressPrefixOptions := &vpcv1.GetVPCAddressPrefixOptions{
				VPCID: &vpcID,
				ID:    &addrPrefixID,
			}
			addrPrefix, _, err := sess.GetVPCAddressPrefix(getvpcAddressPrefixOptions)
			if err != nil {
				return err
			}
			vpcAddressPrefix = *addrPrefix.ID
		}
		return nil
	}
}

func testAccCheckIBMISVPCAddressPrefixConfig(name, prefixName string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
    name = "%s"
}
resource "ibm_is_vpc_address_prefix" "testacc_vpc_address_prefix" {
    name = "%s"
    zone = "%s"
    vpc = "${ibm_is_vpc.testacc_vpc.id}"
	cidr = "%s"
}`, name, prefixName, ISZoneName, ISAddressPrefixCIDR)
}

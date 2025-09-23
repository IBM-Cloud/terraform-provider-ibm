// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
)

func TestAccIBMPublicAddressRangeBasic(t *testing.T) {
	var conf vpcv1.PublicAddressRange
	ipv4AddressCount := "16"
	name := fmt.Sprintf("tf-name-par%d", acctest.RandIntRange(10, 100))
	vpcName := fmt.Sprintf("tf-name-vpc%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPublicAddressRangeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPublicAddressRangeConfigBasic(vpcName, name, ipv4AddressCount),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPublicAddressRangeExists("ibm_is_public_address_range.public_address_range_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_public_address_range.public_address_range_instance", "ipv4_address_count", ipv4AddressCount),
				),
			},
		},
	})
}

func testAccCheckIBMPublicAddressRangeConfigBasic(vpcName, name, ipv4AddressCount string) string {
	return fmt.Sprintf(`
		resource ibm_is_vpc testacc_vpc {
			name = "%s"
		}
		resource "ibm_is_public_address_range" "public_address_range_instance" {			
			name = "%s"
			ipv4_address_count = "%s"
			target {
    			vpc {
      				id = ibm_is_vpc.testacc_vpc.id
    			}
    			zone {
      				name = "%s"
    			}
  			}
		}
	`, vpcName, name, ipv4AddressCount, acc.ISZoneName)
}

func testAccCheckIBMPublicAddressRangeExists(n string, obj vpcv1.PublicAddressRange) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getPublicAddressRangeOptions := &vpcv1.GetPublicAddressRangeOptions{}

		getPublicAddressRangeOptions.SetID(rs.Primary.ID)

		publicAddressRange, _, err := vpcClient.GetPublicAddressRange(getPublicAddressRangeOptions)
		if err != nil {
			return err
		}

		obj = *publicAddressRange
		return nil
	}
}

func testAccCheckIBMPublicAddressRangeDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_public_address_range" {
			continue
		}

		getPublicAddressRangeOptions := &vpcv1.GetPublicAddressRangeOptions{}

		getPublicAddressRangeOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := vpcClient.GetPublicAddressRange(getPublicAddressRangeOptions)

		if err == nil {
			return fmt.Errorf("PublicAddressRange still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for PublicAddressRange (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

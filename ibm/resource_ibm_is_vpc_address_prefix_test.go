package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

func TestAccIBMISVPCAddressPrefix_basic(t *testing.T) {
	var vpcAddressPrefix *models.AddressPoolPrefix
	name1 := fmt.Sprintf("terraformvpcuat_create_step_name_%d", acctest.RandInt())
	prefixName := fmt.Sprintf("terraformvpcuat_create_prefix_name_%d", acctest.RandInt())
	prefixName1 := fmt.Sprintf("terraformvpcuat_create_prefix_name_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPCAddressPrefixDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCAddressPrefixConfig(name1, prefixName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCAddressPrefixExists("ibm_is_vpc_address_prefix.testacc_vpc_address_prefix", &vpcAddressPrefix),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_address_prefix.testacc_vpc_address_prefix", "name", prefixName),
				),
			},
			{
				Config: testAccCheckIBMISVPCAddressPrefixConfig(name1, prefixName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCAddressPrefixExists("ibm_is_vpc_address_prefix.testacc_vpc_address_prefix", &vpcAddressPrefix),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc_address_prefix.testacc_vpc_address_prefix", "name", prefixName1),
				),
			},
		},
	})
}

func testAccCheckIBMISVPCAddressPrefixDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	vpcC := network.NewVPCClient(sess)

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
		_, err = vpcC.GetAddressPrefix(vpcID, addrPrefixID)

		if err == nil {
			return fmt.Errorf("vpc still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISVPCAddressPrefixExists(n string, vpcAddressPrefix **models.AddressPoolPrefix) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).ISSession()
		vpcC := network.NewVPCClient(sess)

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		vpcID := parts[0]
		addrPrefixID := parts[1]
		addrPrefix, err := vpcC.GetAddressPrefix(vpcID, addrPrefixID)

		if err != nil {
			return err
		}

		*vpcAddressPrefix = addrPrefix
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

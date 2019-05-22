package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/sl"
)

func TestAccIBMCDN_Basic(t *testing.T) {
	var cdn datatypes.Network_CdnMarketplace_Configuration_Mapping

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCDNDestroy,
		Steps: []resource.TestStep{
			{
				Config: testingcdn,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCDNExists("ibm_cdn.test_cdn111", &cdn),
					resource.TestCheckResourceAttr(
						"ibm_cdn.test_cdn111", "hostname", hostname),
					resource.TestCheckResourceAttr(
						"ibm_cdn.test_cdn111", "vendor_name", vendor_name),
					resource.TestCheckResourceAttr(
						"ibm_cdn.test_cdn111", "origin_address", origin_address),
					resource.TestCheckResourceAttr(
						"ibm_cdn.test_cdn111", "origin_type", origin_type),
				),
				Destroy: false,
			},
		},
	})
}

func testAccCheckIBMCDNDestroy(s *terraform.State) error {
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(testAccProvider.Meta().(ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cdn" {
			continue
		}

		cdnId := sl.String(rs.Primary.ID)

		// Try to find the domain
		_, err := service.ListDomainMappingByUniqueId(cdnId)

		if err == nil {
			return fmt.Errorf("CDN mapping with id %d still exists", cdnId)
		}
	}

	return nil
}

func testAccCheckIBMCDNExists(n string, cdn *datatypes.Network_CdnMarketplace_Configuration_Mapping) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}
		cdnId := sl.String(rs.Primary.ID)

		service := services.GetNetworkCdnMarketplaceConfigurationMappingService(testAccProvider.Meta().(ClientSession).SoftLayerSession())

		foundId, err := service.ListDomainMappingByUniqueId(cdnId)

		if err != nil {
			return err
		}
		resourceId := *foundId[0].UniqueId
		if resourceId != rs.Primary.ID {
			return errors.New("Record not found")
		}
		return nil
	}
}

var testingcdn = `
  resource "ibm_cdn" "test_cdn111" {
	hostname = "www.test1.com"
	vendor_name = "akamai"
	origin_address = "222.222.222.2"
	origin_type = "HOST_SERVER"
  }
`

var hostname = "www.test1.com"
var vendor_name = "akamai"
var origin_address = "222.222.222.2"
var origin_type = "HOST_SERVER"

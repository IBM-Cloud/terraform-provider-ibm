package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMPrivateDNSZone_Basic(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnszone%s.com", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSZoneDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPrivateDNSZoneBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSZoneExists("ibm_dns_zone.test-pdns-zone-zone", resultprivatedns),
					resource.TestCheckResourceAttr("ibm_dns_zone.test-pdns-zone-zone", "name", name),
				),
			},
		},
	})
}

func TestAccIBMPrivateDNSZoneImport(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnszone%s.com", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	//var resultendpoint apigatewaysdk.V2Endpoint
	//name := fmt.Sprintf("tftest-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSZoneDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPrivateDNSZoneBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSZoneExists("ibm_dns_zone.test-pdns-zone-zone", resultprivatedns),
					resource.TestCheckResourceAttr("ibm_dns_zone.test-pdns-zone-zone", "name", name),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_dns_zone.test-pdns-zone-zone",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"type"},
			},
		},
	})
}

func testAccCheckIBMPrivateDNSZoneBasic(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		name = "VNF VPC Development"
	}
	resource "ibm_resource_instance" "test-pdns-zone-instance" {
		name = "test-pdns-zone-instance"
		resource_group_id = data.ibm_resource_group.rg.id
		location = "global"
		service = "dns-svcs"
		plan = "free-plan"
	}
	resource "ibm_dns_zone" "test-pdns-zone-zone" {
		depends_on = ["ibm_resource_instance.test-pdns-zone-instance"]
		name = "%s"
		instance_id = ibm_resource_instance.test-pdns-zone-instance.guid
		description = "testdescription"
		label = "testlabel"
	}
	  `, name)
}

func testAccCheckIBMPrivateDNSZoneDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dns_zone" {
			continue
		}
		pdnsClient, err := testAccProvider.Meta().(ClientSession).PrivateDnsClientSession()
		if err != nil {
			return err
		}
		parts := rs.Primary.ID
		partslist := strings.Split(parts, "/")

		getZoneOptions := pdnsClient.NewGetDnszoneOptions(partslist[0], partslist[1])
		_, res, err := pdnsClient.GetDnszone(getZoneOptions)
		if err != nil &&
			res.StatusCode != 403 &&
			!strings.Contains(err.Error(), "The service instance was disabled, any access is not allowed.") {

			return fmt.Errorf("testAccCheckIBMPrivateDNSZoneDestroy: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMPrivateDNSZoneExists(n string, result string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		pdnsClient, err := testAccProvider.Meta().(ClientSession).PrivateDnsClientSession()
		if err != nil {
			return err
		}
		parts := rs.Primary.ID
		partslist := strings.Split(parts, "/")

		getZoneOptions := pdnsClient.NewGetDnszoneOptions(partslist[0], partslist[1])
		r, res, err := pdnsClient.GetDnszone(getZoneOptions)

		if err != nil &&
			res.StatusCode != 403 &&
			!strings.Contains(err.Error(), "The service instance was disabled, any access is not allowed.") {
			return fmt.Errorf("testAccCheckIBMPrivateDNSZoneExists: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}

		result = *r.ID
		return nil
	}
}

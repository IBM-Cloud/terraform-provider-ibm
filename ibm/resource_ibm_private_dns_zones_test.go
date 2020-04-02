package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"strings"
	"testing"
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
					testAccCheckIBMPrivateDNSZoneExists("ibm_dns_zone.id", resultprivatedns),
					resource.TestCheckResourceAttr("ibm_dns_zone.name", "name", name),
				),
			},
		},
	})
}

func testAccCheckIBMPrivateDNSZoneBasic(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		name = "VNF VPC Development"
	}
	resource "ibm_resource_instance" "pdns-1" {
		name = "test-terraform-pdns"
		resource_group_id = data.ibm_resource_group.rg.id
		location = "global"
		service = "dns-svcs"
		plan = "free-plan"
	}
	resource "ibm_dns_zone" "pdns-1-zone" {
		depends_on = ["ibm_resource_instance.pdns-1"]
		name = "%s"
		instance_id = ibm_resource_instance.pdns-1.guid
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
		_, _, err = pdnsClient.GetDnszone(getZoneOptions)

		if err != nil && !strings.Contains(err.Error(), "Not Found") {
			return fmt.Errorf("Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
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
		r, _, err := pdnsClient.GetDnszone(getZoneOptions)

		if err != nil && !strings.Contains(err.Error(), "Not Found") {
			return fmt.Errorf("Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}

		result = *r.ID
		return nil
	}
}

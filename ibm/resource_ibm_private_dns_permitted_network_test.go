package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMPrivateDNSPermittedNetwork_Basic(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnspn%s.com", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSPermittedNetworkDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPrivateDNSPermittedNetworkBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSPermittedNetworkExists("ibm_dns_permitted_network.id", resultprivatedns),
				),
			},
		},
	})
}

func testAccCheckIBMPrivateDNSPermittedNetworkBasic(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		name = "VNF VPC Development"
	}
	resource "ibm_is_vpc" "test_vpc" {
		name = "test-private-dns-vpc"
		resource_group = data.ibm_resource_group.rg.id
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
	resource "ibm_dns_permitted_network" "pdns-1-permitted-network" {
		depends_on = ["ibm_dns_zone.pdns-1-zone"]
		instance_id = ibm_resource_instance.pdns-1.guid
		zone_id = element(split("/", ibm_dns_zone.pdns-1-zone.id),1)
		vpc_crn = ibm_is_vpc.test_vpc.resource_crn
	}
	  `, name)
}

func testAccCheckIBMPrivateDNSPermittedNetworkDestroy(s *terraform.State) error {
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

		getPermittedNetworkOptions := pdnsClient.NewGetPermittedNetworkOptions(partslist[0], partslist[1], partslist[2])
		_, _, err = pdnsClient.GetPermittedNetwork(getPermittedNetworkOptions)

		if err != nil && !strings.Contains(err.Error(), "Not Found") {
			return fmt.Errorf("Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMPrivateDNSPermittedNetworkExists(n string, result string) resource.TestCheckFunc {

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

		getPermittedNetworkOptions := pdnsClient.NewGetPermittedNetworkOptions(partslist[0], partslist[1], partslist[2])
		r, _, err := pdnsClient.GetPermittedNetwork(getPermittedNetworkOptions)

		if err != nil && !strings.Contains(err.Error(), "Not Found") {
			return fmt.Errorf("Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}

		result = *r.ID
		return nil
	}
}

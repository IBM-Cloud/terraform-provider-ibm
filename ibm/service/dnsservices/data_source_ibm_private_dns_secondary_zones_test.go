package dnsservices_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPrivateDNSCustomResolverSecondaryZonesDataSource_basic(t *testing.T) {
	szDescription := "test-secondary-zone"
	node := "data.ibm_dns_custom_resolver_secondary_zones.test-sz"
	szzone := fmt.Sprintf("tf-secondaryzone-%d.com", acctest.RandIntRange(100, 200))
	vpcname := fmt.Sprintf("d-sz-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("d-sz-subnet-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmDnsCrSecondaryZonesDataSourceConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, szDescription, szzone),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "secondary_zones.0.description"),
					resource.TestCheckResourceAttrSet(node, "secondary_zones.0.zone"),
					resource.TestCheckResourceAttrSet(node, "secondary_zones.0.enabled"),
				),
			},
		},
	})
}

func testAccCheckIbmDnsCrSecondaryZonesDataSourceConfig(vpcname, subnetname, zone, cidr, szDescription, szzone string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "rg" {
		is_default = "true"
	}
	resource "ibm_is_vpc" "test-pdns-cr-vpc" {
		name			= "%s"
		resource_group	= data.ibm_resource_group.rg.id
	}
	resource "ibm_is_subnet" "test-pdns-cr-subnet1" {
		name			= "%s"
		vpc				= ibm_is_vpc.test-pdns-cr-vpc.id
		zone			= "%s"
		ipv4_cidr_block	= "%s"
		resource_group	= data.ibm_resource_group.rg.id
	}
	resource "ibm_resource_instance" "test-pdns-cr-instance" {
		name				= "test-pdns-cr-instance"
		resource_group_id	= data.ibm_resource_group.rg.id
		location			= "global"
		service				= "dns-svcs"
		plan				= "standard-dns"
	}
	resource "ibm_dns_custom_resolver" "test" {
		name		= "testpdnscustomresolver"
		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
		description = "new test CR - TF"
		high_availability = false
		enabled 	= true
		locations {
			subnet_crn	= ibm_is_subnet.test-pdns-cr-subnet1.crn
			enabled		= true
		}
	}
	resource "ibm_dns_custom_resolver_secondary_zone" "test" {
		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
		resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
		description = "%s"
		zone = "%s"
		enabled = true
		transfer_from = ["10.0.0.8"]
	}
	data "ibm_dns_custom_resolver_secondary_zones" "test-sz" {
		depends_on  = [ibm_dns_custom_resolver_secondary_zone.test]
		instance_id	= ibm_dns_custom_resolver.test.instance_id
		resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
	}
	`, vpcname, subnetname, zone, cidr, szDescription, szzone)
}

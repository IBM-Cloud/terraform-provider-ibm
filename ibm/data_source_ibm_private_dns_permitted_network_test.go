package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPrivateDNSNetworkDataSource_basic(t *testing.T) {
	node := "data.ibm_dns_permitted_networks.test.permitted_networks"
	riname := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(100, 200))
	vpcname := fmt.Sprintf("tfc-vpc-name-%d", acctest.RandIntRange(10, 100))
	zonename := fmt.Sprintf("tf-dnszone-%d.com", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMpDNSPermittedNetworksDataSourceConfig(riname, vpcname, zonename),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "permitted_networks.0.permitted_network_id"),
					resource.TestCheckResourceAttrSet(node, "permitted_networks.0.state"),
					resource.TestCheckResourceAttrSet(node, "permitted_networks.0.type"),
				),
			},
		},
	})
}

func testAccCheckIBMpDNSPermittedNetworksDataSourceConfig(riname, vpcname, zonename string) string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		name = "default"
	}

	resource "ibm_resource_instance" "test-pdns-instance" {
		name = "%s"
		resource_group_id = data.ibm_resource_group.rg.id
		location = "global"
		service = "dns-svcs"
		plan = "standard-dns"
	}

	resource "ibm_is_vpc" "test_pdns_vpc" {
		name = "%s"
	}

	resource "ibm_dns_zone" "test-pdns-zone" {
		name        = "%s"
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		description = "testdescription6"
		label       = "testlabel-updated6"
	}

	resource "ibm_dns_permitted_network" "test-pdns-permitted-network-nw" {
		instance_id = ibm_dns_zone.test-pdns-zone.instance_id
		zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
		vpc_crn     = ibm_is_vpc.test_pdns_vpc.crn
	}

	data "ibm_dns_permitted_networks" "test" {
		depends_on = [ibm_dns_permitted_network.test-pdns-permitted-network-nw]
		instance_id = ibm_dns_zone.test-pdns-zone.instance_id
		zone_id = ibm_dns_zone.test-pdns-zone.zone_id
	}`, riname, vpcname, zonename)
}

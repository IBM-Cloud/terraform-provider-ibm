// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISVPC_basic(t *testing.T) {
	var vpc string
	name1 := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	apm := "manual"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCConfig(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "default_network_acl_name", "dnwacln"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "default_security_group_name", "dsgn"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "default_routing_table_name", "drtn"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMISVPCConfigUpdate(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMISVPCConfig1(name2, apm),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc1", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "name", name2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "tags.#", "2"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc.testacc_vpc1", "cse_source_addresses.#"),
				),
			},
		},
	})
}
func TestAccIBMISVPC_dns_manual(t *testing.T) {
	var vpc string
	name1 := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	enableHubTrue := true
	server1Add := "192.168.3.4"
	server2Add := "192.168.0.4"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDnsManualConfig(name1, server1Add, enableHubTrue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc1", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "dns.0.enable_hub", fmt.Sprintf("%t", enableHubTrue)),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "dns.0.resolver.0.type", "manual"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc.testacc_vpc1", "dns.0.resolver.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc.testacc_vpc1", "dns.0.resolver.0.servers.#"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "dns.0.resolver.0.servers.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMISVPCDnsManualConfigUpdate(name2, server1Add, server2Add, enableHubTrue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc1", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "name", name2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "dns.0.enable_hub", fmt.Sprintf("%t", enableHubTrue)),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "dns.0.resolver.0.type", "manual"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc.testacc_vpc1", "dns.0.resolver.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc.testacc_vpc1", "dns.0.resolver.0.servers.#"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "dns.0.resolver.0.servers.#", "2"),
				),
			},
		},
	})
}
func TestAccIBMISVPC_dns_manual2(t *testing.T) {
	var vpc string
	name1 := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	enableHubTrue := true
	server1Add := "192.168.3.4"
	server2Add := "192.168.0.4"
	server3Add := "192.168.128.4"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDnsManual2Config(name1, server1Add, server2Add, server3Add, enableHubTrue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc1", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "dns.0.enable_hub", fmt.Sprintf("%t", enableHubTrue)),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "dns.0.resolver.0.type", "manual"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc.testacc_vpc1", "dns.0.resolver.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc.testacc_vpc1", "dns.0.resolver.0.servers.#"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "dns.0.resolver.0.servers.#", "3"),
				),
			},
			{
				Config: testAccCheckIBMISVPCDnsManual2Config(name2, server1Add, server2Add, server3Add, enableHubTrue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc1", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "name", name2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "dns.0.enable_hub", fmt.Sprintf("%t", enableHubTrue)),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "dns.0.resolver.0.type", "manual"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc.testacc_vpc1", "dns.0.resolver.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc.testacc_vpc1", "dns.0.resolver.0.servers.#"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "dns.0.resolver.0.servers.#", "3"),
				),
			},
		},
	})
}
func TestAccIBMISVPC_dns_system(t *testing.T) {
	var vpc string
	name1 := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	enableHubTrue := true
	enableHubFalse := false
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDnsSystemConfig(name1, enableHubTrue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "dns.0.enable_hub", fmt.Sprintf("%t", enableHubTrue)),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "dns.0.resolver.0.type", "system"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc.testacc_vpc", "dns.0.resolver.#"),
				),
			},
			{
				Config: testAccCheckIBMISVPCDnsSystemConfig(name2, enableHubFalse),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", name2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "dns.0.enable_hub", fmt.Sprintf("%t", enableHubFalse)),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "dns.0.resolver.0.type", "system"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc.testacc_vpc", "dns.0.resolver.#"),
				),
			},
		},
	})
}
func TestAccIBMISVPC_dns_delegated(t *testing.T) {
	var vpc string
	name1 := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	subnet1 := fmt.Sprintf("terraformsubnet-%d", acctest.RandIntRange(10, 100))
	subnet2 := fmt.Sprintf("terraformsubnet-%d", acctest.RandIntRange(10, 100))
	subnet3 := fmt.Sprintf("terraformsubnet-%d", acctest.RandIntRange(10, 100))
	subnet4 := fmt.Sprintf("terraformsubnet-%d", acctest.RandIntRange(10, 100))
	resourecinstance := fmt.Sprintf("terraformresource-%d", acctest.RandIntRange(10, 100))
	resolver1 := fmt.Sprintf("terraformresolver-%d", acctest.RandIntRange(10, 100))
	resolver2 := fmt.Sprintf("terraformresolver-%d", acctest.RandIntRange(10, 100))
	binding := fmt.Sprintf("terraformbinding-%d", acctest.RandIntRange(10, 100))
	enableHubTrue := true
	enableHubFalse := false
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDnsDelegatedConfig(name1, name2, subnet1, subnet2, subnet3, subnet4, resourecinstance, resolver1, resolver2, binding, enableHubTrue, enableHubFalse),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.hub_true", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_true", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_false_delegated", "name", name2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_true", "dns.0.enable_hub", fmt.Sprintf("%t", enableHubTrue)),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_true", "dns.0.resolver.0.type", "system"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_false_delegated", "dns.0.enable_hub", fmt.Sprintf("%t", enableHubFalse)),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_false_delegated", "dns.0.resolver.0.type", "system"),
				),
			},
			{
				Config: testAccCheckIBMISVPCDnsDelegatedUpdate1Config(name1, name2, subnet1, subnet2, subnet3, subnet4, resourecinstance, resolver1, resolver2, binding, enableHubTrue, enableHubFalse),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.hub_true", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_true", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_false_delegated", "name", name2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_true", "dns.0.resolution_binding_count", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_true", "dns.0.enable_hub", fmt.Sprintf("%t", enableHubTrue)),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_true", "dns.0.resolver.0.type", "system"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_false_delegated", "dns.0.enable_hub", fmt.Sprintf("%t", enableHubFalse)),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_false_delegated", "dns.0.resolver.0.type", "delegated"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_false_delegated", "dns.0.resolution_binding_count", "1"),
				),
			},
			{
				Config: testAccCheckIBMISVPCDnsDelegatedUpdate2Config(name1, name2, subnet1, subnet2, subnet3, subnet4, resourecinstance, resolver1, resolver2, binding, enableHubTrue, enableHubFalse),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.hub_true", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_true", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_false_delegated", "name", name2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_true", "dns.0.enable_hub", fmt.Sprintf("%t", enableHubTrue)),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_true", "dns.0.resolution_binding_count", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_true", "dns.0.resolver.0.type", "system"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_false_delegated", "dns.0.enable_hub", fmt.Sprintf("%t", enableHubFalse)),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_false_delegated", "dns.0.resolver.0.type", "system"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_false_delegated", "dns.0.resolution_binding_count", "1"),
				),
			},
		},
	})
}
func TestAccIBMISVPC_dns_delegated_first(t *testing.T) {
	var vpc string
	name1 := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	subnet1 := fmt.Sprintf("terraformsubnet-%d", acctest.RandIntRange(10, 100))
	subnet2 := fmt.Sprintf("terraformsubnet-%d", acctest.RandIntRange(10, 100))
	resourecinstance := fmt.Sprintf("terraformresource-%d", acctest.RandIntRange(10, 100))
	resolver1 := fmt.Sprintf("terraformresolver-%d", acctest.RandIntRange(10, 100))
	binding := fmt.Sprintf("terraformbinding-%d", acctest.RandIntRange(10, 100))
	enableHubTrue := true
	enableHubFalse := false
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDnsDelegatedFirstConfig(name1, name2, subnet1, subnet2, resourecinstance, resolver1, binding, enableHubTrue, enableHubFalse),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.hub_true", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_true", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_false_delegated", "name", name2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_true", "dns.0.enable_hub", fmt.Sprintf("%t", enableHubTrue)),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_true", "dns.0.resolver.0.type", "system"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_false_delegated", "dns.0.enable_hub", fmt.Sprintf("%t", enableHubFalse)),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_false_delegated", "dns.0.resolver.0.type", "delegated"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_false_delegated", "dns.0.resolution_binding_count", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.hub_false_delegated", "dns.0.resolver.0.dns_binding_name", binding),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc.hub_false_delegated", "dns.0.resolver.0.dns_binding_id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc.hub_false_delegated", "dns.0.resolver.0.vpc_id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpc.hub_false_delegated", "dns.0.resolver.0.vpc_name"),
				),
			},
		},
	})
}

func TestAccIBMISVPC_basic_apm(t *testing.T) {
	var vpc string
	name := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	apm1 := "auto"
	apm2 := "manual"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCConfig2(name, apm1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc1", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "address_prefix_management", apm1),
				),
			},
			{
				Config: testAccCheckIBMISVPCConfig2(name, apm2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc1", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "address_prefix_management", apm2),
				),
			},
		},
	})
}

func TestAccIBMISVPC_securityGroups(t *testing.T) {
	var vpc string
	vpcname := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	sgname := fmt.Sprintf("terraformvpcsg-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCSgConfig(vpcname, sgname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", vpcname),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group.testacc_security_group", "name", sgname),
					resource.TestCheckResourceAttrSet("ibm_is_vpc.testacc_vpc", "security_group.0.group_name"),
					resource.TestCheckResourceAttrSet("ibm_is_vpc.testacc_vpc", "security_group.0.group_id"),
				),
			},
		},
	})
}

func TestAccIBMISVPC_noSGACLRules(t *testing.T) {
	var vpc string
	vpcname := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCNoSgAclRulesConfig(vpcname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", vpcname),
					resource.TestCheckNoResourceAttr("ibm_is_vpc.testacc_vpc", "security_group.0.rules.#"),
					resource.TestCheckNoResourceAttr("ibm_is_vpc.testacc_vpc", "security_group.0.rules.#"),
				),
			},
		},
	})
}

func testAccCheckIBMISVPCDestroy(s *terraform.State) error {
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpc" {
			continue
		}

		getvpcoptions := &vpcv1.GetVPCOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetVPC(getvpcoptions)

		if err == nil {
			return fmt.Errorf("vpc still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISVPCExists(n, vpcID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getvpcoptions := &vpcv1.GetVPCOptions{
			ID: &rs.Primary.ID,
		}
		foundvpc, _, err := sess.GetVPC(getvpcoptions)
		if err != nil {
			return err
		}
		vpcID = *foundvpc.ID
		return nil
	}
}

func testAccCheckIBMISVPCConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
		default_network_acl_name = "dnwacln"
		default_security_group_name = "dsgn"
		default_routing_table_name = "drtn"
		tags = ["Tag1", "tag2"]
	}`, name)

}
func testAccCheckIBMISVPCDnsSystemConfig(name string, enableHub bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
		dns {
			enable_hub = %t
		}
	}`, name, enableHub)

}
func testAccCheckIBMISVPCDnsManualConfig(name, server1Add string, enableHub bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
		dns {
			enable_hub = %t
			resolver {
				manual_servers {
					address = "%s"
				}
			}
		}
	}
	`, name, enableHub, server1Add)

}
func testAccCheckIBMISVPCDnsManualConfigUpdate(name, server1Add, server2Add string, enableHub bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
		dns {
			enable_hub = %t
			resolver {
				manual_servers {
					address = "%s"
				}
				manual_servers {
					address = "%s"
				}
			}
		}
	}
	`, name, enableHub, server1Add, server2Add)

}
func testAccCheckIBMISVPCDnsManual2Config(name, server1Add, server2Add, server3Add string, enableHub bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
		dns {
			enable_hub = %t
			resolver {
				manual_servers     {
					address = "%s"
					zone_affinity= "%s-1"
				}
				manual_servers{
					address =  "%s"
					zone_affinity = "%s-2"
				}
				manual_servers{
					address= "%s"
					zone_affinity ="%s-3"
				}
			}
		}
	}
	`, name, enableHub, server1Add, acc.RegionName, server2Add, acc.RegionName, server3Add, acc.RegionName)

}
func testAccCheckIBMISVPCDnsDelegatedConfig(vpcname, vpcname2, subnetname1, subnetname2, subnetname3, subnetname4, resourceinstance, resolver1, resolver2, bindingname string, enableHub, enablehubfalse bool) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default	   =  true
	}
	
	resource ibm_is_vpc hub_true {
		name = "%s"
		dns {
			enable_hub = %t
		}
	}
	
	resource ibm_is_vpc hub_false_delegated {
		name = "%s"
		dns {
			enable_hub = %t
		}
	}
	
	resource "ibm_is_subnet" "hub_true_sub1" {
		name		   				=  "%s"
		vpc      	   				=  ibm_is_vpc.hub_true.id
		zone		   				=  "%s"
		total_ipv4_address_count 	= 16
	}
	resource "ibm_is_subnet" "hub_true_sub2" {
		name		   				=  "%s"
		vpc      	   				=  ibm_is_vpc.hub_true.id
		zone		   				=  "%s"
		total_ipv4_address_count 	= 16
	}
	resource "ibm_is_subnet" "hub_false_delegated_sub1" {
		name		   				=  "%s"
		vpc      	   				=  ibm_is_vpc.hub_false_delegated.id
		zone		   				=  "%s"
		total_ipv4_address_count 	= 16
	}
	resource "ibm_is_subnet" "hub_false_delegated_sub2" {
		name		   				=  "%s"
		vpc      	   				=  ibm_is_vpc.hub_false_delegated.id
		zone		   				=  "%s"
		total_ipv4_address_count 	= 16
	}
	resource "ibm_resource_instance" "dns-cr-instance" {
		name		   		=  "%s"
		resource_group_id  	=  data.ibm_resource_group.rg.id
		location           	=  "global"
		service		   		=  "dns-svcs"
		plan		   		=  "standard-dns"
	}
	resource "ibm_dns_custom_resolver" "test_hub_true" {
		name		   		=  "%s"
		instance_id 	   	=  ibm_resource_instance.dns-cr-instance.guid
		description	   		=  "new test CR - TF"
		high_availability  	=  true
		enabled 	   		=  true
		locations {
				subnet_crn  = ibm_is_subnet.hub_true_sub1.crn
				enabled	 = true
		}
		locations {
				subnet_crn  = ibm_is_subnet.hub_true_sub2.crn
				enabled	 = true
		}
	}
	resource "ibm_dns_custom_resolver" "test_hub_false_delegated" {
		name		   		=  "%s"
		instance_id 	   	=  ibm_resource_instance.dns-cr-instance.guid
		description	   		=  "new test CR - TF"
		high_availability  	=  true
		enabled 	   		=  true
		locations {
				subnet_crn  = ibm_is_subnet.hub_false_delegated_sub1.crn
				enabled	 = true
		}
		locations {
				subnet_crn  = ibm_is_subnet.hub_false_delegated_sub2.crn
				enabled	 = true
		}
	}
	
	resource ibm_is_vpc_dns_resolution_binding dnstrue {
		name = "%s"
		vpc_id=  ibm_is_vpc.hub_false_delegated.id
		vpc {
			id = ibm_is_vpc.hub_true.id
		}
	}
	
	`, vpcname, enableHub, vpcname2, enablehubfalse, subnetname1, acc.ISZoneName, subnetname2, acc.ISZoneName, subnetname3, acc.ISZoneName, subnetname4, acc.ISZoneName, resourceinstance, resolver1, resolver2, bindingname)

}
func testAccCheckIBMISVPCDnsDelegatedFirstConfig(vpcname, vpcname2, subnetname1, subnetname2, resourceinstance, resolver1, bindingname string, enableHub, enablehubfalse bool) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default	   =  true
	}
	
	resource ibm_is_vpc hub_true {
		name = "%s"
		dns {
			enable_hub = %t
		}
	}
	
	resource ibm_is_vpc hub_false_delegated {
		depends_on = [ ibm_dns_custom_resolver.test_hub_true ]
		name = "%s"
		dns {
			enable_hub = %t
			resolver {
				type = "delegated"
				vpc_id = ibm_is_vpc.hub_true.id
				dns_binding_name = "%s"
			}
		}
	}
	
	resource "ibm_is_subnet" "hub_true_sub1" {
		name		   				=  "%s"
		vpc      	   				=  ibm_is_vpc.hub_true.id
		zone		   				=  "%s"
		total_ipv4_address_count 	= 16
	}
	resource "ibm_is_subnet" "hub_true_sub2" {
		name		   				=  "%s"
		vpc      	   				=  ibm_is_vpc.hub_true.id
		zone		   				=  "%s"
		total_ipv4_address_count 	= 16
	}
	resource "ibm_resource_instance" "dns-cr-instance" {
		name		   		=  "%s"
		resource_group_id  	=  data.ibm_resource_group.rg.id
		location           	=  "global"
		service		   		=  "dns-svcs"
		plan		   		=  "standard-dns"
	}
	resource "ibm_dns_custom_resolver" "test_hub_true" {
		name		   		=  "%s"
		instance_id 	   	=  ibm_resource_instance.dns-cr-instance.guid
		description	   		=  "new test CR - TF"
		high_availability  	=  true
		enabled 	   		=  true
		locations {
				subnet_crn  = ibm_is_subnet.hub_true_sub1.crn
				enabled	 = true
		}
		locations {
				subnet_crn  = ibm_is_subnet.hub_true_sub2.crn
				enabled	 = true
		}
	}
	`, vpcname, enableHub, vpcname2, enablehubfalse, bindingname, subnetname1, acc.ISZoneName, subnetname2, acc.ISZoneName, resourceinstance, resolver1)

}
func testAccCheckIBMISVPCDnsDelegatedUpdate1Config(vpcname, vpcname2, subnetname1, subnetname2, subnetname3, subnetname4, resourceinstance, resolver1, resolver2, bindingname string, enableHub, enablehubfalse bool) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default	   =  true
	}
	
	resource ibm_is_vpc hub_true {
		name = "%s"
		dns {
			enable_hub = %t
		}
	}
	
	resource ibm_is_vpc hub_false_delegated {
		name = "%s"
		dns {
			enable_hub = %t
			resolver {
				type = "delegated"
				vpc_id = ibm_is_vpc.hub_true.id
			}
		}
	}
	
	resource "ibm_is_subnet" "hub_true_sub1" {
		name		   				=  "%s"
		vpc      	   				=  ibm_is_vpc.hub_true.id
		zone		   				=  "%s"
		total_ipv4_address_count 	= 16
	}
	resource "ibm_is_subnet" "hub_true_sub2" {
		name		   				=  "%s"
		vpc      	   				=  ibm_is_vpc.hub_true.id
		zone		   				=  "%s"
		total_ipv4_address_count 	= 16
	}
	resource "ibm_is_subnet" "hub_false_delegated_sub1" {
		name		   				=  "%s"
		vpc      	   				=  ibm_is_vpc.hub_false_delegated.id
		zone		   				=  "%s"
		total_ipv4_address_count 	= 16
	}
	resource "ibm_is_subnet" "hub_false_delegated_sub2" {
		name		   				=  "%s"
		vpc      	   				=  ibm_is_vpc.hub_false_delegated.id
		zone		   				=  "%s"
		total_ipv4_address_count 	= 16
	}
	resource "ibm_resource_instance" "dns-cr-instance" {
		name		   		=  "%s"
		resource_group_id  	=  data.ibm_resource_group.rg.id
		location           	=  "global"
		service		   		=  "dns-svcs"
		plan		   		=  "standard-dns"
	}
	resource "ibm_dns_custom_resolver" "test_hub_true" {
		name		   		=  "%s"
		instance_id 	   	=  ibm_resource_instance.dns-cr-instance.guid
		description	   		=  "new test CR - TF"
		high_availability  	=  true
		enabled 	   		=  true
		locations {
				subnet_crn  = ibm_is_subnet.hub_true_sub1.crn
				enabled	 = true
		}
		locations {
				subnet_crn  = ibm_is_subnet.hub_true_sub2.crn
				enabled	 = true
		}
	}
	resource "ibm_dns_custom_resolver" "test_hub_false_delegated" {
		name		   		=  "%s"
		instance_id 	   	=  ibm_resource_instance.dns-cr-instance.guid
		description	   		=  "new test CR - TF"
		high_availability  	=  true
		enabled 	   		=  true
		locations {
				subnet_crn  = ibm_is_subnet.hub_false_delegated_sub1.crn
				enabled	 = true
		}
		locations {
				subnet_crn  = ibm_is_subnet.hub_false_delegated_sub2.crn
				enabled	 = true
		}
	}
	
	resource ibm_is_vpc_dns_resolution_binding dnstrue {
		name = "%s"
		vpc_id=  ibm_is_vpc.hub_false_delegated.id
		vpc {
			id = ibm_is_vpc.hub_true.id
		}
	}
	
	`, vpcname, enableHub, vpcname2, enablehubfalse, subnetname1, acc.ISZoneName, subnetname2, acc.ISZoneName, subnetname3, acc.ISZoneName, subnetname4, acc.ISZoneName, resourceinstance, resolver1, resolver2, bindingname)

}
func testAccCheckIBMISVPCDnsDelegatedUpdate2Config(vpcname, vpcname2, subnetname1, subnetname2, subnetname3, subnetname4, resourceinstance, resolver1, resolver2, bindingname string, enableHub, enablehubfalse bool) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default	   =  true
	}
	
	resource ibm_is_vpc hub_true {
		name = "%s"
		dns {
			enable_hub = %t
		}
	}
	
	resource ibm_is_vpc hub_false_delegated {
		name = "%s"
		dns {
			enable_hub = %t
			resolver {
				type = "system"
				vpc_id = "null"
			}
		}
	}
	
	resource "ibm_is_subnet" "hub_true_sub1" {
		name		   				=  "%s"
		vpc      	   				=  ibm_is_vpc.hub_true.id
		zone		   				=  "%s"
		total_ipv4_address_count 	= 16
	}
	resource "ibm_is_subnet" "hub_true_sub2" {
		name		   				=  "%s"
		vpc      	   				=  ibm_is_vpc.hub_true.id
		zone		   				=  "%s"
		total_ipv4_address_count 	= 16
	}
	resource "ibm_is_subnet" "hub_false_delegated_sub1" {
		name		   				=  "%s"
		vpc      	   				=  ibm_is_vpc.hub_false_delegated.id
		zone		   				=  "%s"
		total_ipv4_address_count 	= 16
	}
	resource "ibm_is_subnet" "hub_false_delegated_sub2" {
		name		   				=  "%s"
		vpc      	   				=  ibm_is_vpc.hub_false_delegated.id
		zone		   				=  "%s"
		total_ipv4_address_count 	= 16
	}
	resource "ibm_resource_instance" "dns-cr-instance" {
		name		   		=  "%s"
		resource_group_id  	=  data.ibm_resource_group.rg.id
		location           	=  "global"
		service		   		=  "dns-svcs"
		plan		   		=  "standard-dns"
	}
	resource "ibm_dns_custom_resolver" "test_hub_true" {
		name		   		=  "%s"
		instance_id 	   	=  ibm_resource_instance.dns-cr-instance.guid
		description	   		=  "new test CR - TF"
		high_availability  	=  true
		enabled 	   		=  true
		locations {
				subnet_crn  = ibm_is_subnet.hub_true_sub1.crn
				enabled	 = true
		}
		locations {
				subnet_crn  = ibm_is_subnet.hub_true_sub2.crn
				enabled	 = true
		}
	}
	resource "ibm_dns_custom_resolver" "test_hub_false_delegated" {
		name		   		=  "%s"
		instance_id 	   	=  ibm_resource_instance.dns-cr-instance.guid
		description	   		=  "new test CR - TF"
		high_availability  	=  true
		enabled 	   		=  true
		locations {
				subnet_crn  = ibm_is_subnet.hub_false_delegated_sub1.crn
				enabled	 = true
		}
		locations {
				subnet_crn  = ibm_is_subnet.hub_false_delegated_sub2.crn
				enabled	 = true
		}
	}
	
	resource ibm_is_vpc_dns_resolution_binding dnstrue {
		name = "%s"
		vpc_id=  ibm_is_vpc.hub_false_delegated.id
		vpc {
			id = ibm_is_vpc.hub_true.id
		}
	}
	
	`, vpcname, enableHub, vpcname2, enablehubfalse, subnetname1, acc.ISZoneName, subnetname2, acc.ISZoneName, subnetname3, acc.ISZoneName, subnetname4, acc.ISZoneName, resourceinstance, resolver1, resolver2, bindingname)

}

func testAccCheckIBMISVPCConfigUpdate(name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
		tags = ["tag1"]
	}`, name)

}

func testAccCheckIBMISVPCConfig1(name string, apm string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
		address_prefix_management = "%s"
		tags = ["Tag1", "tag2"]
	}`, name, apm)

}
func testAccCheckIBMISVPCConfig2(name string, apm string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
		address_prefix_management = "%s"
	}`, name, apm)
}

func testAccCheckIBMISVPCSgConfig(vpcname string, sgname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_security_group" "testacc_security_group" {
		name = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
	  }
	  
	  resource "ibm_is_security_group_rule" "testacc_security_group_rule_udp" {
		group      = ibm_is_security_group.testacc_security_group.id
		direction  = "inbound"
		remote     = "127.0.0.1"
		udp {
			port_min = 805
			port_max = 807
		}
	}  
`, vpcname, sgname)

}

func testAccCheckIBMISVPCNoSgAclRulesConfig(vpcname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
		no_sg_acl_rules = true
	  }
`, vpcname)

}

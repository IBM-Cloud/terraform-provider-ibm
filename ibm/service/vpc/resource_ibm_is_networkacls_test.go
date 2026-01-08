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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestNetworkACLGen1(t *testing.T) {
	var nwACL string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISNetworkACLConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.isExampleACL", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "name", "is-example-acl"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "rules.#", "2"),
				),
			},
		},
	})
}

func TestNetworkACLResourceGroupUpdate(t *testing.T) {
	var nwACL string
	setResourceGroup := false
	setResourceGroup1 := true
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISNetworkACLResourceGroupConfig(setResourceGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.isExampleACL", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", "tf-nwacl-vpc"),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "name", "tf-nwacl-subnet"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "rules.#", "5"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "tags.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "resource_group_name", "Default"),
				),
			},
			{
				Config: testAccCheckIBMISNetworkACLResourceGroupConfig(setResourceGroup1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.isExampleACL", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", "tf-nwacl-vpc"),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "name", "tf-nwacl-subnet"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "rules.#", "5"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "tags.#", "2"),
					resource.TestCheckResourceAttrWith("ibm_is_network_acl.isExampleACL", "resource_group_name", func(v string) error {
						if v == "Default" {
							return fmt.Errorf("Attribute 'resource_group' is still Default")
						}
						return nil
					}),
				),
			},
		},
	})
}
func TestNetworkACLGen2(t *testing.T) {
	var nwACL string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISNetworkACLConfig1(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.isExampleACL", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "name", "is-example-acl"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "rules.#", "6"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "tags.#", "2"),
				),
			},
		},
	})
}

func checkNetworkACLDestroy(s *terraform.State) error {
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_network_acl" {
			continue
		}

		getnwacloptions := &vpcv1.GetNetworkACLOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetNetworkACL(getnwacloptions)
		if err == nil {
			return fmt.Errorf("network acl still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISNetworkACLExists(n, nwACL string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getnwacloptions := &vpcv1.GetNetworkACLOptions{
			ID: &rs.Primary.ID,
		}
		foundNwACL, _, err := sess.GetNetworkACL(getnwacloptions)
		if err != nil {
			return err
		}
		nwACL = *foundNwACL.ID
		return nil
	}
}

func testAccCheckIBMISNetworkACLConfig() string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "tf-nwacl-vpc"
	  }

	resource "ibm_is_network_acl" "isExampleACL" {
		name = "is-example-acl"
		vpc  = ibm_is_vpc.testacc_vpc.id
		rules {
		  name        = "outbound"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "outbound"
		  icmp {
			code = 8
			type = 1
		  }
		  # Optionals :
		  # port_max =
		  # port_min =
		}
		rules {
		  name        = "inbound"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "inbound"
		  icmp {
			code = 8
			type = 1
		  }
		  # Optionals :
		  # port_max =
		  # port_min =
		}
	  }
	`)
}

func testAccCheckIBMISNetworkACLConfig1() string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "tf-nwacl-vpc"
	  }

	resource "ibm_is_network_acl" "isExampleACL" {
		name = "is-example-acl"
		tags = ["Tag1", "tag2"]
		vpc  = ibm_is_vpc.testacc_vpc.id
		rules {
		  name        = "outbound"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "outbound"
		  icmp {
			code = 8
			type = 1
		  }
		  # Optionals :
		  # port_max =
		  # port_min =
		}
		rules {
		  name        = "inbound"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "inbound"
		  icmp {
			code = 8
			type = 1
		  }
		  # Optionals :
		  # port_max =
		  # port_min =
		}
		rules {
		  name        = "icmnew"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "inbound"
		  protocol    = "icmp"
		  code 		  = 8
		  type 		  = 1
		}
		rules {
		  name        = "anyprotocol"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "inbound"
		  protocol    = "any"
		} 
		
		rules {
		  name        = "icmptcpudp"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "inbound"
		  protocol    = "icmp_tcp_udp"
		}
		rules {
		  name        = "individual"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "inbound"
		  protocol    = "number_99"
		}
	  }
	`)
}
func testAccCheckIBMISNetworkACLResourceGroupConfig(resourceGroupSelect bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "tf-nwacl-vpc"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name           				 	= "tf-nwacl-subnet"
		vpc             				= ibm_is_vpc.testacc_vpc.id
		zone            				= "%s"
		total_ipv4_address_count 		= 16
		network_acl     				= ibm_is_network_acl.isExampleACL.id
	}

	resource "ibm_is_network_acl" "isExampleACL" {
		tags = ["Tag1", "tag2"]
		vpc  = ibm_is_vpc.testacc_vpc.id
		resource_group  				= %t ? "%s" : null
		rules {
		  name        = "outbound"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "outbound"
		  icmp {
			code = 8
			type = 1
		  }
		}
		rules {
		  name        = "inbound"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "inbound"
		  icmp {
			code = 8
			type = 1
		  }
		}

		lifecycle {
			create_before_destroy = true
		}
	  }
	`, acc.ISZoneName, resourceGroupSelect, acc.IsResourceGroupID)
}

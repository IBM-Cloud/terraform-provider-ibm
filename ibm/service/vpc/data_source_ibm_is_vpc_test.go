// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISVPCDatasource_basic(t *testing.T) {
	var vpc string
	name := fmt.Sprintf("tfc-vpc-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISVPCConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc", vpc),
					resource.TestCheckResourceAttr(
						"data.ibm_is_vpc.ds_vpc", "name", name),
					resource.TestCheckResourceAttr(
						"data.ibm_is_vpc.ds_vpc", "tags.#", "1"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc.ds_vpc", "cse_source_addresses.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc.ds_vpc", "default_network_acl_name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc.ds_vpc", "default_security_group_name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc.ds_vpc", "default_routing_table_name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc.ds_vpc_by_id", "cse_source_addresses.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc.ds_vpc_by_id", "default_network_acl_name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc.ds_vpc_by_id", "default_security_group_name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc.ds_vpc_by_id", "default_routing_table_name"),
				),
			},
		},
	})
}

func TestAccIBMISVPCDatasource_securityGroup(t *testing.T) {
	var vpc string
	vpcname := fmt.Sprintf("tfc-vpc-name-%d", acctest.RandIntRange(10, 100))
	sgname := fmt.Sprintf("tfc-sg-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISVPCSgConfig(vpcname, sgname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc", vpc),
					resource.TestCheckResourceAttr(
						"data.ibm_is_vpc.testacc_vpc", "name", vpcname),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_vpc.testacc_vpc", "security_group.#"),
				),
			},
		},
	})
}

func testDSCheckIBMISVPCConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
			tags = ["tag1"]
		}
		data "ibm_is_vpc" "ds_vpc" {
		    name = "${ibm_is_vpc.testacc_vpc.name}"
		}
		data "ibm_is_vpc" "ds_vpc_by_id" {
		    identifier = "${ibm_is_vpc.testacc_vpc.id}"
		}`, name)
}

func testDSCheckIBMISVPCSgConfig(vpcname, sgname string) string {
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

	  data "ibm_is_vpc" "testacc_vpc" {
		name = ibm_is_vpc.testacc_vpc.name
	  
	}`, vpcname, sgname)
}

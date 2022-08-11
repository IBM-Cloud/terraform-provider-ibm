// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISVirtualEndpointGateway_Basic(t *testing.T) {
	var endpointGateway string
	vpcname1 := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname1 := fmt.Sprintf("tfvpngw-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	name := "ibm_is_virtual_endpoint_gateway.endpoint_gateway"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckisVirtualEndpointGatewayConfigBasic(vpcname1, subnetname1, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists(name, &endpointGateway),
					resource.TestCheckResourceAttr(name, "name", name1),
				),
			},
		},
	})
}

func TestAccIBMISVirtualEndpointGateway_CharacterCount(t *testing.T) {
	var endpointGateway string
	vpcname1 := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname1 := fmt.Sprintf("tfvpngw-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpngw-createname-%d-%s", acctest.RandIntRange(10, 100), acctest.RandString(38))
	name := "ibm_is_virtual_endpoint_gateway.endpoint_gateway"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckisVirtualEndpointGatewayConfigBasic(vpcname1, subnetname1, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists(name, &endpointGateway),
					resource.TestCheckResourceAttr(name, "name", name1),
				),
			},
		},
	})
}

func TestAccIBMISVirtualEndpointGateway_Basic_SecurityGroups(t *testing.T) {
	var endpointGateway string
	vpcname1 := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname1 := fmt.Sprintf("tfvpngw-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	sgname1 := fmt.Sprintf("tfsg-createname-%d", acctest.RandIntRange(10, 100))
	name := "ibm_is_virtual_endpoint_gateway.endpoint_gateway"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckisVirtualEndpointGatewayConfigBasicSecurityGroups(vpcname1, subnetname1, sgname1, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists(name, &endpointGateway),
					resource.TestCheckResourceAttr(name, "name", name1),
				),
			},
		},
	})
}

func TestAccIBMISVirtualEndpointGateway_Import(t *testing.T) {
	vpcname1 := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname1 := fmt.Sprintf("tfvpngw-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	name := "ibm_is_virtual_endpoint_gateway.endpoint_gateway"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckisVirtualEndpointGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckisVirtualEndpointGatewayConfigBasic(vpcname1, subnetname1, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", name1),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMISVirtualEndpointGateway_FullySpecified(t *testing.T) {
	var monitor string
	vpcname1 := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname1 := fmt.Sprintf("tfvpngw-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	name := "ibm_is_virtual_endpoint_gateway.endpoint_gateway"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckisVirtualEndpointGatewayDestroy,
		Steps: []resource.TestStep{
			{
				ExpectNonEmptyPlan: true,
				Config:             testAccCheckisVirtualEndpointGatewayConfigFullySpecified(vpcname1, subnetname1, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists(name, &monitor),
					resource.TestCheckResourceAttr(name, "name", name1),
				),
			},
		},
	})
}

func TestAccIBMISVirtualEndpointGateway_CreateAfterManualDestroy(t *testing.T) {
	t.Skip()
	var monitorOne, monitorTwo string
	vpcname1 := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname1 := fmt.Sprintf("tfvpngw-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	name := "ibm_is_virtual_endpoint_gateway.endpoint_gateway"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckisVirtualEndpointGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckisVirtualEndpointGatewayConfigBasic(vpcname1, subnetname1, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists(name, &monitorOne),
					testAccisVirtualEndpointGatewayManuallyDelete(&monitorOne),
				),
			},
			{
				Config: testAccCheckisVirtualEndpointGatewayConfigBasic(vpcname1, subnetname1, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists(name, &monitorTwo),
					func(state *terraform.State) error {
						if monitorOne == monitorTwo {
							return fmt.Errorf("load balancer monitor id is unchanged even after we thought we deleted it ( %s )",
								monitorTwo)
						}
						return nil
					},
				),
			},
		},
	})
}

func testAccisVirtualEndpointGatewayManuallyDelete(tfEndpointGwID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}
		tfEndpointGw := *tfEndpointGwID
		opt := sess.NewDeleteEndpointGatewayOptions(tfEndpointGw)
		response, err := sess.DeleteEndpointGateway(opt)
		if err != nil {
			return fmt.Errorf("Delete Endpoint Gateway failed: %v", response)
		}
		return nil
	}
}

func testAccCheckisVirtualEndpointGatewayDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_virtual_endpoint_gateway" {
			continue
		}
		opt := sess.NewGetEndpointGatewayOptions(rs.Primary.ID)
		_, response, err := sess.GetEndpointGateway(opt)
		if err == nil {
			return fmt.Errorf("Endpoint Gateway still exists: %v", response)
		}
	}

	return nil
}

func testAccCheckisVirtualEndpointGatewayExists(n string, tfEndpointGwID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No endpoint gateway ID is set")
		}

		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		opt := sess.NewGetEndpointGatewayOptions(rs.Primary.ID)
		result, response, err := sess.GetEndpointGateway(opt)
		if err != nil {
			return fmt.Errorf("Endpoint Gateway does not exist: %s", response)
		}
		*tfEndpointGwID = *result.ID
		return nil
	}
}

func testAccCheckisVirtualEndpointGatewayConfigBasic(vpcname1, subnetname1, name1 string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
    }
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%[1]s"
		resource_group = data.ibm_resource_group.test_acc.id
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%[2]s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%[3]s"
		ipv4_cidr_block = "%[4]s"
		resource_group = data.ibm_resource_group.test_acc.id
	}
	resource "ibm_is_virtual_endpoint_gateway" "endpoint_gateway" {
		name = "%[5]s"
		target {
		  name          = "ibm-dns-server2"
		  resource_type = "provider_infrastructure_service"
		}
		vpc = ibm_is_vpc.testacc_vpc.id
		resource_group = data.ibm_resource_group.test_acc.id
	}`, vpcname1, subnetname1, acc.ISZoneName, acc.ISCIDR, name1)
}

func testAccCheckisVirtualEndpointGatewayConfigFullySpecified(vpcname1, subnetname1, name1 string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
    }
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%[1]s"
		resource_group = data.ibm_resource_group.test_acc.id
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%[2]s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%[3]s"
		ipv4_cidr_block = "%[4]s"
		resource_group = data.ibm_resource_group.test_acc.id
	}
	resource "ibm_is_virtual_endpoint_gateway" "endpoint_gateway" {
		name = "%[5]s"
		target {
		  name          = "ibm-dns-server2"
		  resource_type = "provider_infrastructure_service"
		}
		vpc = ibm_is_vpc.testacc_vpc.id
		ips {
		  subnet   = ibm_is_subnet.testacc_subnet.id
		  name        = "test-reserved-ip1"
		}
		resource_group = data.ibm_resource_group.test_acc.id
	}`, vpcname1, subnetname1, acc.ISZoneName, acc.ISCIDR, name1)
}

func testAccCheckisVirtualEndpointGatewayConfigBasicSecurityGroups(vpcname1, subnetname1, sgname1, name1 string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
    }
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%[1]s"
		resource_group = data.ibm_resource_group.test_acc.id
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%[2]s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%[3]s"
		ipv4_cidr_block = "%[4]s"
		resource_group = data.ibm_resource_group.test_acc.id
	}
	resource "ibm_is_security_group" "testacc_security_group" {
		name = "%[5]s"
		vpc = ibm_is_vpc.testacc_vpc.id
	}
	resource "ibm_is_virtual_endpoint_gateway" "endpoint_gateway" {
		name = "%[6]s"
		target {
		  name          = "ibm-dns-server2"
		  resource_type = "provider_infrastructure_service"
		}
		vpc = ibm_is_vpc.testacc_vpc.id
		resource_group = data.ibm_resource_group.test_acc.id
		security_groups = [ibm_is_security_group.testacc_security_group.id]
	}`, vpcname1, subnetname1, acc.ISZoneName, acc.ISCIDR, sgname1, name1)
}

// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

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

func TestAccIBMISVirtualEndpointGateway_PPSG(t *testing.T) {
	var endpointGateway string
	accessPolicy := "deny"
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tf-test-lb%dd", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tf-test-ppsg%d", acctest.RandIntRange(10, 100))
	name := "ibm_is_virtual_endpoint_gateway.endpoint_gateway"
	// targetName := fmt.Sprintf("tf-egw-target%d", acctest.RandIntRange(10, 100))
	egwName := fmt.Sprintf("tf-egw%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckisVirtualEndpointGatewayConfigPPSG(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, accessPolicy, name1, egwName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists(name, &endpointGateway),
					resource.TestCheckResourceAttr(name, "name", egwName),
					resource.TestCheckResourceAttr(name, "target.0.name", name1),
					resource.TestCheckResourceAttr(name, "target.0.resource_type", "private_path_service_gateway"),
				),
			},
		},
	})
}
func TestAccIBMISVirtualEndpointGateway_PPSG_With_AccessPolicy_Review_And_Timeout(t *testing.T) {
	var endpointGateway string
	accessPolicy := "review"
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tf-test-lb%dd", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tf-test-ppsg%d", acctest.RandIntRange(10, 100))
	name := "ibm_is_virtual_endpoint_gateway.endpoint_gateway"
	// targetName := fmt.Sprintf("tf-egw-target%d", acctest.RandIntRange(10, 100))
	egwName := fmt.Sprintf("tf-egw%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckisVirtualEndpointGatewayConfigPPSGWithTimeout(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, accessPolicy, name1, egwName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists(name, &endpointGateway),
					resource.TestCheckResourceAttr(name, "name", egwName),
					resource.TestCheckResourceAttr(name, "lifecycle_state", "pending"),
					resource.TestCheckResourceAttr(name, "lifecycle_reasons.0.code", "access_pending"),
					resource.TestCheckResourceAttr(name, "target.0.name", name1),
					resource.TestCheckResourceAttr(name, "target.0.resource_type", "private_path_service_gateway"),
				),
			},
		},
	})
}
func TestAccIBMISVirtualEndpointGateway_AllowDnsResolutionBinding(t *testing.T) {
	var endpointGateway string
	vpcname1 := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	enable_hub := false
	name1 := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	name := "ibm_is_virtual_endpoint_gateway.endpoint_gateway"
	allowDnsResolutionBindingTrue := true
	allowDnsResolutionBindingFalse := false
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckisVirtualEndpointGatewayConfigAllowDnsResolutionBinding(vpcname1, name1, enable_hub, allowDnsResolutionBindingTrue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists(name, &endpointGateway),
					resource.TestCheckResourceAttr(name, "name", name1),
					resource.TestCheckResourceAttr(name, "allow_dns_resolution_binding", fmt.Sprintf("%t", allowDnsResolutionBindingTrue)),
				),
			},
			{
				Config: testAccCheckisVirtualEndpointGatewayConfigAllowDnsResolutionBinding(vpcname1, name1, enable_hub, allowDnsResolutionBindingFalse),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists(name, &endpointGateway),
					resource.TestCheckResourceAttr(name, "name", name1),
					resource.TestCheckResourceAttr(name, "allow_dns_resolution_binding", fmt.Sprintf("%t", allowDnsResolutionBindingFalse)),
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
func TestAccIBMISVirtualEndpointGateway_OptionalName(t *testing.T) {
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
				Config:             testAccCheckisVirtualEndpointGatewayConfigOptionalName(vpcname1, subnetname1, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists(name, &monitor),
					resource.TestCheckResourceAttr(name, "name", name1),
					resource.TestCheckResourceAttrSet(name, "ips.#"),
					resource.TestCheckResourceAttrSet(name, "target.#"),
					resource.TestCheckResourceAttrSet(name, "created_at"),
					resource.TestCheckResourceAttrSet(name, "crn"),
					resource.TestCheckResourceAttr(name, "health_state", "ok"),
					resource.TestCheckResourceAttr(name, "lifecycle_state", "stable"),
					resource.TestCheckResourceAttrSet(name, "resource_group"),
					resource.TestCheckResourceAttrSet(name, "resource_type"),
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

func testAccCheckisVirtualEndpointGatewayConfigPPSG(vpcname, subnetname, zone, cidr, lbname, accessPolicy, name, egwName string) string {
	return testAccCheckIBMIsPrivatePathServiceGatewayConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, accessPolicy, name) + fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
    }
	resource "ibm_is_virtual_endpoint_gateway" "endpoint_gateway" {
		name = "%s"
		target {
		  crn          = ibm_is_private_path_service_gateway.is_private_path_service_gateway.crn
		  resource_type = "private_path_service_gateway"
		}
		vpc = ibm_is_vpc.testacc_vpc.id
		resource_group = data.ibm_resource_group.test_acc.id
	}`, egwName)
}
func testAccCheckisVirtualEndpointGatewayConfigPPSGWithTimeout(vpcname, subnetname, zone, cidr, lbname, accessPolicy, name, egwName string) string {
	return testAccCheckIBMIsPrivatePathServiceGatewayConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, accessPolicy, name) + fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
    }
	resource "ibm_is_virtual_endpoint_gateway" "endpoint_gateway" {
		name = "%s"
		target {
		  crn          = ibm_is_private_path_service_gateway.is_private_path_service_gateway.crn
		  resource_type = "private_path_service_gateway"
		}
		vpc = ibm_is_vpc.testacc_vpc.id
		resource_group = data.ibm_resource_group.test_acc.id
		timeouts {
			create = "3m"
		}
	}`, egwName)
}
func testAccCheckisVirtualEndpointGatewayConfigAllowDnsResolutionBinding(vpcname1, name1 string, enable_hub, allowDnsResolutionBinding bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
		dns {
			enable_hub = %t
		}
	}
	resource "ibm_is_virtual_endpoint_gateway" "endpoint_gateway" {
		name 							= "%s"
		target {
			name          = "ibm-ntp-server"
			resource_type = "provider_infrastructure_service"
		}
		vpc 							= ibm_is_vpc.testacc_vpc.id
		allow_dns_resolution_binding 	= %t
	}`, vpcname1, enable_hub, name1, allowDnsResolutionBinding)
}

func testAccCheckisVirtualEndpointGatewayConfigOptionalName(vpcname1, subnetname1, name1 string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
	  }
	  resource "ibm_is_vpc" "testacc_vpc" {
		name           = "%[1]s"
		resource_group = data.ibm_resource_group.test_acc.id
	  }
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%[2]s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%[3]s"
		ipv4_cidr_block = "%[4]s"
		resource_group  = data.ibm_resource_group.test_acc.id
	  }
	  resource "ibm_is_virtual_endpoint_gateway" "endpoint_gateway" {
		name = "%[5]s"
		target {
		  name          = "ibm-dns-server2"
		  resource_type = "provider_infrastructure_service"
		}
		vpc = ibm_is_vpc.testacc_vpc.id
		ips {
		  subnet = ibm_is_subnet.testacc_subnet.id
		}
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

// for service endpoints
func testAccCheckisVirtualEndpointGatewayConfigServiceEndpoints(vpcname1, subnetname1, name1 string) string {
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
		  name          = "ibm-ntp-server"
		  resource_type = "provider_infrastructure_service"
		}
		vpc = ibm_is_vpc.testacc_vpc.id
		resource_group = data.ibm_resource_group.test_acc.id
	}`, vpcname1, subnetname1, acc.ISZoneName, acc.ISCIDR, name1)
}

func TestAccIBMISVirtualEndpointGateway_ServiceEndpoints(t *testing.T) {
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
				Config: testAccCheckisVirtualEndpointGatewayConfigServiceEndpoints(vpcname1, subnetname1, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists(name, &endpointGateway),
					resource.TestCheckResourceAttr(name, "name", name1),
					resource.TestCheckResourceAttrSet(
						name, "service_endpoints.#"),
				),
			},
		},
	})
}

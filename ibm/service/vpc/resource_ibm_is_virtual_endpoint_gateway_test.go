// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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

// dns resolution binding mode tests
func TestAccIBMISVirtualEndpointGateway_DnsResolutionBindingMode(t *testing.T) {
	var endpointGateway string
	vpcName := fmt.Sprintf("tf-vpe-vpc-%d", acctest.RandIntRange(10, 100))
	gatewayName1 := fmt.Sprintf("tf-vpe-gateway-1-%d", acctest.RandIntRange(10, 100))
	gatewayName2 := fmt.Sprintf("tf-vpe-gateway-2-%d", acctest.RandIntRange(10, 100))
	gatewayName3 := fmt.Sprintf("tf-vpe-gateway-3-%d", acctest.RandIntRange(10, 100))
	gatewayName4 := fmt.Sprintf("tf-vpe-gateway-4-%d", acctest.RandIntRange(10, 100))
	gatewayName5 := fmt.Sprintf("tf-vpe-gateway-5-%d", acctest.RandIntRange(10, 100))
	gatewayName6 := fmt.Sprintf("tf-vpe-gateway-6-%d", acctest.RandIntRange(10, 100))
	gatewayName7 := fmt.Sprintf("tf-vpe-gateway-7-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			// Step 1: Create VPC only
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveVPCOnly(vpcName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_vpc.testacc_vpc", "name", vpcName),
				),
			},
			// Step 2: Add VPE Gateway 1 - Deprecated field with allow_dns_resolution_binding = true
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveConfig(vpcName, gatewayName1, "vpe_gateway_1", "deprecated", "true", ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists("ibm_is_virtual_endpoint_gateway.vpe_gateway_1", &endpointGateway),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_1", "name", gatewayName1),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_1", "allow_dns_resolution_binding", "true"),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_1", "dns_resolution_binding_mode", "primary"),
				),
			},
			// Step 3: Update VPE Gateway 1 - Change deprecated field from true to false (in-place update)
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveConfig(vpcName, gatewayName1, "vpe_gateway_1", "deprecated", "false", ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists("ibm_is_virtual_endpoint_gateway.vpe_gateway_1", &endpointGateway),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_1", "name", gatewayName1),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_1", "allow_dns_resolution_binding", "false"),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_1", "dns_resolution_binding_mode", "disabled"),
				),
			},
			// Step 4: Remove VPE Gateway 1 - VPC only
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveVPCOnly(vpcName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_vpc.testacc_vpc", "name", vpcName),
				),
			},
			// Step 5: Add VPE Gateway 2 - Deprecated field with allow_dns_resolution_binding = false
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveConfig(vpcName, gatewayName2, "vpe_gateway_2", "deprecated", "false", ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists("ibm_is_virtual_endpoint_gateway.vpe_gateway_2", &endpointGateway),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_2", "name", gatewayName2),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_2", "allow_dns_resolution_binding", "false"),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_2", "dns_resolution_binding_mode", "disabled"),
				),
			},
			// Step 6: Remove VPE Gateway 2 - VPC only
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveVPCOnly(vpcName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_vpc.testacc_vpc", "name", vpcName),
				),
			},
			// Step 7: Add VPE Gateway 3 - New field with dns_resolution_binding_mode = "primary"
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveConfig(vpcName, gatewayName3, "vpe_gateway_3", "new", "", "primary"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists("ibm_is_virtual_endpoint_gateway.vpe_gateway_3", &endpointGateway),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_3", "name", gatewayName3),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_3", "dns_resolution_binding_mode", "primary"),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_3", "allow_dns_resolution_binding", "true"),
				),
			},
			// Step 8: Update VPE Gateway 3 - Change new field from primary to disabled (in-place update)
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveConfig(vpcName, gatewayName3, "vpe_gateway_3", "new", "", "disabled"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists("ibm_is_virtual_endpoint_gateway.vpe_gateway_3", &endpointGateway),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_3", "name", gatewayName3),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_3", "dns_resolution_binding_mode", "disabled"),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_3", "allow_dns_resolution_binding", "false"),
				),
			},
			// Step 9: Update VPE Gateway 3 - Change to per_resource_binding (in-place update)
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveConfig(vpcName, gatewayName3, "vpe_gateway_3", "new", "", "per_resource_binding"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists("ibm_is_virtual_endpoint_gateway.vpe_gateway_3", &endpointGateway),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_3", "name", gatewayName3),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_3", "dns_resolution_binding_mode", "per_resource_binding"),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_3", "allow_dns_resolution_binding", "true"),
				),
			},
			// Step 10: Remove VPE Gateway 3 - VPC only
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveVPCOnly(vpcName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_vpc.testacc_vpc", "name", vpcName),
				),
			},
			// Step 11: Add VPE Gateway 4 - Deprecated field with allow_dns_resolution_binding = true
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveConfig(vpcName, gatewayName4, "vpe_gateway_4", "deprecated", "true", ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists("ibm_is_virtual_endpoint_gateway.vpe_gateway_4", &endpointGateway),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_4", "name", gatewayName4),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_4", "allow_dns_resolution_binding", "true"),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_4", "dns_resolution_binding_mode", "primary"),
				),
			},
			// Step 12: MIGRATION - Remove deprecated field, add new field with equivalent value (primary)
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveConfig(vpcName, gatewayName4, "vpe_gateway_4", "new", "", "primary"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists("ibm_is_virtual_endpoint_gateway.vpe_gateway_4", &endpointGateway),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_4", "name", gatewayName4),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_4", "dns_resolution_binding_mode", "primary"),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_4", "allow_dns_resolution_binding", "true"),
				),
			},
			// Step 13: Update to per_resource_binding (new capability not available in deprecated field)
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveConfig(vpcName, gatewayName4, "vpe_gateway_4", "new", "", "disabled"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists("ibm_is_virtual_endpoint_gateway.vpe_gateway_4", &endpointGateway),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_4", "name", gatewayName4),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_4", "dns_resolution_binding_mode", "disabled"),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_4", "allow_dns_resolution_binding", "false"),
				),
			},
			// Step 14: Remove VPE Gateway 4 - VPC only
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveVPCOnly(vpcName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_vpc.testacc_vpc", "name", vpcName),
				),
			},
			// Step 15: Add VPE Gateway 5 - Deprecated field with allow_dns_resolution_binding = false
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveConfig(vpcName, gatewayName5, "vpe_gateway_5", "deprecated", "false", ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists("ibm_is_virtual_endpoint_gateway.vpe_gateway_5", &endpointGateway),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_5", "name", gatewayName5),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_5", "allow_dns_resolution_binding", "false"),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_5", "dns_resolution_binding_mode", "disabled"),
				),
			},
			// Step 16: MIGRATION - Remove deprecated field (false), add new field with equivalent value (disabled)
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveConfig(vpcName, gatewayName5, "vpe_gateway_5", "new", "", "disabled"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists("ibm_is_virtual_endpoint_gateway.vpe_gateway_5", &endpointGateway),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_5", "name", gatewayName5),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_5", "dns_resolution_binding_mode", "disabled"),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_5", "allow_dns_resolution_binding", "false"),
				),
			},
			// Step 17: Update to primary (enabling DNS resolution binding)
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveConfig(vpcName, gatewayName5, "vpe_gateway_5", "new", "", "primary"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists("ibm_is_virtual_endpoint_gateway.vpe_gateway_5", &endpointGateway),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_5", "name", gatewayName5),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_5", "dns_resolution_binding_mode", "primary"),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_5", "allow_dns_resolution_binding", "true"),
				),
			},
			// Step 18: Remove VPE Gateway 5 - VPC only
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveVPCOnly(vpcName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_vpc.testacc_vpc", "name", vpcName),
				),
			},
			// Step 19: Add VPE Gateway 6 - No field specified (test default behavior)
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveConfig(vpcName, gatewayName6, "vpe_gateway_6", "none", "", ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists("ibm_is_virtual_endpoint_gateway.vpe_gateway_6", &endpointGateway),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_6", "name", gatewayName6),
					// Should have computed default values from API
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway.vpe_gateway_6", "dns_resolution_binding_mode"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway.vpe_gateway_6", "allow_dns_resolution_binding"),
				),
			},
			// Step 20: Remove VPE Gateway 6 - VPC only
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveVPCOnly(vpcName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_vpc.testacc_vpc", "name", vpcName),
				),
			},
			// Step 21: Add VPE Gateway 7 - Test all three modes sequentially via updates
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveConfig(vpcName, gatewayName7, "vpe_gateway_7", "new", "", "disabled"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists("ibm_is_virtual_endpoint_gateway.vpe_gateway_7", &endpointGateway),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_7", "dns_resolution_binding_mode", "disabled"),
				),
			},
			// Step 22: Update VPE Gateway 7 to primary
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveConfig(vpcName, gatewayName7, "vpe_gateway_7", "new", "", "primary"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists("ibm_is_virtual_endpoint_gateway.vpe_gateway_7", &endpointGateway),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_7", "dns_resolution_binding_mode", "primary"),
				),
			},
			// Step 23: Update VPE Gateway 7 to per_resource_binding
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveConfig(vpcName, gatewayName7, "vpe_gateway_7", "new", "", "disabled"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists("ibm_is_virtual_endpoint_gateway.vpe_gateway_7", &endpointGateway),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_7", "dns_resolution_binding_mode", "disabled"),
				),
			},
			// Step 24: Update VPE Gateway 7 back to disabled (full cycle)
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveConfig(vpcName, gatewayName7, "vpe_gateway_7", "new", "", "disabled"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayExists("ibm_is_virtual_endpoint_gateway.vpe_gateway_7", &endpointGateway),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway.vpe_gateway_7", "dns_resolution_binding_mode", "disabled"),
				),
			},
			// Step 25: Final cleanup - Remove VPE Gateway 7, VPC remains
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayComprehensiveVPCOnly(vpcName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_vpc.testacc_vpc", "name", vpcName),
				),
			},
		},
	})
}

// Helper function for VPC only config
func testAccCheckIBMISVirtualEndpointGatewayComprehensiveVPCOnly(vpcName string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	`, vpcName)
}

// Helper function for VPC + VPE Gateway config
func testAccCheckIBMISVirtualEndpointGatewayComprehensiveConfig(vpcName, gatewayName, resourceName, fieldType, deprecatedValue, newValue string) string {
	baseConfig := fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	`, vpcName)

	var gatewayConfig string

	if fieldType == "deprecated" {
		// Using deprecated field
		gatewayConfig = fmt.Sprintf(`
	resource "ibm_is_virtual_endpoint_gateway" "%s" {
		name = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id

		target {
			crn           = "%s"
			resource_type = "%s"
		}
		
		allow_dns_resolution_binding = %s
	}`, resourceName, gatewayName, acc.IsEndpointGatewayTargetCRN, acc.IsEndpointGatewayTargetType, deprecatedValue)
	} else if fieldType == "new" {
		// Using new field
		gatewayConfig = fmt.Sprintf(`
	resource "ibm_is_virtual_endpoint_gateway" "%s" {
		name = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id

		target {
			crn           = "%s"
			resource_type = "%s"
		}
		
		dns_resolution_binding_mode = "%s"
	}`, resourceName, gatewayName, acc.IsEndpointGatewayTargetCRN, acc.IsEndpointGatewayTargetType, newValue)
	} else if fieldType == "none" {
		// No DNS field specified - test defaults
		gatewayConfig = fmt.Sprintf(`
	resource "ibm_is_virtual_endpoint_gateway" "%s" {
		name = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id

		target {
			crn           = "%s"
			resource_type = "%s"
		}
	}`, resourceName, gatewayName, acc.IsEndpointGatewayTargetCRN, acc.IsEndpointGatewayTargetType)
	}

	return baseConfig + gatewayConfig
}

func TestAccIBMISVirtualEndpointGateway_DNSHubDelegatedModel(t *testing.T) {
	var vpeHubID, vpeSpokeID string

	prefix := fmt.Sprintf("tf-vpe-%d", acctest.RandIntRange(10, 100))

	spokeBindingName := prefix + "-spoke-binding"

	nameHub := "ibm_is_virtual_endpoint_gateway.vpe_hub"
	nameSpoke := "ibm_is_virtual_endpoint_gateway.vpe_spoke"
	nameSpokeBinding := "ibm_is_virtual_endpoint_gateway_resource_binding.vpe_spoke_binding"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{

			// Step 1: Create hub + spoke VPCs, resolver, and VPEs (hub=primary, spoke=disabled)
			{
				Config: testAccVPEHubSpokeDNSConfigCreateSpoke(
					prefix,
					acc.ISZoneName,
					acc.IsResourceBindingCRN,
					acc.IsEndpointGatewayTargetCRN,
					acc.IsEndpointGatewayTargetType,
				),
				Check: resource.ComposeTestCheckFunc(
					// Hub VPE checks
					testAccCheckisVirtualEndpointGatewayExists(nameHub, &vpeHubID),
					resource.TestCheckResourceAttr(nameHub, "dns_resolution_binding_mode", "primary"),

					// Spoke VPE checks (disabled)
					testAccCheckisVirtualEndpointGatewayExists(nameSpoke, &vpeSpokeID),
					resource.TestCheckResourceAttr(nameSpoke, "dns_resolution_binding_mode", "disabled"),
					resource.TestCheckResourceAttr(nameSpoke, "target.0.resource_type", acc.IsEndpointGatewayTargetType),
					resource.TestCheckResourceAttr(nameSpoke, "target.0.crn", acc.IsEndpointGatewayTargetCRN),
				),
			},

			// Step 2: Update spoke VPE → per_resource_binding
			{
				Config: testAccVPEHubSpokeDNSConfigUpdateSpoke(
					prefix,
					acc.ISZoneName,
					acc.IsResourceBindingCRN,
					acc.IsEndpointGatewayTargetCRN,
					acc.IsEndpointGatewayTargetType,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(nameSpoke, "dns_resolution_binding_mode", "per_resource_binding"),
				),
			},

			// Step 3: Create resource binding for spoke VPE
			{
				Config: testAccVPEHubSpokeDNSConfigWithBinding(
					prefix,
					acc.ISZoneName,
					acc.IsResourceBindingCRN,
					acc.IsEndpointGatewayTargetCRN,
					acc.IsEndpointGatewayTargetType,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(nameSpokeBinding, "name", spokeBindingName),
					resource.TestCheckResourceAttr(nameSpokeBinding, "target.0.crn", acc.IsResourceBindingCRN),
				),
			},
		},
	})
}
func testAccVPEHubSpokeDNSConfig(prefix, zone, bindingCRN, targetCRN, targetType string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "default" {
  is_default = true
}

resource "ibm_is_vpc" "hub" {
  name = "%[1]s-vpc-hub"
  dns { enable_hub = true }
}

resource "ibm_is_vpc" "spoke" {
  depends_on = [ibm_dns_custom_resolver.hub_resolver]
  name = "%[1]s-vpc-spoke"
  dns {
    enable_hub = false
    resolver {
      type   = "delegated"
      vpc_id = ibm_is_vpc.hub.id
    }
  }
}

resource "ibm_is_subnet" "hub_subnet_1" {
  name                     = "%[1]s-hub-sub1"
  vpc                      = ibm_is_vpc.hub.id
  zone                     = "%[2]s"
  total_ipv4_address_count = 16
}

resource "ibm_is_subnet" "hub_subnet_2" {
  name                     = "%[1]s-hub-sub2"
  vpc                      = ibm_is_vpc.hub.id
  zone                     = "%[2]s"
  total_ipv4_address_count = 16
}

resource "ibm_resource_instance" "dns_services" {
  name              = "%[1]s-dns-svcs"
  resource_group_id = data.ibm_resource_group.default.id
  location          = "global"
  service           = "dns-svcs"
  plan              = "standard-dns"
}

resource "ibm_dns_custom_resolver" "hub_resolver" {
  name        = "%[1]s-hub-resolver"
  instance_id = ibm_resource_instance.dns_services.guid
  enabled     = true
  high_availability = true

  locations {
    subnet_crn = ibm_is_subnet.hub_subnet_1.crn
    enabled    = true
  }
  locations {
    subnet_crn = ibm_is_subnet.hub_subnet_2.crn
    enabled    = true
  }
}

resource "ibm_is_virtual_endpoint_gateway" "vpe_hub" {
  name = "%[1]s-vpe-hub"
  vpc  = ibm_is_vpc.hub.id
  dns_resolution_binding_mode = "primary"

  target {
    resource_type = "%[4]s"
    crn           = "%[3]s"
  }
}

`, prefix, zone, targetCRN, targetType)
}
func testAccVPEHubSpokeDNSConfigCreateSpoke(prefix, zone, bindingCRN, targetCRN, targetType string) string {
	base := testAccVPEHubSpokeDNSConfig(prefix, zone, bindingCRN, targetCRN, targetType)
	return base + `
resource "ibm_is_virtual_endpoint_gateway" "vpe_spoke" {
  depends_on = [ ibm_is_virtual_endpoint_gateway.vpe_hub ]
  name = "` + prefix + `-vpe-spoke"
  vpc  = ibm_is_vpc.spoke.id
  dns_resolution_binding_mode = "disabled"

  target {
    resource_type = "` + targetType + `"
    crn           = "` + targetCRN + `"
  }
}
`
}
func testAccVPEHubSpokeDNSConfigUpdateSpoke(prefix, zone, bindingCRN, targetCRN, targetType string) string {
	base := testAccVPEHubSpokeDNSConfig(prefix, zone, bindingCRN, targetCRN, targetType)
	return base + `
resource "ibm_is_virtual_endpoint_gateway" "vpe_spoke" {
  depends_on = [ ibm_is_virtual_endpoint_gateway.vpe_hub ]
  name = "` + prefix + `-vpe-spoke"
  vpc  = ibm_is_vpc.spoke.id
  dns_resolution_binding_mode = "per_resource_binding"

  target {
    resource_type = "` + targetType + `"
    crn           = "` + targetCRN + `"
  }
}
`
}

func testAccVPEHubSpokeDNSConfigWithBinding(prefix, zone, bindingCRN, targetCRN, targetType string) string {
	updated := testAccVPEHubSpokeDNSConfigUpdateSpoke(prefix, zone, bindingCRN, targetCRN, targetType)
	return updated + `
resource "ibm_is_virtual_endpoint_gateway_resource_binding" "vpe_spoke_binding" {
  name                = "` + prefix + `-spoke-binding"
  endpoint_gateway_id = ibm_is_virtual_endpoint_gateway.vpe_spoke.id

  target {
    crn = "` + bindingCRN + `"
  }
}
`
}

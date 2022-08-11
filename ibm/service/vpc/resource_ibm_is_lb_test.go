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

func testAccCheckIBMISLBUdpConfig(vpcname, subnetname, zone, cidr, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name 	= "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
		profile = "network-fixed"
		type 	= "public"
}`, vpcname, subnetname, zone, cidr, name)

}
func TestAccIBMISLB_basic_udp(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfupdate%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBUdpConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_LB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name),
					resource.TestCheckResourceAttrSet(
						"ibm_is_lb.testacc_LB", "hostname"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "udp_supported", "true"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_lb.testacc_LB", "udp_supported"),
				),
			},

			{
				Config: testAccCheckIBMISLBUdpConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_LB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "udp_supported", "true"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_lb.testacc_LB", "udp_supported"),
				),
			},
		},
	})
}

func TestAccIBMISLB_basic(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfupdate%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_LB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name),
					resource.TestCheckResourceAttrSet(
						"ibm_is_lb.testacc_LB", "hostname"),
				),
			},

			{
				Config: testAccCheckIBMISLBConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_LB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name1),
				),
			},
		},
	})
}
func TestAccIBMISLB_basic_rip(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfupdate%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_LB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name),
					resource.TestCheckResourceAttrSet(
						"ibm_is_lb.testacc_LB", "hostname"),
				),
			},

			{
				Config: testAccCheckIBMISLBConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_LB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name1),
					resource.TestCheckResourceAttrSet(
						"ibm_is_lb.testacc_LB", "private_ip.0.address"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_lb.testacc_LB", "private_ip.0.href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_lb.testacc_LB", "private_ip.0.name"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_lb.testacc_LB", "private_ip.0.reserved_ip"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_lb.testacc_LB", "private_ip.0.resource_type"),
				),
			},
		},
	})
}

func TestAccIBMISLB_basic_logging(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBLoggingCongig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_LB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "logging", "true"),
				),
			},
		},
	})
}

func TestAccIBMISLB_basic_securityGroups(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	securityGroup := fmt.Sprintf("tflbsecuritygroup%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBSecurityGroupConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, securityGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_LB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "logging", "false"),
				),
			},
		},
	})
}

func TestAccIBMISLB_basic_network(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	nlbName := fmt.Sprintf("tfnlbcreate%d", acctest.RandIntRange(10, 100))
	nlbName1 := fmt.Sprintf("tfnlbupdate%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBNetworkConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, nlbName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_NLB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_NLB", "name", nlbName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_lb.testacc_NLB", "hostname"),
				),
			},

			{
				Config: testAccCheckIBMISLBNetworkConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, nlbName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_NLB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_NLB", "name", nlbName1),
				),
			},
		},
	})
}

func TestAccIBMISLB_basic_network_vnf(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	nlbName := fmt.Sprintf("tfnlbcreate%d", acctest.RandIntRange(10, 100))
	nlbName1 := fmt.Sprintf("tfnlbupdate%d", acctest.RandIntRange(10, 100))
	routeModeTrue := true
	routeModeFalse := false

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBNetworkRouteModeConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, nlbName, routeModeTrue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_NLB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_NLB", "route_mode", fmt.Sprintf("%t", routeModeTrue)),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_NLB", "profile", "network-fixed"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_NLB", "type", "private"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_NLB", "name", nlbName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_lb.testacc_NLB", "hostname"),
				),
			},

			{
				Config: testAccCheckIBMISLBNetworkRouteModeConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, nlbName1, routeModeFalse),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_NLB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_NLB", "route_mode", fmt.Sprintf("%t", routeModeFalse)),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_NLB", "profile", "network-fixed"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_NLB", "type", "private"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_NLB", "name", nlbName1),
				),
			},
		},
	})
}

func TestAccIBMISLB_basic_private(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflbt-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-create-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfupdate%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBConfigPrivate(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_LB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name),
				),
			},

			{
				Config: testAccCheckIBMISLBConfigPrivate(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_LB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name1),
				),
			},
		},
	})
}

func testAccCheckIBMISLBDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_lb" {
			continue
		}

		getlboptions := &vpcv1.GetLoadBalancerOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetLoadBalancer(getlboptions)
		if err == nil {
			return fmt.Errorf("LB still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISLBExists(n, lb string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getlboptions := &vpcv1.GetLoadBalancerOptions{
			ID: &rs.Primary.ID,
		}
		foundLB, _, err := sess.GetLoadBalancer(getlboptions)
		if err != nil {
			return err
		}
		lb = *foundLB.ID

		return nil
	}
}

func testAccCheckIBMISLBConfig(vpcname, subnetname, zone, cidr, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
}`, vpcname, subnetname, zone, cidr, name)

}

func testAccCheckIBMISLBLoggingCongig(vpcname, subnetname, zone, cidr, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
		logging = true
}`, vpcname, subnetname, zone, cidr, name)

}

func testAccCheckIBMISLBNetworkConfig(vpcname, subnetname, zone, cidr, nlbName string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_NLB" {
		name = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
		profile = "network-fixed"
    }`, vpcname, subnetname, zone, cidr, nlbName)

}

func testAccCheckIBMISLBNetworkRouteModeConfig(vpcname, subnetname, zone, cidr, nlbName string, routeMode bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name 						= "%s"
		vpc 						= ibm_is_vpc.testacc_vpc.id
		zone 						= "%s"
		total_ipv4_address_count 	= 16
	}
	resource "ibm_is_lb" "testacc_NLB" {
		name 			= "%s"
		subnets 		= [ibm_is_subnet.testacc_subnet.id]
		profile 		= "network-fixed"
		route_mode 		= %t
		type 			= "private"
    }`, vpcname, subnetname, zone, nlbName, routeMode)

}

func testAccCheckIBMISLBConfigPrivate(vpcname, subnetname, zone, cidr, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
		type = "private"
}`, vpcname, subnetname, zone, cidr, name)

}

func testAccCheckIBMISLBSecurityGroupConfig(vpcname, subnetname, zone, cidr, name, securityGroup string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_security_group" "testacc_security_group" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
		security_groups = [ibm_is_security_group.testacc_security_group.id]
		logging = false
}`, vpcname, subnetname, zone, cidr, securityGroup, name)

}

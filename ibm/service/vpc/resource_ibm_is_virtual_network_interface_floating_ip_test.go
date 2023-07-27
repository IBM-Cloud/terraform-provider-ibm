// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsVirtualNetworkInterfaceFloatingIPBasic(t *testing.T) {
	var conf vpcv1.FloatingIPReference

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPExists("ibm_is_floating_ip.is_floating_ip", conf),
				),
			},
		},
	})
}

func TestAccIBMIsVirtualNetworkInterfaceFloatingIPAllArgs(t *testing.T) {
	var conf vpcv1.FloatingIPReference
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPExists("ibm_is_floating_ip.is_floating_ip", conf),
					resource.TestCheckResourceAttr("ibm_is_floating_ip.is_floating_ip", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPConfig(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_floating_ip.is_floating_ip", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_floating_ip.is_floating_ip",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_is_floating_ip" "is_floating_ip_instance" {
		}
	`)
}

func testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPConfig(name string) string {
	return fmt.Sprintf(`

		resource "ibm_is_floating_ip" "fip" {
			name = "%s"
			zone = "%s"
		}
		resource "ibm_is_virtual_network_interface" "vni" {
			allow_ip_spoofing = %s
			auto_delete = %s
			enable_infrastructure_nat = %s
			ips {
				address = "192.168.3.4"
				href = "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e/reserved_ips/6d353a0f-aeb1-4ae1-832e-1110d10981bb"
				id = "6d353a0f-aeb1-4ae1-832e-1110d10981bb"
				name = "my-reserved-ip"
			}
			name = "%s"
			primary_ip {
				address = "192.168.3.4"
				href = "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e/reserved_ips/6d353a0f-aeb1-4ae1-832e-1110d10981bb"
				id = "6d353a0f-aeb1-4ae1-832e-1110d10981bb"
				name = "my-reserved-ip"
			}
			resource_group {
				id = "fee82deba12e4c0fb69c3b09d1f12345"
			}
			security_groups {
				crn = "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				deleted {
					more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
				}
				href = "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				id = "be5df5ca-12a0-494b-907e-aa6ec2bfa271"
				name = "my-security-group"
			}
			subnet {
				crn = "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
				href = "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
				id = "7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
			}
		}
		resource "ibm_is_virtual_network_interface_floating_ip" "vni_fip" {
			virtual_network_interface 	= ibm_is_virtual_network_interface.vni.id
			flaoting_ip					= ibm_is_floating_ip.fip.id
		}

	`, name, acc.ISZoneName)
}

func testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPExists(n string, obj vpcv1.FloatingIPReference) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		sess, err := vpcClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}

		getVirtualNetworkInterfaceFloatingIPOptions := &vpcv1.GetNetworkInterfaceFloatingIPOptions{}

		getVirtualNetworkInterfaceFloatingIPOptions.SetID(rs.Primary.ID)

		floatingIP, _, err := sess.GetNetworkInterfaceFloatingIP(getVirtualNetworkInterfaceFloatingIPOptions)
		if err != nil {
			return err
		}

		obj = *floatingIP
		return nil
	}
}

func testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPDestroy(s *terraform.State) error {
	sess, err := vpcClient(acc.TestAccProvider.Meta())
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_floating_ip" {
			continue
		}

		getVirtualNetworkInterfaceFloatingIPOptions := &vpcv1.GetNetworkInterfaceFloatingIPOptions{}

		getVirtualNetworkInterfaceFloatingIPOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := sess.GetNetworkInterfaceFloatingIP(getVirtualNetworkInterfaceFloatingIPOptions)

		if err == nil {
			return fmt.Errorf("VirtualNetworkInterfaceFloatingIP still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for VirtualNetworkInterfaceFloatingIP (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

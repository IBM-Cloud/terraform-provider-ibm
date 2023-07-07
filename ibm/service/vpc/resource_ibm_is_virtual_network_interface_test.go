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

func TestAccIBMIsVirtualNetworkInterfaceBasic(t *testing.T) {
	var conf vpcv1.VirtualNetworkInterface

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVirtualNetworkInterfaceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVirtualNetworkInterfaceExists("ibm_is_virtual_network_interface.is_virtual_network_interface", conf),
				),
			},
		},
	})
}

func TestAccIBMIsVirtualNetworkInterfaceAllArgs(t *testing.T) {
	var conf vpcv1.VirtualNetworkInterface
	allowIPSpoofing := "false"
	autoDelete := "false"
	enableInfrastructureNat := "false"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	allowIPSpoofingUpdate := "true"
	autoDeleteUpdate := "true"
	enableInfrastructureNatUpdate := "true"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVirtualNetworkInterfaceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceConfig(allowIPSpoofing, autoDelete, enableInfrastructureNat, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVirtualNetworkInterfaceExists("ibm_is_virtual_network_interface.is_virtual_network_interface", conf),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.is_virtual_network_interface", "allow_ip_spoofing", allowIPSpoofing),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.is_virtual_network_interface", "auto_delete", autoDelete),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.is_virtual_network_interface", "enable_infrastructure_nat", enableInfrastructureNat),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.is_virtual_network_interface", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceConfig(allowIPSpoofingUpdate, autoDeleteUpdate, enableInfrastructureNatUpdate, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.is_virtual_network_interface", "allow_ip_spoofing", allowIPSpoofingUpdate),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.is_virtual_network_interface", "auto_delete", autoDeleteUpdate),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.is_virtual_network_interface", "enable_infrastructure_nat", enableInfrastructureNatUpdate),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.is_virtual_network_interface", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_virtual_network_interface.is_virtual_network_interface",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsVirtualNetworkInterfaceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_is_virtual_network_interface" "is_virtual_network_interface_instance" {
		}
	`)
}

func testAccCheckIBMIsVirtualNetworkInterfaceConfig(allowIPSpoofing string, autoDelete string, enableInfrastructureNat string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_is_virtual_network_interface" "is_virtual_network_interface_instance" {
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
	`, allowIPSpoofing, autoDelete, enableInfrastructureNat, name)
}

func testAccCheckIBMIsVirtualNetworkInterfaceExists(n string, obj vpcv1.VirtualNetworkInterface) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		sess, err := vpcClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}

		getVirtualNetworkInterfaceOptions := &vpcv1.GetVirtualNetworkInterfaceOptions{}

		getVirtualNetworkInterfaceOptions.SetID(rs.Primary.ID)

		virtualNetworkInterface, _, err := sess.GetVirtualNetworkInterface(getVirtualNetworkInterfaceOptions)
		if err != nil {
			return err
		}

		obj = *virtualNetworkInterface
		return nil
	}
}

func testAccCheckIBMIsVirtualNetworkInterfaceDestroy(s *terraform.State) error {
	sess, err := vpcClient(acc.TestAccProvider.Meta())
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_virtual_network_interface" {
			continue
		}

		getVirtualNetworkInterfaceOptions := &vpcv1.GetVirtualNetworkInterfaceOptions{}

		getVirtualNetworkInterfaceOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := sess.GetVirtualNetworkInterface(getVirtualNetworkInterfaceOptions)

		if err == nil {
			return fmt.Errorf("VirtualNetworkInterface still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for VirtualNetworkInterface (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

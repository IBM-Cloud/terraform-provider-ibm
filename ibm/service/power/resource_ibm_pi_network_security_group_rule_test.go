// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/service/power"
)

func TestAccIBMPINetworkSecurityGroupRuleBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkSecurityGroupRuleConfigAddRule(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPINetworkSecurityGroupRuleExists("ibm_pi_network_security_group_rule.network_security_group_rule"),
					resource.TestCheckResourceAttrSet("ibm_pi_network_security_group_rule.network_security_group_rule", power.Arg_NetworkSecurityGroupID),
				),
			},
		},
	})
}

func TestAccIBMPINetworkSecurityGroupRuleTCP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkSecurityGroupRuleConfigAddRuleTCP(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPINetworkSecurityGroupRuleExists("ibm_pi_network_security_group_rule.network_security_group_rule"),
					resource.TestCheckResourceAttrSet("ibm_pi_network_security_group_rule.network_security_group_rule", power.Arg_NetworkSecurityGroupID),
				),
			},
		},
	})
}

func TestAccIBMPINetworkSecurityGroupRulePorts(t *testing.T) {
	destinationPortBegin := "1200"
	sourcePortBegin := "1000"
	destinationPortEnd := "2000"
	sourcePortEnd := "2000"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkSecurityGroupRuleConfigPorts(sourcePortBegin, sourcePortEnd, destinationPortBegin, destinationPortEnd),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPINetworkSecurityGroupRuleExists("ibm_pi_network_security_group_rule.network_security_group_rule"),
					resource.TestCheckResourceAttrSet("ibm_pi_network_security_group_rule.network_security_group_rule", power.Arg_NetworkSecurityGroupID),
				),
			},
		},
	})
}

func TestAccIBMPINetworkSecurityGroupRuleRemove(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkSecurityGroupRuleConfigRemoveRule(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPINetworkSecurityGroupRuleRemoved("ibm_pi_network_security_group_rule.network_security_group_rule", acc.Pi_network_security_group_rule_id),
					resource.TestCheckResourceAttrSet("ibm_pi_network_security_group_rule.network_security_group_rule", power.Arg_NetworkSecurityGroupID),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkSecurityGroupRuleConfigAddRule() string {
	return fmt.Sprintf(`
		resource "ibm_pi_network_security_group_rule" "network_security_group_rule" {
  			pi_cloud_instance_id = "%[1]s"
  			pi_network_security_group_id = "%[2]s"
 			pi_action = "allow"
			pi_protocol {
				type = "all"
			}
			pi_remote {
				id = "%[3]s"
				type = "%[4]s"
			}
		}`, acc.Pi_cloud_instance_id, acc.Pi_network_security_group_id, acc.Pi_remote_id, acc.Pi_remote_type)
}

func testAccCheckIBMPINetworkSecurityGroupRuleConfigAddRuleTCP() string {
	return fmt.Sprintf(`
		resource "ibm_pi_network_security_group_rule" "network_security_group_rule" {
  			pi_cloud_instance_id = "%[1]s"
  			pi_network_security_group_id = "%[2]s"
 			pi_action = "allow"
			pi_destination_ports {
				minimum = 1200
				maximum = 37466
			}
			pi_source_ports {
				minimum = 1000
				maximum = 19500
			}
			pi_protocol {
				tcp_flags {
					flag = "ack"
				}
				tcp_flags {
					flag = "syn"
				}
				tcp_flags {
					flag = "fin"
				}
				type       = "tcp"
			}
			pi_remote {
				id = "%[3]s"
				type = "%[4]s"
			}
		}`, acc.Pi_cloud_instance_id, acc.Pi_network_security_group_id, acc.Pi_remote_id, acc.Pi_remote_type)
}

func testAccCheckIBMPINetworkSecurityGroupRuleConfigPorts(sourcePortBegin string, sourcePortEnd string, destinationPortBegin string, destinationPortEnd string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_network_security_group_rule" "network_security_group_rule" {
  			pi_cloud_instance_id = "%[1]s"
  			pi_network_security_group_id = "%[2]s"
 			pi_action = "allow"
			pi_protocol {
				type = "tcp"
			}
			pi_source_port {
				minimum = %[5]s
				maximum = %[6]s
			}
			pi_destination_port {
				minimum = %[7]s
				maximum = %[8]s
			}
			pi_remote {
				id = "%[3]s"
				type = "%[4]s"
			}
		}`, acc.Pi_cloud_instance_id, acc.Pi_network_security_group_id, acc.Pi_remote_id, acc.Pi_remote_type, sourcePortBegin, sourcePortEnd, destinationPortBegin, destinationPortEnd)
}

func testAccCheckIBMPINetworkSecurityGroupRuleConfigRemoveRule() string {
	return fmt.Sprintf(`
		resource "ibm_pi_network_security_group_rule" "network_security_group_rule" {
			pi_cloud_instance_id = "%[1]s"
			pi_network_security_group_id = "%[2]s"
			pi_network_security_group_rule_id = "%[3]s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_network_security_group_id, acc.Pi_network_security_group_rule_id)
}

func testAccCheckIBMPINetworkSecurityGroupRuleExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		cloudInstanceID, nsgID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		nsgClient := instance.NewIBMIPINetworkSecurityGroupClient(context.Background(), sess, cloudInstanceID)
		_, err = nsgClient.Get(nsgID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIBMPINetworkSecurityGroupRuleRemoved(n string, ruleID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		cloudInstanceID, nsgID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		nsgClient := instance.NewIBMIPINetworkSecurityGroupClient(context.Background(), sess, cloudInstanceID)
		networkSecurityGroup, err := nsgClient.Get(nsgID)
		if err != nil {
			return err
		}
		foundRule := false
		if networkSecurityGroup.Rules != nil {
			for _, rule := range networkSecurityGroup.Rules {
				if *rule.ID == ruleID {
					foundRule = true
					break
				}
			}
		}
		if foundRule {
			return fmt.Errorf("NSG rule still exists")
		}
		return nil
	}
}

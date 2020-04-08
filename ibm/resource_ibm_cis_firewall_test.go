package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMCisFirewall_Basic(t *testing.T) {

	var record string
	name := "lockdowns"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisFirewallDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisFirewallConfigCisDS_Basic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisFirewallExists("ibm_cis_firewall.lockdown", &record),
					resource.TestCheckResourceAttr(
						"ibm_cis_firewall.lockdown", "firewall_type", "lockdowns"),
					resource.TestCheckResourceAttr(
						"ibm_cis_firewall.lockdown", "lockdown.0.configurations.0.value", "127.0.0.1"),
				),
			},
			{
				Config: testAccCheckIBMCisFirewallConfigCisDS_Update(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisFirewallExists("ibm_cis_firewall.lockdown", &record),
					resource.TestCheckResourceAttr(
						"ibm_cis_firewall.lockdown", "firewall_type", "lockdowns"),
					resource.TestCheckResourceAttr(
						"ibm_cis_firewall.lockdown", "lockdown.0.configurations.0.value", "127.0.0.3"),
				),
			},
		},
	})
}
func TestAccIBMCisFirewall_Import(t *testing.T) {

	var record string
	name := "lockdowns"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisFirewallDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCisFirewallConfigCisDS_Basic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisFirewallExists("ibm_cis_firewall.lockdown", &record),
					resource.TestCheckResourceAttr(
						"ibm_cis_firewall.lockdown", "firewall_type", "lockdowns"),
					resource.TestCheckResourceAttr(
						"ibm_cis_firewall.lockdown", "lockdown.0.configurations.0.value", "127.0.0.1"),
				),
			},
			{
				Config: testAccCheckIBMCisFirewallConfigCisDS_Update(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisFirewallExists("ibm_cis_firewall.lockdown", &record),
					resource.TestCheckResourceAttr(
						"ibm_cis_firewall.lockdown", "firewall_type", "lockdowns"),
					resource.TestCheckResourceAttr(
						"ibm_cis_firewall.lockdown", "lockdown.0.configurations.0.value", "127.0.0.3"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cis_firewall.lockdown",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
func testAccCheckIBMCisFirewallDestroy(s *terraform.State) error {
	cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_firewall" {
			continue
		}
		firewallType, recordID, zoneID, cisID, _ := convertTfToCisFourVar(rs.Primary.ID)
		err = cisClient.Firewall().DeleteFirewall(cisID, zoneID, firewallType, recordID)
		if err != nil {
			return fmt.Errorf("Record still exists")
		}
	}

	return nil
}

func testAccCheckIBMCisFirewallExists(n string, tfRecordID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		tfRecord := *tfRecordID
		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		firewallType, recordID, zoneID, cisID, _ := convertTfToCisFourVar(rs.Primary.ID)
		foundRecordPtr, err := cisClient.Firewall().GetFirewall(cisID, zoneID, firewallType, recordID)
		if err != nil {
			return err
		}

		foundRecord := *foundRecordPtr
		if foundRecord.ID != recordID {
			return fmt.Errorf("Record not found")
		}

		tfRecord = convertCisToTfFourVar(firewallType, foundRecord.ID, zoneID, cisID)
		*tfRecordID = tfRecord
		return nil
	}
}

func testAccCheckIBMCisFirewallConfigCisDS_Basic(name string) string {
	return fmt.Sprintf(`
	data "ibm_cis_domain" "cis_domain" {
		cis_id = data.ibm_cis.cis.id
		domain = "cis-terraform.com"
	}
	  
	data "ibm_resource_group" "test_acc" {
		name = "Default"
	}
	  
	data "ibm_cis" "cis" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "test-domain"
	}
	resource "ibm_cis_firewall" "lockdown" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
		firewall_type = "%s"
		lockdown {
		  paused      = "false"
		  description = "testdescription"
		  urls = ["www.cis-terraform.com"]
		  configurations {
			target = "ip"
			value  = "127.0.0.1"
		  }
		  priority=2
		}
	  }
	  `, name)
}

func testAccCheckIBMCisFirewallConfigCisDS_Update(name string) string {
	return fmt.Sprintf(`
	data "ibm_cis_domain" "cis_domain" {
		cis_id = data.ibm_cis.cis.id
		domain = "cis-terraform.com"
	}
	  
	data "ibm_resource_group" "test_acc" {
		name = "Default"
	}
	  
	data "ibm_cis" "cis" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "test-domain"
	}
	resource "ibm_cis_firewall" "lockdown" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
		firewall_type = "%s"
		lockdown {
		  paused      = "false"
		  description = "testdescription"
		  urls = ["www.cis-terraform.com"]
		  configurations {
			target = "ip"
			value  = "127.0.0.3"
		  }
		  priority=2
		}
	  }
	  `, name)
}

// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package contextbasedrestrictions_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/contextbasedrestrictionsv1"
)

func TestAccIBMCbrZoneBasic(t *testing.T) {
	var conf contextbasedrestrictionsv1.Zone

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCbrZoneDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrZoneExists("ibm_cbr_zone.cbr_zone", conf),
				),
			},
		},
	})
}

func TestAccIBMCbrZoneAllArgs(t *testing.T) {
	var conf contextbasedrestrictionsv1.Zone
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	accountID := fmt.Sprintf("12ab34cd56ef78ab90cd12ef34ab56cd")
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	accountIDUpdate := fmt.Sprintf("12ab34cd56ef78ab90cd12ef34ab56cd")
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCbrZoneDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneConfig(name, accountID, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrZoneExists("ibm_cbr_zone.cbr_zone", conf),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "name", name),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "description", description),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneConfig(nameUpdate, accountIDUpdate, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "account_id", accountIDUpdate),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "description", descriptionUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cbr_zone.cbr_zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCbrZoneConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone" {
			name = "Test Zone Resource Config Basic"
			description = "Test Zone Resource Config Basic"
			account_id = "12ab34cd56ef78ab90cd12ef34ab56cd"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
		}
	`)
}

func testAccCheckIBMCbrZoneConfig(name string, accountID string, description string) string {
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone" {
			name = "%s"
			description = "%s"
			account_id = "%s"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
			excluded {
				type = "ipAddress"
				value = "169.23.22.10"
			}
			addresses {
				type = "serviceRef"
				ref {
					service_name = "user-management"
					account_id = "%s"
				}
			}
		}
	`, name, description, accountID, accountID)
}

func testAccCheckIBMCbrZoneExists(n string, obj contextbasedrestrictionsv1.Zone) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		contextBasedRestrictionsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContextBasedRestrictionsV1()
		if err != nil {
			return err
		}

		getZoneOptions := &contextbasedrestrictionsv1.GetZoneOptions{}

		getZoneOptions.SetZoneID(rs.Primary.ID)

		zone, _, err := contextBasedRestrictionsClient.GetZone(getZoneOptions)
		if err != nil {
			return err
		}

		obj = *zone
		return nil
	}
}

func testAccCheckIBMCbrZoneDestroy(s *terraform.State) error {
	contextBasedRestrictionsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContextBasedRestrictionsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cbr_zone" {
			continue
		}

		getZoneOptions := &contextbasedrestrictionsv1.GetZoneOptions{}

		getZoneOptions.SetZoneID(rs.Primary.ID)

		// Try to find the key
		_, response, err := contextBasedRestrictionsClient.GetZone(getZoneOptions)

		if err == nil {
			return fmt.Errorf("cbr_zone still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cbr_zone (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

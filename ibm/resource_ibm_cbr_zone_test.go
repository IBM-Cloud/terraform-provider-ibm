// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/contextbasedrestrictionsv1"
)

func TestAccIBMCbrZoneBasic(t *testing.T) {
	var conf contextbasedrestrictionsv1.Zone

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
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
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCbrZoneDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneConfig(name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrZoneExists("ibm_cbr_zone.cbr_zone", conf),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "name", name),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "description", description),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneConfig(nameUpdate, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "description", descriptionUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cbr_zone.cbr_zone",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"transaction_id"},
			},
		},
	})
}

func testAccCheckIBMCbrZoneConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone" {
			name = "Test Zone Resource Config Basic"
			description = "Test Zone Resource Config Basic"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
		}
	`)
}

func testAccCheckIBMCbrZoneConfig(name string, description string) string {
	return fmt.Sprintf(`

		resource "ibm_cbr_zone" "cbr_zone" {
			name = "%s"
			description = "%s"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
			addresses {
				type = "serviceRef"
				ref {
					service_name = "user-management"
				}
			}
			excluded {
				type = "ipAddress"
				value = "169.23.22.10"
			}
		}
	`, name, description)
}

func testAccCheckIBMCbrZoneExists(n string, obj contextbasedrestrictionsv1.Zone) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		contextBasedRestrictionsClient, err := testAccProvider.Meta().(ClientSession).ContextBasedRestrictionsV1()
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
	contextBasedRestrictionsClient, err := testAccProvider.Meta().(ClientSession).ContextBasedRestrictionsV1()
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

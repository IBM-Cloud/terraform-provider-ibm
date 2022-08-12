// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package contextbasedrestrictions_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCbrZoneDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "zone_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "address_count"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "excluded_count"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "addresses.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "excluded.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "created_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "last_modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "last_modified_by_id"),
				),
			},
		},
	})
}

func TestAccIBMCbrZoneDataSourceAllArgs(t *testing.T) {
	zoneName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	zoneAccountID := "12ab34cd56ef78ab90cd12ef34ab56cd"
	zoneDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneDataSourceConfig(zoneName, zoneAccountID, zoneDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "zone_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "address_count"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "excluded_count"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "addresses.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "addresses.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "addresses.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "excluded.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "excluded.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "excluded.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "created_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "last_modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone", "last_modified_by_id"),
				),
			},
		},
	})
}

func testAccCheckIBMCbrZoneDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone" {
			name = "Test Zone Data Source Config Basic"
			description = "Test Zone Data Source Config Basic"
			account_id = "12ab34cd56ef78ab90cd12ef34ab56cd"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
		}

		data "ibm_cbr_zone" "cbr_zone" {
			zone_id = ibm_cbr_zone.cbr_zone.id
		}
	`)
}

func testAccCheckIBMCbrZoneDataSourceConfig(zoneName string, zoneAccountID string, zoneDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone" {
			name = "%s"
			account_id = "%s"
			description = "%s"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
			excluded {
				type = "ipAddress"
				value = "169.23.22.10"
			}
		}

		data "ibm_cbr_zone" "cbr_zone" {
			zone_id = ibm_cbr_zone.cbr_zone.id
		}
	`, zoneName, zoneAccountID, zoneDescription)
}

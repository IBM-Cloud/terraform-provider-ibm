// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package contextbasedrestrictions_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCbrZoneDataSourceBasic(t *testing.T) {
	accountID, _ := getTestAccountAndZoneID()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCbr(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneDataSourceConfigBasic(accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "zone_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "address_count"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "excluded_count"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "addresses.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "excluded.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "created_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "last_modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "last_modified_by_id"),
				),
			},
		},
	})
}

func TestAccIBMCbrZoneDataSourceAllArgs(t *testing.T) {
	zoneName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	zoneAccountID, _ := getTestAccountAndZoneID()
	zoneDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCbr(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneDataSourceConfig(zoneName, zoneAccountID, zoneDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "zone_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "address_count"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "excluded_count"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "addresses.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "addresses.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "addresses.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "excluded.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "excluded.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "excluded.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "created_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "last_modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone.cbr_zone_instance", "last_modified_by_id"),
				),
			},
		},
	})
}

func testAccCheckIBMCbrZoneDataSourceConfigBasic(accountID string) string {
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone_instance" {
			name = "Test Zone Data Source Config Basic"
			description = "Test Zone Data Source Config Basic"
			account_id = "%s"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
		}

		data "ibm_cbr_zone" "cbr_zone_instance" {
			zone_id = ibm_cbr_zone.cbr_zone_instance.id
		}
	`, accountID)
}

func testAccCheckIBMCbrZoneDataSourceConfig(zoneName string, zoneAccountID string, zoneDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone_instance" {
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

		data "ibm_cbr_zone" "cbr_zone_instance" {
			zone_id = ibm_cbr_zone.cbr_zone_instance.id
		}
	`, zoneName, zoneAccountID, zoneDescription)
}

// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package contextbasedrestrictions_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCbrZoneAddressesDataSourceBasic(t *testing.T) {
	accountID, _ := getTestAccountAndZoneID()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCbr(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneAddressesDataSourceConfigBasic(accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone_addresses.cbr_zone_addresses", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone_addresses.cbr_zone_addresses", "zone_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone_addresses.cbr_zone_addresses", "zone_addresses_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.0.value"),
				),
			},
		},
	})
}

func TestAccIBMCbrZoneAddressesDataSourceMultiple(t *testing.T) {
	accountID, _ := getTestAccountAndZoneID()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCbr(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneAddressesDataSourceConfig(accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone_addresses.cbr_zone_addresses", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone_addresses.cbr_zone_addresses", "zone_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone_addresses.cbr_zone_addresses", "zone_addresses_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.1.type"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.1.value"),
				),
			},
		},
	})
}

func testAccCheckIBMCbrZoneAddressesDataSourceConfigBasic(accountID string) string {
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone" {
			name = "Test Zone Addresses Data Source Config Basic"
			account_id = "%s"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
		}

		resource "ibm_cbr_zone_addresses" "cbr_zone_addresses" {
			zone_id = ibm_cbr_zone.cbr_zone.id
			addresses {
				type = "subnet"
				value = "10.0.0.0/24"
			}
		}

		data "ibm_cbr_zone_addresses" "cbr_zone_addresses" {
			zone_addresses_id = ibm_cbr_zone_addresses.cbr_zone_addresses.id
		}
	`, accountID)
}

func testAccCheckIBMCbrZoneAddressesDataSourceConfig(accountID string) string {
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone" {
			name = "Test Zone Addresses Data Source Config"
			account_id = "%s"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
		}

		resource "ibm_cbr_zone_addresses" "cbr_zone_addresses" {
			zone_id = ibm_cbr_zone.cbr_zone.id
			addresses {
				type = "subnet"
				value = "10.0.0.0/24"
			}
			addresses {
				type = "ipAddress"
				value = "169.24.22.10"
			}
		}

		data "ibm_cbr_zone_addresses" "cbr_zone_addresses" {
			zone_addresses_id = ibm_cbr_zone_addresses.cbr_zone_addresses.id
		}
	`, accountID)
}

// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package contextbasedrestrictions_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/service/contextbasedrestrictions"
)

func TestAccIBMCbrZoneAddressesBasic(t *testing.T) {
	accountID, _ := getTestAccountAndZoneID()

	baseAddresses := []map[string]string{
		{
			"type":  "ipRange",
			"value": "169.23.22.0-169.23.22.255",
		},
	}

	additionalAddresses := []map[string]string{
		{
			"type":  "subnet",
			"value": "10.0.0.0/24",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCbr(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCbrZoneDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneAddressesConfig(accountID, baseAddresses, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrBaseZoneExists("ibm_cbr_zone.cbr_zone", 1),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.#", "1"),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.0.type", baseAddresses[0]["type"]),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.0.value", baseAddresses[0]["value"]),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneAddressesConfig(accountID, baseAddresses, additionalAddresses),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrBaseZoneExists("ibm_cbr_zone.cbr_zone", 1),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.#", "1"),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.0.type", baseAddresses[0]["type"]),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.0.value", baseAddresses[0]["value"]),
					testAccCheckIBMCbrZoneAddressesExists("ibm_cbr_zone_addresses.cbr_zone_addresses", 1),
					resource.TestCheckResourceAttr("ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.#", "1"),
					resource.TestCheckResourceAttr("ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.0.type", additionalAddresses[0]["type"]),
					resource.TestCheckResourceAttr("ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.0.value", additionalAddresses[0]["value"]),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneAddressesConfig(accountID, baseAddresses, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrBaseZoneExists("ibm_cbr_zone.cbr_zone", 1),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.#", "1"),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.0.type", baseAddresses[0]["type"]),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.0.value", baseAddresses[0]["value"]),
					testAccCheckIBMCbrZoneAddressesDestroyed("ibm_cbr_zone_addresses.cbr_zone_addresses"),
				),
			},
		},
	})
}

func TestAccIBMCbrZoneAddressesUpdate(t *testing.T) {
	accountID, _ := getTestAccountAndZoneID()

	baseAddresses := []map[string]string{
		{
			"type":  "ipRange",
			"value": "169.23.22.0-169.23.22.255",
		},
	}

	additionalAddresses := []map[string]string{
		{
			"type":  "subnet",
			"value": "10.0.0.0/24",
		},
	}

	additionalAddressesUpdate := []map[string]string{
		{
			"type":  "subnet",
			"value": "10.0.0.0/25",
		},
		{
			"type":  "ipAddress",
			"value": "169.24.22.10",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCbr(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCbrZoneDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneAddressesConfig(accountID, baseAddresses, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrBaseZoneExists("ibm_cbr_zone.cbr_zone", 1),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.#", "1"),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.0.type", baseAddresses[0]["type"]),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.0.value", baseAddresses[0]["value"]),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneAddressesConfig(accountID, baseAddresses, additionalAddresses),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrBaseZoneExists("ibm_cbr_zone.cbr_zone", 1),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.#", "1"),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.0.type", baseAddresses[0]["type"]),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.0.value", baseAddresses[0]["value"]),
					testAccCheckIBMCbrZoneAddressesExists("ibm_cbr_zone_addresses.cbr_zone_addresses", 1),
					resource.TestCheckResourceAttr("ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.#", "1"),
					resource.TestCheckResourceAttr("ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.0.type", additionalAddresses[0]["type"]),
					resource.TestCheckResourceAttr("ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.0.value", additionalAddresses[0]["value"]),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneAddressesConfig(accountID, nil, additionalAddressesUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrBaseZoneExists("ibm_cbr_zone.cbr_zone", 0),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.#", "0"),
					testAccCheckIBMCbrZoneAddressesExists("ibm_cbr_zone_addresses.cbr_zone_addresses", 2),
					resource.TestCheckResourceAttr("ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.#", "2"),
					resource.TestCheckResourceAttr("ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.0.type", additionalAddressesUpdate[0]["type"]),
					resource.TestCheckResourceAttr("ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.0.value", additionalAddressesUpdate[0]["value"]),
					resource.TestCheckResourceAttr("ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.1.type", additionalAddressesUpdate[1]["type"]),
					resource.TestCheckResourceAttr("ibm_cbr_zone_addresses.cbr_zone_addresses", "addresses.1.value", additionalAddressesUpdate[1]["value"]),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCbrZoneAddressesConfig(accountID, baseAddresses, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrBaseZoneExists("ibm_cbr_zone.cbr_zone", 1),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.#", "1"),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.0.type", baseAddresses[0]["type"]),
					resource.TestCheckResourceAttr("ibm_cbr_zone.cbr_zone", "addresses.0.value", baseAddresses[0]["value"]),
					testAccCheckIBMCbrZoneAddressesDestroyed("ibm_cbr_zone_addresses.cbr_zone_addresses"),
				),
			},
		},
	})
}

func testAccCheckIBMCbrZoneAddressesConfig(accountID string, base, additional []map[string]string) string {
	var result strings.Builder

	addAddresses := func(addresses []map[string]string) {
		for _, address := range addresses {
			result.WriteString(fmt.Sprintf(`
			addresses {
				type = "%s"
				value = "%s"
			}`,
				address["type"], address["value"]))
		}
	}

	result.WriteString(fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone" {
			name = "Test Zone Addresses Resource Config"
			description = "Test Zone Addresses Resource Config"
			account_id = "%s"`,
		accountID))
	addAddresses(base)
	result.WriteString(`
		}
	`)

	if len(additional) > 0 {
		result.WriteString(`
		resource "ibm_cbr_zone_addresses" "cbr_zone_addresses" {
			zone_id = ibm_cbr_zone.cbr_zone.id`)
		addAddresses(additional)
		result.WriteString(`
		}
	`)
	}
	return result.String()
}

func testAccCheckIBMCbrBaseZoneExists(n string, expectedCount int) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		contextBasedRestrictionsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContextBasedRestrictionsV1()
		if err != nil {
			return err
		}

		getZoneOptions := contextBasedRestrictionsClient.NewGetZoneOptions(rs.Primary.ID)

		zone, _, err := contextBasedRestrictionsClient.GetZone(getZoneOptions)
		if err != nil {
			return err
		}

		addresses := contextbasedrestrictions.FilterAddressList(zone.Addresses, func(id string) bool {
			return id == ""
		})

		actualCount := len(addresses)
		if actualCount != expectedCount {
			return fmt.Errorf("%s has an unexpected number of addresses (actual=%d, expected=%d)", n, actualCount, expectedCount)
		}

		return nil
	}
}

func testAccCheckIBMCbrZoneAddressesExists(n string, expectedCount int) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if expectedCount == 0 {
			if !ok {
				return nil
			}
			return fmt.Errorf("%s still exists", n)
		}

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		contextBasedRestrictionsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContextBasedRestrictionsV1()
		if err != nil {
			return err
		}

		parts := strings.Split(rs.Primary.ID, "/")
		getZoneOptions := contextBasedRestrictionsClient.NewGetZoneOptions(parts[0])

		zone, _, err := contextBasedRestrictionsClient.GetZone(getZoneOptions)
		if err != nil {
			return err
		}

		addresses := contextbasedrestrictions.FilterAddressList(zone.Addresses, func(id string) bool {
			return id == parts[1]
		})

		actualCount := len(addresses)
		if actualCount != expectedCount {
			if actualCount == 0 {
				return fmt.Errorf("%s does not exist", n)
			}
			return fmt.Errorf("%s has an unexpected number of addresses (actual=%d, expected=%d)", n, actualCount, expectedCount)
		}

		return nil
	}
}

func testAccCheckIBMCbrZoneAddressesDestroyed(n string) resource.TestCheckFunc {
	return testAccCheckIBMCbrZoneAddressesExists(n, 0)
}

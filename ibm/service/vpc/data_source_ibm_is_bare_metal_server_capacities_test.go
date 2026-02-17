// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISBareMetalServerCapacitiesDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_bare_metal_server_capacities.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBareMetalServerCapacitiesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "id"),
					resource.TestCheckResourceAttrSet(resName, "capacities.#"),
					resource.TestCheckResourceAttrSet(resName, "capacities.0.name"),
					resource.TestCheckResourceAttrSet(resName, "capacities.0.zones.#"),
				),
			},
		},
	})
}

func TestAccIBMISBareMetalServerCapacitiesDataSource_FilterByProfile(t *testing.T) {
	resName := "data.ibm_is_bare_metal_server_capacities.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBareMetalServerCapacitiesDataSourceProfileFilterConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "id"),
					resource.TestCheckResourceAttrSet(resName, "capacities.#"),
					resource.TestCheckResourceAttrSet(resName, "capacities.0.name"),
					resource.TestCheckResourceAttrSet(resName, "capacities.0.zones.#"),
				),
			},
		},
	})
}

func TestAccIBMISBareMetalServerCapacitiesDataSource_FilterByZone(t *testing.T) {
	resName := "data.ibm_is_bare_metal_server_capacities.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBareMetalServerCapacitiesDataSourceZoneFilterConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "id"),
					resource.TestCheckResourceAttrSet(resName, "capacities.#"),
					resource.TestCheckResourceAttrSet(resName, "capacities.0.name"),
					resource.TestCheckResourceAttrSet(resName, "capacities.0.zones.#"),
				),
			},
		},
	})
}

func TestAccIBMISBareMetalServerCapacitiesDataSource_FilterByBoth(t *testing.T) {
	resName := "data.ibm_is_bare_metal_server_capacities.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBareMetalServerCapacitiesDataSourceBothFiltersConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "id"),
					resource.TestCheckResourceAttrSet(resName, "capacities.#"),
					resource.TestCheckResourceAttrSet(resName, "capacities.0.name"),
					resource.TestCheckResourceAttrSet(resName, "capacities.0.zones.#"),
				),
			},
		},
	})
}

func testAccCheckIBMISBareMetalServerCapacitiesDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_is_bare_metal_server_capacities" "test1" {
		}
	`)
}

func testAccCheckIBMISBareMetalServerCapacitiesDataSourceProfileFilterConfig() string {
	return fmt.Sprintf(`
		data "ibm_is_bare_metal_server_profiles" "test_profiles" {
		}

		data "ibm_is_bare_metal_server_capacities" "test1" {
			profile = data.ibm_is_bare_metal_server_profiles.test_profiles.profiles.0.name
		}
	`)
}

func testAccCheckIBMISBareMetalServerCapacitiesDataSourceZoneFilterConfig() string {
	return fmt.Sprintf(`
		data "ibm_is_zones" "test_zones" {
			region = "us-south"
		}

		data "ibm_is_bare_metal_server_capacities" "test1" {
			zone = data.ibm_is_zones.test_zones.zones.0
		}
	`)
}

func testAccCheckIBMISBareMetalServerCapacitiesDataSourceBothFiltersConfig() string {
	return fmt.Sprintf(`
		data "ibm_is_bare_metal_server_profiles" "test_profiles" {
		}

		data "ibm_is_zones" "test_zones" {
			region = "us-south"
		}

		data "ibm_is_bare_metal_server_capacities" "test1" {
			profile = data.ibm_is_bare_metal_server_profiles.test_profiles.profiles.0.name
			zone    = data.ibm_is_zones.test_zones.zones.0
		}
	`)
}

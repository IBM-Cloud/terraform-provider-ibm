// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMComputeDedicatedHost_Basic(t *testing.T) {
	var dedicatedHost datatypes.Virtual_DedicatedHost
	hostname := fmt.Sprintf("terraformuat_%d", acctest.RandIntRange(10, 100))
	updatedHostname := fmt.Sprintf("terraformuat_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMComputerDedicatedHostDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMComputeDedicatedHostConfigBasic(hostname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeDedicatedHostExists("ibm_compute_dedicated_host.dedicatedHourly", &dedicatedHost),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedHourly", "hostname", hostname),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedHourly", "domain", "uat.com"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedHourly", "datacenter", "dal10"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedHourly", "router_hostname", "bcr01a.dal10"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedHourly", "hourly_billing", "true"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedHourly", "cpu_count", "56"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedHourly", "disk_capacity", "1200"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedHourly", "memory_capacity", "242"),
				),
			},

			{
				Config: testAccCheckIBMComputeDedicatedHostConfigUpdated(updatedHostname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeDedicatedHostExists("ibm_compute_dedicated_host.dedicatedHourly", &dedicatedHost),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedHourly", "hostname", updatedHostname),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedHourly", "domain", "uat.com"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedHourly", "datacenter", "dal10"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedHourly", "router_hostname", "bcr01a.dal10"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedHourly", "hourly_billing", "true"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedHourly", "cpu_count", "56"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedHourly", "disk_capacity", "1200"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedHourly", "memory_capacity", "242"),
				),
			},
		},
	})
}

func TestAccIBMComputerDedicatedHostWithTag(t *testing.T) {
	var dedicatedHost datatypes.Virtual_DedicatedHost
	hostname := fmt.Sprintf("terraformuat_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMComputerDedicatedHostDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMComputeDedicatedHostWithTag(hostname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeDedicatedHostExists("ibm_compute_dedicated_host.dedicatedMonthly", &dedicatedHost),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedMonthly", "hostname", hostname),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedMonthly", "domain", "uat.com"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedMonthly", "datacenter", "dal10"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedMonthly", "router_hostname", "bcr01a.dal10"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedMonthly", "hourly_billing", "false"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedMonthly", "cpu_count", "56"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedMonthly", "disk_capacity", "1200"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedMonthly", "memory_capacity", "242"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedMonthly", "tags.#", "2"),
				),
			},

			{
				Config: testAccCheckIBMComputeDedicatedHostWithUpdateTag(hostname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeDedicatedHostExists("ibm_compute_dedicated_host.dedicatedMonthly", &dedicatedHost),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedMonthly", "hostname", hostname),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedMonthly", "domain", "uat.com"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedMonthly", "datacenter", "dal10"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedMonthly", "router_hostname", "bcr01a.dal10"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedMonthly", "hourly_billing", "false"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedMonthly", "cpu_count", "56"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedMonthly", "disk_capacity", "1200"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedMonthly", "memory_capacity", "242"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.dedicatedMonthly", "tags.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMComputeDedicatedHostImport(t *testing.T) {
	var dedicatedHost datatypes.Virtual_DedicatedHost
	hostname := fmt.Sprintf("terraformuat_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMComputerDedicatedHostDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMComputeDedicatedHostImport(hostname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeDedicatedHostExists("ibm_compute_dedicated_host.import", &dedicatedHost),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.import", "hostname", hostname),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.import", "domain", "uat.com"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.import", "datacenter", "dal10"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.import", "router_hostname", "bcr01a.dal10"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.import", "hourly_billing", "true"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.import", "cpu_count", "56"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.import", "disk_capacity", "1200"),
					resource.TestCheckResourceAttr(
						"ibm_compute_dedicated_host.import", "memory_capacity", "242"),
				),
			},

			{
				ResourceName:      "ibm_compute_dedicated_host.import",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes",
					"hourly_billing",
					"domain", "flavor",
				},
			},
		},
	})
}

func testAccCheckIBMComputerDedicatedHostDestroy(s *terraform.State) error {
	service := services.GetVirtualDedicatedHostService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_compute_dedicated_host" {
			continue
		}

		dedicatedId, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the key
		_, err := service.Id(dedicatedId).GetObject()

		if err == nil {
			return fmt.Errorf("Dedicated host still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("[ERROR] Error waiting for dedicated host (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMComputeDedicatedHostExists(n string, dedicatedHost *datatypes.Virtual_DedicatedHost) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("[ERROR] Not  found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No Record ID is set")
		}

		dedicatedID, _ := strconv.Atoi(rs.Primary.ID)

		service := services.GetVirtualDedicatedHostService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())
		result, err := service.Id(dedicatedID).GetObject()
		if err != nil {
			return err
		}
		if strconv.Itoa(int(*result.Id)) != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		*dedicatedHost = result

		return nil
	}
}

func testAccCheckIBMComputeDedicatedHostConfigBasic(hostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_dedicated_host" "dedicatedHourly" {
	hostname = "%s"
	domain = "uat.com"
	router_hostname = "bcr01a.dal10"
	datacenter = "dal10"
}`, hostname)
}

func testAccCheckIBMComputeDedicatedHostConfigUpdated(updatedHostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_dedicated_host" "dedicatedHourly" {
	hostname = "%s"
	domain = "uat.com"
	datacenter = "dal10"
	router_hostname = "bcr01a.dal10"
}`, updatedHostname)
}

func testAccCheckIBMComputeDedicatedHostWithTag(hostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_dedicated_host" "dedicatedMonthly" {
	hostname = "%s"
	domain = "uat.com"
	datacenter = "dal10"
	router_hostname = "bcr01a.dal10"
	hourly_billing = false
	tags = ["one", "two"]
}`, hostname)
}

func testAccCheckIBMComputeDedicatedHostWithUpdateTag(hostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_dedicated_host" "dedicatedMonthly" {
	hostname = "%s"
	domain = "uat.com"
	datacenter = "dal10"
	router_hostname = "bcr01a.dal10"
	hourly_billing = false
	tags = ["one", "two", "three"]
}`, hostname)
}

func testAccCheckIBMComputeDedicatedHostImport(hostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_dedicated_host" "import" {
	hostname = "%s"
	domain = "uat.com"
	router_hostname = "bcr01a.dal10"
	datacenter = "dal10"
}`, hostname)
}

// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMCisHealthcheck_Basic(t *testing.T) {
	var monitor string
	name := "ibm_cis_healthcheck.health_check"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		// No requirement for CheckDestory of this resource as by reaching this point it must have already been deleted from CIS.
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisHealthcheckConfigCisDSBasic("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisHealthcheckExists(name, &monitor),
					resource.TestCheckResourceAttr(name, "expected_body", "alive"),
				),
			},
		},
	})
}

func TestAccIBMCisHealthcheck_import(t *testing.T) {
	name := "ibm_cis_healthcheck.health_check"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckCisMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisHealthcheckConfigCisDSBasic("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "expected_body", "alive"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes"},
			},
		},
	})
}

func TestAccIBMCisHealthcheck_FullySpecified(t *testing.T) {
	var monitor string
	name := "ibm_cis_healthcheck.health_check"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckCisMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisHealthcheckConfigFullySpecified("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisHealthcheckExists(name, &monitor),
					resource.TestCheckResourceAttr(name, "path", "/custom"),
					resource.TestCheckResourceAttr(name, "retries", "3"),
					resource.TestCheckResourceAttr(name, "expected_codes", "5xx"),
				),
			},
		},
	})
}

func TestAccIBMCisHealthcheck_CreateAfterManualDestroy(t *testing.T) {
	var monitorOne, monitorTwo string
	name := "ibm_cis_healthcheck.health_check"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckCisMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisHealthcheckConfigCisDSBasic("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisHealthcheckExists(name, &monitorOne),
					testAccCisMonitorManuallyDelete(&monitorOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisHealthcheckConfigCisDSBasic("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisHealthcheckExists(name, &monitorTwo),
					func(state *terraform.State) error {
						if monitorOne == monitorTwo {
							return fmt.Errorf("load balancer monitor id is unchanged even after we thought we deleted it ( %s )",
								monitorTwo)
						}
						return nil
					},
				),
			},
		},
	})
}

func testAccCisMonitorManuallyDelete(tfMonitorID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisGLBHealthCheckClientSession()
		if err != nil {
			return err
		}
		tfMonitor := *tfMonitorID
		healthcheckID, cisID, _ := flex.ConvertTftoCisTwoVar(tfMonitor)
		cisClient.Crn = core.StringPtr(cisID)
		opt := cisClient.NewDeleteLoadBalancerMonitorOptions(healthcheckID)
		_, _, err = cisClient.DeleteLoadBalancerMonitor(opt)
		if err != nil {
			return fmt.Errorf("[ERROR] Error deleting IBMCISMonitor Record: %s", err)
		}
		return nil
	}
}

func testAccCheckCisMonitorDestroy(s *terraform.State) error {
	cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisGLBHealthCheckClientSession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_healthcheck" {
			continue
		}
		healthcheckID, cisID, _ := flex.ConvertTftoCisTwoVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		opt := cisClient.NewGetLoadBalancerMonitorOptions(healthcheckID)
		_, _, err = cisClient.GetLoadBalancerMonitor(opt)
		if err == nil {
			return fmt.Errorf("Load balancer Monitor still exists")
		}
	}

	return nil
}

func testAccCheckCisHealthcheckExists(n string, tfMonitorID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[ERROR] Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No Load Balancer Monitor ID is set")
		}

		cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisGLBHealthCheckClientSession()
		if err != nil {
			return err
		}
		healthcheckID, cisID, _ := flex.ConvertTftoCisTwoVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		opt := cisClient.NewGetLoadBalancerMonitorOptions(healthcheckID)
		foundHealthcheck, _, err := cisClient.GetLoadBalancerMonitor(opt)
		if err != nil {
			return fmt.Errorf("Load balancer Monitor still exists")
		}
		*tfMonitorID = flex.ConvertCisToTfTwoVar(*foundHealthcheck.Result.ID, cisID)
		return nil
	}
}

func testAccCheckCisHealthcheckConfigCisDSBasic(resourceID string, cisDomain string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + `
	resource "ibm_cis_healthcheck" "health_check" {
		cis_id         = data.ibm_cis.cis.id
		expected_body  = "alive"
		expected_codes = "2xx"
	  }
	`
}

func testAccCheckCisHealthcheckConfigFullySpecified(resourceID string, cisDomain string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + `
	resource "ibm_cis_healthcheck" "health_check" {
		cis_id         = data.ibm_cis.cis.id
		expected_body  = "dead"
		expected_codes = "5xx"
		method         = "HEAD"
		timeout        = 9
		path           = "/custom"
		interval       = 60
		retries        = 3
		description    = "this is a very weird load balancer"
		headers {
			header = "Host"
			values = ["example.com", "example1.com"]
		  }
		headers {
			header = "Host1"
			values = ["example3.com", "example11.com"]
		  }
	  }
	`
}

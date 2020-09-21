package ibm

import (
	"fmt"
	"log"
	"testing"

	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMCisHealthcheck_Basic(t *testing.T) {
	var monitor string
	name := "ibm_cis_healthcheck.health_check"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCis(t) },
		Providers: testAccProviders,
		// No requirement for CheckDestory of this resource as by reaching this point it must have already been deleted from CIS.
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisHealthcheckConfigCisDSBasic("test", cisDomainStatic),
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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisHealthcheckConfigCisDSBasic("test", cisDomainStatic),
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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisHealthcheckConfigFullySpecified("test", cisDomainStatic),
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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisHealthcheckConfigCisDSBasic("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisHealthcheckExists(name, &monitorOne),
					testAccCisMonitorManuallyDelete(&monitorOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisHealthcheckConfigCisDSBasic("test", cisDomainStatic),
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

func TestAccIBMCisHealthcheck_CreateAfterCisRIManualDestroy(t *testing.T) {
	t.Skip()
	var monitorOne, monitorTwo string
	name := "ibm_cis_healthcheck.health_check"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisHealthcheckConfigCisRIBasic("test", cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisHealthcheckExists(name, &monitorOne),
					testAccCisMonitorManuallyDelete(&monitorOne),
					func(state *terraform.State) error {
						cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
						if err != nil {
							return err
						}
						for _, r := range state.RootModule().Resources {
							if r.Type == "ibm_cis_domain" {
								log.Printf("[WARN] Manually removing domain")
								zoneID, cisID, _ := convertTftoCisTwoVar(r.Primary.ID)
								_ = cisClient.Zones().DeleteZone(cisID, zoneID)
								cisPtr := &cisID
								log.Printf("[WARN]  Manually removing Cis Instance")
								_ = testAccCisInstanceManuallyDeleteUnwrapped(state, cisPtr)
							}

						}
						return nil
					},
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisHealthcheckConfigCisRIBasic("test", cisDomainTest),
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
		cisClient, err := testAccProvider.Meta().(ClientSession).CisGLBHealthCheckClientSession()
		if err != nil {
			return err
		}
		tfMonitor := *tfMonitorID
		healthcheckID, cisID, _ := convertTftoCisTwoVar(tfMonitor)
		cisClient.Crn = core.StringPtr(cisID)
		opt := cisClient.NewDeleteLoadBalancerMonitorOptions(healthcheckID)
		_, _, err = cisClient.DeleteLoadBalancerMonitor(opt)
		if err != nil {
			return fmt.Errorf("Error deleting IBMCISMonitor Record: %s", err)
		}
		return nil
	}
}

func testAccCheckCisMonitorDestroy(s *terraform.State) error {
	cisClient, err := testAccProvider.Meta().(ClientSession).CisGLBHealthCheckClientSession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_healthcheck" {
			continue
		}
		healthcheckID, cisID, _ := convertTftoCisTwoVar(rs.Primary.ID)
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
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Load Balancer Monitor ID is set")
		}

		cisClient, err := testAccProvider.Meta().(ClientSession).CisGLBHealthCheckClientSession()
		if err != nil {
			return err
		}
		healthcheckID, cisID, _ := convertTftoCisTwoVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		opt := cisClient.NewGetLoadBalancerMonitorOptions(healthcheckID)
		foundHealthcheck, _, err := cisClient.GetLoadBalancerMonitor(opt)
		if err != nil {
			return fmt.Errorf("Load balancer Monitor still exists")
		}
		*tfMonitorID = convertCisToTfTwoVar(*foundHealthcheck.Result.ID, cisID)
		return nil
	}
}

func testAccCheckCisHealthcheckConfigCisDSBasic(resourceID string, cisDomain string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_healthcheck" "health_check" {
		cis_id         = data.ibm_cis.cis.id
		expected_body  = "alive"
		expected_codes = "2xx"
	  }
	`)
}

func testAccCheckCisHealthcheckConfigCisRIBasic(resourceID string, cisDomain string) string {
	return testAccCheckCisDomainConfigCisRIbasic(resourceID, cisDomain) + fmt.Sprintf(`
	resource "ibm_cis_healthcheck" "health_check" {
		cis_id         = ibm_cis.cis.id
		expected_body  = "alive"
		expected_codes = "2xx"
	  }
	`)
}

func testAccCheckCisHealthcheckConfigFullySpecified(resourceID string, cisDomain string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
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
	`)
}

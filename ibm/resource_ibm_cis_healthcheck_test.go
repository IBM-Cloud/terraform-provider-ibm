package ibm

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMCisHealthcheck_Basic(t *testing.T) {
	//t.Parallel()
	var monitor string
	name := "ibm_cis_healthcheck.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCis(t) },
		Providers: testAccProviders,
		// No requirement for CheckDestory of this resource as by reaching this point it must have already been deleted from CIS.
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisHealthcheckConfigCisDS_Basic("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisHealthcheckExists(name, &monitor),
					resource.TestCheckResourceAttr(name, "expected_body", "alive"),
					//resource.TestCheckResourceAttr(name, "header.#", "0"),
					// also expect api to generate some values
					//testAccCheckCisHealthcheckDates(name, &monitor, testStartTime),
				),
			},
		},
	})
}

func TestAccIBMCisHealthcheck_import(t *testing.T) {
	name := "ibm_cis_healthcheck.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckCisHealthcheckConfigCisDS_Basic("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "expected_body", "alive"),
				),
			},
			resource.TestStep{
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
	//t.Parallel()
	var monitor string
	name := "ibm_cis_healthcheck.test"

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
					resource.TestCheckResourceAttr(name, "retries", "5"),
					resource.TestCheckResourceAttr(name, "expected_codes", "5xx"),
				),
			},
		},
	})
}

func TestAccIBMCisHealthcheck_CreateAfterManualDestroy(t *testing.T) {
	t.Parallel()
	var monitorOne, monitorTwo string
	name := "ibm_cis_healthcheck.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisHealthcheckConfigCisDS_Basic("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisHealthcheckExists(name, &monitorOne),
					testAccCisMonitorManuallyDelete(&monitorOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisHealthcheckConfigCisDS_Basic("test", cisDomainStatic),
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
	t.Parallel()
	var monitorOne, monitorTwo string
	name := "ibm_cis_healthcheck.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisHealthcheckConfigCisRI_Basic("test", cisDomainTest),
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
								zoneId, cisId, _ := convertTftoCisTwoVar(r.Primary.ID)
								_ = cisClient.Zones().DeleteZone(cisId, zoneId)
								cisPtr := &cisId
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
				Config: testAccCheckCisHealthcheckConfigCisRI_Basic("test", cisDomainTest),
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

func testAccCisMonitorManuallyDelete(tfMonitorId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		if err != nil {
			return err
		}
		tfMonitor := *tfMonitorId
		monitorId, cisId, _ := convertTftoCisTwoVar(tfMonitor)
		err = cisClient.Monitors().DeleteMonitor(cisId, monitorId)
		if err != nil {
			return fmt.Errorf("Error deleting IBMCISMonitor Record: %s", err)
		}
		return nil
	}
}

func testAccCheckCisMonitorDestroy(s *terraform.State) error {
	cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_healthcheck" {
			continue
		}
		healthcheckId, cisId, _ := convertTftoCisTwoVar(rs.Primary.ID)
		_, err = cisClient.Monitors().GetMonitor(cisId, healthcheckId)
		if err == nil {
			return fmt.Errorf("Load balancer Monitor still exists")
		}
	}

	return nil
}

func testAccCheckCisHealthcheckExists(n string, tfMonitorId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Load Balancer Monitor ID is set")
		}

		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		healthcheckId, cisId, _ := convertTftoCisTwoVar(rs.Primary.ID)
		foundHealthcheck, err := cisClient.Monitors().GetMonitor(cisId, healthcheckId)
		if err != nil {
			return err
		}
		*tfMonitorId = convertCisToTfTwoVar(foundHealthcheck.Id, cisId)
		return nil
	}
}

func testAccCheckCisHealthcheckConfigCisDS_Basic(resourceId string, cisDomain string) string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1(resourceId, cisDomain) + fmt.Sprintf(`
resource "ibm_cis_healthcheck" "%[1]s" {
  cis_id = "${data.ibm_cis.%[2]s.id}"
  expected_body = "alive"
  expected_codes = "2xx"
}`, resourceId, cisInstance)
}

func testAccCheckCisHealthcheckConfigCisRI_Basic(resourceId string, cisDomain string) string {
	return testAccCheckCisDomainConfigCisRI_basic(resourceId, cisDomain) + fmt.Sprintf(`
resource "ibm_cis_healthcheck" "%[1]s" {
  cis_id = "${ibm_cis.%[2]s.id}"
  expected_body = "alive"
  expected_codes = "2xx"
  expected_body = "alive"
}`, resourceId, "testacc_ds_cis")
}

func testAccCheckCisHealthcheckConfigFullySpecified(resourceId string, cisDomain string) string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1(resourceId, cisDomain) + fmt.Sprintf(`
resource "ibm_cis_healthcheck" "%[1]s" {
  cis_id = "${data.ibm_cis.%[2]s.id}"
  expected_body = "dead"
  expected_codes = "5xx"
  method = "HEAD"
  timeout = 9
  path = "/custom"
  interval = 60
  retries = 5
  description = "this is a very weird load balancer"
}`, resourceId, cisInstance)
}

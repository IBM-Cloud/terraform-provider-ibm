package ibm

import (
	"fmt"
	"testing"

	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMCisHealthcheck_Basic(t *testing.T) {
	//t.Parallel()
	var monitor v1.Monitor
	name := "ibm_cis_healthcheck.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// No requirement for CheckDestory of this resource as by reaching this point it must have already been deleted from CIS.
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisHealthcheckConfigBasic("test", cis_domain),
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
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckCisHealthcheckConfigBasic("test", cis_domain),
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
	var monitor v1.Monitor
	name := "ibm_cis_healthcheck.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckCisHealthcheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisHealthcheckConfigFullySpecified("test", cis_domain),
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

func testAccCheckCisHealthcheckExists(n string, load *v1.Monitor) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Load Balancer Monitor ID is set")
		}

		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		healthcheckId, _, _ := convertTftoCisTwoVar(rs.Primary.ID)
		foundHealthcheck, err := cisClient.Monitors().GetMonitor(rs.Primary.Attributes["cis_id"], healthcheckId)
		if err != nil {
			return err
		}

		load = foundHealthcheck

		return nil
	}
}

func testAccCheckCisHealthcheckConfigBasic(resourceId string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic(resourceId, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_healthcheck" "%[1]s" {
  cis_id = "${ibm_cis.instance.id}"
  expected_body = "alive"
  expected_codes = "2xx"
}`, resourceId)
}

func testAccCheckCisHealthcheckConfigFullySpecified(resourceId string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic(resourceId, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_healthcheck" "%[1]s" {
  cis_id = "${ibm_cis.instance.id}"
  expected_body = "dead"
  expected_codes = "5xx"
  method = "HEAD"
  timeout = 9
  path = "/custom"
  interval = 60
  retries = 5
  description = "this is a very weird load balancer"
}`, resourceId)
}

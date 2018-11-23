package ibm

import (
	"fmt"
	"testing"

	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCisHealthcheck_Basic(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	t.Parallel()
	var monitor v1.Monitor
	name := "ibm_cis_healthcheck.test"
	//Fail if cis_crn not set
	if cis_crn == "" {
		panic("IBM_CIS_CRN environment variable not set - required to test CIS")
	}
	cisId := cis_crn

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckCisHealthcheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisHealthcheckConfigBasic(cisId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisHealthcheckExists(name, &monitor, cisId),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr(name, "expected_body", "alive"),
					//resource.TestCheckResourceAttr(name, "header.#", "0"),
					// also expect api to generate some values
					//testAccCheckCisHealthcheckDates(name, &monitor, testStartTime),
				),
			},
		},
	})
}

func TestAccCisHealthcheck_FullySpecified(t *testing.T) {
	t.Parallel()
	var monitor v1.Monitor
	name := "ibm_cis_healthcheck.test"
	//Fail if cis_crn not set
	if cis_crn == "" {
		panic("IBM_CIS_CRN environment variable not set - required to test CIS")
	}
	cisId := cis_crn

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckCisHealthcheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisHealthcheckConfigFullySpecified(cisId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisHealthcheckExists(name, &monitor, cisId),
					// checking our overrides of default values worked
					resource.TestCheckResourceAttr(name, "path", "/custom"),
					resource.TestCheckResourceAttr(name, "retries", "5"),
					resource.TestCheckResourceAttr(name, "expected_codes", "5xx"),
				),
			},
		},
	})
}

func TestAccCisHealthcheck_Update(t *testing.T) {
	t.Parallel()
	var monitor v1.Monitor
	var initialId string
	name := "ibm_cis_healthcheck.test"
	cisId := cis_crn

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckCisHealthcheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisHealthcheckConfigBasic(cisId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisHealthcheckExists(name, &monitor, cisId),
				),
			},
			{
				PreConfig: func() {
					initialId = monitor.Id
				},
				Config: testAccCheckCisHealthcheckConfigFullySpecified(cisId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisHealthcheckExists(name, &monitor, cisId),
					func(state *terraform.State) error {
						if initialId != monitor.Id {
							return fmt.Errorf("wanted update but monitor got recreated (id changed %q -> %q)",
								initialId, monitor.Id)
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccCisHealthcheck_CreateAfterManualDestroy(t *testing.T) {
	t.Parallel()
	var monitor v1.Monitor
	var initialId string
	name := "ibm_cis_healthcheck.test"
	cisId := cis_crn

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckCisHealthcheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisHealthcheckConfigBasic(cisId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisHealthcheckExists(name, &monitor, cisId),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisHealthcheckConfigBasic(cisId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisHealthcheckExists(name, &monitor, cisId),
					func(state *terraform.State) error {
						if initialId == monitor.Id {
							return fmt.Errorf("load balancer monitor id is unchanged even after we thought we deleted it ( %s )",
								monitor.Id)
						}
						return nil
					},
				),
			},
		},
	})
}

// func testAccCheckCisHealthcheckDestroy(s *terraform.State, cisId string) error {
// 	cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()

// 	for _, rs := range s.RootModule().Resources {
// 		if rs.Type != "ibm_cis_healthcheck" {
// 			continue
// 		}

// 		_, err := cisClient.Monitors().GetMonitor(cisId, rs.Primary.ID)
// 		if err == nil {
// 			return fmt.Errorf("Load balancer monitor still exists")
// 		}
// 	}

// 	return nil
// }

func testAccCheckCisHealthcheckExists(n string, load *v1.Monitor, cisId string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Load Balancer Monitor ID is set")
		}

		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		foundHealthcheck, err := cisClient.Monitors().GetMonitor(cisId, rs.Primary.ID)
		if err != nil {
			return err
		}

		load = foundHealthcheck

		return nil
	}
}

func testAccCheckCisHealthcheckConfigBasic(cis_crn string) string {
	return fmt.Sprintf(`
resource "ibm_cis_healthcheck" "test" {
  cis_id = "%s"
  expected_body = "alive"
  expected_codes = "2xx"
}`, cis_crn)
}

func testAccCheckCisHealthcheckConfigFullySpecified(cis_crn string) string {
	return fmt.Sprintf(`
resource "ibm_cis_healthcheck" "test" {
  cis_id = "%s"	
  expected_body = "dead"
  expected_codes = "5xx"
  method = "HEAD"
  timeout = 9
  path = "/custom"
  interval = 60
  retries = 5
  description = "this is a very weird load balancer"
}`, cis_crn)
}

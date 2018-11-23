package ibm

import (
	"fmt"
	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccCisPool_Basic(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	t.Parallel()
	var pool v1.Pool
	rnd := acctest.RandString(10)
	name := "ibm_cis_origin_pool." + rnd
	if cis_crn == "" {
		panic("IBM_CIS_CRN environment variable not set - required to test CIS")
	}
	cisId := cis_crn

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckCisPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisPoolConfigBasic(rnd, cisId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &pool, cisId),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr(name, "check_regions.#", "1"),
				),
			},
		},
	})
}

func TestAccCisPool_FullySpecified(t *testing.T) {
	t.Parallel()
	var pool v1.Pool
	rnd := acctest.RandString(10)
	name := "ibm_cis_origin_pool." + rnd
	if cis_crn == "" {
		panic("IBM_CIS_CRN environment variable not set - required to test CIS")
	}
	cisId := cis_crn

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckCisPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisPoolConfigFullySpecified(rnd, cisId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &pool, cisId),
					// checking our overrides of default values worked
					resource.TestCheckResourceAttr(name, "enabled", "false"),
					resource.TestCheckResourceAttr(name, "description", "tfacc-fully-specified"),
					resource.TestCheckResourceAttr(name, "check_regions.#", "1"),
					resource.TestCheckResourceAttr(name, "minimum_origins", "2"),
					resource.TestCheckResourceAttr(name, "notification_email", "admin@outlook.com"),
				),
			},
		},
	})
}

func testAccCheckCisPoolDestroy(s *terraform.State, cisId string) error {
	cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_origin_pool" {
			continue
		}

		_, err = cisClient.Pools().GetPool(cisId, rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Load balancer pool still exists")
		}
	}

	return nil
}

func testAccCheckCisPoolExists(n string, pool *v1.Pool, cisId string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Load Balancer ID is set")
		}

		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		if err != nil {
			return err
		}
		foundPool, err := cisClient.Pools().GetPool(cisId, rs.Primary.ID)
		if err != nil {
			return err
		}

		pool = foundPool

		return nil
	}
}

// using IPs from 192.0.2.0/24 as per RFC5737
func testAccCheckCisPoolConfigBasic(id string, cisId string) string {
	return fmt.Sprintf(`
resource "ibm_cis_origin_pool" "%[1]s" {
  cis_id = "%[2]s"
  name = "my-tf-pool-basic-%[1]s"
  check_regions = ["WEU"]
  description = "tfacc-fully-specified"
  origins {
    name = "example-1"
    address = "www.google.com"
    enabled = true
  }
}
`, id, cisId)
}

func testAccCheckCisPoolConfigFullySpecified(id string, cisId string) string {
	return testAccCheckCisHealthcheckConfigBasic(cisId) + fmt.Sprintf(`
resource "ibm_cis_origin_pool" "%[1]s" {
  cis_id = "%[2]s"
  name = "my-tf-pool-basic-%[1]s"
  notification_email = "admin@outlook.com"
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = false
  }
  origins {
    name = "example-2"
    address = "192.0.2.2"
    enabled = true
  }
  check_regions = ["WEU"]
  description = "tfacc-fully-specified"
  enabled = false
  minimum_origins = 2
  monitor = "${ibm_cis_healthcheck.test.id}"
}`, id, cisId)
}

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
	// cis_domain from environment var IBM_CIS_DOMAIN

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// Remove check destroy as this occurs after the CIS instance is deleted and fails with an auth error
		//CheckDestroy: testAccCheckCisPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisPoolConfigBasic(rnd, cis_domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &pool, cis_domain),
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

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckCisPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisPoolConfigFullySpecified(rnd, cis_domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &pool, cis_domain),
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

func testAccCheckCisPoolDestroy(s *terraform.State, cis_domain string) error {
	cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_origin_pool" {
			continue
		}

		_, err = cisClient.Pools().GetPool(rs.Primary.Attributes["cis_id"], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Load balancer pool still exists")
		}
	}

	return nil
}

func testAccCheckCisPoolExists(n string, pool *v1.Pool, cis_domain string) resource.TestCheckFunc {
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
		foundPool, err := cisClient.Pools().GetPool(rs.Primary.Attributes["cis_id"], rs.Primary.ID)
		if err != nil {
			return err
		}

		pool = foundPool

		return nil
	}
}

// using IPs from 192.0.2.0/24 as per RFC5737
func testAccCheckCisPoolConfigBasic(resourceId string, cis_domain string) string {
	return testAccCheckIBMCISInstance_basic(cis_domain) + fmt.Sprintf(`
resource "ibm_cis_origin_pool" "%[1]s" {
  cis_id = "${ibm_cis.instance.id}"
  name = "my-tf-pool-basic-%[1]s"
  check_regions = ["WEU"]
  description = "tfacc-fully-specified"
  origins {
    name = "example-1"
    address = "www.google.com"
    enabled = true
  }
}
`, resourceId)
}

func testAccCheckCisPoolConfigFullySpecified(resourceId string, cis_domain string) string {
	return testAccCheckCisHealthcheckConfigBasic(resourceId, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_origin_pool" "%[1]s" {
  cis_id = "${ibm_cis.instance.id}"
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
  monitor = "${ibm_cis_healthcheck.%[1]s.id}"
}`, resourceId)
}

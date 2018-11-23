package ibm

import (
	//"errors"
	"fmt"
	"testing"

	//"regexp"

	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCisGlb_Basic(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	t.Parallel()
	if cis_crn == "" {
		panic("IBM_CIS_CRN environment variable not set - required to test CIS")
	}
	if cis_domain == "" {
		panic("IBM_CIS_DOMAIN environment variable not set - required to test CIS")
	}
	cisId := cis_crn
	var glb v1.Glb

	rnd := acctest.RandString(10)
	name := "ibm_cis_global_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckCisGlbDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisGlbConfigBasic(cis_domain, rnd, cisId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glb, cisId),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					//resource.TestCheckResourceAttr(name, "pop_pools.#", "0"),
					//resource.TestCheckResourceAttr(name, "region_pools.#", "0"),
					// also expect api to generate some values
					resource.TestCheckResourceAttr(name, "proxied", "false"), // default value
				),
			},
		},
	})
}

func TestAccCisGlb_SessionAffinity(t *testing.T) {
	t.Parallel()
	if cis_crn == "" {
		panic("IBM_CIS_CRN environment variable not set - required to test CIS")
	}
	cisId := cis_crn
	var glb v1.Glb
	rnd := acctest.RandString(10)
	name := "ibm_cis_global_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckCisGlbDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisGlbConfigSessionAffinity(cis_domain, rnd, cisId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glb, cisId),
					// explicitly verify that our session_affinity has been set
					//resource.TestCheckResourceAttr(name, "session_affinity", "cookie"),
					// dont check that other specified values are set, this will be evident by lack
					// of plan diff some values will get empty values
					//resource.TestCheckResourceAttr(name, "pop_pools.#", "0"),
					//resource.TestCheckResourceAttr(name, "region_pools.#", "0"),
				),
			},
		},
	})
}

func testAccCheckCisGlbExists(n string, glb *v1.Glb, cisId string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Load Balancer ID is set")
		}
		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		foundGlb, err := cisClient.Glbs().GetGlb(cisId, rs.Primary.Attributes["domain_id"], rs.Primary.ID)
		if err != nil {
			return err
		}

		glb = foundGlb

		return nil
	}
}

func testAccCheckCisGlbConfigBasic(cis_domain string, id string, cisId string) string {
	return testAccIBMCisDomainConfig_basic("test", cisId, cis_domain) + testAccCheckCisPoolConfigBasic(id, cisId) + fmt.Sprintf(`
resource "ibm_cis_global_load_balancer" "%[2]s" {
  cis_id = "%[3]s"	
  domain_id = "${ibm_cis_domain.test.id}"
  name = "%[1]s"
  fallback_pool_id = "${ibm_cis_origin_pool.%[2]s.id}"
  default_pool_ids = ["${ibm_cis_origin_pool.%[2]s.id}"]
}`, cis_domain, id, cisId)
}

func testAccCheckCisGlbConfigSessionAffinity(cis_domain string, id string, cisId string) string {
	return testAccIBMCisDomainConfig_basic("test", cisId, cis_domain) + testAccCheckCisPoolConfigBasic(id, cisId) + fmt.Sprintf(`
resource "ibm_cis_global_load_balancer" "%[2]s" {
  cis_id = "%[3]s"
  domain_id = "${ibm_cis_domain.test.id}"
  name = "%[1]s"
  fallback_pool_id = "${ibm_cis_origin_pool.%[2]s.id}"
  default_pool_ids = ["${ibm_cis_origin_pool.%[2]s.id}"]
  session_affinity = "cookie"
}`, cis_domain, id, cisId)
}

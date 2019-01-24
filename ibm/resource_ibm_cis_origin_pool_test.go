package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"log"
	"testing"
)

func TestAccIBMCisPool_Basic(t *testing.T) {
	//t.Parallel()
	var pool string
	rnd := acctest.RandString(10)
	name := "ibm_cis_origin_pool." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCis(t) },
		Providers: testAccProviders,
		// No requirement for CheckDestory of this resource as by reaching this test it must have already been deleted
		// correctly during the resource destroy phase of test. The destroy of resource_ibm_cis used in testAccCheckCisPoolConfigCisDS_Basic
		// will fail if this resource is not correctly deleted.
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisPoolConfigCisDS_Basic(rnd, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &pool),
					resource.TestCheckResourceAttr(name, "check_regions.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMCisPool_import(t *testing.T) {
	name := "ibm_cis_origin_pool.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckCisPoolConfigCisDS_Basic("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "check_regions.#", "1"),
					resource.TestCheckResourceAttr(name, "origins.#", "1"),
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

func TestAccIBMCisPool_FullySpecified(t *testing.T) {
	//t.Parallel()
	var pool string
	rnd := acctest.RandString(10)
	name := "ibm_cis_origin_pool." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisPoolConfigFullySpecified(rnd, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &pool),
					resource.TestCheckResourceAttr(name, "enabled", "false"),
					resource.TestCheckResourceAttr(name, "description", "tfacc-fully-specified"),
					resource.TestCheckResourceAttr(name, "check_regions.#", "1"),
					resource.TestCheckResourceAttr(name, "minimum_origins", "2"),
					resource.TestCheckResourceAttr(name, "notification_email", "admin@outlook.com"),
					resource.TestCheckResourceAttr(name, "origins.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMCisPool_CreateAfterManualDestroy(t *testing.T) {
	//t.Parallel()
	var poolOne, poolTwo string
	testName := "test_acc"
	name := "ibm_cis_origin_pool." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisPoolConfigCisDS_Basic(testName, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &poolOne),
					testAccCisPoolManuallyDelete(&poolOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisPoolConfigCisDS_Basic(testName, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &poolTwo),
					func(state *terraform.State) error {
						if poolOne == poolTwo {
							return fmt.Errorf("id is unchanged even after we thought we deleted it ( %s )",
								poolTwo)
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccIBMCisPool_CreateAfterCisRIManualDestroy(t *testing.T) {
	//t.Parallel()
	var poolOne, poolTwo string
	testName := "test"
	name := "ibm_cis_origin_pool." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisPoolConfigCisRI_Basic(testName, cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &poolOne),
					testAccCisPoolManuallyDelete(&poolOne),
					func(state *terraform.State) error {
						cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
						if err != nil {
							return err
						}
						for _, r := range state.RootModule().Resources {
							if r.Type == "ibm_cis_domain" {
								log.Printf("[WARN] Removing domain")
								zoneId, cisId, _ := convertTftoCisTwoVar(r.Primary.ID)
								_ = cisClient.Zones().DeleteZone(cisId, zoneId)
								cisPtr := &cisId
								log.Printf("[WARN] Removing Cis Instance")
								_ = testAccCisInstanceManuallyDeleteUnwrapped(state, cisPtr)
							}

						}
						return nil
					},
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisPoolConfigCisRI_Basic(testName, cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &poolTwo),
					func(state *terraform.State) error {
						if poolOne == poolTwo {
							return fmt.Errorf("id is unchanged even after we thought we deleted it ( %s )",
								poolTwo)
						}
						return nil
					},
				),
			},
		},
	})
}

func testAccCheckCisPoolDestroy(s *terraform.State) error {
	cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_origin_pool" {
			continue
		}
		poolId, cisId, _ := convertTftoCisTwoVar(rs.Primary.ID)
		_, err = cisClient.Pools().GetPool(cisId, poolId)
		if err == nil {
			return fmt.Errorf("Load balancer pool still exists")
		}
	}

	return nil
}

func testAccCheckCisPoolExists(n string, tfPoolId *string) resource.TestCheckFunc {
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

		poolId, cisId, _ := convertTftoCisTwoVar(rs.Primary.ID)
		foundPoolPtr, err := cisClient.Pools().GetPool(rs.Primary.Attributes["cis_id"], poolId)
		if err != nil {
			return err
		}

		foundPool := *foundPoolPtr
		if foundPool.Id != poolId {
			return fmt.Errorf("Record not found")
		}

		tfPool := convertCisToTfTwoVar(foundPool.Id, cisId)
		*tfPoolId = tfPool
		return nil
	}
}

func testAccCisPoolManuallyDelete(tfPoolId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		log.Printf("[WARN] Manually removing pool")
		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		if err != nil {
			return err
		}
		tfPool := *tfPoolId
		poolId, cisId, _ := convertTftoCisTwoVar(tfPool)
		err = cisClient.Pools().DeletePool(cisId, poolId)
		if err != nil {
			return fmt.Errorf("Error deleting IBMCISPool Record: %s", err)
		}
		return nil
	}
}

func testAccCheckCisPoolConfigCisDS_Basic(resourceId string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1(resourceId, cisDomainStatic) + fmt.Sprintf(`
resource "ibm_cis_origin_pool" "%[1]s" {
  cis_id = "${data.ibm_cis.%[2]s.id}"
  name = "my-tf-pool-basic-%[1]s"
  check_regions = ["WEU"]
  description = "tfacc-fully-specified"
  origins {
    name = "example-1"
    address = "www.google.com"
    enabled = true
    weight = 1
  }
}
`, resourceId, cisInstance)
}

func testAccCheckCisPoolConfigCisRI_Basic(resourceId string, cisDomain string) string {
	return testAccCheckCisDomainConfigCisRI_basic(resourceId, cisDomain) + fmt.Sprintf(`
resource "ibm_cis_origin_pool" "%[1]s" {
  cis_id = "${ibm_cis.%[2]s.id}"
  name = "my-tf-pool-basic-%[1]s"
  check_regions = ["WEU"]
  description = "tfacc-fully-specified"
  origins {
    name = "example-1"
    address = "www.google.com"
    enabled = true
    weight = 1
  }
}
`, resourceId, "testacc_ds_cis")
}

func testAccCheckCisPoolConfigFullySpecified(resourceId string, cisDomainStatic string) string {
	return testAccCheckCisHealthcheckConfigCisDS_Basic(resourceId, cisDomainStatic) + fmt.Sprintf(`
resource "ibm_cis_origin_pool" "%[1]s" {
  cis_id = "${data.ibm_cis.%[2]s.id}"
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
}`, resourceId, cisInstance)
}

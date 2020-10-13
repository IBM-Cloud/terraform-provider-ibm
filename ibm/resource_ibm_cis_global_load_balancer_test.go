package ibm

import (
	//"errors"
	"fmt"
	"log"
	"testing"

	//"regexp"

	//"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMCisGlb_Basic(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	//t.Parallel()
	var glb string

	name := "ibm_cis_global_load_balancer." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckCis(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisGlbDestroy,
		Steps: []resource.TestStep{
			{
				Config:             testAccCheckCisGlbConfigCisDS_Basic("test", cisDomainStatic),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glb),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr(name, "pop_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "region_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "proxied", "false"), // default value
				),
			},
			{
				Config:             testAccCheckCisGlbConfigCisDS_Update("test", cisDomainStatic),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glb),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr(name, "pop_pools.#", "1"),
					resource.TestCheckResourceAttr(name, "region_pools.#", "1"),
					resource.TestCheckResourceAttr(name, "proxied", "false"), // default value
				),
			},
		},
	})
}

func TestAccIBMCisGlb_CreateAfterManualDestroy(t *testing.T) {
	//t.Parallel()
	t.Skip()
	var glbOne, glbTwo string
	name := "ibm_cis_global_load_balancer." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisGlbDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisGlbConfigCisDS_Basic("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glbOne),
					testAccCisGlbManuallyDelete(&glbOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisGlbConfigCisDS_Basic("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glbTwo),
					func(state *terraform.State) error {
						if glbOne == glbTwo {
							return fmt.Errorf("load balancer id is unchanged even after we thought we deleted it ( %s )",
								glbTwo)
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccIBMCisGlb_CreateAfterManualCisRIDestroy(t *testing.T) {
	//t.Parallel()
	t.Skip()
	var glbOne, glbTwo string
	name := "ibm_cis_global_load_balancer." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisGlbDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisGlbConfigCisRI_Basic("test", cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glbOne),
					testAccCisGlbManuallyDelete(&glbOne),
					func(state *terraform.State) error {
						cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
						if err != nil {
							return err
						}
						for _, r := range state.RootModule().Resources {
							if r.Type == "ibm_cis_pool" {
								log.Printf("[WARN]  Manually removing pool")
								poolId, cisId, _ := convertTftoCisTwoVar(r.Primary.ID)
								_ = cisClient.Pools().DeletePool(cisId, poolId)

							}

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
				Config: testAccCheckCisGlbConfigCisRI_Basic("test", cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glbTwo),
					func(state *terraform.State) error {
						if glbOne == glbTwo {
							return fmt.Errorf("load balancer id is unchanged even after we thought we deleted it ( %s )",
								glbTwo)
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccIBMCisGlb_import(t *testing.T) {
	name := "ibm_cis_global_load_balancer.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckCisGlbConfigCisDS_Basic("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "proxied", "false"), // default value
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

func TestAccIBMCisGlb_SessionAffinity(t *testing.T) {
	//t.Parallel()
	var glb string
	name := "ibm_cis_global_load_balancer." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisGlbConfigSessionAffinity("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glb),
					// explicitly verify that our session_affinity has been set
					resource.TestCheckResourceAttr(name, "session_affinity", "cookie"),
					//resource.TestCheckResourceAttr(name, "pop_pools.#", "0"),
					//resource.TestCheckResourceAttr(name, "region_pools.#", "0"),
				),
			},
		},
	})
}

func testAccCheckCisGlbDestroy(s *terraform.State) error {
	cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_global_load_balancer" {
			continue
		}
		glbId, zoneId, cisId, _ := convertTfToCisThreeVar(rs.Primary.ID)
		_, err = cisClient.Glbs().GetGlb(cisId, zoneId, glbId)
		if err == nil {
			return fmt.Errorf("Global Load balancer still exists")
		}
	}

	return nil
}

func testAccCisGlbManuallyDelete(tfGlbId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		if err != nil {
			return err
		}
		tfGlb := *tfGlbId
		log.Printf("[WARN] Manually removing glb")
		glbId, zoneId, cisId, _ := convertTfToCisThreeVar(tfGlb)
		err = cisClient.Glbs().DeleteGlb(cisId, zoneId, glbId)
		if err != nil {
			return fmt.Errorf("Error deleting IBMCISGlb Record: %s", err)
		}
		return nil
	}
}

func testAccCheckCisGlbExists(n string, tfGlbId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Load Balancer ID is set")
		}
		glbId, zoneId, cisId, _ := convertTfToCisThreeVar(rs.Primary.ID)
		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		foundGlb, err := cisClient.Glbs().GetGlb(rs.Primary.Attributes["cis_id"], zoneId, glbId)
		if err != nil {
			return err
		}
		*tfGlbId = convertCisToTfThreeVar(foundGlb.Id, zoneId, cisId)
		return nil
	}
}

func testAccCheckCisGlbConfigCisDS_Basic(id string, cisDomain string) string {
	return testAccCheckCisPoolConfigFullySpecified(id, cisDomain) + fmt.Sprintf(`
	resource "ibm_cis_global_load_balancer" "%[1]s" {
		cis_id           = data.ibm_cis.cis.id
		domain_id        = data.ibm_cis_domain.cis_domain.id
		name             = "%[2]s"
		fallback_pool_id = ibm_cis_origin_pool.origin_pool.id
		default_pool_ids = [ibm_cis_origin_pool.origin_pool.id]
	  }
	`, id, cisDomainStatic, cisInstance)
}

func testAccCheckCisGlbConfigCisDS_Update(id string, cisDomain string) string {
	return testAccCheckCisPoolConfigFullySpecified(id, cisDomain) + fmt.Sprintf(`
	resource "ibm_cis_global_load_balancer" "%[1]s" {
		cis_id           = data.ibm_cis.cis.id
		domain_id        = data.ibm_cis_domain.cis_domain.id
		name             = "%[2]s"
		fallback_pool_id = ibm_cis_origin_pool.origin_pool.id
		default_pool_ids = [ibm_cis_origin_pool.origin_pool.id]
		region_pools{
			region="WEU"
			pool_ids = [ibm_cis_origin_pool.origin_pool.id]
		}
		pop_pools{
			pop="LAX"
			pool_ids = [ibm_cis_origin_pool.origin_pool.id]
		}
	  }
	`, id, cisDomainStatic, cisInstance)
}

func testAccCheckCisGlbConfigCisRI_Basic(id string, cisDomain string) string {
	return testAccCheckCisPoolConfigCisRIBasic(id, cisDomain) + fmt.Sprintf(`
	resource "ibm_cis_global_load_balancer" "%[1]s" {
		cis_id           = ibm_cis.cis.id
		domain_id        = ibm_cis_domain.cis_domain.id
		name             = "%[2]s"
		fallback_pool_id = ibm_cis_origin_pool.origin_pool.id
		default_pool_ids = [ibm_cis_origin_pool.origin_pool.id]
	  }	  
	`, id, cisDomain, "testacc_ds_cis")
}

func testAccCheckCisGlbConfigSessionAffinity(id string, cisDomainStatic string) string {
	return testAccCheckCisPoolConfigFullySpecified(id, cisDomainStatic) + fmt.Sprintf(`
	resource "ibm_cis_global_load_balancer" "%[1]s" {
		cis_id           = data.ibm_cis.cis.id
		domain_id        = data.ibm_cis_domain.cis_domain.id
		name             = "%[2]s"
		fallback_pool_id = ibm_cis_origin_pool.origin_pool.id
		default_pool_ids = [ibm_cis_origin_pool.origin_pool.id]
		session_affinity = "cookie"
	  }
	`, id, cisDomainStatic, cisInstance)
}

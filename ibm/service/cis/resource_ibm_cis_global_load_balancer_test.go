// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"log"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMCisGlb_Basic(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	//t.Parallel()
	var glb string

	name := "ibm_cis_global_load_balancer." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCis(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckCisGlbDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisGlbConfigCisDSBasic("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glb),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr(name, "pop_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "region_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "proxied", "false"), // default value
					resource.TestCheckResourceAttr(name, "steering_policy", "dynamic_latency"),
				),
			},
			{
				Config: testAccCheckCisGlbConfigCisDSUpdate("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glb),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr(name, "pop_pools.#", "1"),
					resource.TestCheckResourceAttr(name, "region_pools.#", "1"),
					resource.TestCheckResourceAttr(name, "proxied", "false"), // default value
					resource.TestCheckResourceAttr(name, "steering_policy", "dynamic_latency"),
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
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckCisGlbDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisGlbConfigCisDSBasic("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glbOne),
					testAccCisGlbManuallyDelete(&glbOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisGlbConfigCisDSBasic("test", acc.CisDomainStatic),
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
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisGlbConfigCisDSBasic("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "proxied", "false"), // default value
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMCisGlb_SessionAffinity(t *testing.T) {
	//t.Parallel()
	var glb string
	name := "ibm_cis_global_load_balancer." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisGlbConfigSessionAffinity("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glb),
					// explicitly verify that our session_affinity has been set
					resource.TestCheckResourceAttr(name, "session_affinity", "cookie"),
				),
			},
		},
	})
}

func testAccCheckCisGlbDestroy(s *terraform.State) error {
	cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisGLBClientSession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_global_load_balancer" {
			continue
		}
		glbID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(rs.Primary.ID)
		if err != nil {
			return err
		}
		cisClient.Crn = core.StringPtr(crn)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)

		opt := cisClient.NewGetLoadBalancerSettingsOptions(glbID)

		_, _, err = cisClient.GetLoadBalancerSettings(opt)
		if err == nil {
			return fmt.Errorf("Global Load balancer still exists")
		}
	}

	return nil
}

func testAccCisGlbManuallyDelete(tfGlbID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisGLBClientSession()
		if err != nil {
			return err
		}
		tfGlb := *tfGlbID
		log.Printf("[WARN] Manually removing glb")
		glbID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(tfGlb)
		if err != nil {
			return err
		}
		cisClient.Crn = core.StringPtr(crn)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewDeleteLoadBalancerOptions(glbID)
		_, _, err = cisClient.DeleteLoadBalancer(opt)
		if err != nil {
			return fmt.Errorf("[ERROR] Error deleting IBMCISGlb Record: %s", err)
		}
		return nil
	}
}

func testAccCheckCisGlbExists(n string, tfGlbID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[ERROR] Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No Load Balancer ID is set")
		}
		cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisGLBClientSession()
		if err != nil {
			return err
		}
		glbID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(rs.Primary.ID)
		if err != nil {
			return err
		}
		cisClient.Crn = core.StringPtr(crn)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)

		opt := cisClient.NewGetLoadBalancerSettingsOptions(glbID)

		result, _, err := cisClient.GetLoadBalancerSettings(opt)
		if err != nil {
			return fmt.Errorf("Global Load balancer exists")
		}
		*tfGlbID = flex.ConvertCisToTfThreeVar(*result.Result.ID, zoneID, crn)
		return nil
	}
}

func testAccCheckCisGlbConfigCisDSBasic(id string, cisDomain string) string {
	return testAccCheckCisPoolConfigFullySpecified(id, cisDomain) + fmt.Sprintf(`
	resource "ibm_cis_global_load_balancer" "%[1]s" {
		cis_id           = data.ibm_cis.cis.id
		domain_id        = data.ibm_cis_domain.cis_domain.id
		name             = "%[2]s"
		fallback_pool_id = ibm_cis_origin_pool.origin_pool.id
		default_pool_ids = [ibm_cis_origin_pool.origin_pool.id]
		steering_policy = "dynamic_latency" 
	  }
	`, id, acc.CisDomainStatic)
}

func testAccCheckCisGlbConfigCisDSUpdate(id string, cisDomain string) string {
	return testAccCheckCisPoolConfigFullySpecified(id, cisDomain) + fmt.Sprintf(`
	resource "ibm_cis_global_load_balancer" "%[1]s" {
		cis_id           = data.ibm_cis.cis.id
		domain_id        = data.ibm_cis_domain.cis_domain.id
		name             = "%[2]s"
		fallback_pool_id = ibm_cis_origin_pool.origin_pool.id
		default_pool_ids = [ibm_cis_origin_pool.origin_pool.id]
		steering_policy = "dynamic_latency"
		region_pools{
			region="WEU"
			pool_ids = [ibm_cis_origin_pool.origin_pool.id]
		}
		pop_pools{
			pop="LAX"
			pool_ids = [ibm_cis_origin_pool.origin_pool.id]
		}
	  }
	`, id, acc.CisDomainStatic)
}

func testAccCheckCisGlbConfigSessionAffinity(id string, CisDomainStatic string) string {
	return testAccCheckCisPoolConfigFullySpecified(id, acc.CisDomainStatic) + fmt.Sprintf(`
	resource "ibm_cis_global_load_balancer" "%[1]s" {
		cis_id           = data.ibm_cis.cis.id
		domain_id        = data.ibm_cis_domain.cis_domain.id
		name             = "%[2]s"
		fallback_pool_id = ibm_cis_origin_pool.origin_pool.id
		default_pool_ids = [ibm_cis_origin_pool.origin_pool.id]
		session_affinity = "cookie"
		steering_policy = "dynamic_latency" 
	  }
	`, id, acc.CisDomainStatic)
}

// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"log"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMCisPool_Basic(t *testing.T) {
	var pool string
	rnd := acctest.RandString(10)
	name := "ibm_cis_origin_pool.origin_pool"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		// No requirement for CheckDestory of this resource as by reaching this test it must have already been deleted
		// correctly during the resource destroy phase of test. The destroy of resource_ibm_cis used in testAccCheckCisPoolConfigCisDSBasic
		// will fail if this resource is not correctly deleted.
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisPoolConfigCisDSBasic(rnd, acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &pool),
					resource.TestCheckResourceAttr(name, "check_regions.#", "1"),
				),
			},
			{
				Config: testAccCheckCisPoolConfigCisDSUpdate(rnd, acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &pool),
					resource.TestCheckResourceAttr(name, "check_regions.#", "1"),
					resource.TestCheckResourceAttr(name, "description", "tfacc-update-specified"),
				),
			},
		},
	})
}

func TestAccIBMCisPool_import(t *testing.T) {
	name := "ibm_cis_origin_pool.origin_pool"
	rnd := acctest.RandString(10)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisPoolConfigCisDSBasic(rnd, acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "check_regions.#", "1"),
					resource.TestCheckResourceAttr(name, "origins.#", "1"),
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

func TestAccIBMCisPool_FullySpecified(t *testing.T) {
	var pool string
	rnd := acctest.RandString(10)
	name := "ibm_cis_origin_pool.origin_pool"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisPoolConfigFullySpecified(rnd, acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &pool),
					resource.TestCheckResourceAttr(name, "enabled", "false"),
					resource.TestCheckResourceAttr(name, "description", "tfacc-fully-specified"),
					resource.TestCheckResourceAttr(name, "check_regions.#", "2"),
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
	t.Skip()
	var poolOne, poolTwo string
	testName := "test_acc"
	name := "ibm_cis_origin_pool.origin_pool"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckCisPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisPoolConfigCisDSBasic(testName, acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &poolOne),
					testAccCisPoolManuallyDelete(&poolOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisPoolConfigCisDSBasic(testName, acc.CisDomainStatic),
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
	cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisGLBPoolClientSession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_origin_pool" {
			continue
		}
		poolID, cisID, _ := flex.ConvertTftoCisTwoVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		opt := cisClient.NewGetLoadBalancerPoolOptions(poolID)
		_, _, err := cisClient.GetLoadBalancerPool(opt)
		if err == nil {
			return fmt.Errorf("Load balancer pool still exists")
		}
	}

	return nil
}

func testAccCheckCisPoolExists(n string, tfPoolID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[ERROR] Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No Load Balancer ID is set")
		}

		cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisGLBPoolClientSession()
		if err != nil {
			return err
		}

		poolID, cisID, _ := flex.ConvertTftoCisTwoVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		opt := cisClient.NewGetLoadBalancerPoolOptions(poolID)
		result, resp, err := cisClient.GetLoadBalancerPool(opt)
		if err != nil {
			return fmt.Errorf("[ERROR] Error getting glb pool: %v", resp)
		}

		foundPool := result.Result
		if *foundPool.ID != poolID {
			return fmt.Errorf("Record not found")
		}

		tfPool := flex.ConvertCisToTfTwoVar(*foundPool.ID, cisID)
		*tfPoolID = tfPool
		return nil
	}
}

func testAccCisPoolManuallyDelete(tfPoolID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		log.Printf("[WARN] Manually removing pool")
		cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisGLBPoolClientSession()
		if err != nil {
			return err
		}
		tfPool := *tfPoolID
		poolID, cisID, _ := flex.ConvertTftoCisTwoVar(tfPool)
		cisClient.Crn = core.StringPtr(cisID)
		opt := cisClient.NewDeleteLoadBalancerPoolOptions(poolID)
		_, resp, err := cisClient.DeleteLoadBalancerPool(opt)
		if err != nil {
			return fmt.Errorf("[WARN] Delete GLB Pools failed : %v", resp)
		}
		return nil
	}
}

func testAccCheckCisPoolConfigCisDSBasic(resourceID string, CisDomainStatic string) string {
	return testAccCheckCisHealthcheckConfigCisDSBasic(resourceID, acc.CisDomainStatic) + fmt.Sprintf(`
	resource "ibm_cis_origin_pool" "origin_pool" {
		cis_id        = data.ibm_cis.cis.id
		name          = "my-tf-pool-basic-%[1]s"
		check_regions = ["WEU"]
		description   = "tfacc-fully-specified"
		origins {
		  name    = "example-1"
		  address = "www.google.com"
		  enabled = true
		  weight  = 1
		}
		enabled = false
		monitor = ibm_cis_healthcheck.health_check.id
	  }
	  `, resourceID)
}

func testAccCheckCisPoolConfigCisDSUpdate(resourceID string, CisDomainStatic string) string {
	return testAccCheckCisHealthcheckConfigCisDSBasic(resourceID, acc.CisDomainStatic) + fmt.Sprintf(`
	resource "ibm_cis_origin_pool" "origin_pool" {
		cis_id        = data.ibm_cis.cis.id
		name          = "my-tf-pool-update-%[1]s"
		check_regions = ["ENAM"]
		description   = "tfacc-update-specified"
		origins {
		  name    = "example-2"
		  address = "www.google2.com"
		  enabled = false
		  weight  = 0.5
		}
		enabled = true
		monitor = ibm_cis_healthcheck.health_check.monitor_id
	  }
	  `, resourceID)
}

func testAccCheckCisPoolConfigFullySpecified(resourceID string, CisDomainStatic string) string {
	return testAccCheckCisHealthcheckConfigCisDSBasic(resourceID, acc.CisDomainStatic) + fmt.Sprintf(`
	resource "ibm_cis_origin_pool" "origin_pool" {
		cis_id             = data.ibm_cis.cis.id
		name               = "my-tf-pool-basic-%[1]s"
		notification_email = "admin@outlook.com"
		origins {
		  name    = "example-1"
		  address = "150.0.0.1"
		  enabled = true
		}
		origins {
		  name    = "example-2"
		  address = "150.0.0.2"
		  enabled = true
		}
		check_regions   = ["WEU", "ENAM"]
		description     = "tfacc-fully-specified"
		enabled         = false
		minimum_origins = 2
		monitor         = ibm_cis_healthcheck.health_check.monitor_id
	  }
	`, resourceID)
}

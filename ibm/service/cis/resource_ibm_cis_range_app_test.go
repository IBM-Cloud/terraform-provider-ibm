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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMCisRangeApp_Basic(t *testing.T) {
	var app string
	name := "ibm_cis_range_app.app"
	originDNS := fmt.Sprintf("test-lb.%s", acc.CisDomainStatic)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisRangeAppConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisRangeAppExists(name, &app),
					resource.TestCheckResourceAttr(name, "origin_direct.#", "1"),
					resource.TestCheckResourceAttr(name, "protocol", "tcp/22"),
					resource.TestCheckResourceAttr(name, "dns_type", "CNAME"),
				),
			},
			{
				Config: testAccCheckCisRangeAppConfigUpdate(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisRangeAppExists(name, &app),
					resource.TestCheckResourceAttr(name, "origin_dns", originDNS),
					resource.TestCheckResourceAttr(name, "dns_type", "CNAME"),
					resource.TestCheckResourceAttr(name, "traffic_type", "direct"),
				),
			},
		},
	})
}

func TestAccIBMCisRangeApp_import(t *testing.T) {
	name := "ibm_cis_range_app.app"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisRangeAppConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "origin_direct.#", "1"),
					resource.TestCheckResourceAttr(name, "protocol", "tcp/22"),
					resource.TestCheckResourceAttr(name, "dns_type", "CNAME"),
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

func TestAccIBMCisRangeApp_CreateAfterManualDestroy(t *testing.T) {
	t.Skip()
	var appOne, appTwo string
	name := "ibm_cis_range_app.app"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckCisRangeAppDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisRangeAppConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisRangeAppExists(name, &appOne),
					testAccCisRangeAppManuallyDelete(&appOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisRangeAppConfigUpdate(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisRangeAppExists(name, &appTwo),
					func(state *terraform.State) error {
						if appOne == appTwo {
							return fmt.Errorf("id is unchanged even after we thought we deleted it ( %s )",
								appTwo)
						}
						return nil
					},
				),
			},
		},
	})
}

func testAccCheckCisRangeAppDestroy(s *terraform.State) error {
	cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisRangeAppClientSession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_range_app" {
			continue
		}
		rangeAppID, zoneID, cisID, _ := flex.ConvertTfToCisThreeVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewGetRangeAppOptions(rangeAppID)
		_, _, err := cisClient.GetRangeApp(opt)
		if err == nil {
			return fmt.Errorf("Range application still exists")
		}
	}

	return nil
}

func testAccCheckCisRangeAppExists(n string, tfRangeAppID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[ERROR] Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No range app ID is set")
		}

		cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisRangeAppClientSession()
		if err != nil {
			return err
		}

		rangeAppID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(crn)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewGetRangeAppOptions(rangeAppID)
		result, resp, err := cisClient.GetRangeApp(opt)
		if err != nil {
			return fmt.Errorf("[ERROR] Error getting range app: %v", resp)
		}

		foundRangeApp := result.Result
		if *foundRangeApp.ID != rangeAppID {
			return fmt.Errorf("Record not found")
		}

		tfRangeApp := flex.ConvertCisToTfThreeVar(*foundRangeApp.ID, zoneID, crn)
		*tfRangeAppID = tfRangeApp
		return nil
	}
}

func testAccCisRangeAppManuallyDelete(tfRangeAppID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		log.Printf("[WARN] Manually removing range app")
		cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisRangeAppClientSession()
		if err != nil {
			return err
		}
		tfRangeApp := *tfRangeAppID
		rangeAppID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(tfRangeApp)
		cisClient.Crn = core.StringPtr(crn)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewDeleteRangeAppOptions(rangeAppID)
		_, resp, err := cisClient.DeleteRangeApp(opt)
		if err != nil {
			return fmt.Errorf("[WARN] Delete range app failed : %v", resp)
		}
		return nil
	}
}

func testAccCheckCisRangeAppConfigBasic() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_range_app" "app" {
		cis_id         = data.ibm_cis.cis.id
		domain_id      = data.ibm_cis_domain.cis_domain.id
		protocol       = "tcp/22"
		dns_type       = "CNAME"
		dns            = "ssh.%[1]s"
		origin_direct  = ["tcp://12.1.1.1:22"]
		ip_firewall    = true
		proxy_protocol = "v1"
		traffic_type   = "direct"
		tls            = "off"
	  }
	  `, acc.CisDomainStatic)
}

func testAccCheckCisRangeAppConfigUpdate() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_origin_pool" "origin_pool" {
		cis_id             = data.ibm_cis.cis.id
		name               = "my-tf-pool-basic"
		notification_email = "admin@outlook.com"
		origins {
		  name    = "example-1"
		  address = "150.0.0.1"
		  enabled = false
		}
		origins {
		  name    = "example-2"
		  address = "150.0.0.2"
		  enabled = true
		  weight  = 0.5
		}
		check_regions   = ["WEU", "ENAM"]
		description     = "tfacc-fully-specified"
		enabled         = false
		minimum_origins = 2
	  }

	  resource "ibm_cis_global_load_balancer" "test_glb" {
		cis_id           = data.ibm_cis.cis.id
		domain_id        = data.ibm_cis_domain.cis_domain.id
		name             = "test-lb.%[1]s"
		fallback_pool_id = ibm_cis_origin_pool.origin_pool.id
		default_pool_ids = [ibm_cis_origin_pool.origin_pool.id]
		region_pools {
		  region   = "WEU"
		  pool_ids = [ibm_cis_origin_pool.origin_pool.id]
		}
		pop_pools {
		  pop      = "LAX"
		  pool_ids = [ibm_cis_origin_pool.origin_pool.id]
		}
		session_affinity = "cookie"
	  }

	  resource "ibm_cis_range_app" "app" {
		cis_id         = data.ibm_cis.cis.id
		domain_id      = data.ibm_cis_domain.cis_domain.id
		protocol       = "tcp/8081"
		dns_type       = "CNAME"
		dns            = "ssh1.%[1]s"
		origin_dns     = ibm_cis_global_load_balancer.test_glb.name
		origin_port    = 8081
		ip_firewall    = true
		proxy_protocol = "v1"
		traffic_type   = "direct"
		tls            = "off"
	  }`, acc.CisDomainStatic)
}

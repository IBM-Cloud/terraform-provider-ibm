// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMCisRateLimit_Basic(t *testing.T) {
	var record string
	resource.Test(t, resource.TestCase{
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisRateLimitDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisRateLimitConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisRateLimitExists("ibm_cis_rate_limit.ratelimit", &record),
					resource.TestCheckResourceAttr(
						"ibm_cis_rate_limit.ratelimit", "action.0.mode", "simulate"),
				),
			},
			{
				Config: testAccCheckIBMCisRateLimitConfigUpdate(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisRateLimitExists("ibm_cis_rate_limit.ratelimit", &record),
					resource.TestCheckResourceAttr(
						"ibm_cis_rate_limit.ratelimit", "action.0.mode", "ban"),
				),
			},
		},
	})
}

func TestAccIBMCisRateLimitWithoutMatchRequest_Basic(t *testing.T) {
	var record string
	resource.Test(t, resource.TestCase{
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisRateLimitDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisRateLimitConfigWithoutRequestBasic1(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisRateLimitExists("ibm_cis_rate_limit.ratelimit", &record),
					resource.TestCheckResourceAttr(
						"ibm_cis_rate_limit.ratelimit", "action.0.mode", "simulate"),
				),
			},
			{
				Config: testAccCheckIBMCisRateLimitConfigUpdate(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisRateLimitExists("ibm_cis_rate_limit.ratelimit", &record),
					resource.TestCheckResourceAttr(
						"ibm_cis_rate_limit.ratelimit", "action.0.mode", "ban"),
				),
			},
		},
	})
}

func TestAccIBMCisRateLimit_MultiDomain(t *testing.T) {
	var record, record2 string
	resource.Test(t, resource.TestCase{
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisRateLimitDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisRateLimitConfigWithMultiDomain(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisRateLimitExists("ibm_cis_rate_limit.ratelimit", &record),
					testAccCheckIBMCisRateLimitExists("ibm_cis_rate_limit.ratelimit2", &record2),
					resource.TestCheckResourceAttr(
						"ibm_cis_rate_limit.ratelimit", "action.0.mode", "ban"),
					resource.TestCheckResourceAttr(
						"ibm_cis_rate_limit.ratelimit", "action.0.mode", "ban"),
				),
			},
		},
	})
}

func TestAccIBMCisRateLimit_Import(t *testing.T) {

	var record string

	resource.Test(t, resource.TestCase{
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisRateLimitDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisRateLimitConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisRateLimitExists("ibm_cis_rate_limit.ratelimit", &record),
					resource.TestCheckResourceAttr(
						"ibm_cis_rate_limit.ratelimit", "action.0.mode", "simulate"),
					resource.TestCheckResourceAttr(
						"ibm_cis_rate_limit.ratelimit", "action.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMCisRateLimitConfigUpdate(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisRateLimitExists("ibm_cis_rate_limit.ratelimit", &record),
					resource.TestCheckResourceAttr(
						"ibm_cis_rate_limit.ratelimit", "action.0.mode", "ban"),
					resource.TestCheckResourceAttr(
						"ibm_cis_rate_limit.ratelimit", "action.#", "1"),
				),
			},
			{
				ResourceName:      "ibm_cis_rate_limit.ratelimit",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
func testAccCheckIBMCisRateLimitDestroy(s *terraform.State) error {
	cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisRLClientSession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_rate_limit" {
			continue
		}
		recordID, zoneID, cisID, _ := flex.ConvertTfToCisThreeVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewGetRateLimitOptions(recordID)
		_, _, err := cisClient.GetRateLimit(opt)
		if err == nil {
			return fmt.Errorf("Record still exists")
		}
	}

	return nil
}

func testAccCheckIBMCisRateLimitExists(n string, tfRecordID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[ERROR] Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No Record ID is set")
		}

		// tfRecord := *tfRecordID
		cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisRLClientSession()
		if err != nil {
			return err
		}
		recordID, zoneID, cisID, _ := flex.ConvertTfToCisThreeVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewGetRateLimitOptions(recordID)
		foundRecordPtr, _, err := cisClient.GetRateLimit(opt)
		if err != nil {
			return err
		}

		foundRecord := foundRecordPtr.Result
		if *foundRecord.ID != recordID {
			return fmt.Errorf("Record not found")
		}

		tfRecord := flex.ConvertCisToTfThreeVar(*foundRecord.ID, zoneID, cisID)
		*tfRecordID = tfRecord
		return nil
	}
}

func testAccCheckIBMCisRateLimitConfigBasic() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + `
	resource "ibm_cis_rate_limit" "ratelimit" {
		cis_id = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
		threshold = 20
		period = 900
		match {
			request {
				url = "*.example.org/path*"
				schemes = ["HTTP", "HTTPS"]
				methods = ["GET", "HEAD","POST", "PUT", "DELETE"]
			}
			response {
				status = [200, 201, 202, 301, 429]
				origin_traffic = false
				headers {
					name= "Cf-Cache-Status"
					op= "eq"
					value= "HIT"
				}
			}
		}
		action {
			mode = "simulate"
			timeout = 43200
			response {
				content_type = "text/plain"
				body = "custom response body"
			}
		}
		correlate {
			by = "nat"
		}
		disabled = false
		description = "example rate limit for a zone"
	}
	  `
}

func testAccCheckIBMCisRateLimitConfigWithoutRequestBasic1() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + `
	resource "ibm_cis_rate_limit" "ratelimit" {
		cis_id = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
		threshold = 20
		period = 900
		action {
			mode = "simulate"
			timeout = 43200
			response {
				content_type = "text/plain"
				body = "custom response body"
			}
		}
		correlate {
			by = "nat"
		}
		disabled = false
		description = "example rate limit for a zone"
	}
	  `
}

func testAccCheckIBMCisRateLimitConfigUpdate() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + `
	resource "ibm_cis_rate_limit" "ratelimit" {
		cis_id = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
		threshold = 20
		period = 900
		match {
			request {
				url = "*.example.org/path*"
				schemes = ["HTTP", "HTTPS"]
				methods = ["GET", "HEAD","POST", "PUT", "DELETE"]
			}
			response {
				status = [200, 201, 202, 301, 429]
				origin_traffic = false
				headers {
					name= "Cf-Cache-Status"
					op= "eq"
					value= "HIT"
				}
			}
		}
		action {
			mode = "ban"
			timeout = 43200
			response {
				content_type = "text/plain"
				body = "custom response body"
			}
		}
		correlate {
			by = "nat"
		}
		disabled = false
		description = "example rate limit for a zone"
	}
	  `
}

func testAccCheckIBMCisRateLimitConfigWithMultiDomain() string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	data "ibm_cis" "cis" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
	}

	resource "ibm_cis_domain" "domain2" {
		cis_id = data.ibm_cis.cis.id
		domain = "testdomain2.%[3]s"
	}

	resource "ibm_cis_domain" "domain" {
		cis_id = data.ibm_cis.cis.id
		domain = "testdomain.%[3]s"
	}

	data "ibm_cis_domain" "domain" {
		cis_id = data.ibm_cis.cis.id
		domain = ibm_cis_domain.domain.domain
	}

	data "ibm_cis_domain" "domain2" {
		cis_id = data.ibm_cis.cis.id
		domain = ibm_cis_domain.domain2.domain
	}

	resource "ibm_cis_rate_limit" "ratelimit" {
		threshold = 20
		period    = 900
		match {
		  request {
			url     = "zeubiiii.ibm.com/*"
			schemes = ["HTTP", "HTTPS"]
			methods = ["GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"]
		  }
		  response {
			status         = [200, 201, 202, 301, 429]
			origin_traffic = false
		  }
		}
		action {
		  mode    = "ban"
		  timeout = 43200
		  response {
			content_type = "text/plain"
			body         = "custom response body"
		  }
		}
		correlate {
		  by = "nat"
		}
		disabled    = false
		description = "LANDINGZONE 2 INTEGRATION"
		cis_id      = data.ibm_cis.cis.id
		domain_id   = data.ibm_cis_domain.domain.id
	}

	resource "ibm_cis_rate_limit" "ratelimit2" {
		threshold = 20
		period    = 900
		match {
		  request {
			url     = "zobaaaaaa.com/*"
			schemes = ["HTTP", "HTTPS"]
			methods = ["GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"]
		  }
		  response {
			status         = [200, 201, 202, 301, 429]
			origin_traffic = false
		  }
		}
		action {
		  mode    = "ban"
		  timeout = 43200
		  response {
			content_type = "text/plain"
			body         = "custom response body"
		  }
		}
		correlate {
		  by = "nat"
		}
		disabled    = false
		description = "SKT INTEGRATION"
		cis_id      = data.ibm_cis.cis.id
		domain_id   = data.ibm_cis_domain.domain2.id
	}
	`, acc.CisResourceGroup, acc.CisInstance, acc.CisDomainStatic)
}

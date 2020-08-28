package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMCisRateLimit_Basic(t *testing.T) {
	var record string
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
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
func TestAccIBMCisRateLimit_Import(t *testing.T) {

	var record string

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisRateLimitDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
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
			resource.TestStep{
				ResourceName:      "ibm_cis_rate_limit.ratelimit",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
func testAccCheckIBMCisRateLimitDestroy(s *terraform.State) error {
	cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_rate_limit" {
			continue
		}
		recordID, zoneID, cisID, _ := convertTfToCisThreeVar(rs.Primary.ID)
		_, err = cisClient.RateLimit().GetRateLimit(cisID, zoneID, recordID)
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
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		tfRecord := *tfRecordID
		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		recordID, zoneID, cisID, _ := convertTfToCisThreeVar(rs.Primary.ID)
		foundRecordPtr, err := cisClient.RateLimit().GetRateLimit(cisID, zoneID, recordID)
		if err != nil {
			return err
		}

		foundRecord := *foundRecordPtr
		if foundRecord.ID != recordID {
			return fmt.Errorf("Record not found")
		}

		tfRecord = convertCisToTfThreeVar(foundRecord.ID, zoneID, cisID)
		*tfRecordID = tfRecord
		return nil
	}
}

func testAccCheckIBMCisRateLimitConfigBasic() string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1() + fmt.Sprintf(`
	
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
	  `)
}

func testAccCheckIBMCisRateLimitConfigUpdate() string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1() + fmt.Sprintf(`
	
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
	  `)
}

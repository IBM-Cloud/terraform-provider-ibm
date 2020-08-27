package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisRateLimitDataSource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCisRateLimitDataSourceConfig_basic1(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cis_rate_limit.ratelimit", "cis_id"),
				),
			},
		},
	})
}

func testAccCheckIBMCisRateLimitDataSourceConfig_basic1() string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1() + fmt.Sprintf(`

	data "ibm_cis_rate_limit" "ratelimit" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
	}
	  
	`)
}

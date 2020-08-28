package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisDNSRecordsDataSource_basic(t *testing.T) {
	node := "data.ibm_network_cis_dns_records.test_dns_records"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCisDNSRecordsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_dns_records.0.cis_id"),
					resource.TestCheckResourceAttrSet(node, "cis_dns_records.0.domain_id"),
				),
			},
		},
	})
}

func testAccCheckIBMCisDNSRecordsDataSourceConfig() string {
	// status filter defaults to empty
	return testAccCheckIBMCisDomainDataSourceConfig_basic1() + fmt.Sprintf(`
	data "ibm_cis_dns_records" "test_dns_records" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
	  }`)
}

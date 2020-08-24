package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMNetworkCISDNSRecordsDataSource_basic(t *testing.T) {
	node := "data.ibm_network_cis_dns_records.test_dns_records"
	cisID := fmt.Sprintf("crn:v1:staging:public:internet-svcs-ci:global:a/01652b251c3ae2787110a995d8db0135:%d::", acctest.RandIntRange(100, 200))
	domainID := fmt.Sprintf("fd61359275217b573edbf741c89%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMNetworkCISDNSRecordsDataSourceConfig(cisID, domainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_dns_records.0.id"),
					resource.TestCheckResourceAttrSet(node, "cis_dns_records.0.cis_id"),
					resource.TestCheckResourceAttrSet(node, "cis_dns_records.0.domain_id"),
				),
			},
		},
	})
}

func testAccCheckIBMNetworkCISDNSRecordsDataSourceConfig(crn, zoneID string) string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	data "ibm_network_cis_dns_records" "test" {
		cis_id    = %s
		domain_id = %s
	  }`, crn, zoneID)
}

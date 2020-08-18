package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMNetworkCISDNSRecordsDataSource_basic(t *testing.T) {
	node := "data.ibm_network_cis_dns_records.test_dns_records"
	rname := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(100, 200))
	zonename := fmt.Sprintf("tf-dnszone-%d.com", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMNetworkCISDNSRecordsDataSourceConfig(rname, zonename),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_dns_records.0.id"),
					resource.TestCheckResourceAttrSet(node, "cis_dns_records.0.crn"),
					resource.TestCheckResourceAttrSet(node, "cis_dns_records.0.type"),
				),
			},
		},
	})
}

func testAccCheckIBMNetworkCISDNSRecordsDataSourceConfig(crn, zoneID string) string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	data "ibm_network_cis_dns_records" "test" {
		crn     = %s
		zone_id = %s
	  }`, crn, zoneID)
}

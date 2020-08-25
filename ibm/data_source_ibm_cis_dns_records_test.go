package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMNetworkCISDNSRecordsDataSource_basic(t *testing.T) {
	node := "data.ibm_network_cis_dns_records.test_dns_records"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMNetworkCISDNSRecordsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_dns_records.0.record_id"),
					resource.TestCheckResourceAttrSet(node, "cis_dns_records.0.cis_id"),
					resource.TestCheckResourceAttrSet(node, "cis_dns_records.0.domain_id"),
				),
			},
		},
	})
}

func testAccCheckIBMNetworkCISDNSRecordsDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`{
		data "ibm_resource_group" "web_group" {
		name = "Default"
	}
	data "ibm_cis" "web_instance" {
		name              = "CISTest"
		resource_group_id = data.ibm_resource_group.web_group.id
	}
		data "ibm_cis_domain" "web_domain" {
		cis_id = data.ibm_cis.web_instance.id
		domain = "cis-terraform.com"
	}
	data "ibm_network_cis_dns_records" "test" {
		cis_id           = data.ibm_cis.web_instance.id
		domain_id        = data.ibm_cis_domain.web_domain.id
	  }`)
}

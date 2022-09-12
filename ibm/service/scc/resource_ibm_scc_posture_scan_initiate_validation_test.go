package scc_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccPostureScanInitiateValidationBasic(t *testing.T) {
	name := "ibm_scc_posture_scan_initiate_validation." + "scans"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSccPostureScanInitiateValidationConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "result", "true"),
				),
			},
			{
				ResourceName: name,
				ImportState:  true,
			},
		},
	})
}

func testAccCheckSccPostureScanInitiateValidationConfigBasic() string {
	return `resource "ibm_scc_posture_scan_initiate_validation" "scans" {
			scope_id = "70324"
			profile_id = "425"
			name = "Test1Sept22_Scan"
			description = "Test1Sept22_Scan on scope 70324"
			frequency = 6300
			no_of_occurrences = 9
		}`
}

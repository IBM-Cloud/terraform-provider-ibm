// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.0-3c13b041-20250605-193116
 */

package drautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmPdrValidateClustertypeDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmPdrValidateClustertypeDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_validate_clustertype.pdr_validate_clustertype_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_validate_clustertype.pdr_validate_clustertype_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_validate_clustertype.pdr_validate_clustertype_instance", "orchestrator_cluster_type"),
				),
			},
		},
	})
}

func testAccCheckIbmPdrValidateClustertypeDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_validate_clustertype" "pdr_validate_clustertype_instance" {
			instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
			orchestrator_cluster_type = "on-premises"
			Accept-Language = "Accept-Language"
			If-None-Match = "If-None-Match"
		}
	`)
}

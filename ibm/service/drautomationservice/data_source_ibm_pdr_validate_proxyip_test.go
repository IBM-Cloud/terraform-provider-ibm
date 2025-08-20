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

func TestAccIbmPdrValidateProxyipDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmPdrValidateProxyipDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_validate_proxyip.pdr_validate_proxyip_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_validate_proxyip.pdr_validate_proxyip_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_validate_proxyip.pdr_validate_proxyip_instance", "proxyip"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_validate_proxyip.pdr_validate_proxyip_instance", "vpc_location"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_validate_proxyip.pdr_validate_proxyip_instance", "vpc_id"),
				),
			},
		},
	})
}

func testAccCheckIbmPdrValidateProxyipDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_validate_proxyip" "pdr_validate_proxyip_instance" {
			instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
			proxyip = "10.30.40.5:3128"
			vpc_location = "us-south"
			vpc_id = "r006-2f3b3ab9-2149-49cc-83a1-30a5d93d59b2"
			If-None-Match = "If-None-Match"
		}
	`)
}

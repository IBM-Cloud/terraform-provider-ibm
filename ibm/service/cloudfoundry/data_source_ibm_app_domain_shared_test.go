// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudfoundry_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMAppDomainSharedDataSource_basic(t *testing.T) {
	t.Skip()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppDomainSharedDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ibm_app_domain_shared.testacc_domain", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMAppDomainSharedDataSourceConfig() string {
	return `
	
		data "ibm_app_domain_shared" "testacc_domain" {
			name = "mybluemix.net"
		}`

}

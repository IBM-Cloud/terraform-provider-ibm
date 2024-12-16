// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMIamTrustedProfileLinksDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamTrustedProfileLinksDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_links.iam_trusted_profile_links", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_links.iam_trusted_profile_links", "profile_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_links.iam_trusted_profile_links", "links.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileLinksDataSourceConfigBasic() string {
	return `
		data "ibm_iam_trusted_profile_links" "iam_trusted_profile_links" {
			profile_id = "profile_id"
		}
	`
}

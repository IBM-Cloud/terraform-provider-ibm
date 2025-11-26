// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package globaltagging_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIamAccessTagDataSource_basic(t *testing.T) {
	name := "access:synthetics"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIamAccessTagReadDataSource(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_access_tag.example_access_tag", "id", name),
					resource.TestCheckResourceAttr("data.ibm_iam_access_tag.example_access_tag", "name", name),
					resource.TestCheckResourceAttr("data.ibm_iam_access_tag.example_access_tag", "tag_type", "access"),
				),
			},
		},
	})
}

func TestAccIamAccessTagDataSource_NotFoundExpectingError(t *testing.T) {
	name := "access:syntheticsxxxxx"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIamAccessTagReadDataSource(name),
				ExpectError: regexp.MustCompile("not found"),
			},
		},
	})
}

func testAccCheckIamAccessTagReadDataSource(name string) string {
	return fmt.Sprintf(`
	data "ibm_iam_access_tag" "example_access_tag" {
		name = "%s"
	}
`, name)
}

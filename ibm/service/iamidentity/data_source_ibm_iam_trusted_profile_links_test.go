// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.98.0-8be2046a-20241205-162752
 */

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/iamidentity"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIamTrustedProfileLinksDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileLinksDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_links.iam_trusted_profile_links_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_links.iam_trusted_profile_links_instance", "profile_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_links.iam_trusted_profile_links_instance", "links.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileLinksDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_iam_trusted_profile_links" "iam_trusted_profile_links_instance" {
			profile_id = "%s"
		}
	`, acc.IAMTrustedProfileID)
}

func TestDataSourceIBMIamTrustedProfileLinksProfileLinkToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		profileLinkLinkModel := make(map[string]interface{})
		profileLinkLinkModel["crn"] = "testString"
		profileLinkLinkModel["namespace"] = "testString"
		profileLinkLinkModel["name"] = "testString"

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["entity_tag"] = "testString"
		model["created_at"] = "2019-01-01T12:00:00.000Z"
		model["modified_at"] = "2019-01-01T12:00:00.000Z"
		model["name"] = "testString"
		model["cr_type"] = "testString"
		model["link"] = []map[string]interface{}{profileLinkLinkModel}

		assert.Equal(t, result, model)
	}

	profileLinkLinkModel := new(iamidentityv1.ProfileLinkLink)
	profileLinkLinkModel.CRN = core.StringPtr("testString")
	profileLinkLinkModel.Namespace = core.StringPtr("testString")
	profileLinkLinkModel.Name = core.StringPtr("testString")

	model := new(iamidentityv1.ProfileLink)
	model.ID = core.StringPtr("testString")
	model.EntityTag = core.StringPtr("testString")
	model.CreatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.ModifiedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.Name = core.StringPtr("testString")
	model.CrType = core.StringPtr("testString")
	model.Link = profileLinkLinkModel

	result, err := iamidentity.DataSourceIBMIamTrustedProfileLinksProfileLinkToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIamTrustedProfileLinksProfileLinkLinkToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["crn"] = "testString"
		model["namespace"] = "testString"
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(iamidentityv1.ProfileLinkLink)
	model.CRN = core.StringPtr("testString")
	model.Namespace = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")

	result, err := iamidentity.DataSourceIBMIamTrustedProfileLinksProfileLinkLinkToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

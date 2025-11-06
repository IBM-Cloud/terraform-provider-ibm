// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMTrustedProfileTemplateDataSourceBasic(t *testing.T) {
	enterpriseAccountId := acc.IamIdentityEnterpriseAccountId
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_desc_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckIamIdentityEnterpriseTemplates(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMTrustedProfileTemplateDataSourceConfigBasic(enterpriseAccountId, name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "template_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "committed"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "profile.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "crn"),
				),
			},
		},
	})
}

func TestAccIBMTrustedProfileTemplateDataSourceAllArgs(t *testing.T) {
	enterpriseAccountId := acc.IamIdentityEnterpriseAccountId
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckIamIdentityEnterpriseTemplates(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMTrustedProfileTemplateDataSourceConfig(enterpriseAccountId, name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "include_history"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "committed"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "profile.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "profile.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "profile.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "profile.0.rules.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "profile.0.identities.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "created_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "last_modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_template.trusted_profile_template", "last_modified_by_id"),
				),
			},
		},
	})
}

func testAccCheckIBMTrustedProfileTemplateDataSourceConfigBasic(enterpriseAccountId string, name string, description string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile_template" "trusted_profile_template" {
			account_id = "%s"
			name = "%s"
			description = "%s"
			profile {
				name = "%s"
			}
		}

		data "ibm_iam_trusted_profile_template" "trusted_profile_template" {
			template_id = ibm_iam_trusted_profile_template.trusted_profile_template.id
		}
	`, enterpriseAccountId, name, description, name)
}

func testAccCheckIBMTrustedProfileTemplateDataSourceConfig(enterpriseAccountId string, name string, description string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile_template" "trusted_profile_template" {
		  account_id = "%s"
		  name = "%s"
		  description = "%s"
		  profile {
			name = "name"
			description = "description"
			rules {
			  name = "name"
			  type = "Profile-SAML"
			  realm_name = "test-realm-101"
			  expiration = 1
			  conditions {
				claim = "claim"
				operator = "EQUALS"
				value = "\"value\""
			  }
			}
			identities {
			  iam_id = "crn-crn:v1:staging:public:iam-identity::a/685e0f537b4548eb8d2c8593881a6b03:::"
			  identifier = "crn:v1:staging:public:iam-identity::a/685e0f537b4548eb8d2c8593881a6b03:::"
			  type = "crn"
			}
		  }
		}

		data "ibm_iam_trusted_profile_template" "trusted_profile_template" {
			template_id = ibm_iam_trusted_profile_template.trusted_profile_template.id
		}
	`, enterpriseAccountId, name, description)
}

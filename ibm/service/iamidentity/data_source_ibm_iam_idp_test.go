// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIamIdpDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acc.TestAccPreCheck(t)
			if acc.IAMIdpID == "" {
				t.Skip("IBM_IAM_IDP_ID is not set; skipping ibm_iam_idp data source test")
			}
		},
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamIdpDataSourceConfigBasic(acc.IAMIdpID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_idp.iam_idp_ds", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_idp.iam_idp_ds", "idp_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_idp.iam_idp_ds", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_idp.iam_idp_ds", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_idp.iam_idp_ds", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_idp.iam_idp_ds", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_idp.iam_idp_ds", "created_at"),
				),
			},
		},
	})
}

// TestAccIBMIamIdpDataSourceFromResource reads an IdP that was just created by
// the ibm_iam_idp resource, to verify end-to-end attribute propagation.
func TestAccIBMIamIdpDataSourceFromResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckIAMIdp(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamIdpDataSourceConfigFromResource(acc.IAMIdpAccountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"data.ibm_iam_idp.iam_idp_lookup",
						"name",
						"ibm_iam_idp.iam_idp_for_ds",
						"name",
					),
					resource.TestCheckResourceAttrPair(
						"data.ibm_iam_idp.iam_idp_lookup",
						"type",
						"ibm_iam_idp.iam_idp_for_ds",
						"type",
					),
					resource.TestCheckResourceAttrSet("data.ibm_iam_idp.iam_idp_lookup", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_idp.iam_idp_lookup", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_idp.iam_idp_lookup", "modified_at"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Helper configs
// ---------------------------------------------------------------------------

func testAccCheckIBMIamIdpDataSourceConfigBasic(idpID string) string {
	return fmt.Sprintf(`
data "ibm_iam_idp" "iam_idp_ds" {
  idp_id = "%s"
}
`, idpID)
}

func testAccCheckIBMIamIdpDataSourceConfigFromResource(accountID string) string {
	return fmt.Sprintf(`
resource "ibm_iam_idp" "iam_idp_for_ds" {
  account_id = "%s"
  name       = "tf-ds-test-idp"
  type       = "saml"
  active     = true

  properties {
    idp {
      entity_id           = "http://www.okta.com/exampleidp"
      redirect_binding_url = "https://example.okta.com/app/example/sso/saml"
    }
  }

  secrets {}
}

data "ibm_iam_idp" "iam_idp_lookup" {
  idp_id = ibm_iam_idp.iam_idp_for_ds.id
}
`, accountID)
}

// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"strconv"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMIamIdpsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckIAMIdp(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamIdpsDataSourceConfigBasic(acc.IAMIdpAccountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_idps.iam_idps_ds", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_idps.iam_idps_ds", "account_id"),
					// The list may be empty if no IDPs exist yet; just verify the attribute is present.
					resource.TestCheckResourceAttrSet("data.ibm_iam_idps.iam_idps_ds", "idps.#"),
				),
			},
		},
	})
}

// TestAccIBMIamIdpsDataSourceAfterCreate verifies the list grows by one after an
// ibm_iam_idp resource is created in the same configuration.
func TestAccIBMIamIdpsDataSourceAfterCreate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckIAMIdp(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamIdpsDataSourceConfigAfterCreate(acc.IAMIdpAccountId),
				Check: resource.ComposeTestCheckFunc(
					// At least one IDP should exist now.
					resource.TestCheckResourceAttr("data.ibm_iam_idps.iam_idps_ds_with_resource", "account_id", acc.IAMIdpAccountId),
					resource.TestCheckResourceAttrSet("data.ibm_iam_idps.iam_idps_ds_with_resource", "idps.#"),
					// The newly created IDP should appear somewhere in the list (order is not guaranteed).
					testAccCheckIBMIamIdpsContainsIdp(
						"data.ibm_iam_idps.iam_idps_ds_with_resource",
						"ibm_iam_idp.iam_idp_for_list",
					),
				),
			},
		},
	})
}

// testAccCheckIBMIamIdpsContainsIdp verifies that the IDP created by resourceName
// appears somewhere in the idps list of the data source datasourceName.
// The list order is not guaranteed by the API, so we search all entries.
func testAccCheckIBMIamIdpsContainsIdp(datasourceName, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		dsState, ok := s.RootModule().Resources[datasourceName]
		if !ok {
			return fmt.Errorf("data source not found: %s", datasourceName)
		}
		rsState, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		wantName := rsState.Primary.Attributes["name"]
		wantID := rsState.Primary.Attributes["idp_id"]

		countStr := dsState.Primary.Attributes["idps.#"]
		count, _ := strconv.Atoi(countStr)
		for i := 0; i < count; i++ {
			if dsState.Primary.Attributes[fmt.Sprintf("idps.%d.name", i)] == wantName &&
				dsState.Primary.Attributes[fmt.Sprintf("idps.%d.idp_id", i)] == wantID {
				return nil
			}
		}
		return fmt.Errorf("IDP %q (id=%s) not found in %s list", wantName, wantID, datasourceName)
	}
}

// ---------------------------------------------------------------------------
// Helper configs
// ---------------------------------------------------------------------------

func testAccCheckIBMIamIdpsDataSourceConfigBasic(accountID string) string {
	return fmt.Sprintf(`
data "ibm_iam_idps" "iam_idps_ds" {
  account_id = "%s"
}
`, accountID)
}

func testAccCheckIBMIamIdpsDataSourceConfigAfterCreate(accountID string) string {
	return fmt.Sprintf(`
resource "ibm_iam_idp" "iam_idp_for_list" {
  account_id = "%s"
  name       = "tf-list-test-idp"
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

data "ibm_iam_idps" "iam_idps_ds_with_resource" {
  account_id = "%s"
  depends_on = [ibm_iam_idp.iam_idp_for_list]
}
`, accountID, accountID)
}

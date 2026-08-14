// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/iamidentity"
)

// ---------------------------------------------------------------------------
// Acceptance tests
// ---------------------------------------------------------------------------

func TestAccIBMIamIdpBasic(t *testing.T) {
	var conf iamidentityv1.Idp
	name := fmt.Sprintf("tf_idp_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_idp_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckIAMIdp(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamIdpDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamIdpConfigBasic(acc.IAMIdpAccountId, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamIdpExists("ibm_iam_idp.iam_idp_instance", conf),
					resource.TestCheckResourceAttr("ibm_iam_idp.iam_idp_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_idp.iam_idp_instance", "type", "saml"),
					resource.TestCheckResourceAttr("ibm_iam_idp.iam_idp_instance", "account_id", acc.IAMIdpAccountId),
					resource.TestCheckResourceAttrSet("ibm_iam_idp.iam_idp_instance", "idp_id"),
					resource.TestCheckResourceAttrSet("ibm_iam_idp.iam_idp_instance", "entity_tag"),
				),
			},
			{
				Config: testAccCheckIBMIamIdpConfigBasic(acc.IAMIdpAccountId, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_idp.iam_idp_instance", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIBMIamIdpAllArgs(t *testing.T) {
	var conf iamidentityv1.Idp
	name := fmt.Sprintf("tf_idp_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_idp_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckIAMIdp(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamIdpDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamIdpConfigAllArgs(acc.IAMIdpAccountId, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamIdpExists("ibm_iam_idp.iam_idp_full", conf),
					resource.TestCheckResourceAttr("ibm_iam_idp.iam_idp_full", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_idp.iam_idp_full", "type", "saml"),
					resource.TestCheckResourceAttr("ibm_iam_idp.iam_idp_full", "active", "true"),
					resource.TestCheckResourceAttr("ibm_iam_idp.iam_idp_full", "properties.0.idp.0.entity_id", "https://idp.example.com/saml/metadata"),
					resource.TestCheckResourceAttr("ibm_iam_idp.iam_idp_full", "properties.0.idp.0.redirect_binding_url", "https://idp.example.com/saml/sso"),
					resource.TestCheckResourceAttr("ibm_iam_idp.iam_idp_full", "properties.0.sp.0.want_assertion_signed", "true"),
					resource.TestCheckResourceAttrSet("ibm_iam_idp.iam_idp_full", "idp_id"),
					resource.TestCheckResourceAttrSet("ibm_iam_idp.iam_idp_full", "entity_tag"),
					resource.TestCheckResourceAttrSet("ibm_iam_idp.iam_idp_full", "created_at"),
				),
			},
			{
				Config: testAccCheckIBMIamIdpConfigAllArgs(acc.IAMIdpAccountId, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_idp.iam_idp_full", "name", nameUpdate),
				),
			},
			{
				ResourceName:      "ibm_iam_idp.iam_idp_full",
				ImportState:       true,
				ImportStateVerify: true,
				// properties/secrets are not returned by GET; skip verifying them on import
				ImportStateVerifyIgnore: []string{"properties", "secrets"},
			},
		},
	})
}

func TestAccIBMIamIdpWithShareScope(t *testing.T) {
	var conf iamidentityv1.Idp
	name := fmt.Sprintf("tf_idp_shared_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckIAMIdpAccountSetting(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamIdpDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamIdpConfigWithShareScope(acc.IAMIdpAccountId, name, acc.IAMIdpConsumerAccountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamIdpExists("ibm_iam_idp.iam_idp_shared", conf),
					resource.TestCheckResourceAttr("ibm_iam_idp.iam_idp_shared", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_idp.iam_idp_shared", "share_scope.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_idp.iam_idp_shared", "share_scope.0.type", "account"),
					resource.TestCheckResourceAttr("ibm_iam_idp.iam_idp_shared", "share_scope.0.id", acc.IAMIdpConsumerAccountId),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Acceptance-test helper configs
// ---------------------------------------------------------------------------

func testAccCheckIBMIamIdpConfigBasic(accountID, name string) string {
	return fmt.Sprintf(`
resource "ibm_iam_idp" "iam_idp_instance" {
  account_id = "%s"
  name       = "%s"
  type       = "saml"

  properties {
    idp {
      entity_id           = "http://www.okta.com/exampleidp"
      redirect_binding_url = "https://example.okta.com/app/example/sso/saml"
    }
  }

  secrets {}
}
`, accountID, name)
}

func testAccCheckIBMIamIdpConfigAllArgs(accountID, name string) string {
	return fmt.Sprintf(`
resource "ibm_iam_idp" "iam_idp_full" {
  account_id = "%s"
  name       = "%s"
  type       = "saml"
  active     = true

  properties {
    idp {
      entity_id           = "https://idp.example.com/saml/metadata"
      redirect_binding_url = "https://idp.example.com/saml/sso"
      want_request_signed  = true
    }
    sp {
      want_assertion_signed = true
      want_response_signed  = true
      encrypt_response      = false
    }
  }

  secrets {}
}
`, accountID, name)
}

func testAccCheckIBMIamIdpConfigWithShareScope(accountID, name, consumerAccountID string) string {
	return fmt.Sprintf(`
resource "ibm_iam_idp" "iam_idp_shared" {
  account_id = "%s"
  name       = "%s"
  type       = "saml"

  properties {
    idp {
      entity_id           = "http://www.okta.com/exampleidp"
      redirect_binding_url = "https://example.okta.com/app/example/sso/saml"
    }
  }

  secrets {}

  share_scope {
    id   = "%s"
    type = "account"
  }
}
`, accountID, name, consumerAccountID)
}

// ---------------------------------------------------------------------------
// Acceptance-test check functions
// ---------------------------------------------------------------------------

func testAccCheckIBMIamIdpExists(n string, obj iamidentityv1.Idp) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		getOpts := iamIdentityClient.NewGetIdpOptions(rs.Primary.ID)
		idp, _, err := iamIdentityClient.GetIdp(getOpts)
		if err != nil {
			return err
		}

		obj = *idp
		return nil
	}
}

func testAccCheckIBMIamIdpDestroy(s *terraform.State) error {
	iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_idp" {
			continue
		}

		getOpts := iamIdentityClient.NewGetIdpOptions(rs.Primary.ID)
		_, response, err := iamIdentityClient.GetIdp(getOpts)
		if err == nil {
			return fmt.Errorf("ibm_iam_idp still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for ibm_iam_idp (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

// ---------------------------------------------------------------------------
// Unit tests for helper functions
// ---------------------------------------------------------------------------

func TestIamIdpFlattenShareScope(t *testing.T) {
	scopes := []iamidentityv1.ShareScope{
		{
			ID:   core.StringPtr("account-123"),
			Type: core.StringPtr("account"),
		},
		{
			ID:   core.StringPtr("enterprise-456"),
			Type: core.StringPtr("enterprise"),
		},
	}

	result := iamidentity.FlattenShareScopeForTest(scopes)
	assert.Len(t, result, 2)
	assert.Equal(t, "account-123", result[0]["id"])
	assert.Equal(t, "account", result[0]["type"])
	assert.Equal(t, "enterprise-456", result[1]["id"])
	assert.Equal(t, "enterprise", result[1]["type"])
}

func TestIamIdpFlattenShareScopeEmpty(t *testing.T) {
	result := iamidentity.FlattenShareScopeForTest([]iamidentityv1.ShareScope{})
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestIamIdpExpandShareScope(t *testing.T) {
	input := []interface{}{
		map[string]interface{}{
			"id":   "account-abc",
			"type": "account",
		},
	}

	result := iamidentity.ExpandShareScopeForTest(input)
	assert.Len(t, result, 1)
	assert.Equal(t, "account-abc", *result[0].ID)
	assert.Equal(t, "account", *result[0].Type)
}

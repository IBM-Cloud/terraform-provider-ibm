// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/IBM-Cloud/bluemix-go/helpers"
	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"gotest.tools/assert"
)

func TestAccIBMAppIDTokenConfigDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppIDTokenConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_token_config.test_config", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("data.ibm_appid_token_config.test_config", "access_token_expires_in", "7200"),
					resource.TestCheckResourceAttr("data.ibm_appid_token_config.test_config", "anonymous_access_enabled", "false"),
					resource.TestCheckResourceAttr("data.ibm_appid_token_config.test_config", "anonymous_token_expires_in", "7200"),
					resource.TestCheckResourceAttr("data.ibm_appid_token_config.test_config", "refresh_token_enabled", "true"),
					resource.TestCheckResourceAttr("data.ibm_appid_token_config.test_config", "refresh_token_expires_in", "7200"),
					resource.TestCheckResourceAttr("data.ibm_appid_token_config.test_config", "access_token_claim.#", "2"),
					resource.TestCheckResourceAttr("data.ibm_appid_token_config.test_config", "id_token_claim.#", "0"),
					// the order here is deterministic: https://github.com/hashicorp/terraform-plugin-sdk/blob/main/helper/schema/set.go#L268
					resource.TestCheckResourceAttr("data.ibm_appid_token_config.test_config", "access_token_claim.0.destination_claim", "employeeId"),
					resource.TestCheckResourceAttr("data.ibm_appid_token_config.test_config", "access_token_claim.0.source", "appid_custom"),
					resource.TestCheckResourceAttr("data.ibm_appid_token_config.test_config", "access_token_claim.0.source_claim", "employeeId"),
					resource.TestCheckResourceAttr("data.ibm_appid_token_config.test_config", "access_token_claim.1.destination_claim", "groupIds"),
					resource.TestCheckResourceAttr("data.ibm_appid_token_config.test_config", "access_token_claim.1.source", "roles"),
					resource.TestCheckResourceAttr("data.ibm_appid_token_config.test_config", "access_token_claim.1.source_claim", ""),
				),
			},
		},
	})
}

func TestFlattenTokenClaims(t *testing.T) {
	testcases := []struct {
		claims   []appid.TokenClaimMapping
		expected []interface{}
	}{
		{
			claims: []appid.TokenClaimMapping{
				{Source: helpers.String("appid_custom"), SourceClaim: helpers.String("sClaim"), DestinationClaim: helpers.String("dClaim")},
				{Source: helpers.String("appid_custom"), DestinationClaim: helpers.String("dClaim")},
			},
			expected: []interface{}{
				map[string]interface{}{"source": "appid_custom", "source_claim": "sClaim", "destination_claim": "dClaim"},
				map[string]interface{}{"source": "appid_custom", "destination_claim": "dClaim"},
			},
		},
	}

	for _, c := range testcases {
		actual := flattenTokenClaims(c.claims)
		assert.DeepEqual(t, actual, c.expected)
	}
}
func flattenTokenClaims(c []appid.TokenClaimMapping) []interface{} {
	var s []interface{}

	for _, v := range c {
		claim := map[string]interface{}{
			"source": *v.Source,
		}

		if v.SourceClaim != nil {
			claim["source_claim"] = *v.SourceClaim
		}

		if v.DestinationClaim != nil {
			claim["destination_claim"] = *v.DestinationClaim
		}

		s = append(s, claim)
	}

	return s
}
func testAccCheckIBMAppIDTokenConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_token_config" "test_config" {
			tenant_id = "%s"
			access_token_expires_in = 7200    
			anonymous_access_enabled = false
			anonymous_token_expires_in = 7200
			refresh_token_enabled = true
			refresh_token_expires_in = 7200
			
			access_token_claim {
				source = "roles"
				destination_claim = "groupIds"
			}

			access_token_claim {
				source = "appid_custom"
				source_claim = "employeeId"
				destination_claim = "employeeId"
			}
		}

		data "ibm_appid_token_config" "test_config" {
			tenant_id = ibm_appid_token_config.test_config.tenant_id
			
			depends_on = [
				ibm_appid_token_config.test_config
			]
		}
	`, tenantID)
}

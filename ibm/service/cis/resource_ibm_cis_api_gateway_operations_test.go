// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccIBMCisApiGatewayOperations_Basic verifies that two operations are created in
// bulk and that each receives a server-assigned operation_id.
func TestAccIBMCisApiGatewayOperations_Basic(t *testing.T) {
	name := "ibm_cis_api_gateway_operations.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisApiGatewayOperationsBasic("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "operations.#", "2"),
					resource.TestCheckResourceAttrSet(name, "operations.0.operation_id"),
					resource.TestCheckResourceAttrSet(name, "operations.1.operation_id"),
					resource.TestCheckResourceAttr(name, "operations.0.method", "POST"),
					resource.TestCheckResourceAttr(name, "operations.1.method", "GET"),
				),
			},
		},
	})
}

// TestAccIBMCisApiGatewayOperations_Update verifies that adding a third operation in
// an update step creates it while the first two are preserved.
func TestAccIBMCisApiGatewayOperations_Update(t *testing.T) {
	name := "ibm_cis_api_gateway_operations.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisApiGatewayOperationsBasic("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "operations.#", "2"),
					resource.TestCheckResourceAttrSet(name, "operations.0.operation_id"),
					resource.TestCheckResourceAttrSet(name, "operations.1.operation_id"),
				),
			},
			{
				Config: testAccCheckCisApiGatewayOperationsUpdated("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "operations.#", "3"),
					resource.TestCheckResourceAttrSet(name, "operations.0.operation_id"),
					resource.TestCheckResourceAttrSet(name, "operations.1.operation_id"),
					resource.TestCheckResourceAttrSet(name, "operations.2.operation_id"),
				),
			},
		},
	})
}

// TestAccIBMCisApiGatewayOperations_Import verifies that the resource state can be
// imported using the zoneID:CRN composite ID.
//
// ImportStateVerify is intentionally disabled for this resource: the resource
// manages a user-defined *subset* of zone operations.  On import, Read returns
// all operations currently registered in the zone (there may be others created
// outside Terraform), so the imported state will legitimately differ from the
// pre-import state.  After importing, the operator re-applies their config to
// reconcile the managed subset.
func TestAccIBMCisApiGatewayOperations_Import(t *testing.T) {
	name := "ibm_cis_api_gateway_operations.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisApiGatewayOperationsBasic("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "operations.#", "2"),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       false,
				ImportStateVerifyIgnore: []string{"operations"},
			},
		},
	})
}

// ── HCL config helpers ────────────────────────────────────────────────────────

func testAccCheckCisApiGatewayOperationsBasic(id, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
resource "ibm_cis_api_gateway_operations" "%[1]s" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id

  operations {
    method   = "POST"
    host     = "%[2]s"
    endpoint = "/api/v1/chat/completions"
  }

  operations {
    method   = "GET"
    host     = "%[2]s"
    endpoint = "/api/v1/models"
  }
}
`, id, cisDomainStatic)
}

func testAccCheckCisApiGatewayOperationsUpdated(id, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
resource "ibm_cis_api_gateway_operations" "%[1]s" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id

  operations {
    method   = "POST"
    host     = "%[2]s"
    endpoint = "/api/v1/chat/completions"
  }

  operations {
    method   = "GET"
    host     = "%[2]s"
    endpoint = "/api/v1/models"
  }

  operations {
    method   = "DELETE"
    host     = "%[2]s"
    endpoint = "/api/v1/sessions"
  }
}
`, id, cisDomainStatic)
}

// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccScopeDataSourceBasic(t *testing.T) {
	scopeName := fmt.Sprintf("tf_scope_name_%d", acctest.RandIntRange(10, 100))
	scopeDescription := fmt.Sprintf("tf_scope_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccScopeDataSourceConfigBasic(acc.SccInstanceID, acc.SccAccountID, scopeName, scopeDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_scope.scc_scope_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_scope.scc_scope_instance", "scope_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_scope.scc_scope_instance", "name"),
				),
			},
		},
	})
}

func testAccCheckIbmSccScopeDataSourceConfigBasic(instanceID string, accountID string, scopeName string, scopeDescription string) string {
	return fmt.Sprintf(`
  resource "ibm_scc_scope" "scc_account_scope" {
    description = "%s"
    environment = "ibm-cloud"
    instance_id = "%s"
    name        = "%s"
    properties  = {
      scope_id    = "%s"
      scope_type  = "account"
    }
  }

  data "ibm_scc_scope" "scc_scope_instance" {
    instance_id = "%s"
    scope_id = ibm_scc_scope.scc_account_scope.scope_id 
  }
  `, scopeDescription, instanceID, scopeName, accountID, instanceID)
}

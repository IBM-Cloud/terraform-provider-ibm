// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package watsonx_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// The COS to KMS authorization policy must already exist for the COS instance.
func TestAccIBMWatsonxStorageDelegationBasic(t *testing.T) {
	env := requireWatsonxEnv(t, true, false)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccWatsonxStorageDelegationConfig(env, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_watsonx_storage_delegation.test", "id", env.COSGUID),
					resource.TestCheckResourceAttr("ibm_watsonx_storage_delegation.test", "kms_key_crn", env.KMSKeyCRN),
					resource.TestCheckResourceAttr("ibm_watsonx_storage_delegation.test", "catalogs", "true"),
					resource.TestCheckResourceAttr("ibm_watsonx_storage_delegation.test", "projects", "true"),
				),
			},
			{
				Config: testAccWatsonxStorageDelegationConfig(env, false),
				Check:  resource.TestCheckResourceAttr("ibm_watsonx_storage_delegation.test", "catalogs", "false"),
			},
			{
				ResourceName:      "ibm_watsonx_storage_delegation.test",
				ImportState:       true,
				ImportStateId:     env.Region + "/" + env.COSGUID,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccWatsonxStorageDelegationConfig(env watsonxEnv, catalogs bool) string {
	return fmt.Sprintf(`
		resource "ibm_watsonx_storage_delegation" "test" {
			region          = "%s"
			cos_instance_id = "%s"
			kms_key_crn     = "%s"
			catalogs        = %t
		}
	`, env.Region, env.COSGUID, env.KMSKeyCRN, catalogs)
}

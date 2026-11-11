// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package watsonx_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMWatsonxProjectBasic(t *testing.T) {
	env := requireWatsonxEnv(t, false, true)
	name := fmt.Sprintf("tf-wx-project-%d", acctest.RandIntRange(10, 100000))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccWatsonxProjectConfig(env, name, "created by acceptance test"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_watsonx_project.test", "id"),
					resource.TestCheckResourceAttr("ibm_watsonx_project.test", "name", name),
					resource.TestMatchResourceAttr("ibm_watsonx_project.test", "location", regexpProjectLocation),
					resource.TestCheckResourceAttrSet("ibm_watsonx_project.test", "project_url"),
				),
			},
			{
				// description and name are updated in place.
				Config: testAccWatsonxProjectConfig(env, name+"-upd", "updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_watsonx_project.test", "name", name+"-upd"),
					resource.TestCheckResourceAttr("ibm_watsonx_project.test", "description", "updated"),
				),
			},
			{
				ResourceName:      "ibm_watsonx_project.test",
				ImportState:       true,
				ImportStateIdFunc: testAccWatsonxImportID(env.Region, "ibm_watsonx_project.test"),
				ImportStateVerify: true,
				// Create-only inputs are not returned by the API.
				ImportStateVerifyIgnore: []string{"tags", "storage_delegated", "mark_as_sensitive", "cos_crn", "endpoint"},
			},
		},
	})
}

func testAccWatsonxProjectConfig(env watsonxEnv, name, description string) string {
	return fmt.Sprintf(`
		resource "ibm_watsonx_project" "test" {
			region          = "%s"
			name            = "%s"
			description     = "%s"
			tags            = ["acc-test"]
			cos_instance_id = "%s"
			cos_crn         = "%s"
			runtime_name    = "%s"
			runtime_guid    = "%s"
			runtime_crn     = "%s"
		}
	`, env.Region, name, description, env.COSGUID, env.COSCRN, env.RuntimeName, env.RuntimeGUID, env.RuntimeCRN)
}

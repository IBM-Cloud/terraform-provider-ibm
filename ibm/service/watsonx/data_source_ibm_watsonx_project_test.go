// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package watsonx_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMWatsonxProjectDataSourceBasic(t *testing.T) {
	env := requireWatsonxEnv(t, false, true)
	name := fmt.Sprintf("tf-wx-ds-%d", acctest.RandIntRange(10, 100000))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccWatsonxProjectConfig(env, name, "data source test") + fmt.Sprintf(`
					data "ibm_watsonx_project" "test" {
						region     = "%s"
						project_id = ibm_watsonx_project.test.id
					}
				`, env.Region),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.ibm_watsonx_project.test", "id", "ibm_watsonx_project.test", "id"),
					resource.TestCheckResourceAttr("data.ibm_watsonx_project.test", "name", name),
					resource.TestCheckResourceAttr("data.ibm_watsonx_project.test", "cos_instance_id", env.COSGUID),
					resource.TestCheckResourceAttr("data.ibm_watsonx_project.test", "tags.#", "1"),
					resource.TestCheckResourceAttrSet("data.ibm_watsonx_project.test", "project_json"),
				),
			},
		},
	})
}

func TestAccIBMWatsonxProjectDataSourceNotFound(t *testing.T) {
	env := requireWatsonxEnv(t, false, false)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					data "ibm_watsonx_project" "test" {
						region     = "%s"
						project_id = "00000000-0000-0000-0000-000000000000"
					}
				`, env.Region),
				ExpectError: regexp.MustCompile("Project Not Found|HTTP 404"),
			},
		},
	})
}

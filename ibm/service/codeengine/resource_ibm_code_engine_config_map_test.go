// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
)

func TestAccIbmCodeEngineConfigMapBasic(t *testing.T) {
	var conf codeenginev2.ConfigMap
	name := fmt.Sprintf("tf-config-map-basic-%d", acctest.RandIntRange(10, 1000))
	data := `{ "key" = "inner" }`
	nameUpdate := fmt.Sprintf("tf-config-map-basic-update-%d", acctest.RandIntRange(10, 1000))

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineConfigMapDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineConfigMapConfigBasic(projectID, name, data),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineConfigMapExists("ibm_code_engine_config_map.code_engine_config_map_instance", conf),
					resource.TestCheckResourceAttrSet("ibm_code_engine_config_map.code_engine_config_map_instance", "config_map_id"),
					resource.TestCheckResourceAttr("ibm_code_engine_config_map.code_engine_config_map_instance", "data.key", "inner"),
					resource.TestCheckResourceAttr("ibm_code_engine_config_map.code_engine_config_map_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_config_map.code_engine_config_map_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_code_engine_config_map.code_engine_config_map_instance", "resource_type", "config_map_v2"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineConfigMapConfigBasic(projectID, nameUpdate, data),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_code_engine_config_map.code_engine_config_map_instance", "config_map_id"),
					resource.TestCheckResourceAttr("ibm_code_engine_config_map.code_engine_config_map_instance", "data.key", "inner"),
					resource.TestCheckResourceAttr("ibm_code_engine_config_map.code_engine_config_map_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_config_map.code_engine_config_map_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_config_map.code_engine_config_map_instance", "resource_type", "config_map_v2"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_code_engine_config_map.code_engine_config_map_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmCodeEngineConfigMapConfigBasic(projectID string, name string, data string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_config_map" "code_engine_config_map_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			name = "%s"
			data = %s
		}
	`, projectID, name, data)
}

func testAccCheckIbmCodeEngineConfigMapExists(n string, obj codeenginev2.ConfigMap) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
		if err != nil {
			return err
		}

		getConfigMapOptions := &codeenginev2.GetConfigMapOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getConfigMapOptions.SetProjectID(parts[0])
		getConfigMapOptions.SetName(parts[1])

		configMap, _, err := codeEngineClient.GetConfigMap(getConfigMapOptions)
		if err != nil {
			return err
		}

		obj = *configMap
		return nil
	}
}

func testAccCheckIbmCodeEngineConfigMapDestroy(s *terraform.State) error {
	codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_code_engine_config_map" {
			continue
		}

		getConfigMapOptions := &codeenginev2.GetConfigMapOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getConfigMapOptions.SetProjectID(parts[0])
		getConfigMapOptions.SetName(parts[1])

		// Try to find the key
		_, response, err := codeEngineClient.GetConfigMap(getConfigMapOptions)

		if err == nil {
			return fmt.Errorf("code_engine_config_map still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for code_engine_config_map (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

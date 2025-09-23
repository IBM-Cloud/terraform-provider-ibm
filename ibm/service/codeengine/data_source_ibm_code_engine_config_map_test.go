// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmCodeEngineConfigMapDataSourceBasic(t *testing.T) {
	configMapName := fmt.Sprintf("tf-data-config-map-basic-%d", acctest.RandIntRange(10, 1000))
	configMapData := `{ "key" = "inner" }`

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineConfigMapDataSourceConfigBasic(projectID, configMapName, configMapData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_config_map.code_engine_config_map_instance", "config_map_id"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_config_map.code_engine_config_map_instance", "data.key", "inner"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_config_map.code_engine_config_map_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_config_map.code_engine_config_map_instance", "name", configMapName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_config_map.code_engine_config_map_instance", "resource_type", "config_map_v2"),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineConfigMapDataSourceConfigBasic(projectID string, configMapName string, configMapData string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_config_map" "code_engine_config_map_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			name = "%s"
			data = %s
		}

		data "ibm_code_engine_config_map" "code_engine_config_map_instance" {
			project_id = ibm_code_engine_config_map.code_engine_config_map_instance.project_id
			name = ibm_code_engine_config_map.code_engine_config_map_instance.name
		}
	`, projectID, configMapName, configMapData)
}

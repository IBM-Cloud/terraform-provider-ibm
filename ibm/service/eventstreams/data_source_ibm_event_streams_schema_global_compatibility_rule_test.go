// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventstreams_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMEventStreamsSchemaGlobalCompatibilityRuleDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEventStreamsSchemaGlobalCompatibilityRuleDataSourceConfig(getTestInstanceName(mzrKey)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMEventStreamsSchemaGlobalCompatibilityRuleProperties("data.ibm_event_streams_schema_global_rule.es_globalrule", ""),
				),
			},
		},
	})
}

func testAccCheckIBMEventStreamsSchemaGlobalCompatibilityRuleDataSourceConfig(instanceName string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default=true
	}
	data "ibm_resource_instance" "es_instance" {
		resource_group_id = data.ibm_resource_group.group.id
		name              = "%s"
	}
	data "ibm_event_streams_schema_global_rule" "es_globalrule" {
		resource_instance_id = data.ibm_resource_instance.es_instance.id
	}`, instanceName)
}

// check properties of the global compatibility rule data source or resource object
func testAccCheckIBMEventStreamsSchemaGlobalCompatibilityRuleProperties(name string, expectConfig string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}
		gcrID := rs.Primary.ID
		if gcrID == "" {
			return fmt.Errorf("[ERROR] Global compatibility rule ID is not set")
		}
		if !strings.HasSuffix(gcrID, ":schema-global-compatibility-rule:") {
			return fmt.Errorf("[ERROR] Global compatibility rule ID %s not expected CRN", gcrID)
		}
		config := rs.Primary.Attributes["config"]
		if config == "" {
			return fmt.Errorf("[ERROR] Global compatibility config is not defined")
		}
		if expectConfig != "" && config != expectConfig {
			return fmt.Errorf("[ERROR] Global compatibility config is %s, expected %s", config, expectConfig)
		}
		return nil
	}
}

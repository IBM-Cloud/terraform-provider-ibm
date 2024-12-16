// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventstreams_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/eventstreams-go-sdk/pkg/schemaregistryv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMEventStreamsSchemaGlobalCompatibilityRuleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEventStreamsGlobalCompatibilityRuleResetToNoneInInstance(),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEventStreamsSchemaGlobalCompatibilityRuleResourceConfig(getTestInstanceName(mzrKey), "FORWARD"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEventStreamsSchemaGlobalCompatibilityRuleProperties("ibm_event_streams_schema_global_rule.es_globalrule", "FORWARD"),
				),
			},
			{
				ResourceName:      "ibm_event_streams_schema_global_rule.es_globalrule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMEventStreamsSchemaGlobalCompatibilityRuleResourceConfig(instanceName string, rule string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default=true
	}
	data "ibm_resource_instance" "es_instance" {
		resource_group_id = data.ibm_resource_group.group.id
		name              = "%s"
	}
	resource "ibm_event_streams_schema_global_rule" "es_globalrule" {
		resource_instance_id = data.ibm_resource_instance.es_instance.id
		config = "%s"
	}`, instanceName, rule)
}

func testAccCheckIBMEventStreamsGlobalCompatibilityRuleResetToNoneInInstance() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		schemaregistryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ESschemaRegistrySession()
		if err != nil {
			return err
		}
		getOpts := &schemaregistryv1.GetGlobalRuleOptions{}
		getOpts.SetRule("COMPATIBILITY")
		rule, _, err := schemaregistryClient.GetGlobalRule(getOpts)
		if err != nil {
			return err
		}
		if rule.Config == nil || *rule.Config != "NONE" {
			return fmt.Errorf("[ERROR] Expected global compatibility rule reset to NONE after deletion")
		}
		return nil
	}
}

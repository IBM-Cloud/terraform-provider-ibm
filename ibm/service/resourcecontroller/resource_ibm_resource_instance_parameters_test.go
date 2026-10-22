// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package resourcecontroller_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Set IBM_RESOURCE_INSTANCE_PARAMETERS_INSTANCE_ID to an instance that accepts the
// JSON in IBM_RESOURCE_INSTANCE_PARAMETERS (for example an SCC Workload Protection
// instance with {"enable_cspm":false}). The test is skipped when either is unset.
func TestAccIBMResourceInstanceParametersBasic(t *testing.T) {
	instanceID := os.Getenv("IBM_RESOURCE_INSTANCE_PARAMETERS_INSTANCE_ID")
	params := os.Getenv("IBM_RESOURCE_INSTANCE_PARAMETERS")
	if instanceID == "" || params == "" {
		t.Skip("IBM_RESOURCE_INSTANCE_PARAMETERS_INSTANCE_ID and IBM_RESOURCE_INSTANCE_PARAMETERS must be set")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInstanceParametersConfig(instanceID, params),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_resource_instance_parameters.test", "id", instanceID),
					resource.TestCheckResourceAttr("ibm_resource_instance_parameters.test", "parameters", params),
					resource.TestCheckResourceAttrSet("ibm_resource_instance_parameters.test", "state"),
				),
			},
			{
				// A second plan with the same config must be empty: server-side
				// fields outside the managed keys never show as drift.
				Config:   testAccResourceInstanceParametersConfig(instanceID, params),
				PlanOnly: true,
			},
		},
	})
}

func TestAccIBMResourceInstanceParametersInvalidJSON(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceInstanceParametersConfig("00000000-0000-0000-0000-000000000000", "not-json"),
				ExpectError: regexp.MustCompile("parameters must be a JSON object"),
			},
		},
	})
}

func testAccResourceInstanceParametersConfig(instanceID, params string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance_parameters" "test" {
			resource_instance_id = "%s"
			parameters           = %q
		}
	`, instanceID, params)
}

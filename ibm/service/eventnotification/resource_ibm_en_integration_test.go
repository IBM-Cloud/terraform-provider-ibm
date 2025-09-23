// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMEnIntegrationAllArgs(t *testing.T) {
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	integrationid := fmt.Sprintf("tf_integrationid_%d", acctest.RandIntRange(10, 100))
	newintegrationid := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnIntegrationConfig(instanceName, integrationid),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_integration.en_integration_resource_1", "integration_id", integrationid),
				),
			},
			{
				Config: testAccCheckIBMEnIntegrationConfig(instanceName, newintegrationid),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_integration.en_integration_resource_1", "integration_id", integrationid),
				),
			},
			{
				ResourceName:      "ibm_en_integration.en_integration_resource_1",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMEnIntegrationConfig(instanceName, integrationid string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_integration_resource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_integration" "en_integration_resource_1" {
		instance_guid = ibm_resource_instance.en_destination_resource.guid
		integration_id       = "%s"
		type = "kms"
		metadata {
			endpoint = "https://us-south.kms.cloud.ibm.com"
			crn = "crn:v1:bluemix:public:kms:us-south:a/tyyeeuuii2637390003hehhhhi:fgsyysbnjiios::"
			root_key_id = "gyyebvhy-34673783-nshuwubw"
		}

	}
	`, instanceName, integrationid)
}

func testAccCheckIBMEnIntegrationDestroy(s *terraform.State) error {

	return nil
}

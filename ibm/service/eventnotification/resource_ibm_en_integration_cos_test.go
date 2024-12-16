// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	en "github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMEnIntegrationCOSAllArgs(t *testing.T) {
	var metadata en.IntegrationGetResponse
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	// integrationid := fmt.Sprintf("tf_integrationid_%d", acctest.RandIntRange(10, 100))
	// newintegrationid := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnIntegrationCOSConfig(instanceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnIntegrationCOSExists("ibm_en_destination_webhook.en_destination_resource_1", metadata),
					resource.TestCheckResourceAttr("ibm_en_integration_cos.en_integration_resource_1", "type", "collect_failed_events"),
				),
			},
			{
				Config: testAccCheckIBMEnIntegrationCOSConfig(instanceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_integration_cos.en_integration_resource_1", "type", "collect_failed_events"),
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

func testAccCheckIBMEnIntegrationCOSConfig(instanceName string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_integration_resource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_integration" "en_integration_resource_1" {
		instance_guid = ibm_resource_instance.en_destination_resource.guid
		type = "collect_failed_events"
		metadata {
			endpoint = "https://s3.us-west.cloud-object-storage.test.appdomain.cloud",
			crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/xxxxxxx6db359a81a1dde8f44bxxxxxx:xxxxxxxx-1d48-xxxx-xxxx-xxxxxxxxxxxx:bucket:cloud-object-storage"
			bucket_name = "cloud-object-storage"
		}

	}
	`, instanceName)
}

func testAccCheckIBMEnIntegrationCOSExists(n string, obj en.IntegrationGetResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		enClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EventNotificationsApiV1()
		if err != nil {
			return err
		}

		options := &en.GetIntegrationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		options.SetInstanceID(parts[0])
		options.SetID(parts[1])

		result, _, err := enClient.GetIntegration(options)
		if err != nil {
			return err
		}

		obj = *result
		return nil
	}
}

func testAccCheckIBMEnIntegrationCOSDestroy(s *terraform.State) error {

	return nil
}

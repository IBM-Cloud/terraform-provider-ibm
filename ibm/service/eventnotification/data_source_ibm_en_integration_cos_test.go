// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMEnIntegrationCOSDataSourceBasic(t *testing.T) {
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	// integrationid := fmt.Sprintf("tf_integrationid_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnIntegrationCOSDataSourceConfigBasic(instanceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_integration_cos.en_integration_data_1", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_integration_cos.en_integration_data_1", "instance_guid"),
					// resource.TestCheckResourceAttrSet("data.ibm_en_integration_cos.en_integration_data_1", "integrationid"),
					resource.TestCheckResourceAttrSet("data.ibm_en_integration_cos.en_integration_data_1", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIBMEnIntegrationCOSDataSourceConfigBasic(instanceName string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_integration_datasource2" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_integration_cos" "en_integration_resource_1" {
		instance_guid = ibm_resource_instance.en_integration_datasource2.guid
		type = "collect_failed_events"
		metadata {
			endpoint = "https://s3.us-west.cloud-object-storage.test.appdomain.cloud",
			crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/xxxxxxx6db359a81a1dde8f44bxxxxxx:xxxxxxxx-1d48-xxxx-xxxx-xxxxxxxxxxxx:bucket:cloud-object-storage"
			bucket_name = "cloud-object-storage"
		}
	}

		data "ibm_en_integration_cos" "en_integration_data_1" {
			instance_guid = ibm_resource_instance.en_integration_datasource2.guid
			integration_id = ibm_en_integration.en_integration_resource_1.integration_id
		}
	`, instanceName)
}

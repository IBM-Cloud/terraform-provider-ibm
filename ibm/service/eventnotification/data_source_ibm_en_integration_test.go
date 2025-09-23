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

func TestAccIBMEnIntegrationDataSourceBasic(t *testing.T) {
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	integrationid := fmt.Sprintf("tf_integrationid_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnIntegrationDataSourceConfigBasic(instanceName, integrationid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_integration.en_integration_data_1", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_integration.en_integration_data_1", "instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_en_integration.en_integration_data_1", "integrationid"),
					resource.TestCheckResourceAttrSet("data.ibm_en_integration.en_integration_data_1", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIBMEnIntegrationDataSourceConfigBasic(instanceName, integrationid string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_integration_datasource2" {
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

		data "ibm_en_integration" "en_integration_data_1" {
			instance_guid = ibm_resource_instance.en_integration_datasource2.guid
			integration_id = ibm_en_integration.en_integration_resource_1.integration_id
		}
	`, instanceName, integrationid)
}

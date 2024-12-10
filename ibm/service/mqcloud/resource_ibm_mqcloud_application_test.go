// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package mqcloud_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/mqcloud-go-sdk/mqcloudv1"
)

func TestAccIbmMqcloudApplicationBasic(t *testing.T) {
	t.Parallel()
	var conf mqcloudv1.ApplicationDetails
	serviceInstanceGuid := acc.MqcloudDeploymentID
	name := "appbasic"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckMqcloud(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmMqcloudApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudApplicationConfigBasic(serviceInstanceGuid, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmMqcloudApplicationExists("ibm_mqcloud_application.mqcloud_application_instance", conf),
					resource.TestCheckResourceAttr("ibm_mqcloud_application.mqcloud_application_instance", "service_instance_guid", serviceInstanceGuid),
					resource.TestCheckResourceAttr("ibm_mqcloud_application.mqcloud_application_instance", "name", name),
				),
			},
			{
				ResourceName:      "ibm_mqcloud_application.mqcloud_application_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmMqcloudApplicationConfigBasic(serviceInstanceGuid string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_mqcloud_application" "mqcloud_application_instance" {
			service_instance_guid = "%s"
			name = "%s"
		}
	`, serviceInstanceGuid, name)
}

func testAccCheckIbmMqcloudApplicationExists(n string, obj mqcloudv1.ApplicationDetails) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		mqcloudClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MqcloudV1()
		if err != nil {
			return err
		}

		getApplicationOptions := &mqcloudv1.GetApplicationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getApplicationOptions.SetServiceInstanceGuid(parts[0])
		getApplicationOptions.SetApplicationID(parts[1])

		applicationDetails, _, err := mqcloudClient.GetApplication(getApplicationOptions)
		if err != nil {
			return err
		}

		obj = *applicationDetails
		return nil
	}
}

func testAccCheckIbmMqcloudApplicationDestroy(s *terraform.State) error {
	mqcloudClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MqcloudV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_mqcloud_application" {
			continue
		}

		getApplicationOptions := &mqcloudv1.GetApplicationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getApplicationOptions.SetServiceInstanceGuid(parts[0])
		getApplicationOptions.SetApplicationID(parts[1])

		// Try to find the key
		_, response, err := mqcloudClient.GetApplication(getApplicationOptions)

		if err == nil {
			return fmt.Errorf("mqcloud_application still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for mqcloud_application (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

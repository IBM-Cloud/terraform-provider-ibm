// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
)

func TestAccIbmSmEnRegistrationBasic(t *testing.T) {
	var conf secretsmanagerv2.NotificationsRegistration

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmEnRegistrationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmEnRegistrationConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmEnRegistrationExists("ibm_sm_en_registration.sm_en_registration", conf),
				),
			},
		},
	})
}

func testAccCheckIbmSmEnRegistrationConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_sm_en_registration" "sm_en_registration"{
  			instance_id   = "%s"
  			region        = "%s"
  			event_notifications_instance_crn = "%s"
  			event_notifications_source_description = "Terraform data source test."
  			event_notifications_source_name = "My Secrets Manager Terraform Test"
}

	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerENInstanceCrn)
}

func testAccCheckIbmSmEnRegistrationExists(n string, obj secretsmanagerv2.NotificationsRegistration) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
		if err != nil {
			return err
		}

		secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

		getNotificationsRegistrationOptions := &secretsmanagerv2.GetNotificationsRegistrationOptions{}

		notificationsRegistration, _, err := secretsManagerClient.GetNotificationsRegistration(getNotificationsRegistrationOptions)
		if err != nil {
			return err
		}

		obj = *notificationsRegistration

		if err := verifyAttr(*notificationsRegistration.EventNotificationsInstanceCrn, acc.SecretsManagerENInstanceCrn, "snotification instance CRN"); err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIbmSmEnRegistrationDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_en_registration" {
			continue
		}

		getNotificationsRegistrationOptions := &secretsmanagerv2.GetNotificationsRegistrationOptions{}

		// Try to find the key
		_, response, err := secretsManagerClient.GetNotificationsRegistration(getNotificationsRegistrationOptions)

		if err == nil {
			return fmt.Errorf("NotificationsRegistrationPrototype still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for NotificationsRegistrationPrototype (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

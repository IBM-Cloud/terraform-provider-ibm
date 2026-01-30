// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package platformnotifications_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/platformnotifications"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/platformnotificationsv1"
	"github.com/stretchr/testify/assert"
)

func testAccPreCheckPlatformNotifications(t *testing.T) {
	if v := os.Getenv("IC_API_KEY"); v == "" {
		t.Fatal("IC_API_KEY must be set for acceptance tests")
	}
	// Provider configuration is handled by the test framework
}

func TestAccIbmNotificationDistributionListDestinationBasic(t *testing.T) {
	var conf platformnotificationsv1.AddDestination
	accountID := acc.NotificationDistributionListAccountId
	destinationId := acc.NotificationDistributionListDestinationId
	destinationType := "event_notifications"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckPlatformNotifications(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmNotificationDistributionListDestinationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmNotificationDistributionListDestinationConfigBasic(accountID, destinationType, destinationId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmNotificationDistributionListDestinationExists("ibm_notification_distribution_list_destination.notification_distribution_list_destination_instance", conf),
					resource.TestCheckResourceAttr("ibm_notification_distribution_list_destination.notification_distribution_list_destination_instance", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_notification_distribution_list_destination.notification_distribution_list_destination_instance", "destination_type", destinationType),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_notification_distribution_list_destination.notification_distribution_list_destination_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmNotificationDistributionListDestinationConfigBasic(accountID string, destinationType string, destinationId string) string {
	return fmt.Sprintf(`
		resource "ibm_notification_distribution_list_destination" "notification_distribution_list_destination_instance" {
			account_id = "%s"
			destination_type = "%s"
			destination_id = "%s"
		}
	`, accountID, destinationType, destinationId)
}

func testAccCheckIbmNotificationDistributionListDestinationExists(n string, obj platformnotificationsv1.AddDestination) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		platformNotificationsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PlatformNotificationsV1()
		if err != nil {
			return err
		}

		getDistributionListDestinationOptions := &platformnotificationsv1.GetDistributionListDestinationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDistributionListDestinationOptions.SetAccountID(parts[0])
		getDistributionListDestinationOptions.SetDestinationID(parts[1])

		addDestinationIntf, _, err := platformNotificationsClient.GetDistributionListDestination(getDistributionListDestinationOptions)
		if err != nil {
			return err
		}

		switch v := addDestinationIntf.(type) {
		case *platformnotificationsv1.AddDestinationEventNotificationDestination:
			obj = platformnotificationsv1.AddDestination{
				DestinationID:   v.DestinationID,
				DestinationType: v.DestinationType,
			}
		case *platformnotificationsv1.AddDestination:
			obj = *v
		default:
			return fmt.Errorf("unexpected destination type: %T", addDestinationIntf)
		}
		return nil
	}
}

func testAccCheckIbmNotificationDistributionListDestinationDestroy(s *terraform.State) error {
	platformNotificationsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PlatformNotificationsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_notification_distribution_list_destination" {
			continue
		}

		getDistributionListDestinationOptions := &platformnotificationsv1.GetDistributionListDestinationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDistributionListDestinationOptions.SetAccountID(parts[0])
		getDistributionListDestinationOptions.SetDestinationID(parts[1])

		// Try to find the key
		_, response, err := platformNotificationsClient.GetDistributionListDestination(getDistributionListDestinationOptions)

		if err == nil {
			return fmt.Errorf("notification_distribution_list_destination still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for notification_distribution_list_destination (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmNotificationDistributionListDestinationMapToAddDestinationPrototype(t *testing.T) {
	// Checking the result is disabled for this model, because it has a discriminator
	// and there are separate tests for each child model below.
	model := make(map[string]interface{})
	model["destination_id"] = acc.NotificationDistributionListDestinationId
	model["destination_type"] = "event_notifications"

	_, err := platformnotifications.ResourceIbmNotificationDistributionListDestinationMapToAddDestinationPrototype(model)
	assert.Nil(t, err)
}

func TestResourceIbmNotificationDistributionListDestinationMapToAddDestinationPrototypeEventNotificationDestinationPrototype(t *testing.T) {
	checkResult := func(result *platformnotificationsv1.AddDestinationPrototypeEventNotificationDestinationPrototype) {
		model := new(platformnotificationsv1.AddDestinationPrototypeEventNotificationDestinationPrototype)
		mock := strfmt.UUID("12345678-1234-1234-1234-123456789012")
		model.DestinationID = &mock
		model.DestinationType = core.StringPtr("event_notifications")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["destination_id"] = "12345678-1234-1234-1234-123456789012"
	model["destination_type"] = "event_notifications"

	result, err := platformnotifications.ResourceIbmNotificationDistributionListDestinationMapToAddDestinationPrototypeEventNotificationDestinationPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

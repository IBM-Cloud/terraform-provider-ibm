// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package distributionlistapi_test

import (
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/distributionlistapi"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/distributionlistapiv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmDistributionListDestinationBasic(t *testing.T) {
	var conf distributionlistapiv1.AddDestinationResponseBodyEventNotificationDestination
	accountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))
	destinationType := "event_notifications"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmDistributionListDestinationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmDistributionListDestinationConfigBasic(accountID, destinationType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmDistributionListDestinationExists("ibm_distribution_list_destination.distribution_list_destination_instance", conf),
					resource.TestCheckResourceAttr("ibm_distribution_list_destination.distribution_list_destination_instance", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_distribution_list_destination.distribution_list_destination_instance", "destination_type", destinationType),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_distribution_list_destination.distribution_list_destination_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmDistributionListDestinationConfigBasic(accountID string, destinationType string) string {
	return fmt.Sprintf(`
		resource "ibm_distribution_list_destination" "distribution_list_destination_instance" {
			account_id = "%s"
			destination_type = "%s"
		}
	`, accountID, destinationType)
}

func testAccCheckIbmDistributionListDestinationExists(n string, obj distributionlistapiv1.AddDestinationResponseBodyEventNotificationDestination) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		distributionListApiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).DistributionListApiV1()
		if err != nil {
			return err
		}

		getDistributionListDestinationOptions := &distributionlistapiv1.GetDistributionListDestinationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDistributionListDestinationOptions.SetAccountID(parts[0])
		getDistributionListDestinationOptions.SetDestinationID(core.UUIDPtr(strfmt.UUID(parts[1])))

		eventNotificationDestinationIntf, _, err := distributionListApiClient.GetDistributionListDestination(getDistributionListDestinationOptions)
		if err != nil {
			return err
		}

		eventNotificationDestination := eventNotificationDestinationIntf.(*distributionlistapiv1.AddDestinationResponseBodyEventNotificationDestination)
		obj = *eventNotificationDestination
		return nil
	}
}

func testAccCheckIbmDistributionListDestinationDestroy(s *terraform.State) error {
	distributionListApiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).DistributionListApiV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_distribution_list_destination" {
			continue
		}

		getDistributionListDestinationOptions := &distributionlistapiv1.GetDistributionListDestinationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDistributionListDestinationOptions.SetAccountID(parts[0])
		getDistributionListDestinationOptions.SetDestinationID(core.UUIDPtr(strfmt.UUID(parts[1])))

		// Try to find the key
		_, response, err := distributionListApiClient.GetDistributionListDestination(getDistributionListDestinationOptions)

		if err == nil {
			return fmt.Errorf("distribution_list_destination still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for distribution_list_destination (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmDistributionListDestinationMapToAddDestinationResponseBodyPrototypeEventNotificationDestination(t *testing.T) {
	checkResult := func(result *distributionlistapiv1.AddDestinationResponseBodyPrototypeEventNotificationDestination) {
		model := new(distributionlistapiv1.AddDestinationResponseBodyPrototypeEventNotificationDestination)
		model.ID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
		model.DestinationType = core.StringPtr("event_notifications")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "12345678-1234-1234-1234-123456789012"
	model["destination_type"] = "event_notifications"

	result, err := distributionlistapi.ResourceIbmDistributionListDestinationMapToAddDestinationResponseBodyPrototypeEventNotificationDestination(model)
	assert.Nil(t, err)
	checkResult(result)
}

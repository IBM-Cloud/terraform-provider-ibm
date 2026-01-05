// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
*/

package distributionlistapi_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/distributionlistapi"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/distributionlistapiv1"
	"github.com/stretchr/testify/assert"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmDistributionListDestinationsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmDistributionListDestinationsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_distribution_list_destinations.distribution_list_destinations_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_distribution_list_destinations.distribution_list_destinations_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_distribution_list_destinations.distribution_list_destinations_instance", "destinations.#"),
				),
			},
		},
	})
}

func testAccCheckIbmDistributionListDestinationsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_distribution_list_destinations" "distribution_list_destinations_instance" {
			account_id = "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"
		}
	`)
}


func TestDataSourceIbmDistributionListDestinationsAddDestinationResponseBodyToMap(t *testing.T) {
	// Checking the result is disabled for this model, because it has a discriminator
	// and there are separate tests for each child model below.
	model := new(distributionlistapiv1.AddDestinationResponseBody)
	model.ID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
	model.DestinationType = core.StringPtr("event_notifications")

	_, err := distributionlistapi.DataSourceIbmDistributionListDestinationsAddDestinationResponseBodyToMap(model)
	assert.Nil(t, err)
}

func TestDataSourceIbmDistributionListDestinationsAddDestinationResponseBodyEventNotificationDestinationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "12345678-1234-1234-1234-123456789012"
		model["destination_type"] = "event_notifications"

		assert.Equal(t, result, model)
	}

	model := new(distributionlistapiv1.AddDestinationResponseBodyEventNotificationDestination)
	model.ID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
	model.DestinationType = core.StringPtr("event_notifications")

	result, err := distributionlistapi.DataSourceIbmDistributionListDestinationsAddDestinationResponseBodyEventNotificationDestinationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

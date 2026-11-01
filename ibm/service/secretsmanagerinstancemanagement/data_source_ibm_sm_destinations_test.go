// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.2-f568b7d6-20261016-181021
*/

package secretsmanagerinstancemanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/secretsmanagerinstancemanagement"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
	"github.com/stretchr/testify/assert"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmDestinationsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmDestinationsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_destinations.sm_destinations_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destinations.sm_destinations_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destinations.sm_destinations_instance", "destinations.#"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_destinations.sm_destinations_instance", "total"),
				),
			},
		},
	})
}

func testAccCheckIbmSmDestinationsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_sm_destinations" "sm_destinations_instance" {
			instance_id = "bfc50c2e-d66d-4f37-9ccf-9713f8325b39"
			state = "not_started"
		}
	`)
}


func TestDataSourceIbmSmDestinationsDestinationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		model["href"] = "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/destinations/ce8aa9a4-8cde-4b63-8156-d4a8b252beeb"
		model["name"] = "testString"
		model["type"] = "ibm_cloud_database"
		model["description"] = "testString"
		model["state"] = "not_started"
		model["created_at"] = "2019-01-01T12:00:00.000Z"
		model["updated_at"] = "2019-01-01T12:00:00.000Z"
		model["created_by"] = "testString"
		model["message"] = "failed to provision"

		assert.Equal(t, result, model)
	}

	model := new(secretsmanagerinstancemanagementv2.Destination)
	model.ID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
	model.Href = core.StringPtr("https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/destinations/ce8aa9a4-8cde-4b63-8156-d4a8b252beeb")
	model.Name = core.StringPtr("testString")
	model.Type = core.StringPtr("ibm_cloud_database")
	model.Description = core.StringPtr("testString")
	model.State = core.StringPtr("not_started")
	model.CreatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.UpdatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.CreatedBy = core.StringPtr("testString")
	model.Message = core.StringPtr("failed to provision")

	result, err := secretsmanagerinstancemanagement.DataSourceIbmSmDestinationsDestinationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

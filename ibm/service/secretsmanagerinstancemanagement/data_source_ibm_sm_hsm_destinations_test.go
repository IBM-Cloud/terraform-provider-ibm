// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.3-3e91ab11-20261028-182643
 */

package secretsmanagerinstancemanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/secretsmanagerinstancemanagement"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmSmHsmDestinationsDataSourceBasic(t *testing.T) {
	hSMDestinationInstanceID := hsmInstanceID(t)
	hSMDestinationName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	hSMDestinationVendor := "thales_luna"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmHsmDestinationsDataSourceConfigBasic(hSMDestinationInstanceID, hSMDestinationName, hSMDestinationVendor),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "hsm_destinations.#"),
					resource.TestCheckResourceAttr("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "hsm_destinations.0.name", hSMDestinationName),
					resource.TestCheckResourceAttr("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "hsm_destinations.0.vendor", hSMDestinationVendor),
				),
			},
		},
	})
}

func TestAccIbmSmHsmDestinationsDataSourceAllArgs(t *testing.T) {
	hSMDestinationInstanceID := hsmInstanceID(t)
	hSMDestinationName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	hSMDestinationDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	hSMDestinationVendor := "thales_luna"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmHsmDestinationsDataSourceConfig(hSMDestinationInstanceID, hSMDestinationName, hSMDestinationDescription, hSMDestinationVendor),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "hsm_destinations.#"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "hsm_destinations.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "hsm_destinations.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "hsm_destinations.0.account_id"),
					resource.TestCheckResourceAttr("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "hsm_destinations.0.name", hSMDestinationName),
					resource.TestCheckResourceAttr("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "hsm_destinations.0.description", hSMDestinationDescription),
					resource.TestCheckResourceAttr("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "hsm_destinations.0.vendor", hSMDestinationVendor),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "hsm_destinations.0.provisioning_state"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "hsm_destinations.0.health_state"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "hsm_destinations.0.health_as_of"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "hsm_destinations.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "hsm_destinations.0.updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destinations.sm_hsm_destinations_instance", "hsm_destinations.0.created_by"),
				),
			},
		},
	})
}

func testAccCheckIbmSmHsmDestinationsDataSourceConfigBasic(hSMDestinationInstanceID string, hSMDestinationName string, hSMDestinationVendor string) string {
	return fmt.Sprintf(`
		resource "ibm_sm_hsm_destination" "sm_hsm_destination_instance" {
			instance_id = "%s"
			name = "%s"
			vendor = "%s"
		}

		data "ibm_sm_hsm_destinations" "sm_hsm_destinations_instance" {
			instance_id = ibm_sm_hsm_destination.sm_hsm_destination_instance.instance_id
		}
	`, hSMDestinationInstanceID, hSMDestinationName, hSMDestinationVendor)
}

func testAccCheckIbmSmHsmDestinationsDataSourceConfig(hSMDestinationInstanceID string, hSMDestinationName string, hSMDestinationDescription string, hSMDestinationVendor string) string {
	return fmt.Sprintf(`
		resource "ibm_sm_hsm_destination" "sm_hsm_destination_instance" {
			instance_id = "%s"
			name = "%s"
			description = "%s"
			vendor = "%s"
		}

		data "ibm_sm_hsm_destinations" "sm_hsm_destinations_instance" {
			instance_id = ibm_sm_hsm_destination.sm_hsm_destination_instance.instance_id
		}
	`, hSMDestinationInstanceID, hSMDestinationName, hSMDestinationDescription, hSMDestinationVendor)
}

func TestDataSourceIbmSmHsmDestinationsHSMDestinationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "b49ad385-0cb7-3694-a71c-2579d6e94b02"
		model["href"] = "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02"
		model["account_id"] = "a/abc123def456"
		model["name"] = "prod-hsm-destination"
		model["description"] = "Primary Thales Luna destination"
		model["vendor"] = "thales_luna"
		model["provisioning_state"] = "completed"
		model["health_state"] = "healthy"
		model["health_as_of"] = "2026-07-09T15:30:00.000Z"
		model["created_at"] = "2026-07-09T15:00:00.000Z"
		model["updated_at"] = "2026-07-09T15:30:00.000Z"
		model["created_by"] = "iam-ServiceId-abc123"
		model["updated_by"] = "iam-ServiceId-abc123"

		assert.Equal(t, result, model)
	}

	model := new(secretsmanagerinstancemanagementv2.HSMDestination)
	model.ID = core.StringPtr("b49ad385-0cb7-3694-a71c-2579d6e94b02")
	model.Href = core.StringPtr("https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02")
	model.AccountID = core.StringPtr("a/abc123def456")
	model.Name = core.StringPtr("prod-hsm-destination")
	model.Description = core.StringPtr("Primary Thales Luna destination")
	model.Vendor = core.StringPtr("thales_luna")
	model.ProvisioningState = core.StringPtr("completed")
	model.HealthState = core.StringPtr("healthy")
	model.HealthAsOf = CreateMockDateTime("2026-07-09T15:30:00.000Z")
	model.CreatedAt = CreateMockDateTime("2026-07-09T15:00:00.000Z")
	model.UpdatedAt = CreateMockDateTime("2026-07-09T15:30:00.000Z")
	model.CreatedBy = core.StringPtr("iam-ServiceId-abc123")
	model.UpdatedBy = core.StringPtr("iam-ServiceId-abc123")

	result, err := secretsmanagerinstancemanagement.DataSourceIbmSmHsmDestinationsHSMDestinationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

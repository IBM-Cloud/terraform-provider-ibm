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
)

func TestAccIbmSmHsmDestinationDataSourceBasic(t *testing.T) {
	hSMDestinationInstanceID := hsmInstanceID(t)
	hSMDestinationName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	hSMDestinationVendor := "thales_luna"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmHsmDestinationDataSourceConfigBasic(hSMDestinationInstanceID, hSMDestinationName, hSMDestinationVendor),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "destination_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "vendor"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "provisioning_state"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "health_state"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "health_as_of"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "created_by"),
				),
			},
		},
	})
}

func TestAccIbmSmHsmDestinationDataSourceAllArgs(t *testing.T) {
	hSMDestinationInstanceID := hsmInstanceID(t)
	hSMDestinationName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	hSMDestinationDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	hSMDestinationVendor := "thales_luna"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmHsmDestinationDataSourceConfig(hSMDestinationInstanceID, hSMDestinationName, hSMDestinationDescription, hSMDestinationVendor),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "destination_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "vendor"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "provisioning_state"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "health_state"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "health_as_of"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination.sm_hsm_destination_instance", "created_by"),
				),
			},
		},
	})
}

func testAccCheckIbmSmHsmDestinationDataSourceConfigBasic(hSMDestinationInstanceID string, hSMDestinationName string, hSMDestinationVendor string) string {
	return fmt.Sprintf(`
		resource "ibm_sm_hsm_destination" "sm_hsm_destination_instance" {
			instance_id = "%s"
			name = "%s"
			vendor = "%s"
		}

		data "ibm_sm_hsm_destination" "sm_hsm_destination_instance" {
			instance_id = ibm_sm_hsm_destination.sm_hsm_destination_instance.instance_id
			destination_id = ibm_sm_hsm_destination.sm_hsm_destination_instance.destination_id
		}
	`, hSMDestinationInstanceID, hSMDestinationName, hSMDestinationVendor)
}

func testAccCheckIbmSmHsmDestinationDataSourceConfig(hSMDestinationInstanceID string, hSMDestinationName string, hSMDestinationDescription string, hSMDestinationVendor string) string {
	return fmt.Sprintf(`
		resource "ibm_sm_hsm_destination" "sm_hsm_destination_instance" {
			instance_id = "%s"
			name = "%s"
			description = "%s"
			vendor = "%s"
		}

		data "ibm_sm_hsm_destination" "sm_hsm_destination_instance" {
			instance_id = ibm_sm_hsm_destination.sm_hsm_destination_instance.instance_id
			destination_id = ibm_sm_hsm_destination.sm_hsm_destination_instance.destination_id
		}
	`, hSMDestinationInstanceID, hSMDestinationName, hSMDestinationDescription, hSMDestinationVendor)
}

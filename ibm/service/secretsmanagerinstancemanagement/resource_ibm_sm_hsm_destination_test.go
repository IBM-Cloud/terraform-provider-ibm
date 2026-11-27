// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanagerinstancemanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func TestAccIbmSmHsmDestinationBasic(t *testing.T) {
	var conf secretsmanagerinstancemanagementv2.HSMDestination
	instanceID := hsmInstanceID(t)
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	vendor := "thales_luna"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmHsmDestinationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmHsmDestinationConfigBasic(instanceID, name, vendor),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmHsmDestinationExists("ibm_sm_hsm_destination.sm_hsm_destination_instance", conf),
					resource.TestCheckResourceAttr("ibm_sm_hsm_destination.sm_hsm_destination_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_sm_hsm_destination.sm_hsm_destination_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_sm_hsm_destination.sm_hsm_destination_instance", "vendor", vendor),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSmHsmDestinationConfigBasic(instanceID, nameUpdate, vendor),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_sm_hsm_destination.sm_hsm_destination_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_sm_hsm_destination.sm_hsm_destination_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_sm_hsm_destination.sm_hsm_destination_instance", "vendor", vendor),
				),
			},
		},
	})
}

func TestAccIbmSmHsmDestinationAllArgs(t *testing.T) {
	var conf secretsmanagerinstancemanagementv2.HSMDestination
	instanceID := hsmInstanceID(t)
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	vendor := "thales_luna"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmHsmDestinationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmHsmDestinationConfig(instanceID, name, description, vendor),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmHsmDestinationExists("ibm_sm_hsm_destination.sm_hsm_destination_instance", conf),
					resource.TestCheckResourceAttr("ibm_sm_hsm_destination.sm_hsm_destination_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_sm_hsm_destination.sm_hsm_destination_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_sm_hsm_destination.sm_hsm_destination_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_sm_hsm_destination.sm_hsm_destination_instance", "vendor", vendor),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSmHsmDestinationConfig(instanceID, nameUpdate, descriptionUpdate, vendor),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_sm_hsm_destination.sm_hsm_destination_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_sm_hsm_destination.sm_hsm_destination_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_sm_hsm_destination.sm_hsm_destination_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_sm_hsm_destination.sm_hsm_destination_instance", "vendor", vendor),
				),
			},
			resource.TestStep{
				ResourceName:            "ibm_sm_hsm_destination.sm_hsm_destination_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"}, // set on create, not readable on import
			},
		},
	})
}

func testAccCheckIbmSmHsmDestinationConfigBasic(instanceID string, name string, vendor string) string {
	return fmt.Sprintf(`
		resource "ibm_sm_hsm_destination" "sm_hsm_destination_instance" {
			instance_id = "%s"
			name = "%s"
			vendor = "%s"
		}
	`, instanceID, name, vendor)
}

func testAccCheckIbmSmHsmDestinationConfig(instanceID string, name string, description string, vendor string) string {
	return fmt.Sprintf(`

		resource "ibm_sm_hsm_destination" "sm_hsm_destination_instance" {
			instance_id = "%s"
			name = "%s"
			description = "%s"
			vendor = "%s"
		}
	`, instanceID, name, description, vendor)
}

func testAccCheckIbmSmHsmDestinationExists(n string, obj secretsmanagerinstancemanagementv2.HSMDestination) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		secretsManagerInstanceManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerInstanceManagementV2()
		if err != nil {
			return err
		}

		getInstanceHsmDestinationOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getInstanceHsmDestinationOptions.SetInstanceID(parts[0])
		getInstanceHsmDestinationOptions.SetDestinationID(parts[1])

		hSMDestination, _, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestination(getInstanceHsmDestinationOptions)
		if err != nil {
			return err
		}

		obj = *hSMDestination
		return nil
	}
}

func testAccCheckIbmSmHsmDestinationDestroy(s *terraform.State) error {
	secretsManagerInstanceManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_hsm_destination" {
			continue
		}

		getInstanceHsmDestinationOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getInstanceHsmDestinationOptions.SetInstanceID(parts[0])
		getInstanceHsmDestinationOptions.SetDestinationID(parts[1])

		// Try to find the key
		_, response, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestination(getInstanceHsmDestinationOptions)

		if err == nil {
			return fmt.Errorf("sm_hsm_destination still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("error checking for sm_hsm_destination (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

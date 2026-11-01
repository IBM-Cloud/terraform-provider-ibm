// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanagerinstancemanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/secretsmanagerinstancemanagement"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
	"github.com/stretchr/testify/assert"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmDestinationBasic(t *testing.T) {
	var conf secretsmanagerinstancemanagementv2.IbmCloudDatabaseDestination
	instanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	typeVar := "ibm_cloud_database"
	crn := fmt.Sprintf("tf_crn_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := "ibm_cloud_database"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmDestinationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmDestinationConfigBasic(instanceID, name, typeVar, crn),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmDestinationExists("ibm_sm_destination.sm_destination_instance", conf),
					resource.TestCheckResourceAttr("ibm_sm_destination.sm_destination_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_sm_destination.sm_destination_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_sm_destination.sm_destination_instance", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_sm_destination.sm_destination_instance", "crn", crn),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSmDestinationConfigBasic(instanceID, nameUpdate, typeVarUpdate, crn),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_sm_destination.sm_destination_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_sm_destination.sm_destination_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_sm_destination.sm_destination_instance", "type", typeVarUpdate),
					resource.TestCheckResourceAttr("ibm_sm_destination.sm_destination_instance", "crn", crn),
				),
			},
		},
	})
}

func TestAccIbmSmDestinationAllArgs(t *testing.T) {
	var conf secretsmanagerinstancemanagementv2.IbmCloudDatabaseDestination
	instanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	typeVar := "ibm_cloud_database"
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	crn := fmt.Sprintf("tf_crn_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := "ibm_cloud_database"
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmDestinationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmDestinationConfig(instanceID, name, typeVar, description, crn),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmDestinationExists("ibm_sm_destination.sm_destination_instance", conf),
					resource.TestCheckResourceAttr("ibm_sm_destination.sm_destination_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_sm_destination.sm_destination_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_sm_destination.sm_destination_instance", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_sm_destination.sm_destination_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_sm_destination.sm_destination_instance", "crn", crn),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSmDestinationConfig(instanceID, nameUpdate, typeVarUpdate, descriptionUpdate, crn),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_sm_destination.sm_destination_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_sm_destination.sm_destination_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_sm_destination.sm_destination_instance", "type", typeVarUpdate),
					resource.TestCheckResourceAttr("ibm_sm_destination.sm_destination_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_sm_destination.sm_destination_instance", "crn", crn),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_sm_destination.sm_destination_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSmDestinationConfigBasic(instanceID string, name string, typeVar string, crn string) string {
	return fmt.Sprintf(`
		resource "ibm_sm_destination" "sm_destination_instance" {
			instance_id = "%s"
			name = "%s"
			type = "%s"
			crn = "%s"
		}
	`, instanceID, name, typeVar, crn)
}

func testAccCheckIbmSmDestinationConfig(instanceID string, name string, typeVar string, description string, crn string) string {
	return fmt.Sprintf(`

		resource "ibm_sm_destination" "sm_destination_instance" {
			instance_id = "%s"
			name = "%s"
			type = "%s"
			description = "%s"
			crn = "%s"
		}
	`, instanceID, name, typeVar, description, crn)
}

func testAccCheckIbmSmDestinationExists(n string, obj secretsmanagerinstancemanagementv2.IbmCloudDatabaseDestination) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		secretsManagerInstanceManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerInstanceManagementV2()
		if err != nil {
			return err
		}

		getInstanceDestinationOptions := &secretsmanagerinstancemanagementv2.GetInstanceDestinationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getInstanceDestinationOptions.SetInstanceID(parts[0])
		getInstanceDestinationOptions.SetDestinationID(parts[1])

		ibmCloudDatabaseDestination, _, err := secretsManagerInstanceManagementClient.GetInstanceDestination(getInstanceDestinationOptions)
		if err != nil {
			return err
		}

		obj = *ibmCloudDatabaseDestination
		return nil
	}
}

func testAccCheckIbmSmDestinationDestroy(s *terraform.State) error {
	secretsManagerInstanceManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_destination" {
			continue
		}

		getInstanceDestinationOptions := &secretsmanagerinstancemanagementv2.GetInstanceDestinationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getInstanceDestinationOptions.SetInstanceID(parts[0])
		getInstanceDestinationOptions.SetDestinationID(parts[1])

		// Try to find the key
		_, response, err := secretsManagerInstanceManagementClient.GetInstanceDestination(getInstanceDestinationOptions)

		if err == nil {
			return fmt.Errorf("sm_destination still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("error checking for sm_destination (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmSmDestinationMapToCreateDestinationRequest(t *testing.T) {
	// Checking the result is disabled for this model, because it has a discriminator
	// and there are separate tests for each child model below.
	model := make(map[string]interface{})
	model["name"] = "testString"
	model["type"] = "ibm_cloud_database"
	model["description"] = "Production PostgreSQL database"
	model["crn"] = "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/e91c8f42b3d74e1a9c2f05d8b67a3e10:3f8b1c7a-9d42-4e6f-b8a5-2c1d9e7f4b83::"

	_, err := secretsmanagerinstancemanagement.ResourceIbmSmDestinationMapToCreateDestinationRequest(model)
	assert.Nil(t, err)
}

func TestResourceIbmSmDestinationMapToCreateDestinationRequestIbmCloudDatabaseDestinationPrototype(t *testing.T) {
	checkResult := func(result *secretsmanagerinstancemanagementv2.CreateDestinationRequestIbmCloudDatabaseDestinationPrototype) {
		model := new(secretsmanagerinstancemanagementv2.CreateDestinationRequestIbmCloudDatabaseDestinationPrototype)
		model.Name = core.StringPtr("testString")
		model.Type = core.StringPtr("ibm_cloud_database")
		model.Description = core.StringPtr("Production PostgreSQL database")
		model.Crn = core.StringPtr("crn:v1:bluemix:public:databases-for-postgresql:us-south:a/e91c8f42b3d74e1a9c2f05d8b67a3e10:3f8b1c7a-9d42-4e6f-b8a5-2c1d9e7f4b83::")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "testString"
	model["type"] = "ibm_cloud_database"
	model["description"] = "Production PostgreSQL database"
	model["crn"] = "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/e91c8f42b3d74e1a9c2f05d8b67a3e10:3f8b1c7a-9d42-4e6f-b8a5-2c1d9e7f4b83::"

	result, err := secretsmanagerinstancemanagement.ResourceIbmSmDestinationMapToCreateDestinationRequestIbmCloudDatabaseDestinationPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

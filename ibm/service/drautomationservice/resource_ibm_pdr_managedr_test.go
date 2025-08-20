// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package drautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/drautomationservice"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func TestAccIbmPdrManagedrBasic(t *testing.T) {
	var conf drautomationservicev1.ServiceInstanceManageDR
	instanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	standByRedeploy := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmPdrManagedrDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmPdrManagedrConfigBasic(instanceID, standByRedeploy),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmPdrManagedrExists("ibm_pdr_managedr.pdr_managedr_instance", conf),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "stand_by_redeploy", standByRedeploy),
				),
			},
		},
	})
}

func TestAccIbmPdrManagedrAllArgs(t *testing.T) {
	var conf drautomationservicev1.ServiceInstanceManageDR
	instanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	standByRedeploy := "true"
	acceptLanguage := fmt.Sprintf("tf_accept_language_%d", acctest.RandIntRange(10, 100))
	ifNoneMatch := fmt.Sprintf("tf_if_none_match_%d", acctest.RandIntRange(10, 100))
	acceptsIncomplete := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmPdrManagedrDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmPdrManagedrConfig(instanceID, standByRedeploy, acceptLanguage, ifNoneMatch, acceptsIncomplete),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmPdrManagedrExists("ibm_pdr_managedr.pdr_managedr_instance", conf),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "stand_by_redeploy", standByRedeploy),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "accept_language", acceptLanguage),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "if_none_match", ifNoneMatch),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "accepts_incomplete", acceptsIncomplete),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_pdr_managedr.pdr_managedr",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmPdrManagedrConfigBasic(instanceID string, standByRedeploy string) string {
	return fmt.Sprintf(`
		resource "ibm_pdr_managedr" "pdr_managedr_instance" {
			instance_id = "%s"
			stand_by_redeploy = "%s"
		}
	`, instanceID, standByRedeploy)
}

func testAccCheckIbmPdrManagedrConfig(instanceID string, standByRedeploy string, acceptLanguage string, ifNoneMatch string, acceptsIncomplete string) string {
	return fmt.Sprintf(`

		resource "ibm_pdr_managedr" "pdr_managedr_instance" {
			instance_id = "%s"
			stand_by_redeploy = "%s"
			accept_language = "%s"
			if_none_match = "%s"
			accepts_incomplete = %s
		}
	`, instanceID, standByRedeploy, acceptLanguage, ifNoneMatch, acceptsIncomplete)
}

func testAccCheckIbmPdrManagedrExists(n string, obj drautomationservicev1.ServiceInstanceManageDR) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		drAutomationServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).DrAutomationServiceV1()
		if err != nil {
			return err
		}

		serviceInstanceFetchManageDrOptions := &drautomationservicev1.ServiceInstanceFetchManageDrOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		serviceInstanceFetchManageDrOptions.SetInstanceID(parts[0])
		serviceInstanceFetchManageDrOptions.SetInstanceID(parts[1])

		serviceInstanceManageDR, _, err := drAutomationServiceClient.ServiceInstanceFetchManageDr(serviceInstanceFetchManageDrOptions)
		if err != nil {
			return err
		}

		obj = *serviceInstanceManageDR
		return nil
	}
}

func testAccCheckIbmPdrManagedrDestroy(s *terraform.State) error {
	drAutomationServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pdr_managedr" {
			continue
		}

		serviceInstanceFetchManageDrOptions := &drautomationservicev1.ServiceInstanceFetchManageDrOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		serviceInstanceFetchManageDrOptions.SetInstanceID(parts[0])
		serviceInstanceFetchManageDrOptions.SetInstanceID(parts[1])

		// Try to find the key
		_, response, err := drAutomationServiceClient.ServiceInstanceFetchManageDr(serviceInstanceFetchManageDrOptions)

		if err == nil {
			return fmt.Errorf("pdr_managedr still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for pdr_managedr (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmPdrManagedrMapToContext(t *testing.T) {
	checkResult := func(result *drautomationservicev1.Context) {
		model := new(drautomationservicev1.Context)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})

	result, err := drautomationservice.ResourceIbmPdrManagedrMapToContext(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmPdrManagedrMapToManageDrParameters(t *testing.T) {
	checkResult := func(result *drautomationservicev1.ManageDrParameters) {
		model := new(drautomationservicev1.ManageDrParameters)
		model.Location = core.StringPtr("us-south")
		model.OptionalParam = core.StringPtr("parameter required by your service")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["location"] = "us-south"
	model["optional_param"] = "parameter required by your service"

	result, err := drautomationservice.ResourceIbmPdrManagedrMapToManageDrParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logsrouter_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logsrouter"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/logsrouterv3"
	"github.com/stretchr/testify/assert"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMLogsRouterTargetBasic(t *testing.T) {
	var conf logsrouterv3.Target
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	destinationCRN := fmt.Sprintf("tf_destination_crn_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	destinationCRNUpdate := fmt.Sprintf("tf_destination_crn_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLogsRouterTargetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterTargetConfigBasic(name, destinationCRN),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMLogsRouterTargetExists("ibm_logs_router_target.logs_router_target_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_router_target.logs_router_target_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_router_target.logs_router_target_instance", "destination_crn", destinationCRN),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterTargetConfigBasic(nameUpdate, destinationCRNUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_router_target.logs_router_target_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_router_target.logs_router_target_instance", "destination_crn", destinationCRNUpdate),
				),
			},
		},
	})
}

func TestAccIBMLogsRouterTargetAllArgs(t *testing.T) {
	var conf logsrouterv3.Target
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	destinationCRN := fmt.Sprintf("tf_destination_crn_%d", acctest.RandIntRange(10, 100))
	region := fmt.Sprintf("tf_region_%d", acctest.RandIntRange(10, 100))
	managedBy := "enterprise"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	destinationCRNUpdate := fmt.Sprintf("tf_destination_crn_%d", acctest.RandIntRange(10, 100))
	regionUpdate := fmt.Sprintf("tf_region_%d", acctest.RandIntRange(10, 100))
	managedByUpdate := "account"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLogsRouterTargetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterTargetConfig(name, destinationCRN, region, managedBy),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMLogsRouterTargetExists("ibm_logs_router_target.logs_router_target_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_router_target.logs_router_target_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_router_target.logs_router_target_instance", "destination_crn", destinationCRN),
					resource.TestCheckResourceAttr("ibm_logs_router_target.logs_router_target_instance", "region", region),
					resource.TestCheckResourceAttr("ibm_logs_router_target.logs_router_target_instance", "managed_by", managedBy),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterTargetConfig(nameUpdate, destinationCRNUpdate, regionUpdate, managedByUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_router_target.logs_router_target_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_router_target.logs_router_target_instance", "destination_crn", destinationCRNUpdate),
					resource.TestCheckResourceAttr("ibm_logs_router_target.logs_router_target_instance", "region", regionUpdate),
					resource.TestCheckResourceAttr("ibm_logs_router_target.logs_router_target_instance", "managed_by", managedByUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_router_target.logs_router_target_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMLogsRouterTargetConfigBasic(name string, destinationCRN string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_router_target" "logs_router_target_instance" {
			name = "%s"
			destination_crn = "%s"
		}
	`, name, destinationCRN)
}

func testAccCheckIBMLogsRouterTargetConfig(name string, destinationCRN string, region string, managedBy string) string {
	return fmt.Sprintf(`

		resource "ibm_logs_router_target" "logs_router_target_instance" {
			name = "%s"
			destination_crn = "%s"
			region = "%s"
			managed_by = "%s"
		}
	`, name, destinationCRN, region, managedBy)
}

func testAccCheckIBMLogsRouterTargetExists(n string, obj logsrouterv3.Target) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		logsRouterClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsRouterV3()
		if err != nil {
			return err
		}

		getTargetOptions := &logsrouterv3.GetTargetOptions{}

		getTargetOptions.SetID(rs.Primary.ID)

		target, _, err := logsRouterClient.GetTarget(getTargetOptions)
		if err != nil {
			return err
		}

		obj = *target
		return nil
	}
}

func testAccCheckIBMLogsRouterTargetDestroy(s *terraform.State) error {
	logsRouterClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsRouterV3()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_router_target" {
			continue
		}

		getTargetOptions := &logsrouterv3.GetTargetOptions{}

		getTargetOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := logsRouterClient.GetTarget(getTargetOptions)

		if err == nil {
			return fmt.Errorf("logs_router_target still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_router_target (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMLogsRouterTargetWriteStatusToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["status"] = "success"
		model["last_failure"] = "2025-05-18T20:15:12.353Z"
		model["reason_for_last_failure"] = "Provided API key could not be found"

		assert.Equal(t, result, model)
	}

	model := new(logsrouterv3.WriteStatus)
	model.Status = core.StringPtr("success")
	model.LastFailure = CreateMockDateTime("2025-05-18T20:15:12.353Z")
	model.ReasonForLastFailure = core.StringPtr("Provided API key could not be found")

	result, err := logsrouter.ResourceIBMLogsRouterTargetWriteStatusToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

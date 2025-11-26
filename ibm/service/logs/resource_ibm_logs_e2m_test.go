// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func TestAccIbmLogsE2mBasic(t *testing.T) {
	var conf logsv0.Event2Metric
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsE2mDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsE2mConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsE2mExists("ibm_logs_e2m.logs_e2m_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_e2m.logs_e2m_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsE2mConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_e2m.logs_e2m_instance", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIbmLogsE2mAllArgs(t *testing.T) {
	var conf logsv0.Event2Metric
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	typeVar := "logs2metrics"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := "logs2metrics"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsE2mDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsE2mConfig(name, description, typeVar),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsE2mExists("ibm_logs_e2m.logs_e2m_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_e2m.logs_e2m_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_e2m.logs_e2m_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_logs_e2m.logs_e2m_instance", "type", typeVar),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsE2mConfig(nameUpdate, descriptionUpdate, typeVarUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_e2m.logs_e2m_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_e2m.logs_e2m_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_logs_e2m.logs_e2m_instance", "type", typeVarUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_e2m.logs_e2m_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsE2mConfigBasic(name string) string {
	return fmt.Sprintf(`
	resource "ibm_logs_e2m" "logs_e2m_instance" {
		instance_id = "%s"
		region      = "%s"
		name        = "%s"
		description = "test description"
		logs_query {
		  applicationname_filters = []
		  severity_filters = [
			"debug", "error"
		  ]
		  subsystemname_filters = []
		}
		type = "logs2metrics"
	}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name)
}

func testAccCheckIbmLogsE2mConfig(name string, description string, typeVar string) string {
	return fmt.Sprintf(`
	
	resource "ibm_logs_e2m" "logs_e2m_instance" {
		instance_id = "%s"
		region      = "%s"
		name        = "%s"
		description = "%s"
		logs_query {
		  applicationname_filters = []
		  severity_filters = [
			"debug", "error"
		  ]
		  subsystemname_filters = []
		}
		type = "%s"
	}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name, description, typeVar)
}

func testAccCheckIbmLogsE2mExists(n string, obj logsv0.Event2Metric) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
		if err != nil {
			return err
		}
		logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getE2mOptions := &logsv0.GetE2mOptions{}

		getE2mOptions.SetID(resourceID[2])

		event2MetricIntf, _, err := logsClient.GetE2m(getE2mOptions)
		if err != nil {
			return err
		}

		event2Metric := event2MetricIntf.(*logsv0.Event2Metric)
		obj = *event2Metric
		return nil
	}
}

func testAccCheckIbmLogsE2mDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_e2m" {
			continue
		}
		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getE2mOptions := &logsv0.GetE2mOptions{}

		getE2mOptions.SetID(resourceID[2])

		// Try to find the key
		_, response, err := logsClient.GetE2m(getE2mOptions)

		if err == nil {
			return fmt.Errorf("logs_e2m still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_e2m (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

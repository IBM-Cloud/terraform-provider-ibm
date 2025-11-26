// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventstreams_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/eventstreams-go-sdk/pkg/adminrestv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMEventStreamsQuotaResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEventStreamsQuotasDeletedFromInstance(testQuotaEntity3, testQuotaEntity4),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEventStreamsQuotaResourceConfig(getTestInstanceName(mzrKey), 0, testQuotaEntity3, 2048, 1024),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEventStreamsQuotaResourceProperties("ibm_event_streams_quota.es_quota0", testQuotaEntity3),
					testAccCheckIBMEventStreamsQuotaWasSetInInstance(testQuotaEntity3, 2048, 1024),
					resource.TestCheckResourceAttrSet("ibm_event_streams_quota.es_quota0", "id"),
					resource.TestCheckResourceAttr("ibm_event_streams_quota.es_quota0", "producer_byte_rate", "2048"),
					resource.TestCheckResourceAttr("ibm_event_streams_quota.es_quota0", "consumer_byte_rate", "1024"),
				),
			},
			{
				Config: testAccCheckIBMEventStreamsQuotaResourceConfig(getTestInstanceName(mzrKey), 1, testQuotaEntity4, 100000000, -1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEventStreamsQuotaResourceProperties("ibm_event_streams_quota.es_quota1", testQuotaEntity4),
					testAccCheckIBMEventStreamsQuotaWasSetInInstance(testQuotaEntity4, 100000000, -1),
					resource.TestCheckResourceAttrSet("ibm_event_streams_quota.es_quota1", "id"),
					resource.TestCheckResourceAttr("ibm_event_streams_quota.es_quota1", "producer_byte_rate", "100000000"),
					resource.TestCheckResourceAttr("ibm_event_streams_quota.es_quota1", "consumer_byte_rate", "-1"),
				),
			},
			{
				ResourceName:      "ibm_event_streams_quota.es_quota1",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMEventStreamsQuotaResourceConfig(instanceName string, tnum int, entity string, producerRate int, consumerRate int) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "group" {
			is_default=true
	  	}
	  	data "ibm_resource_instance" "es_instance" {
			resource_group_id = data.ibm_resource_group.group.id
			name              = "%s"
	  	}
		resource "ibm_event_streams_quota" "es_quota%d" {
			resource_instance_id = data.ibm_resource_instance.es_instance.id
			entity = "%s"
			producer_byte_rate = %d
			consumer_byte_rate = %d
		}`, instanceName, tnum, entity, producerRate, consumerRate)
}

// check properties of the terraform resource object
func testAccCheckIBMEventStreamsQuotaResourceProperties(name string, entity string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}
		quotaID := rs.Primary.ID
		if quotaID == "" {
			return fmt.Errorf("[ERROR] Quota ID is not set")
		}
		if !strings.HasSuffix(quotaID, fmt.Sprintf(":quota:%s", entity)) {
			return fmt.Errorf("[ERROR] Quota ID %s not expected CRN", quotaID)
		}
		return nil
	}
}

// go to the Event Streams instance and check the quota has been set
func testAccCheckIBMEventStreamsQuotaWasSetInInstance(entity string, producerRate int, consumerRate int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		adminClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ESadminRestSession()
		if err != nil {
			return fmt.Errorf("[ERROR] ESadminRestSession returned %v", err)
		}
		entityName := entity
		qd, _, err := adminClient.GetQuota(&adminrestv1.GetQuotaOptions{EntityName: &entityName})
		if err != nil {
			return fmt.Errorf("[ERROR] GetQuota returned %v", err)
		}
		qdp := testGetQuotaValue(qd.ProducerByteRate)
		if producerRate != qdp {
			return fmt.Errorf("[ERROR] quota producer byte rate expected %d, got %d", producerRate, qdp)
		}
		qdc := testGetQuotaValue(qd.ConsumerByteRate)
		if consumerRate != qdc {
			return fmt.Errorf("[ERROR] quota consumer byte rate expected %d, got %d", consumerRate, qdc)
		}
		return nil
	}
}

// go to the Event Streams instance and check the quota has been destroyed
func testAccCheckIBMEventStreamsQuotasDeletedFromInstance(entities ...string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		adminClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ESadminRestSession()
		if err != nil {
			return fmt.Errorf("[ERROR] ESadminRestSession returned %v", err)
		}
		for _, entity := range entities {
			entityName := entity
			qd, response, err := adminClient.GetQuota(&adminrestv1.GetQuotaOptions{EntityName: &entityName})
			if err == nil {
				return fmt.Errorf("[ERROR] Expected no quota for %s, but GetQuota succeeded (%d,%d)", entity, testGetQuotaValue(qd.ProducerByteRate), testGetQuotaValue(qd.ConsumerByteRate))
			}
			if response != nil && response.StatusCode != 404 {
				return fmt.Errorf("[ERROR] Expected 404 NotFound for %s, but GetQuota response was %d", entity, response.StatusCode)
			}
		}
		return nil
	}
}

func testGetQuotaValue(v *int64) int {
	if v == nil {
		return -1
	}
	return int(*v)
}

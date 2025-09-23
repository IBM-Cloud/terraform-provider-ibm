// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventstreams_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	// Data source test requires MZR instance have this quota with producer rate 10000, consumer rate 20000
	testQuotaEntity1 = "iam-ServiceId-00001111-2222-3333-4444-555566667777"
	// Data source test requires MZR instance have this quota with producer rate 4096, consumer rate not defined
	testQuotaEntity2 = "iam-ServiceId-77776666-5555-4444-3333-222211110000"
	// Resource test requires MZR instance NOT have a quota for this
	testQuotaEntity3 = "iam-ServiceId-99998888-7777-6666-5555-444433332222"
	// Resource test requires MZR instance NOT have a quota for this
	testQuotaEntity4 = "default"
)

func TestAccIBMEventStreamsQuotaDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEventStreamsQuotaDataSourceConfig(getTestInstanceName(mzrKey), testQuotaEntity1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMEventStreamsQuotaDataSourceProperties("data.ibm_event_streams_quota.es_quota", testQuotaEntity1, "10000", "20000"),
				),
			},
			{
				Config: testAccCheckIBMEventStreamsQuotaDataSourceConfig(getTestInstanceName(mzrKey), testQuotaEntity2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMEventStreamsQuotaDataSourceProperties("data.ibm_event_streams_quota.es_quota", testQuotaEntity2, "4096", "-1"),
				),
			},
		},
	})
}

func testAccCheckIBMEventStreamsQuotaDataSourceConfig(instanceName string, entity string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default=true
	}
	data "ibm_resource_instance" "es_instance" {
		resource_group_id = data.ibm_resource_group.group.id
		name              = "%s"
	}
	data "ibm_event_streams_quota" "es_quota" {
		resource_instance_id = data.ibm_resource_instance.es_instance.id
		entity = "%s"
	}`, instanceName, entity)
}

// check properties of the terraform data source object
func testAccCheckIBMEventStreamsQuotaDataSourceProperties(name string, entity string, producerRate string, consumerRate string) resource.TestCheckFunc {
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
			return fmt.Errorf("[ERROR] Quota ID for %s not expected CRN", quotaID)
		}
		if producerRate != rs.Primary.Attributes["producer_byte_rate"] {
			return fmt.Errorf("[ERROR] Quota for %s producer_byte_rate = %s, expected %s", entity, rs.Primary.Attributes["producer_byte_rate"], producerRate)
		}
		if consumerRate != rs.Primary.Attributes["consumer_byte_rate"] {
			return fmt.Errorf("[ERROR] Quota for %s consumer_byte_rate = %s, expected %s", entity, rs.Primary.Attributes["consumer_byte_rate"], consumerRate)
		}
		return nil
	}
}

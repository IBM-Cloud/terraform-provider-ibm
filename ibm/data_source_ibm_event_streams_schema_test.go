// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var (
	mySchemaID = "myschema"
)

func TestAccIBMEventStreamsSchemaDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEventStreamsSchemaDataSourceConfigBasic(MZREnterpriseInstanceName, mySchemaID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_http_url"),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_schema.es_schema", "kafka_http_url"),
					resource.TestCheckResourceAttr("data.ibm_event_streams_schema.es_schema", "schema_id", mySchemaID),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_schema.es_schema", "id"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMEventStreamsSchemaDataSourceConfigBasic(SZREnterpriseInstanceName, mySchemaID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_http_url"),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_schema.es_schema", "kafka_http_url"),
					resource.TestCheckResourceAttr("data.ibm_event_streams_schema.es_schema", "schema_id", mySchemaID),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_schema.es_schema", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMEventStreamsSchemaDataSourceConfigBasic(instanceName, schemaID string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "my_group" {
		is_default=true
	  }
	data "ibm_resource_instance" "es_instance" {
		resource_group_id = data.ibm_resource_group.my_group.id
		name              = "%s"
	}
	data "ibm_event_streams_schema" "es_schema" {
		resource_instance_id = data.ibm_resource_instance.es_instance.id
		schema_id = "%s"
	}`, instanceName, schemaID)
}

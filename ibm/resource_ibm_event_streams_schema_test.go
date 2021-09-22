// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"gotest.tools/assert"

	"github.com/IBM/eventstreams-go-sdk/pkg/schemaregistryv1"
)

func TestAccIBMEventStreamsSchemaBasic(t *testing.T) {
	var conf map[string]interface{}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMEventStreamsSchemaDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEventStreamsSchemaConfigBasicWithExistingInstance(SZREnterpriseInstanceName, "szr"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEventStreamsSchemaExists("ibm_event_streams_schema.es_schema", conf, ""),
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_http_url"),
					resource.TestCheckResourceAttrSet("ibm_event_streams_schema.es_schema", "schema"),
					resource.TestCheckResourceAttrSet("ibm_event_streams_schema.es_schema", "id"),
					resource.TestCheckResourceAttrSet("ibm_event_streams_schema.es_schema", "schema_id"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMEventStreamsSchemaConfigBasicWithExistingInstance(MZREnterpriseInstanceName, "mzr"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEventStreamsSchemaExists("ibm_event_streams_schema.es_schema", conf, ""),
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_http_url"),
					resource.TestCheckResourceAttrSet("ibm_event_streams_schema.es_schema", "schema"),
					resource.TestCheckResourceAttrSet("ibm_event_streams_schema.es_schema", "id"),
					resource.TestCheckResourceAttrSet("ibm_event_streams_schema.es_schema", "schema_id"),
				),
			},
		},
	})
}

func TestAccIBMEventStreamsSchemaAllArgs(t *testing.T) {
	var conf map[string]interface{}
	schemaID := fmt.Sprintf("tf_schema_id_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMEventStreamsSchemaDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEventStreamsSchemaWithSchemaIDWithExistingInstance(MZREnterpriseInstanceName, schemaID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEventStreamsSchemaExists("ibm_event_streams_schema.es_schema", conf, schemaID),
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_http_url"),
					resource.TestCheckResourceAttrSet("ibm_event_streams_schema.es_schema", "schema"),
					resource.TestCheckResourceAttrSet("ibm_event_streams_schema.es_schema", "id"),
					resource.TestCheckResourceAttr("ibm_event_streams_schema.es_schema", "schema_id", schemaID),
				),
			},
		},
	})
}

func TestAccIBMEventStreamsSchemaImport(t *testing.T) {
	var conf map[string]interface{}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMEventStreamsSchemaDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEventStreamsSchemaWithSchemaIDWithExistingInstance(MZREnterpriseInstanceName, schemaID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEventStreamsSchemaExists("ibm_event_streams_schema.es_schema", conf, schemaID),
					resource.TestCheckResourceAttrSet("ibm_event_streams_schema.es_schema", "schema"),
					resource.TestCheckResourceAttr("ibm_event_streams_schema.es_schema", "schema_id", schemaID),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_event_streams_schema.es_schema",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMEventStreamsSchemaConfigBasicWithExistingInstance(instanceName string, prefix string) string {
	s := getPlatformResource(instanceName) + "\n" + createEventStreamsSchemaResourceWithoutSchemaID(false, prefix)
	return s
}

func testAccCheckIBMEventStreamsSchemaWithSchemaIDWithExistingInstance(instanceName, schemaID string) string {
	s := getPlatformResource(instanceName) + "\n" + createEventStreamsSchemaResourceWithSchemaID(false, schemaID)
	return s
}

func createEventStreamsSchemaResourceWithoutSchemaID(createInstance bool, prefix string) string {
	var resourceInstanceID string
	if createInstance {
		resourceInstanceID = "ibm_resource_instance.es_instance.id"
	} else {
		resourceInstanceID = "data.ibm_resource_instance.es_instance.id"
	}
	recordName := fmt.Sprintf("record_%s", prefix)
	return fmt.Sprintf(`
	resource "ibm_event_streams_schema" "es_schema" {
		resource_instance_id 	= %s
		schema           		= <<SCHEMA
		{
			"type": "record",
			"name": "%s",
			"fields" : [
			  {"name": "value_1_1", "type": "long"},
			  {"name": "value_2_1", "type": "string"}
			]
		}
		SCHEMA
	}`, resourceInstanceID, recordName)
}

func createEventStreamsSchemaResourceWithSchemaID(createInstance bool, schemaID string) string {
	var resourceInstanceID string
	if createInstance {
		resourceInstanceID = "ibm_resource_instance.es_instance.id"
	} else {
		resourceInstanceID = "data.ibm_resource_instance.es_instance.id"
	}
	return fmt.Sprintf(`
	resource "ibm_event_streams_schema" "es_schema" {
		resource_instance_id 	= %s
		schema_id 			= "%s"
		schema           		= <<SCHEMA
		{
			"type": "record",
			"name": "record_name",
			"fields" : [
			  {"name": "value_1", "type": "long"},
			  {"name": "value_2", "type": "string"}
			]
		}
		SCHEMA
	}`, resourceInstanceID, schemaID)
}

func testAccCheckIBMEventStreamsSchemaExists(n string, obj map[string]interface{}, schemaID string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		schemaregistryClient, err := testAccProvider.Meta().(ClientSession).ESschemaRegistrySession()
		if err != nil {
			return err
		}

		getLatestSchemaOptions := &schemaregistryv1.GetLatestSchemaOptions{}
		id := rs.Primary.ID
		if id == "" {
			return fmt.Errorf("id is not set")
		}
		if schemaID != "" {
			if !strings.HasSuffix(id, schemaID) {
				return fmt.Errorf("Invalid id: %s and schemaID %s", id, schemaID)
			}
		}

		schemaID = getSchemaID(id)
		getLatestSchemaOptions.SetID(schemaID)

		avroSchema, _, err := schemaregistryClient.GetLatestSchema(getLatestSchemaOptions)
		if err != nil {
			return err
		}

		obj = avroSchema
		return nil
	}
}

func testAccCheckIBMEventStreamsSchemaDestroy(s *terraform.State) error {
	schemaregistryClient, err := testAccProvider.Meta().(ClientSession).ESschemaRegistrySession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_event_streams_schema" {
			continue
		}

		getLatestSchemaOptions := &schemaregistryv1.GetLatestSchemaOptions{}

		getLatestSchemaOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := schemaregistryClient.GetLatestSchema(getLatestSchemaOptions)

		if err == nil {
			return fmt.Errorf("event_streams_schema still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for event_streams_schema (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

var (
	schemaID = "tf-schema"
	id       = "crn:v1:staging:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:c822a30e-bfff-4867-85ec-b805eeab1835:schema:tf-schema"
)

func TestGetUniqueSchemaID(t *testing.T) {
	gotID := getUniqueSchemaID(instanceCRN, schemaID)
	assert.Equal(t, id, gotID)
}

func TestGetSchemaID(t *testing.T) {
	gotSchemaID := getSchemaID(id)
	assert.Equal(t, schemaID, gotSchemaID)
}

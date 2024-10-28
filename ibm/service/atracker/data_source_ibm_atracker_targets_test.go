// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package atracker_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMAtrackerTargetsDataSourceBasic(t *testing.T) {
	targetName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	targetTargetType := "cloud_object_storage"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAtrackerTargetsDataSourceConfigBasic(targetName, targetTargetType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets", "targets.#"),
					resource.TestCheckResourceAttr("data.ibm_atracker_targets.atracker_targets", "targets.0.name", targetName),
					resource.TestCheckResourceAttr("data.ibm_atracker_targets.atracker_targets", "targets.0.target_type", targetTargetType),
				),
			},
		},
	})
}

func TestAccIBMAtrackerTargetsDataSourceAllArgs(t *testing.T) {
	targetName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	targetTargetType := "cloud_object_storage"
	targetRegion := "us-south"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAtrackerTargetsDataSourceConfig(targetName, targetTargetType, targetRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets", "targets.#"),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets", "targets.0.id"),
					resource.TestCheckResourceAttr("data.ibm_atracker_targets.atracker_targets", "targets.0.name", targetName),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets", "targets.0.crn"),
					resource.TestCheckResourceAttr("data.ibm_atracker_targets.atracker_targets", "targets.0.target_type", targetTargetType),
					resource.TestCheckResourceAttr("data.ibm_atracker_targets.atracker_targets", "targets.0.region", targetRegion),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets", "targets.0.encryption_key"),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets", "targets.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets", "targets.0.updated_at"),
					resource.TestCheckResourceAttr("data.ibm_atracker_targets.atracker_targets", "targets.0.api_version", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMAtrackerTargetsDataSourceConfigBasic(targetName string, targetTargetType string) string {
	return fmt.Sprintf(`
		resource "ibm_atracker_target" "atracker_target" {
			name = "%s"
			target_type = "%s"
			cos_endpoint {
				endpoint = "s3.private.us-east.cloud-object-storage.appdomain.cloud"
				target_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
				bucket = "my-atracker-bucket"
				api_key = "%s" // pragma: allowlist secret
				service_to_service_enabled = true
			}
		}

		data "ibm_atracker_targets" "atracker_targets" {
			name = ibm_atracker_target.atracker_target.name
		}
	`, targetName, targetTargetType, acc.COSApiKey)
}

func testAccCheckIBMAtrackerTargetsDataSourceConfig(targetName string, targetTargetType string, targetRegion string) string {
	return fmt.Sprintf(`
		resource "ibm_atracker_target" "atracker_target" {
			name = "%s"
			target_type = "%s"
			region = "%s"
			cos_endpoint {
				endpoint = "s3.private.us-east.cloud-object-storage.appdomain.cloud"
				target_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
				bucket = "my-atracker-bucket"
				api_key = "%s" // pragma: allowlist secret
				service_to_service_enabled = true
			}
			logdna_endpoint {
				target_crn = "crn:v1:bluemix:public:logdna:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
				ingestion_key = "%s"
			}
			eventstreams_endpoint {
				target_crn = "crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
				brokers = [ "kafka-x:9094" ]
				topic = "my-topic"
				api_key = "%s" // pragma: allowlist secret
				service_to_service_enabled = false
			}
			cloudlogs_endpoint {
				target_crn = "crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
			}
		}
		data "ibm_atracker_targets" "atracker_targets" {
			name = ibm_atracker_target.atracker_target.name
		}
		
	`, targetName, targetTargetType, targetRegion, acc.COSApiKey, acc.IngestionKey, acc.IesApiKey)
}

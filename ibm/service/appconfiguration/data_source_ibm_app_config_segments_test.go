// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigSegmentsDataSourceBasic(t *testing.T) {

	tags := "development segment"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_app_config_test_%d", acctest.RandIntRange(10, 100))
	segmentID := fmt.Sprintf("tf_segment_id_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigSegmentsDataSourceConfigBasic(instanceName, name, segmentID, description, tags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments_data2", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments_data2", "total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments_data2", "segments.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments_data2", "segments.0.segment_id"),
					resource.TestCheckResourceAttr("data.ibm_app_config_segments.app_config_segments_data2", "segments.0.name", name),
					resource.TestCheckResourceAttr("data.ibm_app_config_segments.app_config_segments_data2", "segments.0.segment_id", segmentID),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigSegmentsDataSourceConfigBasic(instanceName, name, segmentID, description, tags string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "app_config_terraform_test487" {
			name     = "%s"
			location = "us-south"
			service  = "apprapp"
			plan     = "lite"
		}

		resource "ibm_app_config_segment" "app_config_segment_resource2" {
			guid           		= ibm_resource_instance.app_config_terraform_test487.guid
            name            	= "%s"
			segment_id     	    = "%s"
			rules {
				attribute_name	= "countary"
				operator 		= "contains"
				values 			= ["india", "UK"]
			}
			description    	    = "%s"
		}

		data "ibm_app_config_segments" "app_config_segments_data2" {
			guid				= ibm_app_config_segment.app_config_segment_resource2.guid
		}
		`, instanceName, name, segmentID, description)
}

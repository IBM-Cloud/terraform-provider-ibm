// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigSegmentsDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	segmentName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	segmentSegmentID := fmt.Sprintf("tf_segment_id_%d", acctest.RandIntRange(10, 100))
	segmentDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	segmentTags := fmt.Sprintf("tf_tags_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigSegmentsDataSourceConfigBasic(name, segmentName, segmentSegmentID, segmentDescription, segmentTags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments_data3", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments_data3", "total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments_data3", "segments.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments_data3", "segments.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments_data3", "segments.0.segment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments_data3", "segments.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments_data3", "segments.0.tags"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments_data3", "segments.0.created_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments_data3", "segments.0.updated_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments_data3", "segments.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments_data3", "last.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments_data3", "first.#"),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigSegmentsDataSourceConfigBasic(name, segmentName, segmentSegmentID, segmentDescription, segmentTags string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "app_config_terraform_test49"{
		name     = "%s"
		location = "us-south"
		service  = "apprapp"
		plan     = "standard"
	}
	resource "ibm_app_config_segment" "app_config_segment_resource3" {
		name 					= "%s"
		segment_id 		= "%s"
		description 	= "%s"
		tags 					= "%s"
		rules {
			attribute_name 	= "email"
			operator 				= "endsWith"
			values 					= ["@in.mnc.com"]
		}
		guid 					= ibm_resource_instance.app_config_terraform_test49.guid
	}

	data "ibm_app_config_segments" "app_config_segments_data3" {
		expand		 = true
		guid 			 = ibm_app_config_segment.app_config_segment_resource3.guid
	}
	`, name, segmentName, segmentSegmentID, segmentDescription, segmentTags)
}

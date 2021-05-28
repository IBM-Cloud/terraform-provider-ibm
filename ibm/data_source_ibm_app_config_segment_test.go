// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigSegmentDataSourceBasic(t *testing.T) {
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
				Config: testAccCheckIbmAppConfigSegmentDataSourceConfig(name, segmentName, segmentSegmentID, segmentDescription, segmentTags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segment.app_config_segment_data1", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segment.app_config_segment_data1", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segment.app_config_segment_data1", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segment.app_config_segment_data1", "tags"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segment.app_config_segment_data1", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segment.app_config_segment_data1", "created_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segment.app_config_segment_data1", "updated_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segment.app_config_segment_data1", "rules.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segment.app_config_segment_data1", "rules.0.operator"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segment.app_config_segment_data1", "rules.0.attribute_name"),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigSegmentDataSourceConfig(name, segmentName, segmentSegmentID, segmentDescription, segmentTags string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "app_config_terraform_test47"{
		name     = "%s"
		location = "us-south"
		service  = "apprapp"
		plan     = "standard"
	}
	resource "ibm_app_config_segment" "app_config_segment_resource2" {
		name 					= "%s"
		segment_id 		= "%s"
		description 	= "%s"
		tags 					= "%s"
		rules {
			attribute_name 	= "email"
			operator 				= "endsWith"
			values 					= ["@in.mnc.com"]
		}
		guid 					= ibm_resource_instance.app_config_terraform_test47.guid
	}

	data "ibm_app_config_segment" "app_config_segment_data1" {
		guid 			 = ibm_app_config_segment.app_config_segment_resource2.guid
		segment_id = ibm_app_config_segment.app_config_segment_resource2.segment_id
	}
	`, name, segmentName, segmentSegmentID, segmentDescription, segmentTags)
}

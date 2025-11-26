// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

func TestAccIbmAppConfigSegmentDataSource(t *testing.T) {
	instanceName := fmt.Sprintf("tf_app_config_test_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	segmentID := fmt.Sprintf("tf_segment_id_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigSegmentDataSourceConfigBasic(instanceName, name, segmentID, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segment.ibm_app_config_segment_data1", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segment.ibm_app_config_segment_data1", "segment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segment.ibm_app_config_segment_data1", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segment.ibm_app_config_segment_data1", "created_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segment.ibm_app_config_segment_data1", "updated_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segment.ibm_app_config_segment_data1", "href"),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigSegmentDataSourceConfigBasic(instanceName, name, segmentID, description string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "app_config_terraform_test456" {
			name     = "%s"
			location = "us-south"
			service  = "apprapp"
			plan     = "lite"
		}
		resource "ibm_app_config_segment" "app_config_segment_resource1" {
			guid           	    = ibm_resource_instance.app_config_terraform_test456.guid
			name            	= "%s"
			segment_id     	    = "%s"
			rules {
				attribute_name	= "countary"
				operator		 = "contains"
				values 			 = ["india", "UK"]
			}	
			description    	    = "%s"
		}
		
		data "ibm_app_config_segment" "ibm_app_config_segment_data1" {
			guid          = ibm_app_config_segment.app_config_segment_resource1.guid
			segment_id    = ibm_app_config_segment.app_config_segment_resource1.segment_id
		}
		`, instanceName, name, segmentID, description)
}

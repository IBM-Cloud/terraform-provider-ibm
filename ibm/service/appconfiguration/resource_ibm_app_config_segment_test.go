package appconfiguration_test

import (
	"fmt"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func TestAccIbmIbmAppConfigSegmentBasic(t *testing.T) {
	var conf appconfigurationv1.Segment
	instanceName := fmt.Sprintf("tf_app_config_test_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	segmentID := fmt.Sprintf("tf_segment_id_%d", acctest.RandIntRange(10, 100))

	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmAppConfigSegmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigSegmentConfigBasic(instanceName, name, segmentID, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmAppConfigSegmentExists("ibm_app_config_segment.ibm_app_config_segment_resource1", conf),
					resource.TestCheckResourceAttrSet("ibm_app_config_segment.ibm_app_config_segment_resource1", "id"),
					resource.TestCheckResourceAttrSet("ibm_app_config_segment.ibm_app_config_segment_resource1", "name"),
					resource.TestCheckResourceAttrSet("ibm_app_config_segment.ibm_app_config_segment_resource1", "segment_id"),
					resource.TestCheckResourceAttrSet("ibm_app_config_segment.ibm_app_config_segment_resource1", "description"),
				),
			},
			{
				Config: testAccCheckIbmAppConfigSegmentConfigBasic(instanceName, nameUpdate, segmentID, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_app_config_segment.ibm_app_config_segment_resource1", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_app_config_segment.ibm_app_config_segment_resource1", "description", descriptionUpdate),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigSegmentConfigBasic(name, envName, segmentID, description string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "app_config_terraform_test456" {
			name     = "%s"
			location = "us-south"
			service  = "apprapp"
			plan     = "lite"
		}
		resource "ibm_app_config_segment" "ibm_app_config_segment_resource1" {
			guid           	    = ibm_resource_instance.app_config_terraform_test456.guid
			name            	= "%s"
			segment_id     	    = "%s"
			rules {
				attribute_name	= "countary"
				operator = "contains"
				values = ["india", "UK"]
			}	
			description    	    = "%s"

		}`, name, envName, segmentID, description)
}

func testAccCheckIbmAppConfigSegmentExists(n string, obj appconfigurationv1.Segment) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return flex.FmtErrorf("Not found: %s", n)
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return flex.FmtErrorf(fmt.Sprintf("%s", err))
		}
		appconfigClient, err := getAppConfigClient(acc.TestAccProvider.Meta(), parts[0])
		if err != nil {
			return flex.FmtErrorf(fmt.Sprintf("%s", err))
		}

		options := &appconfigurationv1.GetSegmentOptions{}

		options.SetSegmentID(parts[1])

		result, _, err := appconfigClient.GetSegment(options)
		if err != nil {
			return flex.FmtErrorf(fmt.Sprintf("%s", err))
		}

		obj = *result
		return nil
	}
}

func testAccCheckIbmAppConfigSegmentDestroy(s *terraform.State) error {

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_app_config_segment_resource1" {
			continue
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return flex.FmtErrorf(fmt.Sprintf("%s", err))
		}
		appconfigClient, err := getAppConfigClient(acc.TestAccProvider.Meta(), parts[0])
		if err != nil {
			return flex.FmtErrorf(fmt.Sprintf("%s", err))
		}
		options := &appconfigurationv1.GetSegmentOptions{}

		options.SetSegmentID(parts[1])

		// Try to find the key
		_, response, err := appconfigClient.GetSegment(options)

		if err == nil {
			return flex.FmtErrorf("Segment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return flex.FmtErrorf("[ERROR] Error checking for Segment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

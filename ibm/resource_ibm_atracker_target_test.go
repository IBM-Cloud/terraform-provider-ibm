// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/atrackerv1"
)

func TestAccIBMAtrackerTargetBasic(t *testing.T) {
	var conf atrackerv1.Target
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	targetType := "cloud_object_storage"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	targetTypeUpdate := "cloud_object_storage"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMAtrackerTargetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAtrackerTargetConfigBasic(name, targetType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAtrackerTargetExists("ibm_atracker_target.atracker_target", conf),
					resource.TestCheckResourceAttr("ibm_atracker_target.atracker_target", "name", name),
					resource.TestCheckResourceAttr("ibm_atracker_target.atracker_target", "target_type", targetType),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMAtrackerTargetConfigBasic(nameUpdate, targetTypeUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_atracker_target.atracker_target", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_atracker_target.atracker_target", "target_type", targetTypeUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_atracker_target.atracker_target",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMAtrackerTargetConfigBasic(name string, targetType string) string {
	return fmt.Sprintf(`

		resource "ibm_atracker_target" "atracker_target" {
			name = "%s"
			target_type = "%s"
			cos_endpoint {
				endpoint = "s3.private.us-east.cloud-object-storage.appdomain.cloud"
				target_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
				bucket = "my-atracker-bucket"
				api_key = "xxxxxxxxxxxxxx"
			}
		}
	`, name, targetType)
}

func testAccCheckIBMAtrackerTargetExists(n string, obj atrackerv1.Target) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		atrackerClient, err := testAccProvider.Meta().(ClientSession).AtrackerV1()
		if err != nil {
			return err
		}

		getTargetOptions := &atrackerv1.GetTargetOptions{}

		getTargetOptions.SetID(rs.Primary.ID)

		target, _, err := atrackerClient.GetTarget(getTargetOptions)
		if err != nil {
			return err
		}

		obj = *target
		return nil
	}
}

func testAccCheckIBMAtrackerTargetDestroy(s *terraform.State) error {
	atrackerClient, err := testAccProvider.Meta().(ClientSession).AtrackerV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_atracker_target" {
			continue
		}

		getTargetOptions := &atrackerv1.GetTargetOptions{}

		getTargetOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := atrackerClient.GetTarget(getTargetOptions)

		if err == nil {
			return fmt.Errorf("ATracker Target still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for ATracker Target (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

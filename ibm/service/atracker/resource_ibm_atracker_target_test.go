// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package atracker_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/atrackerv2"
)

func TestAccIBMAtrackerTargetBasic(t *testing.T) {
	var conf atrackerv2.Target
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	targetType := "cloud_object_storage"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAtrackerTargetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAtrackerTargetConfigBasic(name, targetType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAtrackerTargetExists("ibm_atracker_target.atracker_target", conf),
					resource.TestCheckResourceAttr("ibm_atracker_target.atracker_target", "name", name),
					resource.TestCheckResourceAttr("ibm_atracker_target.atracker_target", "target_type", targetType),
				),
			},
			{
				Config: testAccCheckIBMAtrackerTargetConfigBasic(nameUpdate, targetType),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_atracker_target.atracker_target", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_atracker_target.atracker_target", "target_type", targetType),
				),
			},
			{
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
				service_to_service_enabled = false
			}
		}
	`, name, targetType)
}

func testAccCheckIBMAtrackerTargetExists(n string, obj atrackerv2.Target) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		atrackerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AtrackerV2()
		if err != nil {
			return err
		}

		getTargetOptions := &atrackerv2.GetTargetOptions{}

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
	atrackerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AtrackerV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_atracker_target" {
			continue
		}

		getTargetOptions := &atrackerv2.GetTargetOptions{}

		getTargetOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := atrackerClient.GetTarget(getTargetOptions)

		if err == nil {
			return fmt.Errorf("Activity Tracker Target still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for Activity Tracker Target (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

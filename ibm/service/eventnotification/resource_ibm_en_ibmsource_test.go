// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	en "github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
)

func TestAccIBMEnIBMSourceAllArgs(t *testing.T) {
	var config en.Source
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	sourceId := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	enabled := true
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnIBMSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnIBMSourceConfig(instanceName, sourceId, enabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnIBMSourceExists("ibm_en_source.en_source_resource_1", config),
					resource.TestCheckResourceAttr("ibm_en_source.en_source_resource_1", "enabled", "enabled"),
				),
			},
			{
				Config: testAccCheckIBMEnIBMSourceConfig(instanceName, sourceId, enabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_source.en_source_resource_1", "enabled", "enabled"),
				),
			},
			{
				ResourceName:      "ibm_en_source.en_source_resource_1",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMEnIBMSourceConfig(instanceName, sourceId string, enabled bool) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_source_resource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_source" "en_source_resource_1" {
		instance_guid = ibm_resource_instance.en_source_resource.guid
		source_id     = "%s"
		enabled = %t
	}
	`, instanceName, sourceId, enabled)
}

func testAccCheckIBMEnIBMSourceExists(n string, obj en.Source) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		enClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EventNotificationsApiV1()
		if err != nil {
			return err
		}

		options := &en.GetSourceOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		options.SetInstanceID(parts[0])
		options.SetID(parts[1])

		result, _, err := enClient.GetSource(options)
		if err != nil {
			return err
		}

		obj = *result
		return nil
	}
}

func testAccCheckIBMEnIBMSourceDestroy(s *terraform.State) error {

	return nil
}

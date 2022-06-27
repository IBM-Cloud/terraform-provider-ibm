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

func TestAccIBMEnSafariDestinationAllArgs(t *testing.T) {
	var config en.Destination
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	newName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	newDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnSafariDestinationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnSafariDestinationConfig(instanceName, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnSafariDestinationExists("ibm_en_destination_safari.en_destination_resource_safari", config),
					resource.TestCheckResourceAttr("ibm_en_destination_safari.en_destination_resource_safari", "name", name),
					resource.TestCheckResourceAttr("ibm_en_destination_safari.en_destination_resource_safari", "type", "push_safari"),
					resource.TestCheckResourceAttr("ibm_en_destination_safari.en_destination_resource_safari", "description", description),
				),
			},
			{
				Config: testAccCheckIBMEnSafariDestinationConfig(instanceName, newName, newDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_destination_safari.en_destination_resource_safari", "name", newName),
					resource.TestCheckResourceAttr("ibm_en_destination_safari.en_destination_resource_safari", "type", "push_safari"),
					resource.TestCheckResourceAttr("ibm_en_destination_safari.en_destination_resource_safari", "description", newDescription),
				),
			},
			{
				ResourceName:      "ibm_en_destination_safari.en_destination_resource_safari",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMEnSafariDestinationConfig(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_destination_resource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_destination_safari" "en_destination_resource_safari" {
		instance_guid = ibm_resource_instance.en_destination_resource.guid
		name                         = "%s"
		type                         = "push_safari"
        certificate                  = "${path.module}/cert.p12"
		icon_16x16                   = "${path.module}/safariicon.png"
		icon_16x16_2x                = "${path.module}/safariicon.png"
		icon_32x32                   = "${path.module}/safariicon.png"
		icon_32x32_2x                = "${path.module}/safariicon.png"
		icon_128x128                 = "${path.module}/safariicon.png"
		icon_128x128_2x              = "${path.module}/safariicon.png"
		icon_16x16_content_type      = "png"
		icon_16x16_2x_content_type   = "png"
		icon_32x32_content_type      = "png"
		icon_32x32_2x_content_type   = "png"
		icon_128x128_content_type    = "png"
		icon_128x128_2x_content_type = "png"
		description = "%s"
		config {
			params {
				cert_type = "p12"
                password = "certpassword"
				website_name = "NodeJS Starter Application"
                url_format_string = "https://ensafaripush.mybluemix.net"
                website_push_id = "web.net.mybluemix.ensafaripush"
                website_url = "https://ensafaripush.mybluemix.net"
			}
		}
	}
	`, instanceName, name, description)
}

func testAccCheckIBMEnSafariDestinationExists(n string, obj en.Destination) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		enClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EventNotificationsApiV1()
		if err != nil {
			return err
		}

		options := &en.GetDestinationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		options.SetInstanceID(parts[0])
		options.SetID(parts[1])

		result, _, err := enClient.GetDestination(options)
		if err != nil {
			return err
		}

		obj = *result
		return nil
	}
}

func testAccCheckIBMEnSafariDestinationDestroy(s *terraform.State) error {
	enClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "en_destination_resource_safari" {
			continue
		}

		options := &en.GetDestinationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		options.SetInstanceID(parts[0])
		options.SetID(parts[1])

		// Try to find the key
		_, response, err := enClient.GetDestination(options)

		if err == nil {
			return fmt.Errorf("en_destination still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for en_destination (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

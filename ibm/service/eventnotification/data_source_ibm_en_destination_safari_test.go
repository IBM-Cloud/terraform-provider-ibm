// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMEnSafariDestinationDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnSafariDestinationDataSourceConfigBasic(instanceName, name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_safari.en_destination_data_6", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_safari.en_destination_data_6", "instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_safari.en_destination_data_6", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_safari.en_destination_data_6", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_safari.en_destination_data_6", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_safari.en_destination_data_6", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_safari.en_destination_data_6", "destination_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_destination_safari.en_destination_data_6", "subscription_count"),
				),
			},
		},
	})
}

func testAccCheckIBMEnSafariDestinationDataSourceConfigBasic(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_destination_datasource2" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_destination_safari" "en_destination_datasource_4" {
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

		data "ibm_en_destination_safari" "en_destination_data_6" {
			instance_guid = ibm_resource_instance.en_destination_datasource2.guid
			destination_id = ibm_en_destination_safari.en_destination_datasource_4.destination_id
		}
	`, instanceName, name, description)
}

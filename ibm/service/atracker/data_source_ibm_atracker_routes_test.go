// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package atracker_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMAtrackerRoutesDataSourceBasic(t *testing.T) {
	routeName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAtrackerRoutesDataSourceConfigBasic(routeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_atracker_routes.atracker_routes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_routes.atracker_routes", "routes.#"),
					resource.TestCheckResourceAttr("data.ibm_atracker_routes.atracker_routes", "routes.0.name", routeName),
				),
			},
		},
	})
}

func testAccCheckIBMAtrackerRoutesDataSourceConfigBasic(routeName string) string {
	return fmt.Sprintf(`
		resource "ibm_atracker_target" "atracker_target" {
			name = "my-cos-target"
			target_type = "cloud_object_storage"
			cos_endpoint {
				endpoint = "s3.private.us-east.cloud-object-storage.appdomain.cloud"
				target_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
				bucket = "my-atracker-bucket"
				api_key = "xxxxxxxxxxxxxx"
			}
		}

		resource "ibm_atracker_route" "atracker_route" {
			name = "%s"
			rules {
				target_ids = [ ibm_atracker_target.atracker_target.id ]
				locations = [ "us-south" ]
			}
		}

		data "ibm_atracker_routes" "atracker_routes" {
			name = ibm_atracker_route.atracker_route.name
		}
	`, routeName)
}

// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAtrackerRoutesDataSourceBasic(t *testing.T) {
	routeName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	routeReceiveGlobalEvents := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAtrackerRoutesDataSourceConfigBasic(routeName, routeReceiveGlobalEvents),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_atracker_routes.atracker_routes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_routes.atracker_routes", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_routes.atracker_routes", "routes.#"),
					resource.TestCheckResourceAttr("data.ibm_atracker_routes.atracker_routes", "routes.0.name", routeName),
					resource.TestCheckResourceAttr("data.ibm_atracker_routes.atracker_routes", "routes.0.receive_global_events", routeReceiveGlobalEvents),
				),
			},
		},
	})
}

func testAccCheckIBMAtrackerRoutesDataSourceConfigBasic(routeName string, routeReceiveGlobalEvents string) string {
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
			receive_global_events = %s
			rules {
				target_ids = [ ibm_atracker_target.atracker_target.id ]
			}
		}

		data "ibm_atracker_routes" "atracker_routes" {
			name = ibm_atracker_route.atracker_route.name
		}
	`, routeName, routeReceiveGlobalEvents)
}

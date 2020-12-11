/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccIBMAtrackerRoutesDataSourceAllArgs(t *testing.T) {
	name := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAtrackerRoutesDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_atracker_routes.atracker_routes", "id"),
					resource.TestCheckResourceAttr("data.ibm_atracker_routes.atracker_routes", "name", name),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_routes.atracker_routes", "routes.#"),
				),
			},
		},
	})

}

func testAccCheckIBMAtrackerRoutesDataSourceConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_atracker_target" "atracker_target" {
			name = "my-atracker-target"
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
			receive_global_events = true
			rules {
				target_ids = [ ibm_atracker_target.atracker_target.id ]
			}
		}

		data "ibm_atracker_routes" "atracker_routes" {
			name = ibm_atracker_route.atracker_route.name
		}
	`, name)
}

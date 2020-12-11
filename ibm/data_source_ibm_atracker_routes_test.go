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

func TestAccIBMAtrackerRoutesDataSourceBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:	func() { testAccPreCheck(t) },
		Providers:	testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAtrackerRoutesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_atracker_routes.atracker_routes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_routes.atracker_routes", "routes.#"),
				),
			},
		},
	})
}

func TestAccIBMAtrackerRoutesDataSourceAllArgs(t *testing.T) {
	name := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:	func() { testAccPreCheck(t) },
		Providers:	testAccProviders,
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
		data "ibm_atracker_routes" "atracker_routes" {
			name = "%s"
		}
	`, name)
}

func testAccCheckIBMAtrackerRoutesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_atracker_routes" "atracker_routes" {
		}
	`, )
}


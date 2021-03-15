/**
 * (C) Copyright IBM Corp. 2021.
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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmIsDedicatedHostGroupDataSourceBasic(t *testing.T) {
	class := "beta"
	family := "memory"
	name := fmt.Sprintf("name%d", acctest.RandIntRange(10, 100))
	resName := "data.ibm_is_dedicated_host_group.dgroup"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmIsDedicatedHostGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIsDedicatedHostGroupDataSourceConfigBasic(class, family, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", name),
					resource.TestCheckResourceAttr(resName, "class", class),
					resource.TestCheckResourceAttr(resName, "family", family),
					resource.TestCheckResourceAttrSet(resName, "zone"),
				),
			},
		},
	})
}

func testAccCheckIbmIsDedicatedHostGroupDataSourceConfigBasic(class string, family string, name string) string {
	return fmt.Sprintf(`
	
	data "ibm_resource_group" "default" {
		name = "Default" ///give your resource grp
	}
	resource "ibm_is_dedicated_host_group" "dhgroup" {
		class = "%s"
		family = "%s"
		name = "%s"
		resource_group = data.ibm_resource_group.default.id
		zone = "us-south-2"
	}
	data "ibm_is_dedicated_host_group" "dgroup" {
		name = ibm_is_dedicated_host_group.dhgroup.name
	}
	`, class, family, name)
}

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

func TestAccIbmIsDedicatedHostDSBasic(t *testing.T) {
	//var conf vpcv1.DedicatedHost
	resName := "data.ibm_is_dedicated_host.dhost"
	name := "testdh02"
	class := "beta"
	family := "memory"
	grpname := fmt.Sprintf("name%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmIsDedicatedHostDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIsDedicatedHostDSConfigBasic(class, family, grpname, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "name"),
					resource.TestCheckResourceAttrSet(resName, "zone"),
					resource.TestCheckResourceAttrSet(resName, "host_group"),
				),
			},
		},
	})
}

func testAccCheckIbmIsDedicatedHostDSConfigBasic(class, family, grpname, name string) string {
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

	resource "ibm_is_dedicated_host" "dedicated-host-test-01" {
		profile = "dh2-56x464"
		host_group = data.ibm_is_dedicated_host_group.dgroup.id
		name = "%s"
	  }
	data "ibm_is_dedicated_host" "dhost"{
		name = "%s"
		host_group = ibm_is_dedicated_host.dedicated-host-test-01.host_group
	}
	`, class, family, grpname, name, name)
}

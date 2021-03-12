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
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIbmIsDedicatedHostGroupBasic(t *testing.T) {
	var conf vpcv1.DedicatedHostGroup
	class := "beta"
	family := "memory"
	name := fmt.Sprintf("name%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmIsDedicatedHostGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIsDedicatedHostGroupConfigBasic(class, family, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsDedicatedHostGroupExists("ibm_is_dedicated_host_group.is_dedicated_host_group", conf),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmIsDedicatedHostGroupConfigBasic(class, family, name),
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func TestAccIbmIsDedicatedHostGroupAllArgs(t *testing.T) {
	var conf vpcv1.DedicatedHostGroup
	class := "beta"
	family := "memory"
	name := fmt.Sprintf("name%d", acctest.RandIntRange(10, 100))

	nameUpdate := fmt.Sprintf("name%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmIsDedicatedHostGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIsDedicatedHostGroupConfig(class, family, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsDedicatedHostGroupExists("ibm_is_dedicated_host_group.is_dedicated_host_group", conf),
					resource.TestCheckResourceAttr("ibm_is_dedicated_host_group.is_dedicated_host_group", "class", class),
					resource.TestCheckResourceAttr("ibm_is_dedicated_host_group.is_dedicated_host_group", "family", family),
					resource.TestCheckResourceAttr("ibm_is_dedicated_host_group.is_dedicated_host_group", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmIsDedicatedHostGroupConfig(class, family, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_dedicated_host_group.is_dedicated_host_group", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_dedicated_host_group.is_dedicated_host_group",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmIsDedicatedHostGroupConfigBasic(class string, family string, name string) string {
	return fmt.Sprintf(`

		data "ibm_resource_group" "default" {
			name = "Default" ///give your resource grp
		}
		resource "ibm_is_dedicated_host_group" "is_dedicated_host_group" {
			class = "%s"
			family = "%s"
			name = "%s"
			resource_group = data.ibm_resource_group.default.id
			zone = "us-south-2"
		}
	`, class, family, name)
}

func testAccCheckIbmIsDedicatedHostGroupConfig(class string, family string, name string) string {
	return fmt.Sprintf(`

		data "ibm_resource_group" "default" {
			name = "Default" ///give your resource grp
		}
		resource "ibm_is_dedicated_host_group" "is_dedicated_host_group" {
			class = "%s"
			family = "%s"
			name = "%s"
			resource_group = data.ibm_resource_group.default.id
			zone = "us-south-2"
		}
	`, class, family, name)
}

func testAccCheckIbmIsDedicatedHostGroupExists(n string, obj vpcv1.DedicatedHostGroup) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := testAccProvider.Meta().(ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getDedicatedHostGroupOptions := &vpcv1.GetDedicatedHostGroupOptions{}

		getDedicatedHostGroupOptions.SetID(rs.Primary.ID)

		dedicatedHostGroup, _, err := vpcClient.GetDedicatedHostGroup(getDedicatedHostGroupOptions)
		if err != nil {
			return err
		}

		obj = *dedicatedHostGroup
		return nil
	}
}

func testAccCheckIbmIsDedicatedHostGroupDestroy(s *terraform.State) error {
	vpcClient, err := testAccProvider.Meta().(ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_dedicated_host_group" {
			continue
		}

		getDedicatedHostGroupOptions := &vpcv1.GetDedicatedHostGroupOptions{}

		getDedicatedHostGroupOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := vpcClient.GetDedicatedHostGroup(getDedicatedHostGroupOptions)

		if err == nil {
			return fmt.Errorf("DedicatedHostGroup still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for DedicatedHostGroup (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

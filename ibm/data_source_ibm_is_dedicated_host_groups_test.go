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

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIbmIsDedicatedHostGroupsDataSourceBasic(t *testing.T) {
	var conf vpcv1.DedicatedHostGroup
	class := "beta"
	family := "memory"
	name := fmt.Sprintf("name%d", acctest.RandIntRange(10, 100))
	resName := "data.ibm_is_dedicated_host_groups.dhgroups"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmIsDedicatedHostGroupsDSDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIsDedicatedHostGroupConfigBasic(class, family, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsDedicatedHostGroupExists("ibm_is_dedicated_host_group.is_dedicated_host_group", conf),
					resource.TestCheckResourceAttr("ibm_is_dedicated_host_group.is_dedicated_host_group", "class", class),
					resource.TestCheckResourceAttr("ibm_is_dedicated_host_group.is_dedicated_host_group", "family", family),
					resource.TestCheckResourceAttr("ibm_is_dedicated_host_group.is_dedicated_host_group", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmIsDedicatedHostGroupsDataSourceConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "groups.0.name"),
					resource.TestCheckResourceAttrSet(resName, "groups.0.class"),
					resource.TestCheckResourceAttrSet(resName, "groups.0.family"),
					resource.TestCheckResourceAttrSet(resName, "groups.0.zone"),
				),
			},
		},
	})
}

func testAccCheckIbmIsDedicatedHostGroupsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
	
	data "ibm_is_dedicated_host_groups" "dhgroups" {
	}
	`)
}

func testAccCheckIbmIsDedicatedHostGroupsDSDestroy(s *terraform.State) error {
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

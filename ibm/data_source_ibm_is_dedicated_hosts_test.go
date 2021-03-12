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

func TestAccIbmIsDedicatedHostsDSBasic(t *testing.T) {
	var conf vpcv1.DedicatedHost
	class := "beta"
	family := "memory"
	groupname := fmt.Sprintf("name%d", acctest.RandIntRange(10, 100))
	dhname := fmt.Sprintf("name%d", acctest.RandIntRange(10, 1000))
	profile := "dh2-56x464"
	resName := "data.ibm_is_dedicated_hosts.dhosts"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsDedicatedHostConfigBasic(class, family, groupname, profile, dhname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIbmIsDedicatedHostExists("ibm_is_dedicated_host.dedicated-host-test-01", conf),
					resource.TestCheckResourceAttr(
						"ibm_is_dedicated_host.dedicated-host-test-01", "name", dhname),
				),
			},
			{
				Config: testAccCheckIbmIsDedicatedHostsDSConfigBasic(class, family, groupname, profile, dhname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "dedicated_hosts.0.name"),
					resource.TestCheckResourceAttrSet(resName, "dedicated_hosts.0.available_memory"),
					resource.TestCheckResourceAttrSet(resName, "dedicated_hosts.0.memory"),
					resource.TestCheckResourceAttrSet(resName, "dedicated_hosts.0.host_group"),
					resource.TestCheckResourceAttrSet(resName, "dedicated_hosts.0.zone"),
				),
			},
		},
	})
}

func testAccCheckIbmIsDedicatedHostsDSConfigBasic(class string, family string, groupname string, profile string, dhname string) string {
	return testAccCheckIbmIsDedicatedHostConfigBasic(class, family, groupname, profile, dhname) + fmt.Sprintf(`
	 
	 data "ibm_is_dedicated_hosts" "dhosts"{
	 }
	 `)
}

func testAccCheckIbmIsDedicatedHostsDSExists(n string, obj vpcv1.DedicatedHost) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := testAccProvider.Meta().(ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getDedicatedHostOptions := &vpcv1.GetDedicatedHostOptions{}

		getDedicatedHostOptions.SetID(rs.Primary.ID)

		dedicatedHost, _, err := vpcClient.GetDedicatedHost(getDedicatedHostOptions)
		if err != nil {
			return err
		}

		obj = *dedicatedHost
		return nil
	}
}

func testAccCheckIbmIsDedicatedHostsDSDestroy(s *terraform.State) error {
	vpcClient, err := testAccProvider.Meta().(ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_dedicated_host" {
			continue
		}

		getDedicatedHostOptions := &vpcv1.GetDedicatedHostOptions{}

		getDedicatedHostOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := vpcClient.GetDedicatedHost(getDedicatedHostOptions)

		if err == nil {
			return fmt.Errorf("DedicatedHost still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for DedicatedHost (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

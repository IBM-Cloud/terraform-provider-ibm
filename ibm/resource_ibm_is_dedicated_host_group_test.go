// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

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
	name := fmt.Sprintf("tfdhgroup%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmIsDedicatedHostGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIsDedicatedHostGroupConfigBasic(dedicatedHostGroupClass, dedicatedHostGroupFamily, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsDedicatedHostGroupExists("ibm_is_dedicated_host_group.is_dedicated_host_group", conf),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmIsDedicatedHostGroupConfigBasic(dedicatedHostGroupClass, dedicatedHostGroupFamily, name),
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func TestAccIbmIsDedicatedHostGroupAllArgs(t *testing.T) {
	var conf vpcv1.DedicatedHostGroup
	class := "beta"
	family := "memory"
	name := fmt.Sprintf("tfdhgroup%d", acctest.RandIntRange(10, 100))

	nameUpdate := fmt.Sprintf("tfdhgroup%d", acctest.RandIntRange(10, 1000))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmIsDedicatedHostGroupDestroy,
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
				Config: testAccCheckIbmIsDedicatedHostGroupConfigBasic(class, family, nameUpdate),
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
			is_default=true
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

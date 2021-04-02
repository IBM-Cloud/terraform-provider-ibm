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

func TestAccIbmIsDedicatedHostBasic(t *testing.T) {
	var conf vpcv1.DedicatedHost
	class := "beta"
	family := "memory"
	groupname := fmt.Sprintf("tfdhost%d", acctest.RandIntRange(10, 100))
	dhname := "testdh02"

	resname := "ibm_is_dedicated_host.dedicated-host-test-01"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmIsDedicatedHostDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmIsDedicatedHostConfigBasic(class, family, groupname, dedicatedHostProfileName, dhname),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsDedicatedHostExists(resname, conf),
					resource.TestCheckResourceAttr(resname, "name", dhname),
				),
			},
			resource.TestStep{
				ResourceName:      resname,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmIsDedicatedHostConfigBasic(class string, family string, groupname string, profile string, dhname string) string {
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

	resource "ibm_is_dedicated_host" "dedicated-host-test-01" {
		profile = "%s"
		host_group = ibm_is_dedicated_host_group.is_dedicated_host_group.id
		name = "%s"
	  }
	`, class, family, groupname, profile, dhname)
}

func testAccCheckIbmIsDedicatedHostExists(n string, obj vpcv1.DedicatedHost) resource.TestCheckFunc {

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

func testAccCheckIbmIsDedicatedHostDestroy(s *terraform.State) error {
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

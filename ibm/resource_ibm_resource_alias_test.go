// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

func TestAccIbmResourceAliasBasic(t *testing.T) {
	var conf resourcecontrollerv2.ResourceAlias
	name := fmt.Sprintf("tf_alias_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_alias_nameUpdate_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckResourceAliasBinding(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmResourceAliasDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmResourceAliasConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmResourceAliasExists("ibm_resource_alias.resource_alias", conf),
					resource.TestCheckResourceAttr("ibm_resource_alias.resource_alias", "name", name),
					resource.TestCheckResourceAttrSet("ibm_resource_alias.resource_alias", "source"),
					resource.TestCheckResourceAttrSet("ibm_resource_alias.resource_alias", "target"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmResourceAliasConfig(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_resource_alias.resource_alias", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_resource_alias.resource_alias",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmResourceAliasConfig(name string) string {
	return fmt.Sprintf(`
		data "ibm_org" "org" {
			name = "%s"
		}
		data "ibm_space" "space" {
			org = data.ibm_org.org.name
			name = "%s"
		}
		resource "ibm_resource_instance" "instance" {
			name     = "tf_test_service"
			service  = "cloud-object-storage"
			plan     = "lite"
			location = "global"
		}
		resource "ibm_resource_alias" "resource_alias" {
			name = "%s"
			source = ibm_resource_instance.instance.guid
			target = "crn:v1:bluemix:public:cf:us-south:o/${data.ibm_org.org.id}::cf-space:${data.ibm_space.space.id}"
		}
	`, cfOrganization, cfSpace, name)
}

func testAccCheckIbmResourceAliasExists(n string, obj resourcecontrollerv2.ResourceAlias) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		resourceControllerClient, err := testAccProvider.Meta().(ClientSession).ResourceControllerV2API()
		if err != nil {
			return err
		}

		getResourceAliasOptions := &resourcecontrollerv2.GetResourceAliasOptions{}

		getResourceAliasOptions.SetID(rs.Primary.ID)

		resourceAlias, _, err := resourceControllerClient.GetResourceAlias(getResourceAliasOptions)
		if err != nil {
			return err
		}

		obj = *resourceAlias
		return nil
	}
}

func testAccCheckIbmResourceAliasDestroy(s *terraform.State) error {
	resourceControllerClient, err := testAccProvider.Meta().(ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_resource_alias" {
			continue
		}

		getResourceAliasOptions := &resourcecontrollerv2.GetResourceAliasOptions{}

		getResourceAliasOptions.SetID(rs.Primary.ID)

		// Try to find the key
		alias, resp, err := resourceControllerClient.GetResourceAlias(getResourceAliasOptions)

		if err == nil {
			if *alias.State == "removed" {
				return nil
			}
			return fmt.Errorf("Resource alias still exists: %s with resp code: %s", rs.Primary.ID, resp)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for resource key (%s) to be destroyed: %s with resp code: %s", rs.Primary.ID, err, resp)
		}
	}

	return nil
}

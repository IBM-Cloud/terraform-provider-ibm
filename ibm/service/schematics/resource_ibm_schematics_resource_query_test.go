// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func TestAccIBMSchematicsResourceQueryBasic(t *testing.T) {
	var conf schematicsv1.ResourceQueryRecord

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsResourceQueryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSchematicsResourceQueryConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsResourceQueryExists("ibm_schematics_resource_query.schematics_resource_query", conf),
				),
			},
		},
	})
}

func TestAccIBMSchematicsResourceQueryAllArgs(t *testing.T) {
	var conf schematicsv1.ResourceQueryRecord
	typeVar := "vsi"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := "vsi"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsResourceQueryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSchematicsResourceQueryConfig(typeVar, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsResourceQueryExists("ibm_schematics_resource_query.schematics_resource_query", conf),
					resource.TestCheckResourceAttr("ibm_schematics_resource_query.schematics_resource_query", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_schematics_resource_query.schematics_resource_query", "name", name),
				),
			},
			{
				Config: testAccCheckIBMSchematicsResourceQueryConfig(typeVarUpdate, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_schematics_resource_query.schematics_resource_query", "type", typeVarUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_resource_query.schematics_resource_query", "name", nameUpdate),
				),
			},
			{
				ResourceName:      "ibm_schematics_resource_query.schematics_resource_query",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSchematicsResourceQueryConfigBasic() string {
	return `

		resource "ibm_schematics_resource_query" "schematics_resource_query" {
		}
	`
}

func testAccCheckIBMSchematicsResourceQueryConfig(typeVar string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_resource_query" "schematics_resource_query" {
			type = "%s"
			name = "%s"
			queries {
				query_type = "workspaces"
				query_condition {
					name = "name"
					value = "value"
					description = "description"
				}
				query_select = [ "query_select" ]
			}
		}
	`, typeVar, name)
}

func testAccCheckIBMSchematicsResourceQueryExists(n string, obj schematicsv1.ResourceQueryRecord) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
		if err != nil {
			return err
		}

		getResourcesQueryOptions := &schematicsv1.GetResourcesQueryOptions{}

		getResourcesQueryOptions.SetQueryID(rs.Primary.ID)

		resourceQueryRecord, _, err := schematicsClient.GetResourcesQuery(getResourcesQueryOptions)
		if err != nil {
			return err
		}

		obj = *resourceQueryRecord
		return nil
	}
}

func testAccCheckIBMSchematicsResourceQueryDestroy(s *terraform.State) error {
	schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_schematics_resource_query" {
			continue
		}

		getResourcesQueryOptions := &schematicsv1.GetResourcesQueryOptions{}

		getResourcesQueryOptions.SetQueryID(rs.Primary.ID)

		// Try to find the key
		_, response, err := schematicsClient.GetResourcesQuery(getResourcesQueryOptions)

		if err == nil {
			return fmt.Errorf("schematics_resource_query still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for schematics_resource_query (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

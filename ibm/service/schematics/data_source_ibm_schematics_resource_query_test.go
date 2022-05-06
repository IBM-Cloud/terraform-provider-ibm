// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMSchematicsResourceQueryDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsResourceQueryDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_resource_query.schematics_resource_query", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_resource_query.schematics_resource_query", "query_id"),
				),
			},
		},
	})
}

func TestAccIBMSchematicsResourceQueryDataSourceAllArgs(t *testing.T) {
	resourceQueryRecordType := "vsi"
	resourceQueryRecordName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsResourceQueryDataSourceConfig(resourceQueryRecordType, resourceQueryRecordName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_resource_query.schematics_resource_query", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_resource_query.schematics_resource_query", "query_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_resource_query.schematics_resource_query", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_resource_query.schematics_resource_query", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_resource_query.schematics_resource_query", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_resource_query.schematics_resource_query", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_resource_query.schematics_resource_query", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_resource_query.schematics_resource_query", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_resource_query.schematics_resource_query", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_resource_query.schematics_resource_query", "queries.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_resource_query.schematics_resource_query", "queries.0.query_type"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsResourceQueryDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_schematics_resource_query" "schematics_resource_query" {
		}

		data "ibm_schematics_resource_query" "schematics_resource_query" {
			query_id = "query_id"
		}
	`)
}

func testAccCheckIBMSchematicsResourceQueryDataSourceConfig(resourceQueryRecordType string, resourceQueryRecordName string) string {
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

		data "ibm_schematics_resource_query" "schematics_resource_query" {
			query_id = "query_id"
		}
	`, resourceQueryRecordType, resourceQueryRecordName)
}

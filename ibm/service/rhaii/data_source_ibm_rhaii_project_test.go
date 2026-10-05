// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMRhaiiProjectDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-rhaii-%s", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRhaiiProjectDataSourceConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_rhaii_project.by_name", "name", name),
					resource.TestCheckResourceAttr("data.ibm_rhaii_project.by_name", "service", "instructlab"),
					resource.TestCheckResourceAttrPair("data.ibm_rhaii_project.by_name", "project_id", "ibm_rhaii_project.project", "project_id"),
					resource.TestCheckResourceAttrPair("data.ibm_rhaii_project.by_name", "endpoint", "ibm_rhaii_project.project", "endpoint"),
					resource.TestCheckResourceAttrPair("data.ibm_rhaii_project.by_id", "name", "ibm_rhaii_project.project", "name"),
					resource.TestCheckResourceAttrPair("data.ibm_rhaii_project.by_id", "crn", "ibm_rhaii_project.project", "crn"),
				),
			},
		},
	})
}

func testAccCheckIBMRhaiiProjectDataSourceConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_rhaii_project" "project" {
		name = "%s"
	}

	data "ibm_rhaii_project" "by_name" {
		name              = ibm_rhaii_project.project.name
		resource_group_id = ibm_rhaii_project.project.resource_group_id
	}

	data "ibm_rhaii_project" "by_id" {
		identifier = ibm_rhaii_project.project.project_id
	}
	`, name)
}

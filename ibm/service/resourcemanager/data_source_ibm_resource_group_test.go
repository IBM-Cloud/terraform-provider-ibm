// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package resourcemanager_test

import (
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMResourceGroupDataSource_Basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceGroupDataSourceConfigDefault(),
			},
			{
				Config: testAccCheckIBMResourceGroupDataSourceConfigWithName(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_resource_group.testacc_ds_resource_group_name", "name", "default"),
				),
			},
			{
				Config: testAccCheckIBMResourceGroupDataSourceConfigWithID(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_resource_group.testacc_ds_id", "name", "default"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_group.testacc_ds_id", "id"),
				),
			},
		},
	})
}

func TestAccIBMResourceGroupDataSource_Default_false(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMResourceGroupDataSourceDefaultFalse(),
				ExpectError: regexp.MustCompile(`Missing required properties. Need a resource group name, resource group ID, or the is_default true`),
			},
		},
	})
}

func testAccCheckIBMResourceGroupDataSourceConfigDefault() string {
	return `
	
data "ibm_resource_group" "testacc_ds_resource_group" {
	is_default = "true"
}`

}

func testAccCheckIBMResourceGroupDataSourceConfigWithName() string {
	return `

data "ibm_resource_group" "testacc_ds_resource_group_name" {
	name = "default"
}`

}

func testAccCheckIBMResourceGroupDataSourceDefaultFalse() string {
	return `
	
data "ibm_resource_group" "testacc_ds_resource_group" {
	is_default = "false"
}`

}

func testAccCheckIBMResourceGroupDataSourceConfigWithID() string {
	return `

data "ibm_resource_group" "testacc_ds_resource_group_default" {
	is_default = "true"
}

data "ibm_resource_group" "testacc_ds_id" {
	id = data.ibm_resource_group.testacc_ds_resource_group_default.id
}`

}

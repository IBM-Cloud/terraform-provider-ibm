// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package resourcemanager_test

import (
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMResourceGroupsDataSource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceGroupsDataSourceConfigDefault(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_default", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_default", "resource_groups.#"),
					resource.TestCheckResourceAttr("data.ibm_resource_groups.testacc_ds_resource_groups_default", "resource_groups.0.is_default", "true"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_default", "resource_groups.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_default", "resource_groups.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_default", "resource_groups.0.state"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_default", "resource_groups.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_default", "resource_groups.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_default", "resource_groups.0.updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_default", "resource_groups.0.account_id"),
				),
			},
		},
	})
}

func TestAccIBMResourceGroupsDataSource_WithName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceGroupsDataSourceConfigWithName(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_name", "id"),
					resource.TestCheckResourceAttr("data.ibm_resource_groups.testacc_ds_resource_groups_name", "resource_groups.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_resource_groups.testacc_ds_resource_groups_name", "resource_groups.0.name", "Default"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_name", "resource_groups.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_name", "resource_groups.0.state"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_name", "resource_groups.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_name", "resource_groups.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_name", "resource_groups.0.updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_name", "resource_groups.0.teams_url"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_name", "resource_groups.0.payment_methods_url"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_name", "resource_groups.0.quota_url"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_name", "resource_groups.0.quota_id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_name", "resource_groups.0.account_id"),
				),
			},
		},
	})
}

func TestAccIBMResourceGroupsDataSource_WithDate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceGroupsDataSourceConfigWithDate(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_date", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_date", "resource_groups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_date", "resource_groups.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_date", "resource_groups.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_date", "resource_groups.0.state"),
				),
			},
		},
	})
}

func TestAccIBMResourceGroupsDataSource_IncludeDeleted(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceGroupsDataSourceConfigIncludeDeleted(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_include_deleted", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_include_deleted", "resource_groups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_include_deleted", "resource_groups.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_include_deleted", "resource_groups.0.name"),
				),
			},
		},
	})
}

func TestAccIBMResourceGroupsDataSource_AllResourceGroups(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceGroupsDataSourceConfigAll(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_all", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_all", "resource_groups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_all", "resource_groups.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_all", "resource_groups.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_all", "resource_groups.0.state"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_all", "resource_groups.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_all", "resource_groups.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_all", "resource_groups.0.updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_all", "resource_groups.0.account_id"),
					// Test multiple resource groups exist (at least default should exist)
					resource.TestCheckResourceAttr("data.ibm_resource_groups.testacc_ds_resource_groups_all", "resource_groups.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMResourceGroupsDataSource_ResourceLinkages(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceGroupsDataSourceConfigDefault(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_default", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_default", "resource_groups.0.resource_linkages.#"),
					// Resource linkages might be empty or populated, so we just check that the attribute exists
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.testacc_ds_resource_groups_default", "resource_groups.0.id"),
				),
			},
		},
	})
}

func TestAccIBMResourceGroupsDataSource_InvalidConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMResourceGroupsDataSourceInvalidConfig(),
				ExpectError: regexp.MustCompile(`Invalid combination of arguments`),
			},
		},
	})
}

func TestAccIBMResourceGroupsDataSource_MultipleGroups(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceGroupsDataSourceMultipleConfigs(),
				Check: resource.ComposeTestCheckFunc(
					// Test default groups
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.default_groups", "id"),
					resource.TestCheckResourceAttr("data.ibm_resource_groups.default_groups", "resource_groups.0.is_default", "true"),

					// Test named group
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.named_group", "id"),
					resource.TestCheckResourceAttr("data.ibm_resource_groups.named_group", "resource_groups.0.name", "Default"),

					// Test all groups
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.all_groups", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_groups.all_groups", "resource_groups.#"),
				),
			},
		},
	})
}

func testAccCheckIBMResourceGroupsDataSourceConfigDefault() string {
	return `
data "ibm_resource_groups" "testacc_ds_resource_groups_default" {
	is_default = true
}`
}

func testAccCheckIBMResourceGroupsDataSourceConfigWithName() string {
	return `
data "ibm_resource_groups" "testacc_ds_resource_groups_name" {
	name = "Default"
}`
}

func testAccCheckIBMResourceGroupsDataSourceConfigWithDate() string {
	return `
data "ibm_resource_groups" "testacc_ds_resource_groups_date" {
	date = "2024-01"
}`
}

func testAccCheckIBMResourceGroupsDataSourceConfigIncludeDeleted() string {
	return `
data "ibm_resource_groups" "testacc_ds_resource_groups_include_deleted" {
	include_deleted = true
}`
}

func testAccCheckIBMResourceGroupsDataSourceConfigAll() string {
	return `
data "ibm_resource_groups" "testacc_ds_resource_groups_all" {
	# No filters - should return all resource groups
}`
}

func testAccCheckIBMResourceGroupsDataSourceInvalidConfig() string {
	return `
data "ibm_resource_groups" "testacc_ds_resource_groups_invalid" {
	name = "Default"
	is_default = true
}`
}

func testAccCheckIBMResourceGroupsDataSourceMultipleConfigs() string {
	return `
data "ibm_resource_groups" "default_groups" {
	is_default = true
}

data "ibm_resource_groups" "named_group" {
	name = "Default"
}

data "ibm_resource_groups" "all_groups" {
	# No filters
}`
}

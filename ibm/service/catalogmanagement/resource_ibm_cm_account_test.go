// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/catalogmanagement"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMCmAccountBasic(t *testing.T) {
	var conf catalogmanagementv1.Account

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmAccountDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmAccountConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmAccountExists("ibm_cm_account.cm_account", conf),
					resource.TestCheckResourceAttr("ibm_cm_account.cm_account", "hide_ibm_cloud_catalog", "true"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCmAccountConfigBasicUpdate(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmAccountExists("ibm_cm_account.cm_account", conf),
					resource.TestCheckResourceAttr("ibm_cm_account.cm_account", "hide_ibm_cloud_catalog", "false"),
				),
			},
		},
	})
}

func testAccCheckIBMCmAccountConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
			label = "test_tf_account_catalog_label_1"
			kind = "offering"
		}

		resource "ibm_cm_account" "cm_account" {
			hide_ibm_cloud_catalog = true
			terraform_engines {
				name             = "my-tfe-instance"
				type             = "terraform-enterprise"
				public_endpoint  = "foo"
				private_endpoint = "foo"
				api_token        = "foo"
				da_creation {
					enabled                    = true
					default_private_catalog_id = ibm_cm_catalog.cm_catalog.id
					polling_info {
						scopes {
							name = "foo-project"
							type = "project"
						}
					}
				}
			}
		}
	`)
}

func testAccCheckIBMCmAccountConfigBasicUpdate() string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
			label = "test_tf_account_catalog_label_1"
			kind = "offering"
		}

		resource "ibm_cm_account" "cm_account" {
			hide_ibm_cloud_catalog = false
			terraform_engines {
				name             = "my-tfe-instance"
				type             = "terraform-enterprise"
				public_endpoint  = "foo"
				private_endpoint = "foo"
				api_token        = "foo"
				da_creation {
					enabled                    = true
					default_private_catalog_id = ibm_cm_catalog.cm_catalog.id
					polling_info {
						scopes {
							name = "foo-project"
							type = "project"
						}
					}
				}
			}
		}
	`)
}

func testAccCheckIBMCmAccountExists(n string, obj catalogmanagementv1.Account) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		catalogManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CatalogManagementV1()
		if err != nil {
			return err
		}

		getCatalogAccountOptions := &catalogmanagementv1.GetCatalogAccountOptions{}

		account, _, err := catalogManagementClient.GetCatalogAccount(getCatalogAccountOptions)
		if err != nil {
			return err
		}

		obj = *account
		return nil
	}
}

func testAccCheckIBMCmAccountDestroy(s *terraform.State) error {
	return nil
}

func TestResourceIBMCmAccountFiltersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		filterTermsModel := make(map[string]interface{})
		filterTermsModel["filter_terms"] = []string{"testString"}

		categoryFilterModel := make(map[string]interface{})
		categoryFilterModel["include"] = true
		categoryFilterModel["filter"] = []map[string]interface{}{filterTermsModel}

		idFilterModel := make(map[string]interface{})
		idFilterModel["include"] = []map[string]interface{}{filterTermsModel}
		idFilterModel["exclude"] = []map[string]interface{}{filterTermsModel}

		model := make(map[string]interface{})
		model["include_all"] = true
		model["category_filters"] = categoryFilterModel
		model["id_filters"] = []map[string]interface{}{idFilterModel}

		assert.Equal(t, result, model)
	}

	filterTermsModel := new(catalogmanagementv1.FilterTerms)
	filterTermsModel.FilterTerms = []string{"testString"}

	categoryFilterModel := new(catalogmanagementv1.CategoryFilter)
	categoryFilterModel.Include = core.BoolPtr(true)
	categoryFilterModel.Filter = filterTermsModel

	idFilterModel := new(catalogmanagementv1.IDFilter)
	idFilterModel.Include = filterTermsModel
	idFilterModel.Exclude = filterTermsModel

	model := new(catalogmanagementv1.Filters)
	model.IncludeAll = core.BoolPtr(true)
	model.CategoryFilters = map[string]catalogmanagementv1.CategoryFilter{"key1": *categoryFilterModel}
	model.IDFilters = idFilterModel

	result, err := catalogmanagement.ResourceIBMCmAccountFiltersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCmAccountCategoryFilterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		filterTermsModel := make(map[string]interface{})
		filterTermsModel["filter_terms"] = []string{"testString"}

		model := make(map[string]interface{})
		model["include"] = true
		model["filter"] = []map[string]interface{}{filterTermsModel}

		assert.Equal(t, result, model)
	}

	filterTermsModel := new(catalogmanagementv1.FilterTerms)
	filterTermsModel.FilterTerms = []string{"testString"}

	model := new(catalogmanagementv1.CategoryFilter)
	model.Include = core.BoolPtr(true)
	model.Filter = filterTermsModel

	result, err := catalogmanagement.ResourceIBMCmAccountCategoryFilterToMap("testString", model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCmAccountFilterTermsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["filter_terms"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(catalogmanagementv1.FilterTerms)
	model.FilterTerms = []string{"testString"}

	result, err := catalogmanagement.ResourceIBMCmAccountFilterTermsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCmAccountIDFilterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		filterTermsModel := make(map[string]interface{})
		filterTermsModel["filter_terms"] = []string{"testString"}

		model := make(map[string]interface{})
		model["include"] = []map[string]interface{}{filterTermsModel}
		model["exclude"] = []map[string]interface{}{filterTermsModel}

		assert.Equal(t, result, model)
	}

	filterTermsModel := new(catalogmanagementv1.FilterTerms)
	filterTermsModel.FilterTerms = []string{"testString"}

	model := new(catalogmanagementv1.IDFilter)
	model.Include = filterTermsModel
	model.Exclude = filterTermsModel

	result, err := catalogmanagement.ResourceIBMCmAccountIDFilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

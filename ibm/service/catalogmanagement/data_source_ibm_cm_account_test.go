// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.99.1-daeb6e46-20250131-173156
 */

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/service/catalogmanagement"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMCmAccountDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmAccountDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_account.cm_account_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMCmAccountDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
			label = "test_tf_account_catalog_label_1"
			kind = "offering"
		}

		resource "ibm_cm_account" "cm_account_instance" {
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

		data "ibm_cm_account" "cm_account_instance" {
			depends_on = [
				ibm_cm_account.cm_account_instance
			]
		}
	`)
}

func TestDataSourceIBMCmAccountFiltersToMap(t *testing.T) {
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

	result, err := catalogmanagement.DataSourceIBMCmAccountFiltersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCmAccountCategoryFilterToMap(t *testing.T) {
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

	result, err := catalogmanagement.DataSourceIBMCmAccountCategoryFilterToMap("testString", model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCmAccountFilterTermsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["filter_terms"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(catalogmanagementv1.FilterTerms)
	model.FilterTerms = []string{"testString"}

	result, err := catalogmanagement.DataSourceIBMCmAccountFilterTermsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCmAccountIDFilterToMap(t *testing.T) {
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

	result, err := catalogmanagement.DataSourceIBMCmAccountIDFilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

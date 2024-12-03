// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.0-d6dec9d7-20241008-212902
 */

package db2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/db2"
	"github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmDb2SaasWhitelistDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmDb2SaasWhitelistDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_whitelist.db2_saas_whitelist_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_whitelist.db2_saas_whitelist_instance", "x_deployment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_whitelist.db2_saas_whitelist_instance", "ip_addresses.#"),
				),
			},
		},
	})
}

func testAccCheckIbmDb2SaasWhitelistDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_db2_saas_whitelist" "db2_saas_whitelist_instance" {
			x-deployment-id = "x-deployment-id"
		}
	`)
}

func TestDataSourceIbmDb2SaasWhitelistIpAddressToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["address"] = "127.0.0.1"
		model["description"] = "A sample IP address"

		assert.Equal(t, result, model)
	}

	model := new(db2saasv1.IpAddress)
	model.Address = core.StringPtr("127.0.0.1")
	model.Description = core.StringPtr("A sample IP address")

	result, err := db2.DataSourceIbmDb2SaasWhitelistIpAddressToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
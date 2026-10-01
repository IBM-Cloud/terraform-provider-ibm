// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.116.0-df613dbc-20260803-154903
 */

package drautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/drautomationservice"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func TestAccIBMPdrManagedVMListDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrManagedVMListDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_managed_vm_list.pdr_managed_vm_list_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_managed_vm_list.pdr_managed_vm_list_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_managed_vm_list.pdr_managed_vm_list_instance", "managed_vm_list.%"),
				),
			},
		},
	})
}

func testAccCheckIBMPdrManagedVMListDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_managed_vm_list" "pdr_managed_vm_list_instance" {
			instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
			Accept-Language = "en-US"
		}
	`)
}

func TestDataSourceIBMPdrManagedVMListManagedVMDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["core"] = "0.50"
		model["dr_average_time"] = "10"
		model["memory"] = "4"
		model["region"] = "lon04"
		model["workgroup_name"] = "Workgroup1"
		model["workspace_name"] = "Workspace_dallas01"
		model["dr_region"] = "nyc02"
		model["vm_name"] = "example_vm"

		assert.Equal(t, result, model)
	}

	model := new(drautomationservicev1.ManagedVMDetails)
	model.Core = core.StringPtr("0.50")
	model.DrAverageTime = core.StringPtr("10")
	model.Memory = core.StringPtr("4")
	model.Region = core.StringPtr("lon04")
	model.WorkgroupName = core.StringPtr("Workgroup1")
	model.WorkspaceName = core.StringPtr("Workspace_dallas01")
	model.DrRegion = core.StringPtr("nyc02")
	model.VMName = core.StringPtr("example_vm")

	result, err := drautomationservice.DataSourceIBMPdrManagedVMListManagedVMDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

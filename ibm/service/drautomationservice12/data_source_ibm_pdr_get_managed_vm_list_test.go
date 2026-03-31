// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package drautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/drautomationservice"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM/dra-go-sdk/drautomationservicev1"
)

func TestAccIBMPdrGetManagedVMListDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrGetManagedVMListDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_managed_vm_list.pdr_get_managed_vm_list_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_managed_vm_list.pdr_get_managed_vm_list_instance", "instance_id"),
				),
			},
		},
	})
}

func testAccCheckIBMPdrGetManagedVMListDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_get_managed_vm_list" "pdr_get_managed_vm_list_instance" {
			instance_id = "xxxx2ec4-xxxx-4f84-xxxx-c2aa834dd4ed"
		}
	`)
}

func TestDataSourceIBMPdrGetManagedVMListManagedVMDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["core"] = "0.50"
		model["dr_average_time"] = "10"
		model["dr_region"] = "nyc02"
		model["memory"] = "4"
		model["region"] = "lon04"
		model["vm_name"] = "example_vm"
		model["workgroup_name"] = "Workgroup1"
		model["workspace_name"] = "Workspace_dallas01"

		assert.Equal(t, result, model)
	}

	model := new(drautomationservicev1.ManagedVMDetails)
	model.Core = core.StringPtr("0.50")
	model.DrAverageTime = core.StringPtr("10")
	model.DrRegion = core.StringPtr("nyc02")
	model.Memory = core.StringPtr("4")
	model.Region = core.StringPtr("lon04")
	model.VMName = core.StringPtr("example_vm")
	model.WorkgroupName = core.StringPtr("Workgroup1")
	model.WorkspaceName = core.StringPtr("Workspace_dallas01")

	result, err := drautomationservice.DataSourceIBMPdrGetManagedVMListManagedVMDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

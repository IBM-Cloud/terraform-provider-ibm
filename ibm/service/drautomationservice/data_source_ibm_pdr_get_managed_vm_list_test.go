// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.0-3c13b041-20250605-193116
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

func TestAccIbmPdrGetManagedVmListDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmPdrGetManagedVmListDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_managed_vm_list.pdr_get_managed_vm_list_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_managed_vm_list.pdr_get_managed_vm_list_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_managed_vm_list.pdr_get_managed_vm_list_instance", "managed_vms.#"),
				),
			},
		},
	})
}

func testAccCheckIbmPdrGetManagedVmListDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_get_managed_vm_list" "pdr_get_managed_vm_list_instance" {
			instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
		}
	`)
}

func TestDataSourceIbmPdrGetManagedVmListManagedVmListToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["vm_id"] = "5f4f7dd3-6d44-4086-9950-933bd550e24d"
		model["vm_name"] = "abc-aix-vm1"

		assert.Equal(t, result, model)
	}

	model := new(drautomationservicev1.ManagedVmList)
	model.VmID = core.StringPtr("5f4f7dd3-6d44-4086-9950-933bd550e24d")
	model.VmName = core.StringPtr("abc-aix-vm1")

	result, err := drautomationservice.DataSourceIbmPdrGetManagedVmListManagedVmListToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

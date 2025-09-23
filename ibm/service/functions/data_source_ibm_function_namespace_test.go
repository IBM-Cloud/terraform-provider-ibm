// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package functions_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFunctionNamespaceDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("Namespace-%d", acctest.RandIntRange(10, 100))
	resourceGroupName := "default"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckFunctionNamespaceDataSource(name, resourceGroupName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_namespace.namespace", "name", name),
					resource.TestCheckResourceAttr("ibm_function_namespace.namespace", "location", "us-south"),
				),
			},
		},
	})
}

func testAccCheckFunctionNamespaceDataSource(name, resourceGroupName string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "test_acc" {
                name = "%s"
        }

	resource "ibm_function_namespace" "namespace" {
                name                = "%s"
                resource_group_id   = data.ibm_resource_group.test_acc.id
	}

        data "ibm_function_namespace" "test_namespace" {
		name              = ibm_function_namespace.namespace.name
		
}`, resourceGroupName, name)

}

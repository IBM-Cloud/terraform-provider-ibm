// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccFunctionNamespaceDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("Namespace-%d", acctest.RandIntRange(10, 100))
	resourceGroupName := "default"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{

			resource.TestStep{
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

/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccFunctionActionDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionActionDataSource(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "exec.0.kind", "python:3"),
					resource.TestCheckResourceAttr("data.ibm_function_action.action", "name", name),
				),
			},
		},
	})
}

func testAccCheckFunctionActionDataSource(name, namespace string) string {
	return fmt.Sprintf(`
	
	resource "ibm_function_action" "pythonzip" {
		name      = "%s"
		namespace = "%s"
		
		exec {
		  kind = "python:3"
		  code = base64encode("test-fixtures/pythonaction.zip")
		}
	  }
	  
	  data "ibm_function_action" "action" {
		name      = ibm_function_action.pythonzip.name
		namespace = ibm_function_action.pythonzip.namespace
	  }
	  
`, name, namespace)

}

// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package functions_test

import (
	"fmt"
	"os"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFunctionActionDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{

			{
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
		  code = base64encode("../../test-fixtures/pythonaction.zip")
		}
	  }
	  
	  data "ibm_function_action" "action" {
		name      = ibm_function_action.pythonzip.name
		namespace = ibm_function_action.pythonzip.namespace
	  }
	  
`, name, namespace)

}

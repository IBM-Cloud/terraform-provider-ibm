/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

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

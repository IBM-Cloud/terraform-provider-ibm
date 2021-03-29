// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFunctionTriggerDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("terraform_trigger_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionTriggerDataSource(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "publish", "false"),
					resource.TestCheckResourceAttr("data.ibm_function_trigger.datatrigger", "name", name),
				),
			},
		},
	})
}

func testAccCheckFunctionTriggerDataSource(name, namespace string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_trigger" "trigger" {
	name      = "%s"		  
	namespace = "%s"
}

data "ibm_function_trigger" "datatrigger" {
	name      = ibm_function_trigger.trigger.name
	namespace = ibm_function_trigger.trigger.namespace
}
`, name, namespace)

}

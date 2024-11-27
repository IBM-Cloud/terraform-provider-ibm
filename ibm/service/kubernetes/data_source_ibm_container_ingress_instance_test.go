package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMContainerIngressInstanceDatasourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerIngressInstanceDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_ingress_instance.test_ds_instance", "instance_crn"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerIngressInstanceDataSourceConfig() string {
	return fmt.Sprintf(`
	resource "ibm_container_ingress_instance" "test_acc_instance" {
		instance_crn    = "%s"
		secret_group_id = "%s"
		is_default = "%t"
		cluster  = "%s"
	}

	data "ibm_container_ingress_instance" "test_ds_instance" {
	  instance_name = ibm_container_ingress_instance.test_acc_instance.instance_name
	  cluster = "%s"
	}`, acc.InstanceCRN, acc.SecretGroupID, true, acc.ClusterName, acc.ClusterName)
}

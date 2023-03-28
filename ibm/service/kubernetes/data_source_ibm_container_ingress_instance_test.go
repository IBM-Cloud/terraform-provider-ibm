package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMContainerIngressInstanceDatasourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-cluster-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerNLBDNSDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_ingress_instance.instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerIngressInstanceDataSourceConfig(name string) string {
	return testAccCheckIBMContainerVpcClusterBasic(name) + `
	data "ibm_container_ingress_instance" "instance" {
	    cluster = ibm_container_vpc_cluster.cluster.id
	}
`
}

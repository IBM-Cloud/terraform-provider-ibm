package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMContainerClusterFeature_Basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerClusterFeature_basic(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster_feature.feature", "public_service_endpoint", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster_feature.feature", "private_service_endpoint", "true"),
				),
			},
			{
				Config: testAccCheckIBMContainerClusterFeature_update(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster_feature.feature", "public_service_endpoint", "false"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster_feature.feature", "private_service_endpoint", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerClusterFeature_basic(clusterName string) string {
	return fmt.Sprintf(`
resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"
  default_pool_size = 1
  machine_type    = "%s"
  hardware       = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
}
resource ibm_container_cluster_feature feature {
  cluster                  = "${ibm_container_cluster.testacc_cluster.id}"
  private_service_endpoint  = "true"
  timeouts {
    create = "180m"
  }
}`, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerClusterFeature_update(clusterName string) string {
	return fmt.Sprintf(`
resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"
  default_pool_size = 1
  machine_type    = "%s"
  hardware       = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
}
resource ibm_container_cluster_feature feature {
  cluster                  = "${ibm_container_cluster.testacc_cluster.id}"
  public_service_endpoint  = "false"
  timeouts {
    update = "180m"
  }
}`, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}

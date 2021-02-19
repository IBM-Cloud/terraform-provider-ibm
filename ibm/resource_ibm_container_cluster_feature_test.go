/*
* IBM Confidential
* Object Code Only Source Materials
* 5747-SM3
* (c) Copyright IBM Corp. 2017,2021
*
* The source code for this program is not published or otherwise divested
* of its trade secrets, irrespective of what has been deposited with the
* U.S. Copyright Office.
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMContainerClusterFeature_Basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

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
  name              = "%s"
  datacenter        = "%s"
  default_pool_size = 1
  machine_type      = "%s"
  hardware          = "shared"
  public_vlan_id    = "%s"
  private_vlan_id   = "%s"
  timeouts {
	  create = "120m"
  }
}

resource "ibm_container_cluster_feature" "feature" {
  cluster                  = ibm_container_cluster.testacc_cluster.id
  private_service_endpoint = "true"
  timeouts {
    create = "720m"
  }
}`, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerClusterFeature_update(clusterName string) string {
	return fmt.Sprintf(`
resource "ibm_container_cluster" "testacc_cluster" {
  name              = "%s"
  datacenter        = "%s"
  default_pool_size = 1
  machine_type      = "%s"
  hardware          = "shared"
  public_vlan_id    = "%s"
  private_vlan_id   = "%s"
  timeouts {
	create = "120m"
  }
}

resource "ibm_container_cluster_feature" "feature" {
  cluster                 = ibm_container_cluster.testacc_cluster.id
  public_service_endpoint = "false"
  timeouts {
    update = "720m"
  }
}`, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}

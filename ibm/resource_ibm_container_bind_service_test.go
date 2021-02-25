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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMContainerBindServiceBasic(t *testing.T) {

	serviceName := fmt.Sprintf("tf-cluster-bind-%d", acctest.RandIntRange(10, 100))
	clusterName := fmt.Sprintf("tf-cluster-bind-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerBindServiceBasic(clusterName, serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_bind_service.bind_service", "namespace_id", "default"),
				),
			},
			{
				ResourceName:            "ibm_container_bind_service.bind_service",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cluster_name_id", "role"},
			},
		},
	})
}

func testAccCheckIBMContainerBindServiceBasic(clusterName, serviceName string) string {
	return fmt.Sprintf(`
  
resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"
  machine_type    = "%s"
  hardware        = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  wait_till       = "MasterNodeReady"
}

resource "ibm_resource_instance" "cos_instance" {
  name     = "%s"
  service  = "cloud-object-storage"
  plan     = "standard"
  location = "global"
}

resource "ibm_container_bind_service" "bind_service" {
  cluster_name_id     = ibm_container_cluster.testacc_cluster.id
  service_instance_id = ibm_resource_instance.cos_instance.guid
  namespace_id        = "default"
  role                = "Writer"
}
	`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, serviceName)
}

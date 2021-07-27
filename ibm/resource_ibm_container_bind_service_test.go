// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
					resource.TestCheckResourceAttr("ibm_container_bind_service.bind_service", "namespace_id", "default"),
					resource.TestCheckResourceAttrSet("ibm_container_bind_service.bind_service", "cluster_name_id"),
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

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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMContainerBindService_basic(t *testing.T) {

	serviceName := fmt.Sprintf("terraform-%d", acctest.RandIntRange(10, 100))
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerBindService_basic(clusterName, serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_bind_service.bind_service", "namespace_id", "default"),
				),
			},
		},
	})
}

func TestAccIBMContainerBindService_withTag(t *testing.T) {

	serviceName := fmt.Sprintf("terraform-%d", acctest.RandIntRange(10, 100))
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerBindServiceWithTag(clusterName, serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_container_bind_service.bind_service", "namespace_id", "default"),
					resource.TestCheckResourceAttr("ibm_container_bind_service.bind_service", "cluster_name_id", clusterName),
					resource.TestCheckResourceAttr("ibm_container_bind_service.bind_service", "tags.#", "1"),
				)},
		},
	})
}

func TestAccIBMContainerBindService_WithoutOptionalFields(t *testing.T) {

	serviceName := fmt.Sprintf("terraform-%d", acctest.RandIntRange(10, 100))
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerBindService_WithoutOptionalFields(clusterName, serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_bind_service.bind_service", "namespace_id", "default"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerBindService_WithoutOptionalFields(clusterName, serviceName string) string {
	return fmt.Sprintf(`

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

  machine_type    = "%s"
  hardware        = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  region          = "%s"
}

resource "ibm_resource_instance" "cos_instance" {
  name     = "%s"
  service  = "cloud-object-storage"
  plan     = "standard"
  location = "global"
}

resource "ibm_container_bind_service" "bind_service" {
  cluster_name_id     = ibm_container_cluster.testacc_cluster.id
  service_instance_id = element(split(":", ibm_resource_instance.cos_instance.id), 7)
  namespace_id        = "default"
  role                = "Writer"
}
	`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, csRegion, serviceName)
}

func testAccCheckIBMContainerBindService_basic(clusterName, serviceName string) string {
	return fmt.Sprintf(`
  
resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"
  machine_type    = "%s"
  hardware        = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
}

resource "ibm_resource_instance" "cos_instance" {
  name     = "%s"
  service  = "cloud-object-storage"
  plan     = "standard"
  location = "global"
}

resource "ibm_container_bind_service" "bind_service" {
  cluster_name_id     = ibm_container_cluster.testacc_cluster.id
  service_instance_id = element(split(":", ibm_resource_instance.cos_instance.id), 7)
  namespace_id        = "default"
  role                = "Writer"
}
	`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, serviceName)
}

func testAccCheckIBMContainerBindServiceWithTag(clusterName, serviceName string) string {
	return fmt.Sprintf(`
  
resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

  machine_type    = "%s"
  hardware        = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
}

resource "ibm_resource_instance" "cos_instance" {
  name     = "%s"
  service  = "cloud-object-storage"
  plan     = "standard"
  location = "global"
}

resource "ibm_container_bind_service" "bind_service" {
  cluster_name_id     = ibm_container_cluster.testacc_cluster.id
  service_instance_id = element(split(":", ibm_resource_instance.cos_instance.id), 7)
  namespace_id        = "default"
  role                = "Writer"
  tags                = ["test"]
}
	`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, serviceName)
}

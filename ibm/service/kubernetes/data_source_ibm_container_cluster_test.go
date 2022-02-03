// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"strconv"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMContainerClusterDataSource_basic(t *testing.T) {
	clusterName := fmt.Sprintf("tf-cluster-%d", acctest.RandIntRange(10, 100))
	serviceName := fmt.Sprintf("tf-cluster-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerClusterDataSource(clusterName, serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_cluster.testacc_ds_cluster", "id"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_bind_service.bind_service", "id"),
					testAccIBMClusterVlansCheck("data.ibm_container_cluster.testacc_ds_cluster"),
				),
			},
		},
	})
}

func testAccIBMClusterVlansCheck(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		var (
			vlansSize int
			err       error
		)

		if vlansSize, err = strconv.Atoi(a["vlans.#"]); err != nil {
			return err
		}
		if vlansSize < 1 {
			return fmt.Errorf("[ERROR] No subnets found")
		}
		return nil
	}
}
func testAccCheckIBMContainerClusterDataSource(clusterName, serviceName string) string {
	return testAccCheckIBMContainerBindServiceBasic(clusterName, serviceName) + `
	data "ibm_container_cluster" "testacc_ds_cluster" {
		cluster_name_id = ibm_container_cluster.testacc_cluster.id
	}
	data "ibm_container_bind_service" "bind_service" {
		cluster_name_id       = ibm_container_bind_service.bind_service.cluster_name_id
		service_instance_id = ibm_container_bind_service.bind_service.service_instance_id
		namespace_id          = "default"
	}
	`
}

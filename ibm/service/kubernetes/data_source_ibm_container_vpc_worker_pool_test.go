// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMContainerVPCClusterWorkerPoolDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-worker-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_worker_pool", "id"),
				),
			},
		},
	})
}

func TestAccIBMContainerVPCClusterWorkerPoolDataSource_dedicatedhost(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-worker-%d", acctest.RandIntRange(10, 100))
	hostpoolID := acc.HostPoolID
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceConfigDedicatedHost(name, hostpoolID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_vpc_worker_pool.vpc_worker_pool", "host_pool_id", hostpoolID),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceConfig(name string) string {
	return testAccCheckIBMVpcContainerWorkerPoolBasic(name) + `
	data "ibm_container_vpc_cluster_worker_pool" "testacc_ds_worker_pool" {
	    cluster = "${ibm_container_vpc_cluster.cluster.id}"
	    worker_pool_name = "${ibm_container_vpc_worker_pool.test_pool.worker_pool_name}"
	}
`
}

func TestAccIBMContainerVPCClusterWorkerPoolDataSourceEnvvar(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-wp-%d", acctest.RandIntRange(10, 100))
	testChecks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttrSet("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_worker_pool", "id"),
		resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_worker_pool", "autoscale_enabled", "false"),
	}
	if acc.CrkID != "" {
		testChecks = append(testChecks,
			resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_worker_pool", "crk", acc.CrkID),
			resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_worker_pool", "kms_instance_id", acc.KmsInstanceID),
		)
	}
	if acc.WorkerPoolSecondaryStorage != "" {
		testChecks = append(testChecks, resource.TestCheckResourceAttr(
			"data.ibm_container_vpc_cluster_worker_pool.testacc_ds_worker_pool", "secondary_storage.0.name", acc.WorkerPoolSecondaryStorage))
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceEnvvar(name),
				Check:  resource.ComposeTestCheckFunc(testChecks...),
			},
		},
	})
}

func TestAccIBMContainerVPCClusterWorkerPoolDataSourceKmsAccountEnvvar(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-wp-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceKmsAccountEnvvar(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_kms_worker_pool", "id"),
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_kms_worker_pool", "crk", acc.CrkID),
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_kms_worker_pool", "kms_instance_id", acc.KmsInstanceID),
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster_worker_pool.testacc_ds_kms_worker_pool", "kms_account_id", acc.KmsAccountID),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceConfigDedicatedHost(name, hostpoolID string) string {
	return testAccCheckIBMVpcContainerWorkerPoolDedicatedHostCreate(
		acc.ClusterName, name, "bx2d.4x16", acc.IksClusterSubnetID, acc.IksClusterVpcID, acc.IksClusterResourceGroupID, hostpoolID) + `
	data "ibm_container_vpc_worker_pool" "vpc_worker_pool" {
	    cluster = "${ibm_container_vpc_worker_pool.vpc_worker_pool.cluster}"
	    worker_pool_name = "${ibm_container_vpc_worker_pool.vpc_worker_pool.worker_pool_name}"
		depends_on = [
			ibm_container_vpc_worker_pool.vpc_worker_pool
		]
	}
`
}

func testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceEnvvar(name string) string {
	return testAccCheckIBMVpcContainerWorkerPoolBasic(name) + `
	data "ibm_container_vpc_cluster_worker_pool" "testacc_ds_worker_pool" {
	    cluster = "${ibm_container_vpc_worker_pool.test_pool.cluster}"
	    worker_pool_name = "${ibm_container_vpc_worker_pool.test_pool.worker_pool_name}"
	}
`
}

func testAccCheckIBMContainerVPCClusterWorkerPoolDataSourceKmsAccountEnvvar(name string) string {
	return testAccCheckIBMVpcContainerWorkerPoolKmsAccount(name) + `
	data "ibm_container_vpc_cluster_worker_pool" "testacc_ds_kms_worker_pool" {
	    cluster = "${ibm_container_vpc_worker_pool.test_pool.cluster}"
	    worker_pool_name = "${ibm_container_vpc_worker_pool.test_pool.worker_pool_name}"
	}
`
}
func TestAccIBMContainerVpcOpenshiftClusterWorkerPoolDataSource(t *testing.T) {
	name := fmt.Sprintf("tf-vpc-worker-%d", acctest.RandIntRange(10, 100))
	openshiftFlavour := "bx2.16x64"
	openShiftworkerCount := "2"
	operatingSystem := "REDHAT_8_64"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcOpenshiftClusterWorkerPoolDataSourceConfig(name, openshiftFlavour, openShiftworkerCount, operatingSystem),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_vpc_cluster_worker_pool.default_worker_pool", "operating_system", operatingSystem),
				),
			},
		},
	})
}

func testAccCheckIBMContainerVpcOpenshiftClusterWorkerPoolDataSourceConfig(name, openshiftFlavour, openShiftworkerCount, operatingSystem string) string {
	return testAccCheckIBMContainerOcpClusterBasic(name, openshiftFlavour, openShiftworkerCount, operatingSystem) + `
	data "ibm_container_vpc_cluster_worker_pool" "default_worker_pool" {
	    cluster = ibm_container_vpc_cluster.cluster.id
	    worker_pool_name = "default"
	}
`
}

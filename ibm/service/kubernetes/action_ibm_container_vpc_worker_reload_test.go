// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// You need to set up env vars:
// export IBM_CLUSTER_ID
// export IBM_CLUSTER_VPC_RESOURCE_GROUP_ID
// export IBMCLOUD_API_KEY
// export IC_API_KEY
// This test is here to help to focus on given resources, but requires everything else existing already
func TestAccIBMContainerVPCWorkerReloadEnvvar(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
        "null": {
            Source:            "hashicorp/null",
            VersionConstraint: ">= 3.0.0",
        },
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerVpcWorkerReloadEnvvar(acc.IksClusterID, acc.IksClusterResourceGroupID, "1h", false),
				Check:  testAccCheckIBMContainerVpcWorkerState(acc.IksClusterID, "undeploying", "undeployed", "reload_pending", "reloading"),
			},
		},
	})
}

func testAccCheckIBMContainerVpcWorkerReloadEnvvar(clusterName, resourceGroupID, customTimeout string, noWait bool) string {
	timeout := ""
	if customTimeout != "" {
		timeout = fmt.Sprintf(`timeout = "%s"`, customTimeout)
	}
	noWaitConfig := ""
	if noWait {
		noWaitConfig = fmt.Sprintf(`no_wait = "%s"`, "true")
	}
	return fmt.Sprintf(`
	# Get the worker from the cluster
	data "ibm_container_vpc_cluster" "cluster_data" {
		cluster_name_id              = %s
		resource_group_id = %s
	}

	# Reload the first worker using the worker ID (which is the server ID for  workers)
	action "ibm_container_vpc_worker_reload" "test_reload" {
		config {
			cluster_name_id           = ibm_container_vpc_cluster.cluster.id
			worker_id = data.ibm_container_vpc_cluster.cluster_data.workers[0]
			%s
			%s
		}
	}

	resource "null_resource" "trigger_reload" {
		provisioner "local-exec" {
			command = "echo 'All infrastructure ready. Triggering worker reload now...'"
		}

		lifecycle {
			action_trigger {
				events  = [after_create]
				actions = [action.ibm_container_vpc_worker_reload.test_reload]
			}
		}
	}
	`, clusterName, resourceGroupID, timeout, noWaitConfig)
}

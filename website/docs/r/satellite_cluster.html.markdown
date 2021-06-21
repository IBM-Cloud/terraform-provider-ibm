---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_cluster"
description: |-
  Manages IBM Cloud satellite cluster.
---

# ibm\_satellite_cluster

Create, update, or delete [IBM Cloud Satellite Cluster](https://cloud.ibm.com/docs/openshift?topic=openshift-satellite-clusters). Set up an Red Hat® OpenShift® on IBM Cloud™ clusters in an IBM Cloud™ Satellite location, and use the hosts of your own infrastructure that you added to your location as the worker nodes for the cluster.


## Example Usage

###  Create satellite cluster

```hcl
resource "ibm_satellite_cluster" "create_cluster" {
	name                   = "%s"  
	location               = var.location
	enable_config_admin    = true
	kube_version           = "4.5_openshift"
	resource_group_id      = data.ibm_resource_group.rg.id
	wait_for_worker_update = true
	dynamic "zones" {
		for_each = var.zones
		content {
			id	= zones.value
		}
	}
}

```

## Timeouts

ibm_satellite_cluster provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

* `create` - (Default 120 minutes) Used for creating Instance.
* `read`   - (Default 10 minutes) Used for reading Instance.
* `update` - (Default 120 minutes) Used for updating Instance.
* `delete` - (Default 30 minutes) Used for deleting Instance.

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The unique name for the new IBM Cloud Satellite cluster.
* `location` - (Required, string) The name or ID of the Satellite location.
* `kube_version` - (Optional, string) The OpenShift Container Platform version.
* `zones` - (Optional, array of strings)  The name of the zones to create the default worker pool in
* `worker_count` - (Optional, string) The number of worker nodes to create per zone in the default worker pool.
* `enable_config_admin` - (Optional, bool) User provided value to indicate opt-in agreement to SatCon admin agent.
* `host_labels` - (Optional, array of strings) Key-value pairs to label the host, such as cpu=4 to describe the host capabilities.
* `default_worker_pool_labels` - Labels on all the workers in the default worker pool.
* `pull_secret` - (Optional, string) The RedHat pull secret to create the OpenShift cluster.
* `zone` - (Optional, list) Zone for the worker pool in a multizone cluster. Nested `zone` blocks have the following structure:
    * `id` - The name of the zone.
* `resource_group_id` - (Optional, string) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`.
* `tags` - (Optional, array of strings) List of tags associated with this cluster.
*  `wait_for_worker_update` - (Optional, bool) Set to true to wait for kube version of woker nodes to update during the wokrer node kube version update. NOTE: setting wait_for_worker_update to false is not recommended. This results in upgrading all the worker nodes in the cluster at the same time causing the cluster downtime
* `patch_version` - (Optional, string) Set this to update the worker nodes with the required patch version. 
   The patch_version should be in the format - `patch_version_fixpack_version`. Learn more about the Kuberentes version [here](https://cloud.ibm.com/docs/containers?topic=containers-cs_versions).
    **NOTE**: To update the patch/fixpack versions of the worker nodes, Run the command `ibmcloud ks workers -c <cluster_name_or_id> --output json`, fetch the required patch & fixpack versions from `kubeVersion.target` and set the patch_version parameter.
* `retry_patch_version` - (Optional, int) This argument helps to retry the update of patch_version if the previous update fails. Increment the value to retry the update of patch_version on worker nodes.
* `tags` - (Optional, array of strings) Tags associated with the container cluster instance.
* `pod_subnet` - Specify a custom subnet CIDR to provide private IP addresses for pods. The subnet must be at least '/23' or larger. For more info, refer [here](https://cloud.ibm.com/docs/containers?topic=containers-cli-plugin-kubernetes-service-cli#pod-subnet).
* `service_subnet` -  Specify a custom subnet CIDR to provide private IP addresses for services. The subnet must be at least '/24' or larger. For more info, refer [here](https://cloud.ibm.com/docs/containers?topic=containers-cli-plugin-kubernetes-service-cli#service-subnet).


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the cluster.
* `crn` - CRN of the cluster.
* `ingress_hostname` - The Ingress hostname.
* `ingress_secret` - The Ingress secret.
* `state` - State.
* `master_status` - Status of kubernetes master.
* `master_url` - The Master server URL.
* `private_service_endpoint_url` - Private service endpoint url.
* `public_service_endpoint_url` - Public service endpoint url.
* `private_service_endpoint_enabled` - Private service endpoint status.
* `public_service_endpoint_enabled` - Public service endpoint status.
* `state` - The lifecycle state of the cluster.
* `resource_group_name` - The lifecycle state of the cluster.

**NOTE:**

1. The following arguments are immutable and cannot be changed:

* `name` -  The unique name for the new IBM Cloud Satellite cluster.
* `location` -  The name or ID of the Satellite location.
* `resource_group_id` -  The ID of the resource group.

2. Host assignment to workerpool:

*  When you attach a host to a Satellite location, the host automatically assigned to worker pools in satellite resources.
   Auto-assignment works based on matching host labels (https://cloud.ibm.com/docs/satellite?topic=satellite-hosts#host-autoassign-ov).
*  For manual assignment, Use `ibm_satellite_host` resource to assign the host to workerpools.


## Import

`ibm_satellite_cluster` can be imported using the location id or name.

Example:

```
$ terraform import ibm_satellite_cluster.cluster cluster

```

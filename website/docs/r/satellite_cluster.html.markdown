---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_cluster"
description: |-
  Manages IBM Cloud satellite cluster.
---

# ibm_satellite_cluster

Create, update, or delete [IBM Cloud Satellite Cluster](https://cloud.ibm.com/docs/openshift?topic=openshift-satellite-clusters). Set up an Red Hat OpenShift® on IBM Cloud clusters in an IBM Cloud Satellite location, and use the hosts of your own infrastructure that you added to your location as the worker nodes for the cluster.


## Example usage

###  Create satellite cluster

```terraform
resource "ibm_satellite_cluster" "create_cluster" {
	name                   = "%s"  
	location               = var.location
	enable_config_admin    = true
	kube_version           = "4.6_openshift"
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

The `ibm_satellite_cluster` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- `create` - (Default 120 minutes) Used for creating instance.
- `read`   - (Default 10 minutes) Used for reading instance.
- `update` - (Default 120 minutes) Used for updating instance.
- `delete` - (Default 30 minutes) Used for deleting instance.

## Argument reference

Review the argument references that you can specify for your resource. 

- `name` - (Required, String) The unique name for the new IBM Cloud Satellite cluster.
- `location` - (Required, String) The name or ID of the Satellite location.
- `kube_version` - (Optional, String) The Red Hart OpenShift Container Platform version.
- `zones` - (Optional, Array of Strings)  The name of the zones to create the default worker pool.
- `worker_count` - (Optional, String) The number of worker nodes to create per zone in the default worker pool.
- `enable_config_admin` - (Optional, Bool) User provided value to indicate opt-in agreement to SatCon admin agent.
- `host_labels` - (Optional, Array of Strings) Key-value pairs to label the host, such as `cpu=4` to describe the host capabilities.
- `default_worker_pool_labels` - (Optional, String) The labels on all the workers in the default worker pool.
- `pull_secret` - (Optional, String) The Red Hat pull secret to create the OpenShift cluster.
- `zone` - (Optional, List) The zone for the worker pool in a multi-zone cluster. 

   Nested scheme for `zone`:
    - `id` - The name of the zone.
- `resource_group_id` - (Optional, String) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`.
- `tags` - (Optional, Array of Strings) List of tags associated with this cluster.
-  `wait_for_worker_update` - (Optional, Bool) Set to **true** to wait for kube version of woker nodes to update during the worker node kube version update. **NOTE** setting `wait_for_worker_update` to **false** is not recommended. This results in upgrading all the worker nodes in the cluster at the same time causing the cluster downtime.
- `patch_version` - (Optional, String) Set this to update the worker nodes with the required patch version. 
   The `patch_version` should be in the format - `patch_version_fixpack_version`. For more information, see [Kuberentes version](https://cloud.ibm.com/docs/containers?topic=containers-cs_versions).
    **NOTE**: To update the patch/fixpack versions of the worker nodes, Run the command `ibmcloud ks workers -c <cluster_name_or_id> --output json`, fetch the required patch & fixpack versions from `kubeVersion.target` and set the patch_version parameter.
- `retry_patch_version` - (Optional, Integer) This argument helps to retry the update of `patch_version` if the previous update fails. Increment the value to retry the update of `patch_version` on worker nodes.
- `tags` - (Optional, Array of Strings) Tags associated with the container cluster instance.
- `pod_subnet` - Specify a custom subnet CIDR to provide private IP addresses for pods. The subnet must be at least `/23` or larger. For more information, see [Configuring VPC subnets](https://cloud.ibm.com/docs/containers?topic=containers-vpc-subnets).
- `service_subnet` -  Specify a custom subnet CIDR to provide private IP addresses for services. The subnet must be at least `/24` or larger. For more information, see [Configuring VPC subnets](https://cloud.ibm.com/docs/containers?topic=containers-vpc-subnets#vpc_basics).


## Attributes reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the cluster.
- `crn` - (String) The CRN of the cluster.
- `ingress_hostname` - (String) The Ingress hostname.
- `ingress_secret` - (String) The Ingress secret.
- `state` - (String) State.
- `master_status` - (String) The status of the Kubernetes master.
- `master_url` - (String) The master server URL.
- `private_service_endpoint_url` - (String) The private service endpoint URL.
- `public_service_endpoint_url` - (String) The public service endpoint URL.
- `private_service_endpoint_enabled` - (String) The private service endpoint status.
- `public_service_endpoint_enabled` - (String) The public service endpoint status.
- `state` - (String) The lifecycle state of the cluster.
- `resource_group_name` - (String) The lifecycle state of the cluster.

**Note**

1. The following arguments are immutable and cannot be changed:

- `name` -  The unique name for the new IBM Cloud Satellite cluster.
- `location` -  The name or ID of the Satellite location.
- `resource_group_id` -  The ID of the resource group.

2. Host assignment to workerpool:

-  When you attach a host to a Satellite location, the host automatically assigned to worker pools in satellite resources.
   Auto-assignment works based on matching host labels (https://cloud.ibm.com/docs/satellite?topic=satellite-assigning-hosts#host-autoassign-ov).
-  For manual assignment, Use `ibm_satellite_host` resource to assign the host to workerpools.


## Import

The `ibm_satellite_cluster` can be imported using the location id or name.

Example:

```
$ terraform import ibm_satellite_cluster.cluster cluster

```

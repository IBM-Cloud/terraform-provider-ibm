---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_cluster_worker_pool"
description: |-
  Manages IBM Cloud satellite cluster worker pool.
---

# ibm\_satellite\_cluster\_worker\_pool

 Create or delete a worker pool. The worker pool will be attached to the specified cluster[IBM Cloud Satellite Cluster Worker Pool](https://cloud.ibm.com/docs/satellite?topic=openshift-satellite-clusters#satcluster-worker-pools).

## Example Usage

###  Create satellite cluster worker pool

```hcl
resource "ibm_satellite_cluster_worker_pool" "create_cluster_wp" {
	name               = var.worker_pool_name
	cluster	           = var.cluster
	worker_count       = var.worker_count 
	resource_group_id  = data.ibm_resource_group.rg.id
	dynamic "zones" {
		for_each = var.zones
		content {
      		id	= zones.value
    	}
  	}
	host_labels        = var.host_labels
}	
```

###  Create satellite cluster worker pool without workers

```hcl
resource "ibm_satellite_cluster_worker_pool" "create_cluster_wp" {
	name               = var.worker_pool_name
	cluster	           = data.ibm_satellite_cluster.read_cluster.id
	dynamic "zones" {
		for_each = var.zones
		content {
      		id	= zones.value
    	}
  	}
}	
```

## Timeouts

ibm_satellite_cluster_worker_pool provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

* `create` - (Default 120 minutes) Used for creating Instance.
* `read`   - (Default 10 minutes) Used for reading Instance.
* `update` - (Default 120 minutes) Used for updating Instance.
* `delete` - (Default 90 minutes) Used for deleting Instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required, Forces new resource, string) The name of the worker pool.
* `cluster` - (Required, Forces new resource, string) The name or id of the cluster.
* `worker_count` - (Optional, Int) The number of worker nodes per zone in the worker pool.
* `flavor` - (Optional, string) The flavor defines the amount of virtual CPU, memory, and disk space that is set up in each worker node.
* `isolation` - (Optional, string) Isolation for the worker node.
* `disk_encryption` - (Optional, string) Disk encryption for worker node.
* `zones` - (Required, set) A nested block describing the zones of this worker_pool. Nested zones blocks have the following structure:
  * `id` - (Required, string) Name of the zone.
* `host_labels` - (Optional, array of strings) Key-value pairs to label the host, such as cpu=4 to describe the host capabilities.
* `worker_pool_labels` - Labels on all the workers in the worker pool.
* `resource_group_id` - (Optional, Forces new resource, string) The ID of the resource group.  You can retrieve the value from data source 
* `entitlement` - (Optional, string) The openshift cluster entitlement avoids the OCP licence charges incurred. Use cloud paks with OCP Licence entitlement to add the Openshift cluster worker pool.

 
## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the worker pool resource. The id is composed of \<cluster_name_id\>/\<worker_pool_id\>.<br/>

**NOTE:**

Host assignment to workerpool:

*  When you attach a host to a Satellite location, the host automatically assigned to worker pools in satellite resources.
   Auto-assignment works based on matching host labels (https://cloud.ibm.com/docs/satellite?topic=satellite-hosts#host-autoassign-ov).
*  For manual assignment, Use `ibm_satellite_host` resource to assign the host to workerpools.

## Import

`ibm_satellite_cluster_worker_pool` can be imported using cluster_name_id & worker_pool_id eg

```
$ terraform import ibm_satellite_cluster_worker_pool.example mycluster/5c4f4d06e0dc402084922dea70850e3b-7cafe35
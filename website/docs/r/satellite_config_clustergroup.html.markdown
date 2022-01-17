---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : Satellite cluster group"
description: |-
  Manages IBM Cloud Satellite cluster group.
---

# ibm_satellite_config_clustergroup

Create or delete [IBM Cloud Satellite cluster groups](https://test.cloud.ibm.com/docs/satellite?topic=satellite-setup-clusters-satconfig#setup-clusters-satconfig-groups). The cluster group specifies all Red Hat OpenShift on IBM Cloud clusters that you want to include into the deployment of your Kubernetes resources. The clusters can run in your Satellite location or in IBM Cloud.


## Example usage

###  Create Satellite cluster group

```terraform
resource "ibm_satellite_config_clustergroup" "group1" {
  name = "tf-satellite-clustergroup-1"
}

```

## Timeouts

The `ibm_satellite_config_clustergroup` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- `create` - (Default 5 minutes) Used for creating instance.
- `read`   - (Default 5 minutes) Used for reading instance.
- `update` - (Default 5 minutes) Used for updating instance.
- `delete` - (Default 5 minutes) Used for deleting instance.

## Argument reference

Review the argument references that you can specify for your resource. 

- `name` - (Required, String) The unique name for the new IBM Cloud Satellite cluster group.
- `clusters` - (Optional, List) The clusters need to be attached to the cluster group. 

   Nested scheme for `clusters`:
    - `cluster_id` - (Required, String) The id of the cluster.

## Attributes reference

In addition to the argument reference list, you can access the following attribute references after your resource is created.

- `uuid` - (String) The unique identifier of the cluster group.
- `created` - (String) The creation date of the cluster group.

**Note**

1. The following arguments are immutable and cannot be changed:

- `name` -  The unique name for the new IBM Cloud Satellite cluster group.


## Import

The `ibm_satellite_config_clustergroup` can be imported using the name.

Example:

```
$ terraform import ibm_satellite_config_clustergroup.group1 tf-satellite-clustergroup-1

```

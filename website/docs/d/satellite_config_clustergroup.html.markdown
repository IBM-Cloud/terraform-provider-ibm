---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : Satellite cluster group"
description: |-
  Get information about an IBM Cloud Satellite cluster group.
---

# ibm_satellite_config_clustergroup

Retrieve information about an existing IBM Cloud Satellite Cluster group. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about IBM Cloud Satellite Cluster groups, see [how to setup IBM Cloud Satellite cluster group](https://test.cloud.ibm.com/docs/satellite?topic=satellite-setup-clusters-satconfig#setup-clusters-satconfig-groups).

## Example usage

```terraform
data "ibm_satellite_config_clustergroup" "group1" {
  name = "tf-satellite-clustergroup-1"
}

```

## Argument reference

Review the argument reference that you can specify for your data source. 

- `name` - (Required, String) The unique name of the IBM Cloud Satellite cluster group.

## Attributes reference

In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `uuid` - (String) The unique identifier of the cluster group.
- `created` - (String) The creation date of the cluster group.
- `clusters` - (List) The attached clusters on the cluster group. 

   Nested scheme for `clusters`:
    - `cluster_id` - (String) The id of the cluster.
    - `name` - (String) The name of the cluster.

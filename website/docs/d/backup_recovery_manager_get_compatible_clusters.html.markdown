---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_manager_get_compatible_clusters"
description: |-
  Get information about backup_recovery_manager_get_compatible_clusters
subcategory: "IBM REST API"
---

# ibm_backup_recovery_manager_get_compatible_clusters

Provides a read-only data source to retrieve information about backup_recovery_manager_get_compatible_clusters. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_manager_get_compatible_clusters" "backup_recovery_manager_get_compatible_clusters" {
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `release_version` - (Optional, String) 

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_manager_get_compatible_clusters.
* `compatible_clusters` - (List) 
Nested schema for **compatible_clusters**:
	* `cluster_id` - (Integer) Specifies cluster id.
	* `cluster_incarnation_id` - (Integer) Specifies cluster incarnation id.
	* `cluster_name` - (String) Specifies cluster's name.
	* `current_version` - (String) Specifies the current version of the cluster.


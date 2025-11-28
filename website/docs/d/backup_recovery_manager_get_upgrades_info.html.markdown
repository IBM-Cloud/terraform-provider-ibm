---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_manager_get_upgrades_info"
description: |-
  Get information about backup_recovery_manager_get_upgrades_info
subcategory: "IBM REST API"
---

# ibm_backup_recovery_manager_get_upgrades_info

Provides a read-only data source to retrieve information about a backup_recovery_manager_get_upgrades_info. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_manager_get_upgrades_info" "backup_recovery_manager_get_upgrades_info" {
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `cluster_identifiers` - (Optional, List) Fetch upgrade progress details for a list of cluster identifiers in format clusterId:clusterIncarnationId.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_manager_get_upgrades_info.
* `upgrades_info` - (List) 
Nested schema for **upgrades_info**:
	* `cluster_id` - (Integer) Specifies cluster's id.
	* `cluster_incarnation_id` - (Integer) Specifies cluster's incarnation id.
	* `patch_software_version` - (String) Patch software version against which these logs are generated. This is specified for Patch type only.
	* `software_version` - (String) Upgrade software version against which these logs are generated.
	* `type` - (String) Specifies the type of upgrade on a cluster.
	  * Constraints: Allowable values are: `Upgrade`, `Patch`, `UpgradePatch`.
	* `upgrade_logs` - (List) Upgrade logs per node.
	Nested schema for **upgrade_logs**:
		* `logs` - (List) Upgrade logs for the node.
		Nested schema for **logs**:
			* `log` - (String) One log statement of the complete logs.
			* `time_stamp` - (Integer) Time at which this log got generated.
		* `node_id` - (String) Id of the node.
	* `upgrade_percent_complete` - (Float) Upgrade percentage complete so far.
	* `upgrade_status` - (String) Upgrade status.
	  * Constraints: Allowable values are: `Scheduled`, `Complete`, `InProgress`, `Failed`, `ClusterUnreachable`.


---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_manager_update_cluster_upgrades"
description: |-
  Manages backup_recovery_manager_update_cluster_upgrades.
subcategory: "IBM REST API"
---

# ibm_backup_recovery_manager_update_cluster_upgrades

Create, update, and delete backup_recovery_manager_update_cluster_upgradess with this resource.

## Example Usage

```hcl
resource "ibm_backup_recovery_manager_update_cluster_upgrades" "backup_recovery_manager_update_cluster_upgrades_instance" {
  auth_headers {
		key = "key"
		value = "value"
  }
  clusters {
		cluster_id = 1
		cluster_incarnation_id = 1
		current_version = "current_version"
  }
  patch_upgrade_params {
		auth_headers {
			key = "key"
			value = "value"
		}
		ignore_pre_checks_failure = true
		package_url = "package_url"
		target_version = "target_version"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `auth_headers` - (Optional, Forces new resource, List) Specifies the optional headers for upgrade request.
Nested schema for **auth_headers**:
	* `key` - (Required, String) Specifies the key or name of the header.
	* `value` - (Required, String) Specifies the value of the header.
* `clusters` - (Optional, Forces new resource, List) Array for clusters to be upgraded.
Nested schema for **clusters**:
	* `cluster_id` - (Optional, Integer) Specifies cluster id.
	* `cluster_incarnation_id` - (Optional, Integer) Specifies cluster incarnation id.
	* `current_version` - (Optional, String) Specifies current version of cluster.
* `interval_for_rolling_upgrade_in_hours` - (Optional, Forces new resource, Integer) Specifies the difference of time between two cluster's upgrade.
* `package_url` - (Optional, Forces new resource, String) Specifies URL from which package can be downloaded. Note: This option is only supported in Multi-Cluster Manager (MCM).
* `patch_upgrade_params` - (Optional, Forces new resource, List) Specifies the parameters for patch upgrade request.
Nested schema for **patch_upgrade_params**:
	* `auth_headers` - (Optional, List) Specifies the optional headers for the patch cluster request.
	Nested schema for **auth_headers**:
		* `key` - (Required, String) Specifies the key or name of the header.
		* `value` - (Required, String) Specifies the value of the header.
	* `ignore_pre_checks_failure` - (Optional, Boolean) Specify if pre check results can be ignored.
	  * Constraints: The default value is `false`.
	* `package_url` - (Optional, String) Specifies URL from which patch package can be downloaded. Note: This option is only supported in Multi-Cluster Manager (MCM).
	* `target_version` - (Optional, String) Specifies target patch version to which clusters are to be upgraded.
* `target_version` - (Optional, Forces new resource, String) Specifies target version to which clusters are to be upgraded.
* `time_stamp_to_upgrade_at_msecs` - (Optional, Forces new resource, Integer) Specifies the time in msecs at which the cluster has to be upgraded.
* `type` - (Optional, Forces new resource, String) Specifies the type of upgrade to be performed on a cluster.
  * Constraints: The default value is `Upgrade`. Allowable values are: `Upgrade`, `Patch`, `UpgradePatch`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the backup_recovery_manager_update_cluster_upgrades.


## Import

You can import the `ibm_backup_recovery_manager_update_cluster_upgrades` resource by using `id`. Specifies the ID of the object.

# Syntax
<pre>
$ terraform import ibm_backup_recovery_manager_update_cluster_upgrades.backup_recovery_manager_update_cluster_upgrades &lt;id&gt;
</pre>

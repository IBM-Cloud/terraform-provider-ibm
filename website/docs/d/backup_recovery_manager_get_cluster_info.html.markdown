---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_manager_get_cluster_info"
description: |-
  Get information about backup_recovery_manager_get_cluster_info
subcategory: "IBM REST API"
---

# ibm_backup_recovery_manager_get_cluster_info

Provides a read-only data source to retrieve information about a backup_recovery_manager_get_cluster_info. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_manager_get_cluster_info" "backup_recovery_manager_get_cluster_info" {
}
```


## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_manager_get_cluster_info.
* `cohesity_clusters` - (List) Specifies the array of clusters upgrade details.
Nested schema for **cohesity_clusters**:
	* `auth_support_for_pkg_downloads` - (Boolean) If cluster can support authHeader for upgrade or not.
	* `available_versions` - (List) Specifies the release versions the cluster can upgrade to.
	Nested schema for **available_versions**:
		* `notes` - (String) Specifies release's notes.
		* `patch_details` - (List) Specifies the details of the available patch release.
		Nested schema for **patch_details**:
			* `notes` - (String) Specifies patch release's notes.
			* `release_type` - (String) Release's type.
			* `version` - (String) Specifies patch release's version.
		* `release_stage` - (String) Specifies the stage of a release.
		* `release_type` - (String) Release's type e.g, LTS, Feature, Patch, MCM.
		* `type` - (String) Specifies the type of package or release.
		  * Constraints: Allowable values are: `Upgrade`, `Patch`, `UpgradePatch`.
		* `version` - (String) Specifies release's version.
	* `cluster_id` - (Integer) Specifies cluster id.
	* `cluster_incarnation_id` - (Integer) Specifies cluster incarnation id.
	* `cluster_name` - (String) Specifies cluster's name.
	* `current_patch_version` - (String) Specifies current patch version of the cluster.
	* `current_version` - (String) Specifies if the cluster is connected to management console.
	* `health` - (String) Specifies the health of the cluster.
	  * Constraints: Allowable values are: `Critical`, `NonCritical`.
	* `is_connected_to_helios` - (Boolean) Specifies if the cluster is connected to management console.
	* `location` - (String) Specifies the location of the cluster.
	* `multi_tenancy_enabled` - (Boolean) Specifies if multi tenancy is enabled in the cluster.
	* `node_ips` - (List) Specifies an array of node ips for the cluster.
	* `number_of_nodes` - (Integer) Specifies the number of nodes in the cluster.
	* `patch_target_upgrade_url` - (String) Specifies the patch package URL for the cluster. This is populated for patch update only.
	* `patch_target_version` - (String) Specifies target version to which clusters are upgrading. This is populated for patch update only.
	* `provider_type` - (String) Specifies the type of the cluster provider.
	  * Constraints: Allowable values are: `kMCMCohesity`, `kIBMStorageProtect`.
	* `scheduled_timestamp` - (Integer) Time at which an upgrade is scheduled.
	* `status` - (String) Specifies the upgrade status of the cluster.
	  * Constraints: Allowable values are: `InProgress`, `UpgradeAvailable`, `UpToDate`, `Scheduled`, `Failed`, `ClusterUnreachable`.
	* `target_upgrade_url` - (String) Specifies the upgrade URL for the cluster. This is populated for upgrade only.
	* `target_version` - (String) Specifies target version to which clusters are to be upgraded. This is populated for upgrade only.
	* `total_capacity` - (Integer) Specifies how total memory capacity of the cluster.
	* `type` - (String) Specifies the type of the cluster.
	  * Constraints: Allowable values are: `VMRobo`, `Physical`.
	* `update_type` - (String) Specifies the type of upgrade performed on a cluster. This is to be used with status field to know the status of the upgrade action performed on cluster.
	  * Constraints: Allowable values are: `Upgrade`, `Patch`, `UpgradePatch`.
	* `used_capacity` - (Integer) Specifies how much of the cluster capacity is consumed.
* `sp_clusters` - (List) Specifies the array of clusters claimed from IBM Storage Protect environment.
Nested schema for **sp_clusters**:
	* `cluster_id` - (Integer) Specifies cluster id.
	* `cluster_incarnation_id` - (Integer) Specifies cluster incarnation id.
	* `cluster_name` - (String) Specifies cluster's name.
	* `current_version` - (String) Specifies the currently running version on cluster.
	* `health` - (String) Specifies the health of the cluster.
	  * Constraints: Allowable values are: `Critical`, `NonCritical`.
	* `is_connected_to_helios` - (Boolean) Specifies if the cluster is connected to management console.
	* `node_ips` - (List) Specifies an array of node ips for the cluster.
	* `number_of_nodes` - (Integer) Specifies the number of nodes in the cluster.
	* `provider_type` - (String) Specifies the type of the cluster provider.
	  * Constraints: Allowable values are: `kMCMCohesity`, `kIBMStorageProtect`.
	* `total_capacity` - (Integer) Specifies total capacity of the cluster in bytes.
	* `type` - (String) Specifies the type of the SP cluster.
	* `used_capacity` - (Integer) Specifies how much of the cluster capacity is consumed in bytes.


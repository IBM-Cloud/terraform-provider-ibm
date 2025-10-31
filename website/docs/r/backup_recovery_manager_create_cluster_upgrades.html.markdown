---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_manager_create_cluster_upgrades"
description: |-
  Manages backup_recovery_manager_create_cluster_upgrades.
subcategory: "IBM REST API"
---

# ibm_backup_recovery_manager_create_cluster_upgrades

Create, update, and delete backup_recovery_manager_create_cluster_upgradess with this resource.

## Example Usage

```hcl
resource "ibm_backup_recovery_manager_create_cluster_upgrades" "backup_recovery_manager_create_cluster_upgrades_instance" {
}
```


## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the backup_recovery_manager_create_cluster_upgrades.
* `upgrade_response_list` - (List) Specifies a list of disks to exclude from being protected. This is only applicable to VM objects.
Nested schema for **upgrade_response_list**:
	* `cluster_id` - (Integer) Specifies cluster id.
	* `cluster_incarnation_id` - (Integer) Specifies cluster incarnation id.
	* `error_message` - (String) Specifies error message if failed to schedule upgrade.
	* `is_upgrade_scheduling_successful` - (Boolean) Specifies if upgrade scheduling was successsful.


## Import

You can import the `ibm_backup_recovery_manager_create_cluster_upgrades` resource by using `id`. Specifies the ID of the object.

# Syntax
<pre>
$ terraform import ibm_backup_recovery_manager_create_cluster_upgrades.backup_recovery_manager_create_cluster_upgrades &lt;id&gt;
</pre>

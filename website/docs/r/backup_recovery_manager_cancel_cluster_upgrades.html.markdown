---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_manager_cancel_cluster_upgrades"
description: |-
  Manages backup_recovery_manager_cancel_cluster_upgrades.
subcategory: "IBM REST API"
---

# ibm_backup_recovery_manager_cancel_cluster_upgrades

Create, update, and delete backup_recovery_manager_cancel_cluster_upgradess with this resource.

## Example Usage

```hcl
resource "ibm_backup_recovery_manager_cancel_cluster_upgrades" "backup_recovery_manager_cancel_cluster_upgrades_instance" {
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `cluster_identifiers` - (Optional, Forces new resource, List) Specifies the list of cluster identifiers. The format is clusterId:clusterIncarnationId.
  * Constraints: The list items must match regular expression `/^([0-9]+:[0-9]+)$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the backup_recovery_manager_cancel_cluster_upgrades.
* `cancelled_upgrade_response_list` - (List) Specifies list of cluster scheduled uprgade cancel response.
Nested schema for **cancelled_upgrade_response_list**:
	* `cluster_id` - (Integer) Specifies cluster id.
	* `cluster_incarnation_id` - (Integer) Specifies cluster incarnation id.
	* `error_message` - (String) Specifies an error message if failed to cancel a scheduled upgrade.
	* `is_upgrade_cancel_successful` - (Boolean) Specifies if scheduled upgrade cancelling was successful.


## Import

You can import the `ibm_backup_recovery_manager_cancel_cluster_upgrades` resource by using `id`. Specifies the ID of the object.

# Syntax
<pre>
$ terraform import ibm_backup_recovery_manager_cancel_cluster_upgrades.backup_recovery_manager_cancel_cluster_upgrades &lt;id&gt;
</pre>

---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_connector_agent_upgrade"
description: |-
  Manages Connector agent upgrade.
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_connector_agent_upgrade

Upgrade connector agents with this resource. The upgrade is processed asynchronously by the cluster, and the resource tracks the upgrade run by its `upgrade_id`.

~> **NOTE:** Destroying this resource only removes it from the Terraform state. It does not cancel or roll back the upgrade.

## Example Usage

```hcl
resource "ibm_backup_recovery_connector_agent_upgrade" "backup_recovery_connector_agent_upgrade_instance" {
  tenant_ids     = ["tenantId/"]
  connection_ids = [1]
  retry_forever  = false
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `connection_ids` - (Optional, Forces new resource, List) Specifies the specific connections whose connector agents should be upgraded.
* `retry_forever` - (Optional, Forces new resource, Boolean) When true or omitted, the upgrade is retried indefinitely until it succeeds. When false, the cluster retries up to 3 times before giving up.
  * Constraints: The default value is `true`.
* `tenant_ids` - (Optional, Forces new resource, List) Specifies the tenants whose connector agents should be upgraded (all connections under each listed tenant).

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the upgrade run. It is the same as `upgrade_id`.
* `status` - (String) Specifies the overall state of this upgrade run. Possible values are `NotStarted`, `InProgress`, `Succeeded`, `Failed`.
* `upgrade_id` - (String) UUID assigned to this upgrade run.


## Import

You can import the `ibm_backup_recovery_connector_agent_upgrade` resource by using the upgrade run ID.

# Syntax
<pre>
$ terraform import ibm_backup_recovery_connector_agent_upgrade.backup_recovery_connector_agent_upgrade &lt;id&gt;
</pre>

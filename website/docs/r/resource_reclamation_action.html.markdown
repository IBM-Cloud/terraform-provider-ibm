---
subcategory: "Resource management"
layout: "ibm"
page_title: "IBM: ibm_resource_reclamation_action"
description: |-
  Perform a resource reclamation action on IBM Cloud. This resource supports permanent deletion (`reclaim`) or restoration (`restore`) of a reclamation.
---

# ibm_resource_reclamation_action

The `ibm_resource_reclamation_action` resource allows you to perform actions on a resource reclamation in IBM Cloud. Supported actions are:

- `reclaim`: Permanently delete the resource associated with the reclamation.
- `restore`: Restore the resource from its reclamation state.

This resource executes the action immediately upon creation and reflects the latest reclamation state in Terraform.

## Example Usage

### Reclaim (permanent deletion)

```
resource "ibm_resource_reclamation_action" "delete" {
  id         = "1234abcd-56ef-78gh-90ij-klmnopqrstuv"
  action     = "reclaim"
  request_by = "user@example.com"       # optional
  comment    = "Deleting resource reclaim"  # optional
}
```

### Restore

```
resource "ibm_resource_reclamation_action" "restore" {
  id     = "1234abcd-56ef-78gh-90ij-klmnopqrstuv"
  action = "restore"
}
```

## Argument Reference

- `id` - (Required, String, ForceNew) The unique ID of the reclamation resource.
- `action` - (Required, String, ForceNew) The reclamation action to perform. Valid values: `reclaim` or `restore`.
- `request_by` - (Optional, String) The identifier of the user who requested the action, if different from the authentication token.
- `comment` - (Optional, String) A descriptive comment about the action.

## Attribute Reference

In addition to all input arguments, the following attributes are exported reflecting the current state of the reclamation:

- `entity_id` - (String) The entity ID related to the reclamation.
- `entity_type_id` - (String) The entity type ID.
- `entity_crn` - (String) The full Cloud Resource Name (CRN) associated with the reclamation.
- `resource_instance_id` - (String) The resource instance ID related to the reclamation.
- `resource_group_id` - (String) The resource group ID.
- `account_id` - (String) IBM Cloud account ID.
- `policy_id` - (String) The policy ID that triggered the reclamation.
- `state` - (String) Current state of the reclamation.
- `target_time` - (String) Target time for reclamation retention expiration.
- `created_at` - (String) Creation timestamp.
- `created_by` - (String) Creator identifier.
- `updated_at` - (String) Last updated timestamp.
- `updated_by` - (String) Last updater identifier.
- `custom_properties` - (Map) Additional custom properties associated with the reclamation.

## Timeouts

- `create` - (Default 5 minutes) Timeout for creating the reclamation action (running the action).
- `delete` - (Default 5 minutes) Timeout for deleting the resource in Terraform (after action completion).

## Notes

- This resource performs a one-shot operation on creation. Terraform's `delete` step clears local state and does not revert the action.
- Ensure you have sufficient IAM permissions to perform reclamation actions.
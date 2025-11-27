---
subcategory: "Resource management"
layout: "ibm"
page_title: "IBM: ibm_resource_reclamation_delete"
description: |-
  Permanently delete a reclaimed resource on IBM Cloud. This resource performs the equivalent of 'ibmcloud resource reclamation-delete'.
---

# ibm_resource_reclamation_delete

The `ibm_resource_reclamation_delete` resource allows you to permanently delete a reclaimed resource in IBM Cloud. This is equivalent to running the `ibmcloud resource reclamation-delete` command and performs an irreversible deletion of the resource.

This resource executes the permanent deletion immediately upon creation and reflects the latest reclamation state in Terraform.

~> **Warning:** This action is irreversible. Once a reclamation is permanently deleted, the resource cannot be restored.

## Example Usage

### Basic permanent deletion

```hcl
resource "ibm_resource_reclamation_delete" "example" {
  reclamation_id = "1234abcd-56ef-78gh-90ij-klmnopqrstuv"
  comment        = "Permanent deletion of reclaimed resource"
}
```

### With optional request_by field

```hcl
resource "ibm_resource_reclamation_delete" "example" {
  reclamation_id = "1234abcd-56ef-78gh-90ij-klmnopqrstuv"
  request_by     = "admin@example.com"
  comment        = "Admin-requested permanent deletion"
}
```

### Using with data source to find reclamations

```hcl
data "ibm_resource_reclamations" "my_reclamations" {
  account_id = "your-account-id"
}

resource "ibm_resource_reclamation_delete" "cleanup" {
  count = length(data.ibm_resource_reclamations.my_reclamations.reclamations)
  
  reclamation_id = data.ibm_resource_reclamations.my_reclamations.reclamations[count.index].id
  comment        = "Automated cleanup via Terraform"
}
```

## Argument Reference

The following arguments are supported:

- `reclamation_id` - (Required, String, ForceNew) The unique ID of the reclamation resource to permanently delete. This is the reclamation ID, not the resource instance ID.
- `request_by` - (Optional, String, ForceNew) The identifier of the user who requested the deletion, if different from the authentication token.
- `comment` - (Optional, String, ForceNew) A descriptive comment about the deletion action.

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

- `create` - (Default 5 minutes) Timeout for creating the reclamation deletion (executing the permanent deletion).
- `delete` - (Default 5 minutes) Timeout for removing the resource from Terraform state.

## Import

The `ibm_resource_reclamation_delete` resource cannot be imported since it represents a one-time deletion action.

## Notes

- This resource performs a permanent deletion operation on creation. The deletion is irreversible.
- Terraform's `destroy` operation only removes the resource from the Terraform state and does not affect the already-deleted reclamation.
- Ensure you have sufficient IAM permissions to perform reclamation deletion actions.
- This resource is equivalent to running `ibmcloud resource reclamation-delete <reclamation-id>` from the CLI.

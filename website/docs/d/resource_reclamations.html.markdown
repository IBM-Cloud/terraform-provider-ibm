---
subcategory: "Resource management"
layout: "ibm"
page_title: "IBM: ibm_resource_reclamations"
description: |-
  Retrieve a list of all resource reclamations associated with your IBM Cloud account.
---

# ibm_resource_reclamations

The `ibm_resource_reclamations` data source fetches all resource reclamations for the current IBM Cloud account. This is useful to view reclamations that have been created across all resource instances.

> **Note:** Reclamations represent resources pending deletion or restoration actions in IBM Cloud.

## Example Usage

```
data "ibm_resource_reclamations" "all" {}
```

To access reclamation IDs:

```
output "reclamation_ids" {
  value = [for r in data.ibm_resource_reclamations.all.reclamations : r.id]
}
```

## Argument Reference

This data source does not take any arguments.

## Attribute Reference

In addition to all input arguments, the following attributes are exported:

- `reclamations` - (List of Objects) List of reclamations with each object containing:

  - `id` - (String) The ID of the reclamation.
  - `entity_id` - (String) The ID of the entity for the reclamation.
  - `entity_type_id` - (String) The entity type ID.
  - `entity_crn` - (String) Full Cloud Resource Name (CRN) related to this reclamation.
  - `resource_instance_id` - (String) Resource instance ID associated.
  - `resource_group_id` - (String) Resource group ID.
  - `account_id` - (String) IBM Cloud account ID.
  - `policy_id` - (String) The policy ID that triggered the reclamation.
  - `state` - (String) Current state of the reclamation.
  - `target_time` - (String) Target retention expiration time (RFC3339).
  - `created_at` - (String) Creation timestamp of the reclamation.
  - `created_by` - (String) Creator of the reclamation.
  - `updated_at` - (String) Last updated timestamp.
  - `updated_by` - (String) Last updater of the reclamation.
  - `custom_properties` - (Map) Additional custom properties associated with the reclamation.

## Notes

- Reclamations are managed by IBM Cloud resource controller service.
- Ensure appropriate permissions to view reclamations in your account.

---
subcategory: "Container Registry"
layout: "ibm"
page_title: "IBM: ibm_cr_namespaces"
description: |-
  Reads IBM Cloud Container Registry namespaces.
---
# ibm_cr_namespaces

Lists all IBM Cloud Container Registry namespaces in your account in the targeted region. For more information about Container Registry, see [About IBM Cloud Container Registry](https://cloud.ibm.com/docs/Registry?topic=Registry-registry_overview).

## Example usage

The following example retrieves namespaces for your account in the targeted region.

```terraform
data "ibm_cr_namespaces" "test" {}

```

## Argument reference

Input parameters are not supported for this data source.

## Attribute reference

Review the attribute references that are exported.

- `id` - (String) The unique identifier of the ibm_cr_namespaces datasource.
- `namespaces` - (List) List of namespaces that are available in the account.

  Nested scheme for `namespace`:
  - `account` - (String) The IBM Cloud account that owns the namespace.
  - `crn` - (String) If the namespace is assigned to a resource group, the IBM Cloud CRN representing the namespace.
  - `created_date` - (Timestamp) The creation date of the namespace.
  - `name` - (String) The name of the namespace.
  - `resource_group_id` - (String) The ID of the resource group to which the namespace is assigned.
  - `resource_created_date` - (Timestamp) The date that the namespace was assigned to a resource group.
  - `updated_date` - (Timestamp) The date that the namespace was last updated.

---
subcategory: "Container Registry"
layout: "ibm"
page_title: "IBM: cr_namespaces"
description: |-
  Reads IBM Container Registry namespaces.
---

# ibm_cr_namespaces
Lists an IBM Cloud Container Registry namespaces of an account in the targeted region.. For more information, about container registry, see [about IBM Cloud Container Registry](https://cloud.ibm.com/docs/Registry?topic=Registry-registry_overview).

## Example usage
The following example shows how to configure a `namespace`.

```terraform
data "ibm_cr_namespaces" "test" {}

```

## Argument reference
The input parameters are not supported for this data source. 

## Attribute reference
Review the attribute references that are exported.

- `id` - (String) The unique identifier of the namespace datasource.
- `namespaces` - (List) List of namespaces available in the account.

  Nested scheme for `namespace`:
  - `account` - (String) The IBM Cloud account that owns the namespace.
  - `crn` - (String) The `CRN` of the namespace. If the namespace has been assigned to a resource group.
  - `created_date` - (Timestamp) The created time of the namespace.
  - `name` - (String) The name of the namespace to create.
  - `resource_group_id` - (String) ID of the resource group to which the namespace is assigned.
  - `resource_created_date` - (Timestamp) When the namespace was assigned to a resource group.
  - `updated_date` - (Timestamp) The updated time of the namespace.

---
subcategory: "Container Registry"
layout: "ibm"
page_title: "IBM: cr_namespaces"
description: |-
  Reads IBM Container Registry Namespaces.
---

# ibm\_cr_namespaces

Lists IBM Cloud Container Registry namespaces of an account in the targeted region.

## Example Usage

```terraform
data "ibm_cr_namespaces" "test" {}

```

## Argument Reference

The following arguments are supported:

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Id of the Namespace Datasource.
* `namespaces` - List of namespaces available in the account.
    * `name` - The Name of the namespace.
    * `resource_group_id` -  Id of the resource group to which the namespace has to be assigned.
    * `account` - The IBM Cloud account that owns the namespace.
    * `crn` - If the namespace has been assigned to a resource group, this is the IBM Cloud CRN representing the namespace.
    * `created_date` - When the namespace was created.
    * `resource_created_date` - When the namespace was assigned to a resource group.
    * `updated_date` - When the namespace was last updated.

---
subcategory: "Container Registry"
layout: "ibm"
page_title: "IBM: cr_namespaces"
description: |-
  Reads IBM Container Registry Namespaces.
---

# ibm\_cr_namespaces

Lists a Container Registry Namespaces of an account. 

## Example Usage

In the following example, you can configure a alb:

```hcl
data "ibm_cr_namespaces" "test" {}

```

## Argument Reference

The following arguments are supported:

## Attribute Reference

The following attributes are exported:

* `id` - Id of the Namespace Datasource.
* `namespaces` - List of namespaces available in the account.
    * `name` - The Name of the Namespace that has to be created.
    * `resource_group_id` -  Id of the resource group to which the namespace has to be assigned.
    * `crn` - Crn of the namespace.
    * `created_on` - The Created Time of the Namespace.
    * `updated_on` - The Updated Time of the Namespace.

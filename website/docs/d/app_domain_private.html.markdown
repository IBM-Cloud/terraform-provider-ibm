---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: ibm_app_domain_private"
description: |-
  Get information about an IBM Cloud domain private.
---

# ibm\_app_domain_private

Import the details of an existing IBM Cloud private domain as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_app_domain_private" "private_domain" {
  name = "foo.com"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the private domain.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the private domain.  

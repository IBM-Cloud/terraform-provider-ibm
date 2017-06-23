---
layout: "ibm"
page_title: "IBM: ibm_app_domain_private"
sidebar_current: "docs-ibm-datasource-app-domian-private"
description: |-
  Get information about an IBM Bluemix domain private.
---

# ibm\_app_domain_private

Import the details of an existing IBM Bluemix private domain as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration by using interpolation syntax. 

## Example Usage

```hcl
data "ibm_app_domain_private" "private_domain" {
  name = "foo.com"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the private domain.

## Attributes Reference

The following attributes are exported:

* `id` - The unique identifier of the private domain.  

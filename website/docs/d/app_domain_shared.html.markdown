---
layout: "ibm"
page_title: "IBM: ibm_app_domain_shared"
sidebar_current: "docs-ibm-datasource-app-domain-shared"
description: |-
  Get information about an IBM Bluemix shared domain.
---

# ibm\_app_domain_shared

Import the details of an existing IBM Bluemix shared domain as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_app_domain_shared" "shared_domain" {
  name = "foo.com"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the shared domain.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the shared domain.  

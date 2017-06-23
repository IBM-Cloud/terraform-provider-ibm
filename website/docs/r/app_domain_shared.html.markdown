---
layout: "ibm"
page_title: "IBM: app_domain_shared"
sidebar_current: "docs-ibm-resource-app-domain-shared"
description: |-
  Manages IBM Shared Domain.
---

# ibm\_app_domain_shared

Create, update, or delete shared domain on IBM Bluemix.

## Example Usage

```hcl
resource "ibm_app_domain_shared" "domain" {
  name              = "foo.com"
  router_group_guid = "3hG5jkjk4k34JH5666"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the domain.
* `router_group_guid` - (Optional, string) The GUID of the router group.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the shared domain.


---
layout: "ibm"
page_title: "IBM: app_domain_shared"
sidebar_current: "docs-ibm-resource-app-domain-shared"
description: |-
  Manages IBM Shared Domain.
---

# ibm\_app_domain_shared

Provides a shared domain resource. This allows shared domains to be created, updated, and deleted.

## Example Usage

```hcl
resource "ibm_app_domain_shared" "domain" {
  name              = "foo.com"
  router_group_guid = "3hG5jkjk4k34JH5666"
  tags              = ["tag1", "tag2"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the domain.
* `router_group_guid` - (Optional, string) The GUID of the router group.
* `tags` - (Optional, array of strings) Tags associated with the application shared domain instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the shared domain.

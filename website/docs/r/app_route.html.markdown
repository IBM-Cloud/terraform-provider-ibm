---
layout: "ibm"
page_title: "IBM: app_route"
sidebar_current: "docs-ibm-resource-app-route"
description: |-
  Manages IBM Application route.
---

# ibm\_app_route

Create, update, or delete route on IBM Bluemix.

## Example Usage

```hcl
data "ibm_space" "spacedata" {
  space = "space"
  org   = "someexample.com"
}

data "ibm_app_domain_shared" "domain" {
  name = "mybluemix.net"
}

resource "ibm_app_route" "route" {
  domain_guid = "${data.ibm_app_domain_shared.domain.id}"
  space_guid  = "${data.ibm_space.spacedata.id}"
  host        = "somehost172"
  path        = "/app"
}
```

## Argument Reference

The following arguments are supported:

* `domain_guid` - (Required, string) The GUID of the associated domain. The values can be retrieved from data source `ibm_app_domain_shared` or `ibm_app_domain_private`.
* `space_guid` - (Required, string) The GUID of the space where you want to create the route. The values can be retrieved from data source `ibm_space`, or by running the `bx iam space <space_name> --guid` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).
* `host` - (Optional, string) The host portion of the route. Required for shared-domains.
* `port` - (Optional, integer) The port of the route. Supported for domains of TCP router groups only.
* `path` - (Optional, string) The path for a route as raw text. Paths must be between 2 and 128 characters. Paths must start with a forward slash (/). Paths cannot contain a question mark (?).

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the route.

---
layout: "ibm"
page_title: "IBM: app_route"
sidebar_current: "docs-ibm-resource-app-route"
description: |-
  Manages IBM Application route.
---

# ibm\_app_route

Provides a route resource. This allows routes to be created, updated, and deleted.

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

* `domain_guid` - (Required, string) The GUID of the associated domain. You can retrieve the value from data source `ibm_app_domain_shared` or `ibm_app_domain_private`.
* `space_guid` - (Required, string) The GUID of the space where you want to create the route. You can retrieve the value from data source `ibm_space` or by running the `ibmcloud iam space <space_name> --guid` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `host` - (Optional, string) The host portion of the route. Host is required for shared-domains.
* `port` - (Optional, integer) The port of the route. Port is supported for domains of TCP router groups only.
* `path` - (Optional, string) The path for a route as raw text. Paths must be 2 - 128 characters. Paths must start with a forward slash (/) and cannot contain a question mark (?).
* `tags` - (Optional, array of strings) Tags associated with the route instance.  
    **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the route.

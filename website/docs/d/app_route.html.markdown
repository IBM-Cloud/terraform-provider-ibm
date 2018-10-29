---
layout: "ibm"
page_title: "IBM: ibm_app_route"
sidebar_current: "docs-ibm-datasource-app-route"
description: |-
  Get information about an IBM Cloud route.
---

# ibm\_app_route

Import the details of an existing IBM Cloud route as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_app_route" "route" {
  domain_guid = "${data.ibm_app_domain_shared.domain.id}"
  space_guid  = "${data.ibm_space.spacedata.id}"
  host        = "somehost"
  path        = "/app"
}
```

## Argument Reference

The following arguments are supported:

* `domain_guid` - (Required, string) The GUID of the associated domain. You can retrieve the value from the `ibm_app_domain_shared` data source.
* `space_guid` - (Required, string) The GUID of the space where you want to create the route. You can retrieve the value from the `ibm_space` data source or by running the `ibmcloud iam space <space-name> --guid` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `host` - (Optional, string) The host portion of the route. Required for shared domains.
* `port` - (Optional, integer) The port of the route. Supported for domains of TCP router groups only.
* `path` - (Optional, string) The path for a route as raw text. Paths must contain 2 - 128 characters. Paths must start with a forward slash (/). Paths must not contain a question mark (?).


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the route.  

---
layout: "ibm"
page_title: "IBM: app"
sidebar_current: "docs-ibm-datasource-app"
description: |-
  Get information about an IBM Application.
---

# ibm\_app

Import the details of an existing IBM Bluemix app as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration by using interpolation syntax. 

## Example Usage

```hcl
data "ibm_app" "testacc_ds_app" {
  name       = "my-app"
  space_guid = "${ibm_app.app.space_guid}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the application. The value can be retrieved by running the `bx app list` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).
* `space_guid` - (Required, string) The GUID of the Bluemix space where the application is deployed. The value can be retrieved with the `ibm_space` data source, or by running the `bx iam space <space-name> --guid` command in the Bluemix CLI.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the application.
* `memory` - Memory that is allocated to the application, specified in megabytes.
* `instances` - The number of instances of the application.
* `disk_quota` - The disk quota for an instance of the application, specified in megabytes.
* `buildpack` - Buildpack used by the application. It can be a) Blank to indicate auto-detection, b) A Git URL pointing to a buildpack, or c) The name of an installed buildpack.
* `environment_json` - Key/value pairs of all the environment variables. Does not include any system or service variables.
* `route_guid` - The route GUIDs that are bound to the application.
* `service_instance_guid` - The service instance GUIDs that are bound to the application.
* `package_state` - The state of the application package, such as staged, pending.
* `state` - The state of the application.

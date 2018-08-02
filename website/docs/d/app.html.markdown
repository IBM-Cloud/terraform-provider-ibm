---
layout: "ibm"
page_title: "IBM: app"
sidebar_current: "docs-ibm-datasource-app"
description: |-
  Get information about an IBM Application.
---

# ibm\_app

Import the details of an existing IBM Cloud app as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_app" "testacc_ds_app" {
  name       = "my-app"
  space_guid = "${ibm_app.app.space_guid}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the application. You can retrieve the value by running the `bx app list` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `space_guid` - (Required, string) The GUID of the IBM Cloud space where the application is deployed. You can retrieve the value with the `ibm_space` data source or by running the `bx iam space <space-name> --guid` command in the IBM Cloud CLI.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the application.
* `memory` - The memory, specified in megabytes, that is allocated to the application.
* `instances` - The number of instances of the application.
* `disk_quota` - The disk quota for an instance of the application, specified in megabytes.
* `buildpack` - The buildpack used by the application. It can be any of the following:
    * Blank, to indicate auto-detection.
    * A Git URL pointing to a buildpack.
    * The name of an installed buildpack.
* `environment_json` - Key/value pairs of all the environment variables. Does not include any system or service variables.
* `route_guid` - The route GUIDs that are bound to the application.
* `service_instance_guid` - The service instance GUIDs that are bound to the application.
* `package_state` - The state of the application package, such as staged or pending.
* `state` - The state of the application.
* `health_check_http_endpoint` - Endpoint called to determine if the app is healthy.
* `health_check_type` - Type of health check.
* `health_check_timeout` - Health check timeout in seconds.

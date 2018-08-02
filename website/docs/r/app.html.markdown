---
layout: "ibm"
page_title: "IBM: app"
sidebar_current: "docs-ibm-resource-app"
description: |-
  Manages IBM Application.
---

# ibm\_app

Provides an application resource. This allows applications to be created, updated, and deleted.

## Example Usage

```hcl
data "ibm_space" "space" {
  org   = "example.com"
  space = "dev"
}

resource "ibm_app" "app" {
  name                 = "my-app"
  space_guid           = "${data.ibm_space.space.id}"
  app_path             = "hello.zip"
  wait_timeout_minutes = 90
  buildpack            = "sdk-for-nodejs"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the application. You can retrieve the value by running the `bx app list` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `memory` - (Optional, integer) The amount of memory, specified in megabytes, that each instance has. If you don't specify a value, the system assigns pre-defined values based on the quota allocated to the application. You can check the default values by running `bx cf org <org-name>`. The command lists the quotas that are defined in your organization and space. If space quotas are defined, you can get them by running `bx cf space-quota <space-quota-name>`, where <quota-name> is the name of the quota. Otherwise you can check the organization quotas by running `bx cf quota <quota-name>`.
* `instances` - (Optional, integer) The number of instances of the application.
* `disk_quota` - (Optional, integer) The maximum amount of disk, specified in megabytes, available to an instance of an application. The default value is [1024 MB](http://bosh.io/jobs/cloud_controller_ng?source=github.com/cloudfoundry/cf-release&version=234#p=cc.default_app_disk_in_mb). Check with your cloud provider if the value has been set differently.
* `space_guid` - (Required, string) The GUID of the space where the application is deployed. You can retrieve the value from data source `ibm_space` or by running the `bx iam space <space-name> --guid` command in the IBM Cloud CLI.
* `buildpack` - (Optional, string) The buildpack to compile or prepare the application. You can provide its values in the following ways:
  * Leave the value blank for auto-detection.
  * Point to the Git URL for a buildpack. For example, https://github.com/cloudfoundry/nodejs-buildpack.git.
  * List the name of an installed buildpack. For example, `go_buildpack`.
* `environment_json` - (Optional, map) The key/value pairs of all the environment variables to run in your application. Do not provide any key/value pairs for system or service variables.
* `command` - (Optional, string) The initial command for the app.
* `route_guid` - (Optional, set) The route GUIDs that bind you want to the application. The route must be in the same space as the application.
* `service_instance_guid` - (Optional, set) The service instance GUIDs that you want to bind to the application.
* `wait_time_minutes` - (Optional, integer) The duration, expressed in minutes, to wait for the application to restage or start. The default value is `20`. A value of `0` means that there is no wait period.
* `app_path` - (Required, string) The path to the compressed file of the application. The compressed file must contain all the application files directly within it instead of within a top-level folder. To create the compressed file, go to the directory where your application files are and run `zip -r myapplication.zip *`.
* `app_version`	 - (Optional, string) The version of the application. If you make changes to the content in the application compressed file specified by _app_path_, Terraform can't detect the changes. You can let Terraform know that your file content has changed by either changing the application compressed file name or by using this argument to indicate the version of the file.
* `health_check_http_endpoint` - (Optional, string) Endpoint called to determine if the app is healthy.
* `health_check_type` - (Optional, string) Type of health check to perform. Default `port`. Valid types are `port` and `process`.
* `health_check_timeout` - (Optional, integer) Timeout in seconds for health checking of an staged app when starting up.
* `tags` - (Optional, array of strings) Tags associated with the application instance.
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the application.

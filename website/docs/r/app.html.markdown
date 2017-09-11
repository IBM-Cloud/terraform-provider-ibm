---
layout: "ibm"
page_title: "IBM: app"
sidebar_current: "docs-ibm-resource-app"
description: |-
  Manages IBM Application.
---

# ibm\_app

Create, update, or delete IBM application on IBM Bluemix.

## Example Usage

```hcl	
data "ibm_space" "space" {
  org   = "example.com"
  space = "dev"
}

resource "ibm_app" "app" {
  name              = "my-app"
  space_guid        = "${data.ibm_space.space.id}"
  app_path          = "hello.zip"
  wait_time_minutes = 90
  buildpack         = "sdk-for-nodejs"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the application. The value can be retrieved by running the `bx app list` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).
* `memory` - (Optional, integer) The amount of memory (in megabytes) that each instance should have. If you don't specify a value, the system assigns pre-defined values based on the quota allocated to the application. You can check the default values by running `bx cf org <org-name>`. The command lists the quotas defined in your org and space.
If space quotas are defined, you can get them by running `bx cf space-quota <space-quota-name>`. Otherwise you can check the organization quotas by running `bx cf quota <quota-name>`.
* `instances` - (Optional, integer) The number of instances of the application.
* `disk_quota` - (Optional, integer) The maximum amount of disk (in megabytes) available to an instance of an application. Default value: [1024 MB](http://bosh.io/jobs/cloud_controller_ng?source=github.com/cloudfoundry/cf-release&version=234#p=cc.default_app_disk_in_mb). Please check with your cloud provider if the value has been set differently.
* `space_guid` - (Required, string) Define the GUID of the space where the application is deployed. The value can be retrieved from data source `ibm_space`, or by running the `bx iam space <space-name> --guid` command in the Bluemix CLI.
* `buildpack` - (Optional, string) Buildpack to build the application. You can provide its values in the following ways:
  * Leave the value blank for auto-detection.
  * Point to the Git URL for a buildpack. For example, https://github.com/cloudfoundry/nodejs-buildpack.git.
  * List the name of an installed buildpack. For example, `go_buildpack`.
* `environment_json` - (Optional, map) Key/value pairs of all the environment variables to run in your application. Does not include any system or service variables.
* `command` - (Optional, string) The initial command for the app.
* `route_guid` - (Optional, set) Define the route GUIDs which should be bound to the application. Route should be in the same space as application.
* `service_instance_guid` - (Optional, set) Define the service instance GUIDs that should be bound to this application.
* `wait_time_minutes` - (Optional, integer) Define the timeout to wait for the application to restage/start. Default value: 20 minutes. A value of 0 means no wait period.
* `app_path` - (Required, string) Define the path to the zip file of the application. The zip must contain all the application files directly within it and not inside a top-level folder. Typically, you should go to the directory where your application files reside and issue `zip -r myapplication.zip *`.
* `app_version`	 - (Optional, string) Version of the application. If the application content in the file specified by _app_path_ changes, Terraform can't detect it. You can either change the application zip file name to let Terraform know that your zip content has changed, or you can use this attribute to let the provider know that the content changed without changing the _app_path_.
* `tags` - (Optional, array of strings) Set tags on the application instance.

**NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the application.

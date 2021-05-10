---

subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: app"
description: |-
  Manages IBM application.
---


# `ibm_app`

Create, update, or delete a Cloud Foundry app. For more information, about IBM Cloud Pak for application, see [about IBM applications](https://cloud.ibm.com/docs/cloud-pak-applications?topic=cloud-pak-applications-about).


## Example usage
The following example creates the `my-app` Node.js Cloud Foundry app. 


```
data "ibm_space" "space" {
  org   = "example.com"
  space = "dev"
}

resource "ibm_app" "app" {
  name                 = "my-app"
  space_guid           = data.ibm_space.space.id
  app_path             = "hello.zip"
  wait_timeout_minutes = 90
  buildpack            = "sdk-for-nodejs"
}
```


## Argument reference
Review the input parameters that you can specify for your resource. 

- `app_path`- (Required, String) The path to the compressed file of the app. The compressed file must contain all the app files without subdirectories. To create the compressed file, go to the directory where your app files are and run `zip -r myapplication.zip *`.
- `app_version`	 (Optional, String) The version of the app. If you make changes to the content in the app compressed file specified by _app_path_,  Terraform can't detect the changes. You can let  Terraform know that your file content has changed by either changing the application compressed file name or by using this argument to indicate the version of the file.
- `buildpack` - (String)  Optional-The buildpack to use when compiling or preparing the app. Provide the buildpack in one of the following ways: <ul><li>Leave it blank for auto-detection.</li><li>Enter the GitHub URL to a buildpack, such as `https://github.com/cloudfoundry/nodejs-buildpack.git`.</li><li>List the name of an installed buildpack. For example, `go_buildpack`</li></ul>.
- `command` - (Optional, String)  The initial command to run when the app starts.
- `disk_quota` - (Optional, Integer) The maximum amount of disk space, specified in megabytes, that is available to an app instance. The default value is [1024 MB](http://bosh.io/jobs/cloud_controller_ng?source=github.com/cloudfoundry/cf-release&version=234#p=cc.default_app_disk_in_mb).
- `environment_json` (Optional, Map) A list of environment variables to run your app, specified as key-value pairs. Do not provide any key-value pairs for system or service variables.
- `health_check_http_endpoint` - (Optional, String) The endpoint that you want to use to determine if the app is healthy.
- `health_check_type` - (String)  Optional-The type of health check that you want to perform. Supported values are `port`, and `process`. The default values is `port`.
- `health_check_timeout` - (Optional, Integer) The number of seconds to wait for the health check to respond during the start of your app before the health check is considered failed.
- `instances` - (Optional, Integer) The number of app instances that you want to create.
- `memory` - (Optional, Integer) The amount of memory, specified in megabytes, that is allocated to each instance. If you don't specify a value, the system assigns pre-defined values based on the app quota. You can check the default values by running `ibmcloud cf org <org-name>`. The command lists the quotas that are defined in your Cloud Foundry organization and space. If space quotas are defined, you can see them by running `ibmcloud cf space-quota <space-quota-name>`, where `<quota-name>` is the name of the quota. To check the organization quotas, run `ibmcloud cf quota <quota-name>`.
- `name` - (Required, String) The name of the app that you want to create, update, or delete. You can retrieve the value by running the `ibmcloud app list` command in the IBM Cloud CLI.
- `route_guid`(Optional, Sets)  The GUIDs of the routes that you want to bind to the application. The route must be in the same Cloud Foundry space as the app.
- `service_instance_guid`(Optional, Sets)  The GUID of the service instance that you want to bind to the app.
- `space_guid`- (Required, String) The GUID of the space where the app is deployed. You can retrieve the value from data source `ibm_space` or by running the `ibmcloud iam space <space-name> guid` command in the IBM Cloud CLI.
- `tags` (Array of Strings) Optional- The tags that you want to add to your app instance. Tags can help you find your app more easily.  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.
- `wait_time_minutes` - (Optional, Integer) The duration, expressed in minutes, to wait for the app to restage or start. The default value is `20`. A value of `0` means that there is no wait period.


## Attribute reference
Review the output parameters that you can access after your resource is created. 

- `id` - (String) The unique identifier of the application.

---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: app"
description: |-
  Get information about an IBM Application.
---

# `ibm_app`

Retrieve information about an existing Cloud Foundry app. For more information, about a Cloud Foundry application, see [getting started with Cloud Foundry Public](https://cloud.ibm.com/docs/cloud-foundry-public?topic=cloud-foundry-public-getting-started).


## Example usage
The following example retrieves information about the `my-app` Cloud Foundry app.  


```
data "ibm_app" "testacc_ds_app" {
  name       = "my-app"
  space_guid = ibm_app.app.space_guid
}
```

## Argument reference
Review the input parameters that you can specify for your data source. 

- `name` - (Required, String) The name of the app. You can retrieve the value by running the `ibmcloud app list` command in the IBM Cloud CLI.
- `space_guid` - (Required, String) The GUID of the IBM Cloud space where the app is deployed. You can retrieve the value with the `ibm_space` data source or by running the `ibmcloud iam space <space-name> guid` command in the IBM Cloud CLI.


## Attribute reference
Review the output parameters that you can access after you retrieved your data source. 

- `buildpack` - (String) The buildpack that is used by the app. Supported values are: <ul><li>Blank, indicates auto-detection</li><li>A Git URL that points to a buildpack</li><li>The name of an installed buildpack</li></ul>.
- `disk_quota`- (Integer) The maximum amount of disk space that an app instance can use, specified in megabytes.
- `environment_json`- (List of strings) A list of environment variables that the app uses. Environment variables are listed as key-value pairs and do not include system or service variables.
- `health_check_http_endpoint` - (String) The endpoint that is used to perform an HTTP health check and determine if the app is healthy.
- `health_check_type` - (String) Type of health check that is performed.
- `health_check_timeout`- (Integer) The timeout in seconds that the app remains unresponsive before the app is considered to be unhealthy.
- `id` - (String) The unique identifier of the app.
- `instances`- (Integer) The number of app instances that are deployed.
- `memory`- (Integer) The amount of memory, specified in megabytes, that is allocated to the app.
- `package_state` - (String) The state of the app package, such as `staged` or `pending`.
- `route_guid` - (String) The GUIDs of the routes that are assigned to the app.
- `service_instance_guid` - (String) The GUIDs of the service instances that are bound to the app.
- `state` - (String) The state of the app.

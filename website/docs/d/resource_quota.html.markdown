---
layout: "ibm"
page_title: "IBM: ibm_resource_quota"
sidebar_current: "docs-ibm-datasource-resource-quota"
description: |-
  Get information about an IBM Cloud resource quota.
---

# ibm\_resource_quota

Import the details of an existing quota for an IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_resource_quota" "rsquotadata" {
  name = "Trial Quota"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the quota for the IBM Cloud resource. You can retrieve the value by running the `ibmcloud resource quotas` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the quota.
* `type` - Type of the quota.
* `max_apps` - Defines the total app limit.
* `max_instances_per_app` - Defines the total instances limit per app.
* `max_app_instance_memory` - Defines the total memory of app instance.
* `total_app_memory` - Defines the total memory for app.
* `max_service_instances` - Defines the total service instances limit.
* `vsi_limit` - Defines the VSI limit.
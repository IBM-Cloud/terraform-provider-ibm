---

subcategory: "Resource management"
layout: "ibm"
page_title: "IBM: ibm_resource_quota"
description: |-
  Get information about an IBM Cloud resource quota.
---

# ibm_resource_quota
Retrieve information for an existing quota for an IBM Cloud as a read-only data source. For more information, about resource quote, see [ibmcloud resource quota](https://cloud.ibm.com/docs/account?topic=cli-ibmcloud_commands_resource#ibmcloud_resource_quota).

## Example usage

```terraform
data "ibm_resource_quota" "rsquotadata" {
  name = "Trial Quota"
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `name` - (Required, String) The name of the quota for the IBM Cloud resource. You can retrieve the value by executing the `ibmcloud resource quotas` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier of the quota.
- `max_apps` - (String) Defines the total app limit.
- `max_service_instances` - (String) Defines the total service instances limit.
- `max_instances_per_app` - (String) Defines the total instances limit per app.
- `max_app_instance_memory` - (String) Defines the total memory of app instance.
- `type` - (String) Type of the quota.
- `total_app_memory` - (String) Defines the total memory for app.
- `vsi_limit` - (String) Defines the VSI limit.

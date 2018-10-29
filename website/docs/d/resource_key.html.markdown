---
layout: "ibm"
page_title: "IBM: ibm_resource_key"
sidebar_current: "docs-ibm-datasource-resource-key"
description: |-
  Get information about a resource key from IBM Cloud.
---

# ibm\_resource_key

Import the details of an existing IBM resource key from IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_resource_key" "resourceKeydata" {
  name                  = "myobjectKey"
  resource_instance_id  = "${ibm_resource_instance.resource.id}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the resource key. You can retrieve the value by running the `ibmcloud resource service-keys` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `resource_instance_id` - (Optional, string) The id of the resource instance that the resource key is associated with. You can retrieve the value by running the `ibmcloud resource service-instances` command in the IBM Cloud CLI.  
  **NOTE**: Conflicts with `resource_alias_id`.
* `resource_alias_id` - (Optional, string) The id of the resource alias that the resource key is associated with. You can retrieve the value by running the `ibmcloud resource service-alias` command in the IBM Cloud CLI.  
  **NOTE**: Conflicts with `resource_instance_id`.
* `most_recent` - (Optional, boolean) If there are multiple resource keys, you can set this argument to `true` to import only the most recently created key.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the resource key.
* `credentials` - The credentials associated with the key.
* `role` - The user role.
* `status` - Status of resource key.  

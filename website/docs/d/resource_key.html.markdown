---

subcategory: "Resource management"
layout: "ibm"
page_title: "IBM: ibm_resource_key"
description: |-
  Get information about a resource key from IBM Cloud.
---

# ibm_resource_key

Retrieve information about an existing IBM resource key from IBM Cloud as a read-only data source. For more information, about resource key, see [ibmcloud resource service-keys](https://cloud.ibm.com/docs/account?topic=cli-ibmcloud_commands_resource#ibmcloud_resource_service_keys).

## Example usage

```terraform
data "ibm_resource_key" "resourceKeydata" {
  name                  = "myobjectKey"
  resource_instance_id  = ibm_resource_instance.resource.id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `most_recent` - (Optional, Bool) If there are multiple resource keys, you can set this argument to `true` to import only the most recently created key.
- `name` - (Required, String) The name of the resource key. You can retrieve the value by executing the `ibmcloud resource service-keys` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).
- `resource_instance_id` - (Optional, string) The ID of the resource instance that the resource key is associated with. You can retrieve the value by executing the `ibmcloud resource service-instances` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started). **Note**: Conflicts with `resource_alias_id`.
- `resource_alias_id` - (Optional, String) The ID of the resource alias that the resource key is associated with. You can retrieve the value by executing the `ibmcloud resource service-alias` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started). **Note** Conflicts with `resource_instance_id`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `credentials` - The credentials associated with the key.
- `id` - The unique identifier of the resource key.
- `role` - The user role.
- `status` - The status of the resource key.  

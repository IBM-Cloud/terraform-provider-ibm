---

subcategory: "Resource management"
layout: "ibm"
page_title: "IBM: ibm_resource_instance"
description: |-
  Get information about a resource instance from IBM Cloud.
---

# ibm_resource_instance
Retrieve information about an existing IBM resource instance from IBM Cloud as a read-only data source. For more information, about resource instance, see [ibmcloud resource service-instance](https://cloud.ibm.com/docs/account?topic=cli-ibmcloud_commands_resource#ibmcloud_resource_service_instance).

## Example usage

```terraform
data "ibm_resource_group" "group" {
  name = "default"
}

data "ibm_resource_instance" "testacc_ds_resource_instance" {
  name              = "myobjectstore"
  location          = "global"
  resource_group_id = data.ibm_resource_group.group.id
  service           = "cloud-object-storage"
}
data "ibm_resource_instance" "testacc_ds_resource_instance_identifier" {
  identifier        = ibm_resource_instance.instance.id
}
```

## Argument reference

The following arguments are supported:

- `identifier` - The GUID of the resource instance. Conflicts with other arguments.

- `location` - (Optional, String) The location or the environment in which the instance exists.
- `name` - (Optional, String) The name of the resource instance.
- `resource_group_id` - (Optional, String) The ID of the resource group where the resource instance exists. If not provided it takes the default resource group.
- `service` - (Optional, String) The service type of the instance. You can retrieve the value by executing the `ibmcloud catalog service-marketplace` or `ibmcloud catalog search` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `extensions` - (String) The extended metadata as a map associated with the resource instance.
- `guid`- (String) The GUID of the resource instance.
- `id` - (String) The unique identifier of the resource instance.
- `parameters_json` - (String) The parameters associated with the instance in json format.
- `plan` - (String) The plan for the service offering used by this resource instance.
- `status` - (String) The status of resource instance.
- `onetime_credentials` - (Bool) A boolean that dictates if the onetime_credentials is true or false.

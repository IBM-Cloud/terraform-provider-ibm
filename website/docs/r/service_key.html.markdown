---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM : service_key"
description: |-
  Manages IBM service key.
---

# `ibm_service_key`

Create, update, or delete a service key for your Cloud Foundry service instance. For more information, about creating organization, spaces, and a service key, see [Getting started with IBM Cloud Foundry Enterprise Environment](https://cloud.ibm.com/docs/cli?topic=cli-ibmcloud_commands_services#ibmcloud_service_key_create).


## Example usage
The following example creates the `mycloudantkey` service key. 


```
data "ibm_service_instance" "service_instance" {
  name = "mycloudant"
}

resource "ibm_service_key" "serviceKey" {
  name                  = "mycloudantkey"
  service_instance_guid = data.ibm_service_instance.service_instance.id
}
```


## Argument reference
Review the input parameters that you can specify for your resource. 

- `name` - (Required, String) A descriptive name for the service key.
- `parameters` (Optional, Map) Arbitrary parameters to pass along to the service broker. Must be a JSON object.
- `service_instance_guid` - (Required, String) The GUID of the service instance for which you create the service key.
- `tags` (Optional, Array of Strings) The tags that you want to add to the service key instance. Tags can help you find the service keys more easily later.

## Attribute reference
Review the output parameters that you can access after your resource is created. 

- `credentials` - (String) The credentials associated with the key.
- `id` - (String) The unique identifier of the new service key.


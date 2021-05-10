---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: ibm_service_key"
description: |-
  Get information about a service key from IBM Cloud.
---

# `ibm_service_key`

Retrieve information about existing service credentials that a Cloud Foundry service instance uses. For more information, about creating organization, spaces, and a service key, see [Getting started with IBM Cloud Foundry Enterprise Environment](https://cloud.ibm.com/docs/cli?topic=cli-ibmcloud_commands_services#ibmcloud_service_key_create).


## Example usage
The following example retrieves service key information for the `mycloudantdb` service instance. 


```
data "ibm_space" "space" {
  org   = "example.com"
  space = "dev"
}

data "ibm_service_key" "serviceKeydata" {
  name                  = "mycloudantdbKey"
  service_instance_name = "mycloudantdb"
  space_guid            = data.ibm_space.space.id
}
```

## Argument reference
Review the input parameters that you can specify for your data source. 

- `name` - (Required, String) The name of the service key. You can retrieve the value by running the `ibmcloud service keys` command in the IBM Cloud CLI.
- `service_instance_name` - (Required, String) The name of the service instance that the service key is associated with. You can retrieve the value by running the `ibmcloud service list` command in the IBM Cloud CLI.
- `space_guid` - (Required, String) The GUID of the Cloud Foundry space where the service instance exists. You can retrieve the value from the data source `ibm_space`.

## Attribute reference
Review the output parameters that you can access after you retrieved your data source. 

- `credentials` - (String) The credentials associated with the key.
- `id` - (String) The unique identifier of the service key.



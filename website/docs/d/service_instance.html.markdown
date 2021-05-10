---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: ibm_service_instance"
description: |-
  Get information about a service instance from IBM Cloud.
---

# `ibm_service_instance`

Retrieve information about a Cloud Foundry service instance. For more information, about creating organization, spaces, and an instance, see [Getting started with IBM Cloud Foundry Enterprise Environment](https://cloud.ibm.com/docs/cloud-foundry?topic=cloud-foundry-getting-started).


## Example usage
The following example retrieves information about the `mycloudantdb` instance. 


```
data "ibm_space" "space" {
  org   = "myorg"
  space = "dev"
}

data "ibm_service_instance" "serviceInstance" {
  name = "mycloudantdb"
  space_guid   = data.ibm_space.space.id
}
```

## Argument reference
Review the input parameters that you can specify for your data source. 

- `name` - (Required, String) The name of the service instance. You can retrieve the value by running the `ibmcloud service list` command.
- `space_guid` - (Required, String) The GUID of the Cloud Foundry space where the service instance is deployed to. You can retrieve the value from data source `ibm_space`.


## Attribute reference
Review the output parameters that you can access after you retrieved your data source. 

- `credentials` - (String) The credentials provided by the service broker to use this service.
- `id` - (String) The unique identifier of the service instance.
- `service_keys` - (String) The service keys associated with this service.
- `service_plan_guid` - (String) The plan GUID for the service offering used by this service instance.



---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM : service_instance"
description: |-
  Manages IBM service instance.
---

# `ibm_service_instance`

Create, update, or delete a Cloud Foundry service instance. For more information, about creating organization, spaces, and an instance, see [Getting started with IBM Cloud Foundry Enterprise Environment](https://cloud.ibm.com/docs/cloud-foundry?topic=cloud-foundry-getting-started).


## Example usage
The following example creates the `speech_to_text` Cloud Foundry service instance. 


```
data "ibm_space" "spacedata" {
  space = "prod"
  org   = "myorg.com"
}

resource "ibm_service_instance" "service_instance" {
  name       = "myspeech"
  space_guid = data.ibm_space.spacedata.id
  service    = "speech_to_text"
  plan       = "lite"
  tags       = ["cluster-service", "cluster-bind"]
}
```


## Argument reference
Review the input parameters that you can specify for your resource. 

- `name` - (Required, String) A descriptive name for the service instance.
- `parameters` (Optional, Map)  Arbitrary parameters to pass to the service broker. The value must be a JSON object.
- `plan` - (Required, String) The name of the service plan that you want. You can retrieve the value by running the `ibmcloud service offerings` command in the IBM Cloud CLI.
- `service` - (Required, String) The name of the service offering. You can retrieve the value by running the `ibmcloud service offerings` command in the IBM Cloud CLI.
- `space_guid` - (Required, String) The GUID of the Cloud Foundry space where you want to create the service. You can retrieve the value from data source `ibm_space`.
- `tags` (Array of Strings) Optional- The tags that you want to add to the service instance. Tags can help you to find the instance more easily later.
- `wait_time_minutes` - (Optional, Integer) The number of minutes to wait for the service instance to become available before declaring it as created. The same number of minutes is used for the deletion to finish. The default value is `10`.


## Attribute reference
Review the output parameters that you can access after your resource is created. 

- `credentials` - (Map) The credentials provided by the service broker to use the service.
- `dashboard_url`- (String) The dashboard URL of the new service instance.
- `id` - (String) The unique identifier of the new service instance.
- `service_keys` - (String) The service keys associated with the service.
- `service_plan_guid` - (String) The plan of the service offering used by this service instance.



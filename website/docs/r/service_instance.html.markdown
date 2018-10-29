---
layout: "ibm"
page_title: "IBM : service_instance"
sidebar_current: "docs-ibm-resource-service-instance"
description: |-
  Manages IBM Service Instance.
---

# ibm\_service_instance

Provides a service instance resource. This allows service instances to be created, updated, and deleted.

## Example Usage

```hcl
data "ibm_space" "spacedata" {
  space = "prod"
  org   = "somexample.com"
}

resource "ibm_service_instance" "service_instance" {
  name       = "test"
  space_guid = "${data.ibm_space.spacedata.id}"
  service    = "speech_to_text"
  plan       = "lite"
  tags       = ["cluster-service", "cluster-bind"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) A descriptive name used to identify the service instance.
* `space_guid` - (Required, string) The GUID of the space where you want to create the service. You can retrieve the value from data source `ibm_space`.
* `service` - (Required, string) The name of the service offering. You can retrieve the value by running the `ibmcloud service offerings` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `plan` - (Required, string) The name of the plan type supported by service. You can retrieve the value by running the `ibmcloud service offerings` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `tags` - (Optional, array of strings) Tags associated with the public IP instance.
* `parameters` - (Optional, map) Arbitrary parameters to pass to the service broker. The value must be a JSON object.
* `wait_time_minutes` - (Optional, integer) The duration, expressed in minutes, to wait for the service instance to become available before declaring it as created. It is also the same amount of time waited for deletion to finish. The default value is `10`.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the new service instance.
* `credentials` - The credentials provided by the service broker to use the service.
* `service_keys` - The service keys associated with the service.
* `service_plan_guid` - The plan of the service offering used by this service instance.

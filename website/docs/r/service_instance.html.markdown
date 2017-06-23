---
layout: "ibm"
page_title: "IBM : service_instance"
sidebar_current: "docs-ibm-resource-service-instance"
description: |-
  Manages IBM Service Instance.
---

# ibm\_service_instance

Crate, update, or delete service instances on IBM Bluemix.

## Example Usage

```hcl
data "ibm_space" "spacedata" {
  space = "prod"
  org   = "somexample.com"
}

resource "ibm_service_instance" "service_instance" {
  name       = "test"
  space_guid = "${data.ibm_space.spacedata.id}"
  service    = "cloudantNoSQLDB"
  plan       = "Lite"
  tags       = ["cluster-service", "cluster-bind"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) A descriptive name used to identify the service instance.
* `space_guid` - (Required, string) The GUID of the space where you want to create the service. The values can be retrieved from data source `ibm_space`.
* `service` - (Required, string) The name of the service offering. The value can be retrieved by running the `bx service offerings` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).
* `plan` - (Required, string) The name of the plan type supported by service. The value can be retrieved by running the `bx service offerings` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).
* `tags` - (Optional, list) User-provided tags.
* `parameters` - (Optional, map) Arbitrary parameters to pass along to the service broker. Must be a JSON object.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the new service instance.
* `credentials` - The credentials associated with the service instance.
* `service_plan_guid` - The plan of the service offering used by this service instance 

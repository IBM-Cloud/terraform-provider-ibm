---
layout: "ibm"
page_title: "IBM: ibm_service_instance"
sidebar_current: "docs-ibm-datasource-service-instance"
description: |-
  Get information about a service instance from IBM Cloud.
---

# ibm\_service_instance

Import the details of an existing IBM service instance from IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_space" "space" {
  org   = "example.com"
  space = "dev"
}

data "ibm_service_instance" "serviceInstance" {
  name = "mycloudantdb"
  space_guid   = "${data.ibm_space.space.id}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the service instance. You can retrieve the value by running the `ibmcloud service list` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).

* `space_guid` - (Required, string) The GUID of the space where the service instance exists. You can retrieve the value from data source `ibm_space`.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the service instance.
* `credentials` - The credentials provided by the service broker to use this service.
* `service_keys` - The service keys associated with this service.
* `service_plan_guid` - The plan GUID for the service offering used by this service instance.

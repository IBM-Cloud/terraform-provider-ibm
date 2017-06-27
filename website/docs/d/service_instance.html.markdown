---
layout: "ibm"
page_title: "IBM: ibm_service_instance"
sidebar_current: "docs-ibm-datasource-service-instance"
description: |-
  Get information about a service instance from IBM Bluemix.
---

# ibm\_service_instance

Import the details of an existing IBM service instance from IBM Bluemix as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration by using interpolation syntax. 

## Example Usage

```hcl
data "ibm_service_instance" "serviceInstance" {
  name = "mycloudantdb"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the service instance. The value can be retrieved by running the `bx service list` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).

## Attributes Reference

The following attributes are exported:

* `id` - The unique identifier of the service instance. 
* `credentials` - The service broker-provided credentials to use this service.
* `service_keys` - The service keys associated with this service.
* `service_plan_guid` - The plan of the service offering used by this service instance.

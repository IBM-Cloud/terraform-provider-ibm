---
layout: "ibm"
page_title: "IBM: ibm_service_key"
sidebar_current: "docs-ibm-datasource-service-key"
description: |-
  Get information about a service key from IBM Bluemix.
---

# ibm\_service_key

Import the details of an existing IBM service key from IBM Bluemix as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration by using interpolation syntax. 

## Example Usage

```hcl
data "ibm_service_key" "serviceKeydata" {
  name                  = "mycloudantdbKey"
  service_instance_name = "mycloudantdb"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the service key. The value can be retrieved by running the `bx service keys` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).
* `service_instance_name` - (Required) The name of the service instance that the service key is associated with. The value can be retrieved by running the `bx service list` command in the Bluemix CLI.

## Attributes Reference

The following attributes are exported:

* `credentials` - The credentials associated with the key.  

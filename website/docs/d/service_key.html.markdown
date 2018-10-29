---
layout: "ibm"
page_title: "IBM: ibm_service_key"
sidebar_current: "docs-ibm-datasource-service-key"
description: |-
  Get information about a service key from IBM Cloud.
---

# ibm\_service_key

Import the details of an existing IBM service key from IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_space" "space" {
  org   = "example.com"
  space = "dev"
}

data "ibm_service_key" "serviceKeydata" {
  name                  = "mycloudantdbKey"
  service_instance_name = "mycloudantdb"
  space_guid            = "${data.ibm_space.space.id}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the service key. You can retrieve the value by running the `ibmcloud service keys` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `service_instance_name` - (Required, string) The name of the service instance that the service key is associated with. You can retrieve the value by running the `ibmcloud service list` command in the IBM Cloud CLI.
* `space_guid` - (Required, string) The GUID of the space where the service instance exists. You can retrieve the value from the data source `ibm_space`.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the service key.
* `credentials` - The credentials associated with the key.  

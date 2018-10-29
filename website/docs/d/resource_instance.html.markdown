---
layout: "ibm"
page_title: "IBM: ibm_resource_instance"
sidebar_current: "docs-ibm-datasource-resource-instance"
description: |-
  Get information about a resource instance from IBM Cloud.
---

# ibm\_resource_instance

Import the details of an existing IBM resource instance from IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_resource_group" "group" {
  name = "default"
}

data "ibm_resource_instance" "testacc_ds_resource_instance" {
  name              = "myobjectstore"
  location          = "global"
  resource_group_id = "${data.ibm_resource_group.group.id}"
  service           = "cloud-object-storage"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the resource instance.

* `resource_group_id` - (Optional, string) The id of the resource group where the resource instance exists. If not provided it takes the default resource group.

* `location` - (Optional, string) The location or the environment in which instance exists.

* `service` - (Optional, string) The service type of the instance. You can retrieve the value by running the `ibmcloud catalog service-marketplace` or `ibmcloud catalog search` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the resource instance.
* `status` - The status of resource instance.
* `plan` - The plan for the service offering used by this resource instance.

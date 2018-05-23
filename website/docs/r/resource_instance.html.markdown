---
layout: "ibm"
page_title: "IBM : resource_instance"
sidebar_current: "docs-ibm-resource-resource-instance"
description: |-
  Manages IBM Resource Instance.
---

# ibm\_resource_instance

Provides a Resource Instance resource. This allows Resource Instances to be created, updated, and deleted.

## Example Usage

```hcl
data "ibm_resource_group" "group" {
  name = "test"
}

resource "ibm_resource_instance" "resource_instance" {
  name              = "test"
  service           = "cloud-object-storage"
  plan              = "lite"
  location          = "global"
  resource_group_id = "${data.ibm_resource_group.group.id}"
  tags              = ["tag1", "tag2"]

  parameters = {
    "HMAC" = true
  }
  //User can increase timeouts 
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}
```

## Timeouts

ibm_resource_instance provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for Creating Instance.
* `update` - (Default 10 minutes) Used for Updating Instance.
* `delete` - (Default 10 minutes) Used for Deleting Instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required, string) A descriptive name used to identify the resource instance.
* `service` - (Required, string) The name of the service offering. You can retrieve the value by running the `bx catalog service-marketplace` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `plan` - (Required, string) The name of the plan type supported by service. You can retrieve the value by running the `bx catalog service <servicename>` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `location` - (Required, string) Target location or environment to create the resource instance.
* `resource_group_id` - (Optional, string) The ID of the resource group where you want to create the service. You can retrieve the value from data source `ibm_resource_group`. If not provided creates the service in default resource group.
* `tags` - (Optional, array of strings) Tags associated with the instance.
* `parameters` - (Optional, map) Arbitrary parameters to create instance. The value must be a JSON object.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the new resource instance.
* `status` - Status of resource instance.

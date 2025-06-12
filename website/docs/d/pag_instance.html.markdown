---
layout: "ibm"
page_title: "IBM : ibm_pag_instance"
description: |-
  Get information about Privileged Access Gateway instance.
subcategory: "Privileged Access Gateway"
---

# ibm_pag_instance
Retrieve information about an existing IBM Privileged Access Gateway (PAG) instance from IBM Cloud as a read-only data source. 

## Example usage

```terraform
data "ibm_resource_group" "group" {
  name = "default"
}

data "ibm_pag_instance" "testacc_ds_pag_instance" {
  name              = "myPagInstance"
  resource_group_id = data.ibm_resource_group.group.id
  service           = "privileged-access-gateway"
}
```

## Argument reference

The following arguments are supported:

- `location` - (Optional, String) The location or the environment in which the PAG instance exists.
- `name` - (Required, String) The name of the PAG instance.
- `resource_group_id` - (Optional, String) The ID of the resource group where the PAG instance exists. If not provided it takes the default resource group.
- `service` - (Required, String) The service type of the PAG instance.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `guid`- (String) The GUID of the resource instance.
- `id` - (String) The unique identifier of the resource instance.
- `parameters_json` - (String) The parameters associated with the instance in json format.
- `plan` - (String) The plan for the service offering used by this resource instance.
- `status` - (String) The status of resource instance.

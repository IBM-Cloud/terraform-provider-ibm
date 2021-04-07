---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_api_key_reset"
description: |-
  Resets Kubernetes API Key.
---

# ibm\_container_api_key_reset

Resets Kubernetes API Key

## Example Usage

In the following example, you can reset kubernetes api key:

```hcl
resource "ibm_container_api_key_reset" "reset" {
    region ="us-east"
    resource_group_id = "766f3584b2c840ee96d856bc04551da8"
    reset_api_key=2

}

```

## Argument Reference

The following arguments are supported:

* `region` - (Required, Forces new resource, string) The region in which api key has to be reset.
* `resource_group_id` - (Optional, Forces new resource, string) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`. If not provided defaults to default resource group.
* `reset_api_key` - (Optional, int) Default: 1. This attribute determines if the apikey has to be reset or not. Inorder to avoid state dependencies this attribute has been added. To reset apikey on same region and resource_group_id this attribute has to be incremented.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The Resource ID. Id is a combination of `<region>/<resource_group_id>`
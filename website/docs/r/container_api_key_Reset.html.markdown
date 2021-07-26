---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_api_key_reset"
description: |-
  Resets Kubernetes API Key.
---

# ibm_container_api_key_reset
Create, update, or delete Kubernetes API key. For more information, about Kubernetes API key, see [assigning cluster access](https://cloud.ibm.com/docs/containers?topic=containers-users#access-checklist).

## Example usage
In the following example, you can reset kubernetes api key:

```terraform
resource "ibm_container_api_key_reset" "reset" {
    region ="us-east"
    resource_group_id = "766f3584b2c840ee96d856bc04551da8"
    reset_api_key=2

}

```

## Argument reference
Review the argument references that you can specify for your resource. 

- `region` - (Required, Forces new resource, String) The region in which API key has to be reset.
- `resource_group_id` - (Optional, Forces new resource, String) The ID of the resource group. You can retrieve the value from data source `ibm_resource_group`. If not provided defaults to default resource group.
- `reset_api_key`  - (Optional, Integer) Determines the API key need reset or not. This attribute is added to avoid the state dependencies. You need to increment the attribute to reset the API key on same `region` and `resource_group_id`. The default value is `1`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The resource ID. ID is a combination of `<region>/<resource_group_id>`.

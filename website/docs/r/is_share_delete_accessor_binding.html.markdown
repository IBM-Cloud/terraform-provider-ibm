---
layout: "ibm"
page_title: "IBM : is_share_delete_accessor_binding"
description: |-
  Delete a share accessor binding.
subcategory: "VPC infrastructure"
---

# is_share_accessor_binding_operations

Provides a resource for managing the share accessor binding operations like delete binding.

~> **NOTE**
`ibm_share_delete_accessor_binding` is used for deleting share accessor binding


## Example Usage

```terraform
resource "ibm_is_share" "example" {
  name = "my-share"
  size = 200
  profile = "dp2"
  zone = "us-south-2"
}
```
## Example Usage (Create a accessor share)

```terraform
resource "ibm_is_share" "example1" {
    origin_share {
      crn = ibm_is_share.example.crn
    }
    name = "my-replica1"
}
```

## Example Usage

```hcl
// list bindings
data "ibm_is_share_accessor_bindings" "example" {
  share = ibm_is_share.example.id
}

resource "ibm_share_delete_accessor_binding" "example" {
    share = ibm_is_share.example.id
    accessor_binding = data.ibm_is_share_accessor_bindings.example.accessor_bindings.0.id
}
```

## Argument Reference

The following arguments are supported:

- `share` - (Required, string) The file share identifier.
- `accessor_binding` - (Required, string) The share accessor binding ID

## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the Share.

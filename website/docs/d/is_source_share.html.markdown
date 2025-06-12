---
layout: "ibm"
page_title: "IBM : is_source_share"
description: |-
  Get information about Share
subcategory: "VPC infrastructure"
---

# ibm_is_source_share

Provides a read-only data source for the source share. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}
resource "ibm_is_share" "example" {
  name    = "example-share"
  size    = 200
  profile = "dp2"
  zone    = "us-south-2"
}
resource "ibm_is_share" "example1" {
  zone                  = "us-south-3"
  source_share          = ibm_is_share.example.id
  name                  = "my-replica1"
  profile               = "tier-3iops"
  replication_cron_spec = "0 */5 * * *"
}
data "ibm_is_source_share" "example" {
  share_replica = ibm_is_share.example1.id
}
```

## Argument Reference

The following arguments are supported:

- `share_replica` - (Optional, String) The file share identifier.

**Note** One of the aurgument is mandatory

## Attribute Reference

The following attributes are exported:

- `crn` - The CRN for this share.
- `href` - The URL for this share.
- `id` - (String) The ID of the file share.
- `name` - The unique user-defined name for this file share.
- `resource_type` - The type of resource referenced.

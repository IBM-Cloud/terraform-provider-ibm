---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : ibm_satellite_link"
description: |-
  Manages satellite link.
---

# ibm_satellite_link

Provides a resource for ibm_satellite_link. This allows ibm_satellite_link to be created, updated and deleted.

## Example usage

```terraform
resource "ibm_satellite_link" "satellite_link" {
  crn = "crn:v1:staging:public:satellite:us-south:a/1ae4eb57181a46ceade4846519678888::location:brbats7009sqna3dtest"
  location_id = "brbats7009sqna3dtest"
}
```

## Argument reference

The following arguments are supported:

* `crn` - (Required, string) CRN of the Location.
* `location` - (Required, string) Location ID.

## Attribute reference

In addition to all arguments above, the following attributes are exported:

* `created_at` - Timestamp of creation of location.
* `description` - Description of the location.
* `id` - The unique identifier of the ibm_satellite_link.
* `last_change` - Timestamp of latest modification of location.
  * Constraints: Allowable values are: enabled, disabled
* `performance` - The last performance data of the Location.
* `satellite_link_host` - Satellite Link hostname of the location.
* `status` - Enabled/Disabled.
* `ws_endpoint` - The ws endpoint of the location.

## Import

You can import the `ibm_satellite_link` resource by using `location`. Unique identifier for this location.

```
$ terraform import ibm_satellite_link.satellite_link brbats7009sqna3dtest
```

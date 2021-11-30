---
layout: "ibm"
page_title: "IBM : ibm_cbr_zone"
description: |-
  Manages cbr_zone.
subcategory: "Context Based Restrictions"
---

# ibm_cbr_zone

Provides a resource for cbr_zone. This allows cbr_zone to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cbr_zone" "cbr_zone" {
	name = "Test Zone Resource Config Basic"
	description = "Test Zone Resource Config Basic"
	addresses {
		type = "ipRange"
		value = "169.23.22.0-169.23.22.255"
	}
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `addresses` - (Required, List) The list of addresses in the zone.
  * Constraints: The maximum length is `1000` items. The minimum length is `1` item.
Nested scheme for **addresses**:
    * `ref` - (Optional, List) A service reference value.
    Nested scheme for **ref**:
        * `service_instance` - (Optional, String) The service instance.
          * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `^[0-9a-z\-\/]+$`.
        * `service_name` - (Optional, String) The service name.
          * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `^[0-9a-z\-]+$`.
        * `service_type` - (required, String) The service type.
          * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `^[0-9a-z_]+$`.
    * `type` - (required, String) The type of address.
      * Constraints: Allowable values are: `ipAddress`, `ipRange`, `subnet`, `vpc`, `serviceRef`.
    * `value` - (required, String) The IP address.
      * Constraints: The maximum length is `45` characters. The minimum length is `7` characters. The value must match regular expression `^[a-zA-Z0-9:.]+$`.
* `description` - (Optional, String) The description of the zone.
  * Constraints: The maximum length is `300` characters. The minimum length is `0` characters. The value must match regular expression `^[\x20-\xFE]*$`.
* `excluded` - (Optional, List) The list of excluded addresses in the zone. Only addresses of type `ipAddress`, `ipRange`, and `subnet` can be excluded.
  * Constraints: The maximum length is `1000` items.
Nested scheme for **excluded**:
    * `ref` - (Optional, List) A service reference value.
    Nested scheme for **ref**:
        * `service_instance` - (Optional, String) The service instance.
          * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `^[0-9a-z\-/]+$`.
        * `service_name` - (Optional, String) The service name.
          * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `^[0-9a-z\-]+$`.
        * `service_type` - (required, String) The service type.
          * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `^[0-9a-z_]+$`.
    * `type` - (required, String) The type of address.
      * Constraints: Allowable values are: `ipAddress`, `ipRange`, `subnet`, `vpc`, `serviceRef`.
    * `value` - (required, String) The IP address.
      * Constraints: The maximum length is `45` characters. The minimum length is `7` characters. The value must match regular expression `/^[a-zA-Z0-9:.]+$/`.
* `name` - (required, String) The name of the zone.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 \\-_]+$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cbr_zone.
* `address_count` - (Integer) The number of addresses in the zone.
* `created_at` - (String) The time the resource was created.
* `created_by_id` - (String) IAM ID of the user or service which created the resource.
* `crn` - (String) The zone CRN.
* `excluded_count` - (Integer) The number of excluded addresses in the zone.
* `href` - (String) The href link to the resource.
* `last_modified_at` - (String) The last time the resource was modified.
* `last_modified_by_id` - (String) IAM ID of the user or service which modified the resource.

* `version` - Version of the cbr_zone.

## Import

You can import the `ibm_cbr_zone` resource by using `id`. The globally unique ID of the zone.

# Syntax
```
$ terraform import ibm_cbr_zone.cbr_zone <id>
```

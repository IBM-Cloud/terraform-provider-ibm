---
layout: "ibm"
page_title: "IBM : ibm_cbr_zone"
description: |-
  Manages cbr_zone.
subcategory: "Context Based Restrictions"
---

# ibm_cbr_zone

Create, update, and delete cbr_zones with this resource.

## Example Usage

```hcl
resource "ibm_cbr_zone" "cbr_zone_instance" {
  account_id = "12ab34cd56ef78ab90cd12ef34ab56cd"
  addresses {
    type = "ipAddress"
    value = "169.23.56.234"
  }
  addresses {
    type = "ipRange"
    value = "169.23.22.0-169.23.22.255"
  }
  excluded {
    type  = "ipAddress"
    value = "169.23.22.10"
  }
  excluded {
    type  = "ipAddress"
    value = "169.23.22.11"
  }
  description = "this is an example of zone"
  excluded {
		type = "ipAddress"
		value = "value"
  }
  name = "an example of zone"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `account_id` - (Optional, String) The id of the account owning this zone.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-]+$/`.
* `addresses` - (Optional, List) The list of addresses in the zone.
  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
Nested schema for **addresses**:
	* `ref` - (Optional, List) A service reference value.
	Nested schema for **ref**:
		* `account_id` - (Required, String) The id of the account owning the service.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-]+$/`.
		* `location` - (Optional, String) The location.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\-]+$/`.
		* `service_instance` - (Optional, String) The service instance.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\-\/]+$/`.
		* `service_name` - (Optional, String) The service name.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\-]+$/`.
		* `service_type` - (Optional, String) The service type.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z_]+$/`.
	* `type` - (Optional, String) The type of address.
	  * Constraints: Allowable values are: `ipAddress`, `ipRange`, `subnet`, `vpc`, `serviceRef`.
	* `value` - (Optional, String) The IP address.
	  * Constraints: The maximum length is `45` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9:.]+$/`.
* `description` - (Optional, String) The description of the zone.
  * Constraints: The maximum length is `300` characters. The minimum length is `0` characters. The value must match regular expression `/^[\x20-\xFE]*$/`.
* `excluded` - (Optional, List) The list of excluded addresses in the zone. Only addresses of type `ipAddress`, `ipRange`, and `subnet` can be excluded.
  * Constraints: The maximum length is `1000` items.
Nested schema for **excluded**:
	* `ref` - (Optional, List) A service reference value.
	Nested schema for **ref**:
		* `account_id` - (Required, String) The id of the account owning the service.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-]+$/`.
		* `location` - (Optional, String) The location.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\-]+$/`.
		* `service_instance` - (Optional, String) The service instance.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\-/]+$/`.
		* `service_name` - (Optional, String) The service name.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\-]+$/`.
		* `service_type` - (Optional, String) The service type.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z_]+$/`.
	* `type` - (Optional, String) The type of address.
	  * Constraints: Allowable values are: `ipAddress`, `ipRange`, `subnet`, `vpc`, `serviceRef`.
	* `value` - (Optional, String) The IP address.
	  * Constraints: The maximum length is `45` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9:.]+$/`.
* `name` - (Optional, String) The name of the zone.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 \-_]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the cbr_zone.
* `address_count` - (Integer) The number of addresses in the zone.
* `created_at` - (String) The time the resource was created.
* `created_by_id` - (String) IAM ID of the user or service which created the resource.
* `crn` - (String) The zone CRN.
* `excluded_count` - (Integer) The number of excluded addresses in the zone.
* `href` - (String) The href link to the resource.
* `last_modified_at` - (String) The last time the resource was modified.
* `last_modified_by_id` - (String) IAM ID of the user or service which modified the resource.

* `etag` - ETag identifier for cbr_zone.

## Import

You can import the `ibm_cbr_zone` resource by using `id`. The globally unique ID of the zone.

# Syntax
<pre>
$ terraform import ibm_cbr_zone.cbr_zone &lt;id&gt;
</pre>

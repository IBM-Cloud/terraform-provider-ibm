---
layout: "ibm"
page_title: "IBM : ibm_cbr_zone"
description: |-
  Get information about cbr_zone
subcategory: "Context Based Restrictions"
---

# ibm_cbr_zone

Provides a read-only data source to retrieve information about a cbr_zone. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_cbr_zone" "cbr_zone" {
	zone_id = "zone_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `zone_id` - (Required, Forces new resource, String) The ID of a zone.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[a-fA-F0-9]{32}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the cbr_zone.
* `account_id` - (String) The id of the account owning this zone.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-]+$/`.
* `address_count` - (Integer) The number of addresses in the zone.
* `addresses` - (List) The list of addresses in the zone.
  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
Nested schema for **addresses**:
	* `ref` - (List) A service reference value.
	Nested schema for **ref**:
		* `account_id` - (String) The id of the account owning the service.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-]+$/`.
		* `location` - (String) The location.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\-]+$/`.
		* `service_instance` - (String) The service instance.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\-/]+$/`.
		* `service_name` - (String) The service name.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\-]+$/`.
		* `service_type` - (String) The service type.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z_]+$/`.
	* `type` - (String) The type of address.
	  * Constraints: Allowable values are: `ipAddress`, `ipRange`, `subnet`, `vpc`, `serviceRef`.
	* `value` - (String) The IP address.
	  * Constraints: The maximum length is `45` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9:.]+$/`.
* `created_at` - (String) The time the resource was created.
* `created_by_id` - (String) IAM ID of the user or service which created the resource.
* `crn` - (String) The zone CRN.
* `description` - (String) The description of the zone.
  * Constraints: The maximum length is `300` characters. The minimum length is `0` characters. The value must match regular expression `/^[\x20-\xFE]*$/`.
* `excluded` - (List) The list of excluded addresses in the zone. Only addresses of type `ipAddress`, `ipRange`, and `subnet` can be excluded.
  * Constraints: The maximum length is `1000` items.
Nested schema for **excluded**:
	* `id` - (String) The address id (for use by terraform only).
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character.
	* `ref` - (List) A service reference value.
	Nested schema for **ref**:
		* `account_id` - (String) The id of the account owning the service.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-]+$/`.
		* `location` - (String) The location.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\-]+$/`.
		* `service_instance` - (String) The service instance.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\-/]+$/`.
		* `service_name` - (String) The service name.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\-]+$/`.
		* `service_type` - (String) The service type.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z_]+$/`.
	* `type` - (String) The type of address.
	  * Constraints: Allowable values are: `ipAddress`, `ipRange`, `subnet`, `vpc`, `serviceRef`.
	* `value` - (String) The IP address.
	  * Constraints: The maximum length is `45` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9:.]+$/`.
* `excluded_count` - (Integer) The number of excluded addresses in the zone.
* `href` - (String) The href link to the resource.
* `id` - (String) The globally unique ID of the zone.
* `last_modified_at` - (String) The last time the resource was modified.
* `last_modified_by_id` - (String) IAM ID of the user or service which modified the resource.
* `name` - (String) The name of the zone.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 \-_]+$/`.


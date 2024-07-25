---
layout: "ibm"
page_title: "IBM : ibm_cbr_zone_addresses"
description: |-
  Get information about cbr_zone_addresses
subcategory: "Context Based Restrictions"
---

# ibm_cbr_zone_addresses

Provides a read-only data source for cbr_zone_addresses. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cbr_zone_addresses" "cbr_zone_addresses" {
	zone_addresses_id = "zone_addresses_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `zone_addresses_id` - (Required, Forces new resource, String) The ID of a zone addresses resource.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the cbr_zone_addresses.

* `zone_id` - (String) The id of the zone in which the addresses are included.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[a-fA-F0-9]{32}$/`.

* `addresses` - (List) The list of addresses included in the zone.
  * Constraints: The maximum length is `1000` items. The minimum length is `1` items.
Nested scheme for **addresses**:
    * `ref` - (List) A service reference value.
    Nested scheme for **ref**:
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


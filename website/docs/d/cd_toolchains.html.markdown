---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchains"
description: |-
  Get information about cd_toolchains
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchains

Provides a read-only data source to retrieve information about cd_toolchains. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

~> **Warning:** Continuous Delivery (CD) will be discontinued in these regions on 12 February 2027: `au-syd`, `ca-mon`, `ca-tor`, `eu-es`, `jp-osa`, `us-east`. Follow the [migration guide](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-cd-migrate-region) to avoid disruption. [Learn more](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-cd-migrate-region)

## Example Usage

```hcl
data "ibm_cd_toolchains" "cd_toolchains" {
	name = "TestToolchainV2"
	resource_group_id = "6a9a01f2cff54a7f966f803d92877123"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `name` - (Optional, String) Exact name of toolchain to look up. This parameter is case sensitive.
  * Constraints: The maximum length is `128` characters. The minimum length is `0` characters. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9-._ ])+$/`.
* `resource_group_id` - (Required, String) The resource group ID where the toolchains exist.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[0-9a-f]{32}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the cd_toolchains.
* `toolchains` - (List) Toolchain results returned from the collection.
  * Constraints: The maximum length is `200` items. The minimum length is `0` items.
Nested schema for **toolchains**:
	* `account_id` - (String) Account ID where toolchain can be found.
	* `created_at` - (String) Toolchain creation timestamp.
	* `created_by` - (String) Identity that created the toolchain.
	* `crn` - (String) Toolchain CRN.
	* `description` - (String) Describes the toolchain.
	  * Constraints: The maximum length is `500` characters. The minimum length is `0` characters. The value must match regular expression `/^(.*?)$/`.
	* `href` - (String) URI that can be used to retrieve toolchain.
	* `id` - (String) Toolchain ID.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.
	* `location` - (String) Toolchain region.
	* `name` - (String) Toolchain name.
	  * Constraints: The maximum length is `128` characters. The minimum length is `0` characters. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9-._ ])+$/`.
	* `resource_group_id` - (String) Resource group where the toolchain is located.
	  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[0-9a-f]{32}$/`.
	* `ui_href` - (String) URL of a user-facing user interface for this toolchain.
	* `updated_at` - (String) Latest toolchain update timestamp.


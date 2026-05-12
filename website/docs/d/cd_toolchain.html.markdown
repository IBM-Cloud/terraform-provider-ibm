---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain"
description: |-
  Get information about cd_toolchain
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain

Provides a read-only data source to retrieve information about a cd_toolchain. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_cd_toolchain" "cd_toolchain" {
	toolchain_id = "b7c3f7f6-0613-4a2c-a4b0-3189ac95bdc9"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the cd_toolchain.
* `account_id` - (String) Account ID where toolchain can be found.

* `created_at` - (String) Toolchain creation timestamp.

* `created_by` - (String) Identity that created the toolchain.

* `crn` - (String) Toolchain CRN.

* `description` - (String) Describes the toolchain.
  * Constraints: The maximum length is `500` characters. The minimum length is `0` characters. The value must match regular expression `/^(.*?)$/`.

* `href` - (String) URI that can be used to retrieve toolchain.

* `location` - (String) Toolchain region.

* `name` - (String) Toolchain name.
  * Constraints: The maximum length is `128` characters. The minimum length is `0` characters. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9-._ ])+$/`.

* `resource_group_id` - (Forces new resource, String) Resource group where the toolchain is located.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[0-9a-f]{32}$/`.

* `ui_href` - (String) URL of a user-facing user interface for this toolchain.

* `updated_at` - (String) Latest toolchain update timestamp.

* `tags` - (Array of Strings) Tags associated with the toolchain.


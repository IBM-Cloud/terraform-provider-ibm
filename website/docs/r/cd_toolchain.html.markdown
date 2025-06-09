---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain"
description: |-
  Manages cd_toolchain.
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain

Create, update, and delete cd_toolchains with this resource.

## Example Usage

```hcl
resource "ibm_cd_toolchain" "cd_toolchain_instance" {
  description = "A sample toolchain to test the API"
  name = "TestToolchainV2"
  resource_group_id = "6a9a01f2cff54a7f966f803d92877123"
  tags = ["tag1", "tag2"]
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `description` - (Optional, String) Describes the toolchain.
  * Constraints: The maximum length is `500` characters. The minimum length is `0` characters. The value must match regular expression `/^(.*?)$/`.
* `name` - (Required, String) Toolchain name.
  * Constraints: The maximum length is `128` characters. The minimum length is `0` characters. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9-._ ])+$/`.
* `resource_group_id` - (Required, Forces new resource, String) Resource group where the toolchain is located.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[0-9a-f]{32}$/`.
* `tags` - (Optional, Array of Strings) Tags associated with the toolchain.


## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the cd_toolchain.
* `account_id` - (String) Account ID where toolchain can be found.
* `created_at` - (String) Toolchain creation timestamp.
* `created_by` - (String) Identity that created the toolchain.
* `crn` - (String) Toolchain CRN.
* `href` - (String) URI that can be used to retrieve toolchain.
* `location` - (String) Toolchain region.
* `ui_href` - (String) URL of a user-facing user interface for this toolchain.
* `updated_at` - (String) Latest toolchain update timestamp.


## Import

You can import the `ibm_cd_toolchain` resource by using `id`. Toolchain ID.

# Syntax
<pre>
$ terraform import ibm_cd_toolchain.cd_toolchain &lt;id&gt;
</pre>

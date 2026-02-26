---
layout: "ibm"
page_title: "IBM : ibm_pha_api_key"
description: |-
  Manages pha_api_key.
subcategory: "PowerhaAutomation Service"
---

# ibm_pha_api_key

Create, update, and delete pha_api_keys with this resource.

## Example Usage

```hcl
resource "ibm_pha_api_key" "pha_api_key_instance" {
  accept_language = "en-US"
  api_key = "adfadfdsafsdfdsf"
  if_none_match = "abcdef"
  pha_instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `accept_language` - (Optional, Forces new resource, String) The language requested for the return document.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `api_key` - (Optional, Forces new resource, String) The API key associated with the request.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._: -]+$/`.
* `if_none_match` - (Optional, Forces new resource, String) ETag for conditional requests (optional).
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `pha_instance_id` - (Required, Forces new resource, String) instance id of instance to provision.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the pha_api_key.
* `pha_instance_id` - (String) Unique identifier for the API key record.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._: -]+$/`.

* `etag` - ETag identifier for pha_api_key.

## Import

You can import the `ibm_pha_api_key` resource by using `id`.
The `id` property can be formed from `pha_instance_id`, and `pha_instance_id` in the following format:

<pre>
&lt;pha_instance_id&gt;/&lt;pha_instance_id&gt;
</pre>
* `pha_instance_id`: A string in the format `8eefautr-4c02-0009-0086-8bd4d8cf61b6`. instance id of instance to provision.
* `pha_instance_id`: A string in the format `9676767890`. Unique identifier for the API key record.

# Syntax
<pre>
$ terraform import ibm_pha_api_key.pha_api_key &lt;pha_instance_id&gt;/&lt;pha_instance_id&gt;
</pre>

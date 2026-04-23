---
layout: "ibm"
page_title: "IBM : ibm_pha_api_key"
description: |-
  Manages pha_api_key.
subcategory: "PowerhaAutomation Service"
---

# ibm_pha_api_key

Add and retrieve PHA API keys for the specified PowerHA instance.

## Example Usage

```hcl
resource "ibm_pha_api_key" "pha_api_key_instance" {
  accept_language = "en-US"
  api_key = "adfadfdsafsdfdsf"
  if_none_match = "abcdef"
  instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `accept_language` - (Optional, String) The language requested for the return document.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `api_key` - (Optional, String) The API key associated with the request.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._: -]+$/`.
* `if_none_match` - (Optional, String) ETag for conditional requests (optional).
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `instance_id` - (Required, String) Unique identifier of the provisioned instance.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the pha_api_key.
* `instance_id` - (String) Unique identifier for the API key record.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._: -]+$/`.

* `etag` - ETag identifier for pha_api_key.

## Import

Import is not supported

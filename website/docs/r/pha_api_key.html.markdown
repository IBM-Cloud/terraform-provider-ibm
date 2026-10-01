---
layout: "ibm"
page_title: "IBM : ibm_pha_api_key"
description: |-
  Manages pha_api_key.
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
---

# ibm_pha_api_key

Adds an API key to this resource and associates it with a specific PowerHA deployment.

## Example Usage

```hcl
resource "ibm_pha_api_key" "pha_api_key_instance" {
  accept_language = "en-US"
  if_none_match = "abcdef"
  pha_instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
  api_key = "xxxx_asnj_cdcdw_csdcdcwcwewdwe_cwsw"
}
```
### Path Parameters

* `pha_instance_id` - (Required, Forces new resource, String) The unique identifier of the powerha service instance.
  * Constraints: The maximum length is `1048` characters. The minimum length is `36` characters. The value must match regular expression `/^.*$/`.

### Query Parameters

* `accept_language` - (Optional, Forces new resource, String) The language in which the response should be returned.
* `accepts_incomplete` - (Optional, Forces new resource, Boolean) Indicates whether the request can be accepted before the operation is complete.
* `if_none_match` - (Optional, Forces new resource, String) ETag for conditional requests (optional).
  * Constraints: The maximum length is `50` characters. The minimum length is `6` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.

## Argument Reference

You can specify the following arguments for this resource.

* `api_key` - (Required, Forces new resource, String) The API key to be stored or registered.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `href` - (String) Resource reference.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `id` - (String) Unique identifier for the API key record.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:\/ -]+$/`.
* `status` - (String) Status of the API key retrieval request.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._: -]+$/`.

* `etag` - ETag identifier for pha_api_key.

## Import

You can import the `ibm_pha_api_key` resource by using `id`.
The `id` property can be formed from `pha_instance_id`, and `pha_instance_id` in the following format:

<pre>
&lt;pha_instance_id&gt;/&lt;pha_instance_id&gt;
</pre>
* `pha_instance_id`: A string in the format `8eefautr-4c02-0009-0086-8bd4d8cf61b6`. Unique identifier of the provisioned instance.
* `pha_instance_id`: A string in the format `9676_fwdfwfc_cdscvsvc_csd_7890`. Unique identifier for the API key record.

# Syntax
<pre>
$ terraform import ibm_pha_api_key.pha_api_key &lt;pha_instance_id&gt;/&lt;pha_instance_id&gt;
</pre>

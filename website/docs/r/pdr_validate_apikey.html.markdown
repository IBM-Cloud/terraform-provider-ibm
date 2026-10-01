---
layout: "ibm"
page_title: "IBM : ibm_pdr_validate_apikey"
description: |-
  Manages pdr_validate_apikey.
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
---

# ibm_pdr_validate_apikey

Add and update the API key for DR deployments.

## Example Usage

```hcl
resource "ibm_pdr_validate_apikey" "pdr_validate_apikey_instance" {
  accept_language = "en-US"
  instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
  api_key = "xxxx_asnj_cdcdw_csdcdcwcwewdwe_cwsw"
}
```
### Path Parameters

* `instance_id` - (Required, Forces new resource, String) The unique identifier of the DR service instance.
  * Constraints: The maximum length is `1048` characters. The minimum length is `36` characters. The value must match regular expression `/^.*$/`.

### Query Parameters

* `accept_language` - (Optional, Forces new resource, String) The language in which the response should be returned.
* `accepts_incomplete` - (Optional, Forces new resource, Boolean) Indicates whether the request can be accepted before the operation is complete.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the pdr_validate_apikey.
* `description` - (String) Validation result message.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[\\x20-\\x7E]*$/`.
* `href` - (String) Resource reference.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `api_key` - (Required, Forces new resource, String) The API key to be stored or registered.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.  
* `status` - (String) Status of the API key.
  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.


## Import

You can import the `ibm_pdr_validate_apikey` resource by using `id`.
The `id` property can be formed from `instance_id`, and `instance_id` in the following format:

<pre>
&lt;instance_id&gt;/&lt;instance_id&gt;
</pre>
* `instance_id`: A string in the format `123456d3-1122-3344-b67d-4389b44b7bf9`. Service Instance ID.
* `instance_id`: A string in the format `crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::`. Unique identifier of the API key.

# Syntax
<pre>
$ terraform import ibm_pdr_validate_apikey.pdr_validate_apikey &lt;instance_id&gt;/&lt;instance_id&gt;
</pre>

---
layout: "ibm"
page_title: "IBM : ibm_pdr_validate_apikey"
description: |-
  Manages pdr_validate_apikey.
subcategory: "DrAutomation Service"
---

# ibm_pdr_validate_apikey

Create, update, and delete pdr_validate_apikeys with this resource.

## Example Usage

```hcl
resource "ibm_pdr_validate_apikey" "pdr_validate_apikey_instance" {
  instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `accept_language` - (Optional, String) The language requested for the return document.
* `if_none_match` - (Optional, String) ETag for conditional requests (optional).
* `instance_id` - (Required, Forces new resource, String) instance id of instance to provision.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the pdr_validate_apikey.
* `description` - (String) Validation result message.
* `instance_id` - (String) 
* `status` - (String) Status of the API key.

* `etag` - ETag identifier for pdr_validate_apikey.

## Import

You can import the `ibm_pdr_validate_apikey` resource by using `id`.
The `id` property can be formed from `instance_id`, and `instance_id` in the following format:

<pre>
&lt;instance_id&gt;/&lt;instance_id&gt;
</pre>
* `instance_id`: A string in the format `crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::`. instance id of instance to provision.
* `instance_id`: A string in the format `crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::`.

# Syntax
<pre>
$ terraform import ibm_pdr_validate_apikey.pdr_validate_apikey &lt;instance_id&gt;/&lt;instance_id&gt;
</pre>

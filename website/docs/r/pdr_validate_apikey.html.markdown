---
layout: "ibm"
page_title: "IBM : ibm_pdr_validate_apikey"
description: |-
  Manages pdr_validate_apikey.
subcategory: "DrAutomation Service"
---

# ibm_pdr_validate_apikey

 Create, Retrive and Updating the current API key details for the specified service instance.

## Example Usage

```hcl
resource "ibm_pdr_validate_apikey" "pdr_validate_apikey_instance" {
  instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
  api_key = "adsfsdfdfsadfda"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `accept_language` - (Optional, String) The language requested for the return document.
* `instance_id` - (Required, Forces new resource, String) The ID of the Power DR Automation service instance.
* `api_key` - (Required, String) The new API key value that will replace the existing one.
 
## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the pdr_validate_apikey.
* `description` - (String) Validation result message.
* `instance_id` - (String) Unique identifier of the API key.
* `status` - (String) Status of the API key.


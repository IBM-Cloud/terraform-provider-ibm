---
layout: "ibm"
page_title: "IBM : ibm_pdr_validate_proxyip"
description: |-
  Get information about pdr_validate_proxyip
subcategory: "DrAutomation Service"
---

# ibm_pdr_validate_proxyip

Provides a read-only data source to retrieve information about a pdr_validate_proxyip. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pdr_validate_proxyip" "pdr_validate_proxyip" {
	instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
	proxyip = "10.30.40.5:3128"
	vpc_id = "r006-2f3b3ab9-2149-49cc-83a1-30a5d93d59b2"
	vpc_location = "us-south"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `if_none_match` - (Optional, String) ETag for conditional requests (optional).
* `instance_id` - (Required, Forces new resource, String) instance id of instance to provision.
* `proxyip` - (Required, String) proxyip value.
* `vpc_id` - (Required, String) vpc id value.
* `vpc_location` - (Required, String) vpc location value.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_validate_proxyip.
* `description` - (String) 
* `status` - (String) 
* `warning` - (Boolean) Indicates whether the proxy IP is valid but has an advisory (e.g., not in reserved IPs).
  * Constraints: The default value is `false`.


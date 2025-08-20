---
layout: "ibm"
page_title: "IBM : ibm_pdr_managedr"
description: |-
  Manages pdr_managedr.
subcategory: "DrAutomation Service"
---

# ibm_pdr_managedr

Create, update, and delete pdr_managedrs with this resource.

## Example Usage

```hcl
resource "ibm_pdr_managedr" "pdr_managedr_instance" {
  instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
  stand_by_redeploy = "true"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `accept_language` - (Optional, Forces new resource, String) The language requested for the return document.
* `accepts_incomplete` - (Optional, Forces new resource, Boolean) A value of true indicates that both the IBM Cloud platform and the requesting client support asynchronous deprovisioning.
  * Constraints: The default value is `true`.
* `if_none_match` - (Optional, Forces new resource, String) ETag for conditional requests (optional).
* `instance_id` - (Required, Forces new resource, String) instance id of instance to provision.
* `stand_by_redeploy` - (Required, Forces new resource, String) Flag to indicate if standby should be redeployed (must be "true" or "false").
  * Constraints: The default value is `false`. Allowable values are: `true`, `false`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the pdr_managedr.
* `dashboard_url` - (String) 
* `instance_id` - (String) 

* `etag` - ETag identifier for pdr_managedr.

## Import

You can import the `ibm_pdr_managedr` resource by using `id`.
The `id` property can be formed from `instance_id`, and `instance_id` in the following format:

<pre>
&lt;instance_id&gt;/&lt;instance_id&gt;
</pre>
* `instance_id`: A string in the format `crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::`. instance id of instance to provision.
* `instance_id`: A string in the format `crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::`.

# Syntax
<pre>
$ terraform import ibm_pdr_managedr.pdr_managedr &lt;instance_id&gt;/&lt;instance_id&gt;
</pre>

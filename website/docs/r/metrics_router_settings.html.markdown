---
layout: "ibm"
page_title: "IBM : ibm_metrics_router_settings"
description: |-
  Manages metrics_router_settings.
subcategory: "Metrics Routing API Version 3"
---

# ibm_metrics_router_settings

Create, update, and delete metrics_router_settingss with this resource.

## Example Usage

```hcl
resource "ibm_metrics_router_settings" "metrics_router_settings_instance" {
  backup_metadata_region = "us-east"
  default_targets {
		id = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
		crn = "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
		name = "a-mr-target-us-south"
		target_type = "sysdig_monitor"
  }
  permitted_target_regions = us-south
  primary_metadata_region = "us-south"
  private_api_endpoint_only = false
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `backup_metadata_region` - (Optional, String) To backup all your meta data in a different region.
  * Constraints: The maximum length is `256` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 \\-_]+$/`.
* `default_targets` - (Optional, List) A list of default target references. Enterprise-managed targets are not supported.
  * Constraints: The maximum length is `2` items. The minimum length is `0` items.
Nested schema for **default_targets**:
	* `crn` - (Required, String) The CRN of a pre-defined metrics-router target.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 \\-._:\/]+$/`.
	* `id` - (Required, String) The target uuid for a pre-defined metrics router target.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 \\-._:]+$/`.
	* `name` - (Required, String) The name of a pre-defined metrics-router target.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character.
	* `target_type` - (Required, String) The type of the target.
	  * Constraints: Allowable values are: `sysdig_monitor`.
* `permitted_target_regions` - (Optional, List) If present then only these regions may be used to define a target.
  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9 \\-_]+$/`. The maximum length is `16` items. The minimum length is `0` items.
* `primary_metadata_region` - (Optional, String) To store all your meta data in a single region.
  * Constraints: The maximum length is `256` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 \\-_]+$/`.
* `private_api_endpoint_only` - (Optional, Boolean) If you set this true then you cannot access api through public network.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the metrics_router_settings.


## Import

You can import the `ibm_metrics_router_settings` resource by using `primary_metadata_region`. To store all your meta data in a single region.

# Syntax
<pre>
$ terraform import ibm_metrics_router_settings.metrics_router_settings &lt;primary_metadata_region&gt;
</pre>

# Example
```
$ terraform import ibm_metrics_router_settings.metrics_router_settings us-south
```

---
layout: "ibm"
page_title: "IBM : ibm_logs_enrichment"
description: |-
  Manages logs_enrichment.
subcategory: "Cloud Logs"
---

# ibm_logs_enrichment

Create, update, and delete logs_enrichments with this resource.

## Example Usage

```hcl
resource "ibm_logs_enrichment" "logs_enrichment_instance" {
	instance_id = ibm_resource_instance.logs_instance.guid
  	region      = ibm_resource_instance.logs_instance.location
	field_name  = "ip"
	enrichment_type {
		geo_ip  {  }
	}
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String)  Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.
* `enrichment_type` - (Required, Forces new resource, List) The enrichment type.
Nested schema for **enrichment_type**:
	* `custom_enrichment` - (Optional, List) The custom enrichment.
	Nested schema for **custom_enrichment**:
		* `id` - (Optional, Integer) The ID of the custom enrichment.
		  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
	* `geo_ip` - (Optional, List) The geo ip enrichment.
	Nested schema for **geo_ip**:
	* `suspicious_ip` - (Optional, List) The suspicious ip enrichment.
	Nested schema for **suspicious_ip**:
* `field_name` - (Required, Forces new resource, String) The enrichment field name.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_enrichment resource.
* `enrichment_id` - The unique identifier of the logs_enrichment.


## Import

You can import the `ibm_logs_enrichment` resource by using `id`. You can import the `ibm_logs_e2m` resource by using `id`. `id` combination of `region`, `instance_id` and `enrichment_id`.

# Syntax
<pre>
$ terraform import ibm_logs_enrichment.logs_enrichment < region >/< instance_id >/< enrichment_id >
</pre>

# Example
```
$ terraform import ibm_logs_enrichment.logs_enrichment eu-gb/3dc02998-0b50-4ea8-b68a-4779d716fa1f/1
```

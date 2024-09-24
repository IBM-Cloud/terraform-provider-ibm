---
layout: "ibm"
page_title: "IBM : ibm_logs_enrichments"
description: |-
  Get information about logs_enrichments
subcategory: "Cloud Logs"
---

# ibm_logs_enrichments

Provides a read-only data source to retrieve information about logs_enrichments. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_enrichments" "logs_enrichments" {
	instance_id = ibm_resource_instance.logs_instance.guid
  	region      = ibm_resource_instance.logs_instance.location
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String)  Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.


## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_enrichments.
* `enrichments` - (List) The enrichments.
  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
Nested schema for **enrichments**:
	* `enrichment_type` - (List) The enrichment type.
	Nested schema for **enrichment_type**:
		* `custom_enrichment` - (List) The custom enrichment.
		Nested schema for **custom_enrichment**:
			* `id` - (Integer) The ID of the custom enrichment.
			  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
		* `geo_ip` - (List) The geo ip enrichment.
		Nested schema for **geo_ip**:
		* `suspicious_ip` - (List) The suspicious ip enrichment.
		Nested schema for **suspicious_ip**:
	* `field_name` - (String) The enrichment field name.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `id` - (Integer) The enrichment ID.
	  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.


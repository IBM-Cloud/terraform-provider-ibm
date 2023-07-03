---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : Secondary Zones"
description: |-
  Manages IBM Cloud Infrastructure private domain name service secondary zones.
---

# ibm_dns_custom_resolver_secondary_zones

Provides a read-only data source for secondary zones. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information about secondary zones, refer to [list-secondary-zones](https://cloud.ibm.com/apidocs/dns-svcs#list-secondary-zones).

## Example usage

```terraform
data "ibm_dns_custom_resolver_secondary_zones" "test-sz" {
	instance_id	= ibm_dns_custom_resolver.test.instance_id
	resolver_id	= ibm_dns_custom_resolver.test.custom_resolver_id
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_id` - (Required, String) The GUID of the DNS Services instance.
- `resolver_id` - (Required, String) The unique identifier of a custom resolver.

## Attribute reference

In addition to the argument references list, you can access the following attribute references after your data sources are created.

- `secondary_zones` (List) List of secondary zones.

	Nested scheme for `secondary_zones`:
	- `description` - (String) Descriptive text of the secondary zone.
	- `zone` - (String) The name of the zone.
	- `enabled` - (String) Enable/Disable the secondary zone.
	- `transfer_from` - (List) The addresses of DNS servers where the secondary zone data is transferred from.

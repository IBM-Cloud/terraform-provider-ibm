---
subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: ibm_dns_secondary"
description: |-
  Get information about an IBM DNS secondary resource.
---

# ibm_dns_secondary
Retrieve information of an existing DNS secondary zone as a read-only data source. For more information, about DNS secondary resource, see [managing secondary DNS zones](https://cloud.ibm.com/docs/dns?topic=dns-manage-secondary-dns-zones).

## Example usage

```terraform
data "ibm_dns_secondary" "secondary_id" {
    name = "test-secondary.com"
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `zone_name` - (Required, String) The name of the secondary zone.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier of the secondary.
- `master_ip_address` - (String) The IP address of the master name server where a secondary DNS zone is transferred from.
- `status_id` - (String) The status of a secondary DNS record.
- `status_text` - (String) The textual representation of a secondary DNS zone's status.
- `transfer_frequency`- (Integer) Signifies how often a secondary DNS zone transferred in minutes.

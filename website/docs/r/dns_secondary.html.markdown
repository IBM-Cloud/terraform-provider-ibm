---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: dns_secondary"
description: |-
  Manages IBM DNS secondary zone.
---

# ibm_dns_secondary
The `ibm_dns_secondary` resource represents a single secondary DNS zone managed on SoftLayer. Each record created within the secondary DNS service defines which zone is transferred, what server it is transferred from, and the frequency that zone transfers occur at. Zone transfers are performed automatically based on the transfer frequency set on the secondary DNS record. For more information, about DNS secondary zone, see [managing secondary DNS zones](https://cloud.ibm.com/docs/dns?topic=dns-manage-secondary-dns-zones).

## Example usage

```terraform
resource "ibm_dns_secondary" "dns-secondary-test" {
    zone_name = "dns-secondary-test.com"
    master_ip_address = "127.0.0.10"
    transfer_frequency = 10
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `master_ip_address` - (Required, String) The IP address of the master name server where a secondary DNS zone is transferred from.
- `tags`- (Optional, Array of Strings) Tags associated with the DNS secondary instance.  **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.
- `transfer_frequency` - (Required, Integer) Signifies how often a secondary DNS zone should be transferred in minutes.
- `zone_name` - (Required, Forces new resource, String) The name of the zone that is transferred.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`- (String) A secondary zone's internal identifier.
- `status_id`- (String) The status of a secondary DNS record.
- `status_text`- (String) The textual representation of a secondary DNS zone's status.

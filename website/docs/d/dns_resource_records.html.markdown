---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : Private DNS Resource Records"
description: |-
  Manages IBM Cloud infrastructure private domain name service resource records.
---

# ibm_dns_resource_records

Retrieve details about existing IBM Cloud private domain name service records. For more information, about DNS records, see [managing DNS record](https://cloud.ibm.com/docs/dns-svcs?topic=dns-svcs-managing-dns-records).


## Example usage

```terraform
data "ibm_dns_resource_records" "ds_pdns_resource_records" {
  instance_id = "resource_instance_guid"
  zone_id = "resource_dns_resource_records_zone_id"
}
```

## Argument reference
Review the argument reference that you can specify for your data source. 

- `instance_id` - (Required, String) The GUID of the private DNS service instance.
- `zone_id` - (Required, String) The ID of the zone that you added to the private DNS service instance.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `dns_resource_records`- (List) A list of all private domain name service resource records.

  Nested scheme for `dns_resource_records`:
  - `id` - (String) The unique identifier of the private DNS resource record.
  - `name` - (String) The name of a private DNS resource record.
  - `type` - (String) The type of the private DNS resource record. Supported values are `A`, `AAAA`, `CNAME`, `PTR`, `TXT`, `MX`, and `SRV`.
  - `rdata` - (String) The resource data of a private DNS resource record.
  - `ttl`- (Integer) The time-to-live value of the DNS resource record.

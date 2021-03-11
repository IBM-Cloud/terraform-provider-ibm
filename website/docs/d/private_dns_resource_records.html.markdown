---

subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : "
description: |-
  Manages IBM Cloud Infrastructure Private Domain Name Service Resource Records.
---

# ibm\_dns_resource_records

Import the details of an existing IBM Cloud Infrastructure private domain name service resource records as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_dns_resource_records" "ds_pdns_resource_records" {
  instance_id = "resource_instance_guid"
  zone_id = "resource_dns_resource_records_zone_id"
}

```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, string) The resource instance id of the private DNS on which zones were created.
* `zone_id` - (Required, string) The unique identifier of the private DNS zone.



## Attribute Reference

The following attributes are exported:

* `dns_resource_records` - List of all private domain name service resource records in the IBM Cloud Infrastructure.
  * `id` - The unique identifier of the private DNS resource record.
  * `name` - The name of a DNS resource record.
  * `type` - The type of the DNS resource record. Supported Resource Record types are: A, AAAA, CNAME, PTR, TXT, MX, SRV.
  * `rdata` - The resource data of a DNS resource record.
  * `ttl` - The time-to-live value of the DNS resource record.

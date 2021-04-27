---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: dns_reverse_record"
description: |-
  Manages IBM DNS reverse records.
---

# ibm\_dns_reverse_record

Provides a single DNS reverse record managed on IBM Cloud Classic Infrastructure (SoftLayer). Record contain general information about the reverse record, such as the hostname, ip address and time to leave(ttl).

The IBM Cloud Classic Infrastructure (SoftLayer) object  [SoftLayer_Dns_Domain_ResourceRecord](https://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord) is used for most CRUD operations.

## Example Usage
```hcl
resource "ibm_dns_reverse_record" "testreverserecord" {
    ipaddress="123.123.123.123"
    hostname="www.example.com"
    ttl=900
}
```

## Argument Reference

The following arguments are supported:

* `ipaddress` - (Required, Forces new resource, string) The IP address or a hostname of a domain's resource record.

* `hostname` - (Required, string) The host defined by a reverse record.

* `ttl` - (Optional, integer) The time to live (TTL) duration, expressed in seconds, of a resource record. Default value is 604800 seconds.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of reverse dns record.
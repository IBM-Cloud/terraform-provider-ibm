---
layout: "ibm"
page_title: "IBM: dns_reverse_record"
sidebar_current: "docs-ibm-resource-dns-reverse-record"
description: |-
  Manages IBM DNS reverse records.
---

# ibm\_dns_reverse_record

Provides a single DNS reverse record managed on IBM Cloud Infrastructure (SoftLayer). Record contain general information about the reverse record, such as the hostname, ip address and time to leave(ttl).

The IBM Cloud Infrastructure (SoftLayer) object  [SoftLayer_Dns_Domain_ResourceRecord](https://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord) is used for most CRUD operations.

## Example Usage
```hcl
resource "ibm_dns_reverse_record" "testreverserecord" {
    dns_ipaddress="123.123.123.123"
    dns_hostname="www.example.com"
    dns_ttl=900
}
```

## Argument Reference

The following arguments are supported:

* `dns_ipaddress` - (Required, string) The IP address or a hostname of a domain's resource record.

* `dns_hostname` - (Required, string) The host defined by a reverse record.

* `dns_ttl` - (Optional, integer) The time to live (TTL) duration, expressed in seconds, of a resource record. Default value is 604800 seconds.


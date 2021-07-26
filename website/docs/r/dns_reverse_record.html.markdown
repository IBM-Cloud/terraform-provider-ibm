---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: dns_reverse_record"
description: |-
  Manages IBM DNS reverse record.
---

# ibm_dns_reverse_record
Provides a single DNS reverse record managed on IBM Cloud Classic Infrastructure (SoftLayer). Record contains general information about the reverse record, such as the hostname, IP address and time to leave.

The IBM Cloud Classic Infrastructure (SoftLayer) object  [SoftLayer_Dns_Domain_ResourceRecord](https://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord) is used for most create-retrieve-update-delete (`CRUD`) operations.

## Example usage
```terraform
resource "ibm_dns_reverse_record" "testreverserecord" {
    ipaddress="123.123.123.123"
    hostname="www.example.com"
    ttl=900
}
```


## Argument reference
Review the argument references that you can specify for your resource. 

- `ipaddress` - (Required, Forces new resource, String)The IP address or a hostname of a domain's resource record.
- `hostname` - (Required, String)The host defined by a reverse record.
- `ttl`- (Optional, Integer) The time to live (TTL) duration, expressed in seconds, of a resource record. Default value is 604800 seconds.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`- (String) The unique identifier of reverse dns record.

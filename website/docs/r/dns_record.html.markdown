---
layout: "ibm"
page_title: "IBM: dns_record"
sidebar_current: "docs-ibm-resource-dns-record"
description: |-
  Manages IBM DNS resource records.
---

# ibm\_dns_record

Represents a single-resource record entry in `ibm_dns_domain`. Each resource record contains a `host` and `data` property to define the name and target data of a resource.

The Bluemix Infrastructure (SoftLayer) object  [SoftLayer_Dns_Domain_ResourceRecord](https://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord) is used for most CRUD operations. For SRV record types, the Bluemix Infrastructure (SoftLayer) object [SoftLayer_Dns_Domain_ResourceRecord_SrvType](https://sldn.softlayer.com/reference/services/SoftLayer_Dns_Domain_ResourceRecord_SrvType) is used.

You cannot create SOA nor NS record types, as these are automatically created by Bluemix Infrastructure (SoftLayer) when the domain is created.

## Example Usage

### `A` Record | [SLDN](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_AType)

```hcl
resource "ibm_dns_domain" "main" {
    name = "main.example.com"
}

resource "ibm_dns_record" "www" {
    data = "123.123.123.123"
    domain_id = "${ibm_dns_domain.main.id}"
    host = "www.example.com"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "a"
}
```

### `AAAA` Record | [SLDN](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_AaaaType)

```hcl
resource "ibm_dns_record" "aaaa" {
    data = "fe80:0000:0000:0000:0202:b3ff:fe1e:8329"
    domain_id = "${ibm_dns_domain.main.id}"
    host = "www.example.com"
    responsible_person = "user@softlayer.com"
    ttl = 1000
    type = "aaaa"
}
```

### `CNAME` Record | [SLDN](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_CnameType)

```hcl
resource "ibm_dns_record" "cname" {
    data = "real-host.example.com."
    domain_id = "${ibm_dns_domain.main.id}"
    host = "alias.example.com"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "cname"
}
```

### `MX` Record | [SLDN](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_MxType)

```hcl
resource "sibm_dns_record" "recordMX-1" {
    data = "mail-1"
    domain_id = "${ibm_dns_domain.main.id}"
    host = "@"
    mx_priority = "10"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "mx"
}
```

### `SPF` Record | [SLDN](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_SpfType)

```hcl
resource "ibm_dns_record" "recordSPF" {
    data = "v=spf1 mx:mail.example.org ~all"
    domain_id = "${ibm_dns_domain.main.id}"
    host = "mail-1"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "spf"
}
```

### `TXT` Record | [SLDN](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_TxtType/)

```hcl
resource "ibm_dns_record" "recordTXT" {
    data = "host"
    domain_id = "${ibm_dns_domain.main.id}"
    host = "A SPF test host"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "txt"
}
```

### `SRV` Record | [SLDN](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_SrvType)

```hcl
resource "ibm_dns_record" "recordSRV" {
    data = "ns1.example.org"
    domain_id = "${ibm_dns_domain.main.id}"
    host = "hosta-srv.com"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "srv"
    port = 8080
    priority = 3
    protocol = "_tcp"
    weight = 3
    service = "_mail"
}
```

### `PTR` Record

Review the [Bluemix Infrastructure (SoftLayer) docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_PtrType/) to properly implement the `PTR` record. 

```hcl
resource "ibm_dns_record" "recordPTR" {
    data = "ptr.example.com"
    domain_id = "${ibm_dns_domain.main.id}"
# The host is the last octet of IP address in the range of the subnet
    host = "45"  
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "ptr"
}
```

## Argument Reference

The following arguments are supported:

* `data` - (Required, string) The value of a domain's resource record. This can be an IP address or a hostname. Fully qualified host and domain name data must end with the `.` character.
* `domain_id` - (Required, integer) The identifier belonging to the domain that a resource record is associated with.
* `expire` - (Integer) The duration, expressed in seconds, that a secondary name server (or servers) holds a zone before it is no longer considered authoritative.
* `host` - (Required, string) The host defined by a resource record. The `@` symbol denotes a wildcard.
* `minimum_ttl` - (Integer) The duration, expressed in seconds, that a domain's resource records are valid. This is also known as a minimum time to live (TTL), and can be overridden by an individual resource record's TTL.
* `mx_priority` - (Integer) Useful in cases where a domain has more than one mail exchanger, the priority property is the priority of the MTA that delivers mail for a domain. A lower number denotes a higher priority, and mail will attempt to deliver through that MTA before moving to lower priority mail servers. Priority is defaulted to 10 upon resource record creation.
* `refresh` - (Integer) The duration, expressed in seconds, that a secondary name server waits to check for a new copy of a DNS zone from the domain's primary name server. If a zone file has changed, the secondary DNS server updates its copy of the zone to match the primary DNS server's zone.
* `responsible_person` - (Required, string) The email address of the person responsible for a domain, with the `@` symbol replaced with a `.`. Example usage: root@example.org SOA responsibility would be expressed as `root.example.org.`.
* `retry` - (Integer) The duration, expressed in seconds, that a domain's primary name server (or servers) wait before attempting to refresh a domain's zone with the secondary name server. The retry action is triggered after a previous failed attempt to refresh by a secondary name server. 
* `ttl` - (Required, integer) The time to live (TTL) duration, expressed in seconds, of a resource record. TTL is used by a name server to determine how long to cache a resource record. An SOA record's TTL value defines the domain's overall TTL.
* `type` - (Required, string) The type of domain resource record. Accepted values are as follows:
    * `a` for address records
    * `aaaa` for address records
    * `cname` for canonical name records
    * `mx` for mail exchanger records
    * `ptr` for pointer records in reverse domains
    * `spf` for sender policy framework records
    * `srv` for service records
* `txt` - (String) Used for text records.
* `service` - (`SRV` records only, string) The symbolic name of the desired service. 
* `protocol` - (`SRV` records only, string) The protocol of the desired service; this is usually TCP or UDP.
* `port` - (`SVR` records only, integer) The TCP or UDP port on which the service is to be found.
* `priority` - (`SVR` records only, integer) The priority of the target host. The lowest numerical value is given the highest priority.
* `weight` - (`SVR` records only, integer) A relative weight for records that have the same priority.

## Attributes Reference

The following attributes are exported:

* `id` - A domain resource record's internal identifier.

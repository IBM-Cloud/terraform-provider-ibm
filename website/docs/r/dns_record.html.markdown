---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: dns_record"
description: |-
  Manages IBM DNS resource record.
---

# ibm_dns_record
Provides a single-resource record entry in `ibm_dns_domain`. Each resource record contains a `host` and a `data` property to define the name and target data of a resource.

The IBM Cloud Classic Infrastructure (SoftLayer) object  [SoftLayer_Dns_Domain_ResourceRecord](https://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord) is used for most create-retrieve-update-delete (`CRUD`) operations. The IBM Cloud Classic Infrastructure (SoftLayer) object [SoftLayer_Dns_Domain_ResourceRecord_SrvType](https://sldn.softlayer.com/reference/services/SoftLayer_Dns_Domain_ResourceRecord_SrvType) is used for service record types (`SRV`).

The SOA and NS records are automatically created by IBM Cloud Classic Infrastructure (SoftLayer) when the domain is created, you don't need to create those manually.

## Example usage

### `A` Record
Review the [IBM Cloud Classic Infrastructure Docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_AType) to implement the `A` record.

```terraform
resource "ibm_dns_domain" "main" {
    name = "main.example.com"
}

resource "ibm_dns_record" "www" {
    data = "123.123.123.123"
    domain_id = ibm_dns_domain.main.id
    host = "www.example.com"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "a"
}
```

### `AAAA` Record
Review the [IBM Cloud Classic Infrastructure documentation](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_AaaaType) to implement the `AAAA` record.

```terraform
resource "ibm_dns_record" "aaaa" {
    data = "fe80:0000:0000:0000:0202:b3ff:fe1e:8329"
    domain_id = ibm_dns_domain.main.id
    host = "www.example.com"
    responsible_person = "user@softlayer.com"
    ttl = 1000
    type = "aaaa"
}
```

### `CNAME` Record
Review the [IBM Cloud Classic Infrastructure documentation](https://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_CnameType/) to implement the `CNAME` record.

```terraform
resource "ibm_dns_record" "cname" {
    data = "real-host.example.com."
    domain_id = ibm_dns_domain.main.id
    host = "alias.example.com"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "cname"
}
```

### `NS` Record
Review the [IBM Cloud Classic Infrastructure Docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_NsType) to implement the `NS` record.

```terraform
resource "ibm_dns_record" "recordNS" {
    data = "ns.example.com."
    domain_id = ibm_dns_domain.main.id
    host = "example.com"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "ns"
}
```

### `MX` Record
Review the [IBM Cloud Classic Infrastructure Docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_MxType) to implement the `MX` record.

```terraform
resource "sibm_dns_record" "recordMX-1" {
    data = "mail-1"
    domain_id = ibm_dns_domain.main.id
    host = "@"
    mx_priority = "10"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "mx"
}
```

### `SOA` Record
Review the [IBM Cloud Classic Infrastructure Docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_SoaType) to implement the `SOA` record.

```terraform
resource "ibm_dns_record" "recordSOA" {
    data = "ns1.example.com. abuse.example.com. 2018101002 7200 600 1728000 43200"
    domain_id = ibm_dns_domain.main.id
    host = "example.com"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "soa"
}
```

### `SPF` Record
Review the [IBM Cloud Classic Infrastructure Docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_SpfType) to implement the `SPF` record.

```terraform
resource "ibm_dns_record" "recordSPF" {
    data = "v=spf1 mx:mail.example.org ~all"
    domain_id = ibm_dns_domain.main.id
    host = "mail-1"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "spf"
}
```

### `TXT` Record
Review the [IBM Cloud Classic Infrastructure Docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_TxtType/) to implement the `TXT` record.

```terraform
resource "ibm_dns_record" "recordTXT" {
    data = "host"
    domain_id = ibm_dns_domain.main.id
    host = "A SPF test host"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "txt"
}
```

### `SRV` Record
Review the [IBM Cloud Classic Infrastructure  Docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_SrvType) to implement the `SRV` record.

```terraform
resource "ibm_dns_record" "recordSRV" {
    data = "ns1.example.org"
    domain_id = ibm_dns_domain.main.id
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
Review the [IBM Cloud Classic Infrastructure  Docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_PtrType/) to implement the `PTR` record.

```terraform
resource "ibm_dns_record" "recordPTR" {
    data = "ptr.example.com"
    domain_id = ibm_dns_domain.main.id
# The host is the last octet of IP address in the range of the subnet
    host = "45"  
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "ptr"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `data` - (Required, String) The IP address or a hostname of a domain's resource record. Fully qualified host and domain name data must end with the `.` character.
- `domain_id` - (Required, Forces new resource, Integer) The ID for the domain associated with the resource record.
- `expire` - (Optional, Integer)T he duration, expressed in seconds, that a secondary name server (or servers) holds a zone before it is no longer considered authoritative.
- `host`- (Required, String) The host defined by a resource record. The `@` symbol denotes a wildcard.
- `minimum_ttl` - (Optional, Integer) The duration, expressed in seconds, that a domain's resource records are valid. This is also known as a minimum time to live (TTL), and can be overridden by an individual resource record's TTL.
- `mx_priority` - (Optional, Integer) The priority of the mail exchanger that delivers mail for a domain. This is useful in cases where a domain has more than one mail exchanger. A lower number denotes a higher priority, and mail attempt to deliver through the highest priority exchanger before moving to lower priority exchanger. The default value is `0`.
- `protocol` - (Required, String)`SRV` records only. The protocol of the service that you want to use, such as `TCP` or `UDP`.
- `port` - (Required, Integer) `SRV` records only. The TCP or UDP port on which the service be found.
- `priority` - (Required, Integer) `SRV` records only. The priority of the target host. The lowest numerical value is given the highest priority. The default value is `0`.
- `refresh` - (Optional, Integer) The duration expressed in seconds, that a secondary name server waits to check the domain's primary name server for a new copy of a DNS zone. If a zone file has changed, the secondary DNS server updates its copy of the zone to match the primary DNS server's zone.
- `responsible_person` - (Required, String)The email address of the person responsible for a domain. Replace the `@` symbol in the address with a `.`. For example, root@example.org would be expressed as `root.example.org.`.
- `retry` - (Optional, Integer)The duration, expressed in seconds, that the domain's primary name server (or servers) waits before attempting to refresh the domain's zone with the secondary name server. A failed attempt to refresh by a secondary name server triggers the retry action.
- `service` - (Required, String)`SRV` records only. The symbolic name of the service.
- `tags`- (Optional, Array of Strings) Tags associated with the DNS domain record instance.     
   **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.
- `txt`- (Optional, String) Used for text records.
- `ttl` - (Required, Integer)  -The time to live (TTL) duration, expressed in seconds, of a resource record. A name server uses TTL to determine how long to cache a resource record. An SOA record's TTL value defines the domain's overall TTL.
- `type`- (Required, Forces new resource, String) The type of domain resource record. Accepted values are as follows: 
  - `a` for address records.
  - `aaaa` for address records.
  - `cname` for canonical name records.
  - `mx` for mail exchanger records.
  - `ptr` for pointer records in reverse domains.
  - `spf` for sender policy framework records.
  - `srv` for service records.
- `weight` - (Required, Integer) `SRV` records only. A relative weight for records that have the same priority. The default value is `0`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`- (String) The internal identifier of the domain resource record.

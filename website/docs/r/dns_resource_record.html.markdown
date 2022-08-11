---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : dns_resource_record"
description: |-
  Manages IBM Private DNS Resource records.
---

# ibm_dns_resource_record

Create, update, or delete a DNS record. For more information, see [managing DNS records](https://cloud.ibm.com/docs/dns-svcs?topic=dns-svcs-managing-dns-records). 


## Example usage

```terraform
resource "ibm_dns_resource_record" "test-pdns-resource-record-a" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
  type        = "A"
  name        = "testA"
  rdata       = "1.2.3.4"
  ttl         = 3600
}

resource "ibm_dns_resource_record" "test-pdns-resource-record-aaaa" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
  type        = "AAAA"
  name        = "testAAAA"
  rdata       = "2001:0db8:0012:0001:3c5e:7354:0000:5db5"
}

resource "ibm_dns_resource_record" "test-pdns-resource-record-cname" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
  type        = "CNAME"
  name        = "testCNAME"
  rdata       = "test.com"
}

resource "ibm_dns_resource_record" "test-pdns-resource-record-ptr" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
  type        = "PTR"
  name        = "1.2.3.4"
  rdata       = "testA.test.com"
}

resource "ibm_dns_resource_record" "test-pdns-resource-record-mx" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
  type        = "MX"
  name        = "testMX"
  rdata       = "mailserver.test.com"
  preference  = 10
}

resource "ibm_dns_resource_record" "test-pdns-resource-record-srv" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
  type        = "SRV"
  name        = "testSRV"
  rdata       = "tester.com"
  priority    = 100
  weight      = 100
  port        = 8000
  service     = "_sip"
  protocol    = "udp"
}

resource "ibm_dns_resource_record" "test-pdns-resource-record-txt" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
  type        = "TXT"
  name        = "testTXT"
  rdata       = "textinformation"
  ttl         = 900
}
```


## Argument reference
Review the argument reference that you can specify for your resource. 

- `instance_id` - (Required, String) The GUID of the private DNS instance.
- `name` - (Required, String) The name of the DNS record. 
- `preference` - (Optional, Integer) Required for `MX` records. If you create an `MX` record, enter the preference of the record.
- `priority` - (Optional, Integer) Required for `SRV` records-If you create an `SRV` record, enter the priority of the record.
- `port` - (Optional, Integer) Required for `SRV` records. If you create an `SRV` record, enter the TCP or UDP port of the target server.
- `protocol` - (Optional, Integer) Required for `SRV` records. If you create an `SRV` record, enter the name of the protocol that you want.
- `service` - (Optional, Integer) Required for `SRV` records. If you create an `SRV` record, enter the name of the service that you want. The name must start with an underscore (`_`).
- `rdata` - (Required, String) The resource data of a DNS resource record.
- `ttl` - (Optional, Integer) The time to live (TTL) value of the DNS record to be created.
- `type` - (Required, String) The type of DNS record that you want to create. Supported values are `A`, `AAAA`, `CNAME`, `PTR`, `TXT`, `MX`, and `SRV`.
- `weight` - (Optional, Integer) Required for `SRV` records. If you create an `SRV` record, enter the weight of the record. The weight of distributing queries among multiple target servers.
- `zone_id` - (Required, String) The ID of the DNS zone where you want to create a DNS record.


## Attribute reference
In addition to all arguments listed, you can access the following attribute references after your resource is created.

- `created_on` - (Timestamp) The time when the DNS record was created. 
- `id` - (String) The unique identifier of the DNS record. The ID is composed of `<instance_id>/<zone_id>/<dns_record_id>`.
- `modified_on` - (Timestamp) The time when the DNS record was modified.
- `resource_record_id` - (String) The ID of the DNS record. 
- `zone_id` - (String) The ID of the zone where the DNS record was created. 

## Import
The `ibm_dns_resource_record` resource can be imported by using the DNS record ID, zone ID and resource record ID. 

**Syntax**

```
$ terraform import ibm_dns_resource_record.example <instance_id>/<zone_id>/<dns_record_id>
```

**Example**

```
$ terraform import ibm_dns_resource_record.example 6ffda12064634723b079acdb018ef308/5ffda12064634723b079acdb018ef308/6463472064634723b079acdb018a1206
```

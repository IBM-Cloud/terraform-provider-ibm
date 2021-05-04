---

subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : dns_resource_record"
description: |-
  Manages IBM Private DNS Resource Records.
---

# ibm\_dns_resource_record

Provides a private dns resource record resource. This allows dns resource records to be created, and updated and deleted.

## Example Usage

```hcl

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

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, string) The guid of the private DNS instance.
* `zone_id` - (Required, string)  The ID of the DNS zone.
* `type` - (Required, string) The type of the DNS resource record to be created. Supported Resource Record types are: A, AAAA, CNAME, PTR, TXT, MX, SRV. Update is not allowed for this attribute.
* `name` -  (Required, string) The name of a DNS resource record.
* `rdata` -  (Required, string) The resource data of a DNS resource record.
* `ttl` - (Optional, int) The time-to-live value of the DNS record to be created.
* `preference` - (Optional, int) The preference of the MX record. Mandatory field for MX record type.
* `priority` - (Optional, int) The priority of the record. Mandatory field for SRV record type.
* `weight` - (Optional, int) The weight of distributing queries among multiple target servers. Mandatory field for SRV record
* `port` - (Optional, int) The port number of the target server. Mandatory field for SRV record.
* `service` - (Optional, int) The symbolic name of the desired service, start with an underscore (_). Mandatory field for SRV record.
* `protocol` - (Optional, int) The symbolic name of the desired protocol. Madatory field for SRV record.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the private DNS resource record. The id is composed of <instance_id>/<zone_id>/<resource_record_id>.
* `zone_id` - The unique identifier of the private DNS zone.
* `resource_record_id` - The unique identifier of the private DNS resource record.
* `created_on` - The time (Created On) of the DNS resource record.
* `modified_on` - The time (Modified On) of the DNS rsource record.

## Import

ibm_dns_resource_record can be imported using private DNS instance ID, zone ID and resource record ID, eg

```
$ terraform import ibm_dns_resource_record.example 6ffda12064634723b079acdb018ef308/5ffda12064634723b079acdb018ef308/6463472064634723b079acdb018a1206
```

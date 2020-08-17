---
layout: "ibm"
page_title: "IBM : dns_records"
sidebar_current: "docs-ibm-resource-dns-records"
description: |-
  Manages IBM Cloud Internet Service DNS Records.
---

# ibm_network_cis_dns_records

Provides dns record resource. This allows dns records to be created, and updated and deleted.

## Example Usage

```hcl

resource "ibm_network_cis_dns_record" "test_dns_a_record" {
  crn     = var.cis_crn
  zone_id = var.zone_id
  name    = "test-exmple"
  type    = "A"
  content = "1.2.3.4"
  ttl     = 900
}

output "a_record_output" {
  value = ibm_network_cis_dns_record.test_dns_a_record
}

resource "ibm_network_cis_dns_record" "test_dns_aaaa_record" {
  crn     = var.cis_crn
  zone_id = var.zone_id
  name    = "test-exmple.aaaa"
  type    = "AAAA"
  content = "2001::4"
  ttl     = 900
}

output "aaaa_record_output" {
  value = ibm_network_cis_dns_record.test_dns_aaaa_record
}

resource "ibm_network_cis_dns_record" "test_dns_cname_record" {
  crn     = var.cis_crn
  zone_id = var.zone_id
  name    = "test-exmple.cname.com"
  type    = "CNAME"
  content = "domain.com"
  ttl     = 900
}

output "cname_record_output" {
  value = ibm_network_cis_dns_record.test_dns_cname_record
}

resource "ibm_network_cis_dns_record" "test_dns_mx_record" {
  crn      = var.cis_crn
  zone_id  = var.zone_id
  name     = "test-exmple.mx"
  type     = "MX"
  content  = "domain.com"
  ttl      = 900
  priority = 5
}

output "mx_record_output" {
  value = ibm_network_cis_dns_record.test_dns_mx_record
}

resource "ibm_network_cis_dns_record" "test_dns_loc_record" {
  crn     = var.cis_crn
  zone_id = var.zone_id
  name    = "test-exmple.loc"
  type    = "LOC"
  ttl     = 900
  data = {
    altitude       = 98
    lat_degrees    = 60
    lat_direction  = "N"
    lat_minutes    = 53
    lat_seconds    = 53
    long_degrees   = 45
    long_direction = "E"
    long_minutes   = 34
    long_seconds   = 34
    precision_horz = 56
    precision_vert = 64
    size           = 68
  }
}

output "loc_record_output" {
  value = ibm_network_cis_dns_record.test_dns_loc_record
}

resource "ibm_network_cis_dns_record" "test_dns_caa_record" {
  crn     = var.cis_crn
  zone_id = var.zone_id
  name    = "test-exmple.caa"
  type    = "CAA"
  ttl     = 900
  data = {
    tag   = "http"
    value = "domain.com"
  }
}

output "caa_record_output" {
  value = ibm_network_cis_dns_record.test_dns_caa_record
}

resource "ibm_network_cis_dns_record" "test_dns_srv_record" {
  crn     = var.cis_crn
  zone_id = var.zone_id
  type = "SRV"
  ttl  = 900
  data = {
    name     = "test-example.srv"
    port     = 1
    priority = 1
    proto    = "_udp"
    service  = "_sip"
    target   = "domain.com"
    weight   = 1
  }
}

output "srv_record_output" {
  value = ibm_network_cis_dns_record.test_dns_srv_record
}

resource "ibm_network_cis_dns_record" "test_dns_spf_record" {
  crn     = var.cis_crn
  zone_id = var.zone_id
  name    = "test-exmple.spf"
  type    = "SPF"
  content = "test"
}

output "spf_record_output" {
  value = ibm_network_cis_dns_record.test_dns_spf_record
}

resource "ibm_network_cis_dns_record" "test_dns_txt_record" {
  crn     = var.cis_crn
  zone_id = var.zone_id
  name    = "test-exmple.txt"
  type    = "TXT"
  content = "test"
}

output "txt_record_output" {
  value = ibm_network_cis_dns_record.test_dns_txt_record
}

resource "ibm_network_cis_dns_record" "test_dns_ns_record" {
  crn     = var.cis_crn
  zone_id = var.zone_id
  name    = "test-exmple.ns"
  type    = "NS"
  content = "ns1.name.ibm.com"
}

output "ns_record_output" {
  value = ibm_network_cis_dns_record.test_dns_ns_record
}

```

## Argument Reference

The following arguments are supported:

- `crn` - (Required, string) The CRN number of instance.
- `zone_id` - (Required, string) The ID of the DNS zone.
- `type` - (Required, string) The type of the DNS record to be created. Supported Record types are: A, AAAA, CNAME, LOC, TXT, MX, SRV, SPF, NS, CAA.
- `name` - (Required, string) The name of a DNS record.
- `content` - (Required, string) The content of a DNS record.
- `ttl` - (Optional, int) The time-to-live value of the DNS record to be created.
- `priority` - (Optional, int) The priority of the record. Mandatory field for SRV record type.
- `data` - (Optional, map) The data of dns record. Field for LOC and SRV record type.
  - `weight` - (Optional, int) The weight of distributing queries among multiple target servers. Mandatory field for SRV record
  - `port` - (Optional, int) The port number of the target server. Mandatory field for SRV record.
  - `service` - (Optional, int) The symbolic name of the desired service, start with an underscore (\_). Mandatory field for SRV record.
  - `protocol` - (Optional, int) The symbolic name of the desired protocol. Madatory field for SRV record.
  - `altitude` - (Optional, int) The LOC altitude. Mondatory field for LOC record.
  - `size` - (Optional, int) The LOC altitude size. Mondatory field for LOC record.
  - `lat_degrees` - (Optional, int) The LOC latitude degrees. Mondatory field for LOC record.
  - `lat_direction` - (Optional, string) The LOC latitude direction ("N", "E", "S", "W"). Mondatory field for LOC record.
  - `lat_minutes` - (Optional, int) The LOC latitude minutes. Mondatory field for LOC record.
  - `lat_seconds` - (Optional, int) The LOC latitude seconds. Mondatory field for LOC record.
  - `long_degrees` - (Optional, int) The LOC Longitude degrees. Mondatory field for LOC record.
  - `long_direction` - (Optional, string) The LOC longitude direction ("N", "E", "S", "W"). Mondatory field for LOC record.
  - `long_minutes` - (Optional, int) The LOC longitude minutes. Mondatory field for LOC record.
  - `long_seconds` - (Optional, int) The LOC longitude seconds. Mondatory field for LOC record.
  - `precision_horz` - (Optional, int) The LOC horizontal precision. Mondatory field for LOC record.
  - `precision_vert` - (Optional, int) The LOC vertical precision. Mondatory field for LOC record.

## Attribute Reference

The following attributes are exported:

- `crn` - The CRN number of instance.
- `zone_id` - The ID of the DNS zone.
- `type` - The type of the DNS record to be created. Supported Record types are: A, AAAA, CNAME, LOC, TXT, MX, SRV, SPF, NS, CAA.
- `name` - The name of a DNS record.
- `content` - The content of a DNS record.
- `ttl` - The time-to-live value of the DNS record to be created.
- `priority` - The priority of the record. Mondatory for MX record.
- `data` - The data of dns record. Field for LOC and SRV record type.
  - `weight` - The weight of distributing queries among multiple target servers.
  - `port` - The port number of the target server.
  - `service` - The symbolic name of the desired service, start with an underscore (\_).
  - `protocol` - The symbolic name of the desired protocol.
  - `altitude` - The LOC altitude.
  - `size` - The LOC altitude size.
  - `lat_degrees` - The LOC latitude degrees.
  - `lat_direction` - The LOC latitude direction ("N", "E", "S", "W").
  - `lat_minutes` - The LOC latitude minutes.
  - `lat_seconds` - The LOC latitude seconds.
  - `long_degrees` - The LOC Longitude degrees.
  - `long_direction` - The LOC longitude direction ("N", "E", "S", "W").
  - `long_minutes` - The LOC longitude minutes.
  - `long_seconds` - The LOC longitude seconds.
  - `precision_horz` - The LOC horizontal precision.
  - `precision_vert` - The LOC vertical precision.
- `created_on` - The time (Created On) of the DNS resource record.
- `modified_on` - The time (Modified On) of the DNS rsource record.

## Import

ibm_network_cis_dns_record can be imported using CRN, zone ID and DNS record ID, eg

```
$ terraform import ibm_network_cis_dns_record.example 6ffda12064634723b079acdb018ef308|5ffda12064634723b079acdb018ef308|6463472064634723b079acdb018a1206
```

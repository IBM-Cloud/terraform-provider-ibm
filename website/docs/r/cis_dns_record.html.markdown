---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_dns_record"
description: |-
  Provides a IBM CIS DNS record resource.
---

# ibm_cis_dns_record

Create, update, or delete an IBM Cloud Internet Services DNS record resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS domain resource. For more information, about CIS DNS record, see [setting up your Domain Name System for CIS](https://cloud.ibm.com/docs/cis?topic=cis-set-up-your-dns-for-cis).

## Example usage 1 : Create A Record

```terraform
resource "ibm_cis_dns_record" "test_dns_a_record" {
  cis_id  = var.cis_crn
  domain_id = var.zone_id
  name    = "test-exmple"
  type    = "A"
  content = "1.2.3.4"
  ttl     = 900
}

output "a_record_output" {
  value = ibm_cis_dns_record.test_dns_a_record
}
```

## Example usage 2 : Create AAAA record

```terraform
resource "ibm_cis_dns_record" "test_dns_aaaa_record" {
  cis_id  = var.cis_crn
  domain_id = var.zone_id
  name    = "test-exmple.aaaa"
  type    = "AAAA"
  content = "2001::4"
  ttl     = 900
}

output "aaaa_record_output" {
  value = ibm_cis_dns_record.test_dns_aaaa_record
}
```

## Example usage 3 : Create CNAME record

```terraform
resource "ibm_cis_dns_record" "test_dns_cname_record" {
  cis_id  = var.cis_crn
  domain_id = var.zone_id
  name    = "test-exmple.cname.com"
  type    = "CNAME"
  content = "domain.com"
  ttl     = 900
}

output "cname_record_output" {
  value = ibm_cis_dns_record.test_dns_cname_record
}
```

## Example usage 4 : Create MX record

```terraform
resource "ibm_cis_dns_record" "test_dns_mx_record" {
  cis_id   = var.cis_crn
  domain_id  = var.zone_id
  name     = "test-exmple.mx"
  type     = "MX"
  content  = "domain.com"
  ttl      = 900
  priority = 5
}

output "mx_record_output" {
  value = ibm_cis_dns_record.test_dns_mx_record
}
```

## Example usage 5 : Create LOC record

```terraform
resource "ibm_cis_dns_record" "test_dns_loc_record" {
  cis_id  = var.cis_crn
  domain_id = var.zone_id
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
  value = ibm_cis_dns_record.test_dns_loc_record
}
```

## Example usage 6 : Create CAA record

```terraform
resource "ibm_cis_dns_record" "test_dns_caa_record" {
  cis_id  = var.cis_crn
  domain_id = var.zone_id
  name    = "test-exmple.caa"
  type    = "CAA"
  ttl     = 900
  data = {
    tag   = "http"
    value = "domain.com"
  }
}

output "caa_record_output" {
  value = ibm_cis_dns_record.test_dns_caa_record
}
```

## Example usage 7 : Create SRV record

```terraform
resource "ibm_cis_dns_record" "test_dns_srv_record" {
  cis_id  = var.cis_crn
  domain_id = var.zone_id
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
  value = ibm_cis_dns_record.test_dns_srv_record
}
```

## Example usage 8 : Create SPF record

```terraform
resource "ibm_cis_dns_record" "test_dns_spf_record" {
  cis_id  = var.cis_crn
  domain_id = var.zone_id
  name    = "test-exmple.spf"
  type    = "SPF"
  content = "test"
}

output "spf_record_output" {
  value = ibm_cis_dns_record.test_dns_spf_record
}
```

## Example usage 9 : Create TXT record

```terraform
resource "ibm_cis_dns_record" "test_dns_txt_record" {
  cis_id  = var.cis_crn
  domain_id = var.zone_id
  name    = "test-exmple.txt"
  type    = "TXT"
  content = "test"
}

output "txt_record_output" {
  value = ibm_cis_dns_record.test_dns_txt_record
}
```

## Example usage 10 : Create NS record

````terraform
resource "ibm_cis_dns_record" "test_dns_ns_record" {
  cis_id  = var.cis_crn
  domain_id = var.zone_id
  name    = "test-exmple.ns"
  type    = "NS"
  content = "ns1.name.ibm.com"
}

output "ns_record_output" {
  value = ibm_cis_dns_record.test_dns_ns_record
}

output "caa_record_output" {
  value = ibm_cis_dns_record.test_dns_caa_record
}

## Example usage 7 : Create SRV record

```terraform
resource "ibm_cis_dns_record" "test_dns_srv_record" {
  cis_id  = var.cis_crn
  domain_id = var.zone_id
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
  value = ibm_cis_dns_record.test_dns_srv_record
}
````

## Example usage 8 : Create SPF record

```terraform
resource "ibm_cis_dns_record" "test_dns_spf_record" {
  cis_id  = var.cis_crn
  domain_id = var.zone_id
  name    = "test-exmple.spf"
  type    = "SPF"
  content = "test"
}

output "spf_record_output" {
  value = ibm_cis_dns_record.test_dns_spf_record
}
```

## Example usage 9 : Create TXT record

```terraform
resource "ibm_cis_dns_record" "test_dns_txt_record" {
  cis_id  = var.cis_crn
  domain_id = var.zone_id
  name    = "test-exmple.txt"
  type    = "TXT"
  content = "test"
}

output "txt_record_output" {
  value = ibm_cis_dns_record.test_dns_txt_record
}
```

## Example usage 10 : Create NS record

```terraform
resource "ibm_cis_dns_record" "test_dns_ns_record" {
  cis_id  = var.cis_crn
  domain_id = var.zone_id
  name    = "test-exmple.ns"
  type    = "NS"
  content = "ns1.name.ibm.com"
}

output "ns_record_output" {
  value = ibm_cis_dns_record.test_dns_ns_record
}
```

## Example usage 11 : Create PTR record

```terraform
resource "ibm_cis_dns_record" "test_dns_ptr_record" {
  cis_id  = var.cis_crn
  domain_id = var.zone_id
  name    = "1.2.3.4"
  type    = "PTR"
  content = "test-exmple.ptr.com"
}

output "ns_record_output" {
  value = ibm_cis_dns_record.test_dns_ptr_record
}
```

## Argument reference
Review the argument references that you can specify for your resource. 


- `content` - (Optional, String) The value of the record. For example, `192.168.127.127`. You need to provide this or data to be specified.
- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain to add a DNS record. It can be a combination of `<domain_id>:<cis_id> or <domain_id>`.
- `data` (Optional, Map) A map of attributes that constitute the record value. This value is required for `LOC`, `CAA` and `SRV` record types.

  Nested scheme for `data`:
  - `altitude` - (Optional, Integer) The `LOC` altitude. Mandatory field for `LOC` record type.
  - `lat_degrees` - (Optional, Integer) The `LOC` latitude degrees. Mandatory field for `LOC` record type.
  - `lat_direction` - (Optional, String) The `LOC` latitude direction `N`, `E`, `S`, `W`. Mandatory field for `LOC` record type.
  - `lat_minutes` - (Optional, Integer) The `LOC` latitude minutes. Mandatory field for `LOC` record type.
  - `lat_seconds` - (Optional, Integer) The `LOC` latitude seconds. Mandatory field for `LOC` record type.
  - `long_degrees` - (Optional, Integer) The `LOC` longitude degrees. Mandatory field for `LOC` record type.
  - `long_direction` - (Optional, String) The `LOC` longitude direction `N`, `E`, `S`, `W`. Mandatory field for `LOC` record type.
  - `long_minutes` - (Optional, Integer) The `LOC` longitude minutes. Mandatory field for `LOC` record type.
  - `long_seconds` - (Optional, Integer) The `LOC` longitude seconds. Mandatory field for `LOC` record type.
  - `port` - (Optional, Integer) The port number of the target server. Mandatory field for `SRV` record type.
  - `protocol` - (Optional, Integer) The symbolic name of the required protocol. Mandatory field for `SRV` record type.
  - `precision_horz` - (Optional, Integer) The `LOC` horizontal precision. Mandatory field for `LOC` record type.
  - `precision_vert` - (Optional, Integer) The `LOC` vertical precision. Mandatory field for `LOC` record type.
  - `priority` - (Optional, Integer) The priority of the record.
  - `service` - (Optional, Integer) The symbolic name of the required service, start with an underscore `_`. Mandatory field for `SRV` record type.
  - `size` - (Optional, Integer) The `LOC` altitude size. Mandatory field for `LOC` record type.
  - `weight` - (Optional, Integer) The weight of distributing queries among multiple target servers. Mandatory field for `SRV` record type.
- `name` - (Required, String) The name of the DNS record.
- `proxied`- (Optional, Bool) Indicates the record gets CIS's origin protection. Default is **false**.
- `priority` - (Optional, String) The priority of the record. Mandatory field for `SRV` record type.
- `type` - (Required, String) The type of the DNS record to be created. Allowed values are `A`, `AAAA`, `CNAME`, `NS`, `MX`, `TXT`, `LOC`, `SRV`, `SPF`, or `CAA`.
- `ttl` - (Optional, Integer) The time to live `(TTL)` record. The automatic is `ttl=1`. if the record is proxied. Terraform provider takes `TTL` in unit seconds. Therefore, it starts with value 120.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `created_on` - (String) The created date of the DNS record.
- `id` - (String) The ID of the record, zone and CRN with `:` separator.
- `modified_on` - (String) The modified date of the DNS record.
- `name` - (String) The name of the DNS record.
- `proxiable`- (Bool) Indicates if the record can be proxied.
- `proxied`- (Bool) Indicates the record gets CIS's origin protection. Default is **false**.
- `record_id` - (String) The DNS record ID.
- `zone_name` - (String) The DNS zone name.

## Import
The `ibm_cis_dns_record` resource can be imported by using the ID. The ID is formed from the DNS record ID, the domain ID, and the CRN (Cloud Resource Name). All values are  Concatenated  by using a `:` character. 

The domain ID and CRN are located on the **Overview** page of the internet services instance in the **Domain** heading of the console, or via using the `ibmcloud cis` command line commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **DNS Record ID** is a 32 digit character string of the form: `489d96f0da6ed76251b475971b097205c`. The ID of an existing DNS record is not available via the console. You can retrieve programmatically via the CIS API or via the command line using the CIS command `ibmcloud cis dns-records <domain_id>` to list the defined DNS records.

**Syntax**

```
$ terraform import ibm_cis_dns_record.myorg <dns_record_id>:<domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_dns_record.myorg  48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```

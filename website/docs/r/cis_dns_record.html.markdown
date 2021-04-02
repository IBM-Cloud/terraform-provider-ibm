---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_dns_record"
description: |-
  Provides a IBM CIS DNS Record resource.
---

# ibm_cis_dns_record

Provides a IBM CIS DNS Record resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS Domain resource.

## Example Usage 1 : Create A Record

```hcl
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

## Example Usage 2 : Create AAAA record

```hcl
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

## Example Usage 3 : Create CNAME record

```hcl
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

## Example Usage 4 : Create MX record

```hcl
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

## Example Usage 5 : Create LOC record

```hcl
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

## Example Usage 6 : Create CAA record

```hcl
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

## Example Usage 7 : Create SRV record

```hcl
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

## Example Usage 8 : Create SPF record

```hcl
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

## Example Usage 9 : Create TXT record

```hcl
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

## Example Usage 10 : Create NS record

````hcl
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

## Example Usage 7 : Create SRV record

```hcl
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

## Example Usage 8 : Create SPF record

```hcl
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

## Example Usage 9 : Create TXT record

```hcl
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

## Example Usage 10 : Create NS record

```hcl
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

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance
- `domain_id` - (Required,string) The ID of the domain to add the DNS record to. IT can either be a combination of <domain_id>:<cis_id> or <domain_id>
- `type` - (Required, string) The type of the DNS record to be created. Supported Record types are: A, AAAA, CNAME, LOC, TXT, MX, SRV, SPF, NS, CAA.
- `name` - (Required, string) The name of a DNS record.
- `content` - (Optional,string) The (string) value of the record, e.g. "192.168.127.127". Either this or `data` must be specified
- `ttl`-(Optional,int) TTL of the record. It should be automatic(i.e ttl=1) if the record is proxied. Terraform provider takes ttl in unit seconds. Therefore, it starts with value 120.
- `priority` - (Optional, int) The priority of the record. Mandatory field for SRV record type.
- `data` - (Optional,map) Map of attributes that constitute the record value. Only for LOC, CAA and SRV record types. Either this or `content` must be specified
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
  - `priority` - (Optional,int) The priority of the record
- `proxied` - (Optional,bool) Whether the record gets CIS's origin protection; defaults to `false`.

## Attributes Reference

The following attributes are exported:

- `id` - The identifier which consists of record id, zone id and crn with `:` seperator.
- `record_id` - The DNS record identifier.
- `name` - The name of a DNS record.
- `proxiable` - Whether the record has option to set proxied.
- `proxied` - Whether the record gets CIS's origin protection; defaults to `false`.
- `created_on` - The DNS record created date.
- `modified_on`- The DNS record modified date.
- `zone_name` - The DNS zone name.

## Import

The `ibm_cis_dns_record` resource can be imported using the `id`. The ID is formed from the `Dns Record ID`, the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated using a `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Dns Record ID** is a 32 digit character string of the form: `489d96f0da6ed76251b475971b097205c`. The id of an existing DNS record is not avaiable via the UI. It can be retrieved programatically via the CIS API or via the CLI using the CIS command to list the defined DNS recordss: `ibmcloud cis dns-records <domain_id>`

```
$ terraform import ibm_cis_dns_record.myorg <dns_record_id>:<domain-id>:<crn>

$ terraform import ibm_cis_dns_record.myorg  48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```

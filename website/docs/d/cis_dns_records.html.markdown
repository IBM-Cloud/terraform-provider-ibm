---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM : Cloud Internet Service DNS Record"
description: |-
Manages IBM Cloud Internet Service DNS Record.

---

# ibm_network_cis_dns_record

Import the details of an existing IBM Cloud Internet Service domain name service record as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl

data "ibm_cis_dns_records" "test" {
  cis_id    = var.cis_crn
  domain_id = var.zone_id
  file      = "records.txt"
}

```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required, string) The resource cis id of the CIS on which zones were created.
- `domain_id` - (Required, string) The resource domain id of the DNS on which zones were created.
- `file` - (Optional, string) The file which dns records to be exported.

## Attribute Reference

The following attributes are exported:

- `cis_dns_records` - The list of DNS records.
  - `id` - The identifier which consists of record id, zone id and crn with `:` seperator.
  - `record_id` - The DNS record identifier.
  - `name` - The name of a DNS record.
  - `proxiable` - Whether the record has option to set proxied.
  - `proxied` - Whether the record gets CIS's origin protection; defaults to `false`.
  - `created_on` - The DNS record created date.
  - `modified_on`- The DNS record modified date.
  - `zone_name` - The DNS zone name.
  - `type` - The type of the DNS record to be created. Supported Record types are: A, AAAA, CNAME, LOC, TXT, MX, SRV, SPF, NS, CAA.
  - `content` - The (string) value of the record.
  - `ttl`-TTL of the record. It should be automatic(i.e ttl=1) if the record is proxied. Terraform provider takes ttl in unit seconds.
  - `priority` - The priority of the record. Mandatory field for SRV record type.
  - `data` - Map of attributes that constitute the record value. Only for LOC, CAA and SRV record types.

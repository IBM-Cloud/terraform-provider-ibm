---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM : Cloud Internet Service DNS Record"
description: |-
  Manages IBM Cloud Internet Service DNS records.
---

# ibm_cis_dns_records
Retrieve information about an IBM Cloud Internet Services domain name service record. For more information, about DNS records, refer to [Managing DNS records](https://cloud.ibm.com/docs/dns-svcs?topic=dns-svcs-managing-dns-records). 

## Example usage

```terraform

data "ibm_cis_dns_records" "test" {
  cis_id    = var.cis_crn
  domain_id = var.zone_id
  file      = "records.txt"
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance on which zones were created.
- `domain_id` - (Required, String) The resource domain ID of the DNS on which zones were created.
- `file`-  (Optional, String) The file that DNS records to be exported.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `cis_dns_records` - (List) The list of DNS records.

  Nested scheme for `cis_dns_records`:
  - `created_on` - (String) The created date of the DNS record.
  - `data` - (String) Map of attributes that constitute the record value. Only for `LOC`, `CAA` and `SRV` record types.
  - `id` - (String) The ID which consists of record ID, zone ID and CRN with `:` separator.
  - `modified_on` - (String) The modified date of the DNS record.
  - `name` - (String) The name of a DNS record.
  - `priority` - (String) The priority of the record. Mandatory field for `SRV` record type.
  - `proxiable` - (String) Whether the record has option to set proxied.
  - `proxied` - (String) Whether the record gets CIS's origin protection; defaults to **false**.
  - `record_id` - (String) The DNS record identifier.
  - `type` - (String) The type of the DNS record to be created. Supported Record types are `A`, `AAAA`, `CNAME`, `LOC`, `TXT`, `MX`, `SRV`, `SPF`, `NS`, `CAA`.
  - `ttl` - (String) TTL of the record. It should be automatic that is `ttl=1`, if the record is proxied. Terraform provider takes `ttl` in unit seconds.
  - `zone_name` - (String) The DNS zone name.

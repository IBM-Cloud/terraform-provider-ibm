---

layout: "ibm"
page_title: "IBM : Cloud Internet Service DNS Record"
sidebar_current: "docs-ibm-datasource-cis-dns-records"
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
}

```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required, string) The resource cis id of the CIS on which zones were created.
- `domain_id` - (Required, string) The resource domain id of the DNS on which zones were created.

## Attribute Reference

The following attributes are exported:

- `id` - The identifier which consists of record id, zone id and crn with `:` seperator.
- `record_id` - The DNS record identifier.
- `name` - The name of a DNS record.
- `proxied` - Whether the record gets CIS's origin protection; defaults to `false`.
- `created_on` - The DNS record created date.
- `modified_on`- The DNS record modified date.
- `zone_name` - The DNS zone name.

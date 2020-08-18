---

layout: "ibm"
page_title: "IBM : Cloud Internet Service DNS Record"
sidebar_current: "docs-ibm-datasources-cis-dns-record"
description: |-
Manages IBM Cloud Internet Service DNS Record.

---

# ibm_network_cis_dns_record

Import the details of an existing IBM Cloud Internet Service domain name service record as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl

data "ibm_network_cis_dns_records" "test" {
  crn     = var.cis_crn
  zone_id = var.zone_id
}

```

## Argument Reference

The following arguments are supported:

- `crn` - (Required, string) The resource crn id of the CIS on which zones were created.
- `zone_id` - (Required, string) The resource zone id of the DNS on which zones were created.

## Attribute Reference

The following attributes are exported:

- `crn` - The CRN number of instance.
- `zone_id` - The ID of the DNS zone.
- `type` - The type of the DNS record to be created. Supported Record types are: A, AAAA, CNAME, LOC, TXT, MX, SRV, SPF, NS, CAA.
- `name` - The name of a DNS record.
- `content` - The content of a DNS record.
- `ttl` - The time-to-live value of the DNS record to be created.
- `priority` - The priority of the record. Mondatory for MX record.
- `created_on` - The time (Created On) of the DNS resource record.
- `modified_on` - The time (Modified On) of the DNS rsource record.

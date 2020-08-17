---

layout: "ibm"
page_title: "IBM : Networking CIS DNS Record"
sidebar_current: "docs-ibm-datasources-network-cis-dns-record"
description: |-
Manages IBM Cloud Internet Service DNS Record.

---

# ibm_network_cis_dns_record

Import the details of an existing IBM Cloud Internet Service domain name service record as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl

data "ibm_network_cis_dns_records" "test" {
  depends_on = [ibm_network_cis_dns_record.test_dns_a_record,
  ibm_network_cis_dns_record.test_dns_aaaa_record,
  ibm_network_cis_dns_record.test_dns_cname_record,
  ibm_network_cis_dns_record.test_dns_mx_record,
  ibm_network_cis_dns_record.test_dns_loc_record,
  ibm_network_cis_dns_record.test_dns_caa_record,
  ibm_network_cis_dns_record.test_dns_srv_record,
  ibm_network_cis_dns_record.test_dns_spf_record,
  ibm_network_cis_dns_record.test_dns_txt_record,
  ibm_network_cis_dns_record.test_dns_ns_record
  ]
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

---
layout: "ibm"
page_title: "IBM: ibm_cis_ip_addresses"
sidebar_current: "docs-ibm-datasource-cis-ip-addresses"
description: |-
  List the IP addresses used by name servers by Cloud Internet Services. Required for setting whitelist addresses for internet facing application ports.
---

# ibm\_cis_ip_addresses

Import the name of an existing domain as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_cis_ip_addresses" "domain_id" {
}
```


## Argument Reference

No arguments are required.

## Attribute Reference

The following attributes are exported:

* `ipv4_cidrs` - The ipv4 address ranges used by CIS for name servers.
* `ipv6_cidrs` - The ipv6 address ranges used by CIS for name servers.

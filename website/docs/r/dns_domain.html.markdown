---
layout: "ibm"
page_title: "IBM: dns_domain"
sidebar_current: "docs-ibm-resource-dns-domain"
description: |-
  Manages IBM DNS domains.
---

# ibm\_dns_domain

This resource represents a single DNS domain managed on Bluemix Infrastructure (SoftLayer). Domains contain general information about the domain name, such as the name and serial. 

Individual records, such as `A`, `AAAA`, `CTYPE`, and `MX` records, are stored in the domain's associated resource records using the `ibm_dns_domain_record` resource.


## Example Usage

```hcl
resource "ibm_dns_domain" "dns-domain-test" {
    name = "dns-domain-test.com"
    target = "127.0.0.10"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) A domain's name, including top-level domain. For example, "example.com". When the domain is created, proper `NS` and `SOA` records are created automatically for it.
* `target` - (Optional, string) The primary target IP address that the domain resolves to. When created, an `A` record with a host value of `@`, and a data-target value of the IP address, are provided and associated with the new domain.

## Attributes Reference

The following attributes are exported

* `id` - The domain record's internal identifier.
* `serial` - A unique number denoting the latest revision of a domain.
* `update_date` - The date that this domain record was last updated.

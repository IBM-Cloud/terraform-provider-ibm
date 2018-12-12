---
layout: "ibm"
page_title: "IBM: dns_domain"
sidebar_current: "docs-ibm-resource-dns-domain"
description: |-
  Manages IBM DNS domains.
---

# ibm\_dns_domain

Provides a single DNS domain managed on IBM Cloud Infrastructure (SoftLayer). Domains contain general information about the domain name, such as the name and serial number.

Individual records, such as `A`, `AAAA`, `CTYPE`, and `MX` records, are stored in the domain's associated resource records using the [`ibm_dns_record` resource](../r/dns_record.html).


## Example Usage

```hcl
resource "ibm_dns_domain" "dns-domain-test" {
    name = "dns-domain-test.com"
    target = "127.0.0.10"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The domain's name, including the top-level domain. For example, "example.com". When the domain is created, proper `NS` and `SOA` records are created automatically for the domain.
* `target` - (Optional, string) The primary target IP address to which the domain resolves. When the domain is created, an `A` record with a host value of `@` and a data-target value of the IP address are provided and associated with the new domain
* `tags` - (Optional, array of strings) Tags associated with the DNS domain instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique internal identifier of the domain record.
* `serial` - A unique number denoting the latest revision of the domain.
* `update_date` - The date that the domain record was last updated.

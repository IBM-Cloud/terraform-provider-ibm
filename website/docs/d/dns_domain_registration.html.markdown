---
subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: ibm_dns_domain_registration"
description: |-
  Get resource identifier for IBM Cloud DNS domain registration.
---

# ibm_dns_domain_registration
Retrieve information of an existing IBM DNS domain registration service. The domain must initially be registered through the console of the IBM Cloud DNS registration service. Typically, the domain registration data source is used in configuration with global load-balancing services. For example, Cloudflare, Akamai or IBM Cloud Internet Services (Cloudflare). For more information, about DNS domain registration, see [getting started with Domain Name registration](https://cloud.ibm.com/docs/dns?topic=dns-getting-started).

## Example usage
The following example shows how you can use this data source to reference the domain ID in the `ibm_dns_registration_nameservers` resource, since the numeric IDs are often unknown.

```terraform
data "ibm_dns_domain_registration" "dnstestdomain" {
    name = "dnstestdomain.com"
}

resource "ibm_dns_domain_registration_nameservers" "dnstestdomain" {
  dns_registration_id = data.ibm_dns_domain_registration.dnstestdomain.id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `name` - (Required, String) The name of the DNS domain registration as it was defined in IBM Cloud Infrastructure DNS Registration Service.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier of the domain.

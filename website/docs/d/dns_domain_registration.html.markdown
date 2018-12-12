---
layout: "ibm"
page_title: "IBM: ibm_dns_domain_registration"
sidebar_current: "docs-ibm-datasource-dns-domain-registration"
description: |-
  Get resource identifier for IBM DNS domain registration.
---

# ibm\_dns_domain_registration

Import an existing domain registration from the IBM DNS Domain Registration Service as a read-only data source. This DNS registration resource can then be used as a base from which DNS configurations can be performed via Terraform. The domain must inititally be registered via the UI of the IBM Cloud DNS Registration Service. Typically the Domain Registration datasource is used in configuration with global load balancing services, e.g. CloudFlare, Akamai or IBM Cloud Internet Services (Cloudflare). For additional usage instructions see the resource `ibm_dns_domain_registration_nameservers`. 

## Example Usage

```hcl
data "ibm_dns_domain_registration" "dnstestdomain" {
    name = "dnstestdomain.com"
}
```

The following example shows how you can use this data source to reference the domain ID in the `ibm_dns_registration_nameservers` resource, since the numeric IDs are often unknown.

```hcl
resource "ibm_dns_domain_registration_nameservers" "dnstestdomain" {
    ...
    dns_registration_id = "${data.ibm_dns_domain_registration.dnstestdomain.id}"
    ...
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the DNS domain registration as it was defined in IBM Cloud Infrastructure DNS Registration Service (SoftLayer).

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the domain.

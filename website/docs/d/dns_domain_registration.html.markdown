---
layout: "ibm"
page_title: "IBM: ibm_dns_domain_registration"
sidebar_current: "docs-ibm-datasource-dns-domain-registration"
description: |-
  Get resource identifier for IBM DNS domain registration.
---

# ibm\_dns_domain_registration

Import the name of an existing DNS domain registration as a read-only data source. You can then reference this data source in other resources within the same configuration by using interpolation syntax. The domain must inititally be registered via the UI of the IBM Cloud DNS Registration Service. The Domain Registration datasource is used in configuration of IBM Cloud Internet Services. See resource ibm_dns_domain_registration_nameservers. 

## Example Usage

```hcl
data "ibm_dns_domain_registration" "dns-domain-test" {
    name = "test-domain.com"
}
```

The following example shows how you can use this data source to reference the domain ID in the `ibm_dns_registration_nameservers` resource, since the numeric IDs are often unknown.

```hcl
resource "ibm_dns_domain_registration_nameservers" "dns-domain-test" {
    ...
    dns_registration_id = "${data.ibm_dns_domain_registration.dns-domain-test.id}"
    ...
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the DNS domain registration, as it was defined in IBM Cloud Infrastructure DNS Registration Service (SoftLayer).

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the domain.

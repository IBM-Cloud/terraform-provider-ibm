---
layout: "ibm"
page_title: "IBM: ibm_dns_domain"
sidebar_current: "docs-ibm-datasource-dns-domain"
description: |-
  Get information about an IBM DNS domain resource.
---

# ibm\_dns_domain

Import the name of an existing domain as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_dns_domain" "domain_id" {
    name = "test-domain.com"
}
```

The following example shows how you can use this data source to reference the domain ID in the `ibm_dns_record` resource, since the numeric IDs are often unknown.

```hcl
resource "ibm_dns_record" "www" {
    ...
    domain_id = "${data.ibm_dns_domain.domain_id.id}"
    ...
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the domain, as it was defined in Bluemix Infrastructure (SoftLayer).

## Attributes Reference

The following attributes are exported:

* `id` - The unique identifier of the domain.

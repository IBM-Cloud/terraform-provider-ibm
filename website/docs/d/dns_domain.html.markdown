---
subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: ibm_dns_domain"
description: |-
  Get information about an IBM Cloud DNS domain resource.
---

# ibm_dns_domain
Retrieve information of an existing domain as a read-only data source. For more information, about DNS resource, see [managing DNS zones in classic infrastructure](https://cloud.ibm.com/docs/dns?topic=dns-manage-dns-zones).

## Example usage
The following example shows how you can use the data source to reference the domain ID in the `ibm_dns_record` resource, as the numeric IDs are often unknown.

```terraform
data "ibm_dns_domain" "domain_id" {
    name = "test-domain.com"
}

resource "ibm_dns_record" "www" {
    domain_id = data.ibm_dns_domain.domain_id.id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `name` - (Required, String) The name of the domain, as defined in IBM Cloud Classic Infrastructure (SoftLayer). 

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier of the domain.

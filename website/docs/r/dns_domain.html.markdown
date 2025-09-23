---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: dns_domain"
description: |-
  Manages IBM DNS domain.
---

# ibm_dns_domain
Provides a single DNS domain managed on IBM Cloud Classic Infrastructure (SoftLayer). Domains contain general information about the domain name, such as the name and serial number. For more information, about DNS services, see [getting started with IBM Cloud DNS Services](https://cloud.ibm.com/docs/dns-svcs?topic=dns-svcs-getting-started).

Individual records, such as `A`, `AAAA`, `CTYPE`, and `MX` records, are stored in the domain's associated resource records by using the [`ibm_dns_record` resource]((https://github.com/Mavrickk3/terraform-provider-ibm/tree/master/website/docs/r/dns_record.html.markdown)).


## Example usage

```terraform
resource "ibm_dns_domain" "dns-domain-test" {
    name = "dns-domain-test.com"
    target = "127.0.0.10"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `name` - (Required, Forces new resource, String) The name of the domain. For example, "example.com". When the domain is created, proper `NS` and `SOA` records are created automatically for the domain.
- `target`- (Optional, String) The primary target IP address to which the domain resolves. When the domain is created, an `A` record with a host value of `@` and a data-target value of the IP address are provided and associated with the new domain.
- `tags`- (Optional, Array of Strings) Tags associated with the DNS domain instance. **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`- (String) The unique internal identifier of the domain record.
- `serial`- (String) A unique number denoting the latest revision of the domain.
- `update_date`- (Timestamp) The date that the domain record was last updated.


---
layout: "ibm"
page_title: "IBM: dns_domain_registration_nameservers"
sidebar_current: "docs-ibm-resource-dns-domain-registration-nameservers"
description: |-
  Manages the nameservers on IBM DNS domain registrations.
---

# ibm\_dns_domain_registration_nameservers

This resource configures the (custom) name servers associated with a DNS domain registered with the IBM Cloud DNS Registration Service. This is used to delegate DNS domain management to another DNS provider typically for caching or DDoS protection. It is used with services including Akamai, CloudFlare or IBM Cloud Internet Services. DNS management for the domain is delegated by updating the IBM DNS registration record service with the name servers of the DNS service provider. 

This resource updates the (custom) name servers specified in the record for the domain in the IBM DNS registration service, with the new name servers. The original name servers (ns1.softlayer.com and ns2.softlayer.com) are saved and restored when the resource is deleted. Creation of this resource directs DNS management for the domain to the new DNS provider and over-rides an DNS records created on IBM Cloud by the `dns_record` resource. 

The only the name_server attribute of the domain record in the DNS Registration Service can be updated. No ability is provided to create of delete a domain registration to avoid accidental loss of the registration. The domain registration to be modified is identified using a read only dns_domain_registration data source. 

The creation of an IBM Cloud Internet Services instance with a `ibm_cis_domain` resource will export two name servers of the form ns001.name.cloud.ibm.com. By intepolation these can be passed to this resource to configure the name servers at the IBM DNS registrar. 


## Example Usage

```hcl
resource "ibm_dns_domain_registration_nameservers" "dnstestdomain" {
    dns_registration_id = "${data.ibm_dns_domain_registration.dnstestdomain.id}"
    name_servers = "${ibm_cis_domain.dnstestdomain.name_servers}" 
}
data "ibm_dns_domain_registration" "dnstestdomain" {
    name = "dnstestdomain.com"
}
resource "ibm_cis_domain" "dnstestdomain" {
   
}
```

Or 

```hcl
resource "ibm_dns_domain_registration_nameservers" "dns-domain-test" {
    dns_registration_id = "${data.ibm_dns_domain_registration.dns-domain-test.id}"
    name_servers = ["ns006.name.cloud.ibm.com", "ns017.name.ibm.cloud.com"] 
}
data "ibm_dns_domain_registration" "dns-domain-test" {
    name = "test-domain.com"
}
```


## Argument Reference

The following arguments are supported:

* `dns_registration_id` - (Required, string) The unique id of the domain's registration. This is exported by the ibm_dns_domain_registration data source. 
* `name_servers` - (Required, Array of strings) E.g. an array of name servers returned from configuration of a domain on a instance of IBM Cloud Internet Services. This is of the format: ["ns006.name.cloud.ibm.com", "ns017.name.cloud.ibm.com"]


## Attribute Reference

The following attributes are exported:

* `id` - The unique internal identifier of the domain registration record.
* `name_servers` - The new name servers pointing to the new DNS management service provider
* `original_name_servers` - The original name servers configured at the time of domain registration.

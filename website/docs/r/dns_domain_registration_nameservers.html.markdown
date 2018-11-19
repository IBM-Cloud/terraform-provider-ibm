---
layout: "ibm"
page_title: "IBM: dns_domain_registration_nameservers"
sidebar_current: "docs-ibm-resource-dns-domain-registration-nameservers"
description: |-
  Manages the nameservers on IBM DNS domain registrations.
---

# ibm\_dns_domain_registration_nameservers

Configures the (custom) name servers associated with a DNS domain registration managed by the IBM Cloud DNS Registration Service. The default IBM Cloud name servers specified when the domain was initially registered are replaced with the values passed when this resource is created. 

This resource is typically used in conjunction with IBM Cloud Internet Services to enable DNS services for the domain to be managed via IBM Cloud Internet Services. All futher configuration of the domain is then performed using the Cloud Internet Services resource instances. To transfer management control, the IBM Cloud DNS domain registration is updated with the Internet Services specific name servers. This step is required before the domain in Cloud Internet Services becomes active and will start serving web traffic. Using interpolation syntax, the computed name servers of the CIS resource are passed into this resource. 


## Example Usage

```hcl
resource "ibm_dns_domain_registration_nameservers" "dns-domain-test" {
    dns_registration_id = "${data.ibm_dns_domain_registration.dns-domain-test.id}"
    name_servers = "${ibm_cloud_internet_services.domain1.name_servers}" 
}
data "ibm_dns_domain_registration" "dns-domain-test" {
    name = "test-domain.com"
}
resource "ibm_cloud_internet_services" "domain1" {
   
}
```

Or 

```hcl
resource "ibm_dns_domain_registration_nameservers" "dns-domain-test" {
    dns_registration_id = "${data.ibm_dns_domain_registration.dns-domain-test.id}"
    name_servers = ["ns006.name.ibm.cloud.com", "ns017.name.ibm.cloud.com"] 
}
data "ibm_dns_domain_registration" "dns-domain-test" {
    name = "test-domain.com"
}
```


## Argument Reference

The following arguments are supported:

* `dns_registration_id` - (Required, string) The unique id of the domain's registration. This comes from the ibm_dns_domain_registration data source. 
* `name_servers` - (Required, Array of strings) An array of the name servers returned from configuration of a domain on a instance of IBM Cloud Internet Services. This is of the format: ["ns006.name.ibm.cloud.com", "ns017.name.ibm.cloud.com"]


## Attribute Reference

The following attributes are exported:

* `id` - The unique internal identifier of the domain registration record.
* `name_servers` - The name servers configured for the domain registration.


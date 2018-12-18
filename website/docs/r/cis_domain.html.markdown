---
layout: "ibm"
page_title: "IBM: app"
sidebar_current: "docs-ibm-resource-cis"
description: |-
  Assigns managmement of a DNS domain to Cloud Internet Services (CIS) 
---

# cis_domain

Creates a DNS Domain resource that represents a DNS domain assigned to CIS. A domain is the basic resource for working with Cloud Internet Services and is typically the first resouce that is assigned to the CIS service instance. The domain will not become `active` until the DNS Registrar is updated with the CIS name servers in the exported variable `name_servers`. Refer to the resource `dns_domain_registration_nameservers`for updating the DNS Registrars name servers. 

## Example Usage

```hcl
resource "ibm_cis_domain" "example" {
    domain = "example.com"
    cis_id = "${ibm_cis.instance.id}"
}

resource "ibm_cis" "instance" {
  name              = "test-domain"
  plan              = "standard"
```

## Argument Reference

The following arguments are supported: 

* `domain` - (Required) The DNS domain name which will be added to CIS and managed.
* `cis_id` - (Required) The ID of the CIS service instance



## Attributes Reference

The following attributes are exported:

* `id` - The domain ID.
* `paused` - Boolean of whether this domain is paused (traffic bypasses CIS). Default: false.
* `status` - Status of the domain. Valid values: `active`, `pending`, `initializing`, `moved`, `deleted`, `deactivated`. After creation, the status will remain pending until the DNS Registrar is updated with the CIS name servers, exported in the 'name_servers' variable. 
* `name_servers` - The IBM CIS assigned name servers, to be passed by interpolation to the resource dns_`domain_registration_nameservers`.
* `original_name_servers` - The name servers from when the Domain was initially registered with the DNS Registrar.  

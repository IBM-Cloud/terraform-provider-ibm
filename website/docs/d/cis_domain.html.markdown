---
layout: "ibm"
page_title: "IBM : Cloud Internet Services Domain"
sidebar_current: "docs-ibm-datasource-cis-domain"
description: |-
  Get information on an IBM Cloud Internet Services Domain.
---

# ibm\_cis_domain

Imports a read only copy of an existing Internet Services domain resource. This allows new/additional CIS sub-resources to be added to an existing CIS domain registration, specifically DNS records and global load balancers. It is used in conjunction with the CIS data-source. 

## Example Usage

```hcl
data "ibm_cis_domain" "cis_instance_domain" {
  domain = "example.com"
  cis_id = "${ibm_cis.instance.id}"
}
data "ibm_cis" "cis_instance" {
  name              = "test"
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required) The DNS domain name which will be added to CIS and managed.
* `cis_id` - (Required) The ID of the CIS service instance

## Attribute Reference

The following attributes are exported:

* `id` - The domain ID.
* `paused` - Boolean of whether this domain is paused (traffic bypasses CIS). Default: false.
* `status` - Status of the domain. Valid values: `active`, `pending`, `initializing`, `moved`, `deleted`, `deactivated`. After creation, the status will remain pending until the DNS Registrar is updated with the CIS name servers, exported in the 'name_servers' variable. 
* `name_servers` - The IBM CIS assigned name servers, to be passed by interpolation to the resource dns_`domain_registration_nameservers`.
* `original_name_servers` - The name servers from when the Domain was initially registered with the DNS Registrar.  
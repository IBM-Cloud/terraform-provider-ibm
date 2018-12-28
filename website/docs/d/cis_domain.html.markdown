---
layout: "ibm"
page_title: "IBM : Cloud Internet Services Domain"
sidebar_current: "docs-ibm-datasource-cis_domain"
description: |-
  Get information on an IBM Cloud Internet Services Domain.
---

# ibm\_cis

Imports the a read only copy of the details of an existing Internet Services resource. This allows CIS sub-resources to be added to an existing CIS instance. This includes domains, DNS records, pools, healthchecks and global load balancers. 

## Example Usage

```hcl
data "ibm_cis" "cis_instance" {
  domain = "example.com"
  cis_id = "${ibm_cis.instance.id}"
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
---
layout: "ibm"
page_title: "IBM: app"
sidebar_current: "docs-ibm-resource-cis"
description: |-
  Assigns managmement of a DNS domain to Cloud Internet Services (CIS) 
---

# ibm_cis_domain

Creates a DNS Domain resource that represents a DNS domain assigned to CIS. A domain is the basic resource for working with Cloud Internet Services and is typically the first resouce that is assigned to the CIS service instance. The domain will not become `active` until the DNS Registrar is updated with the CIS name servers in the exported variable `name_servers`. Refer to the resource `dns_domain_registration_nameservers`for updating the IBM Cloud DNS Registrars name servers. 

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

## Import

The `ibm_cis_domain` resource can be imported using the `id`. The ID is formed from the `Domain ID` of the domain concatentated using a `:` character with the `CRN` (Cloud Resource Name). 

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `bx cis` CLI commands.

* **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f` 

* **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`


```
$ terraform import ibm_cis_domain.myorg <domain-id>:<crn>

$ terraform import ibm_cis_domain.myorg  9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::

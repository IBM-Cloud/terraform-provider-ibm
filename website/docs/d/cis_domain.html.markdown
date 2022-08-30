---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM : Cloud Internet Services Domain"
description: |-
  Get information on an IBM Cloud Internet Services domain.
---

# ibm_cis_domain
Retrieve information about an existing Internet Services domain resource. This allows new CIS sub-resources to be added to an existing CIS domain registration, specifically DNS records and Global Load Balancers. It is used in conjunction with the CIS data source. For more information, about CIS DNS domain, see [setting up your Domain Name System for CIS](https://cloud.ibm.com/docs/cis?topic=cis-set-up-your-dns-for-cis).

## Example usage

```terraform
data "ibm_cis_domain" "cis_instance_domain" {
  domain = "example.com"
  cis_id = ibm_cis.instance.id
}

data "ibm_cis" "cis_instance" {
  name = "test"
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `domain` - (Required, String) The DNS domain name that is added and managed for an IBM Cloud Internet Services instance.
- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `id` - (String) The unique identifier of your domain.
- `domain_id` - (String) ID of the domain. 
- `name_servers` - (String) The IBM Cloud Internet Services assigned name servers, to be passed by interpolation to the resource dns_domain_registration_nameservers.
- `original_name_servers` - (String) The name servers from when the Domain was initially registered with the DNS Registrar.
- `paused` -  (Bool) If set to **true**, network traffic to this domain is paused. If set to **false**, network traffic to this domain is permitted. The default value is **false**.
- `status` - (String) The status of your domain. Valid values are `active`, `pending`, `initializing`, `moved`, `deleted`, and `deactivated`. After creation, the status remains pending until the DNS Registrar is updated with the CIS name servers, exported in the ‘name_servers’ variable.
- `type` - (String) The type of domain created. `full`- for regular domains, & `partial` for partial domain for CNAME setup.
---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: app"
description: |-
  Assigns managmement of a DNS domain to Cloud Internet Services.
---

# ibm_cis_domain
Creates a DNS Domain resource that represents a DNS domain assigned to Cloud Internet Services (CIS). A domain is the basic resource for working with Cloud Internet Services and is typically the first resouce that is assigned to the CIS service instance. The domain will not become `active` until the DNS Registrar is updated with the CIS name servers in the exported variable `name_servers`. Refer to the resource `dns_domain_registration_nameservers`for updating the IBM Cloud DNS Registrars name servers. For more information, about CIS DNS domain, see [setting up your Domain Name System for CIS](https://cloud.ibm.com/docs/cis?topic=cis-set-up-your-dns-for-cis).

## Example usage - 1 (Regular Domain)
```terraform
resource "ibm_cis_domain" "example" {
  domain = "example.com"
  cis_id = ibm_cis.instance.id
}

resource "ibm_cis" "instance" {
  name = "test-domain"
  plan = "standard-next"
}
```

## Example usage - 2 (Partial Domain)
```terraform
resource "ibm_cis_domain" "example" {
  domain = "example.com"
  cis_id = ibm_cis.instance.id
  type   = "partial"
}

resource "ibm_cis" "instance" {
  name = "test-domain"
  plan = "standard-next"
}
```


## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain` - (Required, String) The DNS domain name that you want to add to your IBM Cloud Internet Services instance.
- `type` - (String) The type of domain to be created. Default value is noted to be `full`- for regular domains, & to create a partial domain for CNAME setup, value to be used is `partial`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `domain_id` - (String) The ID of the domain.
- `id` - (String) The unique identifier of the domain.
- `name_servers` - (String) The name servers that are assigned to your IBM Cloud Internet Services instance.
- `original_name_servers` - (String) The name servers that were used when the domain was first registered with the DNS Registrar.
- `paused`- (Bool) Indicates if the domain is paused and network traffic bypasses your IBM Cloud Internet Services instance. The default values is **false**.
- `status` - (String) The status of the domain. Valid values are `active`, `pending`, `initializing`, `moved`, `deleted`, and `deactivated`. After creation, the status remains pending until the DNS Registrar is updated with the CIS name servers, exported in the `name_servers` variable.
- `verification_key` - (String) The verification key of the domain.
- `cname_suffix` - (String) The cname suffix of the domain.


## Import

The `ibm_cis_domain` resource can be imported using the `id`. The ID is formed from the `Domain ID` of the domain concatentated using a `:` character with the `CRN` (Cloud Resource Name). 

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or by using the `bx cis` CLI commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f` 

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

**Syntax**

```
$ terraform import ibm_cis_domain.myorg <domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_domain.myorg  9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```

---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_routing"
description: |-
  Provides a IBM CIS routing resource.
---

# ibm_cis_routing
Provides an IBM Cloud Internet Services (CIS) routing resource. This resource is associated with an IBM CIS instance and a CIS domain resource. It allows to change routing of a domain of an CIS instance. For more information, about CIS routing, see [routing concepts](https://cloud.ibm.com/docs/cis?topic=cis-cis-routing).

## Example usage

```terraform
# Change Routing of the domain

resource "ibm_cis_routing" "routing" {
	cis_id          = data.ibm_cis.cis.id
	domain_id       = data.ibm_cis_domain.cis_domain.domain_id
	smart_routing   = "on"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain where you want to change routing.
- `smart_routing` - (Optional, String) The smart routing to set enable or disable. Valid values are `on` and `off`.

**Note**

`tiered_caching` is not supported in this provider version.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The record ID. It is a combination of `<domain_id>,<cis_id>` attributes concatenated with `:`.

## Import
The `ibm_cis_routing` resource can be imported using the ID. The ID is formed from the domain ID of the domain and the CRN concatenated  using a `:` character.

The domain ID and CRN will be located on the overview page of the IBM Cloud Internet Services instance of the console domain heading, or by using the `ibmcloud cis` command line commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

**Syntax**

```
$ terraform import ibm_cis_routing.routing <domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_routing.routing 9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```

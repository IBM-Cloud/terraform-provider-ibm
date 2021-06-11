---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_filter"
description: |-
  Provides a IBM CIS Filter.
---

# ibm_cis_filter

Provides a IBM CIS Filter. This resource is associated with an IBM Cloud Internet Services instance and a CIS Domain resource. It allows to create, update, delete filter of a domain of a CIS instance

## Example Usage

```terraform
# Add a filter to the domain

resource "ibm_cis_filter" "test" {
  cis_id      = data.ibm_cis.cis.id
  domain_id   = data.ibm_cis_domain.cis_domain.domain_id
  expression  =  "(http.request.uri eq \"/test-update?number=212\")"
  paused      =  "false"
  description = "Filter-creation"
}

```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance
- `domain_id` - (Required,string) The ID of the domain to add the Filter.
- `expression` - (Required,string) The expression of filter.
- `paused` - (Optional,boolean). Whether this filter is currently disabled.
- `description` - (Required,string) Some useful information about this filter to help identify the purpose of it.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` - The record ID. It is a combination of <`domain_id`>,<`cis_id`> attributes concatenated with ":".

## Import

The `ibm_cis_filter` resource can be imported using the `id`. The ID is formed from the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated using a `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

```
$ terraform import ibm_cis_filter.myorg <domain-id>:<crn>

$ terraform import ibm_cis_filter.myorg 9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
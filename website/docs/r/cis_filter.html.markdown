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
- `description` - (Optional,string) Some useful information about this filter to help identify the purpose of it.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` - Unique identifier for the Filter. It is a combination of <`unique-id`>:<`domain-id`>:<`crn`> attributes concatenated with ":".
- `filter_id` - Unique identifier for the Filter.

## Import

The `ibm_cis_filter` resource can be imported using the `id`. The ID is formed from the `Filter ID`, the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated usinga `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Filter ID** is a 32 digit character string of the form: `1fc7c3247067ee00856729661c7d58c9`. The id of an existing Filteris not avaiable via the UI. It can be retrieved programmatically via the CIS API or via the CLI using the CIS command to list the defined Filter: `ibmcloud cis filter`

```
$ terraform import ibm_cis_filter.myorg <filter_id>:<domain-id>:<crn>

$ terraform import ibm_cis_domain.myorg  57d96f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
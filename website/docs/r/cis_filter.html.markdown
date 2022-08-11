---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_filter"
description: |-
  Provides a IBM Cloud CIS Filter.
---

# ibm_cis_filter

Provides a IBM CIS Filter. This resource is associated with an IBM Cloud Internet Services (CIS) instance and a CIS Domain resource. It allows to create, update, delete filter of a domain of a CIS instance. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis?topic=cis-about-ibm-cloud-internet-services-cis).

## Example usage

```terraform
# Add a filter to the domain

resource "ibm_cis_filter" "test" {
  cis_id      = data.ibm_cis.cis.id
  domain_id   = data.ibm_cis_domain.cis_domain.domain_id
  expression  =  "(http.request.uri eq \"/test-update?number=212\")"
  paused      =  false
  description = "Filter-creation"
}

```

## Argument reference
Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Required, String) The ID of the domain to add the Filter.
- `expression` - (Required, String) The expression of filter.
- `paused` - (Optional, Bool) Whether this filter is currently disabled.
- `description` - (Optional, String) The information about this filter to help identify the purpose of it.

## Attributes Reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of filter resource. It is a combination of <`filter-id`>:<`domain-id`>:<`crn`> attributes concatenated with ":".
- `filter_id` - (String) Unique identifier for the Filter.

## Import

The `ibm_cis_filter` resource can be imported using the `id`. The ID is formed from the `Filter ID`, the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated usinga `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Filter ID** is a 32 digit character string of the form: `d72c91492cc24d8286fb713d406abe91`. 

**Syntax**

```
$ terraform import ibm_cis_filter.myorg <filter_id>:<domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_filter.myorg
d72c91492cc24d8286fb713d406abe91:0b30801280dc2dacac1c3960c33b9ccb:crn:v1:bluemix:public:internet-svcs-ci:global:a/01652b251c3ae2787110a995d8db0135:9054ad06-3485-421a-9300-fe3fb4b79e1d::
```
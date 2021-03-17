---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_custom_page"
description: |-
  Provides a IBM CIS Custom Page resource.
---

# ibm_cis_custom_page

Provides a IBM CIS Custom Page resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS Domain resource. It allows to change Custom Page of a domain of a CIS instance

## Example Usage

```hcl
# Change Custom Page of the domain

resource "ibm_cis_custom_page" "custom_page" {
	cis_id    = data.ibm_cis.cis.id
	domain_id = data.ibm_cis_domain.cis_domain.domain_id
	page_id   = "basic_challenge"
	url       = "https://test.com/index.html"
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance.
- `domain_id` - (Required,string) The ID of the domain to change Custom Page.
- `page_id` - (Required, string) The Custom page identifier. Valid values are `basic_challenge, waf_challenge, waf_block, ratelimit_block, country_challenge, ip_block, under_attack, 500_errors, 1000_errors, always_online`
- `url` - (Required, string) The URL for custom page settings. By default `url` is set with empty string `""`. If this field is being set with empty string, when it is already set with empty string, then it throws error.

## Attributes Reference

The following attributes are exported:

- `id` - The record ID. It is a combination of <`page_id`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":".
- `description` - The description of custom page.
- `required_tokens` - (list) The custom page required token which is expected from `url` page.
- `preview_target` - The custom page target
- `state` - The custom page state. This is set `default` when there is empty `url` and set to `customized` when `url` is set with some url.
- `created_on` - The custom page created date and time.
- `modified_on` - The custom page modified date and time.

## Import

The `ibm_cis_custom_page` resource can be imported using the `id`. The ID is formed from the `page_id`, `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated using a `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **Page ID** is a string of the form: `basic_challenge`

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

```
$ terraform import ibm_cis_custom_page.custom_page <page_id>:<domain-id>:<crn>

$ terraform import ibm_cis_custom_page.custom_page basic_challenge:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```

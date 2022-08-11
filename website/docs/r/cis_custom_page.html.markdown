---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_custom_page"
description: |-
  Provides a IBM CIS custom page resource.
---

# ibm_cis_custom_page

 Provides an IBM Cloud Internet Services custom page resource that is associated with an IBM CIS instance and a CIS domain resource. It allows to create, update, and delete a custom page of a domain of a CIS instance. For more information about custom page, refer to [CIS custom page](https://cloud.ibm.com/docs/cis?topic=cis-custom-page).

## Example usage

```terraform
# Change Custom Page of the domain

resource "ibm_cis_custom_page" "custom_page" {
	cis_id    = data.ibm_cis.cis.id
	domain_id = data.ibm_cis_domain.cis_domain.domain_id
	page_id   = "basic_challenge"
	url       = "https://test.com/index.html"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain to change custom page.
- `page_id` - (Required, String) The custom page identifier. Valid values are `basic_challenge`, `waf_challenge`, `waf_block`, `ratelimit_block`, `country_challenge`, `ip_block`, `under_attack`, `500_errors`, `1000_errors`, `always_online`.
- `url` - (Required, String) The URL for custom page settings. By default URL is set with empty string `""`. Setting a duplicate empty string throws an error.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `created_on` - (String) Created date and time of the custom page.
- `description` - (String) The description of the custom page.
- `id` - (String) The record ID. It is a combination of `<domain_id>,<cis_id>` attributes concatenated with `:`.
- `modified_on` - (String) Modified date and time of the custom page.
- `preview_target` - (String) The custom page target.
- `required_tokens` (List)The custom page required token which is expected from URL page.
- `state` - (String) The custom page state. This is set default when there is an empty URL and can customize when URL is set with some URL.

## Import
The `ibm_cis_custom_page` resource can be imported by using the ID. The ID is formed from the page_id, domain ID of the domain and the CRN concatenated  using a `:` character.

The domain ID and CRN will be located on the overview page of the IBM Cloud Internet Services instance of the console domain heading, or by using the `ibmcloud cis` command line commands.

- **Page ID** is a string of the form: `basic_challenge`

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

**Syntax**

```
$ terraform import ibm_cis_custom_page.custom_page <page_id>:<domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_custom_page.custom_page basic_challenge:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```

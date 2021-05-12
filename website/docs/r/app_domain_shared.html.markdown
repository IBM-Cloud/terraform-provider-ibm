---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: app_domain_shared"
description: |-
  Manages IBM shared domain.
---

# `ibm_app_domain_shared`

Create, update, or delete a shared domain for your Cloud Foundry app. For more information, about an app domain shared, see [Managing your domains](https://cloud.ibm.com/docs/cloud-foundry-public?topic=cloud-foundry-public-custom-domains).


## Example usage
The following example creates the `example.com` shared domain. 

```
resource "ibm_app_domain_shared" "domain" {
  name              = "example.com"
  router_group_guid = "3hG5jkjk4k34JH5666"
  tags              = ["tag1", "tag2"]
}
```

## Argument reference
Review the input parameters that you can specify for your resource. 

- `name` - (Required, String) The name of the domain.
- `router_group_guid` - (Optional, String) The GUID of the router group.
- `tags` (Optional, Array of Strings) The tags that you want to add to the shared domain. Tags can help you find the domain more easily later.


## Attribute reference
Review the output parameters that you can access after your resource is created. 

- `id` - (String) The unique identifier of the shared domain.



---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: app_route"
description: |-
  Manages IBM application route.
---

# `ibm_app_route`

Create, update, or delete a route for your Cloud Foundry app. For more information, about an app route, see [Adding and using a custom domain](https://cloud.ibm.com/docs/cloud-foundry-public?topic=cloud-foundry-public-custom-domains)


## Example usage
The following example creates a route for the `example.com` shared domain. 


```
data "ibm_space" "spacedata" {
  space = "space"
  org   = "myorg.com"
}

data "ibm_app_domain_shared" "domain" {
  name = "example.com"
}

resource "ibm_app_route" "route" {
  domain_guid = data.ibm_app_domain_shared.domain.id
  space_guid  = data.ibm_space.spacedata.id
  host        = "myhost"
  path        = "/app"
}
```


## Argument reference
Review the input parameters that you can specify for your resource. 

- `domain_guid` - (Required, String) The GUID of the associated domain. You can retrieve the value from data source `ibm_app_domain_shared` or `ibm_app_domain_private`.
- `host` - (Optional, String) The hostname of the route. The hostname is required for shared domains.
- `path` - (Optional, String) The path for the route. Paths must be between 2-128 characters, must start with a forward slash (/), and cannot contain a question mark (?).
- `port` - (Optional, String) The port of the route. This option is supported for TCP router group domains only.
- `space_guid` - (Required, String)  The GUID of the Cloud Foundry space where you want to create the route. You can retrieve the value from data source `ibm_space` or by running the `ibmcloud iam space <space_name> guid` command in the IBM Cloud CLI.
- `tags` (Optional, Array of Strings) The tags that you want to add to the route. Tags can help you find the route more easily later. **Note** Currently, `Tags` that are managed locally and not sored on the IBM Cloud service endpoint.

## Attribute reference
Review the output parameters that you can access after your resource is created. 

- `id` - (String) The unique identifier of the route.



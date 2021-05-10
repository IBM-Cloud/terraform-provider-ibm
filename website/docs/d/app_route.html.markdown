---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: ibm_app_route"
description: |-
  Get information about an IBM Cloud route.
---

# `ibm_app_route`

Retrieve information about an existing app route. For more information, about an app route, see [Updating your domain](https://cloud.ibm.com/docs/cloud-foundry-public?topic=cloud-foundry-public-update-domain).


## Example usage
The following example retrieves information about an app route. 


```
data "ibm_app_route" "route" {
  domain_guid = data.ibm_app_domain_shared.domain.id
  space_guid  = data.ibm_space.spacedata.id
  host        = "myhost"
  path        = "/app"
}
```


## Argument reference
Review the input parameters that you can specify for your data source. 

- `domain_guid`- (Required, String) The GUID of the domain that the route belongs to. You can retrieve the value from the `ibm_app_domain_shared` data source.
- `host` - (Optional, String)  The host name of the route. Required for shared domains.
- `path` - (Optional, String)  The path for a route. Paths must contain 2-128 characters. Paths must start with a forward slash (/). Paths must not contain a question mark (?).
- `port` - (Optional, String)  The port of the route. This value is supported for TCP router group domains only.
- `space_guid` - (Required, String) The GUID of the space that the route belongs to. You can retrieve the value from the `ibm_space` data source or by running the `ibmcloud iam space <space-name> guid` command in the IBM Cloud CLI.



## Attribute reference
Review the output parameters that you can access after you retrieved your data source. 

- `id` - (String) The unique identifier of the route.


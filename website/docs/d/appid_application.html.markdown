---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Application"
description: |-
        Retrieves AppID Application information.
---

# ibm_appid_application
Retrieve information about an IBM Cloud AppID Management Services application. For more information, see [application identity and authorization](https://cloud.ibm.com/docs/appid?topic=appid-app)

## Example usage

```terraform
data "ibm_appid_application" "app" {
    tenant_id = var.tenant_id
    client_id = var.client_id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID
- `client_id` - (Required, String) The AppID application identifier

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `discovery_endpoint` - (String) This URL returns OAuth Authorization Server Metadata
- `name` - (String) The application name
- `oauth_server_url` - (String) Base URL for common OAuth endpoints, like `/authorization`, `/token` and `/publickeys`
- `profiles_url` - (String) Base AppID API endpoint
- `type` - (String) The application type. Supported types are `regularwebapp` and `singlepageapp`.
- `secret` - (String, Sensitive) The `secret` is a secret known only to the application and the authorization server

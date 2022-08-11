---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Application"
description: |-
    Provides AppID Application resource.
---

# ibm_appid_application

Create, update, or delete an IBM Cloud AppID Management Services application resource. For more information, see [application identity and authorization](https://cloud.ibm.com/docs/appid?topic=appid-app)

## Example usage

```terraform
resource "ibm_appid_application" "app" {
  tenant_id = var.tenant_id
  name = "example application"
  type = var.application_type // singlepageapp | regularwebapp
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `name` - (Required, String) The AppID application name
- `type` - (Optional, String) The AppID application type, supported values are `singlepageapp` and `regularwebapp` (default)

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created

- `client_id` - (String) The AppID application identifier
- `discovery_endpoint` - (String) This URL returns OAuth Authorization Server Metadata
- `oauth_server_url` - (String) Base URL for common OAuth endpoints, like `/authorization`, `/token` and `/publickeys`
- `profiles_url` - (String) Base AppID API endpoint
- `secret` - (String, Sensitive) The `secret` is a secret known only to the application and the authorization server


## Import

The `ibm_appid_application` resource can be imported by using the AppID tenant ID and application client ID.

**Syntax**

```bash
$ terraform import ibm_appid_application.app <tenant_id>/<client_id>
```
**Example**

```bash
$ terraform import ibm_appid_application.app 5fa344a8-d361-4bc2-9051-58ca253f4b2b/03cd638a-b35a-43f2-a58a-c2d3fe26aaea
```

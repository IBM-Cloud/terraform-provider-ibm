---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Cloud Directory Template"
description: |-
    Provides AppID Cloud Directory Template resource.
---

# ibm_appid_cloud_directory_template

Create, update, or reset an IBM Cloud AppID Management Services Cloud Directory email templates. For more information, see [customizing emails](https://cloud.ibm.com/docs/appid?topic=appid-cd-types)

## Example usage

```terraform
resource "ibm_appid_cloud_directory_template" "tpl" {
  tenant_id = var.tenant_id
  template_name = "USER_VERIFICATION"
  subject = "Please Verify Your Email Address %%{user.displayName}" // note: `%{` has to be escaped, otherwise it will be treated as terraform template directive
  html_body = file("path/to/body.html") // no need to escape %{} within the template files
  plain_text_body = file("path/to/body.txt")
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `template_name` - (Required, String) The type of email template. This can be `USER_VERIFICATION`, `WELCOME`, `PASSWORD_CHANGED`, `RESET_PASSWORD` or `MFA_VERIFICATION`
- `language` - (Required, String) Select language for the template. Format as described at RFC5646. Default: `en`
- `subject` - (Required, String) The subject
- `html_body` - (Optional, String) The HTML body
- `plain_text_body` - (Optional, String) The text body

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created

- `base64_encoded_html_body` - (String) The HTML body of the email encoded in Base64

## Import

The `ibm_appid_cloud_directory_template` resource can be imported by using the AppID tenant ID, template name and language.

**Syntax**

```bash
$ terraform import ibm_appid_cloud_directory_template.tpl <tenant_id>/<template_name>/<language>
```
**Example**

```bash
$ terraform import ibm_appid_cloud_directory_template.tpl 5fa344a8-d361-4bc2-9051-58ca253f4b2b/WELCOME/en
```

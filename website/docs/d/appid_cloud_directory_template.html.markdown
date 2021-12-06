---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Cloud Directory Template"
description: |-
    Retrieves AppID Cloud Directory Template information.
---

# ibm_appid_cloud_directory_template
Retrieve information about an IBM Cloud AppID Management Services Cloud Directory Email Template. For more information, see [customizing emails](https://cloud.ibm.com/docs/appid?topic=appid-cd-types)

## Example usage

```terraform
data "ibm_appid_cloud_directory_template" "tpl" {
    tenant_id = var.tenant_id
    template_name = "USER_VERIFICATION"
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID
- `template_name` - (Required, String) The type of email template. This can be `USER_VERIFICATION`, `WELCOME`, `PASSWORD_CHANGED`, `RESET_PASSWORD` or `MFA_VERIFICATION`
- `language` - (Optional, String) Preferred language for resource. Format as described at RFC5646. Default: `en`

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `subject` - (String) The subject of the email
- `html_body` - (String) The HTML body of the email
- `base64_encoded_html_body` - (String) The HTML body of the email encoded in Base64
- `plain_text_body` - (String) The text body of the email

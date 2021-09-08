---
subcategory: "AppID Management"
layout: "ibm"
page_title: "IBM: AppID Password Regex"
description: |-
    Retrieves AppID Password Regex configuration.
---

# ibm_appid_password_regex
Retrieve an IBM Cloud AppID Password Regex configuration.

## Example usage

```terraform
data "ibm_appid_password_regex" "rgx" {
    tenant_id = var.tenant_id   
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `base64_encoded_regex` (String) The regex expression rule for acceptable password encoded in base64
- `error_message` (String) Custom error message
- `regex` (String) The escaped regex expression rule for acceptable password


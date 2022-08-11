---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_user_profile"
description: |-
  Manages IBM IAM user settings & profile.
---

# ibm_iam_user_profile

Retrieve information about an IAM user profile. For more information, about IAM role action, see [updating company profile details](https://cloud.ibm.com/docs/account?topic=account-contact-info).

## Example usage

```terraform
resource "ibm_iam_user_settings" "user_setting" {
  iam_id = "example@in.ibm.com"
  allowed_ip_addresses = ["192.168.0.2","192.168.0.3","192.168.0.4"]
}

data "ibm_iam_user_profile" "user_profle" {
  iam_id = ibm_iam_user_settings.user_setting.iam_id
}

```

## Argument reference

Review the argument references that you can specify for your data source.

- `iam_id` - (Required, String) The IBM ID or email address of the user.

## Attribute reference

In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `allowed_ip_addresses` (List) List of invited users IP's to access the IBM cloud console.
- `email` - (String) The email address of the user.
- `firstname`-  (String) The first name of the user.
- `ibm_id` - (String) An alphanumeric value identifying the user's IAM ID.
- `id` - (String) The unique identifier or email address of the IAM user.
- `lastname`-  (String) The last name of the user.
- `phonenumber` - (String) The contact number of the user.
- `state` - (String) The state of the user.


  

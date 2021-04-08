---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_user_profile"
description: |-
  Manages IBM IAM user settings & profle.
---

# ibm\_iam_user_profile

Can be used to fetch user profle & settings.

## Example Usage

```hcl
resource "ibm_iam_user_settings" "user_setting" {
  iam_id = "example@in.ibm.com"
  allowed_ip_addresses = ["192.168.0.2","192.168.0.3","192.168.0.4"]
}

data "ibm_iam_user_profile" "user_profle" {
  iam_id = ibm_iam_user_settings.user_setting.iam_id
}

```

## Argument Reference

The following arguments are supported:

* `iam_id` - (Required, string) The iam id or email of user.

## Attribute Reference

The following attributes are exported:

* `allowed_ip_addresses` - List of IPs from which invited user can access the IBM cloud console of the inviter.
* `id` - The unique identifier (email) of the IAM user setting.
* `firstname` - The first name of the user.
* `lastname` - The last name of the user.
* `state` - The state of the user.
* `phonenumber` - phonenumber of the user.
* `email` - The email of the user.


  
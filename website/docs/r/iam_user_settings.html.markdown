---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_user_settings"
description: |-
  Manages IBM IAM user settings.
---

# ibm_iam_user_settings

Provides a resource for IAM User Settings. The IP addresses configured here are the only ones from which a particuler user can log in to IBM Cloud. For more information, about IAM user settings, see [allowing specific IP addresses](https://cloud.ibm.com/docs/account?topic=account-ips).

## Example usage

### Configuring allowed_ip list for a particular user

```terraform
resource "ibm_iam_user_settings" "user_setting" {
  iam_id = "example@in.ibm.com"
  allowed_ip_addresses = ["192.168.0.2","192.168.0.3","192.168.0.4"]
}

```

## Argument reference
Review the argument references that you can specify for your resource. 

- `allowed_ip_addresses` - (Optional, List) Lists the IP addresses in common separated format.
- `iam_id` - (Required, String) The users IAM or Email ID.

## Attributes
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the IAM user setting as `account_id/iam_id`.

---
layout: "ibm"
page_title: "IBM : iam_user_settings"
sidebar_current: "docs-ibm-resource-iam-user-settings"
description: |-
  Manages IBM IAM User Settings.
---

# ibm\_iam_user_settings

Provides a resource for IAM User Settings. The IP addresses configured here are the only ones from which a particuler user can log in to IBM Cloud.

## Example Usage

### Configuraing allowed_ip list for a particular user

```hcl
resource "ibm_iam_user_settings" "user_setting" {
  iam_id = "example@in.ibm.com"
  allowed_ip_addresses = ["192.168.0.2","192.168.0.3","192.168.0.4"]
}

```

## Argument Reference

The following arguments are supported:

* `iam_id` - (Required, string) The user's IAM ID or email ID. 
* `allowed_ip_addresses` - (Optional, list) comma seperated list of IP addresses.

## Attributes

The following attributes are exported:

* `id` - The unique identifier of the IAM user setting.The id is composed of \<account_id\>/\<iam_id\>.

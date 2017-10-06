---
layout: "ibm"
page_title: "IBM : iam_user_policy"
sidebar_current: "docs-ibm-resource-iam-user-policy"
description: |-
  Manages IBM IAM User Policy.
---

# ibm\_iam_policy

Provides a resource for creating and assigning an IAM Access policy to a user. To assign a policy to one user, the user must exist in the account to which you assign the policy. To assign a policy to all of the services under an account, the value for `service name` must be `All Identity and Access enabled services`. To assign a policy to all of the service instances under a service, you must leave the value blank for `service instance`.

**Note**: Currently only one service name and service instance supported.

## Example Usage

```hcl
data "ibm_org" "ds_org" {
  org = "sample"
}

data "ibm_account" "ds_acc" {
  org_guid = "${data.ibm_org.ds_org.id}"
}

resource "ibm_access_policies" "iam_policy" {
  account_guid = "${data.ibm_account.ds_acc.id}"
  ibm_id       = "user@example.com"
  roles        = ["viewer"]
  resources    = [{
    "service_name"     = "sample-service",
    "service_instance" = ["1refjnjb-vr4-vverr"]
  }]
}
```

## Argument Reference

The following arguments are supported:

* `account_guid` - (Required, string) The GUID of the account to which you want to assign the policy. To assign a policy to one user, the user must exist in this account. You can retrieve the value from the `ibm_account` data source or by running the `bx iam accounts` command in the [Bluemix CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `ibm_id` - (Required, string) The IBMid of the user to whom you want to assign the policy.
* `roles` - (Required, array) The IAM Roles you want to assign. Valid values for roles are _viewer_, _editor_, _operator_, and _administrator_. At least one role is required.
* `resources` - (Required, array) A nested block describing the IAM resources to which you want to assign the policy. At least one resource is required.
  * `service_name` - (Required, string) The name of the service. To assign a policy to all of the services under an account, the value must be `All Identity and Access enabled services`.
  * `service_instance` - (Optional, string) The service instance. To assign a policy to all of the service instances under a service, you must leave the value blank.
  * `region` - (Optional, string) The region to which the service belongs.
  * `resource_type` - (Optional, string) The resource type.
  * `resource` - (Optional, string) The resource.
  * `space_guid` - (Optional, string) The GUID of the Bluemix space where the service is deployed. You can retrieve the value with the `ibm_space` data source or by running the `bx iam space <space_name> --guid` command in the Bluemix CLI.
  * `organization_guid` - (Optional, string) The GUID of the Bluemix org. You can retrieve the value from the `ibm_org` data source or by running the `bx iam orgs --guid` command in the [Bluemix CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).

* `tags` - (Optional, array of strings) Tags associated with the IAM user policy instance.
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the policy.
* `etag` - The revision number for updating an object.

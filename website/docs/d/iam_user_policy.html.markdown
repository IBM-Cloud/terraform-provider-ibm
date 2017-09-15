---
layout: "ibm"
page_title: "IBM : iam_user_policy"
sidebar_current: "docs-ibm-datasource-iam-user-policy"
description: |-
  Manages IBM IAM User Policy.
---

# ibm\_iam_user_policy

Import the details of an IAM (Identity and Access Management) user policy on IBM Bluemix as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_org" "ds_org" {
  org = "sample"
}

data "ibm_account" "ds_acc" {
  org_guid = "${data.ibm_org.ds_org.id}"
}

resource "ibm_iam_user_policy" "iam_policy" {
  account_guid = "${data.ibm_account.ds_acc.id}"
  ibm_id       = "user@example.com"
  roles        = ["viewer"]
  resources    = [{"service_name" = "sample-service", "service_instance"=["1refjnjb-vr4-vverr"]}]
}

data "ibm_iam_user_policy" "testacc_iam_policies" {
  account_guid = "${data.ibm_account.ds_acc.id}"
  ibm_id = "${ibm_iam_user_policy.iam_policy.ibm_id}"
}

```

## Argument Reference

The following arguments are supported:

* `account_guid` - (Required, string) * `account_guid` - (Required, string) The GUID for the Bluemix account. You can retrieve the value from the `ibm_account` data source or by running the `bx iam accounts` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).
* `ibm_id` - (Required, string) The IBM ID of the user to whom you want to assign the policy.

## Attribute Reference

The following attributes are exported:

* `policies` - Nested block describing IAM Policies assigned to user in the account.

Nested `policies` blocks have the following structure:

* `id` - IAM Policy ID
* `roles` -  A nested block describing the roles assigned to the policy.
  * `name` - The IAM role assigned to the policy.
* `resources` -  A nested block describing the IAM resources in the policy.
  * `service_name` - The name of the service.
  * `service_instance` - The service instance.
  * `region` - The region to which the service belongs.
  * `resource_type` - The resource type.
  * `resource` - The resource.
  * `space_guid` - The GUID of the Bluemix space.
  * `organization_guid` - The GUID of the Bluemix organization.

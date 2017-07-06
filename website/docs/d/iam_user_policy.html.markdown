---
layout: "ibm"
page_title: "IBM : iam_user_policy"
sidebar_current: "docs-ibm-datasource-iam-user-policy"
description: |-
  Manages IBM IAM User Policy.
---

# ibm\_iam_user_policy


 
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

* `account_guid` - (Required, string) The Guid of the account.The value can be retrieved from the `ibm_account` data source, or by running the `bx iam accounts` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).
* `ibm_id` - (Required, string) The IBM ID of the user whom to assign the policy

## Attributes Reference

The following attributes are exported:

* `policies` - Nested block describing IAM Policies assigned to user in the account

Nested `policies` blocks have the following structure:

* `id` - IAM Policy ID
* `roles` -  Nested block describing the roles assigned to the policy.
* `resources` -  Nested block describing the IAM resources in the policy.

Nested `roles` blocks have the following structure:

* `name` - The IAM Role assigned to policy.
 
Nested `resources` blocks have the following structure:

* `service_name` - Name of the service
* `service_instance` - Service instance 
* `region` - The region to which the service belongs
* `resource_type` - Resource type
* `resource` - Resource 
* `space_guid` - The GUID of the Bluemix space. 
* `organization_guid` - The GUID of the Bluemix org.

    
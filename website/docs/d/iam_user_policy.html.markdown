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
  resources = [{"sample-service"}]
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

* `policies` - IAM Policies
* `policies.id` - IAM Policy ID
* `policies.roles` -  The roles assigned to the policy.
* `policies.roles.id` - The IAM Role assigned to policy.
* `policies.resources` -  Represents IAM resources. 
* `policies.resources.service_name` - Name of the service
* `policies.resources.service_instance` - Service instance 
* `policies.resources.region` - The region to which the service belongs
* `policies.resources.resource_type` - Resource type
* `policies.resources.resource` - Resource 
* `policies.resources.space_guid` - The GUID of the Bluemix space. 
* `policies.resources.organization_guid` - The GUID of the Bluemix org.

    
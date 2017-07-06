---
layout: "ibm"
page_title: "IBM : iam_user_policy"
sidebar_current: "docs-ibm-resource-iam-user-policy"
description: |-
  Manages IBM IAM User Policy.
---

# ibm\_iam_policy

The resource IAM Access policy creates and assigns a policy to user. The user should exist in the account to whom the 
policy is assigned. If IAM User Policy to be assigned to all the services under an account then the service name to be passed
is "All Identity and Access enabled services". If IAM Policy to be assigned to all the service instances under a service 
then the service instance should be passed as blank. 

Note: Currently only one service name and service instance supported.
 
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
  resources    = [{"service_name" = "sample-service", "service_instance"=["1refjnjb-vr4-vverr"]}]
}


```

## Argument Reference

The following arguments are supported:

* `account_guid` - (Required, string) The Guid of the account.The value can be retrieved from the `ibm_account` data source, or by running the `bx iam accounts` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).
* `ibm_id` - (Required, string) The IBM ID of the user whom to assign the policy
* `roles` - (Required, array) Represents IAM Roles. Valid values for roles are _viewer_, _editor_, _operator_ and _administrator_. At least one role is required
* `resources` - (Required,array) Nested block describing an IAM resources to which the policy be created. At least one resource is required

Nested `resources` blocks have the following structure:

* `service_name` - (Required, string) Name of the service
* `service_instance` - (Optional, string) Service instance 
* `region` - (Optional, string) The region to which the service belongs
* `resource_type` - (Optional, string) Resource type
* `resource` - (Optional, string) Resource 
* `space_guid` - (Optional, string) The GUID of the Bluemix space where the service is deployed. The value can be retrieved with the `ibm_space` data source, or by running the `bx iam space <space_name> --guid` command in the Bluemix CLI. 
* `organization_guid` - (Optional, string) The GUID of the Bluemix org. The value can be retrieved from the `ibm_org` data source, or by running the `bx iam orgs --guid` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).

## Attributes Reference

The following attributes are exported:

* `id` - The id of policy created.
* `etag` - The revision number for updating an object
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
is "All Identity and Access enbled services". If IAM Policy to be assigned to all the service instances under a service 
then the service instance to be passed as blank. 

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
  resources    = [{"service_name" = "sample-service", "service_instance"="1refjnjb-vr4-vverr"}]
}


```

## Argument Reference

The following arguments are supported:

* `account_guid` - (Required, string) The Guid of the account.The value can be retrieved from the `ibm_account` data source, or by running the `bx iam accounts` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).
* `ibm_id` - (Required, string) The IBM ID of the user whom to assign the policy
* `resources` - (Required,array) Represents IAM resources. Atleast one resource is required
* `resources.service_name` - (Required, string) Name of the service
* `resources.service_instance` - (Optional, string) Service instance 
* `resources.region` - (Optional, string) The region to which the service belongs
* `resources.resource_type` - (Optional, string) Resource type
* `resources.resource` - (Optional, string) Resource 
* `resources.space_guid` - (Optional, string) The GUID of the Bluemix space where the service is deployed. The value can be retrieved with the `ibm_space` data source, or by running the `bx iam space <space_name> --guid` command in the Bluemix CLI. 
* `resources.organization_guid` - (Optional, string) The GUID of the Bluemix org. The value can be retrieved from the `ibm_org` data source, or by running the `bx iam orgs --guid` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).
* `roles` - (Required, array) Represents IAM Roles. Atleast one role is required 
* `roles.id` - (Required,string) - The IAM Role to be assigned.Valid roles are viewer,Operator,Editor,Administrator
## Attributes Reference

The following attributes are exported:

* `id` - The id of policy created.
* `etag` - The revision number for updating an object
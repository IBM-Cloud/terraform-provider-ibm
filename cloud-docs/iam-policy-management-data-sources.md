---

copyright:
  years: 2021
lastupdated: "2021"

keywords: terraform

subcollection: terraform

---

# IAM Policy Management API data sources
{: #iam-policy-management-data-sources}

Review the data sources that you can use to retrieve information about your IAM Policy Management API resources.
All data sources are imported as read-only information. You can reference the output parameters for each data source by using Terraform interpolation syntax.

Before you start working with your data source, make sure to review the [required parameters](/docs/terraform?topic=terraform-provider-reference#required-parameters)
that you need to specify in the `provider` block of your Terraform configuration file.
{: important}

## `ibm_iam_policy`
{: #iam_policy}

Retrieve information about iam_policy.
{: shortdesc}

### Sample Terraform code
{: #iam_policy-sample}

```
data "ibm_iam_policy" "iam_policy" {
  policy_id = "policy_id"
}
```

### Input parameters
{: #iam_policy-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`policy_id`|String|Required|The policy ID.|

### Output parameters
{: #iam_policy-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The policy ID.|
|`type`|String|The policy type; either 'access' or 'authorization'.|
|`description`|String|Customer-defined description.|
|`subjects`|List|The subjects associated with a policy.|
|`subjects.attributes`|List|List of subject attributes.|
|`subjects.attributes.name`|String|The name of an attribute.|
|`subjects.attributes.value`|String|The value of an attribute.|
|`roles`|List|A set of role cloud resource names (CRNs) granted by the policy.|
|`roles.role_id`|String|The role cloud resource name granted by the policy.|
|`roles.display_name`|String|The display name of the role.|
|`roles.description`|String|The description of the role.|
|`resources`|List|The resources associated with a policy.|
|`resources.attributes`|List|List of resource attributes.|
|`resources.attributes.name`|String|The name of an attribute.|
|`resources.attributes.value`|String|The value of an attribute.|
|`resources.attributes.operator`|String|The operator of an attribute.|
|`href`|String|The href link back to the policy.|
|`created_at`|String|The UTC timestamp when the policy was created.|
|`created_by_id`|String|The iam ID of the entity that created the policy.|
|`last_modified_at`|String|The UTC timestamp when the policy was last modified.|
|`last_modified_by_id`|String|The iam ID of the entity that last modified the policy.|

## `ibm_iam_custom_role`
{: #iam_custom_role}

Retrieve information about iam_custom_role.
{: shortdesc}

### Sample Terraform code
{: #iam_custom_role-sample}

```
data "ibm_iam_custom_role" "iam_custom_role" {
  role_id = "role_id"
}
```

### Input parameters
{: #iam_custom_role-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`role_id`|String|Required|The role ID.|

### Output parameters
{: #iam_custom_role-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The role ID.|
|`display_name`|String|The display name of the role that is shown in the console.|
|`description`|String|The description of the role.|
|`actions`|List|The actions of the role.|
|`crn`|String|The role CRN.|
|`name`|String|The name of the role that is used in the CRN. Can only be alphanumeric and has to be capitalized.|
|`account_id`|String|The account GUID.|
|`service_name`|String|The service name.|
|`created_at`|String|The UTC timestamp when the role was created.|
|`created_by_id`|String|The iam ID of the entity that created the role.|
|`last_modified_at`|String|The UTC timestamp when the role was last modified.|
|`last_modified_by_id`|String|The iam ID of the entity that last modified the policy.|
|`href`|String|The href link back to the role.|


---

copyright:
  years: 2021
lastupdated: "2021"

keywords: terraform

subcollection: terraform

---

# IAM Policy Management API resources
{: #iam-policy-management-resources}

Create, update, or delete IAM Policy Management API resources.
You can reference the output parameters for each resource in other resources or data sources by using Terraform interpolation syntax.

Before you start working with your resource, make sure to review the [required parameters](/docs/terraform?topic=terraform-provider-reference#required-parameters) 
that you need to specify in the `provider` block of your Terraform configuration file.
{: important}

## `ibm_iam_policy`
{: #iam_policy}

Create, update, or delete an iam_policy.
{: shortdesc}

### Sample Terraform code
{: #iam_policy-sample}

```
resource "ibm_iam_policy" "iam_policy" {
  type = "type"
  subjects = { example: "object" }
  roles = { example: "object" }
  resources = { example: "object" }
  description = "placeholder"
  accept_language = "placeholder"
}
```

### Input parameters
{: #iam_policy-input}

Review the input parameters that you can specify for your resource. {: shortdesc}

|Name|Data type|Required/optional|Description|Forces new resource|
|----|-----------|-------|----------|--------------------|
|`type`|String|Required|The policy type; either 'access' or 'authorization'.|No|
|`subjects`|List|Required|The subjects associated with a policy.|No|
|`roles`|List|Required|A set of role cloud resource names (CRNs) granted by the policy.|No|
|`resources`|List|Required|The resources associated with a policy.|No|
|`description`|String|Optional|Customer-defined description.|No|
|`accept_language`|String|Optional|Translation language code.|No|

### Output parameters
{: #iam_policy-output}

Review the output parameters that you can access after your resource is created. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The unique identifier of the iam_policy.|
|`href`|String|The href link back to the policy.|
|`created_at`|String|The UTC timestamp when the policy was created.|
|`created_by_id`|String|The iam ID of the entity that created the policy.|
|`last_modified_at`|String|The UTC timestamp when the policy was last modified.|
|`last_modified_by_id`|String|The iam ID of the entity that last modified the policy.|

### Import
{: #iam_policy-import}

`ibm_iam_policy` can be imported by ID

```
$ terraform import ibm_iam_policy.example sample-id
```

## `ibm_iam_custom_role`
{: #iam_custom_role}

Create, update, or delete an iam_custom_role.
{: shortdesc}

### Sample Terraform code
{: #iam_custom_role-sample}

```
resource "ibm_iam_custom_role" "iam_custom_role" {
  display_name = "display_name"
  actions = [ "actions" ]
  name = "name"
  account_id = "account_id"
  service_name = "service_name"
  description = "placeholder"
  accept_language = "placeholder"
}
```

### Input parameters
{: #iam_custom_role-input}

Review the input parameters that you can specify for your resource. {: shortdesc}

|Name|Data type|Required/optional|Description|Forces new resource|
|----|-----------|-------|----------|--------------------|
|`display_name`|String|Required|The display name of the role that is shown in the console.|No|
|`actions`|List|Required|The actions of the role.|No|
|`name`|String|Required|The name of the role that is used in the CRN. Can only be alphanumeric and has to be capitalized.|No|
|`account_id`|String|Required|The account GUID.|No|
|`service_name`|String|Required|The service name.|No|
|`description`|String|Optional|The description of the role.|No|
|`accept_language`|String|Optional|Translation language code.|No|

### Output parameters
{: #iam_custom_role-output}

Review the output parameters that you can access after your resource is created. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The unique identifier of the iam_custom_role.|
|`crn`|String|The role CRN.|
|`created_at`|String|The UTC timestamp when the role was created.|
|`created_by_id`|String|The iam ID of the entity that created the role.|
|`last_modified_at`|String|The UTC timestamp when the role was last modified.|
|`last_modified_by_id`|String|The iam ID of the entity that last modified the policy.|
|`href`|String|The href link back to the role.|

### Import
{: #iam_custom_role-import}

`ibm_iam_custom_role` can be imported by ID

```
$ terraform import ibm_iam_custom_role.example sample-id
```


---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_user_policy"
description: |-
  Manages IBM IAM user policy.
---

# ibm_iam_user_policy

Create, update, or delete an IAM user policy. To assign a policy to one user, the user must exist in the account to which you assign the policy. For more information, about IAM role action, see [managing access to resources](https://cloud.ibm.com/docs/account?topic=account-assign-access-resources).

## Example usage

### User policy for all Identity and Access enabled services 

```terraform
resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles  = ["Viewer"]
  description = "IAM User Policy"
  
  resource_tags {
    name = "env"
    value = "dev"
  }
  
}

```

### User policy using service with region

```terraform
resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles  = ["Viewer", "Manager"]

  resources {
    service = "cloudantnosqldb"
    region  = "us-south"
  }
}

```
### User policy using resource instance 

```terraform
resource "ibm_resource_instance" "instance" {
  name     = "test"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}

resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles  = ["Manager", "Viewer", "Administrator"]

  resources {
    service              = "kms"
    resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
  }
}

```

### User policy using resource group 

```terraform
data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles  = ["Viewer"]

  resources {
    service           = "containers-kubernetes"
    resource_group_id = data.ibm_resource_group.group.id
  }
}

```

### User policy using resource and resource type 

```terraform
data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles  = ["Administrator"]

  resources {
    resource_type = "resource-group"
    resource      = data.ibm_resource_group.group.id
  }
}

```

### User policy using attributes 

```terraform
data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles  = ["Administrator"]

  resources {
    service = "is"

    attributes = {
      "vpcId" = "*"
    }
  }
}

```

### User policy using resource_attributes

```terraform
resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles           = ["Viewer"]
  resource_attributes {
    name  = "resource"
    value = "test123*"
    operator = "stringMatch"
  }
  resource_attributes {
    name  = "serviceName"
    value = "messagehub"
  }
}
```

### User policy using service_type with region

```terraform
resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles  = ["Viewer"]

  resources {
    service_type = "service"
    region = "us-south"
  }
}

```

### User policy by using service and rule_conditions
`rule_conditions` can be used in conjunction with `pattern` and `rule_operator` to implement user policies with time-based conditions. For information see [Limiting access with time-based conditions](https://cloud.ibm.com/docs/account?topic=account-iam-time-based&interface=ui). **Note** Currently, a policy resource created without `rule_conditions`, `pattern`, and `rule_operator` cannot be updated including those conditions on update.

```terraform
resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles      = ["Viewer"]
  resources {
    service = "kms"
  }
  rule_conditions {
    key = "{{environment.attributes.day_of_week}}"
    operator = "dayOfWeekAnyOf"
    value = ["1+00:00","2+00:00","3+00:00","4+00:00"]
  }
  rule_conditions {
    key = "{{environment.attributes.current_time}}"
    operator = "timeGreaterThanOrEquals"
    value = ["09:00:00+00:00"]
  }
  rule_conditions {
    key = "{{environment.attributes.current_time}}"
    operator = "timeLessThanOrEquals"
    value = ["17:00:00+00:00"]
  }
  rule_operator = "and"
  pattern = "time-based-conditions:weekly:custom-hours"
}
```

### User policy using service_group_id resource attribute

```terraform
resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles  = ["Service ID creator", "User API key creator", "Administrator"]

  resource_attributes {
    name     = "service_group_id"
    operator = "stringEquals"
    value    = "IAM"
  }
}
```

### User Policy by using Attribute Based Condition
`rule_conditions` can be used in conjunction with `pattern = attribute-based-condition:resource:literal-and-wildcard` and `rule_operator` to implement more complex policy conditions. **Note** Currently, a policy resource created without `rule_conditions`, `pattern`, and `rule_operator` cannot be updated including those conditions on update.

```terraform
resource "ibm_iam_user_policy" "policy" {
  ibm_id = "test@in.ibm.com"
  roles  = ["Writer"]
  resource_attributes {
    value = "cloud-object-storage"
    operator = "stringEquals"
    name = "serviceName"
  }
  resource_attributes {
    value = "cos-instance"
    operator = "stringEquals"
    name = "serviceInstance"
  }
  resource_attributes {
    value = "bucket"
    operator = "stringEquals"
    name = "resourceType"
  }
  resource_attributes {
    value = "fgac-tf-test"
    operator = "stringEquals"
    name = "resource"
  }
  rule_conditions {
    operator = "and"
    conditions {
      key = "{{resource.attributes.prefix}}"
      operator = "stringMatch"
      value = ["folder1/subfolder1/*"]
    }
    conditions {
      key = "{{resource.attributes.delimiter}}"
      operator = "stringEqualsAnyOf"
      value = ["/",""]
    }
  }
  rule_conditions {
    key = "{{resource.attributes.path}}"
    operator = "stringMatch"
    value = ["folder1/subfolder1/*"]
  }
  rule_conditions {
    operator = "and"
    conditions {
      key = "{{resource.attributes.delimiter}}"
      operator = "stringExists"
      value = ["false"]
    }
    conditions {
      key = "{{resource.attributes.prefix}}"
      operator = "stringExists"
      value = ["false"]
    }
  }
  rule_operator = "or"
  pattern = "attribute-based-condition:resource:literal-and-wildcard"
  description = "IAM User Policy Attribute Based Condition Creation for test scenario"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `account_management` - (Optional, Bool) Gives access to all account management services if set to **true**. Default value **false**. If you set this option, do not set `resources` at the same time. **Note** Conflicts with `resources` and `resource_attributes`.
- `description`  (Optional, String) The description of the IAM User Policy.
- `ibm_id` - (Required, Forces new resource, String) The IBM ID or Email address of the user.
- `roles` - (Required, List)  A comma separated list of roles. Valid roles are `Writer`, `Reader`, `Manager`, `Administrator`, `Operator`, `Viewer`, and `Editor`. For more information, about supported service specific roles, see  [IAM roles and actions](https://cloud.ibm.com/docs/account?topic=account-iam-service-roles-actions)
- `resources` - (Optional, List) A nested block describes the resource of this policy. **Note** Conflicts with `account_management` and `resource_attributes`.

  Nested scheme for `resources`:
  - `attributes` (Optional, Map)  A set of resource attributes in the format `name=value,name=value`. If you set this option, do not specify `account_management`  and `resource_attributes` at the same time.
  - `resource_instance_id` - (Optional, String) The ID of the resource instance of the policy definition.
  - `region`  (Optional, String) The region of the policy definition.
  - `resource_type` - (Optional, String) The resource type of the policy definition.
  - `resource` - (Optional, String) The resource of the policy definition.
  - `resource_group_id` - (Optional, String) The ID of the resource group. To retrieve the value, run `ibmcloud resource groups` or use the `ibm_resource_group` data source.
  - `service` - (Optional, String) The service name of the policy definition. You can retrieve the value by running the `ibmcloud catalog service-marketplace` or `ibmcloud catalog search` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started). Attributes service, service_type are mutually exclusive.
  - `service_type`  (Optional, String) The service type of the policy definition. **Note** Attributes service, service_type are mutually exclusive.
  - `service_group_id` (Optional, String) The service group id of the policy definition. **Note** Attributes service, service_group_id are mutually exclusive.
- `resource_attributes` - (Optional, List) A nested block describing the resource of this policy. - `resource_attributes` - (Optional, List) A nested block describing the resource of this policy. **Note** Conflicts with `account_management` and `resources`.
  
  Nested scheme for `resource_attributes`:
  - `name` - (Required, String) The name of an Attribute. Supported values are `serviceName`, `serviceInstance`, `region`,`resourceType`, `resource`, `resourceGroupId`, `service_group_id` and other service specific resource attributes.
  - `value` - (Required, String) The value of an attribute.
  - `operator` - (Optional, String) Operator of an attribute. The default value is `stringEquals`. **Note**: Conflicts with `account_management` and `resources`.

- `resource_tags`  (Optional, List)  A nested block describing the access management tags.  **Note** `resource_tags` are only allowed in policy with resource attribute serviceType, where value is equal to service.

  Nested scheme for `resource_tags`:
  - `name` - (Required, String) The key of an access management tag. 
  - `value` - (Required, String) The value of an access management tag.
  - `operator` - (Optional, String) Operator of an attribute. The default value is `stringEquals`.

- `transaction_id`- (Optional, String) The TransactionID can be passed to your request for tracking the calls.

- `rule_conditions` - (Optional, List) A nested block describing the rule conditions of this policy.

  Nested schema for `rule_conditions`:
  - `key` - (Optional, String) The key of a rule condition.
  - `operator` - (Required, String) The operator of a rule condition.
  - `value` - (Optional, List) The value of a rule condition.
  - `conditions` - (Optional, List) A nested block describing additional conditions of this policy.

     Nested schema for `conditions`:
      - `key` - (Required, String) The key of a condition.
      - `operator` - (Required, String) The operator of a condition.
      - `value` - (Required, List) The value of a condition.

- `rule_operator` - (Optional, String) The operator used to evaluate multiple rule conditions, e.g., all must be satisfied with `and`.

- `pattern` - (Optional, String) The pattern that the rule follows, e.g., `time-based-conditions:weekly:all-day`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`  - (String) The unique identifier of the user policy. The ID is composed of `<ibm_id>/<user_policy_id>`.
- `version` - (String) The version of the user policy.


## Import
The user policy can be imported by using the IBMID and user policy ID.

**Syntax**

```
$ terraform import ibm_iam_user_policy.example <ibm_id>/<user_policy_ID>
```

**Example**

```
$ terraform import ibm_iam_user_policy.example test@in.ibm.com/9ebf7018-3d0c-4965-9976-ef8e0c38a7e2
```

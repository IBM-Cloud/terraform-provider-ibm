---
layout: "ibm"
page_title: "IBM : ibm_scc_rule"
description: |-
  Manages scc_rule.
subcategory: "Security and Compliance Center"
---

# ibm_scc_rule

Provides a resource for scc_rule. This allows scc_rule to be created, updated and deleted. For more information about Security and Compliance Center rules, see [Defining Rules](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-rules-define&interface=ui).

## Example Usage

```hcl
resource "ibm_scc_rule" "scc_rule_tf_example" {
  account_id  = "thisIsAFake32CharacterAccountID"
  name        = "Terraform rule"
  description = "Cloud Object Storage buckets can only be created in us-south."
  labels      = ["example"]
  target {
    service_name  = "cloud-object-storage"
    resource_kind = "bucket"
  }
  required_config {
    // example of a Cloud Object Storage configuration
    description = "Cloud Object Storage buckets can only be created in us-south."
    property    = "location"
    operator    = "string_equals"
    value       = "us-south"
  }
  enforcement_actions {
    action = "disallow"
  }
}
```

In the above example, COS buckets must have `location` set to `us-south` to be compliant.

## Argument Reference

Review the argument reference that you can specify for your resource.

* `account_id` - (Required, String) Your IBM Cloud account ID, or the account ID that you want to target.
* `name` - (Required, String) A human-readable alias to assign to your rule.
    * Constraints: The maximum length is `32` characters. The minimum length is `1` character.
* `description` - (Required, String) An extended description of your rule.
    * Constraints: The maximum length is `256` characters. The minimum length is `1` character.
* `labels` - (Optional, List) Labels that you can use to group and search for similar rules, such as those that help you to meet a specific organization guideline.
    * Constraints: The maximum length is `32` items.
* `enforcement_actions` - (Optional, List) The actions that the service must run on your behalf when a request to create or modify the target resource does not comply with your conditions.
    * Constraints: The maximum length is `1` items.
Nested scheme for **enforcement_actions**:
    * `action` - (Required, String) To block a request from completing, use `disallow`.
        * Constraints: Allowable values are: `disallow`.
* `target` - (Required, List) The properties that describe the resource that you want the rule or template to target.
  Nested scheme for **target**:
    * `additional_target_attributes` - (Optional, List) An extra qualifier for the resource kind. When you include additional attributes, only the resources that match the definition are included in the rule or template.
      Nested scheme for **additional_target_attributes**:
        * `name` - (Required, String) The name of the additional attribute that you want to use to further qualify the target. Options differ depending on the service or resource that you are targeting with a rule or template. For more information, refer to the service documentation.
        * `operator` - (Required, String) The way in which the `name` field is compared to its value.There are three types of operators: string, numeric, and boolean.
            * Constraints: Allowable values are:
                * `string_equals`
                * `string_not_equals`
                * `string_match`
                * `string_not_match`
                * `num_equals`
                * `num_not_equals`
                * `num_less_than`
                * `num_less_than_equals`
                * `num_greater_than`
                * `num_greater_than_equals`
                * `is_empty`
                * `is_not_empty`
                * `is_true`
                * `is_false`
        * `value` - (Optional, String) The value that you want to apply to `name` field. Options differ depending on the rule or template that you configure. For more information, refer to the service documentation.
    * `resource_kind` - (Required, String) The type of resource that you want to target.
    * `service_name` - (Required, String) The programmatic name of the IBM Cloud service that you want to target with the rule or template.
        * Constraints: The value must match regular expression `/^[a-z-]*$/`.
* `required_config` - (Required, List)
Nested scheme for **required_config**:
    * `description` - (Optional, String)
    One of the following:
    1. `rule_condition`:
        ~> **NOTE**: Currently the `ips_in_range` and `strings_in_list` cannot be used due to a limitation of the scc-go-sdk
        * `operator` - (Required, String) The way in which the `property` field is compared to its value. To learn more, see the [docs](/docs/security-compliance?topic=security-compliance-what-is-rule#rule-operators).
            * Constraints: Allowable values are:
                * `is_true`
                * `is_false`
                * `is_empty`
                * `is_not_empty`
                * `string_equals`
                * `string_not_equals`
                * `string_match`
                * `string_not_match`
                * `num_equals`
                * `num_not_equals`
                * `num_less_than`
                * `num_less_than_equals`
                * `num_greater_than`
                * `num_greater_than_equals`
        * `property` - (Required, String) A resource configuration variable that describes the property that you want to apply to the target resource.Available options depend on the target service and resource.
        * `value` - (Optional, String) The way in which you want your property to be applied. Value options differ depending on the rule that you configure. If you use a boolean operator, you do not need to input a value.

        example schema for using `rule_condition`:
        ```terraform
        required_config {
            description	= "test config"
            property	= "location"
            operator	= "string_not_equals"
            value	= "eu-de"
        }
        ```
        The above example details a `required_config` that has a single rule_condition
    2. `and/or` - (Optional, List) A list of `rule_condition` that should be set for the rule. If `and` is being used, it means that every `rule_condition` in the list must be true. If `or` is being used, it means that at least one `rule_condition` in the list needs to be true.

        ~> **NOTE**: The required_config must have only one of following: `and`, `or`, or `rule_condtion`. These values cannot be mixed with each other at the same depth (i.e. 'or' and 'and' cannot be defined at the same level/depth)

        example schema for using `and/or`:

        <table>
        <tr>
        <td> Terraform </td> <td> JSON </td>
        </tr>
        <tr>
        <td>

        ```hcl
        required_config {
            description = "test config"
            and {		// rule_condition[0]
                property = "storage_class"
                operator = "string_equals"
                value    = "smart"
            }
            and {		// rule_condition[1]
                property = "location"
                operator = "string_equals"
                value    = "us-south"
            } 
        }
        ``` 

        </td>
        <td>

        ```json
        required_config: {
            "description": "test config",
            "and": [
                {
                    "property": "storage_class",
                    "operator": "string_equals",
                    "value": "smart"
                },
                {
                    "property": "location",
                    "operator": "string_equals",
                    "value": "us-south"
                }
            ]
        }
        ```

        </td>
        </tr>
        </table>
        The above example details a `required_config` that has two `rule_condition`s and it is equivalent to: 
        ```
        rule_condtion[0] && rule_condition[1]
        ```

        Replace both `and` with `or` in the example above if you want the following logic:
        ```
        rule_condition[0] || rule_condition[1]
        ```

        Users can also create nested rules (with a maximum depth of 2 levels).
        Example (with a depth of 2):
        ```hcl
        required_config {
            and {
                // A
            }
            and {
                or {
                    // B
                }
                or {
                    // C
                }
            }
        ```
        The above example is equivalent to: `A && (B || C)`


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the ibm_scc_rule.
* `created_by` - (Optional, String) The unique identifier for the user or application that created the resource.
* `creation_date` - (Optional, String) The date the resource was created.
* `enforcement_actions` - (Required, List) The actions that the service must run on your behalf when a request to create or modify the target resource does not comply with your conditions.
  * Constraints: The maximum length is `1` items.
Nested scheme for **enforcement_actions**:
	* `action` - (Required, String) To block a request from completing, use `disallow`.
	  * Constraints: Allowable values are: `disallow`.
* `modification_date` - (Optional, String) The date the resource was last modified.
* `modified_by` - (Optional, String) The unique identifier for the user or application that last modified the resource.
* `version` - Version of the ibm_scc_rule.
* `rule_type` - (Optional, String) The type of rule. Rules that you create are `user_defined`.
  * Constraints: Allowable values are: `user_defined`.

## Import

You can import the `ibm_scc_rule` resource by using `rule_id`. The UUID that uniquely identifies the rule.

# Syntax
```
$ terraform import ibm_scc_rule.scc_rule <rule_id>
```

# Example
```
$ terraform import ibm_scc_rule.scc_rule rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf
```

---

copyright:
  years: 2021
lastupdated: "2021"

keywords: terraform

subcollection: terraform

---

# Context Based Restrictions resources
{: #context-based-restrictions-resources}

Create, update, or delete Context Based Restrictions resources.
You can reference the output parameters for each resource in other resources or data sources by using Terraform interpolation syntax.

Before you start working with your resource, make sure to review the [required parameters](/docs/terraform?topic=terraform-provider-reference#required-parameters) 
that you need to specify in the `provider` block of your Terraform configuration file.
{: important}

## `ibm_cbr_zone`
{: #cbr_zone}

Create, update, or delete an cbr_zone.
{: shortdesc}

### Sample Terraform code
{: #cbr_zone-sample}

```
resource "ibm_cbr_zone" "cbr_zone" {
  account_id = "12ab34cd56ef78ab90cd12ef34ab56cd"
  description = "this is an example of zone"
  name = "an example of zone"
}
```

### Input parameters
{: #cbr_zone-input}

Review the input parameters that you can specify for your resource. {: shortdesc}

|Name|Data type|Required/optional|Description|Forces new resource|
|----|-----------|-------|----------|--------------------|
|`account_id`|String|Optional|The id of the account owning this zone. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.|No|
|`addresses`|List|Optional|The list of addresses in the zone. The maximum length is `1000` items. The minimum length is `1` item.|No|
|`description`|String|Optional|The description of the zone. The maximum length is `300` characters. The minimum length is `0` characters. The value must match regular expression `/^[\\x20-\\xFE]*$/`.|No|
|`excluded`|List|Optional|The list of excluded addresses in the zone. The maximum length is `1000` items.|No|
|`name`|String|Optional|The name of the zone. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 \\-_]+$/`.|No|
|`transaction_id`|String|Optional|The UUID that is used to correlate and track transactions. If you omit this field, the service generates and sends a transaction ID in the response.**Note:** To help with debugging, we strongly recommend that you generate and supply a `Transaction-Id` with each request. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.|No|

### Output parameters
{: #cbr_zone-output}

Review the output parameters that you can access after your resource is created. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The unique identifier of the cbr_zone.|
|`address_count`|Integer|The number of addresses in the zone.|
|`created_at`|String|The time the resource was created.|
|`created_by_id`|String|IAM ID of the user or service which created the resource.|
|`crn`|String|The zone CRN.|
|`excluded_count`|Integer|The number of excluded addresses in the zone.|
|`href`|String|The href link to the resource.|
|`last_modified_at`|String|The last time the resource was modified.|
|`last_modified_by_id`|String|IAM ID of the user or service which modified the resource.|
|`version`|String|The version of the cbr_zone.|

### Import
{: #cbr_zone-import}

`ibm_cbr_zone` can be imported by ID

```
$ terraform import ibm_cbr_zone.example sample-id
```

## `ibm_cbr_rule`
{: #cbr_rule}

Create, update, or delete an cbr_rule.
{: shortdesc}

### Sample Terraform code
{: #cbr_rule-sample}

```
resource "ibm_cbr_rule" "cbr_rule" {
  description = "this is an example of rule"
}
```

### Input parameters
{: #cbr_rule-input}

Review the input parameters that you can specify for your resource. {: shortdesc}

|Name|Data type|Required/optional|Description|Forces new resource|
|----|-----------|-------|----------|--------------------|
|`contexts`|List|Optional|The contexts this rule applies to. The maximum length is `1000` items. The minimum length is `1` item.|No|
|`description`|String|Optional|The description of the rule. The maximum length is `300` characters. The minimum length is `0` characters. The value must match regular expression `/^[\\x20-\\xFE]*$/`.|No|
|`resources`|List|Optional|The resources this rule apply to. The maximum length is `1` item. The minimum length is `1` item.|No|
|`transaction_id`|String|Optional|The UUID that is used to correlate and track transactions. If you omit this field, the service generates and sends a transaction ID in the response.**Note:** To help with debugging, we strongly recommend that you generate and supply a `Transaction-Id` with each request. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.|No|

### Output parameters
{: #cbr_rule-output}

Review the output parameters that you can access after your resource is created. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The unique identifier of the cbr_rule.|
|`created_at`|String|The time the resource was created.|
|`created_by_id`|String|IAM ID of the user or service which created the resource.|
|`crn`|String|The rule CRN.|
|`href`|String|The href link to the resource.|
|`last_modified_at`|String|The last time the resource was modified.|
|`last_modified_by_id`|String|IAM ID of the user or service which modified the resource.|
|`version`|String|The version of the cbr_rule.|

### Import
{: #cbr_rule-import}

`ibm_cbr_rule` can be imported by ID

```
$ terraform import ibm_cbr_rule.example sample-id
```


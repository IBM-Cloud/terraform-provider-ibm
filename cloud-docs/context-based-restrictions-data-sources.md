---

copyright:
  years: 2021
lastupdated: "2021"

keywords: terraform

subcollection: terraform

---

# Context Based Restrictions data sources
{: #context-based-restrictions-data-sources}

Review the data sources that you can use to retrieve information about your Context Based Restrictions resources.
All data sources are imported as read-only information. You can reference the output parameters for each data source by using Terraform interpolation syntax.

Before you start working with your data source, make sure to review the [required parameters](/docs/terraform?topic=terraform-provider-reference#required-parameters)
that you need to specify in the `provider` block of your Terraform configuration file.
{: important}

## `ibm_cbr_zone`
{: #cbr_zone}

Retrieve information about cbr_zone.
{: shortdesc}

### Sample Terraform code
{: #cbr_zone-sample}

```
data "ibm_cbr_zone" "cbr_zone" {
  zone_id = "zone_id"
}
```

### Input parameters
{: #cbr_zone-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`zone_id`|String|Required|The ID of a zone. The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[a-fA-F0-9]{32}$/`.|

### Output parameters
{: #cbr_zone-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`account_id`|String|The id of the account owning this zone. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.|
|`address_count`|Integer|The number of addresses in the zone.|
|`addresses`|List|The list of addresses in the zone. The maximum length is `1000` items. The minimum length is `1` item.|
|`addresses.type`|String|The type of address. Allowable values are: `ipAddress`, `ipRange`, `subnet`, `vpc`, `serviceRef`.|
|`addresses.value`|String|The IP address. The maximum length is `45` characters. The minimum length is `7` characters. The value must match regular expression `/^[a-zA-Z0-9:.]+$/`.|
|`addresses.ref`|List|A service reference value. This list contains only one item.|
|`addresses.ref.account_id`|String|The id of the account owning the service. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.|
|`addresses.ref.service_type`|String|The service type. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z_]+$/`.|
|`addresses.ref.service_name`|String|The service name. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\\-]+$/`.|
|`addresses.ref.service_instance`|String|The service instance. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\\-\/]+$/`.|
|`created_at`|String|The time the resource was created.|
|`created_by_id`|String|IAM ID of the user or service which created the resource.|
|`crn`|String|The zone CRN.|
|`description`|String|The description of the zone. The maximum length is `300` characters. The minimum length is `0` characters. The value must match regular expression `/^[\\x20-\\xFE]*$/`.|
|`excluded`|List|The list of excluded addresses in the zone. The maximum length is `1000` items.|
|`excluded.type`|String|The type of address. Allowable values are: `ipAddress`, `ipRange`, `subnet`, `vpc`, `serviceRef`.|
|`excluded.value`|String|The IP address. The maximum length is `45` characters. The minimum length is `7` characters. The value must match regular expression `/^[a-zA-Z0-9:.]+$/`.|
|`excluded.ref`|List|A service reference value. This list contains only one item.|
|`excluded.ref.account_id`|String|The id of the account owning the service. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.|
|`excluded.ref.service_type`|String|The service type. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z_]+$/`.|
|`excluded.ref.service_name`|String|The service name. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\\-]+$/`.|
|`excluded.ref.service_instance`|String|The service instance. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\\-\/]+$/`.|
|`excluded_count`|Integer|The number of excluded addresses in the zone.|
|`href`|String|The href link to the resource.|
|`id`|String|The globally unique ID of the zone.|
|`last_modified_at`|String|The last time the resource was modified.|
|`last_modified_by_id`|String|IAM ID of the user or service which modified the resource.|
|`name`|String|The name of the zone. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 \\-_]+$/`.|

## `ibm_cbr_rule`
{: #cbr_rule}

Retrieve information about cbr_rule.
{: shortdesc}

### Sample Terraform code
{: #cbr_rule-sample}

```
data "ibm_cbr_rule" "cbr_rule" {
  rule_id = "rule_id"
}
```

### Input parameters
{: #cbr_rule-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`rule_id`|String|Required|The ID of a rule. The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[a-fA-F0-9]{32}$/`.|

### Output parameters
{: #cbr_rule-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`contexts`|List|The contexts this rule applies to. The maximum length is `1000` items. The minimum length is `1` item.|
|`contexts.attributes`|List|The attributes. The minimum length is `1` item.|
|`contexts.attributes.name`|String|The attribute name. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9]+$/`.|
|`contexts.attributes.value`|String|The attribute value. The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[\\S\\s]+$/`.|
|`created_at`|String|The time the resource was created.|
|`created_by_id`|String|IAM ID of the user or service which created the resource.|
|`crn`|String|The rule CRN.|
|`description`|String|The description of the rule. The maximum length is `300` characters. The minimum length is `0` characters. The value must match regular expression `/^[\\x20-\\xFE]*$/`.|
|`href`|String|The href link to the resource.|
|`id`|String|The globally unique ID of the rule.|
|`last_modified_at`|String|The last time the resource was modified.|
|`last_modified_by_id`|String|IAM ID of the user or service which modified the resource.|
|`resources`|List|The resources this rule apply to. The maximum length is `1` item. The minimum length is `1` item.|
|`resources.attributes`|List|The resource attributes. The minimum length is `1` item.|
|`resources.attributes.name`|String|The attribute name. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9]+$/`.|
|`resources.attributes.value`|String|The attribute value. The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[\\S\\s]+$/`.|
|`resources.attributes.operator`|String|The attribute operator. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9]+$/`.|
|`resources.tags`|List|The optional resource tags. The maximum length is `10` items. The minimum length is `1` item.|
|`resources.tags.name`|String|The tag attribute name. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 _.-]+$/`.|
|`resources.tags.value`|String|The tag attribute value. The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 _*?.-]+$/`.|
|`resources.tags.operator`|String|The attribute operator. The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9]+$/`.|


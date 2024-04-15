# Examples for Cloud Logs API

These examples illustrate how to use the resources and data sources associated with Cloud Logs API.

The following resources are supported:
* ibm_logs_alert
* ibm_logs_rule_group
* ibm_logs_outgoing_webhook
* ibm_logs_policy
* ibm_logs_dashboard
* ibm_logs_e2m
* ibm_logs_view
* ibm_logs_view_folder

The following data sources are supported:
* ibm_logs_alert
* ibm_logs_alerts
* ibm_logs_rule_group
* ibm_logs_rule_groups
* ibm_logs_outgoing_webhooks
* ibm_logs_outgoing_webhook
* ibm_logs_policy
* ibm_logs_policies
* ibm_logs_dashboard
* ibm_logs_e2m
* ibm_logs_e2ms
* ibm_logs_view
* ibm_logs_views
* ibm_logs_view_folder
* ibm_logs_view_folders

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Cloud Logs API resources

### Resource: ibm_logs_alert

```hcl
resource "ibm_logs_alert" "logs_alert_instance" {
  name = var.logs_alert_name
  description = var.logs_alert_description
  is_active = var.logs_alert_is_active
  severity = var.logs_alert_severity
  expiration = var.logs_alert_expiration
  condition = var.logs_alert_condition
  notification_groups = var.logs_alert_notification_groups
  filters = var.logs_alert_filters
  active_when = var.logs_alert_active_when
  notification_payload_filters = var.logs_alert_notification_payload_filters
  meta_labels = var.logs_alert_meta_labels
  meta_labels_strings = var.logs_alert_meta_labels_strings
  tracing_alert = var.logs_alert_tracing_alert
  incident_settings = var.logs_alert_incident_settings
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | Alert name. | `string` | true |
| description | Alert description. | `string` | false |
| is_active | Alert is active. | `bool` | true |
| severity | Alert severity. | `string` | true |
| expiration | Alert expiration date. | `` | false |
| condition | Alert condition. | `` | true |
| notification_groups | Alert notification groups. | `list()` | true |
| filters | Alert filters. | `` | true |
| active_when | When should the alert be active. | `` | false |
| notification_payload_filters | JSON keys to include in the alert notification, if left empty get the full log text in the alert notification. | `list(string)` | false |
| meta_labels | The Meta labels to add to the alert. | `list()` | false |
| meta_labels_strings | The Meta labels to add to the alert as string with ':' separator. | `list(string)` | false |
| tracing_alert | The definition for tracing alert. | `` | false |
| incident_settings | Incident settings, will create the incident based on this configuration. | `` | false |

#### Outputs

| Name | Description |
|------|-------------|
| unique_identifier | Alert unique identifier. |

### Resource: ibm_logs_rule_group

```hcl
resource "ibm_logs_rule_group" "logs_rule_group_instance" {
  name = var.logs_rule_group_name
  description = var.logs_rule_group_description
  creator = var.logs_rule_group_creator
  enabled = var.logs_rule_group_enabled
  rule_matchers = var.logs_rule_group_rule_matchers
  rule_subgroups = var.logs_rule_group_rule_subgroups
  order = var.logs_rule_group_order
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | The name of the rule group. | `string` | true |
| description | A description for the rule group, should express what is the rule group purpose. | `string` | false |
| creator | The creator of the rule group. | `string` | false |
| enabled | Whether or not the rule is enabled. | `bool` | false |
| rule_matchers | // Optional rule matchers which if matched will make the rule go through the rule group. | `list()` | false |
| rule_subgroups | Rule subgroups. Will try to execute the first rule subgroup, and if not matched will try to match the next one in order. | `list()` | true |
| order | // The order in which the rule group will be evaluated. The lower the order, the more priority the group will have. Not providing the order will by default create a group with the last order. | `number` | false |

### Resource: ibm_logs_outgoing_webhook

```hcl
resource "ibm_logs_outgoing_webhook" "logs_outgoing_webhook_instance" {
  type = var.logs_outgoing_webhook_type
  name = var.logs_outgoing_webhook_name
  url = var.logs_outgoing_webhook_url
  ibm_event_notifications = var.logs_outgoing_webhook_ibm_event_notifications
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| type | Outbound webhook type. | `string` | true |
| name | The name of the outbound webhook. | `string` | true |
| url | The URL of the outbound webhook. | `string` | true |
| ibm_event_notifications | The configuration of an IBM Event Notifications outbound webhook. | `` | false |

#### Outputs

| Name | Description |
|------|-------------|
| created_at | The creation time of the outbound webhook. |
| updated_at | The update time of the outbound webhook. |
| external_id | The external ID of the outbound webhook. |

### Resource: ibm_logs_policy

```hcl
resource "ibm_logs_policy" "logs_policy_instance" {
  name = var.logs_policy_name
  description = var.logs_policy_description
  priority = var.logs_policy_priority
  application_rule = var.logs_policy_application_rule
  subsystem_rule = var.logs_policy_subsystem_rule
  archive_retention = var.logs_policy_archive_retention
  log_rules = var.logs_policy_log_rules
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | name of policy. | `string` | true |
| description | description of policy. | `string` | false |
| priority | the data pipeline sources that match the policy rules will go through. | `string` | true |
| application_rule | rule for matching with application. | `` | false |
| subsystem_rule | rule for matching with application. | `` | false |
| archive_retention | archive retention definition. | `` | false |
| log_rules | log rules. | `` | false |

#### Outputs

| Name | Description |
|------|-------------|
| company_id | company id. |
| deleted | soft deletion flag. |
| enabled | enabled flag. |
| order | order of policy in relation to other policies. |
| created_at | created at timestamp. |
| updated_at | updated at timestamp. |

### Resource: ibm_logs_dashboard

```hcl
resource "ibm_logs_dashboard" "logs_dashboard_instance" {
  href = var.logs_dashboard_href
  name = var.logs_dashboard_name
  description = var.logs_dashboard_description
  layout = var.logs_dashboard_layout
  variables = var.logs_dashboard_variables
  filters = var.logs_dashboard_filters
  annotations = var.logs_dashboard_annotations
  absolute_time_frame = var.logs_dashboard_absolute_time_frame
  relative_time_frame = var.logs_dashboard_relative_time_frame
  folder_id = var.logs_dashboard_folder_id
  folder_path = var.logs_dashboard_folder_path
  false = var.logs_dashboard_false
  two_minutes = var.logs_dashboard_two_minutes
  five_minutes = var.logs_dashboard_five_minutes
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| href | Unique identifier for the dashboard. | `string` | false |
| name | Display name of the dashboard. | `string` | true |
| description | Brief description or summary of the dashboard's purpose or content. | `string` | false |
| layout | Layout configuration for the dashboard's visual elements. | `` | true |
| variables | List of variables that can be used within the dashboard for dynamic content. | `list()` | false |
| filters | List of filters that can be applied to the dashboard's data. | `list()` | false |
| annotations | List of annotations that can be applied to the dashboard's visual elements. | `list()` | false |
| absolute_time_frame | Absolute time frame specifying a fixed start and end time. | `` | false |
| relative_time_frame | Relative time frame specifying a duration from the current time. | `string` | false |
| folder_id | Unique identifier of the folder containing the dashboard. | `` | false |
| folder_path | Path of the folder containing the dashboard. | `` | false |
| false |  | `bool` | false |
| two_minutes |  | `bool` | false |
| five_minutes |  | `bool` | false |

### Resource: ibm_logs_e2m

```hcl
resource "ibm_logs_e2m" "logs_e2m_instance" {
  name = var.logs_e2m_name
  description = var.logs_e2m_description
  metric_labels = var.logs_e2m_metric_labels
  metric_fields = var.logs_e2m_metric_fields
  type = var.logs_e2m_type
  spans_query = var.logs_e2m_spans_query
  logs_query = var.logs_e2m_logs_query
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | E2M name. | `string` | true |
| description | E2m description. | `string` | false |
| metric_labels | E2M metric labels. | `list()` | false |
| metric_fields | E2M metric fields. | `list()` | false |
| type | e2m type. | `string` | false |
| spans_query | spans query. | `` | false |
| logs_query | logs query. | `` | false |

#### Outputs

| Name | Description |
|------|-------------|
| create_time | E2M create time. |
| update_time | E2M update time. |
| permutations | represents E2M permutations limit. |
| is_internal | a flag that represents if the e2m is for internal usage. |

### Resource: ibm_logs_view

```hcl
resource "ibm_logs_view" "logs_view_instance" {
  name = var.logs_view_name
  search_query = var.logs_view_search_query
  time_selection = var.logs_view_time_selection
  filters = var.logs_view_filters
  folder_id = var.logs_view_folder_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | View name. | `string` | true |
| search_query | View search query. | `` | true |
| time_selection | View time selection. | `` | true |
| filters | View selected filters. | `` | false |
| folder_id | View folder id. | `` | false |

### Resource: ibm_logs_view_folder

```hcl
resource "ibm_logs_view_folder" "logs_view_folder_instance" {
  name = var.logs_view_folder_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | Folder name. | `string` | true |

## Cloud Logs API data sources

### Data source: ibm_logs_alert

```hcl
data "ibm_logs_alert" "logs_alert_instance" {
  logs_alert_id = var.data_logs_alert_logs_alert_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| logs_alert_id | Alert id. | `` | true |

#### Outputs

| Name | Description |
|------|-------------|
| name | Alert name. |
| description | Alert description. |
| is_active | Alert is active. |
| severity | Alert severity. |
| expiration | Alert expiration date. |
| condition | Alert condition. |
| notification_groups | Alert notification groups. |
| filters | Alert filters. |
| active_when | When should the alert be active. |
| notification_payload_filters | JSON keys to include in the alert notification, if left empty get the full log text in the alert notification. |
| meta_labels | The Meta labels to add to the alert. |
| meta_labels_strings | The Meta labels to add to the alert as string with ':' separator. |
| tracing_alert | The definition for tracing alert. |
| unique_identifier | Alert unique identifier. |
| incident_settings | Incident settings, will create the incident based on this configuration. |

### Data source: ibm_logs_alerts

```hcl
data "ibm_logs_alerts" "logs_alerts_instance" {
}
```

#### Outputs

| Name | Description |
|------|-------------|
| alerts | Alerts. |

### Data source: ibm_logs_rule_group

```hcl
data "ibm_logs_rule_group" "logs_rule_group_instance" {
  group_id = var.data_logs_rule_group_group_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| group_id | The group id. | `` | true |

#### Outputs

| Name | Description |
|------|-------------|
| name | The name of the rule group. |
| description | A description for the rule group, should express what is the rule group purpose. |
| creator | The creator of the rule group. |
| enabled | Whether or not the rule is enabled. |
| rule_matchers | // Optional rule matchers which if matched will make the rule go through the rule group. |
| rule_subgroups | Rule subgroups. Will try to execute the first rule subgroup, and if not matched will try to match the next one in order. |
| order | // The order in which the rule group will be evaluated. The lower the order, the more priority the group will have. Not providing the order will by default create a group with the last order. |

### Data source: ibm_logs_rule_groups

```hcl
data "ibm_logs_rule_groups" "logs_rule_groups_instance" {
}
```

#### Outputs

| Name | Description |
|------|-------------|
| rulegroups | The rule groups. |

### Data source: ibm_logs_outgoing_webhooks

```hcl
data "ibm_logs_outgoing_webhooks" "logs_outgoing_webhooks_instance" {
  type = var.logs_outgoing_webhooks_type
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| type | Outbound webhook type. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| outgoing_webhooks | List of deployed outbound webhooks. |

### Data source: ibm_logs_outgoing_webhook

```hcl
data "ibm_logs_outgoing_webhook" "logs_outgoing_webhook_instance" {
  logs_outgoing_webhook_id = var.data_logs_outgoing_webhook_logs_outgoing_webhook_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| logs_outgoing_webhook_id | Outbound webhook ID. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| type | Outbound webhook type. |
| name | The name of the outbound webhook. |
| url | The URL of the outbound webhook. |
| created_at | The creation time of the outbound webhook. |
| updated_at | The update time of the outbound webhook. |
| external_id | The external ID of the outbound webhook. |
| ibm_event_notifications | The configuration of an IBM Event Notifications outbound webhook. |

### Data source: ibm_logs_policy

```hcl
data "ibm_logs_policy" "logs_policy_instance" {
  logs_policy_id = var.data_logs_policy_logs_policy_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| logs_policy_id | id of policy. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| company_id | company id. |
| name | name of policy. |
| description | description of policy. |
| priority | the data pipeline sources that match the policy rules will go through. |
| deleted | soft deletion flag. |
| enabled | enabled flag. |
| order | order of policy in relation to other policies. |
| application_rule | rule for matching with application. |
| subsystem_rule | rule for matching with application. |
| created_at | created at timestamp. |
| updated_at | updated at timestamp. |
| archive_retention | archive retention definition. |
| log_rules | log rules. |

### Data source: ibm_logs_policies

```hcl
data "ibm_logs_policies" "logs_policies_instance" {
  enabled_only = var.logs_policies_enabled_only
  source_type = var.logs_policies_source_type
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| enabled_only | optionally filter only enabled policies. | `bool` | false |
| source_type | Source type to filter policies by. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| policies | company policies. |

### Data source: ibm_logs_dashboard

```hcl
data "ibm_logs_dashboard" "logs_dashboard_instance" {
  dashboard_id = var.data_logs_dashboard_dashboard_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| dashboard_id | The ID of the dashboard. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| href | Unique identifier for the dashboard. |
| name | Display name of the dashboard. |
| description | Brief description or summary of the dashboard's purpose or content. |
| layout | Layout configuration for the dashboard's visual elements. |
| variables | List of variables that can be used within the dashboard for dynamic content. |
| filters | List of filters that can be applied to the dashboard's data. |
| annotations | List of annotations that can be applied to the dashboard's visual elements. |
| absolute_time_frame | Absolute time frame specifying a fixed start and end time. |
| relative_time_frame | Relative time frame specifying a duration from the current time. |
| folder_id | Unique identifier of the folder containing the dashboard. |
| folder_path | Path of the folder containing the dashboard. |
| false |  |
| two_minutes |  |
| five_minutes |  |

### Data source: ibm_logs_e2m

```hcl
data "ibm_logs_e2m" "logs_e2m_instance" {
  logs_e2m_id = var.data_logs_e2m_logs_e2m_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| logs_e2m_id | id of e2m to be deleted. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| name | E2M name. |
| description | E2m description. |
| create_time | E2M create time. |
| update_time | E2M update time. |
| permutations | represents E2M permutations limit. |
| metric_labels | E2M metric labels. |
| metric_fields | E2M metric fields. |
| type | e2m type. |
| is_internal | a flag that represents if the e2m is for internal usage. |
| spans_query | spans query. |
| logs_query | logs query. |

### Data source: ibm_logs_e2ms

```hcl
data "ibm_logs_e2ms" "logs_e2ms_instance" {
}
```

#### Outputs

| Name | Description |
|------|-------------|
| events2metrics | List of event to metrics definitions. |

### Data source: ibm_logs_view

```hcl
data "ibm_logs_view" "logs_view_instance" {
  logs_view_id = var.data_logs_view_logs_view_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| logs_view_id | View id. | `number` | true |

#### Outputs

| Name | Description |
|------|-------------|
| name | View name. |
| search_query | View search query. |
| time_selection | View time selection. |
| filters | View selected filters. |
| folder_id | View folder id. |

### Data source: ibm_logs_views

```hcl
data "ibm_logs_views" "logs_views_instance" {
}
```

#### Outputs

| Name | Description |
|------|-------------|
| views | List of views. |

### Data source: ibm_logs_view_folder

```hcl
data "ibm_logs_view_folder" "logs_view_folder_instance" {
  logs_view_folder_id = var.data_logs_view_folder_logs_view_folder_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| logs_view_folder_id | Folder id. | `` | true |

#### Outputs

| Name | Description |
|------|-------------|
| name | Folder name. |

### Data source: ibm_logs_view_folders

```hcl
data "ibm_logs_view_folders" "logs_view_folders_instance" {
}
```

#### Outputs

| Name | Description |
|------|-------------|
| view_folders | List of view folders. |

## Assumptions

1. TODO

## Notes

1. TODO

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |

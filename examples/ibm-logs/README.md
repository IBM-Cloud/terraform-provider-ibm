# Examples for Cloud Logs API

These examples illustrate how to use the resources and data sources associated with Cloud Logs API.

The following resources are supported:
* ibm_logs_alert
* ibm_logs_rule_group
* ibm_logs_outgoing_webhook
* ibm_logs_policy
* ibm_logs_dashboard
* ibm_logs_dashboard_folder
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
* ibm_logs_dashboard_folders
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
| type | The type of the deployed Outbound Integrations to list. | `string` | true |
| name | The name of the Outbound Integration. | `string` | true |
| url | The URL of the Outbound Integration. Null for IBM Event Notifications integration. | `string` | false |
| ibm_event_notifications | The configuration of the IBM Event Notifications Outbound Integration. | `` | false |

#### Outputs

| Name | Description |
|------|-------------|
| created_at | The creation time of the Outbound Integration. |
| updated_at | The update time of the Outbound Integration. |
| external_id | The external ID of the Outbound Integration, for connecting with other parts of the system. |

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
| name | Name of policy. | `string` | true |
| description | Description of policy. | `string` | false |
| priority | The data pipeline sources that match the policy rules will go through. | `string` | true |
| application_rule | Rule for matching with application. | `` | false |
| subsystem_rule | Rule for matching with application. | `` | false |
| archive_retention | Archive retention definition. | `` | false |
| log_rules | Log rules. | `` | false |

#### Outputs

| Name | Description |
|------|-------------|
| company_id | Company ID. |
| deleted | Soft deletion flag. |
| enabled | Enabled flag. |
| order | Order of policy in relation to other policies. |
| created_at | Created at date at utc+0. |
| updated_at | Updated at date at utc+0. |

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
| false | Auto refresh interval is set to off. | `` | false |
| two_minutes | Auto refresh interval is set to two minutes. | `` | false |
| five_minutes | Auto refresh interval is set to five minutes. | `` | false |

### Resource: ibm_logs_dashboard_folder

```hcl
resource "ibm_logs_dashboard_folder" "logs_dashboard_folder_instance" {
  name = var.logs_dashboard_folder_name
  parent_id = var.logs_dashboard_folder_parent_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | The dashboard folder name, required. | `string` | true |
| parent_id | The dashboard folder parent ID, optional. If not set, the folder is a root folder, if set, the folder is a subfolder of the parent folder and needs to be a uuid. | `` | false |

### Resource: ibm_logs_e2m

```hcl
resource "ibm_logs_e2m" "logs_e2m_instance" {
  name = var.logs_e2m_name
  description = var.logs_e2m_description
  metric_labels = var.logs_e2m_metric_labels
  metric_fields = var.logs_e2m_metric_fields
  type = var.logs_e2m_type
  logs_query = var.logs_e2m_logs_query
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | Name of the E2M. | `string` | true |
| description | Description of the E2M. | `string` | false |
| metric_labels | E2M metric labels. | `list()` | false |
| metric_fields | E2M metric fields. | `list()` | false |
| type | E2M type. | `string` | false |
| logs_query | E2M logs query. | `` | false |

#### Outputs

| Name | Description |
|------|-------------|
| create_time | E2M create time. |
| update_time | E2M update time. |
| permutations | Represents the limit of the permutations and if the limit was exceeded. |
| is_internal | A flag that represents if the e2m is for internal usage. |

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
| search_query | View search query. | `` | false |
| time_selection | View time selection. | `` | true |
| filters | View selected filters. | `` | false |
| folder_id | View folder ID. | `` | false |

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
| logs_alert_id | Alert ID. | `` | true |

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
| group_id | The group ID. | `` | true |

#### Outputs

| Name | Description |
|------|-------------|
| name | The name of the rule group. |
| description | A description for the rule group, should express what is the rule group purpose. |
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
| type | The type of the deployed Outbound Integrations to list. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| outgoing_webhooks | The list of deployed Outbound Integrations. |

### Data source: ibm_logs_outgoing_webhook

```hcl
data "ibm_logs_outgoing_webhook" "logs_outgoing_webhook_instance" {
  logs_outgoing_webhook_id = var.data_logs_outgoing_webhook_logs_outgoing_webhook_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| logs_outgoing_webhook_id | The ID of the Outbound Integration to delete. | `` | true |

#### Outputs

| Name | Description |
|------|-------------|
| type | The type of the deployed Outbound Integrations to list. |
| name | The name of the Outbound Integration. |
| url | The URL of the Outbound Integration. Null for IBM Event Notifications integration. |
| created_at | The creation time of the Outbound Integration. |
| updated_at | The update time of the Outbound Integration. |
| external_id | The external ID of the Outbound Integration, for connecting with other parts of the system. |
| ibm_event_notifications | The configuration of the IBM Event Notifications Outbound Integration. |

### Data source: ibm_logs_policy

```hcl
data "ibm_logs_policy" "logs_policy_instance" {
  logs_policy_id = var.data_logs_policy_logs_policy_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| logs_policy_id | ID of policy. | `` | true |

#### Outputs

| Name | Description |
|------|-------------|
| company_id | Company ID. |
| name | Name of policy. |
| description | Description of policy. |
| priority | The data pipeline sources that match the policy rules will go through. |
| deleted | Soft deletion flag. |
| enabled | Enabled flag. |
| order | Order of policy in relation to other policies. |
| application_rule | Rule for matching with application. |
| subsystem_rule | Rule for matching with application. |
| created_at | Created at date at utc+0. |
| updated_at | Updated at date at utc+0. |
| archive_retention | Archive retention definition. |
| log_rules | Log rules. |

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
| enabled_only | Optionally filter only enabled policies. | `bool` | false |
| source_type | Source type to filter policies by. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| policies | Company policies. |

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
| false | Auto refresh interval is set to off. |
| two_minutes | Auto refresh interval is set to two minutes. |
| five_minutes | Auto refresh interval is set to five minutes. |

### Data source: ibm_logs_dashboard_folders

```hcl
data "ibm_logs_dashboard_folders" "logs_dashboard_folders_instance" {
}
```

#### Outputs

| Name | Description |
|------|-------------|
| folders | The list of folders. |

### Data source: ibm_logs_e2m

```hcl
data "ibm_logs_e2m" "logs_e2m_instance" {
  logs_e2m_id = var.data_logs_e2m_logs_e2m_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| logs_e2m_id | ID of e2m to be deleted. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| name | Name of the E2M. |
| description | Description of the E2M. |
| create_time | E2M create time. |
| update_time | E2M update time. |
| permutations | Represents the limit of the permutations and if the limit was exceeded. |
| metric_labels | E2M metric labels. |
| metric_fields | E2M metric fields. |
| type | E2M type. |
| is_internal | A flag that represents if the e2m is for internal usage. |
| logs_query | E2M logs query. |

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
| logs_view_id | View ID. | `number` | true |

#### Outputs

| Name | Description |
|------|-------------|
| name | View name. |
| search_query | View search query. |
| time_selection | View time selection. |
| filters | View selected filters. |
| folder_id | View folder ID. |

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
| logs_view_folder_id | Folder ID. | `` | true |

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

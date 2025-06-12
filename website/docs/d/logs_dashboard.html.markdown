---
layout: "ibm"
page_title: "IBM : ibm_logs_dashboard"
description: |-
  Get information about logs_dashboard
subcategory: "Cloud Logs"
---


# ibm_logs_dashboard

Provides a read-only data source to retrieve information about a logs_dashboard. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_dashboard" "logs_dashboard_instance" {
  instance_id  = ibm_logs_dashboard.logs_dashboard_instance.instance_id
  region       = ibm_logs_dashboard.logs_dashboard_instance.region
  dashboard_id = ibm_logs_dashboard.logs_dashboard_instance.dashboard_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String)  Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.
* `dashboard_id` - (Required, Forces new resource, String) The ID of the dashboard.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_dashboard.
* `absolute_time_frame` - (List) Absolute time frame specifying a fixed start and end time.
Nested schema for **absolute_time_frame**:
	* `from` - (String) From is the start of the time frame.
	* `to` - (String) To is the end of the time frame.

* `annotations` - (List) List of annotations that can be applied to the dashboard's visual elements.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **annotations**:
	* `enabled` - (Boolean) Whether the annotation is enabled.
	* `href` - (String) Unique identifier within the dashboard.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
	* `id` - (String) Unique identifier within the dashboard.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
	* `name` - (String) Name of the annotation.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `source` - (List) Source of the annotation events.
	Nested schema for **source**:
		* `logs` - (List) Logs source.
		Nested schema for **logs**:
			* `label_fields` - (List) Labels to display in the annotation.
			  * Constraints: The maximum length is `10` items. The minimum length is `0` items.
			Nested schema for **label_fields**:
				* `keypath` - (List) Path within the dataset scope.
				  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
				* `scope` - (String) Scope of the dataset.
				  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
			* `lucene_query` - (List) Lucene query.
			Nested schema for **lucene_query**:
				* `value` - (String) The Lucene query string.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
			* `message_template` - (String) Template for the annotation message.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
			* `strategy` - (List) Strategy for turning logs data into annotations.
			Nested schema for **strategy**:
				* `duration` - (List) Event start timestamp and duration are extracted from the log entry.
				Nested schema for **duration**:
					* `duration_field` - (List) Field to count distinct values of.
					Nested schema for **duration_field**:
						* `keypath` - (List) Path within the dataset scope.
						  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
						* `scope` - (String) Scope of the dataset.
						  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
					* `start_timestamp_field` - (List) Field to count distinct values of.
					Nested schema for **start_timestamp_field**:
						* `keypath` - (List) Path within the dataset scope.
						  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
						* `scope` - (String) Scope of the dataset.
						  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
				* `instant` - (List) Event timestamp is extracted from the log entry.
				Nested schema for **instant**:
					* `timestamp_field` - (List) Field to count distinct values of.
					Nested schema for **timestamp_field**:
						* `keypath` - (List) Path within the dataset scope.
						  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
						* `scope` - (String) Scope of the dataset.
						  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
				* `range` - (List) Event start and end timestamps are extracted from the log entry.
				Nested schema for **range**:
					* `end_timestamp_field` - (List) Field to count distinct values of.
					Nested schema for **end_timestamp_field**:
						* `keypath` - (List) Path within the dataset scope.
						  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
						* `scope` - (String) Scope of the dataset.
						  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
					* `start_timestamp_field` - (List) Field to count distinct values of.
					Nested schema for **start_timestamp_field**:
						* `keypath` - (List) Path within the dataset scope.
						  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
						* `scope` - (String) Scope of the dataset.
						  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
		* `metrics` - (List) Metrics source.
		Nested schema for **metrics**:
			* `labels` - (List) Labels to display in the annotation.
			  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
			* `message_template` - (String) Template for the annotation message.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
			* `promql_query` - (List) PromQL query.
			Nested schema for **promql_query**:
				* `value` - (String) The PromQL query string.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
			* `strategy` - (List) Strategy for turning metrics data into annotations.
			Nested schema for **strategy**:
				* `start_time_metric` - (List) Take first data point and use its value as annotation timestamp (instead of point own timestamp).
				Nested schema for **start_time_metric**:

* `description` - (String) Brief description or summary of the dashboard's purpose or content.
  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

* `false` - (List) Auto refresh interval is set to off.
Nested schema for **false**:

* `filters` - (List) List of filters that can be applied to the dashboard's data.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **filters**:
	* `collapsed` - (Boolean) Indicates if the filter's UI representation should be collapsed or expanded.
	* `enabled` - (Boolean) Indicates if the filter is currently enabled or not.
	* `source` - (List) Filters to be applied to query results.
	Nested schema for **source**:
		* `logs` - (List) Extra filtering on top of the Lucene query.
		Nested schema for **logs**:
			* `observation_field` - (List) Field to count distinct values of.
			Nested schema for **observation_field**:
				* `keypath` - (List) Path within the dataset scope.
				  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
				* `scope` - (String) Scope of the dataset.
				  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
			* `operator` - (List) Operator to use for filtering the logs.
			Nested schema for **operator**:
				* `equals` - (List) Equality comparison.
				Nested schema for **equals**:
					* `selection` - (List) Selection criteria for the equality comparison.
					Nested schema for **selection**:
						* `all` - (List) Represents a selection of all values.
						Nested schema for **all**:
						* `list` - (List) Represents a selection from a list of values.
						Nested schema for **list**:
							* `values` - (List) List of values for the selection.
							  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
				* `not_equals` - (List) Non-equality comparison.
				Nested schema for **not_equals**:
					* `selection` - (List) Selection criteria for the non-equality comparison.
					Nested schema for **selection**:
						* `list` - (List) Represents a selection from a list of values.
						Nested schema for **list**:
							* `values` - (List) List of values for the selection.
							  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
		* `metrics` - (List) Filtering to be applied to query results.
		Nested schema for **metrics**:
			* `label` - (String) Label associated with the metric.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
			* `operator` - (List) Operator to use for filtering the logs.
			Nested schema for **operator**:
				* `equals` - (List) Equality comparison.
				Nested schema for **equals**:
					* `selection` - (List) Selection criteria for the equality comparison.
					Nested schema for **selection**:
						* `all` - (List) Represents a selection of all values.
						Nested schema for **all**:
						* `list` - (List) Represents a selection from a list of values.
						Nested schema for **list**:
							* `values` - (List) List of values for the selection.
							  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
				* `not_equals` - (List) Non-equality comparison.
				Nested schema for **not_equals**:
					* `selection` - (List) Selection criteria for the non-equality comparison.
					Nested schema for **selection**:
						* `list` - (List) Represents a selection from a list of values.
						Nested schema for **list**:
							* `values` - (List) List of values for the selection.
							  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.

* `five_minutes` - (List) Auto refresh interval is set to five minutes.
Nested schema for **five_minutes**:

* `folder_id` - (List) Unique identifier of the folder containing the dashboard.
Nested schema for **folder_id**:
	* `value` - (String) The UUID value.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

* `folder_path` - (List) Path of the folder containing the dashboard.
Nested schema for **folder_path**:
	* `segments` - (List) The segments of the folder path.
	  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.

* `href` - (String) Unique identifier for the dashboard.
  * Constraints: The maximum length is `21` characters. The minimum length is `21` characters. The value must match regular expression `/^[a-zA-Z0-9]{21}$/`.

* `layout` - (List) Layout configuration for the dashboard's visual elements.
Nested schema for **layout**:
	* `sections` - (List) The sections of the layout.
	  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
	Nested schema for **sections**:
		* `href` - (String) The unique identifier of the section within the layout.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
		* `id` - (List) Unique identifier of the folder containing the dashboard.
		Nested schema for **id**:
			* `value` - (String) The UUID value.
			  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
		* `rows` - (List) The rows of the section.
		  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
		Nested schema for **rows**:
			* `appearance` - (List) The appearance of the row, such as height.
			Nested schema for **appearance**:
				* `height` - (Integer) The height of the row.
			* `href` - (String) The unique identifier of the row within the layout.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
			* `id` - (List) Unique identifier of the folder containing the dashboard.
			Nested schema for **id**:
				* `value` - (String) The UUID value.
				  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
			* `widgets` - (List) The widgets of the row.
			  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
			Nested schema for **widgets**:
				* `created_at` - (String) Creation timestamp.
				* `definition` - (List) Widget definition, contains the widget type and its configuration.
				Nested schema for **definition**:
					* `bar_chart` - (List) Bar chart widget.
					Nested schema for **bar_chart**:
						* `color_scheme` - (String) Supported vaues: classic, severity, cold, negative, green, red, blue.
						  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `colors_by` - (List) Coloring mode.
						Nested schema for **colors_by**:
							* `aggregation` - (List) Each aggregation will have different color and stack color will be derived from aggregation color.
							Nested schema for **aggregation**:
							* `group_by` - (List) Each group will have different color and stack color will be derived from group color.
							Nested schema for **group_by**:
							* `stack` - (List) Each stack will have the same color across all groups.
							Nested schema for **stack**:
						* `data_mode_type` - (String) Data mode type.
						  * Constraints: Allowable values are: `high_unspecified`, `archive`.
						* `group_name_template` - (String) Template for bar labels.
						  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `max_bars_per_chart` - (Integer) Maximum number of bars to present in the chart.
						* `query` - (List) Data source specific query, defines from where and how to fetch the data.
						Nested schema for **query**:
							* `dataprime` - (List) Query based on Dataprime language.
							Nested schema for **dataprime**:
								* `dataprime_query` - (List) Dataprime query.
								Nested schema for **dataprime_query**:
									* `text` - (String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `filters` - (List) Extra filter on top of the Dataprime query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `logs` - (List) Extra filtering on top of the Lucene query.
									Nested schema for **logs**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `operator` - (List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
									* `metrics` - (List) Filtering to be applied to query results.
									Nested schema for **metrics**:
										* `label` - (String) Label associated with the metric.
										  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
										* `operator` - (List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `group_names` - (List) Fields to group by.
								  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `2` items. The minimum length is `1` item.
								* `stacked_group_name` - (String) Field to stack by.
								  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `logs` - (List) Logs specific query.
							Nested schema for **logs**:
								* `aggregation` - (List) Aggregations.
								Nested schema for **aggregation**:
									* `average` - (List) Calculate average value of log field.
									Nested schema for **average**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `count` - (List) Count the number of entries.
									Nested schema for **count**:
									* `count_distinct` - (List) Count the number of distinct values of log field.
									Nested schema for **count_distinct**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `max` - (List) Calculate maximum value of log field.
									Nested schema for **max**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `min` - (List) Calculate minimum value of log field.
									Nested schema for **min**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `percentile` - (List) Calculate percentile value of log field.
									Nested schema for **percentile**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `percent` - (Float) Value in range (0, 100].
									* `sum` - (List) Sum values of log field.
									Nested schema for **sum**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
								* `filters` - (List) Extra filter on top of Lucene query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `observation_field` - (List) Field to count distinct values of.
									Nested schema for **observation_field**:
										* `keypath` - (List) Path within the dataset scope.
										  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
										* `scope` - (String) Scope of the dataset.
										  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `operator` - (List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `group_names_fields` - (List) Fiel to group by.
								  * Constraints: The maximum length is `2` items. The minimum length is `1` item.
								Nested schema for **group_names_fields**:
									* `keypath` - (List) Path within the dataset scope.
									  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
									* `scope` - (String) Scope of the dataset.
									  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
								* `lucene_query` - (List) Lucene query.
								Nested schema for **lucene_query**:
									* `value` - (String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `stacked_group_name_field` - (List) Field to count distinct values of.
								Nested schema for **stacked_group_name_field**:
									* `keypath` - (List) Path within the dataset scope.
									  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
									* `scope` - (String) Scope of the dataset.
									  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
							* `metrics` - (List) Metrics specific query.
							Nested schema for **metrics**:
								* `filters` - (List) Extra filter on top of the PromQL query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `label` - (String) Label associated with the metric.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
									* `operator` - (List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `group_names` - (List) Labels to group by.
								  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `2` items. The minimum length is `1` item.
								* `promql_query` - (List) PromQL query.
								Nested schema for **promql_query**:
									* `value` - (String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `stacked_group_name` - (String) Label to stack by.
								  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `scale_type` - (String) Scale type.
						  * Constraints: Allowable values are: `unspecified`, `linear`, `logarithmic`.
						* `sort_by` - (String) Sorting mode.
						  * Constraints: Allowable values are: `unspecified`, `value`, `name`.
						* `stack_definition` - (List) Stack definition.
						Nested schema for **stack_definition**:
							* `max_slices_per_bar` - (Integer) Maximum number of slices per bar.
							* `stack_name_template` - (String) Template for stack slice label.
							  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `unit` - (String) Unit of the data.
						  * Constraints: Allowable values are: `unspecified`, `microseconds`, `milliseconds`, `seconds`, `bytes`, `kbytes`, `mbytes`, `gbytes`, `bytes_iec`, `kibytes`, `mibytes`, `gibytes`, `eur_cents`, `eur`, `usd_cents`, `usd`.
						* `x_axis` - (List) X axis mode.
						Nested schema for **x_axis**:
							* `time` - (List) Time based axis.
							Nested schema for **time**:
								* `buckets_presented` - (Integer) Buckets presented.
								* `interval` - (String) Time interval.
								  * Constraints: The maximum length is `10` characters. The minimum length is `2` characters. The value must match regular expression `/^[0-9]+[smhdw]?$/`.
							* `value` - (List) Categorical axis.
							Nested schema for **value**:
					* `data_table` - (List) Data table widget.
					Nested schema for **data_table**:
						* `columns` - (List) Columns to display, their order and width.
						  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
						Nested schema for **columns**:
							* `field` - (String) References a field in result set. In case of aggregation, it references the aggregation identifier.
							  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `width` - (Integer) Column width.
						* `data_mode_type` - (String) Data mode type.
						  * Constraints: Allowable values are: `high_unspecified`, `archive`.
						* `order_by` - (List) Column used for ordering the results.
						Nested schema for **order_by**:
							* `field` - (String) The field to order by.
							  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `order_direction` - (String) The direction of the order: ascending or descending.
							  * Constraints: Allowable values are: `unspecified`, `asc`, `desc`.
						* `query` - (List) Data source specific query, defines from where and how to fetch the data.
						Nested schema for **query**:
							* `dataprime` - (List) Query based on Dataprime language.
							Nested schema for **dataprime**:
								* `dataprime_query` - (List) Dataprime query.
								Nested schema for **dataprime_query**:
									* `text` - (String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `filters` - (List) Extra filtering on top of the Dataprime query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `logs` - (List) Extra filtering on top of the Lucene query.
									Nested schema for **logs**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `operator` - (List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
									* `metrics` - (List) Filtering to be applied to query results.
									Nested schema for **metrics**:
										* `label` - (String) Label associated with the metric.
										  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
										* `operator` - (List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
							* `logs` - (List) Logs specific query.
							Nested schema for **logs**:
								* `filters` - (List) Extra filtering on top of the Lucene query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `observation_field` - (List) Field to count distinct values of.
									Nested schema for **observation_field**:
										* `keypath` - (List) Path within the dataset scope.
										  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
										* `scope` - (String) Scope of the dataset.
										  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `operator` - (List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `grouping` - (List) Grouping and aggregation.
								Nested schema for **grouping**:
									* `aggregations` - (List) Aggregations.
									  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
									Nested schema for **aggregations**:
										* `aggregation` - (List) Aggregations.
										Nested schema for **aggregation**:
											* `average` - (List) Calculate average value of log field.
											Nested schema for **average**:
												* `observation_field` - (List) Field to count distinct values of.
												Nested schema for **observation_field**:
													* `keypath` - (List) Path within the dataset scope.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
													* `scope` - (String) Scope of the dataset.
													  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
											* `count` - (List) Count the number of entries.
											Nested schema for **count**:
											* `count_distinct` - (List) Count the number of distinct values of log field.
											Nested schema for **count_distinct**:
												* `observation_field` - (List) Field to count distinct values of.
												Nested schema for **observation_field**:
													* `keypath` - (List) Path within the dataset scope.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
													* `scope` - (String) Scope of the dataset.
													  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
											* `max` - (List) Calculate maximum value of log field.
											Nested schema for **max**:
												* `observation_field` - (List) Field to count distinct values of.
												Nested schema for **observation_field**:
													* `keypath` - (List) Path within the dataset scope.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
													* `scope` - (String) Scope of the dataset.
													  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
											* `min` - (List) Calculate minimum value of log field.
											Nested schema for **min**:
												* `observation_field` - (List) Field to count distinct values of.
												Nested schema for **observation_field**:
													* `keypath` - (List) Path within the dataset scope.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
													* `scope` - (String) Scope of the dataset.
													  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
											* `percentile` - (List) Calculate percentile value of log field.
											Nested schema for **percentile**:
												* `observation_field` - (List) Field to count distinct values of.
												Nested schema for **observation_field**:
													* `keypath` - (List) Path within the dataset scope.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
													* `scope` - (String) Scope of the dataset.
													  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
												* `percent` - (Float) Value in range (0, 100].
											* `sum` - (List) Sum values of log field.
											Nested schema for **sum**:
												* `observation_field` - (List) Field to count distinct values of.
												Nested schema for **observation_field**:
													* `keypath` - (List) Path within the dataset scope.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
													* `scope` - (String) Scope of the dataset.
													  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `id` - (String) Aggregation identifier, must be unique within grouping configuration.
										  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
										* `is_visible` - (Boolean) Whether the aggregation is visible.
										* `name` - (String) Aggregation name, used as column name.
										  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
									* `group_bys` - (List) Fields to group by.
									  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
									Nested schema for **group_bys**:
										* `keypath` - (List) Path within the dataset scope.
										  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
										* `scope` - (String) Scope of the dataset.
										  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
								* `lucene_query` - (List) Lucene query.
								Nested schema for **lucene_query**:
									* `value` - (String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `metrics` - (List) Metrics specific query.
							Nested schema for **metrics**:
								* `filters` - (List) Extra filtering on top of the PromQL query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `label` - (String) Label associated with the metric.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
									* `operator` - (List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `promql_query` - (List) PromQL query.
								Nested schema for **promql_query**:
									* `value` - (String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `results_per_page` - (Integer) Number of results per page.
						* `row_style` - (String) Display style for table rows.
						  * Constraints: Allowable values are: `unspecified`, `one_line`, `two_line`, `condensed`, `json`, `list`.
					* `gauge` - (List) Gauge widget.
					Nested schema for **gauge**:
						* `data_mode_type` - (String) Data mode type.
						  * Constraints: Allowable values are: `high_unspecified`, `archive`.
						* `max` - (Float) Maximum value of the gauge.
						* `min` - (Float) Minimum value of the gauge.
						* `query` - (List) Data source specific query, defines from where and how to fetch the data.
						Nested schema for **query**:
							* `dataprime` - (List) Query based on Dataprime language.
							Nested schema for **dataprime**:
								* `dataprime_query` - (List) Dataprime query.
								Nested schema for **dataprime_query**:
									* `text` - (String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `filters` - (List) Extra filters applied on top of Dataprime query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `logs` - (List) Extra filtering on top of the Lucene query.
									Nested schema for **logs**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `operator` - (List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
									* `metrics` - (List) Filtering to be applied to query results.
									Nested schema for **metrics**:
										* `label` - (String) Label associated with the metric.
										  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
										* `operator` - (List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
							* `logs` - (List) Logs specific query.
							Nested schema for **logs**:
								* `filters` - (List) Extra filters applied on top of Lucene query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `observation_field` - (List) Field to count distinct values of.
									Nested schema for **observation_field**:
										* `keypath` - (List) Path within the dataset scope.
										  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
										* `scope` - (String) Scope of the dataset.
										  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `operator` - (List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `logs_aggregation` - (List) Aggregations.
								Nested schema for **logs_aggregation**:
									* `average` - (List) Calculate average value of log field.
									Nested schema for **average**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `count` - (List) Count the number of entries.
									Nested schema for **count**:
									* `count_distinct` - (List) Count the number of distinct values of log field.
									Nested schema for **count_distinct**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `max` - (List) Calculate maximum value of log field.
									Nested schema for **max**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `min` - (List) Calculate minimum value of log field.
									Nested schema for **min**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `percentile` - (List) Calculate percentile value of log field.
									Nested schema for **percentile**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `percent` - (Float) Value in range (0, 100].
									* `sum` - (List) Sum values of log field.
									Nested schema for **sum**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
								* `lucene_query` - (List) Lucene query.
								Nested schema for **lucene_query**:
									* `value` - (String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `metrics` - (List) Metrics specific query.
							Nested schema for **metrics**:
								* `aggregation` - (String) Aggregation. When AGGREGATION_UNSPECIFIED is selected, widget uses instant query. Otherwise, it uses range query.
								  * Constraints: Allowable values are: `unspecified`, `last`, `min`, `max`, `avg`, `sum`.
								* `filters` - (List) Extra filters applied on top of PromQL query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `label` - (String) Label associated with the metric.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
									* `operator` - (List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `promql_query` - (List) PromQL query.
								Nested schema for **promql_query**:
									* `value` - (String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `show_inner_arc` - (Boolean) Show inner arc (styling).
						* `show_outer_arc` - (Boolean) Show outer arc (styling).
						* `threshold_by` - (String) What threshold color should be applied to: value or background.
						  * Constraints: Allowable values are: `unspecified`, `value`, `background`.
						* `thresholds` - (List) Thresholds for the gauge, values at which the gauge changes color.
						  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
						Nested schema for **thresholds**:
							* `color` - (String) Color.
							  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `from` - (Float) Value at which the color should change.
						* `unit` - (String) Query result value interpretation.
						  * Constraints: Allowable values are: `unspecified`, `number`, `percent`, `microseconds`, `milliseconds`, `seconds`, `bytes`, `kbytes`, `mbytes`, `gbytes`, `bytes_iec`, `kibytes`, `mibytes`, `gibytes`, `eur_cents`, `eur`, `usd_cents`, `usd`.
					* `horizontal_bar_chart` - (List) Horizontal bar chart widget.
					Nested schema for **horizontal_bar_chart**:
						* `color_scheme` - (String) Color scheme name.
						  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `colors_by` - (List) Coloring mode.
						Nested schema for **colors_by**:
							* `aggregation` - (List) Each aggregation will have different color and stack color will be derived from aggregation color.
							Nested schema for **aggregation**:
							* `group_by` - (List) Each group will have different color and stack color will be derived from group color.
							Nested schema for **group_by**:
							* `stack` - (List) Each stack will have the same color across all groups.
							Nested schema for **stack**:
						* `data_mode_type` - (String) Data mode type.
						  * Constraints: Allowable values are: `high_unspecified`, `archive`.
						* `display_on_bar` - (Boolean) Whether to display values on the bars.
						* `group_name_template` - (String) Template for bar labels.
						  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `max_bars_per_chart` - (Integer) Maximum number of bars to display in the chart.
						* `query` - (List) Data source specific query, defines from where and how to fetch the data.
						Nested schema for **query**:
							* `dataprime` - (List) Dataprime specific query.
							Nested schema for **dataprime**:
								* `dataprime_query` - (List) Dataprime query.
								Nested schema for **dataprime_query**:
									* `text` - (String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `filters` - (List) Extra filter on top of the Dataprime query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `logs` - (List) Extra filtering on top of the Lucene query.
									Nested schema for **logs**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `operator` - (List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
									* `metrics` - (List) Filtering to be applied to query results.
									Nested schema for **metrics**:
										* `label` - (String) Label associated with the metric.
										  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
										* `operator` - (List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `group_names` - (List) Fields to group by.
								  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `stacked_group_name` - (String) Field to stack by.
								  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `logs` - (List) Logs specific query.
							Nested schema for **logs**:
								* `aggregation` - (List) Aggregations.
								Nested schema for **aggregation**:
									* `average` - (List) Calculate average value of log field.
									Nested schema for **average**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `count` - (List) Count the number of entries.
									Nested schema for **count**:
									* `count_distinct` - (List) Count the number of distinct values of log field.
									Nested schema for **count_distinct**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `max` - (List) Calculate maximum value of log field.
									Nested schema for **max**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `min` - (List) Calculate minimum value of log field.
									Nested schema for **min**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `percentile` - (List) Calculate percentile value of log field.
									Nested schema for **percentile**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `percent` - (Float) Value in range (0, 100].
									* `sum` - (List) Sum values of log field.
									Nested schema for **sum**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
								* `filters` - (List) Extra filter on top of the Lucene query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `observation_field` - (List) Field to count distinct values of.
									Nested schema for **observation_field**:
										* `keypath` - (List) Path within the dataset scope.
										  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
										* `scope` - (String) Scope of the dataset.
										  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `operator` - (List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `group_names_fields` - (List) Fields to group by.
								  * Constraints: The maximum length is `2` items. The minimum length is `1` item.
								Nested schema for **group_names_fields**:
									* `keypath` - (List) Path within the dataset scope.
									  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
									* `scope` - (String) Scope of the dataset.
									  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
								* `lucene_query` - (List) Lucene query.
								Nested schema for **lucene_query**:
									* `value` - (String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `stacked_group_name_field` - (List) Field to count distinct values of.
								Nested schema for **stacked_group_name_field**:
									* `keypath` - (List) Path within the dataset scope.
									  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
									* `scope` - (String) Scope of the dataset.
									  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
							* `metrics` - (List) Metrics specific query.
							Nested schema for **metrics**:
								* `filters` - (List) Extra filter on top of the PromQL query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `label` - (String) Label associated with the metric.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
									* `operator` - (List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `group_names` - (List) Labels to group by.
								  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `2` items. The minimum length is `1` item.
								* `promql_query` - (List) PromQL query.
								Nested schema for **promql_query**:
									* `value` - (String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `stacked_group_name` - (String) Label to stack by.
								  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `scale_type` - (String) Scale type.
						  * Constraints: Allowable values are: `unspecified`, `linear`, `logarithmic`.
						* `sort_by` - (String) Sorting mode.
						  * Constraints: Allowable values are: `unspecified`, `value`, `name`.
						* `stack_definition` - (List) Stack definition.
						Nested schema for **stack_definition**:
							* `max_slices_per_bar` - (Integer) Maximum number of slices per bar.
							* `stack_name_template` - (String) Template for stack slice label.
							  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `unit` - (String) Unit of the data.
						  * Constraints: Allowable values are: `unspecified`, `microseconds`, `milliseconds`, `seconds`, `bytes`, `kbytes`, `mbytes`, `gbytes`, `bytes_iec`, `kibytes`, `mibytes`, `gibytes`, `eur_cents`, `eur`, `usd_cents`, `usd`.
						* `y_axis_view_by` - (List) Y-axis view mode.
						Nested schema for **y_axis_view_by**:
							* `category` - (List) View by category.
							Nested schema for **category**:
							* `value` - (List) View by value.
							Nested schema for **value**:
					* `line_chart` - (List) Line chart widget.
					Nested schema for **line_chart**:
						* `legend` - (List) Legend configuration.
						Nested schema for **legend**:
							* `columns` - (List) The columns to show in the legend.
							  * Constraints: Allowable list items are: `unspecified`, `min`, `max`, `sum`, `avg`, `last`, `name`. The maximum length is `4096` items. The minimum length is `0` items.
							* `group_by_query` - (Boolean) Whether to group by the query or not.
							* `is_visible` - (Boolean) Whether to show the legend or not.
						* `query_definitions` - (List) Query definitions.
						  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
						Nested schema for **query_definitions**:
							* `color_scheme` - (String) Color scheme for the series.
							  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `data_mode_type` - (String) Data mode type.
							  * Constraints: Allowable values are: `high_unspecified`, `archive`.
							* `id` - (String) Unique identifier of the query within the widget.
							  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
							* `is_visible` - (Boolean) Whether data for this query should be visible on the chart.
							* `name` - (String) Query name.
							  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `query` - (List) Data source specific query, defines from where and how to fetch the data.
							Nested schema for **query**:
								* `dataprime` - (List) Dataprime language based query.
								Nested schema for **dataprime**:
									* `dataprime_query` - (List) Dataprime query.
									Nested schema for **dataprime_query**:
										* `text` - (String) The query string.
										  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
									* `filters` - (List) Filters to be applied to query results.
									  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
									Nested schema for **filters**:
										* `logs` - (List) Extra filtering on top of the Lucene query.
										Nested schema for **logs**:
											* `observation_field` - (List) Field to count distinct values of.
											Nested schema for **observation_field**:
												* `keypath` - (List) Path within the dataset scope.
												  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
												* `scope` - (String) Scope of the dataset.
												  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
											* `operator` - (List) Operator to use for filtering the logs.
											Nested schema for **operator**:
												* `equals` - (List) Equality comparison.
												Nested schema for **equals**:
													* `selection` - (List) Selection criteria for the equality comparison.
													Nested schema for **selection**:
														* `all` - (List) Represents a selection of all values.
														Nested schema for **all**:
														* `list` - (List) Represents a selection from a list of values.
														Nested schema for **list**:
															* `values` - (List) List of values for the selection.
															  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
												* `not_equals` - (List) Non-equality comparison.
												Nested schema for **not_equals**:
													* `selection` - (List) Selection criteria for the non-equality comparison.
													Nested schema for **selection**:
														* `list` - (List) Represents a selection from a list of values.
														Nested schema for **list**:
															* `values` - (List) List of values for the selection.
															  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `metrics` - (List) Filtering to be applied to query results.
										Nested schema for **metrics**:
											* `label` - (String) Label associated with the metric.
											  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
											* `operator` - (List) Operator to use for filtering the logs.
											Nested schema for **operator**:
												* `equals` - (List) Equality comparison.
												Nested schema for **equals**:
													* `selection` - (List) Selection criteria for the equality comparison.
													Nested schema for **selection**:
														* `all` - (List) Represents a selection of all values.
														Nested schema for **all**:
														* `list` - (List) Represents a selection from a list of values.
														Nested schema for **list**:
															* `values` - (List) List of values for the selection.
															  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
												* `not_equals` - (List) Non-equality comparison.
												Nested schema for **not_equals**:
													* `selection` - (List) Selection criteria for the non-equality comparison.
													Nested schema for **selection**:
														* `list` - (List) Represents a selection from a list of values.
														Nested schema for **list**:
															* `values` - (List) List of values for the selection.
															  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `logs` - (List) Logs specific query.
								Nested schema for **logs**:
									* `aggregations` - (List) Aggregations.
									  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
									Nested schema for **aggregations**:
										* `average` - (List) Calculate average value of log field.
										Nested schema for **average**:
											* `observation_field` - (List) Field to count distinct values of.
											Nested schema for **observation_field**:
												* `keypath` - (List) Path within the dataset scope.
												  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
												* `scope` - (String) Scope of the dataset.
												  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `count` - (List) Count the number of entries.
										Nested schema for **count**:
										* `count_distinct` - (List) Count the number of distinct values of log field.
										Nested schema for **count_distinct**:
											* `observation_field` - (List) Field to count distinct values of.
											Nested schema for **observation_field**:
												* `keypath` - (List) Path within the dataset scope.
												  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
												* `scope` - (String) Scope of the dataset.
												  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `max` - (List) Calculate maximum value of log field.
										Nested schema for **max**:
											* `observation_field` - (List) Field to count distinct values of.
											Nested schema for **observation_field**:
												* `keypath` - (List) Path within the dataset scope.
												  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
												* `scope` - (String) Scope of the dataset.
												  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `min` - (List) Calculate minimum value of log field.
										Nested schema for **min**:
											* `observation_field` - (List) Field to count distinct values of.
											Nested schema for **observation_field**:
												* `keypath` - (List) Path within the dataset scope.
												  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
												* `scope` - (String) Scope of the dataset.
												  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `percentile` - (List) Calculate percentile value of log field.
										Nested schema for **percentile**:
											* `observation_field` - (List) Field to count distinct values of.
											Nested schema for **observation_field**:
												* `keypath` - (List) Path within the dataset scope.
												  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
												* `scope` - (String) Scope of the dataset.
												  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
											* `percent` - (Float) Value in range (0, 100].
										* `sum` - (List) Sum values of log field.
										Nested schema for **sum**:
											* `observation_field` - (List) Field to count distinct values of.
											Nested schema for **observation_field**:
												* `keypath` - (List) Path within the dataset scope.
												  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
												* `scope` - (String) Scope of the dataset.
												  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `filters` - (List) Extra filtering on top of the Lucene query.
									  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
									Nested schema for **filters**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `operator` - (List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
									* `group_by` - (List) Group by fields (deprecated).
									  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
									* `group_bys` - (List) Group by fields.
									  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
									Nested schema for **group_bys**:
										* `keypath` - (List) Path within the dataset scope.
										  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
										* `scope` - (String) Scope of the dataset.
										  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `lucene_query` - (List) Lucene query.
									Nested schema for **lucene_query**:
										* `value` - (String) The query string.
										  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `metrics` - (List) Metrics specific query.
								Nested schema for **metrics**:
									* `filters` - (List) Filtering to be applied to query results.
									  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
									Nested schema for **filters**:
										* `label` - (String) Label associated with the metric.
										  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
										* `operator` - (List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
									* `promql_query` - (List) PromQL query.
									Nested schema for **promql_query**:
										* `value` - (String) The query string.
										  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `resolution` - (List) Resolution of the data.
							Nested schema for **resolution**:
								* `buckets_presented` - (Integer) Maximum number of data points to fetch.
								* `interval` - (String) Interval between data points.
								  * Constraints: The maximum length is `10` characters. The minimum length is `2` characters. The value must match regular expression `/^[0-9]+[smhdw]?$/`.
							* `scale_type` - (String) Scale type.
							  * Constraints: Allowable values are: `unspecified`, `linear`, `logarithmic`.
							* `series_count_limit` - (String) Maximum number of series to display.
							  * Constraints: The maximum length is `19` characters. The minimum length is `1` character. The value must match regular expression `/^-?\\d{1,19}$/`.
							* `series_name_template` - (String) Template for series name in legend and tooltip.
							  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `unit` - (String) Unit of the data.
							  * Constraints: Allowable values are: `unspecified`, `microseconds`, `milliseconds`, `seconds`, `bytes`, `kbytes`, `mbytes`, `gbytes`, `bytes_iec`, `kibytes`, `mibytes`, `gibytes`, `eur_cents`, `eur`, `usd_cents`, `usd`.
						* `stacked_line` - (String) Stacked lines.
						  * Constraints: Allowable values are: `unspecified`, `absolute`, `relative`.
						* `tooltip` - (List) Tooltip configuration.
						Nested schema for **tooltip**:
							* `show_labels` - (Boolean) Whether to show labels in the tooltip.
							* `type` - (String) Tooltip type.
							  * Constraints: Allowable values are: `unspecified`, `all`, `single`.
					* `markdown` - (List) Markdown widget.
					Nested schema for **markdown**:
						* `markdown_text` - (String) Markdown text to render.
						  * Constraints: The maximum length is `10000` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `tooltip_text` - (String) Tooltip text on hover.
						  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
					* `pie_chart` - (List) Pie chart widget.
					Nested schema for **pie_chart**:
						* `color_scheme` - (String) Color scheme name.
						  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `data_mode_type` - (String) Data mode type.
						  * Constraints: Allowable values are: `high_unspecified`, `archive`.
						* `group_name_template` - (String) Template for group labels.
						  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `label_definition` - (List) Label settings.
						Nested schema for **label_definition**:
							* `is_visible` - (Boolean) Controls whether to show the label.
							* `label_source` - (String) Source of the label.
							  * Constraints: Allowable values are: `unspecified`, `inner`, `stack`.
							* `show_name` - (Boolean) Controls whether to show the name.
							* `show_percentage` - (Boolean) Controls whether to show the percentage.
							* `show_value` - (Boolean) Controls whether to show the value.
						* `max_slices_per_chart` - (Integer) Maximum number of slices to display in the chart.
						* `min_slice_percentage` - (Integer) Minimum percentage of a slice to be displayed.
						* `query` - (List) Data source specific query, defines from where and how to fetch the data.
						Nested schema for **query**:
							* `dataprime` - (List) Query based on Dataprime language.
							Nested schema for **dataprime**:
								* `dataprime_query` - (List) Dataprime query.
								Nested schema for **dataprime_query**:
									* `text` - (String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `filters` - (List) Extra filters on top of Dataprime query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `logs` - (List) Extra filtering on top of the Lucene query.
									Nested schema for **logs**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `operator` - (List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
									* `metrics` - (List) Filtering to be applied to query results.
									Nested schema for **metrics**:
										* `label` - (String) Label associated with the metric.
										  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
										* `operator` - (List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `group_names` - (List) Fields to group by.
								  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `2` items. The minimum length is `1` item.
								* `stacked_group_name` - (String) Field to stack by.
								  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `logs` - (List) Logs specific query.
							Nested schema for **logs**:
								* `aggregation` - (List) Aggregations.
								Nested schema for **aggregation**:
									* `average` - (List) Calculate average value of log field.
									Nested schema for **average**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `count` - (List) Count the number of entries.
									Nested schema for **count**:
									* `count_distinct` - (List) Count the number of distinct values of log field.
									Nested schema for **count_distinct**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `max` - (List) Calculate maximum value of log field.
									Nested schema for **max**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `min` - (List) Calculate minimum value of log field.
									Nested schema for **min**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `percentile` - (List) Calculate percentile value of log field.
									Nested schema for **percentile**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `percent` - (Float) Value in range (0, 100].
									* `sum` - (List) Sum values of log field.
									Nested schema for **sum**:
										* `observation_field` - (List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
								* `filters` - (List) Extra filters on top of Lucene query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `observation_field` - (List) Field to count distinct values of.
									Nested schema for **observation_field**:
										* `keypath` - (List) Path within the dataset scope.
										  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
										* `scope` - (String) Scope of the dataset.
										  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `operator` - (List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `group_names_fields` - (List) Fields to group by.
								  * Constraints: The maximum length is `2` items. The minimum length is `1` item.
								Nested schema for **group_names_fields**:
									* `keypath` - (List) Path within the dataset scope.
									  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
									* `scope` - (String) Scope of the dataset.
									  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
								* `lucene_query` - (List) Lucene query.
								Nested schema for **lucene_query**:
									* `value` - (String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `stacked_group_name_field` - (List) Field to count distinct values of.
								Nested schema for **stacked_group_name_field**:
									* `keypath` - (List) Path within the dataset scope.
									  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
									* `scope` - (String) Scope of the dataset.
									  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
							* `metrics` - (List) Metrics specific query.
							Nested schema for **metrics**:
								* `filters` - (List) Extra filters on top of PromQL query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `label` - (String) Label associated with the metric.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
									* `operator` - (List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `group_names` - (List) Fields to group by.
								  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `2` items. The minimum length is `1` item.
								* `promql_query` - (List) PromQL query.
								Nested schema for **promql_query**:
									* `value` - (String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `stacked_group_name` - (String) Field to stack by.
								  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `show_legend` - (Boolean) Controls whether to show the legend.
						* `stack_definition` - (List) Stack definition.
						Nested schema for **stack_definition**:
							* `max_slices_per_stack` - (Integer) Maximum number of slices per stack.
							* `stack_name_template` - (String) Template for stack labels.
							  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `unit` - (String) Unit of the data.
						  * Constraints: Allowable values are: `unspecified`, `microseconds`, `milliseconds`, `seconds`, `bytes`, `kbytes`, `mbytes`, `gbytes`, `bytes_iec`, `kibytes`, `mibytes`, `gibytes`, `eur_cents`, `eur`, `usd_cents`, `usd`.
				* `description` - (String) Widget description.
				  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `href` - (String) Widget identifier within the dashboard.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `id` - (List) Unique identifier of the folder containing the dashboard.
				Nested schema for **id**:
					* `value` - (String) The UUID value.
					  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
				* `title` - (String) Widget title.
				  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `updated_at` - (String) Last update timestamp.

* `name` - (String) Display name of the dashboard.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

* `relative_time_frame` - (String) Relative time frame specifying a duration from the current time.
  * Constraints: The maximum length is `10` characters. The minimum length is `2` characters. The value must match regular expression `/^[0-9]+[smhdw]?$/`.

* `two_minutes` - (List) Auto refresh interval is set to two minutes.
Nested schema for **two_minutes**:

* `variables` - (List) List of variables that can be used within the dashboard for dynamic content.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **variables**:
	* `definition` - (List) Definition.
	Nested schema for **definition**:
		* `multi_select` - (List) Multi-select value.
		Nested schema for **multi_select**:
			* `selection` - (List) State of what is currently selected.
			Nested schema for **selection**:
				* `all` - (List) All values are selected, usually translated to wildcard (*).
				Nested schema for **all**:
				* `list` - (List) Specific values are selected.
				Nested schema for **list**:
					* `values` - (List) Selected values.
					  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
			* `source` - (List) Variable value source.
			Nested schema for **source**:
				* `constant_list` - (List) List of constant values.
				Nested schema for **constant_list**:
					* `values` - (List) List of constant values.
					  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
				* `logs_path` - (List) Unique values for a given logs path.
				Nested schema for **logs_path**:
					* `observation_field` - (List) Field to count distinct values of.
					Nested schema for **observation_field**:
						* `keypath` - (List) Path within the dataset scope.
						  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
						* `scope` - (String) Scope of the dataset.
						  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
				* `metric_label` - (List) Unique values for a given metric label.
				Nested schema for **metric_label**:
					* `label` - (String) Metric label to source unique values from.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
					* `metric_name` - (String) Metric name to source unique values from.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
			* `values_order_direction` - (String) The direction of the order: ascending or descending.
			  * Constraints: Allowable values are: `unspecified`, `asc`, `desc`.
	* `display_name` - (String) Name used in variable UI.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `name` - (String) Name of the variable which can be used in templates.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.


---
layout: "ibm"
page_title: "IBM : ibm_logs_dashboard"
description: |-
  Manages logs_dashboard.
subcategory: "Cloud Logs"
---


# ibm_logs_dashboard

Create, update, and delete logs_dashboards with this resource.

## Example Usage

```hcl
resource "ibm_logs_dashboard" "logs_dashboard_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  name        = "example-dashboard"
  description = "example dashboard description"
  layout {
    sections {
      id {
        value = "b9ca2f71-7d7c-10fb-1a08-c78912705095"
      }
      rows {
        id {
          value = "70b12716-cb18-f933-5a89-3061734eaa2f"
        }
        appearance {
          height = 19
        }
        widgets {
          id {
            value = "6118b86d-860c-c2cb-0cdf-effd62e9f331"
          }
          title       = "test"
          description = "test"
          definition {
            line_chart {
              legend {
                is_visible     = true
                group_by_query = true
              }
              tooltip {
                show_labels = false
                type        = "all"
              }
              query_definitions {
                id           = "13139dad-3d45-16e1-fce2-03517daa71c4"
                color_scheme = "cold"
                name         = "Query 1"
                is_visible   = true
                scale_type   = "linear"
                resolution {
                  buckets_presented = 96
                }
                series_count_limit = 20
                query {
                  logs {
                    group_by = []

                    aggregations {
                      min {
                        observation_field {
                          keypath = [
                            "timestamp",
                          ]
                          scope = "metadata"
                        }
                      }
                    }

                    group_bys {
                      keypath = [
                        "severity",
                      ]
                      scope = "metadata"
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
  filters {
    source {
      logs {
        operator {
          equals {
            selection {
              list {}
            }
          }
        }
        observation_field {
          keypath = ["applicationname"]
          scope   = "label"
        }
      }
    }
    enabled   = true
    collapsed = false
  }
  filters {
    source {
      logs {
        # field = "field"
        operator {
          equals {
            selection {
              all {}
            }
          }
        }
        observation_field {
          keypath = ["subsystemname"]
          scope   = "label"
        }
      }
    }
    enabled   = true
    collapsed = false
  }
  relative_time_frame = "900s"
}
```
## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String)  Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.
* `absolute_time_frame` - (Optional, List) Absolute time frame specifying a fixed start and end time.
Nested schema for **absolute_time_frame**:
	* `from` - (Optional, String) From is the start of the time frame.
	* `to` - (Optional, String) To is the end of the time frame.
* `annotations` - (Optional, List) List of annotations that can be applied to the dashboard's visual elements.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **annotations**:
	* `enabled` - (Required, Boolean) Whether the annotation is enabled.
	* `href` - (Optional, String) Unique identifier within the dashboard.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
	* `id` - (Required, String) Unique identifier within the dashboard.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
	* `name` - (Required, String) Name of the annotation.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `source` - (Required, List) Source of the annotation events.
	Nested schema for **source**:
		* `logs` - (Optional, List) Logs source.
		Nested schema for **logs**:
			* `label_fields` - (Optional, List) Labels to display in the annotation.
			  * Constraints: The maximum length is `10` items. The minimum length is `0` items.
			Nested schema for **label_fields**:
				* `keypath` - (Optional, List) Path within the dataset scope.
				  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
				* `scope` - (Optional, String) Scope of the dataset.
				  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
			* `lucene_query` - (Required, List) Lucene query.
			Nested schema for **lucene_query**:
				* `value` - (Optional, String) The Lucene query string.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
			* `message_template` - (Optional, String) Template for the annotation message.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
			* `strategy` - (Required, List) Strategy for turning logs data into annotations.
			Nested schema for **strategy**:
				* `duration` - (Optional, List) Event start timestamp and duration are extracted from the log entry.
				Nested schema for **duration**:
					* `duration_field` - (Required, List) Field to count distinct values of.
					Nested schema for **duration_field**:
						* `keypath` - (Optional, List) Path within the dataset scope.
						  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
						* `scope` - (Optional, String) Scope of the dataset.
						  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
					* `start_timestamp_field` - (Required, List) Field to count distinct values of.
					Nested schema for **start_timestamp_field**:
						* `keypath` - (Optional, List) Path within the dataset scope.
						  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
						* `scope` - (Optional, String) Scope of the dataset.
						  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
				* `instant` - (Optional, List) Event timestamp is extracted from the log entry.
				Nested schema for **instant**:
					* `timestamp_field` - (Required, List) Field to count distinct values of.
					Nested schema for **timestamp_field**:
						* `keypath` - (Optional, List) Path within the dataset scope.
						  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
						* `scope` - (Optional, String) Scope of the dataset.
						  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
				* `range` - (Optional, List) Event start and end timestamps are extracted from the log entry.
				Nested schema for **range**:
					* `end_timestamp_field` - (Required, List) Field to count distinct values of.
					Nested schema for **end_timestamp_field**:
						* `keypath` - (Optional, List) Path within the dataset scope.
						  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
						* `scope` - (Optional, String) Scope of the dataset.
						  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
					* `start_timestamp_field` - (Required, List) Field to count distinct values of.
					Nested schema for **start_timestamp_field**:
						* `keypath` - (Optional, List) Path within the dataset scope.
						  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
						* `scope` - (Optional, String) Scope of the dataset.
						  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
		* `metrics` - (Optional, List) Metrics source.
		Nested schema for **metrics**:
			* `labels` - (Optional, List) Labels to display in the annotation.
			  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
			* `message_template` - (Optional, String) Template for the annotation message.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
			* `promql_query` - (Optional, List) PromQL query.
			Nested schema for **promql_query**:
				* `value` - (Optional, String) The PromQL query string.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
			* `strategy` - (Optional, List) Strategy for turning metrics data into annotations.
			Nested schema for **strategy**:
				* `start_time_metric` - (Optional, List) Take first data point and use its value as annotation timestamp (instead of point own timestamp).
				Nested schema for **start_time_metric**:
* `description` - (Optional, String) Brief description or summary of the dashboard's purpose or content.
  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
* `false` - (Optional, List) Auto refresh interval is set to off.
Nested schema for **false**:
* `filters` - (Optional, List) List of filters that can be applied to the dashboard's data.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **filters**:
	* `collapsed` - (Optional, Boolean) Indicates if the filter's UI representation should be collapsed or expanded.
	* `enabled` - (Optional, Boolean) Indicates if the filter is currently enabled or not.
	* `source` - (Optional, List) Filters to be applied to query results.
	Nested schema for **source**:
		* `logs` - (Optional, List) Extra filtering on top of the Lucene query.
		Nested schema for **logs**:
			* `observation_field` - (Optional, List) Field to count distinct values of.
			Nested schema for **observation_field**:
				* `keypath` - (Optional, List) Path within the dataset scope.
				  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
				* `scope` - (Optional, String) Scope of the dataset.
				  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
			* `operator` - (Optional, List) Operator to use for filtering the logs.
			Nested schema for **operator**:
				* `equals` - (Optional, List) Equality comparison.
				Nested schema for **equals**:
					* `selection` - (Optional, List) Selection criteria for the equality comparison.
					Nested schema for **selection**:
						* `all` - (Optional, List) Represents a selection of all values.
						Nested schema for **all**:
						* `list` - (Optional, List) Represents a selection from a list of values.
						Nested schema for **list**:
							* `values` - (Optional, List) List of values for the selection.
							  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
				* `not_equals` - (Optional, List) Non-equality comparison.
				Nested schema for **not_equals**:
					* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
					Nested schema for **selection**:
						* `list` - (Optional, List) Represents a selection from a list of values.
						Nested schema for **list**:
							* `values` - (Optional, List) List of values for the selection.
							  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
		* `metrics` - (Optional, List) Filtering to be applied to query results.
		Nested schema for **metrics**:
			* `label` - (Optional, String) Label associated with the metric.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
			* `operator` - (Optional, List) Operator to use for filtering the logs.
			Nested schema for **operator**:
				* `equals` - (Optional, List) Equality comparison.
				Nested schema for **equals**:
					* `selection` - (Optional, List) Selection criteria for the equality comparison.
					Nested schema for **selection**:
						* `all` - (Optional, List) Represents a selection of all values.
						Nested schema for **all**:
						* `list` - (Optional, List) Represents a selection from a list of values.
						Nested schema for **list**:
							* `values` - (Optional, List) List of values for the selection.
							  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
				* `not_equals` - (Optional, List) Non-equality comparison.
				Nested schema for **not_equals**:
					* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
					Nested schema for **selection**:
						* `list` - (Optional, List) Represents a selection from a list of values.
						Nested schema for **list**:
							* `values` - (Optional, List) List of values for the selection.
							  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
* `five_minutes` - (Optional, List) Auto refresh interval is set to five minutes.
Nested schema for **five_minutes**:
* `folder_id` - (Optional, List) Unique identifier of the folder containing the dashboard.
Nested schema for **folder_id**:
	* `value` - (Required, String) The UUID value.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `folder_path` - (Optional, List) Path of the folder containing the dashboard.
Nested schema for **folder_path**:
	* `segments` - (Optional, List) The segments of the folder path.
	  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
* `href` - (Optional, String) Unique identifier for the dashboard.
  * Constraints: The maximum length is `21` characters. The minimum length is `21` characters. The value must match regular expression `/^[a-zA-Z0-9]{21}$/`.
* `layout` - (Required, List) Layout configuration for the dashboard's visual elements.
Nested schema for **layout**:
	* `sections` - (Optional, List) The sections of the layout.
	  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
	Nested schema for **sections**:
		* `href` - (Optional, String) The unique identifier of the section within the layout.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
		* `id` - (Required, List) Unique identifier of the folder containing the dashboard.
		Nested schema for **id**:
			* `value` - (Required, String) The UUID value.
			  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
		* `rows` - (Optional, List) The rows of the section.
		  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
		Nested schema for **rows**:
			* `appearance` - (Required, List) The appearance of the row, such as height.
			Nested schema for **appearance**:
				* `height` - (Required, Integer) The height of the row.
			* `href` - (Optional, String) The unique identifier of the row within the layout.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
			* `id` - (Required, List) Unique identifier of the folder containing the dashboard.
			Nested schema for **id**:
				* `value` - (Required, String) The UUID value.
				  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
			* `widgets` - (Optional, List) The widgets of the row.
			  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
			Nested schema for **widgets**:
				* `created_at` - (Optional, String) Creation timestamp.
				* `definition` - (Required, List) Widget definition, contains the widget type and its configuration.
				Nested schema for **definition**:
					* `bar_chart` - (Optional, List) Bar chart widget.
					Nested schema for **bar_chart**:
						* `color_scheme` - (Required, String) Supported vaues: classic, severity, cold, negative, green, red, blue.
						  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `colors_by` - (Required, List) Coloring mode.
						Nested schema for **colors_by**:
							* `aggregation` - (Optional, List) Each aggregation will have different color and stack color will be derived from aggregation color.
							Nested schema for **aggregation**:
							* `group_by` - (Optional, List) Each group will have different color and stack color will be derived from group color.
							Nested schema for **group_by**:
							* `stack` - (Optional, List) Each stack will have the same color across all groups.
							Nested schema for **stack**:
						* `data_mode_type` - (Optional, String) Data mode type.
						  * Constraints: Allowable values are: `high_unspecified`, `archive`.
						* `group_name_template` - (Required, String) Template for bar labels.
						  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `max_bars_per_chart` - (Required, Integer) Maximum number of bars to present in the chart.
						* `query` - (Required, List) Data source specific query, defines from where and how to fetch the data.
						Nested schema for **query**:
							* `dataprime` - (Optional, List) Query based on Dataprime language.
							Nested schema for **dataprime**:
								* `dataprime_query` - (Required, List) Dataprime query.
								Nested schema for **dataprime_query**:
									* `text` - (Optional, String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `filters` - (Optional, List) Extra filter on top of the Dataprime query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `logs` - (Optional, List) Extra filtering on top of the Lucene query.
									Nested schema for **logs**:
										* `observation_field` - (Optional, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `operator` - (Optional, List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (Optional, List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (Optional, List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (Optional, List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (Optional, List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
									* `metrics` - (Optional, List) Filtering to be applied to query results.
									Nested schema for **metrics**:
										* `label` - (Optional, String) Label associated with the metric.
										  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
										* `operator` - (Optional, List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (Optional, List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (Optional, List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (Optional, List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (Optional, List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `group_names` - (Optional, List) Fields to group by.
								  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `2` items. The minimum length is `1` item.
								* `stacked_group_name` - (Optional, String) Field to stack by.
								  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `logs` - (Optional, List) Logs specific query.
							Nested schema for **logs**:
								* `aggregation` - (Required, List) Aggregations.
								Nested schema for **aggregation**:
									* `average` - (Optional, List) Calculate average value of log field.
									Nested schema for **average**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `count` - (Optional, List) Count the number of entries.
									Nested schema for **count**:
									* `count_distinct` - (Optional, List) Count the number of distinct values of log field.
									Nested schema for **count_distinct**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `max` - (Optional, List) Calculate maximum value of log field.
									Nested schema for **max**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `min` - (Optional, List) Calculate minimum value of log field.
									Nested schema for **min**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `percentile` - (Optional, List) Calculate percentile value of log field.
									Nested schema for **percentile**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `percent` - (Required, Float) Value in range (0, 100].
									* `sum` - (Optional, List) Sum values of log field.
									Nested schema for **sum**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
								* `filters` - (Optional, List) Extra filter on top of Lucene query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `observation_field` - (Optional, List) Field to count distinct values of.
									Nested schema for **observation_field**:
										* `keypath` - (Optional, List) Path within the dataset scope.
										  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
										* `scope` - (Optional, String) Scope of the dataset.
										  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `operator` - (Optional, List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (Optional, List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (Optional, List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (Optional, List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (Optional, List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `group_names_fields` - (Optional, List) Fiel to group by.
								  * Constraints: The maximum length is `2` items. The minimum length is `1` item.
								Nested schema for **group_names_fields**:
									* `keypath` - (Optional, List) Path within the dataset scope.
									  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
									* `scope` - (Optional, String) Scope of the dataset.
									  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
								* `lucene_query` - (Optional, List) Lucene query.
								Nested schema for **lucene_query**:
									* `value` - (Optional, String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `stacked_group_name_field` - (Optional, List) Field to count distinct values of.
								Nested schema for **stacked_group_name_field**:
									* `keypath` - (Optional, List) Path within the dataset scope.
									  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
									* `scope` - (Optional, String) Scope of the dataset.
									  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
							* `metrics` - (Optional, List) Metrics specific query.
							Nested schema for **metrics**:
								* `filters` - (Optional, List) Extra filter on top of the PromQL query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `label` - (Optional, String) Label associated with the metric.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
									* `operator` - (Optional, List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (Optional, List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (Optional, List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (Optional, List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (Optional, List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `group_names` - (Optional, List) Labels to group by.
								  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `2` items. The minimum length is `1` item.
								* `promql_query` - (Optional, List) PromQL query.
								Nested schema for **promql_query**:
									* `value` - (Optional, String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `stacked_group_name` - (Optional, String) Label to stack by.
								  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `scale_type` - (Required, String) Scale type.
						  * Constraints: Allowable values are: `unspecified`, `linear`, `logarithmic`.
						* `sort_by` - (Required, String) Sorting mode.
						  * Constraints: Allowable values are: `unspecified`, `value`, `name`.
						* `stack_definition` - (Required, List) Stack definition.
						Nested schema for **stack_definition**:
							* `max_slices_per_bar` - (Optional, Integer) Maximum number of slices per bar.
							* `stack_name_template` - (Optional, String) Template for stack slice label.
							  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `unit` - (Required, String) Unit of the data.
						  * Constraints: Allowable values are: `unspecified`, `microseconds`, `milliseconds`, `seconds`, `bytes`, `kbytes`, `mbytes`, `gbytes`, `bytes_iec`, `kibytes`, `mibytes`, `gibytes`, `eur_cents`, `eur`, `usd_cents`, `usd`.
						* `x_axis` - (Required, List) X axis mode.
						Nested schema for **x_axis**:
							* `time` - (Optional, List) Time based axis.
							Nested schema for **time**:
								* `buckets_presented` - (Optional, Integer) Buckets presented.
								* `interval` - (Optional, String) Time interval.
								  * Constraints: The maximum length is `10` characters. The minimum length is `2` characters. The value must match regular expression `/^[0-9]+[smhdw]?$/`.
							* `value` - (Optional, List) Categorical axis.
							Nested schema for **value**:
					* `data_table` - (Optional, List) Data table widget.
					Nested schema for **data_table**:
						* `columns` - (Optional, List) Columns to display, their order and width.
						  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
						Nested schema for **columns**:
							* `field` - (Required, String) References a field in result set. In case of aggregation, it references the aggregation identifier.
							  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `width` - (Optional, Integer) Column width.
						* `data_mode_type` - (Optional, String) Data mode type.
						  * Constraints: Allowable values are: `high_unspecified`, `archive`.
						* `order_by` - (Optional, List) Column used for ordering the results.
						Nested schema for **order_by**:
							* `field` - (Optional, String) The field to order by.
							  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `order_direction` - (Optional, String) The direction of the order: ascending or descending.
							  * Constraints: Allowable values are: `unspecified`, `asc`, `desc`.
						* `query` - (Required, List) Data source specific query, defines from where and how to fetch the data.
						Nested schema for **query**:
							* `dataprime` - (Optional, List) Query based on Dataprime language.
							Nested schema for **dataprime**:
								* `dataprime_query` - (Required, List) Dataprime query.
								Nested schema for **dataprime_query**:
									* `text` - (Optional, String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `filters` - (Optional, List) Extra filtering on top of the Dataprime query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `logs` - (Optional, List) Extra filtering on top of the Lucene query.
									Nested schema for **logs**:
										* `observation_field` - (Optional, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `operator` - (Optional, List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (Optional, List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (Optional, List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (Optional, List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (Optional, List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
									* `metrics` - (Optional, List) Filtering to be applied to query results.
									Nested schema for **metrics**:
										* `label` - (Optional, String) Label associated with the metric.
										  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
										* `operator` - (Optional, List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (Optional, List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (Optional, List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (Optional, List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (Optional, List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
							* `logs` - (Optional, List) Logs specific query.
							Nested schema for **logs**:
								* `filters` - (Optional, List) Extra filtering on top of the Lucene query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `observation_field` - (Optional, List) Field to count distinct values of.
									Nested schema for **observation_field**:
										* `keypath` - (Optional, List) Path within the dataset scope.
										  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
										* `scope` - (Optional, String) Scope of the dataset.
										  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `operator` - (Optional, List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (Optional, List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (Optional, List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (Optional, List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (Optional, List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `grouping` - (Optional, List) Grouping and aggregation.
								Nested schema for **grouping**:
									* `aggregations` - (Optional, List) Aggregations.
									  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
									Nested schema for **aggregations**:
										* `aggregation` - (Required, List) Aggregations.
										Nested schema for **aggregation**:
											* `average` - (Optional, List) Calculate average value of log field.
											Nested schema for **average**:
												* `observation_field` - (Required, List) Field to count distinct values of.
												Nested schema for **observation_field**:
													* `keypath` - (Optional, List) Path within the dataset scope.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
													* `scope` - (Optional, String) Scope of the dataset.
													  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
											* `count` - (Optional, List) Count the number of entries.
											Nested schema for **count**:
											* `count_distinct` - (Optional, List) Count the number of distinct values of log field.
											Nested schema for **count_distinct**:
												* `observation_field` - (Required, List) Field to count distinct values of.
												Nested schema for **observation_field**:
													* `keypath` - (Optional, List) Path within the dataset scope.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
													* `scope` - (Optional, String) Scope of the dataset.
													  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
											* `max` - (Optional, List) Calculate maximum value of log field.
											Nested schema for **max**:
												* `observation_field` - (Required, List) Field to count distinct values of.
												Nested schema for **observation_field**:
													* `keypath` - (Optional, List) Path within the dataset scope.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
													* `scope` - (Optional, String) Scope of the dataset.
													  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
											* `min` - (Optional, List) Calculate minimum value of log field.
											Nested schema for **min**:
												* `observation_field` - (Required, List) Field to count distinct values of.
												Nested schema for **observation_field**:
													* `keypath` - (Optional, List) Path within the dataset scope.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
													* `scope` - (Optional, String) Scope of the dataset.
													  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
											* `percentile` - (Optional, List) Calculate percentile value of log field.
											Nested schema for **percentile**:
												* `observation_field` - (Required, List) Field to count distinct values of.
												Nested schema for **observation_field**:
													* `keypath` - (Optional, List) Path within the dataset scope.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
													* `scope` - (Optional, String) Scope of the dataset.
													  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
												* `percent` - (Required, Float) Value in range (0, 100].
											* `sum` - (Optional, List) Sum values of log field.
											Nested schema for **sum**:
												* `observation_field` - (Required, List) Field to count distinct values of.
												Nested schema for **observation_field**:
													* `keypath` - (Optional, List) Path within the dataset scope.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
													* `scope` - (Optional, String) Scope of the dataset.
													  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `id` - (Required, String) Aggregation identifier, must be unique within grouping configuration.
										  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
										* `is_visible` - (Required, Boolean) Whether the aggregation is visible.
										* `name` - (Required, String) Aggregation name, used as column name.
										  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
									* `group_bys` - (Optional, List) Fields to group by.
									  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
									Nested schema for **group_bys**:
										* `keypath` - (Optional, List) Path within the dataset scope.
										  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
										* `scope` - (Optional, String) Scope of the dataset.
										  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
								* `lucene_query` - (Optional, List) Lucene query.
								Nested schema for **lucene_query**:
									* `value` - (Optional, String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `metrics` - (Optional, List) Metrics specific query.
							Nested schema for **metrics**:
								* `filters` - (Optional, List) Extra filtering on top of the PromQL query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `label` - (Optional, String) Label associated with the metric.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
									* `operator` - (Optional, List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (Optional, List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (Optional, List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (Optional, List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (Optional, List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `promql_query` - (Required, List) PromQL query.
								Nested schema for **promql_query**:
									* `value` - (Optional, String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `results_per_page` - (Required, Integer) Number of results per page.
						* `row_style` - (Required, String) Display style for table rows.
						  * Constraints: Allowable values are: `unspecified`, `one_line`, `two_line`, `condensed`, `json`, `list`.
					* `gauge` - (Optional, List) Gauge widget.
					Nested schema for **gauge**:
						* `data_mode_type` - (Optional, String) Data mode type.
						  * Constraints: Allowable values are: `high_unspecified`, `archive`.
						* `max` - (Required, Float) Maximum value of the gauge.
						* `min` - (Required, Float) Minimum value of the gauge.
						* `query` - (Required, List) Data source specific query, defines from where and how to fetch the data.
						Nested schema for **query**:
							* `dataprime` - (Optional, List) Query based on Dataprime language.
							Nested schema for **dataprime**:
								* `dataprime_query` - (Required, List) Dataprime query.
								Nested schema for **dataprime_query**:
									* `text` - (Optional, String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `filters` - (Optional, List) Extra filters applied on top of Dataprime query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `logs` - (Optional, List) Extra filtering on top of the Lucene query.
									Nested schema for **logs**:
										* `observation_field` - (Optional, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `operator` - (Optional, List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (Optional, List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (Optional, List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (Optional, List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (Optional, List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
									* `metrics` - (Optional, List) Filtering to be applied to query results.
									Nested schema for **metrics**:
										* `label` - (Optional, String) Label associated with the metric.
										  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
										* `operator` - (Optional, List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (Optional, List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (Optional, List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (Optional, List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (Optional, List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
							* `logs` - (Optional, List) Logs specific query.
							Nested schema for **logs**:
								* `filters` - (Optional, List) Extra filters applied on top of Lucene query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `observation_field` - (Optional, List) Field to count distinct values of.
									Nested schema for **observation_field**:
										* `keypath` - (Optional, List) Path within the dataset scope.
										  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
										* `scope` - (Optional, String) Scope of the dataset.
										  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `operator` - (Optional, List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (Optional, List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (Optional, List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (Optional, List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (Optional, List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `logs_aggregation` - (Optional, List) Aggregations.
								Nested schema for **logs_aggregation**:
									* `average` - (Optional, List) Calculate average value of log field.
									Nested schema for **average**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `count` - (Optional, List) Count the number of entries.
									Nested schema for **count**:
									* `count_distinct` - (Optional, List) Count the number of distinct values of log field.
									Nested schema for **count_distinct**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `max` - (Optional, List) Calculate maximum value of log field.
									Nested schema for **max**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `min` - (Optional, List) Calculate minimum value of log field.
									Nested schema for **min**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `percentile` - (Optional, List) Calculate percentile value of log field.
									Nested schema for **percentile**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `percent` - (Required, Float) Value in range (0, 100].
									* `sum` - (Optional, List) Sum values of log field.
									Nested schema for **sum**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
								* `lucene_query` - (Optional, List) Lucene query.
								Nested schema for **lucene_query**:
									* `value` - (Optional, String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `metrics` - (Optional, List) Metrics specific query.
							Nested schema for **metrics**:
								* `aggregation` - (Required, String) Aggregation. When AGGREGATION_UNSPECIFIED is selected, widget uses instant query. Otherwise, it uses range query.
								  * Constraints: Allowable values are: `unspecified`, `last`, `min`, `max`, `avg`, `sum`.
								* `filters` - (Optional, List) Extra filters applied on top of PromQL query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `label` - (Optional, String) Label associated with the metric.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
									* `operator` - (Optional, List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (Optional, List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (Optional, List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (Optional, List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (Optional, List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `promql_query` - (Required, List) PromQL query.
								Nested schema for **promql_query**:
									* `value` - (Optional, String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `show_inner_arc` - (Required, Boolean) Show inner arc (styling).
						* `show_outer_arc` - (Required, Boolean) Show outer arc (styling).
						* `threshold_by` - (Required, String) What threshold color should be applied to: value or background.
						  * Constraints: Allowable values are: `unspecified`, `value`, `background`.
						* `thresholds` - (Optional, List) Thresholds for the gauge, values at which the gauge changes color.
						  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
						Nested schema for **thresholds**:
							* `color` - (Required, String) Color.
							  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `from` - (Required, Float) Value at which the color should change.
						* `unit` - (Required, String) Query result value interpretation.
						  * Constraints: Allowable values are: `unspecified`, `number`, `percent`, `microseconds`, `milliseconds`, `seconds`, `bytes`, `kbytes`, `mbytes`, `gbytes`, `bytes_iec`, `kibytes`, `mibytes`, `gibytes`, `eur_cents`, `eur`, `usd_cents`, `usd`.
					* `horizontal_bar_chart` - (Optional, List) Horizontal bar chart widget.
					Nested schema for **horizontal_bar_chart**:
						* `color_scheme` - (Required, String) Color scheme name.
						  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `colors_by` - (Optional, List) Coloring mode.
						Nested schema for **colors_by**:
							* `aggregation` - (Optional, List) Each aggregation will have different color and stack color will be derived from aggregation color.
							Nested schema for **aggregation**:
							* `group_by` - (Optional, List) Each group will have different color and stack color will be derived from group color.
							Nested schema for **group_by**:
							* `stack` - (Optional, List) Each stack will have the same color across all groups.
							Nested schema for **stack**:
						* `data_mode_type` - (Optional, String) Data mode type.
						  * Constraints: Allowable values are: `high_unspecified`, `archive`.
						* `display_on_bar` - (Optional, Boolean) Whether to display values on the bars.
						* `group_name_template` - (Optional, String) Template for bar labels.
						  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `max_bars_per_chart` - (Optional, Integer) Maximum number of bars to display in the chart.
						* `query` - (Optional, List) Data source specific query, defines from where and how to fetch the data.
						Nested schema for **query**:
							* `dataprime` - (Optional, List) Dataprime specific query.
							Nested schema for **dataprime**:
								* `dataprime_query` - (Optional, List) Dataprime query.
								Nested schema for **dataprime_query**:
									* `text` - (Optional, String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `filters` - (Optional, List) Extra filter on top of the Dataprime query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `logs` - (Optional, List) Extra filtering on top of the Lucene query.
									Nested schema for **logs**:
										* `observation_field` - (Optional, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `operator` - (Optional, List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (Optional, List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (Optional, List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (Optional, List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (Optional, List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
									* `metrics` - (Optional, List) Filtering to be applied to query results.
									Nested schema for **metrics**:
										* `label` - (Optional, String) Label associated with the metric.
										  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
										* `operator` - (Optional, List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (Optional, List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (Optional, List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (Optional, List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (Optional, List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `group_names` - (Optional, List) Fields to group by.
								  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `stacked_group_name` - (Optional, String) Field to stack by.
								  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `logs` - (Optional, List) Logs specific query.
							Nested schema for **logs**:
								* `aggregation` - (Optional, List) Aggregations.
								Nested schema for **aggregation**:
									* `average` - (Optional, List) Calculate average value of log field.
									Nested schema for **average**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `count` - (Optional, List) Count the number of entries.
									Nested schema for **count**:
									* `count_distinct` - (Optional, List) Count the number of distinct values of log field.
									Nested schema for **count_distinct**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `max` - (Optional, List) Calculate maximum value of log field.
									Nested schema for **max**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `min` - (Optional, List) Calculate minimum value of log field.
									Nested schema for **min**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `percentile` - (Optional, List) Calculate percentile value of log field.
									Nested schema for **percentile**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `percent` - (Required, Float) Value in range (0, 100].
									* `sum` - (Optional, List) Sum values of log field.
									Nested schema for **sum**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
								* `filters` - (Optional, List) Extra filter on top of the Lucene query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `observation_field` - (Optional, List) Field to count distinct values of.
									Nested schema for **observation_field**:
										* `keypath` - (Optional, List) Path within the dataset scope.
										  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
										* `scope` - (Optional, String) Scope of the dataset.
										  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `operator` - (Optional, List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (Optional, List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (Optional, List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (Optional, List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (Optional, List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `group_names_fields` - (Optional, List) Fields to group by.
								  * Constraints: The maximum length is `2` items. The minimum length is `1` item.
								Nested schema for **group_names_fields**:
									* `keypath` - (Optional, List) Path within the dataset scope.
									  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
									* `scope` - (Optional, String) Scope of the dataset.
									  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
								* `lucene_query` - (Optional, List) Lucene query.
								Nested schema for **lucene_query**:
									* `value` - (Optional, String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `stacked_group_name_field` - (Optional, List) Field to count distinct values of.
								Nested schema for **stacked_group_name_field**:
									* `keypath` - (Optional, List) Path within the dataset scope.
									  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
									* `scope` - (Optional, String) Scope of the dataset.
									  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
							* `metrics` - (Optional, List) Metrics specific query.
							Nested schema for **metrics**:
								* `filters` - (Optional, List) Extra filter on top of the PromQL query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `label` - (Optional, String) Label associated with the metric.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
									* `operator` - (Optional, List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (Optional, List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (Optional, List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (Optional, List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (Optional, List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `group_names` - (Optional, List) Labels to group by.
								  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `2` items. The minimum length is `1` item.
								* `promql_query` - (Optional, List) PromQL query.
								Nested schema for **promql_query**:
									* `value` - (Optional, String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `stacked_group_name` - (Optional, String) Label to stack by.
								  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `scale_type` - (Optional, String) Scale type.
						  * Constraints: Allowable values are: `unspecified`, `linear`, `logarithmic`.
						* `sort_by` - (Optional, String) Sorting mode.
						  * Constraints: Allowable values are: `unspecified`, `value`, `name`.
						* `stack_definition` - (Optional, List) Stack definition.
						Nested schema for **stack_definition**:
							* `max_slices_per_bar` - (Optional, Integer) Maximum number of slices per bar.
							* `stack_name_template` - (Optional, String) Template for stack slice label.
							  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `unit` - (Optional, String) Unit of the data.
						  * Constraints: Allowable values are: `unspecified`, `microseconds`, `milliseconds`, `seconds`, `bytes`, `kbytes`, `mbytes`, `gbytes`, `bytes_iec`, `kibytes`, `mibytes`, `gibytes`, `eur_cents`, `eur`, `usd_cents`, `usd`.
						* `y_axis_view_by` - (Optional, List) Y-axis view mode.
						Nested schema for **y_axis_view_by**:
							* `category` - (Optional, List) View by category.
							Nested schema for **category**:
							* `value` - (Optional, List) View by value.
							Nested schema for **value**:
					* `line_chart` - (Optional, List) Line chart widget.
					Nested schema for **line_chart**:
						* `legend` - (Required, List) Legend configuration.
						Nested schema for **legend**:
							* `columns` - (Optional, List) The columns to show in the legend.
							  * Constraints: Allowable list items are: `unspecified`, `min`, `max`, `sum`, `avg`, `last`, `name`. The maximum length is `4096` items. The minimum length is `0` items.
							* `group_by_query` - (Required, Boolean) Whether to group by the query or not.
							* `is_visible` - (Required, Boolean) Whether to show the legend or not.
						* `query_definitions` - (Optional, List) Query definitions.
						  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
						Nested schema for **query_definitions**:
							* `color_scheme` - (Optional, String) Color scheme for the series.
							  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `data_mode_type` - (Optional, String) Data mode type.
							  * Constraints: Allowable values are: `high_unspecified`, `archive`.
							* `id` - (Required, String) Unique identifier of the query within the widget.
							  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
							* `is_visible` - (Required, Boolean) Whether data for this query should be visible on the chart.
							* `name` - (Optional, String) Query name.
							  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `query` - (Required, List) Data source specific query, defines from where and how to fetch the data.
							Nested schema for **query**:
								* `dataprime` - (Optional, List) Dataprime language based query.
								Nested schema for **dataprime**:
									* `dataprime_query` - (Required, List) Dataprime query.
									Nested schema for **dataprime_query**:
										* `text` - (Optional, String) The query string.
										  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
									* `filters` - (Optional, List) Filters to be applied to query results.
									  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
									Nested schema for **filters**:
										* `logs` - (Optional, List) Extra filtering on top of the Lucene query.
										Nested schema for **logs**:
											* `observation_field` - (Optional, List) Field to count distinct values of.
											Nested schema for **observation_field**:
												* `keypath` - (Optional, List) Path within the dataset scope.
												  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
												* `scope` - (Optional, String) Scope of the dataset.
												  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
											* `operator` - (Optional, List) Operator to use for filtering the logs.
											Nested schema for **operator**:
												* `equals` - (Optional, List) Equality comparison.
												Nested schema for **equals**:
													* `selection` - (Optional, List) Selection criteria for the equality comparison.
													Nested schema for **selection**:
														* `all` - (Optional, List) Represents a selection of all values.
														Nested schema for **all**:
														* `list` - (Optional, List) Represents a selection from a list of values.
														Nested schema for **list**:
															* `values` - (Optional, List) List of values for the selection.
															  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
												* `not_equals` - (Optional, List) Non-equality comparison.
												Nested schema for **not_equals**:
													* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
													Nested schema for **selection**:
														* `list` - (Optional, List) Represents a selection from a list of values.
														Nested schema for **list**:
															* `values` - (Optional, List) List of values for the selection.
															  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `metrics` - (Optional, List) Filtering to be applied to query results.
										Nested schema for **metrics**:
											* `label` - (Optional, String) Label associated with the metric.
											  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
											* `operator` - (Optional, List) Operator to use for filtering the logs.
											Nested schema for **operator**:
												* `equals` - (Optional, List) Equality comparison.
												Nested schema for **equals**:
													* `selection` - (Optional, List) Selection criteria for the equality comparison.
													Nested schema for **selection**:
														* `all` - (Optional, List) Represents a selection of all values.
														Nested schema for **all**:
														* `list` - (Optional, List) Represents a selection from a list of values.
														Nested schema for **list**:
															* `values` - (Optional, List) List of values for the selection.
															  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
												* `not_equals` - (Optional, List) Non-equality comparison.
												Nested schema for **not_equals**:
													* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
													Nested schema for **selection**:
														* `list` - (Optional, List) Represents a selection from a list of values.
														Nested schema for **list**:
															* `values` - (Optional, List) List of values for the selection.
															  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `logs` - (Optional, List) Logs specific query.
								Nested schema for **logs**:
									* `aggregations` - (Optional, List) Aggregations.
									  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
									Nested schema for **aggregations**:
										* `average` - (Optional, List) Calculate average value of log field.
										Nested schema for **average**:
											* `observation_field` - (Required, List) Field to count distinct values of.
											Nested schema for **observation_field**:
												* `keypath` - (Optional, List) Path within the dataset scope.
												  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
												* `scope` - (Optional, String) Scope of the dataset.
												  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `count` - (Optional, List) Count the number of entries.
										Nested schema for **count**:
										* `count_distinct` - (Optional, List) Count the number of distinct values of log field.
										Nested schema for **count_distinct**:
											* `observation_field` - (Required, List) Field to count distinct values of.
											Nested schema for **observation_field**:
												* `keypath` - (Optional, List) Path within the dataset scope.
												  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
												* `scope` - (Optional, String) Scope of the dataset.
												  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `max` - (Optional, List) Calculate maximum value of log field.
										Nested schema for **max**:
											* `observation_field` - (Required, List) Field to count distinct values of.
											Nested schema for **observation_field**:
												* `keypath` - (Optional, List) Path within the dataset scope.
												  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
												* `scope` - (Optional, String) Scope of the dataset.
												  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `min` - (Optional, List) Calculate minimum value of log field.
										Nested schema for **min**:
											* `observation_field` - (Required, List) Field to count distinct values of.
											Nested schema for **observation_field**:
												* `keypath` - (Optional, List) Path within the dataset scope.
												  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
												* `scope` - (Optional, String) Scope of the dataset.
												  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `percentile` - (Optional, List) Calculate percentile value of log field.
										Nested schema for **percentile**:
											* `observation_field` - (Required, List) Field to count distinct values of.
											Nested schema for **observation_field**:
												* `keypath` - (Optional, List) Path within the dataset scope.
												  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
												* `scope` - (Optional, String) Scope of the dataset.
												  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
											* `percent` - (Required, Float) Value in range (0, 100].
										* `sum` - (Optional, List) Sum values of log field.
										Nested schema for **sum**:
											* `observation_field` - (Required, List) Field to count distinct values of.
											Nested schema for **observation_field**:
												* `keypath` - (Optional, List) Path within the dataset scope.
												  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
												* `scope` - (Optional, String) Scope of the dataset.
												  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `filters` - (Optional, List) Extra filtering on top of the Lucene query.
									  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
									Nested schema for **filters**:
										* `observation_field` - (Optional, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `operator` - (Optional, List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (Optional, List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (Optional, List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (Optional, List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (Optional, List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
									* `group_by` - (Optional, List) Group by fields (deprecated).
									  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
									* `group_bys` - (Optional, List) Group by fields.
									  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
									Nested schema for **group_bys**:
										* `keypath` - (Optional, List) Path within the dataset scope.
										  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
										* `scope` - (Optional, String) Scope of the dataset.
										  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `lucene_query` - (Optional, List) Lucene query.
									Nested schema for **lucene_query**:
										* `value` - (Optional, String) The query string.
										  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `metrics` - (Optional, List) Metrics specific query.
								Nested schema for **metrics**:
									* `filters` - (Optional, List) Filtering to be applied to query results.
									  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
									Nested schema for **filters**:
										* `label` - (Optional, String) Label associated with the metric.
										  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
										* `operator` - (Optional, List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (Optional, List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (Optional, List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (Optional, List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (Optional, List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
									* `promql_query` - (Optional, List) PromQL query.
									Nested schema for **promql_query**:
										* `value` - (Optional, String) The query string.
										  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `resolution` - (Required, List) Resolution of the data.
							Nested schema for **resolution**:
								* `buckets_presented` - (Optional, Integer) Maximum number of data points to fetch.
								* `interval` - (Optional, String) Interval between data points.
								  * Constraints: The maximum length is `10` characters. The minimum length is `2` characters. The value must match regular expression `/^[0-9]+[smhdw]?$/`.
							* `scale_type` - (Optional, String) Scale type.
							  * Constraints: Allowable values are: `unspecified`, `linear`, `logarithmic`.
							* `series_count_limit` - (Optional, String) Maximum number of series to display.
							  * Constraints: The maximum length is `19` characters. The minimum length is `1` character. The value must match regular expression `/^-?\\d{1,19}$/`.
							* `series_name_template` - (Optional, String) Template for series name in legend and tooltip.
							  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `unit` - (Optional, String) Unit of the data.
							  * Constraints: Allowable values are: `unspecified`, `microseconds`, `milliseconds`, `seconds`, `bytes`, `kbytes`, `mbytes`, `gbytes`, `bytes_iec`, `kibytes`, `mibytes`, `gibytes`, `eur_cents`, `eur`, `usd_cents`, `usd`.
						* `stacked_line` - (Optional, String) Stacked lines.
						  * Constraints: Allowable values are: `unspecified`, `absolute`, `relative`.
						* `tooltip` - (Required, List) Tooltip configuration.
						Nested schema for **tooltip**:
							* `show_labels` - (Optional, Boolean) Whether to show labels in the tooltip.
							* `type` - (Optional, String) Tooltip type.
							  * Constraints: Allowable values are: `unspecified`, `all`, `single`.
					* `markdown` - (Optional, List) Markdown widget.
					Nested schema for **markdown**:
						* `markdown_text` - (Required, String) Markdown text to render.
						  * Constraints: The maximum length is `10000` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `tooltip_text` - (Optional, String) Tooltip text on hover.
						  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
					* `pie_chart` - (Optional, List) Pie chart widget.
					Nested schema for **pie_chart**:
						* `color_scheme` - (Required, String) Color scheme name.
						  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `data_mode_type` - (Optional, String) Data mode type.
						  * Constraints: Allowable values are: `high_unspecified`, `archive`.
						* `group_name_template` - (Optional, String) Template for group labels.
						  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `label_definition` - (Required, List) Label settings.
						Nested schema for **label_definition**:
							* `is_visible` - (Optional, Boolean) Controls whether to show the label.
							* `label_source` - (Optional, String) Source of the label.
							  * Constraints: Allowable values are: `unspecified`, `inner`, `stack`.
							* `show_name` - (Optional, Boolean) Controls whether to show the name.
							* `show_percentage` - (Optional, Boolean) Controls whether to show the percentage.
							* `show_value` - (Optional, Boolean) Controls whether to show the value.
						* `max_slices_per_chart` - (Required, Integer) Maximum number of slices to display in the chart.
						* `min_slice_percentage` - (Required, Integer) Minimum percentage of a slice to be displayed.
						* `query` - (Required, List) Data source specific query, defines from where and how to fetch the data.
						Nested schema for **query**:
							* `dataprime` - (Optional, List) Query based on Dataprime language.
							Nested schema for **dataprime**:
								* `dataprime_query` - (Required, List) Dataprime query.
								Nested schema for **dataprime_query**:
									* `text` - (Optional, String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `filters` - (Optional, List) Extra filters on top of Dataprime query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `logs` - (Optional, List) Extra filtering on top of the Lucene query.
									Nested schema for **logs**:
										* `observation_field` - (Optional, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `operator` - (Optional, List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (Optional, List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (Optional, List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (Optional, List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (Optional, List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
									* `metrics` - (Optional, List) Filtering to be applied to query results.
									Nested schema for **metrics**:
										* `label` - (Optional, String) Label associated with the metric.
										  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
										* `operator` - (Optional, List) Operator to use for filtering the logs.
										Nested schema for **operator**:
											* `equals` - (Optional, List) Equality comparison.
											Nested schema for **equals**:
												* `selection` - (Optional, List) Selection criteria for the equality comparison.
												Nested schema for **selection**:
													* `all` - (Optional, List) Represents a selection of all values.
													Nested schema for **all**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
											* `not_equals` - (Optional, List) Non-equality comparison.
											Nested schema for **not_equals**:
												* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
												Nested schema for **selection**:
													* `list` - (Optional, List) Represents a selection from a list of values.
													Nested schema for **list**:
														* `values` - (Optional, List) List of values for the selection.
														  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `group_names` - (Optional, List) Fields to group by.
								  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `2` items. The minimum length is `1` item.
								* `stacked_group_name` - (Optional, String) Field to stack by.
								  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
							* `logs` - (Optional, List) Logs specific query.
							Nested schema for **logs**:
								* `aggregation` - (Required, List) Aggregations.
								Nested schema for **aggregation**:
									* `average` - (Optional, List) Calculate average value of log field.
									Nested schema for **average**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `count` - (Optional, List) Count the number of entries.
									Nested schema for **count**:
									* `count_distinct` - (Optional, List) Count the number of distinct values of log field.
									Nested schema for **count_distinct**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `max` - (Optional, List) Calculate maximum value of log field.
									Nested schema for **max**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `min` - (Optional, List) Calculate minimum value of log field.
									Nested schema for **min**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `percentile` - (Optional, List) Calculate percentile value of log field.
									Nested schema for **percentile**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
										* `percent` - (Required, Float) Value in range (0, 100].
									* `sum` - (Optional, List) Sum values of log field.
									Nested schema for **sum**:
										* `observation_field` - (Required, List) Field to count distinct values of.
										Nested schema for **observation_field**:
											* `keypath` - (Optional, List) Path within the dataset scope.
											  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
											* `scope` - (Optional, String) Scope of the dataset.
											  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
								* `filters` - (Optional, List) Extra filters on top of Lucene query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `observation_field` - (Optional, List) Field to count distinct values of.
									Nested schema for **observation_field**:
										* `keypath` - (Optional, List) Path within the dataset scope.
										  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
										* `scope` - (Optional, String) Scope of the dataset.
										  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
									* `operator` - (Optional, List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (Optional, List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (Optional, List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (Optional, List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (Optional, List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `group_names_fields` - (Optional, List) Fields to group by.
								  * Constraints: The maximum length is `2` items. The minimum length is `1` item.
								Nested schema for **group_names_fields**:
									* `keypath` - (Optional, List) Path within the dataset scope.
									  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
									* `scope` - (Optional, String) Scope of the dataset.
									  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
								* `lucene_query` - (Optional, List) Lucene query.
								Nested schema for **lucene_query**:
									* `value` - (Optional, String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `stacked_group_name_field` - (Optional, List) Field to count distinct values of.
								Nested schema for **stacked_group_name_field**:
									* `keypath` - (Optional, List) Path within the dataset scope.
									  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
									* `scope` - (Optional, String) Scope of the dataset.
									  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
							* `metrics` - (Optional, List) Metrics specific query.
							Nested schema for **metrics**:
								* `filters` - (Optional, List) Extra filters on top of PromQL query.
								  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
								Nested schema for **filters**:
									* `label` - (Optional, String) Label associated with the metric.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
									* `operator` - (Optional, List) Operator to use for filtering the logs.
									Nested schema for **operator**:
										* `equals` - (Optional, List) Equality comparison.
										Nested schema for **equals**:
											* `selection` - (Optional, List) Selection criteria for the equality comparison.
											Nested schema for **selection**:
												* `all` - (Optional, List) Represents a selection of all values.
												Nested schema for **all**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
										* `not_equals` - (Optional, List) Non-equality comparison.
										Nested schema for **not_equals**:
											* `selection` - (Optional, List) Selection criteria for the non-equality comparison.
											Nested schema for **selection**:
												* `list` - (Optional, List) Represents a selection from a list of values.
												Nested schema for **list**:
													* `values` - (Optional, List) List of values for the selection.
													  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
								* `group_names` - (Optional, List) Fields to group by.
								  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `2` items. The minimum length is `1` item.
								* `promql_query` - (Required, List) PromQL query.
								Nested schema for **promql_query**:
									* `value` - (Optional, String) The query string.
									  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
								* `stacked_group_name` - (Optional, String) Field to stack by.
								  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `show_legend` - (Required, Boolean) Controls whether to show the legend.
						* `stack_definition` - (Required, List) Stack definition.
						Nested schema for **stack_definition**:
							* `max_slices_per_stack` - (Optional, Integer) Maximum number of slices per stack.
							* `stack_name_template` - (Optional, String) Template for stack labels.
							  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `unit` - (Optional, String) Unit of the data.
						  * Constraints: Allowable values are: `unspecified`, `microseconds`, `milliseconds`, `seconds`, `bytes`, `kbytes`, `mbytes`, `gbytes`, `bytes_iec`, `kibytes`, `mibytes`, `gibytes`, `eur_cents`, `eur`, `usd_cents`, `usd`.
				* `description` - (Optional, String) Widget description.
				  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `href` - (Optional, String) Widget identifier within the dashboard.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `id` - (Required, List) Unique identifier of the folder containing the dashboard.
				Nested schema for **id**:
					* `value` - (Required, String) The UUID value.
					  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
				* `title` - (Required, String) Widget title.
				  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `updated_at` - (Optional, String) Last update timestamp.
* `name` - (Required, String) Display name of the dashboard.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
* `relative_time_frame` - (Optional, String) Relative time frame specifying a duration from the current time.
  * Constraints: The maximum length is `10` characters. The minimum length is `2` characters. The value must match regular expression `/^[0-9]+[smhdw]?$/`.
* `two_minutes` - (Optional, List) Auto refresh interval is set to two minutes.
Nested schema for **two_minutes**:
* `variables` - (Optional, List) List of variables that can be used within the dashboard for dynamic content.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **variables**:
	* `definition` - (Required, List) Definition.
	Nested schema for **definition**:
		* `multi_select` - (Optional, List) Multi-select value.
		Nested schema for **multi_select**:
			* `selection` - (Required, List) State of what is currently selected.
			Nested schema for **selection**:
				* `all` - (Optional, List) All values are selected, usually translated to wildcard (*).
				Nested schema for **all**:
				* `list` - (Optional, List) Specific values are selected.
				Nested schema for **list**:
					* `values` - (Optional, List) Selected values.
					  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
			* `source` - (Required, List) Variable value source.
			Nested schema for **source**:
				* `constant_list` - (Optional, List) List of constant values.
				Nested schema for **constant_list**:
					* `values` - (Required, List) List of constant values.
					  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
				* `logs_path` - (Optional, List) Unique values for a given logs path.
				Nested schema for **logs_path**:
					* `observation_field` - (Required, List) Field to count distinct values of.
					Nested schema for **observation_field**:
						* `keypath` - (Optional, List) Path within the dataset scope.
						  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
						* `scope` - (Optional, String) Scope of the dataset.
						  * Constraints: Allowable values are: `unspecified`, `user_data`, `label`, `metadata`.
				* `metric_label` - (Optional, List) Unique values for a given metric label.
				Nested schema for **metric_label**:
					* `label` - (Required, String) Metric label to source unique values from.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
					* `metric_name` - (Required, String) Metric name to source unique values from.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
			* `values_order_direction` - (Required, String) The direction of the order: ascending or descending.
			  * Constraints: Allowable values are: `unspecified`, `asc`, `desc`.
	* `display_name` - (Required, String) Name used in variable UI.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `name` - (Required, String) Name of the variable which can be used in templates.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_dashboard resource.
* `dashboard_id` - The unique identifier of the logs dashboard.


## Import

You can import the `ibm_logs_dashboard` resource by using `id`. `id` combination of `region`, `instance_id` and `dashboard_id`.

# Syntax
<pre>
$ terraform import ibm_logs_dashboard.logs_dashboard < region >/< instance_id >/< dashboard_id >;
</pre>

# Example
```
$ terraform import ibm_logs_dashboard.logs_dashboard eu-gb/3dc02998-0b50-4ea8-b68a-4779d716fa1f/6U1Q8Hpa263Se8PkRKaiE
```

// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	// . "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
)

func TestAccIbmLogsDashboardDataSourceBasic(t *testing.T) {
	dashboardName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsDashboardDataSourceConfigBasic(dashboardName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "dashboard_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "layout.#"),
				),
			},
		},
	})
}

func TestAccIbmLogsDashboardDataSourceAllArgs(t *testing.T) {
	dashboardHref := fmt.Sprintf("tf_href_%d", acctest.RandIntRange(10, 100))
	dashboardName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	dashboardDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	dashboardRelativeTimeFrame := fmt.Sprintf("tf_relative_time_frame_%d", acctest.RandIntRange(10, 100))
	dashboardFalse := "false"
	dashboardTwoMinutes := "false"
	dashboardFiveMinutes := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsDashboardDataSourceConfig(dashboardHref, dashboardName, dashboardDescription, dashboardRelativeTimeFrame, dashboardFalse, dashboardTwoMinutes, dashboardFiveMinutes),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "dashboard_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "layout.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "variables.#"),
					resource.TestCheckResourceAttr("data.ibm_logs_dashboard.logs_dashboard_instance", "variables.0.name", dashboardName),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "variables.0.display_name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "filters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "filters.0.enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "filters.0.collapsed"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "annotations.#"),
					resource.TestCheckResourceAttr("data.ibm_logs_dashboard.logs_dashboard_instance", "annotations.0.href", dashboardHref),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "annotations.0.id"),
					resource.TestCheckResourceAttr("data.ibm_logs_dashboard.logs_dashboard_instance", "annotations.0.name", dashboardName),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "annotations.0.enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "absolute_time_frame.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "relative_time_frame"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "folder_id.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "folder_path.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "false"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "two_minutes"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "five_minutes"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsDashboardDataSourceConfigBasic(dashboardName string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_dashboard" "logs_dashboard_instance" {
			name = "%s"
			layout {
				sections {
					href = "href"
					id {
						value = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
					}
					rows {
						href = "href"
						id {
							value = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
						}
						appearance {
							height = 5
						}
						widgets {
							href = "href"
							id {
								value = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
							}
							title = "Response time"
							description = "The average response time of the system"
							definition {
								line_chart {
									legend {
										is_visible = true
										columns = [ "unspecified" ]
										group_by_query = true
									}
									tooltip {
										show_labels = true
										type = "unspecified"
									}
									query_definitions {
										id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
										query {
											logs {
												lucene_query {
													value = "coralogix.metadata.applicationName:"production""
												}
												group_by = [ "group_by" ]
												aggregations {
													count = {  }
												}
												filters {
													operator {
														equals {
															selection {
																all = {  }
															}
														}
													}
													observation_field {
														keypath = [ "keypath" ]
														scope = "unspecified"
													}
												}
												group_bys {
													keypath = [ "keypath" ]
													scope = "unspecified"
												}
											}
										}
										series_name_template = "{{severity}}"
										series_count_limit = "10"
										unit = "unspecified"
										scale_type = "unspecified"
										name = "CPU usage"
										is_visible = true
										color_scheme = "classic"
										resolution {
											interval = "1m"
											buckets_presented = 100
										}
										data_mode_type = "high_unspecified"
									}
								}
							}
						}
					}
				}
			}
		}

		data "ibm_logs_dashboard" "logs_dashboard_instance" {
			dashboard_id = "dashboard_id"
		}
	`, dashboardName)
}

func testAccCheckIbmLogsDashboardDataSourceConfig(dashboardHref string, dashboardName string, dashboardDescription string, dashboardRelativeTimeFrame string, dashboardFalse string, dashboardTwoMinutes string, dashboardFiveMinutes string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_dashboard" "logs_dashboard_instance" {
			href = "%s"
			name = "%s"
			description = "%s"
			layout {
				sections {
					href = "href"
					id {
						value = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
					}
					rows {
						href = "href"
						id {
							value = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
						}
						appearance {
							height = 5
						}
						widgets {
							href = "href"
							id {
								value = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
							}
							title = "Response time"
							description = "The average response time of the system"
							definition {
								line_chart {
									legend {
										is_visible = true
										columns = [ "unspecified" ]
										group_by_query = true
									}
									tooltip {
										show_labels = true
										type = "unspecified"
									}
									query_definitions {
										id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
										query {
											logs {
												lucene_query {
													value = "coralogix.metadata.applicationName:"production""
												}
												group_by = [ "group_by" ]
												aggregations {
													count = {  }
												}
												filters {
													operator {
														equals {
															selection {
																all = {  }
															}
														}
													}
													observation_field {
														keypath = [ "keypath" ]
														scope = "unspecified"
													}
												}
												group_bys {
													keypath = [ "keypath" ]
													scope = "unspecified"
												}
											}
										}
										series_name_template = "{{severity}}"
										series_count_limit = "10"
										unit = "unspecified"
										scale_type = "unspecified"
										name = "CPU usage"
										is_visible = true
										color_scheme = "classic"
										resolution {
											interval = "1m"
											buckets_presented = 100
										}
										data_mode_type = "high_unspecified"
									}
								}
							}
						}
					}
				}
			}
			variables {
				name = "service_name"
				definition {
					multi_select {
						source {
							logs_path {
								observation_field {
									keypath = [ "keypath" ]
									scope = "unspecified"
								}
							}
						}
						selection {
							all = {  }
						}
						values_order_direction = "unspecified"
					}
				}
				display_name = "Service Name"
			}
			filters {
				source {
					logs {
						operator {
							equals {
								selection {
									all = {  }
								}
							}
						}
						observation_field {
							keypath = [ "keypath" ]
							scope = "unspecified"
						}
					}
				}
				enabled = true
				collapsed = true
			}
			annotations {
				href = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
				id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
				name = "Deployments"
				enabled = true
				source {
					metrics {
						promql_query {
							value = "sum(up)"
						}
						strategy {
							start_time_metric = {  }
						}
						message_template = "message_template"
						labels = [ "labels" ]
					}
				}
			}
			absolute_time_frame {
				from = "2021-01-31T09:44:12Z"
				to = "2021-01-31T09:44:12Z"
			}
			relative_time_frame = "%s"
			folder_id {
				value = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
			}
			folder_path {
				segments = [ "segments" ]
			}
			false = %s
			two_minutes = %s
			five_minutes = %s
		}

		data "ibm_logs_dashboard" "logs_dashboard_instance" {
			dashboard_id = "dashboard_id"
		}
	`, dashboardHref, dashboardName, dashboardDescription, dashboardRelativeTimeFrame, dashboardFalse, dashboardTwoMinutes, dashboardFiveMinutes)
}

// Todo @kavya498: Fix unit testcases
// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstLayoutToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1UUIDModel := make(map[string]interface{})
// 		apisDashboardsV1UUIDModel["value"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"

// 		apisDashboardsV1AstRowAppearanceModel := make(map[string]interface{})
// 		apisDashboardsV1AstRowAppearanceModel["height"] = int(5)

// 		apisDashboardsV1AstWidgetsCommonLegendModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLegendModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsCommonLegendModel["columns"] = []string{"unspecified"}
// 		apisDashboardsV1AstWidgetsCommonLegendModel["group_by_query"] = true

// 		apisDashboardsV1AstWidgetsLineChartTooltipModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartTooltipModel["show_labels"] = true
// 		apisDashboardsV1AstWidgetsLineChartTooltipModel["type"] = "unspecified"

// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_by"] = []string{"testString"}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["aggregations"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_bys"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsLineChartQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartQueryModel["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartLogsQueryModel}

// 		apisDashboardsV1AstWidgetsLineChartResolutionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartResolutionModel["interval"] = "1m"
// 		apisDashboardsV1AstWidgetsLineChartResolutionModel["buckets_presented"] = int(100)

// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["series_name_template"] = "{{severity}}"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["series_count_limit"] = "10"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["unit"] = "unspecified"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["scale_type"] = "unspecified"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["name"] = "CPU usage"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["color_scheme"] = "classic"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["resolution"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartResolutionModel}
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["data_mode_type"] = "high_unspecified"

// 		apisDashboardsV1AstWidgetsLineChartModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartModel["legend"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLegendModel}
// 		apisDashboardsV1AstWidgetsLineChartModel["tooltip"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartTooltipModel}
// 		apisDashboardsV1AstWidgetsLineChartModel["query_definitions"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel}

// 		apisDashboardsV1AstWidgetDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetDefinitionModel["line_chart"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartModel}

// 		apisDashboardsV1AstWidgetModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetModel["href"] = "testString"
// 		apisDashboardsV1AstWidgetModel["id"] = []map[string]interface{}{apisDashboardsV1UUIDModel}
// 		apisDashboardsV1AstWidgetModel["title"] = "Response time"
// 		apisDashboardsV1AstWidgetModel["description"] = "The average response time of the system"
// 		apisDashboardsV1AstWidgetModel["definition"] = []map[string]interface{}{apisDashboardsV1AstWidgetDefinitionModel}

// 		apisDashboardsV1AstRowModel := make(map[string]interface{})
// 		apisDashboardsV1AstRowModel["href"] = "testString"
// 		apisDashboardsV1AstRowModel["id"] = []map[string]interface{}{apisDashboardsV1UUIDModel}
// 		apisDashboardsV1AstRowModel["appearance"] = []map[string]interface{}{apisDashboardsV1AstRowAppearanceModel}
// 		apisDashboardsV1AstRowModel["widgets"] = []map[string]interface{}{apisDashboardsV1AstWidgetModel}

// 		apisDashboardsV1AstSectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstSectionModel["href"] = "testString"
// 		apisDashboardsV1AstSectionModel["id"] = []map[string]interface{}{apisDashboardsV1UUIDModel}
// 		apisDashboardsV1AstSectionModel["rows"] = []map[string]interface{}{apisDashboardsV1AstRowModel}

// 		model := make(map[string]interface{})
// 		model["sections"] = []map[string]interface{}{apisDashboardsV1AstSectionModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1UUIDModel := new(logsv0.ApisDashboardsV1UUID)
// 	apisDashboardsV1UUIDModel.Value = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")

// 	apisDashboardsV1AstRowAppearanceModel := new(logsv0.ApisDashboardsV1AstRowAppearance)
// 	apisDashboardsV1AstRowAppearanceModel.Height = core.Int64Ptr(int64(5))

// 	apisDashboardsV1AstWidgetsCommonLegendModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLegend)
// 	apisDashboardsV1AstWidgetsCommonLegendModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsCommonLegendModel.Columns = []string{"unspecified"}
// 	apisDashboardsV1AstWidgetsCommonLegendModel.GroupByQuery = core.BoolPtr(true)

// 	apisDashboardsV1AstWidgetsLineChartTooltipModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartTooltip)
// 	apisDashboardsV1AstWidgetsLineChartTooltipModel.ShowLabels = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsLineChartTooltipModel.Type = core.StringPtr("unspecified")

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartLogsQuery)
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBy = []string{"testString"}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Aggregations = []logsv0.ApisDashboardsV1CommonLogsAggregationIntf{apisDashboardsV1CommonLogsAggregationModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBys = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	apisDashboardsV1AstWidgetsLineChartQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryValueLogs)
// 	apisDashboardsV1AstWidgetsLineChartQueryModel.Logs = apisDashboardsV1AstWidgetsLineChartLogsQueryModel

// 	apisDashboardsV1AstWidgetsLineChartResolutionModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartResolution)
// 	apisDashboardsV1AstWidgetsLineChartResolutionModel.Interval = core.StringPtr("1m")
// 	apisDashboardsV1AstWidgetsLineChartResolutionModel.BucketsPresented = core.Int64Ptr(int64(100))

// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryDefinition)
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Query = apisDashboardsV1AstWidgetsLineChartQueryModel
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.SeriesNameTemplate = core.StringPtr("{{severity}}")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.SeriesCountLimit = core.StringPtr("10")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Unit = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ScaleType = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Name = core.StringPtr("CPU usage")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ColorScheme = core.StringPtr("classic")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Resolution = apisDashboardsV1AstWidgetsLineChartResolutionModel
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.DataModeType = core.StringPtr("high_unspecified")

// 	apisDashboardsV1AstWidgetsLineChartModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChart)
// 	apisDashboardsV1AstWidgetsLineChartModel.Legend = apisDashboardsV1AstWidgetsCommonLegendModel
// 	apisDashboardsV1AstWidgetsLineChartModel.Tooltip = apisDashboardsV1AstWidgetsLineChartTooltipModel
// 	apisDashboardsV1AstWidgetsLineChartModel.QueryDefinitions = []logsv0.ApisDashboardsV1AstWidgetsLineChartQueryDefinition{*apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel}

// 	apisDashboardsV1AstWidgetDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetDefinitionValueLineChart)
// 	apisDashboardsV1AstWidgetDefinitionModel.LineChart = apisDashboardsV1AstWidgetsLineChartModel

// 	apisDashboardsV1AstWidgetModel := new(logsv0.ApisDashboardsV1AstWidget)
// 	apisDashboardsV1AstWidgetModel.Href = core.StringPtr("testString")
// 	apisDashboardsV1AstWidgetModel.ID = apisDashboardsV1UUIDModel
// 	apisDashboardsV1AstWidgetModel.Title = core.StringPtr("Response time")
// 	apisDashboardsV1AstWidgetModel.Description = core.StringPtr("The average response time of the system")
// 	apisDashboardsV1AstWidgetModel.Definition = apisDashboardsV1AstWidgetDefinitionModel

// 	apisDashboardsV1AstRowModel := new(logsv0.ApisDashboardsV1AstRow)
// 	apisDashboardsV1AstRowModel.Href = core.StringPtr("testString")
// 	apisDashboardsV1AstRowModel.ID = apisDashboardsV1UUIDModel
// 	apisDashboardsV1AstRowModel.Appearance = apisDashboardsV1AstRowAppearanceModel
// 	apisDashboardsV1AstRowModel.Widgets = []logsv0.ApisDashboardsV1AstWidget{*apisDashboardsV1AstWidgetModel}

// 	apisDashboardsV1AstSectionModel := new(logsv0.ApisDashboardsV1AstSection)
// 	apisDashboardsV1AstSectionModel.Href = core.StringPtr("testString")
// 	apisDashboardsV1AstSectionModel.ID = apisDashboardsV1UUIDModel
// 	apisDashboardsV1AstSectionModel.Rows = []logsv0.ApisDashboardsV1AstRow{*apisDashboardsV1AstRowModel}

// 	model := new(logsv0.ApisDashboardsV1AstLayout)
// 	model.Sections = []logsv0.ApisDashboardsV1AstSection{*apisDashboardsV1AstSectionModel}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstLayoutToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstSectionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1UUIDModel := make(map[string]interface{})
// 		apisDashboardsV1UUIDModel["value"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"

// 		apisDashboardsV1AstRowAppearanceModel := make(map[string]interface{})
// 		apisDashboardsV1AstRowAppearanceModel["height"] = int(5)

// 		apisDashboardsV1AstWidgetsCommonLegendModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLegendModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsCommonLegendModel["columns"] = []string{"unspecified"}
// 		apisDashboardsV1AstWidgetsCommonLegendModel["group_by_query"] = true

// 		apisDashboardsV1AstWidgetsLineChartTooltipModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartTooltipModel["show_labels"] = true
// 		apisDashboardsV1AstWidgetsLineChartTooltipModel["type"] = "unspecified"

// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_by"] = []string{"testString"}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["aggregations"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_bys"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsLineChartQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartQueryModel["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartLogsQueryModel}

// 		apisDashboardsV1AstWidgetsLineChartResolutionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartResolutionModel["interval"] = "1m"
// 		apisDashboardsV1AstWidgetsLineChartResolutionModel["buckets_presented"] = int(100)

// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["series_name_template"] = "{{severity}}"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["series_count_limit"] = "10"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["unit"] = "unspecified"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["scale_type"] = "unspecified"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["name"] = "CPU usage"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["color_scheme"] = "classic"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["resolution"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartResolutionModel}
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["data_mode_type"] = "high_unspecified"

// 		apisDashboardsV1AstWidgetsLineChartModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartModel["legend"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLegendModel}
// 		apisDashboardsV1AstWidgetsLineChartModel["tooltip"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartTooltipModel}
// 		apisDashboardsV1AstWidgetsLineChartModel["query_definitions"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel}

// 		apisDashboardsV1AstWidgetDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetDefinitionModel["line_chart"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartModel}

// 		apisDashboardsV1AstWidgetModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetModel["href"] = "testString"
// 		apisDashboardsV1AstWidgetModel["id"] = []map[string]interface{}{apisDashboardsV1UUIDModel}
// 		apisDashboardsV1AstWidgetModel["title"] = "Response time"
// 		apisDashboardsV1AstWidgetModel["description"] = "The average response time of the system"
// 		apisDashboardsV1AstWidgetModel["definition"] = []map[string]interface{}{apisDashboardsV1AstWidgetDefinitionModel}

// 		apisDashboardsV1AstRowModel := make(map[string]interface{})
// 		apisDashboardsV1AstRowModel["href"] = "testString"
// 		apisDashboardsV1AstRowModel["id"] = []map[string]interface{}{apisDashboardsV1UUIDModel}
// 		apisDashboardsV1AstRowModel["appearance"] = []map[string]interface{}{apisDashboardsV1AstRowAppearanceModel}
// 		apisDashboardsV1AstRowModel["widgets"] = []map[string]interface{}{apisDashboardsV1AstWidgetModel}

// 		model := make(map[string]interface{})
// 		model["href"] = "testString"
// 		model["id"] = []map[string]interface{}{apisDashboardsV1UUIDModel}
// 		model["rows"] = []map[string]interface{}{apisDashboardsV1AstRowModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1UUIDModel := new(logsv0.ApisDashboardsV1UUID)
// 	apisDashboardsV1UUIDModel.Value = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")

// 	apisDashboardsV1AstRowAppearanceModel := new(logsv0.ApisDashboardsV1AstRowAppearance)
// 	apisDashboardsV1AstRowAppearanceModel.Height = core.Int64Ptr(int64(5))

// 	apisDashboardsV1AstWidgetsCommonLegendModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLegend)
// 	apisDashboardsV1AstWidgetsCommonLegendModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsCommonLegendModel.Columns = []string{"unspecified"}
// 	apisDashboardsV1AstWidgetsCommonLegendModel.GroupByQuery = core.BoolPtr(true)

// 	apisDashboardsV1AstWidgetsLineChartTooltipModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartTooltip)
// 	apisDashboardsV1AstWidgetsLineChartTooltipModel.ShowLabels = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsLineChartTooltipModel.Type = core.StringPtr("unspecified")

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartLogsQuery)
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBy = []string{"testString"}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Aggregations = []logsv0.ApisDashboardsV1CommonLogsAggregationIntf{apisDashboardsV1CommonLogsAggregationModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBys = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	apisDashboardsV1AstWidgetsLineChartQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryValueLogs)
// 	apisDashboardsV1AstWidgetsLineChartQueryModel.Logs = apisDashboardsV1AstWidgetsLineChartLogsQueryModel

// 	apisDashboardsV1AstWidgetsLineChartResolutionModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartResolution)
// 	apisDashboardsV1AstWidgetsLineChartResolutionModel.Interval = core.StringPtr("1m")
// 	apisDashboardsV1AstWidgetsLineChartResolutionModel.BucketsPresented = core.Int64Ptr(int64(100))

// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryDefinition)
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Query = apisDashboardsV1AstWidgetsLineChartQueryModel
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.SeriesNameTemplate = core.StringPtr("{{severity}}")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.SeriesCountLimit = core.StringPtr("10")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Unit = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ScaleType = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Name = core.StringPtr("CPU usage")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ColorScheme = core.StringPtr("classic")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Resolution = apisDashboardsV1AstWidgetsLineChartResolutionModel
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.DataModeType = core.StringPtr("high_unspecified")

// 	apisDashboardsV1AstWidgetsLineChartModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChart)
// 	apisDashboardsV1AstWidgetsLineChartModel.Legend = apisDashboardsV1AstWidgetsCommonLegendModel
// 	apisDashboardsV1AstWidgetsLineChartModel.Tooltip = apisDashboardsV1AstWidgetsLineChartTooltipModel
// 	apisDashboardsV1AstWidgetsLineChartModel.QueryDefinitions = []logsv0.ApisDashboardsV1AstWidgetsLineChartQueryDefinition{*apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel}

// 	apisDashboardsV1AstWidgetDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetDefinitionValueLineChart)
// 	apisDashboardsV1AstWidgetDefinitionModel.LineChart = apisDashboardsV1AstWidgetsLineChartModel

// 	apisDashboardsV1AstWidgetModel := new(logsv0.ApisDashboardsV1AstWidget)
// 	apisDashboardsV1AstWidgetModel.Href = core.StringPtr("testString")
// 	apisDashboardsV1AstWidgetModel.ID = apisDashboardsV1UUIDModel
// 	apisDashboardsV1AstWidgetModel.Title = core.StringPtr("Response time")
// 	apisDashboardsV1AstWidgetModel.Description = core.StringPtr("The average response time of the system")
// 	apisDashboardsV1AstWidgetModel.Definition = apisDashboardsV1AstWidgetDefinitionModel

// 	apisDashboardsV1AstRowModel := new(logsv0.ApisDashboardsV1AstRow)
// 	apisDashboardsV1AstRowModel.Href = core.StringPtr("testString")
// 	apisDashboardsV1AstRowModel.ID = apisDashboardsV1UUIDModel
// 	apisDashboardsV1AstRowModel.Appearance = apisDashboardsV1AstRowAppearanceModel
// 	apisDashboardsV1AstRowModel.Widgets = []logsv0.ApisDashboardsV1AstWidget{*apisDashboardsV1AstWidgetModel}

// 	model := new(logsv0.ApisDashboardsV1AstSection)
// 	model.Href = core.StringPtr("testString")
// 	model.ID = apisDashboardsV1UUIDModel
// 	model.Rows = []logsv0.ApisDashboardsV1AstRow{*apisDashboardsV1AstRowModel}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstSectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1UUIDToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["value"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1UUID)
// 	model.Value = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1UUIDToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstRowToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1UUIDModel := make(map[string]interface{})
// 		apisDashboardsV1UUIDModel["value"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"

// 		apisDashboardsV1AstRowAppearanceModel := make(map[string]interface{})
// 		apisDashboardsV1AstRowAppearanceModel["height"] = int(5)

// 		apisDashboardsV1AstWidgetsCommonLegendModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLegendModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsCommonLegendModel["columns"] = []string{"unspecified"}
// 		apisDashboardsV1AstWidgetsCommonLegendModel["group_by_query"] = true

// 		apisDashboardsV1AstWidgetsLineChartTooltipModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartTooltipModel["show_labels"] = true
// 		apisDashboardsV1AstWidgetsLineChartTooltipModel["type"] = "unspecified"

// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_by"] = []string{"testString"}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["aggregations"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_bys"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsLineChartQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartQueryModel["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartLogsQueryModel}

// 		apisDashboardsV1AstWidgetsLineChartResolutionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartResolutionModel["interval"] = "1m"
// 		apisDashboardsV1AstWidgetsLineChartResolutionModel["buckets_presented"] = int(100)

// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["series_name_template"] = "{{severity}}"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["series_count_limit"] = "10"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["unit"] = "unspecified"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["scale_type"] = "unspecified"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["name"] = "CPU usage"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["color_scheme"] = "classic"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["resolution"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartResolutionModel}
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["data_mode_type"] = "high_unspecified"

// 		apisDashboardsV1AstWidgetsLineChartModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartModel["legend"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLegendModel}
// 		apisDashboardsV1AstWidgetsLineChartModel["tooltip"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartTooltipModel}
// 		apisDashboardsV1AstWidgetsLineChartModel["query_definitions"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel}

// 		apisDashboardsV1AstWidgetDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetDefinitionModel["line_chart"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartModel}

// 		apisDashboardsV1AstWidgetModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetModel["href"] = "testString"
// 		apisDashboardsV1AstWidgetModel["id"] = []map[string]interface{}{apisDashboardsV1UUIDModel}
// 		apisDashboardsV1AstWidgetModel["title"] = "Response time"
// 		apisDashboardsV1AstWidgetModel["description"] = "The average response time of the system"
// 		apisDashboardsV1AstWidgetModel["definition"] = []map[string]interface{}{apisDashboardsV1AstWidgetDefinitionModel}

// 		model := make(map[string]interface{})
// 		model["href"] = "testString"
// 		model["id"] = []map[string]interface{}{apisDashboardsV1UUIDModel}
// 		model["appearance"] = []map[string]interface{}{apisDashboardsV1AstRowAppearanceModel}
// 		model["widgets"] = []map[string]interface{}{apisDashboardsV1AstWidgetModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1UUIDModel := new(logsv0.ApisDashboardsV1UUID)
// 	apisDashboardsV1UUIDModel.Value = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")

// 	apisDashboardsV1AstRowAppearanceModel := new(logsv0.ApisDashboardsV1AstRowAppearance)
// 	apisDashboardsV1AstRowAppearanceModel.Height = core.Int64Ptr(int64(5))

// 	apisDashboardsV1AstWidgetsCommonLegendModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLegend)
// 	apisDashboardsV1AstWidgetsCommonLegendModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsCommonLegendModel.Columns = []string{"unspecified"}
// 	apisDashboardsV1AstWidgetsCommonLegendModel.GroupByQuery = core.BoolPtr(true)

// 	apisDashboardsV1AstWidgetsLineChartTooltipModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartTooltip)
// 	apisDashboardsV1AstWidgetsLineChartTooltipModel.ShowLabels = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsLineChartTooltipModel.Type = core.StringPtr("unspecified")

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartLogsQuery)
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBy = []string{"testString"}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Aggregations = []logsv0.ApisDashboardsV1CommonLogsAggregationIntf{apisDashboardsV1CommonLogsAggregationModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBys = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	apisDashboardsV1AstWidgetsLineChartQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryValueLogs)
// 	apisDashboardsV1AstWidgetsLineChartQueryModel.Logs = apisDashboardsV1AstWidgetsLineChartLogsQueryModel

// 	apisDashboardsV1AstWidgetsLineChartResolutionModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartResolution)
// 	apisDashboardsV1AstWidgetsLineChartResolutionModel.Interval = core.StringPtr("1m")
// 	apisDashboardsV1AstWidgetsLineChartResolutionModel.BucketsPresented = core.Int64Ptr(int64(100))

// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryDefinition)
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Query = apisDashboardsV1AstWidgetsLineChartQueryModel
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.SeriesNameTemplate = core.StringPtr("{{severity}}")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.SeriesCountLimit = core.StringPtr("10")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Unit = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ScaleType = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Name = core.StringPtr("CPU usage")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ColorScheme = core.StringPtr("classic")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Resolution = apisDashboardsV1AstWidgetsLineChartResolutionModel
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.DataModeType = core.StringPtr("high_unspecified")

// 	apisDashboardsV1AstWidgetsLineChartModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChart)
// 	apisDashboardsV1AstWidgetsLineChartModel.Legend = apisDashboardsV1AstWidgetsCommonLegendModel
// 	apisDashboardsV1AstWidgetsLineChartModel.Tooltip = apisDashboardsV1AstWidgetsLineChartTooltipModel
// 	apisDashboardsV1AstWidgetsLineChartModel.QueryDefinitions = []logsv0.ApisDashboardsV1AstWidgetsLineChartQueryDefinition{*apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel}

// 	apisDashboardsV1AstWidgetDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetDefinitionValueLineChart)
// 	apisDashboardsV1AstWidgetDefinitionModel.LineChart = apisDashboardsV1AstWidgetsLineChartModel

// 	apisDashboardsV1AstWidgetModel := new(logsv0.ApisDashboardsV1AstWidget)
// 	apisDashboardsV1AstWidgetModel.Href = core.StringPtr("testString")
// 	apisDashboardsV1AstWidgetModel.ID = apisDashboardsV1UUIDModel
// 	apisDashboardsV1AstWidgetModel.Title = core.StringPtr("Response time")
// 	apisDashboardsV1AstWidgetModel.Description = core.StringPtr("The average response time of the system")
// 	apisDashboardsV1AstWidgetModel.Definition = apisDashboardsV1AstWidgetDefinitionModel

// 	model := new(logsv0.ApisDashboardsV1AstRow)
// 	model.Href = core.StringPtr("testString")
// 	model.ID = apisDashboardsV1UUIDModel
// 	model.Appearance = apisDashboardsV1AstRowAppearanceModel
// 	model.Widgets = []logsv0.ApisDashboardsV1AstWidget{*apisDashboardsV1AstWidgetModel}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstRowToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstRowAppearanceToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["height"] = int(5)

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstRowAppearance)
// 	model.Height = core.Int64Ptr(int64(5))

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstRowAppearanceToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1UUIDModel := make(map[string]interface{})
// 		apisDashboardsV1UUIDModel["value"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"

// 		apisDashboardsV1AstWidgetsCommonLegendModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLegendModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsCommonLegendModel["columns"] = []string{"unspecified"}
// 		apisDashboardsV1AstWidgetsCommonLegendModel["group_by_query"] = true

// 		apisDashboardsV1AstWidgetsLineChartTooltipModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartTooltipModel["show_labels"] = true
// 		apisDashboardsV1AstWidgetsLineChartTooltipModel["type"] = "unspecified"

// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_by"] = []string{"testString"}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["aggregations"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_bys"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsLineChartQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartQueryModel["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartLogsQueryModel}

// 		apisDashboardsV1AstWidgetsLineChartResolutionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartResolutionModel["interval"] = "1m"
// 		apisDashboardsV1AstWidgetsLineChartResolutionModel["buckets_presented"] = int(100)

// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["series_name_template"] = "{{severity}}"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["series_count_limit"] = "10"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["unit"] = "unspecified"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["scale_type"] = "unspecified"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["name"] = "CPU usage"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["color_scheme"] = "classic"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["resolution"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartResolutionModel}
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["data_mode_type"] = "high_unspecified"

// 		apisDashboardsV1AstWidgetsLineChartModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartModel["legend"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLegendModel}
// 		apisDashboardsV1AstWidgetsLineChartModel["tooltip"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartTooltipModel}
// 		apisDashboardsV1AstWidgetsLineChartModel["query_definitions"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel}

// 		apisDashboardsV1AstWidgetDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetDefinitionModel["line_chart"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartModel}

// 		model := make(map[string]interface{})
// 		model["href"] = "testString"
// 		model["id"] = []map[string]interface{}{apisDashboardsV1UUIDModel}
// 		model["title"] = "Response time"
// 		model["description"] = "The average response time of the system"
// 		model["definition"] = []map[string]interface{}{apisDashboardsV1AstWidgetDefinitionModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1UUIDModel := new(logsv0.ApisDashboardsV1UUID)
// 	apisDashboardsV1UUIDModel.Value = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")

// 	apisDashboardsV1AstWidgetsCommonLegendModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLegend)
// 	apisDashboardsV1AstWidgetsCommonLegendModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsCommonLegendModel.Columns = []string{"unspecified"}
// 	apisDashboardsV1AstWidgetsCommonLegendModel.GroupByQuery = core.BoolPtr(true)

// 	apisDashboardsV1AstWidgetsLineChartTooltipModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartTooltip)
// 	apisDashboardsV1AstWidgetsLineChartTooltipModel.ShowLabels = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsLineChartTooltipModel.Type = core.StringPtr("unspecified")

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartLogsQuery)
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBy = []string{"testString"}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Aggregations = []logsv0.ApisDashboardsV1CommonLogsAggregationIntf{apisDashboardsV1CommonLogsAggregationModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBys = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	apisDashboardsV1AstWidgetsLineChartQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryValueLogs)
// 	apisDashboardsV1AstWidgetsLineChartQueryModel.Logs = apisDashboardsV1AstWidgetsLineChartLogsQueryModel

// 	apisDashboardsV1AstWidgetsLineChartResolutionModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartResolution)
// 	apisDashboardsV1AstWidgetsLineChartResolutionModel.Interval = core.StringPtr("1m")
// 	apisDashboardsV1AstWidgetsLineChartResolutionModel.BucketsPresented = core.Int64Ptr(int64(100))

// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryDefinition)
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Query = apisDashboardsV1AstWidgetsLineChartQueryModel
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.SeriesNameTemplate = core.StringPtr("{{severity}}")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.SeriesCountLimit = core.StringPtr("10")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Unit = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ScaleType = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Name = core.StringPtr("CPU usage")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ColorScheme = core.StringPtr("classic")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Resolution = apisDashboardsV1AstWidgetsLineChartResolutionModel
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.DataModeType = core.StringPtr("high_unspecified")

// 	apisDashboardsV1AstWidgetsLineChartModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChart)
// 	apisDashboardsV1AstWidgetsLineChartModel.Legend = apisDashboardsV1AstWidgetsCommonLegendModel
// 	apisDashboardsV1AstWidgetsLineChartModel.Tooltip = apisDashboardsV1AstWidgetsLineChartTooltipModel
// 	apisDashboardsV1AstWidgetsLineChartModel.QueryDefinitions = []logsv0.ApisDashboardsV1AstWidgetsLineChartQueryDefinition{*apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel}

// 	apisDashboardsV1AstWidgetDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetDefinitionValueLineChart)
// 	apisDashboardsV1AstWidgetDefinitionModel.LineChart = apisDashboardsV1AstWidgetsLineChartModel

// 	model := new(logsv0.ApisDashboardsV1AstWidget)
// 	model.Href = core.StringPtr("testString")
// 	model.ID = apisDashboardsV1UUIDModel
// 	model.Title = core.StringPtr("Response time")
// 	model.Description = core.StringPtr("The average response time of the system")
// 	model.Definition = apisDashboardsV1AstWidgetDefinitionModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetDefinitionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLegendModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLegendModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsCommonLegendModel["columns"] = []string{"unspecified"}
// 		apisDashboardsV1AstWidgetsCommonLegendModel["group_by_query"] = true

// 		apisDashboardsV1AstWidgetsLineChartTooltipModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartTooltipModel["show_labels"] = true
// 		apisDashboardsV1AstWidgetsLineChartTooltipModel["type"] = "unspecified"

// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_by"] = []string{"testString"}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["aggregations"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_bys"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsLineChartQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartQueryModel["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartLogsQueryModel}

// 		apisDashboardsV1AstWidgetsLineChartResolutionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartResolutionModel["interval"] = "1m"
// 		apisDashboardsV1AstWidgetsLineChartResolutionModel["buckets_presented"] = int(100)

// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["series_name_template"] = "{{severity}}"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["series_count_limit"] = "10"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["unit"] = "unspecified"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["scale_type"] = "unspecified"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["name"] = "CPU usage"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["color_scheme"] = "classic"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["resolution"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartResolutionModel}
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["data_mode_type"] = "high_unspecified"

// 		apisDashboardsV1AstWidgetsLineChartModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartModel["legend"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLegendModel}
// 		apisDashboardsV1AstWidgetsLineChartModel["tooltip"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartTooltipModel}
// 		apisDashboardsV1AstWidgetsLineChartModel["query_definitions"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel}

// 		model := make(map[string]interface{})
// 		model["line_chart"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartModel}
// 		model["data_table"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableModel}
// 		model["gauge"] = []map[string]interface{}{apisDashboardsV1AstWidgetsGaugeModel}
// 		model["pie_chart"] = []map[string]interface{}{apisDashboardsV1AstWidgetsPieChartModel}
// 		model["bar_chart"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartModel}
// 		model["horizontal_bar_chart"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartModel}
// 		model["markdown"] = []map[string]interface{}{apisDashboardsV1AstWidgetsMarkdownModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLegendModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLegend)
// 	apisDashboardsV1AstWidgetsCommonLegendModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsCommonLegendModel.Columns = []string{"unspecified"}
// 	apisDashboardsV1AstWidgetsCommonLegendModel.GroupByQuery = core.BoolPtr(true)

// 	apisDashboardsV1AstWidgetsLineChartTooltipModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartTooltip)
// 	apisDashboardsV1AstWidgetsLineChartTooltipModel.ShowLabels = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsLineChartTooltipModel.Type = core.StringPtr("unspecified")

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartLogsQuery)
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBy = []string{"testString"}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Aggregations = []logsv0.ApisDashboardsV1CommonLogsAggregationIntf{apisDashboardsV1CommonLogsAggregationModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBys = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	apisDashboardsV1AstWidgetsLineChartQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryValueLogs)
// 	apisDashboardsV1AstWidgetsLineChartQueryModel.Logs = apisDashboardsV1AstWidgetsLineChartLogsQueryModel

// 	apisDashboardsV1AstWidgetsLineChartResolutionModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartResolution)
// 	apisDashboardsV1AstWidgetsLineChartResolutionModel.Interval = core.StringPtr("1m")
// 	apisDashboardsV1AstWidgetsLineChartResolutionModel.BucketsPresented = core.Int64Ptr(int64(100))

// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryDefinition)
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Query = apisDashboardsV1AstWidgetsLineChartQueryModel
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.SeriesNameTemplate = core.StringPtr("{{severity}}")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.SeriesCountLimit = core.StringPtr("10")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Unit = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ScaleType = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Name = core.StringPtr("CPU usage")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ColorScheme = core.StringPtr("classic")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Resolution = apisDashboardsV1AstWidgetsLineChartResolutionModel
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.DataModeType = core.StringPtr("high_unspecified")

// 	apisDashboardsV1AstWidgetsLineChartModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChart)
// 	apisDashboardsV1AstWidgetsLineChartModel.Legend = apisDashboardsV1AstWidgetsCommonLegendModel
// 	apisDashboardsV1AstWidgetsLineChartModel.Tooltip = apisDashboardsV1AstWidgetsLineChartTooltipModel
// 	apisDashboardsV1AstWidgetsLineChartModel.QueryDefinitions = []logsv0.ApisDashboardsV1AstWidgetsLineChartQueryDefinition{*apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetDefinition)
// 	model.LineChart = apisDashboardsV1AstWidgetsLineChartModel
// 	model.DataTable = apisDashboardsV1AstWidgetsDataTableModel
// 	model.Gauge = apisDashboardsV1AstWidgetsGaugeModel
// 	model.PieChart = apisDashboardsV1AstWidgetsPieChartModel
// 	model.BarChart = apisDashboardsV1AstWidgetsBarChartModel
// 	model.HorizontalBarChart = apisDashboardsV1AstWidgetsHorizontalBarChartModel
// 	model.Markdown = apisDashboardsV1AstWidgetsMarkdownModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetDefinitionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLegendModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLegendModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsCommonLegendModel["columns"] = []string{"unspecified"}
// 		apisDashboardsV1AstWidgetsCommonLegendModel["group_by_query"] = true

// 		apisDashboardsV1AstWidgetsLineChartTooltipModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartTooltipModel["show_labels"] = true
// 		apisDashboardsV1AstWidgetsLineChartTooltipModel["type"] = "unspecified"

// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_by"] = []string{"testString"}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["aggregations"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_bys"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsLineChartQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartQueryModel["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartLogsQueryModel}

// 		apisDashboardsV1AstWidgetsLineChartResolutionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartResolutionModel["interval"] = "1m"
// 		apisDashboardsV1AstWidgetsLineChartResolutionModel["buckets_presented"] = int(100)

// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["series_name_template"] = "{{severity}}"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["series_count_limit"] = "10"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["unit"] = "unspecified"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["scale_type"] = "unspecified"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["name"] = "CPU usage"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["color_scheme"] = "classic"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["resolution"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartResolutionModel}
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["data_mode_type"] = "high_unspecified"

// 		model := make(map[string]interface{})
// 		model["legend"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLegendModel}
// 		model["tooltip"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartTooltipModel}
// 		model["query_definitions"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLegendModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLegend)
// 	apisDashboardsV1AstWidgetsCommonLegendModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsCommonLegendModel.Columns = []string{"unspecified"}
// 	apisDashboardsV1AstWidgetsCommonLegendModel.GroupByQuery = core.BoolPtr(true)

// 	apisDashboardsV1AstWidgetsLineChartTooltipModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartTooltip)
// 	apisDashboardsV1AstWidgetsLineChartTooltipModel.ShowLabels = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsLineChartTooltipModel.Type = core.StringPtr("unspecified")

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartLogsQuery)
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBy = []string{"testString"}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Aggregations = []logsv0.ApisDashboardsV1CommonLogsAggregationIntf{apisDashboardsV1CommonLogsAggregationModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBys = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	apisDashboardsV1AstWidgetsLineChartQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryValueLogs)
// 	apisDashboardsV1AstWidgetsLineChartQueryModel.Logs = apisDashboardsV1AstWidgetsLineChartLogsQueryModel

// 	apisDashboardsV1AstWidgetsLineChartResolutionModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartResolution)
// 	apisDashboardsV1AstWidgetsLineChartResolutionModel.Interval = core.StringPtr("1m")
// 	apisDashboardsV1AstWidgetsLineChartResolutionModel.BucketsPresented = core.Int64Ptr(int64(100))

// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryDefinition)
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Query = apisDashboardsV1AstWidgetsLineChartQueryModel
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.SeriesNameTemplate = core.StringPtr("{{severity}}")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.SeriesCountLimit = core.StringPtr("10")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Unit = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ScaleType = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Name = core.StringPtr("CPU usage")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ColorScheme = core.StringPtr("classic")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Resolution = apisDashboardsV1AstWidgetsLineChartResolutionModel
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.DataModeType = core.StringPtr("high_unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsLineChart)
// 	model.Legend = apisDashboardsV1AstWidgetsCommonLegendModel
// 	model.Tooltip = apisDashboardsV1AstWidgetsLineChartTooltipModel
// 	model.QueryDefinitions = []logsv0.ApisDashboardsV1AstWidgetsLineChartQueryDefinition{*apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonLegendToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["is_visible"] = true
// 		model["columns"] = []string{"unspecified"}
// 		model["group_by_query"] = true

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsCommonLegend)
// 	model.IsVisible = core.BoolPtr(true)
// 	model.Columns = []string{"unspecified"}
// 	model.GroupByQuery = core.BoolPtr(true)

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonLegendToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartTooltipToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["show_labels"] = true
// 		model["type"] = "unspecified"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsLineChartTooltip)
// 	model.ShowLabels = core.BoolPtr(true)
// 	model.Type = core.StringPtr("unspecified")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartTooltipToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartQueryDefinitionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_by"] = []string{"testString"}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["aggregations"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_bys"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsLineChartQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartQueryModel["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartLogsQueryModel}

// 		apisDashboardsV1AstWidgetsLineChartResolutionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartResolutionModel["interval"] = "1m"
// 		apisDashboardsV1AstWidgetsLineChartResolutionModel["buckets_presented"] = int(100)

// 		model := make(map[string]interface{})
// 		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		model["query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartQueryModel}
// 		model["series_name_template"] = "{{severity}}"
// 		model["series_count_limit"] = "10"
// 		model["unit"] = "unspecified"
// 		model["scale_type"] = "unspecified"
// 		model["name"] = "CPU usage"
// 		model["is_visible"] = true
// 		model["color_scheme"] = "classic"
// 		model["resolution"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartResolutionModel}
// 		model["data_mode_type"] = "high_unspecified"

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartLogsQuery)
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBy = []string{"testString"}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Aggregations = []logsv0.ApisDashboardsV1CommonLogsAggregationIntf{apisDashboardsV1CommonLogsAggregationModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBys = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	apisDashboardsV1AstWidgetsLineChartQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryValueLogs)
// 	apisDashboardsV1AstWidgetsLineChartQueryModel.Logs = apisDashboardsV1AstWidgetsLineChartLogsQueryModel

// 	apisDashboardsV1AstWidgetsLineChartResolutionModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartResolution)
// 	apisDashboardsV1AstWidgetsLineChartResolutionModel.Interval = core.StringPtr("1m")
// 	apisDashboardsV1AstWidgetsLineChartResolutionModel.BucketsPresented = core.Int64Ptr(int64(100))

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryDefinition)
// 	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	model.Query = apisDashboardsV1AstWidgetsLineChartQueryModel
// 	model.SeriesNameTemplate = core.StringPtr("{{severity}}")
// 	model.SeriesCountLimit = core.StringPtr("10")
// 	model.Unit = core.StringPtr("unspecified")
// 	model.ScaleType = core.StringPtr("unspecified")
// 	model.Name = core.StringPtr("CPU usage")
// 	model.IsVisible = core.BoolPtr(true)
// 	model.ColorScheme = core.StringPtr("classic")
// 	model.Resolution = apisDashboardsV1AstWidgetsLineChartResolutionModel
// 	model.DataModeType = core.StringPtr("high_unspecified")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartQueryDefinitionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_by"] = []string{"testString"}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["aggregations"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_bys"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartLogsQueryModel}
// 		model["metrics"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartMetricsQueryModel}
// 		model["spans"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartSpansQueryModel}
// 		model["dataprime"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartDataprimeQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartLogsQuery)
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBy = []string{"testString"}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Aggregations = []logsv0.ApisDashboardsV1CommonLogsAggregationIntf{apisDashboardsV1CommonLogsAggregationModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBys = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQuery)
// 	model.Logs = apisDashboardsV1AstWidgetsLineChartLogsQueryModel
// 	model.Metrics = apisDashboardsV1AstWidgetsLineChartMetricsQueryModel
// 	model.Spans = apisDashboardsV1AstWidgetsLineChartSpansQueryModel
// 	model.Dataprime = apisDashboardsV1AstWidgetsLineChartDataprimeQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartLogsQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		model["group_by"] = []string{"testString"}
// 		model["aggregations"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		model["group_bys"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsLineChartLogsQuery)
// 	model.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	model.GroupBy = []string{"testString"}
// 	model.Aggregations = []logsv0.ApisDashboardsV1CommonLogsAggregationIntf{apisDashboardsV1CommonLogsAggregationModel}
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	model.GroupBys = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartLogsQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonLuceneQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["value"] = "coralogix.metadata.applicationName:"production""

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	model.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonLuceneQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		model := make(map[string]interface{})
// 		model["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}
// 		model["count_distinct"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountDistinctModel}
// 		model["sum"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationSumModel}
// 		model["average"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationAverageModel}
// 		model["min"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationMinModel}
// 		model["max"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationMaxModel}
// 		model["percentile"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationPercentileModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	model := new(logsv0.ApisDashboardsV1CommonLogsAggregation)
// 	model.Count = apisDashboardsV1CommonLogsAggregationCountModel
// 	model.CountDistinct = apisDashboardsV1CommonLogsAggregationCountDistinctModel
// 	model.Sum = apisDashboardsV1CommonLogsAggregationSumModel
// 	model.Average = apisDashboardsV1CommonLogsAggregationAverageModel
// 	model.Min = apisDashboardsV1CommonLogsAggregationMinModel
// 	model.Max = apisDashboardsV1CommonLogsAggregationMaxModel
// 	model.Percentile = apisDashboardsV1CommonLogsAggregationPercentileModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationCountToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationCountToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationCountDistinctToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1CommonLogsAggregationCountDistinct)
// 	model.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationCountDistinctToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonObservationFieldToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["keypath"] = []string{"testString"}
// 		model["scope"] = "unspecified"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	model.Keypath = []string{"testString"}
// 	model.Scope = core.StringPtr("unspecified")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonObservationFieldToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationSumToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1CommonLogsAggregationSum)
// 	model.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationSumToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationAverageToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1CommonLogsAggregationAverage)
// 	model.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationAverageToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationMinToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1CommonLogsAggregationMin)
// 	model.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationMinToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationMaxToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1CommonLogsAggregationMax)
// 	model.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationMaxToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationPercentileToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["percent"] = float64(72.5)
// 		model["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1CommonLogsAggregationPercentile)
// 	model.Percent = core.Float64Ptr(float64(72.5))
// 	model.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationPercentileToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationValueCountToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		model := make(map[string]interface{})
// 		model["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	model := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	model.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationValueCountToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationValueCountDistinctToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1CommonLogsAggregationCountDistinctModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationCountDistinctModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["count_distinct"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountDistinctModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonLogsAggregationCountDistinctModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCountDistinct)
// 	apisDashboardsV1CommonLogsAggregationCountDistinctModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCountDistinct)
// 	model.CountDistinct = apisDashboardsV1CommonLogsAggregationCountDistinctModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationValueCountDistinctToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationValueSumToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1CommonLogsAggregationSumModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationSumModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["sum"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationSumModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonLogsAggregationSumModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationSum)
// 	apisDashboardsV1CommonLogsAggregationSumModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueSum)
// 	model.Sum = apisDashboardsV1CommonLogsAggregationSumModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationValueSumToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationValueAverageToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1CommonLogsAggregationAverageModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationAverageModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["average"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationAverageModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonLogsAggregationAverageModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationAverage)
// 	apisDashboardsV1CommonLogsAggregationAverageModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueAverage)
// 	model.Average = apisDashboardsV1CommonLogsAggregationAverageModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationValueAverageToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationValueMinToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1CommonLogsAggregationMinModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationMinModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["min"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationMinModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonLogsAggregationMinModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationMin)
// 	apisDashboardsV1CommonLogsAggregationMinModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueMin)
// 	model.Min = apisDashboardsV1CommonLogsAggregationMinModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationValueMinToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationValueMaxToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1CommonLogsAggregationMaxModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationMaxModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["max"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationMaxModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonLogsAggregationMaxModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationMax)
// 	apisDashboardsV1CommonLogsAggregationMaxModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueMax)
// 	model.Max = apisDashboardsV1CommonLogsAggregationMaxModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationValueMaxToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationValuePercentileToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1CommonLogsAggregationPercentileModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationPercentileModel["percent"] = float64(72.5)
// 		apisDashboardsV1CommonLogsAggregationPercentileModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["percentile"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationPercentileModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonLogsAggregationPercentileModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationPercentile)
// 	apisDashboardsV1CommonLogsAggregationPercentileModel.Percent = core.Float64Ptr(float64(72.5))
// 	apisDashboardsV1CommonLogsAggregationPercentileModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1CommonLogsAggregationValuePercentile)
// 	model.Percentile = apisDashboardsV1CommonLogsAggregationPercentileModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonLogsAggregationValuePercentileToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterLogsFilterToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		model["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	model.Operator = apisDashboardsV1AstFilterOperatorModel
// 	model.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterLogsFilterToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterOperatorToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		model := make(map[string]interface{})
// 		model["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}
// 		model["not_equals"] = []map[string]interface{}{apisDashboardsV1AstFilterNotEqualsModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	model := new(logsv0.ApisDashboardsV1AstFilterOperator)
// 	model.Equals = apisDashboardsV1AstFilterEqualsModel
// 	model.NotEquals = apisDashboardsV1AstFilterNotEqualsModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterOperatorToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterEqualsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		model := make(map[string]interface{})
// 		model["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	model := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	model.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterEqualsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterEqualsSelectionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		model := make(map[string]interface{})
// 		model["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}
// 		model["list"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionListSelectionModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	model := new(logsv0.ApisDashboardsV1AstFilterEqualsSelection)
// 	model.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel
// 	model.List = apisDashboardsV1AstFilterEqualsSelectionListSelectionModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterEqualsSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterEqualsSelectionAllSelectionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterEqualsSelectionAllSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterEqualsSelectionListSelectionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["values"] = []string{"testString"}

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionListSelection)
// 	model.Values = []string{"testString"}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterEqualsSelectionListSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterEqualsSelectionValueAllToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		model := make(map[string]interface{})
// 		model["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	model := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	model.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterEqualsSelectionValueAllToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterEqualsSelectionValueListToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstFilterEqualsSelectionListSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionListSelectionModel["values"] = []string{"testString"}

// 		model := make(map[string]interface{})
// 		model["list"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionListSelectionModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstFilterEqualsSelectionListSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionListSelection)
// 	apisDashboardsV1AstFilterEqualsSelectionListSelectionModel.Values = []string{"testString"}

// 	model := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueList)
// 	model.List = apisDashboardsV1AstFilterEqualsSelectionListSelectionModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterEqualsSelectionValueListToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterNotEqualsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstFilterNotEqualsSelectionListSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterNotEqualsSelectionListSelectionModel["values"] = []string{"testString"}

// 		apisDashboardsV1AstFilterNotEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterNotEqualsSelectionModel["list"] = []map[string]interface{}{apisDashboardsV1AstFilterNotEqualsSelectionListSelectionModel}

// 		model := make(map[string]interface{})
// 		model["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterNotEqualsSelectionModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstFilterNotEqualsSelectionListSelectionModel := new(logsv0.ApisDashboardsV1AstFilterNotEqualsSelectionListSelection)
// 	apisDashboardsV1AstFilterNotEqualsSelectionListSelectionModel.Values = []string{"testString"}

// 	apisDashboardsV1AstFilterNotEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterNotEqualsSelection)
// 	apisDashboardsV1AstFilterNotEqualsSelectionModel.List = apisDashboardsV1AstFilterNotEqualsSelectionListSelectionModel

// 	model := new(logsv0.ApisDashboardsV1AstFilterNotEquals)
// 	model.Selection = apisDashboardsV1AstFilterNotEqualsSelectionModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterNotEqualsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterNotEqualsSelectionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstFilterNotEqualsSelectionListSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterNotEqualsSelectionListSelectionModel["values"] = []string{"testString"}

// 		model := make(map[string]interface{})
// 		model["list"] = []map[string]interface{}{apisDashboardsV1AstFilterNotEqualsSelectionListSelectionModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstFilterNotEqualsSelectionListSelectionModel := new(logsv0.ApisDashboardsV1AstFilterNotEqualsSelectionListSelection)
// 	apisDashboardsV1AstFilterNotEqualsSelectionListSelectionModel.Values = []string{"testString"}

// 	model := new(logsv0.ApisDashboardsV1AstFilterNotEqualsSelection)
// 	model.List = apisDashboardsV1AstFilterNotEqualsSelectionListSelectionModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterNotEqualsSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterNotEqualsSelectionListSelectionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["values"] = []string{"testString"}

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstFilterNotEqualsSelectionListSelection)
// 	model.Values = []string{"testString"}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterNotEqualsSelectionListSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterOperatorValueEqualsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		model := make(map[string]interface{})
// 		model["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	model := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	model.Equals = apisDashboardsV1AstFilterEqualsModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterOperatorValueEqualsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterOperatorValueNotEqualsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstFilterNotEqualsSelectionListSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterNotEqualsSelectionListSelectionModel["values"] = []string{"testString"}

// 		apisDashboardsV1AstFilterNotEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterNotEqualsSelectionModel["list"] = []map[string]interface{}{apisDashboardsV1AstFilterNotEqualsSelectionListSelectionModel}

// 		apisDashboardsV1AstFilterNotEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterNotEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterNotEqualsSelectionModel}

// 		model := make(map[string]interface{})
// 		model["not_equals"] = []map[string]interface{}{apisDashboardsV1AstFilterNotEqualsModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstFilterNotEqualsSelectionListSelectionModel := new(logsv0.ApisDashboardsV1AstFilterNotEqualsSelectionListSelection)
// 	apisDashboardsV1AstFilterNotEqualsSelectionListSelectionModel.Values = []string{"testString"}

// 	apisDashboardsV1AstFilterNotEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterNotEqualsSelection)
// 	apisDashboardsV1AstFilterNotEqualsSelectionModel.List = apisDashboardsV1AstFilterNotEqualsSelectionListSelectionModel

// 	apisDashboardsV1AstFilterNotEqualsModel := new(logsv0.ApisDashboardsV1AstFilterNotEquals)
// 	apisDashboardsV1AstFilterNotEqualsModel.Selection = apisDashboardsV1AstFilterNotEqualsSelectionModel

// 	model := new(logsv0.ApisDashboardsV1AstFilterOperatorValueNotEquals)
// 	model.NotEquals = apisDashboardsV1AstFilterNotEqualsModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterOperatorValueNotEqualsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartMetricsQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterMetricsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterMetricsFilterModel["label"] = "service"
// 		apisDashboardsV1AstFilterMetricsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		model := make(map[string]interface{})
// 		model["promql_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonPromQlQueryModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterMetricsFilterModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonPromQlQuery)
// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterMetricsFilterModel := new(logsv0.ApisDashboardsV1AstFilterMetricsFilter)
// 	apisDashboardsV1AstFilterMetricsFilterModel.Label = core.StringPtr("service")
// 	apisDashboardsV1AstFilterMetricsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsLineChartMetricsQuery)
// 	model.PromqlQuery = apisDashboardsV1AstWidgetsCommonPromQlQueryModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterMetricsFilter{*apisDashboardsV1AstFilterMetricsFilterModel}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartMetricsQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonPromQlQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["value"] = "sum(up)"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsCommonPromQlQuery)
// 	model.Value = core.StringPtr("sum(up)")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonPromQlQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterMetricsFilterToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		model := make(map[string]interface{})
// 		model["label"] = "service"
// 		model["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	model := new(logsv0.ApisDashboardsV1AstFilterMetricsFilter)
// 	model.Label = core.StringPtr("service")
// 	model.Operator = apisDashboardsV1AstFilterOperatorModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterMetricsFilterToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartSpansQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonSpanFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpanFieldModel["metadata_field"] = "unspecified"

// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["metric_field"] = "unspecified"
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["aggregation_type"] = "unspecified"

// 		apisDashboardsV1CommonSpansAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationModel["metric_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationMetricAggregationModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterSpansFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSpansFilterModel["field"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstFilterSpansFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		model := make(map[string]interface{})
// 		model["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		model["group_by"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		model["aggregations"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSpansFilterModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonSpanFieldModel := new(logsv0.ApisDashboardsV1CommonSpanFieldValueMetadataField)
// 	apisDashboardsV1CommonSpanFieldModel.MetadataField = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.MetricField = core.StringPtr("unspecified")
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.AggregationType = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonSpansAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationModel.MetricAggregation = apisDashboardsV1CommonSpansAggregationMetricAggregationModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterSpansFilterModel := new(logsv0.ApisDashboardsV1AstFilterSpansFilter)
// 	apisDashboardsV1AstFilterSpansFilterModel.Field = apisDashboardsV1CommonSpanFieldModel
// 	apisDashboardsV1AstFilterSpansFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsLineChartSpansQuery)
// 	model.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	model.GroupBy = []logsv0.ApisDashboardsV1CommonSpanFieldIntf{apisDashboardsV1CommonSpanFieldModel}
// 	model.Aggregations = []logsv0.ApisDashboardsV1CommonSpansAggregationIntf{apisDashboardsV1CommonSpansAggregationModel}
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterSpansFilter{*apisDashboardsV1AstFilterSpansFilterModel}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartSpansQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonSpanFieldToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["metadata_field"] = "unspecified"
// 		model["tag_field"] = "http.status_code"
// 		model["process_tag_field"] = "process.pid"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1CommonSpanField)
// 	model.MetadataField = core.StringPtr("unspecified")
// 	model.TagField = core.StringPtr("http.status_code")
// 	model.ProcessTagField = core.StringPtr("process.pid")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonSpanFieldToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonSpanFieldValueMetadataFieldToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["metadata_field"] = "unspecified"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1CommonSpanFieldValueMetadataField)
// 	model.MetadataField = core.StringPtr("unspecified")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonSpanFieldValueMetadataFieldToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonSpanFieldValueTagFieldToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["tag_field"] = "http.status_code"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1CommonSpanFieldValueTagField)
// 	model.TagField = core.StringPtr("http.status_code")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonSpanFieldValueTagFieldToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonSpanFieldValueProcessTagFieldToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["process_tag_field"] = "process.pid"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1CommonSpanFieldValueProcessTagField)
// 	model.ProcessTagField = core.StringPtr("process.pid")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonSpanFieldValueProcessTagFieldToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonSpansAggregationToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["metric_field"] = "unspecified"
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["aggregation_type"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["metric_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationMetricAggregationModel}
// 		model["dimension_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationDimensionAggregationModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.MetricField = core.StringPtr("unspecified")
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.AggregationType = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1CommonSpansAggregation)
// 	model.MetricAggregation = apisDashboardsV1CommonSpansAggregationMetricAggregationModel
// 	model.DimensionAggregation = apisDashboardsV1CommonSpansAggregationDimensionAggregationModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonSpansAggregationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonSpansAggregationMetricAggregationToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["metric_field"] = "unspecified"
// 		model["aggregation_type"] = "unspecified"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1CommonSpansAggregationMetricAggregation)
// 	model.MetricField = core.StringPtr("unspecified")
// 	model.AggregationType = core.StringPtr("unspecified")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonSpansAggregationMetricAggregationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonSpansAggregationDimensionAggregationToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["dimension_field"] = "unspecified"
// 		model["aggregation_type"] = "unspecified"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1CommonSpansAggregationDimensionAggregation)
// 	model.DimensionField = core.StringPtr("unspecified")
// 	model.AggregationType = core.StringPtr("unspecified")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonSpansAggregationDimensionAggregationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonSpansAggregationAggregationMetricAggregationToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["metric_field"] = "unspecified"
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["aggregation_type"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["metric_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationMetricAggregationModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.MetricField = core.StringPtr("unspecified")
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.AggregationType = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1CommonSpansAggregationAggregationMetricAggregation)
// 	model.MetricAggregation = apisDashboardsV1CommonSpansAggregationMetricAggregationModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonSpansAggregationAggregationMetricAggregationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonSpansAggregationAggregationDimensionAggregationToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonSpansAggregationDimensionAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationDimensionAggregationModel["dimension_field"] = "unspecified"
// 		apisDashboardsV1CommonSpansAggregationDimensionAggregationModel["aggregation_type"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["dimension_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationDimensionAggregationModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonSpansAggregationDimensionAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationDimensionAggregation)
// 	apisDashboardsV1CommonSpansAggregationDimensionAggregationModel.DimensionField = core.StringPtr("unspecified")
// 	apisDashboardsV1CommonSpansAggregationDimensionAggregationModel.AggregationType = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1CommonSpansAggregationAggregationDimensionAggregation)
// 	model.DimensionAggregation = apisDashboardsV1CommonSpansAggregationDimensionAggregationModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonSpansAggregationAggregationDimensionAggregationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterSpansFilterToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonSpanFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpanFieldModel["metadata_field"] = "unspecified"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		model := make(map[string]interface{})
// 		model["field"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		model["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonSpanFieldModel := new(logsv0.ApisDashboardsV1CommonSpanFieldValueMetadataField)
// 	apisDashboardsV1CommonSpanFieldModel.MetadataField = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	model := new(logsv0.ApisDashboardsV1AstFilterSpansFilter)
// 	model.Field = apisDashboardsV1CommonSpanFieldModel
// 	model.Operator = apisDashboardsV1AstFilterOperatorModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterSpansFilterToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartDataprimeQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonDataprimeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonDataprimeQueryModel["text"] = "source logs"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstFilterSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSourceModel["logs"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}

// 		model := make(map[string]interface{})
// 		model["dataprime_query"] = []map[string]interface{}{apisDashboardsV1CommonDataprimeQueryModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSourceModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonDataprimeQueryModel := new(logsv0.ApisDashboardsV1CommonDataprimeQuery)
// 	apisDashboardsV1CommonDataprimeQueryModel.Text = core.StringPtr("source logs")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstFilterSourceModel := new(logsv0.ApisDashboardsV1AstFilterSourceValueLogs)
// 	apisDashboardsV1AstFilterSourceModel.Logs = apisDashboardsV1AstFilterLogsFilterModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsLineChartDataprimeQuery)
// 	model.DataprimeQuery = apisDashboardsV1CommonDataprimeQueryModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterSourceIntf{apisDashboardsV1AstFilterSourceModel}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartDataprimeQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonDataprimeQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["text"] = "source logs"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1CommonDataprimeQuery)
// 	model.Text = core.StringPtr("source logs")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonDataprimeQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterSourceToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["logs"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		model["spans"] = []map[string]interface{}{apisDashboardsV1AstFilterSpansFilterModel}
// 		model["metrics"] = []map[string]interface{}{apisDashboardsV1AstFilterMetricsFilterModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstFilterSource)
// 	model.Logs = apisDashboardsV1AstFilterLogsFilterModel
// 	model.Spans = apisDashboardsV1AstFilterSpansFilterModel
// 	model.Metrics = apisDashboardsV1AstFilterMetricsFilterModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterSourceToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterSourceValueLogsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["logs"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstFilterSourceValueLogs)
// 	model.Logs = apisDashboardsV1AstFilterLogsFilterModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterSourceValueLogsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterSourceValueSpansToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonSpanFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpanFieldModel["metadata_field"] = "unspecified"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterSpansFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSpansFilterModel["field"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstFilterSpansFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		model := make(map[string]interface{})
// 		model["spans"] = []map[string]interface{}{apisDashboardsV1AstFilterSpansFilterModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonSpanFieldModel := new(logsv0.ApisDashboardsV1CommonSpanFieldValueMetadataField)
// 	apisDashboardsV1CommonSpanFieldModel.MetadataField = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterSpansFilterModel := new(logsv0.ApisDashboardsV1AstFilterSpansFilter)
// 	apisDashboardsV1AstFilterSpansFilterModel.Field = apisDashboardsV1CommonSpanFieldModel
// 	apisDashboardsV1AstFilterSpansFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	model := new(logsv0.ApisDashboardsV1AstFilterSourceValueSpans)
// 	model.Spans = apisDashboardsV1AstFilterSpansFilterModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterSourceValueSpansToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterSourceValueMetricsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterMetricsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterMetricsFilterModel["label"] = "service"
// 		apisDashboardsV1AstFilterMetricsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		model := make(map[string]interface{})
// 		model["metrics"] = []map[string]interface{}{apisDashboardsV1AstFilterMetricsFilterModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterMetricsFilterModel := new(logsv0.ApisDashboardsV1AstFilterMetricsFilter)
// 	apisDashboardsV1AstFilterMetricsFilterModel.Label = core.StringPtr("service")
// 	apisDashboardsV1AstFilterMetricsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	model := new(logsv0.ApisDashboardsV1AstFilterSourceValueMetrics)
// 	model.Metrics = apisDashboardsV1AstFilterMetricsFilterModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterSourceValueMetricsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartQueryValueLogsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_by"] = []string{"testString"}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["aggregations"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_bys"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartLogsQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartLogsQuery)
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBy = []string{"testString"}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Aggregations = []logsv0.ApisDashboardsV1CommonLogsAggregationIntf{apisDashboardsV1CommonLogsAggregationModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBys = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryValueLogs)
// 	model.Logs = apisDashboardsV1AstWidgetsLineChartLogsQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartQueryValueLogsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartQueryValueMetricsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterMetricsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterMetricsFilterModel["label"] = "service"
// 		apisDashboardsV1AstFilterMetricsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		apisDashboardsV1AstWidgetsLineChartMetricsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartMetricsQueryModel["promql_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonPromQlQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartMetricsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterMetricsFilterModel}

// 		model := make(map[string]interface{})
// 		model["metrics"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartMetricsQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonPromQlQuery)
// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterMetricsFilterModel := new(logsv0.ApisDashboardsV1AstFilterMetricsFilter)
// 	apisDashboardsV1AstFilterMetricsFilterModel.Label = core.StringPtr("service")
// 	apisDashboardsV1AstFilterMetricsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	apisDashboardsV1AstWidgetsLineChartMetricsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartMetricsQuery)
// 	apisDashboardsV1AstWidgetsLineChartMetricsQueryModel.PromqlQuery = apisDashboardsV1AstWidgetsCommonPromQlQueryModel
// 	apisDashboardsV1AstWidgetsLineChartMetricsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterMetricsFilter{*apisDashboardsV1AstFilterMetricsFilterModel}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryValueMetrics)
// 	model.Metrics = apisDashboardsV1AstWidgetsLineChartMetricsQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartQueryValueMetricsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartQueryValueSpansToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonSpanFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpanFieldModel["metadata_field"] = "unspecified"

// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["metric_field"] = "unspecified"
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["aggregation_type"] = "unspecified"

// 		apisDashboardsV1CommonSpansAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationModel["metric_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationMetricAggregationModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterSpansFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSpansFilterModel["field"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstFilterSpansFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		apisDashboardsV1AstWidgetsLineChartSpansQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartSpansQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartSpansQueryModel["group_by"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstWidgetsLineChartSpansQueryModel["aggregations"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationModel}
// 		apisDashboardsV1AstWidgetsLineChartSpansQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSpansFilterModel}

// 		model := make(map[string]interface{})
// 		model["spans"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartSpansQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonSpanFieldModel := new(logsv0.ApisDashboardsV1CommonSpanFieldValueMetadataField)
// 	apisDashboardsV1CommonSpanFieldModel.MetadataField = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.MetricField = core.StringPtr("unspecified")
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.AggregationType = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonSpansAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationModel.MetricAggregation = apisDashboardsV1CommonSpansAggregationMetricAggregationModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterSpansFilterModel := new(logsv0.ApisDashboardsV1AstFilterSpansFilter)
// 	apisDashboardsV1AstFilterSpansFilterModel.Field = apisDashboardsV1CommonSpanFieldModel
// 	apisDashboardsV1AstFilterSpansFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	apisDashboardsV1AstWidgetsLineChartSpansQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartSpansQuery)
// 	apisDashboardsV1AstWidgetsLineChartSpansQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsLineChartSpansQueryModel.GroupBy = []logsv0.ApisDashboardsV1CommonSpanFieldIntf{apisDashboardsV1CommonSpanFieldModel}
// 	apisDashboardsV1AstWidgetsLineChartSpansQueryModel.Aggregations = []logsv0.ApisDashboardsV1CommonSpansAggregationIntf{apisDashboardsV1CommonSpansAggregationModel}
// 	apisDashboardsV1AstWidgetsLineChartSpansQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterSpansFilter{*apisDashboardsV1AstFilterSpansFilterModel}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryValueSpans)
// 	model.Spans = apisDashboardsV1AstWidgetsLineChartSpansQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartQueryValueSpansToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartQueryValueDataprimeToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonDataprimeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonDataprimeQueryModel["text"] = "source logs"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstFilterSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSourceModel["logs"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}

// 		apisDashboardsV1AstWidgetsLineChartDataprimeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartDataprimeQueryModel["dataprime_query"] = []map[string]interface{}{apisDashboardsV1CommonDataprimeQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartDataprimeQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSourceModel}

// 		model := make(map[string]interface{})
// 		model["dataprime"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartDataprimeQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonDataprimeQueryModel := new(logsv0.ApisDashboardsV1CommonDataprimeQuery)
// 	apisDashboardsV1CommonDataprimeQueryModel.Text = core.StringPtr("source logs")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstFilterSourceModel := new(logsv0.ApisDashboardsV1AstFilterSourceValueLogs)
// 	apisDashboardsV1AstFilterSourceModel.Logs = apisDashboardsV1AstFilterLogsFilterModel

// 	apisDashboardsV1AstWidgetsLineChartDataprimeQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartDataprimeQuery)
// 	apisDashboardsV1AstWidgetsLineChartDataprimeQueryModel.DataprimeQuery = apisDashboardsV1CommonDataprimeQueryModel
// 	apisDashboardsV1AstWidgetsLineChartDataprimeQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterSourceIntf{apisDashboardsV1AstFilterSourceModel}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryValueDataprime)
// 	model.Dataprime = apisDashboardsV1AstWidgetsLineChartDataprimeQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartQueryValueDataprimeToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartResolutionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["interval"] = "1m"
// 		model["buckets_presented"] = int(100)

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsLineChartResolution)
// 	model.Interval = core.StringPtr("1m")
// 	model.BucketsPresented = core.Int64Ptr(int64(100))

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsLineChartResolutionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["name"] = "count"
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}

// 		apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel["aggregations"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel}
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel["group_bys"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsDataTableLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryModel["grouping"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel}

// 		apisDashboardsV1AstWidgetsDataTableQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableQueryModel["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableLogsQueryModel}

// 		apisDashboardsV1AstWidgetsDataTableColumnModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableColumnModel["field"] = "coralogix.metadata.applicationName"
// 		apisDashboardsV1AstWidgetsDataTableColumnModel["width"] = int(100)

// 		apisDashboardsV1CommonOrderingFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonOrderingFieldModel["field"] = "testString"
// 		apisDashboardsV1CommonOrderingFieldModel["order_direction"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableQueryModel}
// 		model["results_per_page"] = int(10)
// 		model["row_style"] = "unspecified"
// 		model["columns"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableColumnModel}
// 		model["order_by"] = []map[string]interface{}{apisDashboardsV1CommonOrderingFieldModel}
// 		model["data_mode_type"] = "high_unspecified"

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryAggregation)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.Name = core.StringPtr("count")
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.Aggregation = apisDashboardsV1CommonLogsAggregationModel

// 	apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryGrouping)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel.Aggregations = []logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryAggregation{*apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel}
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel.GroupBys = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	apisDashboardsV1AstWidgetsDataTableLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQuery)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryModel.Grouping = apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel

// 	apisDashboardsV1AstWidgetsDataTableQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableQueryValueLogs)
// 	apisDashboardsV1AstWidgetsDataTableQueryModel.Logs = apisDashboardsV1AstWidgetsDataTableLogsQueryModel

// 	apisDashboardsV1AstWidgetsDataTableColumnModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableColumn)
// 	apisDashboardsV1AstWidgetsDataTableColumnModel.Field = core.StringPtr("coralogix.metadata.applicationName")
// 	apisDashboardsV1AstWidgetsDataTableColumnModel.Width = core.Int64Ptr(int64(100))

// 	apisDashboardsV1CommonOrderingFieldModel := new(logsv0.ApisDashboardsV1CommonOrderingField)
// 	apisDashboardsV1CommonOrderingFieldModel.Field = core.StringPtr("testString")
// 	apisDashboardsV1CommonOrderingFieldModel.OrderDirection = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsDataTable)
// 	model.Query = apisDashboardsV1AstWidgetsDataTableQueryModel
// 	model.ResultsPerPage = core.Int64Ptr(int64(10))
// 	model.RowStyle = core.StringPtr("unspecified")
// 	model.Columns = []logsv0.ApisDashboardsV1AstWidgetsDataTableColumn{*apisDashboardsV1AstWidgetsDataTableColumnModel}
// 	model.OrderBy = apisDashboardsV1CommonOrderingFieldModel
// 	model.DataModeType = core.StringPtr("high_unspecified")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["name"] = "count"
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}

// 		apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel["aggregations"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel}
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel["group_bys"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsDataTableLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryModel["grouping"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel}

// 		model := make(map[string]interface{})
// 		model["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableLogsQueryModel}
// 		model["spans"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableSpansQueryModel}
// 		model["metrics"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableMetricsQueryModel}
// 		model["dataprime"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableDataprimeQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryAggregation)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.Name = core.StringPtr("count")
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.Aggregation = apisDashboardsV1CommonLogsAggregationModel

// 	apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryGrouping)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel.Aggregations = []logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryAggregation{*apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel}
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel.GroupBys = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	apisDashboardsV1AstWidgetsDataTableLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQuery)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryModel.Grouping = apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsDataTableQuery)
// 	model.Logs = apisDashboardsV1AstWidgetsDataTableLogsQueryModel
// 	model.Spans = apisDashboardsV1AstWidgetsDataTableSpansQueryModel
// 	model.Metrics = apisDashboardsV1AstWidgetsDataTableMetricsQueryModel
// 	model.Dataprime = apisDashboardsV1AstWidgetsDataTableDataprimeQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableLogsQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["name"] = "count"
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}

// 		apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel["aggregations"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel}
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel["group_bys"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		model["grouping"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryAggregation)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.Name = core.StringPtr("count")
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.Aggregation = apisDashboardsV1CommonLogsAggregationModel

// 	apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryGrouping)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel.Aggregations = []logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryAggregation{*apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel}
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel.GroupBys = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQuery)
// 	model.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	model.Grouping = apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableLogsQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableLogsQueryGroupingToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["name"] = "count"
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["aggregations"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel}
// 		model["group_bys"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryAggregation)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.Name = core.StringPtr("count")
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.Aggregation = apisDashboardsV1CommonLogsAggregationModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryGrouping)
// 	model.Aggregations = []logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryAggregation{*apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel}
// 	model.GroupBys = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableLogsQueryGroupingToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableLogsQueryAggregationToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		model := make(map[string]interface{})
// 		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		model["name"] = "count"
// 		model["is_visible"] = true
// 		model["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryAggregation)
// 	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	model.Name = core.StringPtr("count")
// 	model.IsVisible = core.BoolPtr(true)
// 	model.Aggregation = apisDashboardsV1CommonLogsAggregationModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableLogsQueryAggregationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableSpansQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonSpanFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpanFieldModel["metadata_field"] = "unspecified"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterSpansFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSpansFilterModel["field"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstFilterSpansFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["metric_field"] = "unspecified"
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["aggregation_type"] = "unspecified"

// 		apisDashboardsV1CommonSpansAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationModel["metric_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationMetricAggregationModel}

// 		apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel["name"] = "count"
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationModel}

// 		apisDashboardsV1AstWidgetsDataTableSpansQueryGroupingModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryGroupingModel["group_by"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryGroupingModel["aggregations"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel}

// 		model := make(map[string]interface{})
// 		model["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSpansFilterModel}
// 		model["grouping"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableSpansQueryGroupingModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonSpanFieldModel := new(logsv0.ApisDashboardsV1CommonSpanFieldValueMetadataField)
// 	apisDashboardsV1CommonSpanFieldModel.MetadataField = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterSpansFilterModel := new(logsv0.ApisDashboardsV1AstFilterSpansFilter)
// 	apisDashboardsV1AstFilterSpansFilterModel.Field = apisDashboardsV1CommonSpanFieldModel
// 	apisDashboardsV1AstFilterSpansFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.MetricField = core.StringPtr("unspecified")
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.AggregationType = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonSpansAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationModel.MetricAggregation = apisDashboardsV1CommonSpansAggregationMetricAggregationModel

// 	apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableSpansQueryAggregation)
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel.Name = core.StringPtr("count")
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel.Aggregation = apisDashboardsV1CommonSpansAggregationModel

// 	apisDashboardsV1AstWidgetsDataTableSpansQueryGroupingModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableSpansQueryGrouping)
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryGroupingModel.GroupBy = []logsv0.ApisDashboardsV1CommonSpanFieldIntf{apisDashboardsV1CommonSpanFieldModel}
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryGroupingModel.Aggregations = []logsv0.ApisDashboardsV1AstWidgetsDataTableSpansQueryAggregation{*apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsDataTableSpansQuery)
// 	model.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterSpansFilter{*apisDashboardsV1AstFilterSpansFilterModel}
// 	model.Grouping = apisDashboardsV1AstWidgetsDataTableSpansQueryGroupingModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableSpansQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableSpansQueryGroupingToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonSpanFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpanFieldModel["metadata_field"] = "unspecified"

// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["metric_field"] = "unspecified"
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["aggregation_type"] = "unspecified"

// 		apisDashboardsV1CommonSpansAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationModel["metric_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationMetricAggregationModel}

// 		apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel["name"] = "count"
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationModel}

// 		model := make(map[string]interface{})
// 		model["group_by"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		model["aggregations"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonSpanFieldModel := new(logsv0.ApisDashboardsV1CommonSpanFieldValueMetadataField)
// 	apisDashboardsV1CommonSpanFieldModel.MetadataField = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.MetricField = core.StringPtr("unspecified")
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.AggregationType = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonSpansAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationModel.MetricAggregation = apisDashboardsV1CommonSpansAggregationMetricAggregationModel

// 	apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableSpansQueryAggregation)
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel.Name = core.StringPtr("count")
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel.Aggregation = apisDashboardsV1CommonSpansAggregationModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsDataTableSpansQueryGrouping)
// 	model.GroupBy = []logsv0.ApisDashboardsV1CommonSpanFieldIntf{apisDashboardsV1CommonSpanFieldModel}
// 	model.Aggregations = []logsv0.ApisDashboardsV1AstWidgetsDataTableSpansQueryAggregation{*apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableSpansQueryGroupingToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableSpansQueryAggregationToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["metric_field"] = "unspecified"
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["aggregation_type"] = "unspecified"

// 		apisDashboardsV1CommonSpansAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationModel["metric_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationMetricAggregationModel}

// 		model := make(map[string]interface{})
// 		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		model["name"] = "count"
// 		model["is_visible"] = true
// 		model["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.MetricField = core.StringPtr("unspecified")
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.AggregationType = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonSpansAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationModel.MetricAggregation = apisDashboardsV1CommonSpansAggregationMetricAggregationModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsDataTableSpansQueryAggregation)
// 	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	model.Name = core.StringPtr("count")
// 	model.IsVisible = core.BoolPtr(true)
// 	model.Aggregation = apisDashboardsV1CommonSpansAggregationModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableSpansQueryAggregationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableMetricsQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterMetricsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterMetricsFilterModel["label"] = "service"
// 		apisDashboardsV1AstFilterMetricsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		model := make(map[string]interface{})
// 		model["promql_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonPromQlQueryModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterMetricsFilterModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonPromQlQuery)
// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterMetricsFilterModel := new(logsv0.ApisDashboardsV1AstFilterMetricsFilter)
// 	apisDashboardsV1AstFilterMetricsFilterModel.Label = core.StringPtr("service")
// 	apisDashboardsV1AstFilterMetricsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsDataTableMetricsQuery)
// 	model.PromqlQuery = apisDashboardsV1AstWidgetsCommonPromQlQueryModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterMetricsFilter{*apisDashboardsV1AstFilterMetricsFilterModel}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableMetricsQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableDataprimeQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonDataprimeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonDataprimeQueryModel["text"] = "source logs"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstFilterSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSourceModel["logs"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}

// 		model := make(map[string]interface{})
// 		model["dataprime_query"] = []map[string]interface{}{apisDashboardsV1CommonDataprimeQueryModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSourceModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonDataprimeQueryModel := new(logsv0.ApisDashboardsV1CommonDataprimeQuery)
// 	apisDashboardsV1CommonDataprimeQueryModel.Text = core.StringPtr("source logs")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstFilterSourceModel := new(logsv0.ApisDashboardsV1AstFilterSourceValueLogs)
// 	apisDashboardsV1AstFilterSourceModel.Logs = apisDashboardsV1AstFilterLogsFilterModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsDataTableDataprimeQuery)
// 	model.DataprimeQuery = apisDashboardsV1CommonDataprimeQueryModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterSourceIntf{apisDashboardsV1AstFilterSourceModel}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableDataprimeQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableQueryValueLogsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["name"] = "count"
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}

// 		apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel["aggregations"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel}
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel["group_bys"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsDataTableLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryModel["grouping"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel}

// 		model := make(map[string]interface{})
// 		model["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableLogsQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryAggregation)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.Name = core.StringPtr("count")
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.Aggregation = apisDashboardsV1CommonLogsAggregationModel

// 	apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryGrouping)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel.Aggregations = []logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryAggregation{*apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel}
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel.GroupBys = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	apisDashboardsV1AstWidgetsDataTableLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQuery)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryModel.Grouping = apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsDataTableQueryValueLogs)
// 	model.Logs = apisDashboardsV1AstWidgetsDataTableLogsQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableQueryValueLogsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableQueryValueSpansToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonSpanFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpanFieldModel["metadata_field"] = "unspecified"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterSpansFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSpansFilterModel["field"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstFilterSpansFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["metric_field"] = "unspecified"
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["aggregation_type"] = "unspecified"

// 		apisDashboardsV1CommonSpansAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationModel["metric_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationMetricAggregationModel}

// 		apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel["name"] = "count"
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationModel}

// 		apisDashboardsV1AstWidgetsDataTableSpansQueryGroupingModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryGroupingModel["group_by"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryGroupingModel["aggregations"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel}

// 		apisDashboardsV1AstWidgetsDataTableSpansQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSpansFilterModel}
// 		apisDashboardsV1AstWidgetsDataTableSpansQueryModel["grouping"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableSpansQueryGroupingModel}

// 		model := make(map[string]interface{})
// 		model["spans"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableSpansQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonSpanFieldModel := new(logsv0.ApisDashboardsV1CommonSpanFieldValueMetadataField)
// 	apisDashboardsV1CommonSpanFieldModel.MetadataField = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterSpansFilterModel := new(logsv0.ApisDashboardsV1AstFilterSpansFilter)
// 	apisDashboardsV1AstFilterSpansFilterModel.Field = apisDashboardsV1CommonSpanFieldModel
// 	apisDashboardsV1AstFilterSpansFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.MetricField = core.StringPtr("unspecified")
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.AggregationType = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonSpansAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationModel.MetricAggregation = apisDashboardsV1CommonSpansAggregationMetricAggregationModel

// 	apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableSpansQueryAggregation)
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel.Name = core.StringPtr("count")
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel.Aggregation = apisDashboardsV1CommonSpansAggregationModel

// 	apisDashboardsV1AstWidgetsDataTableSpansQueryGroupingModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableSpansQueryGrouping)
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryGroupingModel.GroupBy = []logsv0.ApisDashboardsV1CommonSpanFieldIntf{apisDashboardsV1CommonSpanFieldModel}
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryGroupingModel.Aggregations = []logsv0.ApisDashboardsV1AstWidgetsDataTableSpansQueryAggregation{*apisDashboardsV1AstWidgetsDataTableSpansQueryAggregationModel}

// 	apisDashboardsV1AstWidgetsDataTableSpansQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableSpansQuery)
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterSpansFilter{*apisDashboardsV1AstFilterSpansFilterModel}
// 	apisDashboardsV1AstWidgetsDataTableSpansQueryModel.Grouping = apisDashboardsV1AstWidgetsDataTableSpansQueryGroupingModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsDataTableQueryValueSpans)
// 	model.Spans = apisDashboardsV1AstWidgetsDataTableSpansQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableQueryValueSpansToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableQueryValueMetricsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterMetricsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterMetricsFilterModel["label"] = "service"
// 		apisDashboardsV1AstFilterMetricsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		apisDashboardsV1AstWidgetsDataTableMetricsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableMetricsQueryModel["promql_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonPromQlQueryModel}
// 		apisDashboardsV1AstWidgetsDataTableMetricsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterMetricsFilterModel}

// 		model := make(map[string]interface{})
// 		model["metrics"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableMetricsQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonPromQlQuery)
// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterMetricsFilterModel := new(logsv0.ApisDashboardsV1AstFilterMetricsFilter)
// 	apisDashboardsV1AstFilterMetricsFilterModel.Label = core.StringPtr("service")
// 	apisDashboardsV1AstFilterMetricsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	apisDashboardsV1AstWidgetsDataTableMetricsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableMetricsQuery)
// 	apisDashboardsV1AstWidgetsDataTableMetricsQueryModel.PromqlQuery = apisDashboardsV1AstWidgetsCommonPromQlQueryModel
// 	apisDashboardsV1AstWidgetsDataTableMetricsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterMetricsFilter{*apisDashboardsV1AstFilterMetricsFilterModel}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsDataTableQueryValueMetrics)
// 	model.Metrics = apisDashboardsV1AstWidgetsDataTableMetricsQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableQueryValueMetricsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableQueryValueDataprimeToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonDataprimeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonDataprimeQueryModel["text"] = "source logs"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstFilterSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSourceModel["logs"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}

// 		apisDashboardsV1AstWidgetsDataTableDataprimeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableDataprimeQueryModel["dataprime_query"] = []map[string]interface{}{apisDashboardsV1CommonDataprimeQueryModel}
// 		apisDashboardsV1AstWidgetsDataTableDataprimeQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSourceModel}

// 		model := make(map[string]interface{})
// 		model["dataprime"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableDataprimeQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonDataprimeQueryModel := new(logsv0.ApisDashboardsV1CommonDataprimeQuery)
// 	apisDashboardsV1CommonDataprimeQueryModel.Text = core.StringPtr("source logs")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstFilterSourceModel := new(logsv0.ApisDashboardsV1AstFilterSourceValueLogs)
// 	apisDashboardsV1AstFilterSourceModel.Logs = apisDashboardsV1AstFilterLogsFilterModel

// 	apisDashboardsV1AstWidgetsDataTableDataprimeQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableDataprimeQuery)
// 	apisDashboardsV1AstWidgetsDataTableDataprimeQueryModel.DataprimeQuery = apisDashboardsV1CommonDataprimeQueryModel
// 	apisDashboardsV1AstWidgetsDataTableDataprimeQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterSourceIntf{apisDashboardsV1AstFilterSourceModel}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsDataTableQueryValueDataprime)
// 	model.Dataprime = apisDashboardsV1AstWidgetsDataTableDataprimeQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableQueryValueDataprimeToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableColumnToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["field"] = "coralogix.metadata.applicationName"
// 		model["width"] = int(100)

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsDataTableColumn)
// 	model.Field = core.StringPtr("coralogix.metadata.applicationName")
// 	model.Width = core.Int64Ptr(int64(100))

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsDataTableColumnToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonOrderingFieldToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["field"] = "testString"
// 		model["order_direction"] = "unspecified"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1CommonOrderingField)
// 	model.Field = core.StringPtr("testString")
// 	model.OrderDirection = core.StringPtr("unspecified")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonOrderingFieldToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterMetricsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterMetricsFilterModel["label"] = "service"
// 		apisDashboardsV1AstFilterMetricsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		apisDashboardsV1AstWidgetsGaugeMetricsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsGaugeMetricsQueryModel["promql_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonPromQlQueryModel}
// 		apisDashboardsV1AstWidgetsGaugeMetricsQueryModel["aggregation"] = "unspecified"
// 		apisDashboardsV1AstWidgetsGaugeMetricsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterMetricsFilterModel}

// 		apisDashboardsV1AstWidgetsGaugeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsGaugeQueryModel["metrics"] = []map[string]interface{}{apisDashboardsV1AstWidgetsGaugeMetricsQueryModel}

// 		apisDashboardsV1AstWidgetsGaugeThresholdModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsGaugeThresholdModel["from"] = float64(0.5)
// 		apisDashboardsV1AstWidgetsGaugeThresholdModel["color"] = "warning"

// 		model := make(map[string]interface{})
// 		model["query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsGaugeQueryModel}
// 		model["min"] = float64(0)
// 		model["max"] = float64(100)
// 		model["show_inner_arc"] = true
// 		model["show_outer_arc"] = true
// 		model["unit"] = "unspecified"
// 		model["thresholds"] = []map[string]interface{}{apisDashboardsV1AstWidgetsGaugeThresholdModel}
// 		model["data_mode_type"] = "high_unspecified"
// 		model["threshold_by"] = "unspecified"

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonPromQlQuery)
// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterMetricsFilterModel := new(logsv0.ApisDashboardsV1AstFilterMetricsFilter)
// 	apisDashboardsV1AstFilterMetricsFilterModel.Label = core.StringPtr("service")
// 	apisDashboardsV1AstFilterMetricsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	apisDashboardsV1AstWidgetsGaugeMetricsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsGaugeMetricsQuery)
// 	apisDashboardsV1AstWidgetsGaugeMetricsQueryModel.PromqlQuery = apisDashboardsV1AstWidgetsCommonPromQlQueryModel
// 	apisDashboardsV1AstWidgetsGaugeMetricsQueryModel.Aggregation = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsGaugeMetricsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterMetricsFilter{*apisDashboardsV1AstFilterMetricsFilterModel}

// 	apisDashboardsV1AstWidgetsGaugeQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsGaugeQueryValueMetrics)
// 	apisDashboardsV1AstWidgetsGaugeQueryModel.Metrics = apisDashboardsV1AstWidgetsGaugeMetricsQueryModel

// 	apisDashboardsV1AstWidgetsGaugeThresholdModel := new(logsv0.ApisDashboardsV1AstWidgetsGaugeThreshold)
// 	apisDashboardsV1AstWidgetsGaugeThresholdModel.From = core.Float64Ptr(float64(0.5))
// 	apisDashboardsV1AstWidgetsGaugeThresholdModel.Color = core.StringPtr("warning")

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsGauge)
// 	model.Query = apisDashboardsV1AstWidgetsGaugeQueryModel
// 	model.Min = core.Float64Ptr(float64(0))
// 	model.Max = core.Float64Ptr(float64(100))
// 	model.ShowInnerArc = core.BoolPtr(true)
// 	model.ShowOuterArc = core.BoolPtr(true)
// 	model.Unit = core.StringPtr("unspecified")
// 	model.Thresholds = []logsv0.ApisDashboardsV1AstWidgetsGaugeThreshold{*apisDashboardsV1AstWidgetsGaugeThresholdModel}
// 	model.DataModeType = core.StringPtr("high_unspecified")
// 	model.ThresholdBy = core.StringPtr("unspecified")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterMetricsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterMetricsFilterModel["label"] = "service"
// 		apisDashboardsV1AstFilterMetricsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		apisDashboardsV1AstWidgetsGaugeMetricsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsGaugeMetricsQueryModel["promql_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonPromQlQueryModel}
// 		apisDashboardsV1AstWidgetsGaugeMetricsQueryModel["aggregation"] = "unspecified"
// 		apisDashboardsV1AstWidgetsGaugeMetricsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterMetricsFilterModel}

// 		model := make(map[string]interface{})
// 		model["metrics"] = []map[string]interface{}{apisDashboardsV1AstWidgetsGaugeMetricsQueryModel}
// 		model["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsGaugeLogsQueryModel}
// 		model["spans"] = []map[string]interface{}{apisDashboardsV1AstWidgetsGaugeSpansQueryModel}
// 		model["dataprime"] = []map[string]interface{}{apisDashboardsV1AstWidgetsGaugeDataprimeQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonPromQlQuery)
// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterMetricsFilterModel := new(logsv0.ApisDashboardsV1AstFilterMetricsFilter)
// 	apisDashboardsV1AstFilterMetricsFilterModel.Label = core.StringPtr("service")
// 	apisDashboardsV1AstFilterMetricsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	apisDashboardsV1AstWidgetsGaugeMetricsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsGaugeMetricsQuery)
// 	apisDashboardsV1AstWidgetsGaugeMetricsQueryModel.PromqlQuery = apisDashboardsV1AstWidgetsCommonPromQlQueryModel
// 	apisDashboardsV1AstWidgetsGaugeMetricsQueryModel.Aggregation = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsGaugeMetricsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterMetricsFilter{*apisDashboardsV1AstFilterMetricsFilterModel}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsGaugeQuery)
// 	model.Metrics = apisDashboardsV1AstWidgetsGaugeMetricsQueryModel
// 	model.Logs = apisDashboardsV1AstWidgetsGaugeLogsQueryModel
// 	model.Spans = apisDashboardsV1AstWidgetsGaugeSpansQueryModel
// 	model.Dataprime = apisDashboardsV1AstWidgetsGaugeDataprimeQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeMetricsQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterMetricsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterMetricsFilterModel["label"] = "service"
// 		apisDashboardsV1AstFilterMetricsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		model := make(map[string]interface{})
// 		model["promql_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonPromQlQueryModel}
// 		model["aggregation"] = "unspecified"
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterMetricsFilterModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonPromQlQuery)
// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterMetricsFilterModel := new(logsv0.ApisDashboardsV1AstFilterMetricsFilter)
// 	apisDashboardsV1AstFilterMetricsFilterModel.Label = core.StringPtr("service")
// 	apisDashboardsV1AstFilterMetricsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsGaugeMetricsQuery)
// 	model.PromqlQuery = apisDashboardsV1AstWidgetsCommonPromQlQueryModel
// 	model.Aggregation = core.StringPtr("unspecified")
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterMetricsFilter{*apisDashboardsV1AstFilterMetricsFilterModel}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeMetricsQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeLogsQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		model["logs_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsGaugeLogsQuery)
// 	model.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	model.LogsAggregation = apisDashboardsV1CommonLogsAggregationModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeLogsQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeSpansQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["metric_field"] = "unspecified"
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["aggregation_type"] = "unspecified"

// 		apisDashboardsV1CommonSpansAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationModel["metric_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationMetricAggregationModel}

// 		apisDashboardsV1CommonSpanFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpanFieldModel["metadata_field"] = "unspecified"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterSpansFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSpansFilterModel["field"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstFilterSpansFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		model := make(map[string]interface{})
// 		model["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		model["spans_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSpansFilterModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.MetricField = core.StringPtr("unspecified")
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.AggregationType = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonSpansAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationModel.MetricAggregation = apisDashboardsV1CommonSpansAggregationMetricAggregationModel

// 	apisDashboardsV1CommonSpanFieldModel := new(logsv0.ApisDashboardsV1CommonSpanFieldValueMetadataField)
// 	apisDashboardsV1CommonSpanFieldModel.MetadataField = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterSpansFilterModel := new(logsv0.ApisDashboardsV1AstFilterSpansFilter)
// 	apisDashboardsV1AstFilterSpansFilterModel.Field = apisDashboardsV1CommonSpanFieldModel
// 	apisDashboardsV1AstFilterSpansFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsGaugeSpansQuery)
// 	model.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	model.SpansAggregation = apisDashboardsV1CommonSpansAggregationModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterSpansFilter{*apisDashboardsV1AstFilterSpansFilterModel}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeSpansQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeDataprimeQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonDataprimeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonDataprimeQueryModel["text"] = "source logs"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstFilterSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSourceModel["logs"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}

// 		model := make(map[string]interface{})
// 		model["dataprime_query"] = []map[string]interface{}{apisDashboardsV1CommonDataprimeQueryModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSourceModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonDataprimeQueryModel := new(logsv0.ApisDashboardsV1CommonDataprimeQuery)
// 	apisDashboardsV1CommonDataprimeQueryModel.Text = core.StringPtr("source logs")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstFilterSourceModel := new(logsv0.ApisDashboardsV1AstFilterSourceValueLogs)
// 	apisDashboardsV1AstFilterSourceModel.Logs = apisDashboardsV1AstFilterLogsFilterModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsGaugeDataprimeQuery)
// 	model.DataprimeQuery = apisDashboardsV1CommonDataprimeQueryModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterSourceIntf{apisDashboardsV1AstFilterSourceModel}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeDataprimeQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeQueryValueMetricsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterMetricsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterMetricsFilterModel["label"] = "service"
// 		apisDashboardsV1AstFilterMetricsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		apisDashboardsV1AstWidgetsGaugeMetricsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsGaugeMetricsQueryModel["promql_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonPromQlQueryModel}
// 		apisDashboardsV1AstWidgetsGaugeMetricsQueryModel["aggregation"] = "unspecified"
// 		apisDashboardsV1AstWidgetsGaugeMetricsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterMetricsFilterModel}

// 		model := make(map[string]interface{})
// 		model["metrics"] = []map[string]interface{}{apisDashboardsV1AstWidgetsGaugeMetricsQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonPromQlQuery)
// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterMetricsFilterModel := new(logsv0.ApisDashboardsV1AstFilterMetricsFilter)
// 	apisDashboardsV1AstFilterMetricsFilterModel.Label = core.StringPtr("service")
// 	apisDashboardsV1AstFilterMetricsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	apisDashboardsV1AstWidgetsGaugeMetricsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsGaugeMetricsQuery)
// 	apisDashboardsV1AstWidgetsGaugeMetricsQueryModel.PromqlQuery = apisDashboardsV1AstWidgetsCommonPromQlQueryModel
// 	apisDashboardsV1AstWidgetsGaugeMetricsQueryModel.Aggregation = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsGaugeMetricsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterMetricsFilter{*apisDashboardsV1AstFilterMetricsFilterModel}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsGaugeQueryValueMetrics)
// 	model.Metrics = apisDashboardsV1AstWidgetsGaugeMetricsQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeQueryValueMetricsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeQueryValueLogsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsGaugeLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsGaugeLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsGaugeLogsQueryModel["logs_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsGaugeLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}

// 		model := make(map[string]interface{})
// 		model["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsGaugeLogsQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsGaugeLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsGaugeLogsQuery)
// 	apisDashboardsV1AstWidgetsGaugeLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsGaugeLogsQueryModel.LogsAggregation = apisDashboardsV1CommonLogsAggregationModel
// 	apisDashboardsV1AstWidgetsGaugeLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsGaugeQueryValueLogs)
// 	model.Logs = apisDashboardsV1AstWidgetsGaugeLogsQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeQueryValueLogsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeQueryValueSpansToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["metric_field"] = "unspecified"
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["aggregation_type"] = "unspecified"

// 		apisDashboardsV1CommonSpansAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationModel["metric_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationMetricAggregationModel}

// 		apisDashboardsV1CommonSpanFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpanFieldModel["metadata_field"] = "unspecified"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterSpansFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSpansFilterModel["field"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstFilterSpansFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		apisDashboardsV1AstWidgetsGaugeSpansQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsGaugeSpansQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsGaugeSpansQueryModel["spans_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationModel}
// 		apisDashboardsV1AstWidgetsGaugeSpansQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSpansFilterModel}

// 		model := make(map[string]interface{})
// 		model["spans"] = []map[string]interface{}{apisDashboardsV1AstWidgetsGaugeSpansQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.MetricField = core.StringPtr("unspecified")
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.AggregationType = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonSpansAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationModel.MetricAggregation = apisDashboardsV1CommonSpansAggregationMetricAggregationModel

// 	apisDashboardsV1CommonSpanFieldModel := new(logsv0.ApisDashboardsV1CommonSpanFieldValueMetadataField)
// 	apisDashboardsV1CommonSpanFieldModel.MetadataField = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterSpansFilterModel := new(logsv0.ApisDashboardsV1AstFilterSpansFilter)
// 	apisDashboardsV1AstFilterSpansFilterModel.Field = apisDashboardsV1CommonSpanFieldModel
// 	apisDashboardsV1AstFilterSpansFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	apisDashboardsV1AstWidgetsGaugeSpansQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsGaugeSpansQuery)
// 	apisDashboardsV1AstWidgetsGaugeSpansQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsGaugeSpansQueryModel.SpansAggregation = apisDashboardsV1CommonSpansAggregationModel
// 	apisDashboardsV1AstWidgetsGaugeSpansQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterSpansFilter{*apisDashboardsV1AstFilterSpansFilterModel}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsGaugeQueryValueSpans)
// 	model.Spans = apisDashboardsV1AstWidgetsGaugeSpansQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeQueryValueSpansToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeQueryValueDataprimeToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonDataprimeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonDataprimeQueryModel["text"] = "source logs"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstFilterSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSourceModel["logs"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}

// 		apisDashboardsV1AstWidgetsGaugeDataprimeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsGaugeDataprimeQueryModel["dataprime_query"] = []map[string]interface{}{apisDashboardsV1CommonDataprimeQueryModel}
// 		apisDashboardsV1AstWidgetsGaugeDataprimeQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSourceModel}

// 		model := make(map[string]interface{})
// 		model["dataprime"] = []map[string]interface{}{apisDashboardsV1AstWidgetsGaugeDataprimeQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonDataprimeQueryModel := new(logsv0.ApisDashboardsV1CommonDataprimeQuery)
// 	apisDashboardsV1CommonDataprimeQueryModel.Text = core.StringPtr("source logs")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstFilterSourceModel := new(logsv0.ApisDashboardsV1AstFilterSourceValueLogs)
// 	apisDashboardsV1AstFilterSourceModel.Logs = apisDashboardsV1AstFilterLogsFilterModel

// 	apisDashboardsV1AstWidgetsGaugeDataprimeQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsGaugeDataprimeQuery)
// 	apisDashboardsV1AstWidgetsGaugeDataprimeQueryModel.DataprimeQuery = apisDashboardsV1CommonDataprimeQueryModel
// 	apisDashboardsV1AstWidgetsGaugeDataprimeQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterSourceIntf{apisDashboardsV1AstFilterSourceModel}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsGaugeQueryValueDataprime)
// 	model.Dataprime = apisDashboardsV1AstWidgetsGaugeDataprimeQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeQueryValueDataprimeToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeThresholdToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["from"] = float64(0.5)
// 		model["color"] = "warning"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsGaugeThreshold)
// 	model.From = core.Float64Ptr(float64(0.5))
// 	model.Color = core.StringPtr("warning")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsGaugeThresholdToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["group_names_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["stacked_group_name_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsPieChartQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsPieChartQueryModel["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsPieChartLogsQueryModel}

// 		apisDashboardsV1AstWidgetsPieChartStackDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsPieChartStackDefinitionModel["max_slices_per_stack"] = int(5)
// 		apisDashboardsV1AstWidgetsPieChartStackDefinitionModel["stack_name_template"] = "{{severity}}"

// 		apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel["label_source"] = "unspecified"
// 		apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel["show_name"] = true
// 		apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel["show_value"] = true
// 		apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel["show_percentage"] = true

// 		model := make(map[string]interface{})
// 		model["query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsPieChartQueryModel}
// 		model["max_slices_per_chart"] = int(5)
// 		model["min_slice_percentage"] = int(1)
// 		model["stack_definition"] = []map[string]interface{}{apisDashboardsV1AstWidgetsPieChartStackDefinitionModel}
// 		model["label_definition"] = []map[string]interface{}{apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel}
// 		model["show_legend"] = true
// 		model["group_name_template"] = "testString"
// 		model["unit"] = "unspecified"
// 		model["color_scheme"] = "classic"
// 		model["data_mode_type"] = "high_unspecified"

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsPieChartLogsQuery)
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.Aggregation = apisDashboardsV1CommonLogsAggregationModel
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.GroupNamesFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.StackedGroupNameField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsPieChartQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsPieChartQueryValueLogs)
// 	apisDashboardsV1AstWidgetsPieChartQueryModel.Logs = apisDashboardsV1AstWidgetsPieChartLogsQueryModel

// 	apisDashboardsV1AstWidgetsPieChartStackDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetsPieChartStackDefinition)
// 	apisDashboardsV1AstWidgetsPieChartStackDefinitionModel.MaxSlicesPerStack = core.Int64Ptr(int64(5))
// 	apisDashboardsV1AstWidgetsPieChartStackDefinitionModel.StackNameTemplate = core.StringPtr("{{severity}}")

// 	apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetsPieChartLabelDefinition)
// 	apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel.LabelSource = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel.ShowName = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel.ShowValue = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel.ShowPercentage = core.BoolPtr(true)

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsPieChart)
// 	model.Query = apisDashboardsV1AstWidgetsPieChartQueryModel
// 	model.MaxSlicesPerChart = core.Int64Ptr(int64(5))
// 	model.MinSlicePercentage = core.Int64Ptr(int64(1))
// 	model.StackDefinition = apisDashboardsV1AstWidgetsPieChartStackDefinitionModel
// 	model.LabelDefinition = apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel
// 	model.ShowLegend = core.BoolPtr(true)
// 	model.GroupNameTemplate = core.StringPtr("testString")
// 	model.Unit = core.StringPtr("unspecified")
// 	model.ColorScheme = core.StringPtr("classic")
// 	model.DataModeType = core.StringPtr("high_unspecified")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["group_names_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["stacked_group_name_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsPieChartLogsQueryModel}
// 		model["spans"] = []map[string]interface{}{apisDashboardsV1AstWidgetsPieChartSpansQueryModel}
// 		model["metrics"] = []map[string]interface{}{apisDashboardsV1AstWidgetsPieChartMetricsQueryModel}
// 		model["dataprime"] = []map[string]interface{}{apisDashboardsV1AstWidgetsPieChartDataprimeQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsPieChartLogsQuery)
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.Aggregation = apisDashboardsV1CommonLogsAggregationModel
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.GroupNamesFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.StackedGroupNameField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsPieChartQuery)
// 	model.Logs = apisDashboardsV1AstWidgetsPieChartLogsQueryModel
// 	model.Spans = apisDashboardsV1AstWidgetsPieChartSpansQueryModel
// 	model.Metrics = apisDashboardsV1AstWidgetsPieChartMetricsQueryModel
// 	model.Dataprime = apisDashboardsV1AstWidgetsPieChartDataprimeQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartLogsQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		model["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		model["group_names_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		model["stacked_group_name_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsPieChartLogsQuery)
// 	model.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	model.Aggregation = apisDashboardsV1CommonLogsAggregationModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	model.GroupNamesFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}
// 	model.StackedGroupNameField = apisDashboardsV1CommonObservationFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartLogsQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartSpansQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["metric_field"] = "unspecified"
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["aggregation_type"] = "unspecified"

// 		apisDashboardsV1CommonSpansAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationModel["metric_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationMetricAggregationModel}

// 		apisDashboardsV1CommonSpanFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpanFieldModel["metadata_field"] = "unspecified"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterSpansFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSpansFilterModel["field"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstFilterSpansFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		model := make(map[string]interface{})
// 		model["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		model["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSpansFilterModel}
// 		model["group_names"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		model["stacked_group_name"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.MetricField = core.StringPtr("unspecified")
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.AggregationType = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonSpansAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationModel.MetricAggregation = apisDashboardsV1CommonSpansAggregationMetricAggregationModel

// 	apisDashboardsV1CommonSpanFieldModel := new(logsv0.ApisDashboardsV1CommonSpanFieldValueMetadataField)
// 	apisDashboardsV1CommonSpanFieldModel.MetadataField = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterSpansFilterModel := new(logsv0.ApisDashboardsV1AstFilterSpansFilter)
// 	apisDashboardsV1AstFilterSpansFilterModel.Field = apisDashboardsV1CommonSpanFieldModel
// 	apisDashboardsV1AstFilterSpansFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsPieChartSpansQuery)
// 	model.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	model.Aggregation = apisDashboardsV1CommonSpansAggregationModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterSpansFilter{*apisDashboardsV1AstFilterSpansFilterModel}
// 	model.GroupNames = []logsv0.ApisDashboardsV1CommonSpanFieldIntf{apisDashboardsV1CommonSpanFieldModel}
// 	model.StackedGroupName = apisDashboardsV1CommonSpanFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartSpansQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartMetricsQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterMetricsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterMetricsFilterModel["label"] = "service"
// 		apisDashboardsV1AstFilterMetricsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		model := make(map[string]interface{})
// 		model["promql_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonPromQlQueryModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterMetricsFilterModel}
// 		model["group_names"] = []string{"testString"}
// 		model["stacked_group_name"] = "pod"

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonPromQlQuery)
// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterMetricsFilterModel := new(logsv0.ApisDashboardsV1AstFilterMetricsFilter)
// 	apisDashboardsV1AstFilterMetricsFilterModel.Label = core.StringPtr("service")
// 	apisDashboardsV1AstFilterMetricsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsPieChartMetricsQuery)
// 	model.PromqlQuery = apisDashboardsV1AstWidgetsCommonPromQlQueryModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterMetricsFilter{*apisDashboardsV1AstFilterMetricsFilterModel}
// 	model.GroupNames = []string{"testString"}
// 	model.StackedGroupName = core.StringPtr("pod")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartMetricsQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartDataprimeQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonDataprimeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonDataprimeQueryModel["text"] = "source logs"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstFilterSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSourceModel["logs"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}

// 		model := make(map[string]interface{})
// 		model["dataprime_query"] = []map[string]interface{}{apisDashboardsV1CommonDataprimeQueryModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSourceModel}
// 		model["group_names"] = []string{"testString"}
// 		model["stacked_group_name"] = "pod"

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonDataprimeQueryModel := new(logsv0.ApisDashboardsV1CommonDataprimeQuery)
// 	apisDashboardsV1CommonDataprimeQueryModel.Text = core.StringPtr("source logs")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstFilterSourceModel := new(logsv0.ApisDashboardsV1AstFilterSourceValueLogs)
// 	apisDashboardsV1AstFilterSourceModel.Logs = apisDashboardsV1AstFilterLogsFilterModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsPieChartDataprimeQuery)
// 	model.DataprimeQuery = apisDashboardsV1CommonDataprimeQueryModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterSourceIntf{apisDashboardsV1AstFilterSourceModel}
// 	model.GroupNames = []string{"testString"}
// 	model.StackedGroupName = core.StringPtr("pod")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartDataprimeQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartQueryValueLogsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["group_names_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["stacked_group_name_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsPieChartLogsQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsPieChartLogsQuery)
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.Aggregation = apisDashboardsV1CommonLogsAggregationModel
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.GroupNamesFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.StackedGroupNameField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsPieChartQueryValueLogs)
// 	model.Logs = apisDashboardsV1AstWidgetsPieChartLogsQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartQueryValueLogsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartQueryValueSpansToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["metric_field"] = "unspecified"
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["aggregation_type"] = "unspecified"

// 		apisDashboardsV1CommonSpansAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationModel["metric_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationMetricAggregationModel}

// 		apisDashboardsV1CommonSpanFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpanFieldModel["metadata_field"] = "unspecified"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterSpansFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSpansFilterModel["field"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstFilterSpansFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		apisDashboardsV1AstWidgetsPieChartSpansQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsPieChartSpansQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsPieChartSpansQueryModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationModel}
// 		apisDashboardsV1AstWidgetsPieChartSpansQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSpansFilterModel}
// 		apisDashboardsV1AstWidgetsPieChartSpansQueryModel["group_names"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstWidgetsPieChartSpansQueryModel["stacked_group_name"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}

// 		model := make(map[string]interface{})
// 		model["spans"] = []map[string]interface{}{apisDashboardsV1AstWidgetsPieChartSpansQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.MetricField = core.StringPtr("unspecified")
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.AggregationType = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonSpansAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationModel.MetricAggregation = apisDashboardsV1CommonSpansAggregationMetricAggregationModel

// 	apisDashboardsV1CommonSpanFieldModel := new(logsv0.ApisDashboardsV1CommonSpanFieldValueMetadataField)
// 	apisDashboardsV1CommonSpanFieldModel.MetadataField = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterSpansFilterModel := new(logsv0.ApisDashboardsV1AstFilterSpansFilter)
// 	apisDashboardsV1AstFilterSpansFilterModel.Field = apisDashboardsV1CommonSpanFieldModel
// 	apisDashboardsV1AstFilterSpansFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	apisDashboardsV1AstWidgetsPieChartSpansQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsPieChartSpansQuery)
// 	apisDashboardsV1AstWidgetsPieChartSpansQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsPieChartSpansQueryModel.Aggregation = apisDashboardsV1CommonSpansAggregationModel
// 	apisDashboardsV1AstWidgetsPieChartSpansQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterSpansFilter{*apisDashboardsV1AstFilterSpansFilterModel}
// 	apisDashboardsV1AstWidgetsPieChartSpansQueryModel.GroupNames = []logsv0.ApisDashboardsV1CommonSpanFieldIntf{apisDashboardsV1CommonSpanFieldModel}
// 	apisDashboardsV1AstWidgetsPieChartSpansQueryModel.StackedGroupName = apisDashboardsV1CommonSpanFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsPieChartQueryValueSpans)
// 	model.Spans = apisDashboardsV1AstWidgetsPieChartSpansQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartQueryValueSpansToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartQueryValueMetricsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterMetricsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterMetricsFilterModel["label"] = "service"
// 		apisDashboardsV1AstFilterMetricsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		apisDashboardsV1AstWidgetsPieChartMetricsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsPieChartMetricsQueryModel["promql_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonPromQlQueryModel}
// 		apisDashboardsV1AstWidgetsPieChartMetricsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterMetricsFilterModel}
// 		apisDashboardsV1AstWidgetsPieChartMetricsQueryModel["group_names"] = []string{"testString"}
// 		apisDashboardsV1AstWidgetsPieChartMetricsQueryModel["stacked_group_name"] = "pod"

// 		model := make(map[string]interface{})
// 		model["metrics"] = []map[string]interface{}{apisDashboardsV1AstWidgetsPieChartMetricsQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonPromQlQuery)
// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterMetricsFilterModel := new(logsv0.ApisDashboardsV1AstFilterMetricsFilter)
// 	apisDashboardsV1AstFilterMetricsFilterModel.Label = core.StringPtr("service")
// 	apisDashboardsV1AstFilterMetricsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	apisDashboardsV1AstWidgetsPieChartMetricsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsPieChartMetricsQuery)
// 	apisDashboardsV1AstWidgetsPieChartMetricsQueryModel.PromqlQuery = apisDashboardsV1AstWidgetsCommonPromQlQueryModel
// 	apisDashboardsV1AstWidgetsPieChartMetricsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterMetricsFilter{*apisDashboardsV1AstFilterMetricsFilterModel}
// 	apisDashboardsV1AstWidgetsPieChartMetricsQueryModel.GroupNames = []string{"testString"}
// 	apisDashboardsV1AstWidgetsPieChartMetricsQueryModel.StackedGroupName = core.StringPtr("pod")

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsPieChartQueryValueMetrics)
// 	model.Metrics = apisDashboardsV1AstWidgetsPieChartMetricsQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartQueryValueMetricsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartQueryValueDataprimeToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonDataprimeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonDataprimeQueryModel["text"] = "source logs"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstFilterSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSourceModel["logs"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}

// 		apisDashboardsV1AstWidgetsPieChartDataprimeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsPieChartDataprimeQueryModel["dataprime_query"] = []map[string]interface{}{apisDashboardsV1CommonDataprimeQueryModel}
// 		apisDashboardsV1AstWidgetsPieChartDataprimeQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSourceModel}
// 		apisDashboardsV1AstWidgetsPieChartDataprimeQueryModel["group_names"] = []string{"testString"}
// 		apisDashboardsV1AstWidgetsPieChartDataprimeQueryModel["stacked_group_name"] = "pod"

// 		model := make(map[string]interface{})
// 		model["dataprime"] = []map[string]interface{}{apisDashboardsV1AstWidgetsPieChartDataprimeQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonDataprimeQueryModel := new(logsv0.ApisDashboardsV1CommonDataprimeQuery)
// 	apisDashboardsV1CommonDataprimeQueryModel.Text = core.StringPtr("source logs")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstFilterSourceModel := new(logsv0.ApisDashboardsV1AstFilterSourceValueLogs)
// 	apisDashboardsV1AstFilterSourceModel.Logs = apisDashboardsV1AstFilterLogsFilterModel

// 	apisDashboardsV1AstWidgetsPieChartDataprimeQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsPieChartDataprimeQuery)
// 	apisDashboardsV1AstWidgetsPieChartDataprimeQueryModel.DataprimeQuery = apisDashboardsV1CommonDataprimeQueryModel
// 	apisDashboardsV1AstWidgetsPieChartDataprimeQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterSourceIntf{apisDashboardsV1AstFilterSourceModel}
// 	apisDashboardsV1AstWidgetsPieChartDataprimeQueryModel.GroupNames = []string{"testString"}
// 	apisDashboardsV1AstWidgetsPieChartDataprimeQueryModel.StackedGroupName = core.StringPtr("pod")

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsPieChartQueryValueDataprime)
// 	model.Dataprime = apisDashboardsV1AstWidgetsPieChartDataprimeQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartQueryValueDataprimeToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartStackDefinitionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["max_slices_per_stack"] = int(5)
// 		model["stack_name_template"] = "{{severity}}"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsPieChartStackDefinition)
// 	model.MaxSlicesPerStack = core.Int64Ptr(int64(5))
// 	model.StackNameTemplate = core.StringPtr("{{severity}}")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartStackDefinitionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartLabelDefinitionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["label_source"] = "unspecified"
// 		model["is_visible"] = true
// 		model["show_name"] = true
// 		model["show_value"] = true
// 		model["show_percentage"] = true

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsPieChartLabelDefinition)
// 	model.LabelSource = core.StringPtr("unspecified")
// 	model.IsVisible = core.BoolPtr(true)
// 	model.ShowName = core.BoolPtr(true)
// 	model.ShowValue = core.BoolPtr(true)
// 	model.ShowPercentage = core.BoolPtr(true)

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsPieChartLabelDefinitionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["group_names_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["stacked_group_name_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsBarChartQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsBarChartQueryModel["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartLogsQueryModel}

// 		apisDashboardsV1AstWidgetsBarChartStackDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsBarChartStackDefinitionModel["max_slices_per_bar"] = int(5)
// 		apisDashboardsV1AstWidgetsBarChartStackDefinitionModel["stack_name_template"] = "{{severity}}"

// 		apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel := make(map[string]interface{})

// 		apisDashboardsV1AstWidgetsCommonColorsByModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonColorsByModel["stack"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel}

// 		apisDashboardsV1AstWidgetsBarChartXAxisXAxisByValueModel := make(map[string]interface{})

// 		apisDashboardsV1AstWidgetsBarChartXAxisModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsBarChartXAxisModel["value"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartXAxisXAxisByValueModel}

// 		model := make(map[string]interface{})
// 		model["query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartQueryModel}
// 		model["max_bars_per_chart"] = int(10)
// 		model["group_name_template"] = "{{severity}}"
// 		model["stack_definition"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartStackDefinitionModel}
// 		model["scale_type"] = "unspecified"
// 		model["colors_by"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonColorsByModel}
// 		model["x_axis"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartXAxisModel}
// 		model["unit"] = "unspecified"
// 		model["sort_by"] = "unspecified"
// 		model["color_scheme"] = "classic"
// 		model["data_mode_type"] = "high_unspecified"

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChartLogsQuery)
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.Aggregation = apisDashboardsV1CommonLogsAggregationModel
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.GroupNamesFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.StackedGroupNameField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsBarChartQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChartQueryValueLogs)
// 	apisDashboardsV1AstWidgetsBarChartQueryModel.Logs = apisDashboardsV1AstWidgetsBarChartLogsQueryModel

// 	apisDashboardsV1AstWidgetsBarChartStackDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChartStackDefinition)
// 	apisDashboardsV1AstWidgetsBarChartStackDefinitionModel.MaxSlicesPerBar = core.Int64Ptr(int64(5))
// 	apisDashboardsV1AstWidgetsBarChartStackDefinitionModel.StackNameTemplate = core.StringPtr("{{severity}}")

// 	apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsByColorsByStack)

// 	apisDashboardsV1AstWidgetsCommonColorsByModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsByValueStack)
// 	apisDashboardsV1AstWidgetsCommonColorsByModel.Stack = apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel

// 	apisDashboardsV1AstWidgetsBarChartXAxisXAxisByValueModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChartXAxisXAxisByValue)

// 	apisDashboardsV1AstWidgetsBarChartXAxisModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChartXAxisTypeValue)
// 	apisDashboardsV1AstWidgetsBarChartXAxisModel.Value = apisDashboardsV1AstWidgetsBarChartXAxisXAxisByValueModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsBarChart)
// 	model.Query = apisDashboardsV1AstWidgetsBarChartQueryModel
// 	model.MaxBarsPerChart = core.Int64Ptr(int64(10))
// 	model.GroupNameTemplate = core.StringPtr("{{severity}}")
// 	model.StackDefinition = apisDashboardsV1AstWidgetsBarChartStackDefinitionModel
// 	model.ScaleType = core.StringPtr("unspecified")
// 	model.ColorsBy = apisDashboardsV1AstWidgetsCommonColorsByModel
// 	model.XAxis = apisDashboardsV1AstWidgetsBarChartXAxisModel
// 	model.Unit = core.StringPtr("unspecified")
// 	model.SortBy = core.StringPtr("unspecified")
// 	model.ColorScheme = core.StringPtr("classic")
// 	model.DataModeType = core.StringPtr("high_unspecified")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["group_names_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["stacked_group_name_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartLogsQueryModel}
// 		model["spans"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartSpansQueryModel}
// 		model["metrics"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartMetricsQueryModel}
// 		model["dataprime"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartDataprimeQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChartLogsQuery)
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.Aggregation = apisDashboardsV1CommonLogsAggregationModel
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.GroupNamesFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.StackedGroupNameField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsBarChartQuery)
// 	model.Logs = apisDashboardsV1AstWidgetsBarChartLogsQueryModel
// 	model.Spans = apisDashboardsV1AstWidgetsBarChartSpansQueryModel
// 	model.Metrics = apisDashboardsV1AstWidgetsBarChartMetricsQueryModel
// 	model.Dataprime = apisDashboardsV1AstWidgetsBarChartDataprimeQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartLogsQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		model["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		model["group_names_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		model["stacked_group_name_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsBarChartLogsQuery)
// 	model.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	model.Aggregation = apisDashboardsV1CommonLogsAggregationModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	model.GroupNamesFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}
// 	model.StackedGroupNameField = apisDashboardsV1CommonObservationFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartLogsQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartSpansQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["metric_field"] = "unspecified"
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["aggregation_type"] = "unspecified"

// 		apisDashboardsV1CommonSpansAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationModel["metric_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationMetricAggregationModel}

// 		apisDashboardsV1CommonSpanFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpanFieldModel["metadata_field"] = "unspecified"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterSpansFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSpansFilterModel["field"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstFilterSpansFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		model := make(map[string]interface{})
// 		model["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		model["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSpansFilterModel}
// 		model["group_names"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		model["stacked_group_name"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.MetricField = core.StringPtr("unspecified")
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.AggregationType = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonSpansAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationModel.MetricAggregation = apisDashboardsV1CommonSpansAggregationMetricAggregationModel

// 	apisDashboardsV1CommonSpanFieldModel := new(logsv0.ApisDashboardsV1CommonSpanFieldValueMetadataField)
// 	apisDashboardsV1CommonSpanFieldModel.MetadataField = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterSpansFilterModel := new(logsv0.ApisDashboardsV1AstFilterSpansFilter)
// 	apisDashboardsV1AstFilterSpansFilterModel.Field = apisDashboardsV1CommonSpanFieldModel
// 	apisDashboardsV1AstFilterSpansFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsBarChartSpansQuery)
// 	model.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	model.Aggregation = apisDashboardsV1CommonSpansAggregationModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterSpansFilter{*apisDashboardsV1AstFilterSpansFilterModel}
// 	model.GroupNames = []logsv0.ApisDashboardsV1CommonSpanFieldIntf{apisDashboardsV1CommonSpanFieldModel}
// 	model.StackedGroupName = apisDashboardsV1CommonSpanFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartSpansQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartMetricsQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterMetricsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterMetricsFilterModel["label"] = "service"
// 		apisDashboardsV1AstFilterMetricsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		model := make(map[string]interface{})
// 		model["promql_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonPromQlQueryModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterMetricsFilterModel}
// 		model["group_names"] = []string{"testString"}
// 		model["stacked_group_name"] = "pod"

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonPromQlQuery)
// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterMetricsFilterModel := new(logsv0.ApisDashboardsV1AstFilterMetricsFilter)
// 	apisDashboardsV1AstFilterMetricsFilterModel.Label = core.StringPtr("service")
// 	apisDashboardsV1AstFilterMetricsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsBarChartMetricsQuery)
// 	model.PromqlQuery = apisDashboardsV1AstWidgetsCommonPromQlQueryModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterMetricsFilter{*apisDashboardsV1AstFilterMetricsFilterModel}
// 	model.GroupNames = []string{"testString"}
// 	model.StackedGroupName = core.StringPtr("pod")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartMetricsQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartDataprimeQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonDataprimeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonDataprimeQueryModel["text"] = "source logs"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstFilterSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSourceModel["logs"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}

// 		model := make(map[string]interface{})
// 		model["dataprime_query"] = []map[string]interface{}{apisDashboardsV1CommonDataprimeQueryModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSourceModel}
// 		model["group_names"] = []string{"testString"}
// 		model["stacked_group_name"] = "severity"

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonDataprimeQueryModel := new(logsv0.ApisDashboardsV1CommonDataprimeQuery)
// 	apisDashboardsV1CommonDataprimeQueryModel.Text = core.StringPtr("source logs")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstFilterSourceModel := new(logsv0.ApisDashboardsV1AstFilterSourceValueLogs)
// 	apisDashboardsV1AstFilterSourceModel.Logs = apisDashboardsV1AstFilterLogsFilterModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsBarChartDataprimeQuery)
// 	model.DataprimeQuery = apisDashboardsV1CommonDataprimeQueryModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterSourceIntf{apisDashboardsV1AstFilterSourceModel}
// 	model.GroupNames = []string{"testString"}
// 	model.StackedGroupName = core.StringPtr("severity")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartDataprimeQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartQueryValueLogsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["group_names_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["stacked_group_name_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartLogsQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChartLogsQuery)
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.Aggregation = apisDashboardsV1CommonLogsAggregationModel
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.GroupNamesFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.StackedGroupNameField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsBarChartQueryValueLogs)
// 	model.Logs = apisDashboardsV1AstWidgetsBarChartLogsQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartQueryValueLogsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartQueryValueSpansToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["metric_field"] = "unspecified"
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["aggregation_type"] = "unspecified"

// 		apisDashboardsV1CommonSpansAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationModel["metric_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationMetricAggregationModel}

// 		apisDashboardsV1CommonSpanFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpanFieldModel["metadata_field"] = "unspecified"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterSpansFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSpansFilterModel["field"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstFilterSpansFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		apisDashboardsV1AstWidgetsBarChartSpansQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsBarChartSpansQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsBarChartSpansQueryModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationModel}
// 		apisDashboardsV1AstWidgetsBarChartSpansQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSpansFilterModel}
// 		apisDashboardsV1AstWidgetsBarChartSpansQueryModel["group_names"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstWidgetsBarChartSpansQueryModel["stacked_group_name"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}

// 		model := make(map[string]interface{})
// 		model["spans"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartSpansQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.MetricField = core.StringPtr("unspecified")
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.AggregationType = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonSpansAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationModel.MetricAggregation = apisDashboardsV1CommonSpansAggregationMetricAggregationModel

// 	apisDashboardsV1CommonSpanFieldModel := new(logsv0.ApisDashboardsV1CommonSpanFieldValueMetadataField)
// 	apisDashboardsV1CommonSpanFieldModel.MetadataField = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterSpansFilterModel := new(logsv0.ApisDashboardsV1AstFilterSpansFilter)
// 	apisDashboardsV1AstFilterSpansFilterModel.Field = apisDashboardsV1CommonSpanFieldModel
// 	apisDashboardsV1AstFilterSpansFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	apisDashboardsV1AstWidgetsBarChartSpansQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChartSpansQuery)
// 	apisDashboardsV1AstWidgetsBarChartSpansQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsBarChartSpansQueryModel.Aggregation = apisDashboardsV1CommonSpansAggregationModel
// 	apisDashboardsV1AstWidgetsBarChartSpansQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterSpansFilter{*apisDashboardsV1AstFilterSpansFilterModel}
// 	apisDashboardsV1AstWidgetsBarChartSpansQueryModel.GroupNames = []logsv0.ApisDashboardsV1CommonSpanFieldIntf{apisDashboardsV1CommonSpanFieldModel}
// 	apisDashboardsV1AstWidgetsBarChartSpansQueryModel.StackedGroupName = apisDashboardsV1CommonSpanFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsBarChartQueryValueSpans)
// 	model.Spans = apisDashboardsV1AstWidgetsBarChartSpansQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartQueryValueSpansToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartQueryValueMetricsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterMetricsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterMetricsFilterModel["label"] = "service"
// 		apisDashboardsV1AstFilterMetricsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		apisDashboardsV1AstWidgetsBarChartMetricsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsBarChartMetricsQueryModel["promql_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonPromQlQueryModel}
// 		apisDashboardsV1AstWidgetsBarChartMetricsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterMetricsFilterModel}
// 		apisDashboardsV1AstWidgetsBarChartMetricsQueryModel["group_names"] = []string{"testString"}
// 		apisDashboardsV1AstWidgetsBarChartMetricsQueryModel["stacked_group_name"] = "pod"

// 		model := make(map[string]interface{})
// 		model["metrics"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartMetricsQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonPromQlQuery)
// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterMetricsFilterModel := new(logsv0.ApisDashboardsV1AstFilterMetricsFilter)
// 	apisDashboardsV1AstFilterMetricsFilterModel.Label = core.StringPtr("service")
// 	apisDashboardsV1AstFilterMetricsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	apisDashboardsV1AstWidgetsBarChartMetricsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChartMetricsQuery)
// 	apisDashboardsV1AstWidgetsBarChartMetricsQueryModel.PromqlQuery = apisDashboardsV1AstWidgetsCommonPromQlQueryModel
// 	apisDashboardsV1AstWidgetsBarChartMetricsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterMetricsFilter{*apisDashboardsV1AstFilterMetricsFilterModel}
// 	apisDashboardsV1AstWidgetsBarChartMetricsQueryModel.GroupNames = []string{"testString"}
// 	apisDashboardsV1AstWidgetsBarChartMetricsQueryModel.StackedGroupName = core.StringPtr("pod")

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsBarChartQueryValueMetrics)
// 	model.Metrics = apisDashboardsV1AstWidgetsBarChartMetricsQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartQueryValueMetricsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartQueryValueDataprimeToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonDataprimeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonDataprimeQueryModel["text"] = "source logs"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstFilterSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSourceModel["logs"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}

// 		apisDashboardsV1AstWidgetsBarChartDataprimeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsBarChartDataprimeQueryModel["dataprime_query"] = []map[string]interface{}{apisDashboardsV1CommonDataprimeQueryModel}
// 		apisDashboardsV1AstWidgetsBarChartDataprimeQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSourceModel}
// 		apisDashboardsV1AstWidgetsBarChartDataprimeQueryModel["group_names"] = []string{"testString"}
// 		apisDashboardsV1AstWidgetsBarChartDataprimeQueryModel["stacked_group_name"] = "severity"

// 		model := make(map[string]interface{})
// 		model["dataprime"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartDataprimeQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonDataprimeQueryModel := new(logsv0.ApisDashboardsV1CommonDataprimeQuery)
// 	apisDashboardsV1CommonDataprimeQueryModel.Text = core.StringPtr("source logs")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstFilterSourceModel := new(logsv0.ApisDashboardsV1AstFilterSourceValueLogs)
// 	apisDashboardsV1AstFilterSourceModel.Logs = apisDashboardsV1AstFilterLogsFilterModel

// 	apisDashboardsV1AstWidgetsBarChartDataprimeQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChartDataprimeQuery)
// 	apisDashboardsV1AstWidgetsBarChartDataprimeQueryModel.DataprimeQuery = apisDashboardsV1CommonDataprimeQueryModel
// 	apisDashboardsV1AstWidgetsBarChartDataprimeQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterSourceIntf{apisDashboardsV1AstFilterSourceModel}
// 	apisDashboardsV1AstWidgetsBarChartDataprimeQueryModel.GroupNames = []string{"testString"}
// 	apisDashboardsV1AstWidgetsBarChartDataprimeQueryModel.StackedGroupName = core.StringPtr("severity")

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsBarChartQueryValueDataprime)
// 	model.Dataprime = apisDashboardsV1AstWidgetsBarChartDataprimeQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartQueryValueDataprimeToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartStackDefinitionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["max_slices_per_bar"] = int(5)
// 		model["stack_name_template"] = "{{severity}}"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsBarChartStackDefinition)
// 	model.MaxSlicesPerBar = core.Int64Ptr(int64(5))
// 	model.StackNameTemplate = core.StringPtr("{{severity}}")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartStackDefinitionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonColorsByToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel := make(map[string]interface{})

// 		model := make(map[string]interface{})
// 		model["stack"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel}
// 		model["group_by"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonColorsByColorsByGroupByModel}
// 		model["aggregation"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonColorsByColorsByAggregationModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsByColorsByStack)

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsBy)
// 	model.Stack = apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel
// 	model.GroupBy = apisDashboardsV1AstWidgetsCommonColorsByColorsByGroupByModel
// 	model.Aggregation = apisDashboardsV1AstWidgetsCommonColorsByColorsByAggregationModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonColorsByToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonColorsByColorsByStackToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsByColorsByStack)

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonColorsByColorsByStackToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonColorsByColorsByGroupByToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsByColorsByGroupBy)

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonColorsByColorsByGroupByToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonColorsByColorsByAggregationToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsByColorsByAggregation)

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonColorsByColorsByAggregationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonColorsByValueStackToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel := make(map[string]interface{})

// 		model := make(map[string]interface{})
// 		model["stack"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsByColorsByStack)

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsByValueStack)
// 	model.Stack = apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonColorsByValueStackToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonColorsByValueGroupByToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonColorsByColorsByGroupByModel := make(map[string]interface{})

// 		model := make(map[string]interface{})
// 		model["group_by"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonColorsByColorsByGroupByModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonColorsByColorsByGroupByModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsByColorsByGroupBy)

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsByValueGroupBy)
// 	model.GroupBy = apisDashboardsV1AstWidgetsCommonColorsByColorsByGroupByModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonColorsByValueGroupByToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonColorsByValueAggregationToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonColorsByColorsByAggregationModel := make(map[string]interface{})

// 		model := make(map[string]interface{})
// 		model["aggregation"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonColorsByColorsByAggregationModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonColorsByColorsByAggregationModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsByColorsByAggregation)

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsByValueAggregation)
// 	model.Aggregation = apisDashboardsV1AstWidgetsCommonColorsByColorsByAggregationModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsCommonColorsByValueAggregationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartXAxisToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsBarChartXAxisXAxisByValueModel := make(map[string]interface{})

// 		model := make(map[string]interface{})
// 		model["value"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartXAxisXAxisByValueModel}
// 		model["time"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartXAxisXAxisByTimeModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsBarChartXAxisXAxisByValueModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChartXAxisXAxisByValue)

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsBarChartXAxis)
// 	model.Value = apisDashboardsV1AstWidgetsBarChartXAxisXAxisByValueModel
// 	model.Time = apisDashboardsV1AstWidgetsBarChartXAxisXAxisByTimeModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartXAxisToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartXAxisXAxisByValueToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsBarChartXAxisXAxisByValue)

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartXAxisXAxisByValueToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartXAxisXAxisByTimeToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["interval"] = "60s"
// 		model["buckets_presented"] = int(10)

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsBarChartXAxisXAxisByTime)
// 	model.Interval = core.StringPtr("60s")
// 	model.BucketsPresented = core.Int64Ptr(int64(10))

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartXAxisXAxisByTimeToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartXAxisTypeValueToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsBarChartXAxisXAxisByValueModel := make(map[string]interface{})

// 		model := make(map[string]interface{})
// 		model["value"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartXAxisXAxisByValueModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsBarChartXAxisXAxisByValueModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChartXAxisXAxisByValue)

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsBarChartXAxisTypeValue)
// 	model.Value = apisDashboardsV1AstWidgetsBarChartXAxisXAxisByValueModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartXAxisTypeValueToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartXAxisTypeTimeToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsBarChartXAxisXAxisByTimeModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsBarChartXAxisXAxisByTimeModel["interval"] = "60s"
// 		apisDashboardsV1AstWidgetsBarChartXAxisXAxisByTimeModel["buckets_presented"] = int(10)

// 		model := make(map[string]interface{})
// 		model["time"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartXAxisXAxisByTimeModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsBarChartXAxisXAxisByTimeModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChartXAxisXAxisByTime)
// 	apisDashboardsV1AstWidgetsBarChartXAxisXAxisByTimeModel.Interval = core.StringPtr("60s")
// 	apisDashboardsV1AstWidgetsBarChartXAxisXAxisByTimeModel.BucketsPresented = core.Int64Ptr(int64(10))

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsBarChartXAxisTypeTime)
// 	model.Time = apisDashboardsV1AstWidgetsBarChartXAxisXAxisByTimeModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsBarChartXAxisTypeTimeToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["group_names_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["stacked_group_name_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsHorizontalBarChartQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsHorizontalBarChartQueryModel["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel}

// 		apisDashboardsV1AstWidgetsHorizontalBarChartStackDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsHorizontalBarChartStackDefinitionModel["max_slices_per_bar"] = int(5)
// 		apisDashboardsV1AstWidgetsHorizontalBarChartStackDefinitionModel["stack_name_template"] = "{{severity}}"

// 		apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel := make(map[string]interface{})

// 		apisDashboardsV1AstWidgetsCommonColorsByModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonColorsByModel["stack"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel}

// 		apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategoryModel := make(map[string]interface{})

// 		apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByModel["category"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategoryModel}

// 		model := make(map[string]interface{})
// 		model["query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartQueryModel}
// 		model["max_bars_per_chart"] = int(5)
// 		model["group_name_template"] = "{{severity}}"
// 		model["stack_definition"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartStackDefinitionModel}
// 		model["scale_type"] = "unspecified"
// 		model["colors_by"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonColorsByModel}
// 		model["unit"] = "unspecified"
// 		model["display_on_bar"] = true
// 		model["y_axis_view_by"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByModel}
// 		model["sort_by"] = "unspecified"
// 		model["color_scheme"] = "classic"
// 		model["data_mode_type"] = "high_unspecified"

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartLogsQuery)
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.Aggregation = apisDashboardsV1CommonLogsAggregationModel
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.GroupNamesFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.StackedGroupNameField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsHorizontalBarChartQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartQueryValueLogs)
// 	apisDashboardsV1AstWidgetsHorizontalBarChartQueryModel.Logs = apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel

// 	apisDashboardsV1AstWidgetsHorizontalBarChartStackDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartStackDefinition)
// 	apisDashboardsV1AstWidgetsHorizontalBarChartStackDefinitionModel.MaxSlicesPerBar = core.Int64Ptr(int64(5))
// 	apisDashboardsV1AstWidgetsHorizontalBarChartStackDefinitionModel.StackNameTemplate = core.StringPtr("{{severity}}")

// 	apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsByColorsByStack)

// 	apisDashboardsV1AstWidgetsCommonColorsByModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsByValueStack)
// 	apisDashboardsV1AstWidgetsCommonColorsByModel.Stack = apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel

// 	apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategoryModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategory)

// 	apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewCategory)
// 	apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByModel.Category = apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategoryModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChart)
// 	model.Query = apisDashboardsV1AstWidgetsHorizontalBarChartQueryModel
// 	model.MaxBarsPerChart = core.Int64Ptr(int64(5))
// 	model.GroupNameTemplate = core.StringPtr("{{severity}}")
// 	model.StackDefinition = apisDashboardsV1AstWidgetsHorizontalBarChartStackDefinitionModel
// 	model.ScaleType = core.StringPtr("unspecified")
// 	model.ColorsBy = apisDashboardsV1AstWidgetsCommonColorsByModel
// 	model.Unit = core.StringPtr("unspecified")
// 	model.DisplayOnBar = core.BoolPtr(true)
// 	model.YAxisViewBy = apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByModel
// 	model.SortBy = core.StringPtr("unspecified")
// 	model.ColorScheme = core.StringPtr("classic")
// 	model.DataModeType = core.StringPtr("high_unspecified")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["group_names_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["stacked_group_name_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel}
// 		model["spans"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartSpansQueryModel}
// 		model["metrics"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartMetricsQueryModel}
// 		model["dataprime"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartDataprimeQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartLogsQuery)
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.Aggregation = apisDashboardsV1CommonLogsAggregationModel
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.GroupNamesFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.StackedGroupNameField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartQuery)
// 	model.Logs = apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel
// 	model.Spans = apisDashboardsV1AstWidgetsHorizontalBarChartSpansQueryModel
// 	model.Metrics = apisDashboardsV1AstWidgetsHorizontalBarChartMetricsQueryModel
// 	model.Dataprime = apisDashboardsV1AstWidgetsHorizontalBarChartDataprimeQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		model["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		model["group_names_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		model["stacked_group_name_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartLogsQuery)
// 	model.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	model.Aggregation = apisDashboardsV1CommonLogsAggregationModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	model.GroupNamesFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}
// 	model.StackedGroupNameField = apisDashboardsV1CommonObservationFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartSpansQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["metric_field"] = "unspecified"
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["aggregation_type"] = "unspecified"

// 		apisDashboardsV1CommonSpansAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationModel["metric_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationMetricAggregationModel}

// 		apisDashboardsV1CommonSpanFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpanFieldModel["metadata_field"] = "unspecified"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterSpansFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSpansFilterModel["field"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstFilterSpansFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		model := make(map[string]interface{})
// 		model["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		model["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSpansFilterModel}
// 		model["group_names"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		model["stacked_group_name"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.MetricField = core.StringPtr("unspecified")
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.AggregationType = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonSpansAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationModel.MetricAggregation = apisDashboardsV1CommonSpansAggregationMetricAggregationModel

// 	apisDashboardsV1CommonSpanFieldModel := new(logsv0.ApisDashboardsV1CommonSpanFieldValueMetadataField)
// 	apisDashboardsV1CommonSpanFieldModel.MetadataField = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterSpansFilterModel := new(logsv0.ApisDashboardsV1AstFilterSpansFilter)
// 	apisDashboardsV1AstFilterSpansFilterModel.Field = apisDashboardsV1CommonSpanFieldModel
// 	apisDashboardsV1AstFilterSpansFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartSpansQuery)
// 	model.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	model.Aggregation = apisDashboardsV1CommonSpansAggregationModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterSpansFilter{*apisDashboardsV1AstFilterSpansFilterModel}
// 	model.GroupNames = []logsv0.ApisDashboardsV1CommonSpanFieldIntf{apisDashboardsV1CommonSpanFieldModel}
// 	model.StackedGroupName = apisDashboardsV1CommonSpanFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartSpansQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartMetricsQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterMetricsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterMetricsFilterModel["label"] = "service"
// 		apisDashboardsV1AstFilterMetricsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		model := make(map[string]interface{})
// 		model["promql_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonPromQlQueryModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterMetricsFilterModel}
// 		model["group_names"] = []string{"testString"}
// 		model["stacked_group_name"] = "service"

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonPromQlQuery)
// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterMetricsFilterModel := new(logsv0.ApisDashboardsV1AstFilterMetricsFilter)
// 	apisDashboardsV1AstFilterMetricsFilterModel.Label = core.StringPtr("service")
// 	apisDashboardsV1AstFilterMetricsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartMetricsQuery)
// 	model.PromqlQuery = apisDashboardsV1AstWidgetsCommonPromQlQueryModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterMetricsFilter{*apisDashboardsV1AstFilterMetricsFilterModel}
// 	model.GroupNames = []string{"testString"}
// 	model.StackedGroupName = core.StringPtr("service")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartMetricsQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartDataprimeQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonDataprimeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonDataprimeQueryModel["text"] = "source logs"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstFilterSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSourceModel["logs"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}

// 		model := make(map[string]interface{})
// 		model["dataprime_query"] = []map[string]interface{}{apisDashboardsV1CommonDataprimeQueryModel}
// 		model["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSourceModel}
// 		model["group_names"] = []string{"testString"}
// 		model["stacked_group_name"] = "testString"

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonDataprimeQueryModel := new(logsv0.ApisDashboardsV1CommonDataprimeQuery)
// 	apisDashboardsV1CommonDataprimeQueryModel.Text = core.StringPtr("source logs")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstFilterSourceModel := new(logsv0.ApisDashboardsV1AstFilterSourceValueLogs)
// 	apisDashboardsV1AstFilterSourceModel.Logs = apisDashboardsV1AstFilterLogsFilterModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartDataprimeQuery)
// 	model.DataprimeQuery = apisDashboardsV1CommonDataprimeQueryModel
// 	model.Filters = []logsv0.ApisDashboardsV1AstFilterSourceIntf{apisDashboardsV1AstFilterSourceModel}
// 	model.GroupNames = []string{"testString"}
// 	model.StackedGroupName = core.StringPtr("testString")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartDataprimeQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartQueryValueLogsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["group_names_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["stacked_group_name_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartLogsQuery)
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.Aggregation = apisDashboardsV1CommonLogsAggregationModel
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.GroupNamesFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.StackedGroupNameField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartQueryValueLogs)
// 	model.Logs = apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartQueryValueLogsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartQueryValueSpansToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["metric_field"] = "unspecified"
// 		apisDashboardsV1CommonSpansAggregationMetricAggregationModel["aggregation_type"] = "unspecified"

// 		apisDashboardsV1CommonSpansAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpansAggregationModel["metric_aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationMetricAggregationModel}

// 		apisDashboardsV1CommonSpanFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpanFieldModel["metadata_field"] = "unspecified"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterSpansFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSpansFilterModel["field"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstFilterSpansFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		apisDashboardsV1AstWidgetsHorizontalBarChartSpansQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsHorizontalBarChartSpansQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartSpansQueryModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonSpansAggregationModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartSpansQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSpansFilterModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartSpansQueryModel["group_names"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartSpansQueryModel["stacked_group_name"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}

// 		model := make(map[string]interface{})
// 		model["spans"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartSpansQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.MetricField = core.StringPtr("unspecified")
// 	apisDashboardsV1CommonSpansAggregationMetricAggregationModel.AggregationType = core.StringPtr("unspecified")

// 	apisDashboardsV1CommonSpansAggregationModel := new(logsv0.ApisDashboardsV1CommonSpansAggregationAggregationMetricAggregation)
// 	apisDashboardsV1CommonSpansAggregationModel.MetricAggregation = apisDashboardsV1CommonSpansAggregationMetricAggregationModel

// 	apisDashboardsV1CommonSpanFieldModel := new(logsv0.ApisDashboardsV1CommonSpanFieldValueMetadataField)
// 	apisDashboardsV1CommonSpanFieldModel.MetadataField = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterSpansFilterModel := new(logsv0.ApisDashboardsV1AstFilterSpansFilter)
// 	apisDashboardsV1AstFilterSpansFilterModel.Field = apisDashboardsV1CommonSpanFieldModel
// 	apisDashboardsV1AstFilterSpansFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	apisDashboardsV1AstWidgetsHorizontalBarChartSpansQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartSpansQuery)
// 	apisDashboardsV1AstWidgetsHorizontalBarChartSpansQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsHorizontalBarChartSpansQueryModel.Aggregation = apisDashboardsV1CommonSpansAggregationModel
// 	apisDashboardsV1AstWidgetsHorizontalBarChartSpansQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterSpansFilter{*apisDashboardsV1AstFilterSpansFilterModel}
// 	apisDashboardsV1AstWidgetsHorizontalBarChartSpansQueryModel.GroupNames = []logsv0.ApisDashboardsV1CommonSpanFieldIntf{apisDashboardsV1CommonSpanFieldModel}
// 	apisDashboardsV1AstWidgetsHorizontalBarChartSpansQueryModel.StackedGroupName = apisDashboardsV1CommonSpanFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartQueryValueSpans)
// 	model.Spans = apisDashboardsV1AstWidgetsHorizontalBarChartSpansQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartQueryValueSpansToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartQueryValueMetricsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterMetricsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterMetricsFilterModel["label"] = "service"
// 		apisDashboardsV1AstFilterMetricsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		apisDashboardsV1AstWidgetsHorizontalBarChartMetricsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsHorizontalBarChartMetricsQueryModel["promql_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonPromQlQueryModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartMetricsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterMetricsFilterModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartMetricsQueryModel["group_names"] = []string{"testString"}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartMetricsQueryModel["stacked_group_name"] = "service"

// 		model := make(map[string]interface{})
// 		model["metrics"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartMetricsQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonPromQlQuery)
// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterMetricsFilterModel := new(logsv0.ApisDashboardsV1AstFilterMetricsFilter)
// 	apisDashboardsV1AstFilterMetricsFilterModel.Label = core.StringPtr("service")
// 	apisDashboardsV1AstFilterMetricsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	apisDashboardsV1AstWidgetsHorizontalBarChartMetricsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartMetricsQuery)
// 	apisDashboardsV1AstWidgetsHorizontalBarChartMetricsQueryModel.PromqlQuery = apisDashboardsV1AstWidgetsCommonPromQlQueryModel
// 	apisDashboardsV1AstWidgetsHorizontalBarChartMetricsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterMetricsFilter{*apisDashboardsV1AstFilterMetricsFilterModel}
// 	apisDashboardsV1AstWidgetsHorizontalBarChartMetricsQueryModel.GroupNames = []string{"testString"}
// 	apisDashboardsV1AstWidgetsHorizontalBarChartMetricsQueryModel.StackedGroupName = core.StringPtr("service")

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartQueryValueMetrics)
// 	model.Metrics = apisDashboardsV1AstWidgetsHorizontalBarChartMetricsQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartQueryValueMetricsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartQueryValueDataprimeToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonDataprimeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonDataprimeQueryModel["text"] = "source logs"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstFilterSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSourceModel["logs"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}

// 		apisDashboardsV1AstWidgetsHorizontalBarChartDataprimeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsHorizontalBarChartDataprimeQueryModel["dataprime_query"] = []map[string]interface{}{apisDashboardsV1CommonDataprimeQueryModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartDataprimeQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterSourceModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartDataprimeQueryModel["group_names"] = []string{"testString"}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartDataprimeQueryModel["stacked_group_name"] = "testString"

// 		model := make(map[string]interface{})
// 		model["dataprime"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartDataprimeQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonDataprimeQueryModel := new(logsv0.ApisDashboardsV1CommonDataprimeQuery)
// 	apisDashboardsV1CommonDataprimeQueryModel.Text = core.StringPtr("source logs")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstFilterSourceModel := new(logsv0.ApisDashboardsV1AstFilterSourceValueLogs)
// 	apisDashboardsV1AstFilterSourceModel.Logs = apisDashboardsV1AstFilterLogsFilterModel

// 	apisDashboardsV1AstWidgetsHorizontalBarChartDataprimeQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartDataprimeQuery)
// 	apisDashboardsV1AstWidgetsHorizontalBarChartDataprimeQueryModel.DataprimeQuery = apisDashboardsV1CommonDataprimeQueryModel
// 	apisDashboardsV1AstWidgetsHorizontalBarChartDataprimeQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterSourceIntf{apisDashboardsV1AstFilterSourceModel}
// 	apisDashboardsV1AstWidgetsHorizontalBarChartDataprimeQueryModel.GroupNames = []string{"testString"}
// 	apisDashboardsV1AstWidgetsHorizontalBarChartDataprimeQueryModel.StackedGroupName = core.StringPtr("testString")

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartQueryValueDataprime)
// 	model.Dataprime = apisDashboardsV1AstWidgetsHorizontalBarChartDataprimeQueryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartQueryValueDataprimeToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartStackDefinitionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["max_slices_per_bar"] = int(5)
// 		model["stack_name_template"] = "{{severity}}"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartStackDefinition)
// 	model.MaxSlicesPerBar = core.Int64Ptr(int64(5))
// 	model.StackNameTemplate = core.StringPtr("{{severity}}")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartStackDefinitionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategoryModel := make(map[string]interface{})

// 		model := make(map[string]interface{})
// 		model["category"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategoryModel}
// 		model["value"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByValueModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategoryModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategory)

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewBy)
// 	model.Category = apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategoryModel
// 	model.Value = apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByValueModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategoryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategory)

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategoryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByValueToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByValue)

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByValueToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewCategoryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategoryModel := make(map[string]interface{})

// 		model := make(map[string]interface{})
// 		model["category"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategoryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategoryModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategory)

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewCategory)
// 	model.Category = apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategoryModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewCategoryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewValueToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByValueModel := make(map[string]interface{})

// 		model := make(map[string]interface{})
// 		model["value"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByValueModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByValueModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByValue)

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewValue)
// 	model.Value = apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByValueModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewValueToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsMarkdownToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["markdown_text"] = "# Database metrics"
// 		model["tooltip_text"] = " # Database metrics"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetsMarkdown)
// 	model.MarkdownText = core.StringPtr("# Database metrics")
// 	model.TooltipText = core.StringPtr(" # Database metrics")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetsMarkdownToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetDefinitionValueLineChartToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLegendModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLegendModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsCommonLegendModel["columns"] = []string{"unspecified"}
// 		apisDashboardsV1AstWidgetsCommonLegendModel["group_by_query"] = true

// 		apisDashboardsV1AstWidgetsLineChartTooltipModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartTooltipModel["show_labels"] = true
// 		apisDashboardsV1AstWidgetsLineChartTooltipModel["type"] = "unspecified"

// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_by"] = []string{"testString"}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["aggregations"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsLineChartLogsQueryModel["group_bys"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsLineChartQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartQueryModel["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartLogsQueryModel}

// 		apisDashboardsV1AstWidgetsLineChartResolutionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartResolutionModel["interval"] = "1m"
// 		apisDashboardsV1AstWidgetsLineChartResolutionModel["buckets_presented"] = int(100)

// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartQueryModel}
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["series_name_template"] = "{{severity}}"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["series_count_limit"] = "10"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["unit"] = "unspecified"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["scale_type"] = "unspecified"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["name"] = "CPU usage"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["color_scheme"] = "classic"
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["resolution"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartResolutionModel}
// 		apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel["data_mode_type"] = "high_unspecified"

// 		apisDashboardsV1AstWidgetsLineChartModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsLineChartModel["legend"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLegendModel}
// 		apisDashboardsV1AstWidgetsLineChartModel["tooltip"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartTooltipModel}
// 		apisDashboardsV1AstWidgetsLineChartModel["query_definitions"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel}

// 		model := make(map[string]interface{})
// 		model["line_chart"] = []map[string]interface{}{apisDashboardsV1AstWidgetsLineChartModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLegendModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLegend)
// 	apisDashboardsV1AstWidgetsCommonLegendModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsCommonLegendModel.Columns = []string{"unspecified"}
// 	apisDashboardsV1AstWidgetsCommonLegendModel.GroupByQuery = core.BoolPtr(true)

// 	apisDashboardsV1AstWidgetsLineChartTooltipModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartTooltip)
// 	apisDashboardsV1AstWidgetsLineChartTooltipModel.ShowLabels = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsLineChartTooltipModel.Type = core.StringPtr("unspecified")

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartLogsQuery)
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBy = []string{"testString"}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Aggregations = []logsv0.ApisDashboardsV1CommonLogsAggregationIntf{apisDashboardsV1CommonLogsAggregationModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsLineChartLogsQueryModel.GroupBys = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	apisDashboardsV1AstWidgetsLineChartQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryValueLogs)
// 	apisDashboardsV1AstWidgetsLineChartQueryModel.Logs = apisDashboardsV1AstWidgetsLineChartLogsQueryModel

// 	apisDashboardsV1AstWidgetsLineChartResolutionModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartResolution)
// 	apisDashboardsV1AstWidgetsLineChartResolutionModel.Interval = core.StringPtr("1m")
// 	apisDashboardsV1AstWidgetsLineChartResolutionModel.BucketsPresented = core.Int64Ptr(int64(100))

// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChartQueryDefinition)
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Query = apisDashboardsV1AstWidgetsLineChartQueryModel
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.SeriesNameTemplate = core.StringPtr("{{severity}}")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.SeriesCountLimit = core.StringPtr("10")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Unit = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ScaleType = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Name = core.StringPtr("CPU usage")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.ColorScheme = core.StringPtr("classic")
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.Resolution = apisDashboardsV1AstWidgetsLineChartResolutionModel
// 	apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel.DataModeType = core.StringPtr("high_unspecified")

// 	apisDashboardsV1AstWidgetsLineChartModel := new(logsv0.ApisDashboardsV1AstWidgetsLineChart)
// 	apisDashboardsV1AstWidgetsLineChartModel.Legend = apisDashboardsV1AstWidgetsCommonLegendModel
// 	apisDashboardsV1AstWidgetsLineChartModel.Tooltip = apisDashboardsV1AstWidgetsLineChartTooltipModel
// 	apisDashboardsV1AstWidgetsLineChartModel.QueryDefinitions = []logsv0.ApisDashboardsV1AstWidgetsLineChartQueryDefinition{*apisDashboardsV1AstWidgetsLineChartQueryDefinitionModel}

// 	model := new(logsv0.ApisDashboardsV1AstWidgetDefinitionValueLineChart)
// 	model.LineChart = apisDashboardsV1AstWidgetsLineChartModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetDefinitionValueLineChartToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetDefinitionValueDataTableToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["name"] = "count"
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}

// 		apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel["aggregations"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel}
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel["group_bys"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsDataTableLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsDataTableLogsQueryModel["grouping"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel}

// 		apisDashboardsV1AstWidgetsDataTableQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableQueryModel["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableLogsQueryModel}

// 		apisDashboardsV1AstWidgetsDataTableColumnModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableColumnModel["field"] = "coralogix.metadata.applicationName"
// 		apisDashboardsV1AstWidgetsDataTableColumnModel["width"] = int(100)

// 		apisDashboardsV1CommonOrderingFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonOrderingFieldModel["field"] = "testString"
// 		apisDashboardsV1CommonOrderingFieldModel["order_direction"] = "unspecified"

// 		apisDashboardsV1AstWidgetsDataTableModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsDataTableModel["query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableQueryModel}
// 		apisDashboardsV1AstWidgetsDataTableModel["results_per_page"] = int(10)
// 		apisDashboardsV1AstWidgetsDataTableModel["row_style"] = "unspecified"
// 		apisDashboardsV1AstWidgetsDataTableModel["columns"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableColumnModel}
// 		apisDashboardsV1AstWidgetsDataTableModel["order_by"] = []map[string]interface{}{apisDashboardsV1CommonOrderingFieldModel}
// 		apisDashboardsV1AstWidgetsDataTableModel["data_mode_type"] = "high_unspecified"

// 		model := make(map[string]interface{})
// 		model["data_table"] = []map[string]interface{}{apisDashboardsV1AstWidgetsDataTableModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryAggregation)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.Name = core.StringPtr("count")
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel.Aggregation = apisDashboardsV1CommonLogsAggregationModel

// 	apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryGrouping)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel.Aggregations = []logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQueryAggregation{*apisDashboardsV1AstWidgetsDataTableLogsQueryAggregationModel}
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel.GroupBys = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	apisDashboardsV1AstWidgetsDataTableLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableLogsQuery)
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsDataTableLogsQueryModel.Grouping = apisDashboardsV1AstWidgetsDataTableLogsQueryGroupingModel

// 	apisDashboardsV1AstWidgetsDataTableQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableQueryValueLogs)
// 	apisDashboardsV1AstWidgetsDataTableQueryModel.Logs = apisDashboardsV1AstWidgetsDataTableLogsQueryModel

// 	apisDashboardsV1AstWidgetsDataTableColumnModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTableColumn)
// 	apisDashboardsV1AstWidgetsDataTableColumnModel.Field = core.StringPtr("coralogix.metadata.applicationName")
// 	apisDashboardsV1AstWidgetsDataTableColumnModel.Width = core.Int64Ptr(int64(100))

// 	apisDashboardsV1CommonOrderingFieldModel := new(logsv0.ApisDashboardsV1CommonOrderingField)
// 	apisDashboardsV1CommonOrderingFieldModel.Field = core.StringPtr("testString")
// 	apisDashboardsV1CommonOrderingFieldModel.OrderDirection = core.StringPtr("unspecified")

// 	apisDashboardsV1AstWidgetsDataTableModel := new(logsv0.ApisDashboardsV1AstWidgetsDataTable)
// 	apisDashboardsV1AstWidgetsDataTableModel.Query = apisDashboardsV1AstWidgetsDataTableQueryModel
// 	apisDashboardsV1AstWidgetsDataTableModel.ResultsPerPage = core.Int64Ptr(int64(10))
// 	apisDashboardsV1AstWidgetsDataTableModel.RowStyle = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsDataTableModel.Columns = []logsv0.ApisDashboardsV1AstWidgetsDataTableColumn{*apisDashboardsV1AstWidgetsDataTableColumnModel}
// 	apisDashboardsV1AstWidgetsDataTableModel.OrderBy = apisDashboardsV1CommonOrderingFieldModel
// 	apisDashboardsV1AstWidgetsDataTableModel.DataModeType = core.StringPtr("high_unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstWidgetDefinitionValueDataTable)
// 	model.DataTable = apisDashboardsV1AstWidgetsDataTableModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetDefinitionValueDataTableToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetDefinitionValueGaugeToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1AstFilterMetricsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterMetricsFilterModel["label"] = "service"
// 		apisDashboardsV1AstFilterMetricsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}

// 		apisDashboardsV1AstWidgetsGaugeMetricsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsGaugeMetricsQueryModel["promql_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonPromQlQueryModel}
// 		apisDashboardsV1AstWidgetsGaugeMetricsQueryModel["aggregation"] = "unspecified"
// 		apisDashboardsV1AstWidgetsGaugeMetricsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterMetricsFilterModel}

// 		apisDashboardsV1AstWidgetsGaugeQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsGaugeQueryModel["metrics"] = []map[string]interface{}{apisDashboardsV1AstWidgetsGaugeMetricsQueryModel}

// 		apisDashboardsV1AstWidgetsGaugeThresholdModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsGaugeThresholdModel["from"] = float64(0.5)
// 		apisDashboardsV1AstWidgetsGaugeThresholdModel["color"] = "warning"

// 		apisDashboardsV1AstWidgetsGaugeModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsGaugeModel["query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsGaugeQueryModel}
// 		apisDashboardsV1AstWidgetsGaugeModel["min"] = float64(0)
// 		apisDashboardsV1AstWidgetsGaugeModel["max"] = float64(100)
// 		apisDashboardsV1AstWidgetsGaugeModel["show_inner_arc"] = true
// 		apisDashboardsV1AstWidgetsGaugeModel["show_outer_arc"] = true
// 		apisDashboardsV1AstWidgetsGaugeModel["unit"] = "unspecified"
// 		apisDashboardsV1AstWidgetsGaugeModel["thresholds"] = []map[string]interface{}{apisDashboardsV1AstWidgetsGaugeThresholdModel}
// 		apisDashboardsV1AstWidgetsGaugeModel["data_mode_type"] = "high_unspecified"
// 		apisDashboardsV1AstWidgetsGaugeModel["threshold_by"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["gauge"] = []map[string]interface{}{apisDashboardsV1AstWidgetsGaugeModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonPromQlQuery)
// 	apisDashboardsV1AstWidgetsCommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1AstFilterMetricsFilterModel := new(logsv0.ApisDashboardsV1AstFilterMetricsFilter)
// 	apisDashboardsV1AstFilterMetricsFilterModel.Label = core.StringPtr("service")
// 	apisDashboardsV1AstFilterMetricsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel

// 	apisDashboardsV1AstWidgetsGaugeMetricsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsGaugeMetricsQuery)
// 	apisDashboardsV1AstWidgetsGaugeMetricsQueryModel.PromqlQuery = apisDashboardsV1AstWidgetsCommonPromQlQueryModel
// 	apisDashboardsV1AstWidgetsGaugeMetricsQueryModel.Aggregation = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsGaugeMetricsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterMetricsFilter{*apisDashboardsV1AstFilterMetricsFilterModel}

// 	apisDashboardsV1AstWidgetsGaugeQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsGaugeQueryValueMetrics)
// 	apisDashboardsV1AstWidgetsGaugeQueryModel.Metrics = apisDashboardsV1AstWidgetsGaugeMetricsQueryModel

// 	apisDashboardsV1AstWidgetsGaugeThresholdModel := new(logsv0.ApisDashboardsV1AstWidgetsGaugeThreshold)
// 	apisDashboardsV1AstWidgetsGaugeThresholdModel.From = core.Float64Ptr(float64(0.5))
// 	apisDashboardsV1AstWidgetsGaugeThresholdModel.Color = core.StringPtr("warning")

// 	apisDashboardsV1AstWidgetsGaugeModel := new(logsv0.ApisDashboardsV1AstWidgetsGauge)
// 	apisDashboardsV1AstWidgetsGaugeModel.Query = apisDashboardsV1AstWidgetsGaugeQueryModel
// 	apisDashboardsV1AstWidgetsGaugeModel.Min = core.Float64Ptr(float64(0))
// 	apisDashboardsV1AstWidgetsGaugeModel.Max = core.Float64Ptr(float64(100))
// 	apisDashboardsV1AstWidgetsGaugeModel.ShowInnerArc = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsGaugeModel.ShowOuterArc = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsGaugeModel.Unit = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsGaugeModel.Thresholds = []logsv0.ApisDashboardsV1AstWidgetsGaugeThreshold{*apisDashboardsV1AstWidgetsGaugeThresholdModel}
// 	apisDashboardsV1AstWidgetsGaugeModel.DataModeType = core.StringPtr("high_unspecified")
// 	apisDashboardsV1AstWidgetsGaugeModel.ThresholdBy = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstWidgetDefinitionValueGauge)
// 	model.Gauge = apisDashboardsV1AstWidgetsGaugeModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetDefinitionValueGaugeToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetDefinitionValuePieChartToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["group_names_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		apisDashboardsV1AstWidgetsPieChartLogsQueryModel["stacked_group_name_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsPieChartQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsPieChartQueryModel["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsPieChartLogsQueryModel}

// 		apisDashboardsV1AstWidgetsPieChartStackDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsPieChartStackDefinitionModel["max_slices_per_stack"] = int(5)
// 		apisDashboardsV1AstWidgetsPieChartStackDefinitionModel["stack_name_template"] = "{{severity}}"

// 		apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel["label_source"] = "unspecified"
// 		apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel["is_visible"] = true
// 		apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel["show_name"] = true
// 		apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel["show_value"] = true
// 		apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel["show_percentage"] = true

// 		apisDashboardsV1AstWidgetsPieChartModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsPieChartModel["query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsPieChartQueryModel}
// 		apisDashboardsV1AstWidgetsPieChartModel["max_slices_per_chart"] = int(5)
// 		apisDashboardsV1AstWidgetsPieChartModel["min_slice_percentage"] = int(1)
// 		apisDashboardsV1AstWidgetsPieChartModel["stack_definition"] = []map[string]interface{}{apisDashboardsV1AstWidgetsPieChartStackDefinitionModel}
// 		apisDashboardsV1AstWidgetsPieChartModel["label_definition"] = []map[string]interface{}{apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel}
// 		apisDashboardsV1AstWidgetsPieChartModel["show_legend"] = true
// 		apisDashboardsV1AstWidgetsPieChartModel["group_name_template"] = "testString"
// 		apisDashboardsV1AstWidgetsPieChartModel["unit"] = "unspecified"
// 		apisDashboardsV1AstWidgetsPieChartModel["color_scheme"] = "classic"
// 		apisDashboardsV1AstWidgetsPieChartModel["data_mode_type"] = "high_unspecified"

// 		model := make(map[string]interface{})
// 		model["pie_chart"] = []map[string]interface{}{apisDashboardsV1AstWidgetsPieChartModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsPieChartLogsQuery)
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.Aggregation = apisDashboardsV1CommonLogsAggregationModel
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.GroupNamesFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}
// 	apisDashboardsV1AstWidgetsPieChartLogsQueryModel.StackedGroupNameField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsPieChartQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsPieChartQueryValueLogs)
// 	apisDashboardsV1AstWidgetsPieChartQueryModel.Logs = apisDashboardsV1AstWidgetsPieChartLogsQueryModel

// 	apisDashboardsV1AstWidgetsPieChartStackDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetsPieChartStackDefinition)
// 	apisDashboardsV1AstWidgetsPieChartStackDefinitionModel.MaxSlicesPerStack = core.Int64Ptr(int64(5))
// 	apisDashboardsV1AstWidgetsPieChartStackDefinitionModel.StackNameTemplate = core.StringPtr("{{severity}}")

// 	apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetsPieChartLabelDefinition)
// 	apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel.LabelSource = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel.IsVisible = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel.ShowName = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel.ShowValue = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel.ShowPercentage = core.BoolPtr(true)

// 	apisDashboardsV1AstWidgetsPieChartModel := new(logsv0.ApisDashboardsV1AstWidgetsPieChart)
// 	apisDashboardsV1AstWidgetsPieChartModel.Query = apisDashboardsV1AstWidgetsPieChartQueryModel
// 	apisDashboardsV1AstWidgetsPieChartModel.MaxSlicesPerChart = core.Int64Ptr(int64(5))
// 	apisDashboardsV1AstWidgetsPieChartModel.MinSlicePercentage = core.Int64Ptr(int64(1))
// 	apisDashboardsV1AstWidgetsPieChartModel.StackDefinition = apisDashboardsV1AstWidgetsPieChartStackDefinitionModel
// 	apisDashboardsV1AstWidgetsPieChartModel.LabelDefinition = apisDashboardsV1AstWidgetsPieChartLabelDefinitionModel
// 	apisDashboardsV1AstWidgetsPieChartModel.ShowLegend = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsPieChartModel.GroupNameTemplate = core.StringPtr("testString")
// 	apisDashboardsV1AstWidgetsPieChartModel.Unit = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsPieChartModel.ColorScheme = core.StringPtr("classic")
// 	apisDashboardsV1AstWidgetsPieChartModel.DataModeType = core.StringPtr("high_unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstWidgetDefinitionValuePieChart)
// 	model.PieChart = apisDashboardsV1AstWidgetsPieChartModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetDefinitionValuePieChartToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetDefinitionValueBarChartToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["group_names_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		apisDashboardsV1AstWidgetsBarChartLogsQueryModel["stacked_group_name_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsBarChartQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsBarChartQueryModel["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartLogsQueryModel}

// 		apisDashboardsV1AstWidgetsBarChartStackDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsBarChartStackDefinitionModel["max_slices_per_bar"] = int(5)
// 		apisDashboardsV1AstWidgetsBarChartStackDefinitionModel["stack_name_template"] = "{{severity}}"

// 		apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel := make(map[string]interface{})

// 		apisDashboardsV1AstWidgetsCommonColorsByModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonColorsByModel["stack"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel}

// 		apisDashboardsV1AstWidgetsBarChartXAxisXAxisByValueModel := make(map[string]interface{})

// 		apisDashboardsV1AstWidgetsBarChartXAxisModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsBarChartXAxisModel["value"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartXAxisXAxisByValueModel}

// 		apisDashboardsV1AstWidgetsBarChartModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsBarChartModel["query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartQueryModel}
// 		apisDashboardsV1AstWidgetsBarChartModel["max_bars_per_chart"] = int(10)
// 		apisDashboardsV1AstWidgetsBarChartModel["group_name_template"] = "{{severity}}"
// 		apisDashboardsV1AstWidgetsBarChartModel["stack_definition"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartStackDefinitionModel}
// 		apisDashboardsV1AstWidgetsBarChartModel["scale_type"] = "unspecified"
// 		apisDashboardsV1AstWidgetsBarChartModel["colors_by"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonColorsByModel}
// 		apisDashboardsV1AstWidgetsBarChartModel["x_axis"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartXAxisModel}
// 		apisDashboardsV1AstWidgetsBarChartModel["unit"] = "unspecified"
// 		apisDashboardsV1AstWidgetsBarChartModel["sort_by"] = "unspecified"
// 		apisDashboardsV1AstWidgetsBarChartModel["color_scheme"] = "classic"
// 		apisDashboardsV1AstWidgetsBarChartModel["data_mode_type"] = "high_unspecified"

// 		model := make(map[string]interface{})
// 		model["bar_chart"] = []map[string]interface{}{apisDashboardsV1AstWidgetsBarChartModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChartLogsQuery)
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.Aggregation = apisDashboardsV1CommonLogsAggregationModel
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.GroupNamesFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}
// 	apisDashboardsV1AstWidgetsBarChartLogsQueryModel.StackedGroupNameField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsBarChartQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChartQueryValueLogs)
// 	apisDashboardsV1AstWidgetsBarChartQueryModel.Logs = apisDashboardsV1AstWidgetsBarChartLogsQueryModel

// 	apisDashboardsV1AstWidgetsBarChartStackDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChartStackDefinition)
// 	apisDashboardsV1AstWidgetsBarChartStackDefinitionModel.MaxSlicesPerBar = core.Int64Ptr(int64(5))
// 	apisDashboardsV1AstWidgetsBarChartStackDefinitionModel.StackNameTemplate = core.StringPtr("{{severity}}")

// 	apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsByColorsByStack)

// 	apisDashboardsV1AstWidgetsCommonColorsByModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsByValueStack)
// 	apisDashboardsV1AstWidgetsCommonColorsByModel.Stack = apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel

// 	apisDashboardsV1AstWidgetsBarChartXAxisXAxisByValueModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChartXAxisXAxisByValue)

// 	apisDashboardsV1AstWidgetsBarChartXAxisModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChartXAxisTypeValue)
// 	apisDashboardsV1AstWidgetsBarChartXAxisModel.Value = apisDashboardsV1AstWidgetsBarChartXAxisXAxisByValueModel

// 	apisDashboardsV1AstWidgetsBarChartModel := new(logsv0.ApisDashboardsV1AstWidgetsBarChart)
// 	apisDashboardsV1AstWidgetsBarChartModel.Query = apisDashboardsV1AstWidgetsBarChartQueryModel
// 	apisDashboardsV1AstWidgetsBarChartModel.MaxBarsPerChart = core.Int64Ptr(int64(10))
// 	apisDashboardsV1AstWidgetsBarChartModel.GroupNameTemplate = core.StringPtr("{{severity}}")
// 	apisDashboardsV1AstWidgetsBarChartModel.StackDefinition = apisDashboardsV1AstWidgetsBarChartStackDefinitionModel
// 	apisDashboardsV1AstWidgetsBarChartModel.ScaleType = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsBarChartModel.ColorsBy = apisDashboardsV1AstWidgetsCommonColorsByModel
// 	apisDashboardsV1AstWidgetsBarChartModel.XAxis = apisDashboardsV1AstWidgetsBarChartXAxisModel
// 	apisDashboardsV1AstWidgetsBarChartModel.Unit = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsBarChartModel.SortBy = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsBarChartModel.ColorScheme = core.StringPtr("classic")
// 	apisDashboardsV1AstWidgetsBarChartModel.DataModeType = core.StringPtr("high_unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstWidgetDefinitionValueBarChart)
// 	model.BarChart = apisDashboardsV1AstWidgetsBarChartModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetDefinitionValueBarChartToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetDefinitionValueHorizontalBarChartToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonLogsAggregationCountModel := make(map[string]interface{})

// 		apisDashboardsV1CommonLogsAggregationModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLogsAggregationModel["count"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationCountModel}

// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonLuceneQueryModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["aggregation"] = []map[string]interface{}{apisDashboardsV1CommonLogsAggregationModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["filters"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["group_names_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel["stacked_group_name_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstWidgetsHorizontalBarChartQueryModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsHorizontalBarChartQueryModel["logs"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel}

// 		apisDashboardsV1AstWidgetsHorizontalBarChartStackDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsHorizontalBarChartStackDefinitionModel["max_slices_per_bar"] = int(5)
// 		apisDashboardsV1AstWidgetsHorizontalBarChartStackDefinitionModel["stack_name_template"] = "{{severity}}"

// 		apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel := make(map[string]interface{})

// 		apisDashboardsV1AstWidgetsCommonColorsByModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsCommonColorsByModel["stack"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel}

// 		apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategoryModel := make(map[string]interface{})

// 		apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByModel["category"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategoryModel}

// 		apisDashboardsV1AstWidgetsHorizontalBarChartModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsHorizontalBarChartModel["query"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartQueryModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartModel["max_bars_per_chart"] = int(5)
// 		apisDashboardsV1AstWidgetsHorizontalBarChartModel["group_name_template"] = "{{severity}}"
// 		apisDashboardsV1AstWidgetsHorizontalBarChartModel["stack_definition"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartStackDefinitionModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartModel["scale_type"] = "unspecified"
// 		apisDashboardsV1AstWidgetsHorizontalBarChartModel["colors_by"] = []map[string]interface{}{apisDashboardsV1AstWidgetsCommonColorsByModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartModel["unit"] = "unspecified"
// 		apisDashboardsV1AstWidgetsHorizontalBarChartModel["display_on_bar"] = true
// 		apisDashboardsV1AstWidgetsHorizontalBarChartModel["y_axis_view_by"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByModel}
// 		apisDashboardsV1AstWidgetsHorizontalBarChartModel["sort_by"] = "unspecified"
// 		apisDashboardsV1AstWidgetsHorizontalBarChartModel["color_scheme"] = "classic"
// 		apisDashboardsV1AstWidgetsHorizontalBarChartModel["data_mode_type"] = "high_unspecified"

// 		model := make(map[string]interface{})
// 		model["horizontal_bar_chart"] = []map[string]interface{}{apisDashboardsV1AstWidgetsHorizontalBarChartModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonLuceneQuery)
// 	apisDashboardsV1AstWidgetsCommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonLogsAggregationCountModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationCount)

// 	apisDashboardsV1CommonLogsAggregationModel := new(logsv0.ApisDashboardsV1CommonLogsAggregationValueCount)
// 	apisDashboardsV1CommonLogsAggregationModel.Count = apisDashboardsV1CommonLogsAggregationCountModel

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartLogsQuery)
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.LuceneQuery = apisDashboardsV1AstWidgetsCommonLuceneQueryModel
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.Aggregation = apisDashboardsV1CommonLogsAggregationModel
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.Filters = []logsv0.ApisDashboardsV1AstFilterLogsFilter{*apisDashboardsV1AstFilterLogsFilterModel}
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.GroupNamesFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}
// 	apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel.StackedGroupNameField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstWidgetsHorizontalBarChartQueryModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartQueryValueLogs)
// 	apisDashboardsV1AstWidgetsHorizontalBarChartQueryModel.Logs = apisDashboardsV1AstWidgetsHorizontalBarChartLogsQueryModel

// 	apisDashboardsV1AstWidgetsHorizontalBarChartStackDefinitionModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartStackDefinition)
// 	apisDashboardsV1AstWidgetsHorizontalBarChartStackDefinitionModel.MaxSlicesPerBar = core.Int64Ptr(int64(5))
// 	apisDashboardsV1AstWidgetsHorizontalBarChartStackDefinitionModel.StackNameTemplate = core.StringPtr("{{severity}}")

// 	apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsByColorsByStack)

// 	apisDashboardsV1AstWidgetsCommonColorsByModel := new(logsv0.ApisDashboardsV1AstWidgetsCommonColorsByValueStack)
// 	apisDashboardsV1AstWidgetsCommonColorsByModel.Stack = apisDashboardsV1AstWidgetsCommonColorsByColorsByStackModel

// 	apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategoryModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategory)

// 	apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewCategory)
// 	apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByModel.Category = apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByYAxisViewByCategoryModel

// 	apisDashboardsV1AstWidgetsHorizontalBarChartModel := new(logsv0.ApisDashboardsV1AstWidgetsHorizontalBarChart)
// 	apisDashboardsV1AstWidgetsHorizontalBarChartModel.Query = apisDashboardsV1AstWidgetsHorizontalBarChartQueryModel
// 	apisDashboardsV1AstWidgetsHorizontalBarChartModel.MaxBarsPerChart = core.Int64Ptr(int64(5))
// 	apisDashboardsV1AstWidgetsHorizontalBarChartModel.GroupNameTemplate = core.StringPtr("{{severity}}")
// 	apisDashboardsV1AstWidgetsHorizontalBarChartModel.StackDefinition = apisDashboardsV1AstWidgetsHorizontalBarChartStackDefinitionModel
// 	apisDashboardsV1AstWidgetsHorizontalBarChartModel.ScaleType = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsHorizontalBarChartModel.ColorsBy = apisDashboardsV1AstWidgetsCommonColorsByModel
// 	apisDashboardsV1AstWidgetsHorizontalBarChartModel.Unit = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsHorizontalBarChartModel.DisplayOnBar = core.BoolPtr(true)
// 	apisDashboardsV1AstWidgetsHorizontalBarChartModel.YAxisViewBy = apisDashboardsV1AstWidgetsHorizontalBarChartYAxisViewByModel
// 	apisDashboardsV1AstWidgetsHorizontalBarChartModel.SortBy = core.StringPtr("unspecified")
// 	apisDashboardsV1AstWidgetsHorizontalBarChartModel.ColorScheme = core.StringPtr("classic")
// 	apisDashboardsV1AstWidgetsHorizontalBarChartModel.DataModeType = core.StringPtr("high_unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstWidgetDefinitionValueHorizontalBarChart)
// 	model.HorizontalBarChart = apisDashboardsV1AstWidgetsHorizontalBarChartModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetDefinitionValueHorizontalBarChartToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstWidgetDefinitionValueMarkdownToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstWidgetsMarkdownModel := make(map[string]interface{})
// 		apisDashboardsV1AstWidgetsMarkdownModel["markdown_text"] = "# Database metrics"
// 		apisDashboardsV1AstWidgetsMarkdownModel["tooltip_text"] = " # Database metrics"

// 		model := make(map[string]interface{})
// 		model["markdown"] = []map[string]interface{}{apisDashboardsV1AstWidgetsMarkdownModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstWidgetsMarkdownModel := new(logsv0.ApisDashboardsV1AstWidgetsMarkdown)
// 	apisDashboardsV1AstWidgetsMarkdownModel.MarkdownText = core.StringPtr("# Database metrics")
// 	apisDashboardsV1AstWidgetsMarkdownModel.TooltipText = core.StringPtr(" # Database metrics")

// 	model := new(logsv0.ApisDashboardsV1AstWidgetDefinitionValueMarkdown)
// 	model.Markdown = apisDashboardsV1AstWidgetsMarkdownModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstWidgetDefinitionValueMarkdownToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstVariableToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstMultiSelectLogsPathSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectLogsPathSourceModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstMultiSelectSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectSourceModel["logs_path"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectLogsPathSourceModel}

// 		apisDashboardsV1AstMultiSelectSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstMultiSelectSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectSelectionAllSelectionModel}

// 		apisDashboardsV1AstMultiSelectModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectModel["source"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectSourceModel}
// 		apisDashboardsV1AstMultiSelectModel["selection"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectSelectionModel}
// 		apisDashboardsV1AstMultiSelectModel["values_order_direction"] = "unspecified"

// 		apisDashboardsV1AstVariableDefinitionModel := make(map[string]interface{})
// 		apisDashboardsV1AstVariableDefinitionModel["multi_select"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectModel}

// 		model := make(map[string]interface{})
// 		model["name"] = "service_name"
// 		model["definition"] = []map[string]interface{}{apisDashboardsV1AstVariableDefinitionModel}
// 		model["display_name"] = "Service Name"

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstMultiSelectLogsPathSourceModel := new(logsv0.ApisDashboardsV1AstMultiSelectLogsPathSource)
// 	apisDashboardsV1AstMultiSelectLogsPathSourceModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstMultiSelectSourceModel := new(logsv0.ApisDashboardsV1AstMultiSelectSourceValueLogsPath)
// 	apisDashboardsV1AstMultiSelectSourceModel.LogsPath = apisDashboardsV1AstMultiSelectLogsPathSourceModel

// 	apisDashboardsV1AstMultiSelectSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstMultiSelectSelectionAllSelection)

// 	apisDashboardsV1AstMultiSelectSelectionModel := new(logsv0.ApisDashboardsV1AstMultiSelectSelectionValueAll)
// 	apisDashboardsV1AstMultiSelectSelectionModel.All = apisDashboardsV1AstMultiSelectSelectionAllSelectionModel

// 	apisDashboardsV1AstMultiSelectModel := new(logsv0.ApisDashboardsV1AstMultiSelect)
// 	apisDashboardsV1AstMultiSelectModel.Source = apisDashboardsV1AstMultiSelectSourceModel
// 	apisDashboardsV1AstMultiSelectModel.Selection = apisDashboardsV1AstMultiSelectSelectionModel
// 	apisDashboardsV1AstMultiSelectModel.ValuesOrderDirection = core.StringPtr("unspecified")

// 	apisDashboardsV1AstVariableDefinitionModel := new(logsv0.ApisDashboardsV1AstVariableDefinitionValueMultiSelect)
// 	apisDashboardsV1AstVariableDefinitionModel.MultiSelect = apisDashboardsV1AstMultiSelectModel

// 	model := new(logsv0.ApisDashboardsV1AstVariable)
// 	model.Name = core.StringPtr("service_name")
// 	model.Definition = apisDashboardsV1AstVariableDefinitionModel
// 	model.DisplayName = core.StringPtr("Service Name")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstVariableToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstVariableDefinitionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstMultiSelectLogsPathSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectLogsPathSourceModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstMultiSelectSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectSourceModel["logs_path"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectLogsPathSourceModel}

// 		apisDashboardsV1AstMultiSelectSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstMultiSelectSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectSelectionAllSelectionModel}

// 		apisDashboardsV1AstMultiSelectModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectModel["source"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectSourceModel}
// 		apisDashboardsV1AstMultiSelectModel["selection"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectSelectionModel}
// 		apisDashboardsV1AstMultiSelectModel["values_order_direction"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["multi_select"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstMultiSelectLogsPathSourceModel := new(logsv0.ApisDashboardsV1AstMultiSelectLogsPathSource)
// 	apisDashboardsV1AstMultiSelectLogsPathSourceModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstMultiSelectSourceModel := new(logsv0.ApisDashboardsV1AstMultiSelectSourceValueLogsPath)
// 	apisDashboardsV1AstMultiSelectSourceModel.LogsPath = apisDashboardsV1AstMultiSelectLogsPathSourceModel

// 	apisDashboardsV1AstMultiSelectSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstMultiSelectSelectionAllSelection)

// 	apisDashboardsV1AstMultiSelectSelectionModel := new(logsv0.ApisDashboardsV1AstMultiSelectSelectionValueAll)
// 	apisDashboardsV1AstMultiSelectSelectionModel.All = apisDashboardsV1AstMultiSelectSelectionAllSelectionModel

// 	apisDashboardsV1AstMultiSelectModel := new(logsv0.ApisDashboardsV1AstMultiSelect)
// 	apisDashboardsV1AstMultiSelectModel.Source = apisDashboardsV1AstMultiSelectSourceModel
// 	apisDashboardsV1AstMultiSelectModel.Selection = apisDashboardsV1AstMultiSelectSelectionModel
// 	apisDashboardsV1AstMultiSelectModel.ValuesOrderDirection = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstVariableDefinition)
// 	model.MultiSelect = apisDashboardsV1AstMultiSelectModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstVariableDefinitionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstMultiSelectLogsPathSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectLogsPathSourceModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstMultiSelectSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectSourceModel["logs_path"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectLogsPathSourceModel}

// 		apisDashboardsV1AstMultiSelectSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstMultiSelectSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectSelectionAllSelectionModel}

// 		model := make(map[string]interface{})
// 		model["source"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectSourceModel}
// 		model["selection"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectSelectionModel}
// 		model["values_order_direction"] = "unspecified"

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstMultiSelectLogsPathSourceModel := new(logsv0.ApisDashboardsV1AstMultiSelectLogsPathSource)
// 	apisDashboardsV1AstMultiSelectLogsPathSourceModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstMultiSelectSourceModel := new(logsv0.ApisDashboardsV1AstMultiSelectSourceValueLogsPath)
// 	apisDashboardsV1AstMultiSelectSourceModel.LogsPath = apisDashboardsV1AstMultiSelectLogsPathSourceModel

// 	apisDashboardsV1AstMultiSelectSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstMultiSelectSelectionAllSelection)

// 	apisDashboardsV1AstMultiSelectSelectionModel := new(logsv0.ApisDashboardsV1AstMultiSelectSelectionValueAll)
// 	apisDashboardsV1AstMultiSelectSelectionModel.All = apisDashboardsV1AstMultiSelectSelectionAllSelectionModel

// 	model := new(logsv0.ApisDashboardsV1AstMultiSelect)
// 	model.Source = apisDashboardsV1AstMultiSelectSourceModel
// 	model.Selection = apisDashboardsV1AstMultiSelectSelectionModel
// 	model.ValuesOrderDirection = core.StringPtr("unspecified")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSourceToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstMultiSelectLogsPathSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectLogsPathSourceModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["logs_path"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectLogsPathSourceModel}
// 		model["metric_label"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectMetricLabelSourceModel}
// 		model["constant_list"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectConstantListSourceModel}
// 		model["span_field"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectSpanFieldSourceModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstMultiSelectLogsPathSourceModel := new(logsv0.ApisDashboardsV1AstMultiSelectLogsPathSource)
// 	apisDashboardsV1AstMultiSelectLogsPathSourceModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstMultiSelectSource)
// 	model.LogsPath = apisDashboardsV1AstMultiSelectLogsPathSourceModel
// 	model.MetricLabel = apisDashboardsV1AstMultiSelectMetricLabelSourceModel
// 	model.ConstantList = apisDashboardsV1AstMultiSelectConstantListSourceModel
// 	model.SpanField = apisDashboardsV1AstMultiSelectSpanFieldSourceModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSourceToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectLogsPathSourceToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstMultiSelectLogsPathSource)
// 	model.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectLogsPathSourceToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectMetricLabelSourceToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["metric_name"] = "http_requests_total"
// 		model["label"] = "service"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstMultiSelectMetricLabelSource)
// 	model.MetricName = core.StringPtr("http_requests_total")
// 	model.Label = core.StringPtr("service")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectMetricLabelSourceToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectConstantListSourceToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["values"] = []string{"testString"}

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstMultiSelectConstantListSource)
// 	model.Values = []string{"testString"}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectConstantListSourceToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSpanFieldSourceToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonSpanFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpanFieldModel["metadata_field"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["value"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonSpanFieldModel := new(logsv0.ApisDashboardsV1CommonSpanFieldValueMetadataField)
// 	apisDashboardsV1CommonSpanFieldModel.MetadataField = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstMultiSelectSpanFieldSource)
// 	model.Value = apisDashboardsV1CommonSpanFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSpanFieldSourceToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSourceValueLogsPathToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstMultiSelectLogsPathSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectLogsPathSourceModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["logs_path"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectLogsPathSourceModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstMultiSelectLogsPathSourceModel := new(logsv0.ApisDashboardsV1AstMultiSelectLogsPathSource)
// 	apisDashboardsV1AstMultiSelectLogsPathSourceModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstMultiSelectSourceValueLogsPath)
// 	model.LogsPath = apisDashboardsV1AstMultiSelectLogsPathSourceModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSourceValueLogsPathToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSourceValueMetricLabelToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstMultiSelectMetricLabelSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectMetricLabelSourceModel["metric_name"] = "http_requests_total"
// 		apisDashboardsV1AstMultiSelectMetricLabelSourceModel["label"] = "service"

// 		model := make(map[string]interface{})
// 		model["metric_label"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectMetricLabelSourceModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstMultiSelectMetricLabelSourceModel := new(logsv0.ApisDashboardsV1AstMultiSelectMetricLabelSource)
// 	apisDashboardsV1AstMultiSelectMetricLabelSourceModel.MetricName = core.StringPtr("http_requests_total")
// 	apisDashboardsV1AstMultiSelectMetricLabelSourceModel.Label = core.StringPtr("service")

// 	model := new(logsv0.ApisDashboardsV1AstMultiSelectSourceValueMetricLabel)
// 	model.MetricLabel = apisDashboardsV1AstMultiSelectMetricLabelSourceModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSourceValueMetricLabelToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSourceValueConstantListToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstMultiSelectConstantListSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectConstantListSourceModel["values"] = []string{"testString"}

// 		model := make(map[string]interface{})
// 		model["constant_list"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectConstantListSourceModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstMultiSelectConstantListSourceModel := new(logsv0.ApisDashboardsV1AstMultiSelectConstantListSource)
// 	apisDashboardsV1AstMultiSelectConstantListSourceModel.Values = []string{"testString"}

// 	model := new(logsv0.ApisDashboardsV1AstMultiSelectSourceValueConstantList)
// 	model.ConstantList = apisDashboardsV1AstMultiSelectConstantListSourceModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSourceValueConstantListToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSourceValueSpanFieldToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonSpanFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonSpanFieldModel["metadata_field"] = "unspecified"

// 		apisDashboardsV1AstMultiSelectSpanFieldSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectSpanFieldSourceModel["value"] = []map[string]interface{}{apisDashboardsV1CommonSpanFieldModel}

// 		model := make(map[string]interface{})
// 		model["span_field"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectSpanFieldSourceModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonSpanFieldModel := new(logsv0.ApisDashboardsV1CommonSpanFieldValueMetadataField)
// 	apisDashboardsV1CommonSpanFieldModel.MetadataField = core.StringPtr("unspecified")

// 	apisDashboardsV1AstMultiSelectSpanFieldSourceModel := new(logsv0.ApisDashboardsV1AstMultiSelectSpanFieldSource)
// 	apisDashboardsV1AstMultiSelectSpanFieldSourceModel.Value = apisDashboardsV1CommonSpanFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstMultiSelectSourceValueSpanField)
// 	model.SpanField = apisDashboardsV1AstMultiSelectSpanFieldSourceModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSourceValueSpanFieldToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSelectionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstMultiSelectSelectionAllSelectionModel := make(map[string]interface{})

// 		model := make(map[string]interface{})
// 		model["all"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectSelectionAllSelectionModel}
// 		model["list"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectSelectionListSelectionModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstMultiSelectSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstMultiSelectSelectionAllSelection)

// 	model := new(logsv0.ApisDashboardsV1AstMultiSelectSelection)
// 	model.All = apisDashboardsV1AstMultiSelectSelectionAllSelectionModel
// 	model.List = apisDashboardsV1AstMultiSelectSelectionListSelectionModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSelectionAllSelectionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstMultiSelectSelectionAllSelection)

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSelectionAllSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSelectionListSelectionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["values"] = []string{"testString"}

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstMultiSelectSelectionListSelection)
// 	model.Values = []string{"testString"}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSelectionListSelectionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSelectionValueAllToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstMultiSelectSelectionAllSelectionModel := make(map[string]interface{})

// 		model := make(map[string]interface{})
// 		model["all"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectSelectionAllSelectionModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstMultiSelectSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstMultiSelectSelectionAllSelection)

// 	model := new(logsv0.ApisDashboardsV1AstMultiSelectSelectionValueAll)
// 	model.All = apisDashboardsV1AstMultiSelectSelectionAllSelectionModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSelectionValueAllToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSelectionValueListToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstMultiSelectSelectionListSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectSelectionListSelectionModel["values"] = []string{"testString"}

// 		model := make(map[string]interface{})
// 		model["list"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectSelectionListSelectionModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstMultiSelectSelectionListSelectionModel := new(logsv0.ApisDashboardsV1AstMultiSelectSelectionListSelection)
// 	apisDashboardsV1AstMultiSelectSelectionListSelectionModel.Values = []string{"testString"}

// 	model := new(logsv0.ApisDashboardsV1AstMultiSelectSelectionValueList)
// 	model.List = apisDashboardsV1AstMultiSelectSelectionListSelectionModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstMultiSelectSelectionValueListToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstVariableDefinitionValueMultiSelectToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstMultiSelectLogsPathSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectLogsPathSourceModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstMultiSelectSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectSourceModel["logs_path"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectLogsPathSourceModel}

// 		apisDashboardsV1AstMultiSelectSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstMultiSelectSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectSelectionAllSelectionModel}

// 		apisDashboardsV1AstMultiSelectModel := make(map[string]interface{})
// 		apisDashboardsV1AstMultiSelectModel["source"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectSourceModel}
// 		apisDashboardsV1AstMultiSelectModel["selection"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectSelectionModel}
// 		apisDashboardsV1AstMultiSelectModel["values_order_direction"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["multi_select"] = []map[string]interface{}{apisDashboardsV1AstMultiSelectModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstMultiSelectLogsPathSourceModel := new(logsv0.ApisDashboardsV1AstMultiSelectLogsPathSource)
// 	apisDashboardsV1AstMultiSelectLogsPathSourceModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstMultiSelectSourceModel := new(logsv0.ApisDashboardsV1AstMultiSelectSourceValueLogsPath)
// 	apisDashboardsV1AstMultiSelectSourceModel.LogsPath = apisDashboardsV1AstMultiSelectLogsPathSourceModel

// 	apisDashboardsV1AstMultiSelectSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstMultiSelectSelectionAllSelection)

// 	apisDashboardsV1AstMultiSelectSelectionModel := new(logsv0.ApisDashboardsV1AstMultiSelectSelectionValueAll)
// 	apisDashboardsV1AstMultiSelectSelectionModel.All = apisDashboardsV1AstMultiSelectSelectionAllSelectionModel

// 	apisDashboardsV1AstMultiSelectModel := new(logsv0.ApisDashboardsV1AstMultiSelect)
// 	apisDashboardsV1AstMultiSelectModel.Source = apisDashboardsV1AstMultiSelectSourceModel
// 	apisDashboardsV1AstMultiSelectModel.Selection = apisDashboardsV1AstMultiSelectSelectionModel
// 	apisDashboardsV1AstMultiSelectModel.ValuesOrderDirection = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstVariableDefinitionValueMultiSelect)
// 	model.MultiSelect = apisDashboardsV1AstMultiSelectModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstVariableDefinitionValueMultiSelectToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFilterToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := make(map[string]interface{})

// 		apisDashboardsV1AstFilterEqualsSelectionModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsSelectionModel["all"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel}

// 		apisDashboardsV1AstFilterEqualsModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterEqualsModel["selection"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsSelectionModel}

// 		apisDashboardsV1AstFilterOperatorModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterOperatorModel["equals"] = []map[string]interface{}{apisDashboardsV1AstFilterEqualsModel}

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstFilterLogsFilterModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterLogsFilterModel["operator"] = []map[string]interface{}{apisDashboardsV1AstFilterOperatorModel}
// 		apisDashboardsV1AstFilterLogsFilterModel["observation_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstFilterSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstFilterSourceModel["logs"] = []map[string]interface{}{apisDashboardsV1AstFilterLogsFilterModel}

// 		model := make(map[string]interface{})
// 		model["source"] = []map[string]interface{}{apisDashboardsV1AstFilterSourceModel}
// 		model["enabled"] = true
// 		model["collapsed"] = true

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionAllSelection)

// 	apisDashboardsV1AstFilterEqualsSelectionModel := new(logsv0.ApisDashboardsV1AstFilterEqualsSelectionValueAll)
// 	apisDashboardsV1AstFilterEqualsSelectionModel.All = apisDashboardsV1AstFilterEqualsSelectionAllSelectionModel

// 	apisDashboardsV1AstFilterEqualsModel := new(logsv0.ApisDashboardsV1AstFilterEquals)
// 	apisDashboardsV1AstFilterEqualsModel.Selection = apisDashboardsV1AstFilterEqualsSelectionModel

// 	apisDashboardsV1AstFilterOperatorModel := new(logsv0.ApisDashboardsV1AstFilterOperatorValueEquals)
// 	apisDashboardsV1AstFilterOperatorModel.Equals = apisDashboardsV1AstFilterEqualsModel

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstFilterLogsFilterModel := new(logsv0.ApisDashboardsV1AstFilterLogsFilter)
// 	apisDashboardsV1AstFilterLogsFilterModel.Operator = apisDashboardsV1AstFilterOperatorModel
// 	apisDashboardsV1AstFilterLogsFilterModel.ObservationField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstFilterSourceModel := new(logsv0.ApisDashboardsV1AstFilterSourceValueLogs)
// 	apisDashboardsV1AstFilterSourceModel.Logs = apisDashboardsV1AstFilterLogsFilterModel

// 	model := new(logsv0.ApisDashboardsV1AstFilter)
// 	model.Source = apisDashboardsV1AstFilterSourceModel
// 	model.Enabled = core.BoolPtr(true)
// 	model.Collapsed = core.BoolPtr(true)

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFilterToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel := make(map[string]interface{})

// 		apisDashboardsV1AstAnnotationMetricsSourceStrategyModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationMetricsSourceStrategyModel["start_time_metric"] = []map[string]interface{}{apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel}

// 		apisDashboardsV1AstAnnotationMetricsSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationMetricsSourceModel["promql_query"] = []map[string]interface{}{apisDashboardsV1CommonPromQlQueryModel}
// 		apisDashboardsV1AstAnnotationMetricsSourceModel["strategy"] = []map[string]interface{}{apisDashboardsV1AstAnnotationMetricsSourceStrategyModel}
// 		apisDashboardsV1AstAnnotationMetricsSourceModel["message_template"] = "testString"
// 		apisDashboardsV1AstAnnotationMetricsSourceModel["labels"] = []string{"testString"}

// 		apisDashboardsV1AstAnnotationSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationSourceModel["metrics"] = []map[string]interface{}{apisDashboardsV1AstAnnotationMetricsSourceModel}

// 		model := make(map[string]interface{})
// 		model["href"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		model["name"] = "Deployments"
// 		model["enabled"] = true
// 		model["source"] = []map[string]interface{}{apisDashboardsV1AstAnnotationSourceModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonPromQlQueryModel := new(logsv0.ApisDashboardsV1CommonPromQlQuery)
// 	apisDashboardsV1CommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel := new(logsv0.ApisDashboardsV1AstAnnotationMetricsSourceStartTimeMetric)

// 	apisDashboardsV1AstAnnotationMetricsSourceStrategyModel := new(logsv0.ApisDashboardsV1AstAnnotationMetricsSourceStrategy)
// 	apisDashboardsV1AstAnnotationMetricsSourceStrategyModel.StartTimeMetric = apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel

// 	apisDashboardsV1AstAnnotationMetricsSourceModel := new(logsv0.ApisDashboardsV1AstAnnotationMetricsSource)
// 	apisDashboardsV1AstAnnotationMetricsSourceModel.PromqlQuery = apisDashboardsV1CommonPromQlQueryModel
// 	apisDashboardsV1AstAnnotationMetricsSourceModel.Strategy = apisDashboardsV1AstAnnotationMetricsSourceStrategyModel
// 	apisDashboardsV1AstAnnotationMetricsSourceModel.MessageTemplate = core.StringPtr("testString")
// 	apisDashboardsV1AstAnnotationMetricsSourceModel.Labels = []string{"testString"}

// 	apisDashboardsV1AstAnnotationSourceModel := new(logsv0.ApisDashboardsV1AstAnnotationSourceValueMetrics)
// 	apisDashboardsV1AstAnnotationSourceModel.Metrics = apisDashboardsV1AstAnnotationMetricsSourceModel

// 	model := new(logsv0.ApisDashboardsV1AstAnnotation)
// 	model.Href = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	model.Name = core.StringPtr("Deployments")
// 	model.Enabled = core.BoolPtr(true)
// 	model.Source = apisDashboardsV1AstAnnotationSourceModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSourceToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel := make(map[string]interface{})

// 		apisDashboardsV1AstAnnotationMetricsSourceStrategyModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationMetricsSourceStrategyModel["start_time_metric"] = []map[string]interface{}{apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel}

// 		apisDashboardsV1AstAnnotationMetricsSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationMetricsSourceModel["promql_query"] = []map[string]interface{}{apisDashboardsV1CommonPromQlQueryModel}
// 		apisDashboardsV1AstAnnotationMetricsSourceModel["strategy"] = []map[string]interface{}{apisDashboardsV1AstAnnotationMetricsSourceStrategyModel}
// 		apisDashboardsV1AstAnnotationMetricsSourceModel["message_template"] = "testString"
// 		apisDashboardsV1AstAnnotationMetricsSourceModel["labels"] = []string{"testString"}

// 		model := make(map[string]interface{})
// 		model["metrics"] = []map[string]interface{}{apisDashboardsV1AstAnnotationMetricsSourceModel}
// 		model["logs"] = []map[string]interface{}{apisDashboardsV1AstAnnotationLogsSourceModel}
// 		model["spans"] = []map[string]interface{}{apisDashboardsV1AstAnnotationSpansSourceModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonPromQlQueryModel := new(logsv0.ApisDashboardsV1CommonPromQlQuery)
// 	apisDashboardsV1CommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel := new(logsv0.ApisDashboardsV1AstAnnotationMetricsSourceStartTimeMetric)

// 	apisDashboardsV1AstAnnotationMetricsSourceStrategyModel := new(logsv0.ApisDashboardsV1AstAnnotationMetricsSourceStrategy)
// 	apisDashboardsV1AstAnnotationMetricsSourceStrategyModel.StartTimeMetric = apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel

// 	apisDashboardsV1AstAnnotationMetricsSourceModel := new(logsv0.ApisDashboardsV1AstAnnotationMetricsSource)
// 	apisDashboardsV1AstAnnotationMetricsSourceModel.PromqlQuery = apisDashboardsV1CommonPromQlQueryModel
// 	apisDashboardsV1AstAnnotationMetricsSourceModel.Strategy = apisDashboardsV1AstAnnotationMetricsSourceStrategyModel
// 	apisDashboardsV1AstAnnotationMetricsSourceModel.MessageTemplate = core.StringPtr("testString")
// 	apisDashboardsV1AstAnnotationMetricsSourceModel.Labels = []string{"testString"}

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationSource)
// 	model.Metrics = apisDashboardsV1AstAnnotationMetricsSourceModel
// 	model.Logs = apisDashboardsV1AstAnnotationLogsSourceModel
// 	model.Spans = apisDashboardsV1AstAnnotationSpansSourceModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSourceToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationMetricsSourceToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel := make(map[string]interface{})

// 		apisDashboardsV1AstAnnotationMetricsSourceStrategyModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationMetricsSourceStrategyModel["start_time_metric"] = []map[string]interface{}{apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel}

// 		model := make(map[string]interface{})
// 		model["promql_query"] = []map[string]interface{}{apisDashboardsV1CommonPromQlQueryModel}
// 		model["strategy"] = []map[string]interface{}{apisDashboardsV1AstAnnotationMetricsSourceStrategyModel}
// 		model["message_template"] = "testString"
// 		model["labels"] = []string{"testString"}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonPromQlQueryModel := new(logsv0.ApisDashboardsV1CommonPromQlQuery)
// 	apisDashboardsV1CommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel := new(logsv0.ApisDashboardsV1AstAnnotationMetricsSourceStartTimeMetric)

// 	apisDashboardsV1AstAnnotationMetricsSourceStrategyModel := new(logsv0.ApisDashboardsV1AstAnnotationMetricsSourceStrategy)
// 	apisDashboardsV1AstAnnotationMetricsSourceStrategyModel.StartTimeMetric = apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationMetricsSource)
// 	model.PromqlQuery = apisDashboardsV1CommonPromQlQueryModel
// 	model.Strategy = apisDashboardsV1AstAnnotationMetricsSourceStrategyModel
// 	model.MessageTemplate = core.StringPtr("testString")
// 	model.Labels = []string{"testString"}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationMetricsSourceToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonPromQlQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["value"] = "sum(up)"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1CommonPromQlQuery)
// 	model.Value = core.StringPtr("sum(up)")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonPromQlQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationMetricsSourceStrategyToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel := make(map[string]interface{})

// 		model := make(map[string]interface{})
// 		model["start_time_metric"] = []map[string]interface{}{apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel := new(logsv0.ApisDashboardsV1AstAnnotationMetricsSourceStartTimeMetric)

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationMetricsSourceStrategy)
// 	model.StartTimeMetric = apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationMetricsSourceStrategyToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationMetricsSourceStartTimeMetric)

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationLogsSourceToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel["timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstAnnotationLogsSourceStrategyModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationLogsSourceStrategyModel["instant"] = []map[string]interface{}{apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel}

// 		model := make(map[string]interface{})
// 		model["lucene_query"] = []map[string]interface{}{apisDashboardsV1CommonLuceneQueryModel}
// 		model["strategy"] = []map[string]interface{}{apisDashboardsV1AstAnnotationLogsSourceStrategyModel}
// 		model["message_template"] = "testString"
// 		model["label_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonLuceneQueryModel := new(logsv0.ApisDashboardsV1CommonLuceneQuery)
// 	apisDashboardsV1CommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel := new(logsv0.ApisDashboardsV1AstAnnotationLogsSourceStrategyInstant)
// 	apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel.TimestampField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstAnnotationLogsSourceStrategyModel := new(logsv0.ApisDashboardsV1AstAnnotationLogsSourceStrategyValueInstant)
// 	apisDashboardsV1AstAnnotationLogsSourceStrategyModel.Instant = apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationLogsSource)
// 	model.LuceneQuery = apisDashboardsV1CommonLuceneQueryModel
// 	model.Strategy = apisDashboardsV1AstAnnotationLogsSourceStrategyModel
// 	model.MessageTemplate = core.StringPtr("testString")
// 	model.LabelFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationLogsSourceToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonLuceneQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["value"] = "coralogix.metadata.applicationName:"production""

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1CommonLuceneQuery)
// 	model.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonLuceneQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationLogsSourceStrategyToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel["timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["instant"] = []map[string]interface{}{apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel}
// 		model["range"] = []map[string]interface{}{apisDashboardsV1AstAnnotationLogsSourceStrategyRangeModel}
// 		model["duration"] = []map[string]interface{}{apisDashboardsV1AstAnnotationLogsSourceStrategyDurationModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel := new(logsv0.ApisDashboardsV1AstAnnotationLogsSourceStrategyInstant)
// 	apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel.TimestampField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationLogsSourceStrategy)
// 	model.Instant = apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel
// 	model.Range = apisDashboardsV1AstAnnotationLogsSourceStrategyRangeModel
// 	model.Duration = apisDashboardsV1AstAnnotationLogsSourceStrategyDurationModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationLogsSourceStrategyToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationLogsSourceStrategyInstantToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationLogsSourceStrategyInstant)
// 	model.TimestampField = apisDashboardsV1CommonObservationFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationLogsSourceStrategyInstantToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationLogsSourceStrategyRangeToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["start_timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		model["end_timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationLogsSourceStrategyRange)
// 	model.StartTimestampField = apisDashboardsV1CommonObservationFieldModel
// 	model.EndTimestampField = apisDashboardsV1CommonObservationFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationLogsSourceStrategyRangeToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationLogsSourceStrategyDurationToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["start_timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		model["duration_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationLogsSourceStrategyDuration)
// 	model.StartTimestampField = apisDashboardsV1CommonObservationFieldModel
// 	model.DurationField = apisDashboardsV1CommonObservationFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationLogsSourceStrategyDurationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationLogsSourceStrategyValueInstantToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel["timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["instant"] = []map[string]interface{}{apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel := new(logsv0.ApisDashboardsV1AstAnnotationLogsSourceStrategyInstant)
// 	apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel.TimestampField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationLogsSourceStrategyValueInstant)
// 	model.Instant = apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationLogsSourceStrategyValueInstantToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationLogsSourceStrategyValueRangeToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstAnnotationLogsSourceStrategyRangeModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationLogsSourceStrategyRangeModel["start_timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		apisDashboardsV1AstAnnotationLogsSourceStrategyRangeModel["end_timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["range"] = []map[string]interface{}{apisDashboardsV1AstAnnotationLogsSourceStrategyRangeModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstAnnotationLogsSourceStrategyRangeModel := new(logsv0.ApisDashboardsV1AstAnnotationLogsSourceStrategyRange)
// 	apisDashboardsV1AstAnnotationLogsSourceStrategyRangeModel.StartTimestampField = apisDashboardsV1CommonObservationFieldModel
// 	apisDashboardsV1AstAnnotationLogsSourceStrategyRangeModel.EndTimestampField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationLogsSourceStrategyValueRange)
// 	model.Range = apisDashboardsV1AstAnnotationLogsSourceStrategyRangeModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationLogsSourceStrategyValueRangeToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationLogsSourceStrategyValueDurationToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstAnnotationLogsSourceStrategyDurationModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationLogsSourceStrategyDurationModel["start_timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		apisDashboardsV1AstAnnotationLogsSourceStrategyDurationModel["duration_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["duration"] = []map[string]interface{}{apisDashboardsV1AstAnnotationLogsSourceStrategyDurationModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstAnnotationLogsSourceStrategyDurationModel := new(logsv0.ApisDashboardsV1AstAnnotationLogsSourceStrategyDuration)
// 	apisDashboardsV1AstAnnotationLogsSourceStrategyDurationModel.StartTimestampField = apisDashboardsV1CommonObservationFieldModel
// 	apisDashboardsV1AstAnnotationLogsSourceStrategyDurationModel.DurationField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationLogsSourceStrategyValueDuration)
// 	model.Duration = apisDashboardsV1AstAnnotationLogsSourceStrategyDurationModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationLogsSourceStrategyValueDurationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSpansSourceToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel["timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstAnnotationSpansSourceStrategyModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationSpansSourceStrategyModel["instant"] = []map[string]interface{}{apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel}

// 		model := make(map[string]interface{})
// 		model["lucene_query"] = []map[string]interface{}{apisDashboardsV1CommonLuceneQueryModel}
// 		model["strategy"] = []map[string]interface{}{apisDashboardsV1AstAnnotationSpansSourceStrategyModel}
// 		model["message_template"] = "testString"
// 		model["label_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonLuceneQueryModel := new(logsv0.ApisDashboardsV1CommonLuceneQuery)
// 	apisDashboardsV1CommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel := new(logsv0.ApisDashboardsV1AstAnnotationSpansSourceStrategyInstant)
// 	apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel.TimestampField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstAnnotationSpansSourceStrategyModel := new(logsv0.ApisDashboardsV1AstAnnotationSpansSourceStrategyValueInstant)
// 	apisDashboardsV1AstAnnotationSpansSourceStrategyModel.Instant = apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationSpansSource)
// 	model.LuceneQuery = apisDashboardsV1CommonLuceneQueryModel
// 	model.Strategy = apisDashboardsV1AstAnnotationSpansSourceStrategyModel
// 	model.MessageTemplate = core.StringPtr("testString")
// 	model.LabelFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSpansSourceToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSpansSourceStrategyToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel["timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["instant"] = []map[string]interface{}{apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel}
// 		model["range"] = []map[string]interface{}{apisDashboardsV1AstAnnotationSpansSourceStrategyRangeModel}
// 		model["duration"] = []map[string]interface{}{apisDashboardsV1AstAnnotationSpansSourceStrategyDurationModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel := new(logsv0.ApisDashboardsV1AstAnnotationSpansSourceStrategyInstant)
// 	apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel.TimestampField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationSpansSourceStrategy)
// 	model.Instant = apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel
// 	model.Range = apisDashboardsV1AstAnnotationSpansSourceStrategyRangeModel
// 	model.Duration = apisDashboardsV1AstAnnotationSpansSourceStrategyDurationModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSpansSourceStrategyToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSpansSourceStrategyInstantToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationSpansSourceStrategyInstant)
// 	model.TimestampField = apisDashboardsV1CommonObservationFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSpansSourceStrategyInstantToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSpansSourceStrategyRangeToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["start_timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		model["end_timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationSpansSourceStrategyRange)
// 	model.StartTimestampField = apisDashboardsV1CommonObservationFieldModel
// 	model.EndTimestampField = apisDashboardsV1CommonObservationFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSpansSourceStrategyRangeToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSpansSourceStrategyDurationToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["start_timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		model["duration_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationSpansSourceStrategyDuration)
// 	model.StartTimestampField = apisDashboardsV1CommonObservationFieldModel
// 	model.DurationField = apisDashboardsV1CommonObservationFieldModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSpansSourceStrategyDurationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSpansSourceStrategyValueInstantToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel["timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["instant"] = []map[string]interface{}{apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel := new(logsv0.ApisDashboardsV1AstAnnotationSpansSourceStrategyInstant)
// 	apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel.TimestampField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationSpansSourceStrategyValueInstant)
// 	model.Instant = apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSpansSourceStrategyValueInstantToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSpansSourceStrategyValueRangeToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstAnnotationSpansSourceStrategyRangeModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationSpansSourceStrategyRangeModel["start_timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		apisDashboardsV1AstAnnotationSpansSourceStrategyRangeModel["end_timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["range"] = []map[string]interface{}{apisDashboardsV1AstAnnotationSpansSourceStrategyRangeModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstAnnotationSpansSourceStrategyRangeModel := new(logsv0.ApisDashboardsV1AstAnnotationSpansSourceStrategyRange)
// 	apisDashboardsV1AstAnnotationSpansSourceStrategyRangeModel.StartTimestampField = apisDashboardsV1CommonObservationFieldModel
// 	apisDashboardsV1AstAnnotationSpansSourceStrategyRangeModel.EndTimestampField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationSpansSourceStrategyValueRange)
// 	model.Range = apisDashboardsV1AstAnnotationSpansSourceStrategyRangeModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSpansSourceStrategyValueRangeToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSpansSourceStrategyValueDurationToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstAnnotationSpansSourceStrategyDurationModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationSpansSourceStrategyDurationModel["start_timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}
// 		apisDashboardsV1AstAnnotationSpansSourceStrategyDurationModel["duration_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["duration"] = []map[string]interface{}{apisDashboardsV1AstAnnotationSpansSourceStrategyDurationModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstAnnotationSpansSourceStrategyDurationModel := new(logsv0.ApisDashboardsV1AstAnnotationSpansSourceStrategyDuration)
// 	apisDashboardsV1AstAnnotationSpansSourceStrategyDurationModel.StartTimestampField = apisDashboardsV1CommonObservationFieldModel
// 	apisDashboardsV1AstAnnotationSpansSourceStrategyDurationModel.DurationField = apisDashboardsV1CommonObservationFieldModel

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationSpansSourceStrategyValueDuration)
// 	model.Duration = apisDashboardsV1AstAnnotationSpansSourceStrategyDurationModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSpansSourceStrategyValueDurationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSourceValueMetricsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonPromQlQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonPromQlQueryModel["value"] = "sum(up)"

// 		apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel := make(map[string]interface{})

// 		apisDashboardsV1AstAnnotationMetricsSourceStrategyModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationMetricsSourceStrategyModel["start_time_metric"] = []map[string]interface{}{apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel}

// 		apisDashboardsV1AstAnnotationMetricsSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationMetricsSourceModel["promql_query"] = []map[string]interface{}{apisDashboardsV1CommonPromQlQueryModel}
// 		apisDashboardsV1AstAnnotationMetricsSourceModel["strategy"] = []map[string]interface{}{apisDashboardsV1AstAnnotationMetricsSourceStrategyModel}
// 		apisDashboardsV1AstAnnotationMetricsSourceModel["message_template"] = "testString"
// 		apisDashboardsV1AstAnnotationMetricsSourceModel["labels"] = []string{"testString"}

// 		model := make(map[string]interface{})
// 		model["metrics"] = []map[string]interface{}{apisDashboardsV1AstAnnotationMetricsSourceModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonPromQlQueryModel := new(logsv0.ApisDashboardsV1CommonPromQlQuery)
// 	apisDashboardsV1CommonPromQlQueryModel.Value = core.StringPtr("sum(up)")

// 	apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel := new(logsv0.ApisDashboardsV1AstAnnotationMetricsSourceStartTimeMetric)

// 	apisDashboardsV1AstAnnotationMetricsSourceStrategyModel := new(logsv0.ApisDashboardsV1AstAnnotationMetricsSourceStrategy)
// 	apisDashboardsV1AstAnnotationMetricsSourceStrategyModel.StartTimeMetric = apisDashboardsV1AstAnnotationMetricsSourceStartTimeMetricModel

// 	apisDashboardsV1AstAnnotationMetricsSourceModel := new(logsv0.ApisDashboardsV1AstAnnotationMetricsSource)
// 	apisDashboardsV1AstAnnotationMetricsSourceModel.PromqlQuery = apisDashboardsV1CommonPromQlQueryModel
// 	apisDashboardsV1AstAnnotationMetricsSourceModel.Strategy = apisDashboardsV1AstAnnotationMetricsSourceStrategyModel
// 	apisDashboardsV1AstAnnotationMetricsSourceModel.MessageTemplate = core.StringPtr("testString")
// 	apisDashboardsV1AstAnnotationMetricsSourceModel.Labels = []string{"testString"}

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationSourceValueMetrics)
// 	model.Metrics = apisDashboardsV1AstAnnotationMetricsSourceModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSourceValueMetricsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSourceValueLogsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel["timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstAnnotationLogsSourceStrategyModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationLogsSourceStrategyModel["instant"] = []map[string]interface{}{apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel}

// 		apisDashboardsV1AstAnnotationLogsSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationLogsSourceModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1CommonLuceneQueryModel}
// 		apisDashboardsV1AstAnnotationLogsSourceModel["strategy"] = []map[string]interface{}{apisDashboardsV1AstAnnotationLogsSourceStrategyModel}
// 		apisDashboardsV1AstAnnotationLogsSourceModel["message_template"] = "testString"
// 		apisDashboardsV1AstAnnotationLogsSourceModel["label_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["logs"] = []map[string]interface{}{apisDashboardsV1AstAnnotationLogsSourceModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonLuceneQueryModel := new(logsv0.ApisDashboardsV1CommonLuceneQuery)
// 	apisDashboardsV1CommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel := new(logsv0.ApisDashboardsV1AstAnnotationLogsSourceStrategyInstant)
// 	apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel.TimestampField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstAnnotationLogsSourceStrategyModel := new(logsv0.ApisDashboardsV1AstAnnotationLogsSourceStrategyValueInstant)
// 	apisDashboardsV1AstAnnotationLogsSourceStrategyModel.Instant = apisDashboardsV1AstAnnotationLogsSourceStrategyInstantModel

// 	apisDashboardsV1AstAnnotationLogsSourceModel := new(logsv0.ApisDashboardsV1AstAnnotationLogsSource)
// 	apisDashboardsV1AstAnnotationLogsSourceModel.LuceneQuery = apisDashboardsV1CommonLuceneQueryModel
// 	apisDashboardsV1AstAnnotationLogsSourceModel.Strategy = apisDashboardsV1AstAnnotationLogsSourceStrategyModel
// 	apisDashboardsV1AstAnnotationLogsSourceModel.MessageTemplate = core.StringPtr("testString")
// 	apisDashboardsV1AstAnnotationLogsSourceModel.LabelFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationSourceValueLogs)
// 	model.Logs = apisDashboardsV1AstAnnotationLogsSourceModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSourceValueLogsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSourceValueSpansToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisDashboardsV1CommonLuceneQueryModel := make(map[string]interface{})
// 		apisDashboardsV1CommonLuceneQueryModel["value"] = "coralogix.metadata.applicationName:"production""

// 		apisDashboardsV1CommonObservationFieldModel := make(map[string]interface{})
// 		apisDashboardsV1CommonObservationFieldModel["keypath"] = []string{"testString"}
// 		apisDashboardsV1CommonObservationFieldModel["scope"] = "unspecified"

// 		apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel["timestamp_field"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		apisDashboardsV1AstAnnotationSpansSourceStrategyModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationSpansSourceStrategyModel["instant"] = []map[string]interface{}{apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel}

// 		apisDashboardsV1AstAnnotationSpansSourceModel := make(map[string]interface{})
// 		apisDashboardsV1AstAnnotationSpansSourceModel["lucene_query"] = []map[string]interface{}{apisDashboardsV1CommonLuceneQueryModel}
// 		apisDashboardsV1AstAnnotationSpansSourceModel["strategy"] = []map[string]interface{}{apisDashboardsV1AstAnnotationSpansSourceStrategyModel}
// 		apisDashboardsV1AstAnnotationSpansSourceModel["message_template"] = "testString"
// 		apisDashboardsV1AstAnnotationSpansSourceModel["label_fields"] = []map[string]interface{}{apisDashboardsV1CommonObservationFieldModel}

// 		model := make(map[string]interface{})
// 		model["spans"] = []map[string]interface{}{apisDashboardsV1AstAnnotationSpansSourceModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisDashboardsV1CommonLuceneQueryModel := new(logsv0.ApisDashboardsV1CommonLuceneQuery)
// 	apisDashboardsV1CommonLuceneQueryModel.Value = core.StringPtr("coralogix.metadata.applicationName:"production"")

// 	apisDashboardsV1CommonObservationFieldModel := new(logsv0.ApisDashboardsV1CommonObservationField)
// 	apisDashboardsV1CommonObservationFieldModel.Keypath = []string{"testString"}
// 	apisDashboardsV1CommonObservationFieldModel.Scope = core.StringPtr("unspecified")

// 	apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel := new(logsv0.ApisDashboardsV1AstAnnotationSpansSourceStrategyInstant)
// 	apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel.TimestampField = apisDashboardsV1CommonObservationFieldModel

// 	apisDashboardsV1AstAnnotationSpansSourceStrategyModel := new(logsv0.ApisDashboardsV1AstAnnotationSpansSourceStrategyValueInstant)
// 	apisDashboardsV1AstAnnotationSpansSourceStrategyModel.Instant = apisDashboardsV1AstAnnotationSpansSourceStrategyInstantModel

// 	apisDashboardsV1AstAnnotationSpansSourceModel := new(logsv0.ApisDashboardsV1AstAnnotationSpansSource)
// 	apisDashboardsV1AstAnnotationSpansSourceModel.LuceneQuery = apisDashboardsV1CommonLuceneQueryModel
// 	apisDashboardsV1AstAnnotationSpansSourceModel.Strategy = apisDashboardsV1AstAnnotationSpansSourceStrategyModel
// 	apisDashboardsV1AstAnnotationSpansSourceModel.MessageTemplate = core.StringPtr("testString")
// 	apisDashboardsV1AstAnnotationSpansSourceModel.LabelFields = []logsv0.ApisDashboardsV1CommonObservationField{*apisDashboardsV1CommonObservationFieldModel}

// 	model := new(logsv0.ApisDashboardsV1AstAnnotationSourceValueSpans)
// 	model.Spans = apisDashboardsV1AstAnnotationSpansSourceModel

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstAnnotationSourceValueSpansToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1CommonTimeFrameToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["from"] = "2019-01-01T12:00:00.000Z"
// 		model["to"] = "2019-01-01T12:00:00.000Z"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1CommonTimeFrame)
// 	model.From = CreateMockDateTime("2019-01-01T12:00:00.000Z")
// 	model.To = CreateMockDateTime("2019-01-01T12:00:00.000Z")

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1CommonTimeFrameToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsDashboardApisDashboardsV1AstFolderPathToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["segments"] = []string{"testString"}

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisDashboardsV1AstFolderPath)
// 	model.Segments = []string{"testString"}

// 	result, err := logs.DataSourceIbmLogsDashboardApisDashboardsV1AstFolderPathToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

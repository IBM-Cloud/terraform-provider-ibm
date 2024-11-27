// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	// . "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
)

func TestAccIbmLogsDashboardDataSourceBasic(t *testing.T) {
	dashboardName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
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
	dashboardName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	dashboardDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	dashboardRelativeTimeFrame := fmt.Sprintf("%ds", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsDashboardDataSourceConfig(dashboardName, dashboardDescription, dashboardRelativeTimeFrame),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "dashboard_id"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "layout.#"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "variables.#"),
					// resource.TestCheckResourceAttr("data.ibm_logs_dashboard.logs_dashboard_instance", "variables.0.name", dashboardName),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "variables.0.display_name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "filters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "filters.0.enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "filters.0.collapsed"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "annotations.#"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "annotations.0.id"),
					// resource.TestCheckResourceAttr("data.ibm_logs_dashboard.logs_dashboard_instance", "annotations.0.name", dashboardName),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "annotations.0.enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "absolute_time_frame.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "relative_time_frame"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "folder_id.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "folder_path.#"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "false"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "two_minutes"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard.logs_dashboard_instance", "five_minutes"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsDashboardDataSourceConfigBasic(dashboardName string) string {
	return fmt.Sprintf(`
	resource "ibm_logs_dashboard" "logs_dashboard_instance" {
		instance_id = "%s"
		region      = "%s"
		name        = "%s"
		description = "test description"
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

	  data "ibm_logs_dashboard" "logs_dashboard_instance" {
		instance_id  = ibm_logs_dashboard.logs_dashboard_instance.instance_id
		region       = ibm_logs_dashboard.logs_dashboard_instance.region
		dashboard_id = ibm_logs_dashboard.logs_dashboard_instance.dashboard_id
	  }
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, dashboardName)
}

func testAccCheckIbmLogsDashboardDataSourceConfig(dashboardName string, dashboardDescription string, dashboardRelativeTimeFrame string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_dashboard" "logs_dashboard_instance" {
		instance_id = "%s"
		region      = "%s"
		name        = "%s"
		description = "%s"
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
		relative_time_frame = "%s"
	  }
	  data "ibm_logs_dashboard" "logs_dashboard_instance" {
		instance_id  = ibm_logs_dashboard.logs_dashboard_instance.instance_id
		region       = ibm_logs_dashboard.logs_dashboard_instance.region
		dashboard_id = ibm_logs_dashboard.logs_dashboard_instance.dashboard_id
	  }
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, dashboardName, dashboardDescription, dashboardRelativeTimeFrame)
}

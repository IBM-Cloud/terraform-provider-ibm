// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	// . "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func TestAccIbmLogsDashboardBasic(t *testing.T) {
	var conf logsv0.Dashboard
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsDashboardDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsDashboardConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsDashboardExists("ibm_logs_dashboard.logs_dashboard_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_dashboard.logs_dashboard_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsDashboardConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_dashboard.logs_dashboard_instance", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIbmLogsDashboardAllArgs(t *testing.T) {
	var conf logsv0.Dashboard
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	relativeTimeFrame := fmt.Sprintf("%ds", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	relativeTimeFrameUpdate := fmt.Sprintf("%ds", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsDashboardDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsDashboardConfig(name, description, relativeTimeFrame),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsDashboardExists("ibm_logs_dashboard.logs_dashboard_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_dashboard.logs_dashboard_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_dashboard.logs_dashboard_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_logs_dashboard.logs_dashboard_instance", "relative_time_frame", relativeTimeFrame),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsDashboardConfig(nameUpdate, descriptionUpdate, relativeTimeFrameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_dashboard.logs_dashboard_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_dashboard.logs_dashboard_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_logs_dashboard.logs_dashboard_instance", "relative_time_frame", relativeTimeFrameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_dashboard.logs_dashboard_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsDashboardConfigBasic(name string) string {
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
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name)
}

func testAccCheckIbmLogsDashboardConfig(name string, description string, relativeTimeFrame string) string {
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
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name, description, relativeTimeFrame)
}

func testAccCheckIbmLogsDashboardExists(n string, obj logsv0.Dashboard) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
		if err != nil {
			return err
		}
		logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getDashboardOptions := &logsv0.GetDashboardOptions{}

		getDashboardOptions.SetDashboardID(resourceID[2])

		dashboardIntf, _, err := logsClient.GetDashboard(getDashboardOptions)
		if err != nil {
			return err
		}

		dashboard := dashboardIntf.(*logsv0.Dashboard)
		obj = *dashboard
		return nil
	}
}

func testAccCheckIbmLogsDashboardDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_dashboard" {
			continue
		}
		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		getDashboardOptions := &logsv0.GetDashboardOptions{}

		getDashboardOptions.SetDashboardID(resourceID[2])

		// Try to find the key
		_, response, err := logsClient.GetDashboard(getDashboardOptions)

		if err == nil {
			return fmt.Errorf("logs_dashboard still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_dashboard (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

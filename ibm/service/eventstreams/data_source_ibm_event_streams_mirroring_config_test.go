// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventstreams_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMEventStreamsMirroringConfigDataSource(t *testing.T) {
	instanceName := "ES Integration Pipeline MZR"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEventStreamsMirroringConfigDataSource(instanceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMEventStreamsMirroringConfigProperties("data.ibm_event_streams_mirroring_config.es_mirroring_config"),
					resource.TestCheckResourceAttr("data.ibm_event_streams_mirroring_config.es_mirroring_config", "mirroring_topic_patterns.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_event_streams_mirroring_config.es_mirroring_config", "mirroring_topic_patterns.0", ".*"),
				),
			},
		},
	})
}

// check properties of the mirroring config data source or resource object
func testAccCheckIBMEventStreamsMirroringConfigProperties(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}
		mcID := rs.Primary.ID
		if mcID == "" {
			return fmt.Errorf("[ERROR] Mirroring config ID is not set")
		}
		if !strings.HasSuffix(mcID, ":mirroring-config:") {
			return fmt.Errorf("[ERROR] Mirroring config ID %s not expected CRN", mcID)
		}
		return nil
	}
}

func testAccCheckIBMEventStreamsMirroringConfigDataSource(instanceName string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "group" {
  is_default = true
}
data "ibm_resource_instance" "es_instance" {
  resource_group_id = data.ibm_resource_group.group.id
  name              = "%s"
}
data "ibm_event_streams_mirroring_config" "es_mirroring_config" {
  resource_instance_id = data.ibm_resource_instance.es_instance.id
}`, instanceName)
}

// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMEnPagerDutyTemplateDataSourceBasic(t *testing.T) {
	templateInstanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	templateName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	templateType := fmt.Sprintf("tf_type_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEnPagerDutyTemplateDataSourceConfigBasic(templateInstanceID, templateName, templateType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_pagerduty_template.en_template_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pagerduty_template.en_template_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pagerduty_template.en_template_instance", "template_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pagerduty_template.en_template_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pagerduty_template.en_template_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pagerduty_template.en_template_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pagerduty_template.en_template_instance", "subscription_count"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pagerduty_template.en_template_instance", "subscription_names.#"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pagerduty_template.en_template_instance", "updated_at"),
				),
			},
		},
	})
}

func TestAccIBMEnPagerDutyTemplateDataSourceAllArgs(t *testing.T) {
	templateInstanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	templateName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	templateDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	templateType := fmt.Sprintf("tf_type_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEnPagerDutyTemplateDataSourceConfig(templateInstanceID, templateName, templateDescription, templateType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_pagerduty_template.en_template_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pagerduty_template.en_template_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pagerduty_template.en_template_instance", "template_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pagerduty_template.en_template_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pagerduty_template.en_template_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pagerduty_template.en_template_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pagerduty_template.en_template_instance", "subscription_count"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pagerduty_template.en_template_instance", "subscription_names.#"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pagerduty_template.en_template_instance", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIBMEnPagerDutyTemplateDataSourceConfigBasic(templateInstanceID string, templateName string, templateType string) string {
	return fmt.Sprintf(`
		resource "ibm_en_pagerduty_template" "en_template_instance" {
			instance_id = "%s"
			name = "%s"
			type = "%s"
		}

		data "ibm_en_pagerduty_template" "en_template_instance" {
			instance_id = ibm_en_pagerduty_template.en_template_instance.instance_id
			template_id = ibm_en_pagerduty_template.en_template_instance.template_id
		}
	`, templateInstanceID, templateName, templateType)
}

func testAccCheckIBMEnPagerDutyTemplateDataSourceConfig(templateInstanceID string, templateName string, templateDescription string, templateType string) string {
	return fmt.Sprintf(`
		resource "ibm_en_pagerduty_template" "en_template_instance" {
			instance_id = "%s"
			name = "%s"
			description = "%s"
			type = "%s"
			params {
			    body = "ewogICJwYXlsb2FkIjogewogICAgInN1bW1hcnkiOiAie3sgZGF0YS5hbGVydF9kZWZpbml0aW9uLm5hbWV9fSIsCiAgICAidGltZXN0YW1wIjogInt7dGltZX19IiwKICAgICJzZXZlcml0eSI6ICJpbmZvIiwKICAgICJzb3VyY2UiOiAie3sgc291cmNlIH19IgogIH0sCiAgImRlZHVwX2tleSI6ICJ7eyBpZCB9fSIsCiAge3sjZXF1YWwgZGF0YS5zdGF0dXMgInRyaWdnZXJlZCJ9fQogICJldmVudF9hY3Rpb24iOiAidHJpZ2dlciIKICAge3svZXF1YWx9fQoKICB7eyNlcXVhbCBkYXRhLnN0YXR1cyAicmVzb2x2ZWQifX0KICAiZXZlbnRfYWN0aW9uIjogInJlc29sdmUiCiAge3svZXF1YWx9fQoKICAge3sjZXF1YWwgZGF0YS5zdGF0dXMgImFja25vd2xlZGdlZCJ9fQogICAiZXZlbnRfYWN0aW9uIjogImFja25vd2xlZGdlIgogICB7ey9lcXVhbH19Cn0="
			}
		}

		data "ibm_en_pagerduty_template" "en_template_instance" {
			instance_id = ibm_en_pagerduty_template.en_template_instance.instance_id
			en_template_id = ibm_en_pagerduty_template.en_template_instance.template_id
		}
	`, templateInstanceID, templateName, templateDescription, templateType)
}

// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtektonpipeline_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/continuous-delivery-go-sdk/cdtektonpipelinev2"
)

func TestAccIBMTektonPipelinePropertyBasic(t *testing.T) {
	var conf cdtektonpipelinev2.Property
	pipelineID := fmt.Sprintf("tf_pipeline_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTektonPipelinePropertyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelinePropertyConfigBasic(pipelineID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMTektonPipelinePropertyExists("ibm_cd_tekton_pipeline_property.tekton_pipeline_property", conf),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "pipeline_id", pipelineID),
				),
			},
		},
	})
}

func TestAccIBMTektonPipelinePropertyAllArgs(t *testing.T) {
	var conf cdtektonpipelinev2.Property
	pipelineID := fmt.Sprintf("tf_pipeline_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	value := fmt.Sprintf("tf_value_%d", acctest.RandIntRange(10, 100))
	defaultVar := fmt.Sprintf("tf_default_%d", acctest.RandIntRange(10, 100))
	typeVar := "SECURE"
	path := fmt.Sprintf("tf_path_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	valueUpdate := fmt.Sprintf("tf_value_%d", acctest.RandIntRange(10, 100))
	defaultVarUpdate := fmt.Sprintf("tf_default_%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := "APPCONFIG"
	pathUpdate := fmt.Sprintf("tf_path_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTektonPipelinePropertyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelinePropertyConfig(pipelineID, name, value, defaultVar, typeVar, path),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMTektonPipelinePropertyExists("ibm_cd_tekton_pipeline_property.tekton_pipeline_property", conf),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "pipeline_id", pipelineID),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "name", name),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "value", value),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "default", defaultVar),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "path", path),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelinePropertyConfig(pipelineID, nameUpdate, valueUpdate, defaultVarUpdate, typeVarUpdate, pathUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "pipeline_id", pipelineID),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "value", valueUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "default", defaultVarUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "type", typeVarUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "path", pathUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_tekton_pipeline_property.tekton_pipeline_property",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMTektonPipelinePropertyConfigBasic(pipelineID string) string {
	return fmt.Sprintf(`

		resource "ibm_cd_tekton_pipeline_property" "tekton_pipeline_property" {
			pipeline_id = "%s"
		}
	`, pipelineID)
}

func testAccCheckIBMTektonPipelinePropertyConfig(pipelineID string, name string, value string, defaultVar string, typeVar string, path string) string {
	return fmt.Sprintf(`

		resource "ibm_cd_tekton_pipeline_property" "tekton_pipeline_property" {
			pipeline_id = "%s"
			name = "%s"
			value = "%s"
			enum = "FIXME"
			default = "%s"
			type = "%s"
			path = "%s"
		}
	`, pipelineID, name, value, defaultVar, typeVar, path)
}

func testAccCheckIBMTektonPipelinePropertyExists(n string, obj cdtektonpipelinev2.Property) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cdTektonPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdTektonPipelineV2()
		if err != nil {
			return err
		}

		getTektonPipelinePropertyOptions := &cdtektonpipelinev2.GetTektonPipelinePropertyOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTektonPipelinePropertyOptions.SetPipelineID(parts[0])
		getTektonPipelinePropertyOptions.SetPropertyName(parts[1])

		property, _, err := cdTektonPipelineClient.GetTektonPipelineProperty(getTektonPipelinePropertyOptions)
		if err != nil {
			return err
		}

		obj = *property
		return nil
	}
}

func testAccCheckIBMTektonPipelinePropertyDestroy(s *terraform.State) error {
	cdTektonPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdTektonPipelineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_tekton_pipeline_property" {
			continue
		}

		getTektonPipelinePropertyOptions := &cdtektonpipelinev2.GetTektonPipelinePropertyOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTektonPipelinePropertyOptions.SetPipelineID(parts[0])
		getTektonPipelinePropertyOptions.SetPropertyName(parts[1])

		// Try to find the key
		_, response, err := cdTektonPipelineClient.GetTektonPipelineProperty(getTektonPipelinePropertyOptions)

		if err == nil {
			return fmt.Errorf("tekton_pipeline_property still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for tekton_pipeline_property (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

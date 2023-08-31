// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
)

func TestAccIbmSccProviderTypeInstanceBasic(t *testing.T) {
	var conf securityandcompliancecenterapiv3.ProviderTypeInstanceItem
	providerTypeID := fmt.Sprintf("tf_provider_type_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccProviderTypeInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProviderTypeInstanceConfigBasic(providerTypeID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccProviderTypeInstanceExists("ibm_scc_provider_type_instance.scc_provider_type_instance", conf),
					resource.TestCheckResourceAttr("ibm_scc_provider_type_instance.scc_provider_type_instance", "provider_type_id", providerTypeID),
					resource.TestCheckResourceAttr("ibm_scc_provider_type_instance.scc_provider_type_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSccProviderTypeInstanceConfigBasic(providerTypeID, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_provider_type_instance.scc_provider_type_instance", "provider_type_id", providerTypeID),
					resource.TestCheckResourceAttr("ibm_scc_provider_type_instance.scc_provider_type_instance", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIbmSccProviderTypeInstanceAllArgs(t *testing.T) {
	var conf securityandcompliancecenterapiv3.ProviderTypeInstanceItem
	providerTypeID := fmt.Sprintf("tf_provider_type_id_%d", acctest.RandIntRange(10, 100))
	xCorrelationID := fmt.Sprintf("tf_x_correlation_id_%d", acctest.RandIntRange(10, 100))
	xRequestID := fmt.Sprintf("tf_x_request_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	xCorrelationIDUpdate := fmt.Sprintf("tf_x_correlation_id_%d", acctest.RandIntRange(10, 100))
	xRequestIDUpdate := fmt.Sprintf("tf_x_request_id_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccProviderTypeInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProviderTypeInstanceConfig(providerTypeID, xCorrelationID, xRequestID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccProviderTypeInstanceExists("ibm_scc_provider_type_instance.scc_provider_type_instance", conf),
					resource.TestCheckResourceAttr("ibm_scc_provider_type_instance.scc_provider_type_instance", "provider_type_id", providerTypeID),
					resource.TestCheckResourceAttr("ibm_scc_provider_type_instance.scc_provider_type_instance", "x_correlation_id", xCorrelationID),
					resource.TestCheckResourceAttr("ibm_scc_provider_type_instance.scc_provider_type_instance", "x_request_id", xRequestID),
					resource.TestCheckResourceAttr("ibm_scc_provider_type_instance.scc_provider_type_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSccProviderTypeInstanceConfig(providerTypeID, xCorrelationIDUpdate, xRequestIDUpdate, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_provider_type_instance.scc_provider_type_instance", "provider_type_id", providerTypeID),
					resource.TestCheckResourceAttr("ibm_scc_provider_type_instance.scc_provider_type_instance", "x_correlation_id", xCorrelationIDUpdate),
					resource.TestCheckResourceAttr("ibm_scc_provider_type_instance.scc_provider_type_instance", "x_request_id", xRequestIDUpdate),
					resource.TestCheckResourceAttr("ibm_scc_provider_type_instance.scc_provider_type_instance", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_provider_type_instance.scc_provider_type_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSccProviderTypeInstanceConfigBasic(providerTypeID string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_provider_type_instance" "scc_provider_type_instance_instance" {
			provider_type_id = "provider_type_id"
			name = "workload-protection-instance-1"
			attributes = {"wp_crn":"crn:v1:staging:public:sysdig-secure:eu-gb:a/14q5SEnVIbwxzvP4AWPCjr2dJg5BAvPb:d1461d1ae-df1eee12fa81812e0-12-aa259::"}
		}
		resource "ibm_scc_provider_type_instance" "scc_provider_type_instance_instance" {
			provider_type_id = "%s"
			name = "%s"
			attributes = {"wp_crn":"crn:v1:staging:public:sysdig-secure:eu-gb:a/14q5SEnVIbwxzvP4AWPCjr2dJg5BAvPb:d1461d1ae-df1eee12fa81812e0-12-aa259::"}
		}
	`, providerTypeID, name)
}

func testAccCheckIbmSccProviderTypeInstanceConfig(providerTypeID string, xCorrelationID string, xRequestID string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_provider_type_instance" "scc_provider_type_instance_instance" {
			provider_type_id = "provider_type_id"
			name = "workload-protection-instance-1"
			attributes = {"wp_crn":"crn:v1:staging:public:sysdig-secure:eu-gb:a/14q5SEnVIbwxzvP4AWPCjr2dJg5BAvPb:d1461d1ae-df1eee12fa81812e0-12-aa259::"}
		}

		resource "ibm_scc_provider_type_instance" "scc_provider_type_instance_instance" {
			provider_type_id = "%s"
			x_correlation_id = "%s"
			x_request_id = "%s"
			name = "%s"
			attributes = {"wp_crn":"crn:v1:staging:public:sysdig-secure:eu-gb:a/14q5SEnVIbwxzvP4AWPCjr2dJg5BAvPb:d1461d1ae-df1eee12fa81812e0-12-aa259::"}
		}
	`, providerTypeID, xCorrelationID, xRequestID, name)
}

func testAccCheckIbmSccProviderTypeInstanceExists(n string, obj securityandcompliancecenterapiv3.ProviderTypeInstanceItem) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		securityAndComplianceCenterApIsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
		if err != nil {
			return err
		}

		getProviderTypeInstanceOptions := &securityandcompliancecenterapiv3.GetProviderTypeInstanceOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getProviderTypeInstanceOptions.SetProviderTypeID(parts[0])
		getProviderTypeInstanceOptions.SetProviderTypeInstanceID(parts[1])

		providerTypeInstanceItem, _, err := securityAndComplianceCenterApIsClient.GetProviderTypeInstance(getProviderTypeInstanceOptions)
		if err != nil {
			return err
		}

		obj = *providerTypeInstanceItem
		return nil
	}
}

func testAccCheckIbmSccProviderTypeInstanceDestroy(s *terraform.State) error {
	securityAndComplianceCenterApIsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_provider_type_instance" {
			continue
		}

		getProviderTypeInstanceOptions := &securityandcompliancecenterapiv3.GetProviderTypeInstanceOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getProviderTypeInstanceOptions.SetProviderTypeID(parts[0])
		getProviderTypeInstanceOptions.SetProviderTypeInstanceID(parts[1])

		// Try to find the key
		_, response, err := securityAndComplianceCenterApIsClient.GetProviderTypeInstance(getProviderTypeInstanceOptions)

		if err == nil {
			return fmt.Errorf("scc_provider_type_instance still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for scc_provider_type_instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func TestAccIBMTrustedProfileTemplateBasic(t *testing.T) {
	enterpriseAccountId := acc.IamIdentityEnterpriseAccountId
	var conf iamidentityv1.TrustedProfileTemplateResponse

	name := fmt.Sprintf("tf_tp_name_%d", acctest.RandIntRange(10, 100))
	profileName := fmt.Sprintf("tf_tp_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_tp_desc_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_desc_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTrustedProfileTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMTrustedProfileTemplateConfigBasic(enterpriseAccountId, name, description, profileName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMTrustedProfileTemplateExists("ibm_iam_trusted_profile_template.trusted_profile_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_template.trusted_profile_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_template.trusted_profile_template", "description", description),
				),
			},
			{
				Config: testAccCheckIBMTrustedProfileTemplateConfigBasic(enterpriseAccountId, name, descriptionUpdate, profileName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMTrustedProfileTemplateExists("ibm_iam_trusted_profile_template.trusted_profile_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_template.trusted_profile_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_template.trusted_profile_template", "description", descriptionUpdate),
				),
			},
		},
	})
}

func TestAccIBMTrustedProfileTemplateVersionBasic(t *testing.T) {
	enterpriseAccountId := acc.IamIdentityEnterpriseAccountId
	var conf iamidentityv1.TrustedProfileTemplateResponse

	name := fmt.Sprintf("tf_tp_name_%d", acctest.RandIntRange(10, 100))
	profileName := fmt.Sprintf("tf_tp_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTrustedProfileTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMTrustedProfileTemplateVersionConfigBasic(enterpriseAccountId, name, profileName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMTrustedProfileTemplateExists("ibm_iam_trusted_profile_template.trusted_profile_template", conf),
					testAccCheckIBMTrustedProfileTemplateExists("ibm_iam_trusted_profile_template.trusted_profile_template_version", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_template.trusted_profile_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_template.trusted_profile_template_version", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_template.trusted_profile_template_version", "description", "New version description"),
				),
			},
		},
	})
}

func TestAccIBMTrustedProfileTemplateAllArgs(t *testing.T) {
	enterpriseAccountId := acc.IamIdentityEnterpriseAccountId
	var conf iamidentityv1.TrustedProfileTemplateResponse
	name := fmt.Sprintf("tf_tp_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_tp_desc_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_desc_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTrustedProfileTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMTrustedProfileTemplateConfig(enterpriseAccountId, name, description, "false"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMTrustedProfileTemplateExists("ibm_iam_trusted_profile_template.trusted_profile_template", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_template.trusted_profile_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_template.trusted_profile_template", "description", description),
				),
			},
			{
				Config: testAccCheckIBMTrustedProfileTemplateConfig(enterpriseAccountId, name, descriptionUpdate, "false"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_template.trusted_profile_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_template.trusted_profile_template", "description", descriptionUpdate),
				),
			},
			{
				Config: testAccCheckIBMTrustedProfileTemplateConfig(enterpriseAccountId, name, descriptionUpdate, "true"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_template.trusted_profile_template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_template.trusted_profile_template", "description", descriptionUpdate),
				),
			},
			{
				ResourceName:      "ibm_iam_trusted_profile_template.trusted_profile_template",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMTrustedProfileTemplateConfigBasic(enterpriseAccountId string, name string, description string, profileName string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile_template" "trusted_profile_template" {
			account_id = "%s"
			name = "%s"
			description = "%s"
			profile {
				name = "%s"
			}
		}
	`, enterpriseAccountId, name, description, profileName)
}

func testAccCheckIBMTrustedProfileTemplateVersionConfigBasic(enterpriseAccountId string, name string, profileName string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile_template" "trusted_profile_template" {
			account_id = "%s"
			name = "%s"
			profile {
				name = "%s"
			}
		}

		resource "ibm_iam_trusted_profile_template" "trusted_profile_template_version" {
			template_id = ibm_iam_trusted_profile_template.trusted_profile_template.id
			account_id = ibm_iam_trusted_profile_template.trusted_profile_template.account_id
			name = "%s"
			description = "New version description"
			profile {
				name = "%s"
			}
		}
	`, enterpriseAccountId, name, profileName, name, profileName)
}

func testAccCheckIBMTrustedProfileTemplateConfig(enterpriseAccountId string, name string, description string, committed string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile_template" "trusted_profile_template" {
			account_id = "%s"
			name = "%s"
			description = "%s"
			committed = "%s"
			profile {
				name = "name"
				description = "description"
				rules {
					name = "name"
					type = "Profile-SAML"
					realm_name = "test-realm-101"
					expiration = 1
					conditions {
						claim = "claim"
						operator = "EQUALS"
						value = "\"value\""
					}
				}
				identities {
					iam_id = "crn-crn:v1:staging:public:iam-identity::a/684e0f537b4548eb8d2c9593881a6b03:::"
					identifier = "crn:v1:staging:public:iam-identity::a/684e0f537b4548eb8d2c9593881a6b03:::"
					type = "crn"
				}
			}
		}
	`, enterpriseAccountId, name, description, committed)
}

func testAccCheckIBMTrustedProfileTemplateExists(n string, obj iamidentityv1.TrustedProfileTemplateResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		id, version := parseProfileResourceId(rs.Primary.ID)
		getProfileTemplateVersionOptions := iamIdentityClient.NewGetProfileTemplateVersionOptions(id, version)

		trustedProfileTemplateResponse, _, err := iamIdentityClient.GetProfileTemplateVersion(getProfileTemplateVersionOptions)
		if err != nil {
			return err
		}

		obj = *trustedProfileTemplateResponse
		return nil
	}
}

func testAccCheckIBMTrustedProfileTemplateDestroy(s *terraform.State) error {
	iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_trusted_profile_template" {
			continue
		}

		id, version := parseProfileResourceId(rs.Primary.ID)
		getProfileTemplateVersionOptions := iamIdentityClient.NewGetProfileTemplateVersionOptions(id, version)

		// Try to find the key
		_, response, err := iamIdentityClient.GetProfileTemplateVersion(getProfileTemplateVersionOptions)

		if err == nil {
			return fmt.Errorf("trusted_profile_template still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for trusted_profile_template (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func parseProfileResourceId(ID string) (templateId, templateVersion string) {
	resourceIdParts := strings.Split(ID, "/")

	if len(resourceIdParts) == 1 {
		return resourceIdParts[0], ""
	}

	return resourceIdParts[0], resourceIdParts[1]
}

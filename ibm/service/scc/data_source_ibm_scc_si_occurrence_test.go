// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccSiOccurrenceDataSourceBasic(t *testing.T) {
	apiOccurrenceProviderID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	apiOccurrencenoteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))
	apiOccurrenceKind := "FINDING"
	apiOccurrenceID := fmt.Sprintf("tf_occurrence_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSccSiOccurrenceDataSourceConfigBasic(acc.Scc_si_account, apiOccurrenceProviderID, apiOccurrencenoteID, apiOccurrenceKind, apiOccurrenceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_occurrence.scc_si_occurrence", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_occurrence.scc_si_occurrence", "provider_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_occurrence.scc_si_occurrence", "occurrence_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_occurrence.scc_si_occurrence", "note_name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_occurrence.scc_si_occurrence", "kind"),
				),
			},
		},
	})
}

func TestAccIBMSccSiOccurrenceDataSourceAllArgs(t *testing.T) {
	apiOccurrenceProviderID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	apiOccurrencenoteID := fmt.Sprintf("tf_note_name_%d", acctest.RandIntRange(10, 100))
	apiOccurrenceKind := "FINDING"
	apiOccurrenceID := fmt.Sprintf("tf_occurrence_id_%d", acctest.RandIntRange(10, 100))
	apiOccurrenceResourceURL := fmt.Sprintf("tf_resource_url_%d", acctest.RandIntRange(10, 100))
	apiOccurrenceRemediation := fmt.Sprintf("tf_remediation_%d", acctest.RandIntRange(10, 100))
	apiOccurrenceReplaceIfExists := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSccSiOccurrenceDataSourceConfig(acc.Scc_si_account, apiOccurrenceProviderID, apiOccurrencenoteID, apiOccurrenceKind, apiOccurrenceID, apiOccurrenceResourceURL, apiOccurrenceRemediation, apiOccurrenceReplaceIfExists),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_occurrence.scc_si_occurrence", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_occurrence.scc_si_occurrence", "provider_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_occurrence.scc_si_occurrence", "occurrence_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_occurrence.scc_si_occurrence", "resource_url"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_occurrence.scc_si_occurrence", "note_name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_occurrence.scc_si_occurrence", "kind"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_occurrence.scc_si_occurrence", "remediation"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_occurrence.scc_si_occurrence", "create_time"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_occurrence.scc_si_occurrence", "update_time"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_occurrence.scc_si_occurrence", "context.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_occurrence.scc_si_occurrence", "finding.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSccSiOccurrenceDataSourceConfigBasic(accountID string, apiOccurrenceProviderID string, apiOccurrencenoteID string, apiOccurrenceKind string, apiOccurrenceID string) string {
	return fmt.Sprintf(`

	resource "ibm_scc_si_note" "finding" {
		account_id        = "%s"
		provider_id       = "%s"
		short_description = "Security Threat"
		long_description  = "Security Threat found in your account"
		kind              = "FINDING"
		note_id           = "%s"
		reported_by {
		  id    = "scc-si-terraform"
		  title = "SCC SI Terraform"
		  url   = "https://cloud.ibm.com"
		}
		finding {
			severity = "LOW"
			next_steps {
			  title = "Security Threat"
			  url   = "https://cloud.ibm.com/security-compliance/findings"
			}
		}
	  }

		resource "ibm_scc_si_occurrence" "scc_si_occurrence" {
			provider_id = "${ibm_scc_si_note.finding.provider_id}"
			note_name = "%s/providers/${ibm_scc_si_note.finding.provider_id}/notes/${ibm_scc_si_note.finding.note_id}"
			kind = "%s"
			occurrence_id = "%s"
			finding {
				severity = "LOW"
				certainty = "LOW"
				next_steps {
					title = "title"
					url = "url"
				}
				network_connection {
					direction = "direction"
					protocol = "protocol"
					client {
						address = "address"
						port = 1
					}
					server {
						address = "address"
						port = 1
					}
				}
				data_transferred {
					client_bytes = 1
					server_bytes = 1
					client_packets = 1
					server_packets = 1
				}
			}
		}

		data "ibm_scc_si_occurrence" "scc_si_occurrence" {
			provider_id = ibm_scc_si_occurrence.scc_si_occurrence.provider_id
			occurrence_id = ibm_scc_si_occurrence.scc_si_occurrence.occurrence_id
		}
	`, accountID, apiOccurrenceProviderID, apiOccurrencenoteID, accountID, apiOccurrenceKind, apiOccurrenceID)
}

func testAccCheckIBMSccSiOccurrenceDataSourceConfig(accountID string, apiOccurrenceProviderID string, apiOccurrencenoteID string, apiOccurrenceKind string, apiOccurrenceID string, apiOccurrenceResourceURL string, apiOccurrenceRemediation string, apiOccurrenceReplaceIfExists string) string {
	return fmt.Sprintf(`

	resource "ibm_scc_si_note" "finding" {
		account_id        = "%s"
		provider_id       = "%s"
		short_description = "Security Threat"
		long_description  = "Security Threat found in your account"
		kind              = "FINDING"
		note_id           = "%s"
		reported_by {
		  id    = "scc-si-terraform"
		  title = "SCC SI Terraform"
		  url   = "https://cloud.ibm.com"
		}
		finding {
			severity = "LOW"
			next_steps {
			  title = "Security Threat"
			  url   = "https://cloud.ibm.com/security-compliance/findings"
			}
		}
	  }

		resource "ibm_scc_si_occurrence" "scc_si_occurrence" {
			provider_id = "${ibm_scc_si_note.finding.provider_id}"
			note_name = "%s/providers/${ibm_scc_si_note.finding.provider_id}/notes/${ibm_scc_si_note.finding.note_id}"
			kind = "%s"
			occurrence_id = "%s"
			resource_url = "%s"
			remediation = "%s"
			context {
				region = "region"
				resource_crn = "resource_crn"
				resource_id = "resource_id"
				resource_name = "resource_name"
				resource_type = "resource_type"
				service_crn = "service_crn"
				service_name = "service_name"
				environment_name = "environment_name"
				component_name = "component_name"
				toolchain_id = "toolchain_id"
			}
			finding {
				severity = "LOW"
				certainty = "LOW"
				next_steps {
					title = "title"
					url = "url"
				}
				network_connection {
					direction = "direction"
					protocol = "protocol"
					client {
						address = "address"
						port = 1
					}
					server {
						address = "address"
						port = 1
					}
				}
				data_transferred {
					client_bytes = 1
					server_bytes = 1
					client_packets = 1
					server_packets = 1
				}
			}
			replace_if_exists = %s
		}

		data "ibm_scc_si_occurrence" "scc_si_occurrence" {
			provider_id = ibm_scc_si_occurrence.scc_si_occurrence.provider_id
			occurrence_id = ibm_scc_si_occurrence.scc_si_occurrence.occurrence_id
		}
	`, accountID, apiOccurrenceProviderID, apiOccurrencenoteID, accountID, apiOccurrenceKind, apiOccurrenceID, apiOccurrenceResourceURL, apiOccurrenceRemediation, apiOccurrenceReplaceIfExists)
}

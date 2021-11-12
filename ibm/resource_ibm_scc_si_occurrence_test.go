// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/scc-go-sdk/findingsv1"
)

func TestAccIBMSccSiOccurrenceBasic(t *testing.T) {
	var conf findingsv1.APIOccurrence
	kind := "FINDING"
	occurrenceID := fmt.Sprintf("tf_occurrence_id_%d", acctest.RandIntRange(10, 100))

	kindUpdate := kind
	providerID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	noteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))
	noteIdUpdate := noteID
	occurrenceIDUpdate := occurrenceID
	accountID := scc_si_account

	noteName := fmt.Sprintf("%s/providers/%s/notes/%s", accountID, providerID, noteID)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSccSiOccurrenceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccSiOccurrenceConfigBasic(providerID, noteID, accountID, kind, occurrenceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccSiOccurrenceExists("ibm_scc_si_occurrence.scc_si_occurrence", conf),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "provider_id", providerID),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "note_name", noteName),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "kind", kind),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "occurrence_id", occurrenceID),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSccSiOccurrenceConfigBasic(providerID, noteIdUpdate, accountID, kindUpdate, occurrenceIDUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "provider_id", providerID),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "note_name", noteName),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "kind", kindUpdate),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "occurrence_id", occurrenceIDUpdate),
				),
			},
		},
	})
}

func TestAccIBMSccSiOccurrenceKpiBasic(t *testing.T) {
	var conf findingsv1.APIOccurrence
	kind := "KPI"
	occurrenceID := fmt.Sprintf("tf_occurrence_id_%d", acctest.RandIntRange(10, 100))

	kindUpdate := kind
	providerID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	noteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))
	noteIdUpdate := noteID
	occurrenceIDUpdate := occurrenceID
	accountID := scc_si_account

	noteName := fmt.Sprintf("%s/providers/%s/notes/%s", accountID, providerID, noteID)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSccSiOccurrenceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccSiOccurrenceConfigKpiBasic(providerID, noteID, accountID, kind, occurrenceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccSiOccurrenceExists("ibm_scc_si_occurrence.scc_si_occurrence", conf),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "provider_id", providerID),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "note_name", noteName),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "kind", kind),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "occurrence_id", occurrenceID),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSccSiOccurrenceConfigKpiBasic(providerID, noteIdUpdate, accountID, kindUpdate, occurrenceIDUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "provider_id", providerID),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "note_name", noteName),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "kind", kindUpdate),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "occurrence_id", occurrenceIDUpdate),
				),
			},
		},
	})
}

func TestAccIBMSccSiOccurrenceFindingNoteOccurrenceKPI(t *testing.T) {
	kind := "FINDING"
	occurrenceID := fmt.Sprintf("tf_occurrence_id_%d", acctest.RandIntRange(10, 100))

	providerID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	noteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))

	accountID := scc_si_account

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSccSiOccurrenceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMSccSiOccurrenceConfigFindingNoteOccurrenceKPI(providerID, noteID, accountID, kind, occurrenceID),
				ExpectError: regexp.MustCompile("Missing field 'finding' for kind 'FINDING"),
			},
		},
	})
}

func TestAccIBMSccSiOccurrenceKPINoteOccurrenceFinding(t *testing.T) {
	kind := "KPI"
	occurrenceID := fmt.Sprintf("tf_occurrence_id_%d", acctest.RandIntRange(10, 100))

	providerID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	noteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))

	accountID := scc_si_account

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSccSiOccurrenceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMSccSiOccurrenceConfigKPINoteOccurrenceFinding(providerID, noteID, accountID, kind, occurrenceID),
				ExpectError: regexp.MustCompile("Missing field 'kpi' for kind 'KPI'"),
			},
		},
	})
}

func TestAccIBMSccSiOccurrenceValidNoteNoOccurrence(t *testing.T) {
	kind := "KPI"
	occurrenceID := fmt.Sprintf("tf_occurrence_id_%d", acctest.RandIntRange(10, 100))

	providerID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	noteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))

	accountID := scc_si_account

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSccSiOccurrenceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMSccSiOccurrenceConfigValidNoteNoOccurrence(providerID, noteID, accountID, kind, occurrenceID),
				ExpectError: regexp.MustCompile("one of `finding,kpi` must be specified"),
			},
		},
	})
}

func TestAccIBMSccSiOccurrenceValidNoteKpiFindingOccurrence(t *testing.T) {
	kind := "KPI"
	occurrenceID := fmt.Sprintf("tf_occurrence_id_%d", acctest.RandIntRange(10, 100))

	providerID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	noteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))

	accountID := scc_si_account

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSccSiOccurrenceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMSccSiOccurrenceConfigValidNoteKpiFindingOccurrence(providerID, noteID, accountID, kind, occurrenceID),
				ExpectError: regexp.MustCompile("only one of `finding,kpi` can be specified"),
			},
		},
	})
}
func TestAccIBMSccSiOccurrenceAllArgs(t *testing.T) {
	var conf findingsv1.APIOccurrence
	providerID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	noteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))
	kind := "FINDING"
	occurrenceID := fmt.Sprintf("tf_occurrence_id_%d", acctest.RandIntRange(10, 100))
	resourceURL := fmt.Sprintf("tf_resource_url_%d", acctest.RandIntRange(10, 100))
	remediation := fmt.Sprintf("tf_remediation_%d", acctest.RandIntRange(10, 100))
	replaceIfExists := "false"
	noteIdUpdate := noteID
	kindUpdate := kind
	occurrenceIDUpdate := occurrenceID
	resourceURLUpdate := fmt.Sprintf("tf_resource_url_%d", acctest.RandIntRange(10, 100))
	remediationUpdate := fmt.Sprintf("tf_remediation_%d", acctest.RandIntRange(10, 100))
	replaceIfExistsUpdate := "true"
	accountID := scc_si_account

	noteName := fmt.Sprintf("%s/providers/%s/notes/%s", accountID, providerID, noteID)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSccSiOccurrenceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccSiOccurrenceConfig(accountID, providerID, noteID, kind, occurrenceID, resourceURL, remediation, replaceIfExists),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccSiOccurrenceExists("ibm_scc_si_occurrence.scc_si_occurrence", conf),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "provider_id", providerID),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "note_name", noteName),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "kind", kind),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "occurrence_id", occurrenceID),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "resource_url", resourceURL),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "remediation", remediation),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSccSiOccurrenceConfig(accountID, providerID, noteIdUpdate, kindUpdate, occurrenceIDUpdate, resourceURLUpdate, remediationUpdate, replaceIfExistsUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "provider_id", providerID),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "note_name", noteName),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "kind", kindUpdate),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "occurrence_id", occurrenceIDUpdate),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "resource_url", resourceURLUpdate),
					resource.TestCheckResourceAttr("ibm_scc_si_occurrence.scc_si_occurrence", "remediation", remediationUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_si_occurrence.scc_si_occurrence",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"replace_if_exists",
				},
			},
		},
	})
}

func testAccCheckIBMSccSiOccurrenceConfigFindingNoteOccurrenceKPI(providerID string, noteID string, accountID string, kind string, occurrenceID string) string {
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
			kpi {
				value = 1.0
				total = 1.0
			}
		}
	`, accountID, providerID, noteID, accountID, kind, occurrenceID)
}

func testAccCheckIBMSccSiOccurrenceConfigKPINoteOccurrenceFinding(providerID string, noteID string, accountID string, kind string, occurrenceID string) string {
	return fmt.Sprintf(`

	resource "ibm_scc_si_note" "finding" {
		account_id        = "%s"
		provider_id       = "%s"
		short_description = "Security Threat"
		long_description  = "Security Threat found in your account"
		kind              = "KPI"
		note_id           = "%s"
		reported_by {
		  id    = "scc-si-terraform"
		  title = "SCC SI Terraform"
		  url   = "https://cloud.ibm.com"
		}
		kpi {
			aggregation_type = "SUM"
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
	`, accountID, providerID, noteID, accountID, kind, occurrenceID)
}

func testAccCheckIBMSccSiOccurrenceConfigValidNoteNoOccurrence(providerID string, noteID string, accountID string, kind string, occurrenceID string) string {
	return fmt.Sprintf(`

	resource "ibm_scc_si_note" "finding" {
		account_id        = "%s"
		provider_id       = "%s"
		short_description = "Security Threat"
		long_description  = "Security Threat found in your account"
		kind              = "KPI"
		note_id           = "%s"
		reported_by {
		  id    = "scc-si-terraform"
		  title = "SCC SI Terraform"
		  url   = "https://cloud.ibm.com"
		}
		kpi {
			aggregation_type = "SUM"
		}
	  }

		resource "ibm_scc_si_occurrence" "scc_si_occurrence" {
			provider_id = "${ibm_scc_si_note.finding.provider_id}"
			note_name = "%s/providers/${ibm_scc_si_note.finding.provider_id}/notes/${ibm_scc_si_note.finding.note_id}"
			kind = "%s"
			occurrence_id = "%s"
		}
	`, accountID, providerID, noteID, accountID, kind, occurrenceID)
}

func testAccCheckIBMSccSiOccurrenceConfigValidNoteKpiFindingOccurrence(providerID string, noteID string, accountID string, kind string, occurrenceID string) string {
	return fmt.Sprintf(`

	resource "ibm_scc_si_note" "finding" {
		account_id        = "%s"
		provider_id       = "%s"
		short_description = "Security Threat"
		long_description  = "Security Threat found in your account"
		kind              = "KPI"
		note_id           = "%s"
		reported_by {
		  id    = "scc-si-terraform"
		  title = "SCC SI Terraform"
		  url   = "https://cloud.ibm.com"
		}
		kpi {
			aggregation_type = "SUM"
		}
	  }

		resource "ibm_scc_si_occurrence" "scc_si_occurrence" {
			provider_id = "${ibm_scc_si_note.finding.provider_id}"
			note_name = "%s/providers/${ibm_scc_si_note.finding.provider_id}/notes/${ibm_scc_si_note.finding.note_id}"
			kind = "%s"
			occurrence_id = "%s"
			kpi {
				value = 1.0
				total = 1.0
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
		}
	`, accountID, providerID, noteID, accountID, kind, occurrenceID)
}

func testAccCheckIBMSccSiOccurrenceConfigKpiBasic(providerID string, noteID string, accountID string, kind string, occurrenceID string) string {
	return fmt.Sprintf(`

	resource "ibm_scc_si_note" "finding" {
		account_id        = "%s"
		provider_id       = "%s"
		short_description = "Security Threat"
		long_description  = "Security Threat found in your account"
		kind              = "KPI"
		note_id           = "%s"
		reported_by {
		  id    = "scc-si-terraform"
		  title = "SCC SI Terraform"
		  url   = "https://cloud.ibm.com"
		}
		kpi {
			aggregation_type = "SUM"
		}
	  }

		resource "ibm_scc_si_occurrence" "scc_si_occurrence" {
			provider_id = "${ibm_scc_si_note.finding.provider_id}"
			note_name = "%s/providers/${ibm_scc_si_note.finding.provider_id}/notes/${ibm_scc_si_note.finding.note_id}"
			kind = "%s"
			occurrence_id = "%s"
			kpi {
				value = 1.0
				total = 1.0
			}
		}
	`, accountID, providerID, noteID, accountID, kind, occurrenceID)
}

func testAccCheckIBMSccSiOccurrenceConfigBasic(providerID string, noteID string, accountID string, kind string, occurrenceID string) string {
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
	`, accountID, providerID, noteID, accountID, kind, occurrenceID)
}

func testAccCheckIBMSccSiOccurrenceConfig(accountID string, providerID string, noteID string, kind string, occurrenceID string, resourceURL string, remediation string, replaceIfExists string) string {
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
	`, accountID, providerID, noteID, accountID, kind, occurrenceID, resourceURL, remediation, replaceIfExists)
}

func testAccCheckIBMSccSiOccurrenceExists(n string, obj findingsv1.APIOccurrence) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		findingsClient, err := testAccProvider.Meta().(ClientSession).FindingsV1()
		if err != nil {
			return err
		}

		getOccurrenceOptions := &findingsv1.GetOccurrenceOptions{}

		parts, err := sepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}
		findingsClient.AccountID = &parts[0]
		getOccurrenceOptions.SetProviderID(parts[1])
		getOccurrenceOptions.SetOccurrenceID(parts[2])

		apiOccurrence, _, err := findingsClient.GetOccurrence(getOccurrenceOptions)
		if err != nil {
			return err
		}

		obj = *apiOccurrence
		return nil
	}
}

func testAccCheckIBMSccSiOccurrenceDestroy(s *terraform.State) error {
	findingsClient, err := testAccProvider.Meta().(ClientSession).FindingsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_si_occurrence" {
			continue
		}

		getOccurrenceOptions := &findingsv1.GetOccurrenceOptions{}

		parts, err := sepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}
		findingsClient.AccountID = &parts[0]
		getOccurrenceOptions.SetProviderID(parts[1])
		getOccurrenceOptions.SetOccurrenceID(parts[2])

		// Try to find the key
		_, response, err := findingsClient.GetOccurrence(getOccurrenceOptions)

		if err == nil {
			return fmt.Errorf("scc_si_occurrence still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for scc_si_occurrence (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
